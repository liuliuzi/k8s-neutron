package networkmange

import (
    //"fmt"
    "encoding/json"
)
type Network struct {
	Id, Name ,HostName, Status, Timestamp string

}


func (n Network) String() string {
	j, _ := json.Marshal(n)
	return string(j)
}
