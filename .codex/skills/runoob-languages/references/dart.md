# Dart - 菜鸟教程

Tutorial: https://www.runoob.com/dart/dart-tutorial.html

---

## Dart 教程

Source: https://www.runoob.com/dart/dart-tutorial.html

## Dart 教程

Dart 是 Google 开发的一门现代化编程语言，主要用于构建高性能、跨平台的应用程序。

Dart 支持面向对象、函数式编程等多种编程范式，可以运行在浏览器、服务器、移动端以及桌面平台。

Dart 是 Flutter 框架的核心语言，通过 Dart 可以快速开发 Android、iOS、Web 和桌面应用。

本教程基于 Dart 3.x 版本测试，适合零基础开发者学习 Dart 编程。

### 学习 Dart 前，您需要了解的知识：

- 基本编程概念
- 变量和数据类型
- 条件判断和循环
- 面向对象编程基础

如果您熟悉 JavaScript、Java、C# 或 Python 等语言，学习 Dart 会更加容易。

### Dart 应用领域

Dart 采用现代语言设计理念，拥有简洁的语法、强大的类型系统以及高效的运行环境。

Dart 的主要应用领域：

- 移动应用开发：通过 Flutter 构建 Android 和 iOS 应用。
- Web 开发：Dart 可以编译为 JavaScript，在浏览器中运行。
- 服务器开发：使用 Dart 构建后端服务和 API。
- 桌面应用：支持 Windows、macOS 和 Linux 应用开发。
- 跨平台开发：一套代码运行在多个平台。

### 第一个实例

### Dart Hello World 实例

void main() {
print("Hello Dart!");
}

### Dart 的特点和优势

- 简单易学： Dart 语法简洁清晰，和 Java、JavaScript 等语言有很多相似之处，学习成本较低。
- 面向对象： Dart 是纯面向对象语言，所有内容都是对象，支持类、继承、接口和混入等特性。
- 强类型系统： Dart 提供静态类型检查，可以帮助开发者减少程序错误，提高代码质量。
- 高性能： Dart 支持 JIT（即时编译）和 AOT（提前编译），可以提供快速的开发体验和运行性能。
- 跨平台开发： Dart 配合 Flutter 可以使用同一套代码开发移动端、Web 和桌面应用。
- 现代化语言特性： Dart 支持异步编程、空安全、扩展方法、模式匹配等现代语言功能。
- 丰富的开发工具： Dart 提供完善的开发环境支持，包括 Dart SDK、调试工具和代码分析工具。
- 强大的 Flutter 生态： Dart 是 Flutter 的官方开发语言，拥有大量 UI 组件和第三方库。
- 空安全机制： Dart 的 Null Safety 可以减少空值导致的程序崩溃问题。
- 开源社区支持： Dart 拥有活跃的开发社区，并持续进行语言更新和优化。

### Dart 与其他语言比较

- Dart vs JavaScript： Dart 类型系统更严格，适合大型项目开发。
- Dart vs Java： Dart 语法更加现代化，开发效率更高。
- Dart vs Python： Dart 性能更接近编译型语言，适合构建应用程序。

### Flutter 与 Dart

Flutter 是 Google 推出的跨平台 UI 开发框架，而 Dart 是 Flutter 使用的开发语言。

通过 Flutter + Dart，开发者可以使用一套代码快速构建多个平台的应用。

- Android 应用
- iOS 应用
- Web 应用
- Windows 应用
- macOS 应用
- Linux 应用

### 参考资料：

官方网站： https://dart.dev/

中文文档： https://dart.cn/

Flutter 官方网站： https://flutter.dev/

Flutter 教程： https://www.runoob.com/flutter/flutter-tutorial.html

---

## Dart 简介

Source: https://www.runoob.com/dart/dart-intro.html

## Dart 简介

Dart 是由 Google 开发的一种现代化编程语言，专为构建高性能、跨平台应用而设计。

本章带你了解 Dart 的起源、设计理念以及它在前端、移动端和服务端的广泛应用场景。

### Dart 的历史与设计目标

Dart 诞生于 2011 年，最初的目标是替代 JavaScript 成为浏览器的原生脚本语言。

虽然这个目标未能实现，但 Dart 在 2017 年迎来了关键转折点——Flutter 框架的发布让 Dart 重新进入开发者视野。

Dart 的设计目标非常明确：

设计目标说明 高性能支持 AOT（提前编译）和 JIT（即时编译），既能快速开发又能生成原生代码 强类型支持类型推断，兼顾类型安全与开发效率 跨平台一套代码可编译为 ARM、x64 机器码或 JavaScript 易学习语法借鉴 Java、JavaScript、C# 等主流语言，降低学习成本 空安全内置 Null Safety 机制，从语言层面避免空指针异常

Dart 2.0 在 2018 年发布，标志着语言走向成熟。Dart 3.0 在 2023 年发布，进一步加强了类型系统和模式匹配等特性。

### Dart 的应用场景

Dart 的应用场景非常广泛，覆盖了从移动端到服务端的全栈开发需求。

#### 移动端开发（Flutter）

这是 Dart 目前最核心的应用场景。

通过 Flutter 框架，Dart 可以构建 iOS 和 Android 双平台的原生应用，一套代码同时运行。

#### Web 开发

Dart 可以编译为 JavaScript，用于构建 Web 前端应用。

AngularDart 是 Google 维护的 Dart 版 Angular 框架，适合大型 Web 项目。

#### 服务端开发

通过 dart:io 库和 Shelf 等框架，Dart 也能编写高性能的服务端应用。

Google 内部有多个服务端项目使用 Dart 编写。

#### 命令行工具

Dart 可以编译为独立的可执行文件，非常适合编写跨平台的 CLI 工具。

应用场景编译目标典型框架/工具 移动端ARM 机器码Flutter Web 前端JavaScriptAngularDart 服务端机器码 / JITShelf、Dart Frog 命令行机器码dart compile exe

### Dart 与 JavaScript 对比

Dart 和 JavaScript 都是面向对象的语言，但 Dart 在语言设计上更加严谨。

下面从几个维度对比两者的差异：

对比维度DartJavaScript 类型系统强类型，支持类型推断弱类型，动态类型 空安全内置 Null Safety无（需 TypeScript 或手动检查） 入口函数必须有 main() 函数脚本直接执行 类与继承基于类的 OOP，单继承基于原型链 编译方式AOT + JIT 双模式JIT 为主 并发模型Isolate（消息传递）Worker 线程 / 事件循环 包管理pub.dev（统一官方仓库）npm（社区仓库）

如果你有 JavaScript 基础，学习 Dart 会非常轻松——它们的异步语法（async/await）几乎完全相同。

但 Dart 的类型系统和空安全机制会让你在大型项目中更有安全感。

### 实例

同样的功能在两种语言中的写法对比：

// JavaScript 版本
function greet(name) {
if (name === null || name === undefined) {
return 'Hello, stranger!';
}
return `Hello, ${name}!`;
}

console.log(greet('runoob'));
console.log(greet(null));

// Dart 版本——空安全让类型更明确
String greet(String? name) {
// name 可能为 null，需要显式处理
if (name == null) {
return 'Hello, stranger!';
}
return 'Hello, $name!';
}

void main() {
print(greet('runoob'));
print(greet(null));
}

```

Hello, runoob!
Hello, stranger!

```

### Dart 与 Flutter 的关系

很多人第一次接触 Dart 是因为 Flutter。

实际上，Flutter 是 Dart 的"杀手级应用"，但 Dart 本身是一门独立的通用编程语言。

Flutter 选择 Dart 的原因有以下几点：

- JIT 编译支持热重载（Hot Reload），修改代码后亚秒级看到效果，极大提升开发效率
- AOT 编译生成原生机器码，应用启动快、运行流畅，无 JavaScript Bridge 的性能损耗
- Dart 的声明式 UI与 Flutter 的 Widget 树结构天然契合
- 统一技术栈：Google 同时维护语言和框架，版本兼容性好

学习 Flutter 之前，建议先掌握 Dart 基础。理解语言特性后，Flutter 的学习曲线会平滑很多。本教程就是为这个目标准备的。

### Dart 版本发展简史

版本发布时间重要特性 Dart 1.02013 年首个稳定版，基本语言特性 Dart 2.02018 年强类型系统，new/const 可选 Dart 2.122021 年引入 Null Safety Dart 3.02023 年模式匹配、Records、类修饰符

本教程基于 Dart 3.x 版本编写，所有代码示例都在 Dart 3.0 及以上版本中验证通过。

---

## Dart 开发环境搭建

Source: https://www.runoob.com/dart/dart-install.html

## Dart 开发环境搭建

学习任何编程语言的第一步，都是在本地搭建一个可用的开发环境。

本章将带你安装 Dart SDK、配置编辑器，并运行你的第一个 Dart 程序。

#### 在线体验

如果你只是想快速体验 Dart，不想在本地安装任何东西，DartPad 是最省事的选择。

DartPad 是 Dart 官方提供的在线编辑器，无需注册，打开浏览器即可编写和运行代码。

访问地址：https://dartpad.dev

#### DartPad 的主要功能

- 在线编写、运行 Dart 代码，实时查看输出
- 内置多个示例程序，适合快速上手
- 支持分享代码链接，方便协作与提问
- 支持 Flutter 代码预览（可以在浏览器里运行简单的 Flutter 界面）

### 安装 Dart SDK

Dart SDK 包含了编译、运行、调试 Dart 程序所需的所有工具。

根据你的操作系统，选择对应的安装方式。

#### macOS 安装

如果你使用 macOS，推荐通过 Homebrew 安装：

```

$ brew tap dart-lang/dart
$ brew install dart

```

安装完成后，验证版本：

```

$ dart --version
Dart SDK version: 3.5.0 (stable)

```

#### Windows 安装

Windows 用户可以通过 Chocolatey 包管理器安装：

```

C:\> choco install dart-sdk

```

或者直接访问 dart.dev/get-dart 下载安装包，双击运行安装向导。

#### Linux 安装

Debian/Ubuntu 系统可以通过 apt 安装：

```

$ sudo apt-get update
$ sudo apt-get install apt-transport-https
$ wget -qO- https://dl-ssl.google.com/linux/linux_signing_key.pub | sudo gpg --dearmor -o /usr/share/keyrings/dart.gpg
$ echo 'deb [signed-by=/usr/share/keyrings/dart.gpg arch=amd64] https://storage.googleapis.com/download.dartlang.org/linux/debian stable main' | sudo tee /etc/apt/sources.list.d/dart_stable.list
$ sudo apt-get update
$ sudo apt-get install dart

```

#### 验证安装

无论使用哪种系统，安装完成后都可以通过以下命令确认 SDK 正常工作：

```

$ dart --version
$ dart --help

```

如果能看到版本号和帮助信息，说明安装成功。

### 使用 DartPad 在线编辑器

如果你暂时不想安装任何软件，DartPad 是最快的上手方式。

DartPad 是 Google 官方提供的在线 Dart 编辑器，无需注册，打开即用。

访问地址：https://dartpad.dev

DartPad 的特点：

特性说明 零配置打开浏览器即可编写和运行 Dart 代码 支持 Flutter可切换到 Flutter 模式，编写带界面的示例 代码分享生成链接分享你的代码片段 语法高亮内置 Dart 语法高亮和代码补全 版本可选支持选择不同 Dart/Flutter 版本

DartPad 非常适合快速验证语法和分享代码片段。但对于完整的项目开发，还是建议配置本地 IDE 环境。

### 配置 VS Code

VS Code 是目前最流行的 Dart 开发编辑器，配合官方插件可以实现极佳的开发体验。

#### 安装步骤

第一步，确保已安装 VS Code（可从 code.visualstudio.com 下载）。

第二步，打开 VS Code，进入扩展商店（快捷键 `Cmd+Shift+X` 或 `Ctrl+Shift+X`）。

第三步，搜索并安装 Dart 扩展（发布者为 Dart Code）。

安装 Dart 扩展时会自动安装 Flutter 扩展（如果检测到 Flutter SDK）。如果你后续要学习 Flutter，建议同时安装 Flutter 扩展。

#### 验证配置

安装完成后，在 VS Code 中新建一个文件，命名为 `hello.dart`。

输入以下代码，如果看到语法高亮和代码提示，说明配置成功。

### 实例

void main() {
print('Hello, RUNOOB!');
}

#### 运行代码

在 VS Code 中运行 Dart 代码有三种方式：

方式操作 右键菜单在编辑区右键 → Run Without Debugging 快捷键`Ctrl+F5`（Windows）或 `Cmd+F5`（Mac） 终端命令打开终端输入 `dart run hello.dart`

```

$ dart run hello.dart
Hello, RUNOOB!

```

#### 其他 IDE 选择

IDE插件名适用场景 IntelliJ IDEADart 插件Java/Kotlin 开发者首选 Android Studio内置支持Flutter 移动端开发 VS CodeDart 扩展通用开发，轻量快速

### Dart 项目结构初探

除了单文件运行，Dart 还支持标准化的项目结构。

通过 dart create 命令可以快速创建一个 Dart 项目骨架：

```

$ dart create my_first_app

```

生成的目录结构如下：

```

my_first_app/
├── bin/
│ └── my_first_app.dart # 程序入口文件
├── lib/
│ └── my_first_app.dart # 库代码（可复用）
├── test/
│ └── my_first_app_test.dart # 测试文件
├── pubspec.yaml # 项目配置文件（依赖管理）
├── analysis_options.yaml # 静态分析规则
└── README.md

```

各目录的用途说明：

目录/文件用途 bin/存放可执行程序的入口文件 lib/存放可被其他项目引用的库代码 test/存放单元测试文件 pubspec.yaml项目的元数据和依赖声明 analysis_options.yaml配置代码静态分析规则

进入项目目录，运行项目：

```

$ cd my_first_app
$ dart run

```

```

Hello world: my_first_app!

```

对于学习基础语法阶段，我们主要使用单文件模式。当你开始写自己的小项目时，再使用 dart create 创建标准项目结构。

---

## 第一个 Dart 程序

Source: https://www.runoob.com/dart/dart-first-program.html

## 第一个 Dart 程序

学习一门编程语言，最好的方式不是先背语法规则，而是先跑起来一个程序，再一行一行地问这是什么意思。

本章我们将从一个最简单的程序出发，逐步扩展到一个稍微完整的小程序，在这个过程中把 Dart 最基础的概念都过一遍。

Dart 语言标准文件后缀 .dart，比如 hello.dart。

### 最简单的程序：Hello, Dart!

打开 DartPad https://dartpad.dev（或本地编辑器创建 hello.dart 文件），输入以下代码：

### 实例

void main() {
print('Hello, Dart!');
}

点击运行，输出：

```

Hello, Dart!

```

本地环境运行命令：

```

dart hello.dart
```

就这 3 行代码。现在我们来把每一个字符都弄清楚。

### 逐行解析

#### void main() {

这是程序的入口函数声明，Dart 程序永远从 main() 开始执行。

拆开来看：

部分含义 void返回值类型。void 表示"这个函数执行完后不返回任何值" main函数名。这是 Dart 约定的入口函数名，不能改成别的 ()参数列表。这里是空的，表示 main 函数不接受任何参数 {函数体的开始。与最后一行的 } 配对，包裹函数的所有代码

main() 是 Dart 的唯一入口。不管你的程序有多少个文件、多少个类，执行总是从 main() 的第一行开始。

#### print('Hello, Dart!');

这是一条语句，作用是把括号内的内容输出到控制台。

拆开来看：

部分含义 print内置函数名，用于向控制台输出一行文字 (...)函数调用的参数列表 'Hello, Dart!'一个字符串字面量，用单引号 ' 包裹 ;语句结束符。每条语句必须以分号结尾 前面的两个空格缩进，表示这行代码属于 main() 函数内部。Dart 约定用 2 个空格缩进

#### }

