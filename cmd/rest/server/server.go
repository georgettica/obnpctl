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
package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/georgettica/obnpctl/pkg/consts"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command.
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A command to run a server in the background",
	RunE:  run,
}

var serverCmdArgs = struct {
	port int
}{
	port: consts.Port,
}

func run(cmd *cobra.Command, args []string) error {
	http.HandleFunc("/", fooHandler)

	headerTimeout := time.Minute
	s := http.Server{
		Addr:              fmt.Sprintf(":%d", serverCmdArgs.port),
		Handler:           nil,
		ReadHeaderTimeout: headerTimeout,
	}

	if err := s.ListenAndServe(); err != nil {
		return fmt.Errorf("error serving on port %d: %w", serverCmdArgs.port, err)
	}

	return nil
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v: Hello\n", time.Now())
	fmt.Printf("%v: Hello\n", time.Now())
}

// GetCmd will return the current command.
func GetCmd() *cobra.Command {
	return serverCmd
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
