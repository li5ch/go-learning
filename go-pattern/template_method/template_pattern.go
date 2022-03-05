package template_method

import "fmt"

type Downloader interface {
	Download(string)
}

type implement interface {
	download()
	save()
}

// 父类封装不变的方法，并把变化的方法抽象到具体的implement方法中
type template struct {
	implement
	uri string
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

func (t *template) save() {
	fmt.Print("default save")
}

type HTTPDownloader struct {
	*template // 关键：必须有父类的引用
}

func NewHTTPDownloader() Downloader {
	// 注意implement
	d := &HTTPDownloader{}
	// 这里用子类作为参数传入父类的new模板方法
	d.implement = newTemplate(d)
	return d
}

// 重写父类方法即可，不变的模板方法可以继续复用
func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}
