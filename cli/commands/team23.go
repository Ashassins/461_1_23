/*
Root command is created as team23. All other commands are built on
top of this command. Creation of new commands requires an init
function per command with rootCmd.AddCommand(<newCmd>)
*/

package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/19chonm/461_1_23/fileio"
	"github.com/19chonm/461_1_23/worker"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "team23",
	Short: "team23 - root command for app",
	Long:  "team23 is the root command to navigate through Team 23's CLI",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// test URL_FILE with: /Users/emile/461_1_23/test/urls_file.txt

// First function to be ran on main. Will check if second argument is either
// an absolute filepath, one of the recognized commands or neither. If neither,
// program will throw error. If argument is an absolute filepath, a direct call
// to functions are executed. No cobra command is created because name varies.
func Execute() {

	if len(os.Args) != 2 {
		fmt.Println(`CLI: Please use one of the recognized commands: 'build', 
		'install', 'test', or 'URL_FILE' where URL_FILE is an absolute path 
		to a file`)
	} else if filepath.IsAbs(os.Args[1]) {
		// Create channels for interthread communication
		url_ch := fileio.MakeUrlChannel()
		rating_ch := fileio.MakeRatingsChannel()

		go fileio.ReadFile(os.Args[1], url_ch)    // Start file reader
		go worker.StartWorkers(url_ch, rating_ch) // Start workers

		// Start output
		ratings := fileio.Sort_modules(rating_ch)
		fileio.Print_sorted_output(ratings)

	} else if os.Args[1] == "build" || os.Args[1] == "install" ||
		os.Args[1] == "test" {

		if err := rootCmd.Execute(); err != nil {
			fmt.Println("CLI: Error using CLI ", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("CLI: Not a recognized command")
		os.Exit(1)
	}

	os.Exit(0)
}
