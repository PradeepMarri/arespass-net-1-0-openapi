package main

import (
	"github.com/arespass/mcp-server/config"
	"github.com/arespass/mcp-server/models"
	tools_about "github.com/arespass/mcp-server/tools/about"
	tools_ec "github.com/arespass/mcp-server/tools/ec"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_about.CreateGet_aboutTool(cfg),
		tools_ec.CreateGet_ecTool(cfg),
	}
}
