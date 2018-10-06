package api

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Scenes struct {
	Api HueApi
}

type GroupAction struct {
	on bool
}

func (api *HueApi) Scenes() Scenes {
	return Scenes{Api: *api}
}

func (Scenes) ValidateSceneID(sceneID string) error {
	if sceneID == "" {
		return errors.New("invalid scene ID")
	}
	return nil
}

func (scenes *Scenes) GetSceneUrl(sceneID string) string {
	return fmt.Sprintf("http://%s/api/%s/groups/%s/action", scenes.Api.IP, scenes.Api.Token, sceneID)
}

func (scenes *Scenes) TurnOn(sceneID string) {
	if err := scenes.ValidateSceneID(sceneID); err != nil {
		log.Error(err.Error())
		return
	}
	log.Infof("Turn scene '%s' on\n", sceneID)
	_, err := scenes.Api.HttpPut(scenes.GetSceneUrl(sceneID), map[string]interface{}{"on": true})
	if err != nil {
		log.Error(err.Error())
	}
}

func (scenes *Scenes) TurnOff(sceneID string) {
	if err := scenes.ValidateSceneID(sceneID); err != nil {
		log.Error(err.Error())
		return
	}
	log.Infof("Turn scene '%s' off\n", sceneID)
	_, err := scenes.Api.HttpPut(scenes.GetSceneUrl(sceneID), map[string]interface{}{"on": false})
	if err != nil {
		log.Error(err.Error())
	}
}
