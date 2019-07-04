package main

import ("fmt"
        "net/http"
        "io/ioutil"
        "encoding/xml"
        "strings"
       )

type SitemapIndex struct {
  Locations []string `xml:"sitemap>loc"`
}

type News struct {
  Titles []string `xml:"url>news>title"`
  Keywords []string `xml:"url>news>keywords"`
  Locations []string `xml:"url>loc"`
}

type NewsMap struct {
  Keyword string
  Location string
}

func main(){
  var s SitemapIndex
  var n News
  news_map := make(map[string]NewsMap)

  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  xml.Unmarshal(bytes, &s)

  for _, Location := range s.Locations {
    Location = strings.TrimSpace(Location)
		resp, err := http.Get(Location)
		if err != nil {
			fmt.Println(err)
		}
    bytes, _ := ioutil.ReadAll(resp.Body)
    xml.Unmarshal(bytes, &n)
    for idx, _ := range n.Titles{
      news_map[n.Titles[idx]]= NewsMap{n.Keywords[idx], n.Titles[idx]}
    }
  }
  for idx, data:= range news_map{
    fmt.Println("\n\n\n",idx)
    fmt.Println("\n",data.Keyword)
    fmt.Println("\n",data.Location)

  }
}
