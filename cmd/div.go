/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// divCmd represents the div command
var divCmd = &cobra.Command{
	Use:   "div",
	Short: "Division Operation",
	Long:  `It divides any number of input numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		performDivOP(cmd, args)
	},
}

func performDivOP(cmd *cobra.Command, args []string) {
	fstatus, _ := cmd.Flags().GetBool("float")

	if fstatus {
		divFloat(args)
	} else {
		divFloat(args)
	}
}

func divFloat(args []string) {
	var div float64
	div = 1

	for _, fval := range args {
		ftemp, err := strconv.ParseFloat(fval, 64)

		if err != nil {
			fmt.Println(err)
		}

		div = ftemp / div
	}

	fmt.Printf("Division of floating point numbers %s is %f", args, div)
}

func init() {
	rootCmd.AddCommand(divCmd)

	// Local Flag for float numbers
	divCmd.Flags().BoolP("float", "f", false, "Floating point division")
}
