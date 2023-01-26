/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Addtion Operation",
	Long:  `It Adds any number of input numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		performAddOP(cmd, args)
	},
}

func performAddOP(cmd *cobra.Command, args []string) {
	fstatus, _ := cmd.Flags().GetBool("float")

	if fstatus {
		addFloat(args)
	} else {
		addInt(args)
	}
}

func addInt(args []string) {
	var sum int

	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)

		if err != nil {
			fmt.Println(err)
		}

		sum += itemp
	}

	fmt.Printf("Addtion of %s is %d", args, sum)
}

func addFloat(args []string) {
	var sum float64

	for _, fval := range args {
		ftemp, err := strconv.ParseFloat(fval, 64)

		if err != nil {
			fmt.Println(err)
		}

		sum += ftemp
	}

	fmt.Printf("Addtion of floating point numbers %s is %f", args, sum)
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Local Flag for float numbers
	addCmd.Flags().BoolP("float", "f", false, "Floating point addtion")
}
