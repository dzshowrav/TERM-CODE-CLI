# Vector Database Libraries — Complete A-to-Z Reference

---

## 1. Chroma DB (chroma-core/chroma)
**GitHub**: https://github.com/chroma-core/chroma | **Stars**: 17K+
**npm**: `chromadb` | **Weekly**: ~120K | **License**: Apache 2.0
**Website**: https://www.trychroma.com/

### 1.1 Installation
```bash
npm install chromadb
# Also need the server:
pip install chromadb     # (Python server required)
```

### 1.2 Complete API

#### Client Setup

```typescript
import { ChromaClient, OpenAIEmbeddingFunction } from 'chromadb';

// === HTTP Client (default) ===
const client = new ChromaClient({
  path: 'http://localhost:8000',       // default
  fetchOptions: { headers: {} },       // custom fetch options
  auth: {
    provider: 'token',
    credentials: 'my-token',
  },                                   // basic auth
  // or auth: { provider: 'basic', credentials: 'username:password' }
});

// === Cloud Client ===
const cloudClient = new ChromaClient({
  path: 'https://your-tenant.chroma.cloud',
  auth: { provider: 'token', credentials: process.env.CHROMA_API_KEY },
});

// === Embedding Function ===
const embedder = new OpenAIEmbeddingFunction({
  openai_api_key: process.env.OPENAI_API_KEY,
  openai_model: 'text-embedding-3-small',   // default: text-embedding-ada-002
  openai_organization: 'org-xxx',           // optional
  openai_dimensions: 256,                   // for v3 models
});

// Custom embedding function:
const customEmbedder = {
  generate: async (texts: string[]) => {
    // return number[][] of embeddings
    return texts.map(t => new Array(384).fill(0));
  },
};
```

#### Collection Operations

```typescript
// === Create Collection ===
const collection = await client.createCollection({
  name: 'my_collection',
  metadata: { 'hnsw:space': 'cosine' },    // 'l2' | 'ip' | 'cosine'
  embeddingFunction: embedder,
});

// Auto-create if not exists
const collection2 = await client.getOrCreateCollection({
  name: 'my_collection',
  metadata: { 'hnsw:space': 'cosine' },
});

// === Get Collection ===
const collection3 = await client.getCollection({
  name: 'my_collection',
  embeddingFunction: embedder,
});

// === List Collections ===
const collections = await client.listCollections();
// [{ name: string, metadata: object }]

// === Delete Collection ===
await client.deleteCollection({ name: 'old_collection' });

// === Count ===
const count = await collection.count();
```

#### Upsert / Add / Update

```typescript
// === Add (no overwrite) ===
await collection.add({
  ids: ['id1', 'id2', 'id3'],
  embeddings: [[0.1, 0.2], [0.3, 0.4], [0.5, 0.6]],  // optional if embedding function set
  metadatas: [
    { source: 'pdf', page: 1 },
    { source: 'web', url: 'https://...' },
    { source: 'pdf', page: 2 },
  ],
  documents: ['Doc 1 content', 'Doc 2 content', 'Doc 3 content'],
});

// === Upsert (overwrite if exists) ===
await collection.upsert({
  ids: ['id1', 'id2'],
  embeddings: [[0.1, 0.2], [0.3, 0.4]],
  metadatas: [{ source: 'updated' }, { source: 'web' }],
  documents: ['Updated doc 1', 'Updated doc 2'],
});

// === Update (only existing) ===
await collection.update({
  ids: ['id1'],
  embeddings: [[0.5, 0.6]],
  metadatas: [{ source: 'updated' }],
  documents: ['Updated content'],
});
```

#### Query

```typescript
// === Similarity Search ===
const results = await collection.query({
  queryEmbeddings: [[0.1, 0.2, 0.3]],       // or use queryTexts with embedding fn
  queryTexts: ['search query'],               // alternative to queryEmbeddings
  nResults: 10,                               // number of results
  where: { source: 'pdf' },                   // metadata filter (equality)
  whereDocument: { $contains: 'keyword' },    // document filter
  include: ['documents', 'metadatas', 'distances', 'embeddings'],
  // Default: ['metadatas', 'documents', 'distances']
});

// results = {
//   ids: [['id1', 'id2', ...]],
//   distances: [[0.1, 0.2, ...]],
//   metadatas: [[{...}, {...}, ...]],
//   documents: [['doc1', 'doc2', ...]],
//   embeddings: [[[vec1], [vec2], ...]],
// }

// === Get by IDs ===
const items = await collection.get({
  ids: ['id1', 'id2'],
  where: { source: 'pdf' },
  limit: 10,
  offset: 0,
  include: ['documents', 'metadatas'],
});

// === Peek (first N) ===
const first10 = await collection.peek({ limit: 10 });
```

