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
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

// s3NewCmd represents the s3New command
var s3NewCmd = &cobra.Command{
	Use:   "s3New",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		bucket := "temp"
		key := "README.md"

		// load credentials from the shared credentials file ~/.aws/credentials.
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-east-1")},
		)

		// Create S3 service client
		svc := s3.New(sess)

		_, err = svc.CreateBucket(&s3.CreateBucketInput{
			Bucket: &bucket,
		}) //error
		if err != nil {
			log.Println("Failed to create bucket", err)
			return
		} //error
		if err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{Bucket: &bucket}); err != nil {
			log.Printf("Failed to wait for bucket to exist %s, %s\n", bucket, err)
			return
		}

		_, err = svc.PutObject(&s3.PutObjectInput{
			Body:   strings.NewReader("Hello World!"),
			Bucket: &bucket,
			Key:    &key,
		})
		if err != nil {
			log.Printf("Failed to upload data to %s/%s, %s\n", bucket, key, err)
			return
		}

		log.Printf("Successfully created bucket %s and uploaded data with key %s\n", bucket, key)

	},
}

func init() {
	rootCmd.AddCommand(s3NewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// s3NewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// s3NewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
