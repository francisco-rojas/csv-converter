package processors

import (
	"strconv"
	"strings"
)

// FileProcessor iterates over a data collection and processes the data according to the mappings.
type FileProcessor struct {
	Filename   string
	Rows       [][]string
	Mappings   map[interface{}]interface{}
	HasHeaders bool
	Headers    []string
}

// ProcessFile iterates over Rows and processes the data accordingly.
func (fp *FileProcessor) ProcessFile() []map[string]interface{} {
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

func (fp *FileProcessor) processEntities(rowNum int, row []string) map[string]interface{} {
	r := make(map[string]interface{})

	for e, m := range fp.Mappings {
		o := map[string]string{"filename": fp.Filename, "rowNum": strconv.Itoa(rowNum), "entity": e.(string), "row": strings.Join(row, ",")}
		p := EntityProcessor{
			Options:    o,
			Row:        row,
			Mappings:   m.(map[interface{}]interface{}),
			HasHeaders: fp.HasHeaders,
			Headers:    fp.Headers,
		}

		r[e.(string)] = p.ProcessEntity()
	}

	return r
}
