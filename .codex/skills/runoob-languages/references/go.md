# Go - 菜鸟教程

Tutorial: https://www.runoob.com/go/go-tutorial.html

---

## Go 语言教程

Source: https://www.runoob.com/go/go-tutorial.html

## Go 语言教程

Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。

Go是从2007年末由Robert Griesemer, Rob Pike, Ken Thompson主持开发，后来还加入了Ian Lance Taylor, Russ Cox等人，并最终于2009年11月开源，在2012年早些时候发布了Go 1稳定版本。现在Go的开发已经是完全开放的，并且拥有一个活跃的社区。

### Go 语言特色

- 简洁、快速、安全
- 并行、有趣、开源
- 内存管理、数组安全、编译迅速

### Go 语言用途

Go 语言被设计成一门应用于搭载 Web 服务器，存储集群或类似用途的巨型中央服务器的系统编程语言。

对于高性能分布式系统领域而言，Go 语言无疑比大多数其它语言有着更高的开发效率。它提供了海量并行的支持，这对于游戏服务端的开发而言是再好不过了。

### 第一个 Go 程序

接下来我们来编写第一个 Go 程序 hello.go（Go 语言源文件的扩展是 .go），代码如下：

### hello.go 文件

package main

import "fmt"

func main() {
fmt.Println("Hello, World!")
}

运行实例 »

要执行 Go 语言代码可以使用 go run 命令。

执行以上代码输出:

```

$ go run hello.go
Hello, World!

```

此外我们还可以使用 go build 命令来生成二进制文件：

```
$ go build hello.go
$ ls
hello hello.go
$ ./hello
Hello, World!
```

---

## Go 简介

Source: https://www.runoob.com/go/go-intro.html

## Go 简介

Go（又称 Golang）是 Google 开发的一种开源编程语言，设计目标是让开发者能够高效地构建简单、可靠且高性能的软件。

Go 语言诞生于 2007 年，由 Google 的三位资深工程师——Robert Griesemer、Rob Pike 和 Ken Thompson 联合设计，并于 2009 年正式开源发布。

Go 的设计初衷是解决 Google 内部面临的大规模软件开发痛点：C++ 编译太慢、Java 运行时太臃肿、脚本语言性能不足。三位设计者希望创造一门兼具编译型语言的性能和动态语言的开发效率的编程语言。

Go 语言吉祥物是一只名为 Gopher 的土拨鼠，象征着 Go 社区的务实与活力。

Go 语言的关键设计目标可以归纳为以下几点：

设计目标说明 编译速度大型项目的编译应在几秒内完成，而非几分钟 并发支持语言级别原生支持并发编程，无需第三方库 简洁性语法最少化，没有类和继承，减少认知负担 内存安全自动垃圾回收（GC），避免手动内存管理的风险 静态类型编译期类型检查，减少运行时错误 跨平台一套代码可编译为 Windows、Linux、macOS 等多平台可执行文件

#### Go 诞生要解决的三个核心问题

问题传统语言的困境Go 的解决方式 编译速度慢大型 C++ 项目编译动辄数分钟乃至数小时依赖管理设计极简，编译速度极快，通常秒级完成 并发编程复杂线程、锁、回调的组合让并发代码难以编写和维护内置 goroutine 和 channel，并发如同写顺序代码 代码难以维护C++ 特性过多、Java 过度设计，大型团队协作困难极简语法，强制格式化，一种问题只有一种写法

#### Go 的核心设计哲学

Go 的官方口号是 Simple, Fast, Reliable，这三个词概括了它全部的设计取舍。

Go 刻意保持语言规范的精简。整个语言规范只有约 60 页，关键字仅有 25 个（相比之下 C++ 有 80+ 个关键字），这意味着一个有经验的开发者在一两天内就能掌握语法全貌，之后将精力集中在业务逻辑上。

### Go 语言的发展历史

Go 语言的发展历程并非一蹴而就，下面梳理了从立项到如今的关键时间节点。

时间版本/事件说明 2007 年项目立项Rob Pike、Ken Thompson、Robert Griesemer 开始设计 Go 语言 2009 年 11 月Go 首次开源以 BSD 协议在 GitHub 上公开发布 2012 年 3 月Go 1.0首个稳定版本发布，承诺向后兼容 2015 年 8 月Go 1.5编译器用 Go 自举（此前用 C 编写），移除 C 依赖 2018 年 8 月Go 1.11引入 Go Modules，初步解决依赖管理问题 2020 年 2 月Go 1.14Go Modules 成为官方推荐的依赖管理方案 2022 年 3 月Go 1.18引入泛型（Generics），这是 Go 发布以来最大的语言特性变更 2024 年 2 月Go 1.22for 循环变量语义优化，增强路由模式匹配 2025 年 8 月Go 1.24持续改进运行时性能、工具链体验与标准库

Go 1.0 发布时，核心团队做出了一个重大承诺：后续所有 1.x 版本的 Go 均保持向后兼容。这意味着十年前编写的 Go 代码，在今天的最新编译器中依然可以直接编译运行。

### Go 语言的核心特性

Go 的设计哲学决定了它的特性集合——精简而强大。以下逐一拆解 Go 最核心的语言特性。

#### 原生并发：Goroutine 与 Channel

并发是 Go 语言最引人注目的特性。Goroutine 是一种比操作系统线程轻量得多的执行单元——每个 Goroutine 的初始栈大小仅为约 2 KB，而操作系统线程通常在 1 MB 以上。

你可以在一台普通机器上同时运行数十万个 Goroutine，这是传统线程模型无法做到的。

Channel（通道）是 Goroutine 之间的通信管道，遵循 Go 的并发哲学：

不要通过共享内存来通信，而要通过通信来共享内存。（Don't communicate by sharing memory; share memory by communicating.）

一个简单的 Goroutine 与 Channel 示例：

### 实例

package main

import (
"fmt"
"time"
)

// worker 模拟一个工作协程，从通道接收任务并处理
func worker(id int, jobs <-chan int, results chan<- int) {
for j := range jobs {
fmt.Printf("worker %d 开始处理任务 %d\n", id, j)
time.Sleep(time.Second) // 模拟耗时操作
results <- j * 2 // 将处理结果发送到 results 通道
fmt.Printf("worker %d 完成任务 %d\n", id, j)
}
}

func main() {
const numJobs = 5
jobs := make(chan int, numJobs) // 创建带缓冲的任务通道
results := make(chan int, numJobs) // 创建带缓冲的结果通道

// 启动 3 个 worker goroutine
for w := 1; w <= 3; w++ {
go worker(w, jobs, results)
}

// 发送 5 个任务到 jobs 通道
for j := 1; j <= numJobs; j++ {
jobs <- j
}
close(jobs) // 关闭通道，通知 worker 不会再有新任务

// 收集所有处理结果
for a := 1; a <= numJobs; a++ {
fmt.Printf("结果: %d\n", <-results)
}
}

#### 静态类型与类型推断

Go 是静态类型语言，所有变量类型在编译时即确定。但与 C/Java 不同的是，Go 支持类型推断——编译器可以自动推导变量类型，开发者无需反复声明类型名。

这种设计在保持类型安全的同时，让代码写起来像动态语言一样简洁。

### 实例

package main

import "fmt"

func main() {
// 完整声明：显式指定类型
var name string = "runoob"

// 类型推断：编译器自动推导类型
var version = 1.24 // 推导为 float64
count := 100 // 短变量声明，推导为 int
message := "Hello Go" // 推导为 string

// 多变量声明
x, y := 10, 20 // 同时声明多个变量

fmt.Println(name, version, count, message, x, y)
}

#### 快速编译

Go 的编译速度极快，大型项目（几十万行代码）也能在几秒内完成编译。

这得益于 Go 精心设计的依赖分析机制：只编译实际被引用的包，且消除了 C/C++ 中头文件重复解析的开销。

#### 内置垃圾回收（GC）

Go 使用并发三色标记-清扫算法实现自动垃圾回收。GC 与用户代码并发执行，暂停时间通常在微秒级别，不会对服务响应时间产生明显影响。

开发者无需手动管理内存，同时又能获得接近 C 语言的运行时性能。

#### 丰富的标准库

Go 的标准库覆盖了现代开发的大部分需求，无需引入第三方依赖即可完成常见任务。下表列出了最常用的标准库包：

包名功能典型用途 net/httpHTTP 客户端与服务端构建 Web 服务、RESTful API encoding/jsonJSON 编解码API 数据序列化与反序列化 database/sql数据库访问接口配合驱动操作 MySQL、PostgreSQL 等 os操作系统交互文件读写、环境变量、进程管理 fmt格式化输入输出打印日志、格式化字符串 sync同步原语Mutex、WaitGroup、Once 等并发控制 testing测试框架单元测试、基准测试、示例测试 time时间处理计时器、定时器、时间格式化 context上下文传递超时控制、请求取消、值传递 ioI/O 抽象统一的读写接口 crypto加密算法哈希、对称/非对称加密、TLS flag命令行参数解析CLI 工具开发

#### 跨平台编译

Go 编译器支持交叉编译，在任意平台上编译出其他目标平台的可执行文件。只需设置 GOOS 和 GOARCH 两个环境变量即可。

```

# 在 macOS 上编译 Linux 可执行文件
$ GOOS=linux GOARCH=amd64 go build -o app-linux main.go

# 在 Linux 上编译 Windows 可执行文件
$ GOOS=windows GOARCH=amd64 go build -o app.exe main.go

```

编译产物是独立的二进制文件，目标机器无需安装 Go 运行时或任何依赖，直接运行即可。

### Go 语言的应用领域

Go 凭借高性能、易部署和原生并发的优势，在以下领域得到了广泛采用。

#### 云原生与基础设施

Go 是现代云原生技术栈的事实标准语言。几乎所有知名的云原生项目均使用 Go 编写。

项目说明所属领域 Docker容器化平台，改变了软件交付方式容器 Kubernetes (K8s)容器编排系统，云原生的事实标准容器编排 etcd分布式键值存储，K8s 的核心组件分布式存储 Prometheus监控与告警系统可观测性 Grafana数据可视化平台（后端）可观测性 Terraform基础设施即代码（IaC）工具DevOps Consul服务发现与配置管理服务网格 Traefik云原生反向代理与负载均衡器网络

#### 微服务与 API 开发

Go 凭借极低的内存占用和出色的并发处理能力，成为构建微服务的理想选择。一个简单的 HTTP 服务仅需几 MB 内存即可运行。

常用的 Go Web 框架包括 Gin、Echo、Fiber 等，它们在性能基准测试中通常大幅领先其他语言的同类框架。

#### CLI 工具开发

Go 编译为单个二进制文件的特性，使其成为 CLI 工具开发的首选。用户无需安装任何运行时，下载一个文件即可使用。

知名项目包括：Hugo（静态网站生成器）、fzf（模糊搜索工具）、gh（GitHub CLI）和 kubectl（Kubernetes 命令行工具）。

#### 网络编程与分布式系统

Go 标准库中的 net 包提供了完善的网络编程支持，goroutine 模型天然适配高并发网络服务（如代理服务器、消息队列、游戏服务器）。gRPC 框架的 Go 实现也是分布式系统中服务间通信的主流方案。

#### 区块链

Go 在区块链领域也有重要应用。go-ethereum（Geth） 是以太坊最主流的客户端实现，Hyperledger Fabric 的核心组件也使用 Go 编写。

#### AI 与机器学习（新兴领域）

虽然 Python 仍是 AI/ML 领域的主导语言，但 Go 正逐步进入这一领域。Ollama（本地大模型运行工具）和 LangChainGo 等项目的兴起，展示了 Go 在 AI 基础设施和 LLM 应用开发中的潜力。

### 快速开始

从零搭建 Go 开发环境只需三步：安装、配置、写代码。

#### 安装 Go

访问 go.dev/dl 下载对应平台的安装包。macOS 用户也可通过 Homebrew 安装。

```

# macOS (Homebrew)
$ brew install go

# 验证安装
$ go version
go version go1.24.0 darwin/arm64

```

#### 第一个 Go 程序

创建项目目录并编写经典的 Hello World 程序。

### 实例

// 文件路径：main.go
// 每个 Go 程序必须属于一个包，main 包是可执行程序的入口

package main

// 导入格式化输出的标准库包
import "fmt"

// main 函数：程序的入口点，无参数，无返回值
func main() {
// Println 打印一行内容并在末尾自动添加换行
fmt.Println("Hello, RUNOOB!")
fmt.Println("欢迎来到 Go 语言的世界")
}

运行方式有两种：直接执行（不产生二进制文件）或编译后执行。

```

# 方式一：直接运行（开发调试时常用）
$ go run main.go
Hello, RUNOOB!
欢迎来到 Go 语言的世界

# 方式二：编译为二进制文件后执行（生产部署时常用）
$ go build -o hello main.go
$ ./hello
Hello, RUNOOB!
欢迎来到 Go 语言的世界

```

#### Go Modules 初始化

对于正式项目，建议使用 Go Modules 管理依赖。模块名通常使用仓库路径（如 GitHub 地址），以便他人引用。

```

# 初始化 Go Module
$ mkdir my-go-project && cd my-go-project
$ go mod init github.com/runoob/my-go-project
go: creating new go.mod: module github.com/runoob/my-go-project

# go.mod 文件内容
module github.com/runoob/my-go-project

go 1.24

```

### 基础语法概览

Go 的语法精简且一致。以下快速浏览最常用的语法元素，每个示例都是可直接运行的完整程序。

#### 变量与常量

Go 提供多种变量声明方式，适应不同场景。常量使用 const 关键字声明。

### 实例

package main

import "fmt"

func main() {
// 方式一：标准声明（显式类型）
var age int = 25

// 方式二：声明后赋值（类型在声明时指定）
var score float64
score = 98.5

// 方式三：短变量声明（最常用，仅限函数内部）
name := "runoob"

// 方式四：多变量同时声明
x, y := 10, 20

// 常量声明（编译时确定，不可修改）
const Pi = 3.14159
const AppName = "RUNOOB"

fmt.Println(name, age, score, x, y, Pi, AppName)
}

#### 控制流程

Go 的控制流程精简为 if、for 和 switch 三种。Go 没有 while 关键字，所有循环均使用 for 实现。

### 实例

package main

import "fmt"

func main() {
// if 语句：条件表达式不需要括号，但大括号必须有
score := 85
if score >= 90 {
fmt.Println("RUNOOB 评级: A")
} else if score >= 60 {
fmt.Println("RUNOOB 评级: B")
} else {
fmt.Println("RUNOOB 评级: C")
}

// for 循环一：标准三要素形式（类似 C 的 for）
fmt.Print("计数: ")
for i := 0; i < 5; i++ {
fmt.Print(i, " ")
}
fmt.Println()

// for 循环二：仅条件（类似 while）
sum := 0
n := 1
for n <= 10 {
sum += n
n++
}
fmt.Println("1 到 10 之和:", sum)

// for 循环三：range 遍历（最常用）
fruits := []string{"apple", "banana", "cherry"}
for index, fruit := range fruits {
fmt.Printf("索引 %d: %s\n", index, fruit)
}

// switch 语句：每个 case 自动 break，无需手动添加
day := "周一"
switch day {
case "周一":
fmt.Println("新的一周开始了")
case "周五":
fmt.Println("周末快到了")
default:
fmt.Println("普通的一天")
}
}

#### 函数

Go 函数使用 func 关键字声明，支持多返回值——这是 Go 错误处理机制的基础。

### 实例

package main

import (
"errors"
"fmt"
)

// add 函数：两个参数同为 int 类型，返回一个 int
func add(a, b int) int {
return a + b
}

// divide 函数：返回两个值——结果和错误
// Go 惯用 (result, error) 模式进行错误处理
func divide(a, b float64) (float64, error) {
if b == 0 {
// 返回零值和自定义错误
return 0, errors.New("除数不能为零")
}
return a / b, nil // nil 表示没有错误
}

// swap 函数：多返回值用于返回多个计算结果
func swap(x, y string) (string, string) {
return y, x
}

func main() {
// 调用普通函数
sum := add(10, 20)
fmt.Println("10 + 20 =", sum)

// 调用多返回值函数并检查错误
result, err := divide(10, 2)
if err != nil {
fmt.Println("错误:", err)
} else {
fmt.Println("10 / 2 =", result)
}

// 测试除零情况
_, err = divide(10, 0) // _ 用于忽略不需要的返回值
if err != nil {
fmt.Println("除零错误:", err) // 预期输出此错误
}

// 多返回值交换
first, second := swap("RUNOOB", "Hello")
fmt.Println("交换后:", first, second)
}

Go 没有 try-catch 异常机制，而是通过函数返回 error 值来显式处理错误。这种设计让错误处理路径一目了然，避免了隐式控制流。

#### 结构体与方法

Go 没有类的概念，使用 struct（结构体）组织数据，通过为类型定义方法来实现面向对象风格的行为。方法是带有接收者（receiver）的函数。

### 实例

package main

import "fmt"

// User 定义用户结构体（类似于其他语言中的类）
type User struct {
Name string // 首字母大写 = 公开字段（可被外部包访问）
Email string
age int // 首字母小写 = 私有字段（仅包内可访问）
}

// Greet 是为 User 类型定义的方法
// (u User) 是值接收者，方法内对 u 的修改不会影响原始值
func (u User) Greet() string {
return fmt.Sprintf("你好，我是 %s，邮箱是 %s", u.Name, u.Email)
}

// Birthday 使用指针接收者，方法内可以修改原始值
func (u *User) Birthday() {
u.age++ // 通过指针直接修改原始结构体的 age 字段
}

func main() {
// 创建结构体实例
user := User{
Name: "runoob",
Email: "runoob@example.com",
age: 25,
}

// 调用方法
fmt.Println(user.Greet())
fmt.Printf("生日前年龄: %d\n", user.age)

user.Birthday() // 使用指针接收者，age 被加 1
fmt.Printf("生日后年龄: %d\n", user.age)
}

#### 接口

Go 的接口是隐式实现的——只要一个类型实现了接口中所有方法，它就自动满足该接口，无需显式声明 implements。这种设计极大地降低了代码耦合度。

### 实例

package main

import (
"fmt"
"math"
)

// Shape 接口：定义几何形状的行为规范
type Shape interface {
Area() float64 // 计算面积
Perimeter() float64 // 计算周长
}

// Circle 圆形结构体
type Circle struct {
Radius float64
}

// Circle 实现 Shape 接口（隐式实现，无需显式声明）
func (c Circle) Area() float64 {
return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
return 2 * math.Pi * c.Radius
}

// Rectangle 矩形结构体
type Rectangle struct {
Width, Height float64
}

// Rectangle 也隐式实现了 Shape 接口
func (r Rectangle) Area() float64 {
return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
return 2 * (r.Width + r.Height)
}

// PrintShape 接受任何实现了 Shape 接口的类型
func PrintShape(s Shape) {
fmt.Printf("面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
c := Circle{Radius: 5} // 半径为 5 的圆
r := Rectangle{Width: 4, Height: 6} // 宽 4 高 6 的矩形

fmt.Println("圆形 (RUNOOB 示例):")
PrintShape(c)

fmt.Println("矩形 (RUNOOB 示例):")
PrintShape(r)
}

Go 的接口设计鼓励「小接口」——通常只包含 1-3 个方法。标准库中的 io.Reader（仅一个 Read 方法）和 io.Writer（仅一个 Write 方法）是最经典的例子。

### 常用示例

以下通过两个接近真实场景的完整示例，展示 Go 在实际开发中的用法。

#### 构建 HTTP 服务器

使用标准库 net/http 即可搭建一个功能完整的 Web 服务，无需任何第三方框架。

### 实例

// 文件路径：server.go
package main

import (
"encoding/json"
"fmt"
"log"
"net/http"
)

// Response 定义 JSON 响应的数据结构
type Response struct {
Code int `json:"code"` // 结构体标签指定 JSON 字段名
Message string `json:"message"`
Data any `json:"data,omitempty"` // omitempty：值为空时不输出该字段
}

// helloHandler 处理 /hello 路由的请求
func helloHandler(w http.ResponseWriter, r *http.Request) {
// 仅允许 GET 请求
if r.Method != http.MethodGet {
http.Error(w, "仅支持 GET 请求", http.StatusMethodNotAllowed)
return
}

// 从查询参数获取 name，默认为 "RUNOOB"
name := r.URL.Query().Get("name")
if name == "" {
name = "RUNOOB"
}

// 构建响应数据
resp := Response{
Code: 200,
Message: fmt.Sprintf("你好，%s！欢迎访问 Go HTTP 服务", name),
}

// 设置响应头并返回 JSON
w.Header().Set("Content-Type", "application/json; charset=utf-8")
json.NewEncoder(w).Encode(resp)
}

func main() {
// 注册路由处理函数
http.HandleFunc("/hello", helloHandler)

addr := ":8080"
fmt.Printf("RUNOOB 服务器启动，监听地址: http://localhost%s\n", addr)
fmt.Printf("访问示例: http://localhost%s/hello?name=Go开发者\n", addr)

// 启动 HTTP 服务（阻塞，直到服务停止）
log.Fatal(http.ListenAndServe(addr, nil))
}

启动服务后，访问对应地址即可获得 JSON 响应：

```

$ go run server.go
RUNOOB 服务器启动，监听地址: http://localhost:8080
访问示例: http://localhost:8080/hello?name=Go开发者

# 使用 curl 测试
$ curl http://localhost:8080/hello?name=Go开发者
{"code":200,"message":"你好，Go开发者！欢迎访问 Go HTTP 服务"}

$ curl http://localhost:8080/hello
{"code":200,"message":"你好，RUNOOB！欢迎访问 Go HTTP 服务"}

```

对于生产级 Web 服务，建议结合实际需求选择合适的框架。小型项目用标准库即可，需要路由分组、中间件等功能时可选用 Gin 或 Echo。

#### 并发 Worker 池

这是 Go 并发的经典应用场景：使用限定数量的 goroutine 并行处理大批量任务。

### 实例

// 文件路径：workerpool.go
package main

import (
"fmt"
"sync"
"time"
)

// worker 从 jobs 通道读取任务，处理后将结果写入 results 通道
// 任务处理完毕时调用 wg.Done() 通知 WaitGroup
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
defer wg.Done() // 无论函数如何退出，都标记该 worker 完成
for job := range jobs {
fmt.Printf("[Worker %d] 处理任务 %d\n", id, job)
time.Sleep(200 * time.Millisecond) // 模拟任务处理耗时
results <- job * job // 计算平方作为结果
}
}

