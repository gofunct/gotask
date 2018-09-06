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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/spf13/cobra"
)

// emailStatsCmd represents the emailStats command
var emailStatsCmd = &cobra.Command{
	Use:   "emailStats",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize a session that the SDK uses to load
		// credentials from the shared credentials file ~/.aws/credentials
		// and configuration from the shared configuration file ~/.aws/config.
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		// Create an SES session.
		svc := ses.New(sess)

		// Attempt to send the email.
		result, err := svc.GetSendStatistics(nil)

		// Display any error message
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		dps := result.SendDataPoints

		fmt.Println("Got", len(dps), "datapoints")
		fmt.Println("")

		for _, dp := range dps {
			fmt.Println("Timestamp: ", dp.Timestamp)
			fmt.Println("Attempts:  ", aws.Int64Value(dp.DeliveryAttempts))
			fmt.Println("Bounces:   ", aws.Int64Value(dp.Bounces))
			fmt.Println("Complaints:", aws.Int64Value(dp.Complaints))
			fmt.Println("Rejects:   ", aws.Int64Value(dp.Rejects))
			fmt.Println("")
		}
	},
}

func init() {
	rootCmd.AddCommand(emailStatsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// emailStatsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// emailStatsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
