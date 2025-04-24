package main

import "time"

type Entries struct {
	Entry []Entry `json:"entry"`
}

type Entry struct {
	Title    string `json:"title"`
	ToDoList []struct {
		Item string `json:"item"`
	} `json:"toDoList"`
	AddedDate   time.Time `json:"addedDate"`
	ClosingDate time.Time `json:"closingDate"`
	Priority    *int      `json:"priority"`
}
