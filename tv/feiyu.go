package tv

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const feiyunurl = "http://www.feiyun.tv/room"

//飞云tv
func Feiyun() {
	defer func() {
		Wg.Done()
		if r := recover(); r != nil {
			fmt.Println("feiyun recovered in f", r)
		}
	}()
	//	fmt.Println("enter feiyun...")
	doc, err := goquery.NewDocument(feiyunurl)
	if err != nil {
		fmt.Println("feiyun newDocument err:", err)
	}
	//	fmt.Println("find feiyun...")
	doc.Find(".list-wrap .list-title").Each(func(i int, contrentSelection *goquery.Selection) {
		total := contrentSelection.Find(".curr").Text()
		//		fmt.Println("totla:", total)
		ibegin := strings.Index(total, "(")
		iend := strings.Index(total, ")")
		//		fmt.Println(ibegin, iend)
		str := string([]byte(total)[ibegin+1 : iend])
		//		fmt.Println("飞云TV:", str)
		Outputmap["飞云TV"] = str
	})
}
