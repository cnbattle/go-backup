package main

import (
	"fmt"
	"github.com/cnbattle/go-backup/core/config"
	"github.com/cnbattle/go-backup/core/mail"
	"github.com/cnbattle/go-backup/core/zip"
	"os"
	"time"
)

func main() {
	var attach []string
	for i := range config.Cfg.Back.Folders {
		dest, err := zip.Zip(config.Cfg.Back.Folders[i])
		if err != nil {
			panic(err)
		}
		attach = append(attach, dest)
	}
	attachAll := append(attach, config.Cfg.Back.Files...)
	subject := time.Now().String() + "的备份"
	SendMail(subject, "详细见附件", attachAll)
	for i := range attach {
		_ = os.Remove(attach[i])
	}

	fmt.Println("success")
}

func SendMail(subject string, body string, attach []string) {
	err := mail.SendMail(config.Cfg.Mail.To, subject, body, attach)
	if err != nil {
		panic(err)
	}
}
