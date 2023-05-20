package main

import (
	"encoding/json"
	"fmt"
	"math"
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

func testInt64(isShort bool) {
	type int64Struct struct {
		L1  int64
		L2  int64
		L3  int64
		L4  int64
		L5  int64
		L6  int64
		L7  int64
		L8  int64
		L9  int64
		L10 int64
	}
	var LS int64Struct
	if isShort {
		LS = int64Struct{
			L1:  1,
			L2:  2,
			L3:  3,
			L4:  4,
			L5:  5,
			L6:  6,
			L7:  7,
			L8:  8,
			L9:  9,
			L10: 0,
		}
	} else {
		// maxInt64: 9223372036854775807, 19位
		LS = int64Struct{
			L1:  math.MaxInt64,
			L2:  math.MaxInt64,
			L3:  math.MaxInt64,
			L4:  math.MaxInt64,
			L5:  math.MaxInt64,
			L6:  math.MaxInt64,
			L7:  math.MaxInt64,
			L8:  math.MaxInt64,
			L9:  math.MaxInt64,
			L10: math.MaxInt64,
		}
	}

	bsonBytes, err := bson.Marshal(LS)
	if err != nil {
		fmt.Printf("bson marshal failed: %v\n", err)
		return
	}
	jsonBytes, err := json.Marshal(LS)
	if err != nil {
		fmt.Printf("json marshal failed: %v\n", err)
		return
	}
	fmt.Printf("bsonBytes length = %d\njsonBytes length = %d\n",
		len(bsonBytes), len(jsonBytes))

	// Marshal bson
	t1 := time.Now().UnixNano()
	for i := 0; i < TestIteration; i++ {
		bson.Marshal(LS)
	}

	// Marshal json
	t2 := time.Now().UnixNano()
	for i := 0; i < TestIteration; i++ {
		json.Marshal(LS)
	}

	// Unmarshal bson
	t3 := time.Now().UnixNano()
	for i := 0; i < TestIteration; i++ {
		_ = bson.Unmarshal(bsonBytes, &LS)
	}

	// Unmarshal json
	t4 := time.Now().UnixNano()
	for i := 0; i < TestIteration; i++ {
		_ = json.Unmarshal(jsonBytes, &LS)
	}

	t5 := time.Now().UnixNano()

	fmt.Printf("bson marshal total time(us): %v\n", (t2-t1)/1000)
	fmt.Printf("json marshal total time(us): %v\n", (t3-t2)/1000)
	fmt.Printf("bson unmarshal total time(us): %v\n", (t4-t3)/1000)
	fmt.Printf("json unmarshal total time(us): %v\n", (t5-t4)/1000)
}

func testDouble(isShort bool) {
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
	var DS doubleStruct
	if isShort {
		DS = doubleStruct{
			D1:  1,
			D2:  2,
			D3:  3,
			D4:  4,
			D5:  5,
			D6:  6,
			D7:  7,
			D8:  8,
			D9:  9,
			D10: 10,
		}
	} else {
		// MaxFloat64: 1.79769313486231570814527423731704356798070e+308
		DS = doubleStruct{
			D1:  math.MaxFloat64,
			D2:  math.MaxFloat64,
			D3:  math.MaxFloat64,
			D4:  math.MaxFloat64,
			D5:  math.MaxFloat64,
			D6:  math.MaxFloat64,
			D7:  math.MaxFloat64,
			D8:  math.MaxFloat64,
			D9:  math.MaxFloat64,
			D10: math.MaxFloat64,
		}
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
	//fmt.Printf("jsonBytes: %s\n", string(jsonBytes))
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

	fmt.Println("\n\n--Test double short--")
	testDouble(true)

	fmt.Println("\n\n--Test double long--")
	testDouble(false)

	fmt.Println("\n\n--Test int64 short--")
	testInt64(true)

	fmt.Println("\n\n--Test int64 long--")
	testInt64(false)

	fmt.Println("\n\n--Test string--")
	testString()

}
