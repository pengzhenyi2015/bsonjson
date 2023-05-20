# 简介

一个简短的 Golang 测试程序，对比 BSON 和 JSON 在同样数据的场景下，**长度**、**编码性能**、**解码性能** 的差异。

使用的代码库：
- JSON: Golang 内置的 JSON 库
- BSON: MongoDB 官方提供的 go-driver

测试的数据类型：
- double: 一个结构体包含 10 个 double 成员，分 “短” double(数值 0-10) 和 “长” double(数值 math.maxFloat64) 2 种场景
- int64: 一个结构体包含 10 个 long 成员，分 “短” int64(数值 0-10) 和 “长” int64(数值 math.maxInt64) 2 种场景
- string：一个结构体包含 10 个 string 成员，每个 string 成员的长度是 117 字节

测试方法：
将每种类型的数据对象，分别使用 JSON 和 BSON 的方式 marshal/unmarshal 各 100,000 次。
对比 **长度**、**编码耗时**、**解码耗时**。

# 运行方式
1. 下载代码
2. go mod tidy
3. go run main.go

# 测试结果

测试环境：腾讯云 CVM， CPU型号：Intel(R) Xeon(R) Platinum 8255C CPU @ 2.50GHz

```
~/bsonjson$ go run main.go
TestIteration: 100000


--Test double short--
bsonBytes length = 126
jsonBytes length = 73
bson marshal total time(us): 135807
json marshal total time(us): 125041
bson unmarshal total time(us): 208203
json unmarshal total time(us): 280744


--Test double long--
bsonBytes length = 126
jsonBytes length = 292
bson marshal total time(us): 130765
json marshal total time(us): 175027
bson unmarshal total time(us): 199971
json unmarshal total time(us): 471656


--Test int64 short--
bsonBytes length = 126
jsonBytes length = 72
bson marshal total time(us): 135033
json marshal total time(us): 51550
bson unmarshal total time(us): 174061
json unmarshal total time(us): 261838


--Test int64 long--
bsonBytes length = 126
jsonBytes length = 252
bson marshal total time(us): 135817
json marshal total time(us): 91171
bson unmarshal total time(us): 200689
json unmarshal total time(us): 407105


--Test string--
bsonBytes length = 1266
jsonBytes length = 1252
bson marshal total time(us): 187195
json marshal total time(us): 269829
bson unmarshal total time(us): 269802
json unmarshal total time(us): 973015
```

## 总结

|类型|BSON 长度|JSON 长度|BSON-Marshal 单次耗时(us)|JSON-Marshal 单次耗时(us)|BSON-Unmarshal 单次耗时(us)|JSON-Unmarshal 单次耗时(us)|
|:-|:-|:-|:-|:-|:-|:-|
|“短” double * 10|**126**|**73**|**1.4**|**1.3**|2.1|2.8|
|“长” double * 10|126|292|1.3|1.8|2.0|4.7|
|"短" int64 * 10|**126**|**72**|**1.4**|**0.5**|1.7|2.6|
|“长” int64 * 10|126|252|**1.4**|**0.9**|2.0|4.1|
|String(117B) * 10|**1266**|**1252**|1.9|2.7|2.7|9.7|

### 1. 存储空间（长度）效率
对于 “长” double/int64 类型，BSON 在空间上，有一定优势。因为 BSON 使用固定 8 字节进行编码。
对于 “短” double/int64 类型，BSON 在空间上，存在劣势。还是因为 BSON 使用固定 8 字节编码的原因。
对于测试使用的 string 类型，BSON 并没有优势。

### 2. 编码（marshal）效率
对于测试使用的 “长” double 类型和 string 类型， BSON 均有10%-30% 的性能提升。
在 “短” double 类型上，BSON 和 JSON 性能相当。
对于 int64 类型，不管数据的长短，BSON 均有明显的性能下降。需要额外关注，并进一步分析。

### 3. 解码（unmarshal）效率
BSON 优势非常明显。对于测试使用的 “长” double/int64 类型有 100% 的性能提升，对于 string 类型有 200% 以上的性能提升。
对于 “短” double/int64 类型，也有20% 以上的性能提升。