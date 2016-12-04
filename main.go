package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ikawaha/kagome/tokenizer"
	"github.com/k0kubun/pp"
)

var sample = `日本国民は、正当に選挙された国会における代表者を通じて行動し、われらとわれらの子孫のために、諸国民との協和による成果と、わが国全土にわたつて自由のもたらす恵沢を確保し、政府の行為によつて再び戦争の惨禍が起ることのないやうにすることを決意し、ここに主権が国民に存することを宣言し、この憲法を確定する。そもそも国政は、国民の厳粛な信託によるものであつて、その権威は国民に由来し、その権力は国民の代表者がこれを行使し、その福利は国民がこれを享受する。これは人類普遍の原理であり、この憲法は、かかる原理に基くものである。われらは、これに反する一切の憲法、法令及び詔勅を排除する。
　日本国民は、恒久の平和を念願し、人間相互の関係を支配する崇高な理想を深く自覚するのであつて、平和を愛する諸国民の公正と信義に信頼して、われらの安全と生存を保持しようと決意した。われらは、平和を維持し、専制と隷従、圧迫と偏狭を地上から永遠に除去しようと努めてゐる国際社会において、名誉ある地位を占めたいと思ふ。われらは、全世界の国民が、ひとしく恐怖と欠乏から免かれ、平和のうちに生存する権利を有することを確認する。
　われらは、いづれの国家も、自国のことのみに専念して他国を無視してはならないのであつて、政治道徳の法則は、普遍的なものであり、この法則に従ふことは、自国の主権を維持し、他国と対等関係に立たうとする各国の責務であると信ずる。
　日本国民は、国家の名誉にかけ、全力をあげてこの崇高な理想と目的を達成することを誓ふ。`

const pos = "名詞"
var count int = 0

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/mask", mask)
		v1.GET("/dic", dic)
	}
	router.Run(":8080")
}

func mask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"response":"OK"})
}

func dic(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"response":"OK"})
}

func morphological() int {
	var dic tokenizer.Dic
	dic = tokenizer.SysDicIPA()
	pp.Print(dic)
	t := tokenizer.NewWithDic(dic)
	tokens := t.Analyze(sample, tokenizer.Normal)
	for _, token := range tokens {
		if token.Pos() == pos {
			//pp.Printf("%s : %s\n", token.Surface, token.Pos())
			count = count + 1
		}
		continue
	}
	return count
}