package service
//
//import (
//	//"github.com/gogf/gf/net/ghttp"
//	"net/smtp"
//	"strings"
//)
//
//
//var SendmailService = Sendmail{}
//
//type Sendmail struct {}
//
////邮件发送函数
//
//func SendToMail( to, subject,body string)error  {
//	host := "smtpdm.aliyun.com:80"
//	user := "bigw.greatld.com"
//	password := "qwer1234"
//	senduserName := "tmpdep"
//	mailtype := "html"
//
//	hp := strings.Split(host,":")
//	auth := smtp.PlainAuth("",user,password,hp[0])
//	var content_type string
//	if mailtype == "html" {
//		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
//	}else {
//		content_type = "Content-Type: text/plain" + "; chartset=UTF-8"
//	}
//	msg := []byte("To: " + to + "\r\nFrom: " + senduserName + "<"+user+">" +"\r\nSubject: " +subject+ "\r\n" + content_type + "\r\n\r\n" + body)
//	send_to := strings.Split(to, ";")
//	err := smtp.SendMail(host, auth, user, send_to, msg)
//	return err
//
//}
//
//func (m Sendmail) SendService() error {
//
//}