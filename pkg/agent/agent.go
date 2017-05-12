package agent

import (
	"fmt"
	"github.com/docker/docker/client"
	"context"
    dockertypes "github.com/docker/docker/api/types"
    "net/http"
    "bytes"
    "io/ioutil"
    //"github.com/liuliuzi/k8s-neutron/pkg/api/pod"
    //"github.com/liuliuzi/k8s-neutron/pkg/api/networkmange"
    "github.com/liuliuzi/k8s-neutron/pkg/types"
    "github.com/liuliuzi/k8s-neutron/pkg/networkmanger"
    "encoding/json"
)


func Update(networkconfigChan chan types.Network){
    //var networkconfig networkmange.Network
    networkconfig:= <- networkconfigChan
	fmt.Println(networkconfig)
    fmt.Println(networkconfig.BrType)
    if networkconfig.BrType==types.Ovs{
        ovsdr:=networkmanger.OvsNetworkManger{}
        err:=ovsdr.Add(networkconfig)
        if err!=nil{
            fmt.Println(err)
        }
    }

}

func SyncNetwork(dc *client.Client ,  networkconfigChan chan types.Network) {
	podConfigList,err:=getPodConfigList()
    fmt.Println(podList)

	//get container list
	containerList, err := dc.ContainerList(context.Background(), dockertypes.ContainerListOptions{})
    if err != nil {
        panic(err)
    }

    for _, container := range containerList {
        fmt.Printf("%s %s\n", container.ID, container.Image)
        for _, podConfig := range podConfigList{
            if container.id==podConfig.id{
                fmt.Printf("find ----------------\n")
                currentNetworkIdList:=getNetworksFromContainer(container.id)
                currentNetworkConfigIdList:=getNetworksFromPod(podConfig.Networks)
                for _,currentNetwork:=range currentNetworks{
                    if 
                }
        }
        networkList=getContainerNetworkList(container.ID)
    }

    

}

func getNetworksForPod(pod types.Pod ,networkconfigChan chan types.Network){
    for _, network := range pod.Networks{
        fmt.Println(network.Id)
        syncNetworkForPod(pod,network,networkconfigChan)
    }
}

func syncNetworkForPod(pod types.Pod, network types.Network,networkconfigChan chan types.Network){
    //judege network is in pod
    //1,judege br exit
    //
    networkconfigChan <- network
}

//func listNetworkForPod(pod pod.Pod, network networkmange.Network){
    
//}



func getPodConfigList() ([]types.Pod, error){
	//get pod list
    ServerUrl:="http://10.140.163.102:8090/pods"
    var querystring = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
    req, err := http.NewRequest("GET", ServerUrl, bytes.NewBuffer(querystring))
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil,err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil,err
    }
    pods := make([]types.Pod,0)
    //fmt.Println(string(body)
    err=json.Unmarshal([]byte(body), &pods)
    if err != nil {
        fmt.Println("Unmarshal failed" )
        return nil,err
    }
    return pods,nil

}