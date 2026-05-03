# Minimal GO MCP Server

This is a minimal implementation of a mcp server (written in Go) as an example project to learn the language and expand my knowledge towards ai. The server listens for incoming requests and responds with a greeting message when the "hello" tool is called.

## Getting Started
To run the server, follow these steps:

1. Build the server:
    ```bash
    $ make build
    ```
2. Run the server:
    ```bash
    $ ./bin/http
    ```

# Features
| Tool Name | Description |
|-----------|-------------|
| hello     | A simple tool that greets the user by name. |

## Example Request

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "tools/call",
  "params": {
    "name": "hello",
    "arguments": {
      "name": "Constantin"
    }
  }
}
```

## Example Response

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "content": [
      {
        "type": "text",
        "text": "Hello, Constantin! Thanks for using the hello tool."
      }
    ]
  }
}
```