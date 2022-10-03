package main

import (
	"sync"

	"github.com/cmwylie19/kubefs/pkg/server"
	"github.com/spf13/cobra"
)

var (

	// WaitGroup is used to wait for the program to finish goroutines.
	wg sync.WaitGroup
	// cfgPath is the path to the EnvoyGateway configuration file.
	key  string
	cert string
	port string
	dir  string
)

// getServerCommand returns the server cobra command to be executed.
func getServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "serve",
		Aliases: []string{"s"},
		Short:   "Serve Media Controller",
		RunE: func(cmd *cobra.Command, args []string) error {
			s := &server.Server{}
			return s.Serve()
			//		return s.Serve(tlsKey, tlsCert, port)
		},
	}

	cmd.PersistentFlags().StringVarP(&dir, "dir", "d", "", "Directory from which to manage files")

	return cmd
}
