package main

import (
	"fmt"
	"os"
	"runtime"
	//"github.com/spf13/pflag"
	//"github.com/golang/glog"
	"github.com/urfave/cli"

	//"github.com/liuliuzi/stf/cmd/stf-agent/app"
	//"github.com/liuliuzi/stf/cmd/stf-agent/app/options"
	//"github.com/liuliuzi/stf/pkg/util"

)
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	s := options.NewAgentServer()
	s.AddFlags(pflag.CommandLine)

	//util.InitLogs()
	//defer util.FlushLogs()

	//glog.Info("valid agent app start")

	if err := app.Run(s); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}