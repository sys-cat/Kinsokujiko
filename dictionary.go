package Kinsokujiko

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Item is dictionary Item
type Item struct {
	Surf  string
	Slice string
	Kana  string
	Pos   string
}

// Dictionary is Slice any Item
type Dictionary []Item

var dicPath = "_dic/dic.txt"

func Update(dic Dictionary) (string, error) {
	_, err := _add_item(dic)
	if err != nil {
		fmt.Println(err)
	}
	for index, item := range dic {
		if !_check_item(item) {
			return "Error", errors.New(fmt.Sprintf("%d 行目に問題があるようです", index))
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

func _add_item(dic Dictionary) (string, error) {
	dir, _ := os.Getwd()
	dicPath := fmt.Sprintf("%s/%s", dir, dicPath)
	file, err := os.OpenFile(dicPath, os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()
	read, _ := ioutil.ReadAll(file)
	res := fmt.Sprintf("ファイルの詳細：%+v", read)
	return res, err
}

func Show() (Dictionary, error) {
	return Dictionary{}, nil
}

func _get_dic() {}
