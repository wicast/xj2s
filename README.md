# XJ2S: A small tool for Golang to generate Golang struct form a xml/json file

The generated struct can be used for a `Unmarshal()` function from a `encoding/xml` or `encoding/json` package

### Installation
you should setup your $GOPATH properely  
`go get github.com/wicast/xj2s/cmd/...`

And a executable called `xmljson2struct` will be appear in your $GOPATH/bin directory.

### Usage
```
Usage: xmljson2struct [-flags] [file]
  -f string
    	Parse from a file given a name
  -stdin
    	Parse from stdin
  -t string
    	Type to parses
	avaliable type:xml,json (default "xml")
```

### XML Parsing Example
```
$ curl -s https://github.com/timeline|xmljson2struct -stdin
type Feed struct {
	TitleEntry	[]Title	`xml:"entry>title"`
	Email	[]string	`xml:"entry>author>email"`
	Id	string	`xml:"id"`
	Title	string	`xml:"title"`
	Content	[]Content	`xml:"entry>content"`
	Published	[]string	`xml:"entry>published"`
	Updated	string	`xml:"updated"`
	IdEntry	[]string	`xml:"entry>id"`
	Thumbnail	[]Thumbnail	`xml:"entry>thumbnail"`
	Xmlns	string	`xml:"xmlns,attr"`
	UpdatedEntry	[]string	`xml:"entry>updated"`
	Name	[]string	`xml:"entry>author>name"`
	Media	string	`xml:"media,attr"`
	Link	[]Link	`xml:"link"`
	LinkEntry	[]Link	`xml:"entry>link"`
	Uri	[]string	`xml:"entry>author>uri"`
	Lang	string	`xml:"lang,attr"`
}

type Link struct {
	Type	string	`xml:"type,attr"`
	Rel	string	`xml:"rel,attr"`
	Href	string	`xml:"href,attr"`
}
type Content struct {
	Type	string	`xml:"type,attr"`
	Text	string	`xml:",chardata"`
}
type LinkEntry struct {
	Type	string	`xml:"type,attr"`
	Rel	string	`xml:"rel,attr"`
	Href	string	`xml:"href,attr"`
}
type TitleEntry struct {
	Type	string	`xml:"type,attr"`
	Text	string	`xml:",chardata"`
}
type Thumbnail struct {
	Url	string	`xml:"url,attr"`
	Height	string	`xml:"height,attr"`
	Width	string	`xml:"width,attr"`
}
```

The programe will use as less struct as possible,no nesting struct.The first entry is the Root struct.In the Root struct,every field present a useful node in xml.  
The rest struct is used for those having attributes.
