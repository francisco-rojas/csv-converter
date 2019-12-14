package processor

import (
	"go-csv-converter/converter"
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
	s, ok := ep.Mappings["separator"].(string)

	var i int
	h, hok := ep.Mappings["header"].(string)
	// column header
	if hok {
		i = utils.Index(ep.Headers, h)
	} else {
		// column position
		i = ep.Mappings["header"].(int)
	}

	if !ok {
		panic(converter.ErrNoSeparatorProvided{Name: h})
	}

	if i == -1 || i >= len(ep.Row) {
		panic(converter.ErrInvalidColumn{Header: ep.Mappings["header"]})
	}

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
