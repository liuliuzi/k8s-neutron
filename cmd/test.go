package main

import "fmt"
import "encoding/json"
import "github.com/liuliuzi/k8s-neutron/pkg/api/pod"

func main() {

	keysBody:=[]byte(`[{ "Id": "192.168.1.12","Name": "","HostName": "","Status": "","Timestamp": "","Networks": null  },  {"Id": "192.168.1.13","Name": "","HostName": "","Status": "","Timestamp": "","Networks": [{"Id": "192.168.1.8","Name": "","HostName": "","Status": "","Timestamp": ""},{"Id": "192.168.1.10","Name": "","HostName": "","Status": "","Timestamp": ""}]}]`)
    
    pods := make([]pod.Pod,0)
    json.Unmarshal(keysBody, &pods)
    fmt.Printf("%#v", pods)
}