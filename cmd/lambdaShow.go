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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/spf13/cobra"
)

// lambdaShowCmd represents the lambdaShow command
var lambdaShowCmd = &cobra.Command{
	Use:   "lambdaShow",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize a session
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		// Create Lambda service client
		svc := lambda.New(sess, &aws.Config{Region: aws.String("us-west-2")})

		result, err := svc.ListFunctions(nil)

		if err != nil {
			fmt.Println("Cannot list functions")
			os.Exit(0)
		}

		fmt.Println("Functions:")

		for _, f := range result.Functions {
			fmt.Println("Name:        " + aws.StringValue(f.FunctionName))
			fmt.Println("Description: " + aws.StringValue(f.Description))
			fmt.Println("")
		}
	},
}

func init() {
	rootCmd.AddCommand(lambdaShowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lambdaShowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lambdaShowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
