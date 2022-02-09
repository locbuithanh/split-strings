package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	buf, err := os.ReadFile("sentences_lower.txt")
	if err != nil {
		log.Println(err)
		return
	}
	str := strings.ReplaceAll(string(buf), "\r\n", "\n")
	s := strings.SplitN(str, "\n", -1)
	res := ""
	re := regexp.MustCompile(`iễu|ưỡi|uỗi|uấy|oáy|ươu|oẹo|yểu|oèo|iây|uẫy|uéo` +
		`|uai|iời|iấu|oai|ưởi|oải|uya|uỷu|oéo|iai|iũa|iườ|uái|uyễ` +
		`|uội|iàu|oái|uầy|yếu|uôi|uay|iày|iải|iấy|ượu|iỏi|oay` +
		`|iểu|ểu|iệu|ệu|iêu|yêu|êu|uyể|yể|uyế|yế` +
		`|iã|ỉu|ẫy|uỵ|ìu|ửu|ừu|ĩu|uề|oó|ỡi|ọe|ũy|ué|oẵ|iâ|uẫ|uắ|uã|iở|iỗ|iơ|iă|iô|iu` +
		`|uỗ|ụa|ũa|oo|uỹ|ảu|uể|oè|iặ|iũ|iò|oắ|oằ|ẽo|ựu|ọa|õi|òe|óe|ụy` +
		`|iậ|uở|uơ|oé|uỷ|oẻ|iư|uạ|ũi|íu|oã|ỉa|ẻo|ìa|oẹ|uỳ|uà|ùy|iọ|ủi` +
		`|uệ|ỏa|oe|ãi|ụi|oă|ẹo|ùi|òi|ue|ía|iỏ|èo|ưỡ|ữu|úy|ỵ`)
	for _, i := range s {
		res += fmt.Sprintf("%s|%s\n", i, strings.Join((re.FindAllString(i, -1)), ", "))
	}
	f, err := os.Create("test.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	f.WriteString(res)
}

// func main() {
// 	buf, err := os.ReadFile("input.txt")
// 	// buf, err := os.ReadFile("select.txt")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	str := strings.ToLower(string((buf)))
// 	// str := string((buf))
// 	s := split(replace(replaceMarks(str)))
// 	var res string
// 	count := 1287
// 	// count := 0
// 	for _, i := range s {
// 		if regexp.MustCompile("[0-9]+").MatchString(i) {
// 			continue
// 		}
// 		i = strings.TrimSpace(strings.Trim(i, `-,;: `)) + "."
// 		// i = firstUpper(i)
// 		len := utf8.RuneCountInString(i)
// 		if filter(res, i, len) {
// 			count++
// 			res += fmt.Sprintf("%d|%d|%s\n", count, len, i)
// 		}
// 		if count == 2800 {
// 			break
// 		}
// 	}
// 	f, err := os.Create("output.txt")
// 	// f, err := os.Create("output3_lower.txt")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer f.Close()
// 	f.WriteString(returnMarks(res))
// 	fmt.Println(strings.Count(res, "\n"))
// }

func firstUpper(str string) string {
	if len(str) == 0 {
		return ""
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}

func split(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == '.' || r == '\n'
	})
}

func replaceMarks(s string) string {
	r := strings.NewReplacer("?", "?.", "!", "!.")
	return r.Replace(s)
}

func returnMarks(s string) string {
	r := strings.NewReplacer("?.", "?", "!.", "!")
	return r.Replace(s)
}

func replace(s string) string {
	r := strings.NewReplacer(
		"tp hcm", "thành phố hồ chí minh",
		"TP HCM", "thành phố Hồ Chí Minh",
		"covid-19", "cô vít",
		"Covid-19", "Cô vít",
		"covid", "cô vít",
		"Covid", "Cô vít",
		" 1 ", " một ",
		" 2 ", " hai ",
		" 3 ", " ba ",
		" 4 ", " bốn ",
		" 5 ", " năm ",
		" 6 ", " sáu ",
		" 7 ", " bảy ",
		" 8 ", " tám ",
		" 9 ", " chín ",
		" 10 ", " mười ",
		" 16 ", " mười sáu ",
		"kg", "cân",
		"ôtô", "ô tô",
		"noel", "giáng sinh",
		"Noel", "Giáng sinh",
		"vaccine", "vắc xin",
		"Vaccine", "Vắc xin",
		"online", "trực tuyến",
		"Online", "Trực tuyến",
		"australia", "úc",
		"Australia", "Úc",
		"click", "kích",
		"internet", "in tơ nét",
		"Internet", "In tơ nét",
		"bus", "buýt",
		"taxi", "tắc xi",
		"jesse", "lâm",
		"Jesse", "Lâm",
		"jess", "lâm",
		"Jess", "Lâm",
		"nylon", "ni lông",
		"virus", "vi rút",
		"Virus", "Vi rút",
		"singapore", "xinh-ga-po",
		"Singapore", "Xinh-ga-po",
		"thcs", "trung học cơ sở",
		"THCS", "trung học cơ sở",
		"thpt", "trung học phổ thông",
		"THPT", "trung học phổ thông",
		"facebook", "phây búc",
		"Facebook", "Phây búc",
		"visa", "thị thực",
		"Visa", "Thị thực",
		"canada", "ca-na-đa",
		"Canada", "Ca-na-đa",
		"wikipedia", "bách khoa toàn thư mở",
		"Wikipedia", "Bách khoa toàn thư mở",
		"start-up", "công ty khởi nghiệp",
		"Start-up", "Công ty khởi nghiệp",
		"gs", "giáo sư",
		"GS", "Giáo sư",
		"delta", "mới",
		"Delta", "mới",
		"app", "ứng dụng",
	)
	return r.Replace(s)
}

func filter(res, s string, len int) bool {
	return len > 30 && len < 95 && !strings.Contains(res, s) && string(rune(s[0])) != "0"
	// return len < 95 && !strings.Contains(res, s) && string(rune(s[0])) != "0"
}
