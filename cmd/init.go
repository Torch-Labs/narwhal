/*
Copyright © 2021 Sharith Godamanna

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”),
to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CommandFlag struct {
	long, short, help, flag string
	required                bool
}

var commandFlags = []CommandFlag{
	{long: "name", short: "n", help: "New project name (required)", required: true},
	{long: "gitrepo", short: "g", help: "Git repo, in the format git@github.com:<org>/<repo>.git (required)", required: true},
	{long: "remotehost", short: "r", help: "Remote host, ex: rapsberrypi (required)", required: true},
	{long: "remoteuser", short: "u", help: "Remote user, ex: pi (required)", required: true},
	{long: "envpath", short: "e", help: "Path to .env file, default cwd .env (Will be copied using scp), If not specified nothing will be copied (optinal)", required: false},
}

const (
	name = iota
	gitrepo
	remotehost
	remoteuser
	envpath
)

// var name_flg string
// var gitrepo_flg string
// var remotehost_flg string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new narwhal project",
	Long: `
- Creates a config file with all the required attributs for a narwhal project.
- Initialized SSH remote server and saves git repo information.
- This command must be run first before a project can be deployed.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if commandFlags[name].flag != "" {
			fmt.Println("Name")
		}
		if commandFlags[gitrepo].flag != "" {
			fmt.Println("Git Repo")
		}
		if commandFlags[remotehost].flag != "" {
			fmt.Println("Remote Host")
		}
		if commandFlags[remoteuser].flag != "" {
			fmt.Println("Remote user")
		}
		if commandFlags[envpath].flag != "" {
			fmt.Println("Env path")
		}
	},
}

func init() {

	for _, item := range commandFlags {
		// name flag (Required) project name
		initCmd.Flags().StringVarP(&item.flag, item.long, item.short, "", item.help)
		if item.required {
			initCmd.MarkFlagRequired(item.long)
		}
		viper.BindPFlag(item.long, rootCmd.PersistentFlags().Lookup(item.long))
	}

	rootCmd.AddCommand(initCmd)

}
