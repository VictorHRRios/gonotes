package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

func commandAdd(params ...string) error {
	jsonPath, err := getEntriesFilePath()
	if err != nil {
		return err
	}

	readNoteFile, err := read(jsonPath)
	if err != nil {
		return err
	}
	fmt.Print(readNoteFile)
	fmt.Printf("params: %v\n", params)
	if len(params) == 1 {
		for key, value := range readNoteFile.Entry {
			if value.Title == params[0] {
				readNoteFile.Entry[key].ToDoList = append(readNoteFile.Entry[key].ToDoList, ToDoList{
					AddedDate:   time.Now(),
					ClosingDate: time.Now(),
					Item:        "prueba",
					Priority:    nil,
				})
				if err := update(readNoteFile, jsonPath); err != nil {
					return err
				}
				return nil
			}
		}
		return fmt.Errorf("Entry not found\n")
	}

	fmt.Printf("Title of the entry: ")
	reader := bufio.NewReader(os.Stdin)
	title, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	newEntry := Entry{
		ID:        len(readNoteFile.Entry),
		AddedDate: time.Now(),
		Title:     strings.TrimSpace(title),
		ToDoList:  nil,
	}
	if err := write(readNoteFile, newEntry, jsonPath); err != nil {
		return err
	}
	return nil
}

func read(filePath string) (Entries, error) {
	jsonFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return Entries{}, fmt.Errorf("run command init to initalize: %v\n", err)
	}
	defer jsonFile.Close()
	var entries Entries
	decoder := json.NewDecoder(jsonFile)

	if err = decoder.Decode(&entries); err != nil {
		return Entries{}, nil
	}
	return entries, nil
}

func update(e Entries, filePath string) error {
	jsonFile, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "    ")
	if err = encoder.Encode(e); err != nil {
		return err
	}
	return nil
}

func write(e Entries, newEntry Entry, filePath string) error {
	jsonFile, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	e.Entry = append(e.Entry, newEntry)

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "    ")
	if err = encoder.Encode(e); err != nil {
		return err
	}
	return nil
}

func getEntriesFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filePath := path.Join(home, entriesFile)
	return filePath, nil
}
