package converter

import (
	"go-csv-converter/utils"
	"strings"
)

// String cleans up a string
type String struct {
	RawData string
	Options interface{}
}

// Run converts Data to string
func (s *String) Run() interface{} {
	return s.data()
}

// Data returns the data to be converted
func (s *String) data() interface{} {
	s.RawData = strings.Trim(s.RawData, " ")

	if !s.isEmpty() {
		return s.RawData
	}

	m, ok := s.Options.(map[interface{}]interface{})
	if !ok {
		return s.zeroValue()
	}

	if m["default"] == nil {
		return s.zeroValue()
	}

	return m["default"]
}

func (s *String) isEmpty() bool {
	if s.RawData == "" {
		return true
	}

	m, ok := s.Options.(map[interface{}]interface{})
	if !ok {
		return s.RawData == ""
	}

	v, ok := m["empty_values"].([]string)
	if !ok {
		return s.RawData == ""
	}

	return utils.InArray(s.RawData, v)
}

func (s *String) zeroValue() interface{} {
	return ""
}
