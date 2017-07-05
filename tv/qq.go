package tv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const qqurl = "http://egame.qq.com/livelist?layoutid=hot&name=%E7%83%AD%E9%97%A8%E7%9B%B4%E6%92%AD"
const qqmoreurl = "http://share.egame.qq.com/cgi-bin/pgg_skey_async_fcgi?&cgi_module=pgg_live_read_svr&cgi_method=get_live_list&p_tk=cMErwTMynmGdltrgrwC1dhsVBupV9Gg7ctCwz13NDj4_&g_tk=1661463292"

var qqret QQResult

type QQResult struct {
	Ecode int `json:"ecode"`
	Data  struct {
		Num0 struct {
			Module  string `json:"module"`
			Method  string `json:"method"`
			RetMsg  string `json:"retMsg"`
			RetCode int    `json:"retCode"`
			RetBody struct {
				Data struct {
					IsGetOver int `json:"is_get_over"`
					LiveData  struct {
						LiveList []struct {
							AnchorFaceURL string `json:"anchor_face_url"`
							AnchorID      int    `json:"anchor_id"`
							AnchorName    string `json:"anchor_name"`
							Appid         string `json:"appid"`
							Appname       string `json:"appname"`
							City          string `json:"city"`
							Ext           struct {
							} `json:"ext"`
							FansCount int    `json:"fans_count"`
							Online    int    `json:"online"`
							Tag       string `json:"tag"`
							Title     string `json:"title"`
							VideoInfo struct {
								Dst      string `json:"dst"`
								Provider int    `json:"provider"`
								URL      string `json:"url"`
								VAttr    struct {
									DualID      int    `json:"dual_id"`
									DualType    int    `json:"dual_type"`
									HvDirection int    `json:"hv_direction"`
									Source      string `json:"source"`
									VCacheTmMax int    `json:"v_cache_tm_max"`
									VCacheTmMin int    `json:"v_cache_tm_min"`
									VHeight     int    `json:"v_height"`
									VPlayMode   int    `json:"v_play_mode"`
									VWidth      int    `json:"v_width"`
								} `json:"v_attr"`
								Vid       string `json:"vid"`
								VideoType int    `json:"video_type"`
							} `json:"video_info"`
						} `json:"live_list"`
					} `json:"live_data"`
					Total int `json:"total"`
				} `json:"data"`
				Message  string `json:"message"`
				Result   int    `json:"result"`
				TimeCost int    `json:"time_cost"`
			} `json:"retBody"`
		} `json:"0"`
	} `json:"data"`
	LoginCost int `json:"login_cost"`
	TimeCost  int `json:"time_cost"`
}

//http://share.egame.qq.com/cgi-bin/pgg_skey_async_fcgi?
//&cgi_module=pgg_live_read_svr
//&cgi_method=get_live_list
//&p_tk=cMErwTMynmGdltrgrwC1dhsVBupV9Gg7ctCwz13NDj4_
//&g_tk=1661463292

//http://share.egame.qq.com/cgi-bin/pgg_skey_async_fcgi?&cgi_module=pgg_live_read_svr&cgi_method=get_live_list&p_tk=cMErwTMynmGdltrgrwC1dhsVBupV9Gg7ctCwz13NDj4_&g_tk=1661463292
//http://share.egame.qq.com/cgi-bin/pgg_skey_async_fcgi?&cgi_module=pgg_live_read_svr&cgi_method=get_live_list&p_tk=cMErwTMynmGdltrgrwC1dhsVBupV9Gg7ctCwz13NDj4_&g_tk=1661463292

//key:param
//value:{"0":{"module":"pgg_live_read_svr","method":"get_live_list","param":{"layout_id":"hot","page_num":2,"page_size":20,"other_uid":0}}}

//key:app_info
//value:{"platform":4,"terminal_type":2,"egame_id":"egame_official"}

func Qq() {
	defer func() {
		Wg.Done()
		if r := recover(); r != nil {
			fmt.Println("qq recovered in f", r)
		}
	}()
	num := 2
	data := make(url.Values)
	param := "{\"0\":{\"module\":\"pgg_live_read_svr\",\"method\":\"get_live_list\",\"param\":{\"layout_id\":\"hot\",\"page_num\":" + strconv.Itoa(num) + ",\"page_size\":20,\"other_uid\":0}}}"
	app_info := "{platform\": 4, \"terminal_type\": 2, \"egame_id\": \"egame_official\"}"
	data["param"] = []string{param}
	data["app_info"] = []string{app_info}

	//把post表单发送给目标服务器
	res, err := http.PostForm(qqmoreurl, data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()

	result, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(result, &qqret)

	count := qqret.Data.Num0.RetBody.Data.Total
	Outputmap["企鹅电竞"] = strconv.Itoa(count)
}
