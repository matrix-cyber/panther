package csvstream

/**
 * Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import "encoding/csv"

type StreamingCSVReader struct {
	CVSReader *csv.Reader
	logLine   string
}

func NewStreamingCSVReader() (scr *StreamingCSVReader) {
	scr = &StreamingCSVReader{}
	reader := csv.NewReader(scr)
	reader.FieldsPerRecord = -1 // variable fields allowed
	reader.ReuseRecord = true   // for performance
	scr.CVSReader = reader
	return scr
}

// for io.Reader signature
func (scr *StreamingCSVReader) Read(b []byte) (n int, err error) {
	n = copy(b, scr.logLine)
	if n < len(scr.logLine) { // partial copy
		scr.logLine = scr.logLine[n:] // the rest for next call
	} else { // full copy
		scr.logLine = ""
		if n < len(b) { // terminate line if there is room
			b[n] = '\n'
			n++
		}
	}
	return n, nil
}

func (scr *StreamingCSVReader) Parse(log string) ([]string, error) {
	scr.logLine = log
	return scr.CVSReader.Read()
}
