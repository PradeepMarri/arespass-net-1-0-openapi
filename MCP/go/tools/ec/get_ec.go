package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/arespass/mcp-server/config"
	"github.com/arespass/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_ecHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["password"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("password=%v", val))
		}
		if val, ok := args["outputFormat"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("outputFormat=%v", val))
		}
		if val, ok := args["penalty"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("penalty=%v", val))
		}
		if val, ok := args["reqId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("reqId=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/ec%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.Ec
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGet_ecTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_ec",
		mcp.WithDescription("The entropy calculator - alias ec -, analyzes a password and calculates its entropy.

Entropy calculator requests are billed.
"),
		mcp.WithString("password", mcp.Required(), mcp.Description("**The password to be analyzed.**\n\nMinimum length is 4 characters; maximum length is 128 characters.\n\nBeware that certain characters like '&#35;', '&#61;' or '&#63;' must be properly encoded.\n\nFor more information about this issue, please refer to RFC 3986 (\"*Uniform Resource Identifier (URI): Generic Syntax*\"), sections 2.1, 2.2 and 2.4.\n")),
		mcp.WithString("outputFormat", mcp.Description("**The format of the returned analysis.**\n\nAllowed values are *json*, *xml* and *yaml*.\n\nThe default value is *xml*.\n")),
		mcp.WithString("penalty", mcp.Description("**The penalty applied to each character that is part of a word, number sequence, alphabet sequence, etc.**\n\nThe penalty is a float number in the range [0, 1]. Full penalty, 0; no penalty, 1.\n\nThe character used as decimal separator is always '&#46;'. Hence, a parameter value like *0,33* would be illegal.\n\nThe default value is *0.25*.\n")),
		mcp.WithString("reqId", mcp.Description("**An identifier for this request.**\n\nThe request identifier is a string that must match the regular expression */(?i)^[a-z0-9]{8,16}$/*.\n\nThis identifier is echoed in the returned response. Its value has no effect on the password analysis.\n\nIf this parameter is unset, a randomly generated identifier will be automatically assigned to this request.\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_ecHandler(cfg),
	}
}
