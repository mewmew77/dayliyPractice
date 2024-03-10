package smtp

import (
	"net/smtp"
	"strconv"
)

// Sender 发件人邮箱结构体
type Sender struct {
	MailAddr string //发件人邮箱地址
	Passwd   string //邮箱密码，授权码，126/163邮箱需要登录邮箱获取
	Name     string //发件人昵称
	Host     string //邮箱服务器
	Port     int    //服务端口
}

func sendEmail(sender Sender, receiver []string) error {
	auth := smtp.PlainAuth("", sender.MailAddr, sender.Passwd, sender.Host)
	msg := []byte("Subject: 标题测试\r\n\r\n" + "正文测试\r\n")
	err := smtp.SendMail(sender.Host+":"+strconv.Itoa(sender.Port), auth, sender.MailAddr, receiver, msg)
	if err != nil {
		return err
	}
	return nil
}
