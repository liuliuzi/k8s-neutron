package pod

import (
    //"fmt"
    "encoding/json"
    "github.com/liuliuzi/k8s-neutron/pkg/api/networkmange"
)
type Pod struct {
	Id, Name ,HostName, Status, Timestamp string
	Networks [] networkmange.Network

}




func (n Pod) String() string {
	j, _ := json.Marshal(n)
	return string(j)
}
