package common

import (
	"strings"
)

const replacementChar = "*"

var badwords = []string{
	"chink",
	"cipa",
	"coon",
	"dyke",
	"fag",
	"fagging",
	"faggitt",
	"faggot",
	"faggs",
	"fagot",
	"fagots",
	"fags",
	"fuck",
	"hoar",
	"hoare",
	"hoer",
	"homo",
	"hore",
	"kike",
	"n1gga",
	"n1gger",
	"nigg3r",
	"nigg4h",
	"nigga",
	"niggah",
	"niggas",
	"nigger",
	"niggers",
	"whoar",
	"whore",
}

func ContainsCurseWords(str string) bool {

	if str == "" {
		return false
	}
	str = strings.ToLower(str)
	for _, word := range badwords {
		if strings.Contains(str, word) {
			return true
		}
	}
	return false
}
