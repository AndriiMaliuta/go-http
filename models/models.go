package models

type Cat struct {
	Name  string
	Age   int
	Color string
}

type City struct {
	Name       string
	Population int64
}

type Person struct {
	Id     int64  `json:"personId"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}
