package cmd

import (
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var sceneID string

var lightsCmd = &cobra.Command{
	Use:   "lights",
	Short: "Actions for your lights",
	Long:  `Send actions to your lights. For example: 'hue lights on' to turn on your lights.`,
}

func init() {
	rootCmd.AddCommand(lightsCmd)

	lightsCmd.PersistentFlags().StringVar(&sceneID, "sceneID", "", "Scene ID to use for actions")
	viper.BindPFlag("sceneID", lightsCmd.PersistentFlags().Lookup("sceneID"))
}
