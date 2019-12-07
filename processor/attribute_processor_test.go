package processor

import (
	"go-csv-converter/converter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttribute_converterFor(t *testing.T) {
	type fields struct {
		Row        []string
		Options    map[string]string
		Mappings   map[interface{}]interface{}
		HasHeaders bool
		Headers    []string
	}
	type args struct {
		c string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   converter.Converter
	}{
		{
			name: "when string converter provided",
			fields: fields{
				Row:     []string{"John", "Doe"},
				Options: map[string]string{},
				Mappings: map[interface{}]interface{}{
					"header":     1,
					"converters": map[interface{}]interface{}{},
				},
			},
			args: args{"string"},
			want: &converter.String{RawData: "Doe", Options: map[string]string{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ap := &Attribute{
				Row:        tt.fields.Row,
				Options:    tt.fields.Options,
				Mappings:   tt.fields.Mappings,
				HasHeaders: tt.fields.HasHeaders,
				Headers:    tt.fields.Headers,
			}

			v := ap.converterFor(tt.args.c)
			assert.Equal(t, tt.want, v)
		})
	}
}

func TestAttribute_converterForErrors(t *testing.T) {
	type fields struct {
		Row        []string
		Options    map[string]string
		Mappings   map[interface{}]interface{}
		HasHeaders bool
		Headers    []string
	}
	type args struct {
		c string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "when invalid converter provided",
			fields: fields{
				Row:     []string{"John", "Doe"},
				Options: map[string]string{},
				Mappings: map[interface{}]interface{}{
					"header":     1,
					"converters": map[interface{}]interface{}{},
				},
			},
			args: args{"invalid"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ap := &Attribute{
				Row:        tt.fields.Row,
				Options:    tt.fields.Options,
				Mappings:   tt.fields.Mappings,
				HasHeaders: tt.fields.HasHeaders,
				Headers:    tt.fields.Headers,
			}

			assert.PanicsWithValue(t, converter.ErrInvalidConverter{Name: tt.args.c}, func() { ap.converterFor(tt.args.c) })
		})
	}
}

func TestAttribute_data(t *testing.T) {
	type fields struct {
		Row        []string
		Options    map[string]string
		Mappings   map[interface{}]interface{}
		HasHeaders bool
		Headers    []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "when no headers provided",
			fields: fields{
				Row:     []string{"John", "Doe"},
				Options: map[string]string{},
				Mappings: map[interface{}]interface{}{
					"header": 0,
					"converters": map[interface{}]interface{}{
						"string": nil,
					},
				},
			},
			want: "John",
		},
		{
			name: "when headers provided",
			fields: fields{
				Row:     []string{"John", "Doe"},
				Options: map[string]string{},
				Mappings: map[interface{}]interface{}{
					"header": "Last Name",
					"converters": map[interface{}]interface{}{
						"string": nil,
					},
				},
				HasHeaders: true,
				Headers:    []string{"First Name", "Last Name"},
			},
			want: "Doe",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ap := &Attribute{
				Row:        tt.fields.Row,
				Options:    tt.fields.Options,
				Mappings:   tt.fields.Mappings,
				HasHeaders: tt.fields.HasHeaders,
				Headers:    tt.fields.Headers,
			}

			assert.Equal(t, tt.want, ap.data())
		})
	}
}

func TestAttribute_dataErrors(t *testing.T) {
	type fields struct {
		Row        []string
		Options    map[string]string
		Mappings   map[interface{}]interface{}
		HasHeaders bool
		Headers    []string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "when invalid header provided",
			fields: fields{
				Row:     []string{"John", "Doe"},
				Options: map[string]string{},
				Mappings: map[interface{}]interface{}{
					"header": "Last Name",
					"converters": map[interface{}]interface{}{
						"string": nil,
					},
				},
				HasHeaders: true,
				Headers:    []string{"Name", "Surname"},
			},
		},
		{
			name: "when invalid header provided",
			fields: fields{
				Row:     []string{"John", "Doe"},
				Options: map[string]string{},
				Mappings: map[interface{}]interface{}{
					"header": 5,
					"converters": map[interface{}]interface{}{
						"string": nil,
					},
				},
				HasHeaders: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ap := &Attribute{
				Row:        tt.fields.Row,
				Options:    tt.fields.Options,
				Mappings:   tt.fields.Mappings,
				HasHeaders: tt.fields.HasHeaders,
				Headers:    tt.fields.Headers,
			}

			assert.PanicsWithValue(t, converter.ErrInvalidColumn{Header: ap.Mappings["header"]}, func() { ap.data() })
		})
	}
}
