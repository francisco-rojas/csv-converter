package converter

// Converter is the interface that wraps the Converters methods
type Converter interface {
	isEmpty() bool
	Run() interface{}
	data() interface{}
	zeroValue() interface{}
}
