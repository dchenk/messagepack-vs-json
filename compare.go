package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//go:generate msgp -io=false -tests=false

func main() {

	var s Small
	data, err := ioutil.ReadFile("data_small.json")
	if err != nil {
		printErr(err)
		return
	}
	if err = json.Unmarshal(data, &s); err != nil {
		printErr(err)
		return
	}
	marshalledSmallJSON, err := json.Marshal(&s)
	if err != nil {
		printErr(err)
		return
	}
	marshalledSmallMsgp, err := s.MarshalMsg(nil)
	if err != nil {
		printErr(err)
		return
	}
	fmt.Println("Small:")
	fmt.Println("  JSON:", len(marshalledSmallJSON))
	fmt.Println("  MSGP", len(marshalledSmallMsgp))

	var m Medium
	data, err = ioutil.ReadFile("data_medium.json")
	if err != nil {
		printErr(err)
		return
	}
	if err = json.Unmarshal(data, &m); err != nil {
		printErr(err)
		return
	}
	marshalledMedJSON, err := json.Marshal(&m)
	if err != nil {
		printErr(err)
		return
	}
	marshalledMedMsgp, err := m.MarshalMsg(nil)
	if err != nil {
		printErr(err)
		return
	}
	fmt.Println("Medium:")
	fmt.Println("  JSON:", len(marshalledMedJSON))
	fmt.Println("  MSGP", len(marshalledMedMsgp))

	var lg Large
	data, err = ioutil.ReadFile("data_large.json")
	if err != nil {
		printErr(err)
		return
	}
	if err = json.Unmarshal(data, &lg); err != nil {
		printErr(err)
		return
	}
	marshalledLargeJSON, err := json.Marshal(&lg)
	if err != nil {
		printErr(err)
		return
	}
	marshalledLargeMsgp, err := lg.MarshalMsg(nil)
	if err != nil {
		printErr(err)
		return
	}
	fmt.Println("Large:")
	fmt.Println("  JSON:", len(marshalledLargeJSON))
	fmt.Println("  MSGP", len(marshalledLargeMsgp))

}

// Small message schema
type Small struct {
	Hello  string  `json:"hello" msgp:"hello"`
	Name   string  `json:"name" msgp:"name"`
	Age    int     `json:"age" msgp:"age"`
	Weight float64 `json:"weight" msgp:"weight"`
}

// Medium message schema
type Medium []struct {
	Hello   string   `json:"hello" msgp:"hello"`
	Name    string   `json:"name" msgp:"name"`
	Age     int      `json:"age" msgp:"age"`
	Weight  float64  `json:"weight" msgp:"weight"`
	Height  float64  `json:"height" msgp:"height"`
	Hobbies []string `json:"hobbies" msgp:"hobbies"`
	Extra   struct {
		Location string `json:"location" msgp:"location"`
		Bio      string `json:"bio" msgp:"bio"`
	} `json:"extra" msgp:"extra"`
}

// Large message schema
type Large []struct {
	Hello   string   `json:"hello" msgp:"hello"`
	Name    string   `json:"name" msgp:"name"`
	Age     int      `json:"age" msgp:"age"`
	Weight  float64  `json:"weight" msgp:"weight"`
	Height  float64  `json:"height" msgp:"height"`
	Hobbies []string `json:"hobbies" msgp:"hobbies"`
	Extra   struct {
		Location    string `json:"location" msgp:"location"`
		Bio         string `json:"bio" msgp:"bio"`
		Number      int    `json:"number" msgp:"number"`
		NegativeNum int    `json:"negative_num" msgp:"negative_num"`
	} `json:"extra" msgp:"extra"`
}

func printErr(e error) {
	fmt.Println("ERROR:", e)
}
