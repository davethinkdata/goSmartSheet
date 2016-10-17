package goSmartSheet

import (
	"encoding/json"
)

//Cell is a SmartSheet cell
type Cell struct {
	ColumnID     int64      `json:"columnId"`
	Value        *CellValue `json:"value,omitempty"` //TODO: should this be a pointer?
	DisplayValue string     `json:"displayValue,omitempty"`
}

//CellValue represents the possible strongly typed values that could exist in a SS cell
//another good article on it..
//http://attilaolah.eu/2013/11/29/json-decoding-in-go/
type CellValue struct {
	Value json.RawMessage

	StringVal string
	IntVal    int
	FloatVal  float32
}

//MarshalJSON is a custom marshaller for CellValue
func (c *CellValue) MarshalJSON() ([]byte, error) {
	if c.StringVal != "" {
		return json.Marshal(c.StringVal)
	}

	if c.IntVal != 0 {
		return json.Marshal(c.IntVal)
	}

	if c.FloatVal != 0 {
		return json.Marshal(c.FloatVal)
	}

	return json.Marshal(c.Value) //default raw message
}

//UnmarshalJSON is a custom unmarshaller for CellValue
func (c *CellValue) UnmarshalJSON(b []byte) (err error) {
	s := ""
	if err = json.Unmarshal(b, &s); err == nil {
		c.StringVal = s
		return
	}
	var i int
	if err = json.Unmarshal(b, &i); err == nil {
		c.IntVal = i
		return
	}
	var f float32
	if err = json.Unmarshal(b, &f); err == nil {
		c.FloatVal = f
		return
	}

	c.Value = json.RawMessage(b) //default to raw message
	return
}

//MarshalJSON s a custom marshaller to deal with the raw message
// func (c *Cell) MarshalJSON() ([]byte, error) {
// 	b := new(bytes.Buffer)

// 	fmt.Fprintf(b, `{"columnId":`)
// 	var numB []byte
// 	numB = strconv.AppendInt(numB, c.ColumnID, 10)
// 	b.Write(numB)
// 	fmt.Fprintf(b, `,`)

// 	//custom logic for raw message (just get string of bytes)
// 	fmt.Fprintf(b, `"value":"%v"`, string(c.Value))

// 	if c.DisplayValue != "" {
// 		fmt.Fprintf(b, `,"displayValue":"%v",`, string(c.ColumnID))
// 	}

// 	fmt.Fprintf(b, `}`)

// 	//log.Println(string(b.Bytes()))

// 	return b.Bytes(), nil
// }

/* //http://eagain.net/articles/go-dynamic-json/
type CellString struct {
	Value string `json:"value,omitempty"`
}*/