func main() {
const numJobs = 10 // 总任务数
const numWorkers = 3 // 并发 worker 数量

jobs := make(chan int, numJobs) // 带缓冲的任务通道
results := make(chan int, numJobs) // 带缓冲的结果通道
var wg sync.WaitGroup // 用于等待所有 worker 完成

fmt.Println("=== RUNOOB Worker Pool 示例 ===")

// 启动固定数量的 worker goroutine
for w := 1; w <= numWorkers; w++ {
wg.Add(1) // 每启动一个 worker，WaitGroup 计数加 1
go worker(w, jobs, results, &wg)
}

// 向通道发送所有任务
for j := 1; j <= numJobs; j++ {
jobs <- j
}
close(jobs) // 关闭 jobs 通道，worker 的 range 循环会自行退出

// 等待所有 worker 完成后，关闭结果通道
go func() {
wg.Wait()
close(results)
}()

// 收集并打印所有结果
fmt.Println("结果汇总:")
for result := range results {
fmt.Printf("%d ", result)
}
fmt.Println("\n所有任务处理完成")
}

```

$ go run workerpool.go
=== RUNOOB Worker Pool 示例 ===
[Worker 3] 处理任务 1
[Worker 1] 处理任务 3
[Worker 2] 处理任务 2
[Worker 1] 处理任务 4
[Worker 3] 处理任务 5
[Worker 2] 处理任务 6
...
结果汇总:
1 4 9 16 25 36 49 64 81 100
所有任务处理完成

```

### 常用工具与命令

Go 自带一套完整的开发工具链，无需安装额外的构建工具。以下是最常用的 go 命令速查：

命令功能常用场景 go build编译源代码生成可执行文件，用于部署 go run编译并运行快速测试代码，开发调试 go test执行测试运行 _test.go 文件中的测试函数 go mod init初始化模块创建新项目时生成 go.mod 文件 go mod tidy整理依赖下载缺失的依赖，移除未使用的依赖 go fmt格式化代码统一代码风格（Go 社区强制使用） go vet静态分析检查代码中的可疑构造 go get安装依赖包下载并安装指定的包 go doc查看文档在终端查看包或函数的文档 go env查看环境变量查看 GOPATH、GOROOT 等配置 go clean清理构建缓存释放磁盘空间

go fmt 是 Go 社区的一项强制性约定——所有 Go 代码都应通过 go fmt 格式化。这消除了关于代码风格的争论，让代码库保持高度一致性。大多数编辑器（VS Code、GoLand 等）都支持保存时自动运行 go fmt。

### 注意事项与常见问题

以下是 Go 初学者最容易遇到的几个问题和最佳实践建议。

#### GOPATH 与 Go Modules

Go 1.16 起，Go Modules 已成为默认模式。新项目不再需要放在 GOPATH 目录下，可以在任意位置创建。

如果你看到教程中提到 GOPATH 的复杂配置，通常可以忽略——使用 Go Modules 就能满足绝大多数开发需求。

#### 错误处理不可忽略

Go 编译器不会强制你处理 error，但忽略错误是一种不良实践。至少应在开发阶段记录或打印错误信息。

### 实例

package main

import (
"fmt"
"os"
)

func main() {
// 不推荐：忽略错误
data, _ := os.ReadFile("config.txt")
fmt.Println(string(data))

// 推荐：显式处理错误
data, err := os.ReadFile("config.txt")
if err != nil {
fmt.Printf("读取文件失败: %v\n", err)
return
}
fmt.Println(string(data))
}

#### Goroutine 泄漏

启动 goroutine 后，如果它永远阻塞且无法退出，就会造成 goroutine 泄漏。长期运行的服务中，goroutine 泄漏会逐渐耗尽内存。

最佳实践：

- 确保每个 goroutine 都有退出路径
- 使用 context.Context 传递取消信号
- 使用 sync.WaitGroup 追踪 goroutine 生命周期

#### 切片（slice）的底层数组共享

多个切片可能共享同一个底层数组。当你对一个切片执行 append 或修改操作时，需要注意是否会影响其他引用同一底层数组的切片。

当不确定时，使用 copy() 创建独立副本。

#### nil 接口与 nil 指针的区别

这是一个经典的 Go 陷阱：一个接口值包含一个 nil 指针时，接口本身不是 nil。这意味着 if err != nil 可能返回 true，但实际指针为 nil。

在返回 error 时，应直接 return nil 而非 return 一个包含 nil 指针的 error 接口变量。

Go 的类型系统简洁但有自己的「性格」。从其他语言转过来的开发者最好花时间阅读 Effective Go 官方文档（golang.org/doc/effective_go），了解 Go 惯用的编码方式和设计模式。

#### 代码组织

Go 使用包（package）组织代码，而不是文件夹目录结构。

同一个目录下的所有 .go 文件必须属于同一个包，且包名通常与目录名一致（main 包和 _test 测试文件除外）。Go 没有「子包」的概念——不同目录的包即使路径相近，也是完全独立的。

### 总结

Go 语言以极简的语法、卓越的性能和原生的并发支持，在云原生、微服务、DevOps 等领域确立了不可替代的地位。

它不是一门追求特性的语言，而是一门追求工程效率的语言——让团队写更少的代码，得到更稳定的系统。

Go 语言的优势和适用场景总结如下：

维度Go 的表现 学习曲线语法精简，有编程基础的开发者可在一周内上手 编译速度大型项目秒级编译，极快的开发反馈循环 运行时性能接近 C/C++，远超 Python/Ruby/JavaScript 内存占用一个基础 HTTP 服务仅需 5-10 MB 内存 并发处理goroutine 极其轻量，数万并发不是问题 部署难度编译为独立二进制文件，复制即可运行 最适合云原生服务、微服务、CLI 工具、分布式系统、网络编程 不太适合GUI 桌面应用、移动端开发、操作系统内核、嵌入式底层

如果你正在构建一个面向并发的后端服务、一个 CLI 工具或一个基础设施项目，Go 是非常值得认真考虑的选择。

---

## Go 语言环境安装

Source: https://www.runoob.com/go/go-environment.html

## Go 语言环境安装

Go 语言支持以下系统：

- Linux
- FreeBSD
- Mac OS X（也称为 Darwin）
- Windows

安装包下载地址为：https://go.dev/dl/。

如果打不开可以使用这个地址：https://golang.google.cn/dl/。

各个系统对应的包名：

安装包文件名 适用系统 处理器架构 系统版本要求 说明 `go1.26.3.windows-amd64.msi` Microsoft Windows x86-64 (Intel 64-bit) Windows 10 或更新 Windows 64 位系统图形化安装程序 `go1.26.3.darwin-arm64.pkg` Apple macOS ARM64 (Apple Silicon 64-bit) macOS 12 或更新 适用于 M 系列芯片 Mac（包括你的 Mac mini 4）的安装包 `go1.26.3.darwin-amd64.pkg` Apple macOS x86-64 (Intel 64-bit) macOS 12 或更新 适用于搭载 Intel 处理器的老款 Mac `go1.26.3.linux-amd64.tar.gz` Linux x86-64 (Intel 64-bit) Linux 3.2 或更新 绝大多数 64 位 Linux 系统的压缩包 `go1.26.3.src.tar.gz` 任意系统 源码 - Go 语言源代码包，需自行编译

### UNIX/Linux/Mac OS X, 和 FreeBSD 安装

以下介绍了在UNIX/Linux/Mac OS X, 和 FreeBSD系统下使用源码安装方法：

1、下载二进制包：go1.4.linux-amd64.tar.gz。

2、将下载的二进制包解压至 /usr/local目录。

```
tar -C /usr/local -xzf go1.4.linux-amd64.tar.gz
```

3、将 /usr/local/go/bin 目录添加至 PATH 环境变量：

```

export PATH=$PATH:/usr/local/go/bin

```

以上只能暂时添加 PATH，关闭终端下次再登录就没有了。

我们可以编辑 ~/.bash_profile 或者 /etc/profile，并将以下命令添加该文件的末尾，这样就永久生效了：

```

export PATH=$PATH:/usr/local/go/bin

```

添加后需要执行：

```
source ~/.bash_profile
或
source /etc/profile
```

注意：MAC 系统下你可以使用 .pkg 结尾的安装包直接双击来完成安装，安装目录在 /usr/local/go/ 下。

### Windows 系统下安装

Windows 下可以使用 .msi 后缀(在下载列表中可以找到该文件，如go1.4.2.windows-amd64.msi)的安装包来安装。

默认情况下 .msi 文件会安装在 c:\Go 目录下。你可以将 c:\Go\bin 目录添加到 Path 环境变量中。添加后你需要重启命令窗口才能生效。

#### 安装测试

创建工作目录 C:\>Go_WorkSpace。

### test.go 文件代码：

package main

import "fmt"

func main() {
fmt.Println("Hello, World!")
}

使用 go 命令执行以上代码输出结果如下：

```

C:\Go_WorkSpace>go run test.go

Hello, World!

```

---

## 使用 VSCode 开发 Go

Source: https://www.runoob.com/go/go-vscode.html

## 使用 VSCode 开发 Go

Visual Studio Code（简称 VS Code）是一个由微软开发的免费、开源的代码编辑器，支持多种编程语言，并提供了代码高亮、智能代码补全、代码重构、调试等功能。

借助 Visual Studio Code 的 Go 官方扩展插件，你可以获得智能提示、代码跳转、符号检索、单元测试、程序调试等一系列助力 Go 项目开发的实用能力。

如果你还不了解 VS Code 或者还未安装，可以参考：VSCode 教程。

#### 安装 Go 插件

打开 VS Code 扩展市场，搜索：Go Extension for VS Code。

编辑器内容，点击左侧的扩展，然后搜索 Go，然后安装官方插件：

安装完成后，我们就可以开始编写 Go 的代码，可以输入 fmt. 试试智能提示功能：

### 智能提示(IntelliSense)

智能提示由 Go 官方团队维护的 gopls 语言服务提供，可通过 gopls 配置项自定义运行规则。

#### 1. 语义高亮

VS Code默认基于TextMate语法做代码高亮，如需效果更好的语义级高亮，在配置中开启`ui.semanticTokens`：

```
"gopls": { "ui.semanticTokens": true }
```

#### 2. 代码自动补全

编辑Go源码时，插件实时弹出代码补全候选，支持当前包、已导入包、尚未导入包的成员联想；输入包名+.即可唤起对应包内方法/常量提示。

快捷键：`Ctrl+空格(⌃Space)` 手动唤起补全弹窗

#### 3. 悬浮文档提示

鼠标悬浮在变量、函数、结构体上，弹窗展示注释文档、函数签名等详情。

#### 4. 函数参数提示

键入函数左括号`(`时自动弹出形参签名，输入实参时下划线跟随切换当前填写参数；

快捷键：`Shift+Cmd+空格(⇧⌘Space)` 光标在括号内时手动触发参数提示

### 代码导航

编辑器右键菜单内置全套代码导航功能：

- F12 跳转定义：进入类型/变量源码定义处
- 跳转类型定义：定位变量底层的类型源码
- Alt+F12(⌥F12) 预览定义：弹窗预览源码，不切换文件
- Shift+F12(⇧F12) 查找引用：列出当前符号所有被引用位置
- Shift+Alt+H(⇧⌥H) 调用层级：查看函数的调用方与被调用方链路
- Cmd+F12(⌘F12) 查看实现：接口查看所有实现类、实体类查看实现的全部接口

打开命令面板`Shift+Cmd+P(⇧⌘P)`，通过符号检索快速跳转：

- `Shift+Cmd+O(⇧⌘O)`：当前文件内检索符号
- `Cmd+T(⌘T)`：全工作区全局检索符号

使用指令 `Go: Toggle Test File`，快速在业务代码与对应测试文件来回切换。

### 代码格式化

- 快捷键：`Shift+Alt+F(⇧⌥F)` 一键格式化当前文件
- 也可通过命令面板/右键菜单执行「格式化文档」

#### 格式化配置

- 关闭保存自动格式化

```
"[go]": {
"editor.formatOnSave": false
}
```

- 指定Go插件为Go文件默认格式化工具

```
"[go]": {
"editor.defaultFormatter": "golang.go"
}
```

- 启用gofumpt严格格式化风格（比原生gofmt规范更严苛）

```
"gopls": {
"formatting.gofumpt": true
}
```

格式化底层依赖gopls实现

### 单元测试

VS Code侧边栏测试面板 + 代码上方CodeLens快捷按钮，支持函数/文件/包/整个项目的单元测试、基准测试、性能采样。
命令面板输入 `Go: test` 可查看全部测试指令，常用：

- Go: Test Function At Cursor：运行光标所在函数测试
- Go: Test File：运行当前文件全部测试用例
- Go: Test Package：运行当前代码包全部测试
- Go: Test All Packages in Workspace：全项目批量测试

前3条指令可基于gotests自动生成测试用例骨架；插件支持配置：

- `go.testOnSave`：保存代码自动执行测试
- `go.coverOnSave`：保存自动生成测试覆盖率
- `go.testFlags`：自定义go test运行参数

### 调试Debug

Go插件依托Delve调试器实现Go代码调试，独有能力：

- 本地调试 + 远程服务调试
- 使用Delve表达式语法在调试面板查看自定义数据
- 在调试控制台执行dlv原生指令，动态修改调试配置、查看变量
- 配置`hideSystemGoroutines`隐藏/展示系统协程
- 右键代码打开汇编视图Disassembly View
- 实验性功能：函数调用调试、core文件分析、Mozilla rr反向调试

### 更多相关工具

---

## Go 语言结构

Source: https://www.runoob.com/go/go-program-structure.html

## Go 语言结构

在我们开始学习 Go 编程语言的基础构建模块前，让我们先来了解 Go 语言最简单程序的结构。

### Go Hello World 实例

Go 语言的基础组成有以下几个部分：

- 包声明
- 引入包
- 函数
- 变量
- 语句 & 表达式
- 注释

接下来让我们来看下简单的代码，该代码输出了"Hello World!":

### 实例

package main

import "fmt"

func main() {
/* 这是我的第一个简单的程序 */
fmt.Println("Hello, World!")
}

让我们来看下以上程序的各个部分：

- 第一行代码 package main 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
- 下一行 import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
- 下一行 func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
- 下一行 /*...*/ 是注释，在程序执行时将被忽略。单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。多行注释也叫块注释，均已以 /* 开头，并以 */ 结尾，且不可以嵌套使用，多行注释一般用于包的文档描述或注释成块的代码片段。
- 下一行 fmt.Println(...) 可以将字符串输出到控制台，并在最后自动增加换行字符 \n。
使用 fmt.Print("hello, world\n") 可以得到相同的结果。
Print 和 Println 这两个函数也支持使用变量，如：fmt.Println(arr)。如果没有特别指定，它们会以默认的打印格式将变量 arr 输出到控制台。
- 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。

### 执行 Go 程序

让我们来看下如何编写 Go 代码并执行它。步骤如下：

- 打开编辑器如Sublime2，将以上代码添加到编辑器中。
- 将以上代码保存为 hello.go
- 打开命令行，并进入程序文件保存的目录中。
- 输入命令 go run hello.go 并按回车执行代码。
- 如果操作正确你将在屏幕上看到 "Hello World!" 字样的输出。

```

$ go run hello.go
Hello, World!

```

- 我们还可以使用 go build 命令来生成二进制文件：

```
$ go build hello.go
$ ls
hello hello.go
$ ./hello
Hello, World!
```

#### 注意

需要注意的是 { 不能单独放在一行，所以以下代码在运行时会产生错误：

### 实例

package main

import "fmt"

func main()
{ // 错误，{ 不能在单独的行上
fmt.Println("Hello, World!")
}

---

## Go 语言基础语法

Source: https://www.runoob.com/go/go-basic-syntax.html

## Go 语言基础语法

上一章节我们已经了解了 Go 语言的基本组成结构，本章节我们将学习 Go 语言的基础语法。

### Go 标记

Go 程序可以由多个标记组成，可以是关键字，标识符，常量，字符串，符号。如以下 GO 语句由 6 个标记组成：

```

fmt.Println("Hello, World!")

```

6 个标记是(每行一个)：

```

1. fmt
2. .
3. Println
4. (
5. "Hello, World!"
6. )

```

### 行分隔符

在 Go 程序中，一行代表一个语句结束。每个语句不需要像 C 家族中的其它语言一样以分号 ; 结尾，因为这些工作都将由 Go 编译器自动完成。

如果你打算将多个语句写在同一行，它们则必须使用 ; 人为区分，但在实际开发中我们并不鼓励这种做法。

以下为两个语句：

fmt.Println("Hello, World!")
fmt.Println("菜鸟教程：runoob.com")

### 注释

注释不会被编译，每一个包应该有相关注释。

单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。多行注释也叫块注释，均已以 /* 开头，并以 */ 结尾。如：

```

// 单行注释
/*
Author by 菜鸟教程
我是多行注释
*/

```

### 标识符

标识符用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字。

以下是有效的标识符：

```

mahesh kumar abc move_name a_123
myname50 _temp j a23b9 retVal

```

以下是无效的标识符：

- 1ab（以数字开头）
- case（Go 语言的关键字）
- a+b（运算符是不允许的）

### 字符串连接

Go 语言的字符串连接可以通过 + 实现：

### 实例

package main
import "fmt"
func main() {
fmt.Println("Google" + "Runoob")
}

以上实例输出结果为：

```
GoogleRunoob
```

### 关键字

下面列举了 Go 代码中会使用到的 25 个关键字或保留字：

breakdefaultfuncinterfaceselect casedefergomapstruct chanelsegotopackageswitch constfallthroughifrangetype continueforimportreturnvar

除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：

append bool byte cap close complex complex64 complex128 uint16 copy false float32 float64 imag int int8 int16 uint32 int32 int64 iota len make new nil panic uint64 print println real recover string true uint uint8 uintptr

程序一般由关键字、常量、变量、运算符、类型和函数组成。

程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。

程序中可能会使用到这些标点符号：.、,、;、: 和 …。

### Go 语言的空格

在 Go 语言中，空格通常用于分隔标识符、关键字、运算符和表达式，以提高代码的可读性。

Go 语言中变量的声明必须使用空格隔开，如：

```
var x int
const Pi float64 = 3.14159265358979323846
```

在运算符和操作数之间要使用空格能让程序更易阅读：

无空格：

```
fruit=apples+oranges;
```

在变量与运算符间加入空格，程序看起来更加美观，如：

```

fruit = apples + oranges;

```

在关键字和表达式之间要使用空格。

例如：

```
if x > 0 {
// do something
}
```

在函数调用时，函数名和左边等号之间要使用空格，参数之间也要使用空格。

例如：

```
result := add(2, 3)
```

### 格式化字符串

Go 语言中使用 fmt.Sprintf 或 fmt.Printf 格式化字符串并赋值给新串：

- Sprintf 根据格式化参数生成格式化的字符串并返回该字符串。
- Printf 根据格式化参数生成格式化的字符串并写入标准输出。

### Sprintf 实例

package main

import (
"fmt"
)

func main() {
// %d 表示整型数字，%s 表示字符串
var stockcode=123
var enddate="2020-12-31"
var url="Code=%d&endDate=%s"
var target_url=fmt.Sprintf(url,stockcode,enddate)
fmt.Println(target_url)
}

输出结果为：

```
Code=123&endDate=2020-12-31
```

### Printf 实例

package main

import (
"fmt"
)

func main() {
// %d 表示整型数字，%s 表示字符串
var stockcode=123
var enddate="2020-12-31"
var url="Code=%d&endDate=%s"
fmt.Printf(url,stockcode,enddate)
}

输出结果为：

```
Code=123&endDate=2020-12-31
```

更多内容参见：

- Go fmt.Sprintf 格式化字符串
- Go fmt.Printf 格式化字符串

---

## Go 语言数据类型

Source: https://www.runoob.com/go/go-data-types.html

## Go 语言数据类型

在 Go 编程语言中，数据类型用于声明函数和变量。

数据类型的出现是为了把数据分成所需内存大小不同的数据，编程的时候需要用大数据的时候才需要申请大内存，就可以充分利用内存。

Go 语言按类别有以下几种数据类型：

序号类型和描述 1布尔型
布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。 2数字类型
整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。 3字符串类型:
字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。 4派生类型:
包括：

- (a) 指针类型（Pointer）
- (b) 数组类型
- (c) 结构化类型(struct)
- (d) Channel 类型
- (e) 函数类型
- (f) 切片类型
- (g) 接口类型（interface）
- (h) Map 类型

### 数字类型

Go 也有基于架构的类型，例如：int、uint 和 uintptr。

序号类型和描述 1uint8
无符号 8 位整型 (0 到 255) 2uint16
无符号 16 位整型 (0 到 65535) 3uint32
无符号 32 位整型 (0 到 4294967295) 4uint64
无符号 64 位整型 (0 到 18446744073709551615) 5int8
有符号 8 位整型 (-128 到 127) 6int16
有符号 16 位整型 (-32768 到 32767) 7int32
有符号 32 位整型 (-2147483648 到 2147483647) 8int64
有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)

