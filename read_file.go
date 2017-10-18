// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"io/ioutil"
)

func readFile(name string) ([]byte, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	var escaped bool // string value flag
	var comments int // 1 line; 2 multi line
	var returns []byte

	length := len(data)
	for i, w := 0, 0; i < length; i += w {
		w = 1

		switch comments {
		case 1:
			if data[i] == '\n' {
				comments = 0
				escaped = false
			}
			continue
		case 2:
			if data[i] != '*' || length == i+1 {
				continue
			}
			if data[i+1] != '/' {
				continue
			}
			w = 2
			comments = 0
			escaped = false
			continue
		}
		switch data[i] {
		case '"':
			{
				if escaped {
					escaped = false
				} else {
					escaped = true
				}
				returns = append(returns, data[i])
			}
		case '/':
			{
				if escaped || length == i+1 {
					returns = append(returns, data[i])
					break
				}
				switch data[i+1] {
				case '/':
					w = 2
					comments = 1
				case '*':
					w = 2
					comments = 2
				default:
					returns = append(returns, data[i])
				}
			}
		default:
			if escaped || !isWhitespace(data[i]) {
				returns = append(returns, data[i])
			}

		}
	}
	return returns, nil
}

/*
SPACE (\u0020)
NO-BREAK SPACE (\u00A0)
OGHAM SPACE MARK (\u1680)
EN QUAD (\u2000)
EM QUAD (\u2001)
EN SPACE (\u2002)
EM SPACE (\u2003)
THREE-PER-EM SPACE (\u2004)
FOUR-PER-EM SPACE (\u2005)
SIX-PER-EM SPACE (\u2006)
FIGURE SPACE (\u2007)
PUNCTUATION SPACE (\u2008)
THIN SPACE (\u2009)
HAIR SPACE (\u200A)
NARROW NO-BREAK SPACE (\u202F)
MEDIUM MATHEMATICAL SPACE (\u205F)
and IDEOGRAPHIC SPACE (\u3000)
Byte Order Mark (\uFEFF)
*/
func isWhitespace(c byte) bool {
	str := string(c)

	switch str {
	case " ", "\t", "\n", "\u000B", "\u000C",
		"\u000D", "\u00A0", "\u1680", "\u2000",
		"\u2001", "\u2002", "\u2003", "\u2004",
		"\u2005", "\u2006", "\u2007", "\u2008",
		"\u2009", "\u200A", "\u202F", "\u205F",
		"\u2060", "\u3000", "\uFEFF":
		return true
	}
	return false
}
