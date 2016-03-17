package main

import (
  "fmt"
  "flag"
  "github.com/yukihir0/mecab-go"
)

func main() {
  var input string
  flag.StringVar(&input, "input", "", "Of the analyzed text")
  flag.Parse()

  args := mecab.NewArgs()
  args.DicDir = "/usr/local/lib/mecab/dic/mecab-ipadic-neologd"
  parser, err := mecab.InitializeParser(args)
  if err != nil {
    panic(err)
  }
  defer parser.Release()
  nodes, err := parser.Parse(input)
  if err != nil {
    panic(err)
  }
  for _, node := range nodes {
    fmt.Println(node)
    /*
    if node.Pos == "名詞" && (node.Pos1 == "一般" || node.Pos1 == "固有名詞") {
      fmt.Println(node.Surface)
    }
    */
  }
}
