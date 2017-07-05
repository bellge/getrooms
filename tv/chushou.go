package tv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const chushouurl = "http://chushou.tv/live/list.htm"
const urlmore = "http://chushou.tv/live/down-v2.htm?&breakpoint="

var chushouMore ChushouMore
var index1, index2, index3 int
var flag bool

type ChushouMore struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Count int `json:"count"`
		Items []struct {
			ID      int    `json:"id"`
			PanelID int    `json:"panelId"`
			Type    int    `json:"type"`
			Name    string `json:"name"`
			Meta    struct {
				OnlineCount     int    `json:"onlineCount"`
				Creator         string `json:"creator"`
				Gender          string `json:"gender"`
				GameName        string `json:"gameName"`
				SubscriberCount int    `json:"subscriberCount"`
				GameTargetKey   string `json:"gameTargetKey"`
				Avatar          string `json:"avatar"`
				GameCtxPath     string `json:"gameCtxPath"`
				Live            bool   `json:"live"`
				Professional    int    `json:"professional"`
			} `json:"meta"`
			Cover             string      `json:"cover"`
			TargetKey         string      `json:"targetKey"`
			Desc              string      `json:"desc"`
			Order             int         `json:"order"`
			Style             int         `json:"style"`
			CategoryID        int         `json:"categoryId"`
			StartTime         interface{} `json:"startTime"`
			EndTime           interface{} `json:"endTime"`
			CreatedTime       interface{} `json:"createdTime"`
			UpdatedTime       interface{} `json:"updatedTime"`
			Display           int         `json:"display"`
			Hot               int         `json:"hot"`
			Scope             int         `json:"scope"`
			CornerMark        int         `json:"cornerMark"`
			CategorySpecified bool        `json:"categorySpecified"`
			Show              bool        `json:"show"`
		} `json:"items"`
		Breakpoint string `json:"breakpoint"`
	} `json:"data"`
}

//一个一个直播间的数，可以获取直播间的更多信息
func Chushou() {
	defer func() {
		ch <- true
		if r := recover(); r != nil {
			fmt.Println("chushou recovered in f", r)
		}
	}()
	index3 = 1
	doc, err := goquery.NewDocument(chushouurl)
	if err != nil {
		fmt.Println(err)
	}

	doc.Find(".liveCon .liveOne").Each(func(i int, contentSelection *goquery.Selection) {

		title := contentSelection.Find(".livePlayerName").Text()
		livecount := contentSelection.Find(".liveCount").Text()
		livePlayerName := contentSelection.Find(".livePlayerName").Text()
		liveDetail, _ := contentSelection.Find(".liveDetail a").Attr("href")
		herf, _ := contentSelection.Find(".home_live_block a").Eq(1).Attr("href")
		herf2 := "http://chushou.tv" + herf
		if flag {
			fmt.Println(index3, i+1, title, livecount, livePlayerName, liveDetail, herf2)
		}

		index3++
	})

	databreak, _ := doc.Find(".more").Attr("data-break")
	//	fmt.Println(databreak)
	csmoreJson(databreak)
	//	fmt.Println("触手TV:", index3)
	Outputmap["触手TV"] = strconv.Itoa(index3)

}

func csmoreJson(databreak string) {
	chushouMore = ChushouMore{}
	url := urlmore + databreak
	rep, err := http.Get(url)
	if err != nil {
		return
	}
	result, _ := ioutil.ReadAll(rep.Body)
	json.Unmarshal(result, &chushouMore)
	for i, v := range chushouMore.Data.Items {
		roomurl := "http://chushou.tv/room/" + v.TargetKey + ".htm"
		if flag {
			fmt.Println(index3, i+1, v.Name, v.Meta.OnlineCount, roomurl)
		}

		index3++
	}
	if chushouMore.Data.Breakpoint != databreak {
		csmoreJson(chushouMore.Data.Breakpoint)
	}
	return
}

//直接获取总数，简单快速
func Chushou2() {
	defer func() {
		//ch <- true
		Wg.Done()
		if r := recover(); r != nil {
			fmt.Println("chushou recovered in f", r)
		}
	}()
	doc, err := goquery.NewDocument("http://chushou.tv/live/list.htm")
	if err != nil {
		fmt.Println(err)
	}
	doc.Find(".block_title").Each(func(i int, contentSelection *goquery.Selection) {

		livenum := contentSelection.Find(".liveNum").Text()
		Outputmap["触手TV"] = livenum
	})

}
