/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/balchua/consistent-demo/pkg/balancer"
	"github.com/balchua/consistent-demo/pkg/config"
	"github.com/balchua/consistent-demo/pkg/logging"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	cfgFile       string
	clusterConfig *config.ClusterConfig
)

// balancerCmd represents the balancer command
var balancerCmd = &cobra.Command{
	Use:   "balancer",
	Short: "consistent hash balancer",
	Long:  `consistent hash balancer`,
	Run:   startBalancer,
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(balancerCmd)
	balancerCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.consistent-demo.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".consistent-demo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".consistent-demo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		zap.S().Debugf("Using config file: %s", viper.ConfigFileUsed())
	}

	clusterConfig = &config.ClusterConfig{}

	if err := viper.Unmarshal(clusterConfig); err != nil {
		zap.S().Fatalf("%v", err)
	}

}

func startBalancer(cmd *cobra.Command, args []string) {
	b := balancer.NewBalancer(*clusterConfig)
	router := balancer.NewServer(b)
	app := fiber.New()

	app.Get("/calculate/:key", router.Pick)
	app.Post("/node", router.AddNode)
	app.Delete("/node", router.RemoveNode)

	if err := app.Listen(":3000"); err != nil {
		logging.Errorf("%v", err)
	}
}
