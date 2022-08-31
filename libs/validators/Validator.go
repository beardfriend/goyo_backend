package validators

import "regexp"

func ValidateRegex(regex, value string) bool {
	reg := regexp.MustCompile(regex)
	return reg.Match([]byte(value))
}

const (
	UrlRegex                  = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
	PhoneRegex                = `^01([0|1|6|7|8|9])([0-9]{3,4})([0-9]{4})$`
	ForbiddenSpecialCharRegex = "([^\"#%'()+/:;<=>?\\[\\]^{|}~]+)$"
	AtLeastOneCharOneNumRegex = "^(?:[0-9~!@$%^&*]+[a-zA-Z!@~!@$%^&*]|[a-zA-Z~!@$%^&*]+[0-9~!@$%^&*])[a-zA-Z0-9~!@$%^&*]*$"
)
