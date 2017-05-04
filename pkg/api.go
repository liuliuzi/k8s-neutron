package pkg
import (
    "github.com/coreos/etcd/client"
    "fmt"

    
)

type ApiRuntime struct {
    Root           string

    Etcdclient     client.KeysAPI

}
func NewApiRuntime() (*ApiRuntime, error) {
    fmt.Println("NewApiRuntime123")
    apiRuntime := &ApiRuntime{
        Root:           "root2",
        //NodeManager:    &NodeManager{"root1",nil},
    }
    
    fmt.Println(apiRuntime.Root)
    return apiRuntime, nil
}


