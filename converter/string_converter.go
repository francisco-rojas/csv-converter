package converter

import "strings"

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
	d := strings.Trim(s.RawData, " ")

	if !s.isEmpty() {
		return d
	}

	m, ok := s.Options.(map[interface{}]interface{})
	if ok && m["default"] == nil {
		return s.zeroValue()
	}

	return m["default"]
}

func (s *String) isEmpty() bool {
	return s.RawData == ""
}

func (s *String) zeroValue() interface{} {
	return ""
}