main() 函数体的结束标志，与开头的 { 配对。

#### 整体结构图

```

void main() { ← 函数签名（返回类型 + 函数名 + 参数）
print('Hello, ← 函数体：调用 print 函数
Dart!'); ← 传入字符串参数，以 ; 结尾
} ← 函数体结束

```

### 字符串的两种写法

Dart 中字符串可以用单引号或双引号包裹，两者完全等价。

### 实例

void main() {
print('Hello, Dart!'); // 单引号
print("Hello, Dart!"); // 双引号，效果相同
}

```

Hello, Dart!
Hello, Dart!

```

什么时候用哪种引号？可以参考以下场景：

场景推荐示例 一般情况单引号（Dart 社区约定）'Hello' 字符串内含有单引号双引号"It's a beautiful day!" 字符串内含有双引号单引号'他说："你好"'

字符串内容本身含有单引号时，用双引号包裹更方便，省去转义。Dart 社区约定优先使用单引号，本教程后续也以单引号为主。

### 注释

注释是写给人看的说明文字，Dart 编译器会完全忽略它。

### 实例

void main() {
// 这是单行注释，从 // 开始到行尾
print('Hello, Dart!'); // 也可以写在语句后面

/*
这是多行注释。
可以跨越多行，
常用于临时注释掉一大段代码。
*/

/// 这是文档注释，用于描述函数、类或变量的作用。
/// IDE 会将文档注释显示为悬浮提示。
}

注释类型语法用途 单行注释// 内容解释某一行代码 多行注释/* 内容 */注释掉一段代码，或写较长说明 文档注释/// 内容描述 API，生成文档用

好注释解释"为什么"，而不是"做了什么"。代码本身已经说明了"做了什么"，注释应该补充背后的原因或意图。

### 变量：给数据起名字

程序不可能只输出固定文字，我们需要变量来存储和使用数据。

### 实例

void main() {
String name = 'RUNOOB';
int age = 25;
double height = 1.68;
bool isStudent = true;

print(name);
print(age);
print(height);
print(isStudent);
}

```

RUNOOB
25
1.68
true

```

#### 变量声明的格式

```

类型名 变量名 = 初始值 ;
│ │ │ └─ 语句结束符
│ │ └─ 赋值运算符
│ └─ 你给这个数据起的名字
└─ 告诉 Dart 这个变量存储什么类型的数据

```

#### 基本数据类型速览

类型说明示例 String文本字符串'RUNOOB'、"hello" int整数25、-10、0 double小数（浮点数）1.68、3.14 bool布尔值（真/假）true、false

数据类型的详细内容将在后续章节展开，这里先有个初步印象。

### 用 var 自动推断类型

每次都写类型名有点繁琐。

Dart 支持用 var 关键字自动推断变量类型——编译器会根据你赋的值来判断类型。

### 实例

void main() {
var name = 'RUNOOB'; // 推断为 String
var age = 25; // 推断为 int
var height = 1.68; // 推断为 double
var isStudent = true; // 推断为 bool

print(name);
print(age);
}

```

RUNOOB
25

```

var 推断的类型在编译期就确定了，之后不能再赋值成别的类型。

### 实例

void main() {
var name = 'RUNOOB';
name = 'Bob'; // 可以，仍然是 String
// name = 123; // 错误：不能把 int 赋给 String 类型的变量
}

什么时候用 var，什么时候写明类型？

场景推荐原因 局部变量，类型一眼就能看出来var更简洁 函数参数、类的字段、返回值类型写明类型代码更清晰，自文档化

### 字符串插值：把变量嵌入字符串

如果想把变量的值嵌入到字符串里，Dart 提供了非常简洁的字符串插值语法——用 $变量名 或 ${表达式}。

### 实例

void main() {
var name = 'RUNOOB';
var age = 25;

print('我叫 $name，今年 $age 岁。');
print('明年我将 ${age + 1} 岁。');
}

```

我叫 RUNOOB，今年 25 岁。
明年我将 26 岁。

```

语法用途示例 $变量名直接嵌入变量值，简洁$name ${表达式}嵌入任意表达式（计算、方法调用等）${age + 1}、${name.toUpperCase()}

相比 print('我叫 ' + name + '，今年 ' + age.toString() + ' 岁。') 这种字符串拼接写法，插值语法更清晰、更安全，是 Dart 的惯用风格。

### 一个稍完整的小程序

现在我们把前面学到的东西综合起来，写一个稍微真实一点的程序：一个简单的自我介绍生成器。

### 实例

void main() {
// 个人信息
var name = 'RUNOOB';
var age = 25;
var city = '上海';
var isStudent = false;
var hobby = 'Dart 编程';

// 生成介绍语：三元运算符判断身份
var occupation = isStudent ? '学生' : '职场人';

print('=============================');
print(' 个人简介');
print('=============================');
print('姓名：$name');
print('年龄：$age 岁');
print('城市：$city');
print('身份：$occupation');
print('爱好：$hobby');
print('-----------------------------');
print('大家好，我是来自$city的$name，');
print('今年 $age 岁，是一名$occupation。');
print('我的爱好是${hobby}，很高兴认识大家！');
print('=============================');
}

```

=============================
个人简介
=============================
姓名：RUNOOB
年龄：25 岁
城市：上海
身份：职场人
爱好：Dart 编程
-----------------------------
大家好，我是来自上海的RUNOOB，
今年 25 岁，是一名职场人。
我的爱好是Dart 编程，很高兴认识大家！
=============================

```

#### 新出现的语法：三元运算符

### 实例

var occupation = isStudent ? '学生' : '职场人';

这是 Dart 的三元运算符，格式为：

```

条件 ? 条件为真时的值 : 条件为假时的值

```

等价于以下 if/else 写法：

### 实例

String occupation;
if (isStudent) {
occupation = '学生';
} else {
occupation = '职场人';
}

三元运算符在值只有两种可能时非常简洁，后续章节会详细介绍 if/else 的完整用法。

### 代码风格：Dart 的书写规范

好的代码风格让代码更易读、更专业。

Dart 有一套官方推荐的规范，初学阶段掌握以下几点就够了。

#### 缩进用 2 个空格

### 实例

// 正确：2 个空格
void main() {
print('Hello');
}

// 不推荐：4 个空格（其他语言常见，但 Dart 社区偏好 2 格）

#### 变量和函数名用小驼峰（lowerCamelCase）

### 实例

var firstName = 'RUNOOB'; // 小驼峰：正确
// var first_name = 'RUNOOB'; // 下划线风格：非 Dart 惯例
// var FirstName = 'RUNOOB'; // 大驼峰：保留给类名

#### 类名用大驼峰（UpperCamelCase）

### 实例

class UserProfile { } // 正确
// class user_profile { } // 错误
// class userProfile { } // 错误

#### 常量用小驼峰（不是全大写）

### 实例

const maxRetries = 3; // Dart 惯例：正确
// const MAX_RETRIES = 3; // 这是 C/Java 风格，Dart 不推荐

#### 使用 dart format 自动格式化

不用死记规范，只需在项目目录下运行：

```

$ dart format .

```

Dart 官方格式化工具会自动把代码调整为标准风格，保存时 VS Code 也可以配置自动格式化。

### 常见错误与解读

初学者写第一个程序时经常遇到以下错误，来认识一下它们。

#### 错误一：忘记分号

### 实例

void main() {
print('Hello, Dart!') // 缺少 ;
}

报错信息：

```

Error: Expected ';' after this.

```

解决方法：在语句末尾加上 ;。

#### 错误二：括号不配对

### 实例

void main() {
print('Hello, Dart!' // 缺少右括号 )
}

报错信息：

```

Error: Expected ')' before this.

```

解决方法：仔细检查每个 ( 是否都有对应的 )。VS Code 会高亮显示不匹配的括号。

#### 错误三：字符串引号不匹配

```

void main() {
print('Hello, Dart!"); // 一个单引号、一个双引号
}
```

报错信息：

```

Error: Unterminated string literal.

```

解决方法：确保字符串两端用同一种引号。

#### 错误四：main 拼写错误

### 实例

void Main() { // M 大写了
print('Hello, Dart!');
}

Dart 区分大小写，Main 和 main 是完全不同的名字。

程序能编译，但运行时会报错：

```

Error: No 'main' method found.

```

解决方法：入口函数必须是小写的 main。

#### 错误五：变量在使用前未赋值

### 实例

void main() {
String name;
print(name); // name 未初始化
}

Dart 空安全机制的报错：

```

Error: Non-nullable variable 'name' must be assigned before it can be used.

```

解决方法：声明变量时同时赋初值，或声明为可空类型 String? name（后续章节详细介绍）。

---

## Dart 基础语法

Source: https://www.runoob.com/dart/dart-basic-syntax.html

## Dart 基础语法

每种编程语言都有自己的一套基本规则，Dart 也不例外。

本章介绍 Dart 程序的基本结构、注释写法、语句与表达式的区别，以及标识符的命名规范。

### 程序结构与 main() 函数

Dart 文件名后缀为 .dart，每个 Dart 程序都必须有一个入口点——main() 函数。

程序从 main() 函数的第一行开始执行，到最后一行的花括号结束。

### 实例

一个最简 Dart 程序的结构：

// 程序入口：main() 函数是 Dart 程序的起点
void main() {
// 在这里编写你的代码
print('欢迎来到 RUNOOB Dart 教程！');
}

```

欢迎来到 RUNOOB Dart 教程！

```

让我们拆解这段代码的每个部分：

代码元素说明 void返回值类型，void 表示该函数不返回任何值 main函数名，这是 Dart 规定的入口函数名，不可更改 ()参数列表，main() 可以不接收参数，也可以接收命令行参数 { }函数体，花括号内的代码是函数执行的内容 print()内置函数，用于向控制台输出文本 ;语句结束符，Dart 中每条语句必须以分号结尾

Dart 严格要求每条语句以分号（;）结尾。忘记分号是初学者最常见的语法错误。好在 VS Code 的 Dart 插件会自动提示缺少分号的位置。

#### 带命令行参数的 main() 函数

main() 函数可以接收一个字符串列表作为命令行参数：

### 实例

// List<String> args 接收命令行传入的参数列表
void main(List<String> args) {
// 输出参数个数
print('参数个数：${args.length}');

// 遍历并打印每个参数
for (var i = 0; i &lt; args.length; i++) {
print('参数 $i: ${args[i]}');
}
}

```

$ dart run hello.dart runoob dart 教程
参数个数：3
参数 0: runoob
参数 1: dart
参数 2: 教程

```

这里用到了 List 类型和 for 循环，我们会在后续章节详细介绍。

### 注释写法

注释是对代码的解释说明，编译器会忽略注释内容。

好的注释能让别人（包括几个月后的你自己）更快理解代码的意图。

Dart 支持三种注释方式：

类型语法适用场景 单行注释// 注释内容简短说明某一行代码的作用 多行注释/* 注释内容 */跨多行的详细说明 文档注释/// 文档注释为函数、类生成 API 文档

### 实例

/// 这是一个文档注释，用于描述程序的功能
/// 文档注释可以用 [dart doc] 命令生成 HTML 文档
void main() {
// 单行注释：输出欢迎信息
print('欢迎学习 Dart！'); // 行内注释：这也是合法的

/*
* 多行注释：适合写较长的说明
* 第二行注释
* 第三行注释
*/
print('RUNOOB 教程 - Dart 基础语法');
}

```

欢迎学习 Dart！
RUNOOB 教程 - Dart 基础语法

```

注释的原则是"解释为什么，而不是做了什么"。代码本身已经说明了做了什么，注释应该解释为什么这样写，或者补充代码无法表达的背景信息。

### 语句与表达式

在 Dart 中，代码由语句（Statement）和表达式（Expression）组成。

理解两者的区别有助于你更好地组织代码逻辑。

#### 表达式（Expression）

表达式是一段可以被求值的代码，它会产出一个值。

### 实例

void main() {
// 以下每行都是一个表达式，它们都产出一个值
42; // 字面量表达式，值为 42
3 + 5; // 算术表达式，值为 8
'Hello'.length; // 属性访问表达式，值为 5
(10 > 5); // 关系表达式，值为 true
}

#### 语句（Statement）

语句是执行某个动作的指令，它不产出值。

### 实例

void main() {
// 变量声明语句：声明一个变量并赋值
var name = 'RUNOOB';

// 函数调用语句：调用 print() 执行输出动作
print(name);

// 控制流语句：根据条件决定执行哪段代码
if (name.length > 3) {
print('名称长度大于 3');
}
}

```

RUNOOB
名称长度大于 3

```

简单来说：表达式产生值，语句执行动作。

你可以在需要值的地方使用表达式，但不能使用语句。

一个常见的误区是把所有代码都当成"语句"。实际上，Dart 中很多结构（如函数调用、变量赋值）本身就是表达式，这意味着它们可以被嵌套使用。这是 Dart 和其他 C 系语言的共同特点。

### 标识符命名规范

标识符（Identifier）是你给变量、函数、类等代码元素起的名字。

Dart 对标识符有一些规则和约定：

#### 命名规则（必须遵守）

标识符只能由字母、数字、下划线（_）和美元符号（$）组成。

标识符不能以数字开头。

标识符不能是 Dart 的保留关键字（如 class、if、var 等）。

#### 命名约定（推荐遵守）

代码元素命名风格示例 类名大驼峰（UpperCamelCase）StudentInfo、HttpClient 枚举类型大驼峰（UpperCamelCase）Color、Status 库名、文件名小写 + 下划线（lowercase_with_underscores）my_utils、user_service 变量名小驼峰（lowerCamelCase）userName、maxCount 函数名小驼峰（lowerCamelCase）getUserInfo()、calculateTotal() 常量名小驼峰（lowerCamelCase）piValue、defaultPort 私有成员下划线前缀_internalState、_privateMethod()

### 实例

各种命名风格的实际应用：

// 类名：大驼峰
class UserAccount {
// 公开成员：小驼峰
String userName;

// 私有成员：下划线前缀（只在本文件中可访问）
String _password;

// 构造函数
UserAccount(this.userName, this._password);

// 公开方法：小驼峰
bool verifyPassword(String input) {
return input == _password;
}

// 私有方法：下划线前缀
void _encryptPassword() {
// 加密逻辑...
}
}

// 入口函数：小写
void main() {
// 变量名：小驼峰
var userAccount = UserAccount('runoob', '123456');

// 常量：小驼峰（Dart 风格）
const maxLoginAttempts = 3;

print('用户名: ${userAccount.userName}');
print('最大登录尝试次数: $maxLoginAttempts');
}

```

用户名: runoob
最大登录尝试次数: 3

```

Dart 的命名约定中，常量使用小驼峰（lowerCamelCase），而不是其他语言常见的全大写+下划线（UPPER_CASE）。这是 Dart 社区的独特风格，请务必注意。

---

## Dart 变量与数据类型

Source: https://www.runoob.com/dart/dart-variables-and-types.html

## Dart 变量与数据类型

变量是程序中存储数据的基本单元，数据类型决定了变量能存储什么值。

本章介绍 Dart 中的变量声明方式、基本数据类型以及重要的空安全机制。

### 变量的声明方式：var、final、const

Dart 提供了三种声明变量的关键字，分别适用于不同的场景。

理解它们的区别是学好 Dart 的第一步。

#### var——类型推断变量

使用 var 声明的变量，类型由初始值自动推断，且类型一旦确定就不能更改。

### 实例

void main() {
// 类型推断：name 的类型被推断为 String
var name = 'RUNOOB';

// 类型推断：age 的类型被推断为 int
var age = 10;

print('$name 已经 $age 年了');

// 以下写法会报错——类型确定后不能再赋其他类型的值
// name = 123; // 错误：不能将 int 赋值给 String
}

```

RUNOOB 已经 10 年了

```

var 不是"无类型"或"动态类型"，它的类型在赋值那一刻就固定了。这和 JavaScript 的 var 完全不同。

#### final——运行时常量

使用 final 声明的变量只能赋值一次，但赋值时机可以延迟到运行时。

### 实例

void main() {
// final 变量：一旦赋值就不能再修改
final currentTime = DateTime.now();
print('当前时间: $currentTime');

// 延迟赋值也是允许的
final String greeting;
greeting = 'Hello, RUNOOB!'; // 第一次赋值 OK
// greeting = 'Hello again'; // 错误：final 变量不能再次赋值

print(greeting);
}

```

当前时间: 2026-06-12 10:30:00.000
Hello, RUNOOB!

```

#### const——编译时常量

const 比 final 更严格：它的值必须在编译时就能确定。

### 实例

void main() {
// const 必须在编译时就能确定值
const siteName = 'RUNOOB';
const maxUsers = 100;

print('站点: $siteName, 最大用户数: $maxUsers');

// 以下写法会报错——DateTime.now() 是运行时才能确定的值
// const currentTime = DateTime.now(); // 错误！
}

```

站点: RUNOOB, 最大用户数: 100

```

#### 三者对比总结

关键字可修改赋值时机典型场景 var可以任意值会变化的普通变量 final不可修改运行时确定从 API 获取的配置、当前时间 const不可修改编译时确定数学常量、固定配置、颜色值

一个实用建议：默认使用 final。当你发现需要修改它时，再改为 var。这个习惯能让代码更安全，减少意外修改导致的 bug。

### 数字类型：int 与 double

Dart 的数字类型分为两种：int（整数）和 double（浮点数）。

它们的共同父类是 num。

### 实例

void main() {
// int：整数，范围是 -2^63 到 2^63-1
int viewCount = 10000;
int negativeNumber = -42;

// double：双精度浮点数
double price = 9.99;
double pi = 3.1415926;

// num 可以同时接收 int 和 double
num someNumber = 10;
someNumber = 10.5; // 合法

print('浏览次数: $viewCount');
print('价格: ¥$price');
print('圆周率: $pi');
}

```

浏览次数: 10000
价格: ¥9.99
圆周率: 3.1415926

```

#### 数字常用方法

### 实例

import 'dart:math'; // 导入数学库

void main() {
int a = -10;
double b = 3.7;

// 绝对值
print('a 的绝对值: ${a.abs()}'); // 10

// 四舍五入取整
print('b 四舍五入: ${b.round()}'); // 4

// 向下取整
print('b 向下取整: ${b.floor()}'); // 3

// 向上取整
print('b 向上取整: ${b.ceil()}'); // 4

// 字符串转数字
int parsed = int.parse('42');
double parsedDouble = double.parse('3.14');
print('解析结果: $parsed, $parsedDouble');

// 生成随机数（需要 import 'dart:math'）
var random = Random().nextInt(100);
print('随机数(0-99): $random');
}

```

a 的绝对值: 10
b 四舍五入: 4
b 向下取整: 3
b 向上取整: 4
解析结果: 42, 3.14
随机数(0-99): 67

```

### 字符串：String 与插值语法

String 用于表示文本，Dart 的字符串支持单引号和双引号两种写法。

字符串插值（String Interpolation）是 Dart 最方便的特性之一。

### 实例

void main() {
// 字符串可以用单引号或双引号
String singleQuoted = 'Hello';
String doubleQuoted = "World";

// 多行字符串：使用三个引号
String multiLine = '''
这是第一行
这是第二行
这是第三行 - RUNOOB
''';

// 字符串插值：使用 $变量名 或 ${表达式}
String site = 'RUNOOB';
int years = 10;

// $变量名：直接引用变量
print('欢迎来到 $site');

// ${表达式}：嵌入复杂表达式
print('$site 已经陪伴我们 ${years + 1} 年了');

// 转义字符
print('换行符: 第一行\n第二行');
print('制表符: 列1\t列2');
}

```

欢迎来到 RUNOOB
RUNOOB 已经陪伴我们 11 年了
换行符: 第一行
第二行
制表符: 列1 列2

```

#### 字符串常用方法

### 实例

void main() {
String text = ' Hello, RUNOOB! ';

// 去除首尾空格
print('trim: [${text.trim()}]');

// 转大写 / 转小写
print('大写: ${text.toUpperCase()}');
print('小写: ${text.toLowerCase()}');

// 是否包含某个子串
print('包含 RUNOOB? ${text.contains('RUNOOB')}');

// 是否以某字符串开头/结尾
print('以空格开头? ${text.startsWith(' ')}');
print('以!结尾? ${text.trim().endsWith('!')}');

// 字符串替换
print('替换: ${text.replaceAll('RUNOOB', 'DART')}');

// 字符串分割
String csv = 'Dart,Flutter,Google';
List<String> items = csv.split(',');
print('分割结果: $items');

// 字符串拼接
print('拼接: ' + 'Hello' + ', ' + 'World');
}

```

trim: [Hello, RUNOOB!]
大写: HELLO, RUNOOB!
小写: hello, runoob!
包含 RUNOOB? true
以空格开头? true
以!结尾? true
替换: Hello, DART!
分割结果: [Dart, Flutter, Google]
拼接: Hello, World

```

### 布尔值 bool

bool 类型只有两个值：true 和 false。

它主要用于条件判断，配合 if 语句使用。

### 实例

void main() {
bool isLoggedIn = true;
bool hasPermission = false;

// 布尔表达式
bool canAccess = isLoggedIn && hasPermission; // 与：两者都为 true
bool isGuest = !isLoggedIn; // 非：取反

print('可以访问: $canAccess');
print('是访客: $isGuest');

// 比较运算产生 bool 值
int age = 20;
bool isAdult = age >= 18;
print('是否成年: $isAdult');

// 直接用于条件判断
if (isAdult) {
print('欢迎访问 RUNOOB');
}
}

```

可以访问: false
是访客: false
是否成年: true
欢迎访问 RUNOOB

```

Dart 是强类型语言，条件判断中只接受 bool 类型。这和 JavaScript 不同——在 Dart 中，if(0) 或 if('') 这样的写法会直接报错。

### 空安全（Null Safety）

空安全是 Dart 2.12 引入的最重要特性。

它的核心思想是：默认情况下，所有变量都不能为 null，除非你明确声明。

#### 可空类型与非空类型

在类型后面加 ? 表示该变量可以为 null。

### 实例

void main() {
// 非空类型：不能赋值为 null
String name = 'RUNOOB';
// name = null; // 错误！name 不能为 null

// 可空类型：允许为 null
String? nickname;
nickname = null; // OK
nickname = 'R'; // 也 OK

print('昵称: $nickname');

// int? 也可空
int? score;
print('分数: $score'); // 输出 null
}

```

昵称: R
分数: null

```

#### 空值处理操作符

操作符说明示例 ??如果左侧为 null，使用右侧值name ?? '默认名' ??=如果变量为 null，则赋值为右侧name ??= '默认名' ?.安全访问：对象为 null 时跳过调用user?.name !断言非 null（谨慎使用）name!

### 实例

void main() {
String? username;

// ?? 操作符：如果 username 为 null，使用默认值
String displayName = username ?? 'RUNOOB 访客';
print(displayName); // RUNOOB 访客

// ??= 操作符：如果为 null 则赋值
username ??= '新用户';
print(username); // 新用户

// ?. 安全访问：只在非 null 时调用
String? maybeNull;
print(maybeNull?.length); // null（不会报错）

maybeNull = 'Hello';
print(maybeNull?.length); // 5

// ! 断言：我确定它不是 null（为 null 时会抛异常）
String definitelyNotNull = 'RUNOOB';
print(definitelyNotNull!.length); // 6
}

```

RUNOOB 访客
新用户
null
5
6

```

尽量避免使用 ! 断言操作符。如果你频繁使用 !，说明你的类型设计可能有问题。优先使用 ??、??= 和 ?. 来安全地处理空值。

### 类型转换

在 Dart 中，不同类型的值可以相互转换。

### 实例

void main() {
// 数字 → 字符串
int count = 42;
String countStr = count.toString();
print('数字转字符串: $countStr (类型: ${countStr.runtimeType})');

// 字符串 → 数字
String priceStr = '19.99';
double price = double.parse(priceStr);
print('字符串转数字: ¥$price');

// int &#x2194; double
int integer = 10;
double decimal = integer.toDouble();
print('int 转 double: $decimal');

double pi = 3.14;
int piInt = pi.toInt(); // 截断小数部分
print('double 转 int (截断): $piInt');

// 任何值 → 字符串
bool flag = true;
print('bool 转字符串: ${flag.toString()}');

// 字符串 → bool（需要手动判断）
String boolStr = 'true';
bool parsed = boolStr == 'true';
print('字符串转 bool: $parsed');
}

```

数字转字符串: 42 (类型: String)
字符串转数字: ¥19.99
int 转 double: 10.0
double 转 int (截断): 3
bool 转字符串: true
字符串转 bool: true

```

---

## Dart 控制流

Source: https://www.runoob.com/dart/dart-control-flow.html

## Dart 控制流

控制流决定了程序中的代码按照什么顺序执行。

本章介绍 Dart 中的条件判断（if/else、switch）和循环语句（for、while），以及 break 与 continue 的使用。

### if / else / else if 条件判断

if 语句是最基本的控制流结构，它根据条件决定是否执行某段代码。

### 实例

基本的 if 判断：

void main() {
int score = 85;

// 基本 if：条件为 true 时执行
if (score >= 60) {
print('RUNOOB 恭喜，你及格了！');
}

// if-else：二选一
if (score >= 90) {
print('优秀');
} else {
print('继续加油');
}

// if-else if-else：多分支选择
if (score >= 90) {
print('等级：A');
} else if (score >= 80) {
print('等级：B');
} else if (score >= 70) {
print('等级：C');
} else if (score >= 60) {
print('等级：D');
} else {
print('等级：F（不及格）');
}
}

```

RUNOOB 恭喜，你及格了！
继续加油
等级：B

```

Dart 的条件表达式必须是 bool 类型。不能像 JavaScript 那样用 0、''、null 等作为条件——这些写法在 Dart 中会直接报编译错误。

#### if 语句的嵌套

if 语句内部可以再嵌套 if，处理更复杂的逻辑。

### 实例

void main() {
bool isLoggedIn = true;
bool isAdmin = false;
String username = 'runoob';

if (isLoggedIn) {
// 外层判断：是否登录
if (isAdmin) {
// 内层判断：是否是管理员
print('欢迎回来，管理员 $username');
} else {
print('欢迎回来，$username');
}
} else {
print('请先登录');
}
}

```

欢迎回来，runoob

```

嵌套层次过多会降低可读性，建议不超过 3 层。

如果逻辑确实复杂，可以考虑用提前返回（return）或拆分函数来简化。

### 实例

用提前返回替代深层嵌套：

// 用提前返回替代深层嵌套，可读性更好
String getAccessMessage(bool isLoggedIn, bool isAdmin, String name) {
// 先处理"不满足条件就返回"的情况
if (!isLoggedIn) {
return '请先登录';
}

if (!isAdmin) {
return '你没有管理权限';
}

// 最后才是主逻辑
return '欢迎回来，管理员 $name';
}

void main() {
print(getAccessMessage(true, false, 'runoob'));
print(getAccessMessage(true, true, 'runoob'));
}

```

你没有管理权限
欢迎回来，管理员 runoob

```

### switch / case 多分支选择

当一个变量有多个固定的可能值时，switch 比 if-else if 更清晰。

### 实例

void main() {
String day = '星期三';

switch (day) {
case '星期一':
print('新的一周开始了');
break; // break 结束当前 case，防止"穿透"到下一个 case
case '星期二':
print('第二天的奋斗');
break;
case '星期三':
print('一周过半了');
break;
case '星期四':
print('快到周五了');
break;
case '星期五':
print('TGIF！');
break;
case '星期六':
case '星期日':
// 多个 case 可以共用一段代码
print('周末愉快！');
break;
default:
// 所有 case 都不匹配时执行
print('未知的日期');
}
}

```

一周过半了

```

Dart 的 switch 不会自动"穿透"（fall-through），每个非空的 case 必须以 break、return、throw 或 continue 结尾。如果想实现穿透效果，使用 continue 配合标签（label）。

#### switch 表达式（Dart 3.0+）

Dart 3.0 引入了 switch 表达式，可以直接产出值，写法更简洁。

### 实例

void main() {
String day = '星期三';

// switch 表达式：直接返回值
String mood = switch (day) {
'星期一' => '有点困',
'星期二' => '开始适应',
'星期三' => '干劲十足',
'星期四' => '期待周末',
'星期五' => '心情愉悦',
'星期六' || '星期日' => '自由自在',
_ => '未知' // _ 是通配符，相当于 default
};

print('RUNOOB 心情: $mood');
}

```

RUNOOB 心情: 干劲十足

```

switch 表达式要求覆盖所有可能的情况，编译器会帮你检查是否有遗漏。

### for 循环

for 循环是最常用的循环结构，适合在知道循环次数时使用。

#### 标准 for 循环

### 实例

void main() {
// 标准 for 循环：初始化; 条件; 更新
for (int i = 1; i <= 5; i++) {
print('RUNOOB 第 $i 次循环');
}

// 循环变量也可以在外部声明
int j = 0;
for (; j < 3; j++) {
print('j = $j');
}
}

```

RUNOOB 第 1 次循环
RUNOOB 第 2 次循环
RUNOOB 第 3 次循环
RUNOOB 第 4 次循环
RUNOOB 第 5 次循环
j = 0
j = 1
j = 2

```

#### for-in 循环（遍历集合）

for-in 是遍历 List、Set 等集合最简洁的方式。

### 实例

void main() {
List<String> fruits = ['苹果', '香蕉', '橙子', '葡萄'];

// for-in：逐个取出集合中的元素
for (var fruit in fruits) {
print('RUNOOB 水果: $fruit');
}
}

```

RUNOOB 水果: 苹果
RUNOOB 水果: 香蕉
RUNOOB 水果: 橙子
RUNOOB 水果: 葡萄

```

#### forEach 方法

List 和 Set 还提供了 forEach 方法，使用回调函数处理每个元素。

### 实例

void main() {
List<int> scores = [85, 92, 78, 60];

// forEach：对每个元素执行回调函数
scores.forEach((score) {
String result = score >= 60 ? '及格' : '不及格';
print('RUNOOB 分数: $score -> $result');
});
}

```

RUNOOB 分数: 85 -> 及格
RUNOOB 分数: 92 -> 及格
RUNOOB 分数: 78 -> 及格
RUNOOB 分数: 60 -> 及格

```

### while 与 do-while 循环

while 循环在每次迭代前检查条件，适合循环次数不确定的场景。

### 实例

void main() {
// while：先判断条件，再执行
int count = 3;
while (count > 0) {
print('RUNOOB 倒计时: $count');
count--; // 记得更新条件变量，否则会死循环
}
print('发射！');
}

```

RUNOOB 倒计时: 3
RUNOOB 倒计时: 2
RUNOOB 倒计时: 1
发射！

```

do-while 与 while 的区别是：它至少会执行一次循环体，因为条件在最后才检查。

### 实例

void main() {
// do-while：先执行一次，再判断条件
int num = 0;
do {
print('RUNOOB 至少执行一次，num = $num');
num++;
} while (num < 0); // 条件一开始就是 false

print('循环结束，num = $num');
}

```

RUNOOB 至少执行一次，num = 0
循环结束，num = 1

```

do-while 保证循环体至少执行一次，即使条件一开始就不满足。这在"先尝试再判断"的场景中很有用，比如读取用户输入直到有效。

#### while vs for：如何选择

场景推荐循环原因 循环次数确定for初始化、条件、更新写在一起，一目了然 循环次数不确定while只关注条件，更灵活 至少执行一次do-while先执行再判断 遍历集合for-in 或 forEach最简洁的集合遍历方式

### break 与 continue

break 用于立即退出整个循环，continue 用于跳过当前迭代进入下一轮。

### 实例

void main() {
// break：找到目标后立即停止循环
List<String> names = ['Alice', 'Bob', 'runoob', 'David'];

for (var name in names) {
print('检查: $name');
if (name == 'runoob') {
print('找到 RUNOOB 了！');
break; // 退出循环，不再检查后面的元素
}
}

print('---分隔线---');

// continue：跳过当前迭代
for (int i = 1; i <= 5; i++) {
if (i == 3) {
continue; // 跳过 i=3 这次迭代
}
print('RUNOOB 第 $i 次');
}
}

```

检查: Alice
检查: Bob
检查: runoob
找到 RUNOOB 了！
---分隔线---
RUNOOB 第 1 次
RUNOOB 第 2 次
RUNOOB 第 4 次
RUNOOB 第 5 次

```

#### 带标签的 break 与 continue

在嵌套循环中，标签可以让你跳出指定的外层循环。

### 实例

void main() {
// 标签 outerLoop 标记外层循环
outerLoop:
for (int i = 1; i <= 3; i++) {
for (int j = 1; j <= 3; j++) {
print('i=$i, j=$j');
if (i == 2 && j == 2) {
print('找到目标，退出外层循环');
break outerLoop; // 跳出标记为 outerLoop 的外层循环
}
}
}
print('循环结束');
}

```

i=1, j=1
i=1, j=2
i=1, j=3
i=2, j=1
i=2, j=2
找到目标，退出外层循环
循环结束

```

标签（label）应该谨慎使用。大多数情况下，将嵌套循环逻辑提取为独立的函数，用 return 来退出会更清晰。

---

## Dart 集合类型

Source: https://www.runoob.com/dart/dart-collections.html

## Dart 集合类型

集合是用来存放多个数据的数据结构。

Dart 内置了三种核心集合类型：List（列表）、Set（集合）和 Map（映射），本章将逐一介绍它们的用法。

### List：有序列表

List 是最常用的集合类型，用于存放一组有序的元素，允许重复。

你可以通过索引（从 0 开始）来访问 List 中的任意元素。

#### 创建 List

### 实例

void main() {
// 使用字面量创建 List
List<String> sites = ['RUNOOB', 'Google', 'GitHub'];
print('网站列表: $sites');

// 类型推断（省略泛型）
var numbers = [1, 2, 3, 4, 5];
print('数字列表: $numbers');

// 创建空列表
var emptyList = <String>[];

// 使用 List.filled 创建固定长度的列表
var zeros = List.filled(3, 0); // [0, 0, 0]
print('填充列表: $zeros');

// 使用 List.generate 生成列表
var squares = List.generate(5, (i) => i * i);
print('平方数列表: $squares');
}

```

网站列表: [RUNOOB, Google, GitHub]
数字列表: [1, 2, 3, 4, 5]
填充列表: [0, 0, 0]
平方数列表: [0, 1, 4, 9, 16]

```

#### 访问和修改 List 元素

### 实例

void main() {
var fruits = ['苹果', '香蕉', '橙子'];

// 通过索引访问（从 0 开始）
print('第一个水果: ${fruits[0]}'); // 苹果
print('最后一个: ${fruits[fruits.length - 1]}'); // 橙子

// 修改元素
fruits[1] = '葡萄';
print('修改后: $fruits'); // [苹果, 葡萄, 橙子]

// 获取长度
print('水果数量: ${fruits.length}');

// 检查是否为空
print('列表为空? ${fruits.isEmpty}');
print('列表非空? ${fruits.isNotEmpty}');
}

```

第一个水果: 苹果
最后一个: 橙子
修改后: [苹果, 葡萄, 橙子]
水果数量: 3
列表为空? false
列表非空? true

```

#### List 常用操作

### 实例

void main() {
var list = ['RUNOOB', 'Dart'];

// 添加元素
list.add('Flutter'); // 添加单个
list.addAll(['Google', 'AI']); // 添加多个
print('添加后: $list');

// 插入元素（在指定位置）
list.insert(1, '教程');
print('插入后: $list');

// 删除元素
list.remove('AI'); // 按值删除
list.removeAt(0); // 按索引删除
list.removeLast(); // 删除最后一个
print('删除后: $list');

// 查找元素
bool hasDart = list.contains('Dart');
int index = list.indexOf('Dart');
print('包含 Dart? $hasDart, 位置: $index');

// 排序
var nums = [3, 1, 4, 1, 5, 9];
nums.sort();
print('排序后: $nums');

// 反转
var reversed = nums.reversed.toList();
print('反转后: $reversed');

// 截取子列表
var subList = nums.sublist(0, 3);
print('前3个: $subList');
}

```

添加后: [RUNOOB, Dart, Flutter, Google, AI]
插入后: [RUNOOB, 教程, Dart, Flutter, Google, AI]
删除后: [教程, Dart, Flutter]
包含 Dart? true, 位置: 1
排序后: [1, 1, 3, 4, 5, 9]
反转后: [9, 5, 4, 3, 1, 1]
前3个: [1, 1, 3]

```

#### List 的函数式方法

Dart 的 List 支持 map、where、reduce 等函数式方法，让数据处理更加简洁。

### 实例

void main() {
var scores = [55, 78, 92, 60, 45, 88];

// where：过滤（保留满足条件的元素）
var passed = scores.where((s) => s >= 60);
print('及格分数: $passed');

// map：映射（将每个元素转换为新值）
var grades = scores.map((s) => s >= 60 ? '及格' : '不及格');
print('评级: $grades');

// where + map 链式调用
var highScores = scores
.where((s) => s >= 80)
.map((s) => '高分: $s')
.toList();
print('高分列表: $highScores');

// reduce：累积计算
var total = scores.reduce((sum, s) => sum + s);
print('RUNOOB 总分: $total');

// fold：带初始值的累积计算
var avg = scores.fold(0, (sum, s) => sum + s) / scores.length;
print('平均分: ${avg.toStringAsFixed(1)}');

// any / every：存在性判断
bool hasFullMark = scores.any((s) => s == 100);
bool allPassed = scores.every((s) => s >= 60);
print('有满分吗? $hasFullMark');
print('全部及格? $allPassed');
}

```

及格分数: (78, 92, 60, 88)
评级: (不及格, 及格, 及格, 及格, 不及格, 及格)
高分列表: [高分: 92, 高分: 88]
RUNOOB 总分: 418
平均分: 69.7
有满分吗? false
全部及格? false

```

map() 和 where() 返回的是 Iterable（惰性求值），需要用 toList() 转换为 List 才会真正执行计算。如果你只需要遍历一次，直接用 Iterable 即可，不需要转换。

### Set：唯一元素集合

Set 和 List 类似，但 Set 中的每个元素只能出现一次，不允许重复。

Set 是无序的，不能通过索引访问元素。

### 实例

void main() {
// 创建 Set（元素自动去重）
Set<String> tags = {'Dart', 'Flutter', 'Dart', 'RUNOOB'};
print('标签集合: $tags'); // {Dart, Flutter, RUNOOB}，重复的 Dart 被去掉了
print('标签数量: ${tags.length}'); // 3

// 添加元素
tags.add('Google');
tags.add('Dart'); // Dart 已存在，不会重复添加
print('添加后: $tags');

// 删除元素
tags.remove('Google');
print('删除后: $tags');

// 检查是否包含
print('包含 RUNOOB? ${tags.contains('RUNOOB')}');

// 集合运算
var setA = {1, 2, 3, 4};
var setB = {3, 4, 5, 6};

print('交集: ${setA.intersection(setB)}'); // {3, 4}
print('并集: ${setA.union(setB)}'); // {1, 2, 3, 4, 5, 6}
print('差集: ${setA.difference(setB)}'); // {1, 2}
}

```

标签集合: {Dart, Flutter, RUNOOB}
标签数量: 3
添加后: {Dart, Flutter, RUNOOB, Google}
删除后: {Dart, Flutter, RUNOOB}
包含 RUNOOB? true
交集: {3, 4}
并集: {1, 2, 3, 4, 5, 6}
差集: {1, 2}

```

Set 的典型使用场景是"去重"——当你不需要重复元素时，用 Set 比用 List 手动去重要高效得多。

### Map：键值对映射

Map 用于存储键值对（Key-Value Pair），每个键对应一个值。

键在 Map 中必须是唯一的，值可以重复。

#### 创建和访问 Map

### 实例

void main() {
// 使用字面量创建 Map
Map<String, String> siteInfo = {
'name': 'RUNOOB',
'url': 'https://www.runoob.com',
'type': '编程教程',
};

// 通过键访问值
print('站点名称: ${siteInfo['name']}');
print('站点 URL: ${siteInfo['url']}');

// 访问不存在的键返回 null
print('描述: ${siteInfo['description']}'); // null

// 添加/修改键值对
siteInfo['language'] = '中文';
siteInfo['name'] = 'RUNOOB.COM'; // 修改已有键的值
print('更新后: $siteInfo');

// 获取所有键和所有值
print('所有键: ${siteInfo.keys}');
print('所有值: ${siteInfo.values}');

// 检查键是否存在
print('有 url 键? ${siteInfo.containsKey('url')}');
print('有 desc 键? ${siteInfo.containsKey('desc')}');
}

```

站点名称: RUNOOB
站点 URL: https://www.runoob.com
描述: null
更新后: {name: RUNOOB.COM, url: https://www.runoob.com, type: 编程教程, language: 中文}
所有键: (name, url, type, language)
所有值: (RUNOOB.COM, https://www.runoob.com, 编程教程, 中文)
有 url 键? true
有 desc 键? false

```

#### Map 常用操作

### 实例

void main() {
var scores = {
'runoob': 95,
'Alice': 87,
'Bob': 72,
};

// 遍历 Map
scores.forEach((name, score) {
print('$name: $score 分');
});

// 删除键值对
scores.remove('Bob');
print('删除 Bob 后: $scores');

// 获取值或默认值
int aliceScore = scores['Alice'] ?? 0;
int eveScore = scores['Eve'] ?? 0; // 不存在，返回默认值 0
print('Alice: $aliceScore, Eve: $eveScore');

// putIfAbsent：键不存在时才添加
scores.putIfAbsent('runoob', () => 100); // 已存在，不添加
scores.putIfAbsent('David', () => 80); // 不存在，添加
print('最终: $scores');

// 获取长度
print('人数: ${scores.length}');
}

```

runoob: 95 分
Alice: 87 分
Bob: 72 分
删除 Bob 后: {runoob: 95, Alice: 87}
Alice: 87, Eve: 0
最终: {runoob: 95, Alice: 87, David: 80}
人数: 3

```

Map 的键可以是任何类型（String、int 等），但必须保证键的相等性有意义。使用自定义对象作为键时，需要确保正确实现了 == 和 hashCode。

### 集合展开运算符 ... 与 ...?

展开运算符可以将一个集合的元素"展开"到另一个集合中。

### 实例

void main() {
var basics = ['Dart', 'Flutter'];
var advanced = ['异步编程', '状态管理'];

// ... 展开运算符：将另一个集合的元素插入
var allCourses = ['RUNOOB 入门', ...basics, ...advanced];
print('所有课程: $allCourses');

// ...? 空安全展开：如果集合为 null，则跳过
List<String>? optionalList; // 可能为 null
var safeList = ['第一项', ...?optionalList];
print('安全展开: $safeList'); // 只有第一项

optionalList = ['额外内容'];
safeList = ['第一项', ...?optionalList];
print('有值展开: $safeList');
}

```

所有课程: [RUNOOB 入门, Dart, Flutter, 异步编程, 状态管理]
安全展开: [第一项]
有值展开: [第一项, 额外内容]

```

#### 集合中的 if 和 for（集合控制流）

Dart 允许在集合字面量中直接使用 if 和 for，这是非常实用的语法糖。

### 实例

void main() {
bool showAdmin = true;

// 集合中的 if：根据条件决定是否包含元素
var menuItems = [
'首页',
'教程',
if (showAdmin) '管理后台', // 条件为 true 时才加入
'关于',
];
print('菜单: $menuItems');

// 集合中的 for：从另一个集合生成元素
var numbers = [1, 2, 3];
var doubled = [
for (var n in numbers) n * 2, // 遍历生成新元素
];
print('翻倍: $doubled');

// if + for 可以组合使用
var tags = ['Dart', 'Flutter'];
var links = [
'RUNOOB 首页',
for (var tag in tags) 'https://runoob.com/$tag',
];
print('链接: $links');
}

```

菜单: [首页, 教程, 管理后台, 关于]
翻倍: [2, 4, 6]
链接: [RUNOOB 首页, https://runoob.com/Dart, https://runoob.com/Flutter]

```

### 三种集合对比

特性ListSetMap 有序性有序无序无序（但有迭代顺序） 允许重复允许不允许键不允许重复，值允许 索引访问支持 []不支持通过键访问 [] 典型场景有序列表、序列去重、标签集合键值映射、配置项 创建字面量[]{}{key: value}

#### 注意事项

Set 和 Map 都使用 {} 作为字面量。

区别在于：

- Map 使用 key: value 形式（有冒号）
- Set 直接写元素（没有冒号）

例如：

```

var map = {'name': 'Tom'}; // Map
var set = {'Tom', 'Jerry'}; // Set
```

需要特别注意：

```

var x = {};
```

这不会创建空 Set，而会被 Dart 推断为：

```

Map<dynamic, dynamic>
```

因为空花括号无法判断元素类型，Dart 默认按 Map 处理。

如果要创建空 Set，必须显式指定：

```

var x = <String>{};
```

或：

```

Set<String> x = {};
```

---

## Dart 函数

Source: https://www.runoob.com/dart/dart-functions.html

## Dart 函数

函数是组织代码的基本单元，它将一段可复用的逻辑封装起来，需要时调用即可。

Dart 的函数系统非常灵活，支持命名参数、可选参数、箭头函数和高阶函数等特性。

### 函数定义与返回值

函数由返回类型、函数名、参数列表和函数体组成。

### 实例

// 定义一个简单的函数
// String 是返回类型，greet 是函数名，(String name) 是参数列表
String greet(String name) {
return '你好，$name！欢迎来到 RUNOOB。';
}

// 没有返回值的函数使用 void
void printWelcome() {
print('=== RUNOOB Dart 教程 ===');
}

void main() {
printWelcome();

// 调用函数并接收返回值
String message = greet('小明');
print(message);
}

```

=== RUNOOB Dart 教程 ===
你好，小明！欢迎来到 RUNOOB。

```

#### 返回值

每个函数都有一个返回值。

如果没有显式 return，函数隐式返回 null（但空安全下这会导致类型不匹配，所以通常用 void 表示无返回值）。

### 实例

// 返回 int 类型的函数
int add(int a, int b) {
return a + b;
}

// 使用箭头语法简写单表达式函数
int multiply(int a, int b) => a * b;

// void 表示没有有意义的返回值
void log(String msg) {
print('[RUNOOB 日志] $msg');
}

void main() {
print('3 + 5 = ${add(3, 5)}');
print('3 × 5 = ${multiply(3, 5)}');
log('函数学习完成');
}

```

3 + 5 = 8
3 × 5 = 15
[RUNOOB 日志] 函数学习完成

```

### 命名参数与可选参数

Dart 的参数分为两种：必选参数（required）和可选参数。

可选参数又分为命名参数（named）和位置参数（positional）。

#### 命名参数（Named Parameters）

命名参数用花括号 {} 包裹，调用时通过参数名传递，顺序可以任意。

### 实例

// 花括号 {} 中的是命名参数，默认是可选的
// required 关键字标记为必填
String createUser({
required String name, // 必填的命名参数
int age = 0, // 可选，默认值为 0
String? email, // 可选，可以为 null
bool isVip = false, // 可选，默认值为 false
}) {
var info = '用户名: $name, 年龄: $age';
if (email != null) {
info += ', 邮箱: $email';
}
if (isVip) {
info += ' [VIP用户]';
}
return info;
}

void main() {
// 命名参数：通过参数名传递，顺序无关
print(createUser(name: 'runoob', age: 10));
print(createUser(name: 'admin', email: 'admin@runoob.com', isVip: true));
// 必填参数不能省略
// createUser(); // 错误：缺少必填参数 name
}

```

用户名: runoob, 年龄: 10
用户名: admin, 邮箱: admin@runoob.com [VIP用户]

```

#### 位置参数（Positional Parameters）

位置参数用方括号 [] 包裹，调用时按顺序传递。

### 实例

// 方括号 [] 中的是可选位置参数
String buildUrl(String host, [String path = '/', int port = 80]) {
return 'http://$host:$port$path';
}

void main() {
// 只传必填参数，可选参数使用默认值
print(buildUrl('www.runoob.com'));

// 传 2 个参数
print(buildUrl('www.runoob.com', '/dart'));

// 传 3 个参数
print(buildUrl('localhost', '/api', 8080));
}

```

http://www.runoob.com:80/
http://www.runoob.com:80/dart
http://localhost:8080/api

```

#### 命名参数 vs 位置参数：如何选择

场景推荐方式原因 参数很多（3个以上）命名参数调用时参数名自文档化，不易传错顺序 参数含义不明显命名参数true/false 这类值不说明用途，加名字更清晰 参数很少且含义明确位置参数如 add(1, 2)，简洁明了 Flutter Widget 构建命名参数Flutter 标准风格

一个函数不能同时使用位置可选参数和命名可选参数。你只能选一种。

### 默认参数值

可选参数可以指定默认值，当调用者没有传该参数时使用默认值。

### 实例

// 所有可选参数都设置了默认值
String formatMessage(
String content, {
String prefix = '[RUNOOB]',
String suffix = '',
bool uppercase = false,
}) {
var result = '$prefix $content $suffix';
return uppercase ? result.toUpperCase() : result;
}

void main() {
// 全部使用默认值
print(formatMessage('Dart 教程更新了'));

// 覆盖部分默认值
print(formatMessage('重要通知', prefix: '[公告]', suffix: '!!!'));

// 覆盖默认值 + 开启大写
print(formatMessage('error', prefix: '[错误]', uppercase: true));
}

```

[RUNOOB] Dart 教程更新了
[公告] 重要通知 !!!
[错误] ERROR

```

默认值必须是编译时常量。也就是说，默认值不能是 DateTime.now() 或某个函数的返回值（除非该函数是 const 的）。如果需要在运行时确定默认值，可以在函数体内用 ?? 处理。

### 箭头函数与匿名函数

当函数体只有一条表达式时，可以用箭头语法（=>）简写。

匿名函数是没有名字的函数，通常作为参数传递给其他函数。

### 实例

void main() {
var numbers = [1, 2, 3, 4, 5];

// 匿名函数（完整写法）
numbers.forEach((number) {
print('RUNOOB 数字: $number');
});

print('---');

// 箭头函数（单表达式简写）
var doubled = numbers.map((n) => n * 2);
print('翻倍: $doubled');

// 箭头函数在函数定义中的使用
// 当函数体只有一个 return 语句时，可以用 => 简写
int square(int x) => x * x;
print('5 的平方: ${square(5)}');

// 存储匿名函数到变量
var sayHello = (String name) => '你好，$name！';
print(sayHello('runoob'));
}

```

RUNOOB 数字: 1
RUNOOB 数字: 2
RUNOOB 数字: 3
RUNOOB 数字: 4
RUNOOB 数字: 5
---
翻倍: (2, 4, 6, 8, 10)
5 的平方: 25
你好，runoob！

```

### 高阶函数与闭包

高阶函数是指可以接收函数作为参数、或将函数作为返回值的函数。

闭包是指函数可以访问其词法作用域中的变量，即使该函数在作用域外被调用。

#### 高阶函数

### 实例

// 高阶函数：接收一个函数作为参数
List<int> filterList(
List<int> items, bool Function(int) predicate) {
return items.where(predicate).toList();
}

// 高阶函数：返回一个函数
Function makeMultiplier(int factor) {
// 返回一个闭包：记住了外部的 factor
return (int n) => n * factor;
}

void main() {
var numbers = [10, 15, 20, 25, 30];

// 传入一个匿名函数作为筛选条件
var bigNumbers = filterList(numbers, (n) => n > 18);
print('RUNOOB 大于18的数: $bigNumbers');

// 获取返回的函数
var doubleIt = makeMultiplier(2);
var tripleIt = makeMultiplier(3);

print('5 × 2 = ${doubleIt(5)}');
print('5 × 3 = ${tripleIt(5)}');
}

```

RUNOOB 大于18的数: [20, 25, 30]
5 × 2 = 10
5 × 3 = 15

```

#### 闭包（Closure）

闭包是一个函数对象，它可以访问其词法作用域中的变量，即使该函数在原始作用域之外被调用。

### 实例

// 闭包典型应用：创建计数器
Function makeCounter() {
int count = 0; // 这个变量被返回的函数"捕获"
return () {
count++;
return count;
};
}

void main() {
// counterA 和 counterB 各自拥有独立的 count
var counterA = makeCounter();
var counterB = makeCounter();

print('RUNOOB 计数器 A: ${counterA()}'); // 1
print('RUNOOB 计数器 A: ${counterA()}'); // 2
print('RUNOOB 计数器 B: ${counterB()}'); // 1（独立的计数）
print('RUNOOB 计数器 A: ${counterA()}'); // 3
}

```

RUNOOB 计数器 A: 1
RUNOOB 计数器 A: 2
RUNOOB 计数器 B: 1
RUNOOB 计数器 A: 3

```

闭包的关键在于：内部函数"记住"了外部函数的变量，即使外部函数已经执行完毕。每次调用 makeCounter() 都会创建一个全新的 count 变量和闭包，所以多个计数器互不影响。

---

## Dart 类与对象

Source: https://www.runoob.com/dart/dart-classes-and-objects.html

## Dart 类与对象

类是面向对象编程的核心概念，它是对现实世界中事物的抽象描述。

对象是类的具体实例。本章介绍 Dart 中类的定义、构造函数、成员变量与方法、getter/setter 和静态成员。

### 类的定义与实例化

类是创建对象的模板，定义了对象有哪些属性（数据）和行为（方法）。

### 实例

// 定义一个 User 类
class User {
// 成员变量（属性）
String name = '';
int age = 0;

// 成员方法（行为）
void introduce() {
print('你好，我是 $name，今年 $age 岁。');
}

bool isAdult() {
return age >= 18;
}
}

void main() {
// 实例化：用 new 关键字（Dart 2.0+ 可省略 new）
var user1 = User();
user1.name = 'runoob';
user1.age = 10;

// 访问成员变量和方法
print('用户名: ${user1.name}');
user1.introduce();
print('是否成年: ${user1.isAdult()}');

// 创建另一个对象，彼此独立
var user2 = User();
user2.name = '小明';
user2.age = 20;
user2.introduce();
}

```

用户名: runoob
你好，我是 runoob，今年 10 岁。
是否成年: false
你好，我是 小明，今年 20 岁。

```

Dart 2.0 开始，实例化时 new 关键字是可选的。社区约定是省略 new，保持代码简洁。本教程后续示例都将省略 new。

### 构造函数

构造函数是在创建对象时自动调用的特殊方法，用于初始化对象的状态。

Dart 提供了多种构造函数形式，非常灵活。

#### 默认构造函数

如果你没有定义构造函数，Dart 会自动提供一个无参的默认构造函数。

#### 普通构造函数

### 实例

class Student {
String name;
int age;

// 普通构造函数：与类同名
Student(String name, int age) {
this.name = name;
this.age = age;
}

void printInfo() {
print('学生: $name, 年龄: $age');
}
}

void main() {
var s1 = Student('runoob', 10);
var s2 = Student('小明', 15);
s1.printInfo();
s2.printInfo();
}

```

学生: runoob, 年龄: 10
学生: 小明, 年龄: 15

```

#### 语法糖构造函数

当构造函数的参数直接赋值给成员变量时，Dart 提供了简写语法。

### 实例

class Student {
String name;
int age;

// 语法糖：this.参数名 自动将参数赋值给同名成员变量
// 这种方式比手动写 this.name = name 更简洁
Student(this.name, this.age);

void printInfo() {
print('RUNOOB 学生: $name, 年龄: $age');
}
}

void main() {
var s = Student('runoob', 10);
s.printInfo();
}

```

RUNOOB 学生: runoob, 年龄: 10

```

#### 命名构造函数

一个类可以有多个命名构造函数，用于不同的初始化逻辑。

### 实例

class Point {
double x;
double y;

// 默认构造函数
Point(this.x, this.y);

// 命名构造函数：创建原点
Point.origin()
: x = 0,
y = 0;

// 命名构造函数：从单个值创建（x 和 y 相同）
Point.diagonal(double value)
: x = value,
y = value;

@override
String toString() => 'Point($x, $y)';
}

void main() {
var p1 = Point(3, 4);
var p2 = Point.origin();
var p3 = Point.diagonal(5);

print('RUNOOB 坐标: $p1, $p2, $p3');
}

```

RUNOOB 坐标: Point(3.0, 4.0), Point(0.0, 0.0), Point(5.0, 5.0)

```

#### 初始化列表

初始化列表在构造函数体执行之前运行，用于初始化 final 字段或进行参数校验。

### 实例

class Person {
final String name; // final 字段必须在构造函数中初始化
final int birthYear;
final int age;

// 初始化列表：在冒号后面、函数体前面
// 初始化列表在构造函数体执行之前运行
Person(this.name, this.birthYear)
: age = 2026 - birthYear, // 计算属性
assert(birthYear > 1900, '出生年份不合理'); // 断言校验

void printInfo() {
print('RUNOOB 用户: $name, 出生: $birthYear, 年龄: $age');
}
}

void main() {
var p = Person('runoob', 2016);
p.printInfo();

// 下面这行会触发 assert 失败（开发模式下）
// var p2 = Person('error', 1800);
}

```

RUNOOB 用户: runoob, 出生: 2016, 年龄: 10

```

assert 只在开发模式（debug mode）下生效，生产环境（release mode）中会被忽略。它用于在开发阶段尽早发现逻辑错误。

### 成员变量与方法

成员变量存储对象的状态，成员方法定义对象的行为。

### 实例

class BankAccount {
// 私有成员：以下划线开头，仅在本文件中可访问
String _accountNumber;
double _balance = 0;

// 公开成员
String ownerName;

BankAccount(this.ownerName, this._accountNumber);

// 公开方法
void deposit(double amount) {
if (amount > 0) {
_balance += amount;
print('存入 ¥$amount，当前余额: ¥$_balance');
}
}

bool withdraw(double amount) {
if (amount > 0 && amount <= _balance) {
_balance -= amount;
print('取出 ¥$amount，当前余额: ¥$_balance');
return true;
}
print('余额不足，取款失败');
return false;
}

// 私有方法：仅供类内部使用
void _logTransaction(String type, double amount) {
print('[内部日志] $type: ¥$amount');
}
}

void main() {
var account = BankAccount('RUNOOB', '6222-0000-1234');

account.deposit(1000);
account.withdraw(300);
account.withdraw(800); // 余额不足

// 外部无法访问私有成员
// print(account._balance); // 错误：_balance 是私有的
}

```

存入 ¥1000，当前余额: ¥1000
取出 ¥300，当前余额: ¥700
余额不足，取款失败

```

### Getter 与 Setter

Getter 和 Setter 是对成员变量进行读写操作的特殊方法。

它们让你在访问属性时执行额外的逻辑，同时保持属性访问的简洁语法。

### 实例

class Circle {
double _radius; // 私有变量，存储半径

Circle(this._radius);

// Getter：获取直径（计算属性）
double get diameter => _radius * 2;

// Getter：获取面积
double get area => 3.14159 * _radius * _radius;

// Getter 和 Setter 配合：对外暴露 radius 属性
double get radius => _radius;

set radius(double value) {
// Setter 中可以添加校验逻辑
if (value <= 0) {
throw ArgumentError('半径必须大于 0');
}
_radius = value;
}

// 只读 Getter：没有对应 Setter
String get info => '圆(半径: $_radius)';
}

void main() {
var circle = Circle(5);

// 使用 Getter（像访问普通属性一样）
print('RUNOOB 圆形: ${circle.info}');
print('直径: ${circle.diameter}');
print('面积: ${circle.area.toStringAsFixed(2)}');

// 使用 Setter
circle.radius = 10;
print('修改后直径: ${circle.diameter}');

// 尝试设置非法值会抛出异常
// circle.radius = -1; // 抛出 ArgumentError
}

```

RUNOOB 圆形: 圆(半径: 5.0)
直径: 10.0
面积: 78.54
修改后直径: 20.0

```

Getter 和 Setter 的优势在于：调用者不需要知道 radius 背后是简单字段还是计算属性。如果将来需要给 radius 添加校验逻辑，只需把字段改为 getter/setter，调用代码无需修改。这就是"封装"的好处。

### 静态成员

静态成员属于类本身，而不是类的某个实例。

所有实例共享同一个静态变量，静态方法不需要创建对象就能调用。

### 实例

class Counter {
// 静态变量：所有实例共享
static int totalCount = 0;

// 实例变量：每个实例独立
String label;

Counter(this.label) {
// 每次创建实例时，静态计数器加 1
totalCount++;
}

// 静态方法：通过类名调用
static void printTotal() {
print('RUNOOB 总实例数: $totalCount');
}

// 静态方法常用于工具函数
static bool isValidLabel(String label) {
return label.isNotEmpty && label.length <= 20;
}
}

void main() {
// 静态方法：直接通过类名调用，无需创建实例
print('标签是否有效: ${Counter.isValidLabel('test')}');

var c1 = Counter('计数器A');
Counter.printTotal(); // 1

var c2 = Counter('计数器B');
var c3 = Counter('计数器C');
Counter.printTotal(); // 3

// 静态变量：通过类名访问
print('最终计数: ${Counter.totalCount}');
}

```

标签是否有效: true
RUNOOB 总实例数: 1
RUNOOB 总实例数: 3
最终计数: 3

```

静态方法中不能访问实例成员（因为没有 this）。静态方法适合放工具函数和工厂逻辑，但不要过度使用——静态方法无法被子类重写，过度使用会降低代码的可测试性和灵活性。

---

## Dart 继承与多态

Source: https://www.runoob.com/dart/dart-inheritance.html

## Dart 继承与多态

继承是面向对象编程的三大特性之一，它允许一个类基于另一个类进行扩展。

本章介绍 Dart 中的 extends 单继承、方法重写、super 关键字和抽象类。

### extends 单继承

Dart 只支持单继承——每个类只能有一个直接父类。

但继承链可以很长：子类继承父类，父类再继承祖父类，层层传递。

### 实例

// 父类（基类、超类）
class Animal {
String name;

Animal(this.name);

void eat() {
print('$name 在吃东西');
}

void sleep() {
print('$name 在睡觉');
}
}

// 子类：使用 extends 继承 Animal
class Dog extends Animal {
String breed; // 品种

// 子类构造函数需要调用父类构造函数
Dog(String name, this.breed) : super(name);

// 子类新增的方法
void bark() {
print('$name（$breed）在汪汪叫！');
}
}

void main() {
var dog = Dog('旺财', '金毛');

// 调用从父类继承的方法
dog.eat();
dog.sleep();

// 调用子类自己的方法
dog.bark();

// 子类对象也是父类类型
Animal animal = dog; // Dog 是 Animal，可以向上转型
animal.eat();
}

```

旺财 在吃东西
旺财 在睡觉
旺财（金毛）在汪汪叫！
旺财 在吃东西

```

子类可以访问父类中所有非私有的成员变量和方法。

私有成员（以下划线开头）不会被继承。

Dart 中没有 public、protected、private 关键字。以下划线（_）开头的成员是库级私有（library-private），只能在同一文件中访问。这与 Java/C++ 的访问控制不同。

### 方法重写 @override

子类可以重写（override）父类的方法，提供自己的实现。

使用 @override 注解明确表示重写意图，编译器会帮你检查是否正确重写。

### 实例

class Animal {
String name;
Animal(this.name);

// 父类的方法
void makeSound() {
print('$name 发出了一些声音');
}

@override
String toString() {
return '动物: $name';
}
}

class Cat extends Animal {
Cat(String name) : super(name);

// @override 注解：告诉编译器我是在重写父类方法
@override
void makeSound() {
print('$name 喵喵叫～');
}
}

class Duck extends Animal {
Duck(String name) : super(name);

@override
void makeSound() {
print('$name 嘎嘎叫！');
}
}

void main() {
var cat = Cat('小花');
var duck = Duck('唐老鸭');

cat.makeSound();
duck.makeSound();

// 多态：同一个方法，不同对象有不同表现
List<Animal> animals = [
Cat('咪咪'),
Duck('鸭鸭'),
Cat('球球'),
];

print('--- RUNOOB 动物园 ---');
for (var animal in animals) {
animal.makeSound(); // 运行时会调用实际类型的方法
}
}

```

小花 喵喵叫～
唐老鸭 嘎嘎叫！
--- RUNOOB 动物园 ---
咪咪 喵喵叫～
鸭鸭 嘎嘎叫！
球球 喵喵叫～

```

上面这个例子展示了多态（Polymorphism）：同一个方法调用 animal.makeSound()，实际执行的行为取决于对象的真实类型。

使用 @override 注解有两个好处：一是让代码意图更清晰，二是如果父类没有同名方法（比如拼写错误），编译器会报错提醒。建议始终加 @override。

### super 关键字

super 代表父类对象，子类通过 super 来调用父类的构造函数、方法和属性。

### 实例

class Vehicle {
String brand;
int year;

Vehicle(this.brand, this.year);

void start() {
print('$brand 车辆启动了');
}

void displayInfo() {
print('品牌: $brand, 年份: $year');
}
}

class ElectricCar extends Vehicle {
int batteryLevel;

ElectricCar(String brand, int year, this.batteryLevel)
: super(brand, year); // 调用父类构造函数

@override
void start() {
// 在子类方法中调用父类的方法
super.start(); // 先执行父类的启动逻辑
print('电池电量: $batteryLevel%');
print('电动机无声启动完成');
}

@override
void displayInfo() {
super.displayInfo(); // 先显示父类的基础信息
print('电池电量: $batteryLevel%'); // 再补充子类的特有信息
}
}

void main() {
var car = ElectricCar('RUNOOB EV', 2026, 85);
car.start();
print('---');
car.displayInfo();
}

```

RUNOOB EV 车辆启动了
电池电量: 85%
电动机无声启动完成
---
品牌: RUNOOB EV, 年份: 2026
电池电量: 85%

```

#### super 的使用场景总结

场景语法说明 调用父类构造函数: super(args)必须在初始化列表中调用 调用父类方法super.methodName()可在子类重写方法中调用 访问父类属性super.propertyName访问父类的成员变量

### 抽象类 abstract

抽象类是不能被直接实例化的类，它只定义接口规范，由子类实现具体逻辑。

### 实例

// abstract 关键字声明抽象类
abstract class Shape {
// 抽象方法：没有方法体，子类必须实现
double getArea();
double getPerimeter();

// 抽象类中也可以有具体方法
void describe() {
print('这是一个形状，面积: ${getArea()}, 周长: ${getPerimeter()}');
}
}

class Rectangle extends Shape {
double width;
double height;

Rectangle(this.width, this.height);

@override
double getArea() => width * height;

@override
double getPerimeter() => 2 * (width + height);
}

class Circle extends Shape {
double radius;

Circle(this.radius);

@override
double getArea() => 3.14159 * radius * radius;

@override
double getPerimeter() => 2 * 3.14159 * radius;
}

void main() {
// Shape shape = Shape(); // 错误！抽象类不能实例化

var rect = Rectangle(10, 5);
var circle = Circle(7);

rect.describe();
circle.describe();

// 抽象类作为类型使用
List<Shape> shapes = [
Rectangle(3, 4),
Circle(5),
Rectangle(6, 2),
];

print('--- RUNOOB 形状统计 ---');
double totalArea = 0;
for (var shape in shapes) {
totalArea += shape.getArea();
}
print('总面积: ${totalArea.toStringAsFixed(2)}');
}

```

这是一个形状，面积: 50.0, 周长: 30.0
这是一个形状，面积: 153.93791, 周长: 43.98226
--- RUNOOB 形状统计 ---
总面积: 108.54

```

抽象类的核心价值是"定义契约"。它告诉所有子类：你必须实现这些方法，否则编译不通过。这让团队协作更安全——实现者不会遗漏方法，调用者确信方法一定存在。

---

## Dart 接口与 Mixin

Source: https://www.runoob.com/dart/dart-interfaces-and-mixins.html

## Dart 接口与 Mixin

虽然 Dart 只支持单继承，但通过接口和 Mixin 可以灵活地实现多重代码复用。

本章介绍如何使用 implements 实现接口、mixin 的代码复用以及 with 关键字。

### implements 实现接口

在 Dart 中，每个类都隐式定义了一个接口。

使用 implements 关键字可以让一个类实现另一个类的接口，且可以实现多个。

### 实例

// 定义一个"可打印"的接口
class Printable {
// 接口中的方法没有默认实现
// 实现者必须重写所有成员
void printContent() {
// 这个方法体在 implements 时会被忽略
}
}

// 定义一个"可保存"的接口
class Savable {
void save() {}
}

// 使用 implements 实现多个接口
class Document implements Printable, Savable {
String title;
String content;

Document(this.title, this.content);

// 必须实现 Printable 接口要求的所有方法
@override
void printContent() {
print('--- 文档: $title ---');
print(content);
print('--- RUNOOB 打印结束 ---');
}

// 必须实现 Savable 接口要求的所有方法
@override
void save() {
print('文档 "$title" 已保存到磁盘');
}
}

void main() {
var doc = Document('Dart 教程', '这是一份 Dart 入门指南');

// 多态：Document 既是 Printable 也是 Savable
Printable printable = doc;
printable.printContent();

Savable savable = doc;
savable.save();
}

```

--- 文档: Dart 教程 ---
这是一份 Dart 入门指南
--- RUNOOB 打印结束 ---
文档 "Dart 教程" 已保存到磁盘

```

implements 和 extends 的关键区别：extends 继承父类的所有实现（方法体），而 implements 只继承接口签名，你必须重新实现所有方法。此外，extends 只能继承一个类，implements 可以实现多个接口。

#### extends vs implements 对比

特性extendsimplements 数量限制只能继承一个可以实现多个 方法实现继承父类的实现必须自己实现所有方法 构造函数可以调用 super()不继承构造函数 使用场景"是一个"的关系"能做什么"的契约

### 多接口实现

Dart 可以实现多个接口，这是突破单继承限制的主要方式。

### 实例

// 定义多个行为接口
class Flyable {
void fly() {}
}

class Swimmable {
void swim() {}
}

class Walkable {
void walk() {}
}

// 一只鸭子可以飞、游、走
class Duck implements Flyable, Swimmable, Walkable {
String name;

Duck(this.name);

@override
void fly() {
print('$name 在天空中飞翔');
}

@override
void swim() {
print('$name 在水面上游泳');
}

@override
void walk() {
print('$name 在陆地上摇摇晃晃地走');
}
}

// 一条鱼只能游
class Fish implements Swimmable {
@override
void swim() {
print('鱼儿在水中游动');
}
}

void main() {
var duck = Duck('RUNOOB 鸭');

// 鸭子可以扮演多种角色
(duck as Flyable).fly();
(duck as Swimmable).swim();
(duck as Walkable).walk();

// 类型检查
print('鸭子能飞吗? ${duck is Flyable}');
print('鸭子能游吗? ${duck is Swimmable}');

var fish = Fish();
print('鱼能飞吗? ${fish is Flyable}');
}

```

RUNOOB 鸭 在天空中飞翔
RUNOOB 鸭 在水面上游泳
RUNOOB 鸭 在陆地上摇摇晃晃地走
鸭子能飞吗? true
鸭子能游吗? true
鱼能飞吗? false

```

### Mixin 代码复用

Mixin 是一种在多个类之间复用代码的机制，它解决了"多个类需要共享相同方法但又不适合继承"的问题。

使用 mixin 关键字定义一个 Mixin，使用 with 关键字将其混入到类中。

### 实例

// 使用 mixin 关键字定义
mixin Logger {
// Mixin 中的方法可以被多个类复用
void log(String message) {
print('[${DateTime.now()}] $message');
}

void logError(String message) {
print('[ERROR ${DateTime.now()}] $message');
}
}

mixin TimestampFormatter {
String formatTimestamp(DateTime dt) {
return '${dt.year}-${dt.month.toString().padLeft(2, '0')}'
'-${dt.day.toString().padLeft(2, '0')}';
}
}

// 使用 with 关键字混入 Mixin
class UserService with Logger, TimestampFormatter {
void createUser(String name) {
log('创建用户: $name');
var now = DateTime.now();
print('创建时间: ${formatTimestamp(now)}');
}
}

class OrderService with Logger {
void createOrder(String orderId) {
log('创建订单: $orderId');
}

void cancelOrder(String orderId) {
logError('取消订单失败: $orderId');
}
}

void main() {
var userService = UserService();
userService.createUser('runoob');

var orderService = OrderService();
orderService.createOrder('ORD-001');
orderService.cancelOrder('ORD-001');
}

```

[2026-06-12 10:30:00.000] 创建用户: runoob
创建时间: 2026-06-12
[2026-06-12 10:30:00.001] 创建订单: ORD-001
[ERROR 2026-06-12 10:30:00.001] 取消订单失败: ORD-001

```

UserService 和 OrderService 来自不同的业务领域，不适合用继承来共享 Logger。

通过 Mixin，它们都能获得日志功能，而且互不影响。

Mixin 不能有构造函数，也不能被实例化。它的定位很纯粹：提供可复用的方法，不涉及状态初始化。

#### Mixin 的限制条件

Mixin 可以使用 on 关键字限制它只能被特定类型的类使用。

### 实例

// 基础类
class Animal {
String name;
Animal(this.name);
}

// 这个 Mixin 只能混入到 Animal 及其子类中
mixin FlyableMixin on Animal {
void fly() {
print('$name 飞起来了！');
}
}

// Bird 是 Animal，可以用 FlyableMixin
class Bird extends Animal with FlyableMixin {
Bird(String name) : super(name);
}

// 下面这行会报错：Car 不是 Animal，不能用 FlyableMixin
// class Car with FlyableMixin {} // 错误！

void main() {
var bird = Bird('RUNOOB 小鸟');
bird.fly();

// FlyableMixin 中的方法可以访问 Animal 的属性 name
print('小鸟的名字: ${bird.name}');
}

```

RUNOOB 小鸟 飞起来了！
小鸟的名字: RUNOOB 小鸟

```

#### class、Mixin、extends、implements 的完整组合

Dart 中一个类可以同时使用 extends、with 和 implements：

### 实例

mixin Jumpable {
void jump() => print('跳起来！');
}

mixin Runnable {
void run() => print('跑步前进！');
}

class Animal {
void breathe() => print('呼吸...');
}

// 继承 Animal，混入 Jumpable 和 Runnable，实现 Comparable
class Athlete extends Animal
with Jumpable, Runnable
implements Comparable<Athlete> {
String name;
int score;

Athlete(this.name, this.score);

@override
int compareTo(Athlete other) => score.compareTo(other.score);

void showSkills() {
breathe();
run();
jump();
}
}

void main() {
var a1 = Athlete('RUNOOB 选手A', 95);
var a2 = Athlete('选手B', 88);

a1.showSkills();
print('${a1.name} vs ${a2.name}: ${a1.compareTo(a2) > 0 ? "A 胜" : "B 胜"}');
}

```

呼吸...
跑步前进！
跳起来！
RUNOOB 选手A vs 选手B: A 胜

```

完整声明顺序：class 子类 extends 父类 with Mixin1, Mixin2 implements 接口1, 接口2

---

## Dart 泛型

Source: https://www.runoob.com/dart/dart-generics.html

## Dart 泛型

泛型（Generics）让你在编写代码时不指定具体类型，而是用类型参数占位，在使用时再确定类型。

泛型的核心价值是类型安全和代码复用——一份代码适配多种类型，同时编译器帮你检查类型错误。

### 为什么需要泛型

先看一个没有泛型时的痛点。

### 实例

没有泛型的写法：

// 没有泛型：只能存放 Object，取出时需要手动转型
class IntBox {
int value;
IntBox(this.value);
}

class StringBox {
String value;
StringBox(this.value);
}

// 每新增一种类型就要写一个新的 Box 类，代码大量重复

void main() {
var intBox = IntBox(42);
var stringBox = StringBox('RUNOOB');

print('整数: ${intBox.value}');
print('字符串: ${stringBox.value}');
}

有没有办法写一个 Box 类，既能装 int、又能装 String，还能在编译时保证类型安全？

这就是泛型要解决的问题。

### 泛型类

泛型类在类名后用尖括号 <> 声明类型参数，类内部使用这个类型参数。

### 实例

// <T> 是类型参数，T 是一个占位符，代表"某种类型"
// 约定俗成的命名：T（Type）、E（Element）、K（Key）、V（Value）
class Box<T> {
T value;

Box(this.value);

T getValue() {
return value;
}

void setValue(T newValue) {
value = newValue;
}

void printValue() {
print('Box 中的值: $value (类型: ${value.runtimeType})');
}
}

void main() {
// 使用泛型类时指定具体类型
var intBox = Box<int>(42);
var stringBox = Box<String>('RUNOOB Dart 教程');
var doubleBox = Box<double>(3.14);

intBox.printValue();
stringBox.printValue();
doubleBox.printValue();

// 类型安全：下面的代码会在编译时报错
// intBox.setValue('hello'); // 错误：String 不能赋值给 int

// 类型推断：Dart 可以根据构造函数参数推断类型
var autoBox = Box('自动推断为 String');
print('推断类型: ${autoBox.value.runtimeType}');
}

```

Box 中的值: 42 (类型: int)
Box 中的值: RUNOOB Dart 教程 (类型: String)
Box 中的值: 3.14 (类型: double)
推断类型: String

```

#### 多个类型参数

泛型类可以有多个类型参数。

### 实例

// K 表示键的类型，V 表示值的类型
class Pair<K, V> {
K key;
V value;

Pair(this.key, this.value);

@override
String toString() => 'Pair($key: $value)';
}

void main() {
var pair1 = Pair<String, int>('runoob', 10);
var pair2 = Pair('score', 95.5); // 类型推断

print(pair1);
print(pair2);
}

```

Pair(runoob: 10)
Pair(score: 95.5)

```

### 泛型方法

除了泛型类，函数和方法也可以使用泛型。

### 实例

// 泛型函数：返回类型和参数类型一致
// 在返回值类型前声明 <T>
T firstElement<T>(List<T> list) {
if (list.isEmpty) {
throw ArgumentError('列表不能为空');
}
return list[0];
}

// 泛型函数：交换两个值
void swap<T>(List<T> list, int i, int j) {
T temp = list[i];
list[i] = list[j];
list[j] = temp;
}

void main() {
// 调用时 Dart 会自动推断类型
var names = ['runoob', 'Dart', 'Flutter'];
print('第一个名字: ${firstElement(names)}');

var scores = [95, 88, 72];
print('第一个分数: ${firstElement(scores)}');

// 交换元素
var items = ['A', 'B', 'C'];
print('交换前: $items');
swap(items, 0, 2);
print('交换后: $items');
}

```

第一个名字: runoob
第一个分数: 95
交换前: [A, B, C]
交换后: [C, B, A]

```

### 类型约束 extends

有时候你希望类型参数只能是某些类型的子类型。

使用 extends 关键字可以对泛型参数施加约束。

### 实例

// 定义一个数值计算器
// <T extends num> 约束 T 必须是 num 或其子类（int、double）
class Calculator<T extends num> {
T a;
T b;

Calculator(this.a, this.b);

// 因为约束了 T extends num，所以可以安全地使用 num 的方法
double add() => (a + b).toDouble();
double subtract() => (a - b).toDouble();
double multiply() => (a * b).toDouble();
double divide() => a / b;

void printOperations() {
print('RUNOOB 计算器: $a 和 $b');
print(' 加: ${add()}');
print(' 减: ${subtract()}');
print(' 乘: ${multiply()}');
print(' 除: ${divide().toStringAsFixed(2)}');
}
}

void main() {
var intCalc = Calculator<int>(10, 3);
intCalc.printOperations();

var doubleCalc = Calculator<double>(3.5, 1.5);
doubleCalc.printOperations();

// 下面这行会报错：String 不是 num 的子类
// var strCalc = Calculator<String>('a', 'b'); // 错误！
}

```

RUNOOB 计算器: 10 和 3
加: 13.0
减: 7.0
乘: 30.0
除: 3.33
RUNOOB 计算器: 3.5 和 1.5
加: 5.0
减: 2.0
乘: 5.25
除: 2.33

```

类型约束让你既能保持泛型的灵活性，又能使用约束类型的方法。比如上面约束 T extends num，就可以在方法体内使用 +、-、*、/ 等运算符，因为 num 类型支持这些操作。

### 泛型集合的使用

Dart 的核心集合类（List、Set、Map）都是泛型的。

使用泛型集合可以获得编译期的类型检查。

### 实例

void main() {
// 指定元素类型为 String
List<String> names = ['runoob', 'Dart', 'Flutter'];
// names.add(42); // 错误：int 不能添加到 List<String>

// 指定键值类型
Map<String, int> scores = {
'runoob': 95,
'Alice': 87,
};
// scores['Bob'] = '优秀'; // 错误：String 不能赋值给 int

// 指定 Set 元素类型
Set<double> prices = {9.99, 19.99, 29.99};

// 泛型集合的函数式操作
List<int> numbers = [1, 2, 3, 4, 5, 6];

// 类型安全的过滤和映射
List<int> evenNumbers = numbers.where((n) => n % 2 == 0).toList();
List<String> labels = numbers.map((n) => 'RUNOOB-$n').toList();

print('偶数: $evenNumbers');
print('标签: $labels');

// 使用 fold 进行类型安全的累积计算
int sum = numbers.fold<int>(0, (prev, n) => prev + n);
print('总和: $sum');
}

```

偶数: [2, 4, 6]
标签: [RUNOOB-1, RUNOOB-2, RUNOOB-3, RUNOOB-4, RUNOOB-5, RUNOOB-6]
总和: 21

```

### 实际应用：泛型仓库模式

泛型在实际开发中最经典的应用之一是"仓库模式"（Repository Pattern）。

### 实例

// 定义一个通用仓库接口
abstract class Repository<T> {
void add(T item);
void remove(T item);
T? findById(String id);
List<T> getAll();
}

// 用户模型
class User {
String id;
String name;

User(this.id, this.name);

@override
String toString() => 'User($id, $name)';
}

// 商品模型
class Product {
String id;
String name;
double price;

Product(this.id, this.name, this.price);

@override
String toString() => 'Product($name, ¥$price)';
}

// 具体实现：用户仓库
class UserRepository implements Repository<User> {
final List<User> _users = [];

@override
void add(User user) => _users.add(user);

@override
void remove(User user) => _users.remove(user);

@override
User? findById(String id) {
try {
return _users.firstWhere((u) => u.id == id);
} catch (_) {
return null;
}
}

@override
List<User> getAll() => List.unmodifiable(_users);
}

// 具体实现：商品仓库（同样的接口，不同的类型）
class ProductRepository implements Repository<Product> {
final List<Product> _products = [];

@override
void add(Product product) => _products.add(product);

@override
void remove(Product product) => _products.remove(product);

@override
Product? findById(String id) {
try {
return _products.firstWhere((p) => p.id == id);
} catch (_) {
return null;
}
}

@override
List<Product> getAll() => List.unmodifiable(_products);
}

void main() {
var userRepo = UserRepository();
userRepo.add(User('u1', 'runoob'));
userRepo.add(User('u2', 'Alice'));

var productRepo = ProductRepository();
productRepo.add(Product('p1', 'Dart 教程', 29.99));
productRepo.add(Product('p2', 'Flutter 教程', 49.99));

print('用户列表: ${userRepo.getAll()}');
print('商品列表: ${productRepo.getAll()}');
print('查找用户 u1: ${userRepo.findById('u1')}');
print('查找商品 p2: ${productRepo.findById('p2')}');
}

```

用户列表: [User(u1, runoob), User(u2, Alice)]
商品列表: [Product(Dart 教程, ¥29.99), Product(Flutter 教程, ¥49.99)]
查找用户 u1: User(u1, runoob)
查找商品 p2: Product(Flutter 教程, ¥49.99)

```

泛型仓库模式让你用统一的接口处理不同类型的数据，既保证了类型安全，又避免了重复代码。

---

## Dart 枚举与符号

Source: https://www.runoob.com/dart/dart-enums-and-symbols.html

## Dart 枚举与符号

枚举（enum）用于定义一组有名字的常量值，让代码更具可读性。

本章介绍 Dart 枚举的定义方式、增强枚举（带方法和属性）以及 Symbol 和 Rune 类型。

### enum 枚举定义

枚举是一种特殊的类，用于表示固定数量的常量值。

最典型的场景是表示状态、方向、颜色等有限选项。

### 实例

// 定义枚举
enum Status {
pending, // 待处理
approved, // 已批准
rejected, // 已拒绝
cancelled, // 已取消
}

void main() {
// 使用枚举值
Status currentStatus = Status.pending;
print('当前状态: $currentStatus');

// 通过 .index 获取索引（从 0 开始）
print('索引值: ${currentStatus.index}'); // 0

// 通过 .values 获取所有枚举值
print('所有状态: ${Status.values}');

// 遍历所有枚举值
for (var status in Status.values) {
print('RUNOOB 状态: $status, 索引: ${status.index}');
}

// 通过字符串获取枚举值
Status parsed = Status.values.byName('approved');
print('解析出的状态: $parsed');

// 在 switch 中使用枚举
switch (currentStatus) {
case Status.pending:
print('订单待处理');
break;
case Status.approved:
print('订单已批准');
break;
case Status.rejected:
print('订单被拒绝');
break;
case Status.cancelled:
print('订单已取消');
break;
}
}

```

当前状态: Status.pending
索引值: 0
所有状态: [Status.pending, Status.approved, Status.rejected, Status.cancelled]
RUNOOB 状态: Status.pending, 索引: 0
RUNOOB 状态: Status.approved, 索引: 1
RUNOOB 状态: Status.rejected, 索引: 2
RUNOOB 状态: Status.cancelled, 索引: 3
解析出的状态: Status.approved
订单待处理

```

枚举和 switch 是绝配。当你在 switch 中处理所有枚举值时，Dart 编译器会检查你是否覆盖了所有情况。这能避免遗漏分支导致的 bug。

### 增强枚举（带方法和属性）

Dart 3.0 引入了增强枚举（Enhanced Enums），允许枚举拥有字段、方法和构造函数。

### 实例

// 增强枚举：可以拥有成员变量、构造函数和方法
enum HttpStatus {
// 枚举值通过构造函数传递参数
ok(200, '请求成功'),
created(201, '资源已创建'),
badRequest(400, '请求错误'),
unauthorized(401, '未授权'),
notFound(404, '资源未找到'),
serverError(500, '服务器内部错误');

// 成员变量：每个枚举值都有 code 和 message
final int code;
final String message;

// 构造函数：必须是 const
const HttpStatus(this.code, this.message);

// 枚举方法
bool get isSuccess => code >= 200 && code < 300;
bool get isClientError => code >= 400 && code < 500;
bool get isServerError => code >= 500;

// 静态方法：根据状态码查找枚举
static HttpStatus? fromCode(int code) {
try {
return HttpStatus.values.firstWhere((s) => s.code == code);
} catch (_) {
return null;
}
}
}

void main() {
var status = HttpStatus.notFound;

print('RUNOOB HTTP 状态: $status');
print('状态码: ${status.code}');
print('消息: ${status.message}');
print('是否成功: ${status.isSuccess}');
print('是否客户端错误: ${status.isClientError}');

// 使用静态方法查找
var found = HttpStatus.fromCode(201);
if (found != null) {
print('找到状态: ${found.message}');
}

var unknown = HttpStatus.fromCode(999);
print('未知状态码: $unknown');
}

```

RUNOOB HTTP 状态: HttpStatus.notFound
状态码: 404
消息: 资源未找到
是否成功: false
是否客户端错误: true
找到状态: 资源已创建
未知状态码: null

```

增强枚举让枚举从"命名常量集合"升级为"有行为的类型"，可以封装与该枚举相关的逻辑。

增强枚举的构造函数必须是 const，成员变量必须是 final。这是因为枚举值本身就是编译时常量。

### Symbol 类型

Symbol 代表 Dart 程序中的标识符（变量名、函数名等）。

它在日常开发中不常用，主要在反射（Reflection）和代码生成场景中使用。

### 实例

void main() {
// 使用 # 前缀创建 Symbol
Symbol sym1 = #runoob;
Symbol sym2 = #helloWorld;

print('Symbol 1: $sym1'); // Symbol("runoob")
print('Symbol 2: $sym2'); // Symbol("helloWorld")

// Symbol 比较
Symbol sym3 = #runoob;
print('sym1 == sym3: ${sym1 == sym3}'); // true

// Symbol 可以用作 Map 的键
var metadata = {
#author: 'RUNOOB',
#version: '1.0.0',
#description: 'Dart 教程',
};
print('作者: ${metadata[#author]}');
}

```

Symbol 1: Symbol("runoob")
Symbol 2: Symbol("helloWorld")
sym1 == sym3: true
作者: RUNOOB

```

Symbol 在 Dart 编译后通常会被压缩（minified），所以不要依赖 Symbol("runoob").toString() 的结果来做逻辑判断。Symbol 的正确用法是通过 dart:mirrors 库进行反射操作。

### Rune 与 Unicode

Rune 是 Dart 中表示 Unicode 码点的类型。

在 Dart 中，字符串是 UTF-16 编码的序列。

对于基本多语言平面（BMP）之外的字符（如 emoji），需要使用 Rune 来处理。

### 实例

void main() {
// 获取字符串的 Unicode 码点
String text = 'RUNOOB';
print('字符: $text');
print('码点列表: ${text.runes.toList()}');

// 遍历每个字符和它的码点
for (var char in text.runes) {
print('${String.fromCharCode(char)} -> U+${char.toRadixString(16).toUpperCase().padLeft(4, '0')}');
}

// 从码点创建字符
var heart = String.fromCharCode(0x2764);
print('心形: $heart');

// 处理多码点字符（如某些 emoji）
var smile = '\u{1F600}'; // 使用 \u{} 语法表示超出 BMP 的字符
print('笑脸: $smile');

// 获取多码点字符的码点
print('笑脸码点: ${smile.runes.map((r) => 'U+${r.toRadixString(16).toUpperCase()}').toList()}');
}

```

字符: RUNOOB
码点列表: [82, 85, 78, 79, 79, 66]
R -> U+0052
U -> U+0055
N -> U+004E
O -> U+004F
O -> U+004F
B -> U+0042
心形: &#x2764;
笑脸: &#x1f600;
笑脸码点: [U+1F600]

```

大部分日常开发中你不需要直接操作 Rune，字符串本身已经能很好地处理 Unicode。

Rune 主要在需要精确操作字符码点时使用，比如文本编辑器、字体渲染等底层场景。

---

## Dart 异常处理

Source: https://www.runoob.com/dart/dart-exception-handling.html

## Dart 异常处理

异常是程序运行时发生的错误事件。

良好的异常处理能让程序在遇到错误时优雅地降级，而不是直接崩溃。

本章介绍 Dart 中的 try/catch/finally、throw 抛出异常、on 捕获指定类型以及 rethrow 重新抛出。

### try / catch / finally 基础

try 块中放置可能抛出异常的代码，catch 块捕获并处理异常，finally 块无论是否异常都会执行。

### 实例

void main() {
// 基本的 try-catch
try {
int result = 10 ~/ 0; // 除零操作会抛出异常
print('结果: $result'); // 这行不会执行
} catch (e) {
// e 是异常对象
print('捕获到异常: $e');
}

// 获取异常和堆栈信息
try {
List<int> numbers = [1, 2, 3];
print(numbers[10]); // 访问越界索引
} catch (e, stackTrace) {
// e 是异常，stackTrace 是调用堆栈
print('RUNOOB 异常: $e');
print('堆栈信息:\n$stackTrace');
}

// finally：无论是否异常都会执行
try {
print('尝试执行操作...');
int result = 5 ~/ 0;
print('这行不会执行: $result');
} catch (e) {
print('出错: $e');
} finally {
// finally 块中的代码始终执行，常用于清理资源
print('清理资源完成（finally 总是执行）');
}
}

```

捕获到异常: IntegerDivisionByZeroException
RUNOOB 异常: RangeError (index): Invalid value: Not in inclusive range 0..2: 10
堆栈信息:
#0 List.[] (dart:core-patch/growable_array.dart:264)
#1 main (file:///...)
...
尝试执行操作...
出错: IntegerDivisionByZeroException
清理资源完成（finally 总是执行）

```

finally 块最适合用于资源清理，比如关闭文件、释放数据库连接等。即使 try 或 catch 中有 return 语句，finally 仍然会执行。

### throw 抛出异常

使用 throw 关键字可以主动抛出异常。

Dart 允许抛出任何对象作为异常（不仅仅是 Exception 的子类），但最佳实践是抛出 Exception 或 Error 的子类。

### 实例

// 自定义异常类
class InvalidAgeException implements Exception {
final int age;
final String message;

InvalidAgeException(this.age)
: message = '无效的年龄: $age（年龄必须在 0 到 150 之间）';

@override
String toString() => 'InvalidAgeException: $message';
}

// 使用自定义异常的函数
void validateAge(int age) {
if (age < 0) {
throw InvalidAgeException(age);
}
if (age > 150) {
throw InvalidAgeException(age);
}
print('年龄 $age 验证通过');
}

// 抛出内置异常
int divide(int a, int b) {
if (b == 0) {
throw ArgumentError('除数不能为 0'); // 使用内置异常
}
return a ~/ b;
}

void main() {
// 测试自定义异常
try {
validateAge(-5);
} catch (e) {
print(e);
}

try {
validateAge(200);
} catch (e) {
print(e);
}

validateAge(25); // 正常情况

// 测试抛出内置异常
try {
divide(10, 0);
} catch (e) {
print('RUNOOB 错误: $e');
}
}

```

InvalidAgeException: 无效的年龄: -5（年龄必须在 0 到 150 之间）
InvalidAgeException: 无效的年龄: 200（年龄必须在 0 到 150 之间）
年龄 25 验证通过
RUNOOB 错误: Invalid argument(s): 除数不能为 0

```

Exception 和 Error 的区别：Exception 是可预期的、程序可以处理的错误（如网络超时、文件不存在）；Error 是不可预期的、通常表示程序 bug 的问题（如类型错误、空指针）。你的代码应该捕获 Exception 而非 Error。

### on 捕获指定类型

on 关键字可以按异常类型进行捕获，让错误处理更精细。

### 实例

void main() {
// 按类型捕获不同的异常
try {
// 模拟不同类型的异常
int result = performOperation('divide_by_zero');
print('结果: $result');
} on IntegerDivisionByZeroException {
// 只捕获除零异常
print('RUNOOB 错误: 不能除以零');
} on FormatException catch (e) {
// 捕获格式异常，同时获取异常信息
print('RUNOOB 格式错误: ${e.message}');
} on RangeError {
// 只捕获范围错误
print('RUNOOB 错误: 索引越界');
} catch (e) {
// 兜底：捕获所有其他异常
print('未知错误: $e');
} finally {
print('操作结束\n');
}

// 按顺序匹配的示例
print('--- 测试不同的操作 ---');
for (var op in ['divide_by_zero', 'parse_error', 'out_of_range', 'normal']) {
try {
performOperation(op);
} on IntegerDivisionByZeroException {
print('$op -> 除零异常');
} on FormatException {
print('$op -> 格式异常');
} on RangeError {
print('$op -> 范围异常');
} catch (e) {
print('$op -> 其他异常: $e');
}
}
}

int performOperation(String operation) {
switch (operation) {
case 'divide_by_zero':
return 10 ~/ 0; // 抛出 IntegerDivisionByZeroException
case 'parse_error':
int.parse('not_a_number'); // 抛出 FormatException
return 0;
case 'out_of_range':
var list = [1, 2, 3];
return list[100]; // 抛出 RangeError
case 'normal':
return 42;
default:
throw Exception('未知操作');
}
}

```

RUNOOB 错误: 不能除以零
操作结束

--- 测试不同的操作 ---
divide_by_zero -> 除零异常
parse_error -> 格式异常
out_of_range -> 范围异常

```

on 和 catch 的排列顺序很重要：Dart 会按照从上到下的顺序匹配异常类型。更具体的异常类型应该放在前面，更通用的放在后面。如果把 catch (e) 放在最前面，后面的 on 就永远不会被执行。

### rethrow 重新抛出

有时你需要在捕获异常后做一些处理（如记录日志），然后让异常继续向上传播。

rethrow 关键字用于重新抛出当前捕获的异常，保留原始堆栈信息。

### 实例

// 模拟一个多层调用的场景
void databaseOperation() {
print(' 正在连接数据库...');
throw Exception('RUNOOB 数据库连接失败');
}

void serviceLayer() {
try {
print(' 服务层: 调用数据库操作');
databaseOperation();
} catch (e) {
// 先记录日志，然后重新抛出
print(' [日志] 数据库操作异常: $e');
rethrow; // 保留原始异常和堆栈，继续向上传播
}
}

void controllerLayer() {
try {
print('控制器层: 处理请求');
serviceLayer();
} catch (e, stackTrace) {
// 最外层捕获并处理
print('控制器层捕获到异常');
print('错误信息: $e');
print('给用户返回: "服务暂时不可用，请稍后重试"');
}
}

void main() {
controllerLayer();
print('\n--- rethrow vs throw e 的区别 ---');

// 演示 rethrow 和 throw e 的区别
try {
try {
throw FormatException('原始异常');
} catch (e) {
// throw e：会重置堆栈信息，丢失原始调用位置
print('使用 throw e（丢失原始堆栈）...');
throw e; // 堆栈从这里开始
}
} catch (e, stack) {
print('捕获: $e');
print('堆栈指向这里（不是原始位置）:\n$stack');
}
}

```

控制器层: 处理请求
服务层: 调用数据库操作
正在连接数据库...
[日志] 数据库操作异常: Exception: RUNOOB 数据库连接失败
控制器层捕获到异常
错误信息: Exception: RUNOOB 数据库连接失败
给用户返回: "服务暂时不可用，请稍后重试"

--- rethrow vs throw e 的区别 ---
使用 throw e（丢失原始堆栈）...
捕获: FormatException: 原始异常
堆栈指向这里（不是原始位置）:
...

```

rethrow 和 throw e 的区别很关键：rethrow 保留原始异常的全部堆栈信息，方便定位问题根源；throw e 会重置堆栈，从当前位置重新开始。在记录日志后需要继续传播异常时，请使用 rethrow。

### 异常处理的最佳实践

实践说明 只捕获能处理的异常不要为了"不报错"而盲目 catch 所有异常。如果不知道如何处理，让它向上传播 按类型捕获使用 on 指定异常类型，避免用一个 catch 吞掉所有异常 finally 清理资源文件、网络连接等资源在 finally 中关闭 不要吞掉异常捕获了异常但什么都不做（空 catch 块）是危险的，至少记录日志 使用 rethrow需要记录日志后继续传播异常时，用 rethrow 而非 throw e 异常消息要具体throw Exception('数据库连接失败: 192.168.1.1:3306') 比 throw Exception('出错了') 更有用

---

## Dart 包与库管理

Source: https://www.runoob.com/dart/dart-packages-and-libraries.html

## Dart 包与库管理

当项目规模增长时，将代码拆分为多个文件和模块是必不可少的。

本章介绍 Dart 的包管理系统 pub、依赖配置文件 pubspec.yaml、import/export 语法以及如何创建自定义库。

### pubspec.yaml 配置

pubspec.yaml 是每个 Dart 项目的核心配置文件，声明了项目的元数据和依赖。

一个典型的 pubspec.yaml 文件结构如下：

### 实例

# 文件路径：pubspec.yaml
name: my_dart_app # 项目名称（必填，小写+下划线）
description: 一个 Dart 示例项目 # 项目描述
version: 1.0.0 # 版本号
# publish_to: none # 如果不想发布到 pub.dev，取消这行注释

environment:
sdk: '>=3.0.0 <4.0.0' # Dart SDK 版本范围

dependencies:
# 项目运行时依赖的第三方包
http: ^1.1.0 # HTTP 客户端
path: ^1.8.0 # 路径工具

dev_dependencies:
# 仅开发时需要的依赖（测试、代码检查等）
test: ^1.24.0 # 测试框架
lints: ^2.0.0 # 官方 lint 规则

# 可选：可执行文件入口
# executables:
# my_app: main

版本号的写法说明：

写法含义示例 ^1.1.0兼容 1.1.0 到 2.0.0（不含）最常用，推荐使用 1.1.0精确版本不太灵活 >=1.1.0 <1.5.0版本范围需要精确控制时使用 any任意版本不推荐

^ 符号（caret）是 Dart 的默认版本约束方式。^1.1.0 等价于 >=1.1.0 <2.0.0。这意味着可以自动升级次版本和补丁版本，但不能升级主版本（主版本升级可能包含破坏性变更）。

#### 安装依赖

```

$ dart pub get # 安装依赖
$ dart pub upgrade # 升级依赖到最新兼容版本
$ dart pub outdated # 查看哪些依赖有更新

```

### pub.dev 查找依赖

pub.dev 是 Dart 和 Flutter 的官方包仓库，类似于 npm（JavaScript）或 PyPI（Python）。

你可以在 https://pub.dev 上搜索和浏览数以万计的 Dart 包。

常用的第三方包举例：

包名用途说明 httpHTTP 请求官方维护的 HTTP 客户端 path路径操作跨平台的路径处理工具 test单元测试官方测试框架 json_serializableJSON 序列化自动生成 JSON 转换代码 dioHTTP 客户端比 http 更强大的第三方 HTTP 库 riverpod状态管理Flutter 项目的热门状态管理方案

### 实例

添加 http 包并使用：

// 首先在 pubspec.yaml 的 dependencies 中添加：
// http: ^1.1.0
// 然后运行：dart pub get

import 'package:http/http.dart' as http;

void main() async {
// 发送一个 GET 请求
var url = Uri.parse('https://www.runoob.com');
var response = await http.get(url);

print('状态码: ${response.statusCode}');
print('响应体长度: ${response.body.length} 字符');
}

### import 导入语法

import 用于在一个 Dart 文件中引入其他库的代码。

Dart 支持多种导入方式，每种适用于不同的场景。

#### 导入 Dart 内置库

### 实例

// dart: 前缀表示 Dart 内置的核心库
import 'dart:math'; // 数学库（随机数、三角函数等）
import 'dart:convert'; // 编码转换库（JSON、Base64 等）
import 'dart:io'; // I/O 库（文件、网络等）

void main() {
// 使用 dart:math 中的函数
print('圆周率: $pi');
print('2 的 10 次方: ${pow(2, 10)}');

// 使用 dart:convert 中的函数
var jsonStr = '{"name": "runoob", "age": 10}';
var decoded = jsonDecode(jsonStr);
print('解析后的 JSON: $decoded');
print('用户名: ${decoded['name']}');
}

```

圆周率: 3.141592653589793
2 的 10 次方: 1024
解析后的 JSON: {name: runoob, age: 10}
用户名: runoob

```

#### 导入第三方包

### 实例

// package: 前缀表示 pub.dev 上的第三方包
import 'package:http/http.dart';
import 'package:path/path.dart' as p;

#### 导入本地文件

### 实例

// 相对路径导入：导入同项目中的其他 Dart 文件
import 'src/utils.dart'; // 同目录下的子目录
import '../models/user.dart'; // 上级目录中的文件
import 'constants.dart'; // 同目录下的文件

#### import 的修饰符

修饰符语法用途 前缀（as）import 'lib.dart' as myLib;给库起别名，避免命名冲突 只导入部分（show）import 'lib.dart' show foo, bar;只导入指定的名称 排除部分（hide）import 'lib.dart' hide foo;导入除指定名称外的所有内容 延迟加载（deferred as）import 'lib.dart' deferred as lib;按需加载，减少启动时间

### 实例

各种导入修饰符的实际用法：

import 'dart:math'; // 标准导入
import 'dart:math' as math; // 前缀导入：使用 math.pow() 而非 pow()
import 'dart:math' show pi, sqrt; // 只导入 pi 和 sqrt
import 'dart:math' hide Random; // 导入除 Random 外的所有内容

void main() {
// 标准导入
print('sin(0) = ${sin(0)}');

// 前缀导入：需要通过前缀访问
print('cos(0) = ${math.cos(0)}');

// show 导入：只能使用 pi 和 sqrt
print('π = $pi');
print('sqrt(16) = ${sqrt(16)}');
// print(sin(0)); // 错误：sin 没有被导入

// hide 导入：Random 不可用，其他都可以
print('max(3, 7) = ${max(3, 7)}');
// Random(); // 错误：Random 被排除了
}

```

sin(0) = 0.0
cos(0) = 1.0
π = 3.141592653589793
sqrt(16) = 4.0
max(3, 7) = 7

```

show 和 hide 不只是为了便利——它们也是代码质量的工具。使用 show 可以让依赖关系更清晰：你能一眼看出这个文件使用了库中的哪些符号。

### export 导出语法

export 用于将其他库的公开 API 重新暴露出去，主要用于创建聚合库。

假设你有以下文件结构：

```

lib/
├── my_package.dart # 入口文件（聚合导出）
├── src/
│ ├── models/
│ │ └── user.dart # User 类
│ ├── services/
│ │ └── api_service.dart # API 服务
│ └── utils/
│ └── helpers.dart # 工具函数

```

### 实例

使用 export 创建聚合库入口：

// 文件：lib/my_package.dart
// 聚合导出：将所有公开 API 统一暴露

export 'src/models/user.dart';
export 'src/services/api_service.dart';
// 只导出 helpers 中的部分内容
export 'src/utils/helpers.dart' show formatDate, validateEmail;

这样使用者只需要导入一个文件：

### 实例

// 使用者只需导入入口文件
import 'package:my_package/my_package.dart';

void main() {
// 可以直接使用所有 export 的类
var user = User('runoob');
var api = ApiService();
}

### 自定义库的创建

创建一个自定义库不需要特殊语法——每个 Dart 文件就是一个库。

但你可以使用 library 关键字明确声明库名，以及使用 part/part of 拆分大型库。

#### 使用 library 声明库名

### 实例

// 文件：lib/calculator.dart
// library 声明库名（可选，但有助于文档化）
library calculator;

/// 加法运算
int add(int a, int b) => a + b;

/// 减法运算
int subtract(int a, int b) => a - b;

/// 乘法运算
int multiply(int a, int b) => a * b;

/// 除法运算，除数为 0 时抛出异常
double divide(int a, int b) {
if (b == 0) {
throw ArgumentError('除数不能为 0');
}
return a / b;
}

#### 使用 part 拆分大型库

当一个库的代码太长时，可以用 part 将其拆分为多个物理文件，但它们在逻辑上仍然属于同一个库。

### 实例

主文件声明 part：

// 文件：lib/user_system.dart
// 主库文件
library user_system;

// 声明组成这个库的其他文件
part 'src/user_model.dart';
part 'src/user_service.dart';
part 'src/user_validator.dart';

// 库级别的公开 API
String libraryVersion = '1.0.0';

### 实例

part 文件声明 part of：

// 文件：lib/src/user_model.dart
// part of 声明自己属于哪个库
part of '../user_system.dart';

// 可以访问主库中的 libraryVersion
class User {
String name;
User(this.name);

void printVersion() {
print('RUNOOB User 模块版本: $libraryVersion');
}
}

part/part of 在现代 Dart 开发中不常用，多数团队更倾向于使用 import/export 来组织代码。part 的缺点是 part 文件之间共享所有私有成员，破坏了封装性。除非有明确需求，否则推荐使用 import/export。

### Dart 核心库速览

库名导入方式主要功能 dart:core自动导入基础类型、集合、异常等 dart:mathimport 'dart:math';数学常量和函数 dart:convertimport 'dart:convert';JSON、UTF-8、Base64 编解码 dart:ioimport 'dart:io';文件、网络、进程操作 dart:asyncimport 'dart:async';Future、Stream 等异步工具 dart:collectionimport 'dart:collection';更多集合类型（Queue 等） dart:developerimport 'dart:developer';调试和性能分析工具

dart:core 是自动导入的，你不需要写 import 'dart:core' 就能使用 int、String、List、Map 等基础类型。

---

## Dart typedef 与函数类型

Source: https://www.runoob.com/dart/dart-typedef.html

## Dart typedef 与函数类型

typedef 是 Dart 中的类型别名机制，它让你为复杂的类型定义简短易读的名称。

本章介绍 typedef 的两种用法：传统的函数类型别名，以及 Dart 3.0 引入的通用类型别名。

### typedef 函数类型别名

当一个函数类型被频繁使用时，typedef 可以给它一个清晰的名字。

这对于回调函数、事件处理器等场景特别有用。

### 实例

// 定义一个函数类型别名
// 表示：接收两个 int 参数，返回 int 的函数
typedef IntOperation = int Function(int a, int b);

// 定义一个回调类型
// 表示：接收 String 消息，无返回值的函数
typedef MessageCallback = void Function(String message);

// 使用 typedef 作为参数类型
int performOperation(int x, int y, IntOperation operation) {
return operation(x, y);
}

void logMessage(String msg, MessageCallback callback) {
print('准备输出消息...');
callback('[RUNOOB] $msg');
}

void main() {
// 传入符合 IntOperation 签名的函数
int add(int a, int b) => a + b;
int multiply(int a, int b) => a * b;

print('10 + 5 = ${performOperation(10, 5, add)}');
print('10 × 5 = ${performOperation(10, 5, multiply)}');

// 也可以直接传入匿名函数
int result = performOperation(20, 4, (a, b) => a ~/ b);
print('20 ÷ 4 = $result');

// 使用 MessageCallback
logMessage('操作完成', (msg) {
print('收到消息: $msg');
});
}

```

10 + 5 = 15
10 × 5 = 50
20 ÷ 4 = 5
准备输出消息...
收到消息: [RUNOOB] 操作完成

```

没有 typedef 时，你需要在每个参数位置重复书写冗长的函数类型签名。

有了 typedef，类型声明变得清晰、统一，修改时也只需要改一处。

typedef 只是给类型起了一个别名，它不会创建新的类型。IntOperation 和 int Function(int, int) 在类型系统中完全等价。

### 函数类型作为参数

在 Dart 中，函数是一等公民（First-class Citizen），可以像普通值一样传递。

### 实例

// 直接使用函数类型声明参数（不使用 typedef）
void processNumbers(
List<int> numbers,
bool Function(int) filter,
String Function(int) formatter,
) {
var filtered = numbers.where(filter);
for (var n in filtered) {
print(formatter(n));
}
}

// 返回一个函数的高阶函数
int Function(int) makeMultiplier(int factor) {
// 返回的闭包捕获了 factor
return (int n) => n * factor;
}

void main() {
var scores = [55, 78, 92, 60, 45, 88];

print('--- 及格分数 ---');
processNumbers(
scores,
(n) => n >= 60, // 过滤条件
(n) => 'RUNOOB 分数: $n 分', // 格式化
);

print('--- 高分（> 80）---');
processNumbers(
scores,
(n) => n > 80,
(n) => '高分: $n 分',
);

// 函数作为返回值
var doubler = makeMultiplier(2);
var tripler = makeMultiplier(3);
print('5 × 2 = ${doubler(5)}');
print('5 × 3 = ${tripler(5)}');
}

```

--- 及格分数 ---
RUNOOB 分数: 78 分
RUNOOB 分数: 92 分
RUNOOB 分数: 60 分
RUNOOB 分数: 88 分
--- 高分（> 80）---
高分: 92 分
高分: 88 分
5 × 2 = 10
5 × 3 = 15

```

### 回调模式

回调模式是函数类型最常见的应用场景。

它将"做什么"的控制权交给调用者，让代码更灵活和可复用。

### 实例

// 定义回调类型
typedef ResultCallback<T> = void Function(T result);
typedef ErrorCallback = void Function(String error);

// 模拟异步操作
void fetchUserData(
String userId, {
required ResultCallback<Map<String, dynamic>> onSuccess,
required ErrorCallback onError,
}) {
// 模拟网络请求
print('正在获取用户数据...');

// 模拟成功/失败
if (userId == 'runoob') {
var data = {
'id': 'runoob',
'name': 'RUNOOB 用户',
'level': 'VIP',
};
onSuccess(data);
} else {
onError('用户 $userId 不存在');
}
}

void main() {
// 使用回调处理结果
fetchUserData(
'runoob',
onSuccess: (data) {
print('获取成功！');
print('用户名: ${data['name']}');
print('等级: ${data['level']}');
},
onError: (error) {
print('获取失败: $error');
},
);

print('---');

// 测试失败情况
fetchUserData(
'unknown',
onSuccess: (data) {
print('获取成功');
},
onError: (error) {
print('获取失败: $error');
},
);
}

```

正在获取用户数据...
获取成功！
用户名: RUNOOB 用户
等级: VIP
---
正在获取用户数据...
获取失败: 用户 unknown 不存在

```

回调模式的常见应用场景：

场景回调类型示例 网络请求onSuccess(data) / onError(error) UI 事件onTap() / onLongPress() 数据转换map()、where()、reduce() 中的回调 定时器Timer(callback, duration) 动画完成onComplete()

虽然回调模式很灵活，但在处理多层异步操作时会导致"回调地狱"。Dart 提供了 async/await 来更优雅地处理异步流程，我们将在第 17 章详细介绍。

### Dart 3.0 通用类型别名

Dart 3.0 扩展了 typedef 的能力，现在你可以为任何类型创建别名，不仅仅是函数类型。

### 实例

// Dart 3.0：typedef 可以为任何类型创建别名

// 复杂集合类型的别名
typedef JsonMap = Map<String, dynamic>;
typedef UserList = List<Map<String, dynamic>>;

// 泛型别名
typedef Result<T> = ({T data, String? error});

// 函数类型别名（传统用法）
typedef Validator<T> = String? Function(T value);

// 使用类型别名
void processJson(JsonMap json) {
print('处理 JSON: $json');
print('键的数量: ${json.length}');
}

void validateAndPrint<T>(T value, Validator<T> validator) {
var error = validator(value);
if (error != null) {
print('验证失败: $error');
} else {
print('验证通过: $value');
}
}

void main() {
// 使用 JsonMap 别名
JsonMap userData = {
'name': 'runoob',
'age': 10,
'isVip': true,
};
processJson(userData);

// 使用 Result 别名
Result<String> successResult = (data: '操作成功', error: null);
Result<String> errorResult = (data: '', error: '网络连接超时');
print('成功: ${successResult.data}');
print('失败: ${errorResult.error}');

// 使用泛型 Validator
Validator<String> nameValidator = (value) {
if (value.isEmpty) return '名称不能为空';
if (value.length < 3) return '名称至少 3 个字符';
return null; // null 表示验证通过
};

validateAndPrint('RUNOOB', nameValidator);
validateAndPrint('AB', nameValidator);
}

```

处理 JSON: {name: runoob, age: 10, isVip: true}
键的数量: 3
成功: 操作成功
失败: 网络连接超时
验证通过: RUNOOB
验证失败: 名称至少 3 个字符

```

Dart 3.0 的类型别名大大减少了冗长的类型声明，让代码更简洁、更具可读性。

类型别名（typedef）和类型本身在运行时完全等价，它们只是编译期的"昵称"。这意味着你不能用 typedef 来区分两个结构相同但语义不同的类型——它们会被视为同一类型。

---

## Dart 异步编程

Source: https://www.runoob.com/dart/dart-async.html

## Dart 异步编程

异步编程是现代应用开发中最核心的能力之一。

无论是网络请求、文件读写还是数据库操作，都需要异步处理来避免阻塞主线程。

本章介绍 Future 的概念、async/await 语法以及如何处理异步错误。

### 为什么需要异步编程

先理解同步和异步的区别。

同步代码按顺序一行一行执行，上一行完成之前下一行不会开始。

如果某个操作需要等待（如网络请求），整个程序就会被卡住。

### 实例

对比同步和异步代码的执行效果：

// 模拟一个耗时的操作
String fetchDataSync() {
// 同步代码：假设这里要等待 2 秒
// 在这 2 秒内，整个程序什么都做不了
sleep(Duration(seconds: 2)); // 阻塞式等待
return '数据获取完成';
}

// 异步版本：不阻塞主线程
Future<String> fetchDataAsync() async {
// 异步等待：这 2 秒内程序可以继续执行其他任务
await Future.delayed(Duration(seconds: 2));
return '数据获取完成';
}

void main() async {
print('开始同步获取数据...');
var syncResult = fetchDataSync();
print('同步结果: $syncResult');
print('（注意：同步期间程序被卡住了）');

print('\n开始异步获取数据...');
print('发起请求后，程序继续执行其他任务...');
var asyncResult = await fetchDataAsync();
print('异步结果: $asyncResult');
print('（异步期间程序没有被卡住）');
}

```

开始同步获取数据...
同步结果: 数据获取完成
（注意：同步期间程序被卡住了）

开始异步获取数据...
发起请求后，程序继续执行其他任务...
异步结果: 数据获取完成
（异步期间程序没有被卡住）

```

在 Dart 中，所有 I/O 操作（文件、网络、数据库）都应该使用异步方式。同步 I/O 会阻塞整个 Isolate，在 Flutter 应用中会导致 UI 卡顿。

### Future 的概念

Future 代表一个"未来某个时刻会完成"的值。

它像一个承诺：我现在没有结果，但将来一定会给你一个 T 类型的值（或者一个错误）。

Future 有三种状态：

状态说明 未完成（Uncompleted）异步操作还在进行中 完成且有值（Completed with data）操作成功，Future 携带了结果值 完成但有错误（Completed with error）操作失败，Future 携带了错误信息

### 实例

创建和使用 Future：

void main() {
print('程序开始');

// 创建一个 Future
Future<String> future = Future(() {
// 这个函数会在未来某个时刻执行
return 'RUNOOB 异步结果';
});

print('Future 已创建，但还没完成');

// 注册回调：当 Future 完成时执行
future.then((result) {
print('收到结果: $result');
});

print('程序继续执行...');
}

```

程序开始
Future 已创建，但还没完成
程序继续执行...
收到结果: RUNOOB 异步结果

```

注意输出的顺序：先打印"程序继续执行"，后打印"收到结果"。

这说明 Future 的回调是异步执行的，不会阻塞后续代码。

### async / await 语法

async 和 await 是 Dart 处理异步操作的核心语法。

async 标记一个函数为异步函数，await 等待一个 Future 完成并获取其结果。

### 实例

// async 关键字标记函数为异步函数
// 异步函数的返回类型必须是 Future<T>
Future<String> fetchUserName(int id) async {
// await 等待 Future 完成，并获取其值
// 在等待期间，当前函数暂停执行，但不会阻塞其他代码
await Future.delayed(Duration(seconds: 1));
return '用户$id';
}

Future<int> fetchUserAge(int id) async {
await Future.delayed(Duration(milliseconds: 500));
return 20 + id;
}

// 顺序执行多个异步操作
Future<void> printUserInfo(int id) async {
print('开始获取用户信息...');

// 逐个等待：先获取名称，再获取年龄
var name = await fetchUserName(id);
print('获取到名称: $name');

var age = await fetchUserAge(id);
print('获取到年龄: $age');

print('RUNOOB 用户: $name, 年龄: $age');
}

// 并发执行多个异步操作
Future<void> printUserInfoParallel(int id) async {
print('开始并发获取用户信息...');

// 同时发起两个请求，不互相等待
var nameFuture = fetchUserName(id);
var ageFuture = fetchUserAge(id);

// 等待两个请求都完成
var results = await Future.wait([nameFuture, ageFuture]);
var name = results[0] as String;
var age = results[1] as int;

print('RUNOOB 用户(并发): $name, 年龄: $age');
}

void main() async {
// 顺序执行（总时间 ≈ 1s + 0.5s = 1.5s）
await printUserInfo(1);

print('---');

// 并发执行（总时间 ≈ max(1s, 0.5s) = 1s）
await printUserInfoParallel(2);
}

```

开始获取用户信息...
获取到名称: 用户1
获取到年龄: 21
RUNOOB 用户: 用户1, 年龄: 21
---
开始并发获取用户信息...
RUNOOB 用户(并发): 用户2, 年龄: 22

```

如果多个异步操作之间没有依赖关系，应该使用 Future.wait 让它们并发执行，而不是逐个 await。这能显著提升性能——总耗时等于最慢的操作，而非所有操作之和。

#### async/await 的核心规则

规则说明 async 函数返回 Future即使函数返回 T，async 后会自动包装为 Future<T> await 只能在 async 函数中使用普通函数不能使用 await await 暂停当前函数但不阻塞其他代码的执行 await 获取 Future 的结果如果 Future 有错误，await 会抛出异常

### then / catchError 链式调用

除了 async/await，Dart 还支持使用 then() 和 catchError() 进行链式调用。

这是传统的 Promise 风格写法。

### 实例

// 模拟网络请求的函数
Future<String> fetchData(String url) {
return Future.delayed(Duration(seconds: 1), () {
if (url.isEmpty) {
throw Exception('URL 不能为空');
}
return '来自 $url 的数据';
});
}

void main() {
print('开始请求...');

// then / catchError 链式调用
fetchData('https://runoob.com/api')
.then((data) {
// 请求成功
print('获取到数据: $data');
return '处理后的: $data'; // 返回值会传递给下一个 then
})
.then((processed) {
// 处理上一个 then 的返回值
print(processed);
})
.catchError((error) {
// 统一捕获链中的任何错误
print('请求出错: $error');
})
.whenComplete(() {
// 无论成功还是失败都会执行（类似 finally）
print('请求结束（无论成功或失败）');
});

print('请求已发出，程序继续...');
}

```

开始请求...
请求已发出，程序继续...
获取到数据: 来自 https://runoob.com/api 的数据
处理后的: 来自 https://runoob.com/api 的数据
请求结束（无论成功或失败）

```

#### async/await vs then/catchError 如何选择

场景推荐方式原因 线性异步流程async/await代码像同步一样清晰 并发多个请求async/await + Future.wait语义明确 链式数据转换then()每个 then 做一步转换，链式清晰 单一回调then()比 async/await 更简洁

大多数情况下推荐 async/await，它的可读性更好。但 then() 在简单的链式处理中依然有价值。两者可以混用，但不建议——一个函数要么用 async/await，要么用 then/catchError，混用会让代码难以理解。

### 处理异步错误

异步操作中可能发生错误，需要用 try-catch 来处理。

async/await 让错误处理变得和同步代码一样简单。

### 实例

// 模拟可能失败的网络请求
Future<String> fetchUserProfile(int userId) async {
await Future.delayed(Duration(seconds: 1));

if (userId <= 0) {
throw Exception('无效的用户 ID: $userId');
}

if (userId == 404) {
throw HttpException('用户不存在');
}

return '用户 $userId 的个人资料';
}

// 自定义异常
class HttpException implements Exception {
final String message;
HttpException(this.message);

@override
String toString() => 'HttpException: $message';
}

Future<void> loadUserProfile(int userId) async {
print('正在加载用户 $userId 的资料...');

try {
var profile = await fetchUserProfile(userId);
print('加载成功: $profile');
} on HttpException catch (e) {
// 按类型捕获特定的异步错误
print('HTTP 错误: $e');
print('提示用户：请检查用户是否存在');
} on Exception catch (e) {
// 捕获其他异常
print('一般错误: $e');
} finally {
print('加载操作结束');
}
}

void main() async {
await loadUserProfile(1); // 成功
print('---');
await loadUserProfile(404); // HttpException
print('---');
await loadUserProfile(-1); // 一般 Exception
}

```

正在加载用户 1 的资料...
加载成功: 用户 1 的个人资料
加载操作结束
---
正在加载用户 404 的资料...
HTTP 错误: HttpException: 用户不存在
提示用户：请检查用户是否存在
加载操作结束
---
正在加载用户 -1 的资料...
一般错误: Exception: 无效的用户 ID: -1
加载操作结束

```

async 函数中的错误处理规则：

- 如果 async 函数中抛出异常，该异常会被自动包装到 Future 中
- 调用者可以用 try-catch 捕获 await 的异常
- 如果使用 then()，异常会传递到 catchError()
- 未捕获的异步异常不会导致程序崩溃，但会被视为未处理的 Future 错误

永远不要忽略异步错误。每个 async 函数调用都应该有对应的错误处理。未处理的异步错误在开发模式下会打印警告，在生产环境中可能被静默忽略，导致难以排查的 bug。

---

## Dart Stream 流

Source: https://www.runoob.com/dart/dart-streams.html

## Dart Stream 流

Stream（流）是 Dart 中处理连续异步事件序列的机制。

如果说 Future 是一次性的异步结果，那么 Stream 就是多次的、持续不断的异步数据流。

本章介绍 Stream 的概念、await for 监听、StreamController 创建以及单订阅与广播流的区别。

### Stream 与事件序列

Stream 就像一条传送带，数据会随着时间推移一个个到来。

你不需要一次性等待所有数据，而是来一个处理一个。

Stream 的典型应用场景：

场景示例 用户输入事件按钮点击、鼠标移动、键盘输入 文件读取逐行读取大文件 网络数据WebSocket 消息、实时 API 定时器每秒触发一次的定时事件 状态变化Flutter 中的状态管理流

### 实例

Stream 的基本使用——listen 监听：

void main() {
// 创建一个 Stream：每隔 1 秒发射一个数字
var stream = Stream<int>.periodic(
Duration(seconds: 1),
(count) => count + 1, // count 从 0 开始
);

print('开始监听 Stream...');

// listen 订阅 Stream
var subscription = stream.listen(
(data) {
// 每当有新数据到达时调用
print('RUNOOB 收到数据: $data');
},
onError: (error) {
// Stream 发生错误时调用
print('错误: $error');
},
onDone: () {
// Stream 关闭时调用
print('Stream 已关闭');
},
);

// 5 秒后取消订阅
Future.delayed(Duration(seconds: 5), () {
subscription.cancel();
print('已取消订阅');
});
}

```

开始监听 Stream...
RUNOOB 收到数据: 1
RUNOOB 收到数据: 2
RUNOOB 收到数据: 3
RUNOOB 收到数据: 4
RUNOOB 收到数据: 5
已取消订阅

```

listen() 返回一个 StreamSubscription 对象，你可以用它来控制订阅：

方法作用 subscription.pause()暂停接收数据 subscription.resume()恢复接收数据 subscription.cancel()取消订阅 subscription.isPaused是否处于暂停状态

### await for 监听

除了 listen() 回调方式，还可以使用 await for 循环来消费 Stream。

await for 让 Stream 的处理逻辑像普通的 for 循环一样清晰。

### 实例

// 生成一个有限的数据流
Stream<int> countStream(int max) async* {
for (int i = 1; i <= max; i++) {
await Future.delayed(Duration(milliseconds: 300));
yield i; // yield 发射数据到 Stream
}
}

Future<void> main() async {
print('开始用 await for 消费 Stream...');

// await for：等待每个数据到达，逐个处理
await for (var value in countStream(5)) {
print('RUNOOB 计数: $value');
}

print('Stream 消费完毕');
}

```

开始用 await for 消费 Stream...
RUNOOB 计数: 1
RUNOOB 计数: 2
RUNOOB 计数: 3
RUNOOB 计数: 4
RUNOOB 计数: 5
Stream 消费完毕

```

await for 的特点：

- 和普通 for 循环写法相似，但每次迭代会等待下一个数据到达
- 当 Stream 关闭时，循环自动结束
- 只能在 async 函数中使用
- 如果需要中途退出，使用 break（和普通 for 循环一样）

await for 适合"需要顺序处理所有数据"的场景，如读取文件所有行。listen() 适合"需要持续响应事件"的场景，如按钮点击。两者可以互相替代，但各有擅长的领域。

### StreamController 创建 Stream

StreamController 是手动创建和控制 Stream 的工具。

你可以随时向其中添加数据、错误，或者关闭它。

### 实例

import 'dart:async';

// 使用 StreamController 实现一个简单的倒计时器
class CountdownTimer {
final StreamController<int> _controller = StreamController<int>();
Timer? _timer;
int _remaining = 0;

// 暴露 Stream 给外部订阅
Stream<int> get tickStream => _controller.stream;

// 开始倒计时
void start(int seconds) {
_remaining = seconds;

// 立即发送初始值
_controller.add(_remaining);

_timer = Timer.periodic(Duration(seconds: 1), (timer) {
_remaining--;

if (_remaining > 0) {
_controller.add(_remaining); // 发送数据
} else {
_controller.add(0); // 发送最后的 0
_controller.close(); // 关闭 Stream
timer.cancel();
}
});
}

// 取消倒计时
void cancel() {
_timer?.cancel();
_controller.addError('倒计时被取消'); // 发送错误
_controller.close();
}

// 释放资源
void dispose() {
_controller.close();
}
}

Future<void> main() async {
var timer = CountdownTimer();

// 订阅倒计时事件
timer.tickStream.listen(
(remaining) {
print('RUNOOB 倒计时: $remaining 秒');
},
onError: (error) {
print('错误: $error');
},
onDone: () {
print('倒计时结束！');
},
);

timer.start(5);

// 等待倒计时完成
await Future.delayed(Duration(seconds: 6));
timer.dispose();
}

```

RUNOOB 倒计时: 5 秒
RUNOOB 倒计时: 4 秒
RUNOOB 倒计时: 3 秒
RUNOOB 倒计时: 2 秒
RUNOOB 倒计时: 1 秒
RUNOOB 倒计时: 0 秒
倒计时结束！

```

#### async* 生成器函数

如果你只需要简单地生成一个数据序列，async* 比 StreamController 更方便。

### 实例

// async* 标记这是一个异步生成器函数
// 返回类型必须是 Stream
Stream<String> readLinesAsync() async* {
var lines = ['第一行', '第二行', '第三行', 'RUNOOB'];

for (var line in lines) {
await Future.delayed(Duration(milliseconds: 500));
yield line; // yield 发射数据到 Stream
}
// 函数结束时 Stream 自动关闭
}

// 带错误处理的异步生成器
Stream<int> generateNumbersWithError() async* {
for (int i = 1; i <= 5; i++) {
await Future.delayed(Duration(milliseconds: 300));

if (i == 3) {
throw Exception('数字 3 出错了！');
}
yield i;
}
}

Future<void> main() async {
print('逐行读取：');
await for (var line in readLinesAsync()) {
print(' $line');
}

print('\n带错误的生成器：');
try {
await for (var num in generateNumbersWithError()) {
print(' 数字: $num');
}
} catch (e) {
print(' 捕获错误: $e');
}
}

```

逐行读取：
第一行
第二行
第三行
RUNOOB

带错误的生成器：
数字: 1
数字: 2
捕获错误: Exception: 数字 3 出错了！

```

### 单订阅 vs 广播流

Dart 的 Stream 分为两种类型：单订阅流（Single-subscription）和广播流（Broadcast）。

#### 单订阅流（Single-subscription Stream）

这是默认的 Stream 类型。

它只能被一个监听者订阅，适合"从头到尾消费一次"的场景。

### 实例

void main() {
// 单订阅流（默认）
var stream = Stream.fromIterable([1, 2, 3]);

// 第一个订阅：OK
stream.listen((data) => print('订阅者1: $data'));

// 第二个订阅：错误！单订阅流不能被多次订阅
// stream.listen((data) => print('订阅者2: $data')); // 运行时错误
}

#### 广播流（Broadcast Stream）

广播流允许多个监听者同时订阅，适合事件广播场景。

### 实例

import 'dart:async';

void main() {
// 创建一个广播流
var controller = StreamController<int>.broadcast();

// 多个订阅者可以同时监听
controller.stream.listen(
(data) => print('RUNOOB 订阅者A: 收到 $data'),
);

controller.stream.listen(
(data) => print('RUNOOB 订阅者B: 收到 $data'),
);

// 发送数据，两个订阅者都会收到
controller.add(1);
controller.add(2);
controller.add(3);

// 延迟订阅也能收到后续数据（但不会收到之前的数据）
Future.delayed(Duration(seconds: 1), () {
controller.stream.listen(
(data) => print('RUNOOB 迟到的订阅者C: 收到 $data'),
);
controller.add(4);
controller.close();
});
}

```

RUNOOB 订阅者A: 收到 1
RUNOOB 订阅者B: 收到 1
RUNOOB 订阅者A: 收到 2
RUNOOB 订阅者B: 收到 2
RUNOOB 订阅者A: 收到 3
RUNOOB 订阅者B: 收到 3
RUNOOB 订阅者A: 收到 4
RUNOOB 订阅者B: 收到 4
RUNOOB 迟到的订阅者C: 收到 4

```

两种 Stream 的对比：

特性单订阅流广播流 订阅者数量只能一个可以多个 数据重放从头开始消费只接收订阅后的数据 典型场景文件读取、HTTP 响应按钮点击、状态通知 创建方式StreamController()StreamController.broadcast()

广播流的"迟到订阅者"只能收到订阅之后的事件，之前的事件已经错过了。这是广播流和单订阅流最重要的行为差异。

### Stream 常用转换方法

Stream 提供了一系列方法来转换和处理数据，类似于 List 的函数式方法。

### 实例

void main() async {
var numbers = Stream.fromIterable([1, 2, 3, 4, 5, 6]);

// map：转换每个数据
var doubled = numbers.map((n) => n * 2);
print('翻倍:');
await for (var n in doubled) {
print(' $n');
}

// where：过滤数据
var source = Stream.fromIterable([10, 15, 20, 25, 30]);
var even = source.where((n) => n % 2 == 0);
print('偶数:');
await for (var n in even) {
print(' $n');
}

// take：只取前 N 个
var infinite = Stream.periodic(
Duration(milliseconds: 100),
(i) => i + 1,
);
print('只取前 3 个:');
await for (var n in infinite.take(3)) {
print(' RUNOOB: $n');
}

// skip：跳过前 N 个
var data = Stream.fromIterable([1, 2, 3, 4, 5]);
print('跳过前 2 个:');
await for (var n in data.skip(2)) {
print(' $n');
}

// distinct：去重
var duplicates = Stream.fromIterable([1, 2, 2, 3, 3, 3]);
print('去重:');
await for (var n in duplicates.distinct()) {
print(' $n');
}
}

```

翻倍:
2
4
6
8
10
12
偶数:
10
20
30
只取前 3 个:
RUNOOB: 1
RUNOOB: 2
RUNOOB: 3
跳过前 2 个:
3
4
5
去重:
1
2
3

```

---

## Dart 并发与 Isolate

Source: https://www.runoob.com/dart/dart-isolates.html

## Dart 并发与 Isolate

并发是指同时处理多个任务的能力。

Dart 的并发模型不同于传统的多线程，它使用 Isolate（隔离区）来实现并行计算。

本章介绍 Dart 的并发模型、Isolate 的概念与使用以及消息传递机制。

### Dart 并发模型

大多数编程语言使用共享内存的多线程模型，多个线程共享同一个内存空间。

这种模型的缺点是容易产生竞态条件（race condition）和死锁（deadlock）。

Dart 采用了不同的策略：每个 Isolate 拥有自己独立的内存堆，Isolate 之间不共享内存。

它们通过消息传递来通信，这从根本上避免了数据竞争问题。

Dart 的并发模型有三个层次：

层次机制适用场景 事件循环单线程异步（Event Loop）I/O 操作、定时器、用户交互 Isolate独立内存 + 消息传递CPU 密集型计算 Future/Stream异步编程语法大多数日常开发场景

大多数 Dart 程序的并发需求都可以通过 async/await 和 Future/Stream 来满足。Isolate 只在你需要进行大量 CPU 密集型计算时才有必要使用。过早引入 Isolate 会让代码变复杂，收益却不大。

### Isolate 概念与使用

Isolate 是 Dart 的并发单元，每个 Isolate 有自己独立的内存和事件循环。

主程序本身就在一个 Isolate（主 Isolate）中运行。

#### 使用 Isolate.spawn 创建新 Isolate

### 实例

import 'dart:isolate';

// 这个函数会在新的 Isolate 中运行
// SendPort 用于向主 Isolate 发送消息
void heavyComputation(SendPort sendPort) {
print('新 Isolate 开始计算...');

// 模拟 CPU 密集型计算
int sum = 0;
for (int i = 1; i <= 10000000; i++) {
sum += i;
}

// 将计算结果发送回主 Isolate
sendPort.send(sum);

print('新 Isolate 计算完成，结果已发送');
}

Future<void> main() async {
print('主 Isolate 启动');

// 创建一个 ReceivePort 用于接收消息
var receivePort = ReceivePort();

// spawn 创建新的 Isolate
await Isolate.spawn(heavyComputation, receivePort.sendPort);

print('主 Isolate 在等待结果的同时可以做其他事情...');

// 等待新 Isolate 的计算结果
var result = await receivePort.first;
print('RUNOOB 计算结果: 1 到 10000000 的和 = $result');

receivePort.close();
print('主 Isolate 结束');
}

```

主 Isolate 启动
主 Isolate 在等待结果的同时可以做其他事情...
新 Isolate 开始计算...
新 Isolate 计算完成，结果已发送
RUNOOB 计算结果: 1 到 10000000 的和 = 50000005000000
主 Isolate 结束

```

#### 对比：有无 Isolate 的性能差异

### 实例

// 在主 Isolate 中执行 CPU 密集任务（会阻塞）
int fibonacci(int n) {
if (n <= 1) return n;
return fibonacci(n - 1) + fibonacci(n - 2);
}

void main() {
// 在主 Isolate 中计算——会阻塞所有其他操作
var startTime = DateTime.now();

// 注意：如果 n 太大，这个计算会非常耗时
// 实际开发中应该把这种计算放到单独的 Isolate 中
var result = fibonacci(40);

var elapsed = DateTime.now().difference(startTime);
print('RUNOOB 斐波那契(40) = $result');
print('耗时: ${elapsed.inMilliseconds}ms');
print('（注意：在计算期间，主 Isolate 无法处理其他任务）');
}

在 Flutter 应用中，如果在主 Isolate 中执行耗时的同步计算，会导致 UI 冻结。解决方案就是将计算任务放到单独的 Isolate 中，计算完成后通过消息将结果传回并更新 UI。

### 消息传递机制

Isolate 之间通过 SendPort 和 ReceivePort 进行消息传递。

消息必须是可序列化的（基本类型、String、List、Map 等），不能传递函数或闭包。

#### 双向通信

### 实例

import 'dart:isolate';

// 在新 Isolate 中运行的工作函数
void workerIsolate(SendPort mainSendPort) {
// 创建自己的 ReceivePort 来接收主 Isolate 的消息
var workerReceivePort = ReceivePort();

// 先把 worker 的 SendPort 发给主 Isolate，建立双向通道
mainSendPort.send(workerReceivePort.sendPort);

print('Worker Isolate: 等待任务...');

// 监听主 Isolate 发来的任务
workerReceivePort.listen((message) {
if (message is List<int>) {
// 收到任务：对列表中的每个数求平方
print('Worker Isolate: 收到数据 $message');
var result = message.map((n) => n * n).toList();

// 通过 mainSendPort 发回结果
mainSendPort.send(result);
} else if (message == 'exit') {
print('Worker Isolate: 收到退出信号，关闭');
workerReceivePort.close();
mainSendPort.send('goodbye');
}
});
}

Future<void> main() async {
print('主 Isolate: 启动');

// 创建主接收端口
var mainReceivePort = ReceivePort();

// 启动 Worker Isolate
await Isolate.spawn(workerIsolate, mainReceivePort.sendPort);

// 等待 Worker 发来它的 SendPort（建立双向通信）
SendPort? workerSendPort;
await for (var msg in mainReceivePort) {
if (msg is SendPort) {
workerSendPort = msg;
print('主 Isolate: 已建立与 Worker 的双向通信');
break;
}
}

// 通过 Worker 的 SendPort 发送任务
workerSendPort!.send([1, 2, 3, 4, 5]);
workerSendPort.send([10, 20, 30]);

// 接收 Worker 的计算结果
int responseCount = 0;
await for (var msg in mainReceivePort) {
if (msg is List<int>) {
print('主 Isolate: 收到结果 $msg');
responseCount++;
if (responseCount == 2) break;
}
}

// 发送退出信号
workerSendPort.send('exit');
await for (var msg in mainReceivePort) {
if (msg == 'goodbye') {
print('主 Isolate: Worker 已退出');
break;
}
}

mainReceivePort.close();
print('RUNOOB 双向通信演示结束');
}

```

主 Isolate: 启动
主 Isolate: 已建立与 Worker 的双向通信
Worker Isolate: 等待任务...
Worker Isolate: 收到数据 [1, 2, 3, 4, 5]
主 Isolate: 收到结果 [1, 4, 9, 16, 25]
Worker Isolate: 收到数据 [10, 20, 30]
主 Isolate: 收到结果 [100, 400, 900]
Worker Isolate: 收到退出信号，关闭
主 Isolate: Worker 已退出
RUNOOB 双向通信演示结束

```

#### 使用 Isolate.run 简化（Dart 3.0+）

Dart 3.0 引入了 Isolate.run()，大大简化了单次计算任务的使用。

### 实例

import 'dart:isolate';

// 一个耗时的计算函数
int complexCalculation(int n) {
int result = 0;
for (int i = 0; i < n; i++) {
result += i * i;
}
return result;
}

Future<void> main() async {
print('开始计算...');

// Isolate.run：一行代码搞定，自动创建 Isolate、执行、返回结果
var result = await Isolate.run(() => complexCalculation(10000000));
print('RUNOOB 计算结果: $result');

print('计算完成');

// 对比：如果不使用 Isolate
print('（如果在主 Isolate 中直接计算，会阻塞其他操作）');
}

```

开始计算...
RUNOOB 计算结果: 333333283333335000000
计算完成
（如果在主 Isolate 中直接计算，会阻塞其他操作）

```

Isolate.run() 适合"执行一次计算并返回结果"的场景。如果你需要持续的、双向的通信（比如长时间运行的后台服务），则需要使用 Isolate.spawn() + SendPort/ReceivePort。

#### Isolate 消息传递的限制

可以传递不能传递 null、bool、int、double、String函数、闭包 List、Map、Set（元素也可传递）Stream、Future SendPort（用于建立通信链）大多数非基本类型的对象 Capability（权限令牌）文件句柄、网络 Socket TransferableTypedData（高效传递大块数据）自定义类（除非可序列化）

### 并发模型对比

特性Dart Isolate传统多线程（如 Java） 内存模型独立内存，不共享共享内存 通信方式消息传递（SendPort）共享变量 + 锁 数据竞争不存在（无共享）需要锁机制保护 死锁风险几乎不存在存在 创建开销相对较大相对较小 适用场景CPU 密集型并行计算通用并发

Dart 的 Isolate 模型虽然避免了数据竞争，但代价是创建和通信的开销较大。不要创建成千上万个 Isolate——通常几个 Isolate 就足够了。对于 I/O 密集型任务，使用 async/await 是最佳选择。

---

## Dart 单元测试

Source: https://www.runoob.com/dart/dart-unit-testing.html

## Dart 单元测试

单元测试是保证代码质量的基础手段。

它验证每个最小功能单元（通常是函数或方法）的行为是否符合预期。

本章介绍 Dart 的 test 包使用、test() 测试用例编写、expect 断言以及 group() 分组测试。

### test 包的安装与配置

Dart 官方提供了 test 包来编写和运行测试。

首先在 pubspec.yaml 中添加依赖：

```

# 文件路径：pubspec.yaml
dev_dependencies:
test: ^1.24.0
```

然后运行以下命令安装：

```

$ dart pub get

```

测试文件通常放在项目的 test/ 目录下，文件名以 _test.dart 结尾。

一个典型的项目结构：

```

my_project/
├── lib/
│ └── calculator.dart # 被测试的代码
├── test/
│ └── calculator_test.dart # 测试文件
└── pubspec.yaml

```

### 编写 test() 测试用例

test() 函数是编写测试用例的基本单元。

它接收两个参数：测试描述（字符串）和测试函数体。

### 实例

被测试的代码（lib/calculator.dart）：

// 文件路径：lib/calculator.dart
class Calculator {
int add(int a, int b) => a + b;
int subtract(int a, int b) => a - b;
int multiply(int a, int b) => a * b;

double divide(int a, int b) {
if (b == 0) {
throw ArgumentError('除数不能为 0');
}
return a / b;
}

bool isEven(int n) => n % 2 == 0;

List<int> filterPositive(List<int> numbers) {
return numbers.where((n) => n > 0).toList();
}
}

### 实例

测试文件（test/calculator_test.dart）：

// 文件路径：test/calculator_test.dart
import 'package:test/test.dart';
import 'package:my_project/calculator.dart';

void main() {
// 创建测试用的 Calculator 实例
var calculator = Calculator();

// test() 函数：第一个参数是测试描述，第二个是测试函数
test('add() 两个正数相加', () {
var result = calculator.add(3, 5);
// expect() 断言：验证实际结果是否等于期望值
expect(result, equals(8));
});

test('add() 包含负数相加', () {
expect(calculator.add(-3, 5), equals(2));
expect(calculator.add(-3, -5), equals(-8));
});

test('subtract() 减法运算', () {
expect(calculator.subtract(10, 3), equals(7));
});

test('multiply() 乘法运算', () {
expect(calculator.multiply(4, 5), equals(20));
// 乘以零
expect(calculator.multiply(100, 0), equals(0));
});

test('divide() 正常除法', () {
expect(calculator.divide(10, 2), equals(5.0));
expect(calculator.divide(7, 2), equals(3.5));
});

test('isEven() 偶数判断', () {
expect(calculator.isEven(4), isTrue);
expect(calculator.isEven(5), isFalse);
});
}

运行测试：

```

$ dart test

```

```

00:00 +6: All tests passed!

```

### expect 断言

expect() 是测试中最核心的函数，它验证实际值是否满足某个条件。

基本语法：expect(actual, matcher)。

#### 常用 Matcher

Matcher用途示例 equals(expected)验证值相等expect(result, equals(42)) isTrue / isFalse验证布尔值expect(flag, isTrue) isNull / isNotNull验证 nullexpect(value, isNull) contains(value)包含某个元素（列表）或子串（字符串）expect(list, contains('a')) isA()验证类型expect(obj, isA()) throws()验证抛出异常expect(() => f(), throwsException) isNotEmpty非空expect(list, isNotEmpty) hasLength(n)验证长度expect(list, hasLength(3)) greaterThan(n)大于expect(score, greaterThan(60)) closeTo(value, delta)浮点数近似相等expect(3.14, closeTo(3.1, 0.1))

### 实例

各种 Matcher 的实际应用：

import 'package:test/test.dart';

void main() {
test('RUNOOB 各种断言示例', () {
// 基本相等
expect(2 + 2, equals(4));

// 布尔值
expect('hello'.contains('h'), isTrue);
expect(''.isEmpty, isTrue);

// null 检查
String? name;
expect(name, isNull);
name = 'RUNOOB';
expect(name, isNotNull);

// 类型检查
expect('RUNOOB', isA<String>());
expect(42, isA<int>());

// 列表/字符串包含
expect([1, 2, 3], contains(2));
expect('Hello, RUNOOB!', contains('RUNOOB'));

// 长度
expect([1, 2, 3], hasLength(3));
expect('Dart', hasLength(4));

// 数值比较
expect(100, greaterThan(50));
expect(30, lessThan(60));
expect(75, greaterThanOrEqualTo(60));

// 浮点数比较（避免精度问题）
expect(0.1 + 0.2, closeTo(0.3, 0.001));

// 集合非空
expect([1, 2], isNotEmpty);

// 列表相等
expect([1, 2, 3], equals([1, 2, 3]));

// Map 包含某个键
var user = {'name': 'runoob', 'age': 10};
expect(user, containsPair('name', 'runoob'));
expect(user.keys, contains('age'));
});
}

```

00:00 +1: All tests passed!

```

#### 测试异常

### 实例

import 'package:test/test.dart';

int divide(int a, int b) {
if (b == 0) throw ArgumentError('除数不能为 0');
return a ~/ b;
}

void main() {
test('divide() 除零应该抛出异常', () {
// 验证抛出任意异常
expect(() => divide(10, 0), throwsException);

// 验证抛出特定类型的异常
expect(() => divide(10, 0), throwsArgumentError);

// 验证异常消息
expect(
() => divide(10, 0),
throwsA(predicate((e) =>
e is ArgumentError &&
e.message.contains('除数'))),
);
});

test('divide() 正常情况不抛异常', () {
expect(() => divide(10, 2), returnsNormally);
expect(divide(10, 2), equals(5));
});
}

```

00:00 +2: All tests passed!

```

测试异常时，需要将可能抛异常的代码包装在匿名函数中（() => code），而不是直接调用。如果直接写 expect(divide(10, 0), throwsException)，divide 会立即抛出异常，expect 根本没机会执行。

### 分组测试 group()

group() 用于将相关的测试用例组织在一起，让测试结构更清晰。

### 实例

import 'package:test/test.dart';

// 被测试的函数
class StringUtils {
static String capitalize(String s) {
if (s.isEmpty) return s;
return s[0].toUpperCase() + s.substring(1);
}

static String reverse(String s) {
return s.split('').reversed.join('');
}

static bool isPalindrome(String s) {
var clean = s.toLowerCase().replaceAll(' ', '');
return clean == reverse(clean);
}

static int countWords(String s) {
if (s.trim().isEmpty) return 0;
return s.trim().split(RegExp(r'\s+')).length;
}
}

void main() {
// group() 嵌套组织测试
group('StringUtils', () {
// 子分组
group('capitalize()', () {
test('正常单词首字母大写', () {
expect(StringUtils.capitalize('hello'), equals('Hello'));
});

test('已大写的单词不变', () {
expect(StringUtils.capitalize('Hello'), equals('Hello'));
});

test('空字符串返回空字符串', () {
expect(StringUtils.capitalize(''), equals(''));
});

test('单个字符', () {
expect(StringUtils.capitalize('a'), equals('A'));
});
});

group('reverse()', () {
test('反转正常字符串', () {
expect(StringUtils.reverse('RUNOOB'), equals('BOONUR'));
});

test('反转回文字符串不变', () {
expect(StringUtils.reverse('aba'), equals('aba'));
});

test('空字符串', () {
expect(StringUtils.reverse(''), equals(''));
});
});

group('isPalindrome()', () {
test('回文字符串返回 true', () {
expect(StringUtils.isPalindrome('racecar'), isTrue);
expect(StringUtils.isPalindrome('A man a plan a canal Panama'), isTrue);
});

test('非回文字符串返回 false', () {
expect(StringUtils.isPalindrome('hello'), isFalse);
});

test('空字符串视为回文', () {
expect(StringUtils.isPalindrome(''), isTrue);
});
});

group('countWords()', () {
test('正常句子', () {
expect(StringUtils.countWords('Hello World Dart'), equals(3));
});

test('多余空格', () {
expect(StringUtils.countWords(' Hello World '), equals(2));
});

test('空字符串', () {
expect(StringUtils.countWords(''), equals(0));
expect(StringUtils.countWords(' '), equals(0));
});
});
});
}

```

00:00 +12: All tests passed!

```

group() 可以嵌套，创建层次化的测试结构。

这让测试报告更易读，也方便快速定位失败的测试。

一个好的分组策略是按"被测试的模块/类/方法"来组织。这样当某个测试失败时，你能立即知道是哪个功能出了问题。

#### setUp 和 tearDown

setUp 在每个 test 之前运行，tearDown 在每个 test 之后运行。

它们用于准备测试环境和清理资源。

### 实例

import 'package:test/test.dart';

// 模拟一个需要初始化和清理的类
class Database {
bool isConnected = false;

void connect() {
isConnected = true;
print(' 数据库已连接');
}

void disconnect() {
isConnected = false;
print(' 数据库已断开');
}

String query(String sql) {
if (!isConnected) throw Exception('未连接数据库');
return '查询结果: $sql';
}
}

void main() {
group('Database 测试', () {
late Database db; // late 延迟初始化

// 每个 test 之前执行
setUp(() {
db = Database();
db.connect();
print(' [setUp] 准备测试环境');
});

// 每个 test 之后执行
tearDown(() {
db.disconnect();
print(' [tearDown] 清理测试环境');
});

test('query() 正常查询', () {
var result = db.query('SELECT * FROM users');
expect(result, contains('RUNOOB') ? result.contains('users') : result.contains('users'));
// 简化断言
expect(result, isNotEmpty);
});

test('query() 连接后可以执行多次查询', () {
var r1 = db.query('SELECT 1');
var r2 = db.query('SELECT 2');
expect(r1, isNotEmpty);
expect(r2, isNotEmpty);
});

test('未连接时查询会抛异常', () {
db.disconnect(); // 手动断开
expect(() => db.query('SELECT 1'), throwsException);
});
});
}

```

数据库已连接
[setUp] 准备测试环境
数据库已断开
[tearDown] 清理测试环境
数据库已连接
[setUp] 准备测试环境
数据库已断开
[tearDown] 清理测试环境
数据库已连接
[setUp] 准备测试环境
数据库已断开
[tearDown] 清理测试环境
00:00 +3: All tests passed!

```

注意输出中 setUp 和 tearDown 的执行顺序——每个 test 前后都会执行一次。

setUp 和 tearDown 确保每个测试从一个干净的状态开始，不受其他测试的影响。这是测试独立性的关键保证。

### 异步测试

Dart test 包原生支持异步测试。

测试函数可以返回 Future，框架会自动等待 Future 完成。

### 实例

import 'package:test/test.dart';

// 异步函数：模拟网络请求
Future<String> fetchUserData(int userId) async {
await Future.delayed(Duration(milliseconds: 100));

if (userId <= 0) {
throw Exception('无效的用户 ID');
}

return '用户 $userId 的数据';
}

Future<List<int>> fetchNumbers() async {
await Future.delayed(Duration(milliseconds: 50));
return [1, 2, 3, 4, 5];
}

void main() {
group('异步测试', () {
test('fetchUserData() 正常获取数据', () async {
var data = await fetchUserData(1);
expect(data, equals('用户 1 的数据'));
});

test('fetchUserData() 无效 ID 抛异常', () async {
expect(
() => fetchUserData(-1),
throwsA(isA<Exception>()),
);
});

test('fetchNumbers() 返回正确的列表', () async {
var numbers = await fetchNumbers();
expect(numbers, hasLength(5));
expect(numbers, contains(3));
expect(numbers.first, equals(1));
});

// 多个异步操作的测试
test('连续异步操作', () async {
var data1 = await fetchUserData(1);
var data2 = await fetchUserData(2);

expect(data1, isNot(equals(data2)));
expect(data1, contains('1'));
expect(data2, contains('2'));
});
});
}

```

00:00 +4: All tests passed!

```

### 运行测试的常用命令

命令功能 dart test运行所有测试 dart test test/calculator_test.dart运行指定测试文件 dart test --name="add"只运行名称包含 "add" 的测试 dart test --concurrency=4并发运行测试（加速） dart test --reporter=expanded详细输出模式 dart test --coverage=coverage生成测试覆盖率数据

### 测试最佳实践

实践说明 每个测试只验证一件事一个 test() 对应一个行为，失败时定位更快 测试名称描述行为而非实现"add() 两个正数相加" 比 "add() 测试" 更好 AAA 模式Arrange（准备）→ Act（执行）→ Assert（断言） 先写失败的测试确认测试能捕获错误，再写代码让它通过 边界条件测试空值、零、负数、极大值、极小值 测试之间保持独立一个测试的结果不应影响另一个测试

测试不是负担，而是保险。你写的测试越多，重构时就越有信心。当一个测试失败时，你不需要猜是哪里出了问题——测试会精确地告诉你。

### 本章小结

本章介绍了 Dart 单元测试的完整流程：test 包的安装、test() 测试用例的编写、expect 断言的使用、group() 分组组织、setUp/tearDown 环境管理以及异步测试。

单元测试是专业开发者的必修课，养成写测试的习惯会让你的代码质量持续提升。

### 全教程总结

恭喜你完成了 Dart 编程语言的入门学习！

让我们回顾一下这 21 章的学习路径：

阶段章节核心内容 一、入门基础1-4 章Dart 概览、环境搭建、第一个程序、基础语法 二、核心类型与控制流5-8 章变量与数据类型、运算符、控制流、集合 三、函数与面向对象9-13 章函数、类与对象、继承与多态、接口与 Mixin、泛型 四、进阶特性14-17 章枚举与符号、异常处理、包与库管理、typedef 五、异步编程与测试18-21 章异步编程、Stream 流、并发与 Isolate、单元测试

掌握了这些内容，你已经具备了使用 Dart 进行日常开发的能力。

下一步可以深入学习 Flutter 框架，将 Dart 知识应用到移动端和 Web 端的 UI 开发中。
