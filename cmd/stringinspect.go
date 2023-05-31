/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
)

// define onlyDigits var for -d flag
var onlyDigits bool

// stringinspectCmd represents the stringinspect command
var stringinspectCmd = &cobra.Command{
	Use:   "stringinspect",
	Short: "Inspect a string",
	Long:  `Inspect a string for learning Golang`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		i := args[0]

		//fmt.Println(i)

		res, kind := Inspect(i, onlyDigits)

		// Configure Emoji
		emojiValue := "Hedgehog"
		if onlyDigits == true {
			emojiValue = "Bat"
		}

		// Add variable to make output pluralß
		pluralS := "s"

		// If there only one character - no plural
		if res == 1 {
			pluralS = ""
			if emojiValue == "Hedgehog" {
				fmt.Printf("'%s' has %d %s%s %v\n", i, res, kind, pluralS, emoji.Hedgehog)
			} else if emojiValue == "Bat" {
				fmt.Printf("'%s' has %d %s%s %[5]s\n", i, res, kind, pluralS, emoji.Bat)
			}

		} else if res > 1 {
			// If there is more than one character - make plural
			if emojiValue == "Hedgehog" {
				fmt.Printf("'%s' has %d %s%s %[5]s\n", i, res, kind, pluralS, emoji.Hedgehog)
			} else if emojiValue == "Bat" {
				fmt.Printf("'%s' has %d %s%s %[5]s\n", i, res, kind, pluralS, emoji.Bat)
			}

		} else {
			// If there are no characters - change message
			fmt.Printf("%[2]v There is no %[1]s %[2]v\n", kind, emoji.CrossMark)
		}

	},
}

func init() {
	// Added flag
	stringinspectCmd.Flags().BoolVarP(&onlyDigits, "digits", "d", false, "Count only digits")
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

	// Input is false be default, so this function will always run
	if !digits {
		return len(input), "hedgehog"
	}

	// Is in use when the digit flag is added
	return inspectNumbers(input), "bat"
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