#### Metadata Filter Operators

```typescript
const filters = {
  // Where (metadata)
  where: { key: 'value' },                         // equals (string)
  where: { key: 42 },                               // equals (number)
  where: { key: { $eq: 'value' } },                 // equals
  where: { key: { $ne: 'value' } },                 // not equals
  where: { key: { $gt: 10 } },                      // greater than
  where: { key: { $gte: 10 } },                     // greater or equal
  where: { key: { $lt: 100 } },                     // less than
  where: { key: { $lte: 100 } },                    // less or equal
  where: { key: { $in: ['a', 'b', 'c'] } },         // in array

  // Logical
  where: { $and: [{ a: 1 }, { b: 2 }] },            // AND
  where: { $or: [{ a: 1 }, { b: 2 }] },             // OR

  // Nested
  where: { 'metadata.key': { $contains: 'sub' } }, // not supported on metadata

  // WhereDocument
  whereDocument: { $contains: 'text' },              // string contains
  whereDocument: { $not_contains: 'exclude' },       // string does not contain
};
```

#### Delete

```typescript
await collection.delete({
  ids: ['id1', 'id2', 'id3'],
  where: { source: 'temporary' },          // optional filter
});
```

#### Collection Metadata

```typescript
// Update metadata
await collection.modify({
  name: 'new_name',                        // rename
  metadata: { 'hnsw:space': 'ip' },        // update metadata
});

// Update metadata on collection creation:
// 'hnsw:space' - 'l2' (default), 'ip' (inner product), 'cosine'
// 'hnsw:construction_ef' - 100 (default)
// 'hnsw:M' - 16 (default), 'hnsw:search_ef' - 10 (default)
// 'chroma:immediate' - no auto-compaction
```

#### Tenant & Database (v0.5+)

```typescript
// === Multi-tenant ===
await client.createTenant({ name: 'my-tenant' });
await client.getTenant({ name: 'my-tenant' });

// === Database per tenant ===
await client.createDatabase({ name: 'my-db', tenantName: 'my-tenant' });

// Client scoped
const tenantClient = new ChromaClient({
  tenant: 'my-tenant',
  database: 'my-db',
});
```

### 1.3 Error Handling

```typescript
import { ChromaClient } from 'chromadb';

try {
  await client.getCollection({ name: 'nonexistent' });
} catch (error: any) {
  if (error.status === 404) {
    console.log('Collection not found');
  }
  if (error.name === 'ChromaUniqueError') {
    console.log('Duplicate ID');
  }
}
```

---

## 2. Qdrant JS Client (qdrant/qdrant-client)
**GitHub**: https://github.com/qdrant/qdrant-js | **Stars**: 700+
**npm**: `@qdrant/js-client-rest` | **Weekly**: ~60K | **License**: Apache 2.0
**Website**: https://qdrant.tech/

### 2.1 Installation
```bash
npm install @qdrant/js-client-rest
```

### 2.2 Complete API

#### Client Setup

```typescript
import { QdrantClient } from '@qdrant/js-client-rest';

// Local
const client = new QdrantClient({
  host: 'localhost',
  port: 6333,
  https: false,
  apiKey: 'optional-api-key',           // for cloud
});

// Cloud
const cloudClient = new QdrantClient({
  url: 'https://xxx.us-east-1-0.aws.cloud.qdrant.io:6333',
  apiKey: process.env.QDRANT_API_KEY,
});

// With timeout
const client2 = new QdrantClient({
  host: 'localhost',
  port: 6333,
  timeout: 30000,                       // 30s timeout
  checkCompatibility: true,             // check server version
});
```

#### Collections

