package main

import (
	"fmt"
	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/table"
)

// CREATE TABLE test(c1 bigint, c2 bigint, PRIMARY KEY(c1)) PARTITION BY hash(c1) partitions 2;
func main() {
	const (
		configUrl    = "xxx"
		fullUserName = "root@sys#obcluster"
		passWord     = ""
		sysUserName  = "root"
		sysPassWord  = ""
		tableName    = "test"
	)

	cfg := config.NewDefaultClientConfig()
	cli, err := client.NewClient(configUrl, fullUserName, passWord, sysUserName, sysPassWord, cfg)
	if err != nil {
		panic(err)
	}

	err = cli.AddRowkey(tableName, []string{"c1"})
	if err != nil {
		panic(err)
	}
	rowkey := []*table.Column{table.NewColumn("c1", int64(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int64(1))}
	affectRows, err := cli.Insert(
		tableName,
		rowkey,
		mutateColumns,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(affectRows)
}