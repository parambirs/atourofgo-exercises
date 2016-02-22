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

func Crawl(url string, depth int, fetcher Fetcher) {
  done := make(chan bool)
  // fmt.Println("Starting to crawl url:", url, "depth:", depth)
  go Crawl_(url, depth, fetcher, done)
  // fmt.Println("Waiting for main done signal")
  <-done
  // fmt.Println("Returning from main Crawl")
  return
}

// Crawl uses fetcher to recursively crawl
// pages starting with a url, to a maximum depth.
func Crawl_(url string, depth int, fetcher Fetcher, done chan bool) {
  if depth <= 0 {
    // fmt.Println("returning early for url:", url, "depth:", depth)
    done <- true
    return
  }

  body, urls, err := fetcher.Fetch(url)
  if err != nil {
    fmt.Println(err)
    done <- true
    return
  }
  fmt.Printf("got: %s %q\n", url, body)

  doneCh := make(chan bool, len(urls))
  // fmt.Println("spawning and waiting for", len(urls), "goroutines")

  for _, u := range urls {
    go Crawl_(u, depth-1, fetcher, doneCh)
  }
  // wait for all goroutines to finish
  for i := 0; i < len(urls); i++ {
    // fmt.Println("Waiting for url:", url, " depth:", depth)
    <-doneCh
  }

  // fmt.Println("returning late for url:", url, "depth: ", depth)
  done <- true
  return
}

func main() {
  Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
  body string
  urls []string
}

var cache = make(map[string]*fakeResult)
var lock = &sync.Mutex{}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
  lock.Lock()
  defer lock.Unlock()
  if cachedRes, ok := cache[url]; ok {
    fmt.Println("Found in cache, url:", url)
    return cachedRes.body, cachedRes.urls, nil
  }
  if res, ok := f[url]; ok {
    fmt.Println("Had to fetch url:", url)
    cache[url] = res
    return res.body, res.urls, nil
  }
  return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher {
  "http://golang.org/": &fakeResult{
    "The Go Programming Language",
    []string{
      "http://golang.org/pkg/",
      "http://golang.org/cmd/",
    },
  },
  "http://golang.org/pkg/": &fakeResult{
    "Packages",
    []string{
      "http://golang.org/",
      "http://golang.org/cmd/",
      "http://golang.org/pkg/fmt/",
      "http://golang.org/pkg/os/",
    },
  },
  "http://golang.org/pkg/fmt/": &fakeResult{
    "Package fmt",
    []string{
      "http://golang.org/",
      "http://golang.org/pkg/",
    },
  },
  "http://golang.org/pkg/os/": &fakeResult{
    "Package os",
    []string{
      "http://golang.org/",
      "http://golang.org/pkg/",
    },
  },
}


/*
> go run exercise-web-crawler.go
Had to fetch url: http://golang.org/
got: http://golang.org/ "The Go Programming Language"
not found: http://golang.org/cmd/
Had to fetch url: http://golang.org/pkg/
got: http://golang.org/pkg/ "Packages"
Had to fetch url: http://golang.org/pkg/os/
got: http://golang.org/pkg/os/ "Package os"
Found in cache, url: http://golang.org/pkg/
got: http://golang.org/pkg/ "Packages"
Found in cache, url: http://golang.org/
got: http://golang.org/ "The Go Programming Language"
not found: http://golang.org/cmd/
not found: http://golang.org/cmd/
Had to fetch url: http://golang.org/pkg/fmt/
got: http://golang.org/pkg/fmt/ "Package fmt"
Found in cache, url: http://golang.org/pkg/
got: http://golang.org/pkg/ "Packages"
Found in cache, url: http://golang.org/
got: http://golang.org/ "The Go Programming Language"
Found in cache, url: http://golang.org/pkg/
got: http://golang.org/pkg/ "Packages"
Found in cache, url: http://golang.org/
got: http://golang.org/ "The Go Programming Language"
*/