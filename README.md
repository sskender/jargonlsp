# Jargon LSP Server

Understand internal jargon instantly, directly within your editor.

## Overview

Working on a large codebase filled with company-specific terminology, acronyms, or domain-specific jargon? Keeping track of all these can be painful, especially when scattered across multiple documents and teams.

**Jargon LSP Server** is a lightweight, Language Server Protocol ([LSP](https://en.wikipedia.org/wiki/Language_Server_Protocol)) server written in Go. It allows you to query definitions of unknown terms directly in your code editor, powered by a customizable dictionary file. The server is editor-agnostic and filetype-agnostic. It integrates seamlessly with any LSP-capable editor, such as Neovim, VSCode, Sublime, etc.

*Tested primarily with Neovim, but should work across different editors.*

## Why LSP?

You might wonder: why implement this as an LSP server instead of a native plugin?

- **Editor-Agnostic**: One implementation, usable across all editors that support LSP.

- **Filetype-Agnostic**: Works with any file type - source code, configs, documentation, etc.

- **Extensible & Scalable**: Easier to evolve, debug, and add features (e.g., TCP support in the future).

- **Learning Value**: A practical project to dive deep into LSP internals, communication protocols, and editor integrations - Neovim especially.

## Features

- Customizable Dictionary: Load any JSON dictionary file with term-definition mappings.

- Hover Support: Get definitions by hovering over tokens.

- Multiple Dictionaries: Use per-project dictionaries by configuring dictionary paths.

- Minimal Dependencies: No overhead - runs anywhere Go runs.

## Installation

### Prerequisites

- Go v1.18+

### Build & Install

```sh
git clone https://github.com/sskender/jargonlsp
cd jargonlsp
go build -o jargonlsp
sudo mv jargonlsp /usr/local/bin/
```

### Prepare Dictionary

Create a JSON dictionary file structured like:

```json
{
    "AMM": "A decentralized asset trading pool that enables market participants to buy or sell cryptocurrencies. Uniswap is the most well-known AMM.",
    "APY": "Annual Percentage Yield, a time-based measurement of the Return On Investment (ROI) on an asset.",
    "dApp": "A decentralized Web3 software application that normally runs on a blockchain.",
    "DAO": "Distributed Autonomous Organization.",
    "HODL": "HODL was initially a spelling error of ‘hold’ that became a term that was embraced as an inside joke by the early adopters of Bitcoin and Ethereum.",
    "Oracle": "A trusted feed of data, such as the current market prices of an asset or assets, that provides confidence to users that the data are timely, accurate, and untampered.",
    "ROI": "Return On Investment. The gains or losses on an investment.",
    "TVL": "The Total Value Locked into a Smart Contract or set of Smart Contracts that may be deployed or stored at one or more exchanges or markets.",
}
```

You can maintain one central glossary or project-specific dictionaries, and even optionally automate updates by parsing internal documentation.

### Editor Setup

Example for Neovim:

```lua
local lspconfig = require("lspconfig")
local configs = require("lspconfig.configs")

configs.jargonlsp = {
    default_config = {
        cmd = { "jargonlsp", "--dictionary=glossary.json" },
        filetypes = { "json", "python" },
        root_dir = vim.fn.getcwd(),
        capabilities = lspconfig.util.default_config.capabilities,
    }
}

lspconfig.jargonlsp.setup({
    on_attach = function(client, bufnr)
        print("JargonLSP attached successfully")
    end
})
```

Use the `hover` action to view definitions under your cursor.
If you are already using LSP hovers for other servers, you can bind this one to a separate key to avoid conflicts.

## Future Work

- TCP Support: Current implementation uses standard input/output (stdio) communication. Future versions will add optional TCP-based communication to support more flexible setups like remote usage.

- Advanced Matching: Token normalization, fuzzy matching, and support for multi-word phrases.

## Disclaimer & Considerations

- The LSP server does not attempt to parse language syntax. It operates purely on token boundaries.

- Dictionary accuracy and updates are user-maintained. Automation scripts for parsing company glossaries are recommended.

- Be mindful of editor-specific LSP quirks when configuring LSP settings.
