package main

import (
	"bufio"
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
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
	logger.Println("go lsp stopped")
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic("faild to initialize logfile" + err.Error())
	}

	return log.New(logfile, "[golsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