```typescript
// === Create Collection ===
await client.createCollection('my_collection', {
  vectors: {
    size: 384,                          // vector dimension
    distance: 'Cosine',                 // 'Cosine' | 'Euclid' | 'Dot'
    on_disk: false,                     // store vectors on disk
  },
  // Optional config:
  shard_number: 1,
  replication_factor: 1,
  write_consistency_factor: 1,
  on_disk_payload: true,               // on-disk payload
  hnsw_config: {
    m: 16,
    ef_construct: 100,
    full_scan_threshold: 10000,
    max_indexing_threads: 0,
    on_disk: false,
  },
  wal_config: {
    wal_capacity_mb: 32,
    wal_segments_ahead: 0,
  },
  optimizers_config: {
    deleted_threshold: 0.2,
    vacuum_min_vector_number: 1000,
    default_segment_number: 0,
    max_segment_size: null,
    memmap_threshold_kb: 20000,
    indexing_threshold: 20000,
    flush_interval_sec: 5,
    max_optimization_threads: 4,
  },
  quantization_config: {
    // Scalar quantization
    scalar: {
      type: 'int8',
      quantile: 0.5,
      always_ram: true,
    },
    // Or Product quantization:
    product: {
      compression: 'x4',          // x4, x8, x16, x32, x64
      always_ram: true,
    },
  },
  init_from: { collection: 'base_collection' }, // copy from
});

// === Get Collection ===
const info = await client.getCollection('my_collection');
// info = { status, optimizer_status, vectors_count, indexed_vectors_count,
//          points_count, segments_count, config, payload_schema }

// === List Collections ===
const { collections } = await client.getCollections();
// [{ name: 'my_collection', status: 'green' }, ...]

// === Update Collection ===
await client.updateCollection('my_collection', {
  optimizers_config: { indexing_threshold: 50000 },
  // Can update: optimizers_config, params (hnsw_config, quantization),
  //             sparse_vectors
});

// === Delete Collection ===
await client.deleteCollection('my_collection');

// === Collection Exists ===
await client.collectionExists('my_collection');
// { exists: true | false }
```

#### Points (CRUD)

```typescript
// === Upsert Points ===
await client.upsert('my_collection', {
  wait: true,                          // wait for indexing
  ordering: 'weak',                    // 'weak' | 'medium' | 'strong'
  points: [
    {
      id: 1,                            // number or string (UUID)
      vector: [0.1, 0.2, 0.3, ...],
      payload: {
        title: 'Document 1',
        source: 'pdf',
        page: 1,
        tags: ['ai', 'ml'],
      },
    },
    {
      id: '550e8400-e29b-41d4-a716-446655440000',
      vector: [0.4, 0.5, 0.6, ...],
      payload: { title: 'Document 2' },
    },
  ],
});

// Named vectors (multiple vectors per point)
await client.upsert('my_collection', {
  points: [{
    id: 1,
    vector: {
      text: [0.1, 0.2, ...],          // name: vector
      image: [0.3, 0.4, ...],         // another vector
    },
    payload: { title: 'Multi-vector doc' },
  }],
});

// === Get Points ===
const points = await client.retrieve('my_collection', {
  ids: [1, 2, 3],
  with_payload: true,
  with_vector: true,
});

// === Update Points ===
await client.setPayload('my_collection', {
  payload: { new_field: 'value' },
  points: [1, 2, 3],
});

// Overwrite payload
await client.overwritePayload('my_collection', {
  payload: { title: 'New Title' },
  points: [1],
});

// Delete payload fields
await client.deletePayload('my_collection', {
  keys: ['temp_field'],
  points: [1, 2],
});

// Clear payload
await client.clearPayload('my_collection', {
  points: [1, 2, 3],
});

// === Delete Points ===
await client.delete('my_collection', {
  points: [1, 2, 3],                     // by IDs
  // OR by filter:
  filter: {
    must: [{ key: 'source', match: { value: 'temporary' } }],
  },
});
```

#### Search

```typescript
// === Basic Search ===
const results = await client.search('my_collection', {
  vector: [0.1, 0.2, 0.3, ...],
  limit: 10,
  offset: 0,
  with_payload: true,
  with_vector: false,
  score_threshold: 0.5,                  // min score
  params: {
    hnsw_ef: 128,                         // search ef
    exact: false,                         // exact search (slower, more accurate)
    quantization: null,                   // override quantization
  },
});

// result items:
// { id: 1, version: 0, score: 0.95, payload: {...}, vector: null }

// === Filtered Search ===
await client.search('my_collection', {
  vector: [0.1, 0.2, 0.3],
  limit: 10,
  filter: {
    must: [
      { key: 'source', match: { value: 'pdf' } },
      { key: 'page', range: { gte: 1, lte: 100 } },
    ],
    must_not: [
      { key: 'tags', match: { value: 'draft' } },
    ],
    should: [
      { key: 'author', match: { value: 'John' } },
    ],
    min_should: 1,
  },
});

// === Named Vector Search ===
// When using named vectors:
await client.search('my_collection', {
  vector: {
    name: 'text',
    vector: [0.1, 0.2, ...],
  },
  limit: 10,
});

// === Group Search (v1.10+) ===
await client.searchGroups('my_collection', {
  vector: [0.1, 0.2, 0.3],
  limit: 100,
  group_by: 'source',                    // field to group by
  group_size: 3,                          // results per group
});
```

