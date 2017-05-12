package pod
import (
	"net/http"
	"github.com/emicklei/go-restful"
	"fmt"
	//"github.com/liuliuzi/k8s-neutron/pkg"
	"golang.org/x/net/context"
	"strings"
	"github.com/liuliuzi/k8s-neutron/pkg"
	"encoding/json"
	"github.com/liuliuzi/k8s-neutron/pkg/types"
	//"github.com/bitly/go-simplejson"
	//"github.com/emicklei/go-restful/swagger"
)

type PodService struct {
	// normally one would use DAO (data access object)
	Pods map[string]types.Pod
	ApiRuntime   *pkg.ApiRuntime
	Prefix       string
}


func (ps PodService) Register() {
	ws := new(restful.WebService)
	ws.
		Path("/pods").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well
	ws.Route(ws.GET("/").To(ps.findAllpods))
	ws.Route(ws.GET("/{pod-id}").To(ps.findpod)) // on the respopse
	ws.Route(ws.PUT("/{pod-id}").To(ps.updatepod)) // from the request
	ws.Route(ws.POST("/").To(ps.createpod)) // from the request
	ws.Route(ws.DELETE("/{pod-id}").To(ps.removepod))
	restful.Add(ws)
}

func (ps PodService) findAllpods(request *restful.Request, response *restful.Response) {
	fmt.Println("findAllpods")
    value, err := ps.storageGetKey()
    if err != nil {
	    fmt.Println(err)
	    return
	}
	pods := []types.Pod{}
	for _, key := range value {
            key=strings.Replace(key,ps.Prefix,"",20)
            value, _ := ps.storageGet(key)
            pod := new(types.Pod)
            err = json.Unmarshal([]byte(value), pod)
			if err != nil {
			    fmt.Println(err)
			    return
			}
            pods=append(pods,*pod)
        }

	response.WriteEntity(pods)
}


func (ps PodService) findpod(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("pod-id")
	value, err := ps.storageGet(id)
	pod := new(types.Pod)
	err = json.Unmarshal([]byte(value), pod)
	if err != nil {
	    fmt.Println(err)
	}else{
	    fmt.Println(value)
	}

	if len(value) == 0 {
		response.WriteErrorString(http.StatusNotFound, "pod could not be found.")
	} else {
		fmt.Println("findpod")
		response.WriteEntity(pod)
	}
}

func (ps *PodService) updatepod(request *restful.Request, response *restful.Response) {
	pod := new(types.Pod)
	err := request.ReadEntity(&pod)
	if err == nil {
		ps.storageUpdate(pod.Id,pod.String())
		fmt.Println("updatepod")
		response.WriteEntity(pod)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}


func (ps *PodService) createpod(request *restful.Request, response *restful.Response) {
	fmt.Println("createpod")
	pod := new(types.Pod)
	err := request.ReadEntity(&pod)
	if err == nil {
		fmt.Println("createpod succ")
		//ps.storageSet(pod.Id,pod)
		ps.storageSet(pod.Id,pod.String())
		//ps.storageSet(pod.Id,{'Id':'123566666','Name':'liu'})

		//ps.storageSet(pod.Id,kk)
		response.WriteHeaderAndEntity(http.StatusCreated, pod)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func (ps *PodService) removepod(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("pod-id")
	ps.storageDelete(id)
}


func (ps *PodService) storageSet(key string , value string ) error {
	if ps.ApiRuntime==nil {
		return nil
	}
	resp, err := ps.ApiRuntime.Etcdclient.Set(context.Background(), ps.Prefix+key, value, nil)
    if err != nil {
    	fmt.Println("storageSet failed")
        fmt.Println(err)
    }else{
    	fmt.Println("storageSet success")
    	fmt.Println(resp)
    }
    return err
}

func (ps *PodService) storageGetKey() ([]string ,error) {

	resp, err := ps.ApiRuntime.Etcdclient.Get(context.Background(), ps.Prefix, nil)

    if err != nil {
    	fmt.Println("storageGetKey failed")
        fmt.Println(err)
        return nil,err
    }else{
    	ret:=[]string{}
    	fmt.Println("storageGetKey success")
    	for _, pod := range resp.Node.Nodes  {
            fmt.Println( pod)
            ret=append(ret,pod.Key)
        }

        return ret,nil

    }

}
func (ps *PodService) storageGet(key string ) (string ,error) {
	fmt.Println("storageGet", key)

	resp, err := ps.ApiRuntime.Etcdclient.Get(context.Background(), ps.Prefix+key, nil)
    if err != nil {
    	fmt.Println("storageGet failed")
        fmt.Println(err)
    }else{
    	fmt.Println("storageGet success")
    	fmt.Println(resp)
    }
    return resp.Node.Value,err
}


func (ps *PodService) storageUpdate(key string , value string ) (string ,error) {
	resp, err := ps.ApiRuntime.Etcdclient.Update(context.Background(), ps.Prefix+key,value)
    if err != nil {
    	fmt.Println("storageUpdate failed")
        fmt.Println(err)
    }else{
    	fmt.Println("storageUpdate success")
    	fmt.Println(resp)
    }
    return resp.Node.Value,err

}
func (ps *PodService) storageDelete(key string ) error {
	resp, err := ps.ApiRuntime.Etcdclient.Delete(context.Background(), ps.Prefix+key, nil)
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

