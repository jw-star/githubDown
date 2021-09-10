package util

import (
	"github.com/schollz/progressbar/v3"
	"io"
	"net/http"
	"os"
)

// DownloadFileProgress 进度条 正在下载 100% |████████████████████████████████| (28/28 MB, 2.898 MB/s)
//DownloadFileProgress 下载在线文件到指定目录
func DownloadFileProgress(url, filename string) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() {_ = r.Body.Close()}()

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bar := progressbar.DefaultBytes(
		r.ContentLength,
		"正在下载",
	)

	io.Copy(io.MultiWriter(f, bar), r.Body)
}




