/*
Copyright © 2019 Dustin Ratcliffe <dustin.k.ratcliffe@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/spkg/bom"
)

type Csv struct {
	Headers []string
	Records []Record
}

type Record struct {
	Headers []string
	Fields  map[string]string
}

func Parse(file *os.File) Csv {
	if file == nil {
		return Csv{}
	}

	reader := csv.NewReader(bom.NewReader(file))
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	records := make([]Record, 0)
	var headers []string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if headers == nil {
			headers = line
		} else {
			record := Record{Headers: headers, Fields: make(map[string]string)}
			for i := range headers {
				record.Fields[headers[i]] = line[i]
			}
			records = append(records, record)
		}
	}
	return Csv{Headers: headers, Records: records}
}
