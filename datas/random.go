package datas

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func init() {
	// to change value when random
	rand.Seed(time.Now().UnixNano())
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}

	return a[rand.Intn(n)]
}

func randomCharacter() string {
	return randomStringFromSet("a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z")
}

func randomNumber() string {
	return randomStringFromSet("0", "1", "2", "3", "4", "5", "6", "7", "8", "9")
}

func randomID() string {
	return uuid.New().String()
}

func randomLocale() string {
	return randomStringFromSet(
		"en-AU",
		"zh-CN",
		"en-AU",
		"de-CH",
	)
}

func randomTimezone() string {
	return randomStringFromSet(
		"Papua New Guinea",
		"Svalbard and Jan Mayen Islands",
		"Swaziland",
		"Reunion",
		"Gibraltar",
		"Mayotte",
	)
}

func randomGender() string {
	return randomStringFromSet("Mr", "Miss")
}

func randomRole() string {
	return randomStringFromSet(
		"end-user",
		"admin",
		"agent",
	)
}

func randomDateTime() string {
	now := time.Now()
	typeDate := rand.Intn(3)
	num := rand.Intn(13)

	switch typeDate {
	case 0:
		now = now.AddDate(0, 0, -num)
		break
	case 1:
		now = now.AddDate(0, -num, 0)
		break
	default:
		now = now.AddDate(0, -num, -num)
		break
	}

	return now.Format("2006-01-02T15:04:05Z07:00")
}

func randomEmail(alias string, LastName string) string {
	suffixEmail := randomStringFromSet(
		"@gmail.com",
		"@yahoo.com",
		"@flotonic.com",
	)

	return alias + LastName + suffixEmail
}

func RandomPhone() string {
	phone := ""
	for i := 0; i < 4; i++ {
		phone += randomNumber()
	}

	phone += "-"
	for i := 0; i < 3; i++ {
		phone += randomNumber()
	}

	phone += "-"
	for i := 0; i < 3; i++ {
		phone += randomNumber()
	}

	return phone
}

func RandomType() string {
	return randomStringFromSet(
		"question",
		"task",
		"problem",
	)
}

func randomName(n int) string {
	num := rand.Intn(5) + n
	str := ""
	for i := 0; i < num; i++ {
		str += randomCharacter()
	}
	return str
}

func randomText(n int) string {
	subject := ""
	num := rand.Intn(n) + 2
	for i := 0; i < num; i++ {
		name := randomName(3)
		if i == num {
			subject += name
		} else {
			subject += name + " "
		}

	}

	return subject
}

func randomPriority() string {
	return randomStringFromSet(
		"high",
		"low",
		"normal",
	)
}

func randomStatus() string {
	return randomStringFromSet(
		"hold",
		"pending",
		"closed",
	)
}

func randomVia() string {
	return randomStringFromSet(
		"chat",
		"web",
		"voice",
	)
}
func GetUrl(suffix string) string {
	return "http://initech.tokoin.io.com/api/v2/" + suffix
}

func randomDomain(n int) []string {
	domains := []string{}
	for i := 0; i < n; i++ {
		name := randomName(3)
		domains = append(domains, fmt.Sprint(name, ".com"))
	}
	return domains
}

func randomTag(n int) []string {
	tags := []string{}
	for i := 0; i < n; i++ {
		name := randomName(3)
		tags = append(tags, name)
	}
	return tags
}
