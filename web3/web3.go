package web3

import (
  "github.com/bitly/go-simplejson"
  "github.com/parnurzeal/gorequest"
  "log"
  // "fmt"
)

const (
  apiUrl string = "https://mainnet-rpc.tlxscan.com/" // 8545 (default port)
)

func Call(args ...string) (*simplejson.Json) {
  method := args[0]
  params := "[]"
  if len(args) > 1 {
    params = args[1]
  }

  postBody := `{"jsonrpc":"2.0","method":"`+method+`","params":`+params+`}`
  // fmt.Println("postBody: " + postBody)

  _, body, errs := gorequest.New().Post(apiUrl).
    Send(postBody).
    End()

  if errs != nil {
    panic(errs)
  }

  js, err := simplejson.NewJson([]byte(body))
  if err != nil {
      log.Fatalln(err)
  }

  return js
}
