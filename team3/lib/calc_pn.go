package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// []struct -> float64

type Dic struct {
	Words []WordPn `json:"dictionary"`
}

type WordPn struct {
	Surface  string  `json:"surface"`
	Readging string  `json:"readging"`
	Pos      string  `json:"pos"`
	Pn       float64 `json:"pn"`
}

type Word struct {
	Surface  string `json:"surface"`
	Readging string `json:"readging"`
	Pos      string `json:"pos"`
}

var dic []WordPn

func calcPn(words []*Word) float64 {

	pn := 0.0

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(dic); j++ {
			if words[i].Surface == dic[j].Surface {
				pn += dic[j].Pn
			}
		}
	}

	return pn
}

func loadDic() {
	file, e := ioutil.ReadFile("./pn_ja.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		// os.Exit(1)
	}
	// fmt.Printf("%s\n", string(file))

	var jsontype Dic
	json.Unmarshal(file, &jsontype)
	dic = jsontype.Words
}

func main() {
	loadDic()

	words := []*Word{
		&Word{
			Surface:  "go",
			Readging: "せかい",
			Pos:      "名詞",
		},
		&Word{
			Surface:  "は",
			Readging: "",
			Pos:      "",
		},
		&Word{
			Surface:  "楽しい",
			Readging: "",
			Pos:      "",
		},
		&Word{
			Surface:  "そして",
			Readging: "",
			Pos:      "",
		},
		&Word{
			Surface:  "幸せ",
			Readging: "",
			Pos:      "",
		},
	}
	fmt.Printf("pn = %f\n", calcPn(words))

	words = []*Word{
		&Word{
			Surface:  "javascript",
			Readging: "",
			Pos:      "",
		},
		&Word{
			Surface:  "は",
			Readging: "",
			Pos:      "",
		},
		&Word{
			Surface:  "悲しい",
			Readging: "",
			Pos:      "",
		},
	}
	fmt.Printf("pn = %f\n", calcPn(words))
}
