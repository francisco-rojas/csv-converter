package processors

import (
	"go-csv-converter/utils"
	"strings"
)

// EntityProcessor iterates over a data collection and processes the data according to the mappings.
type EntityProcessor struct {
	Row        []string
	Options    map[string]string
	Mappings   map[interface{}]interface{}
	HasHeaders bool
	Headers    []string
}

// ProcessEntity iterates over the attributes of an entity and processes the data accordingly.
func (ep *EntityProcessor) ProcessEntity() map[string]interface{} {
	if _, ok := ep.Mappings["nested"].(bool); ok {
		return ep.nestedEntity()
	}

	r := make(map[string]interface{})

	for a, m := range ep.Mappings {
		ep.Options["attr"] = a.(string)
		p := AttributeProcessor{
			Options:    ep.Options,
			Row:        ep.Row,
			Mappings:   m.(map[interface{}]interface{}),
			HasHeaders: ep.HasHeaders,
			Headers:    ep.Headers,
		}

		r[a.(string)] = p.ProcessAttribute()
	}

	return r
}

func (ep *EntityProcessor) nestedEntity() map[string]interface{} {
	s := ep.Mappings["separator"].(string)
	if s == "" {
		s = ","
	}

	h := ep.Mappings["header"].(string)
	i := utils.Index(ep.Headers, h)
	nr := ep.Row[i]
	r := strings.Split(nr, s)

	p := EntityProcessor{
		Options:    ep.Options,
		Row:        r,
		Mappings:   ep.Mappings["mappings"].(map[interface{}]interface{}),
		HasHeaders: ep.HasHeaders,
		Headers:    ep.Headers,
	}

	return p.ProcessEntity()
}
