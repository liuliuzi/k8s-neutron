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

    //"golang.org/x/net/context"

)


func Run (c *cli.Context) error {
    fmt.Println("aaaaaaaaaaaaaaaaaaaa")
    name := "Nefertiti"
    if c.NArg() > 0 {
      name = c.Args().Get(0)
    }
    fmt.Println(c.String("etcd"),name)

    ns := networkmange.NetworkService{"/k8s-neutron/networks/"}
    ns.Register()
    ps := pod.PodService{"/k8s-neutron/pods/"}
    ps.Register()


    config := swagger.Config{
        WebServices:    restful.RegisteredWebServices(), // you control what services are visible
        WebServicesUrl: "http://localhost:8090",
        ApiPath:        "/apidocs.json",

        // Optionally, specifiy where the UI is located
        SwaggerPath:     "/apidocs/",
        SwaggerFilePath: "/Users/emicklei/Projects/swagger-ui/dist"}
    swagger.InstallSwaggerService(config)

    return http.ListenAndServe(":8090", nil)

    /*
    Etcdclient:=pkg.Etcdclient("127.0.0.1",4001)
    resp, err :=Etcdclient.Set(context.Background(), "cc", "123value", nil)
    if err != nil {
        fmt.Println("storageSet failed")
        fmt.Println(err)
    }else{
        fmt.Println("storageSet success")
        fmt.Println(resp)
    }

*/
    //return nil
}




