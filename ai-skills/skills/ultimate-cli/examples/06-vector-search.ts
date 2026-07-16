#!/usr/bin/env node
/**
 * ============================================================
 *  VECTOR SEARCH CLI — Chroma + Qdrant + LanceDB
 * ============================================================
 * Libraries: chromadb, @qdrant/js-client-rest, vectordb,
 *            @inquirer/prompts, chalk, ora
 *
 * A unified CLI for vector database operations:
 *   - Create/delete collections
 *   - Ingest documents (PDF, text, code)
 *   - Semantic search with filters
 *   - Multi-provider support
 *   - Export/import data
 *
 * Run: npx ts-node examples/06-vector-search.ts ingest ./docs
 *      npx ts-node examples/06-vector-search.ts search "machine learning"
 *      npx ts-node examples/06-vector-search.ts list
 * ============================================================
 */

import { ChromaClient, OpenAIEmbeddingFunction } from 'chromadb';
import { QdrantClient } from '@qdrant/js-client-rest';
import { Command } from 'commander';
import { input, select, confirm, password } from '@inquirer/prompts';
import chalk from 'chalk';
import fg from 'fast-glob';
import fs from 'fs/promises';
import path from 'path';

// ─── Types ───────────────────────────────────────────────
type VectorDB = 'chroma' | 'qdrant' | 'lancedb';

interface SearchResult {
  id: string | number;
  score: number;
  text?: string;
  metadata?: Record<string, any>;
}

interface Document {
  id: string;
  text: string;
  metadata: Record<string, any>;
}

// ─── Config ──────────────────────────────────────────────
const CHROMA_URL = 'http://localhost:8000';
const QDRANT_URL = 'http://localhost:6333';

// ─── Embedding Function ─────────────────────────────────
// In production, use OpenAI / local embeddings
function createEmbedder(apiKey?: string) {
  return {
    generate: async (texts: string[]): Promise<number[][]> => {
      // Simple hash-based embedding for demo
      // In production: use OpenAIEmbeddingFunction
      return texts.map((text) => {
        const vec = new Array(384).fill(0);
        for (let i = 0; i < text.length; i++) {
          vec[i % 384] += text.charCodeAt(i) / 255;
        }
        // Normalize
        const mag = Math.sqrt(vec.reduce((s, v) => s + v * v, 0)) || 1;
        return vec.map((v) => v / mag);
      });
    },
  };
}

// ─── Document Processing ────────────────────────────────
async function loadDocuments(
  patterns: string[],
  baseDir: string
): Promise<Document[]> {
  const documents: Document[] = [];
  let id = 0;

  for (const pattern of patterns) {
    const files = await fg(pattern, {
      cwd: baseDir,
      absolute: true,
      ignore: ['**/node_modules/**', '**/.git/**', '**/dist/**'],
      onlyFiles: true,
    });

    for (const file of files) {
      try {
        const content = await fs.readFile(file, 'utf-8');
        const ext = path.extname(file);
        const relPath = path.relative(baseDir, file);

        // Split into chunks
        const chunks = chunkText(content, 500, 50);
        for (const chunk of chunks) {
          documents.push({
            id: `${relPath}::${id++}`,
            text: chunk,
            metadata: {
              source: relPath,
              extension: ext,
              directory: path.dirname(relPath),
              file_size: content.length,
              chunk_index: chunks.indexOf(chunk),
            },
          });
        }
      } catch (err) {
        // Binary file, skip
      }
    }
  }

  return documents;
}

function chunkText(text: string, chunkSize: number, overlap: number): string[] {
  const chunks: string[] = [];
  let start = 0;
  while (start < text.length) {
    const end = Math.min(start + chunkSize, text.length);
    chunks.push(text.slice(start, end));
    start += chunkSize - overlap;
  }
  return chunks;
}

// ─── ChromaDB Operations ────────────────────────────────
class ChromaDBOps {
  private client: ChromaClient;
  private embedder: any;

  constructor() {
    this.client = new ChromaClient({ path: CHROMA_URL });
    this.embedder = createEmbedder();
  }

  async listCollections(): Promise<string[]> {
    const collections = await this.client.listCollections();
    return collections.map((c: any) => c.name);
  }

  async createCollection(name: string): Promise<void> {
    await this.client.createCollection({
      name,
      metadata: { 'hnsw:space': 'cosine' },
    });
  }

  async deleteCollection(name: string): Promise<void> {
    await this.client.deleteCollection({ name });
  }

