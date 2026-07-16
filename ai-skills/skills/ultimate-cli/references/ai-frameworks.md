# AI & LLM Frameworks — Complete A-to-Z Reference

---

## 1. LangChain.js (langchain-ai/langchainjs)
**GitHub**: https://github.com/langchain-ai/langchainjs | **Stars**: 13K+
**npm**: `langchain` | **Weekly**: ~5.5M | **License**: MIT
**Website**: https://js.langchain.com/docs/

### 1.1 Installation
```bash
npm install langchain @langchain/openai @langchain/anthropic @langchain/community
# Vector stores:
npm install @langchain/pinecone @langchain/chroma @langchain/qdrant
# Document loaders:
npm install @langchain/document-loaders
```

### 1.2 Complete API

#### LLM Models

```typescript
import { ChatOpenAI } from '@langchain/openai';

// Chat model (recommended)
const model = new ChatOpenAI({
  modelName: 'gpt-4o-mini',    // or 'gpt-4o', 'gpt-3.5-turbo'
  temperature: 0.7,
  maxTokens: 1000,
  topP: 1,
  frequencyPenalty: 0,
  presencePenalty: 0,
  timeout: 30000,
  maxRetries: 2,
  streaming: true,
  callbacks: [{
    handleLLMNewToken(token: string) {
      console.log(token);
    },
  }],
  configuration: {
    baseURL: 'https://api.openai.com/v1',  // custom endpoint
    apiKey: process.env.OPENAI_API_KEY,
    defaultHeaders: { 'Custom-Header': 'value' },
  },
});

// Other providers
import { ChatAnthropic } from '@langchain/anthropic';
const anthropic = new ChatAnthropic({
  modelName: 'claude-3-5-sonnet-20241022',
  temperature: 0,
  maxTokens: 4096,
  anthropicApiKey: process.env.ANTHROPIC_API_KEY,
  streaming: true,
});

// Ollama
import { ChatOllama } from '@langchain/community/chat_models/ollama';
const ollama = new ChatOllama({
  model: 'llama3',
  baseUrl: 'http://localhost:11434',
  temperature: 0,
});

// Google Gemini
import { ChatGoogleGenerativeAI } from '@langchain/google-genai';
const gemini = new ChatGoogleGenerativeAI({
  modelName: 'gemini-pro',
  apiKey: process.env.GOOGLE_API_KEY,
});

// Groq (via community)
import { ChatGroq } from '@langchain/community/chat_models/groq';
const groq = new ChatGroq({
  modelName: 'mixtral-8x7b-32768',
  apiKey: process.env.GROQ_API_KEY,
});
```

#### Prompts

```typescript
import { ChatPromptTemplate } from '@langchain/core/prompts';

// === Simple template ===
const prompt = ChatPromptTemplate.fromMessages([
  ['system', 'You are a helpful assistant.'],
  ['user', 'Hello, my name is {name}.'],
]);
const formatted = await prompt.formatMessages({ name: 'Alice' });
// => [SystemMessage, HumanMessage]

// === Few-shot ===
import { FewShotPromptTemplate } from '@langchain/core/prompts';
const fewShotPrompt = new FewShotPromptTemplate({
  examples: [
    { input: 'hi', output: 'Hello!' },
    { input: 'bye', output: 'Goodbye!' },
  ],
  examplePrompt: ChatPromptTemplate.fromTemplate('Input: {input}\nOutput: {output}'),
  prefix: 'Here are some examples:',
  suffix: 'Input: {input}\nOutput:',
  inputVariables: ['input'],
});

// === Pipeline prompt ===
const composed = prompt.pipe(anotherPrompt);

// === Prompt with partial variables ===
const partialPrompt = await ChatPromptTemplate.fromTemplate('{greeting}, {name}!')
  .partial({ greeting: 'Hello' });

// === System/Human/Message types ===
import { SystemMessage, HumanMessage, AIMessage } from '@langchain/core/messages';
new SystemMessage('You are a bot');
new HumanMessage('Hello');
new AIMessage('Hi there!');
new HumanMessage({ content: 'Hello', name: 'user' });
```

#### Chains

