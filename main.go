package main

import (
	"./web"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sentryDSN := os.Getenv("SENTRY_DSN")

	if sentryDSN == "" {
		log.Fatal("센트리 설정이 추가되지 않았습니다")
	}

	sentryErr := raven.SetDSN(sentryDSN)

	if sentryErr != nil {
		log.Fatal("Error loading .env file")
	}

	webPort := os.Getenv("WEB_PORT")

	if webPort == "" {
		log.Fatal("실행할 포트가 설정되지 않았습니다")
	}

	router := web.NewRouter()

	fmt.Println(webPort + "로 시작되었습니다")

	log.Fatal(http.ListenAndServe(":"+webPort, router))
}
