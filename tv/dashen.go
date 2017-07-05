package tv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const dashenurl = "http://www.dashen.tv/live/get_all_live_list/1"

var dashenret DashenRet

type DashenRet struct {
	ResultCode int `json:"result_code"`
	ResultData struct {
		Result []struct {
			AnchorID     string `json:"anchor_id"`
			ImageURL     string `json:"image_url"`
			Popularity   string `json:"popularity"`
			HeadImageURL string `json:"head_image_url"`
			Nick         string `json:"nick"`
			Sex          string `json:"sex"`
			Name         string `json:"name"`
		} `json:"result"`
		Count int `json:"count"`
	} `json:"result_data"`
}

//大神TV
func Dashen() {
	defer func() {
		Wg.Done()
		if r := recover(); r != nil {
			fmt.Println("dashentv recovered in f", r)
		}
	}()
	dashenret = DashenRet{}
	rep, err := http.Get(dashenurl)
	if err != nil {
		return
	}
	defer rep.Body.Close()
	result, _ := ioutil.ReadAll(rep.Body)
	json.Unmarshal(result, &dashenret)
	count := dashenret.ResultData.Count
	//	fmt.Println("大神TV:", count)
	Outputmap["大神TV"] = strconv.Itoa(count)
}
