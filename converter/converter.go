package converter

import "fmt"

// Converter is the interface that wraps the Converters methods
type Converter interface {
	isEmpty() bool
	Run() interface{}
	data() interface{}
	zeroValue() interface{}
}

// ErrInvalidConverter is raised when the converter provided in the mappings doesn't exist
type ErrInvalidConverter struct {
	Name string
}

func (e *ErrInvalidConverter) Error() string {
	return fmt.Sprintf("%s is not a valid converter", e.Name)
}

// ErrInvalidColumn is raised when the column header/index provided in the mappings doesn't exist in the file
type ErrInvalidColumn struct {
	Header interface{}
}

func (e *ErrInvalidColumn) Error() string {
	return fmt.Sprintf("there is no column '%v' in the file", e.Header)
}

// ErrNoSeparatorProvided is raised when no separator is provided for splitting nested csv strings
type ErrNoSeparatorProvided struct {
	Name string
}

func (e *ErrNoSeparatorProvided) Error() string {
	return fmt.Sprintf("No separator was providing for the nested entity '%s'", e.Name)
}
