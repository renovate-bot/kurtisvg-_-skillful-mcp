package tools

import (
	"testing"

	"skillful-mcp/internal/clientmanager"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestListServerNamesSorted(t *testing.T) {
	mgr := clientmanager.NewFromSessions(map[string]*mcp.ClientSession{
		"charlie": nil,
		"alpha":   nil,
		"bravo":   nil,
	})

	names := mgr.ListServerNames()
	expected := []string{"alpha", "bravo", "charlie"}
	if len(names) != len(expected) {
		t.Fatalf("got %d names, want %d", len(names), len(expected))
	}
	for i, name := range names {
		if name != expected[i] {
			t.Errorf("names[%d] = %q, want %q", i, name, expected[i])
		}
	}
}

func TestListServerNamesEmpty(t *testing.T) {
	mgr := clientmanager.NewFromSessions(map[string]*mcp.ClientSession{})
	names := mgr.ListServerNames()
	if len(names) != 0 {
		t.Errorf("expected empty, got %v", names)
	}
}
