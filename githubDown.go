package main

/*
   下载github 最新 releases 文件到服务器
*/
import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"githubDownload/entity"
	"githubDownload/util"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	//"github.com/robfig/cron"
	"gopkg.in/yaml.v3"
)



// */1 * * * * go run githubDown.go -r /2dust/v2rayN/ -p /caddy/test/
// go run githubDown.go -r /2dust/v2rayN/ -p /caddy/test/
// go run githubDown.go -r /BoyceLig/Clash_Chinese_Patch/ -p /caddy/test/
// go run githubDown.go -r /2dust/v2rayNG/ -p  /caddy/test/  -remove v2rayNG
// go run githubDown.go -r /Kr328/ClashForAndroid/ -p /caddy/test/
// go run githubDown.go -r /Fndroid/clash_for_windows_pkg/ -p /caddy/test/ -n 6,7 -remove Clash

func main() {
	yamlFile, err :=ioutil.ReadFile("/config/conf.yml")
	if err != nil {
		log.Println(err.Error())
	}
	yml:=&Yml{}
	yaml.Unmarshal(yamlFile,yml)
	c := cron.New()
	for _,val := range yml.Jobs {
		option := newDownOption()
		option.CronStr=val.CronStr
		option.Repo=val.Repo
		option.LocalPath=val.LocalPath
		option.NewName=val.NewName
		option.Indexs=val.Indexs
		if len(val.RemvStrs)>0 {
			option.RemvStrs=val.RemvStrs
		}
		log.Println(option)
		////定时任务
		c.AddFunc(option.CronStr, func() {
			downShedule(option)
		})
	}
	c.Start()
	select {} //阻塞主线程停止
}



//downShedule 定时任务
func downShedule(option *DownOption) {
	log.Println("开始执行任务",option.Repo)
	files, _ := ioutil.ReadDir(option.LocalPath)

	for _, f := range files {
		for _, r := range option.RemvStrs {
			if r != "" && strings.Contains(f.Name(), r) {
				fmt.Println("删除关键词" + r + "文件")
				os.Remove(option.LocalPath + f.Name())
			}
		}
	}

	r, err := http.Get("https://api.github.com/repos" + option.Repo + "releases/latest")
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()
	body, _ := ioutil.ReadAll(r.Body)
	var xxm entity.Mybody
	err = json.Unmarshal(body, &xxm)
	indexArr := strings.Split(option.Indexs, ",")
	filename := ""
	for _, i := range indexArr {
		//字符串转int
		i, _ := strconv.Atoi(i)
		url := xxm.Assets[i].BrowserDownloadURL
		if option.NewName != "" {
			filename = option.NewName
		} else {
			filename = xxm.Assets[i].Name
		}

		log.Println("正在下载---------" + filename)
		util.DownloadFileProgress(url, option.LocalPath+filename)
		log.Println("下载完成---------" , filename)
	}
}

type Yml struct {
	Jobs []DownOption `yaml:"jobs"`
}

//DownOption downShedule方法的参数
type DownOption struct {
	Repo string `yaml:"repo"`
	LocalPath string `yaml:"localPath"`
	Indexs string `yaml:"indexs"`
	RemvStrs []string `yaml:"remvStrs"`
	NewName string `yaml:"newName"`
	CronStr string `yaml:"cronStr"`
}

//newDownOption 设置方法参数默认值
func newDownOption() *DownOption {
	return &DownOption{
		Indexs: "0",
	}
}


