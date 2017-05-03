package ecs

import "github.com/aws/aws-sdk-go/aws/session"

type Config struct {
	BloxURL string           `json:"bloxUrl"`
	Session *session.Session `json:"session"`
}
