package networkmange
import (
	"net/http"
	"github.com/emicklei/go-restful"
	"fmt"
	//"github.com/liuliuzi/k8s-neutron/pkg"
	//"golang.org/x/net/context"
	"strings"
	//"encoding/json"
	//"github.com/bitly/go-simplejson"
	//"github.com/emicklei/go-restful/swagger"
)

type NetworkService struct {
	// normally one would use DAO (data access object)
	//Networks map[string]Network
	//ApiRuntime   *pkg.ApiRuntime
	Prefix       string
}

func (ns NetworkService) Init(){
	ns.Prefix="/stf/network/"
	fmt.Println(ns.Prefix)
}

func (ns NetworkService) Register() {
	ws := new(restful.WebService)
	ws.
		Path("/networks").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well
	ws.Route(ws.GET("/").To(ns.findAllNetworks))
	ws.Route(ws.GET("/{Network-id}").To(ns.findNetwork)) // on the response
	ws.Route(ws.PUT("/{Network-id}").To(ns.updateNetwork)) // from the request
	ws.Route(ws.POST("/").To(ns.createNetwork)) // from the request
	ws.Route(ws.DELETE("/{Network-id}").To(ns.removeNetwork))
	restful.Add(ws)
}

func (ns NetworkService) findAllNetworks(request *restful.Request, response *restful.Response) {
	fmt.Println("findAllNetworks")
	value, _ := ns.storageGetKey()
	ret:="["
	for _, key := range value {
            fmt.Println( key)
            ret=ret+`"`+key+`":`
            key=strings.Replace(key,ns.Prefix,"",20)
            value, _ := ns.storageGet(key)
            //fmt.Println (value)
            ret=ret+value

        }
    ret=ret+"]"

	response.WriteEntity(ret)
}


func (ns NetworkService) findNetwork(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("Network-id")
	value, err := ns.storageGet(id)
	if err != nil {
	    fmt.Println(err)
	}else{
	    fmt.Println(value)
	}

	if len(value) == 0 {
		response.WriteErrorString(http.StatusNotFound, "Network could not be found.")
	} else {
		fmt.Println("findNetwork")
		response.WriteEntity(value)
	}
}

func (u *NetworkService) updateNetwork(request *restful.Request, response *restful.Response) {
	nw := new(Network)
	err := request.ReadEntity(&nw)
	if err == nil {
		//ns.storageUpdate(nw.Id,nw.String())
		fmt.Println("updateNetwork")
		response.WriteEntity(nw)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}


func (u *NetworkService) createNetwork(request *restful.Request, response *restful.Response) {

	nw := new(Network)
	err := request.ReadEntity(&nw)
	if err == nil {
		//ns.storageSet(nw.Id,nw)
		//ns.storageSet(nw.Id,nw.String())
		//ns.storageSet(nw.Id,{'Id':'123566666','Name':'liu'})

		//ns.storageSet(nw.Id,kk)
		response.WriteHeaderAndEntity(http.StatusCreated, nw)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func (u *NetworkService) removeNetwork(request *restful.Request, response *restful.Response) {
	//id := request.PathParameter("Network-id")
	//ns.storageDelete(id)
	fmt.Println("removeNetwork")
}

/*
func (u *NetworkService) storageSet(key string , value string ) error {
	resp, err := ns.ApiRuntime.Etcdclient.Set(context.Background(), ns.Prefix+key, value, nil)
    if err != nil {
    	fmt.Println("storageSet failed")
        fmt.Println(err)
    }else{
    	fmt.Println("storageSet success")
    	fmt.Println(resp)
    }
    return err
}*/

func (u *NetworkService) storageGetKey() ([]string ,error) {
	return  []string{ "str1", "str2", "str3", } , nil
	/*
	resp, err := ns.ApiRuntime.Etcdclient.Get(context.Background(), ns.Prefix, nil)

    if err != nil {
    	fmt.Println("storageGetKey failed")
        fmt.Println(err)
        return nil,err
    }else{
    	ret:=[]string{}
    	fmt.Println("storageGetKey success")
    	for _, Network := range resp.Network.Networks {
            fmt.Println( Network)
            ret=append(ret,Network.Key)
        }

        return ret,nil

    }
    */
}
func (u *NetworkService) storageGet(key string ) (string ,error) {
	fmt.Println("storageGet", key)
	return "aa",nil
	/*
	resp, err := ns.ApiRuntime.Etcdclient.Get(context.Background(), ns.Prefix+key, nil)
    if err != nil {
    	fmt.Println("storageGet failed")
        fmt.Println(err)
    }else{
    	fmt.Println("storageGet success")
    	fmt.Println(resp)
    }
    return resp.Network.Value,err
    */
}
/*

func (u *NetworkService) storageUpdate(key string , value string ) (string ,error) {
	resp, err := ns.ApiRuntime.Etcdclient.Update(context.Background(), ns.Prefix+key,value)
    if err != nil {
    	fmt.Println("storageUpdate failed")
        fmt.Println(err)
    }else{
    	fmt.Println("storageUpdate success")
    	fmt.Println(resp)
    }
    return resp.Network.Value,err
}
func (u *NetworkService) storageDelete(key string ) error {
	resp, err := ns.ApiRuntime.Etcdclient.Delete(context.Background(), ns.Prefix+key, nil)
    if err != nil {
    	fmt.Println("storageDelete failed")
        fmt.Println(err)
        return err
    }else{
    	fmt.Println("storageDelete success")
    	fmt.Println(resp)
    	return nil
    }
}

*/