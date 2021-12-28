package main

/*
  @Author : lanyulei
  @Desc : 发送邮件
*/

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/spf13/viper"

	"gopkg.in/gomail.v2"
)

func server(mailTo []string, ccTo []string, subject, body string, args ...string) error {
	//定义邮箱服务器连接信息，如果是网易邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": "dmmo@mail.com",
		"pass": "",
		"host": "smtp.mail.com",
		"port": "25",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], viper.GetString("工单完结"))) //这种方式可以添加别名，即“XX官方”
	m.SetHeader("To", mailTo...)                                                    //发送给多个用户
	m.SetHeader("Cc", ccTo...)                                                      //发送给多个用户
	m.SetHeader("Subject", subject)                                                 //设置邮件主题
	m.SetBody("text/html", body)                                                    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)
	return err

}

func main() {
	mailTo := []string{"dm-tech@mail.com"}
	ccTo := []string{"zhaosan@mail.com", "zhaowu@mail.com", "lisi@mail.com"}

	subject := flag.String("s", "subject", "主题")
	body := flag.String("b", "body", "收件内容")
	flag.Parse()
	err := server(mailTo, ccTo, *subject, *body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("send successfully")
}
