package mail

import (
	"crypto/tls"
	"github.com/cnbattle/go-backup/core/config"
	"gopkg.in/gomail.v2"
)

type conn struct {
	User string
	Pass string
	Host string
	Port int
}

var mailConn *conn

func init() {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn = &conn{
		User: config.Cfg.Mail.User,
		Pass: config.Cfg.Mail.Password,
		Host: config.Cfg.Mail.Host,
		Port: config.Cfg.Mail.Port,
	}
}

// SendMail 发送邮件
func SendMail(mailTo []string, subject string, body string, attach []string) error {
	m := gomail.NewMessage()
	//这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("From", "<"+mailConn.User+">")
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文
	for i := range attach {
		m.Attach(attach[i]) // 设置附件
	}
	d := gomail.NewDialer(mailConn.Host, mailConn.Port, mailConn.User, mailConn.Pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	return err
}
