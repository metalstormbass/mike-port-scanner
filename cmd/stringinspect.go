/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// stringinspectCmd represents the stringinspect command
var stringinspectCmd = &cobra.Command{
	Use:   "stringinspect",
	Short: "Inspect a string",
	Long:  `Inspect a string for learning Golang`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		i := args[0]

		//fmt.Println(i)

		res, kind := Inspect(i, false)

		//fmt.Println(res)
		//fmt.Println(kind)

		// Add variable to make output plural
		pluralS := "s"

		// If there only one character - no plural
		if res == 1 {
			pluralS = ""
			fmt.Printf("'%s' has %d %s%s.\n", i, res, kind, pluralS)
		} else if res > 1 {
			// If there is more than one character - make plural
			fmt.Printf("'%s' has %d %s%s.\n", i, res, kind, pluralS)
		} else {
			// If there are no characters - change message
			fmt.Printf("There is %s.\n", kind)
		}

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

	// Input is false be default, so this function will always run
	if !digits {
		// Added Error Handling
		if len(input) == 0 {
			return len(input), "no hedgehog"
		}
	}

	// Returns length of input and aribtrary string
	return len(input), "hedgehog"
	//inspectNumbers(input)

	//return inspectNumbers(input), "digit"
}

func inspectNumbers(input string) (count int) {
	println("This is the input: ", input)
	count = 123451234
	//for _, c := range input {
	//	_, err := strconv.Atoi(string(c))
	//	if err == nil {
	//		count++
	//	}
	//}
	return count
}
