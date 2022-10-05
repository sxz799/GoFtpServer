package main

import (
	"fmt"
	"goftp.io/server/core"
	"goftp.io/server/driver/file"
	"log"
	"net"
)

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}
	return ip
}
func GetLocalIP() {
	interfaces, _ := net.Interfaces()
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			fmt.Println("ftp地址：ftp://" + ip.String())
		}
	}
}

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
	GetLocalIP()
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("服务启动失败:", err)
	}

}
