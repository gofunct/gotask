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
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

// s3DeleteCmd represents the s3Delete command
var s3DeleteCmd = &cobra.Command{
	Use:   "s3Delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) != 2 {
			exitErrorf("bucket name required\nUsage: %s bucket_name", os.Args[0])
		}

		bucket := os.Args[1]

		// Initialize a session in us-west-2 that the SDK will use to load
		// credentials from the shared credentials file ~/.aws/credentials.
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-west-2")},
		)

		// Create S3 service client
		svc := s3.New(sess)

		// Delete the S3 Bucket
		// It must be empty or else the call fails
		_, err = svc.DeleteBucket(&s3.DeleteBucketInput{
			Bucket: aws.String(bucket),
		})
		if err != nil {
			exitErrorf("Unable to delete bucket %q, %v", bucket, err)
		}

		// Wait until bucket is deleted before finishing
		fmt.Printf("Waiting for bucket %q to be deleted...\n", bucket)

		err = svc.WaitUntilBucketNotExists(&s3.HeadBucketInput{
			Bucket: aws.String(bucket),
		})
		if err != nil {
			exitErrorf("Error occurred while waiting for bucket to be deleted, %v", bucket)
		}

		fmt.Printf("Bucket %q successfully deleted\n", bucket)

	},
}

func init() {
	rootCmd.AddCommand(s3DeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// s3DeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// s3DeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
