/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/Tak1za/tasker/access"
	"github.com/Tak1za/tasker/helper"
	"github.com/Tak1za/tasker/models"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get tasks",
	Long:  `Get the list of tasks present along with nitty gritty details about the tasks`,
	RunE: func(cmd *cobra.Command, rgs []string) error {
		id, _ := cmd.Flags().GetInt("id")
		if id != 0 {
			fmt.Println("ID provided: ", id)
			return nil
		} else {
			payload, err := access.GetTasks()
			if err != nil {
				return err
			}
			detailed, _ := cmd.Flags().GetBool("detailed")
			if detailed {
				results, err := helper.GetTasks(payload, true)
				if err != nil {
					return err
				}

				for _, d := range results.([]models.ToDoList) {
					d.String()
				}
				return nil
			} else {
				results, err := helper.GetTasks(payload, false)
				if err != nil {
					return err
				}

				for _, d := range results.([]string) {
					fmt.Println(d)
				}
				return nil
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolP("detailed", "d", false, "Get detailed information")
	getCmd.Flags().IntP("id", "i", 0, "Get task by ID")
}
