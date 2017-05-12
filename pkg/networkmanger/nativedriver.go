package networkmanger

import (
	"fmt"
	"strings"
    "github.com/liuliuzi/k8s-neutron/pkg/util/cli"
    "github.com/liuliuzi/k8s-neutron/pkg/types"

)

const brctlPath       = "/sbin/brctl"

type NativeNetworkManger struct {

}

func (nnm NativeNetworkManger)Add(networkconfig types.Network) error {
    fmt.Println("start create br ",networkconfig.Br)
    err:=nnm.createBridge(networkconfig.Br)
    if err!=nil{
    	return err
    }
    return nil
}

func (nnm NativeNetworkManger)Show(networkconfig types.Network) error {
    fmt.Println(networkconfig)
    fmt.Println("test")
    return nil
}

func (nnm NativeNetworkManger)createBridge(brName string) error {
	// create bridge if necessary
    err := nnm.ensureBridge(brName)
    if err != nil {
        return fmt.Errorf("failed to create bridge %q: %v", brName, err)
    }
    return nil
}

func (nnm NativeNetworkManger)ensureBridge(brName string) error {
	if brName==""{
		return fmt.Errorf("no br name")
	}
	out, err := cli.New().Command(brctlPath, "show").CombinedOutput()
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
        _, err := cli.New().Command(brctlPath, "addbr",brName).CombinedOutput()
        if err != nil {
            return fmt.Errorf("failed to ls bridge %q: %v", out, err)
        }
    }
    return nil
}

