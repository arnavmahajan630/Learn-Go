# Simple Agent

A minimal Go example project that demonstrates a tiny command-based agent.

## Overview

This repository contains a small Go program with commands implemented under the `cmd` package. It's intended as a beginner-friendly starting point for exploring how to structure a simple CLI/agent in Go.

## Prerequisites

- Go 1.18 or newer installed

## Quick start

Clone the repo and run the program:

```bash
go run main.go
```

Build a binary:

```bash
go build -o simple-agent
./simple-agent
```

Run tests:

```bash
go test ./cmd
```

## Project structure

- [go.mod](go.mod) — module definition
- [main.go](main.go) — program entrypoint
- [cmd/ai.go](cmd/ai.go) — AI-related command code
- [cmd/home.go](cmd/home.go) — home/landing command
- [cmd/service.go](cmd/service.go) — core service logic
- [cmd/service_test.go](cmd/service_test.go) — unit tests for service logic

## Development notes

- Keep code organized under `cmd` for command implementations and small services.
- Use `go fmt` and `go vet` to maintain code quality.

## Contributing

Contributions are welcome. Open issues or submit pull requests with small, focused changes.

## Future improvements (scope)

Planned enhancements to increase flexibility and output quality:

- **Model selection:** allow users to choose the model they want to chat with (remote APIs or local/runnable models) via configuration or CLI flags.
- **Chat histories:** persist conversation histories per user/session so the agent can maintain context across turns and sessions.
- **Vector database / RAG:** add embeddings + a vector database (e.g., Pinecone, Milvus, Weaviate, or a local FAISS-based store) to enable retrieval-augmented generation for more accurate and relevant responses.
- **Configurable retrieval strategy:** support windowed context, similarity thresholds, and hybrid search (embedding + keyword) to improve result quality.
- **Privacy & storage controls:** provide options for local-only histories, retention policies, and opt-in telemetry so users control where data is stored.
- **Pluggable backends:** design interfaces for swapping model providers, embedding services, and vector stores without changing app code.

## License

This project is provided as-is. Add a license file if you plan to re-use or publish.