```typescript
// === LLM Chain ===
import { LLMChain } from 'langchain/chains';

const chain = new LLMChain({
  llm: model,
  prompt: ChatPromptTemplate.fromMessages([
    ['system', 'You translate to {language}'],
    ['user', '{text}'],
  ]),
});
const result = await chain.call({ language: 'French', text: 'Hello' });

// === Simple Sequence ===
import { SequentialChain } from 'langchain/chains';

// === Conversation ===
import { ConversationChain } from 'langchain/chains';
const conversation = new ConversationChain({ llm: model, verbose: true });
await conversation.call({ input: 'Hi' });
await conversation.call({ input: 'What did I just say?' });

// === Runnable Interface (LCEL — LangChain Expression Language) ===
import { RunnableSequence, RunnablePassthrough } from '@langchain/core/runnables';

const chain2 = RunnableSequence.from([
  {
    input: new RunnablePassthrough(),
    context: async () => 'Some context',
  },
  prompt,
  model,
]);
const r = await chain2.invoke('Hello');

// or with pipe operator:
const pipeChain = prompt.pipe(model);
const result2 = await pipeChain.invoke({ topic: 'AI' });
```

#### Runnable Primitives

```typescript
import { RunnableSequence, RunnablePassthrough, RunnableMap } from '@langchain/core/runnables';

// === RunnablePassthrough (pass through / assign) ===
const chain3 = RunnableSequence.from([
  {
    question: new RunnablePassthrough(),
    context: retriever,
  },
  prompt,
  model,
]);

// === RunnableMap ===
const mapChain = RunnableMap.from({
  step1: chainA,
  step2: chainB,
});
// Returns combined results

// === RunnableBranch ===
import { RunnableBranch } from '@langchain/core/runnables';
const branch = RunnableBranch.from([
  [(x) => x > 0, positiveChain],
  [(x) => x < 0, negativeChain],
  fallbackChain,
]);

// === RunnableLambda ===
import { RunnableLambda } from '@langchain/core/runnables';
const toUpper = new RunnableLambda({ func: (x: string) => x.toUpperCase() });

// === RunnableWithMessageHistory ===
import { RunnableWithMessageHistory } from '@langchain/core/runnables';
const withHistory = new RunnableWithMessageHistory({
  runnable: model,
  getMessageHistory: (sessionId: string) => new ChatMessageHistory(),
  inputMessagesKey: 'input',
  historyMessagesKey: 'history',
});
```

#### Output Parsers

```typescript
import { StringOutputParser } from '@langchain/core/output_parsers';
import { StructuredOutputParser } from 'langchain/output_parsers';
import { JsonOutputParser } from '@langchain/core/output_parsers';
import { CommaSeparatedListOutputParser } from '@langchain/core/output_parsers';
import z from 'zod';

// === String ===
const parser = new StringOutputParser();
const result = await pipeChain.pipe(parser).invoke({ text: 'hi' });

// === JSON ===
const jsonParser = new JsonOutputParser();

// === Structured (Zod) ===
const schema = z.object({
  name: z.string(),
  age: z.number(),
  hobbies: z.array(z.string()),
});
const structParser = StructuredOutputParser.fromZodSchema(schema);
const prompt2 = ChatPromptTemplate.fromTemplate(
  'Extract info: {text}\n{format_instructions}'
);
const chain4 = prompt2.pipe(model).pipe(structParser);

// === List ===
const listParser = new CommaSeparatedListOutputParser();

// === Custom ===
import { BaseOutputParser } from '@langchain/core/output_parsers';
class CustomParser extends BaseOutputParser<string> {
  lc_namespace = ['custom'];
  async parse(text: string): Promise<string> {
    return text.trim();
  }
  getFormatInstructions(): string {
    return 'Return plain text';
  }
}
```

#### Embeddings

```typescript
import { OpenAIEmbeddings } from '@langchain/openai';

const embeddings = new OpenAIEmbeddings({
  modelName: 'text-embedding-3-small', // or 'text-embedding-3-large'
  dimensions: 256,          // for v3 models, reduce dimensions
  batchSize: 512,
  stripNewLines: true,
  timeout: 30000,
  maxRetries: 3,
});

// Embed single
const vector: number[] = await embeddings.embedQuery('Hello world');

// Embed multiple
const vectors: number[][] = await embeddings.embedDocuments([
  'Doc 1',
  'Doc 2',
  'Doc 3',
]);

// Other embedding models:
import { HuggingFaceInferenceEmbeddings } from '@langchain/community/embeddings/hf';
import { GoogleGenerativeAIEmbeddings } from '@langchain/google-genai';
import { OllamaEmbeddings } from '@langchain/community/embeddings/ollama';
```

