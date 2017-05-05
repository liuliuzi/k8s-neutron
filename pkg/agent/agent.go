package agent

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
    "github.com/liuliuzi/k8s-neutron/pkg/networkmanger"
    "encoding/json"
)


func Update(networkconfigChan chan networkmange.Network){
    networkconfig<-networkconfigChan
	fmt.Println(networkconfig)
    var networkManger networkmanger
    if networkconfig.BrType==networkmange.Ovs

}

func SyncNetwork(dc *client.Client ,  networkconfigChan chan networkmange.Network) {
	podList,err:=getPodList()
    fmt.Println(podList)

	//get container list
	containers, err := dc.ContainerList(context.Background(), types.ContainerListOptions{})
    if err != nil {
        panic(err)
    }

    for _, container := range containers {
        fmt.Printf("%s %s\n", container.ID, container.Image)
        for _, pod := range podList{
        	//fmt.Printf("%s \n", pod.ID, container.Image)
        	//if pod.Id==container.ID{
        		fmt.Printf("find ----------------\n")
                getNetworksForPod(pod, networkconfigChan)
        	//}
        }
    }

}

func getNetworksForPod(pod pod.Pod ,networkconfigChan chan networkmange.Network){
    for _, network := range pod.Networks{
        fmt.Println(network.Id)
        syncNetworkForPod(pod,network,networkconfigChan)
    }
}

func syncNetworkForPod(pod pod.Pod, network networkmange.Network,networkconfigChan chan networkmange.Network){
    //judege network is in pod
    //1,judege br exit
    //
    networkconfigChan <- network
}

//func listNetworkForPod(pod pod.Pod, network networkmange.Network){
    
//}



func getPodList() ([]pod.Pod, error){
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
    pods := make([]pod.Pod,0)
    //fmt.Println(string(body)
    err=json.Unmarshal([]byte(body), &pods)
    if err != nil {
        fmt.Println("Unmarshal failed" )
        return nil,err
    }
    return pods,nil

}