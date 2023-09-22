/*
Copyright © 2023 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cobra",
	Short: "A sample Cobra project",
	Long: `A sample Cobra project for
	Mastering Go, 4th edition.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringP("directory", "d", "/tmp", "Path")
	rootCmd.PersistentFlags().Uint("depth", 2, "Depth of search")
	viper.BindPFlag("directory", rootCmd.PersistentFlags().Lookup("directory"))
	viper.BindPFlag("depth", rootCmd.PersistentFlags().Lookup("depth"))
}
