package ecs

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	blox "github.com/segmentio/ecs-blox/blox/client"
)

type Config struct {
	BloxURL string           `json:"bloxUrl"`
	Session *session.Session `json:"session"`
}

func (config *Config) getEcsClient() *ecs.ECS {
	return ecs.New(config.Session)
}

func (config *Config) getBloxClient() *blox.BloxClusterStateService {
	t := blox.DefaultTransportConfig().WithHost(config.BloxURL)
	return blox.NewHTTPClientWithConfig(nil, t)
}
