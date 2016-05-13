// Copyright © 2016 mijime
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
	"os"

	"github.com/mijime/lipio/pipe"
	"github.com/spf13/cobra"
)

// pipeCmd represents the pipe command
var pipeCmd = &cobra.Command{
	Use:   "pipe",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Println(pipe.NotMatchError)
			os.Exit(1)
			return
		}

		p, pipeErr := pipe.NewPipe(pipe.Option{Scheme: args[0]})

		if pipeErr != nil {
			log.Println(pipeErr)
			os.Exit(1)
			return
		}

		_, execErr := p.Execute(os.Stdout, os.Stdin)

		if execErr != nil {
			log.Println(execErr)
			os.Exit(1)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(pipeCmd)
}
