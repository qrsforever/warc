package archiver

import (
    "os"
    "fmt"
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"time"
    "golang.org/x/net/proxy"
)

var httpClient *http.Client
var socks5Proxy = ""

func init() {
	jar, _ := cookiejar.New(nil)

    httpClient = &http.Client{
        Timeout: 30 * time.Second,
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: true,
            },
        },
        Jar: jar,
    }
}

func GetHttpClient() *http.Client {
    s5proxy := os.Getenv("SOCKS5_PROXY")
    if s5proxy == socks5Proxy {
        fmt.Fprintln(os.Stdout, "warc not change proxy")
        return httpClient
    }
	jar, _ := cookiejar.New(nil)
    if s5proxy == "" {
        fmt.Fprintln(os.Stdout, "warc change to no proxy")
        httpClient = &http.Client{
            Timeout: 30 * time.Second,
            Transport: &http.Transport{
                TLSClientConfig: &tls.Config{
                    InsecureSkipVerify: true,
                },
            },
            Jar: jar,
        }
    } else {
        fmt.Fprintln(os.Stdout, "warc change to no proxy:", s5proxy)
        dialer, err := proxy.SOCKS5("tcp", s5proxy, nil, proxy.Direct)
        if err != nil {
            fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
            return nil
        }
        httpClient = &http.Client{
            Timeout: 30 * time.Second,
            Transport: &http.Transport{
                TLSClientConfig: &tls.Config{
                    InsecureSkipVerify: true,
                },
                Dial: dialer.Dial,
            },
            Jar: jar,
        }
    }
    socks5Proxy = s5proxy
    return httpClient
}
