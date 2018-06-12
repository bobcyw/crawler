package singleThread

import "net/http"

type SDownloader struct{
	client *http.Client
}

func (inst *SDownloader) Downloader(url string) (*http.Response, error){
	return inst.GetClient().Get(url)
}

func (inst *SDownloader) SetClient(client *http.Client){
	inst.client = client
}

func (inst *SDownloader) GetClient() *http.Client{
	if inst.client == nil{
		inst.client = http.DefaultClient
	}
	return inst.client
}