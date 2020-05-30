/*
Copyright © 2020 Varun Gupta varungupta2015135@gmail.com

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
	"errors"

	"github.com/Tak1za/tasker/access"
	"github.com/Tak1za/tasker/models"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task and start marking its progress`,
	RunE: func(cmd *cobra.Command, args []string) error {
		t, _ := cmd.Flags().GetString("task")
		if t != "" {
			var task models.ToDoListDB
			task.Task = t
			s, _ := cmd.Flags().GetBool("status")
			if s {
				task.Status = s
			}
			_, err := access.AddTask(task)
			if err != nil {
				return err
			}
		} else {
			return errors.New("A task message needs to be added")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("status", "s", false, "Set status")
	addCmd.Flags().StringP("task", "t", "", "Add task")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
