package appserver

import (
	"fmt"
	"time"
	"github.com/urfave/cli"
	//"k8s.io/kubernetes/pkg/util/wait"
    "github.com/liuliuzi/k8s-neutron/pkg/util/wait"
	"github.com/liuliuzi/k8s-neutron/pkg/agent"

    "github.com/docker/docker/client"
    //"github.com/liuliuzi/k8s-neutron/pkg/api/pod"
    "github.com/liuliuzi/k8s-neutron/pkg/api/networkmange"
    //"encoding/json"
)




func Run (c *cli.Context) error {
    name := "Nefertiti"
    if c.NArg() > 0 {
      name = c.Args().Get(0)
    }
    fmt.Println(c.String("etcd"),name)
/*
    keysBody:=[]byte(`[{ "Id": "192.168.1.12","Name": "","HostName": "","Status": "","Timestamp": "","Networks": null  },  {"Id": "192.168.1.13","Name": "","HostName": "","Status": "","Timestamp": "","Networks": [{"Id": "192.168.1.8","Name": "","HostName": "","Status": "","Timestamp": ""},{"Id": "192.168.1.10","Name": "","HostName": "","Status": "","Timestamp": ""}]}]`)
    
    pods := make([]pod.Pod,0)
    json.Unmarshal(keysBody, &pods)
    fmt.Printf("%#v", pods)
/////////////////*/
    dc, err := client.NewEnvClient()
    if err != nil {
        panic(err)
    }

    startAgent(dc)
    select{ }
    return nil
}

func startAgent(dc *client.Client){
    networkconfigChan := make(chan networkmange.Network)
    //nm :=new agent.NetworkManger
    go wait.Until(func() { agent.Update(networkconfigChan) }, 0, wait.NeverStop)
	go wait.Until(func() { agent.SyncNetwork(dc,networkconfigChan) }, 5*time.Second, wait.NeverStop)
}



