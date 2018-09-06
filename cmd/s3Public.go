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
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

// s3PublicCmd represents the s3Public command
var s3PublicCmd = &cobra.Command{
	Use:   "s3Public",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(os.Args) < 2 {
			exitErrorf("Bucket name required.\nUsage: go run", os.Args[0], "BUCKET")
		}

		bucket := os.Args[1]

		// private | public-read | public-read-write | authenticated-read
		// See https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL for details
		acl := "public-read"

		// Initialize a session that the SDK uses to load
		// credentials from the shared credentials file ~/.aws/credentials
		// and region from the shared configuration file ~/.aws/config.
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		// Create S3 service client
		svc := s3.New(sess)

		params := &s3.PutBucketAclInput{
			Bucket: &bucket,
			ACL:    &acl,
		}

		// Set bucket ACL
		_, err := svc.PutBucketAcl(params)
		if err != nil {
			exitErrorf(err.Error())
		}

		fmt.Println("Bucket " + bucket + " is now public")

	},
}

func init() {
	rootCmd.AddCommand(s3PublicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// s3PublicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// s3PublicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
