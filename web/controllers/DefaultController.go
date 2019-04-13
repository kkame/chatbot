package controllers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "준비중입니다")
	fmt.Println("로그 준비중입니다")

}
