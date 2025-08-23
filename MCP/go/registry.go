package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_apimanagementservice "github.com/apimanagementclient/mcp-server/tools/apimanagementservice"
	tools_apimanagementoperations "github.com/apimanagementclient/mcp-server/tools/apimanagementoperations"
	tools_apimanagementserviceskus "github.com/apimanagementclient/mcp-server/tools/apimanagementserviceskus"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_apimanagementservice.CreateApimanagementservice_checknameavailabilityTool(cfg),
		tools_apimanagementoperations.CreateApimanagementoperations_listTool(cfg),
		tools_apimanagementservice.CreateApimanagementservice_listbyresourcegroupTool(cfg),
		tools_apimanagementservice.CreateApimanagementservice_getTool(cfg),
		tools_apimanagementservice.CreateApimanagementservice_updateTool(cfg),
		tools_apimanagementservice.CreateApimanagementservice_createorupdateTool(cfg),
		tools_apimanagementservice.CreateApimanagementservice_deleteTool(cfg),
		tools_apimanagementservice.CreateApimanagementservice_backupTool(cfg),
		tools_apimanagementservice.CreateApimanagementservice_getssotokenTool(cfg),
		tools_apimanagementservice.CreateApimanagementservice_listTool(cfg),
		tools_apimanagementservice.CreateApimanagementservice_applynetworkconfigurationupdatesTool(cfg),
		tools_apimanagementservice.CreateApimanagementservice_restoreTool(cfg),
		tools_apimanagementserviceskus.CreateApimanagementserviceskus_listavailableserviceskusTool(cfg),
	}
}
