package kinsokujiko

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
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

func Show(path string) (Dictionary, error) {
	return _get_dic(path)
}

func _get_dic(path string) (Dictionary, error) {
	var dicPath string
	if path != "" {
		dicPath = path
	}
	if dicPath == "" {
		dir, _ := os.Getwd()
		dicPath = fmt.Sprintf("%s/%s", dir, dicPath)
	}
	log.Println(fmt.Sprintf("open dictionary file is %s", dicPath))
	file, err := os.OpenFile(dicPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return Dictionary{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var dic Dictionary
	r := regexp.MustCompile(`^\#`)
	for scanner.Scan() {
		sp := strings.Split(scanner.Text(), ",")
		if len(sp) >= 4 {
			if !r.MatchString(sp[0]) { // 先頭に#が付くとコメント
				dic = append(dic, Item{
					sp[0],
					sp[1],
					sp[2],
					sp[3],
				})
			}
		}
	}

	return dic, err
}
