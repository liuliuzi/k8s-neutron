package networkmange

import (
	"fmt"
	"github.com/docker/docker/client"
	"context"
    "github.com/docker/docker/api/types"
    "net/http"
    "bytes"
    "io/ioutil"
    "github.com/liuliuzi/k8s-neutron/pkg/api/pod"
    "github.com/liuliuzi/k8s-neutron/pkg/api/networkmange"

)


type OvsNetworkManger struct {

}

func (nm OvsNetworkManger)Add(networkconfigChan chan networkmange.Network){
    fmt.Println(<-networkconfigChan)
	fmt.Println("test")
}

func (nm OvsNetworkManger)show(networkconfig networkmange.Network){
    fmt.Println(<-networkconfigChan)
    fmt.Println("test")
}