#### 浮点型

序号类型和描述 1float32
IEEE-754 32位浮点型数 2float64
IEEE-754 64位浮点型数 3complex64
32 位实数和虚数 4complex128
64 位实数和虚数

### 其他数字类型

以下列出了其他更多的数字类型：

序号类型和描述 1byte
类似 uint8 2rune
类似 int32 3uint
32 或 64 位 4int
与 uint 一样大小 5uintptr
无符号整型，用于存放一个指针

---

## Go 语言变量

Source: https://www.runoob.com/go/go-variables.html

## Go 语言变量

变量来源于数学，是计算机语言中能储存计算结果或能表示值抽象概念。

变量可以通过变量名访问。

Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。

声明变量的一般形式是使用 var 关键字：

```
var identifier type
```

可以一次声明多个变量：

```
var identifier1, identifier2 type
```

### 实例

package main
import "fmt"
func main() {
var a string = "Runoob"
fmt.Println(a)

var b, c int = 1, 2
fmt.Println(b, c)
}

以上实例输出结果为：

```
Runoob
1 2
```

#### 变量声明

第一种，指定变量类型，如果没有初始化，则变量默认为零值。

```

var v_name v_type
v_name = value

```

零值就是变量没有做初始化时系统默认设置的值。

### 实例

package main
import "fmt"
func main() {

// 声明一个变量并初始化
var a = "RUNOOB"
fmt.Println(a)

// 没有初始化就为零值
var b int
fmt.Println(b)

// bool 零值为 false
var c bool
fmt.Println(c)
}

以上实例执行结果为：

```

RUNOOB
0
false

```

- 数值类型（包括complex64/128）为 0
- 布尔类型为 false
- 字符串为 ""（空字符串）
- 以下几种类型为 nil：

```
var a *int
var a []int
var a map[string] int
var a chan int
var a func(string) int
var a error // error 是接口
```

### 实例

package main

import "fmt"

func main() {
var i int
var f float64
var b bool
var s string
fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

输出结果是：

```
0 0 false ""
```

第二种，根据值自行判定变量类型。

```

var v_name = value

```

### 实例

package main
import "fmt"
func main() {
var d = true
fmt.Println(d)
}

输出结果是：

```
true
```

第三种，如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误，格式：

```
v_name := value
```

例如：

```

var intVal int
intVal :=1 // 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明

```

直接使用下面的语句即可：

```
intVal := 1 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
```

intVal := 1 相等于：

```
var intVal int
intVal =1
```

可以将 var f string = "Runoob" 简写为 f := "Runoob"：

### 实例

package main
import "fmt"
func main() {
f := "Runoob" // var f string = "Runoob"

fmt.Println(f)
}

输出结果是：

```
Runoob
```

#### 多变量声明

```

//类型相同多个变量, 非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3

var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断

vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误


// 这种因式分解关键字的写法一般用于声明全局变量
var (
vname1 v_type1
vname2 v_type2
)

```

### 实例

package main
import "fmt"

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
a int
b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main(){
g, h := 123, "hello"
fmt.Println(x, y, a, b, c, d, e, f, g, h)
}

以上实例执行结果为：

```

0 0 0 false 1 2 123 hello 123 hello

```

### 值类型和引用类型

所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值：

当使用等号 `=` 将一个变量的值赋值给另一个变量时，如：`j = i`，实际上是在内存中将 i 的值进行了拷贝：

你可以通过 &i 来获取变量 i 的内存地址，例如：0xf840000040（每次的地址都可能不一样）。

值类型变量通常存储在栈中，尤其是当它们是局部变量时。当值类型变量的值需要在函数作用域之外使用时，Go 会将其分配到堆内存中。

内存地址会根据机器的不同而有所不同，甚至相同的程序在不同的机器上执行后也会有不同的内存地址。因为每台机器可能有不同的存储器布局，并且位置分配也可能不同。

更复杂的数据通常会需要使用多个字，这些数据一般使用引用类型保存。

一个引用类型的变量 r1 存储的是 r1 的值所在的内存地址（数字），或内存地址中第一个字所在的位置。

这个内存地址称之为指针，这个指针实际上也被存在另外的某一个值中。

同一个引用类型的指针指向的多个字可以是在连续的内存地址中（内存布局是连续的），这也是计算效率最高的一种存储形式；也可以将这些字分散存放在内存中，每个字都指示了下一个字所在的内存地址。

当使用赋值语句 r2 = r1 时，只有引用（地址）被复制。

如果 r1 的值被改变了，那么这个值的所有引用都会指向被修改后的内容，在这个例子中，r2 也会受到影响。

### 简短形式，使用 := 赋值操作符

我们知道可以在变量的初始化时省略变量的类型而由系统自动推断，声明语句写上 var 关键字其实是显得有些多余了，因此我们可以将它们简写为 a := 50 或 b := false。

a 和 b 的类型（int 和 bool）将由编译器自动推断。

这是使用变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。

#### 注意事项

如果在相同的代码块中，我们不可以再次对于相同名称的变量使用初始化声明，例如：a := 20 就是不被允许的，编译器会提示错误 no new variables on left side of :=，但是 a = 20 是可以的，因为这是给相同的变量赋予一个新的值。

如果你在定义变量 a 之前使用它，则会得到编译错误 undefined: a。

如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误，例如下面这个例子当中的变量 a：

### 实例

package main

import "fmt"

func main() {
var a string = "abc"
fmt.Println("hello, world")
}

尝试编译这段代码将得到错误 a declared but not used。

此外，单纯地给 a 赋值也是不够的，这个值必须被使用，所以使用

```
fmt.Println("hello, world", a)
```

会移除错误。

但是全局变量是允许声明但不使用的。 同一类型的多个变量可以声明在同一行，如：

```

var a, b, c int

```

多变量可以在同一行进行赋值，如：

```

var a, b int
var c string
a, b, c = 5, 7, "abc"

```

上面这行假设了变量 a，b 和 c 都已经被声明，否则的话应该这样使用：

```

a, b, c := 5, 7, "abc"
```

右边的这些值以相同的顺序赋值给左边的变量，所以 a 的值是 5， b 的值是 7，c 的值是 "abc"。

这被称为 并行 或 同时 赋值。

如果你想要交换两个变量的值，则可以简单地使用 a, b = b, a，两个变量的类型必须是相同。

空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。

_ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。

并行赋值也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：val, err = Func1(var1)。

---

## Go 语言常量

Source: https://www.runoob.com/go/go-constants.html

## Go 语言常量

常量是一个简单值的标识符，在程序运行时，不会被修改的量。常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

单词释义： `const` 是 `constant`（常量）的缩写，表示程序中不可变的值。

### 基本语法与参数

常量的定义格式：

```

const identifier [type] = value

```

#### 语法说明

- identifier：常量名称，遵循 Go 标识符命名规则（字母、数字、下划线，且不能以数字开头）。
- type（可选）：类型名，如果省略则由编译器自动推断。
- value：常量的值。

#### 定义方式

- 显式类型定义：`const b string = "abc"`
- 隐式类型定义：`const b = "abc"`
- 批量定义：`const c_name1, c_name2 = value1, value2`

### 实例

#### 示例 1：基础常量定义

### 实例

package main

import "fmt"

func main() {
const LENGTH int = 10
const WIDTH int = 5
var area int
const a, b, c = 1, false, "str" // 多重赋值

area = LENGTH * WIDTH
fmt.Printf("面积为 : %d", area)
println()
println(a, b, c)
}

运行结果预期:

```

面积为 : 50
1 false str

```

代码解析:

- `const LENGTH int = 10` 定义了一个整型常量 LENGTH，值为 10。
- `const WIDTH int = 5` 定义了一个整型常量 WIDTH，值为 5。
- `const a, b, c = 1, false, "str"` 使用多重赋值同时定义多个不同类型的常量。
- `area = LENGTH * WIDTH` 计算面积并赋值给变量 area。

### 常量组与枚举

Go 语言没有枚举（enum）关键字，通常使用常量组来实现枚举效果：

### 实例

package main

import "fmt"

// 定义性别枚举
const (
Unknown = 0
Female = 1
Male = 2
)

func main() {
fmt.Printf("未知=%d, 女性=%d, 男性=%d\n", Unknown, Female, Male)
}

运行结果预期:

```

未知=0, 女性=1, 男性=2

```

代码解析:

- 使用常量组可以定义一组相关的常量值。
- 数字 0、1、2 分别代表未知性别、女性和男性。

### 内置函数在常量中的应用

常量表达式中可以使用 `len()`、`cap()`、`unsafe.Sizeof()` 等内置函数计算表达式的值。

注意：常量表达式中函数必须是内置函数。

### 实例

package main

import (
"fmt"
"unsafe"
)

const (
a = "abc"
b = len(a)
c = unsafe.Sizeof(a)
)

func main() {
fmt.Printf("a=%s, b=%d, c=%d\n", a, b, c)
}

运行结果预期:

```

a=abc, b=3, c=16

```

代码解析:

- `a = "abc"`：字符串常量。
- `b = len(a)`：使用 len() 获取字符串长度，结果为 3。
- `c = unsafe.Sizeof(a)`：使用 unsafe.Sizeof() 获取变量内存占用，结果为 16 字节（字符串结构体在 64 位系统上占 16 字节）。

### iota 常量计数器

`iota` 是 Go 语言的特殊常量，也称为常量计数器。它可以在 const 声明块中自动生成递增的序列值。

重要特性：

- `iota` 在 `const` 关键字出现时将被重置为 0。
- const 中每新增一行常量声明，iota 自动计数加 1。
- 可以理解为 const 语句块中的行索引（从 0 开始）。

#### 示例 1：基础 iota 用法

### 实例

package main

import "fmt"

// iota 从 0 开始，每行加 1
const (
a = iota // 0
b // 1（省略值，默认使用上一行的 iota）
c // 2
)

func main() {
fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
}

运行结果预期:

```

a=0, b=1, c=2

```

代码解析:

- 当常量值省略时，默认使用上一行的 iota 值。
- 因此 b=1, c=2。

#### 示例 2：iota 详细用法

### 实例

package main

import "fmt"

func main() {
const (
a = iota // 0
b // 1
c // 2
d = "ha" // 独立值，iota += 1
e // "ha"，iota += 1
f = 100 // 独立值，iota += 1
g // 100，iota += 1
h = iota // 7，恢复使用 iota 值
i // 8
)
fmt.Printf("a=%d, b=%d, c=%d, d=%s, e=%s, f=%d, g=%d, h=%d, i=%d\n", a, b, c, d, e, f, g, h, i)
}

运行结果预期:

```

a=0, b=1, c=2, d=ha, e=ha, f=100, g=100, h=7, i=8

```

代码解析:

- 当显式赋值时，使用该值，但 iota 仍然递增。
- 当省略赋值时，使用上一行的值。
- 在 `h = iota` 时恢复使用 iota 的当前值（7）。

#### 示例 3：iota 与位移运算

### 实例

package main

import "fmt"

const (
i = 1 << iota // 1 << 0 = 1
j = 3 << iota // 3 << 1 = 6
k // 3 << 2 = 12
l // 3 << 3 = 24
)

func main() {
fmt.Printf("i=%d, j=%d, k=%d, l=%d\n", i, j, k, l)
}

运行结果预期:

```

i=1, j=6, k=12, l=24

```

代码解析:

- `<<` 是左移运算符，`x << n` 等价于 `x * 2^n`。
- `i = 1 << 0 = 1`：1 左移 0 位，结果为 1。
- `j = 3 << 1 = 6`：3 左移 1 位，二进制 11 变为 110，即 6。
- `k = 3 << 2 = 12`：3 左移 2 位，二进制 11 变为 1100，即 12。
- `l = 3 << 3 = 24`：3 左移 3 位，二进制 11 变为 11000，即 24。

#### 补充说明

- `1 << n = 1 * 2^n`
- `3 << n = 3 * 2^n`

---

## Go 语言运算符

Source: https://www.runoob.com/go/go-operators.html

## Go 语言运算符

运算符用于在程序运行时执行数学或逻辑运算。

Go 语言内置的运算符有：

- 算术运算符
- 关系运算符
- 逻辑运算符
- 位运算符
- 赋值运算符
- 其他运算符

接下来让我们来详细看看各个运算符的介绍。

### 算术运算符

下表列出了所有Go语言的算术运算符。假定 A 值为 10，B 值为 20。

运算符描述实例 +相加 A + B 输出结果 30 -相减 A - B 输出结果 -10 *相乘 A * B 输出结果 200 /相除 B / A 输出结果 2 %求余 B % A 输出结果 0 ++自增 A++ 输出结果 11 --自减 A-- 输出结果 9

以下实例演示了各个算术运算符的用法：

### 实例

package main

import "fmt"

func main() {

var a int = 21
var b int = 10
var c int

c = a + b
fmt.Printf("第一行 - c 的值为 %d\n", c )
c = a - b
fmt.Printf("第二行 - c 的值为 %d\n", c )
c = a * b
fmt.Printf("第三行 - c 的值为 %d\n", c )
c = a / b
fmt.Printf("第四行 - c 的值为 %d\n", c )
c = a % b
fmt.Printf("第五行 - c 的值为 %d\n", c )
a++
fmt.Printf("第六行 - a 的值为 %d\n", a )
a=21 // 为了方便测试，a 这里重新赋值为 21
a--
fmt.Printf("第七行 - a 的值为 %d\n", a )
}

以上实例运行结果：

```

第一行 - c 的值为 31
第二行 - c 的值为 11
第三行 - c 的值为 210
第四行 - c 的值为 2
第五行 - c 的值为 1
第六行 - a 的值为 22
第七行 - a 的值为 20

```

### 关系运算符

下表列出了所有Go语言的关系运算符。假定 A 值为 10，B 值为 20。

运算符描述实例 == 检查两个值是否相等，如果相等返回 True 否则返回 False。 (A == B) 为 False != 检查两个值是否不相等，如果不相等返回 True 否则返回 False。 (A != B) 为 True >检查左边值是否大于右边值，如果是返回 True 否则返回 False。 (A > B) 为 False <检查左边值是否小于右边值，如果是返回 True 否则返回 False。 (A < B) 为 True >=检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。 (A >= B) 为 False <= 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。 (A <= B) 为 True

以下实例演示了关系运算符的用法：

### 实例

package main

import "fmt"

func main() {
var a int = 21
var b int = 10

if( a == b ) {
fmt.Printf("第一行 - a 等于 b\n" )
} else {
fmt.Printf("第一行 - a 不等于 b\n" )
}
if ( a < b ) {
fmt.Printf("第二行 - a 小于 b\n" )
} else {
fmt.Printf("第二行 - a 不小于 b\n" )
}

if ( a > b ) {
fmt.Printf("第三行 - a 大于 b\n" )
} else {
fmt.Printf("第三行 - a 不大于 b\n" )
}
/* Lets change value of a and b */
a = 5
b = 20
if ( a <= b ) {
fmt.Printf("第四行 - a 小于等于 b\n" )
}
if ( b >= a ) {
fmt.Printf("第五行 - b 大于等于 a\n" )
}
}

以上实例运行结果：

```

第一行 - a 不等于 b
第二行 - a 不小于 b
第三行 - a 大于 b
第四行 - a 小于等于 b
第五行 - b 大于等于 a

```

### 逻辑运算符

下表列出了所有Go语言的逻辑运算符。假定 A 值为 True，B 值为 False。

运算符描述实例 && 逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。 (A && B) 为 False ||逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。 (A || B) 为 True !逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。 !(A && B) 为 True

以下实例演示了逻辑运算符的用法：

### 实例

package main

import "fmt"

func main() {
var a bool = true
var b bool = false
if ( a && b ) {
fmt.Printf("第一行 - 条件为 true\n" )
}
if ( a || b ) {
fmt.Printf("第二行 - 条件为 true\n" )
}
/* 修改 a 和 b 的值 */
a = false
b = true
if ( a && b ) {
fmt.Printf("第三行 - 条件为 true\n" )
} else {
fmt.Printf("第三行 - 条件为 false\n" )
}
if ( !(a && b) ) {
fmt.Printf("第四行 - 条件为 true\n" )
}
}

以上实例运行结果：

```

第二行 - 条件为 true
第三行 - 条件为 false
第四行 - 条件为 true

```

### 位运算符

位运算符对整数在内存中的二进制位进行操作。

下表列出了位运算符 &, |, 和 ^ 的计算：

pqp & qp | qp ^ q 00000 01011 11110 10011

假定 A = 60; B = 13; 其二进制数转换为：

```

A = 0011 1100

B = 0000 1101

-----------------

A&B = 0000 1100

A|B = 0011 1101

A^B = 0011 0001

```

Go 语言支持的位运算符如下表所示。假定 A 为60，B 为13：

运算符描述实例 & 按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。 (A & B) 结果为 12, 二进制为 0000 1100 |按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或 (A | B) 结果为 61, 二进制为 0011 1101 ^ 按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。 (A ^ B) 结果为 49, 二进制为 0011 0001 << 左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。 A << 2 结果为 240 ，二进制为 1111 0000 >> 右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。 A >> 2 结果为 15 ，二进制为 0000 1111

以下实例演示了位运算符的用法：

### 实例

package main

import "fmt"

func main() {

var a uint = 60 /* 60 = 0011 1100 */
var b uint = 13 /* 13 = 0000 1101 */
var c uint = 0

c = a & b /* 12 = 0000 1100 */
fmt.Printf("第一行 - c 的值为 %d\n", c )

c = a | b /* 61 = 0011 1101 */
fmt.Printf("第二行 - c 的值为 %d\n", c )

c = a ^ b /* 49 = 0011 0001 */
fmt.Printf("第三行 - c 的值为 %d\n", c )

c = a << 2 /* 240 = 1111 0000 */
fmt.Printf("第四行 - c 的值为 %d\n", c )

c = a >> 2 /* 15 = 0000 1111 */
fmt.Printf("第五行 - c 的值为 %d\n", c )
}

以上实例运行结果：

```

第一行 - c 的值为 12
第二行 - c 的值为 61
第三行 - c 的值为 49
第四行 - c 的值为 240
第五行 - c 的值为 15

```

### 赋值运算符

下表列出了所有Go语言的赋值运算符。

运算符描述实例 =简单的赋值运算符，将一个表达式的值赋给一个左值 C = A + B 将 A + B 表达式结果赋值给 C +=相加后再赋值 C += A 等于 C = C + A -=相减后再赋值 C -= A 等于 C = C - A *=相乘后再赋值 C *= A 等于 C = C * A /=相除后再赋值 C /= A 等于 C = C / A %=求余后再赋值 C %= A 等于 C = C % A <<=左移后赋值 C <<= 2 等于 C = C << 2 >>=右移后赋值 C >>= 2 等于 C = C >> 2 &=按位与后赋值 C &= 2 等于 C = C & 2 ^=按位异或后赋值 C ^= 2 等于 C = C ^ 2 |=按位或后赋值 C |= 2 等于 C = C | 2

以下实例演示了赋值运算符的用法：

### 实例

package main

import "fmt"

func main() {
var a int = 21
var c int

c = a
fmt.Printf("第 1 行 - = 运算符实例，c 值为 = %d\n", c )

c += a
fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c )

c -= a
fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c )

c *= a
fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c )

c /= a
fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c )

c = 200;

c <<= 2
fmt.Printf("第 6行 - <<= 运算符实例，c 值为 = %d\n", c )

c >>= 2
fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c )

c &= 2
fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c )

c ^= 2
fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c )

c |= 2
fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c )

}

以上实例运行结果：

```

第 1 行 - = 运算符实例，c 值为 = 21
第 2 行 - += 运算符实例，c 值为 = 42
第 3 行 - -= 运算符实例，c 值为 = 21
第 4 行 - *= 运算符实例，c 值为 = 441
第 5 行 - /= 运算符实例，c 值为 = 21
第 6行 - <<= 运算符实例，c 值为 = 800
第 7 行 - >>= 运算符实例，c 值为 = 200
第 8 行 - &= 运算符实例，c 值为 = 0
第 9 行 - ^= 运算符实例，c 值为 = 2
第 10 行 - |= 运算符实例，c 值为 = 2

```

### 其他运算符

下表列出了Go语言的其他运算符。

运算符描述实例 &返回变量存储地址&a; 将给出变量的实际地址。 *指针变量。*a; 是一个指针变量

以下实例演示了其他运算符的用法：

### 实例

package main

import "fmt"

func main() {
var a int = 4
var b int32
var c float32
var ptr *int

/* 运算符实例 */
fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );

/* & 和 * 运算符实例 */
ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
fmt.Printf("a 的值为 %d\n", a);
fmt.Printf("*ptr 为 %d\n", *ptr);
}

以上实例运行结果：

```

第 1 行 - a 变量类型为 = int
第 2 行 - b 变量类型为 = int32
第 3 行 - c 变量类型为 = float32
a 的值为 4
*ptr 为 4

```

### 运算符优先级

有些运算符拥有较高的优先级，二元运算符的运算方向均是从左至右。下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：

优先级运算符 5 * / % << >> & &^ 4 + - | ^ 3 == != < <= > >= 2 && 1 ||

当然，你可以通过使用括号来临时提升某个表达式的整体运算优先级。

以上实例运行结果：

### 实例

package main

import "fmt"

func main() {
var a int = 20
var b int = 10
var c int = 15
var d int = 5
var e int;

e = (a + b) * c / d; // ( 30 * 15 ) / 5
fmt.Printf("(a + b) * c / d 的值为 : %d\n", e );

e = ((a + b) * c) / d; // (30 * 15 ) / 5
fmt.Printf("((a + b) * c) / d 的值为 : %d\n" , e );

e = (a + b) * (c / d); // (30) * (15/5)
fmt.Printf("(a + b) * (c / d) 的值为 : %d\n", e );

e = a + (b * c) / d; // 20 + (150/5)
fmt.Printf("a + (b * c) / d 的值为 : %d\n" , e );
}

