package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	buf, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
		return
	}

	str := strings.ToLower(string((buf)))
	s := split(replace(str))
	var res string
	count := 0
	d := 0
	for _, i := range s {
		i = strings.Trim(i, `-,;: `) + "."
		len := utf8.RuneCountInString(i)
		if filter(res, i, len) {
			count++
			res += fmt.Sprintf("%d|%d|%s\n", count, len, i)
			if len > 90 {
				d++
			}
		}
	}
	f, err := os.Create("output.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	f.WriteString(res)
	fmt.Println(strings.Count(res, "\n"), d)
}

func split(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == '.' || r == '?' || r == '!' || r == '\n'
	})
}

func replace(s string) string {
	r := strings.NewReplacer(
		"tp hcm", "thành phố hồ chí minh",
		"covid-19", "cô vít",
		"covid", "cô vít",
		" 1 ", " một ",
		" 2 ", " hai ",
		" 3 ", " ba ",
		" 4 ", " bốn ",
		" 5 ", " năm ",
		" 6 ", " sáu ",
		" 7 ", " bảy ",
		" 8 ", " tám ",
		" 9 ", " chín ",
		" 16 ", " mười sáu ",
		"kg", "cân",
		"vaccine", "vắc xin",
		"online", "trực tuyến",
		"australia", "úc",
		"click", "kích",
	)
	return r.Replace(s)
}

func filter(res, s string, len int) bool {
	return len > 20 && len < 95 && !strings.Contains(res, s) && string(rune(s[0])) != "0"
}
