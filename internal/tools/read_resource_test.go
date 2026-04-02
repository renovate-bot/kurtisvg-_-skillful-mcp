package tools

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestReadResourceReturnsContent(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	session := newFakeSession(t, ctx, func(s *mcp.Server) {
		s.AddResource(&mcp.Resource{URI: "test://hello", Name: "hello"}, func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
			return &mcp.ReadResourceResult{
				Contents: []*mcp.ResourceContents{
					{URI: "test://hello", MIMEType: "text/plain", Text: "Hello, World!"},
				},
			}, nil
		})
	})

	result, err := session.ReadResource(ctx, &mcp.ReadResourceParams{URI: "test://hello"})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Contents) != 1 {
		t.Fatalf("expected 1 content, got %d", len(result.Contents))
	}
	if result.Contents[0].Text != "Hello, World!" {
		t.Errorf("text = %q, want 'Hello, World!'", result.Contents[0].Text)
	}
}

func TestReadResourceUnknownURI(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	session := newFakeSession(t, ctx)
	_, err := session.ReadResource(ctx, &mcp.ReadResourceParams{URI: "test://nonexistent"})
	if err == nil {
		t.Fatal("expected error for unknown resource URI")
	}
}
