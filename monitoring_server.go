package main

import (
    "fmt"
    "os"
    "strings"
    "time"
    "net/http"
)

var protocols = map[string]bool {
    "https:": true,
    "http:": false,
}

func Probe(url string, samples_file string) (int) {
    client := &http.Client {
        CheckRedirect: func(request * http.Request, via [] * http.Request) error {
            return http.ErrUseLastResponse
        },
        Timeout: 30 * time.Second,
    }

    request, _ := http.NewRequest("GET", url, nil)

    file, _ := os.Create(samples_file)
    defer file.Close()

    fmt.Fprintf(file, "URL=%s\n", url)
    
    for {
        start := time.Now()
        response, error := client.Do(request)
        elapsed := time.Since(start)

        status := response.StatusCode

        if elapsed.Seconds() >= 30 || error != nil {
            fmt.Println("took longer than 30 seconds")
            fmt.Println(status)

            status := -1
            fmt.Fprintf(file, "%d, %d\n", time.Now().Unix(), status)
        } else {
            fmt.Println("took less than 30 seconds")
            fmt.Println(status)

            fmt.Fprintf(file, "%d, %d\n", time.Now().Unix(), status)

            sleep := (30 * time.Second).Seconds() - elapsed.Seconds()
            time.Sleep(time.Duration(sleep) * time.Second)
        }

        defer response.Body.Close()
    }
    return 0
}

func main() {
    url := os.Args[1:][0]
    samples_file := os.Args[1:][1]
    
    is_https, is_url := protocols[strings.Split(url, "/")[0]]
    if !is_url || is_https {
        fmt.Println("please use the prefix http:// in your url")
        os.Exit(1)
    }
    
    Probe(url, samples_file)
}
