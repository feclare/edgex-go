//
// Copyright (c) 2018 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package logging

import (
	"bufio"
	"encoding/json"
	"io"
	"os"

	"github.com/edgexfoundry/edgex-go/support/domain"
)

type fileLog struct {
	filename string
	out      io.Writer
}

func (fl *fileLog) add(le support_domain.LogEntry) {
	if fl.out == nil {
		var err error
		fl.out, err = os.OpenFile(fl.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			// Could not open file
			fl.out = nil
			return
		}
	}

	res, err := json.Marshal(le)
	if err != nil {
		return
	}
	fl.out.Write(res)
	fl.out.Write([]byte("\n"))
	//fmt.Println("file: ", le)
}

func (fl *fileLog) remove(criteria matchCriteria) int {
	return 0
}

func (fl *fileLog) find(criteria matchCriteria) []support_domain.LogEntry {
	var logs []support_domain.LogEntry
	f, err := os.Open(fl.filename)
	if err != nil {
		//fmt.Println("Error opening log file: ", fl.filename, err)
		return nil
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var le support_domain.LogEntry

		line := scanner.Bytes()
		err := json.Unmarshal(line, &le)
		if err == nil {
			if criteria.match(le) {
				logs = append(logs, le)
			}
		}
	}
	return logs
}
