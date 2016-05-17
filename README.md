[![Build Status](https://travis-ci.org/golibri/feed.svg?branch=master)](https://travis-ci.org/golibri/website)
[![Code Climate](https://codeclimate.com/github/golibri/feed/badges/gpa.svg)](https://codeclimate.com/github/golibri/website)
[![GoDoc](https://godoc.org/github.com/golibri/feed?status.svg)](https://godoc.org/github.com/golibri/website)
[![Built with Spacemacs](https://cdn.rawgit.com/syl20bnr/spacemacs/442d025779da2f62fc86c2082703697714db6514/assets/spacemacs-badge.svg)](http://github.com/syl20bnr/spacemacs)

# golibri/feed
Parse any RSS/Atom feed and get the relevant article URLs.

# Requirements
`go get -u github.com/golibri/fetch`

`go get -u github.com/PuerkitoBio/goquery`

`go get -u github.com/mmcdole/gofeed`


# installation
`go get -u github.com/golibri/feed`

# usage
````go
import "github.com/golibri/feed"

func main() {
    f := feed.FromURL("http://example.com/whatever")
    // OR:
    f := feed.Parse("feed-xml-string")
    // f is a Feed object, see below

    // to get all article URLs directly:
    links := f.Links() // []string
}
````

# data fields
A **Feed** has the following data fields:

````go
type Feed struct {
    URL   string
    Body  string
    Items []Item
}

type Item struct {
    Title string
    URL   string
    Text  string
    Image string
}
````

# license
LGPLv3. (You can use it in commercial projects as you like, but improvements/bugfixes must flow back to this lib.)
