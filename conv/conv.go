package conv

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/azer/snakecase"
)

func IntOr(source string, deflt int) int {
	out, err := strconv.Atoi(source)
	if err != nil {
		return deflt
	}
	return out
}

func CaseSnake(text string) string {
	return snakecase.SnakeCase(text)
}

var unwantedChars = regexp.MustCompile(`[^a-zA-Z0-9\-]+`)

func CaseURL(text string, conf ...UrlConfig) string {
	var words []string
	l := 0
	for s := text; s != ""; s = s[l:] {
		l = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
		if l <= 0 {
			l = len(s)
		}
		words = append(words, s[:l])
	}

	url := strings.ToLower(strings.Join(words, "-"))
	url = strings.Replace(url, "--", "-", -1) // pure hack. todo: reg-ex

	// remove extra chars
	if len(conf) > 0 && conf[0].RemoveExtraCharacters {
		url = unwantedChars.ReplaceAllString(url, "")
	}

	return url
}

func CaseSentence(text string) string {
	if len(text) > 0 {
		u := []rune(text)
		u[0] = unicode.ToUpper(u[0])
		return string(u)
	}

	return text
}

type UrlConfig struct {
	RemoveExtraCharacters bool // any chars other than A-Za-z0-9 and -
}
