package main

import (
	"bufio"
	"encoding/json"
	"golsp/analysis"
	"golsp/lsp"
	"golsp/rpc"
	"io"
	"log"
	"os"
)

func main() {
	logger := getLogger("/Users/nirdosh/Personal/golsp/golsp_log.txt")
	logger.Println("go lsp started...")
	// give scanner something to read from
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("failed to decode msg: %s\n", err)
			continue
		}

		handleMessage(logger, writer, state, method, contents)
	}
	logger.Println("go lsp stopped")
}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, contents []byte) {
	logger.Printf("received msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Println("initialize err:", err)
			return
		}
		logger.Printf("connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)

		// now we need to reply initialize reponse to the editor
		msg := lsp.NewInitializeResponse(request.ID)
		writeResponse(writer, msg)
		logger.Println("initialize response sent!")
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Println("textDocument/didOpen err:", err)
			return
		}
		logger.Printf("opened: %s\n", request.Params.TextDocument.URI)
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)

	case "textDocument/didChange":
		var request lsp.DidChangeTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Println("textDocument/didChange err:", err)
			return
		}

		logger.Printf("changed: %s\n", request.Params.TextDocument.URI)
		for _, change := range request.Params.ContentChanges {
			state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Println("textDocument/hover err:", err)
			return
		}

		// create a response which will be displayed by the editor while hovering
		response := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		writeResponse(writer, response)
	}
}

// write message to given writer
// writer can be anything. e.g. stdout or http response
func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic("faild to initialize logfile" + err.Error())
	}

	return log.New(logfile, "[golsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
