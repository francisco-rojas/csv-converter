package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	vs := []string{"first name", "last name", "company"}
	r := Index(vs, "company")

	assert.Equal(t, 2, r)
}

func TestInArray(t *testing.T) {
	type args struct {
		s string
		a []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "when string is containted in array",
			args: args{
				s: "Lorem",
				a: []string{"Veritas", "Ipsum", "Lorem"},
			},
			want: true,
		},
		{
			name: "when string is NOT containted in array",
			args: args{
				s: "Lorem",
				a: []string{"Veritas", "Ipsum", "Sapienta"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := InArray(tt.args.s, tt.args.a)
			assert.Equal(t, tt.want, r)
		})
	}
}
