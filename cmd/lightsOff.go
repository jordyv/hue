package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"hue/api"

	"github.com/spf13/cobra"
)

var lightsOffCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn off the lights",
	Run: func(cmd *cobra.Command, args []string) {
		api, err := api.New(viper.GetString("ip"), viper.GetString("token"))
		if err != nil {
			log.Error(err.Error())
			return
		}
		scenes := api.Scenes()
		scenes.TurnOff(viper.GetString("sceneID"))
	},
}

func init() {
	lightsCmd.AddCommand(lightsOffCmd)
}
