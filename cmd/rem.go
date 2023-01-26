/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

// remCmd represents the rem command
var remCmd = &cobra.Command{
	Use:   "rem",
	Short: "Modulo Operation",
	Long:  `It provides remainder for input numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		performRemOP(cmd, args)
	},
}

func performRemOP(cmd *cobra.Command, args []string) {
	fstatus, _ := cmd.Flags().GetBool("float")

	if fstatus {
		remFloat(args)
	} else {
		remFloat(args)
	}
}

func remFloat(args []string) {
	var rem float64

	if len(args) > 2 || len(args) <= 1 {
		fmt.Println("Invalid number of arguments")
		return 
	} else {
		fval1, err1 := strconv.ParseFloat(args[0], 64)
		fval2, err2 := strconv.ParseFloat(args[1], 64)

		if err1 != nil && err2 != nil {
			fmt.Println("Invalid input")
			return 
		} else {
			rem = math.Mod(fval1, fval2)
		}
	}

	fmt.Printf("Modulo of floating point numbers %s is %f", args, rem)
}

func init() {
	rootCmd.AddCommand(remCmd)

	// Local Flag for float numbers
	remCmd.Flags().BoolP("float", "f", false, "Floating point modulo")
}
