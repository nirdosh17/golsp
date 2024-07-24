# golsp
A language server in Go with following capabilities implemented:
- Hover
- Definition
- Completion
- Code actions
- Diagnostics

## How it works?
**Language Server Protocol (LSP)** was built by Microsoft initially for VS Code and is now an open standard for language servers and development tools to communicate with each other.

LSP defines message formats to perform following things:
- Handshake with the code editor (e.g. Neovim, VSCode) and notify list of capabilities the language server can do e.g. hover, definition search, completion, code actions e.t.c.
- Then, code editor will send details of events as different message types. For example:
  - initialize
  - textDocument/didOpen
  - textDocument/didChange
  - textDocument/hover
  - textDocument/completion
- Each message contains details about the file and the context (either full file content or partial).

  LSP's job is to parse these messages and build appropriate response in the format defined in LSP.

  For example, if LSP receives a hover message, we can build response something like below:
  ```json
    {
      "jsonrpc":"2.0",
      "id":2,
      "result": {
        "contents": "Doc: Readme.md Characters: 1333"
      }
    }
  ```
  Text editor understands this message and displays text inside the `contents` field to the user.


## Language server setup
- A language server does not start by it's own, it is controlled by the code editor.
- We specify how to run the LSP executable to the editor.
- In the project, we are using Go to build the language server. First, we compile the code into a binary called `markdownlsp`.
- Here is how we configure it in Neovim using lua:
  ```lua
  local client = vim.lsp.start_client {
    name = "ng_markdownlsp",
    -- location to LSP binary 'markdownlsp'
    cmd = { "/Users/nirdosh/Personal/golsp/markdownlsp" },
  }

  if not client then
    vim.notify "failed to load ng_markdownlsp plugin"
    return
  end

  -- telling neovim to attach this LSP to the buffer whenever you open a markdown file
  -- Type :LspInfo to see which LSP has been attached to current buffer
  vim.api.nvim_create_autocmd('FileType', {
    pattern = 'markdown',
    callback = function()
      vim.lsp.buf_attach_client(0, client)
    end,
  })
  ```

### LSP-Editor Communication Channels
We can communicate with the editor via different channels and most common is StandardIO where the LSP is present locally and controlled by the editor. In this project, we are using `stdio`.

Other alternatives are: TCP and Websockets.

---
### References:
- [Buiding a LSP in GO](https://www.youtube.com/watch?v=YsdlcQoHqPY) by **TJ DeVries**.
- [LSP Specs](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification)
