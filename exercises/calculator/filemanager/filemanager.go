package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	inputPath string
	outputPath string
}

func (fm FileManager) ReadLines() ([]string, error) {	
	file, err := os.Open(fm.inputPath)

	if err != nil {
		return nil, errors.New("Failed loading file.")
	}
	
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed reading file")
	}

	return lines, nil
}

func (fm FileManager) WriteData(data interface{}) (error) {
	file, err := os.Create(fm.outputPath)

	if err != nil {
		return errors.New("Failed to create new file")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	return err
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		inputPath: inputPath,
		outputPath: outputPath,
	}
}
