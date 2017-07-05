package main

import (
	"flag"
	"fmt"

	"time"

	"github.com/bellge/getrooms/core"
	"github.com/bellge/getrooms/tv"
)

var interval = flag.Int("i", 30, "input a integer as interval(min)")

func parse() {
	flag.Parse()
	flag.Usage()

}

func run() {

	fmt.Printf("【每%d分钟抓取一次数据（默认30分钟）】\n", *interval)
	fmt.Println("开始抓取数据...\n")

	for {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		tv.Outputmap["time"] = fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05"))

		tv.Wg.Add(4)
		go tv.Chushou2()
		go tv.Feiyun()
		go tv.Dashen()
		go tv.Qq()
		tv.Wg.Wait()

		core.W2xls()

		fmt.Println("----------------------")
		time.Sleep(time.Minute * time.Duration(*interval))

	}
}

func main() {
	parse()
	run()
}
