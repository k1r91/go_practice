package main

import (
	"fmt"
	"net"
	 iofs "io/fs"
	"github.com/hirochachacha/go-smb2"
)

func main() {
	conn, err := net.Dial("tcp", "nw-dfs-efs02.megafon.ru:445")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     "AFSMBFOSARGO",
			Password: "gg5x/0a=YXnyAX3=`A/7",
			Domain: "megafon.ru",
		},
	}

	s, err := d.Dial(conn)
	if err != nil {
		panic(err)
	}
	defer s.Logoff()
	names, err := s.ListSharenames()
	if err != nil {
		panic(err)
	}

	for _, name := range names {
		fmt.Println(name)
	}
	fs, err := s.Mount(`Disk1$`)
	if err != nil {
		panic(err)
	}
	err = iofs.WalkDir(fs.DirFS("АО Апатит"), ".", func(path string, d iofs.DirEntry, err error) error {
		fmt.Println(path, d, err)
		return nil
	})
	if err != nil {
		panic(err)
	}
}