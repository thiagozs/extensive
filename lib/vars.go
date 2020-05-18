package lib

var (
	names = map[int]string{}

	casas = []string{
		"",
		"mil",
		"milhões",
		"bilhões",
		"trilhões",
		"quatrilhões",
		"quintilhões",
		"sextilhões",
		"septilhões",
		"octilhões",
		"nonilhões",
		"decilhões",
		"undecilhões",
		"duodecilhões",
		"tredecilhões",
	}

	dezenas = map[int]string{
		11: "onze",
		12: "doze",
		13: "treze",
		14: "quatorze",
		15: "quinze",
		16: "dezesseis",
		17: "dezessete",
		18: "dezoito",
		19: "dezenove",
	}

	unidade = map[int]string{
		0: "zero",
		1: "um",
		2: "dois",
		3: "três",
		4: "quatro",
		5: "cinco",
		6: "seis",
		7: "sete",
		8: "oito",
		9: "nove",
	}

	dezena = map[int]string{
		10: "dez",
		20: "vinte",
		30: "trinta",
		40: "quarenta",
		50: "cinquenta",
		60: "sessenta",
		70: "setenta",
		80: "oitenta",
		90: "noventa",
	}

	centena = map[int]string{
		100: "cento",
		200: "duzentos",
		300: "trezentos",
		400: "quatrocentos",
		500: "quinhentos",
		600: "seiscentos",
		700: "setecentos",
		800: "oitocentos",
		900: "novecentos",
	}
)
