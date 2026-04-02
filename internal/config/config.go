package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type TransportType string

const (
	TransportSTDIO TransportType = "stdio"
	TransportHTTP  TransportType = "http"
	TransportSSE   TransportType = "sse"
)

type Config struct {
	MCPServers map[string]ServerConfig `json:"mcpServers"`
}

type ServerConfig struct {
	Type    string            `json:"type,omitempty"`
	Command string            `json:"command,omitempty"`
	Args    []string          `json:"args,omitempty"`
	Env     map[string]string `json:"env,omitempty"`
	URL     string            `json:"url,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}

func (s *ServerConfig) TransportType() (TransportType, error) {
	switch s.Type {
	case "", "stdio":
		return TransportSTDIO, nil
	case "http":
		return TransportHTTP, nil
	case "sse":
		return TransportSSE, nil
	default:
		return "", fmt.Errorf("unknown transport type: %q", s.Type)
	}
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing config file: %w", err)
	}

	return &cfg, nil
}

func (c *Config) Validate() error {
	for name, srv := range c.MCPServers {
		tt, err := srv.TransportType()
		if err != nil {
			return fmt.Errorf("server %q: %w", name, err)
		}
		switch tt {
		case TransportSTDIO:
			if srv.Command == "" {
				return fmt.Errorf("server %q: STDIO transport requires 'command'", name)
			}
		case TransportHTTP, TransportSSE:
			if srv.URL == "" {
				return fmt.Errorf("server %q: %s transport requires 'url'", name, tt)
			}
		}
	}
	return nil
}
