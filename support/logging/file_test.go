//
// Copyright (c) 2018 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package logging

import (
	"os"
	"testing"

	"github.com/edgexfoundry/edgex-go/support/domain"
)

const (
	testFilename   string = "test.log"
	sampleService1 string = "tservice1"
	sampleService2 string = "tservice2"
	message1       string = "message1"
	message2       string = "message2"
)

func TestFileFind(t *testing.T) {
	var keywords1 = []string{"1"}
	var keywords2 = []string{"2"}
	var keywords12 = []string{"2", "1"}

	var labels1 = []string{"label1"}
	var labels2 = []string{"label2"}
	var labels12 = []string{"label2", "label1"}

	var tests = []struct {
		name     string
		criteria matchCriteria
		result   int
	}{
		{"empty", matchCriteria{}, 6},
		{"keywords1", matchCriteria{Keywords: keywords1}, 3},
		{"keywords2", matchCriteria{Keywords: keywords2}, 3},
		{"keywords12", matchCriteria{Keywords: keywords12}, 6},
		{"labels1", matchCriteria{Labels: labels1}, 5},
		{"labels2", matchCriteria{Labels: labels2}, 1},
		{"labels12", matchCriteria{Labels: labels12}, 6},
	}

	fl := fileLog{filename: testFilename}

	// Remove test log, the test needs an empty file
	os.Remove(testFilename)

	// Remove test log when test ends
	defer os.Remove(testFilename)

	le := support_domain.LogEntry{
		Level:         support_domain.TRACE,
		OriginService: sampleService1,
		Message:       message1,
		Labels:        labels1,
	}
	fl.add(le)
	le.Message = message2
	fl.add(le)
	le.Message = message1
	fl.add(le)
	le.Message = message2
	le.OriginService = sampleService2
	le.Message = message2
	fl.add(le)
	le.Message = message1
	fl.add(le)
	le.Message = message2
	le.Labels = labels2
	fl.add(le)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logs := fl.find(tt.criteria)
			if logs == nil {
				t.Errorf("Should not be nil")
			}
			if len(logs) != tt.result {
				t.Errorf("Should return %d log entries, returned %d",
					tt.result, len(logs))
			}
		})
	}
}