#### Vector Stores

```typescript
import { Chroma } from '@langchain/community/vectorstores/chroma';
import { MemoryVectorStore } from 'langchain/vectorstores/memory';

// === In-memory ===
const memoryStore = await MemoryVectorStore.fromTexts(
  ['Doc 1', 'Doc 2', 'Doc 3'],
  [{ id: 1 }, { id: 2 }, { id: 3 }],   // metadata
  embeddings,
);

// === Chroma ===
const chromaStore = new Chroma(embeddings, {
  collectionName: 'my_collection',
  url: 'http://localhost:8000',
  collectionMetadata: { 'hnsw:space': 'cosine' },
});

// Add documents
await chromaStore.addDocuments(docs);

// Similarity search
const results = await chromaStore.similaritySearch('query', 5);
// Each result: Document { pageContent, metadata }

// With score
const scored = await chromaStore.similaritySearchWithScore('query', 5);
// [[Document, score], ...]

// Filtered search
const filtered = await chromaStore.similaritySearch('query', 5, {
  filter: { source: 'pdf' },
});

// Max marginal relevance (diverse results)
const diverse = await chromaStore.maxMarginalRelevanceSearch('query', 5, { fetchK: 20 });
```

#### Document Loaders

```typescript
import { TextLoader } from 'langchain/document_loaders/fs/text';
import { PDFLoader } from '@langchain/community/document_loaders/fs/pdf';
import { CSVLoader } from '@langchain/community/document_loaders/fs/csv';
import { JSONLoader } from 'langchain/document_loaders/fs/json';
import { DirectoryLoader } from 'langchain/document_loaders/fs/directory';

// === Text ===
const textLoader = new TextLoader('file.txt');
const docs = await textLoader.load();

// === PDF ===
const pdfLoader = new PDFLoader('file.pdf', {
  splitPages: true,
  parsedItemSeparator: '\n',
});
const pdfDocs = await pdfLoader.load();

// === CSV ===
const csvLoader = new CSVLoader('data.csv', {
  column: 'text',    // which column to use as content
  separator: ',',
});
const csvDocs = await csvLoader.load();

// === Directory (auto-detect loaders) ===
const dirLoader = new DirectoryLoader('./docs', {
  '.txt': (path) => new TextLoader(path),
  '.pdf': (path) => new PDFLoader(path),
  '.csv': (path) => new CSVLoader(path),
});
const allDocs = await dirLoader.load();

// === Web ===
import { CheerioWebBaseLoader } from 'langchain/document_loaders/web/cheerio';
const webLoader = new CheerioWebBaseLoader('https://example.com');
const webDocs = await webLoader.load();
```

#### Text Splitters

```typescript
import { RecursiveCharacterTextSplitter } from 'langchain/text_splitter';
import { TokenTextSplitter } from 'langchain/text_splitter';
import { CharacterTextSplitter } from 'langchain/text_splitter';

// === Recursive splitter (best for general text) ===
const splitter = new RecursiveCharacterTextSplitter({
  chunkSize: 1000,
  chunkOverlap: 200,
  separators: ['\n\n', '\n', '.', ' ', ''],
  keepSeparator: false,
  lengthFunction: (text) => text.length,
});

const chunks: Document[] = await splitter.splitDocuments(docs);
const texts: string[] = await splitter.splitText('Long text...');

// === Token splitter ===
const tokenSplitter = new TokenTextSplitter({
  encodingName: 'gpt2',     // tokenizer encoding
  chunkSize: 512,
  chunkOverlap: 50,
});

// === Markdown splitter ===
import { MarkdownTextSplitter } from 'langchain/text_splitter';
const mdSplitter = new MarkdownTextSplitter({
  chunkSize: 1000,
  chunkOverlap: 100,
});

// === Code splitter ===
import { LatexTextSplitter, PythonTextSplitter } from 'langchain/text_splitter';
```

#### Retrievers

