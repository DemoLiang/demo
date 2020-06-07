package main

import "log"

type Downloader interface {
	Download(uri string)
}

type template struct {
	implement
	uri string
}

type implement interface {
	download()
	save()
}

func newTemplate(impl implement)*template{
	return &template{
		implement:impl,
	}
}

func (t *template)Download(uri string){
	t.uri = uri
	log.Printf("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	log.Printf("finish downloading\n")
}

func (t *template)save(){
	log.Printf("default save\n")
}


type HTTPDownloader struct {
	*template
}

func NewHTTPDownloader() Downloader{
	downloader:=&HTTPDownloader{}
	template:=newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *HTTPDownloader)download(){
	log.Printf("download %s via http\n",d.uri)
}

func (*HTTPDownloader)save(){
	log.Printf("http save\n")
}


type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader{
	downloader:=&FTPDownloader{}
	template:=newTemplate(downloader)
	downloader.template = template
	return downloader
}

func(d *FTPDownloader)download(){
	log.Printf("download %s via ftp\n",d.uri)
}

func main(){
	var downloader Downloader = NewHTTPDownloader()
	downloader.Download("http://example.com/abc.zip")

	var downloader2 Downloader = NewFTPDownloader()
	downloader2.Download("ftp://example.com/abc.zip")
	return
}