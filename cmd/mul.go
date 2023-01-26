/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// mulCmd represents the mul command
var mulCmd = &cobra.Command{
	Use:   "mul",
	Short: "Multiplication Operation",
	Long:  `It multiplies any number of input numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		performMulOP(cmd, args)
	},
}

func performMulOP(cmd *cobra.Command, args []string) {
	fstatus, _ := cmd.Flags().GetBool("float")

	if fstatus {
		mulFloat(args)
	} else {
		mulInt(args)
	}
}

func mulInt(args []string) {
	var mul int
	mul = 1

	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)

		if err != nil {
			fmt.Println(err)
		}

		mul *= itemp
	}

	fmt.Printf("Multiplication of %s is %d", args, mul)
}

func mulFloat(args []string) {
	var mul float64
	mul = 1

	for _, fval := range args {
		ftemp, err := strconv.ParseFloat(fval, 64)

		if err != nil {
			fmt.Println(err)
		}

		mul *= ftemp
	}

	fmt.Printf("Multiplication of floating point numbers %s is %f", args, mul)
}

func init() {
	rootCmd.AddCommand(mulCmd)

	// Local Flag for float numbers
	mulCmd.Flags().BoolP("float", "f", false, "Floating point multiplication")
}
