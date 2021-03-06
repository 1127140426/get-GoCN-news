package main

import (
	"github.com/greenpipig/get-GoCN-news/GoCN-news"
	"github.com/greenpipig/get-GoCN-news/cron"
	"github.com/greenpipig/get-GoCN-news/getNews"
	"github.com/greenpipig/get-GoCN-news/log"
	"os/exec"
	"time"
)

func mainFunc() {
	url, title := getNews.FetchUrl("https://gocn.vip/topics/node18")
	log.Infof("newsPageUrl:%s,newsPageTitle:%s", url, title)
	if url == "" || title == "" {
		return
	}
	newsList, newsUrlList := getNews.FetchTotalNew(url)
	judgeFirst := GoCN_news.WriteToMd(newsList, newsUrlList, title)
	if judgeFirst {
		command := `./update.sh`
		cmd := exec.Command("/bin/bash", "-c", command)
		_, err := cmd.Output()
		if err != nil {
			log.Infof("Execute Shell:%s failed with error:%s", command, err.Error())
			return
		}
	}
}

func main() {
	cron.ReloadJob("update mdFile", mainFunc, 1*time.Minute)
}
