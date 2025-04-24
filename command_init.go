package main

import (
	"os"
)

func commandInit(s ...string) error {
	filePath, err := getEntriesFilePath()
	if err != nil {
		return err
	}
	_, err = os.Create(filePath)
	if err != nil {
		return err
	}
	return nil
}