#### Filter Types (Complete)

```typescript
const filter: QdrantFilter = {
  must: [
    // === Match (exact) ===
    { key: 'status', match: { value: 'active' } },         // string match
    { key: 'count', match: { value: 42 } },                 // number match
    { key: 'flag', match: { value: true } },                // boolean match
    { key: 'tags', match: { any: ['a', 'b', 'c'] } },     // any of
    { key: 'category', match: { except: ['x', 'y'] } },    // except these

    // === Range ===
    { key: 'price', range: { gt: 10, lt: 100 } },           // 10 < price < 100
    { key: 'age', range: { gte: 18 } },                      // price >= 18
    { key: 'score', range: { lte: 100 } },                   // score <= 100

    // === Geo ===
    { key: 'location', geo_radius: {
      center: { lat: 40.71, lon: -74.00 },
      radius: 1000,                                           // meters
    }},
    { key: 'location', geo_bounding_box: {
      top_left: { lat: 41.0, lon: -75.0 },
      bottom_right: { lat: 40.0, lon: -73.0 },
    }},

    // === Nested ===
    { key: 'metadata', nested: {
      filter: { must: [{ key: 'metadata.key', match: { value: 'val' } }] },
    }},

    // === IsEmpty / IsNull ===
    { key: 'optional_field', is_empty: true },
    { key: 'nullable_field', is_null: true },

    // === HasId ===
    { has_id: [1, 2, 3] },
  ],
  must_not: [
    // same match/range objects as must
  ],
  should: [
    // same match/range objects
  ],
  min_should: 0,
};
```

#### Scroll (Batch Retrieve)

```typescript
const scrollResult = await client.scroll('my_collection', {
  filter: { must: [{ key: 'source', match: { value: 'pdf' } }] },
  limit: 100,
  offset: null,                           // pagination cursor
  with_payload: true,
  with_vector: false,
  order_by: 'page',                       // order by payload field
});

// scrollResult = {
//   points: [...],
//   next_page_offset: '...' | null,     // for pagination
// }

// Paginate:
let offset: string | number | null = null;
do {
  const page = await client.scroll('my_collection', {
    limit: 100,
    offset,
  });
  processPage(page.points);
  offset = page.next_page_offset;
} while (offset !== null);
```

#### Batch Operations

```typescript
// === Update Batch ===
await client.batchUpdate('my_collection', {
  operations: [
    {
      upsert: {
        points: [{ id: 10, vector: [...], payload: {} }],
      },
    },
    {
      delete: {
        filter: { must: [{ key: 'temp', match: { value: true } }] },
      },
    },
  ],
  wait: true,
});
```

#### Count

```typescript
const countResult = await client.count('my_collection', {
  filter: { must: [{ key: 'source', match: { value: 'pdf' } }] },
  exact: true,
});

// { count: 150 }
```

#### Snapshots

```typescript
// Collection snapshot
const snapshot = await client.createSnapshot('my_collection');
const snapshots = await client.listSnapshots('my_collection');
await client.deleteSnapshot('my_collection', snapshot.name);

// Full storage snapshot
const storageSnapshot = await client.createFullSnapshot();
await client.listFullSnapshots();
```

#### Alias Management

```typescript
// Create alias
await client.createAlias({
  actions: [{
    create_alias: {
      collection_name: 'my_collection',
      alias_name: 'production',
    },
  }],
});

// Update alias (swap collections atomically)
await client.updateAliases({
  actions: [
    {
      delete_alias: { alias_name: 'old_alias' },
    },
    {
      create_alias: { collection_name: 'new_collection', alias_name: 'old_alias' },
    },
  ],
});

// List aliases
const aliases = await client.listAliases();
const collectionAliases = await client.listCollectionAliases('my_collection');
```

#### Cluster & Lock

```typescript
// Cluster info
const clusterInfo = await client.clusterInfo();

// Collection cluster info
const clusterInfo2 = await client.collectionClusterInfo('my_collection');

// Lock
await client.lockStorage({ reason: 'Maintenance' });
await client.unlockStorage();
const lock = await client.getLockStorage();
```

#### Error Handling

```typescript
import { QdrantClient } from '@qdrant/js-client-rest';

try {
  await client.search('nonexistent', { vector: [0.1], limit: 10 });
} catch (error: any) {
  if (error.status === 404) {
    console.log('Collection not found');
  }
  if (error instanceof QdrantClientError) {
    console.log(error.code, error.message);
  }
}
```

---

## 3. LanceDB (lancedb/lancedb)
**GitHub**: https://github.com/lancedb/lancedb | **Stars**: 5K+
**npm**: `vectordb` | **Weekly**: ~5K | **License**: Apache 2.0
**Website**: https://lancedb.github.io/lancedb/

