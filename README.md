# 简介

一个简短的 Golang 测试程序，对比 BSON 和 JSON 在同样数据的场景下，**长度**、**编码性能**、**解码性能** 的差异。

使用的代码库：
- JSON: Golang 内置的 JSON 库
- BSON: MongoDB 官方提供的 go-driver

测试的数据类型：
- double: 一个结构体包含 10 个 double 成员
- string：一个结构体包含 10 个 string 成员，每个 string 成员的长度是 1KB

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


--Test double--
bsonBytes length = 126
jsonBytes length = 161
bson marshal total time(us): 137548
json marshal total time(us): 162248
bson unmarshal total time(us): 198119
json unmarshal total time(us): 388782


--Test string--
bsonBytes length = 1266
jsonBytes length = 1252
bson marshal total time(us): 180113
json marshal total time(us): 268086
bson unmarshal total time(us): 273701
json unmarshal total time(us): 979400
```

## 总结
### 1. 存储空间（长度）效率
对于测试使用的 “长” double 类型，BSON 的空间有一定优势。因为 BSON 使用固定字节进行编码。

对于测试使用的 string 类型，BSON 并没有优势。

### 2. 编码（marshal）效率
对于测试使用的 “长” double 类型和 string 类型， BSON 均有10%-30% 的性能提升。

### 3. 解码（unmarshal）效率
BSON 优势非常明显。对于测试使用的 “长” double 类型有 100% 的性能提升，对于 string 类型有 200% 以上的性能提升。