package webcrawler

import (
	"fmt"
	"sync"
)

/**
- Mutex : protect from concurrent
- WaitGroup : to figure out if goroutine is completed

*/

func TestWebCrawler() {
	Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()
}


type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Void struct {
}
type StringSet map[string]Void

func (s *StringSet) put (key string) {
	(*s)[key] = Void{}
}
func (s *StringSet) remove (key string) {
	delete(*s, key)
}
func (s *StringSet) contains (key string) bool {
	_, ok := (*s)[key]
	return ok
}

type SafeStringSet struct {
	StringSet
	sync.Mutex
}

func (s *SafeStringSet) put (key string) {
	s.Lock()
	defer s.Unlock()

	s.StringSet.put(key)

}
func (s *SafeStringSet) remove (key string) {
	s.Lock()
	defer s.Unlock()
	s.StringSet.remove(key)
}
func (s *SafeStringSet) contains (key string) bool {
	s.Lock()
	defer s.Unlock()
	return s.StringSet.contains(key)
}


type UrlCache struct {
	SafeStringSet
}

var cache UrlCache = UrlCache {SafeStringSet{StringSet{}, sync.Mutex{}}}
var wg sync.WaitGroup
// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// This implementation doesn't do either:

	if depth <= 0 {
		return
	}

	if cache.contains(url) {
		return
	}
	cache.put(url)

	wg.Add(1)
	go func() {
		fetchAndCrawlSubs(url, depth)
		wg.Done()
	}()
}

func fetchAndCrawlSubs(url string, depth int) {
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
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
