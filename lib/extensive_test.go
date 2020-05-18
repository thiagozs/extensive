package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputExtensive(t *testing.T) {
	ex := NewExtensive()

	tests := []struct {
		input  string
		output string
		expect bool
	}{
		{
			input:  "aaaa",
			output: "",
			expect: true,
		},
		{
			input:  "-aaaa",
			output: "",
			expect: true,
		},
		{
			input:  "12345",
			output: "doze mil e trezentos e quarenta e cinco",
			expect: true,
		},
		{
			input:  "1234567890",
			output: "um bilhão e duzentos e trinta e quatro milhões e quinhentos e sessenta e sete mil e oitocentos noventa",
			expect: true,
		},
		{
			input:  "123",
			output: "cento e vinte e três",
			expect: true,
		},
		{
			input:  "12",
			output: "doze",
			expect: true,
		},
		{
			input:  "1",
			output: "um",
			expect: true,
		},
		{
			input:  "20",
			output: "vinte",
			expect: true,
		},
		{
			input:  "1234567890119",
			output: "um trilhão e duzentos e trinta e quatro bilhões e quinhentos e sessenta e sete milhões e oitocentos noventa mil e cento e dezenove",
			expect: true,
		},
		{
			input:  "200",
			output: "duzentos",
			expect: true,
		},
		{
			input:  "1200",
			output: "um mil e duzentos",
			expect: true,
		},
		{
			input:  "10019",
			output: "dez mil e dezenove",
			expect: true,
		},
		{
			input:  "100200019",
			output: "cem milhões e duzentos mil e dezenove",
			expect: true,
		},
		{
			input:  "1001001001",
			output: "um bilhão e um milhão e um mil e um",
			expect: true,
		},
		{
			input:  "-1",
			output: "menos um",
			expect: true,
		},
		{
			input:  "999999",
			output: "novecentos e noventa e nove mil e novecentos e noventa e nove",
			expect: true,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			text := ex.Convert(test.input)
			assert.Equal(t, test.output, text)
		})
	}
}
