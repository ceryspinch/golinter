package common

import (
	"encoding/json"
	"os"
)

type LintResult struct {
	FilePath string `json:"file"`
	Line     int    `json:"line"`
	Message  string `json:"message"`
}

var LintResults []LintResult

func AppendResultToJSON(result LintResult, filePath string) error {
	// Open the file in append mode and create it if it doesn't exist
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Marshal the result to JSON
	jsonOutput, err := json.Marshal(result)
	if err != nil {
		return err
	}

	// Append the JSON data to the file
	if _, err := file.Write(appendNewline(jsonOutput)); err != nil {
		return err
	}

	return nil
}

func appendNewline(data []byte) []byte {
	return append(data, '\n')
}
