package smtp

import "testing"

func TestSendMail(t *testing.T) {
	sender := Sender{
		MailAddr: "xxx@126.com",
		Passwd:   "xxxxxxx",
		Name:     "xx",
		Host:     "smtp.126.com",
		Port:     25,
	}
	target := []string{"xxx@126.com"}

	err := sendEmail(sender, target)
	if err != nil {
		t.Fatal("邮件发送失败！！！", err)
	}
}
