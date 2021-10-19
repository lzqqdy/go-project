package notice

import (
	"strconv"

	"go-project/pkg/logger"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

// SendMail 外部调用发送邮件方法
func SendMail(address string, sub string, msg string) {
	body := `
		<html>
		<body>
		<div style="min-height:550px; padding: 100px 55px 200px;">` + msg + `</div>
		</body>
		</html>
		`
	err := coreSendMail(address, sub, body)
	if err != nil {
		logger.Logger("send.email").Error(err)
	}
}

// coreSendMail 邮件发送核心方法
func coreSendMail(mailTo string, subject string, body string) error {
	//定义邮箱服务器连接信息
	mailConn := map[string]string{
		"user":     viper.GetString("rtmp.user"),
		"pass":     viper.GetString("rtmp.pass"),
		"host":     viper.GetString("rtmp.host"),
		"port":     viper.GetString("rtmp.port"),
		"nickname": viper.GetString("rtmp.nickname"),
	}
	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(mailConn["user"], mailConn["nickname"]))
	m.SetHeader("To", mailTo)       //发送用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err
}