### 3.1 Installation
```bash
npm install vectordb
# Note: Node.js bindings may have native dependency requirements
```

### 3.2 Complete API

**Note**: LanceDB JS API is newer and may differ slightly from Python. API below reflects current stable release.

#### Connection

```typescript
import * as lancedb from 'vectordb';

// === Local ===
const db = await lancedb.connect('./data/lancedb');
// Directory is created if it doesn't exist

// === Cloud (coming soon) ===
// const db = await lancedb.connect({
//   uri: 's3://bucket/lancedb',
//   region: 'us-east-1',
// });
```

#### Table Operations

```typescript
// === Create Table from data ===
const data = [
  { vector: [0.1, 0.2, 0.3], text: 'Document 1', source: 'pdf', page: 1 },
  { vector: [0.4, 0.5, 0.6], text: 'Document 2', source: 'web', url: 'example.com' },
];

const table = await db.createTable('my_table', data);

// With schema override (TypeScript)
import { Schema, Field, Float64, Utf8, Int32 } from 'vectordb';

const schema = new Schema([
  new Field('vector', new Float64()),
  new Field('text', new Utf8()),
  new Field('source', new Utf8()),
  new Field('page', new Int32()),
]);

const table2 = await db.createTable('typed_table', data, schema);

// === Open Table ===
const table3 = await db.openTable('my_table');

// === List Tables ===
const tableNames: string[] = await db.tableNames();

// === Drop Table ===
await db.dropTable('old_table');
```

#### Add / Update

```typescript
// === Add (append) ===
await table.add([
  { vector: [0.7, 0.8, 0.9], text: 'Document 3', source: 'note' },
]);

// === Update (by filter) ===
await table.update({
  where: 'source = "pdf"',
  values: { page: 100 },                    // set snapshot = 99
});
```

#### Search

```typescript
// === Vector Search ===
const results = await table.search([0.1, 0.2, 0.3])
  .limit(10)
  .execute();

// results = [
//   { vector: [...], text: '...', source: '...', _distance: 0.15 },
//   ...
// ]

// === Filtered Search ===
const filtered = await table.search([0.1, 0.2, 0.3])
  .where('source = "pdf"')
  .limit(5)
  .execute();

// === Metric Type ===
const results2 = await table.search([0.1, 0.2, 0.3])
  .metric('cosine')        // 'l2' (default), 'cosine', 'dot'
  .limit(10)
  .execute();

// === Prefilter ===
const results3 = await table.search([0.1, 0.2, 0.3])
  .where('source = "pdf"')
  .prefilter(true)         // apply filter before search
  .limit(10)
  .execute();
```

#### Delete

```typescript
await table.delete('source = "temporary"');
await table.delete('page IS NULL');
```

#### Query (SQL-like)

```typescript
// === SQL Query ===
const results = await db.query('SELECT * FROM my_table WHERE source = "pdf" LIMIT 10');

// === Count ===
const count = await table.countRows();

// === List all rows ===
const allData = await table.toArray();
```

#### Indices

```typescript
// === Create IVF-PQ Index ===
await table.createIndex({
  type: 'ivf_pq',              // IVF with Product Quantization
  metric: 'cosine',            // 'l2' | 'cosine' | 'dot'
  num_partitions: 256,         // IVF centroids
  num_sub_vectors: 96,         // PQ sub-vectors
});

// === HNSW Index (if supported) ===
await table.createIndex({
  type: 'hnsw',
  metric: 'cosine',
  ef_construction: 200,
  max_neighbors: 50,
});
```

#### Statistics

```typescript
// Table stats
const stats = await table.stats();
// { num_rows: number, num_indices: number, ... }

// Version
// LanceDB maintains versioning — you can query old versions
```

### 3.3 Error Handling

```typescript
import * as lancedb from 'vectordb';

try {
  await db.openTable('nonexistent');
} catch (error: any) {
  if (error.message.includes('not found')) {
    console.log('Table does not exist');
  }
  if (error instanceof lancedb.LanceError) {
    console.log('Lance error:', error.message);
  }
}
```

### 3.4 Configuration Notes

```typescript
// Connection options (second arg):
const db = await lancedb.connect('./data', {
  // Read consistency
  read_consistency_interval: {
    secs: 5,
    nanos: 0,
  },

  // AWS S3 config (for cloud storage)
  // storage_options: {
  //   region: 'us-east-1',
  //   access_key_id: '...',
  //   secret_access_key: '...',
  // },
});
```