```typescript
// === Vector store retriever ===
const retriever = chromaStore.asRetriever({
  searchType: 'similarity',      // 'similarity' | 'mmr'
  k: 5,
  fetchK: 20,                    // for MMR
  lambda: 0.5,                   // diversity vs similarity (MMR)
});

// === Multi-query retriever ===
import { MultiQueryRetriever } from 'langchain/retrievers/multi_query';
const multiRetriever = MultiQueryRetriever.fromLLM({
  retriever,
  llm: model,
});

// === Contextual compression ===
import { ContextualCompressionRetriever } from 'langchain/retrievers/contextual_compression';
import { LLMChainExtractor } from 'langchain/retrievers/document_compressors';

const compressor = LLMChainExtractor.fromLLM(model);
const compressionRetriever = new ContextualCompressionRetriever({
  baseCompressor: compressor,
  baseRetriever: retriever,
});

// === Hybrid search ===
// Some stores support hybrid:
import { Chroma } from '@langchain/community/vectorstores/chroma';
const hybridRetriever = chromaStore.asRetriever({
  searchType: 'similarity',
  // If store supports it, pass alpha for hybrid (0=keyword, 1=vector)
});

// === Ensemble retriever ===
import { EnsembleRetriever } from 'langchain/retrievers/ensemble';
const ensemble = new EnsembleRetriever({
  retrievers: [retriever, keywordRetriever],
  weights: [0.5, 0.5],
});
```

#### Tools

```typescript
import { DynamicTool } from '@langchain/core/tools';

const tool = new DynamicTool({
  name: 'current_time',
  description: 'Returns the current time',
  func: async () => new Date().toISOString(),
});

// === Built-in tools ===
import { TavilySearchResults } from '@langchain/community/tools/tavily_search';
import { Calculator } from '@langchain/community/tools/calculator';
import { SerpAPI } from '@langchain/community/tools/serpapi';

const search = new TavilySearchResults({ maxResults: 3 });
const calc = new Calculator();
```

#### Agents

```typescript
import { AgentExecutor, createToolCallingAgent } from 'langchain/agents';
import { ChatPromptTemplate } from '@langchain/core/prompts';

const tools = [search, calc, tool];

const prompt2 = ChatPromptTemplate.fromMessages([
  ['system', 'You are a helpful assistant.'],
  ['placeholder', '{chat_history}'],
  ['human', '{input}'],
  ['placeholder', '{agent_scratchpad}'],
]);

const agent = await createToolCallingAgent({
  llm: model,
  tools,
  prompt: prompt2,
});

const executor = new AgentExecutor({
  agent,
  tools,
  verbose: true,
  maxIterations: 10,        // prevent infinite loops
  returnIntermediateSteps: true,
  earlyStoppingMethod: 'generate',
  handleParsingErrors: true,
});

const result = await executor.invoke({
  input: 'What is the current time in Tokyo?',
  chat_history: [],
});
// { input: '...', output: '...', intermediateSteps: [...] }
```

#### Chat History

```typescript
import { BufferMemory } from 'langchain/memory';
import { ChatMessageHistory } from 'langchain/stores/message/in_memory';

const memory = new BufferMemory({
  returnMessages: true,
  memoryKey: 'chat_history',
  inputKey: 'input',
  outputKey: 'response',
  chatHistory: new ChatMessageHistory(),
});

// Add messages
await memory.chatHistory.addMessage(new HumanMessage('Hello'));
await memory.chatHistory.addMessage(new AIMessage('Hi!'));

// === Other stores ===
import { RedisChatMessageHistory } from '@langchain/redis';
import { PostgresChatMessageHistory } from '@langchain/community/stores/message/postgres';
```

#### Callbacks

```typescript
import { CallbackManager } from '@langchain/core/callbacks/manager';

const callbackManager = CallbackManager.configure({
  handleLLMStart: async (llm, prompts) => {
    console.log('LLM starting with prompts:', prompts);
  },
  handleLLMEnd: async (output) => {
    console.log('LLM finished:', output);
  },
  handleToolStart: async (tool, input) => {
    console.log('Tool start:', tool.name, input);
  },
  handleToolEnd: async (output) => {
    console.log('Tool output:', output);
  },
  handleChainStart: async (chain, inputs) => {
    console.log('Chain start:', chain.name);
  },
  handleChainEnd: async (outputs) => {
    console.log('Chain end:', outputs);
  },
  handleLLMNewToken: (token: string) => {
    process.stdout.write(token);
  },
});
```

#### Streaming

```typescript
const stream = await model.stream('Tell me a story');
for await (const chunk of stream) {
  console.log(chunk.content);
}

// With chain
const streamChain = prompt.pipe(model).pipe(new StringOutputParser());
const stream2 = await streamChain.stream({ topic: 'AI' });
for await (const chunk of stream2) {
  process.stdout.write(chunk);
}

// Event streaming (LangChain v0.2+)
const events = await chain.streamEvents({ input: 'Hello' }, { version: 'v2' });
for await (const event of events) {
  // { event: 'on_chain_start', name: '...', data: { ... }, tags: [] }
  // { event: 'on_chat_model_stream', data: { chunk: ... } }
}
```