  async ingest(collectionName: string, docs: Document[]): Promise<void> {
    const collection = await this.client.getOrCreateCollection({
      name: collectionName,
    });

    const batchSize = 100;
    for (let i = 0; i < docs.length; i += batchSize) {
      const batch = docs.slice(i, i + batchSize);
      const embeddings = await this.embedder.generate(batch.map((d) => d.text));
      await collection.add({
        ids: batch.map((d) => d.id),
        embeddings,
        metadatas: batch.map((d) => d.metadata),
        documents: batch.map((d) => d.text),
      });
      process.stdout.write(chalk.dim(`\r  Indexed ${Math.min(i + batchSize, docs.length)}/${docs.length} chunks`));
    }
    process.stdout.write('\n');
  }

  async search(collectionName: string, query: string, limit: number, filter?: any): Promise<SearchResult[]> {
    const collection = await this.client.getCollection({ name: collectionName });
    const [queryEmbedding] = await this.embedder.generate([query]);

    const results = await collection.query({
      queryEmbeddings: [queryEmbedding],
      nResults: limit,
      where: filter,
      include: ['documents', 'metadatas', 'distances'],
    });

    return (results.ids[0] || []).map((id: string, idx: number) => ({
      id,
      score: 1 - (results.distances?.[0]?.[idx] || 0),
      text: results.documents?.[0]?.[idx],
      metadata: results.metadatas?.[0]?.[idx],
    }));
  }
}

// ─── Qdrant Operations ──────────────────────────────────
class QdrantOps {
  private client: QdrantClient;

  constructor() {
    this.client = new QdrantClient({ host: 'localhost', port: 6333 });
  }

  async listCollections(): Promise<string[]> {
    const { collections } = await this.client.getCollections();
    return collections.map((c: any) => c.name);
  }

  async createCollection(name: string): Promise<void> {
    await this.client.createCollection(name, {
      vectors: { size: 384, distance: 'Cosine' },
    });
  }

  async deleteCollection(name: string): Promise<void> {
    await this.client.deleteCollection(name);
  }

  async ingest(collectionName: string, docs: Document[]): Promise<void> {
    const batchSize = 100;
    for (let i = 0; i < docs.length; i += batchSize) {
      const batch = docs.slice(i, i + batchSize);
      const embeddings = await createEmbedder().generate(batch.map((d) => d.text));
      await this.client.upsert(collectionName, {
        wait: true,
        points: batch.map((d, j) => ({
          id: i + j,
          vector: embeddings[j],
          payload: {
            text: d.text,
            ...d.metadata,
          },
        })),
      });
      process.stdout.write(chalk.dim(`\r  Indexed ${Math.min(i + batchSize, docs.length)}/${docs.length} chunks`));
    }
    process.stdout.write('\n');
  }

  async search(collectionName: string, query: string, limit: number, filter?: any): Promise<SearchResult[]> {
    const [queryEmbedding] = await createEmbedder().generate([query]);

    const results = await this.client.search(collectionName, {
      vector: queryEmbedding,
      limit,
      with_payload: true,
      filter: filter ? { must: Object.entries(filter).map(([k, v]) => ({ key: k, match: { value: v } })) } : undefined,
    });

    return results.map((r: any) => ({
      id: r.id,
      score: r.score,
      text: r.payload?.text,
      metadata: r.payload,
    }));
  }
}

// ─── CLI Logic ──────────────────────────────────────────
async function getDB(name: VectorDB): Promise<ChromaDBOps | QdrantOps> {
  if (name === 'chroma') return new ChromaDBOps();
  if (name === 'qdrant') return new QdrantOps();
  throw new Error('LanceDB not configured in this demo');
}

// ─── Commands ───────────────────────────────────────────

async function listCollections(dbName: VectorDB): Promise<void> {
  const db = await getDB(dbName);
  const collections = await db.listCollections();
  console.log(chalk.bold.cyan(`\n  📚 Collections (${dbName})\n`));
  if (collections.length === 0) {
    console.log(chalk.dim('  No collections\n'));
  } else {
    collections.forEach((name) => console.log(`  ${chalk.green('▸')} ${name}`));
    console.log();
  }
}

async function createCollectionCmd(dbName: VectorDB, name: string): Promise<void> {
  const db = await getDB(dbName);
  await db.createCollection(name);
  console.log(chalk.green(`\n  ✅ Collection '${name}' created on ${dbName}\n`));
}

