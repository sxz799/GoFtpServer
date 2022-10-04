package main

import (
	"fmt"
	"goftp.io/server/core"
	"goftp.io/server/driver/file"
	"log"
)

func main() {

	username := "ftp"
	password := "ftp"
	var a, b string
	fmt.Println("请输入FTP用户名(默认: ftp )：")

	fmt.Scanln(&a)
	if len(a) < 1 {
		fmt.Println("未输入用户名，将使用默认用户名 ftp")
	}
	fmt.Println()
	fmt.Println("请输入 FTP  密码   (默认: ftp )：")

	fmt.Scanln(&b)
	if len(b) < 1 {
		fmt.Println("未输入密码，将使用默认密码 ftp")
	}

	Name := "Go FTP Server"
	rootPath := "./" //FTP根目录
	Port := 21       //FTP 端口
	var perm = core.NewSimplePerm("goFtp", "goFtp")

	// Server options without hostname or port
	opt := &core.ServerOpts{
		Name: Name,
		Factory: &file.DriverFactory{
			RootPath: rootPath,
			Perm:     perm,
		},
		Auth: &core.SimpleAuth{
			Name:     username, // FTP 账号
			Password: password, // FTP 密码
		},
		Port: Port,
	}
	// start ftp server
	s := core.NewServer(opt)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
