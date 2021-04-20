/*
	Package cmd
	Copyright © 2021 András Horváth horvath.andras216@gmail.com
*/
package cmd

import (
	"github.com/horvathandris/go-type-registry/parser"
	"github.com/spf13/cobra"
	"log"
	"os"
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
		if singleFile {
			if input == output {
				log.Fatalln("Input and output cannot be the same file.")
				return
			}
			parser.StartFile(input, output)
		} else {
			parser.StartDir(input, input+"/"+output)
		}
		log.Printf("Created type registry at %v .\n", output)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

var input string
var output string
var singleFile bool

func init() {
	cobra.OnInitialize()

	log.SetFlags(0)
	log.SetOutput(os.Stdout)
	log.SetPrefix("✨ ")

	rootCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "the input directory or .go file, from which the registry is created")
	_ = rootCmd.MarkPersistentFlagRequired("input")

	rootCmd.PersistentFlags().BoolVarP(&singleFile, "singlefile", "s", false, "true if you only want to parse a single file, false by default")

	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "the output .go file, where the registry is created")
	_ = rootCmd.MarkPersistentFlagRequired("output")
}
