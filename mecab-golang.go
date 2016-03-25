package main

import(
  "fmt"
  "github.com/bluele/mecab-golang"
  "os"
  "strings"
)

func parseToNode(m *mecab.MeCab) {
  tg, err  := m.NewTagger()
  if err != nil {
    panic(err)
  }
  defer tg.Destroy()

  lt, err := m.NewLattice("なのはちゃんかわいい")
  if err != nil {
    panic(err)
  }
  defer lt.Destroy()

  node := tg.ParseToNode(lt)
  fmt.Println(node.Surface())
  for {
    features := strings.Split(node.Feature(), ",")
    fmt.Println(features)
    if features[0] == "名詞" || features[0] == "固有名詞" {
      fmt.Println(fmt.Sprintf("%s : %s", node.Surface(), node.Feature()))
    }
    if node.Next() != nil {
      break
    }
  }
}

func main() {
  m, err := mecab.New("-Owakati -d " + os.Getenv("NEOLOGD"))
  if err != nil {
    panic(err)
  }
  defer m.Destroy()
  parseToNode(m)
}
