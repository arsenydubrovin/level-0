package config

import (
	"fmt"
	"strconv"
)

const (
	natsHostEnvName      = "NATS_HOST"
	natsPortEnvName      = "NATS_PORT"
	stanClusterIdEnvName = "STAN_CLUSTER_ID"
	stanSubjectEnvName   = "STAN_SUBJECT"
)

type StanConfig interface {
	NatsURL() string
	StanClusterID() string
	StanSubject() string
}

type stanConfig struct {
	natsHost  string
	natsPort  int
	clusterId string
	subject   string
}

func NewStanConfig() (cfg StanConfig, err error) {
	natsHost, err := getEnvVariable(natsHostEnvName)
	if err != nil {
		return nil, err
	}

	natsPortStr, err := getEnvVariable(natsPortEnvName)
	if err != nil {
		return nil, err
	}

	natsPort, err := strconv.Atoi(natsPortStr)
	if err != nil {
		return nil, err
	}

	clusterId, err := getEnvVariable(stanClusterIdEnvName)
	if err != nil {
		return nil, err
	}

	subject, err := getEnvVariable(stanSubjectEnvName)
	if err != nil {
		return nil, err
	}

	return &stanConfig{
		natsHost:  natsHost,
		natsPort:  natsPort,
		clusterId: clusterId,
		subject:   subject,
	}, nil
}

func (cfg *stanConfig) NatsURL() string {
	return fmt.Sprintf("nats://%s:%d", cfg.natsHost, cfg.natsPort)
}

func (cfg *stanConfig) StanClusterID() string {
	return cfg.clusterId
}

func (cfg *stanConfig) StanSubject() string {
	return cfg.subject
}
