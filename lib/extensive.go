package lib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Mapper type cast
type Mapper map[int]string
type Extensive struct {
}

// NewExtensive return a new writer
func NewExtensive() *Extensive {
	names = mergeMaps(names, unidade, dezenas, dezena, centena)
	return &Extensive{}
}

func (e Extensive) Convert(number string) string {
	// negative context
	var output string
	if strings.HasPrefix(number, "-") {
		output += "menos "
	}

	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		panic(err)
	}

	nstr := reg.ReplaceAllString(number, "")
	if nstr != "" {
		if nstr == "0" {
			output = ""
		}
		output += numberWriter(nstr)
	} else {
		output = ""
	}

	return strings.TrimSpace(output)
}

func mergeMaps(left Mapper, elems ...Mapper) Mapper {
	for _, elem := range elems {
		for key, rightVal := range elem {
			if _, present := left[key]; !present {
				left[key] = rightVal
			}
		}
	}
	return left
}

func rangeMaps(number int) string {
	result := ""
	if v, ok := names[number]; ok {
		result = v
	}
	return result
}

func getGroup(number string) string {

	var output string

	if number == "000" {
		return output
	}

	if number == "100" {
		output = "cem"
		return output
	}

	switch len(number) {
	case 1:
		output = unit(number)
	case 2:
		output = dozen(number)
	case 3:
		output = hundred(number)
	}
	return output
}

func unit(number string) string {
	return rangeMaps(getInt(number[:1]))
}

func dozen(number string) string {
	var output string
	if strings.HasPrefix(number, "0") {
		//unit
		return rangeMaps(getInt(number[:1]))
	} else {
		//dozen
		_, dozen, unit, sum := getInfo(number)

		if !(sum > 10 && sum < 20) {
			output += rangeMaps(dozen)
			if unit != 0 {
				output += " e " + rangeMaps(unit)
			}
			return output
		}
		output += rangeMaps(sum)
	}
	return output
}

func hundred(number string) string {
	var output string
	if strings.HasPrefix(number, "00") {
		//unit
		unit := getInt(number[len(number)-1:])
		output += rangeMaps(unit)
	} else if strings.HasPrefix(number, "0") {
		// dozen
		_, dozen, unit, sum := getInfo(number)
		if !(sum > 10 && sum < 20) {
			output += rangeMaps(dozen)
			if unit != 0 {
				output += " e " + rangeMaps(unit)
			}
			return output
		}
		output += rangeMaps(sum)
	} else {
		// hundred
		hundred, dozen, unit, sum := getInfo(number)
		if dozen > 0 && unit > 0 {
			output += rangeMaps(hundred) + " e"
		} else {
			output += rangeMaps(hundred)
		}
		if sum == 0 {
			return output
		}
		if !(sum > 10 && sum < 20) && dozen > 0 {
			output += " " + rangeMaps(dozen)
			if unit != 0 {
				output += " e " + rangeMaps(unit)
			}
			return output
		}
		output += " " + rangeMaps(sum)
	}
	return output
}

func getInfo(number string) (int, int, int, int) {
	var (
		hundred int
		dozen   int
		unit    int
		sum     int
	)
	if len(number) == 3 {
		hundred = getInt(number[:1]) * 100
		dozen = getInt(number[1:2]) * 10
	} else if len(number) == 2 {
		dozen = getInt(number[:1]) * 10
	}
	unit = getInt(number[len(number)-1:])
	sum = dozen + unit
	return hundred, dozen, unit, sum
}

func getInt(number string) int {
	index, _ := strconv.Atoi(string(number))
	return index
}

// reverse split 2D matrix for search group of numbers
func split(buf string, limit int) [][]string {
	var divided [][]string
	var arr = strings.Split(buf, "")
	size := (len(arr) + limit - 1) / limit
	if size != 3 {
		size = 3
	}
	for i := len(arr); i >= 0; i -= size {
		end := i - size
		if end < 0 {
			end = 0
		}
		divided = append(divided, arr[end:i])
	}
	return divided
}

func numberWriter(number string) string {

	var output string        // output write
	nums := split(number, 3) //find cluster of numbers
	//fmt.Printf("%#v\n", nums)

	for i, num := range nums {
		// flat  the numbers
		fnum := strings.Join(num, "")
		// if array and string is not empty keep going
		if len(num) > 0 && fnum != "" {
			// get minemonic
			cas := casas[i]
			and := "e"
			group := getGroup(fnum)
			if group != "" {
				if i == 0 {
					and = ""
				}
				if singular(fnum) {
					cas = strings.Replace(cas, "ões", "ão", -1)
				}
				test := fmt.Sprintf("%s %s", group, cas)
				output = fmt.Sprintf("%s %s %s", test, and, output)
			}
		}
	}
	return output
}

func singular(number string) bool {
	hundred, dozen, unit, _ := getInfo(number)
	total := hundred + dozen + unit
	return total == 1
}
