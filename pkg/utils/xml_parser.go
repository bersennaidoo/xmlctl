package utils

import (
	"fmt"

	"github.com/beevik/etree"
)

type XmlParser struct {
	Data []byte
}

func (xp XmlParser) Parse() {

	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(xp.Data); err != nil {
		panic(err)
	}

  root := doc.SelectElement("bookstore")
  fmt.Println("ROOT element:", root.Tag)

	for _, book := range root.SelectElements("book") {
		fmt.Println("CHILD element:", book.Tag)
		if title := book.SelectElement("title"); title != nil {
			lang := title.SelectAttrValue("lang", "unknown")
			fmt.Printf("  TITLE: %s (%s)\n", title.Text(), lang)
		}
		for _, attr := range book.Attr {
			fmt.Printf("  ATTR: %s=%s\n", attr.Key, attr.Value)
		}
	}
}