以上实例运行结果：

```

(a + b) * c / d 的值为 : 90
((a + b) * c) / d 的值为 : 90
(a + b) * (c / d) 的值为 : 90
a + (b * c) / d 的值为 : 50

```

---

## Go 语言条件语句

Source: https://www.runoob.com/go/go-decision-making.html

## Go 语言条件语句

条件语句需要开发者通过指定一个或多个条件，并通过测试条件是否为 true 来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句。

下图展示了程序语言中条件语句的结构：

Go 语言提供了以下几种条件判断语句：

语句描述 if 语句if 语句 由一个布尔表达式后紧跟一个或多个语句组成。 if...else 语句if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。 if 嵌套语句你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。 switch 语句switch 语句用于基于不同条件执行不同动作。 select 语句select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。

注意：Go 没有三目运算符，所以不支持 ?: 形式的条件判断。

---

## Go 语言循环语句

Source: https://www.runoob.com/go/go-loops.html

## Go 语言循环语句

在不少实际问题中有许多具有规律性的重复操作，因此在程序中就需要重复执行某些语句。

以下为大多编程语言循环程序的流程图：

Go 语言提供了以下几种类型循环处理语句：

循环类型 描述 for 循环 重复执行语句块 循环嵌套 在 for 循环中嵌套一个或多个 for 循环

### 循环控制语句

循环控制语句可以控制循环体内语句的执行过程。

GO 语言支持以下几种循环控制语句：

控制语句 描述 break 语句 经常用于中断当前 for 循环或跳出 switch 语句 continue 语句 跳过当前循环的剩余语句，然后继续进行下一轮循环。 goto 语句 将控制转移到被标记的语句。

### 无限循环

如果循环中条件语句永远不为 false 则会进行无限循环，我们可以通过 for 循环语句中只设置一个条件表达式来执行无限循环：

### 实例

package main

import "fmt"

func main() {
for true {
fmt.Printf("这是无限循环。\n");
}
}

---

## Go 语言函数

Source: https://www.runoob.com/go/go-functions.html

## Go 语言函数

函数是基本的代码块，用于执行一个任务。

Go 语言最少有个 main() 函数。

你可以通过函数来划分不同功能，逻辑上每个函数执行的是指定的任务。

函数声明告诉了编译器函数的名称，返回类型，和参数。

Go 语言标准库提供了多种可动用的内置的函数。例如，len() 函数可以接受不同类型参数并返回该类型的长度。如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的元素个数。

### 函数定义

Go 语言函数定义格式如下：

```
func function_name( [parameter list] ) [return_types] {
函数体
}

```

函数定义解析：

- func：函数由 func 开始声明
- function_name：函数名称，参数列表和返回值类型构成了函数签名。
- parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
- return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
- 函数体：函数定义的代码集合。

#### 实例

以下实例为 max() 函数的代码，该函数传入两个整型参数 num1 和 num2，并返回这两个参数的最大值：

### 实例

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
/* 声明局部变量 */
var result int

if (num1 > num2) {
result = num1
} else {
result = num2
}
return result
}

### 函数调用

当创建函数时，你定义了函数需要做什么，通过调用该函数来执行指定任务。

调用函数，向函数传递参数，并返回值，例如：

### 实例

package main

import "fmt"

func main() {
/* 定义局部变量 */
var a int = 100
var b int = 200
var ret int

/* 调用函数并返回最大值 */
ret = max(a, b)

fmt.Printf( "最大值是 : %d\n", ret )
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
/* 定义局部变量 */
var result int

if (num1 > num2) {
result = num1
} else {
result = num2
}
return result
}

以上实例在 main() 函数中调用 max（）函数，执行结果为：

```
最大值是 : 200

```

### 函数返回多个值

Go 函数可以返回多个值，例如：

### 实例

package main

import "fmt"

func swap(x, y string) (string, string) {
return y, x
}

func main() {
a, b := swap("Google", "Runoob")
fmt.Println(a, b)
}

以上实例执行结果为：

```
Runoob Google

```

### 函数参数

函数如果使用参数，该变量可称为函数的形参。

形参就像定义在函数体内的局部变量。

调用函数，可以通过两种方式来传递参数：

传递类型 描述 值传递 值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。 引用传递 引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

### 函数用法

函数用法 描述 函数作为另外一个函数的实参 函数定义后可作为另外一个函数的实参数传入 闭包 闭包是匿名函数，可在动态编程中使用 方法 方法就是一个包含了接受者的函数

---

## Go 语言变量作用域

Source: https://www.runoob.com/go/go-scope-rules.html

## Go 语言变量作用域

作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。

Go 语言中变量可以在三个地方声明：

- 函数内定义的变量称为局部变量
- 函数外定义的变量称为全局变量
- 函数定义中的变量称为形式参数

接下来让我们具体了解局部变量、全局变量和形式参数。

### 局部变量

在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。

以下实例中 main() 函数使用了局部变量 a, b, c：

### 实例

package main

import "fmt"

func main() {
/* 声明局部变量 */
var a, b, c int

/* 初始化参数 */
a = 10
b = 20
c = a + b

fmt.Printf ("结果： a = %d, b = %d and c = %d\n", a, b, c)
}

以上实例执行输出结果为：

```
结果： a = 10, b = 20 and c = 30

```

### 全局变量

在函数体外声明的变量称之为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用。

全局变量可以在任何函数中使用，以下实例演示了如何使用全局变量：

### 实例

package main

import "fmt"

/* 声明全局变量 */
var g int

func main() {

/* 声明局部变量 */
var a, b int

/* 初始化参数 */
a = 10
b = 20
g = a + b

fmt.Printf("结果： a = %d, b = %d and g = %d\n", a, b, g)
}

以上实例执行输出结果为：

```
结果： a = 10, b = 20 and g = 30

```

Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。实例如下：

### 实例

package main

import "fmt"

/* 声明全局变量 */
var g int = 20

func main() {
/* 声明局部变量 */
var g int = 10

fmt.Printf ("结果： g = %d\n", g)
}

以上实例执行输出结果为：

```
结果： g = 10

```

### 形式参数

形式参数会作为函数的局部变量来使用。实例如下：

### 实例

package main

import "fmt"

/* 声明全局变量 */
var a int = 20;

func main() {
/* main 函数中声明局部变量 */
var a int = 10
var b int = 20
var c int = 0

fmt.Printf("main()函数中 a = %d\n", a);
c = sum( a, b);
fmt.Printf("main()函数中 c = %d\n", c);
}

/* 函数定义-两数相加 */
func sum(a, b int) int {
fmt.Printf("sum() 函数中 a = %d\n", a);
fmt.Printf("sum() 函数中 b = %d\n", b);

return a + b;
}

以上实例执行输出结果为：

```
main()函数中 a = 10
sum() 函数中 a = 10
sum() 函数中 b = 20
main()函数中 c = 30

```

### 初始化局部和全局变量

不同类型的局部和全局变量默认值为：

数据类型 初始化默认值 int 0 float32 0 pointer nil

---

## Go 语言数组

Source: https://www.runoob.com/go/go-arrays.html

## Go 语言数组

Go 语言提供了数组类型的数据结构。

数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整型、字符串或者自定义类型。

相对于去声明 number0, number1, ..., number99 的变量，使用数组形式 numbers[0], numbers[1] ..., numbers[99] 更加方便且易于扩展。

数组元素可以通过索引（位置）来读取（或者修改），索引从 0 开始，第一个元素索引为 0，第二个索引为 1，以此类推。

### 声明数组

Go 语言数组声明需要指定元素类型及元素个数，语法格式如下：

```
var arrayName [size]dataType

```
其中，arrayName 是数组的名称，size 是数组的大小，dataType 是数组中元素的数据类型。

以下定义了数组 balance 长度为 10 类型为 float32：

```
var balance [10]float32

```

### 初始化数组

以下演示了数组初始化：

以下实例声明一个名为 numbers 的整数数组，其大小为 5，在声明时，数组中的每个元素都会根据其数据类型进行默认初始化，对于整数类型，初始值为 0。

```
var numbers [5]int
```

还可以使用初始化列表来初始化数组的元素：

```
var numbers = [5]int{1, 2, 3, 4, 5}
```

以上代码声明一个大小为 5 的整数数组，并将其中的元素分别初始化为 1、2、3、4 和 5。

另外，还可以使用 := 简短声明语法来声明和初始化数组：

```
numbers := [5]int{1, 2, 3, 4, 5}
```

以上代码创建一个名为 numbers 的整数数组，并将其大小设置为 5，并初始化元素的值。

注意：在 Go 语言中，数组的大小是类型的一部分，因此不同大小的数组是不兼容的，也就是说 [5]int 和 [10]int 是不同的类型。

以下定义了数组 balance 长度为 5 类型为 float32，并初始化数组的元素：

```
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

```

我们也可以通过字面量在声明数组的同时快速初始化数组：

```
balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

```

如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度：

```

var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
或
balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
```

如果设置了数组的长度，我们还可以通过指定下标来初始化元素：

```
// 将索引为 1 和 3 的元素初始化
balance := [5]float32{1:2.0,3:7.0}
```

初始化数组中 {} 中的元素个数不能大于 [] 中的数字。

如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：

```
balance[4] = 50.0

```

以上实例读取了第五个元素。数组元素可以通过索引（位置）来读取（或者修改），索引从 0 开始，第一个元素索引为 0，第二个索引为 1，以此类推。

### 访问数组元素

数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值。例如：

```

var salary float32 = balance[9]

```

以上实例读取了数组 balance 第 10 个元素的值。

以下演示了数组完整操作（声明、赋值、访问）的实例：

### 实例 1

package main

import "fmt"

func main() {
var n [10]int /* n 是一个长度为 10 的数组 */
var i,j int

/* 为数组 n 初始化元素 */
for i = 0; i < 10; i++ {
n[i] = i + 100 /* 设置元素为 i + 100 */
}

/* 输出每个数组元素的值 */
for j = 0; j < 10; j++ {
fmt.Printf("Element[%d] = %d\n", j, n[j] )
}
}

以上实例执行结果如下：

```
Element[0] = 100
Element[1] = 101
Element[2] = 102
Element[3] = 103
Element[4] = 104
Element[5] = 105
Element[6] = 106
Element[7] = 107
Element[8] = 108
Element[9] = 109

```

### 实例 2

package main

import "fmt"

func main() {
var i,j,k int
// 声明数组的同时快速初始化数组
balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

/* 输出数组元素 */ ...
for i = 0; i < 5; i++ {
fmt.Printf("balance[%d] = %f\n", i, balance[i] )
}

balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
/* 输出每个数组元素的值 */
for j = 0; j < 5; j++ {
fmt.Printf("balance2[%d] = %f\n", j, balance2[j] )
}

// 将索引为 1 和 3 的元素初始化
balance3 := [5]float32{1:2.0,3:7.0}
for k = 0; k < 5; k++ {
fmt.Printf("balance3[%d] = %f\n", k, balance3[k] )
}
}

以上实例执行结果如下：

```
balance[0] = 1000.000000
balance[1] = 2.000000
balance[2] = 3.400000
balance[3] = 7.000000
balance[4] = 50.000000
balance2[0] = 1000.000000
balance2[1] = 2.000000
balance2[2] = 3.400000
balance2[3] = 7.000000
balance2[4] = 50.000000
balance3[0] = 0.000000
balance3[1] = 2.000000
balance3[2] = 0.000000
balance3[3] = 7.000000
balance3[4] = 0.000000
```

### 更多内容

数组对 Go 语言来说是非常重要的，以下我们将介绍数组更多的内容：

内容 描述 多维数组 Go 语言支持多维数组，最简单的多维数组是二维数组 向函数传递数组 你可以向函数传递数组参数

---

## Go 语言指针

Source: https://www.runoob.com/go/go-pointers.html

## Go 语言指针

Go 语言中指针是很容易学习的，Go 语言中使用指针可以更简单的执行一些任务。

接下来让我们来一步步学习 Go 语言指针。

我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。

Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。

以下实例演示了变量在内存中地址：

### 实例

package main

import "fmt"

func main() {
var a int = 10

fmt.Printf("变量的地址: %x\n", &a )
}

执行以上代码输出结果为：

```
变量的地址: 20818a220

```

现在我们已经了解了什么是内存地址和如何去访问它。接下来我们将具体介绍指针。

### 什么是指针

一个指针变量指向了一个值的内存地址。

类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：

```
var var_name *var-type

```

var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。以下是有效的指针声明：

```
var ip *int /* 指向整型*/
var fp *float32 /* 指向浮点型 */

```

本例中这是一个指向 int 和 float32 的指针。

### 如何使用指针

指针使用流程：

- 定义指针变量。
- 为指针变量赋值。
- 访问指针变量中指向地址的值。

在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。

### 实例

package main

import "fmt"

func main() {
var a int= 20 /* 声明实际变量 */
var ip *int /* 声明指针变量 */

ip = &a /* 指针变量的存储地址 */

fmt.Printf("a 变量的地址是: %x\n", &a )

/* 指针变量的存储地址 */
fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

/* 使用指针访问值 */
fmt.Printf("*ip 变量的值: %d\n", *ip )
}

以上实例执行输出结果为：

```
a 变量的地址是: 20818a220
ip 变量储存的指针地址: 20818a220
*ip 变量的值: 20

```

### Go 空指针

当一个指针被定义后没有分配到任何变量时，它的值为 nil。

nil 指针也称为空指针。

nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

一个指针变量通常缩写为 ptr。

查看以下实例：

### 实例

package main

import "fmt"

func main() {
var ptr *int

fmt.Printf("ptr 的值为 : %x\n", ptr )
}

以上实例输出结果为：

```
ptr 的值为 : 0

```

空指针判断：

```
if(ptr != nil) /* ptr 不是空指针 */
if(ptr == nil) /* ptr 是空指针 */

```

### Go指针更多内容

接下来我们将为大家介绍Go语言中更多的指针应用：

内容 描述 Go 指针数组 你可以定义一个指针数组来存储地址 Go 指向指针的指针 Go 支持指向指针的指针 Go 向函数传递指针参数 通过引用或地址传参，在函数调用时可以改变其值

---

## Go 语言结构体

Source: https://www.runoob.com/go/go-structures.html

## Go 语言结构体

Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。

结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。

结构体表示一项记录，比如保存图书馆的书籍记录，每本书有以下属性：

- Title ：标题
- Author ： 作者
- Subject：学科
- ID：书籍ID

### 定义结构体

结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：

```
type struct_variable_type struct {
member definition
member definition
...
member definition
}

```

一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：

```

variable_name := structure_variable_type {value1, value2...valuen}
或
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}

```

实例如下：

### 实例

package main

import "fmt"

type Books struct {
title string
author string
subject string
book_id int
}

func main() {

// 创建一个新的结构体
fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

// 也可以使用 key => value 格式
fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

// 忽略的字段为 0 或 空
fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}

输出结果为：

```
{Go 语言 www.runoob.com Go 语言教程 6495407}
{Go 语言 www.runoob.com Go 语言教程 6495407}
{Go 语言 www.runoob.com 0}
```

### 访问结构体成员

如果要访问结构体成员，需要使用点号 . 操作符，格式为：

```
结构体.成员名"
```

结构体类型变量使用 struct 关键字定义，实例如下：

### 实例

package main

import "fmt"

type Books struct {
title string
author string
subject string
book_id int
}

func main() {
var Book1 Books /* 声明 Book1 为 Books 类型 */
var Book2 Books /* 声明 Book2 为 Books 类型 */

/* book 1 描述 */
Book1.title = "Go 语言"
Book1.author = "www.runoob.com"
Book1.subject = "Go 语言教程"
Book1.book_id = 6495407

/* book 2 描述 */
Book2.title = "Python 教程"
Book2.author = "www.runoob.com"
Book2.subject = "Python 语言教程"
Book2.book_id = 6495700

/* 打印 Book1 信息 */
fmt.Printf( "Book 1 title : %s\n", Book1.title)
fmt.Printf( "Book 1 author : %s\n", Book1.author)
fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

/* 打印 Book2 信息 */
fmt.Printf( "Book 2 title : %s\n", Book2.title)
fmt.Printf( "Book 2 author : %s\n", Book2.author)
fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
}

以上实例执行运行结果为：

```
Book 1 title : Go 语言
Book 1 author : www.runoob.com
Book 1 subject : Go 语言教程
Book 1 book_id : 6495407
Book 2 title : Python 教程
Book 2 author : www.runoob.com
Book 2 subject : Python 语言教程
Book 2 book_id : 6495700

```

### 结构体作为函数参数

你可以像其他数据类型一样将结构体类型作为参数传递给函数。并以以上实例的方式访问结构体变量：

### 实例

package main

import "fmt"

type Books struct {
title string
author string
subject string
book_id int
}

func main() {
var Book1 Books /* 声明 Book1 为 Books 类型 */
var Book2 Books /* 声明 Book2 为 Books 类型 */

/* book 1 描述 */
Book1.title = "Go 语言"
Book1.author = "www.runoob.com"
Book1.subject = "Go 语言教程"
Book1.book_id = 6495407

/* book 2 描述 */
Book2.title = "Python 教程"
Book2.author = "www.runoob.com"
Book2.subject = "Python 语言教程"
Book2.book_id = 6495700

/* 打印 Book1 信息 */
printBook(Book1)

/* 打印 Book2 信息 */
printBook(Book2)
}

func printBook( book Books ) {
fmt.Printf( "Book title : %s\n", book.title)
fmt.Printf( "Book author : %s\n", book.author)
fmt.Printf( "Book subject : %s\n", book.subject)
fmt.Printf( "Book book_id : %d\n", book.book_id)
}

以上实例执行运行结果为：

```
Book title : Go 语言
Book author : www.runoob.com
Book subject : Go 语言教程
Book book_id : 6495407
Book title : Python 教程
Book author : www.runoob.com
Book subject : Python 语言教程
Book book_id : 6495700

```

### 结构体指针

你可以定义指向结构体的指针类似于其他指针变量，格式如下：

```
var struct_pointer *Books

```

以上定义的指针变量可以存储结构体变量的地址。查看结构体变量地址，可以将 & 符号放置于结构体变量前：

```
struct_pointer = &Book1

```

使用结构体指针访问结构体成员，使用 "." 操作符：

```
struct_pointer.title

```

接下来让我们使用结构体指针重写以上实例，代码如下：

### 实例

package main

import "fmt"

type Books struct {
title string
author string
subject string
book_id int
}

func main() {
var Book1 Books /* 声明 Book1 为 Books 类型 */
var Book2 Books /* 声明 Book2 为 Books 类型 */

/* book 1 描述 */
Book1.title = "Go 语言"
Book1.author = "www.runoob.com"
Book1.subject = "Go 语言教程"
Book1.book_id = 6495407

/* book 2 描述 */
Book2.title = "Python 教程"
Book2.author = "www.runoob.com"
Book2.subject = "Python 语言教程"
Book2.book_id = 6495700

/* 打印 Book1 信息 */
printBook(&Book1)

/* 打印 Book2 信息 */
printBook(&Book2)
}
func printBook( book *Books ) {
fmt.Printf( "Book title : %s\n", book.title)
fmt.Printf( "Book author : %s\n", book.author)
fmt.Printf( "Book subject : %s\n", book.subject)
fmt.Printf( "Book book_id : %d\n", book.book_id)
}

以上实例执行运行结果为：

```
Book title : Go 语言
Book author : www.runoob.com
Book subject : Go 语言教程
Book book_id : 6495407
Book title : Python 教程
Book author : www.runoob.com
Book subject : Python 语言教程
Book book_id : 6495700

```

---

## Go 语言切片(Slice)

Source: https://www.runoob.com/go/go-slice.html

## Go 语言切片(Slice)

Go 语言切片是对数组的抽象。

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

### 定义切片

你可以声明一个未指定大小的数组来定义切片：

```
var identifier []type

```

切片不需要说明长度。

或使用 make() 函数来创建切片:

```
var slice1 []type = make([]type, len)

也可以简写为

slice1 := make([]type, len)

```

也可以指定容量，其中 capacity 为可选参数。

```
make([]T, length, capacity)

```

这里 len 是数组的长度并且也是切片的初始长度。

#### 切片初始化

```
s :=[] int {1,2,3 }

```

直接初始化切片，[] 表示是切片类型，{1,2,3} 初始化值依次是 1,2,3，其 cap=len=3。

```
s := arr[:]

```

初始化切片 s，是数组 arr 的引用。

```
s := arr[startIndex:endIndex]

```

将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片。

```
s := arr[startIndex:]

```

默认 endIndex 时将表示一直到arr的最后一个元素。

```
s := arr[:endIndex]

```

默认 startIndex 时将表示从 arr 的第一个元素开始。

```
s1 := s[startIndex:endIndex]

```

通过切片 s 初始化切片 s1。

```
s :=make([]int,len,cap)

```

通过内置函数 make() 初始化切片s，[]int 标识为其元素类型为 int 的切片。

### len() 和 cap() 函数

切片是可索引的，并且可以由 len() 方法获取长度。

切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

以下为具体实例：

### 实例

package main

import "fmt"

func main() {
var numbers = make([]int,3,5)

printSlice(numbers)
}

