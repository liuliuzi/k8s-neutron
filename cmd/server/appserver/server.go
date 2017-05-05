package appserver

import (
	"fmt"
//	"time
    "net/http"
	"github.com/urfave/cli"
    "github.com/emicklei/go-restful"
    "github.com/emicklei/go-restful/swagger"
    "github.com/liuliuzi/k8s-neutron/pkg/api/networkmange"
    "github.com/liuliuzi/k8s-neutron/pkg/api/pod"
    "github.com/liuliuzi/k8s-neutron/pkg"


)

//var apiRuntime,_=pkg.NewApiRuntime()

func Run (c *cli.Context) error {
    fmt.Println("aaaaaaaaaaaaaaaaaaaa")
    var apiRuntime,_=pkg.NewApiRuntime()
    name := "Nefertiti"
    if c.NArg() > 0 {
      name = c.Args().Get(0)
    }
    fmt.Println(c.String("etcd"),name)


    var err error
    apiRuntime.Etcdclient, err = pkg.Etcdclient("10.140.163.102",2379)
    fmt.Println(apiRuntime.Etcdclient)

    if err!=nil {
        fmt.Println("etcd connecet err : %s", err)
        return err
    } else {
        fmt.Println("etcd conneceted")
    }

    ns := networkmange.NetworkService{"/k8s-neutron/networks/"}
    ns.Register()
    ps := pod.PodService{map[string]pod.Pod{},apiRuntime,"/k8s-neutron/pods/"}
    ps.Register()


    config := swagger.Config{
        WebServices:    restful.RegisteredWebServices(), // you control what services are visible
        WebServicesUrl: "http://10.140.163.102:8090",
        ApiPath:        "/apidocs.json",

        // Optionally, specifiy where the UI is located
        SwaggerPath:     "/apidocs/",
        SwaggerFilePath: "/Users/emicklei/Projects/swagger-ui/dist"}
    swagger.InstallSwaggerService(config)

    return http.ListenAndServe(":8090", nil)


}




