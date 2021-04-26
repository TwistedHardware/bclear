package main

import (
	"github.com/uadmin/uadmin"
	"time"
)

type Task struct {
	uadmin.Model
	Name string `uadmin:"search"`
	Description string `uadmin:"html;search"`
	Deadline time.Time `uadmin:"filter"`
	Progress float64 `uadmin:"progress_bar:40:red,70:yellow,100:lime"`
	Attachment string `uadmin:"image"`
}

func main() {
	uadmin.Register(
		Task{},
	)
	uadmin.StartServer()
}
