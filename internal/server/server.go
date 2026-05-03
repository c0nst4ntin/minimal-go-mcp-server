package server

import (
	"github.com/c0nst4ntin/go-mcp-server-test/internal/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func NewServer() *mcp.Server {
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "go-mcp-server",
			Version: "1.0.0",
		},
		&mcp.ServerOptions{
			Capabilities: &mcp.ServerCapabilities{
				Tools: &mcp.ToolCapabilities{
					ListChanged: false,
				},
			},
		},
	)

	mcp.AddTool(server, tools.HelloTool, tools.HelloHandler)

	return server
}