async function deleteCollectionCmd(dbName: VectorDB, name: string): Promise<void> {
  const db = await getDB(dbName);
  const confirm = await select({
    message: `Delete collection '${name}'?`,
    choices: [
      { name: 'Yes, delete', value: true },
      { name: 'Cancel', value: false },
    ],
  });
  if (confirm) {
    await db.deleteCollection(name);
    console.log(chalk.green(`\n  ✅ Collection '${name}' deleted\n`));
  }
}

async function ingest(
  dbName: VectorDB,
  collection: string,
  patterns: string[],
  baseDir: string
): Promise<void> {
  console.log(chalk.bold.cyan('\n  📥 Ingesting Documents\n'));
  console.log(chalk.dim(`  DB:    ${dbName}`));
  console.log(chalk.dim(`  Col:   ${collection}`));
  console.log(chalk.dim(`  Path:  ${baseDir}`));
  console.log(chalk.dim(`  Globs: ${patterns.join(', ')}\n`));

  const db = await getDB(dbName);
  const docs = await loadDocuments(patterns, baseDir);

  if (docs.length === 0) {
    console.log(chalk.yellow('  No documents found\n'));
    return;
  }

  console.log(chalk.dim(`  Loaded ${docs.length} chunks\n`));
  await db.ingest(collection, docs);
  console.log(chalk.green(`\n  ✅ Ingested ${docs.length} chunks into '${collection}'\n`));
}

async function search(
  dbName: VectorDB,
  collection: string,
  query: string,
  limit: number,
  filter?: string
): Promise<void> {
  console.log(chalk.bold.cyan('\n  🔍 Vector Search\n'));
  console.log(chalk.dim(`  DB:    ${dbName}`));
  console.log(chalk.dim(`  Col:   ${collection}`));
  console.log(chalk.dim(`  Query: ${query}`));
  console.log();

  const db = await getDB(dbName);
  let parsedFilter: any = undefined;
  if (filter) {
    try {
      parsedFilter = JSON.parse(filter);
    } catch {
      console.log(chalk.red('  ✖ Invalid filter JSON\n'));
      return;
    }
  }

  const results = await db.search(collection, query, limit, parsedFilter);

  if (results.length === 0) {
    console.log(chalk.yellow('  No results found\n'));
    return;
  }

  results.forEach((r, i) => {
    console.log(`  ${chalk.cyan(`#${i + 1}`)} ${chalk.yellow(`(${(r.score * 100).toFixed(1)}%)`)}`);
    if (r.metadata?.source) {
      console.log(chalk.dim(`     ${r.metadata.source}`));
    }
    const text = r.text || '';
    const preview = text.length > 200 ? text.slice(0, 200) + '...' : text;
    console.log(`     ${preview}`);
    console.log();
  });
}

// ─── CLI ────────────────────────────────────────────────
const program = new Command()
  .name('vector-search')
  .description('Unified Vector Database CLI')
  .version('1.0.0');

// Global option for DB provider
program
  .option('--db <provider>', 'Vector DB (chroma|qdrant|lancedb)', 'chroma');

program
  .command('list')
  .description('List collections')
  .action(async (opts, cmd) => {
    const globalOpts = program.opts();
    await listCollections(globalOpts.db as VectorDB);
  });

program
  .command('create <name>')
  .description('Create a collection')
  .action(async (name) => {
    const globalOpts = program.opts();
    await createCollectionCmd(globalOpts.db as VectorDB, name);
  });

program
  .command('delete <name>')
  .description('Delete a collection')
  .action(async (name) => {
    const globalOpts = program.opts();
    await deleteCollectionCmd(globalOpts.db as VectorDB, name);
  });

program
  .command('ingest')
  .description('Ingest documents into vector store')
  .argument('<collection>', 'Collection name')
  .argument('<patterns...>', 'Glob patterns')
  .option('--base-dir <dir>', 'Base directory', '.')
  .action(async (collection, patterns, opts) => {
    const globalOpts = program.opts();
    await ingest(globalOpts.db as VectorDB, collection, patterns, opts.baseDir);
  });

program
  .command('search <collection> <query>')
  .description('Search for similar documents')
  .option('-n, --limit <n>', 'Number of results', '10')
  .option('--filter <json>', 'Metadata filter (JSON)')
  .action(async (collection, query, opts) => {
    const globalOpts = program.opts();
    await search(globalOpts.db as VectorDB, collection, query, parseInt(opts.limit), opts.filter);
  });

program.parse(process.argv);
