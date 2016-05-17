package feed

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/golibri/fetch"
	"github.com/mmcdole/gofeed"
	"strings"
)

// Feed represents an Atom/RSS feed's important data
type Feed struct {
	URL   string
	Body  string
	Items []Item
}

// Item is a single entry within a Feed
type Item struct {
	Title string
	URL   string
	Text  string
	Image string
}

// FromURL parses a RSS/Atom feed directly from a given URL
func FromURL(URL string) (Feed, error) {
	page, err := fetch.Get(URL)
	if err != nil {
		return Feed{URL: URL}, err
	}
	result, err := Parse(page.Body)
	if err != nil {
		return result, err
	}
	result.URL = URL
	result.Body = page.Body
	return result, nil
}

// Parse executes an XML string containing RSS/Atom Data and returns a Feed{}
func Parse(XML string) (Feed, error) {
	fp := gofeed.NewParser()
	data, err := fp.ParseString(XML)
	if err != nil {
		return parseManually(XML)
	}
	f := Feed{Body: XML}
	for _, i := range data.Items {
		item := Item{i.Title, i.Link, i.Description, i.Link}
		item.URL = strings.TrimSpace(item.URL)
		if i.Image != nil {
			item.Image = i.Image.URL
		}
		f.Items = append(f.Items, item)
	}
	return f, nil
}

func parseManually(XML string) (Feed, error) {
	doc := docFromString(XML)
	items := []Item{}
	doc.Find("entry,item").Each(func(i int, s *goquery.Selection) {
		item := Item{URL: findURLfromItem(s)}
		items = append(items, item)
	})
	return Feed{Items: items}, nil
}

func docFromString(str string) goquery.Document {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(str)
	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		return goquery.Document{}
	}
	return *doc
}

func findURLfromItem(s *goquery.Selection) string {
	str, exists := s.Attr("href")
	if exists {
		return str
	}
	str, exists = s.Attr("link")
	if exists {
		return str
	}
	str = s.Find("link").First().Text()
	return str
}

// Links collects all item URLs and returns them as an array
func (f *Feed) Links() []string {
	list := make([]string, len(f.Items))
	for i, item := range f.Items {
		list[i] = item.URL
	}
	return list
}

// NewLinks shows only URLs that are present in this Feed and not in the old one
func (f *Feed) NewLinks(oldFeed *Feed) []string {
	new := f.Links()
	old := oldFeed.Links()
	list := []string{}
	var included bool
	for _, newURL := range new {
		included = false
		for _, oldURL := range old {
			if newURL == oldURL {
				included = true
			}
		}
		if included == false {
			list = append(list, newURL)
		}
	}
	return list
}
