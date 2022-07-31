/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/balchua/consistent-demo/pkg/logging"
	"github.com/balchua/consistent-demo/pkg/simple"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

var (
	port int32

	// simpleCmd represents the simple command
	simpleCmd = &cobra.Command{
		Use:   "simple",
		Short: "A simple service listening to a port",
		Long:  `Do nothing stuff`,
		Run:   startService,
	}
)

func init() {
	rootCmd.AddCommand(simpleCmd)
	simpleCmd.Flags().Int32VarP(&port, "port", "p", 10000, "port for the calculate service")
}

func startService(cmd *cobra.Command, args []string) {
	svc := simple.NewService()
	app := fiber.New()

	app.Post("/do/", svc.Do)
	logging.Info(fmt.Sprintf(":%d", port))
	app.Listen(fmt.Sprintf(":%d", port))
}