func printSlice(x []int){
fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

以上实例运行输出结果为:

```
len=3 cap=5 slice=[0 0 0]

```

### 空(nil)切片

一个切片在未初始化之前默认为 nil，长度为 0，实例如下：

### 实例

package main

import "fmt"

func main() {
var numbers []int

printSlice(numbers)

if(numbers == nil){
fmt.Printf("切片是空的")
}
}

func printSlice(x []int){
fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

以上实例运行输出结果为:

```
len=0 cap=0 slice=[]
切片是空的

```

### 切片截取

可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]，实例如下：

### 实例

package main

import "fmt"

func main() {
/* 创建切片 */
numbers := []int{0,1,2,3,4,5,6,7,8}
printSlice(numbers)

/* 打印原始切片 */
fmt.Println("numbers ==", numbers)

/* 打印子切片从索引1(包含) 到索引4(不包含)*/
fmt.Println("numbers[1:4] ==", numbers[1:4])

/* 默认下限为 0*/
fmt.Println("numbers[:3] ==", numbers[:3])

/* 默认上限为 len(s)*/
fmt.Println("numbers[4:] ==", numbers[4:])

numbers1 := make([]int,0,5)
printSlice(numbers1)

/* 打印子切片从索引 0(包含) 到索引 2(不包含) */
number2 := numbers[:2]
printSlice(number2)

/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
number3 := numbers[2:5]
printSlice(number3)

}

func printSlice(x []int){
fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

执行以上代码输出结果为：

```
len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]
numbers == [0 1 2 3 4 5 6 7 8]
numbers[1:4] == [1 2 3]
numbers[:3] == [0 1 2]
numbers[4:] == [4 5 6 7 8]
len=0 cap=5 slice=[]
len=2 cap=9 slice=[0 1]
len=3 cap=7 slice=[2 3 4]

```

### append() 和 copy() 函数

如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。

下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法。

### 实例

package main

import "fmt"

func main() {
var numbers []int
printSlice(numbers)

/* 允许追加空切片 */
numbers = append(numbers, 0)
printSlice(numbers)

/* 向切片添加一个元素 */
numbers = append(numbers, 1)
printSlice(numbers)

/* 同时添加多个元素 */
numbers = append(numbers, 2,3,4)
printSlice(numbers)

/* 创建切片 numbers1 是之前切片的两倍容量*/
numbers1 := make([]int, len(numbers), (cap(numbers))*2)

/* 拷贝 numbers 的内容到 numbers1 */
copy(numbers1,numbers)
printSlice(numbers1)
}

func printSlice(x []int){
fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

以上代码执行输出结果为：

```
len=0 cap=0 slice=[]
len=1 cap=1 slice=[0]
len=2 cap=2 slice=[0 1]
len=5 cap=6 slice=[0 1 2 3 4]
len=5 cap=12 slice=[0 1 2 3 4]

```

---

## Go 语言范围(Range)

Source: https://www.runoob.com/go/go-range.html

## Go 语言范围(Range)

Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：

```
for key, value := range oldMap {
newMap[key] = value
}

```

以上代码中的 key 和 value 是可以省略。

如果只想读取 key，格式如下：

```
for key := range oldMap
```

或者这样：

for key, _ := range oldMap

如果只想读取 value，格式如下：

```
for _, value := range oldMap
```

### 实例

#### 数组和切片

遍历简单的切片，2**%d 的结果为 2 对应的次方数：

### 实例

package main

import "fmt"

// 声明一个包含 2 的幂次方的切片
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
// 遍历 pow 切片，i 是索引，v 是值
for i, v := range pow {
// 打印 2 的 i 次方等于 v
fmt.Printf("2**%d = %d\n", i, v)
}
}

以上实例运行输出结果为：

```

2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128

```

#### 字符串

range 迭代字符串时，返回每个字符的索引和 Unicode 代码点（rune）。

### 实例

package main

import "fmt"

func main() {
for i, c := range "hello" {
fmt.Printf("index: %d, char: %c\n", i, c)
}
}

以上实例运行输出结果为:

```
index: 0, char: h
index: 1, char: e
index: 2, char: l
index: 3, char: l
index: 4, char: o
```

#### 映射（Map）

for 循环的 range 格式可以省略 key 和 value，如下实例：

### 实例

package main

import "fmt"

func main() {
// 创建一个空的 map，key 是 int 类型，value 是 float32 类型
map1 := make(map[int]float32)

// 向 map1 中添加 key-value 对
map1[1] = 1.0
map1[2] = 2.0
map1[3] = 3.0
map1[4] = 4.0

// 遍历 map1，读取 key 和 value
for key, value := range map1 {
// 打印 key 和 value
fmt.Printf("key is: %d - value is: %f\n", key, value)
}

// 遍历 map1，只读取 key
for key := range map1 {
// 打印 key
fmt.Printf("key is: %d\n", key)
}

// 遍历 map1，只读取 value
for _, value := range map1 {
// 打印 value
fmt.Printf("value is: %f\n", value)
}
}

以上实例运行输出结果为:

```
key is: 4 - value is: 4.000000
key is: 1 - value is: 1.000000
key is: 2 - value is: 2.000000
key is: 3 - value is: 3.000000
key is: 1
key is: 2
key is: 3
key is: 4
value is: 1.000000
value is: 2.000000
value is: 3.000000
value is: 4.000000

```

#### 通道（Channel）

range 遍历从通道接收的值，直到通道关闭。

### 实例

package main

import "fmt"

func main() {
ch := make(chan int, 2)
ch <- 1
ch <- 2
close(ch)

for v := range ch {
fmt.Println(v)
}
}

以上实例运行输出结果为:

```
1
2
```

#### 忽略值

在遍历时可以使用 _ 来忽略索引或值。

### 实例

package main

import "fmt"

func main() {
nums := []int{2, 3, 4}

// 忽略索引
for _, num := range nums {
fmt.Println("value:", num)
}

// 忽略值
for i := range nums {
fmt.Println("index:", i)
}
}

以上实例运行输出结果为:

```
value: 2
value: 3
value: 4
index: 0
index: 1
index: 2
```

#### 其他

range 遍历其他数据结构：

### 实例

package main
import "fmt"
func main() {
//这是我们使用 range 去求一个 slice 的和。使用数组跟这个很类似
nums := []int{2, 3, 4}
sum := 0
for _, num := range nums {
sum += num
}
fmt.Println("sum:", sum)
//在数组上使用 range 将传入索引和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
for i, num := range nums {
if num == 3 {
fmt.Println("index:", i)
}
}
//range 也可以用在 map 的键值对上。
kvs := map[string]string{"a": "apple", "b": "banana"}
for k, v := range kvs {
fmt.Printf("%s -> %s\n", k, v)
}

//range也可以用来枚举 Unicode 字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
for i, c := range "go" {
fmt.Println(i, c)
}
}

以上实例运行输出结果为：

```

sum: 9
index: 1
a -> apple
b -> banana
0 103
1 111

```

---

## Go 语言Map(集合)

Source: https://www.runoob.com/go/go-map.html

## Go 语言Map(集合)

Map 是一种无序的键值对的集合。

Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。

Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，遍历 Map 时返回的键值对的顺序是不确定的。

在获取 Map 的值时，如果键不存在，返回该类型的零值，例如 int 类型的零值是 0，string 类型的零值是 ""。

Map 是引用类型，如果将一个 Map 传递给一个函数或赋值给另一个变量，它们都指向同一个底层数据结构，因此对 Map 的修改会影响到所有引用它的变量。

#### 定义 Map

可以使用内建函数 make 或使用 map 关键字来定义 Map:

```
/* 使用 make 函数 */
map_variable := make(map[KeyType]ValueType, initialCapacity)
```

其中 KeyType 是键的类型，ValueType 是值的类型，initialCapacity 是可选的参数，用于指定 Map 的初始容量。Map 的容量是指 Map 中可以保存的键值对的数量，当 Map 中的键值对数量达到容量时，Map 会自动扩容。如果不指定 initialCapacity，Go 语言会根据实际情况选择一个合适的值。

### 实例

// 创建一个空的 Map
m := make(map[string]int)

// 创建一个初始容量为 10 的 Map
m := make(map[string]int, 10)

也可以使用字面量创建 Map：

```

// 使用字面量创建 Map
m := map[string]int{
"apple": 1,
"banana": 2,
"orange": 3,
}

```

获取元素：

```
// 获取键值对
v1 := m["apple"]
v2, ok := m["pear"] // 如果键不存在，ok 的值为 false，v2 的值为该类型的零值
```

修改元素：

```
// 修改键值对
m["apple"] = 5
```

获取 Map 的长度：

```
// 获取 Map 的长度
len := len(m)
```

遍历 Map：

```
// 遍历 Map
for k, v := range m {
fmt.Printf("key=%s, value=%d\n", k, v)
}
```

删除元素：

```
// 删除键值对
delete(m, "banana")
```

#### 实例

下面实例演示了创建和使用map:

### 实例

package main

import "fmt"

func main() {
var siteMap map[string]string /*创建集合 */
siteMap = make(map[string]string)

/* map 插入 key - value 对,各个国家对应的首都 */
siteMap [ "Google" ] = "谷歌"
siteMap [ "Runoob" ] = "菜鸟教程"
siteMap [ "Baidu" ] = "百度"
siteMap [ "Wiki" ] = "维基百科"

/*使用键输出地图值 */
for site := range siteMap {
fmt.Println(site, "首都是", siteMap [site])
}

/*查看元素在集合中是否存在 */
name, ok := siteMap [ "Facebook" ] /*如果确定是真实的,则存在,否则不存在 */
/*fmt.Println(capital) */
/*fmt.Println(ok) */
if (ok) {
fmt.Println("Facebook 的 站点是", name)
} else {
fmt.Println("Facebook 站点不存在")
}
}

以上实例运行结果为：

```

Wiki 首都是 维基百科
Google 首都是 谷歌
Runoob 首都是 菜鸟教程
Baidu 首都是 百度
Facebook 站点不存在

```

### delete() 函数

delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。实例如下：

### 实例

package main

import "fmt"

func main() {
/* 创建map */
countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

fmt.Println("原始地图")

/* 打印地图 */
for country := range countryCapitalMap {
fmt.Println(country, "首都是", countryCapitalMap [ country ])
}

/*删除元素*/ delete(countryCapitalMap, "France")
fmt.Println("法国条目被删除")

fmt.Println("删除元素后地图")

/*打印地图*/
for country := range countryCapitalMap {
fmt.Println(country, "首都是", countryCapitalMap [ country ])
}
}

以上实例运行结果为：

```

原始地图
India 首都是 New delhi
France 首都是 Paris
Italy 首都是 Rome
Japan 首都是 Tokyo
法国条目被删除
删除元素后地图
Italy 首都是 Rome
Japan 首都是 Tokyo
India 首都是 New delhi

```

---

## Go 语言递归函数

Source: https://www.runoob.com/go/go-recursion.html

## Go 语言递归函数

递归是一种函数直接或间接调用自身的编程技术。

递归函数通常包含两个部分：

- 基准条件（Base Case）：这是递归的终止条件，防止函数无限调用自身。
- 递归条件（Recursive Case）：这是函数调用自身的部分，用于将问题分解为更小的子问题。

在 Go 语言中，递归的使用与其他语言类似，但需要注意 Go 的一些特性。

语法格式如下：

func recursion() {
recursion() /* 函数调用自身 */
}

func main() {
recursion()
}

Go 语言支持递归，但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。

递归函数对于解决数学上的问题是非常有用的，就像计算阶乘，生成斐波那契数列等。

### 阶乘

阶乘是一个正整数的乘积，表示为 `n!`。例如：

```
5! = 5 * 4 * 3 * 2 * 1 = 120

```

以下实例通过 Go 语言的递归函数实例阶乘：

### 实例

package main

import "fmt"

// 递归函数计算阶乘
func factorial(n int) int {
// 基准条件
if n == 0 {
return 1
}
// 递归条件
return n * factorial(n-1)
}

func main() {
fmt.Println(factorial(5)) // 输出: 120
}

##### 代码解释

- 基准条件：当 `n` 等于 0 时，函数返回 1，因为 `0!` 定义为 1。
- 递归条件：函数返回 `n` 乘以 `factorial(n-1)` 的结果，逐步将问题分解为更小的子问题。

以上实例执行输出结果为：

```

120

```

### 斐波那契数列

以下实例通过 Go 语言的递归函数实现斐波那契数列：

### 实例

package main

import "fmt"

func fibonacci(n int) int {
if n < 2 {
return n
}
return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
var i int
for i = 0; i < 10; i++ {
fmt.Printf("%d\t", fibonacci(i))
}
}

以上实例执行输出结果为：

```

0 1 1 2 3 5 8 13 21 34

```

### 求平方根

以下实例通过 Go 语言使用递归方法实现求平方根的代码：

### 实例

package main

import (
"fmt"
)

func sqrtRecursive(x, guess, prevGuess, epsilon float64) float64 {
if diff := guess*guess - x; diff < epsilon && -diff < epsilon {
return guess
}

newGuess := (guess + x/guess) / 2
if newGuess == prevGuess {
return guess
}

return sqrtRecursive(x, newGuess, guess, epsilon)
}

func sqrt(x float64) float64 {
return sqrtRecursive(x, 1.0, 0.0, 1e-9)
}

func main() {
x := 25.0
result := sqrt(x)
fmt.Printf("%.2f 的平方根为 %.6f\n", x, result)
}

以上实例中，sqrtRecursive 函数使用递归方式实现平方根的计算。

sqrtRecursive 函数接受四个参数：

- x 表示待求平方根的数
- guess 表示当前猜测的平方根值
- prevGuess 表示上一次的猜测值
- epsilon 表示精度要求（即接近平方根的程度）

递归的终止条件是当前猜测的平方根与上一次猜测的平方根非常接近，差值小于给定的精度 epsilon。

在 sqrt 函数中，我们调用 sqrtRecursive 来计算平方根，并传入初始值和精度要求，然后在 main 函数中，我们调用 sqrt 函数来求解平方根，并将结果打印出来。

执行以上代码输出结果为：

```
25.00 的平方根为 5.000000
```

### 递归的优缺点

#### 优点

- 简洁性：递归代码通常比迭代代码更简洁，易于理解。
- 问题分解：递归天然适合解决可以分解为相似子问题的问题，如树遍历、分治算法等。

#### 缺点

- 性能开销：递归调用会占用栈空间，可能导致栈溢出，尤其是在深度递归时。
- 调试困难：递归代码可能较难调试，尤其是在递归深度较大时。

### 递归与迭代

递归和迭代是解决问题的两种不同方法。递归通过函数调用自身来解决问题，而迭代则通过循环结构（如 `for` 循环）来重复执行代码块。

#### 递归 vs 迭代

特性 递归 迭代 代码简洁性 通常更简洁 可能更冗长 性能 可能较慢，占用栈空间 通常更快，占用较少内存 适用场景 适合分解为子问题的问题 适合线性或简单重复的问题

### 递归的常见应用

递归在许多算法和数据结构中都有广泛应用，例如：

- 树和图的遍历：如深度优先搜索（DFS）。
- 分治算法：如归并排序、快速排序。
- 动态规划：如斐波那契数列的计算。

#### 文件目录遍历

### 实例

import (
"fmt"
"os"
"path/filepath"
)

func walkDir(dir string, indent string) {
entries, err := os.ReadDir(dir)
if err != nil {
return
}

for _, entry := range entries {
fmt.Println(indent + entry.Name())
if entry.IsDir() {
walkDir(filepath.Join(dir, entry.Name()), indent+" ")
}
}
}

func main() {
walkDir(".", "")
}

在 Go 中使用递归时，应特别注意基线条件（终止条件）的正确性，避免无限递归。对于性能敏感或可能深度递归的场景，建议考虑迭代实现或使用 channel/goroutine 等 Go 特有机制。

---

## Go 语言类型转换

Source: https://www.runoob.com/go/go-type-casting.html

## Go 语言类型转换

类型转换用于将一种数据类型的变量转换为另外一种类型的变量。

Go 语言类型转换基本格式如下：

```

type_name(expression)

```

type_name 为类型，expression 为表达式。

#### 数值类型转换

将整型转换为浮点型：

var a int = 10
var b float64 = float64(a)

以下实例中将整型转化为浮点型，并计算结果，将结果赋值给浮点型变量：

### 实例

package main

import "fmt"

func main() {
var sum int = 17
var count int = 5
var mean float32

mean = float32(sum)/float32(count)
fmt.Printf("mean 的值为: %f\n",mean)
}

以上实例执行输出结果为：

```

mean 的值为: 3.400000

```

#### 字符串类型转换

将一个字符串转换成另一个类型，可以使用以下语法：

```

var str string = "10"
var num int
num, _ = strconv.Atoi(str)
```

以上代码将字符串变量 str 转换为整型变量 num。

注意，strconv.Atoi 函数返回两个值，第一个是转换后的整型值，第二个是可能发生的错误，我们可以使用空白标识符 _ 来忽略这个错误。

以下实例将字符串转换为整数

### 实例

package main

import (
"fmt"
"strconv"
)

func main() {
str := "123"
num, err := strconv.Atoi(str)
if err != nil {
fmt.Println("转换错误:", err)
} else {
fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
}
}

以上实例执行输出结果为：

```

字符串 '123' 转换为整数为：123

```

以下实例将整数转换为字符串：

### 实例

package main

import (
"fmt"
"strconv"
)

func main() {
num := 123
str := strconv.Itoa(num)
fmt.Printf("整数 %d 转换为字符串为：'%s'\n", num, str)
}

以上实例执行输出结果为：

```
整数 123 转换为字符串为：'123'
```

以下实例将字符串转换为浮点数：

### 实例

package main

import (
"fmt"
"strconv"
)

func main() {
str := "3.14"
num, err := strconv.ParseFloat(str, 64)
if err != nil {
fmt.Println("转换错误:", err)
} else {
fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str, num)
}
}

以上实例执行输出结果为：

```
字符串 '3.14' 转为浮点型为：3.140000
```

以下实例将浮点数转换为字符串：

### 实例

package main

import (
"fmt"
"strconv"
)

func main() {
num := 3.14
str := strconv.FormatFloat(num, 'f', 2, 64)
fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", num, str)
}

以上实例执行输出结果为：

```

浮点数 3.140000 转为字符串为：'3.14'
```

#### 接口类型转换

接口类型转换有两种情况：类型断言和类型转换。

#### 类型断言

类型断言用于将接口类型转换为指定类型，其语法为：

```

value.(type)
或者
value.(T)
```

其中 value 是接口类型的变量，type 或 T 是要转换成的类型。

如果类型断言成功，它将返回转换后的值和一个布尔值，表示转换是否成功。

### 实例

package main

import "fmt"

func main() {
var i interface{} = "Hello, World"
str, ok := i.(string)
if ok {
fmt.Printf("'%s' is a string\n", str)
} else {
fmt.Println("conversion failed")
}
}

以上实例中，我们定义了一个接口类型变量 i，并将它赋值为字符串 "Hello, World"。然后，我们使用类型断言将 i 转换为字符串类型，并将转换后的值赋值给变量 str。最后，我们使用 ok 变量检查类型转换是否成功，如果成功，我们打印转换后的字符串；否则，我们打印转换失败的消息。

#### 类型转换

类型转换用于将一个接口类型的值转换为另一个接口类型，其语法为：

```
T(value)
```

T 是目标接口类型，value 是要转换的值。

在类型转换中，我们必须保证要转换的值和目标接口类型之间是兼容的，否则编译器会报错。

### 实例

package main

import "fmt"

// 定义一个接口 Writer
type Writer interface {
Write([]byte) (int, error)
}

// 实现 Writer 接口的结构体 StringWriter
type StringWriter struct {
str string
}

// 实现 Write 方法
func (sw *StringWriter) Write(data []byte) (int, error) {
sw.str += string(data)
return len(data), nil
}

func main() {
// 创建一个 StringWriter 实例并赋值给 Writer 接口变量
var w Writer = &StringWriter{}

// 将 Writer 接口类型转换为 StringWriter 类型
sw := w.(*StringWriter)

// 修改 StringWriter 的字段
sw.str = "Hello, World"

// 打印 StringWriter 的字段值
fmt.Println(sw.str)
}

解析：

- 定义接口和结构体：

- `Writer` 接口定义了 `Write` 方法。
- `StringWriter` 结构体实现了 `Write` 方法。
- 类型转换：

- 将 `StringWriter` 实例赋值给 `Writer` 接口变量 `w`。
- 使用 `w.(*StringWriter)` 将 `Writer` 接口类型转换为 `StringWriter` 类型。
- 访问字段：

- 修改 `StringWriter` 的字段 `str`，并打印其值。

#### 空接口类型

空接口 interface{} 可以持有任何类型的值。在实际应用中，空接口经常被用来处理多种类型的值。

### 实例

package main

import (
"fmt"
)

func printValue(v interface{}) {
switch v := v.(type) {
case int:
fmt.Println("Integer:", v)
case string:
fmt.Println("String:", v)
default:
fmt.Println("Unknown type")
}
}

func main() {
printValue(42)
printValue("hello")
printValue(3.14)
}

在这个例子中，printValue 函数接受一个空接口类型的参数，并使用类型断言和类型选择来处理不同的类型。

---

## Go 语言接口

Source: https://www.runoob.com/go/go-interfaces.html

## Go 语言接口

接口（interface）是 Go 语言中的一种类型，用于定义行为的集合，它通过描述类型必须实现的方法，规定了类型的行为契约。

Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

Go 的接口设计简单却功能强大，是实现多态和解耦的重要工具。

接口可以让我们将不同的类型绑定到一组公共的方法上，从而实现多态和灵活的设计。

#### 接口的特点

隐式实现：

- Go 中没有关键字显式声明某个类型实现了某个接口。
- 只要一个类型实现了接口要求的所有方法，该类型就自动被认为实现了该接口。

接口类型变量：

- 接口变量可以存储实现该接口的任意值。
- 接口变量实际上包含了两个部分：

- 动态类型：存储实际的值类型。
- 动态值：存储具体的值。

零值接口：

- 接口的零值是 `nil`。
- 一个未初始化的接口变量其值为 `nil`，且不包含任何动态类型或值。

空接口：

- 定义为 `interface{}`，可以表示任何类型。

#### 接口的常见用法

- 多态：不同类型实现同一接口，实现多态行为。
- 解耦：通过接口定义依赖关系，降低模块之间的耦合。
- 泛化：使用空接口 `interface{}` 表示任意类型。

### 接口定义和实现

接口定义使用关键字 interface，其中包含方法声明。

### 实例

/* 定义接口 */
type interface_name interface {
method_name1 [return_type]
method_name2 [return_type]
method_name3 [return_type]
...
method_namen [return_type]
}

/* 定义结构体 */
type struct_name struct {
/* variables */
}

/* 实现接口方法 */
func (struct_name_variable struct_name) method_name1() [return_type] {
/* 方法实现 */
}
...
func (struct_name_variable struct_name) method_namen() [return_type] {
/* 方法实现*/
}

定义一个简单接口:

```
type Shape interface {
Area() float64
Perimeter() float64
}
```

- `Shape` 是一个接口，定义了两个方法：`Area` 和 `Perimeter`。
- 任意类型只要实现了这两个方法，就被认为实现了 `Shape` 接口。

实现接口: 类型通过实现接口要求的所有方法来实现接口。

### 实例

package main

import (
"fmt"
"math"
)

// 定义接口
type Shape interface {
Area() float64
Perimeter() float64
}

// 定义一个结构体
type Circle struct {
Radius float64
}

// Circle 实现 Shape 接口
func (c Circle) Area() float64 {
return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
return 2 * math.Pi * c.Radius
}

func main() {
c := Circle{Radius: 5}
var s Shape = c // 接口变量可以存储实现了接口的类型
fmt.Println("Area:", s.Area())
fmt.Println("Perimeter:", s.Perimeter())
}

执行以上代码，输出结果为：

```
Area: 78.53981633974483
Perimeter: 31.41592653589793
```

### 空接口

空接口 `interface{}` 是 Go 的特殊接口，表示所有类型的超集。

- 任意类型都实现了空接口。
- 常用于需要存储任意类型数据的场景，如泛型容器、通用参数等。

### 实例

package main

import "fmt"

func printValue(val interface{}) {
fmt.Printf("Value: %v, Type: %T\n", val, val)
}

func main() {
printValue(42) // int
printValue("hello") // string
printValue(3.14) // float64
printValue([]int{1, 2}) // slice
}

执行以上代码，输出结果为：

```
Value: 42, Type: int
Value: hello, Type: string
Value: 3.14, Type: float64
Value: [1 2], Type: []int
```

### 类型断言

类型断言用于从接口类型中提取其底层值。

基本语法:

```

value := iface.(Type)
```

- `iface` 是接口变量。
- `Type` 是要断言的具体类型。
- 如果类型不匹配，会触发 `panic`。

### 实例

package main

import "fmt"

func main() {
var i interface{} = "hello"
str := i.(string) // 类型断言
fmt.Println(str) // 输出：hello
}

#### 带检查的类型断言

为了避免 panic，可以使用带检查的类型断言：

```

value, ok := iface.(Type)
```

- `ok` 是一个布尔值，表示断言是否成功。
- 如果断言失败，`value` 为零值，`ok` 为 `false`。

### 实例

package main

import "fmt"

func main() {
var i interface{} = 42
if str, ok := i.(string); ok {
fmt.Println("String:", str)
} else {
fmt.Println("Not a string")
}
}

执行以上代码，输出结果为：

```
Not a string
```

### 类型选择（type switch）

type switch 是 Go 中的语法结构，用于根据接口变量的具体类型执行不同的逻辑。

### 实例

package main

import "fmt"

func printType(val interface{}) {
switch v := val.(type) {
case int:
fmt.Println("Integer:", v)
case string:
fmt.Println("String:", v)
case float64:
fmt.Println("Float:", v)
default:
fmt.Println("Unknown type")
}
}

func main() {
printType(42)
printType("hello")
printType(3.14)
printType([]int{1, 2, 3})
}

执行以上代码，输出结果为：

```
Integer: 42
String: hello
Float: 3.14
Unknown type
```

### 接口组合

接口可以通过嵌套组合，实现更复杂的行为描述。

### 实例

package main

import "fmt"

type Reader interface {
Read() string
}

type Writer interface {
Write(data string)
}

type ReadWriter interface {
Reader
Writer
}

type File struct{}

func (f File) Read() string {
return "Reading data"
}

func (f File) Write(data string) {
fmt.Println("Writing data:", data)
}

func main() {
var rw ReadWriter = File{}
fmt.Println(rw.Read())
rw.Write("Hello, Go!")
}

### 动态值和动态类型

接口变量实际上包含了两部分：

- 动态类型：接口变量存储的具体类型。
- 动态值：具体类型的值。

动态值和动态类型示例：

### 实例

package main

import "fmt"

func main() {
var i interface{} = 42
fmt.Printf("Dynamic type: %T, Dynamic value: %v\n", i, i)
}

执行以上代码，输出结果为：

```

Dynamic type: int, Dynamic value: 42
```

### 接口的零值

接口的零值是 nil。

当接口变量的动态类型和动态值都为 nil 时，接口变量为 nil。

接口零值示例：

### 实例

package main

import "fmt"

func main() {
var i interface{}
fmt.Println(i == nil) // 输出：true
}

### 练习实例

以下两个实例演示了接口的使用：

### 实例 1

package main

import (
"fmt"
)

type Phone interface {
call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
fmt.Println("I am iPhone, I can call you!")
}

func main() {
var phone Phone

phone = new(NokiaPhone)
phone.call()

phone = new(IPhone)
phone.call()

}

在上面的例子中，我们定义了一个接口 Phone，接口里面有一个方法 call()。然后我们在 main 函数里面定义了一个 Phone 类型变量，并分别为之赋值为 NokiaPhone 和 IPhone。然后调用 call() 方法，输出结果如下：

```

I am Nokia, I can call you!
I am iPhone, I can call you!

```

第二个接口实例：

### 实例

package main

import "fmt"

type Shape interface {
area() float64
}

type Rectangle struct {
width float64
height float64
}

func (r Rectangle) area() float64 {
return r.width * r.height
}

type Circle struct {
radius float64
}

func (c Circle) area() float64 {
return 3.14 * c.radius * c.radius
}

func main() {
var s Shape

s = Rectangle{width: 10, height: 5}
fmt.Printf("矩形面积: %f\n", s.area())

s = Circle{radius: 3}
fmt.Printf("圆形面积: %f\n", s.area())
}

以上实例中，我们定义了一个 Shape 接口，它定义了一个方法 area()，该方法返回一个 float64 类型的面积值。然后，我们定义了两个结构体 Rectangle 和 Circle，它们分别实现了 Shape 接口的 area() 方法。在 main() 函数中，我们首先定义了一个 Shape 类型的变量 s，然后分别将 Rectangle 和 Circle 类型的实例赋值给它，并通过 area() 方法计算它们的面积并打印出来，输出结果如下：

```
矩形面积: 50.000000
圆形面积: 28.260000
```

需要注意的是，接口类型变量可以存储任何实现了该接口的类型的值。在示例中，我们将 Rectangle 和 Circle 类型的实例都赋值给了 Shape 类型的变量 s，并通过 area() 方法调用它们的面积计算方法。

---

## Go 语言泛型

Source: https://www.runoob.com/go/go-generics.html

## Go 语言泛型

泛型是 Go 语言在 1.18 版本中引入的重要特性，它让开发者能够编写更加灵活和可重用的代码。

泛型主要通过以下两个核心概念来实现：

- 类型参数（Type Parameters）：允许你在函数或类型定义中使用一个或多个类型作为参数。
- 类型约束（Type Constraints）：指定类型参数必须满足的条件，确保在函数内部可以安全地操作这些类型。

概念 作用 示例 类型参数 在函数或类型名后声明，表示待定的类型。 `[T any]` 类型约束 定义类型参数必须满足的条件（如支持的操作符或方法）。 `int，float64，comparable，constraints.Ordered，any` `any` 约束类型参数为任何类型。 `[T any]` `comparable` 约束类型参数为可比较的类型。 `[K comparable]`

泛型（Generics）允许我们编写不依赖特定数据类型的代码。

在引入泛型之前，如果我们想要处理不同类型的数据，通常需要为每种类型编写重复的函数。

传统方式的局限性：

### 实例

// 处理 int 类型的函数
func MaxInt(a, b int) int {
if a > b {
return a
}
return b
}

// 处理 float64 类型的函数
func MaxFloat(a, b float64) float64 {
if a > b {
return a
}
return b
}

使用泛型的解决方案：

### 实例

// 一个函数处理多种类型
func Max[T comparable](a, b T) T {
if a > b {
return a
}
return b
}

### 泛型语法详解

#### 类型参数声明

泛型函数和类型通过类型参数列表来声明，语法为 `[类型参数 约束]`。

### 实例

// 基本语法结构
func 函数名[T 约束](参数 T) 返回值类型 {
// 函数体
}

type 类型名[T 约束] struct {
// 结构体字段
}

#### 类型参数命名约定

- 通常使用大写字母：`T`、`K`、`V`、`E` 等
- `T`：表示 Type（类型）
- `K`：表示 Key（键）
- `V`：表示 Value（值）
- `E`：表示 Element（元素）

### 约束（Constraints）

约束定义了类型参数必须满足的条件，是泛型的核心概念。

#### 内置约束

##### 1. `any` 约束

`any` 是空接口 `interface{}` 的别名，表示任何类型都可以。

### 实例

func PrintAny[T any](value T) {
fmt.Printf("Value: %v, Type: %T\n", value, value)
}

// 使用示例
PrintAny(42) // Value: 42, Type: int
PrintAny("hello") // Value: hello, Type: string
PrintAny(3.14) // Value: 3.14, Type: float64

##### 2. `comparable` 约束

`comparable` 表示类型支持 `==` 和 `!=` 操作符。

### 实例

func FindIndex[T comparable](slice []T, target T) int {
for i, v := range slice {
if v == target {
return i
}
}
return -1
}

// 使用示例
numbers := []int{1, 2, 3, 4, 5}
fmt.Println(FindIndex(numbers, 3)) // 输出: 2

names := []string{"Alice", "Bob", "Charlie"}
fmt.Println(FindIndex(names, "Bob")) // 输出: 1

##### 3. 联合约束（Union Constraints）

使用 `|` 运算符组合多个类型。

### 实例

// 数字类型约束
type Number interface {
int | int8 | int16 | int32 | int64 |
uint | uint8 | uint16 | uint32 | uint64 |
float32 | float64
}

func Add[T Number](a, b T) T {
return a + b
}

// 使用示例
fmt.Println(Add(10, 20)) // 输出: 30
fmt.Println(Add(3.14, 2.71)) // 输出: 5.85

#### 自定义约束

##### 1. 方法约束

定义需要特定方法的约束。

### 实例

// 定义 Stringer 约束
type Stringer interface {
String() string
}

func PrintString[T Stringer](value T) {
fmt.Println(value.String())
}

// 实现自定义类型
type Person struct {
Name string
Age int
}

func (p Person) String() string {
return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

// 使用示例
person := Person{Name: "Alice", Age: 25}
PrintString(person) // 输出: Alice (25 years old)

##### 2. 复杂约束

结合类型和方法要求。

### 实例

// 要求类型是数字且实现 String() 方法
type NumericStringer interface {
Number
String() string
}

### 泛型函数实践

#### 1. 通用工具函数

### 实例

// 交换两个值
func Swap[T any](a, b T) (T, T) {
return b, a
}

// 判断切片是否包含元素
func Contains[T comparable](slice []T, target T) bool {
for _, item := range slice {
if item == target {
return true
}
}
return false
}

// 去重函数
func Unique[T comparable](slice []T) []T {
seen := make(map[T]bool)
result := []T{}

for _, item := range slice {
if !seen[item] {
seen[item] = true
result = append(result, item)
}
}
return result
}

// 使用示例
func main() {
// Swap 示例
a, b := 10, 20
a, b = Swap(a, b)
fmt.Printf("a=%d, b=%d\n", a, b) // 输出: a=20, b=10

// Contains 示例
numbers := []int{1, 2, 3, 4, 5}
fmt.Println(Contains(numbers, 3)) // 输出: true

// Unique 示例
duplicates := []int{1, 2, 2, 3, 4, 4, 5}
unique := Unique(duplicates)
fmt.Println(unique) // 输出: [1 2 3 4 5]
}

#### 2. 数学运算函数

### 实例

// 求切片最大值
func Max[T Number](slice []T) T {
if len(slice) == 0 {
var zero T
return zero
}

max := slice[0]
for _, value := range slice[1:] {
if value > max {
max = value
}
}
return max
}

// 求切片最小值
func Min[T Number](slice []T) T {
if len(slice) == 0 {
var zero T
return zero
}

min := slice[0]
for _, value := range slice[1:] {
if value < min {
min = value
}
}
return min
}

// 求切片平均值
func Average[T Number](slice []T) float64 {
if len(slice) == 0 {
return 0
}

var sum T
for _, value := range slice {
sum += value
}
return float64(sum) / float64(len(slice))
}

// 使用示例
func main() {
ints := []int{1, 5, 3, 9, 2}
floats := []float64{1.1, 5.5, 3.3, 9.9, 2.2}

fmt.Printf("Max int: %d\n", Max(ints)) // 输出: 9
fmt.Printf("Min float: %.1f\n", Min(floats)) // 输出: 1.1
fmt.Printf("Average: %.2f\n", Average(floats)) // 输出: 4.40
}

### 泛型类型

#### 1. 泛型结构体

### 实例

// 泛型栈实现
type Stack[T any] struct {
elements []T
}

// 入栈
func (s *Stack[T]) Push(value T) {
s.elements = append(s.elements, value)
}

// 出栈
func (s *Stack[T]) Pop() (T, bool) {
if len(s.elements) == 0 {
var zero T
return zero, false
}

lastIndex := len(s.elements) - 1
value := s.elements[lastIndex]
s.elements = s.elements[:lastIndex]
return value, true
}

// 查看栈顶元素
func (s *Stack[T]) Peek() (T, bool) {
if len(s.elements) == 0 {
var zero T
return zero, false
}
return s.elements[len(s.elements)-1], true
}

// 判断栈是否为空
func (s *Stack[T]) IsEmpty() bool {
return len(s.elements) == 0
}

// 使用示例
func main() {
// 整数栈
intStack := Stack[int]{}
intStack.Push(1)
intStack.Push(2)
intStack.Push(3)

fmt.Println(intStack.Pop()) // 输出: 3 true

// 字符串栈
stringStack := Stack[string]{}
stringStack.Push("hello")
stringStack.Push("world")

fmt.Println(stringStack.Pop()) // 输出: world true
}

#### 2. 泛型映射（Map）

### 实例

// 线程安全的泛型映射
type SafeMap[K comparable, V any] struct {
data map[K]V
mutex sync.RWMutex
}

// 创建新的 SafeMap
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
return &SafeMap[K, V]{
data: make(map[K]V),
}
}

// 设置键值对
func (m *SafeMap[K, V]) Set(key K, value V) {
m.mutex.Lock()
defer m.mutex.Unlock()
m.data[key] = value
}

// 获取值
func (m *SafeMap[K, V]) Get(key K) (V, bool) {
m.mutex.RLock()
defer m.mutex.RUnlock()
value, exists := m.data[key]
return value, exists
}

// 删除键
func (m *SafeMap[K, V]) Delete(key K) {
m.mutex.Lock()
defer m.mutex.Unlock()
delete(m.data, key)
}

// 获取所有键
func (m *SafeMap[K, V]) Keys() []K {
m.mutex.RLock()
defer m.mutex.RUnlock()

keys := make([]K, 0, len(m.data))
for key := range m.data {
keys = append(keys, key)
}
return keys
}

// 使用示例
func main() {
// 创建字符串到整数的映射
scores := NewSafeMap[string, int]()
scores.Set("Alice", 95)
scores.Set("Bob", 87)

if score, exists := scores.Get("Alice"); exists {
fmt.Printf("Alice's score: %d\n", score) // 输出: Alice's score: 95
}

fmt.Println("Keys:", scores.Keys()) // 输出: Keys: [Alice Bob]
}

### 类型推断

Go 编译器能够自动推断类型参数，让代码更加简洁。

### 实例

// 无需显式指定类型
func main() {
// 类型推断示例
fmt.Println(Max([]int{1, 2, 3})) // 编译器推断 T 为 int
fmt.Println(Max([]float64{1.1, 2.2})) // 编译器推断 T 为 float64

// 显式指定类型（有时需要）
var result int = Max[int]([]int{1, 2, 3})
fmt.Println(result)
}

### 实践练习

#### 练习 1：实现泛型过滤器

编写一个 `Filter` 函数，根据条件过滤切片元素。

### 实例

// 你的实现代码在这里
func Filter[T any](slice []T, predicate func(T) bool) []T {
// 实现过滤逻辑
}

// 测试代码
func main() {
numbers := []int{1, 2, 3, 4, 5, 6}
even := Filter(numbers, func(n int) bool {
return n%2 == 0
})
fmt.Println(even) // 应该输出: [2 4 6]
}

#### 练习 2：实现泛型映射函数

编写一个 `Map` 函数，将切片中的每个元素转换为另一种类型。

### 实例

// 你的实现代码在这里
func Map[T any, U any](slice []T, mapper func(T) U) []U {
// 实现映射逻辑
}

// 测试代码
func main() {
numbers := []int{1, 2, 3, 4, 5}
strings := Map(numbers, func(n int) string {
return fmt.Sprintf("Number: %d", n)
})
fmt.Println(strings)
}

### 常见问题与注意事项

#### 1. 性能考虑

泛型在编译时进行类型特化，运行时性能与手写特定类型代码相当。

#### 2. 类型约束的选择

- 使用 `any` 时最灵活，但功能受限
- 使用 `comparable` 支持相等比较
- 使用联合约束限制可用的具体类型

#### 3. 错误处理

### 实例

// 良好的错误处理实践
func SafeMax[T Number](slice []T) (T, error) {
if len(slice) == 0 {
var zero T
return zero, errors.New("slice is empty")
}
return Max(slice), nil
}

---

## Go 错误处理

Source: https://www.runoob.com/go/go-error-handling.html

## Go 错误处理

Go 语言通过内置的错误接口提供了非常简单的错误处理机制。

Go 语言的错误处理采用显式返回错误的方式，而非传统的异常处理机制。这种设计使代码逻辑更清晰，便于开发者在编译时或运行时明确处理错误。

Go 的错误处理主要围绕以下机制展开：

- `error` 接口：标准的错误表示。
- 显式返回值：通过函数的返回值返回错误。
- 自定义错误：可以通过标准库或自定义的方式创建错误。
- `panic` 和 `recover`：处理不可恢复的严重错误。

### error 接口

Go 标准库定义了一个 error 接口，表示一个错误的抽象。

error 类型是一个接口类型，这是它的定义：

```

type error interface {
Error() string
}

```

- 实现 `error` 接口：任何实现了 `Error()` 方法的类型都可以作为错误。
- `Error()` 方法返回一个描述错误的字符串。

#### 使用 errors 包创建错误

我们可以在编码中通过实现 error 接口类型来生成错误信息。

创建一个简单错误：

### 实例

package main

import (
"errors"
"fmt"
)

func main() {
err := errors.New("this is an error")
fmt.Println(err) // 输出：this is an error
}

函数通常在最后的返回值中返回错误信息，使用 errors.New 可返回一个错误信息：

```

func Sqrt(f float64) (float64, error) {
if f < 0 {
return 0, errors.New("math: square root of negative number")
}
// 实现
}

```

在下面的例子中，我们在调用 Sqrt 的时候传递的一个负数，然后就得到了 non-nil 的 error 对象，将此对象与 nil 比较，结果为 true，所以 fmt.Println(fmt 包在处理 error 时会调用 Error 方法)被调用，以输出错误，请看下面调用的示例代码：

```

result, err:= Sqrt(-1)

if err != nil {
fmt.Println(err)
}

```

### 显式返回错误

Go 中，错误通常作为函数的返回值返回，开发者需要显式检查并处理。

显式返回错误：

### 实例

package main

import (
"errors"
"fmt"
)

func divide(a, b int) (int, error) {
if b == 0 {
return 0, errors.New("division by zero")
}
return a / b, nil
}

func main() {
result, err := divide(10, 0)
if err != nil {
fmt.Println("Error:", err)
} else {
fmt.Println("Result:", result)
}
}

输出：

```

Error: division by zero
```

### 自定义错误

通过定义自定义类型，可以扩展 error 接口。

自定义错误类型：

### 实例

package main

import (
"fmt"
)

type DivideError struct {
Dividend int
Divisor int
}

func (e *DivideError) Error() string {
return fmt.Sprintf("cannot divide %d by %d", e.Dividend, e.Divisor)
}

func divide(a, b int) (int, error) {
if b == 0 {
return 0, &DivideError{Dividend: a, Divisor: b}
}
return a / b, nil
}

func main() {
_, err := divide(10, 0)
if err != nil {
fmt.Println(err) // 输出：cannot divide 10 by 0
}
}

### fmt 包与错误格式化

`fmt` 包提供了对错误的格式化输出支持：

- `%v`：默认格式。
- `%+v`：如果支持，显示详细的错误信息。
- `%s`：作为字符串输出。

### 实例

package main

import (
"fmt"
)

// 定义一个 DivideError 结构
type DivideError struct {
dividee int
divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
strFormat := `
Cannot proceed, the divider is zero.
dividee: %d
divider: 0
`
return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
if varDivider == 0 {
dData := DivideError{
dividee: varDividee,
divider: varDivider,
}
errorMsg = dData.Error()
return
} else {
return varDividee / varDivider, ""
}

}

func main() {

// 正常情况
if result, errorMsg := Divide(100, 10); errorMsg == "" {
fmt.Println("100/10 = ", result)
}
// 当除数为零的时候会返回错误信息
if _, errorMsg := Divide(100, 0); errorMsg != "" {
fmt.Println("errorMsg is: ", errorMsg)
}

}

执行以上程序，输出结果为：

```

100/10 = 10
errorMsg is:
Cannot proceed, the divider is zero.
dividee: 100
divider: 0


```

### 使用 errors.Is 和 errors.As

从 Go 1.13 开始，`errors` 包引入了 `errors.Is` 和 `errors.As` 用于处理错误链：

#### `errors.Is`

检查某个错误是否是特定错误或由该错误包装而成。

### 实例

package main

import (
"errors"
"fmt"
)

var ErrNotFound = errors.New("not found")

func findItem(id int) error {
return fmt.Errorf("database error: %w", ErrNotFound)
}

func main() {
err := findItem(1)
if errors.Is(err, ErrNotFound) {
fmt.Println("Item not found")
} else {
fmt.Println("Other error:", err)
}
}

#### errors.As

将错误转换为特定类型以便进一步处理。

### 实例

package main

import (
"errors"
"fmt"
)

type MyError struct {
Code int
Msg string
}

func (e *MyError) Error() string {
return fmt.Sprintf("Code: %d, Msg: %s", e.Code, e.Msg)
}

func getError() error {
return &MyError{Code: 404, Msg: "Not Found"}
}

func main() {
err := getError()
var myErr *MyError
if errors.As(err, &myErr) {
fmt.Printf("Custom error - Code: %d, Msg: %s\n", myErr.Code, myErr.Msg)
}
}

### panic 和 recover

Go 的 panic 用于处理不可恢复的错误，recover 用于从 panic 中恢复。

panic:

- 导致程序崩溃并输出堆栈信息。
- 常用于程序无法继续运行的情况。

recover:

- 捕获 `panic`，避免程序崩溃。

### 实例

package main

import "fmt"

func safeFunction() {
defer func() {
if r := recover(); r != nil {
fmt.Println("Recovered from panic:", r)
}
}()
panic("something went wrong")
}

func main() {
fmt.Println("Starting program...")
safeFunction()
fmt.Println("Program continued after panic")
}

执行以上代码，输出结果为：

```
Starting program...
Recovered from panic: something went wrong
Program continued after panic
```

---

## Go 并发

Source: https://www.runoob.com/go/go-concurrent.html

## Go 并发

并发是指程序同时执行多个任务的能力。

Go 语言支持并发，通过 goroutines 和 channels 提供了一种简洁且高效的方式来实现并发。

Goroutines：

- Go 中的并发执行单位，类似于轻量级的线程。
- Goroutine 的调度由 Go 运行时管理，用户无需手动分配线程。
- 使用 `go` 关键字启动 Goroutine。
- Goroutine 是非阻塞的，可以高效地运行成千上万个 Goroutine。

Channel：

- Go 中用于在 Goroutine 之间通信的机制。
- 支持同步和数据共享，避免了显式的锁机制。
- 使用 `chan` 关键字创建，通过 `<-` 操作符发送和接收数据。

Scheduler（调度器）：

Go 的调度器基于 GMP 模型，调度器会将 Goroutine 分配到系统线程中执行，并通过 M 和 P 的配合高效管理并发。

- G：Goroutine。
- M：系统线程（Machine）。
- P：逻辑处理器（Processor）。

### Goroutine

goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

goroutine 语法格式：

```
go 函数名( 参数列表 )
```

例如：

```
go f(x, y, z)
```

开启一个新的 goroutine:

```
f(x, y, z)
```

Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。 同一个程序中的所有 goroutine 共享同一个地址空间。

### 实例

package main

import (
"fmt"
"time"
)

func sayHello() {
for i := 0; i < 5; i++ {
fmt.Println("Hello")
time.Sleep(100 * time.Millisecond)
}
}

func main() {
go sayHello() // 启动 Goroutine
for i := 0; i < 5; i++ {
fmt.Println("Main")
time.Sleep(100 * time.Millisecond)
}
}

执行以上代码，你会看到输出的 Main 和 Hello。输出是没有固定先后顺序，因为它们是两个 goroutine 在执行：

```
Main
Hello
Main
Hello
...

```

### 通道（Channel）

通道（Channel）是用于 Goroutine 之间的数据传递。

通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。

使用 `make` 函数创建一个 channel，使用 `<-` 操作符发送和接收数据。如果未指定方向，则为双向通道。

```
ch <- v // 把 v 发送到通道 ch
v := <-ch // 从 ch 接收数据
// 并把值赋给 v
```

声明一个通道很简单，我们使用chan关键字即可，通道在使用前必须先创建：

```
ch := make(chan int)
```

注意：默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。

以下实例通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和：

### 实例

package main

import "fmt"

func sum(s []int, c chan int) {
sum := 0
for _, v := range s {
sum += v
}
c <- sum // 把 sum 发送到通道 c
}

func main() {
s := []int{7, 2, 8, -9, 4, 0}

c := make(chan int)
go sum(s[:len(s)/2], c)
go sum(s[len(s)/2:], c)
x, y := <-c, <-c // 从通道 c 中接收

fmt.Println(x, y, x+y)
}

输出结果为：

```
-5 17 12
```

#### 通道缓冲区

通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：

```
ch := make(chan int, 100)
```

带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。

不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。

注意：如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。

### 实例

package main

import "fmt"

func main() {
// 这里我们定义了一个可以存储整数类型的带缓冲通道
// 缓冲区大小为2
ch := make(chan int, 2)

// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
// 而不用立刻需要去同步读取数据
ch <- 1
ch <- 2

// 获取这两个数据
fmt.Println(<-ch)
fmt.Println(<-ch)
}

执行输出结果为：

```
1
2
```

#### Go 遍历通道与关闭通道

Go 通过 range 关键字来实现遍历读取到的数据，类似于与数组或切片。格式如下：

```
v, ok := <-ch
```

如果通道接收不到数据后 ok 就为 false，这时通道就可以使用 close() 函数来关闭。

### 实例

package main

import (
"fmt"
)

func fibonacci(n int, c chan int) {
x, y := 0, 1
for i := 0; i < n; i++ {
c <- x
x, y = y, x+y
}
close(c)
}

func main() {
c := make(chan int, 10)
go fibonacci(cap(c), c)
// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
// 会结束，从而在接收第 11 个数据的时候就阻塞了。
for i := range c {
fmt.Println(i)
}
}

执行输出结果为：

```
0
1
1
2
3
5
8
13
21
34
```

#### Select 语句

`select` 语句使得一个 goroutine 可以等待多个通信操作。`select` 会阻塞，直到其中的某个 case 可以继续执行：

### 实例

package main

import "fmt"

func fibonacci(c, quit chan int) {
x, y := 0, 1
for {
select {
case c <- x:
x, y = y, x+y
case <-quit:
fmt.Println("quit")
return
}
}
}

func main() {
c := make(chan int)
quit := make(chan int)

go func() {
for i := 0; i < 10; i++ {
fmt.Println(<-c)
}
quit <- 0
}()
fibonacci(c, quit)
}

以上代码中，`fibonacci` goroutine 在 channel `c` 上发送斐波那契数列，当接收到 `quit` channel 的信号时退出。

执行输出结果为：

```
0
1
1
2
3
5
8
13
21
34
quit
```

### 使用 WaitGroup

sync.WaitGroup 用于等待多个 Goroutine 完成。

同步多个 Goroutine：

### 实例

package main

import (
"fmt"
"sync"
)

func worker(id int, wg *sync.WaitGroup) {
defer wg.Done() // Goroutine 完成时调用 Done()
fmt.Printf("Worker %d started\n", id)
fmt.Printf("Worker %d finished\n", id)
}

func main() {
var wg sync.WaitGroup

for i := 1; i <= 3; i++ {
wg.Add(1) // 增加计数器
go worker(i, &wg)
}

wg.Wait() // 等待所有 Goroutine 完成
fmt.Println("All workers done")
}

以上代码，执行输出结果如下：

```
Worker 1 started
Worker 1 finished
Worker 2 started
Worker 2 finished
Worker 3 started
Worker 3 finished
All workers done
```

### 高级特性

Buffered Channel：

创建有缓冲的 Channel。

```
ch := make(chan int, 2)
```

Context：

用于控制 Goroutine 的生命周期。

```
context.WithCancel、context.WithTimeout。
```

Mutex 和 RWMutex：

sync.Mutex 提供互斥锁，用于保护共享资源。

```
var mu sync.Mutex
mu.Lock()
// critical section
mu.Unlock()
```

### 并发编程小结

Go 语言通过 Goroutine 和 Channel 提供了强大的并发支持，简化了传统线程模型的复杂性。配合调度器和同步工具，可以轻松实现高性能并发程序。

- Goroutines 是轻量级线程，使用 `go` 关键字启动。
- Channels 用于 goroutines 之间的通信。
- Select 语句 用于等待多个 channel 操作。

#### 常见问题

死锁 (Deadlock)：

- 示例：所有 Goroutine 都在等待，但没有任何数据可用。
- 解决：避免无限等待、正确关闭通道。

数据竞争 (Data Race)：

- 示例：多个 Goroutine 同时访问同一变量。
- 解决：使用 Mutex 或 Channel 同步访问。

---

## Go 语言文件处理

Source: https://www.runoob.com/go/go-file-handle.html

## Go 语言文件处理

文件处理是 Go 语言中最常见的操作之一——读取配置、写入日志、数据持久化都离不开它。Go 的标准库提供了一套简洁而强大的文件 I/O 接口，覆盖从单次读写到流式处理的各种场景。

与文件处理相关的核心包有 5 个，各有分工： Go 文件处理核心包 os 核心：创建、读写、删除、权限 io 通用 Reader/Writer 接口 bufio 缓冲读写，性能优化 filepath 跨平台路径处理 ioutil ⚠ Go 1.16 已弃用 → os / io 典型使用流程 os.Open() bufio.Scanner 处理数据 defer Close()

版本提示：`ioutil` 包在 Go 1.16 已弃用，其功能已迁移到 `os`（如 `os.ReadFile`、`os.WriteFile`）和 `io`（如 `io.ReadAll`）。新代码应直接使用 `os` 和 `io` 包。

### 1. 文件创建

`os.Create` 创建一个新文件。如果文件已存在，会被截断（清空内容）。返回的文件对象必须关闭以释放系统资源：

### 实例

package main

import (
"log"
"os"
)

func main() {
// os.Create 创建文件（已存在则清空）
file, err := os.Create("test.txt")
if err != nil {
log.Fatal(err)
}
defer file.Close() // defer 确保函数结束时关闭文件

// 写入内容验证创建成功
file.WriteString("文件创建成功\n")
log.Println("文件创建成功")
}

`defer file.Close()` 是 Go 文件操作的最佳实践。`defer` 会确保即使后续代码出现 panic，文件也会被正确关闭，避免文件描述符泄漏。

### 2. 文件打开与关闭

`os` 包提供三种打开文件的方式，适用于不同场景：

函数 模式 说明 `os.Open(name)` 只读 最简单的方式，只读打开 `os.Create(name)` 读写 + 创建/截断 创建新文件或清空已有文件 `os.OpenFile(name, flag, perm)` 自定义 可指定读写、追加、创建等标志

### 实例

package main

import (
"fmt"
"os"
)

func main() {
// 方式1：只读打开
file, err := os.Open("example.txt")
if err != nil {
fmt.Println("打开失败:", err)
return
}
defer file.Close()
fmt.Println("文件打开成功")

// 方式2：OpenFile 自定义模式
// os.O_WRONLY 只写
// os.O_CREATE 不存在则创建
// os.O_APPEND 追加模式
// os.O_TRUNC 存在则截断
f, err := os.OpenFile("log.txt",
os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
if err != nil {
fmt.Println("打开失败:", err)
return
}
defer f.Close()
f.WriteString("追加一行日志\n")
}

#### OpenFile 标志位说明

标志 说明 `os.O_RDONLY` 只读（默认） `os.O_WRONLY` 只写 `os.O_RDWR` 读写 `os.O_APPEND` 追加模式，写入内容追加到文件末尾 `os.O_CREATE` 文件不存在时创建 `os.O_TRUNC` 打开时清空文件内容 `os.O_EXCL` 与 CREATE 配合，文件已存在时报错

### 3. 文件读取

Go 提供了三种主要的文件读取方式，根据文件大小和处理需求选择： 三种文件读取方式 一次性读取 os.ReadFile("file.txt") 适合小文件（配置、JSON等） 自动打开/关闭，最简洁 逐行读取 bufio.NewScanner(file) 适合大文件（日志、CSV等） 内存友好，逐行处理 流式读取 io.ReadAll(reader) 适合网络响应、已打开的文件 通用 Reader 接口

#### 3.1 一次性读取（小文件推荐）

`os.ReadFile` 是最简洁的方式——自动打开、读取、关闭文件，一步到位：

### 实例

package main

import (
"fmt"
"log"
"os"
)

func main() {
// 一次性读取整个文件（Go 1.16+）
data, err := os.ReadFile("config.json")
if err != nil {
log.Fatal(err)
}

fmt.Println(string(data))
}

#### 3.2 逐行读取（大文件推荐）

对于大文件，使用 `bufio.Scanner` 逐行处理，避免一次性加载到内存：

### 实例

package main

import (
"bufio"
"fmt"
"log"
"os"
)

func main() {
file, err := os.Open("large-file.log")
if err != nil {
log.Fatal(err)
}
defer file.Close()

// 逐行扫描
scanner := bufio.NewScanner(file)
lineNum := 0
for scanner.Scan() {
lineNum++
line := scanner.Text() // 获取当前行内容
fmt.Printf("第 %d 行: %s\n", lineNum, line)
}

// 检查扫描过程中是否出错
if err := scanner.Err(); err != nil {
log.Fatal("读取错误:", err)
}
}

注意：`bufio.Scanner` 默认最大行长度为 64KB。如果文件中有超长行，需要在调用 `Scan()` 前设置 `scanner.Buffer(make([]byte, 0), maxSize)` 来增大缓冲区。

#### 3.3 使用 io.ReadAll 读取

当你已经有一个打开的 `io.Reader`（如网络响应、已打开的文件），可以使用 `io.ReadAll`：

### 实例

package main

import (
"fmt"
"io"
"log"
"os"
)

func main() {
file, err := os.Open("example.txt")
if err != nil {
log.Fatal(err)
}
defer file.Close()

// 从 Reader 读取所有数据
data, err := io.ReadAll(file)
if err != nil {
log.Fatal(err)
}

fmt.Println(string(data))
}

### 4. 文件写入

Go 提供了多种写入方式，从简单的一次性写入到高性能的缓冲写入：

#### 4.1 一次性写入

`os.WriteFile` 将数据一次性写入文件（覆盖原有内容）：

### 实例

package main

import (
"log"
"os"
)

func main() {
content := []byte("Hello, Go!\n这是第二行\n")

// 0644: 所有者读写，其他用户只读
err := os.WriteFile("output.txt", content, 0644)
if err != nil {
log.Fatal(err)
}
log.Println("写入成功")
}

#### 4.2 使用 File 对象写入

通过文件对象写入，可以分多次写入内容：

### 实例

package main

import (
"fmt"
"log"
"os"
)

func main() {
file, err := os.Create("output.txt")
if err != nil {
log.Fatal(err)
}
defer file.Close()

// 方式1：写入字符串
file.WriteString("直接写入字符串\n")

// 方式2：写入字节切片
data := []byte("写入字节切片\n")
file.Write(data)

// 方式3：格式化写入
fmt.Fprintf(file, "格式化写入: %d + %d = %d\n", 3, 4, 7)
}

#### 4.3 缓冲写入（大量数据推荐）

`bufio.Writer` 会先将数据写入内存缓冲区，攒够后再批量写入磁盘，显著减少 I/O 次数。务必在结束前调用 `Flush()`：

### 实例

package main

import (
"bufio"
"log"
"os"
)

func main() {
file, err := os.Create("buffered-output.txt")
if err != nil {
log.Fatal(err)
}
defer file.Close()

// 创建带缓冲的写入器
writer := bufio.NewWriter(file)

// 写入多行（先进入缓冲区，不会立即写磁盘）
for i := 0; i < 1000; i++ {
writer.WriteString("这是第 " + itoa(i) + " 行\n")
}

// Flush 将缓冲区剩余数据写入文件
if err := writer.Flush(); err != nil {
log.Fatal("刷新缓冲区失败:", err)
}
log.Println("缓冲写入完成")
}

// 简单的 int 转 string 辅助函数
func itoa(n int) string {
return fmt.Sprintf("%d", n)
}

### 5. 文件追加写入

使用 `os.O_APPEND` 标志在文件末尾追加内容，而不覆盖已有数据：

### 实例

package main

import (
"log"
"os"
"time"
)

func main() {
// O_APPEND: 追加模式
// O_CREATE: 不存在则创建
// O_WRONLY: 只写
file, err := os.OpenFile("app.log",
os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
log.Fatal(err)
}
defer file.Close()

// 每次运行都会追加，不会覆盖
timestamp := time.Now().Format("2006-01-02 15:04:05")
file.WriteString("[" + timestamp + "] 应用启动\n")
}

对比：`os.Create` 会清空已有内容；`os.OpenFile` 配合 `O_APPEND` 会追加内容。日志场景务必使用追加模式。

### 6. 文件删除与重命名

### 实例

package main

import (
"log"
"os"
)

func main() {
// 重命名 / 移动文件
err := os.Rename("old.txt", "new.txt")
if err != nil {
log.Fatal("重命名失败:", err)
}
log.Println("重命名成功")

// 删除文件
err = os.Remove("temp.txt")
if err != nil {
log.Fatal("删除失败:", err)
}
log.Println("删除成功")
}

### 7. 文件信息与检查

#### 7.1 获取文件信息

### 实例

package main

import (
"fmt"
"log"
"os"
)

func main() {
info, err := os.Stat("test.txt")
if err != nil {
log.Fatal(err)
}

fmt.Println("文件名: ", info.Name())
fmt.Println("大小: ", info.Size(), "字节")
fmt.Println("权限: ", info.Mode())
fmt.Println("修改时间: ", info.ModTime())
fmt.Println("是目录吗: ", info.IsDir())
}

#### 7.2 检查文件是否存在

Go 没有专门的"文件存在"函数，通过 `os.Stat` + `os.IsNotExist` 判断：

### 实例

package main

import (
"fmt"
"os"
)

// fileExists 检查文件是否存在
func fileExists(path string) bool {
_, err := os.Stat(path)
return !os.IsNotExist(err)
}

func main() {
if fileExists("test.txt") {
fmt.Println("文件存在")
} else {
fmt.Println("文件不存在")
}
}

### 8. 目录操作

### 实例

package main

import (
"fmt"
"log"
"os"
)

func main() {
// 创建单个目录
err := os.Mkdir("newdir", 0755)
if err != nil {
log.Fatal(err)
}

// 递归创建多级目录（父目录不存在也会自动创建）
err = os.MkdirAll("path/to/deep/dir", 0755)
if err != nil {
log.Fatal(err)
}

// 读取目录内容
entries, err := os.ReadDir(".")
if err != nil {
log.Fatal(err)
}
for _, entry := range entries {
mark := "&#x1f4c4;"
if entry.IsDir() {
mark = "&#x1f4c1;"
}
fmt.Printf("%s %s\n", mark, entry.Name())
}

// 删除空目录
os.Remove("newdir")

// 递归删除目录及其所有内容（&#x26a0; 谨慎使用）
os.RemoveAll("path")
}

### 9. 高级操作

#### 9.1 文件复制

使用 `io.Copy` 在两个文件之间高效复制数据：

### 实例

package main

import (
"io"
"log"
"os"
)

// copyFile 复制文件的通用函数
func copyFile(src, dst string) (int64, error) {
srcFile, err := os.Open(src)
if err != nil {
return 0, err
}
defer srcFile.Close()

dstFile, err := os.Create(dst)
if err != nil {
return 0, err
}
defer dstFile.Close()

// io.Copy 自动管理缓冲区，高效复制
return io.Copy(dstFile, srcFile)
}

func main() {
n, err := copyFile("source.txt", "destination.txt")
if err != nil {
log.Fatal(err)
}
log.Printf("复制完成，共 %d 字节", n)
}

#### 9.2 临时文件与目录

### 实例

package main

import (
"fmt"
"log"
"os"
)

func main() {
// 创建临时文件（自动命名，前缀为 "app-"）
tmpFile, err := os.CreateTemp("", "app-*.txt")
if err != nil {
log.Fatal(err)
}
defer os.Remove(tmpFile.Name()) // 用完后清理
fmt.Println("临时文件:", tmpFile.Name())

// 创建临时目录
tmpDir, err := os.MkdirTemp("", "app-*")
if err != nil {
log.Fatal(err)
}
defer os.RemoveAll(tmpDir) // 用完后清理
fmt.Println("临时目录:", tmpDir)
}

#### 9.3 递归遍历目录

`filepath.Walk` 可以递归遍历目录树中的所有文件和子目录：

### 实例

package main

import (
"fmt"
"log"
"os"
"path/filepath"
)

func main() {
root := "." // 从当前目录开始遍历

err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
if err != nil {
return err // 遇到错误时可以选择跳过或返回
}

// 打印路径和大小
size := ""
if !info.IsDir() {
size = fmt.Sprintf(" (%d bytes)", info.Size())
}
fmt.Println(path + size)
return nil
})

if err != nil {
log.Fatal(err)
}
}

