package main

import (
	"context"
	"fmt"

	"github.com/ZaneXie/gen/examples/biz"
	"github.com/ZaneXie/gen/examples/conf"
	"github.com/ZaneXie/gen/examples/dal"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()
}

func main() {
	// start your project here
	fmt.Println("hello world")
	defer fmt.Println("bye~")

	biz.Query(context.Background())
}