#### Ecosystem

```typescript
// === LangServe (deploy chains as API) ===
// @langchain/langgraph-sdk
// import { Client } from '@langchain/langgraph-sdk';

// === LangGraph ===
// @langchain/langgraph
// StateGraph, MessageGraph, etc. (separate package)

// === Evaluation ===
// @langchain/evaluation — not yet fully released for JS
```

---

## 2. AI SDK by Vercel (vercel/ai)
**GitHub**: https://github.com/vercel/ai | **Stars**: 12K+
**npm**: `ai` | **Weekly**: ~2M | **License**: Apache 2.0
**Website**: https://sdk.vercel.ai/docs

### 2.1 Installation
```bash
npm install ai @ai-sdk/openai @ai-sdk/anthropic @ai-sdk/google
```

### 2.2 Complete API

#### Core Function: `generateText`

```typescript
import { generateText } from 'ai';
import { openai } from '@ai-sdk/openai';

const result = await generateText({
  model: openai('gpt-4o-mini'),
  system: 'You are a helpful assistant.',
  prompt: 'What is the capital of France?',
  maxTokens: 100,
  temperature: 0.7,
  topP: 1,
  presencePenalty: 0,
  frequencyPenalty: 0,
  stopSequences: ['\n'],
  seed: 42,
  maxRetries: 2,
  abortSignal: AbortSignal.timeout(10000),
  headers: { 'Custom-Header': 'value' },

  // Tool calling
  tools: {
    getWeather: {
      description: 'Get weather for a location',
      parameters: z.object({
        location: z.string(),
        unit: z.enum(['c', 'f']).default('c'),
      }),
      execute: async ({ location, unit }) => {
        return { temperature: 22, unit };
      },
    },
  },
  toolChoice: 'auto',           // 'auto' | 'required' | { type: 'tool', toolName: '...' }
  maxSteps: 5,                   // max tool call rounds

  // Transformers
  experimental_telemetry: { isEnabled: true, metadata: {} },
  experimental_output: undefined,  // for structured output
});

result.text;                    // 'Paris'
result.finishReason;            // 'stop' | 'length' | 'content-filter' | 'tool-calls' | 'error'
result.usage;                   // { promptTokens: number, completionTokens: number, totalTokens: number }
result.toolCalls;               // ToolCall[]
result.toolResults;             // ToolResult[]
result.steps;                   // each step of generation
result.warnings;                // warnings from provider
result.response;                // raw response
```

#### Core Function: `streamText`

```typescript
import { streamText } from 'ai';
import { openai } from '@ai-sdk/openai';

const result = streamText({
  model: openai('gpt-4o-mini'),
  prompt: 'Tell me a story',
  onChunk: ({ chunk }) => {
    if (chunk.type === 'text-delta') {
      process.stdout.write(chunk.textDelta);
    }
  },
  onFinish: ({ text, finishReason, usage }) => {
    console.log('Done:', { text, finishReason, usage });
  },
});

// Read from stream
const stream = result.textStream;
for await (const chunk of stream) {
  process.stdout.write(chunk);
}

// Full text promise
const fullText = await result.text;   // resolves when done
```

#### Core Function: `generateObject`

```typescript
import { generateObject } from 'ai';
import { openai } from '@ai-sdk/openai';
import z from 'zod';

const result = await generateObject({
  model: openai('gpt-4o-mini'),
  schema: z.object({
    name: z.string(),
    age: z.number(),
    hobbies: z.array(z.string()),
  }),
  prompt: 'Extract info: John is 30 and likes coding, reading',
});

result.object;     // { name: 'John', age: 30, hobbies: ['coding', 'reading'] }
result.finishReason;
result.usage;

// With list
const { object } = await generateObject({
  model: openai('gpt-4o-mini'),
  schema: z.object({
    name: z.string(),
    email: z.string().email(),
  }),
  output: 'array',             // generate array of objects
  prompt: 'Create 3 fake users',
});
// object → [{ name, email }, ...]

// With JSON schema instead of Zod:
const result2 = await generateObject({
  model: openai('gpt-4o-mini'),
  output: 'object',
  schemaName: 'person',
  schemaDescription: 'A person',
  // schema can be JSON Schema object
});
```

