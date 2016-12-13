package rpconv

import (
	"fmt"
	"math"
	"strings"
)

var (
	BahasaNumberMap = map[int]string{
		1: "SATU",
		2: "DUA",
		3: "TIGA",
		4: "EMPAT",
		5: "LIMA",
		6: "ENAM",
		7: "TUJUH",
		8: "DELAPAN",
		9: "SEMBILAN",
	}
)

const (
	tens       = "PULUH"
	singletens = "BELAS"
	hundreds   = "RATUS"
	thousands  = "RIBU"
	millions   = "JUTA"
	billions   = "MILYAR"
)

type (
	Rupiah struct {
		billions  hundred
		millions  hundred
		thousands hundred
		ones      hundred
	}

	hundred struct {
		val int
	}
)

//return error if input is less than 0 or greater equal 1 trillion
func Convert(inp float64) (string, error) {
	if inp < 0 {
		return "", fmt.Errorf("Only supports number greater or equal 0")
	}
	if inp >= math.Pow(1000, 4) {
		return "", fmt.Errorf("Only supports number lesser than 10^12")
	}
	return fmt.Sprintf("%s", new(inp)), nil
}

func new(inp float64) *Rupiah {
	var result Rupiah

	for power := 3; power >= 0; power-- {
		if hndrd := int(inp) / int(math.Pow(1000, float64(power))); hndrd > 0 {
			switch power {
			case 3:
				result.billions = hundred{val: hndrd}
			case 2:
				result.millions = hundred{val: hndrd}
			case 1:
				result.thousands = hundred{val: hndrd}
			case 0:
				result.ones = hundred{val: hndrd}
			}
			//deduct input by this value
			inp = inp - float64(hndrd)*math.Pow(1000, float64(power))
		}
	}

	return &result
}

func (rp Rupiah) String() string {
	str := []string{}

	if rp.billions.val > 0 {
		str = append(str, fmt.Sprintf("%s %s", rp.billions, billions))
	}

	if rp.millions.val > 0 {
		str = append(str, fmt.Sprintf("%s %s", rp.millions, millions))
	}

	if rp.thousands.val > 0 {
		str = append(str, fmt.Sprintf("%s %s", rp.thousands, thousands))
	}

	if rp.ones.val > 0 {
		str = append(str, fmt.Sprintf("%s", rp.ones))
	}

	if len(str) == 0 {
		str = append(str, "NOL")
	}

	str = append(str, "RUPIAH")

	return strings.Join(str, " ")
}

func (h hundred) String() string {
	str := []string{}

	//check the hundred
	if hundrd := h.val / 100; hundrd > 0 {
		//make 2 conditions: one for 1 and one for the rest
		switch hundrd {
		case 1:
			str = append(str, fmt.Sprintf("SE%s", hundreds))
		default:
			str = append(str, fmt.Sprintf("%s %s",
				BahasaNumberMap[hundrd], hundreds))
		}
		//now decrement the value by x00
		h.val = h.val - hundrd*100
	}

	//check the ten
	if ten := h.val / 10; ten > 0 {
		//also make 2 conditions for 1
		switch ten {
		case 1:
			//special case if tens is equal 1
			restVal := h.val - ten*10
			str = append(str, createSingleTens(restVal))
			return strings.Join(str, " ")
		default:
			str = append(str, fmt.Sprintf("%s %s",
				BahasaNumberMap[ten], tens))
		}
		//decrement the value by x0
		h.val = h.val - ten*10
	}

	if h.val > 0 {
		str = append(str, BahasaNumberMap[h.val])
	}

	return strings.Join(str, " ")
}

//another special case in Bahasa for single ten
//e.g. 1 -> sebelas, 2 -> dua belas
func createSingleTens(val int) string {
	switch val {
	case 0:
		return fmt.Sprintf("SE%s", tens)
	case 1:
		return fmt.Sprintf("SE%s", singletens)
	default:
		return fmt.Sprintf("%s %s", BahasaNumberMap[val], singletens)
	}
}
