/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Subtraction Operation",
	Long:  `It subtracts any number of input numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		performSubOP(cmd, args)
	},
}

func performSubOP(cmd *cobra.Command, args []string) {
	fstatus, _ := cmd.Flags().GetBool("float")

	if fstatus {
		subFloat(args)
	} else {
		subInt(args)
	}
}

func subInt(args []string) {
	var sub int
	temp, numErr := strconv.Atoi(args[0])

	if numErr != nil {
		fmt.Println("Invalid Input")
		return
	} else {
		sub = temp * 2
	}

	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)

		if err != nil {
			fmt.Println(err)
		}

		sub -= itemp
	}

	fmt.Printf("Subtraction of %s is %d", args, sub)
}

func subFloat(args []string) {
	var sub float64
	temp, numErr := strconv.ParseFloat(args[0], 64)

	if numErr != nil {
		fmt.Println("Invalid Input")
		return
	} else {
		sub = temp * 2
	}

	for _, fval := range args {
		ftemp, err := strconv.ParseFloat(fval, 64)

		if err != nil {
			fmt.Println(err)
		}

		sub -= ftemp
	}

	fmt.Printf("Subtraction of floating point numbers %s is %f", args, sub)
}

func init() {
	rootCmd.AddCommand(subCmd)

	// Local Flag for float numbers
	subCmd.Flags().BoolP("float", "f", false, "Floating point subtraction")
}
