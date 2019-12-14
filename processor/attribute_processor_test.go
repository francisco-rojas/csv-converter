package processor

import (
	"go-csv-converter/converter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttribute_ProcessAttribute(t *testing.T) {
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
		want   interface{}
	}{
		{
			name: "when column header provided",
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
		{
			name: "when column index provided",
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
			name: "trims spaces from a string attribute",
			fields: fields{
				Row:     []string{"John", "   Doe   "},
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

			assert.Equal(t, tt.want, ap.ProcessAttribute())
		})
	}
}

func TestAttribute_ProcessAttributePanics(t *testing.T) {
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
		want   interface{}
	}{
		{
			name: "panics when invalid column header provided",
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
			want: converter.ErrInvalidColumn{Header: "Last Name"},
		},
		{
			name: "panics when invalid column index provided",
			fields: fields{
				Row:     []string{"John", "Doe"},
				Options: map[string]string{},
				Mappings: map[interface{}]interface{}{
					"header": "5",
					"converters": map[interface{}]interface{}{
						"string": nil,
					},
				},
			},
			want: converter.ErrInvalidColumn{Header: "5"},
		},
		{
			name: "panics when invalid converter provided",
			fields: fields{
				Row:     []string{"John", "Doe"},
				Options: map[string]string{},
				Mappings: map[interface{}]interface{}{
					"header": "Last Name",
					"converters": map[interface{}]interface{}{
						"invalid": nil,
					},
				},
				HasHeaders: true,
				Headers:    []string{"First Name", "Last Name"},
			},
			want: converter.ErrInvalidConverter{Name: "invalid"},
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

			assert.PanicsWithValue(t, tt.want, func() { ap.ProcessAttribute() })
		})
	}
}
