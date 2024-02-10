/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dockit/internal/command/configure"
	dcontext "dockit/internal/command/docker/context"

	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure [path]",
	Short: "Configures access to a remote dockit instance",
	Long: `The 'configure' command sets up access to a remote docker server
by providing the dockit configuration file. It creates a new docker context
and sets it as active, enabling seamless interaction with the remote Dockit instance.`,
	Example:               "  dockit configure /path/to/gz.tar",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := configure.Run(args[0]); err != nil {
			return err
		}

		if err := dcontext.Create(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
