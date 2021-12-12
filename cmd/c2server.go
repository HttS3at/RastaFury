/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"RastaShellGenerator/c2server"
	"github.com/spf13/cobra"
)

var port string

// c2serverCmd represents the c2server command
var c2serverCmd = &cobra.Command{
	Use:   "c2server",
	Short: "Start a c2 server listener",
	Long:  `is a listener for handle the incoming conections and workd with encrypted commands`,
	Run: func(cmd *cobra.Command, args []string) {
		c2server.Server(port)
	},
}

func init() {
	rootCmd.AddCommand(c2serverCmd)
	c2serverCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to listen incoming conections")
}
