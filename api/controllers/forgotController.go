// forgotController.go
package controllers

import (
	"auth-api/database"
	"auth-api/models"
	"fmt"
	"math/rand"
	"net/smtp"

	"github.com/gofiber/fiber/v2"
)

func Forgot(c *fiber.Ctx) error {
	var data map[string]string

	// リクエストデータをパースする
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	token := RandStringRunes(12)
	passwordReset := models.PasswordReset{
		Email: data["email"],
		Token: token,
	}

	// DBに保存
	database.DB.Create(&passwordReset)

	// SMTPメール送信
	from := "selfnote-owner@yahoo.co.jp"
	to := []string{
		data["email"],
	}
	sendFrom := fmt.Sprintf("From: %s\n", from)
	subject := fmt.Sprintf("Subject; %s\n", "Password Reset")
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	// Vue.jsのアドレス
	url := "http://localhost:8080/reset/" + token
	message := fmt.Sprintf("Click <a href=\"%s\">here</a> to reset password!", url)
	err := smtp.SendMail(
		"smtp:1025", // コンテナサービス名+port
		nil,
		from,
		to,
		[]byte(sendFrom+subject+mime+message),
	)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "SUCCESS",
	})
}

// ランダム文字列を返す関数
func RandStringRunes(n int) string {
	var lettersRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = lettersRunes[rand.Intn(len(lettersRunes))]
	}
	return string(b)
}
