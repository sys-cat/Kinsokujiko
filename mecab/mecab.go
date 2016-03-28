package mecab

import(
  "github.com/bluele/mecab-golang"
  "os"
  "strings"
  "errors"
)

type Result struct {
  result string
}

func parseToNode(m *mecab.MeCab, mask string, list []string)(result string, err error) {
  tg, err  := m.NewTagger()
  if err != nil {
    panic(err)
  }
  defer tg.Destroy()
  lt, err := m.NewLattice(mask)
  if err != nil {
    panic(err)
  }
  defer lt.Destroy()
  var rs Result
  node := tg.ParseToNode(lt)
  for {
    n := node.Surface()
    features := strings.Split(node.Feature(), ",")
    if features[0] == "名詞" || features[0] == "固有名詞" {
      for _,v := range list {
        if v == n {
          n = strings.Repeat("*", len(n))
        }
      }
      //fmt.Println(fmt.Sprintf("%s : %s", node.Surface(), node.Feature()))
    }
    rs.result += n
    if node.Next() != nil {
      return rs.result, nil
    }
  }
}

func Masking(val string, list []string)(result string, err error) {
  if val == "" || len(list) == 0 {
    return "", errors.New("no data.")
  }
  m, err := mecab.New("-Owakati -d " + os.Getenv("NEOLOGD"))
  if err != nil {
    panic(err)
  }
  defer m.Destroy()
  return parseToNode(m, val, list)
}
