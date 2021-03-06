
package csv

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/DustyRat/post-it/pkg/client"

	"github.com/spkg/bom"
)

// Csv ...
type Csv struct {
	Headers []string
	Records []Record
}

// Record ...
type Record struct {
	Headers []string
	Body    []byte
	Fields  map[string]string
}

// Request ...
func (r *Record) Request(method, rawurl string) (*client.Request, error) {
	return client.NewRequest(method, rawurl, http.Header{}, bytes.NewBuffer(r.Body), r.Fields)
}

// Parse ...
func Parse(file *os.File, body string) Csv {
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
			if b, ok := record.Fields[body]; ok {
				record.Body = []byte(b)
			}
			records = append(records, record)
		}
	}
	return Csv{Headers: headers, Records: records}
}
