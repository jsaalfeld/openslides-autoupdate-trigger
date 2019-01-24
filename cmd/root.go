package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"

	"../service"
)

var cfgFile string
var host string
var port int
var username string
var password string

// RootCmd is the base command
var RootCmd = &cobra.Command{
	Use:   "openslides-autoupdate-trigger",
	Short: "Script for triggering autoupdates on OpenSlides",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Output(1, fmt.Sprintf("The Trigger is working on instance %s:%d\n", viper.GetString("network.host"), viper.GetInt("network.port")))
		service.Start(viper.GetString("network.host"), viper.GetInt("network.port"), viper.GetString("authentication.username"), viper.GetString("authentication.password"), viper.GetInt("actions.secondsOfActivity"), viper.GetInt("actions.secondsOfInactivity"), viper.GetString("actions.actionLevel"))
	},
}

// Execute is called once by main.main()
// it handles the flags
func Execute() {
	log.Output(1, "Welcome to the OpenSlides Autoupdate Trigger")
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "c", "config file (default is ./run.toml)")
	RootCmd.Flags().StringVarP(&host, "host", "H", "", "Host address to listen to.")
	RootCmd.Flags().IntVarP(&port, "port", "p", -1, "Port to listen to.")
	RootCmd.Flags().StringVarP(&username, "username", "u", "", "Username to connect with OS")
	RootCmd.Flags().StringVarP(&password, "secret", "s", "", "Password to connect with OS")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	initDefaults()
	viper.SetConfigName("openslides-autoupdate-trigger")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Output(1, "No config file found")
	}

	initStringFlag("host", "network")
	initIntFlag("port", "network")
	initStringFlag("username", "authentication")
	initStringFlag("password", "authentication")
}

func initDefaults() {
	viper.SetDefault("network", map[string]interface{}{
		"host": "localhost",
		"port": 4200,
	})
	viper.SetDefault("authentication", map[string]interface{}{
		"username": "admin",
		"password": "admin",
	})
	viper.SetDefault("actions", map[string]interface{}{
		"secondsOfActivity":   60,
		"secondsOfInactivity": 300,
		"actionLevel":         "medium",
	})
}

func initStringFlag(key, group string) {
	flag := RootCmd.Flags().Lookup(key)
	if flag != nil && flag.Changed {
		viper.Set(group+"."+key, flag.Value)
	}
}

func initIntFlag(key, group string) {
	flag := RootCmd.Flags().Lookup(key)
	if flag != nil && flag.Changed {
		value, err := strconv.ParseInt(flag.Value.String(), 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Invalid input for %s: %s", key, flag.Value.String()))
		}
		viper.Set(group+"."+key, value)
	}
}
