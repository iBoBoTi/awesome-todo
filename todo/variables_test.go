package todo

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"testing"


	"github.com/stretchr/testify/assert"
)

func TestAddToDo(t *testing.T) {
	todoList := ToDoList{}
	todoList.Tasks = make([]ToDo, 0)

	tests := []struct {
		name      string
		taskInput string
	}{
		{"task1", "say a prayer"},
		{"task2", "do some exercise"},
		{"task3", "take tea"},
		{"task4", "meditate"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			AddToDo(test.taskInput)
			file, _ := ioutil.ReadFile("todos.json")
			_ = json.Unmarshal(file, &todoList)
			assert.Equal(t, todoList.Tasks[len(todoList.Tasks)-1].Title, test.taskInput)
			assert.Equal(t, todoList.Tasks[len(todoList.Tasks)-1].IsComplete, false)
		})
	}
}

func TestListToDo(t *testing.T) {

}

func TestCompleteToDo(t *testing.T) {
	todoList := ToDoList{}
	todoList.Tasks = make([]ToDo, 0)
	file, _ := ioutil.ReadFile("todos.json")
	_ = json.Unmarshal(file, &todoList)

	tests := []struct {
		name       string
		completeNo int
	}{
		{"complete1", rand.Intn(len(todoList.Tasks)-1)+1},
		{"complete2", rand.Intn(len(todoList.Tasks)-1)+1},
		{"complete3", rand.Intn(len(todoList.Tasks)-1)+1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			CompleteToDo(test.completeNo)
			file, _ := ioutil.ReadFile("todos.json")
			_ = json.Unmarshal(file, &todoList)
			assert.Equal(t, todoList.Tasks[test.completeNo-1].IsComplete, true)
		})
	}
}

func TestDeleteToDo(t *testing.T) {
	todoList := ToDoList{}
	todoList.Tasks = make([]ToDo, 0)
	file, _ := ioutil.ReadFile("todos.json")
	_ = json.Unmarshal(file, &todoList)

	tests := []struct {
		name     string
		deleteNo int
	}{
		{"delete1", rand.Intn(len(todoList.Tasks)-1)+1},
		{"delete2", rand.Intn(len(todoList.Tasks)-1)+1},
		{"delete3", rand.Intn(len(todoList.Tasks)-1)+1},
	}
	for _, test := range tests {

		initialLength := len(todoList.Tasks)
		
		t.Run(test.name, func(t *testing.T) {
			DeleteToDo(test.deleteNo)
			file, _ := ioutil.ReadFile("todos.json")
			_ = json.Unmarshal(file, &todoList)
			assert.Equal(t, len(todoList.Tasks), initialLength-1)
		})
	}
}
