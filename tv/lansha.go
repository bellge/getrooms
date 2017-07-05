package tv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const lashaurl = "http://www.lansha.tv/getLiveList.html?show=1&pageIndex="

var lanshaMore LanshaMore

type LanshaMore struct {
	Data []struct {
		Avatar    string `json:"avatar"`
		BogerName string `json:"bogerName"`
		GameName  string `json:"gameName"`
		GameURL   string `json:"gameURL"`
		NickName  string `json:"nickName"`
		RoomImg   string `json:"roomImg"`
		RoomName  string `json:"roomName"`
		RoomURL   string `json:"roomURL"`
		ViewNum   int    `json:"viewNum"`
	} `json:"data"`
	Status int `json:"status"`
}

//蓝鲨
func Lansha() {
	defer func() {
		ch <- true
		if r := recover(); r != nil {
			fmt.Println("lansha recovered in f", r)
		}
	}()
	index2 = 1
	doc, err := goquery.NewDocument("http://www.lansha.tv/liveList.html")
	if err != nil {
		fmt.Println(err)
	}
	doc.Find(".game-list .game-one").Each(func(i int, contentSelection *goquery.Selection) {
		title := contentSelection.Find(".room-name-title").Text()
		roomid, _ := contentSelection.Find("a").Attr("href")
		url2 := "http://www.lansha.tv" + roomid
		if flag {
			fmt.Println(index2, i+1, title, url2)
		}

		index2++
	})
	lsmoreJson(2)
	//	fmt.Println("蓝鲨TV:", index2)
	Outputmap["蓝鲨TV"] = strconv.Itoa(index2)

}

func lsmoreJson(pageIndex int) {
	flags := false
	lanshaMore = LanshaMore{}
	url := lashaurl + strconv.Itoa(pageIndex)
	rep, err := http.Get(url)
	if err != nil {
		return
	}
	defer rep.Body.Close()
	result, _ := ioutil.ReadAll(rep.Body)
	json.Unmarshal(result, &lanshaMore)
	for i, v := range lanshaMore.Data {
		url := "http://www.lansha.tv" + v.RoomURL
		if flag {
			fmt.Println(index2, i+1, v.RoomName, url)
		}

		index2++
		flags = true
	}
	if flags {
		pageIndex++
		lsmoreJson(pageIndex)
	}

}
