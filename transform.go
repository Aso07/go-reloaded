package main

import (
	"strconv"
	"strings"
)
func process(text string) string {
	words := strings.Fields(text)
	words = joinTags(words)
	result := []string{}

	for _, word := range words {
			if word == "(up)" {
				result[len(result)-1] = toUpper(result[len(result)-1])

			}else if word == "(low)" {
				result[len(result)-1] = toLower(result[len(result)-1])

			}else if word == "(cap)" {
				result[len(result)-1] = capitalize(result[len(result)-1])

			}else if strings.HasPrefix(word, "(up,") {
				count := getCount(word)
				for i := len(result) - count; i < len(result); i++ {
					result[i] = toUpper(result[i])
				}
			}else if strings.HasPrefix(word, "(low,") {
				count := getCount(word)
				for i := len(result) - count; i < len(result); i++ {
					result[i] = toLower(result[i])
				}
			}else if strings.HasPrefix(word, "(cap,") {
				count := getCount(word)
				for i := len(result) - count; i < len(result); i++ {
					result[i] = capitalize(result[i])
			}			
			
			}else if word == "(hex)" {
				converted, err := fromHex(result[len(result)-1])
				if err == nil {
					result[len(result)-1] = converted
				} 

			}else if word == "(bin)" {
				converted, err := fromBin(result[len(result)-1])
				if err == nil {
					result[len(result)-1] = converted
				}
			}else{
				result = append(result, word)
			}
			
		}
		result = fixPunctuation(result)
		result = fixArticles(result)
		return strings.Join(result, " ")
		


	
}
func toUpper(s string) string {
	return strings.ToUpper(s)
}
func toLower(s string) string {
	return strings.ToLower(s)
}
func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}
func fromHex(s string) (string, error) {
result, err := strconv.ParseInt(s, 16,64)
if err != nil {
	return "", err
}
return strconv.Itoa(int(result)), nil
}
func fromBin(s string) (string, error) {
	result, err := strconv.ParseInt(s, 2,64)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(result)), nil
}


func getCount(tag string) int {
	word := strings.Split(tag, ",")
	cleaned := strings.Trim(strings.TrimSpace(word[1]),")")
	n, err := strconv.Atoi(cleaned)
	if err != nil || n < 1{
		return 1
	}
	return n
}

func joinTags(words []string) []string {
	result := []string{}
	
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(") && strings.HasSuffix(words[i], ",") && i+1 < len(words) {
		result = append(result, words[i]+words[i+1])
		i++
		}else{
			result = append(result, words[i])
		}

	}
	return result
}

func fixPunctuation(words []string) []string {
	result := []string{}
	for _, word := range words {
		if strings.Trim(word, ".,!?:;") == "" {
			result[len(result)-1] = result[len(result)-1] + word
		}else{
			result = append(result, word)
		}


	}
	return result

}

func fixArticles(words []string) []string {
	result := []string{}
	for i, word := range words {
		if (word == "a" || word == "A") && i+1 < len(words) && strings.ContainsAny(words[i+1][:1], "aeiouAEIOUhH") {
			if word == "a" {
				word = "an"

		}else{
			word = "An"
			
		}
		result = append(result, word)

	}else{
	result = append(result, word)
}
}
return result
}