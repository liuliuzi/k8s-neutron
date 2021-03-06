package main

import (
	"fmt"
	"os"
	"runtime"
	//"github.com/spf13/pflag"
	//"github.com/golang/glog"
	"github.com/urfave/cli"
	"github.com/liuliuzi/k8s-neutron/cmd/server/appserver"


)
func main() {
	fmt.Println("bbbbbbbbbbbbbbbbbb")
	runtime.GOMAXPROCS(runtime.NumCPU())
    app := cli.NewApp()
    app.Flags = []cli.Flag {
	    cli.StringFlag{
	      Name: "etcd",
	      Value: "192.168.1.1:4000",
	      Usage: "etcd url",
	    },
	}

    app.Action = appserver.Run

    if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}