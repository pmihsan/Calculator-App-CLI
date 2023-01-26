/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "my-calc",
	Short: "A Simple Calculator",
	Long: `my-calc is simple command line utility developed by cobra.
	
It performs 
				
	addtion
	subraction
	multiplication
	division
				
It could be modified to implement more mathematical operations`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("SIMPLE CALCULATOR\n\n")
		fmt.Printf("To Perform arithmetic operation\nUse any of the below format with arguments\n\n")

		fmt.Printf("For Integers\n\n\t")
		fmt.Printf("my-calc [op] [args]\n\n")

		fmt.Printf("For Floating Point Numbers\n\n\t")
		fmt.Printf("my-calc [op] [flags] [args]\n\n")

		fmt.Printf("ARGS: \n\n\t")
		fmt.Printf("They can be upto any length.\n\tThey can be either integer or float.\n\n")

		fmt.Printf("FLAGS: \n\n\t")
		fmt.Println("Use -f for floation point numbers")
		fmt.Printf("\tNo flags for decimal numbers\n\n")

		fmt.Printf("OPERATION(OP): \n\n")
		fmt.Println("\tadd => For Addtion")
		fmt.Println("\tsub => For Subtraction")
		fmt.Println("\tmul => For Multiplication")
		fmt.Println("\tdiv => For Division")
		fmt.Println("\trem => For Modulo")
		fmt.Printf("\teval => For Evaluating a string\n\n")

		fmt.Printf("EXAMPLES: \n\n")
		fmt.Println("Decimal")
		fmt.Printf("\t1)my-calc add 1 2 3 4 5\n\tAddtion of [1 2 3 4 5] is 15\n\n");
		fmt.Printf("\t2)my-calc add -f 1.5 3.2 4.3\n\tAddtion of floating point numbers [1.5 3.2 4.3] is 9.000000\n\n");

		fmt.Println("Evaluate")
		fmt.Printf("\t1)my-calc eval 1 + 3 * 5 - 2  + 10\n\tExpression [1 + 3 * 5 - 2 + 10] evaluates to 24.000000\n\n")

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.my-calc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
