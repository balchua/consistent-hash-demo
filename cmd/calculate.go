/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/balchua/consistent-demo/pkg/calculator"
	"github.com/balchua/consistent-demo/pkg/logging"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

var (
	port int32
)

// calculateCmd represents the calculate command
var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Calculate a number",
	Long:  `Calculate a number`,
	Run:   startCalculate,
}

func init() {
	rootCmd.AddCommand(calculateCmd)

	calculateCmd.Flags().Int32VarP(&port, "port", "p", 10000, "port for the calculate service")
}

func startCalculate(cmd *cobra.Command, args []string) {
	calc := calculator.NewCalculator()
	app := fiber.New()

	app.Post("/calculate/", calc.Calculate)
	logging.Info(fmt.Sprintf(":%d", port))
	app.Listen(fmt.Sprintf(":%d", port))
}
