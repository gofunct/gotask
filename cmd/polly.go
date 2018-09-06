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
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/spf13/cobra"
)

// pollyCmd represents the polly command
var pollyCmd = &cobra.Command{
	Use:   "polly",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(os.Args) != 2 {
			fmt.Println("You must supply an alarm name")
			os.Exit(1)
		}

		// The name of the text file to convert to MP3
		fileName := os.Args[1]

		// Open text file and get it's contents as a string
		contents, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println("Got error opening file " + fileName)
			fmt.Print(err.Error())
			os.Exit(1)
		}

		// Convert bytes to string
		s := string(contents[:])

		// Initialize a session that the SDK uses to load
		// credentials from the shared credentials file. (~/.aws/credentials).
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		// Create Polly client
		svc := polly.New(sess)

		// Output to MP3 using voice Joanna
		input := &polly.SynthesizeSpeechInput{OutputFormat: aws.String("mp3"), Text: aws.String(s), VoiceId: aws.String("Joanna")}

		output, err := svc.SynthesizeSpeech(input)
		if err != nil {
			fmt.Println("Got error calling SynthesizeSpeech:")
			fmt.Print(err.Error())
			os.Exit(1)
		}

		// Save as MP3
		names := strings.Split(fileName, ".")
		name := names[0]
		mp3File := name + ".mp3"

		outFile, err := os.Create(mp3File)
		if err != nil {
			fmt.Println("Got error creating " + mp3File + ":")
			fmt.Print(err.Error())
			os.Exit(1)
		}

		defer outFile.Close()
		_, err = io.Copy(outFile, output.AudioStream)
		if err != nil {
			fmt.Println("Got error saving MP3:")
			fmt.Print(err.Error())
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(pollyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pollyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pollyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
