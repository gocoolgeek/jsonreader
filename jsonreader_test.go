package jsonreader

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadConfigs(t *testing.T) {
	var err error
	var validJSON1 string
	validJSON1 = "{    \"glossary\": {" +
		"\"title\": \"example glossary\"," +
		"\"GlossDiv\": {" +
		"\"title\": \"S\"," +
		"\"GlossList\": {" +
		"\"GlossEntry\": {" +
		"\"ID\": \"SGML\"," +
		"\"SortAs\": \"SGML\"," +
		"\"GlossTerm\": \"Standard Generalized Markup Language\"," +
		"\"Acronym\": \"SGML\"," +
		"\"Abbrev\": \"ISO 8879:1986\"," +
		"\"GlossDef\": {" +
		"\"para\": \"A meta-markup language, used to create markup languages such as DocBook.\"," +
		"\"GlossSeeAlso\": [\"GML\", \"XML\"]" +
		"}," +
		"\"GlossSee\": \"markup\"" +
		"}" +
		"}" +
		"}" +
		"}" +
		"}"

	var validJSON2 string
	validJSON2 = "{" +
		"\"dbconfig\" : {" +
		"\"username\" : \"name\"," +
		"\"password\" : \"pwd\"," +
		"\"database\" : \"blah\"" +
		"}" +
		"}"

	var validJSON3 string
	validJSON3 = "{\"language\" : \"Go\"}"

	os.MkdirAll("/tmp/conf", 0777)

	d1 := []byte(validJSON1)
	err = ioutil.WriteFile("/tmp/conf/json1.json", d1, 0644)
	if err != nil {
		t.Error(err)
	}

	d2 := []byte(validJSON2)
	err = ioutil.WriteFile("/tmp/conf/json2.json", d2, 0644)
	if err != nil {
		t.Error(err)
	}

	d3 := []byte(validJSON3)
	err = ioutil.WriteFile("/tmp/conf/json3.json", d3, 0644)
	if err != nil {
		t.Error(err)
	}

	err = Load("/tmp/conf/")
	if err != nil {
		t.Error(err)
	}

	glossary := GetMap("glossary")
	glossdiv := TransformInterfaceToMap(glossary["GlossDiv"])
	glosslist := TransformInterfaceToMap(glossdiv["GlossList"])
	glossentry := TransformInterfaceToMap(glosslist["GlossEntry"])
	glosssee := glossentry["GlossSee"]
	if glosssee != "markup" {
		t.Error(errors.New("GlossSee should have a value of markup"))
	}

	glossary1 := GetMap("blahglah")
	if len(glossary1) != 0 {
		t.Error(errors.New("blahglah should be empty, GetMap failing"))
	}

	glosssee1 := GetValue("glossary.GlossDiv.GlossList.GlossEntry.GlossSee")
	if glosssee1 != "markup" {
		t.Error(errors.New("GlossSee should have a value of markup, GetValue failing"))
	}

	glosssee2 := GetValue("glossary.GlossDiv.GlossList.GlossEntry.GlossSee.Blah")
	if glosssee2 != "" {
		t.Error(errors.New("GlossSee should be nil, GetValue failing"))
	}

	lname := GetValue("language")
	if lname != "Go" {
		t.Error(errors.New("language should have a value of Go, GetValue failing"))
	}

	err = Load("/tmp/conf/json2.json")
	if err != nil {
		t.Error(err)
	}

	dbconfig := GetMap("dbconfig")
	database := dbconfig["database"]
	if database != "blah" {
		t.Error(errors.New("Database should have a value of blah, GetMap failing"))
	}
	fname := GetValue("language")
	if fname != "Go" {
		t.Error(errors.New("language should have a value of Go, GetValue failing"))
	}

	os.RemoveAll("/tmp/conf/")
}
