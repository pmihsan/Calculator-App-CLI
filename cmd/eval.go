/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/spf13/cobra"
)

// evalCmd represents the eval command
var evalCmd = &cobra.Command{
	Use:   "eval",
	Short: "Evaluate a String of operations",
	Long: `Performs an evaluation of operations provided in the given input string`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetBool("float")
		if fstatus {
			exprEval(cmd, args)
		} else {
			exprEval(cmd, args)
		}
	},
}

func exprEval(cmd *cobra.Command, args []string){


	if len(args) <= 2 || len(args[0]) <= 0 {
		fmt.Println("Not enough arguments")
		return 
	} 

	src := ""

	for _, ival := range args {
		src += ival
	}

	expression, err1 := govaluate.NewEvaluableExpression(src);

	if err1 != nil {
		fmt.Println("Invalid Input")
		return
	} else {
		params := make(map[string]interface{}, 8)
		params["x"] = 8 

		result, err2 := expression.Evaluate(params)

		if err2 != nil {
			fmt.Println("Invalid Operation")
			return
		} else {
			fmt.Printf("Expression %s evaluates to %f",args, result)
		}
	}
}

func init() {
	rootCmd.AddCommand(evalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// evalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// evalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// For Local Flags
	evalCmd.Flags().BoolP("float", "f", false, "Floating point evaluation")
}
