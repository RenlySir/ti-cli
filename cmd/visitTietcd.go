package cmd

import (
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func connectEtcd() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"172.16.7.138:2400", "172.16.7.138:22400", "172.16.7.138:32400"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()
}
