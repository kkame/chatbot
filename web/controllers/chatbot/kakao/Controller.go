package kakao

import (
	"fmt"
	"net/http"
	"strings"
)

func Post(w http.ResponseWriter, r *http.Request) {

	responseText := ""

	keys := r.URL.Query()

	//isGroupChat := keys.Get("isGroupChat")
	room := keys.Get("room")
	sender := keys.Get("sender")
	msg := keys.Get("msg")
	msg = strings.TrimSpace(msg)

	responseText = googling(sender, msg)

	if responseText == "" {

		if room == "laravel" {

			responseText = laravel(sender, msg)
		}
	}

	//ImageDB := keys.Get("ImageDB")
	//fmt.Println(ImageDB)

	_, err := fmt.Fprintln(w, responseText)

	if err != nil {

		fmt.Println(err)
	}
}
