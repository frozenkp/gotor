/*
code original from http://www.devdungeon.com/content/making-tor-http-requests-go
*/
package gotor

import (
  "log"
  "net/http"
  "net/url"
  "time"

  "golang.org/x/net/proxy"
)

// Specify Tor proxy ip and port
var torProxy string = "socks5://127.0.0.1:9050" // 9150 w/ Tor Browser

// New return a tor client.
//
// Using "socks5://127.0.0.1:9050" as proxy in default, you can put your own proxy
// in paramemter to replace it.
func New(timeOut time.Duration, InputProxy ...string) *http.Client{
  // Parse Tor proxy URL string to a URL type
  var proxyURL string
  if len(InputProxy) != 0 {
    proxyURL = InputProxy[0]
  }else{
    proxyURL = torProxy
  }
  torProxyUrl, err := url.Parse(proxyURL)
  if err != nil {
    log.Fatal("Error parsing Tor proxy URL:", torProxy, ".", err)
  }

  // Create proxy dialer using Tor SOCKS proxy
  torDialer, err := proxy.FromURL(torProxyUrl, proxy.Direct)
  if err != nil {
    log.Fatal("Error setting Tor proxy.", err)
  }

  // Set up a custom HTTP transport to use the proxy and create the client
  torTransport := &http.Transport{Dial: torDialer.Dial}
  client := &http.Client{Transport: torTransport, Timeout: timeOut}

  return client
}
