package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "strconv"
    "encoding/json"
)

type Upstream struct {
    Name string `json:"name"`
}

type UpstreamList []*Upstream

type Upstreams struct {
    Upstreams *UpstreamList `json:"upstreams"`
}

func deleteJson() ([]byte, error) {
    var i int = 0
    // 16 upstreams
    var upstreams  UpstreamList
    for i = 0; i < 2; i++ {
	upstream := Upstream{Name: "dyhost"+ strconv.Itoa(i)}
	upstreams = append(upstreams, & upstream)
    }

    jsonUps, err := json.Marshal(upstreams)
    if err != nil {
	fmt.Println(err)
	return []byte{}, err
    }

    fmt.Println(string(jsonUps))
    return jsonUps, nil
}

func main(){
    url := "http://127.0.0.1:8081/upstream/delete"
    
    //json序列化
    //post := "{\"UserId\":\"" + "123412" +"\",\"Password\":\"" + "1234123" + "\"}"
    post := "server 127.0.0.1:8089;server 127.0.0.1:8088;"
    ups, _ := deleteJson()
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

