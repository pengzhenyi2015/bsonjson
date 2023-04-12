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