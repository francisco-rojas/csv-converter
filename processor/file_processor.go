package processor

import (
	"fmt"
	"strconv"
	"strings"
)

// File iterates over a data collection and processes the data according to the mappings.
type File struct {
	Filename   string
	Rows       [][]string
	Mappings   map[interface{}]interface{}
	HasHeaders bool
	Headers    []string
}

// ProcessFile iterates over Rows and processes the data accordingly.
func (fp *File) ProcessFile() []map[string]interface{} {
	if fp.HasHeaders {
		fp.Headers = fp.Rows[0]
		fp.Rows = fp.Rows[1:]
	}

	var r []map[string]interface{}

	for rowNum, row := range fp.Rows {
		if fp.HasHeaders {
			rowNum = rowNum + 1
		}
		r = append(r, fp.processEntities(rowNum, row))
	}

	return r
}

func (fp *File) processEntities(rowNum int, row []string) map[string]interface{} {
	r := make(map[string]interface{})

	for e, m := range fp.Mappings {
		o := map[string]string{"filename": fp.Filename, "rowNum": strconv.Itoa(rowNum), "entity": e.(string), "row": strings.Join(row, ",")}
		p := Entity{
			Options:    o,
			Row:        row,
			Mappings:   m.(map[interface{}]interface{}),
			HasHeaders: fp.HasHeaders,
			Headers:    fp.Headers,
		}

		r[e.(string)] = p.ProcessEntity()
	}

	fmt.Println(r)
	return r
}
