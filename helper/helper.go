package helper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Christopher John
func Upper(str string) string {
	words := strings.Fields(str)

	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i++
		} else if strings.HasPrefix(words[i], "(up") {
			words[i+1] = strings.TrimSuffix(words[i+1], ")")
			n, err := strconv.Atoi(words[i+1])
			if err != nil {
				continue
			}
			for j := 1; j <= n && i-j >= 0; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			i++
		}

	}

	return strings.Join(words, " ")
}

func FixAToAn(text string) string {
	words := strings.Fields(text)
	consonantSound := []string{
		"university", "unicorn", "use", "user", "unit", "euro", "one",
	}
	vowelSound := []string{
		"hour", "honor", "heir", "honest",
	}

	for i, word := range words {
		if word == "a" || word == "A" {
			if i+1 < len(words) {
				next := strings.ToLower(words[i+1])
				first := next[0]
				isVowelStart := strings.ContainsRune("aeiou", rune(first))
				for _, cs := range consonantSound {
					if strings.HasPrefix(next, cs) {
						isVowelStart = false
						break
					}
				}
				for _, vs := range vowelSound {
					if strings.HasPrefix(next, vs) {
						isVowelStart = true
						break
					}
				}
				if isVowelStart {
					if word == "a" {
						words[i] = "an"
					} else {
						words[i] = "An"
					}
				}
			}
		}
	}

	return strings.Join(words, " ")
}

// Benedict Onyeke(bonyeke)
func FixSingleQuotes(text string) string {
	text = strings.Trim(text, "'")

	text = strings.TrimSpace(text)

	return "'" + text + "'"
}

// Code From Roland Elaigwu(relaigwu)

// func FixQuotes(s string) string {
// 	r := []rune(s)
// 	var res []rune
// 	inQuote := false

// 	for i := 0; i < len(r); i++ {

// 		if r[i] == '\'' {

// 			if !inQuote {
// 				if len(res) > 0 && res[len(res)-1] != ' ' {
// 					res = append(res, ' ')
// 				}
// 				res = append(res, '\'')
// 				inQuote = true

// 				i++
// 				for i < len(r) && r[i] == ' ' {
// 					i++
// 				}
// 				i--

// 			} else {
// 				if len(res) > 0 && res[len(res)-1] == ' ' {
// 					res = res[:len(res)-1]
// 				}
// 				res = append(res, '\'')
// 				inQuote = false
// 			}
// 			continue
// 		}

// 		res = append(res, r[i])
// 	}

// 	return string(res)
// }

// Code from Excel
func ToLower(s string) string {
	slice := strings.Fields(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] == "(low)" {
			slice[i-1] = strings.ToLower(slice[i-1])

			slice = append(slice[:i], slice[i+1:]...)
		} else if strings.HasPrefix(slice[i], "(low") {
			slice[i+1] = strings.TrimSuffix(slice[i+1], ")")
			n, err := strconv.Atoi(slice[i+1])
			if err != nil {
				continue
			}
			for j := 1; j <= n && i-j >= 0; j++ {
				slice[i-j] = strings.ToLower(slice[i-j])
			}
			slice = append(slice[:i], slice[i+2:]...)
			i--
		}
	}
	return strings.Join(slice, " ")
}

// patrick christian

func ConvertNumbers(s string) string {
	reHex := regexp.MustCompile(`(-?[0-9A-Fa-f]+)\s*\(hex\)`)

	s = reHex.ReplaceAllStringFunc(s, func(m string) string {
		parts := reHex.FindStringSubmatch(m)
		if len(parts) < 2 {
			return m
		}
		n, err := strconv.ParseInt(parts[1], 16, 64)
		if err != nil {
			return m
		}
		return fmt.Sprint(n)
	})
	return s
}

// David Ayuba

func ToCap(dswayy string) string {
	s := strings.Fields(dswayy)
	for i := 0; i < len(s); i++ {
		if s[i] == "(cap)" {
			s[i-1] = strings.Title(s[i-1])
			s = append(s[:i], s[i+1:]...)
		} else if strings.HasPrefix(s[i], "(cap") {
			s[i+1] = strings.TrimSuffix(s[i+1], ")")
			n, err := strconv.Atoi(s[i+1])
			if err != nil {
				continue
			}
			for j := 1; j <= n && i-j >= 0; j++ {
				s[i-j] = strings.Title(s[i-j])
			}
			s = append(s[:i], s[i+2:]...)
			i--
		}
	}

	return strings.Join(s, " ")
}

// abraham zion

func BinTodecimal(bin string) string {

	words := strings.Fields(bin)
	for i, j := range words {
		if j == "(bin)" && i > 0 {
			result, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				fmt.Println("Error")
				continue
			}
			words[i-1] = strconv.Itoa(int(result))
			words[i] = ""
		}

	}
	return strings.TrimSpace(strings.Join(strings.Fields(strings.Join(words, " ")), " "))

}

func FixPunct(s string) string {
	s = strings.NewReplacer(" .", ".", " ?", "?", " ,", ", ").Replace(s)
	return s
}
