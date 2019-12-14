package processor

import (
	"go-csv-converter/converter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntity_ProcessEntity(t *testing.T) {
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
		want   map[string]interface{}
	}{
		{
			name: "when column headers provided",
			fields: fields{
				Row:     []string{"John", "Doe"},
				Options: map[string]string{},
				Mappings: map[interface{}]interface{}{
					"firstName": map[interface{}]interface{}{
						"header": "First Name",
						"converters": map[interface{}]interface{}{
							"string": nil,
						},
					},
					"lastName": map[interface{}]interface{}{
						"header": "Last Name",
						"converters": map[interface{}]interface{}{
							"string": nil,
						},
					},
				},
				HasHeaders: true,
				Headers:    []string{"First Name", "Last Name"},
			},
			want: map[string]interface{}{
				"firstName": "John",
				"lastName":  "Doe",
			},
		},
		{
			name: "when column index provided",
			fields: fields{
				Row:     []string{"John", "Doe"},
				Options: map[string]string{},
				Mappings: map[interface{}]interface{}{
					"firstName": map[interface{}]interface{}{
						"header": 0,
						"converters": map[interface{}]interface{}{
							"string": nil,
						},
					},
					"lastName": map[interface{}]interface{}{
						"header": 1,
						"converters": map[interface{}]interface{}{
							"string": nil,
						},
					},
				},
			},
			want: map[string]interface{}{
				"firstName": "John",
				"lastName":  "Doe",
			},
		},
		{
			name: "when nested csv provided",
			fields: fields{
				Row:     []string{"John", "Doe", "NY,2nd Ave.,10521"},
				Options: map[string]string{},
				Mappings: map[interface{}]interface{}{
					"nested":    true,
					"separator": ",",
					"header":    "Address",
					"mappings": map[interface{}]interface{}{
						"state": map[interface{}]interface{}{
							"header": 0,
							"converters": map[interface{}]interface{}{
								"string": nil,
							},
						},
						"street": map[interface{}]interface{}{
							"header": 1,
							"converters": map[interface{}]interface{}{
								"string": nil,
							},
						},
						"zip": map[interface{}]interface{}{
							"header": 2,
							"converters": map[interface{}]interface{}{
								"string": nil,
							},
						},
					},
				},
				HasHeaders: true,
				Headers:    []string{"First Name", "Last Name", "Address"},
			},
			want: map[string]interface{}{
				"state":  "NY",
				"street": "2nd Ave.",
				"zip":    "10521",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ep := &Entity{
				Row:        tt.fields.Row,
				Options:    tt.fields.Options,
				Mappings:   tt.fields.Mappings,
				HasHeaders: tt.fields.HasHeaders,
				Headers:    tt.fields.Headers,
			}

			assert.Equal(t, tt.want, ep.ProcessEntity())
		})
	}
}

func TestEntity_ProcessEntityPanics(t *testing.T) {
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
			name: "when no separator provided for nested csv",
			fields: fields{
				Row:     []string{"John", "Doe", "NY,2nd Ave.,10521"},
				Options: map[string]string{},
				Mappings: map[interface{}]interface{}{
					"nested": true,
					"header": "Address",
					"mappings": map[interface{}]interface{}{
						"state": map[interface{}]interface{}{
							"header": 0,
							"converters": map[interface{}]interface{}{
								"string": nil,
							},
						},
						"street": map[interface{}]interface{}{
							"header": 1,
							"converters": map[interface{}]interface{}{
								"string": nil,
							},
						},
						"zip": map[interface{}]interface{}{
							"header": 2,
							"converters": map[interface{}]interface{}{
								"string": nil,
							},
						},
					},
				},
				HasHeaders: true,
				Headers:    []string{"First Name", "Last Name", "Address"},
			},
			want: converter.ErrNoSeparatorProvided{Name: "Address"},
		},
		{
			name: "when invalid column header provided",
			fields: fields{
				Row:     []string{"John", "Doe", "NY,2nd Ave.,10521"},
				Options: map[string]string{},
				Mappings: map[interface{}]interface{}{
					"nested":    true,
					"header":    "invalid",
					"separator": ",",
					"mappings": map[interface{}]interface{}{
						"state": map[interface{}]interface{}{
							"header": 0,
							"converters": map[interface{}]interface{}{
								"string": nil,
							},
						},
						"street": map[interface{}]interface{}{
							"header": 1,
							"converters": map[interface{}]interface{}{
								"string": nil,
							},
						},
						"zip": map[interface{}]interface{}{
							"header": 2,
							"converters": map[interface{}]interface{}{
								"string": nil,
							},
						},
					},
				},
				HasHeaders: true,
				Headers:    []string{"First Name", "Last Name", "Address"},
			},
			want: converter.ErrInvalidColumn{Header: "invalid"},
		},
		{
			name: "when invalid column index provided",
			fields: fields{
				Row:     []string{"John", "Doe", "NY,2nd Ave.,10521"},
				Options: map[string]string{},
				Mappings: map[interface{}]interface{}{
					"nested":    true,
					"header":    5,
					"separator": ",",
					"mappings": map[interface{}]interface{}{
						"state": map[interface{}]interface{}{
							"header": 0,
							"converters": map[interface{}]interface{}{
								"string": nil,
							},
						},
						"street": map[interface{}]interface{}{
							"header": 1,
							"converters": map[interface{}]interface{}{
								"string": nil,
							},
						},
						"zip": map[interface{}]interface{}{
							"header": 2,
							"converters": map[interface{}]interface{}{
								"string": nil,
							},
						},
					},
				},
			},
			want: converter.ErrInvalidColumn{Header: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ep := &Entity{
				Row:        tt.fields.Row,
				Options:    tt.fields.Options,
				Mappings:   tt.fields.Mappings,
				HasHeaders: tt.fields.HasHeaders,
				Headers:    tt.fields.Headers,
			}

			assert.PanicsWithValue(t, tt.want, func() { ep.ProcessEntity() })
		})
	}
}
