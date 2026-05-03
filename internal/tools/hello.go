package tools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// HelloTool represents the hello greeting tool definition
var HelloTool = &mcp.Tool{
	Name:        "hello",
	Description: "A simple greeting tool that returns a personalized hello message",
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type":        "string",
				"description": "The name of the person to greet",
			},
		},
		"required": []string{"name"},
	},
}

// HelloInput represents the input for the hello tool
type HelloInput struct {
	Name string `json:"name"`
}

// HelloHandler handles the hello tool execution
func HelloHandler(_ context.Context, _ *mcp.CallToolRequest, input HelloInput) (*mcp.CallToolResult, any, error) {
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("Hello, %s! Thanks for using the hello tool.", input.Name)},
		},
	}, nil, nil
}