#### 9.4 跨平台路径拼接

永远不要用字符串拼接路径（如 `dir + "/" + file`），使用 `filepath.Join` 确保跨平台兼容：

```
// 错误：Windows 上路径分隔符是 \
path := "dir" + "/" + "file.txt"

// 正确：自动适配当前操作系统
path := filepath.Join("dir", "file.txt")

```

### 场景速查表

场景 推荐方式 关键代码 读取小文件 `os.ReadFile` `data, err := os.ReadFile("file.txt")` 逐行读取大文件 `bufio.Scanner` `scanner := bufio.NewScanner(file)` 写入小文件 `os.WriteFile` `os.WriteFile("f.txt", data, 0644)` 大量数据写入 `bufio.Writer` `writer := bufio.NewWriter(file)` 追加日志 `os.OpenFile` + `O_APPEND` `os.OpenFile("log", os.O_APPEND|os.O_WRONLY, 0644)` 文件复制 `io.Copy` `io.Copy(dst, src)` 遍历目录 `filepath.Walk` `filepath.Walk(".", callback)` 路径拼接 `filepath.Join` `filepath.Join("dir", "file.txt")` 检查文件存在 `os.Stat` + `os.IsNotExist` `_, err := os.Stat(path)` 创建临时文件 `os.CreateTemp` `os.CreateTemp("", "prefix-*.txt")`

