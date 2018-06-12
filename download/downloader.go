package download

import "net/http"

type Downloader interface{
	Downloader(url string) (*http.Response, error)
	SetClient(client *http.Client)
	GetClient() *http.Client
}