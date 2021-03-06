/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/iBoBoTi/awesome-todo/todo"

	"github.com/spf13/cobra"
)

// AddCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a new task to the todo list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			// read from stdin and add to todo list
			todo.AddToDo(todo.ReadFromStdin())
			return
		}
		todo.AddToDo(args[0])
	},
}

func init() {
	rootCmd.AddCommand(AddCmd)
}
