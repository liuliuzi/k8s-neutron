package types

import (
    //"fmt"
    "encoding/json"
    //"github.com/liuliuzi/k8s-neutron/pkg/types"
)
type Pod struct {
	Id, Name ,HostName, Status, Timestamp string
	Networks [] Network

}


func (n Pod) String() string {
	j, _ := json.Marshal(n)
	return string(j)
}