文件权限说明：`0644` 表示所有者可读写（6=4+2），组和其他用户只读（4）。目录通常使用 `0755`（多了执行权限，否则无法进入目录）。

---

## Go 语言正则表达式

Source: https://www.runoob.com/go/go-regex.html

## Go 语言正则表达式

正则表达式（Regular Expression，简称 regex 或 regexp）是一种用于匹配字符串的强大工具。

正则表达式通过定义一种模式（pattern），可以快速搜索、替换或提取符合该模式的字符串，详细可以参见正则表达式教程。

在 Go 语言中，正则表达式通过 `regexp` 包来实现。

### Go 语言中的 `regexp` 包

Go 语言的标准库提供了 `regexp` 包，用于处理正则表达式。以下是 `regexp` 包中常用的函数和方法：

- `Compile` 和 `MustCompile`
用于编译正则表达式。`Compile` 返回一个 `*Regexp` 对象和一个错误，而 `MustCompile` 在编译失败时会直接 panic。
- `MatchString`
检查字符串是否匹配正则表达式。
- `FindString` 和 `FindAllString`
用于查找匹配的字符串。`FindString` 返回第一个匹配项，`FindAllString` 返回所有匹配项。
- `ReplaceAllString`
用于替换匹配的字符串。
- `Split`
根据正则表达式分割字符串。

