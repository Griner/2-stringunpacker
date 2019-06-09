package stringunpacker

import (
	"fmt"
	"strings"
)

func itoa(i rune) (int, error) {
	switch i {
	case '0':
		return 0, nil
	case '1':
		return 1, nil
	case '2':
		return 2, nil
	case '3':
		return 3, nil
	case '4':
		return 4, nil
	case '5':
		return 5, nil
	case '6':
		return 6, nil
	case '7':
		return 7, nil
	case '8':
		return 8, nil
	case '9':
		return 9, nil
	default:
		return 0, fmt.Errorf("i isn't number")
	}
}

func UnpackString(packed string) (string, error) {

	var symbolInfo struct {
		s      rune
		c      int
		escape bool
	}
	var unpacked strings.Builder

	for _, symbol := range packed {

		if symbol == '\\' && !symbolInfo.escape {
			symbolInfo.escape = true

			if symbolInfo.s != '\x00' {
				if symbolInfo.c == 0 {
					symbolInfo.c = 1
				}
				unpacked.WriteString(strings.Repeat(string(symbolInfo.s), symbolInfo.c))
			}

			continue
		}

		if symbolInfo.escape {
			symbolInfo.s = symbol
			symbolInfo.c = 0
			symbolInfo.escape = false
			continue
		}

		if i, err := itoa(symbol); err == nil {
			symbolInfo.c = symbolInfo.c*10 + i
		} else {
			if symbolInfo.s != '\x00' {
				if symbolInfo.c == 0 {
					symbolInfo.c = 1
				}
				unpacked.WriteString(strings.Repeat(string(symbolInfo.s), symbolInfo.c))
			}

			symbolInfo.s = symbol
			symbolInfo.c = 0
		}
	}

	if symbolInfo.s != '\x00' {
		if symbolInfo.c == 0 {
			symbolInfo.c = 1
		}
		unpacked.WriteString(strings.Repeat(string(symbolInfo.s), symbolInfo.c))
	}

	if unpacked.String() == "" {
		return "", fmt.Errorf("Invalid string")
	}

	//log.Println(unpacked.String())

	return unpacked.String(), nil
}
