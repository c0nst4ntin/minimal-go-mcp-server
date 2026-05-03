package main

import (
	"fmt"
	"net/http"

	"github.com/c0nst4ntin/go-mcp-server-test/internal/server"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	handler := mcp.NewStreamableHTTPHandler(func(_ *http.Request) *mcp.Server {
		return server.NewServer()
	}, nil)

	mux := http.NewServeMux()
	mux.Handle("/mcp", handler)

	httpServer := &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: mux,
	}

	fmt.Println("Starting server on http://127.0.0.1:3000")
	httpServer.ListenAndServe()
}
