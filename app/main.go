package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

func main() {
	url := "https://golang.org/"

	// Getリクエストの送信
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s\n", res.StatusCode, res.Status)
	}

	// 読み取り
	buf, _ := ioutil.ReadAll(res.Body)

	// 文字コード判定
	det := chardet.NewTextDetector()
	detResult, _ := det.DetectBest(buf)
	log.Println(detResult.Charset)

	// 文字コード変換
	bufReader := bytes.NewReader(buf)
	reader, _ := charset.NewReaderLabel(detResult.Charset, bufReader)

	// HTML の ドキュメント取得
	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		log.Fatalln(err)
	}

	// テキスト取得
	result := doc.Find("title").Text()

	// 特定の属性値抜き出し
	//result, _ := doc.Find("iframe").Attr("src")

	fmt.Println(result)
}
