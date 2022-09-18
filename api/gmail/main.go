package gmail

import (
	config "api/setting"
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

type GAmll struct {
	Email    string
	Password string
	Key      smtp.Auth
}

func (g *GAmll) Login(Email, Password string) {
	g.Email = Email
	g.Key = smtp.PlainAuth("", Email, Password, "smtp.gmail.com")

}

func (g *GAmll) LoginCon(StructData *config.Data_Config) {
	g.Email = StructData.Senderemail.Email
	g.Key = smtp.PlainAuth("", g.Email, StructData.Senderemail.Password, "smtp.gmail.com")

}
func (g GAmll) SEndlogin(Username, tag, otp, sendto string) {
	t, err := template.ParseFiles("template/template.html")
	buf := new(bytes.Buffer)
	data := struct {
		Name string
		OTP  string
		Tag  string
	}{
		Name: Username,
		OTP:  otp,
		Tag:  tag,
	}
	if er := t.Execute(buf, data); er != nil {

	}
	if err != nil {
		fmt.Printf("smtp error: %s", err)
		return
	}

	from := g.Email

	to := sendto

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +

		"Subject: Hello there\n" + "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" + buf.String()
	r := smtp.SendMail("smtp.gmail.com:587",
		g.Key,
		from, []string{to}, []byte(msg))

	if r != nil {
		fmt.Printf("smtp error: %s", r)
		return
	}

}
func (g GAmll) SEndCAll(Username, tag, LiNK, sendto string) {
	t, err := template.ParseFiles("template/SEnd.html")
	buf := new(bytes.Buffer)
	data := struct {
		Name string
		OTP  string
		Tag  string
	}{
		Name: Username,
		OTP:  LiNK,
		Tag:  tag,
	}
	if er := t.Execute(buf, data); er != nil {
	}
	if err != nil {
		fmt.Printf("smtp error: %s", err)
		return
	}
	from := g.Email
	to := sendto
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +

		"Subject: Hello there\n" + "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" + buf.String()
	r := smtp.SendMail("smtp.gmail.com:587",
		g.Key,
		from, []string{to}, []byte(msg))

	if r != nil {
		fmt.Printf("smtp error: %s", r)
		return
	}

}
