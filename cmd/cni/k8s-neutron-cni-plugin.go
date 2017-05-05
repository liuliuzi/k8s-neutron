// Copyright 2014 CNI authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
    "encoding/json"
    "fmt"
    "os"
    "runtime"
    "net/http"
    "bytes"
    "io/ioutil"
//    "strings"
    "github.com/containernetworking/cni/pkg/skel"
    "github.com/containernetworking/cni/pkg/types"
    "github.com/containernetworking/cni/pkg/version"
)

//const defaultBrName = "cni0"
const dockerPath    = "/usr/bin/docker"
const ovsPath       = "/usr/bin/ovs-vsctl"
const ipPath        = "/bin/ip"

type NetConf struct {
    types.NetConf
    //ServerUrl       string `json:"ServerUrl"`
    ServerUrl       "http://127.0.0.1:8060"
}

func init() {
    // this ensures that main runs only on main thread (thread group leader).
    // since namespace ops (unshare, setns) are done for a single thread, we
    // must ensure that the goroutine does not jump from OS thread to thread
    runtime.LockOSThread()
}

func loadNetConf(bytes []byte) (*NetConf, error) {
    n := &NetConf{
        //BrName: defaultBrName,
    }
    if err := json.Unmarshal(bytes, n); err != nil {
        return nil, fmt.Errorf("failed to load netconf: %v", err)
    }
    return n, nil
}

func notice(args *skel.CmdArgs, mode string) error {
    n, err := loadNetConf(args.StdinData)
    if err != nil {
        return err
    }
    containerID:=os.Getenv("CNI_CONTAINERID")
    podName:=os.Getenv("K8S_POD_NAME")
    cniNetNs:=os.Getenv("CNI_NETNS")
    cniArgs:=os.Getenv("CNI_ARGS")
    fmt.Print(n,containerID,podName)
    var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
    req, err := http.NewRequest("POST", n.ServerUrl, bytes.NewBuffer(jsonStr))
    //req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return  err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }
    fmt.Println(string(body))
    return nil
}

func cmdAdd(args *skel.CmdArgs) error {
    err := notice(args,"add")
    if err != nil {
        return err
    }
    return nil
}


func cmdDel(args *skel.CmdArgs) error {
    err := notice(args,"del")
    if err != nil {
        return err
    }
    return nil
}

func main() {
    skel.PluginMain(cmdAdd, cmdDel, version.Legacy)
}
