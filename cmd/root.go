/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"xiliangzi_pro/cmd/apiServer"
	"xiliangzi_pro/cmd/websocket"

	"github.com/spf13/cobra"
)

var commands = []*cobra.Command{
	websocket.Cmd,
	apiServer.Cmd,
}

func Run() {
	root := cobra.Command{Use: "apiServer"}
	root.AddCommand(commands...)
	root.Execute()
}
