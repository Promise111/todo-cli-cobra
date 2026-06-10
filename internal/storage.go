package internal

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}

var dirName = "data"

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	_, statErr := os.Stat(dirName)
	if !os.IsNotExist(statErr) && statErr != nil {
		return statErr
	}

	if os.IsNotExist(statErr) {
		dirErr := os.Mkdir(dirName, 0755)
		if dirErr != nil {
			return dirErr
		}
	}

	return os.WriteFile(dirName+"/"+s.FileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(dirName + "/" + s.FileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileData, data)
	if err != nil {
		return err
	}
	return nil
}
