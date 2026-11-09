package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	hostname := os.Args[1]

	reader := reader()
	for {
		input, err := reader.ReadString('\n')
		if strings.TrimSpace(input) == "exit" {
			break
		}
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		cmd := Command{
			Cmd: strings.TrimSpace(input),
		}

		output, err := executeCommand(cmd, "http://"+hostname)

		if err != nil {
			fmt.Println("Error executing command:", err)
			return
		}

		fmt.Println(output)
	}
}

type Command struct {
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}

func reader() *bufio.Reader {
	reader := bufio.NewReader(os.Stdin)

	return reader
}

func executeCommand(cmd Command, serverURL string) (string, error) {
	data, err := json.Marshal(cmd)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(serverURL+"/execute/1/1", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
