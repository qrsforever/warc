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

func init() {
	jar, _ := cookiejar.New(nil)

    socks5_proxy := os.Getenv("SOCKS5_PROXY")
    if socks5_proxy != "" {
        fmt.Fprintln(os.Stdout, "using socks5:", socks5_proxy)
        dialer, err := proxy.SOCKS5("tcp", socks5_proxy, nil, proxy.Direct)
        if err != nil {
            fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
            os.Exit(1)
        }
        httpClient = &http.Client{
            Timeout: time.Minute,
            Transport: &http.Transport{
                TLSClientConfig: &tls.Config{
                    InsecureSkipVerify: true,
                },
                Dial: dialer.Dial,
            },
            Jar: jar,
        }
    } else {
        httpClient = &http.Client{
            Timeout: time.Minute,
            Transport: &http.Transport{
                TLSClientConfig: &tls.Config{
                    InsecureSkipVerify: true,
                },
            },
            Jar: jar,
        }
    }
}
