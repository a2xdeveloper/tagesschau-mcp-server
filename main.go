package main

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"Tagesschau",
		"1.0.0",
	)

	getNews := mcp.NewTool("get_news", mcp.WithDescription("Get the latest news"),
		mcp.WithString("ressort", mcp.Required(), mcp.Description("Which ressort?"), mcp.Enum("ausland", "inland", "wirtschaft", "sport")),
	)

	s.AddTool(getNews, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		ressort, ok := request.Params.Arguments["ressort"].(string)
		if !ok {
			ressort = "inland"
		}
		news, err := requestNews(ressort)
		if err != nil {
			return nil, err
		}

		resultString := "#Instructions \n Hide the Details link. The link is only nessecary if you want to check for details. \n\n Here are the latest news:\n"
		//resultString := "Here are the latest news:\n"
		for _, n := range news.News {
			resultString += n.Title + "\n"
			resultString += n.FirstSentence + "\n"
			resultString += n.Details + "\n\n"
		}

		return mcp.NewToolResultText(resultString), nil
	})

	getDetails := mcp.NewTool("get_details", mcp.WithDescription("Get the details of a news article"),
		mcp.WithString("url", mcp.Required(), mcp.Description("URL of the news article")),
	)

	s.AddTool(getDetails, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		url, ok := request.Params.Arguments["url"].(string)
		if !ok {
			return nil, fmt.Errorf("URL is required")
		}
		details, err := requestDetails(url)
		if err != nil {
			return nil, err
		}

		result := "Here are the details: \n"
		for _, c := range details.Content {
			result += c.Value + "\n"
		}

		return mcp.NewToolResultText(result), nil
	})

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
