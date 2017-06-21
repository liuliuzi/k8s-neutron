package iplink
import (
	"fmt"
	//"strings"
    "github.com/liuliuzi/k8s-neutron/pkg/util/cli"
)

const (
    ippath  = "/sbin/ip"
)

func GetNetList(ns string) error {
	var out []byte
	var err error
	//nsflag:=""
	if ns!=""{
		out, err = cli.New().Command(ippath, "netns","exec" ,ns,"ip","link","show").CombinedOutput()
	}else{
		out, err = cli.New().Command(ippath, "link","show").CombinedOutput()
	}
    if err != nil {
        return fmt.Errorf("failed ip link show e %q: %v", out, err)
    }
    fmt.Println(string(out))
    return nil
    /*
    if string(out) != "" {
        outlines:=strings.Split(string(out), "\n")
        for _, numline := range outlines {
            if strings.Contains(numline, brName){
                return nil
            }
        }*/
}