package networkmanger
import (
	"fmt"
	"strings"
    //"github.com/liuliuzi/k8s-neutron/pkg/api/networkmange"
    "github.com/liuliuzi/k8s-neutron/pkg/util/cli"
    "github.com/liuliuzi/k8s-neutron/pkg/types"

)

const ovsPath       = "/usr/bin/ovs-vsctl"
const pipeworkPath  = "/usr/bin/pipework"

type OvsNetworkManger struct {

}

func (nm OvsNetworkManger)Add(networkconfig types.Network) error {
    fmt.Println("start create br ",networkconfig.Br)
    err:=nm.createBridge(networkconfig.Br)
    if err!=nil{
    	return err
    }
    return nil
}

func (nm OvsNetworkManger)Show(networkconfig types.Network) error {
    fmt.Println(networkconfig)
    return nil
}

func (nm OvsNetworkManger)createBridge(brName string) error {
	// create bridge if necessary
    err := nm.ensureBridge(brName)
    if err != nil {
        return fmt.Errorf("failed to create bridge %q: %v", brName, err)
    }
    return nil
}

func (nm OvsNetworkManger)ensureBridge(brName string) error {
	if brName==""{
		return fmt.Errorf("no br name")
	}
	out, err := cli.New().Command(ovsPath, "show").CombinedOutput()
    if err != nil {
        return fmt.Errorf("failed to ls bridge %q: %v", out, err)
    }
    if string(out) != "" {
        outlines:=strings.Split(string(out), "\n")
        for _, numline := range outlines {
            if strings.Contains(numline, brName){
                return nil
            }
        }
        _, err := cli.New().Command(ovsPath, "add-br",brName).CombinedOutput()
        if err != nil {
            return fmt.Errorf("failed to ls bridge %q: %v", out, err)
        }
    }
    return nil
}

func (nm OvsNetworkManger)addVeth(brName string, containerID string, network types.Network, argu string)  error {
    if network.Type==types.Direct{
        return fmt.Errorf("unimplement network type%s", network.Type)
        //out, err := New().Command(pipeworkPath, brName, containerID,ip).CombinedOutput()
    }else if network.Type==types.Vlan{
        return fmt.Errorf("unimplement network type%s", network.Type)
        //out, err := New().Command(pipeworkPath, brName, containerID,ip,"@"+vlanTag).CombinedOutput()
    }else if network.Type==types.Vxlan{
        return fmt.Errorf("unimplement network type%s", network.Type)
        //out, err := New().Command(pipeworkPath, brName, containerID,ip,"@"+vlanTag).CombinedOutput()
    }else if network.Type==types.Gre{
        return fmt.Errorf("unimplement network type%s", network.Type)
        //out, err := New().Command(pipeworkPath, brName, containerID,ip,"@"+vlanTag).CombinedOutput()
    }else{
        return fmt.Errorf("unspport network type %s", network.Type)
    }

/*
    if err != nil {
        return fmt.Errorf("failed to add veth %q: %v", out, err)
    }*/
    return  nil
}
