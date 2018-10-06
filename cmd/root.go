package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var ip string
var token string
var verbose bool

var rootCmd = &cobra.Command{
	Use:   "hue",
	Short: "Golang CLI application to communicate with your Philips Hue light system",
	Long: `Golang CLI application to communicate with your Philips Hue light system

Check https://github.com/jordyv/hue for more information`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hue.yaml)")
	rootCmd.PersistentFlags().StringVar(&ip, "ip", "", "IP address of your Hue bridge")
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "Token of your Hue bridge")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose logging")

	viper.BindPFlag("ip", rootCmd.PersistentFlags().Lookup("ip"))
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".hue" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".hue")
	}

	viper.SetEnvPrefix("hue")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file:", viper.ConfigFileUsed())
	}

	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	log.Debugf("Loaded configuration: %s", viper.AllSettings())
}
