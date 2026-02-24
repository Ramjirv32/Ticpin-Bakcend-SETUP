package config

import (
"fmt"
"math/rand"
"os"
"strconv"

gomail "gopkg.in/gomail.v2"
)

func sendOTP(from, pass, to, subject, body string) error {
port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
if port == 0 {
port = 587
}

m := gomail.NewMessage()
m.SetHeader("From", from)
m.SetHeader("To", to)
m.SetHeader("Subject", subject)
m.SetBody("text/html", body)

d := gomail.NewDialer("smtp.gmail.com", port, from, pass)
return d.DialAndSend(m)
}

func SendPlayOTP(toEmail, otp string) error {
from := os.Getenv("PLAY_EMAIL")
pass := os.Getenv("PLAY_APP_PASSWORD")
body := fmt.Sprintf("<h2>Your Ticpin Play OTP: <b>%s</b></h2><p>Valid for 10 minutes.</p>", otp)
return sendOTP(from, pass, toEmail, "Ticpin Play OTP Verification", body)
}

func GenerateOTP() string {
return fmt.Sprintf("%06d", rand.Intn(1000000))
}
