/*
Copyright © 2021 Sharith Godamanna
*/
package cmd

import (
	"fmt"

	"github.com/Torch-Labs/narwhal/config"
	"github.com/Torch-Labs/narwhal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CommandFlag struct {
	long, short, help, flag string
	required                bool
}

// all flags for this command
var commandFlags = []CommandFlag{
	{long: "conf", short: "c", help: "Config file (optional), If not specified .narwhal_config.yml will be used", required: false},
	{long: "envpath", short: "e", help: "Path to .env file, default cwd .env (Will be copied using scp), If not specified nothing will be copied (optional)", required: false},
}

// const index map for the flags
const (
	conf = iota
	envpath
)

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
		if commandFlags[conf].flag != "" {
			fmt.Println("Conf set")
		} else {
			confExists := utils.CheckConfigExists()
			if !confExists {
				utils.PrintError("No config file detected")
				return
			}
			config := config.Config{}
			err := config.SetFromBytes(utils.GetConfigBytes())
			if err != nil {
				utils.PrintError("Invalid config file detected")
				return
			}
		}
		if commandFlags[envpath].flag != "" {
			fmt.Println("Env set")
		} else {
			fmt.Println("Env not set")
		}
	},
}

func init() {

	for i := 0; i < len(commandFlags); i++ {
		item := &commandFlags[i]
		// initialize flags
		initCmd.Flags().StringVarP(&item.flag, item.long, item.short, "", item.help)
		if item.required {
			initCmd.MarkFlagRequired(item.long)
		}
		// BindPFlag binds a specific key to a pflag (as used by cobra).
		viper.BindPFlag(item.long, rootCmd.PersistentFlags().Lookup(item.long))
	}

	rootCmd.AddCommand(initCmd)

}
