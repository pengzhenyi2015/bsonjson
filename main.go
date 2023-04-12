package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/pprof"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	TestIteration = 100000
	ProfFile      = "cpu.prof"
)

func testString() {
	type stringStruct struct {
		S1  string
		S2  string
		S3  string
		S4  string
		S5  string
		S6  string
		S7  string
		S8  string
		S9  string
		S10 string
	}
	SS := stringStruct{
		S1:  "Mrs Erlynne, the mother of Lady Windermere - her daughter does not know about her - wants to be introduced in society",
		S2:  "Mrs Erlynne, the mother of Lady Windermere - her daughter does not know about her - wants to be introduced in society",
		S3:  "Mrs Erlynne, the mother of Lady Windermere - her daughter does not know about her - wants to be introduced in society",
		S4:  "Mrs Erlynne, the mother of Lady Windermere - her daughter does not know about her - wants to be introduced in society",
		S5:  "Mrs Erlynne, the mother of Lady Windermere - her daughter does not know about her - wants to be introduced in society",
		S6:  "Mrs Erlynne, the mother of Lady Windermere - her daughter does not know about her - wants to be introduced in society",
		S7:  "Mrs Erlynne, the mother of Lady Windermere - her daughter does not know about her - wants to be introduced in society",
		S8:  "Mrs Erlynne, the mother of Lady Windermere - her daughter does not know about her - wants to be introduced in society",
		S9:  "Mrs Erlynne, the mother of Lady Windermere - her daughter does not know about her - wants to be introduced in society",
		S10: "Mrs Erlynne, the mother of Lady Windermere - her daughter does not know about her - wants to be introduced in society",
	}

	bsonBytes, err := bson.Marshal(SS)
	if err != nil {
		fmt.Printf("bson marshal failed: %v\n", err)
		return
	}
	jsonBytes, err := json.Marshal(SS)
	if err != nil {
		fmt.Printf("json marshal failed: %v\n", err)
		return
	}
	fmt.Printf("bsonBytes length = %d\njsonBytes length = %d\n",
		len(bsonBytes), len(jsonBytes))

	// Marshal bson
	t1 := time.Now().UnixNano()
	for i := 0; i < TestIteration; i++ {
		bson.Marshal(SS)
	}

	// Marshal json
	t2 := time.Now().UnixNano()
	for i := 0; i < TestIteration; i++ {
		json.Marshal(SS)
	}

	// Unmarshal bson
	t3 := time.Now().UnixNano()
	for i := 0; i < TestIteration; i++ {
		_ = bson.Unmarshal(bsonBytes, &SS)
	}

	// Unmarshal json
	t4 := time.Now().UnixNano()
	for i := 0; i < TestIteration; i++ {
		_ = json.Unmarshal(jsonBytes, &SS)
	}

	t5 := time.Now().UnixNano()

	fmt.Printf("bson marshal total time(us): %v\n", (t2-t1)/1000)
	fmt.Printf("json marshal total time(us): %v\n", (t3-t2)/1000)
	fmt.Printf("bson unmarshal total time(us): %v\n", (t4-t3)/1000)
	fmt.Printf("json unmarshal total time(us): %v\n", (t5-t4)/1000)

}

func testDouble() {
	type doubleStruct struct {
		D1  float64
		D2  float64
		D3  float64
		D4  float64
		D5  float64
		D6  float64
		D7  float64
		D8  float64
		D9  float64
		D10 float64
	}
	DS := doubleStruct{
		D1:  123456789.0,
		D2:  1234567.89,
		D3:  123456.789,
		D4:  12345.6789,
		D5:  1234.56789,
		D6:  123.456789,
		D7:  12.3456789,
		D8:  1.23456789,
		D9:  0.123456789,
		D10: 123456789.0,
	}

	bsonBytes, err := bson.Marshal(DS)
	if err != nil {
		fmt.Printf("bson marshal failed: %v\n", err)
		return
	}
	jsonBytes, err := json.Marshal(DS)
	if err != nil {
		fmt.Printf("json marshal failed: %v\n", err)
		return
	}
	fmt.Printf("bsonBytes length = %d\njsonBytes length = %d\n",
		len(bsonBytes), len(jsonBytes))

	// Marshal bson
	t1 := time.Now().UnixNano()
	for i := 0; i < TestIteration; i++ {
		bson.Marshal(DS)
	}

	// Marshal json
	t2 := time.Now().UnixNano()
	for i := 0; i < TestIteration; i++ {
		json.Marshal(DS)
	}

	// Unmarshal bson
	t3 := time.Now().UnixNano()
	for i := 0; i < TestIteration; i++ {
		_ = bson.Unmarshal(bsonBytes, &DS)
	}

	// Unmarshal json
	t4 := time.Now().UnixNano()
	for i := 0; i < TestIteration; i++ {
		_ = json.Unmarshal(jsonBytes, &DS)
	}

	t5 := time.Now().UnixNano()

	fmt.Printf("bson marshal total time(us): %v\n", (t2-t1)/1000)
	fmt.Printf("json marshal total time(us): %v\n", (t3-t2)/1000)
	fmt.Printf("bson unmarshal total time(us): %v\n", (t4-t3)/1000)
	fmt.Printf("json unmarshal total time(us): %v\n", (t5-t4)/1000)
}

func main() {
	fmt.Printf("TestIteration: %v\n", TestIteration)

	fileHandler, err := os.Create(ProfFile)
	if err != nil {
		fmt.Printf("create prof file failed:%v\n", err)
		return
	}
	if err = pprof.StartCPUProfile(fileHandler); err != nil {
		fmt.Printf("start cpu prof failed:%v\n", err)
		return
	}

	defer func() {
		pprof.StopCPUProfile()
		_ = fileHandler.Close()
	}()

	fmt.Println("\n\n--Test double--")
	testDouble()

	fmt.Println("\n\n--Test string--")
	testString()

}
