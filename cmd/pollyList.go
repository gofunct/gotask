// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/spf13/cobra"
)

// pollyListCmd represents the pollyList command
var pollyListCmd = &cobra.Command{
	Use:   "pollyList",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize a session that the SDK uses to load
		// credentials from the shared credentials file. (~/.aws/credentials).
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		// Create Polly client
		svc := polly.New(sess)

		resp, err := svc.ListLexicons(nil)
		if err != nil {
			fmt.Println("Got error calling ListLexicons:")
			fmt.Print(err.Error())
			os.Exit(1)
		}

		for _, l := range resp.Lexicons {
			fmt.Println(*l.Name)
			fmt.Println("  Alphabet: " + *l.Attributes.Alphabet)
			fmt.Println("  Language: " + *l.Attributes.LanguageCode)
			fmt.Println("")
		}
	},
}

func init() {
	rootCmd.AddCommand(pollyListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pollyListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pollyListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
