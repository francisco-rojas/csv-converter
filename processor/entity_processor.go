package processor

import (
	"go-csv-converter/utils"
	"strings"
)

// Entity iterates over a data collection and processes the data according to the mappings.
type Entity struct {
	Row        []string
	Options    map[string]string
	Mappings   map[interface{}]interface{}
	HasHeaders bool
	Headers    []string
}

// ProcessEntity iterates over the attributes of an entity and processes the data accordingly.
func (ep *Entity) ProcessEntity() map[string]interface{} {
	if _, ok := ep.Mappings["nested"].(bool); ok {
		return ep.nestedEntity()
	}

	r := make(map[string]interface{})

	for a, m := range ep.Mappings {
		ep.Options["attr"] = a.(string)
		p := Attribute{
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

func (ep *Entity) nestedEntity() map[string]interface{} {
	s := ep.Mappings["separator"].(string)
	if s == "" {
		s = ","
	}

	h := ep.Mappings["header"].(string)
	i := utils.Index(ep.Headers, h)
	nr := ep.Row[i]
	r := strings.Split(nr, s)

	p := Entity{
		Options:    ep.Options,
		Row:        r,
		Mappings:   ep.Mappings["mappings"].(map[interface{}]interface{}),
		HasHeaders: ep.HasHeaders,
		Headers:    ep.Headers,
	}

	return p.ProcessEntity()
}
