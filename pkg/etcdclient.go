package pkg
import (
    "time"
    "strconv"
    "github.com/coreos/etcd/client"
    "github.com/golang/glog"
)
func Etcdclient(ip string, port int ) client.KeysAPI {
        cfg := client.Config{
        Endpoints:               []string{"http://"+ip+":"+strconv.Itoa(port)},
        Transport:               client.DefaultTransport,
        // set timeout per request to fail fast when the target endpoint is unavailable
        HeaderTimeoutPerRequest: time.Second,
    }
    c, err := client.New(cfg)
    if err != nil {
        glog.Fatal(err)
    }
    kapi := client.NewKeysAPI(c)
    return  kapi
}