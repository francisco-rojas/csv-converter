package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString_isEmpty(t *testing.T) {
	type fields struct {
		RawData string
		Options interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "when empty string provided",
			fields: fields{
				RawData: "",
				Options: nil,
			},
			want: true,
		},
		{
			name: "when valid string provided",
			fields: fields{
				RawData: "data",
				Options: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &String{
				RawData: tt.fields.RawData,
				Options: tt.fields.Options,
			}
			assert.Equal(t, tt.want, s.isEmpty())
		})
	}
}

func TestString_zeroValue(t *testing.T) {
	type fields struct {
		RawData string
		Options interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name:   "returns emtpy string",
			fields: fields{},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &String{
				RawData: tt.fields.RawData,
				Options: tt.fields.Options,
			}
			assert.Equal(t, tt.want, s.zeroValue())
		})
	}
}

func TestString_data(t *testing.T) {
	type fields struct {
		RawData string
		Options interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "when valid input provided",
			fields: fields{
				RawData: " Lorem1 ",
				Options: nil,
			},
			want: "Lorem1",
		},
		{
			name: "when empty string and no default value provided",
			fields: fields{
				RawData: "  ",
				Options: nil,
			},
			want: "",
		},
		{
			name: "when empty string and default value provided",
			fields: fields{
				RawData: "  ",
				Options: map[interface{}]interface{}{
					"default": "Lorem2",
				},
			},
			want: "Lorem2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &String{
				RawData: tt.fields.RawData,
				Options: tt.fields.Options,
			}
			d := s.data()
			assert.Equal(t, tt.want, d)
		})
	}
}
