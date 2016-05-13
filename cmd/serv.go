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
	"fmt"
	"net/http"
	"os"

	"github.com/mijime/lipio/serv"
	"github.com/spf13/cobra"
)

var (
	path, addr string
)

// servCmd represents the serv command
var servCmd = &cobra.Command{
	Use:   "serv",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		h, handleErr := serv.NewHandler(serv.Option{})

		if handleErr != nil {
			fmt.Fprintln(os.Stderr, handleErr)
			return
		}

		http.Handle(path, h)
		servErr := http.ListenAndServe(addr, nil)

		if servErr != nil {
			fmt.Fprintln(os.Stderr, servErr)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(servCmd)

	servCmd.Flags().StringVarP(&addr, "addr", "a", ":8080", "Help message for toggle")
	servCmd.Flags().StringVarP(&path, "path", "p", "/", "Help message for toggle")
}
