package entity

//Mybody github api响应json对应的结构体
type Mybody struct {
	Assets []Assets `json:"assets"`
}


type Assets struct {
	BrowserDownloadURL string `json:"browser_download_url"`
	Name string `json:"name"`
}

