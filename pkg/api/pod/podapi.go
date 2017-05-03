package pod
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

type PodService struct {
	// normally one would use DAO (data access object)
	//pods map[string]pod
	//ApiRuntime   *pkg.ApiRuntime
	Prefix       string
}

func (ns PodService) Init(){
	ns.Prefix="/stf/pod/"
	fmt.Println(ns.Prefix)
}

func (ns PodService) Register() {
	ws := new(restful.WebService)
	ws.
		Path("/pods").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well
	ws.Route(ws.GET("/").To(ns.findAllpods))
	ws.Route(ws.GET("/{pod-id}").To(ns.findpod)) // on the respopse
	ws.Route(ws.PUT("/{pod-id}").To(ns.updatepod)) // from the request
	ws.Route(ws.POST("/").To(ns.createpod)) // from the request
	ws.Route(ws.DELETE("/{pod-id}").To(ns.removepod))
	restful.Add(ws)
}

func (ns PodService) findAllpods(request *restful.Request, response *restful.Response) {
	fmt.Println("findAllpods")
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


func (ns PodService) findpod(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("pod-id")
	value, err := ns.storageGet(id)
	if err != nil {
	    fmt.Println(err)
	}else{
	    fmt.Println(value)
	}

	if len(value) == 0 {
		response.WriteErrorString(http.StatusNotFound, "pod could not be found.")
	} else {
		fmt.Println("findpod")
		response.WriteEntity(value)
	}
}

func (u *PodService) updatepod(request *restful.Request, response *restful.Response) {
	nw := new(Pod)
	err := request.ReadEntity(&nw)
	if err == nil {
		//ns.storageUpdate(nw.Id,nw.String())
		fmt.Println("updatepod")
		response.WriteEntity(nw)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}


func (u *PodService) createpod(request *restful.Request, response *restful.Response) {

	nw := new(Pod)
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

func (u *PodService) removepod(request *restful.Request, response *restful.Response) {
	//id := request.PathParameter("pod-id")
	//ns.storageDelete(id)
	fmt.Println("removepod")
}

/*
func (u *PodService) storageSet(key string , value string ) error {
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

func (u *PodService) storageGetKey() ([]string ,error) {
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
    	for _, pod := range resp.pod.pods {
            fmt.Println( pod)
            ret=append(ret,pod.Key)
        }

        return ret,nil

    }
    */
}
func (u *PodService) storageGet(key string ) (string ,error) {
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
    return resp.pod.Value,err
    */
}
/*

func (u *PodService) storageUpdate(key string , value string ) (string ,error) {
	resp, err := ns.ApiRuntime.Etcdclient.Update(context.Background(), ns.Prefix+key,value)
    if err != nil {
    	fmt.Println("storageUpdate failed")
        fmt.Println(err)
    }else{
    	fmt.Println("storageUpdate success")
    	fmt.Println(resp)
    }
    return resp.pod.Value,err
}
func (u *PodService) storageDelete(key string ) error {
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