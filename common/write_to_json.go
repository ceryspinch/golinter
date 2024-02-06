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

// OpenJSONFile opens the JSON file for writing and initializes it with an open square bracket.
func OpenJSONFile(filePath string) (*os.File, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}

	// Write an open square bracket to the file
	if _, err = file.WriteString("["); err != nil {
		file.Close()
		return nil, err
	}

	return file, nil
}

func CloseJSONFile(filePath string) error {
	// Open the file in append mode
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// Write an closing square bracket to the file
	if _, err = file.WriteString("]"); err != nil {
		file.Close()
		return err
	}

	return nil
}

func AppendResultToJSON(result LintResult, filePath string) error {
	// Open the file in append mode
	file, err := os.OpenFile(filePath, os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	// Marshal the result to JSON
	jsonOutput, err := json.MarshalIndent(result, "", "    ")
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
	return append(append(data, ','), '\n')
}
