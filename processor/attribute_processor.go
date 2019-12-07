package processor

import (
	"fmt"
	"go-csv-converter/converter"
	"go-csv-converter/utils"
)

// Attribute iterates over a data collection and processes the data according to the mappings.
type Attribute struct {
	Row        []string
	Options    map[string]string
	Mappings   map[interface{}]interface{}
	HasHeaders bool
	Headers    []string
}

// ProcessAttribute iterates over the attributes of an attribute and processes the data accordingly.
func (ap *Attribute) ProcessAttribute() interface{} {
	m, ok := ap.Mappings["converters"].(map[interface{}]interface{})
	if !ok {
		return ap.data()
	}

	var d interface{}
	for conv, opts := range m {
		d, _ = ap.invokeConverter(conv.(string), opts)
	}

	return d
}

func (ap *Attribute) invokeConverter(c string, opts interface{}) (interface{}, error) {
	o, ok := opts.(map[interface{}]interface{})
	if ok {
		for k, v := range o {
			ap.Options[k.(string)] = fmt.Sprint(v)
		}
	}

	e, err := ap.converterFor(c)
	if err != nil {
		return nil, err
	}

	return e.Run(), nil
}

func (ap *Attribute) data() string {
	var i int
	h, ok := ap.Mappings["header"].(string)
	// column header
	if ok {
		i = utils.Index(ap.Headers, h)
	} else {
		// column position
		i = ap.Mappings["header"].(int)
	}

	if i == -1 || i >= len(ap.Row) {
		panic(converter.ErrInvalidColumn{Header: ap.Mappings["header"]})
	}

	return ap.Row[i]
}

func (ap *Attribute) converterFor(c string) (converter.Converter, error) {
	a := map[string]converter.Converter{
		"string": &converter.String{RawData: ap.data(), Options: ap.Options},
		// "boolean":   &converter.String{RawData: ap.data(), Options: ap.Options},
		// "uppercase": &converter.String{RawData: ap.data(), Options: ap.Options},
		// "date":      &converter.String{RawData: ap.data(), Options: ap.Options},
		// "integer":   &converter.String{RawData: ap.data(), Options: ap.Options},
		// "decimal":   &converter.String{RawData: ap.data(), Options: ap.Options},
		// "float":     &converter.String{RawData: ap.data(), Options: ap.Options},
	}

	conv := a[c]
	if conv == nil {
		return nil, &converter.ErrInvalidConverter{Name: c}
	}

	return a[c], nil
}
