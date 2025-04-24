package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"
)

const entriesFile = ".entries"

func commandAdd(params ...string) error {
	fmt.Printf("Title of the entry: ")
	reader := bufio.NewReader(os.Stdin)
	title, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	//fmt.Printf("Closing date of the entry: ")
	//closingDate, err := reader.ReadString('\n')
	//if err != nil {
	//return err
	//}
	//closingDate = strings.TrimSpace(closingDate)
	//parsedClosingDate, err := time.Parse("2006-Jan-02", closingDate)
	//if err != nil {
	//return err
	//}

	jsonPath, err := getEntriesFilePath()
	if err != nil {
		return err
	}

	readNoteFile, err := read(jsonPath)
	if err != nil {
		return err
	}
	newEntry := Entry{
		Title:       title,
		ToDoList:    nil,
		AddedDate:   time.Now(),
		ClosingDate: time.Now(),
		Priority:    nil,
	}
	if err := write(readNoteFile, newEntry, jsonPath); err != nil {
		return err
	}
	return nil
}

func read(filePath string) (Entries, error) {
	jsonFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return Entries{}, err
	}
	defer jsonFile.Close()
	var entries Entries
	decoder := json.NewDecoder(jsonFile)

	if err = decoder.Decode(&entries); err != nil {
		return Entries{}, nil
	}
	return entries, nil
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
