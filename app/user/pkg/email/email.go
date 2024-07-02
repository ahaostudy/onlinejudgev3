package email

import (
	"fmt"
	"github.com/ahaostudy/onlinejudge/app/user/conf"
	"net/smtp"
	"regexp"

	"github.com/jordan-wright/email"
)

func Send(subject, html string, toEmails ...string) error {
	confEmail := conf.GetConf().Email
	e := email.NewEmail()

	e.From = confEmail.From
	e.To = toEmails
	e.Subject = subject
	e.HTML = []byte(html)
	auth := smtp.PlainAuth("", confEmail.Email, confEmail.Auth, confEmail.Host)

	return e.Send(confEmail.Addr, auth)
}

func SendCaptcha(captcha string, toEmails ...string) error {
	subject := "【OnlineJudge】邮箱验证"
	html := fmt.Sprintf(`<div style="text-align: center;">
		<h2 style="color: #333;">欢迎使用，你的验证码为：</h2>
		<h1 style="margin: 1.2em 0;">%s</h1>
		<p style="font-size: 12px; color: #666;">请在5分钟内完成验证，过期失效，请勿告知他人，以防个人信息泄露</p>
	</div>`, captcha)
	return Send(subject, html, toEmails...)
}

func ExtractUsernameFromEmail(email string) (string, bool) {
	pattern := `([^@]+)@`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(email)
	if len(match) == 2 {
		return match[1], true
	} else {
		return "", false
	}
}
