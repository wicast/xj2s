# DISCONTINUED!!

# XJ2S: A small tool for Golang to generate Golang struct from a xml/json file

The generated struct can be used for a `Unmarshal()` function from a `encoding/xml` or `encoding/json` package

###### Notice:
This tool is for xml/json data parsing,not for modifing xml/json data struct.

If you want to modify a xml/json data,My recommandation is to use `github.com/clbanning/mxj` package,`xj2s`is using this package for backend.

### Installation
you should setup your $GOPATH properely  
`go get github.com/wicast/xj2s/cmd/...`

And a executable called `xmljson2struct` will be appear in your $GOPATH/bin directory.

### Usage
```
Usage: xmljson2struct [-flags] [file]
Flags:
  -f string
    	Parse from a file given a name
  -n	Generate structs whit nesting style
	notice:json haven't implement the > style yet,so only nesting style is working for json
  -root string
    	For struct root name when using json,Default is JsonRoot (default "JsonRoot")
  -t string
    	Type to parse
	avaliable type:xml,json (default "xml")
```

### XML Parsing Example
```
$ curl -s https://github.com/timeline|xmljson2struct 
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

### JSON Parsing Example
Not implemented yet.


## TODO
- [ ] XML conflict issue [#2](https://github.com/wicast/xj2s/issues/2).
- [ ] JSON/XML nesting style.
- [ ] Rust support?
