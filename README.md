# tagesschau-mcp-server

## Overview
This is an MCP (Model Context Protocol) server for [tagesschau.de](https://www.tagesschau.de), providing access to the latest news articles and details from the tagesschau platform.

## Features
- Fetch the latest news articles from various categories.
- Retrieve detailed information about specific news articles.

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/a2xdevelopment/tagesschau-mcp-server.git
   cd tagesschau-mcp-server
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the server:
   ```bash
   go build -o tagesschau 
   ```

## Usage
Once the server is running, you can access the following endpoints:
- **Get Latest News**: Fetch the latest news articles.
- **Get Article Details**: Retrieve details for a specific article by providing its URL.

## Tagesschau Configuration Example
To configure the tagesschau MCP server, add the following settings to your MCP configuration file:

```json
{
  "mcpServers": {
    "tagesschau": {
      "command": "/path/to/tagesschau",
      "args": [],
      "env": {}
    }
  }
}
```

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments
- [tagesschau.de](https://www.tagesschau.de) for providing the news API.
