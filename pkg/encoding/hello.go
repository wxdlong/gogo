package encoding

import (
	"encoding/xml"
)

type Stats struct {
	XMLName        xml.Name
	PageCount      int `xml:"page-count,attr"`
	WordCount      int `xml:"word-count,attr"`
	CharacterCount int `xml:"character-count,attr"`
}

type Meta struct {
	XMLName xml.Name
	Date    string `xml:"creation-date"`
	Title   string `xml:"title"`
	Stats   Stats  `xml:"document-statistic"`
}

type DocumentMeta struct {
	XMLName xml.Name
	Meta    Meta `xml:"meta"`
}

var data = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<office:document-meta>
    <office:meta>
        <meta:creation-date>2016-10-18T15:05:19.137453293</meta:creation-date>
        <dc:title>Document</dc:title>
        <dc:date>2017-01-17T00:59:04.731054728</dc:date>
        <meta:document-statistic meta:page-count="1" meta:word-count="3" meta:character-count="9"/>
    </office:meta>
</office:document-meta>`)
