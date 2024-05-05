package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func main() {
	// Firebase Admin SDK 초기화
	opt := option.WithCredentialsFile("serviceAccountKey.json") // Service Account Key JSON 파일 경로
	config := &firebase.Config{
		ProjectID: " ", // 여기에 실제 프로젝트 ID를 입력하세요
	}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// FCM 클라이언트 가져오기
	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error getting FCM client: %v\n", err)
	}

	// 메시지 생성
	message := &messaging.MulticastMessage{
		Tokens: []string{" "}, // 대상 기기의 FCM 토큰들
		Notification: &messaging.Notification{
			Title: "제목",
			Body:  "내용",
		},
	}

	// 메시지 전송
	response, err := client.SendMulticast(context.Background(), message)
	if err != nil {
		log.Fatalf("error sending multicast message: %v\n", err)
	}

	// 전송 결과 출력
	fmt.Println("Successfully sent multicast message:", response)
}
