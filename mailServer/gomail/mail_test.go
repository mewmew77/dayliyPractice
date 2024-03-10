package gomail

import "testing"

func TestSendMail(t *testing.T) {
	sender := Sender{
		MailAddr: "xxx@126.com",
		Passwd:   "xxxxxx",
		Name:     "xx",
		Host:     "smtp.126.com",
		Port:     25,
	}
	target := []string{"xx@126.com"}
	msg := MessageInfo{
		Subject: "gomail测试",
		Content: "测试 gomail 实现的邮件发送",
	}
	err := sendMail(sender, target, msg)
	if err != nil {
		t.Fatal("邮件发送失败！！！", err)
	}
}
