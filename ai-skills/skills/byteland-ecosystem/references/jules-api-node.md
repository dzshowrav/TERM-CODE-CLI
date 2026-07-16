# jules-api-node — Reference

> **NPM**: `jules-api-node`
> **GitHub**: https://github.com/ByteLandTechnology/jules-api-node
> **License**: MIT

The **unofficial** TypeScript client for the Jules API. Auto-generated from the Jules OpenAPI specification using `@hey-api/openapi-ts`.

---

## Available Methods

| Method | Description |
|--------|-------------|
| `approvePlan` | Approve a plan |
| `listSessions` | List all sessions |
| `createSession` | Create a new session |
| `getSession` | Get session details |
| `sendMessage` | Send a message to a session |
| `getActivity` | Get activity details |
| `listActivities` | List all activities |
| `getSource` | Get source details |
| `listSources` | List all sources |

---

## Usage

```ts
import { createSession } from "jules-api-node";

async function initializeSession() {
  try {
    const response = await createSession({
      headers: {
        "X-Goog-Api-Key": "YOUR_API_KEY_HERE",
      },
      body: {
        // session data
      },
    });
    console.log("Session created:", response.data);
    return response.data;
  } catch (error) {
    console.error("Error creating session:", error);
  }
}
```

## Authentication

Pass your API key in the `X-Goog-Api-Key` header for every request.

## Build from Source

```bash
npm install
npm run build  # Regenerates client from jules-openapi.yaml
```

## API Reference

For parameters and return types, refer to the TypeScript definitions in the package or the [OpenAPI specification](https://github.com/ByteLandTechnology/blob/jules-api-node/main/jules-openapi.yaml).
