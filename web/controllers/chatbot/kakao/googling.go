package kakao

import "net/url"

func googling(sender string, msg string) (responseText string) {

	if len(msg) > 8 && msg[:7] == "@구글" {

		responseText = "https://www.google.com/search?q=" + url.QueryEscape(msg[8:])
	}

	return responseText
}
