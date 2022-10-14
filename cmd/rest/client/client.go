/*
Copyright © 2022 Ron Green

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package client

import (
	"fmt"
	"net/http"

	"github.com/georgettica/obnpctl/pkg/consts"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command.
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A command to send a REST command",
	RunE:  run,
}

var clientCmdArgs = struct {
	url string
}{
	url: fmt.Sprintf("http://localhost:%d", consts.Port),
}

func run(cmd *cobra.Command, args []string) error {
	_, err := http.Get(clientCmdArgs.url)
	if err != nil {
		return fmt.Errorf("could not run the GET command to '%s': %w", clientCmdArgs.url, err)
	}

	return nil
}

// GetCmd will return the current command.
func GetCmd() *cobra.Command {
	return clientCmd
}

func init() {
	clientCmd.Flags().StringVarP(&clientCmdArgs.url,
		"url",
		"u",
		clientCmdArgs.url,
		"the url to which the command will send the GET request to")
}
