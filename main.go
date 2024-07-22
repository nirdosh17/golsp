package main

import (
	"bufio"
	"encoding/json"
	"golsp/lsp"
	"golsp/rpc"
	"log"
	"os"
)

func main() {
	logger := getLogger("/Users/nirdosh/Personal/golsp/golsp_log.txt")
	logger.Println("go lsp started...")
	// give scanner something to read from
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("failed to decode msg: %s\n", err)
			continue
		}

		handleMessage(logger, method, contents)
	}
	logger.Println("go lsp stopped")
}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("received msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Println("failed to parse initialize req:", err)
		}
		logger.Printf("connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)

		// now we need to reply initialize reponse to the editor
		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)
		writer := os.Stdout
		writer.Write([]byte(reply))

		logger.Println("initialize response sent!")
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Println("failed to parse initialize req:", err)
		}
		logger.Printf("text doc open notif: %s %s\n",
			request.Params.TextDocument.URI,
			request.Params.TextDocument.Text)
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic("faild to initialize logfile" + err.Error())
	}

	return log.New(logfile, "[golsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
