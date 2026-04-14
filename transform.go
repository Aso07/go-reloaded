package main
import(
	"strings"
	"strconv"
)
func process(text string) string {
	words := strings.Fields(text)
	result := []string{}

	insideQuote := false
	justOpened := false
	justClosed := false


	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "'" {
    insideQuote = !insideQuote

	if insideQuote {
		justOpened = true
	}else{
		justClosed = true
	}
	continue
		}


		if word == "(up)" && len(result) > 0 {
			result[len(result)-1] = toUpper(result[len(result)-1])
			continue
			}

			if word == "(low)" && len(result) > 0 {
				result[len(result)-1] = toLower(result[len(result)-1])
				continue
				}

				if word == "(cap)" && len(result) > 0 {
					result[len(result)-1] = capitalize(result[len(result)-1])
					continue
					}

					if strings.HasPrefix(word, "(up,") { 
					numstr := strings.TrimSuffix(strings.TrimPrefix(word, "(up,"), ")")
					n, _ := strconv.Atoi(strings.TrimSpace(numstr))

					for j := 1; j <= n && len(result)-j >= 0; j++ {
						result[len(result)-j] = toUpper(result[len(result)-j]) 
					}
					continue
					}

					if strings.HasPrefix(word, "(low,") {
						numstr := strings.TrimSuffix(strings.TrimPrefix(word, "(low,"), ")") 
						n, _ := strconv.Atoi(strings.TrimSpace(numstr))

						for j := 1; j <= n && len(result)-j >= 0; j++ {
							result[len(result)-j] = toLower(result[len(result)-j])
						}
						continue
					}

					if strings.HasPrefix(word, "(cap,") {
						numstr := strings.TrimSuffix(strings.TrimPrefix(word, "(cap,"), ")")
						n, _ := strconv.Atoi(strings.TrimSpace(numstr))

						for j := 1; j <= n && len(result)-j >= 0; j++ {
							result[len(result)-j] = capitalize(result[len(result)-j])
						}
						continue
					}

					if word == "(hex)" && len(result) > 0 {
					converted, err := fromHex(result[len(result)-1])
					if err == nil {
						result[len(result)-1] = converted
					}
					continue
					}
					if word == "(bin)" && len(result) > 0 {
						converted, err := fromBin(result[len(result)-1])
						if err == nil {
							result[len(result)-1] = converted
						}
						continue
					}

					if justOpened {
						word = "'" + word
						justOpened = false
					}


					result = append(result, word)

					if justClosed && len(result) > 0 {
						result[len(result)-1] = result[len(result)-1] + "'"
						justClosed = false
					}
				}
				return strings.Join(result, " ")
			}
func toUpper(s string) string {
	return strings.ToUpper(s)
}
func toLower(s string) string {
	return strings.ToLower(s)
}
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}
func fromHex(s string) (string, error) {
	result, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(result)), nil
}
func fromBin(s string) (string, error) {
	result, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(result)), nil
}

func fixPunct(s string) string {

	s = strings.ReplaceAll(s, " .", ".")
	s = strings.ReplaceAll(s, " ,", ",")
	s = strings.ReplaceAll(s, " !", "!")
	s = strings.ReplaceAll(s, " ;", ";")
	s = strings.ReplaceAll(s, " :", ":")
	s = strings.ReplaceAll(s, " ?", "?")

	words := strings.Fields(s)
	result := []string{}

	for _, word := range words {
		if (word == "." || word == "," || word == "!" || word == ":" || word == ";" || word == "?") && len(result) > 0 {
			result[len(result)-1] = result[len(result)-1] + word
		}else{
			result = append(result, word)
		}

	}
	return strings.Join(result, " ")
}

func fixAAn(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words) -1; i++ {
		if strings.ToLower(words[i]) == "a" {
			nextWord := strings.ToLower(words[i+1])

			if len(nextWord) > 0 {
				firstLetter := nextWord[0]
				if firstLetter == 'a' || firstLetter == 'e' || firstLetter == 'i' || firstLetter == 'o' || firstLetter == 'u' || firstLetter == 'h' {
					if words[i] == "A" {
						words[i] = "An"
					}else{
						words[i] = "an"
					}

				}
			}
		}

	}

	return strings.Join(words, " ")
}