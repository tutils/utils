package basen

import "errors"

var Base62CharSet = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

type Converter struct {
	base      int
	encodeMap []byte
	decodeMap [256]int
}

func (c *Converter) ToBaseN(number int) []byte {
	var baseN []byte
	for number > 0 {
		a := number / c.base
		baseN = append(baseN, c.encodeMap[number-a*c.base])
		number = a
	}
	for i, m, n := 0, len(baseN), len(baseN)/2; i < n; i++ {
		j := m - i - 1
		baseN[i], baseN[j] = baseN[j], baseN[i]
	}
	return baseN
}

func (c *Converter) ToNumber(baseN []byte) (int, error) {
	number := 0
	for _, b := range baseN {
		if n := c.decodeMap[b]; n < 0 {
			return 0, errors.New("basen: illegal byte")
		} else {
			number = number*c.base + n
		}

	}
	return number, nil
}

func NewConverter(charSet []byte) *Converter {
	decodeMap := [256]int{}
	for i, n := 0, len(decodeMap); i < n; i++ {
		decodeMap[i] = -1
	}

	for i, b := range charSet {
		decodeMap[b] = i
	}
	return &Converter{
		base:      len(charSet),
		encodeMap: charSet,
		decodeMap: decodeMap,
	}
}
