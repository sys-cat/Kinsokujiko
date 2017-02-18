package Kinsokujiko

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Item is dictionary Item
type Item struct {
	Surf  string `json:"surf" form:"surf"`
	Slice string `json:"slice" form:"slice"`
	Kana  string `json:"kana" form:"kana"`
	Pos   string `json:"pos" form:"pos"`
}

// Dictionary is Slice any Item
type Dictionary []Item

var dicPath = "_dic/dic.txt"

func Update(dic Dictionary) (bool, error) {
	for index, item := range dic {
		if !_check_item(item) {
			return false, errors.New(fmt.Sprintf("%d 行目に問題があるようです", index))
		}
	}
	return _add_item(dic)
}

func _check_item(it Item) bool {
	var res = false
	if strings.Count(it.Surf, "") > 1 {
		res = true
	}
	if strings.Count(it.Slice, "") > 1 {
		res = true
	}
	if strings.Count(it.Kana, "") > 1 {
		res = true
	}
	if strings.Count(it.Pos, "") > 1 {
		res = true
	}
	return res
}

func _add_item(dic Dictionary) (bool, error) {
	dir, _ := os.Getwd()
	dicPath := fmt.Sprintf("%s/%s", dir, dicPath)
	file, err := os.OpenFile(dicPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return false, err
	}
	defer file.Close()
	read, _ := ioutil.ReadAll(file)
	items := string(read)
	for _, varible := range dic {
		items = items + fmt.Sprintf("%s,%s,%s,%s\n", varible.Surf, varible.Slice, varible.Kana, varible.Pos)
	}
	rerr := ioutil.WriteFile(dicPath, []byte(items), os.ModePerm)
	res := true
	if rerr != nil {
		res = false
	}
	return res, rerr
}

func Show() (Dictionary, error) {
	return _get_dic()
}

func _get_dic() (Dictionary, error) {
	dir, _ := os.Getwd()
	dicPath := fmt.Sprintf("%s/%s", dir, dicPath)
	file, err := os.OpenFile(dicPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return Dictionary{}, err
	}
	defer file.Close()
	scnr := bufio.NewScanner(file)
	var dic Dictionary
	for scnr.Scan() {
		sp := strings.Split(scnr.Text(), ",")
		dic = append(dic, Item{
			sp[0],
			sp[1],
			sp[2],
			sp[3],
		})
	}

	return dic, err
}
