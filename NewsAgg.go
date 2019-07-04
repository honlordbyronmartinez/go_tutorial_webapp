package main

import ("fmt"
        "net/http"
        "io/ioutil"
        "encoding/xml"
        "strings"
        "html/template"
       )

type NewsMap struct {
  Keyword string
  Location string
 }

type NewsAggPage struct {
  Title string
  News map[string]NewsMap
 }

type SitemapIndex struct {
  Locations []string `xml:"sitemap>loc"`
}

type News struct {
  Titles []string `xml:"url>news>title"`
  Keywords []string `xml:"url>news>keywords"`
  Locations []string `xml:"url>loc"`
}

func indexHandler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request){
  var s SitemapIndex
  var n News
  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  xml.Unmarshal(bytes, &s)
  news_map := make(map[string]NewsMap)

  for _, Location := range s.Locations {
    Location = strings.TrimSpace(Location)
    resp, err := http.Get(Location)
    if err != nil {
      fmt.Println(err)
    }
    bytes, _ := ioutil.ReadAll(resp.Body)
    xml.Unmarshal(bytes, &n)
    for idx, _ := range n.Keywords{
      news_map[n.Titles[idx]]= NewsMap{n.Keywords[idx], n.Locations[idx]}
    }
  }

  p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}
  t, _ := template.ParseFiles("newsaggtemplate.html")
  t.Execute(w,p)
}

func main()  {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/agg/", newsAggHandler)
  http.ListenAndServe(":8000", nil)
}
