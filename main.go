package main

import (
	"flag"
	"fmt"
	"github.com/canhlinh/hlsdl"
	"os"
)

var(
	m3u8Url        string // url地址
	filename       string // 文件名
	coroutineCount int    // 协程数
)

func init() {
	flag.StringVar(&m3u8Url,"u","","url address")
	flag.StringVar(&filename,"n","","filename")
	flag.IntVar(&coroutineCount,"c",64,"coroutines count")

	// 检查文件夹并建立
	os.Mkdir("./download",os.ModePerm)
}

func main() {
	flag.Parse()
	if m3u8Url =="" {
		fmt.Println("url is empty")
		return
	}
	if filename=="" {
		fmt.Println("filename is empty")
		return
	}

	hlsDL := hlsdl.New(m3u8Url, nil, filename, coroutineCount,true)
	filePath , err := hlsDL.Download()
	if err != nil {
		panic(err)
	}

	os.Rename("./"+filePath, "download/"+filename+".ts")
	os.Remove("./"+filename)

	fmt.Println(filename + " download done")
}