#### Core Function: `streamObject`

```typescript
const result = streamObject({
  model: openai('gpt-4o-mini'),
  schema: z.object({ name: z.string(), steps: z.array(z.string()) }),
  prompt: 'Plan a birthday party',
});

for await (const partialObject of result.partialObjectStream) {
  console.log(partialObject);
}

const finalObject = await result.object;
```

#### Core Function: `embed`

```typescript
import { embed, embedMany } from 'ai';
import { openai } from '@ai-sdk/openai';

const { embedding } = await embed({
  model: openai.embedding('text-embedding-3-small'),
  value: 'Hello world',
  maxRetries: 3,
});

const { embeddings } = await embedMany({
  model: openai.embedding('text-embedding-3-small'),
  values: ['Doc 1', 'Doc 2', 'Doc 3'],
});
```

#### Core Function: `tool`

```typescript
import { tool } from 'ai';

const weatherTool = tool({
  description: 'Get weather',
  parameters: z.object({
    location: z.string().describe('City name'),
  }),
  execute: async ({ location }) => {
    const response = await fetch(`https://api.weather.com/${location}`);
    return response.json();
  },
});
```

#### Model Providers

```typescript
import { openai } from '@ai-sdk/openai';
import { anthropic } from '@ai-sdk/anthropic';
import { google } from '@ai-sdk/google';
import { groq } from '@ai-sdk/groq';
import { mistral } from '@ai-sdk/mistral';
import { ollama } from 'ollama-ai-provider';

openai('gpt-4o');                    // string model ref
openai('gpt-4o', {               // with options
  structuredOutputs: true,
  parallelToolCalls: true,
  cacheControl: true,
});

anthropic('claude-3-5-sonnet-20241022');
google('gemini-1.5-pro');
groq('mixtral-8x7b-32768');
mistral('mistral-large-latest');
ollama('llama3');
```

#### Middleware

```typescript
import { generateText, type LanguageModelV1Middleware } from 'ai';

const middleware: LanguageModelV1Middleware = {
  transformParams: ({ params, type }) => {
    if (type === 'generate-text') {
      params.maxTokens = 100;
    }
    return params;
  },
  wrapGenerate: async ({ doGenerate }) => {
    const start = Date.now();
    const result = await doGenerate();
    console.log('Generation took', Date.now() - start, 'ms');
    return result;
  },
  wrapStream: async ({ doStream }) => {
    const { stream, ...rest } = await doStream();
    // wrap stream
    return { stream, ...rest };
  },
};
```

#### Telemetry

```typescript
generateText({
  model: openai('gpt-4'),
  prompt: 'Hello',
  experimental_telemetry: {
    isEnabled: true,
    functionId: 'my-function',
    metadata: { userId: '123' },
    recordInputs: true,
    recordOutputs: true,
  },
});
```

#### Provider Agnostic

```typescript
// The model parameter is typed so you can swap providers:
import { type LanguageModel } from 'ai';

async function chat(model: LanguageModel, prompt: string) {
  return generateText({ model, prompt });
}

// Use with any provider:
await chat(openai('gpt-4o'), 'Hi');
await chat(anthropic('claude-3-5-sonnet'), 'Hi');
```

---

## 3. LlamaIndex.TS (run-llama/LlamaIndexTS)
**GitHub**: https://github.com/run-llama/LlamaIndexTS | **Stars**: 1.4K+
**npm**: `llamaindex` | **Weekly**: ~25K | **License**: MIT
**Website**: https://ts.llamaindex.ai/

### 3.1 Installation
```bash
npm install llamaindex
```

### 3.2 Complete API

#### LLM

```typescript
import { Ollama, OpenAI, TogetherLLM, Anthropic } from 'llamaindex';

const llm = new OpenAI({
  model: 'gpt-4o-mini',
  temperature: 0.1,
  maxTokens: 1024,
  apiKey: process.env.OPENAI_API_KEY,
});

// Chat
const response = await llm.chat({
  messages: [
    { role: 'system', content: 'You are a helpful assistant.' },
    { role: 'user', content: 'Hello!' },
  ],
});
console.log(response.message.content);

// Complete (non-chat)
const completion = await llm.complete('The capital of France is ');
```

#### Embeddings

```typescript
import { OpenAIEmbedding } from 'llamaindex';

