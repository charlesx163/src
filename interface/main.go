package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	writer := customWriter{}
	io.Copy(writer, resp.Body)
}

type customWriter struct{}
type GitHubResponse []struct {
	FullName string `json:"full_name"`
}

func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GitHubResponse
	json.Unmarshal(p, &resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), nil
}
