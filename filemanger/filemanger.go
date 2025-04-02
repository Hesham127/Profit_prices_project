package filemanger

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Filemanger struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm Filemanger) ReadFile() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("failed to scan line")
	}
	return lines, nil
}

func (fm Filemanger) WriteResult(data any) error {
	file, err := os.Create(fmt.Sprintf("%v.json", fm.OutputFilePath))

	if err != nil {
		return errors.New("failed to create the file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to convert data to json")
	}

	return nil

}
func New(inputFilePath, outputFilePath string) Filemanger {
	return Filemanger{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}
