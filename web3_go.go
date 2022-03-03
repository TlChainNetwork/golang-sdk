package main

import (
    "fmt"
    "github.com/makevoid/web3_go/web3"
    "github.com/bitly/go-simplejson"
    "encoding/json"
    "encoding/hex"
	  "github.com/ethereum/go-ethereum/crypto"
    "strings"
)

func main() {
  account := tl20()
  fmt.Println("account:", account)

  accs := accounts()
  fmt.Println("accounts:", accs)

  bal := getBalance(account)
  fmt.Println("balance of account:", account, "=", bal)

  contract := "contract test { function multiply(uint a) returns(uint d) {   return a * 7;   } }"
  resp_c := compile(contract)
  fmt.Println("compiled contract infos (abi):")
  info_c := getAbi(resp_c)
  pp(info_c)


  sha := crypto.Sha3( []byte("data()") )
  fmt.Println("sha3('data()'):", sha)
  shaHex := hex.EncodeToString(sha[:])
  fmt.Println("sha3('data()'):", shaHex)
  fmt.Println("method data() signature:", shaHex[0:8])


  data := "0x"+shaHex[0:8]
  to   := "0x0000000000000000000000000000000000000000"

  resp := call(data, to)
  fmt.Println("call method(), resp:", resp)
  str := resp.(string)
  str = strings.TrimPrefix(str, "0x")
  str = string(str)
  value, err := hex.DecodeString(str)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("value (as bytes):",  value)
  fmt.Println("value (as string):", string(value))

}

func compile(contract string) (*simplejson.Json) {
  res := web3.Call("eth_compileSolidity", `["`+contract+`"]`).Get("result")
  return res
}

func getAbi(compiledResp *simplejson.Json) (interface {}) {
  return compiledResp.Get("test").Get("info").Get("abiDefinition").MustArray()
}

func call(data string, to string) (interface {}) {
  callArgs := `[{ "to": "`+to+`", "data": "`+data+`" }]`
  res := web3.Call("eth_call", callArgs)
  errStr := res.Get("error").MustString()
  if errStr != "" {
    fmt.Println("Error: " + errStr)
    pp(res)
  }
  resp := res.Get("result").MustString()
  return resp
}

func tl20() (string) {
  res := web3.Call("eth_tl20").Get("result").MustString()
  return res
}

func accounts() ([]interface {}) {
  res := web3.Call("eth_accounts").Get("result").MustArray()
  return res
}

func getBalance(address string) (string) {
  pp(web3.Call("eth_getBalance", "[\""+address+"\"]").Get("result"))
  res := web3.Call("eth_getBalance", "[\""+address+"\"]").Get("result").MustString()
  return res
}

// pretty print

func pp(data interface {}) {
  js, err := json.MarshalIndent(data, "", "  ")
  if err != nil {
    fmt.Println("error:", err)
  }
  fmt.Print(string(js))
  println()
}
