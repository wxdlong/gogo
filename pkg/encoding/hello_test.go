package encoding

import (
	"encoding/xml"
	"fmt"
	"testing"
)

type ABC struct {
	Root xml.Name `xml:"Wxdlong"`
	Comm xml.Comment
}

type Address struct {
	City, State string
}
type Person struct {
	XMLName   xml.Name `xml:"urn:3gpp:metadata:2009:MBMS:schemaVersion person"`
	Id        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Age       int      `xml:"age"`
	Height    float32  `xml:"height,omitempty"`
	Married   bool
	Address
	Comment string `xml:",comment"`
}

//输出xml带域名空间, 和注释
//<person xmlns="urn:3gpp:metadata:2009:MBMS:schemaVersion" id="13">
//<name>
//<first>John</first>
//<last>Doe</last>
//</name>
//<age>42</age>
//<Married>false</Married>
//<City>Hanga Roa</City>
//<State>Easter Island</State>
//<!-- Need more details. -->
//</person>
func TestXml1(t *testing.T) {

	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println(string(output))
}
