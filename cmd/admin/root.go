package admin

import (
	"github.com/hkonitzer/ohmab/internal/pkg/common/config"
	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "admin",
	Short: "Manage various aspects of the application",
	Long:  `With admin you can manage the users (e.g. create/delete).`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}
func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yml", "config file name (default is config.yml)")
	RootCmd.PersistentFlags().StringVar(&cfgFile, "configpath", ".", "config file path (default is .)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	_, err := config.Get()
	cobra.CheckErr(err)
}
