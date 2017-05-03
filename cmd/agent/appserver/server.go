package appserver

import (
	"fmt"
	"time"
	"github.com/urfave/cli"
	"k8s.io/kubernetes/pkg/util/wait"
	"github.com/liuliuzi/k8s-neutron/pkg/agent"
)


func Run (c *cli.Context) error {
    name := "Nefertiti"
    if c.NArg() > 0 {
      name = c.Args().Get(0)
    }
    fmt.Println(c.String("etcd"),name)

    startAgent()
    select{ }
    return nil
}

func startAgent(){
	//go wait.Until(func() { k.Run(podCfg.Updates()) }, 0, wait.NeverStop)
	go wait.Until(func() {  agent.SyncNetwork() }, 5*time.Second, wait.NeverStop)
}


