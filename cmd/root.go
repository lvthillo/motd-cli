/*
Copyright © 2019 NAME HERE lorenz.vanthillo@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

//variable used as flag later on
var name string
var greeting string
var preview bool
var prompt bool
var debug bool = false

// rootCmd represents the base command when called without any subcommands
// pointer to a struct
var rootCmd = &cobra.Command{
	Use:   "motd",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// Show usage if flag combination isn't valid (exit in this scenario)
		// If no arguments passed, show usage
		if prompt == false && (name == "" || greeting == "") {
			cmd.Usage()
			os.Exit(1)
		}

		// Optionally print flags and exit if DEBUG is set
		if debug {
			fmt.Println("Name:", name)
			fmt.Println("Greeting:", greeting)
			fmt.Println("Prompt:", prompt)
			fmt.Println("Preview:", preview)

			os.Exit(0)
		}

		// Conditionally read from stdin
		if prompt { //if prompt is true
			name, greeting = renderPrompt()
		}

		// Generate message
		message := buildMessage(name, greeting)

		// Either preview message or write to file
		if preview {
			fmt.Println(message)
		} else {
			// Open file to write and create if it does not exist
			f, err := os.OpenFile("./file.txt", os.O_WRONLY|os.O_CREATE, 0644)

			if err != nil {
				fmt.Println("Error: Unable to open to ./file.txt")
				fmt.Println(err)
				os.Exit(1)
			}

			defer f.Close()

			// Empty file.txt
			err = os.Truncate("./file.txt", 0)

			if err != nil {
				fmt.Println("Error: Failed to truncate ./file.txt")
				fmt.Println(err)
				os.Exit(1)
			}

			// Write message to file.txt
			_, err = f.Write([]byte(message))

			if err != nil {
				fmt.Println("Error: Failed to write to ./file.txt")
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

func buildMessage(name, greeting string) string {
	return fmt.Sprintf("%s, %s", greeting, name)
}

func renderPrompt() (name, greeting string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Your Greeting: ")
	greeting, _ = reader.ReadString('\n')
	greeting = strings.TrimSpace(greeting)

	fmt.Print("Your Name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	return
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	//; multiple actions on same line
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Executed when file is loaded (run before Execute is called)
func init() {
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "name to use within the message")
	rootCmd.Flags().StringVarP(&greeting, "greeting", "g", "", "phrase to use within the greeting")
	rootCmd.Flags().BoolVarP(&preview, "preview", "v", false, "use preview to output message without writing to ./file.txt")
	rootCmd.Flags().BoolVarP(&prompt, "prompt", "p", false, "use prompt to input name and message")

	if os.Getenv("DEBUG") != "" {
		debug = true
	}
}
