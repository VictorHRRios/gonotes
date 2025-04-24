package main

import "time"

type Entries struct {
	Entry []Entry `json:"entry"`
}

type Entry struct {
	ID        int        `json:"id"`
	AddedDate time.Time  `json:"addedDate"`
	Title     string     `json:"title"`
	ToDoList  []ToDoList `json:"toDoList"`
}

type ToDoList struct {
	AddedDate   time.Time `json:"addedDate"`
	ClosingDate time.Time `json:"closingDate"`
	Item        string    `json:"item"`
	Priority    *int      `json:"priority"`
}
