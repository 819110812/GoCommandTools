/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	auto_selector "GoCommandTools/pkg/auto-selector"
	"github.com/spf13/cobra"
	"os"
)

// autoSelectorCmd represents the autoSelector command
var autoSelectorCmd = &cobra.Command{
	Use:     "autoSelector",
	Aliases: []string{"as"},
	Short:   "randomly select string from a list",
	Long:    `randomly select string from a list`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Println("please input a list of string")
			os.Exit(1)
		}
		res, err := auto_selector.AutoSelect(args)
		if err != nil {
			cmd.Println("got error: ", err.Error())
			os.Exit(1)
		}
		cmd.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(autoSelectorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// autoSelectorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// autoSelectorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
