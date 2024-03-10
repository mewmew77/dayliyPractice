gomail - 使用gomail实现邮件发送

```
go get gopkg.in/gomail.v2
```

git地址：  https://github.com/go-gomail/gomail

踩坑参考： https://blog.csdn.net/getyouwant/article/details/97374871

QQ 邮箱：
SMTP 服务器地址：smtp.qq.com（SSL协议端口：465/587, 非SSL协议端口：25）

163 邮箱：
SMTP 服务器地址：smtp.163.com（SSL协议端口：465/994，非SSL协议端口：25）

126 邮箱：
SMTP 服务器地址：smtp.126.com（SSL协议端口：465/994，非SSL协议端口：25）

各类邮箱配置：

```
163邮箱

POP3：pop.126.com

SMTP：smtp.126.com SMTP端口号：25


126邮箱

POP3：pop.126.com    端口：110    SSL端口：995

SMTP：smtp.126.com  端口号：25    SSL协议：465 / 994

IMAP: imap.126.com  端口：143  SSL协议：993


QQ邮箱

POP3：pop.qq.com

SMTP：smtp.qq.com   端口号：25 （使用SSL时，端口号465或587）



新浪
免费邮箱
POP3：pop.sina.com
SMTP：smtp.sina.com
SMTP端口号：25
VIP邮箱
POP3：pop.vip.sina.com
SMTP：smtp.vip.sina.com
SMTP端口号：25

搜狐邮箱
POP3：pop.sohu.com
SMTP：smtp.sohu.com
SMTP端口号：25



雅虎邮箱
POP3：pop.mail.yahoo.cn
SMTP：smtp.mail.yahoo.cn
SMTP端口号：25

TOM邮箱
POP3：pop.tom.com
SMTP：smtp.tom.com
SMTP端口号：25

Gmail邮箱
POP3：pop.gmail.com
SMTP：smtp.gmail.com
SMTP端口号：587 或 25

263邮箱
域名：263.net
POP3：pop.263.net
SMTP：smtp.263.net
SMTP端口号：25
域名：x263.net
POP3：pop.x263.net
SMTP：smtp.x263.net
SMTP端口号：25
域名：263.net.cn
POP3：pop.263.net.cn
SMTP：smtp.263.net.cn
SMTP端口号：25
域名：炫我型
POP3：pop.263xmail.com
SMTP：smtp.263xmail.com
SMTP端口号：25

21CN邮箱
免费邮箱: 
　 POP3：pop.21cn.com 
　 SMTP：smtp.21cn.com
SMTP端口号：25
商务邮箱： 
　 POP3：pop.21cn.net 
　 SMTP：smtp.21cn.net
SMTP端口号：25 
　经济邮箱： 
　 POP3：pop.21cn.com 
　 SMTP：smtp.21cn.com
SMTP端口号：25 
　企业邮箱： 
　 POP3：pop-ent.21cn.com 
　 SMTP：smtp-ent.21cn.com
SMTP端口号：25
快感邮箱: 
　 POP3：VIP.21cn.com
SMTP：vip.21cn.com
SMTP端口号：25

中华网任我邮邮箱
POP3：rwpop.china.com
SMTP：rwsmtp.china.com
SMTP端口号：25

中华网时尚、商务邮箱
POP3：pop.china.com
SMTP：smtp.china.com
SMTP端口号：25
原文链接：https://blog.csdn.net/qq_29180565/article/details/91906635```