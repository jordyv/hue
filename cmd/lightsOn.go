package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"hue/api"
)

var lightsOnCmd = &cobra.Command{
	Use:   "on",
	Short: "Turn on the lights",
	Run: func(cmd *cobra.Command, args []string) {
		api, err := api.New(viper.GetString("ip"), viper.GetString("token"))
		if err != nil {
			log.Error(err.Error())
			return
		}
		scenes := api.Scenes()
		scenes.TurnOn(viper.GetString("sceneID"))
	},
}

func init() {
	lightsCmd.AddCommand(lightsOnCmd)
}