const embedModel = new OpenAIEmbedding({
  model: 'text-embedding-3-small',
  dimensions: 256,
});

const embedding = await embedModel.getTextEmbedding('Hello world');
const embeddings = await embedModel.getTextEmbeddings(['A', 'B', 'C']);
```

#### Documents

```typescript
import { Document } from 'llamaindex';

const doc = new Document({
  text: 'This is document text.',
  id_: 'doc1',
  metadata: { source: 'web', author: 'John' },
});

// Load from file
import { SimpleDirectoryReader } from 'llamaindex';
const reader = new SimpleDirectoryReader();
const docs = await reader.loadData('./data');
```

#### Indexing

```typescript
import { VectorStoreIndex, storageContextFromDefaults } from 'llamaindex';

// From documents
const index = await VectorStoreIndex.fromDocuments(docs, {
  serviceContext: { llm, embedModel },
  storageContext: await storageContextFromDefaults({ persistDir: './storage' }),
});

// From existing
const loadedIndex = await VectorStoreIndex.init({
  storageContext: await storageContextFromDefaults({ persistDir: './storage' }),
});
```

#### Querying

```typescript
const queryEngine = index.asQueryEngine({
  similarityTopK: 5,
  responseMode: 'compact',     // 'refine' | 'compact' | 'tree_summarize' | 'simple'
});

const response = await queryEngine.query({
  query: 'What is this document about?',
});

console.log(response.message.content);
console.log(response.sourceNodes);   // source nodes
```

#### Chat Engine

```typescript
const chatEngine = index.asChatEngine({
  messageHistory: [],
});

const response = await chatEngine.chat({
  message: 'Tell me more',
});
```

#### Ingestion Pipeline

```typescript
import { IngestionPipeline } from 'llamaindex';

const pipeline = new IngestionPipeline({
  transformations: [
    // text splitters, embedding, etc.
  ],
});

const nodes = await pipeline.run({ documents: docs });
```

#### Tool Calling

```typescript
import { QueryEngineTool, ToolMetadata } from 'llamaindex';

const tool = new QueryEngineTool({
  queryEngine,
  metadata: {
    name: 'docs_tool',
    description: 'Useful for querying documentation',
  },
});

const agent = new FunctionCallAgent({
  llm,
  tools: [tool],
  systemPrompt: 'You are a helpful assistant.',
});

const result = await agent.chat({ message: 'What is in the docs?' });
```

---

## 4. MCP (Model Context Protocol) — Client & Server
**GitHub**: https://github.com/modelcontextprotocol/typescript-sdk | **Stars**: 6K+
**npm**: `@modelcontextprotocol/sdk` | **License**: MIT
**Website**: https://modelcontextprotocol.io/

### 4.1 Installation
```bash
npm install @modelcontextprotocol/sdk
```

### 4.2 Complete API

#### Server

```typescript
import { Server } from '@modelcontextprotocol/sdk/server/index.js';
import { StdioServerTransport } from '@modelcontextprotocol/sdk/server/stdio.js';
import {
  CallToolRequestSchema,
  ListToolsRequestSchema,
  ListResourcesRequestSchema,
  ReadResourceRequestSchema,
  ListPromptsRequestSchema,
  GetPromptRequestSchema,
  ErrorCode,
  McpError,
} from '@modelcontextprotocol/sdk/types.js';

const server = new Server(
  {
    name: 'my-mcp-server',
    version: '1.0.0',
  },
  {
    capabilities: {
      tools: {},           // enable tools
      resources: {},       // enable resources
      prompts: {},         // enable prompts
      logging: {},         // enable logging
    },
  }
);
```

##### Tools

```typescript
server.setRequestHandler(ListToolsRequestSchema, async () => ({
  tools: [
    {
      name: 'get_weather',
      description: 'Get weather for a location',
      inputSchema: {
        type: 'object',
        properties: {
          location: { type: 'string', description: 'City name' },
          unit: { type: 'string', enum: ['c', 'f'] },
        },
        required: ['location'],
      },
    },
  ],
}));

server.setRequestHandler(CallToolRequestSchema, async (request) => {
  if (request.params.name === 'get_weather') {
    const { location, unit } = request.params.arguments as {
      location: string;
      unit?: string;
    };
    return {
      content: [
        {
          type: 'text',
          text: `The weather in ${location} is sunny, 22°${unit || 'c'}`,
        },
      ],
      isError: false,
    };
  }
  throw new McpError(ErrorCode.MethodNotFound, 'Tool not found');
});
```

##### Resources

```typescript
server.setRequestHandler(ListResourcesRequestSchema, async () => ({
  resources: [
    {
      uri: 'file:///docs/readme.txt',
      name: 'Readme',
      description: 'Project readme',
      mimeType: 'text/plain',
    },
  ],
}));

