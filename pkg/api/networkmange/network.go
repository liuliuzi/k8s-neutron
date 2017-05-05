package networkmange

import (
    //"fmt"
    "encoding/json"
)

type networkType int
type Status int
type BrType int

const (
    Direct networkType = iota // value --> 0
    Vlan              // value --> 1
    Vxlan            // value --> 2
    Gre           // value --> 3
)

const (
    Init Status = iota
    Connect
    Erroe
)

const (
    Ovs BrType = iota
    native
)


func (this networkType) String() string {
    switch this {
    case Direct:
        return "Direct"
    case Vlan:
        return "Vlan"
    case Vxlan:
        return "Vxlan"
    case Gre:
        return "Gre"
    default:
        return "Unknow"
    }
}

type Network struct {
	Id, Name ,HostName, Timestamp,Br string
	Status Status
	BrType BrType
	Type  networkType
	argu string        // Direct  null
	                   // Vlan    segment ID
		               // Vxlan   vxlan ID
					   // Gre     remotIP list

}


func (n Network) String() string {
	j, _ := json.Marshal(n)
	return string(j)
}
