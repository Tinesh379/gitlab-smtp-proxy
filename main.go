package main

import (
	"bytes"
	"html/template"
	"log"
	"gopkg.in/gomail.v2"
)
const(
    fromEmailAddr string = "noreply@gmail.com"
)
type Smtp struct{
    server  string
    port int
    user string
    password string
}
func formHtmlTemplate(templatePath string) string {
    t, err := template.ParseFiles(templatePath)
    if err != nil{
        log.Fatal("error parsing template:",err)
    }
    var body bytes.Buffer
    err = t.ExecuteTemplate(&body, "emailTemplate.html",struct{ Name string}{ Name: "Tinesh",})
    if err != nil{
        log.Fatal("error executing the template:", err)
    }
    return body.String()
}
func (s Smtp)sendGoEmail(emailBody string, subject string, attachment string, toList ... string){
    
    m := gomail.NewMessage()
    m.SetHeader("From", fromEmailAddr)
    m.SetHeader("To", toList...)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", emailBody)
    m.Attach(attachment)

d := gomail.NewDialer(s.server, s.port, s.user, s.password)

// Send the email to Bob, Cora and Dan.
if err := d.DialAndSend(m); err != nil {
	panic(err)
}
}
func main(){
    log.Println("Started SMTP Proxy Application")
    subject := "Test Subject"
    attach := "emailTemplate.html"
    emailToList := "tineshbabukatta@gmail.com"
    smtp := Smtp{
        server: "smtp.gmail.com",
        port: 5349,
        user: "tineshbabukatta@gmail.com",
        password: "finalPassword",
    }
    emailBody := formHtmlTemplate("emailTemplate.html")
    smtp.sendGoEmail(emailBody, subject, attach, emailToList)
    
}