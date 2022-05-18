package todo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ToDo struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	IsComplete bool   `json:"is_complete"`
}

type ToDoList struct {
	Tasks []ToDo
}

func NewToDo(task string, complete bool) *ToDo {
	return &ToDo{
		Title:      task,
		IsComplete: complete,
	}
}

func AddToDo(task string) {
	task = strings.TrimSpace(strings.ToLower(task))
	file, err := ioutil.ReadFile("todos.json")
	if err != nil {
		log.Println("I got here")
		log.Fatal(err)
	}

	// Unmarshal read data to a type of ToDoList
	todoList := ToDoList{}
	todoList.Tasks = make([]ToDo, 0)
	if string(file) != "" {
		err = json.Unmarshal(file, &todoList)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Set up the new todo struct and add it to the unmarshalled list
	newTask := NewToDo(task, false)
	switch len(todoList.Tasks) == 0 {
	case true:
		newTask.ID = 1
	default:
		newTask.ID = todoList.Tasks[len(todoList.Tasks)-1].ID + 1
	}
	todoList.Tasks = append(todoList.Tasks, *newTask)

	// Marshal the updated list back to JSON
	jsonData, err := json.MarshalIndent(&todoList, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	// Write the updated list to the file
	err = ioutil.WriteFile("todos.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func CompleteToDo(todoNum int) {
	file, err := ioutil.ReadFile("todos.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal read data to a type of ToDoList
	todoList := ToDoList{}
	todoList.Tasks = make([]ToDo, 0)
	if string(file) != "" {
		err = json.Unmarshal(file, &todoList)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Find the task with the given ID and set it to complete
	for i, _ := range todoList.Tasks {
		if todoNum == i+1 {
			todoList.Tasks[i].IsComplete = true
		}
	}

	// Marshal the updated list back to JSON
	jsonData, err := json.MarshalIndent(&todoList, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	// Write the updated list to the file
	err = ioutil.WriteFile("todos.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func ListToDo() {
	file, err := ioutil.ReadFile("todos.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal read data to a type of ToDoList
	todoList := ToDoList{}
	todoList.Tasks = make([]ToDo, 0)
	if string(file) != "" {
		err = json.Unmarshal(file, &todoList)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Print the list of tasks
	for i, task := range todoList.Tasks {
		switch task.IsComplete {
		case true:
			fmt.Printf("(%v*) %s\n", i+1, task.Title)
		default:
			fmt.Printf("(%v) %s\n", i+1, task.Title)
		}
	}
}

func DeleteToDo(todoNum int) {
	file, err := ioutil.ReadFile("todos.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal read data to a type of ToDoList
	todoList := ToDoList{}
	todoList.Tasks = make([]ToDo, 0)
	if string(file) != "" {
		err = json.Unmarshal(file, &todoList)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Find the task with the given ID and delete it
	for i, _ := range todoList.Tasks {
		if todoNum == i+1 {
			todoList.Tasks = append(todoList.Tasks[:i], todoList.Tasks[i+1:]...)
		}
	}

	// Marshal the updated list back to JSON
	jsonData, err := json.MarshalIndent(&todoList, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	// Write the updated list to the file
	err = ioutil.WriteFile("todos.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFromStdin() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
