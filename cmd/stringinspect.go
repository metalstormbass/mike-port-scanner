/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// stringinspectCmd represents the stringinspect command
var stringinspectCmd = &cobra.Command{
	Use:   "stringinspect",
	Short: "Inspect a string",
	Long:  `Inspect a string for learning Golang`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stringinspect called")
		MainHandle()
	},
}

func init() {
	rootCmd.AddCommand(stringinspectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stringinspectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stringinspectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// ADDING SOME STUFF TO PLAY WITH

func MainHandle() {
	fmt.Println("stringinspect called - testing flow to MainHandle function")

}

func Reverse(input string) (result string) {
	for _, c := range input {
		result = string(c) + result
	}
	return result
}

func Inspect(input string, digits bool) (count int, kind string) {
	if !digits {
		return len(input), "char"
	}
	return inspectNumbers(input), "digit"
}

func inspectNumbers(input string) (count int) {
	for _, c := range input {
		_, err := strconv.Atoi(string(c))
		if err == nil {
			count++
		}
	}
	return count
}
