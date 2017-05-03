package pod

import (
    //"fmt"
    "encoding/json"
)
type Pod struct {
	Id, Name ,HostName, Status, Timestamp string

}


func (n Pod) String() string {
	j, _ := json.Marshal(n)
	return string(j)
}
