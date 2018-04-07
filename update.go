package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "strconv"
    "encoding/json"
)

type MServer struct {
    Server string `json:"server"`
    Port int64 `json:"port"`
    Weight int64 `json:"weight"`
}

type ServerList *[]*MServer 

type Upstream struct {
    Name string `json:"name"`
    Servers ServerList `json:"servers"`
}

type UpstreamList []*Upstream

type Upstreams struct {
    Upstreams *UpstreamList `json:"upstreams"`
}

func updateJson() ([]byte, error) {
    var i int = 0
    // 16 upstreams
    var upstreams  UpstreamList
    for i = 0; i < 2; i++ {
	upstream := Upstream{Name: "dyhost"+ strconv.Itoa(i)}
        servers := []*MServer{}
	// 2 serevers
	var j int  = 0
	for j = 0; j <2; j ++ {
	    server := MServer{ Server: "127.0.0.1:"+strconv.Itoa(8088+j), Weight: 100, }
	    servers = append(servers, &server)
	}
	upstream.Servers = &servers
	upstreams = append(upstreams, & upstream)
    }

    //results := Upstreams{Upstreams: &upstreams}
    
    jsonUps, err := json.Marshal(upstreams)
    if err != nil {
	fmt.Println(err)
	return []byte{}, err
    }

    fmt.Println(string(jsonUps))
    return jsonUps, nil
}

func main(){
    url := "http://127.0.0.1:8081/upstream/update"
    
    //json序列化
    //post := "{\"UserId\":\"" + "123412" +"\",\"Password\":\"" + "1234123" + "\"}"
    post := "server 127.0.0.1:8089;server 127.0.0.1:8088;"
    ups, _ := updateJson()
    post = string(ups)

    fmt.Println(url, "post", post)
    
    var jsonStr = []byte(post)
    fmt.Println("jsonStr", jsonStr)
    fmt.Println("new_str", bytes.NewBuffer(jsonStr))

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    // req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}

