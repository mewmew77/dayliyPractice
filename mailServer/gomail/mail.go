package gomail

import (
	"errors"
	"gopkg.in/gomail.v2"
	"net"
)

// Sender 发件人邮箱结构体
type Sender struct {
	MailAddr string //发件人邮箱地址
	Passwd   string //邮箱密码，授权码，126/163邮箱需要登录邮箱获取
	Name     string //发件人昵称
	Host     string //邮箱服务器
	Port     int    //服务端口
}

// MessageInfo 邮件内容结构体
type MessageInfo struct {
	Subject string //邮件主题
	Content string //邮件内容
}

func sendMail(sender Sender, targetMailAddr []string, message MessageInfo) error {
	_, err := net.LookupHost(sender.Host)
	if err != nil {
		return errors.New("wrong host")
	}
	m := gomail.NewMessage()
	//m.SetHeader("From",sender.MailAddr) //发送方没有昵称
	m.SetHeader("From", m.FormatAddress(sender.MailAddr, sender.Name)) //设置发送方
	m.SetHeader("To", targetMailAddr...)                               //发送给多个用户
	m.SetHeader("Subject", message.Subject)                            //设置邮件主题
	m.SetBody("text/html", message.Content)                            //设置邮件正文

	d := gomail.NewDialer(sender.Host, sender.Port, sender.MailAddr, sender.Passwd)
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err = d.DialAndSend(m)
	return err
}
