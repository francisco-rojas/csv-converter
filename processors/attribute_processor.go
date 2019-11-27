package processors

// AttributeProcessor iterates over a data collection and processes the data according to the mappings.
type AttributeProcessor struct {
	Row        []string
	Options    map[string]string
	Mappings   map[interface{}]interface{}
	HasHeaders bool
	Headers    []string
}

// ProcessAttribute iterates over the attributes of an attribute and processes the data accordingly.
func (ap *AttributeProcessor) ProcessAttribute() interface{} {
	var r interface{}

	return r
}
