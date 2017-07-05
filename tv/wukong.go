package tv

import (
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const wukongurl = "http://www.5kong.tv"

//悟空
func Wukong() {
	defer func() {
		ch <- true
		if r := recover(); r != nil {
			fmt.Println("wukong recovered in f", r)
		}
	}()
	index1 = 1
	var urlmap map[string]string
	urlmap = make(map[string]string)
	doc, err := goquery.NewDocument("http://www.5kong.tv/category")
	if err != nil {
		fmt.Println(err)
	}

	doc.Find(".list .item").Each(func(i int, contrentSelection *goquery.Selection) {
		href, _ := contrentSelection.Find("a").Attr("href")
		suburl := wukongurl + href
		name := contrentSelection.Find(".game-zone-name").Text()
		urlmap[name] = suburl
		if flag {
			fmt.Println(i+1, name, suburl)
		}

	})

	for name, suburl := range urlmap {
		if flag {
			fmt.Println("-----", name, suburl)
		}

		wkMore(suburl)
	}

	//	fmt.Println("悟空TV:", index1)
	Outputmap["悟空TV"] = strconv.Itoa(index1)

}

func wkMore(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
	}

	doc.Find(".live-recommend ul li").Each(func(i int, selection *goquery.Selection) {
		roomid, _ := selection.Find("a").Attr("href")
		liveurl := wukongurl + roomid
		title := selection.Find(".anchor-name").Text()
		if flag {
			fmt.Println(index1, i+1, title, liveurl)
		}

		index1++
	})

}