server.setRequestHandler(ReadResourceRequestSchema, async (request) => {
  const uri = request.params.uri;
  return {
    contents: [
      {
        uri,
        mimeType: 'text/plain',
        text: 'File contents...',
      },
    ],
  };
});
```

##### Prompts

```typescript
server.setRequestHandler(ListPromptsRequestSchema, async () => ({
  prompts: [
    {
      name: 'greeting',
      description: 'Generate a greeting',
      arguments: [
        {
          name: 'name',
          description: 'Name to greet',
          required: true,
        },
      ],
    },
  ],
}));

server.setRequestHandler(GetPromptRequestSchema, async (request) => {
  if (request.params.name === 'greeting') {
    return {
      messages: [
        {
          role: 'user',
          content: {
            type: 'text',
            text: `Hello ${request.params.arguments?.name}!`,
          },
        },
      ],
    };
  }
  throw new McpError(ErrorCode.MethodNotFound, 'Prompt not found');
});
```

##### Logging

```typescript
// Send log messages to client
await server.sendLoggingMessage({
  level: 'info',      // 'debug' | 'info' | 'warning' | 'error'
  data: 'Server starting...',
});
```

##### Transport

```typescript
// === Stdio ===
const transport = new StdioServerTransport();
await server.connect(transport);

// === SSE (experimental) ===
import { SSEServerTransport } from '@modelcontextprotocol/sdk/server/sse.js';
const transport2 = new SSEServerTransport('/messages', response);
await server.connect(transport2);

// === Custom ===
const transport3 = new CustomTransport();
await server.connect(transport3);
```

#### Client

```typescript
import { Client } from '@modelcontextprotocol/sdk/client/index.js';
import { StdioClientTransport } from '@modelcontextprotocol/sdk/client/stdio.js';
import {
  ListToolsResultSchema,
  CallToolResultSchema,
  ListResourcesResultSchema,
  ReadResourceResultSchema,
} from '@modelcontextprotocol/sdk/types.js';

const client = new Client(
  {
    name: 'mcp-client',
    version: '1.0.0',
  },
  {
    capabilities: {},
  }
);

// Connect to server
const transport = new StdioClientTransport({
  command: 'node',
  args: ['./server.js'],
});
await client.connect(transport);

// List tools
const toolsResult = await client.request(
  { method: 'tools/list' },
  ListToolsResultSchema
);
console.log(toolsResult.tools);

// Call tool
const result = await client.request(
  {
    method: 'tools/call',
    params: { name: 'get_weather', arguments: { location: 'Paris' } },
  },
  CallToolResultSchema
);

// List resources
const resources = await client.request(
  { method: 'resources/list' },
  ListResourcesResultSchema
);

// Read resource
const resource = await client.request(
  {
    method: 'resources/read',
    params: { uri: 'file:///docs/readme.txt' },
  },
  ReadResourceResultSchema
);
```

#### Error Handling

```typescript
import { McpError, ErrorCode } from '@modelcontextprotocol/sdk/types.js';

// Predefined codes:
ErrorCode.ParseError;              // -32700
ErrorCode.InvalidRequest;          // -32600
ErrorCode.MethodNotFound;          // -32601
ErrorCode.InvalidParams;           // -32602
ErrorCode.InternalError;           // -32603

// Throw errors
throw new McpError(ErrorCode.InternalError, 'Something went wrong');
throw new McpError(ErrorCode.InvalidParams, 'Missing required field');
```

#### TypeScript Types

```typescript
import type {
  Tool,
  Resource,
  Prompt,
  CallToolRequest,
  ReadResourceRequest,
  GetPromptRequest,
  TextContent,
  ImageContent,
  EmbeddedResource,
} from '@modelcontextprotocol/sdk/types.js';

// Content types
const textContent: TextContent = { type: 'text', text: 'hello' };
const imageContent: ImageContent = { type: 'image', data: 'base64...', mimeType: 'image/png' };
const embeddedResource: EmbeddedResource = {
  type: 'resource',
  resource: { uri: '...', text: '...' },
};
```
