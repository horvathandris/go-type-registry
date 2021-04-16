/*
	Package cmd
	Copyright © 2021 András Horváth horvath.andras216@gmail.com
*/
package cmd

import (
	"github.com/horvathandris/go-type-registry/parser"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-type-registry",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		parser.Start(inFileName, outFileName)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

var inFileName string
var outFileName string

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().StringVarP(&inFileName, "input", "i", "", "the input .go file, from which the registry is created")
	_ = rootCmd.MarkPersistentFlagRequired("input")

	rootCmd.PersistentFlags().StringVarP(&outFileName, "output", "o", "", "the output .go file, where the registry is created")
	_ = rootCmd.MarkPersistentFlagRequired("output")
}
