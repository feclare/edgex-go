//
// Copyright (c) 2018 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package logging

import (
	"github.com/edgexfoundry/edgex-go/support/domain"
)

const (
	applicationName    = "support-logging"
	defaultPort        = 48061
	defaultPersistence = "PersistenceFile"
	defaultLogFilename = "support-logging.log"

	PersistenceMongo = "mongodb"
	PersistenceFile  = "file"
)

type Config struct {
	Port        int
	Persistence string

	// Used by PersistenceFile
	LogFilename string
}

type persistence interface {
	add(logEntry support_domain.LogEntry)
	remove(criteria matchCriteria) int
	find(criteria matchCriteria) []support_domain.LogEntry
}

type matchCriteria struct {
	OriginServices  []string
	MessageKeywords []string
	LogLevels       []string
	Labels          []string
	Keywords        []string
	Start           int64
	End             int64
	Limit           int
}

func GetDefaultConfig() Config {
	return Config{
		Port:        defaultPort,
		Persistence: defaultPersistence,
		LogFilename: defaultLogFilename,
	}
}

func (matchCriteria) match(le support_domain.LogEntry) bool {

	return true
}