### 正则表达式的基本语法

以下是一些常用的正则表达式语法：

- `.`：匹配任意单个字符（除了换行符）。
- `*`：匹配前面的字符 0 次或多次。
- `+`：匹配前面的字符 1 次或多次。
- `?`：匹配前面的字符 0 次或 1 次。
- `\d`：匹配数字字符（等价于 `[0-9]`）。
- `\w`：匹配字母、数字或下划线（等价于 `[a-zA-Z0-9_]`）。
- `\s`：匹配空白字符（包括空格、制表符、换行符等）。
- `[]`：匹配括号内的任意一个字符（例如 `[abc]` 匹配 `a`、`b` 或 `c`）。
- `^`：匹配字符串的开头。
- `$`：匹配字符串的结尾。

### 示例代码

以下是一些使用 Go 语言正则表达式的示例：

#### 示例 1：检查字符串是否匹配正则表达式

### 实例

package main

import (
"fmt"
"regexp"
)

func main() {
pattern := `^[a-zA-Z0-9]+$`
regex := regexp.MustCompile(pattern)

str := "Hello123"
if regex.MatchString(str) {
fmt.Println("字符串匹配正则表达式")
} else {
fmt.Println("字符串不匹配正则表达式")
}
}

#### 示例 2：查找匹配的字符串

### 实例

package main

import (
"fmt"
"regexp"
)

func main() {
pattern := `\d+`
regex := regexp.MustCompile(pattern)

str := "我有 3 个苹果和 5 个香蕉"
matches := regex.FindAllString(str, -1)
fmt.Println("找到的数字：", matches)
}

#### 示例 3：替换匹配的字符串

### 实例

package main

import (
"fmt"
"regexp"
)

func main() {
pattern := `\s+`
regex := regexp.MustCompile(pattern)

str := "Hello World"
result := regex.ReplaceAllString(str, " ")
fmt.Println("替换后的字符串：", result)
}

#### 示例 4：分割字符串

### 实例

package main

import (
"fmt"
"regexp"
)

func main() {
pattern := `,`
regex := regexp.MustCompile(pattern)

str := "apple,banana,orange"
parts := regex.Split(str, -1)
fmt.Println("分割后的字符串：", parts)
}

### 注意事项

- 性能问题
正则表达式的匹配和替换操作可能会消耗较多资源，尤其是在处理大量数据时。建议在性能敏感的场景下谨慎使用。
- 转义字符
在 Go 语言中，正则表达式中的反斜杠 `\` 需要写成 `\\`，因为反斜杠在字符串中也是转义字符。
- 错误处理
使用 `Compile` 函数时，务必检查返回的错误，以避免程序崩溃。

---

## Go 类型断言

Source: https://www.runoob.com/go/go-type-assertion.html

## Go 类型断言

在 Go 语言中，类型断言（Type Assertion）是一种用于检查接口值的实际类型的机制。

类型断言是 Go 语言中处理接口类型的重要工具，它允许我们从接口值中提取出具体的类型，并对其进行操作。

类型断言通常用于处理接口类型的变量，因为接口变量可以存储任何实现了该接口的具体类型的值。

#### 基本语法

类型断言的基本语法如下：

```

value, ok := interfaceValue.(Type)
```

- `interfaceValue` 是一个接口类型的变量。
- `Type` 是你想要断言的类型。
- `value` 是断言成功后的具体类型的值。
- `ok` 是一个布尔值，表示断言是否成功。

如果断言成功，`value` 将是 `interfaceValue` 的实际值，`ok` 为 `true`；如果断言失败，`value` 将是 `Type` 的零值，`ok` 为 `false`。

### 实例

package main

import "fmt"

func main() {
var i interface{} = "Hello, Go!"

// 尝试将 i 断言为 string 类型
s, ok := i.(string)
if ok {
fmt.Println("断言成功:", s)
} else {
fmt.Println("断言失败")
}

// 尝试将 i 断言为 int 类型
n, ok := i.(int)
if ok {
fmt.Println("断言成功:", n)
} else {
fmt.Println("断言失败")
}
}

#### 输出结果

```
断言成功: Hello, Go!
断言失败

```

### 类型断言的另一种形式

除了上述的 `value, ok := interfaceValue.(Type)` 形式，Go 还支持另一种形式的类型断言，它不返回布尔值，而是直接在断言失败时引发 panic。

这种形式的语法如下：

```

value := interfaceValue.(Type)
```

#### 示例代码

### 实例

package main

import "fmt"

func main() {
var i interface{} = "Hello, Go!"

// 直接断言为 string 类型
s := i.(string)
fmt.Println("断言成功:", s)

// 直接断言为 int 类型（会引发 panic）
n := i.(int)
fmt.Println("断言成功:", n)
}

#### 输出结果

```
断言成功: Hello, Go!
panic: interface conversion: interface {} is string, not int

```

### 类型断言的常见用途

#### 1. 处理多种类型的接口值

Go 还提供了特殊的 type switch 语法来测试多种类型：

```
switch v := i.(type) {
case T1:
// v的类型是T1
case T2:
// v的类型是T2
default:
// 默认情况
}
```

当接口变量可能存储多种类型的值时，类型断言可以帮助我们根据实际类型执行不同的操作。

### 实例

func printType(i interface{}) {
switch v := i.(type) {
case int:
fmt.Println("这是一个整数:", v)
case string:
fmt.Println("这是一个字符串:", v)
default:
fmt.Println("未知类型")
}
}

#### 2. 从接口中提取具体类型

在处理接口类型的变量时，我们可能需要将其转换为具体的类型以便进行进一步的操作。

### 实例

func processInterface(i interface{}) {
if s, ok := i.(string); ok {
fmt.Println("处理字符串:", s)
} else if n, ok := i.(int); ok {
fmt.Println("处理整数:", n)
} else {
fmt.Println("无法处理的类型")
}
}

### 注意事项

- 类型断言只能用于接口类型：类型断言只能用于接口类型的变量，不能用于非接口类型的变量。
- 避免 panic：在使用不返回布尔值的类型断言时，务必确保类型断言不会失败，否则会引发 panic。
- 类型断言的性能：类型断言在运行时进行类型检查，因此可能会带来一定的性能开销。在性能敏感的场景中，应谨慎使用。

---

## Go 继承

Source: https://www.runoob.com/go/go-inheritance.html

## Go 继承

在面向对象编程（OOP）中，继承是一种机制，允许一个类（子类）从另一个类（父类）继承属性和方法。通过继承，子类可以复用父类的代码，并且可以在不修改父类的情况下扩展或修改其行为。

Go 语言并不是一种传统的面向对象编程语言，它没有类和继承的概念。

Go 使用结构体（struct）和接口（interface）来实现类似的功能。

### Go 中的 "继承"

Go 语言没有传统面向对象语言中的类(class)和继承(inheritance)概念，而是通过组合(composition)和接口(interface)来实现类似的功能。

#### 1. 组合（Composition）

组合是 Go 中实现代码复用的主要方式。通过将一个结构体嵌入到另一个结构体中，子结构体可以"继承"父结构体的字段和方法。

### 实例

package main

import "fmt"

// 父结构体
type Animal struct {
Name string
}

// 父结构体的方法
func (a *Animal) Speak() {
fmt.Println(a.Name, "says hello!")
}

// 子结构体
type Dog struct {
Animal // 嵌入 Animal 结构体
Breed string
}

func main() {
dog := Dog{
Animal: Animal{Name: "Buddy"},
Breed: "Golden Retriever",
}

dog.Speak() // 调用父结构体的方法
fmt.Println("Breed:", dog.Breed)
}

##### 代码解释

- `Animal` 是父结构体，包含一个字段 `Name` 和一个方法 `Speak`。
- `Dog` 是子结构体，通过嵌入 `Animal` 结构体，继承了 `Animal` 的字段和方法。
- 在 `main` 函数中，我们创建了一个 `Dog` 实例，并调用了 `Speak` 方法。

#### 2. 接口（Interface）

接口是 Go 中实现多态的主要方式。通过定义接口，不同的结构体可以实现相同的方法，从而实现类似继承的多态行为。

##### 示例代码

### 实例

package main

import "fmt"

// 定义接口
type Speaker interface {
Speak()
}

// 父结构体
type Animal struct {
Name string
}

// 实现接口方法
func (a *Animal) Speak() {
fmt.Println(a.Name, "says hello!")
}

// 子结构体
type Dog struct {
Animal
Breed string
}

func main() {
var speaker Speaker

dog := Dog{
Animal: Animal{Name: "Buddy"},
Breed: "Golden Retriever",
}

speaker = &dog
speaker.Speak() // 通过接口调用方法
}

##### 代码解释

- `Speaker` 是一个接口，定义了一个 `Speak` 方法。
- `Animal` 结构体实现了 `Speaker` 接口。
- `Dog` 结构体通过嵌入 `Animal` 结构体，间接实现了 `Speaker` 接口。
- 在 `main` 函数中，我们将 `Dog` 实例赋值给 `Speaker` 接口，并通过接口调用 `Speak` 方法。

#### Go 与经典继承的区别

特性经典继承Go 的方式代码复用通过继承通过组合(嵌入结构体)多态通过继承和方法重写通过接口实现关系"是一个"(is-a)关系"有一个"(has-a)或"实现了"关系灵活性继承关系固定可以运行时组合

完整继承模拟：

### 实例

package main

import "fmt"

// 基类
type Vehicle struct {
Brand string
}

func (v *Vehicle) Start() {
fmt.Println(v.Brand, "started")
}

// 派生类
type Car struct {
Vehicle // 嵌入Vehicle
Model string
}

// 重写Start方法
func (c *Car) Start() {
fmt.Println(c.Brand, c.Model, "car started")
}

func main() {
v := Vehicle{Brand: "Toyota"}
c := Car{
Vehicle: Vehicle{Brand: "Honda"},
Model: "Civic",
}

v.Start() // Toyota started
c.Start() // Honda Civic car started
c.Vehicle.Start() // Honda started
}

Go 的这种设计避免了传统继承的许多问题，如脆弱的基类问题，同时提供了更大的灵活性。

---

## Go 语言开发工具

Source: https://www.runoob.com/go/go-ide.html

## Go 语言开发工具

### VSCode

VScode 安装教程参见：https://www.runoob.com/w3cnote/vscode-tutorial.html

然后我们打开 VSCode 的扩展（Ctrl+Shift+P）：

搜索 go：

点击安装，安装完成后我们就可以使用代码提示、测试、调试等功能了。

### GoLand

GoLand 是 Jetbrains 家族的 Go 语言 IDE，有 30 天的免费试用期。

安装也很简单访问 Goland 的下载页面，根据你当期的系统环境三大平台（Mac、Linux、Windows）下载对应的软件。

### LiteIDE

LiteIDE 是一款开源、跨平台的轻量级 Go 语言集成开发环境（IDE）。

#### 支持的 操作系统

- Windows x86 (32-bit or 64-bit)
- Linux x86 (32-bit or 64-bit)

下载地址 ：http://sourceforge.net/projects/liteide/files/

源码地址 ：https://github.com/visualfc/liteide

### Eclipse

Eclipse 也是非常常用的开发利器，以下介绍如何使用 Eclipse 来编写 Go 程序。

Eclipse 编辑 Go 的主界面

- 首先下载并安装好 Eclipse
- 下载 goclipse 插件 https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#installation
- 下载 gocode，用于 go 的代码补全提示

gocode 的 github 地址：

```
https://github.com/nsf/gocode
```

在 Windows下要安装 git，通常用 msysgit。

再在 cmd 下安装：

```
go get -u github.com/nsf/gocode
```

也可以下载代码，直接用 go build 来编译，会生成 gocode.exe
- 下载 MinGW 并按要求装好
- 配置插件

Windows->Reference->Go

(1)、配置 Go 的编译器

设置 Go 的一些基础信息

(2)、配置 Gocode（可选，代码补全），设置 Gocode 路径为之前生成的 gocode.exe 文件

设置 gocode 信息

(3)、配置 GDB（可选，做调试用），设置 GDB 路径为 MingW 安装目录下的 gdb.exe 文件

设置 GDB 信息
- 测试是否成功

新建一个 go 工程，再建立一个 hello.go。如下图：

新建项目编辑文件

调试如下（要在 console 中用输入命令来调试）：

图 1.16 调试 Go 程序

---

## Go 语言测验知识挑战

Source: https://www.runoob.com/go/go-quiz.html

## Go 语言测验

## 知识挑战

通过这些互动问题测试学习情况

### 测验完成！

0/0

重新开始

上一题 下一题

### 更多测验

1. Go 测验 1 2. Go 测验 2 3. Go 测验 3 4. Go 测验 4 5. Go 测验 5 6. Go 测验 6 7. Go 测验 7 8. Go 测验 8 9. Go 测验 9 10. Go 测验 10

---

## Go Modules

Source: https://www.runoob.com/go/go-modules.html

## Go Modules

Go Modules 是 Go 语言的官方依赖管理工具，自 Go 1.11 版本开始引入，在 Go 1.16 版本成为默认的依赖管理模式。

Go Modules 解决了 Go 语言长期以来在依赖管理方面的痛点，为开发者提供了版本控制、依赖隔离和可重复构建等核心功能。

Go Modules 是一组相关 Go 包的集合，它们被版本化并作为一个独立的单元进行管理。每个模块都有一个明确的版本标识，允许开发者在项目中精确指定所需依赖的版本。

#### 核心概念解析

模块（Module）：包含 `go.mod` 文件的目录树，该文件定义了模块的路径、Go 版本要求和依赖关系。

版本（Version）：遵循语义化版本控制（Semantic Versioning）的标识符，格式为 `vMAJOR.MINOR.PATCH`。

依赖图（Dependency Graph）：模块及其所有传递依赖的层次结构，Go 工具会自动解析和维护。

### 为什么需要 Go Modules？

#### 传统 GOPATH 的问题

在 Go Modules 出现之前，Go 使用 GOPATH 模式，存在以下局限性：

- 工作空间限制：所有项目必须放在 GOPATH 目录下
- 版本管理困难：无法精确控制依赖版本
- 依赖冲突：多个项目可能使用同一依赖的不同版本
- 可重复构建挑战：难以确保不同环境下的构建一致性

#### Go Modules 的优势

传统 GOPATH vs Go Modules 对比：

特性 GOPATH 模式 Go Modules 项目位置限制 必须放在 GOPATH 下 任意位置均可 版本控制 有限支持 完整的语义化版本控制 依赖隔离 全局共享 项目级隔离 可重复构建 困难 自动保障 离线工作 不支持 支持本地缓存

### 核心文件解析

#### go.mod 文件

`go.mod` 是模块的定义文件，包含以下主要部分：

### 实例

module example.com/mymodule // 模块路径

go 1.21 // Go 版本要求

require (
github.com/gin-gonic/gin v1.9.1
golang.org/x/text v0.12.0
)

replace golang.org/x/text => ../local/text // 本地替换

exclude github.com/old/module v1.0.0 // 排除特定版本

#### go.sum 文件

`go.sum` 文件记录依赖模块的加密哈希值，用于验证模块内容的完整性：

### 实例

github.com/bytedance/sonic v1.9.1 h1:ei0tVql02GmiYGRCTUcI6g...
github.com/bytedance/sonic v1.9.1/go.mod h1:iZcSUejdk5C4OW...

### 基本命令详解

#### 模块初始化

### 实例

# 创建新模块
go mod init example.com/myproject

# 在现有项目中初始化
cd /path/to/project
go mod init

#### 依赖管理

### 实例

# 添加依赖（自动选择最新版本）
go get github.com/gin-gonic/gin

# 添加特定版本
go get github.com/gin-gonic/gin@v1.9.1

# 更新到最新版本
go get -u github.com/gin-gonic/gin

# 更新所有依赖
go get -u all

# 下载依赖到本地缓存
go mod download

# 整理 go.mod 文件
go mod tidy

#### 依赖查询

### 实例

# 查看所有依赖
go list -m all

# 查看特定依赖的可用版本
go list -m -versions github.com/gin-gonic/gin

# 查看为什么需要某个依赖
go mod why github.com/gin-gonic/gin

### 实际工作流程

#### 1. 新项目初始化

### 实例

# 创建项目目录
mkdir myproject && cd myproject

# 初始化模块
go mod init github.com/username/myproject

# 编写代码并导入依赖
# 然后运行以下命令自动处理依赖
go mod tidy

#### 2. 依赖版本控制策略

### 实例

// go.mod 中的版本指定方式
require (
github.com/lib/pq v1.10.9 // 精确版本
golang.org/x/text v0.3.7 // 精确版本
github.com/stretchr/testify v1.8.0 // 测试依赖
)

// 间接依赖由 Go 工具自动管理

#### 3. 版本选择机制

Go Modules 使用最小版本选择（MVS）算法：

### 高级特性

#### 版本替换（Replace）

```
// 用本地路径替换远程依赖
replace github.com/some/dependency => ../local/dependency

// 用不同版本替换
replace github.com/some/dependency => github.com/some/dependency v2.0.0

// 用 fork 仓库替换
replace github.com/some/dependency => github.com/myfork/dependency v1.0.0
```

#### 排除特定版本

```
exclude (
github.com/problematic/module v1.0.0
github.com/another/badmodule v2.1.0
)
```

#### 私有仓库支持

```

# 配置私有仓库认证
git config --global url."https://user:token@github.com".insteadOf "https://github.com"

# 或使用环境变量
export GOPRIVATE=github.com/mycompany/*
```

### 最佳实践

#### 1. 版本管理策略

```
# 开发阶段使用最新版本
go get -u ./...

# 发布前锁定版本
go mod tidy
go mod vendor # 可选：创建vendor目录

# 定期更新依赖
go get -u all
go mod tidy
```

#### 2. 协作开发规范

```
# 提交前确保 go.mod 和 go.sum 一致
go mod tidy
go mod verify

# 检查未使用的依赖
go mod tidy -v
```

#### 3. CI/CD 集成

```
# GitHub Actions 示例
jobs:
test:
runs-on: ubuntu-latest
steps:
- uses: actions/checkout@v3
- uses: actions/setup-go@v3
with:
go-version: '1.21'
- run: go mod download
- run: go test ./...
```

### 常见问题与解决方案

#### 问题 1: 依赖下载失败

解决方案：

```
# 设置代理
go env -w GOPROXY=https://goproxy.cn,direct

# 清理缓存并重试
go clean -modcache
go mod download
```

#### 问题 2: 版本冲突

解决方案：

```
# 查看依赖图
go mod graph

# 分析冲突原因
go mod why -m conflicting/package

# 使用 replace 指令解决
```

#### 问题 3: 私有模块认证

解决方案：

```

# 配置 netrc 文件
machine github.com
login username
password token

# 或使用 SSH 替代 HTTPS
git config --global url."git@github.com:".insteadOf "https://github.com/"
```

### 实践练习

#### 练习 1: 创建第一个模块

1、创建新目录并初始化模块：

```
mkdir hello-world && cd hello-world
go mod init example.com/hello
```

2、创建 `main.go`：

### 实例

package main

import (
"fmt"
"rsc.io/quote"
)

func main() {
fmt.Println(quote.Hello())
}

3、运行并观察依赖管理：

```

go run main.go
go mod tidy
cat go.mod
```

#### 练习 2: 版本控制实践

1、添加特定版本的依赖：

```

go get golang.org/x/text@v0.3.7
```

2、尝试更新到最新版本：

```

go get -u golang.org/x/text
```

3、查看版本变化：

```

go list -m all | grep text
```
