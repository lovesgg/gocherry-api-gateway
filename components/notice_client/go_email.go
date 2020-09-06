package notice_client

import (
	"gocherry-api-gateway/admin/services"
	"gopkg.in/gomail.v2"
	"strconv"
)

/**
发送邮件告警
notice_client.EmailSend([]string{"9932851@qq.com"}, "333", "333")
*/
func EmailSend(mailTo []string, subject string, body string) bool {
	config := services.GetAppConfig()
	mailConn := map[string]string{
		"user": config.Common.EmailUser,
		"pass": config.Common.EmailPass,
		"host": config.Common.EmailHost,
		"port": config.Common.EtcdPort,
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "XX官方")) //这种方式可以添加别名，即“XX官方”
	m.SetHeader("To", mailTo...)                                   //发送给多个用户
	m.SetHeader("Subject", subject)                                //设置邮件主题
	m.SetBody("text/html", body)                                   //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	if err != nil {
		return false
	} else {
		return true
	}
}
