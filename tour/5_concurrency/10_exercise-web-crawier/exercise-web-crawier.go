/* 练习：Web 爬虫
在这个练习中，我们将会使用 Go 的并发特性来并行化一个 Web 爬虫。

修改 Crawl 函数来并行地抓取 URL，并且保证不重复。

提示：你可以用一个 map 来缓存已经获取的 URL，但是要注意 map 本身并不是并发安全的！ */


package main

import (
  "fmt"
  "sync"
)
type Fetcher interface {
  // Fetch returns the body of URL and
  // a slice of URLs found on that page.
  Fetch(url string) (body string, urls []string, err error)
}
var cache = make(Cache)
var wg sync.WaitGroup // required to let all launched goroutines finish before main() exits
var mux sync.Mutex
// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
  defer wg.Done()
  if cache.get(url) {
    fmt.Printf("xx Skipping: %s\n", url)
    return
  }
  fmt.Printf("** Crawling: %s\n", url)
  cache.set(url, true)
 
  if depth <= 0 {
    return
  }
  body, urls, err := fetcher.Fetch(url)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("found: %s %q\n", url, body)
  for _, u := range urls {
    wg.Add(1)
    go Crawl(u, depth-1, fetcher)
  }
  return
}
func main() {
  wg.Add(1)
  Crawl("https://golang.org/", 4, fetcher)
  wg.Wait()
}
type Cache map[string]bool
func (ch Cache) get(key string) bool {
  mux.Lock()
  defer mux.Unlock()
  return cache[key]
}
func (ch Cache) set(key string, value bool) {
  mux.Lock()
  defer mux.Unlock()
  cache[key] = value
}
// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult
type fakeResult struct {
  body string
  urls []string
}
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
  if res, ok := f[url]; ok {
    return res.body, res.urls, nil
  }
  return "", nil, fmt.Errorf("not found: %s", url)
}
// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
  "https://golang.org/": &fakeResult{
   "The Go Programming Language",
  []string{
   "https://golang.org/pkg/",
   "https://golang.org/cmd/",
  },
 },
 "https://golang.org/pkg/": &fakeResult{
  "Packages",
  []string{
   "https://golang.org/",
   "https://golang.org/cmd/",
   "https://golang.org/pkg/fmt/",
   "https://golang.org/pkg/os/",
  },
 },
 "https://golang.org/pkg/fmt/": &fakeResult{
  "Package fmt",
  []string{
   "https://golang.org/",
   "https://golang.org/pkg/",
  },
 },
 "https://golang.org/pkg/os/": &fakeResult{
  "Package os",
  []string{
   "https://golang.org/",
   "https://golang.org/pkg/",
  },
 },
}

/* 
** Crawling: https://golang.org/
found: https://golang.org/ "The Go Programming Language"
** Crawling: https://golang.org/cmd/
not found: https://golang.org/cmd/
** Crawling: https://golang.org/pkg/
found: https://golang.org/pkg/ "Packages"
** Crawling: https://golang.org/pkg/os/
found: https://golang.org/pkg/os/ "Package os"
xx Skipping: https://golang.org/pkg/
xx Skipping: https://golang.org/
** Crawling: https://golang.org/pkg/fmt/
found: https://golang.org/pkg/fmt/ "Package fmt"
xx Skipping: https://golang.org/pkg/
xx Skipping: https://golang.org/cmd/
xx Skipping: https://golang.org/
xx Skipping: https://golang.org/
*/