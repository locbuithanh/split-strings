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
	if err != nil {
		log.Println(err)
		return
	}
	var res string
	count := 0
	for _, i := range s {
		i = strings.TrimSpace(i) + "."
		len := utf8.RuneCountInString(i)
		if len > 20 && len < 100 && !strings.Contains(res, i) {
			count++
			res += fmt.Sprintf("%d|%d|%s\n", count, len, i)
		}
	}
	f, err := os.Create("output.txt")
	defer f.Close()
	f.WriteString(res)
	fmt.Println(strings.Count(res, "\n"))
}

func split(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == '.' || r == '?' || r == '!' || r == '\n'
	})
}

func replace(s string) string {
	r := strings.NewReplacer("tp hcm", "thành phố hồ chí minh",
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
	)
	return r.Replace(s)
}