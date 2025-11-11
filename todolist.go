package main

import (
	"fmt"
	"bufio"
	"os"
)

func outputTodoList(todoList **[]string) {
	fmt.Print("\n")
	for _, todoItem := range **todoList {
		fmt.Println(todoItem)
	}
}

func addTodoItem(todoList **[]string) {
	var newTodoItem string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nInput new to-do item > ")
	newTodoItem, _ = reader.ReadString('\n')
	**todoList = append(**todoList, newTodoItem)
}

func listChoicer(todoList *[]string, running *bool) {
	var choice int
	fmt.Print("\nPlease select a course of action:\n\n1. Add to-do item.\n2. View to-do items.\n3. Quit.\n\n> ")
	fmt.Scan(&choice);

	switch choice {
	case 1:
		addTodoItem(&todoList)
	case 2:
		outputTodoList(&todoList)
	case 3:
		*running = false
	}
}

func main() {
	var todoList = []string {"- Clean my room", "- Do the dishes"}
	var running = true

	for running {
		listChoicer(&todoList, &running)
	}
}