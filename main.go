package main

import (
    "fmt"
    "strings"
    "io/ioutil"
    "net/http"
    "net/http/cookiejar"
    "encoding/json"
)

type Course struct {
    Name string
    Id string
}

type Profile struct {
    Uid string
    Password string
    Courses []Course
}

func main() {
    var v Profile
    data, _ := ioutil.ReadFile("./profile.json")
    _ = json.Unmarshal([]byte(data), &v)

    var client http.Client
    jar, _ := cookiejar.New(nil)
    client.Jar = jar

    req, _ := http.NewRequest("GET" ,"<url>", nil)
    req.Header.Add("User-Agent", "<UA>")
    req.Header.Add("Host", "<host>")

    res, _ := client.Do(req)

    exec := "<exec>"

    payload := fmt.Sprintf("username=%s&password=%s&execution=%s&_eventId=%s&geolocation=",
                           v.Uid, v.Password, exec, "submit")
    req, _ = http.NewRequest("POST", "<url>", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded;param=value")
    res, _ = client.Do(req)

    req, _ = http.NewRequest("GET", "<url>", nil)
    res, _ = client.Do(req)

    req.Header.Add("Referer", "<url>")
    req.Header.Add("X-Requested-With", "XMLHttpRequest")

    for i := 0; i < len(v.Courses); i++ {
        url := "<url>" + v.Courses[i].Id
        req, _ = http.NewRequest("GET", url, nil) 
        res, _ = client.Do(req)
        body, _ = ioutil.ReadAll(res.Body)
        fmt.Printf("%s: %s\n", v.Courses[i].Name, string(body))
    }
}