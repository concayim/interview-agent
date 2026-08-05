# C# - 菜鸟教程

Tutorial: https://www.runoob.com/csharp/csharp-tutorial.html

---

## C# 教程

Source: https://www.runoob.com/csharp/csharp-tutorial.html

## C# 教程

C# 是一个简单的、现代的、通用的、面向对象的编程语言，它是由微软（Microsoft）开发的。

本教程将告诉您基础的 C# 编程，同时将向您讲解 C# 编程语言相关的各种先进理念。

现在开始学习 C#！

C# 在线工具

### 谁适合阅读本教程？

本教程有助于初学者理解基础的 C# 编程。在学习完本教程后，您将到达一个中级的 C# 编程水平。

### 阅读本教程前，您需要了解的知识：

C# 编程是基于 C 和 C++ 编程语言的，因此如果您对 C 和 C++ 编程有基本的了解，将有助于您学习 C# 编程语言。

### 编译/执行 C# 程序

菜鸟教程提供了在线的 C# 在线编译环境，您只需进行简单的点击动作，即可在高端的服务器上体验真实的编程经验。这是完全免费的在线工具。

### 实例

using System;
namespace HelloWorldApplication
{
/* 类名为 HelloWorld */
class HelloWorld
{
/* main函数 */
static void Main(string[] args)
{
/* 我的第一个 C# 程序 */
Console.WriteLine("Hello World!");
Console.ReadKey();
}
}
}

运行实例 »

### C# 有用的资源

本教程列出了 C# 网站、书籍和文章。

#### C# 有用的网站

- C# Programming Guide - 介绍了有关关键的 C# 语言特征以及如何通过 .NET 框架访问 C# 的详细信息。
- Visual Studio - 下载作为 C# 集成开发环境的 Visual Studio 的最新版本。
- Go Mono - Mono 是一个允许开发人员简单地创建跨平台应用程序的软件平台。
- C Sharp (programming language) - 维基百科解释 C#（编程语言）。

---

## C# 简介

Source: https://www.runoob.com/csharp/csharp-intro.html

## C# 简介

C#（读作"C Sharp"）是一门由微软开发的编程语言。

C# 诞生于 2000 年前后，由微软工程师 Anders Hejlsberg 主导设计。

C# 是一门面向对象的语言。你可以把面向对象理解为一种思维方式：把现实世界中的事物（比如一个用户、一辆汽车、一个订单）抽象成代码里的对象，再通过操作这些对象来完成各种任务，这种方式让代码更贴近真实业务逻辑，也更容易维护和扩展。

时至今日，C# 已经过十余个版本的迭代演进，功能越来越强大，语法也越来越简洁优雅。

### C# 能做什么？

很多初学者会问：学了 C# 能干什么？答案是：几乎什么都能做。

- 桌面应用：开发 Windows 上的各种软件，比如办公工具、管理系统、工具类应用。
- Web 后端开发：借助 ASP.NET Core 框架，构建高性能的网站和 API 接口，支撑淘宝、微博那量级的请求也不在话下。
- 游戏开发：大名鼎鼎的游戏引擎 Unity 使用的就是 C#，《王者荣耀》《原神》等众多手游背后都有 C# 的身影。
- 移动应用：通过 .NET MAUI 框架，用一套代码同时开发 iOS 和 Android 应用。
- 云计算与微服务：在 Azure 等云平台上构建弹性伸缩的企业级服务。
- 人工智能与数据处理：结合 ML.NET 等库，进行机器学习模型的训练与推理。

### C# 和 .NET 是什么关系？

学 C# 时，你一定会频繁看到一个名词：.NET（读作"dot net"）。很多人刚开始会把它们搞混，这里做个简单类比：

如果把 C# 比作"驾驶技术"，那么 .NET 就是"汽车和道路基础设施"。你用 C# 写出的代码，需要运行在 .NET 这个平台上才能跑起来。

.NET 为 C# 提供了运行时环境（负责执行代码、管理内存）和海量的标准类库（提供现成的工具，比如读写文件、发送网络请求、操作数据库等），让你不必从零造轮子。

早期 .NET 只能在 Windows 上运行，而现在的 .NET 6 / .NET 8 已经完全跨平台，Windows、macOS、Linux 都能跑，这也让 C# 的应用场景大大拓展。

### C# 和其他语言相比怎么样？

作为初学者，你可能听说过 Java、Python、C++ 等语言，那 C# 和它们相比有什么特点？

- 对比 C/C++：C# 语法上借鉴了 C 和 C++ 的风格（大括号、分号等），但屏蔽了指针操作和手动内存管理，入门更安全，不容易犯内存相关的低级错误。
- 对比 Java：两者非常相似，有人说"C# 是微软的 Java"。C# 在语言特性上普遍比 Java 更现代，比如更早引入了 Lambda 表达式、async/await 异步语法等。
- 对比 Python：Python 更偏向数据科学和快速脚本，入门极简但性能较弱；C# 是强类型语言，更适合构建大型、高性能的工程项目，代码在大规模协作时也更易维护。

### C# 适合初学者吗？

完全适合。C# 的设计哲学之一就是"易于学习，难以用错"。它有以下几点对初学者非常友好：

- 语法清晰：代码结构工整，逻辑一目了然，不像 C++ 那样充满"陷阱"。
- 工具链完善：配合 Visual Studio 或 VS Code，代码补全、错误提示、调试工具一应俱全，写代码效率极高。
- 文档丰富：微软官方文档详尽且持续更新，中文资料和社区教程也非常丰富。
- 就业前景好：企业级开发、游戏开发、金融系统等领域对 C# 工程师有大量需求，薪资待遇也颇具竞争力。

总的来说，C# 是一门兼顾易学性与工程深度的语言——你可以用它写出第一个"Hello World"，也可以用它支撑起日均亿级访问的大型系统。无论是想入门编程，还是追求职业发展，C# 都是一个值得投入的选择。

### C# 强大的编程功能

C# 在语法上借鉴了 C 和 C++ 的传统，同时引入了与 Java 相似的面向对象机制，并在此基础上进行了大量创新，提供了一套功能强大且表达力丰富的现代语言特性，深受广大开发者青睐。

以下是 C# 的一些核心功能亮点：

- 布尔条件（Boolean Conditions）：支持简洁的条件判断与逻辑控制流，是所有程序逻辑的基础。
- 自动垃圾回收（Automatic Garbage Collection）：由运行时自动管理内存，开发者无需手动释放内存，大幅降低内存泄漏风险。
- 标准库（Standard Library）：内置极其丰富的基础类库，涵盖文件操作、网络请求、数据结构、加密等常见开发需求，开箱即用。
- 程序集版本控制（Assembly Versioning）：支持组件的独立打包、部署与版本管理，便于大型项目的模块化维护。
- 属性（Properties）与事件（Events）：提供优雅的封装机制和事件驱动编程模型，让代码逻辑更清晰、更安全。
- 委托（Delegates）与事件管理（Events Management）：实现灵活的回调机制与解耦设计，是构建可扩展系统的重要工具。
- 泛型（Generics）：编写类型安全、可复用的通用代码，避免重复劳动，使用方式简单直观。
- 索引器（Indexers）：允许自定义对象像数组一样通过索引访问，提升 API 的表达力与易用性。
- 条件编译（Conditional Compilation）：根据编译符号灵活控制哪些代码参与编译，方便区分开发环境与生产环境。
- 多线程支持（Multithreading）：内置简洁的并发编程模型，配合 async/await 语法，可以轻松构建高响应性的多线程应用。
- LINQ 与 Lambda 表达式：以声明式风格查询和处理集合、数据库、XML 等数据源，代码量少、可读性强，是 C# 最受欢迎的特性之一。
- 深度集成 Windows 与 .NET 生态：与 Windows 平台及 .NET 体系无缝协作，同时支持跨平台部署，兼顾灵活性与生态广度。

---

## C# 开发环境

Source: https://www.runoob.com/csharp/csharp-environment-setup.html

## C# 开发环境

工欲善其事，必先利其器，在正式开始写 C# 代码之前，我们需要先在电脑上搭建好开发环境。

C# 开发环境搭建，整个过程并不复杂——只需要安装两样东西：一个运行平台（.NET SDK）和一个代码编辑器，你就可以开始你的 C# 编程之旅了。

### 第一步：了解 .NET 与 C# 的关系

在安装工具之前，我们先搞清楚一个常见的疑问：C# 和 .NET 是同一个东西吗？

不是。它们的关系可以这样理解：

- C# 是你写代码用的编程语言，就像你用来写文章的"文字"。
- .NET 是运行这些代码的平台和工具集，就像让文字得以印刷、传播的"印刷机和纸张"。

你用 C# 写好的程序，必须依赖 .NET 才能在电脑上运行。因此，搭建 C# 开发环境的第一步，就是安装 .NET SDK。

.NET 目前的主流版本是 .NET 8（长期支持版），它支持 Windows、macOS 和 Linux 三大平台，不再像早年那样只能运行在 Windows 上。

### 第二步：安装 .NET SDK

.NET SDK（软件开发工具包）包含了编译器、运行时和各种开发工具，是 C# 开发的基础。安装步骤如下：

- 访问 .NET 官方下载页面：https://dotnet.microsoft.com/download
- 选择最新的 LTS（长期支持）版本，推荐初学者优先选择 LTS 版，稳定且资料丰富。
- 根据你的操作系统（Windows / macOS / Linux）下载对应的安装包，按提示完成安装。

安装完成后，打开命令行（Windows 用"命令提示符"或"PowerShell"，macOS/Linux 用"终端"），输入以下命令验证是否安装成功：

```
dotnet --version
```

如果输出了一个版本号（如 `8.0.xxx`），说明 .NET SDK 已经正确安装。

### 第三步：选择一款代码编辑器

有了 .NET SDK，你理论上已经可以用记事本写代码、用命令行编译运行了。但实际开发中，我们都会使用专门的代码编辑器或 IDE（集成开发环境），它们提供代码补全、错误提示、调试等功能，大大提升开发效率。

对于 C# 初学者，推荐以下两款工具：

#### ① Visual Studio（推荐 Windows 用户）

Visual Studio 是微软为 C# 量身打造的旗舰 IDE，功能最为完整，调试体验极佳，是目前企业中使用最广泛的 C# 开发工具。

- 优点：功能强大，集代码编写、调试、测试、发布于一体，特别适合开发 Windows 桌面应用和企业级项目。
- 缺点：安装包较大（几 GB），启动较慢，仅支持 Windows 和 macOS。
- 价格：个人学习使用 Community（社区版）免费，功能已经非常完整。
- 下载地址：https://visualstudio.microsoft.com/zh-hans/downloads/

💡 安装时，在"工作负载"页面勾选 "ASP.NET 和 Web 开发" 或 ".NET 桌面开发"，即可自动安装 C# 所需的全部组件，无需单独配置。

#### ② Visual Studio Code（推荐跨平台用户）

VS Code 是一款轻量级的代码编辑器，体积小、启动快，支持 Windows、macOS 和 Linux。通过安装 C# 扩展插件，同样可以获得完善的 C# 开发体验。

- 优点：轻量灵活，跨平台，插件生态丰富，适合追求简洁环境的开发者。
- 缺点：需要手动安装插件，部分高级调试功能不如 Visual Studio 完善。
- 价格：完全免费开源。
- 下载地址：https://code.visualstudio.com/

💡 安装 VS Code 后，在扩展商店中搜索并安装 C# Dev Kit 插件（微软官方出品），即可获得智能提示、调试、项目管理等完整功能。

#### 如何选择？

- 如果你使用 Windows，且主要目标是学习 C# 基础或开发 Windows 应用 → 推荐 Visual Studio Community。
- 如果你使用 macOS 或 Linux，或者希望编辑器更轻量 → 推荐 Visual Studio Code + C# Dev Kit。

#### 更多相关工具

### 第四步：写下你的第一个 C# 程序

环境搭好之后，我们来验证一下——用命令行创建并运行一个最简单的 C# 项目：

打开命令行，依次执行以下命令：

```

dotnet new console -n HelloWorld
cd HelloWorld
dotnet run

```

如果你看到输出：

```
Hello, World!
```

恭喜你！你的 C# 开发环境已经搭建成功，并且成功运行了第一个程序。

### 在 macOS 和 Linux 上开发 C#

如前所述，现代 .NET（.NET 5 及以上版本）已经原生支持跨平台运行，macOS 和 Linux 用户直接安装 .NET SDK 即可，与 Windows 体验基本一致。

此外，你可能还会听到一个叫 Mono 的东西。Mono 是早期社区推出的 .NET 开源跨平台实现，曾经是 Linux 和 macOS 上运行 C# 的主要方式。但随着微软官方 .NET 的全面跨平台化，Mono 的使用场景已经大幅缩减，目前主要用于 Unity 游戏引擎等特定场合。对于新项目，建议直接使用官方 .NET SDK，无需安装 Mono。

如需了解 Mono 的更多信息，可访问：http://www.mono-project.com/

---

## C# 程序结构

Source: https://www.runoob.com/csharp/csharp-program-structure.html

## C# 程序结构

在我们学习 C# 编程语言的基础构件块之前，让我们先看一下 C# 的最小的程序结构，以便作为接下来章节的参考。

### C# Hello World 实例

一个 C# 程序主要包括以下部分：

- 命名空间声明（Namespace declaration）
- 一个 class
- Class 方法
- Class 属性
- 一个 Main 方法
- 语句（Statements）& 表达式（Expressions）
- 注释

C# 文件的后缀为 .cs。

以下创建一个 test.cs 文件，文件包含了可以打印出 "Hello World" 的简单代码：

### test.cs 文件代码：

using System;
namespace HelloWorldApplication
{
class HelloWorld
{
static void Main(string[] args)
{
/* 我的第一个 C# 程序*/
Console.WriteLine("Hello World");
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Hello World

```

让我们看一下上面程序的各个部分：

- 程序的第一行 using System; - using 关键字用于在程序中包含 System 命名空间。 一个程序一般有多个 using 语句。
- 下一行是 namespace 声明。一个 namespace 里包含了一系列的类。HelloWorldApplication 命名空间包含了类 HelloWorld。
- 下一行是 class 声明。类 HelloWorld 包含了程序使用的数据和方法声明。类一般包含多个方法。方法定义了类的行为。在这里，HelloWorld 类只有一个 Main 方法。
- 下一行定义了 Main 方法，是所有 C# 程序的 入口点。Main 方法说明当执行时 类将做什么动作。
- 下一行 /*...*/ 将会被编译器忽略，且它会在程序中添加额外的 注释。
- Main 方法通过语句 Console.WriteLine("Hello World"); 指定了它的行为。

WriteLine 是一个定义在 System 命名空间中的 Console 类的一个方法。该语句会在屏幕上显示消息 "Hello World"。
- 最后一行 Console.ReadKey(); 是针对 VS.NET 用户的。这使得程序会等待一个按键的动作，防止程序从 Visual Studio .NET 启动时屏幕会快速运行并关闭。

以下几点值得注意：

- C# 是大小写敏感的。
- 所有的语句和表达式必须以分号（;）结尾。
- 程序的执行从 Main 方法开始。
- 与 Java 不同的是，文件名可以不同于类的名称。

### 编译 & 执行 C# 程序

如果您使用 Visual Studio.Net 编译和执行 C# 程序，请按下面的步骤进行：

- 启动 Visual Studio。
- 在菜单栏上，选择 File -> New -> Project。
- 从模板中选择 Visual C#，然后选择 Windows。
- 选择 Console Application。
- 为您的项目制定一个名称，然后点击 OK 按钮。
- 新项目会出现在解决方案资源管理器（Solution Explorer）中。
- 在代码编辑器（Code Editor）中编写代码。
- 点击 Run 按钮或者按下 F5 键来运行程序。会出现一个命令提示符窗口（Command Prompt window），显示 Hello World。

您也可以使用命令行代替 Visual Studio IDE 来编译 C# 程序：

- 打开一个文本编辑器，添加上面提到的代码。
- 保存文件为 helloworld.cs。
- 打开命令提示符工具，定位到文件所保存的目录。
- 键入 csc helloworld.cs 并按下 enter 键来编译代码。
- 如果代码没有错误，命令提示符会进入下一行，并生成 helloworld.exe 可执行文件。
- 接下来，键入 helloworld 来执行程序。
- 您将看到 "Hello World" 打印在屏幕上。

---

## C# 基本语法

Source: https://www.runoob.com/csharp/csharp-basic-syntax.html

## C# 基本语法

C# 是一种面向对象的编程语言。在面向对象的程序设计方法中，程序由各种相互交互的对象组成。相同种类的对象通常具有相同的类型，或者说，是在相同的 class 中。

例如，以 Rectangle（矩形）对象为例。它具有 length 和 width 属性。根据设计，它可能需要接受这些属性值、计算面积和显示细节。

让我们来看看一个 Rectangle（矩形）类的实现，并借此讨论 C# 的基本语法：

### 实例

using System;
namespace RectangleApplication
{
class Rectangle
{
// 成员变量
double length;
double width;
public void Acceptdetails()
{
length = 4.5;
width = 3.5;
}
public double GetArea()
{
return length * width;
}
public void Display()
{
Console.WriteLine("Length: {0}", length);
Console.WriteLine("Width: {0}", width);
Console.WriteLine("Area: {0}", GetArea());
}
}

class ExecuteRectangle
{
static void Main(string[] args)
{
Rectangle r = new Rectangle();
r.Acceptdetails();
r.Display();
Console.ReadLine();
}
}
}

尝试一下 »

当上面的代码被编译和执行时，它会产生下列结果：

```

Length: 4.5
Width: 3.5
Area: 15.75

```

### using 关键字

在任何 C# 程序中的第一条语句都是：

```

using System;

```

using 关键字用于在程序中包含命名空间。一个程序可以包含多个 using 语句。

### class 关键字

class 关键字用于声明一个类。

### C# 中的注释

注释是用于解释代码。编译器会忽略注释的条目。在 C# 程序中，多行注释以 /* 开始，并以字符 */ 终止，如下所示：

```

/* 这个程序演示
C# 的注释
使用 */

```

单行注释是用 // 符号表示。例如：

```

// 这一行是注释

```

### 成员变量

变量是类的属性或数据成员，用于存储数据。在上面的程序中，Rectangle 类有两个成员变量，名为 length 和 width。

### 成员函数

函数是一系列执行指定任务的语句。类的成员函数是在类内声明的。我们举例的类 Rectangle 包含了三个成员函数： AcceptDetails、GetArea 和 Display。

### 实例化一个类

在上面的程序中，类 ExecuteRectangle 是一个包含 Main() 方法和实例化 Rectangle 类的类。

### 标识符

标识符是用来识别类、变量、函数或任何其它用户定义的项目。在 C# 中，类的命名必须遵循如下基本规则：

- 标识符必须以字母、下划线或 @ 开头，后面可以跟一系列的字母、数字（ 0 - 9 ）、下划线（ _ ）、@。
- 标识符中的第一个字符不能是数字。
- 标识符必须不包含任何嵌入的空格或符号，比如 ? - +! # % ^ & * ( ) [ ] { } . ; : " ' / \。
- 标识符不能是 C# 关键字。除非它们有一个 @ 前缀。 例如，@if 是有效的标识符，但 if 不是，因为 if 是关键字。
- 标识符必须区分大小写。大写字母和小写字母被认为是不同的字母。
- 不能与C#的类库名称相同。

### C# 关键字

关键字是 C# 编译器预定义的保留字。这些关键字不能用作标识符，但是，如果您想使用这些关键字作为标识符，可以在关键字前面加上 @ 字符作为前缀。

在 C# 中，有些关键字在代码的上下文中有特殊的意义，如 get 和 set，这些被称为上下文关键字（contextual keywords）。

下表列出了 C# 中的保留关键字（Reserved Keywords）和上下文关键字（Contextual Keywords）：

保留关键字 abstractasbaseboolbreakbytecase catchcharcheckedclassconstcontinuedecimal defaultdelegatedodoubleelseenumevent explicitexternfalsefinallyfixedfloatfor foreachgotoifimplicitinin (generic
modifier)int interfaceinternalislocklongnamespacenew nullobjectoperatoroutout
(generic
modifier)overrideparams privateprotectedpublicreadonlyrefreturnsbyte sealedshortsizeofstackallocstaticstringstruct switchthisthrowtruetrytypeofuint ulonguncheckedunsafeushortusingvirtualvoid volatilewhile 上下文关键字 addaliasascendingdescendingdynamicfromget globalgroupintojoinletorderbypartial
(type) partial
(method)removeselectset

### 顶级语句（Top-Level Statements）

在 C# 9.0 版本中，引入了顶级语句（Top-Level Statements）的概念，这是一种新的编程范式，允许开发者在文件的顶层直接编写语句，而不需要将它们封装在方法或类中。

特点：

- 无需类或方法：顶级语句允许你直接在文件的顶层编写代码，无需定义类或方法。
- 文件作为入口点：包含顶级语句的文件被视为程序的入口点，类似于 C# 之前的 `Main` 方法。
- 自动 `Main` 方法：编译器会自动生成一个 `Main` 方法，并将顶级语句作为 `Main` 方法的主体。
- 支持局部函数：尽管不需要定义类，但顶级语句的文件中仍然可以定义局部函数。
- 更好的可读性：对于简单的脚本或工具，顶级语句提供了更好的可读性和简洁性。
- 适用于小型项目：顶级语句非常适合小型项目或脚本，可以快速编写和运行代码。
- 与现有代码兼容：顶级语句可以与现有的 C# 代码库一起使用，不会影响现有代码。

传统 C# 代码 - 在使用顶级语句之前，你必须像这样编写一个 C# 程序：

### 实例

using System;

namespace MyApp
{
class Program
{
static void Main(string[] args)
{
Console.WriteLine("Hello, World!");
}
}
}

使用顶级语句的 C# 代码 - 使用顶级语句，可以简化为：

### 实例

using System;

Console.WriteLine("Hello, World!");

顶级语句支持所有常见的 C# 语法，包括声明变量、定义方法、处理异常等。

### 实例

using System;
using System.Linq;

// 顶级语句中的变量声明
int number = 42;
string message = "The answer to life, the universe, and everything is";

// 输出变量
Console.WriteLine($"{message} {number}.");

// 定义和调用方法
int Add(int a, int b) => a + b;
Console.WriteLine($"Sum of 1 and 2 is {Add(1, 2)}.");

// 使用 LINQ
var numbers = new[] { 1, 2, 3, 4, 5 };
var evens = numbers.Where(n => n % 2 == 0).ToArray();
Console.WriteLine("Even numbers: " + string.Join(", ", evens));

// 异常处理
try
{
int zero = 0;
int result = number / zero;
}
catch (DivideByZeroException ex)
{
Console.WriteLine("Error: " + ex.Message);
}

#### 注意事项

- 文件限制：顶级语句只能在一个源文件中使用。如果在一个项目中有多个使用顶级语句的文件，会导致编译错误。
- 程序入口：如果使用顶级语句，则该文件会隐式地包含 Main 方法，并且该文件将成为程序的入口点。
- 作用域限制：顶级语句中的代码共享一个全局作用域，这意味着可以在顶级语句中定义的变量和方法可以在整个文件中访问。

顶级语句在简化代码结构、降低学习难度和加快开发速度方面具有显著优势，特别适合于编写简单程序和脚本。

---

## C# 数据类型

Source: https://www.runoob.com/csharp/csharp-data-types.html

## C# 数据类型

在 C# 中，变量分为以下几种类型：

- 值类型（Value types）
- 引用类型（Reference types）
- 指针类型（Pointer types）

### 值类型（Value types）

值类型变量可以直接分配给一个值。它们是从类 System.ValueType 中派生的。

值类型直接包含数据。比如 int、char、float，它们分别存储数字、字符、浮点数。当您声明一个 int 类型时，系统分配内存来存储值。

下表列出了 C# 2010 中可用的值类型：

类型描述范围默认值 bool布尔值True 或 FalseFalse byte8 位无符号整数0 到 2550 char16 位 Unicode 字符U +0000 到 U +ffff'\0' decimal128 位精确的十进制值，28-29 有效位数(-7.9 x 1028 到 7.9 x 1028) / 100 到 28 0.0M double64 位双精度浮点型(+/-)5.0 x 10-324 到 (+/-)1.7 x 103080.0D float32 位单精度浮点型-3.4 x 1038 到 + 3.4 x 10380.0F int32 位有符号整数类型-2,147,483,648 到 2,147,483,6470 long64 位有符号整数类型-9,223,372,036,854,775,808 到 9,223,372,036,854,775,807 0L sbyte8 位有符号整数类型-128 到 1270 short16 位有符号整数类型-32,768 到 32,7670 uint32 位无符号整数类型0 到 4,294,967,2950 ulong64 位无符号整数类型0 到 18,446,744,073,709,551,6150 ushort16 位无符号整数类型0 到 65,5350

如需得到一个类型或一个变量在特定平台上的准确尺寸，可以使用 sizeof 方法。表达式 sizeof(type) 产生以字节为单位存储对象或类型的存储尺寸。下面举例获取任何机器上 int 类型的存储尺寸：

### 实例

using System;

namespace DataTypeApplication
{
class Program
{
static void Main(string[] args)
{
Console.WriteLine("Size of int: {0}", sizeof(int));
Console.ReadLine();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Size of int: 4

```

### 引用类型（Reference types）

引用类型不包含存储在变量中的实际数据，但它们包含对变量的引用。

换句话说，它们指的是一个内存位置。使用多个变量时，引用类型可以指向一个内存位置。如果内存位置的数据是由一个变量改变的，其他变量会自动反映这种值的变化。内置的 引用类型有：object、dynamic 和 string。

#### 对象（Object）类型

对象（Object）类型 是 C# 通用类型系统（Common Type System - CTS）中所有数据类型的终极基类。Object 是 System.Object 类的别名。所以对象（Object）类型可以被分配任何其他类型（值类型、引用类型、预定义类型或用户自定义类型）的值。但是，在分配值之前，需要先进行类型转换。

当一个值类型转换为对象类型时，则被称为 装箱；另一方面，当一个对象类型转换为值类型时，则被称为 拆箱。

```

object obj;
obj = 100; // 这是装箱

```

#### 动态（Dynamic）类型

您可以存储任何类型的值在动态数据类型变量中。这些变量的类型检查是在运行时发生的。

声明动态类型的语法：

```

dynamic <variable_name> = value;

```

例如：

```

dynamic d = 20;

```

动态类型与对象类型相似，但是对象类型变量的类型检查是在编译时发生的，而动态类型变量的类型检查是在运行时发生的。

#### 字符串（String）类型

字符串（String）类型 允许您给变量分配任何字符串值。字符串（String）类型是 System.String 类的别名。它是从对象（Object）类型派生的。字符串（String）类型的值可以通过两种形式进行分配：引号和 @引号。

例如：

```

String str = "runoob.com";

```

一个 @引号字符串：

```

@"runoob.com";

```

C# string 字符串的前面可以加 @（称作"逐字字符串"）将转义字符（\）当作普通字符对待，比如：

```

string str = @"C:\Windows";

```

等价于：

```

string str = "C:\\Windows";

```

@ 字符串中可以任意换行，换行符及缩进空格都计算在字符串长度之内。

```

string str = @"<script type=""text/javascript"">
<!--
-->
</script>";

```

用户自定义引用类型有：class、interface 或 delegate。我们将在以后的章节中讨论这些类型。

### 指针类型（Pointer types）

指针类型变量存储另一种类型的内存地址。C# 中的指针与 C 或 C++ 中的指针有相同的功能。

声明指针类型的语法：

```

type* identifier;

```

例如：

```

char* cptr;
int* iptr;

```

我们将在章节"不安全的代码"中讨论指针类型。

---

## C# 类型转换

Source: https://www.runoob.com/csharp/csharp-type-conversion.html

## C# 类型转换

在 C# 中，类型转换是将一个数据类型的值转换为另一个数据类型的过程。

C# 中的类型转换可以分为两种：隐式类型转换和显式类型转换（也称为强制类型转换）。

#### 隐式类型转换

隐式转换是不需要编写代码来指定的转换，编译器会自动进行。

隐式转换是指将一个较小范围的数据类型转换为较大范围的数据类型时，编译器会自动完成类型转换，这些转换是 C# 默认的以安全方式进行的转换, 不会导致数据丢失。

例如，从 int 到 long，从 float 到 double 等。

从小的整数类型转换为大的整数类型，从派生类转换为基类。将一个 byte 类型的变量赋值给 int 类型的变量，编译器会自动将 byte 类型转换为 int 类型，不需要显示转换。

### 实例

byte b = 10;
int i = b; // 隐式转换，不需要显式转换

将一个整数赋值给一个长整数，或者将一个浮点数赋值给一个双精度浮点数，这种转换不会导致数据丢失：

### 实例

int intValue = 42;
long longValue = intValue; // 隐式转换，从 int 到 long

#### 显式转换

显式类型转换，即强制类型转换，需要程序员在代码中明确指定。

显式转换是指将一个较大范围的数据类型转换为较小范围的数据类型时，或者将一个对象类型转换为另一个对象类型时，需要使用强制类型转换符号进行显示转换，强制转换会造成数据丢失。

例如，将一个 int 类型的变量赋值给 byte 类型的变量，需要显示转换。

### 实例

int i = 10;
byte b = (byte)i; // 显式转换，需要使用强制类型转换符号

强制转换为整数类型：

### 实例

double doubleValue = 3.14;
int intValue = (int)doubleValue; // 强制从 double 到 int，数据可能损失小数部分

强制转换为浮点数类型：

### 实例

int intValue = 42;
float floatValue = (float)intValue; // 强制从 int 到 float，数据可能损失精度

强制转换为字符串类型：

### 实例

int intValue = 123;
string stringValue = intValue.ToString(); // 将 int 转换为字符串

下面的实例显示了一个显式的类型转换：

### 实例

using System;

namespace TypeConversionApplication
{
class ExplicitConversion
{
static void Main(string[] args)
{
double d = 5673.74;
int i;

// 强制转换 double 为 int
i = (int)d;
Console.WriteLine(i);
Console.ReadKey();

}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

5673

```

### C# 类型转换方法

C# 提供了下列内置的类型转换方法：

序号方法 & 描述 1ToBoolean
如果可能的话，把类型转换为布尔型。 2ToByte
把类型转换为字节类型。 3ToChar
如果可能的话，把类型转换为单个 Unicode 字符类型。 4ToDateTime
把类型（整数或字符串类型）转换为 日期-时间 结构。 5ToDecimal
把浮点型或整数类型转换为十进制类型。 6ToDouble
把类型转换为双精度浮点型。 7ToInt16
把类型转换为 16 位整数类型。 8ToInt32
把类型转换为 32 位整数类型。 9ToInt64
把类型转换为 64 位整数类型。 10ToSbyte
把类型转换为有符号字节类型。 11ToSingle
把类型转换为小浮点数类型。 12ToString
把类型转换为字符串类型。 13ToType
把类型转换为指定类型。 14ToUInt16
把类型转换为 16 位无符号整数类型。 15ToUInt32
把类型转换为 32 位无符号整数类型。 16ToUInt64
把类型转换为 64 位无符号整数类型。

这些方法都定义在 System.Convert 类中，使用时需要包含 System 命名空间。它们提供了一种安全的方式来执行类型转换，因为它们可以处理 null值，并且会抛出异常，如果转换不可能进行。

例如，使用 Convert.ToInt32 方法将字符串转换为整数：

```

string str = "123";
int number = Convert.ToInt32(str); // 转换成功，number为123
```

如果字符串不是有效的整数表示，Convert.ToInt32 将抛出 FormatException。

下面的实例把不同值的类型转换为字符串类型：

### 实例

using System;

namespace TypeConversionApplication
{
class StringConversion
{
static void Main(string[] args)
{
// 定义一个整型变量
int i = 75;

// 定义一个浮点型变量
float f = 53.005f;

// 定义一个双精度浮点型变量
double d = 2345.7652;

// 定义一个布尔型变量
bool b = true;

// 将整型变量转换为字符串并输出
Console.WriteLine(i.ToString());

// 将浮点型变量转换为字符串并输出
Console.WriteLine(f.ToString());

// 将双精度浮点型变量转换为字符串并输出
Console.WriteLine(d.ToString());

// 将布尔型变量转换为字符串并输出
Console.WriteLine(b.ToString());

// 等待用户按键后关闭控制台窗口
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

75
53.005
2345.7652
True

```

在进行类型转换时需要注意以下几点：

- 隐式转换只能将较小范围的数据类型转换为较大范围的数据类型，不能将较大范围的数据类型转换为较小范围的数据类型；
- 显式转换可能会导致数据丢失或精度降低，需要进行数据类型的兼容性检查；
- 对于对象类型的转换，需要进行类型转换的兼容性检查和类型转换的安全性检查。

### 类型转换方法

C# 提供了多种类型转换方法，例如使用 Convert 类、Parse 方法和 TryParse 方法，这些方法可以帮助处理不同的数据类型之间的转换。

#### 使用 Convert 类

Convert 类提供了一组静态方法，可以在各种基本数据类型之间进行转换。

### 实例

string str = "123";
int num = Convert.ToInt32(str);

#### 使用 Parse 方法

Parse 方法用于将字符串转换为对应的数值类型，如果转换失败会抛出异常。

### 实例

string str = "123.45";
double d = double.Parse(str);

#### 使用 TryParse 方法

TryParse 方法类似于 Parse，但它不会抛出异常，而是返回一个布尔值指示转换是否成功。

### 实例

string str = "123.45";
double d;
bool success = double.TryParse(str, out d);

if (success) {
Console.WriteLine("转换成功: " + d);
} else {
Console.WriteLine("转换失败");
}

### 自定义类型转换

C# 还允许你定义自定义类型转换操作，通过在类型中定义 implicit 或 explicit 关键字。

### 实例

using System;

public class Fahrenheit
{
public double Degrees { get; set; }

public Fahrenheit(double degrees)
{
Degrees = degrees;
}

// 隐式转换从Fahrenheit到double
public static implicit operator double(Fahrenheit f)
{
return f.Degrees;
}

// 显式转换从double到Fahrenheit
public static explicit operator Fahrenheit(double d)
{
return new Fahrenheit(d);
}
}

public class Program
{
public static void Main()
{
Fahrenheit f = new Fahrenheit(98.6);
Console.WriteLine("Fahrenheit object: " + f.Degrees + " degrees");

double temp = f; // 隐式转换
Console.WriteLine("After implicit conversion to double: " + temp + " degrees");

Fahrenheit newF = (Fahrenheit)temp; // 显式转换
Console.WriteLine("After explicit conversion back to Fahrenheit: " + newF.Degrees + " degrees");
}
}

以上例子中，我们定义了一个 Fahrenheit 类，并实现了从 Fahrenheit 到 double 的隐式转换和从 double 到 Fahrenheit 的显式转换。

输出结果将显示如下：

```
Fahrenheit object: 98.6 degrees
After implicit conversion to double: 98.6 degrees
After explicit conversion back to Fahrenheit: 98.6 degrees
```

### 总结

在 C# 中，内置的类型转换方法主要通过以下几种方式实现：隐式转换、显式转换（强制转换）、使用 Convert 类的方法、Parse 方法和 TryParse 方法，这些方法广泛应用于不同数据类型之间的转换。

以下是 C# 内置类型转换方法的表格：

方法类别方法描述隐式转换自动进行的转换无需显式指定，通常用于安全的类型转换，如从较小类型到较大类型显式转换（强制转换）`(type)value`需要显式指定，通常用于可能导致数据丢失或转换失败的情况`Convert` 类方法`Convert.ToBoolean(value)`将指定类型转换为 `Boolean``Convert.ToByte(value)`将指定类型转换为 `Byte``Convert.ToChar(value)`将指定类型转换为 `Char``Convert.ToDateTime(value)`将指定类型转换为 `DateTime``Convert.ToDecimal(value)`将指定类型转换为 `Decimal``Convert.ToDouble(value)`将指定类型转换为 `Double``Convert.ToInt16(value)`将指定类型转换为 `Int16`（短整型）`Convert.ToInt32(value)`将指定类型转换为 `Int32`（整型）`Convert.ToInt64(value)`将指定类型转换为 `Int64`（长整型）`Convert.ToSByte(value)`将指定类型转换为 `SByte``Convert.ToSingle(value)`将指定类型转换为 `Single`（单精度浮点型）`Convert.ToString(value)`将指定类型转换为 `String``Convert.ToUInt16(value)`将指定类型转换为 `UInt16`（无符号短整型）`Convert.ToUInt32(value)`将指定类型转换为 `UInt32`（无符号整型）`Convert.ToUInt64(value)`将指定类型转换为 `UInt64`（无符号长整型）`Parse` 方法`Boolean.Parse(string)`将字符串解析为 `Boolean``Byte.Parse(string)`将字符串解析为 `Byte``Char.Parse(string)`将字符串解析为 `Char``DateTime.Parse(string)`将字符串解析为 `DateTime``Decimal.Parse(string)`将字符串解析为 `Decimal``Double.Parse(string)`将字符串解析为 `Double``Int16.Parse(string)`将字符串解析为 `Int16``Int32.Parse(string)`将字符串解析为 `Int32``Int64.Parse(string)`将字符串解析为 `Int64``SByte.Parse(string)`将字符串解析为 `SByte``Single.Parse(string)`将字符串解析为 `Single``UInt16.Parse(string)`将字符串解析为 `UInt16``UInt32.Parse(string)`将字符串解析为 `UInt32``UInt64.Parse(string)`将字符串解析为 `UInt64``TryParse` 方法`Boolean.TryParse(string, out bool)`尝试将字符串解析为 `Boolean`，返回布尔值表示是否成功`Byte.TryParse(string, out byte)`尝试将字符串解析为 `Byte`，返回布尔值表示是否成功`Char.TryParse(string, out char)`尝试将字符串解析为 `Char`，返回布尔值表示是否成功`DateTime.TryParse(string, out DateTime)`尝试将字符串解析为 `DateTime`，返回布尔值表示是否成功`Decimal.TryParse(string, out decimal)`尝试将字符串解析为 `Decimal`，返回布尔值表示是否成功`Double.TryParse(string, out double)`尝试将字符串解析为 `Double`，返回布尔值表示是否成功`Int16.TryParse(string, out short)`尝试将字符串解析为 `Int16`，返回布尔值表示是否成功`Int32.TryParse(string, out int)`尝试将字符串解析为 `Int32`，返回布尔值表示是否成功`Int64.TryParse(string, out long)`尝试将字符串解析为 `Int64`，返回布尔值表示是否成功`SByte.TryParse(string, out sbyte)`尝试将字符串解析为 `SByte`，返回布尔值表示是否成功`Single.TryParse(string, out float)`尝试将字符串解析为 `Single`，返回布尔值表示是否成功`UInt16.TryParse(string, out ushort)`尝试将字符串解析为 `UInt16`，返回布尔值表示是否成功`UInt32.TryParse(string, out uint)`尝试将字符串解析为 `UInt32`，返回布尔值表示是否成功`UInt64.TryParse(string, out ulong)`尝试将字符串解析为 `UInt64`，返回布尔值表示是否成功

---

## C# 变量

Source: https://www.runoob.com/csharp/csharp-variables.html

## C# 变量

一个变量只不过是一个供程序操作的存储区的名字。

在 C# 中，变量是用于存储和表示数据的标识符，在声明变量时，您需要指定变量的类型，并且可以选择性地为其分配一个初始值。

在 C# 中，每个变量都有一个特定的类型，类型决定了变量的内存大小和布局，范围内的值可以存储在内存中，可以对变量进行一系列操作。

我们已经讨论了各种数据类型。C# 中提供的基本的值类型大致可以分为以下几类：

类型举例 整数类型sbyte、byte、short、ushort、int、uint、long、ulong 和 char 浮点型float, double 十进制类型decimal 布尔类型true 或 false 值，指定的值 空字符串 string 空类型可为空值的数据类型

C# 允许定义其他值类型的变量，比如 enum，也允许定义引用类型变量，比如 class。这些我们将在以后的章节中进行讨论。在本章节中，我们只研究基本变量类型。

C# 4.0引入了动态类型 (dynamic)，它允许在运行时推断变量的类型。这在一些特殊情况下很有用，但通常最好使用静态类型以获得更好的性能和编译时类型检查。

```
dynamic dynamicVariable = "This can be any type";
```

### C# 中的变量定义

C# 中变量定义的语法：

```

<data_type> <variable_list>;

```

在这里，data_type 必须是一个有效的 C# 数据类型，可以是 char、int、float、double 或其他用户自定义的数据类型。variable_list 可以由一个或多个用逗号分隔的标识符名称组成。

一些有效的变量定义如下所示：

```

int i, j, k;
char c, ch;
float f, salary;
double d;

```

您可以在变量定义时进行初始化：

```

int i = 100;

```

#### 变量的命名规则

在 C# 中，变量的命名需要遵循一些规则：

- 变量名可以包含字母、数字和下划线。
- 变量名必须以字母或下划线开头。
- 变量名区分大小写。
- 避免使用 C# 的关键字作为变量名。

```
int myVariable = 10;
string _userName = "John";
```

### C# 中的变量初始化

变量通过在等号后跟一个常量表达式进行初始化（赋值）。初始化的一般形式为：

```

variable_name = value;

```

变量可以在声明时被初始化（指定一个初始值）。初始化由一个等号后跟一个常量表达式组成，如下所示：

```

<data_type> <variable_name> = value;

```

一些实例：

```

int d = 3, f = 5; /* 初始化 d 和 f. */
byte z = 22; /* 初始化 z. */
double pi = 3.14159; /* 声明 pi 的近似值 */
char x = 'x'; /* 变量 x 的值为 'x' */

```

正确地初始化变量是一个良好的编程习惯，否则有时程序会产生意想不到的结果。

请看下面的实例，使用了各种类型的变量：

### 实例

using System;
namespace VariableDefinition
{
class Program
{
static void Main(string[] args)
{
short a;
int b ;
double c;

/* 实际初始化 */
a = 10;
b = 20;
c = a + b;
Console.WriteLine("a = {0}, b = {1}, c = {2}", a, b, c);
Console.ReadLine();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

a = 10, b = 20, c = 30

```

### 接受来自用户的值 System 命名空间中的 Console 类提供了一个函数 ReadLine()，用于接收来自用户的输入，并把它存储到一个变量中。

例如：

```

int num;
num = Convert.ToInt32(Console.ReadLine());

```

函数 Convert.ToInt32() 把用户输入的数据转换为 int 数据类型，因为 Console.ReadLine() 只接受字符串格式的数据。

### C# 中的 Lvalues 和 Rvalues

C# 中的两种表达式：

- lvalue：lvalue 表达式可以出现在赋值语句的左边或右边。
- rvalue：rvalue 表达式可以出现在赋值语句的右边，不能出现在赋值语句的左边。

变量是 lvalue 的，所以可以出现在赋值语句的左边。数值是 rvalue 的，因此不能被赋值，不能出现在赋值语句的左边。下面是一个有效的语句：

```

int g = 20;

```

下面是一个无效的语句，会产生编译时错误：

```

10 = 20;

```

---

## C# 变量作用域

Source: https://www.runoob.com/csharp/csharp-variable-scope.html

## C# 变量作用域

在 C# 中，变量的作用域定义了变量的可见性和生命周期。

变量的作用域通常由花括号 {} 定义的代码块来确定。

以下是关于C#变量作用域的一些基本规则：

#### 局部变量

在方法、循环、条件语句等代码块内声明的变量是局部变量，它们只在声明它们的代码块中可见。

### 实例

void MyMethod()
{
int localVar = 10; // 局部变量
// ...
}
// localVar 在这里不可见

#### 块级作用域

在 C# 7及更高版本中，引入了块级作用域，即使用大括号 {} 创建的任何块都可以定义变量的作用域。

### 实例

{
int blockVar = 20; // 块级作用域
// ...
}
// blockVar 在这里不可见

#### 方法参数作用域

方法的参数也有其自己的作用域，它们在整个方法中都是可见的。

### 实例

void MyMethod(int parameter)
{
// parameter 在整个方法中可见
// ...
}

#### 全局变量

在类的成员级别定义的变量是成员变量，它们在整个类中可见，如果在命名空间级别定义，那么它们在整个命名空间中可见。

### 实例

class MyClass
{
int memberVar = 30; // 成员变量，在整个类中可见
}

#### 静态变量作用域

静态变量是在类级别上声明的，但它们的作用域也受限于其定义的类。

### 实例

class MyClass
{
static int staticVar = 40; // 静态变量，在整个类中可见
}

#### 循环变量作用域

在 for 循环中声明的循环变量在循环体内可见。

### 实例

for (int i = 0; i < 5; i++)
{
// i 在循环体内可见
}
// i 在这里不可见

总体而言，变量的作用域有助于管理变量的可见性和生命周期，确保变量在其有效范围内使用，也有助于防止命名冲突。

---

## C# 常量

Source: https://www.runoob.com/csharp/csharp-constants.html

## C# 常量

常量是固定值，程序执行期间不会改变。常量可以是任何基本数据类型，比如整数常量、浮点常量、字符常量或者字符串常量，还有枚举常量。

常量可以被当作常规的变量，只是它们的值在定义后不能被修改。

### 整数常量

整数常量可以是十进制、八进制或十六进制的常量。前缀指定基数：0x 或 0X 表示十六进制，0 表示八进制，没有前缀则表示十进制。

整数常量也可以有后缀，可以是 U 和 L 的组合，其中，U 和 L 分别表示 unsigned 和 long。后缀可以是大写或者小写，多个后缀以任意顺序进行组合。

这里有一些整数常量的实例：

```

212 /* 合法 */
215u /* 合法 */
0xFeeL /* 合法 */
078 /* 非法：8 不是一个八进制数字 */
032UU /* 非法：不能重复后缀 */

```

以下是各种类型的整数常量的实例：

```

85 /* 十进制 */
0213 /* 八进制 */
0x4b /* 十六进制 */
30 /* int */
30u /* 无符号 int */
30l /* long */
30ul /* 无符号 long */

```

### 浮点常量

一个浮点常量是由整数部分、小数点、小数部分和指数部分组成。您可以使用小数形式或者指数形式来表示浮点常量。

这里有一些浮点常量的实例：

```

3.14159 /* 合法 */
314159E-5L /* 合法 */
510E /* 非法：不完全指数 */
210f /* 非法：没有小数或指数 */
.e55 /* 非法：缺少整数或小数 */

```

使用浮点形式表示时，必须包含小数点、指数或同时包含两者。使用指数形式表示时，必须包含整数部分、小数部分或同时包含两者。有符号的指数是用 e 或 E 表示的。

### 字符常量

字符常量是括在单引号里，例如，'x'，且可存储在一个简单的字符类型变量中。一个字符常量可以是一个普通字符（例如 'x'）、一个转义序列（例如 '\t'）或者一个通用字符（例如 '\u02C0'）。

在 C# 中有一些特定的字符，当它们的前面带有反斜杠时有特殊的意义，可用于表示换行符（\n）或制表符 tab（\t）。在这里，列出一些转义序列码：

转义序列含义 \\\ 字符 \'' 字符 \"" 字符 \?? 字符 \aAlert 或 bell \b退格键（Backspace） \f换页符（Form feed） \n换行符（Newline） \r回车 \t水平制表符 tab \v垂直制表符 tab \ooo一到三位的八进制数 \xhh . . .一个或多个数字的十六进制数

以下是一些转义序列字符的实例：

```

namespace EscapeChar
{
class Program
{
static void Main(string[] args)
{
Console.WriteLine("Hello\tWorld\n\n");
Console.ReadLine();
}
}
}

```

当上面的代码被编译和执行时，它会产生下列结果：

```

Hello World

```

### 字符串常量

字符串常量是括在双引号 "" 里，或者是括在 @"" 里。字符串常量包含的字符与字符常量相似，可以是：普通字符、转义序列和通用字符

使用字符串常量时，可以把一个很长的行拆成多个行，可以使用空格分隔各个部分。

这里是一些字符串常量的实例。下面所列的各种形式表示相同的字符串。

```

string a = "hello, world"; // hello, world
string b = @"hello, world"; // hello, world
string c = "hello \t world"; // hello world
string d = @"hello \t world"; // hello \t world
string e = "Joe said \"Hello\" to me"; // Joe said "Hello" to me
string f = @"Joe said ""Hello"" to me"; // Joe said "Hello" to me
string g = "\\\\server\\share\\file.txt"; // \\server\share\file.txt
string h = @"\\server\share\file.txt"; // \\server\share\file.txt
string i = "one\r\ntwo\r\nthree";
string j = @"one
two
three";

```

### 定义常量

常量是使用 const 关键字来定义的 。定义一个常量的语法如下：

```

const <data_type> <constant_name> = value;

```

下面的代码演示了如何在程序中定义和使用常量：

### 实例

using System;

public class ConstTest
{
class SampleClass
{
public int x;
public int y;
public const int c1 = 5;
public const int c2 = c1 + 5;

public SampleClass(int p1, int p2)
{
x = p1;
y = p2;
}
}

static void Main()
{
SampleClass mC = new SampleClass(11, 22);
Console.WriteLine("x = {0}, y = {1}", mC.x, mC.y);
Console.WriteLine("c1 = {0}, c2 = {1}",
SampleClass.c1, SampleClass.c2);
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

x = 11, y = 22
c1 = 5, c2 = 10

```

---

## C# 运算符

Source: https://www.runoob.com/csharp/csharp-operators.html

## C# 运算符

运算符是一种告诉编译器执行特定的数学或逻辑操作的符号。C# 有丰富的内置运算符，分类如下：

- 算术运算符
- 关系运算符
- 逻辑运算符
- 位运算符
- 赋值运算符
- 其他运算符

本教程将逐一讲解算术运算符、关系运算符、逻辑运算符、位运算符、赋值运算符及其他运算符。

### 算术运算符

下表显示了 C# 支持的所有算术运算符。假设变量 A 的值为 10，变量 B 的值为 20，则：

运算符描述实例 +把两个操作数相加 A + B 将得到 30 -从第一个操作数中减去第二个操作数 A - B 将得到 -10 *把两个操作数相乘 A * B 将得到 200 /分子除以分母 B / A 将得到 2 %取模运算符，整除后的余数 B % A 将得到 0 ++自增运算符，整数值增加 1 A++ 将得到 11 --自减运算符，整数值减少 1 A-- 将得到 9

#### 实例

请看下面的实例，了解 C# 中所有可用的算术运算符：

### 实例

using System;

namespace OperatorsAppl
{
class Program
{
static void Main(string[] args)
{
int a = 21;
int b = 10;
int c;

c = a + b;
Console.WriteLine("Line 1 - c 的值是 {0}", c);
c = a - b;
Console.WriteLine("Line 2 - c 的值是 {0}", c);
c = a * b;
Console.WriteLine("Line 3 - c 的值是 {0}", c);
c = a / b;
Console.WriteLine("Line 4 - c 的值是 {0}", c);
c = a % b;
Console.WriteLine("Line 5 - c 的值是 {0}", c);

// ++a 先进行自增运算再赋值
c = ++a;
Console.WriteLine("Line 6 - c 的值是 {0}", c);

// 此时 a 的值为 22
// --a 先进行自减运算再赋值
c = --a;
Console.WriteLine("Line 7 - c 的值是 {0}", c);
Console.ReadLine();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Line 1 - c 的值是 31
Line 2 - c 的值是 11
Line 3 - c 的值是 210
Line 4 - c 的值是 2
Line 5 - c 的值是 1
Line 6 - c 的值是 22
Line 7 - c 的值是 21

```

- c = a++: 先将 a 赋值给 c，再对 a 进行自增运算。
- c = ++a: 先将 a 进行自增运算，再将 a 赋值给 c 。
- c = a--: 先将 a 赋值给 c，再对 a 进行自减运算。
- c = --a: 先将 a 进行自减运算，再将 a 赋值给 c 。

### 实例

using System;

namespace OperatorsAppl
{
class Program
{
static void Main(string[] args)
{
int a = 1;
int b;

// a++ 先赋值再进行自增运算
b = a++;
Console.WriteLine("a = {0}", a);
Console.WriteLine("b = {0}", b);
Console.ReadLine();

// ++a 先进行自增运算再赋值
a = 1; // 重新初始化 a
b = ++a;
Console.WriteLine("a = {0}", a);
Console.WriteLine("b = {0}", b);
Console.ReadLine();

// a-- 先赋值再进行自减运算
a = 1; // 重新初始化 a
b= a--;
Console.WriteLine("a = {0}", a);
Console.WriteLine("b = {0}", b);
Console.ReadLine();

// --a 先进行自减运算再赋值
a = 1; // 重新初始化 a
b= --a;
Console.WriteLine("a = {0}", a);
Console.WriteLine("b = {0}", b);
Console.ReadLine();
}
}
}

运行实例 »

执行以上程序，输出结果为：

```

a = 2
b = 1
a = 2
b = 2
a = 0
b = 1
a = 0
b = 0

```

### 关系运算符

下表显示了 C# 支持的所有关系运算符。假设变量 A 的值为 10，变量 B 的值为 20，则：

运算符描述实例 ==检查两个操作数的值是否相等，如果相等则条件为真。 (A == B) 不为真。 !=检查两个操作数的值是否相等，如果不相等则条件为真。 (A != B) 为真。 >检查左操作数的值是否大于右操作数的值，如果是则条件为真。 (A > B) 不为真。 <检查左操作数的值是否小于右操作数的值，如果是则条件为真。 (A < B) 为真。 >=检查左操作数的值是否大于或等于右操作数的值，如果是则条件为真。 (A >= B) 不为真。 <=检查左操作数的值是否小于或等于右操作数的值，如果是则条件为真。 (A <= B) 为真。

#### 实例

请看下面的实例，了解 C# 中所有可用的关系运算符：

### 实例

using System;

class Program
{
static void Main(string[] args)
{
int a = 21;
int b = 10;

if (a == b)
{
Console.WriteLine("Line 1 - a 等于 b");
}
else
{
Console.WriteLine("Line 1 - a 不等于 b");
}
if (a < b)
{
Console.WriteLine("Line 2 - a 小于 b");
}
else
{
Console.WriteLine("Line 2 - a 不小于 b");
}
if (a > b)
{
Console.WriteLine("Line 3 - a 大于 b");
}
else
{
Console.WriteLine("Line 3 - a 不大于 b");
}
/* 改变 a 和 b 的值 */
a = 5;
b = 20;
if (a <= b)
{
Console.WriteLine("Line 4 - a 小于或等于 b");
}
if (b >= a)
{
Console.WriteLine("Line 5 - b 大于或等于 a");
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Line 1 - a 不等于 b
Line 2 - a 不小于 b
Line 3 - a 大于 b
Line 4 - a 小于或等于 b
Line 5 - b 大于或等于 a

```

### 逻辑运算符

下表显示了 C# 支持的所有逻辑运算符。假设变量 A 为布尔值 true，变量 B 为布尔值 false，则：

运算符描述实例 &&称为逻辑与运算符。如果两个操作数都非零，则条件为真。 (A && B) 为假。 ||称为逻辑或运算符。如果两个操作数中有任意一个非零，则条件为真。 (A || B) 为真。 !称为逻辑非运算符。用来逆转操作数的逻辑状态。如果条件为真则逻辑非运算符将使其为假。 !(A && B) 为真。

#### 实例

请看下面的实例，了解 C# 中所有可用的逻辑运算符：

### 实例

using System;

namespace OperatorsAppl
{
class Program
{
static void Main(string[] args)
{
bool a = true;
bool b = true;

if (a && b)
{
Console.WriteLine("Line 1 - 条件为真");
}
if (a || b)
{
Console.WriteLine("Line 2 - 条件为真");
}
/* 改变 a 和 b 的值 */
a = false;
b = true;
if (a && b)
{
Console.WriteLine("Line 3 - 条件为真");
}
else
{
Console.WriteLine("Line 3 - 条件不为真");
}
if (!(a && b))
{
Console.WriteLine("Line 4 - 条件为真");
}
Console.ReadLine();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Line 1 - 条件为真
Line 2 - 条件为真
Line 3 - 条件不为真
Line 4 - 条件为真

```

### 位运算符

位运算符作用于位，并逐位执行操作。&、 | 和 ^ 的真值表如下所示：

pqp & qp | qp ^ q 00000 01011 11110 10011

假设如果 A = 60，且 B = 13，现在以二进制格式表示，它们如下所示：

A = 0011 1100

B = 0000 1101

-----------------

A&B = 0000 1100

A|B = 0011 1101

A^B = 0011 0001

~A = 1100 0011

下表列出了 C# 支持的位运算符。假设变量 A 的值为 60，变量 B 的值为 13，则：

运算符描述实例 &如果同时存在于两个操作数中，二进制 AND 运算符复制一位到结果中。 (A & B) 将得到 12，即为 0000 1100 |如果存在于任一操作数中，二进制 OR 运算符复制一位到结果中。 (A | B) 将得到 61，即为 0011 1101 ^如果存在于其中一个操作数中但不同时存在于两个操作数中，二进制异或运算符复制一位到结果中。 (A ^ B) 将得到 49，即为 0011 0001 ~按位取反运算符是一元运算符，具有"翻转"位效果，即0变成1，1变成0，包括符号位。(~A ) 将得到 -61，即为 1100 0011，一个有符号二进制数的补码形式。 <<二进制左移运算符。左操作数的值向左移动右操作数指定的位数。 A << 2 将得到 240，即为 1111 0000 >>二进制右移运算符。左操作数的值向右移动右操作数指定的位数。 A >> 2 将得到 15，即为 0000 1111

#### 实例

请看下面的实例，了解 C# 中所有可用的位运算符：

### 实例

using System;
namespace OperatorsAppl
{
class Program
{
static void Main(string[] args)
{
int a = 60; /* 60 = 0011 1100 */
int b = 13; /* 13 = 0000 1101 */
int c = 0;

c = a & b; /* 12 = 0000 1100 */
Console.WriteLine("Line 1 - c 的值是 {0}", c );

c = a | b; /* 61 = 0011 1101 */
Console.WriteLine("Line 2 - c 的值是 {0}", c);

c = a ^ b; /* 49 = 0011 0001 */
Console.WriteLine("Line 3 - c 的值是 {0}", c);

c = ~a; /*-61 = 1100 0011 */
Console.WriteLine("Line 4 - c 的值是 {0}", c);

c = a << 2; /* 240 = 1111 0000 */
Console.WriteLine("Line 5 - c 的值是 {0}", c);

c = a >> 2; /* 15 = 0000 1111 */
Console.WriteLine("Line 6 - c 的值是 {0}", c);
Console.ReadLine();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Line 1 - c 的值是 12
Line 2 - c 的值是 61
Line 3 - c 的值是 49
Line 4 - c 的值是 -61
Line 5 - c 的值是 240
Line 6 - c 的值是 15

```

### 赋值运算符

下表列出了 C# 支持的赋值运算符：

运算符描述实例 =简单的赋值运算符，把右边操作数的值赋给左边操作数 C = A + B 将把 A + B 的值赋给 C +=加且赋值运算符，把右边操作数加上左边操作数的结果赋值给左边操作数 C += A 相当于 C = C + A -=减且赋值运算符，把左边操作数减去右边操作数的结果赋值给左边操作数 C -= A 相当于 C = C - A *=乘且赋值运算符，把右边操作数乘以左边操作数的结果赋值给左边操作数 C *= A 相当于 C = C * A /=除且赋值运算符，把左边操作数除以右边操作数的结果赋值给左边操作数 C /= A 相当于 C = C / A %=求模且赋值运算符，求两个操作数的模赋值给左边操作数 C %= A 相当于 C = C % A <<=左移且赋值运算符 C <<= 2 等同于 C = C << 2 >>=右移且赋值运算符 C >>= 2 等同于 C = C >> 2 &=按位与且赋值运算符 C &= 2 等同于 C = C & 2 ^=按位异或且赋值运算符 C ^= 2 等同于 C = C ^ 2 |=按位或且赋值运算符 C |= 2 等同于 C = C | 2

#### 实例

请看下面的实例，了解 C# 中所有可用的赋值运算符：

### 实例

using System;

namespace OperatorsAppl
{
class Program
{
static void Main(string[] args)
{
int a = 21;
int c;

c = a;
Console.WriteLine("Line 1 - = c 的值 = {0}", c);

c += a;
Console.WriteLine("Line 2 - += c 的值 = {0}", c);

c -= a;
Console.WriteLine("Line 3 - -= c 的值 = {0}", c);

c *= a;
Console.WriteLine("Line 4 - *= c 的值 = {0}", c);

c /= a;
Console.WriteLine("Line 5 - /= c 的值 = {0}", c);

c = 200;
c %= a;
Console.WriteLine("Line 6 - %= c 的值 = {0}", c);

c <<= 2;
Console.WriteLine("Line 7 - <<= c 的值 = {0}", c);

c >>= 2;
Console.WriteLine("Line 8 - >>= c 的值 = {0}", c);

c &= 2;
Console.WriteLine("Line 9 - &= c 的值 = {0}", c);

c ^= 2;
Console.WriteLine("Line 10 - ^= c 的值 = {0}", c);

c |= 2;
Console.WriteLine("Line 11 - |= c 的值 = {0}", c);
Console.ReadLine();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Line 1 - = c 的值 = 21
Line 2 - += c 的值 = 42
Line 3 - -= c 的值 = 21
Line 4 - *= c 的值 = 441
Line 5 - /= c 的值 = 21
Line 6 - %= c 的值 = 11
Line 7 - <<= c 的值 = 44
Line 8 - >>= c 的值 = 11
Line 9 - &= c 的值 = 2
Line 10 - ^= c 的值 = 0
Line 11 - |= c 的值 = 2

```

### 其他运算符

下表列出了 C# 支持的其他一些重要的运算符，包括 sizeof、typeof 和 ? :。

运算符描述实例 sizeof()返回数据类型的大小。sizeof(int)，将返回 4. typeof()返回 class 的类型。typeof(StreamReader); &返回变量的地址。&a; 将得到变量的实际地址。 *变量的指针。*a; 将指向一个变量。 ? :条件表达式 如果条件为真 ? 则为 X : 否则为 Y is判断对象是否为某一类型。If( Ford is Car) // 检查 Ford 是否是 Car 类的一个对象。 as强制转换，即使转换失败也不会抛出异常。Object obj = new StringReader("Hello");
StringReader r = obj as StringReader;

#### 实例

### 实例

using System;

namespace OperatorsAppl
{

class Program
{
static void Main(string[] args)
{

/* sizeof 运算符的实例 */
Console.WriteLine("int 的大小是 {0}", sizeof(int));
Console.WriteLine("short 的大小是 {0}", sizeof(short));
Console.WriteLine("double 的大小是 {0}", sizeof(double));

/* 三元运算符的实例 */
int a, b;
a = 10;
b = (a == 1) ? 20 : 30;
Console.WriteLine("b 的值是 {0}", b);

b = (a == 10) ? 20 : 30;
Console.WriteLine("b 的值是 {0}", b);
Console.ReadLine();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

int 的大小是 4
short 的大小是 2
double 的大小是 8
b 的值是 30
b 的值是 20

```

typeof 关键字用于获取一个类型的类型对象，它通常用于反射和动态创建类型实例。

下面是一个使用 typeof 的简单示例：

### 实例

using System;

class Program
{
static void Main(string[] args)
{
Type type = typeof(string);
Console.WriteLine(type.FullName);
Console.ReadKey();
}
}

在上面的代码中，我们使用 typeof 关键字来获取 string 类型的类型对象，并将其存储在 Type 类型的变量 type 中，然后，我们使用 FullName 属性打印该类型的完全限定名。

当上面的代码被编译和执行时，它会产生下列结果：

```
System.String
```

### C# 中的运算符优先级

运算符的优先级确定表达式中项的组合。这会影响到一个表达式如何计算。某些运算符比其他运算符有更高的优先级，例如，乘除运算符具有比加减运算符更高的优先级。

例如 x = 7 + 3 * 2，在这里，x 被赋值为 13，而不是 20，因为运算符 * 具有比 + 更高的优先级，所以首先计算乘法 3*2，然后再加上 7。

下表将按运算符优先级从高到低列出各个运算符，具有较高优先级的运算符出现在表格的上面，具有较低优先级的运算符出现在表格的下面。在表达式中，较高优先级的运算符会优先被计算。

优先级简易概括：有括号先括号，后乘除在加减，然后位移再关系，逻辑完后条件，最后一个逗号 , 。

类别 运算符 结合性 后缀 () [] -> . ++ - - 从左到右 一元 + - ! ~ ++ - - (type)* & sizeof 从右到左 乘除 * / % 从左到右 加减 + - 从左到右 移位 << >> 从左到右 关系 < <= > >= 从左到右 相等 == != 从左到右 位与 AND & 从左到右 位异或 XOR ^ 从左到右 位或 OR | 从左到右 逻辑与 AND && 从左到右 逻辑或 OR || 从左到右 条件 ?: 从右到左 赋值 = += -= *= /= %=>>= <<= &= ^= |= 从右到左 逗号 , 从左到右

#### 实例

### 实例

using System;

namespace OperatorsAppl
{

class Program
{
static void Main(string[] args)
{
int a = 20;
int b = 10;
int c = 15;
int d = 5;
int e;
e = (a + b) * c / d; // ( 30 * 15 ) / 5
Console.WriteLine("(a + b) * c / d 的值是 {0}", e);

e = ((a + b) * c) / d; // (30 * 15 ) / 5
Console.WriteLine("((a + b) * c) / d 的值是 {0}", e);

e = (a + b) * (c / d); // (30) * (15/5)
Console.WriteLine("(a + b) * (c / d) 的值是 {0}", e);

e = a + (b * c) / d; // 20 + (150/5)
Console.WriteLine("a + (b * c) / d 的值是 {0}", e);
Console.ReadLine();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

(a + b) * c / d 的值是 90
((a + b) * c) / d 的值是 90
(a + b) * (c / d) 的值是 90
a + (b * c) / d 的值是 50

```

算术运算符和赋值运算符的优先级：

### 实例

int a = 10, b = 5, c = 2;
int result = a + b * c; // 先计算乘法，再计算加法
Console.WriteLine(result); // 输出20

result = (a + b) * c; // 先计算加法，再计算乘法
Console.WriteLine(result); // 输出30

a += b * c; // 先计算乘法，再执行加法赋值
Console.WriteLine(a); // 输出20

逻辑运算符的优先级：

### 实例

bool a = true, b = false, c = true;
bool result = a || b && c; // 先计算与运算，再计算或运算
Console.WriteLine(result); // 输出true

result = (a || b) && c; // 先计算或运算，再计算与运算
Console.WriteLine(result); // 输出true

result = a || (b && c); // 先计算与运算，再计算或运算
Console.WriteLine(result); // 输出true

条件运算符的优先级：

### 实例

int a = 10, b = 5;
string result = a > b ? "a大于b" : "a不大于b"; // 先判断大小关系，再执行条件语句
Console.WriteLine(result); // 输出"a大于b"

注意：由于括号可以改变运算符优先级，所以在实际应用中建议尽可能使用括号来明确运算顺序，提高代码的可读性和准确性。

---

## C# 判断

Source: https://www.runoob.com/csharp/csharp-decision.html

## C# 判断

判断结构要求程序员指定一个或多个要评估或测试的条件，以及条件为真时要执行的语句（必需的）和条件为假时要执行的语句（可选的）。

下面是大多数编程语言中典型的判断结构的一般形式：

### 判断语句

C# 提供了以下类型的判断语句。点击链接查看每个语句的细节。

语句描述 if 语句一个 if 语句 由一个布尔表达式后跟一个或多个语句组成。 if...else 语句一个 if 语句 后可跟一个可选的 else 语句，else 语句在布尔表达式为假时执行。 嵌套 if 语句您可以在一个 if 或 else if 语句内使用另一个 if 或 else if 语句。 switch 语句一个 switch 语句允许测试一个变量等于多个值时的情况。 嵌套 switch 语句您可以在一个 switch 语句内使用另一个 switch 语句。 C# Null 条件运算符C# 6.0 引入了 Null 条件运算符 ?.，用于安全地访问属性或方法，而不必写一堆 if 判断。

### ? : 运算符

我们已经在前面的章节中讲解了 条件运算符 ? :，可以用来替代 if...else 语句。它的一般形式如下：

```

Exp1 ? Exp2 : Exp3;

```

其中，Exp1、Exp2 和 Exp3 是表达式。请注意，冒号的使用和位置。

? 表达式的值是由 Exp1 决定的。如果 Exp1 为真，则计算 Exp2 的值，结果即为整个 ? 表达式的值。如果 Exp1 为假，则计算 Exp3 的值，结果即为整个 ? 表达式的值。

---

## C# 循环

Source: https://www.runoob.com/csharp/csharp-loops.html

## C# 循环

有的时候，可能需要多次执行同一块代码。一般情况下，语句是顺序执行的：函数中的第一个语句先执行，接着是第二个语句，依此类推。

编程语言提供了允许更为复杂的执行路径的多种控制结构。

循环语句允许我们多次执行一个语句或语句组，下面是大多数编程语言中循环语句的一般形式：

### 循环类型

C# 提供了以下几种循环类型。点击链接查看每个类型的细节。

循环类型描述 while 循环当给定条件为真时，重复语句或语句组。它会在执行循环主体之前测试条件。 for/foreach 循环多次执行一个语句序列，简化管理循环变量的代码。 do...while 循环除了它是在循环主体结尾测试条件外，其他与 while 语句类似。 嵌套循环您可以在 while、for 或 do..while 循环内使用一个或多个循环。

### 循环控制语句

循环控制语句更改执行的正常序列。当执行离开一个范围时，所有在该范围中创建的自动对象都会被销毁。

C# 提供了下列的控制语句。点击链接查看每个语句的细节。

控制语句描述 break 语句终止 loop 或 switch 语句，程序流将继续执行紧接着 loop 或 switch 的下一条语句。 continue 语句跳过本轮循环，开始下一轮循环。

### 无限循环

如果条件永远不为假，则循环将变成无限循环。for 循环在传统意义上可用于实现无限循环。由于构成循环的三个表达式中任何一个都不是必需的，您可以将某些条件表达式留空来构成一个无限循环。

### 实例

using System;

namespace Loops
{

class Program
{
static void Main(string[] args)
{
for (; ; )
{
Console.WriteLine("Hey! I am Trapped");
}

}
}
}

当条件表达式不存在时，它被假设为真。您也可以设置一个初始值和增量表达式，但是一般情况下，程序员偏向于使用 for(;;) 结构来表示一个无限循环。

---

## C# 封装

Source: https://www.runoob.com/csharp/csharp-encapsulation.html

## C# 封装

封装 被定义为"把一个或多个项目封闭在一个物理的或者逻辑的包中"。在面向对象程序设计方法论中，封装是为了防止对实现细节的访问。

抽象和封装是面向对象程序设计的相关特性。抽象允许相关信息可视化，封装则使开发者实现所需级别的抽象。

C# 封装根据具体的需要，设置使用者的访问权限，并通过 访问修饰符 来实现。

一个 访问修饰符 定义了一个类成员的范围和可见性。C# 支持的访问修饰符如下所示：

- public：所有对象都可以访问；
- private：对象本身在对象内部可以访问；
- protected：只有该类对象及其子类对象可以访问
- internal：同一个程序集的对象可以访问；
- protected internal：访问限于当前程序集或派生自包含类的类型。

### Public 访问修饰符

Public 访问修饰符允许一个类将其成员变量和成员函数暴露给其他的函数和对象。任何公有成员可以被外部的类访问。

下面的实例说明了这点：

### 实例

using System;

namespace RectangleApplication
{
class Rectangle
{
//成员变量
public double length;
public double width;

public double GetArea()
{
return length * width;
}
public void Display()
{
Console.WriteLine("长度： {0}", length);
Console.WriteLine("宽度： {0}", width);
Console.WriteLine("面积： {0}", GetArea());
}
}// Rectangle 结束

class ExecuteRectangle
{
static void Main(string[] args)
{
Rectangle r = new Rectangle();
r.length = 4.5;
r.width = 3.5;
r.Display();
Console.ReadLine();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

长度： 4.5
宽度： 3.5
面积： 15.75

```

在上面的实例中，成员变量 length 和 width 被声明为 public，所以它们可以被函数 Main() 使用 Rectangle 类的实例 r 访问。

成员函数 Display() 和 GetArea() 可以直接访问这些变量。

成员函数 Display() 也被声明为 public，所以它也能被 Main() 使用 Rectangle 类的实例 r 访问。

### Private 访问修饰符

Private 访问修饰符允许一个类将其成员变量和成员函数对其他的函数和对象进行隐藏。只有同一个类中的函数可以访问它的私有成员。即使是类的实例也不能访问它的私有成员。

下面的实例说明了这点：

### 实例

using System;

namespace RectangleApplication
{
class Rectangle
{
// 私有成员变量
private double length;
private double width;

// 公有方法，用于从用户输入获取矩形的长度和宽度
public void AcceptDetails()
{
Console.WriteLine("请输入长度：");
length = Convert.ToDouble(Console.ReadLine());
Console.WriteLine("请输入宽度：");
width = Convert.ToDouble(Console.ReadLine());
}

// 公有方法，用于计算矩形的面积
public double GetArea()
{
return length * width;
}

// 公有方法，用于显示矩形的属性和面积
public void Display()
{
Console.WriteLine("长度： {0}", length);
Console.WriteLine("宽度： {0}", width);
Console.WriteLine("面积： {0}", GetArea());
}
}//end class Rectangle

class ExecuteRectangle
{
static void Main(string[] args)
{
// 创建 Rectangle 类的实例
Rectangle r = new Rectangle();

// 通过公有方法 AcceptDetails() 从用户输入获取矩形的长度和宽度
r.AcceptDetails();

// 通过公有方法 Display() 显示矩形的属性和面积
r.Display();

Console.ReadLine();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

请输入长度：
5
请输入宽度：
3
长度： 5
宽度： 3
面积： 15

```

说明：

- `length` 和 `width` 被声明为私有成员变量，以防止直接在类的外部访问和修改。
- `AcceptDetails` 方法允许用户输入矩形的长度和宽度，这是通过公有方法来操作私有变量的一个例子。
- `GetArea` 方法用于计算矩形的面积，而 `Display` 方法用于显示矩形的属性和面积。
- 在 `ExecuteRectangle` 类中，通过创建 `Rectangle` 类的实例，然后调用其公有方法，来执行操作。这样，主程序无法直接访问和修改矩形的长度和宽度，而是通过类提供的公有接口来进行操作，实现了封装。

### Protected 访问修饰符

Protected 访问修饰符允许子类访问它的基类的成员变量和成员函数。这样有助于实现继承。我们将在继承的章节详细讨论这个。更详细地讨论这个。

### Internal 访问修饰符

Internal 访问修饰符允许一个类将其成员变量和成员函数暴露给当前程序中的其他函数和对象。换句话说，带有 internal 访问修饰符的任何成员可以被定义在该成员所定义的应用程序内的任何类或方法访问。

下面的实例说明了这点：

### 实例

using System;

namespace RectangleApplication
{
class Rectangle
{
//成员变量
internal double length;
internal double width;

double GetArea()
{
return length * width;
}
public void Display()
{
Console.WriteLine("长度： {0}", length);
Console.WriteLine("宽度： {0}", width);
Console.WriteLine("面积： {0}", GetArea());
}
}//end class Rectangle
class ExecuteRectangle
{
static void Main(string[] args)
{
Rectangle r = new Rectangle();
r.length = 4.5;
r.width = 3.5;
r.Display();
Console.ReadLine();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

长度： 4.5
宽度： 3.5
面积： 15.75

```

在上面的实例中，请注意成员函数 GetArea() 声明的时候不带有任何访问修饰符。如果没有指定访问修饰符，则使用类成员的默认访问修饰符，即为 private。

### Protected Internal 访问修饰符

Protected Internal 访问修饰符允许在本类,派生类或者包含该类的程序集中访问。这也被用于实现继承。

---

## C# 方法（Methods）

Source: https://www.runoob.com/csharp/csharp-methods.html

## C# 方法（Methods）

方法（Method）是将一组相关语句组织在一起、用来执行特定任务的代码块。它是 C# 中实现代码复用和逻辑封装的核心机制。每一个 C# 程序至少有一个 `Main` 方法作为程序入口。

使用方法需要两步：定义方法和调用方法。

### 1. 定义方法

C# 中定义方法的基本语法如下：

```
<访问修饰符> <返回类型> <方法名>(<参数列表>)
{
方法体
}
```

方法由以下几个元素组成：

- 访问修饰符（Access Modifier）：控制方法的可见性，如 `public`、`private`、`protected`、`internal`。
- 返回类型（Return Type）：方法返回值的数据类型。如果方法不返回任何值，使用 `void`。
- 方法名（Method Name）：方法的唯一标识符，采用 PascalCase 命名规范（如 `FindMax`、`GetUserName`）。
- 参数列表（Parameter List）：方法接收的输入，用圆括号括起来。参数是可选的，一个方法可以没有参数。
- 方法体（Method Body）：包含实现功能的代码，用花括号 `{}` 包围。

下图展示了一个方法的完整结构： C# 方法结构解析 public int FindMax ( int num1, int num2 ) 访问修饰符 public / private 控制可见性 返回类型 int / void 返回值的数据类型 方法名 PascalCase 唯一标识符 参数列表 类型 + 名称 多个用逗号分隔 { 方法体 — 实现功能的代码 } 方法体（Method Body） 完整方法签名： public int FindMax(int a, int b) 修饰符 → 返回类型 → 方法名 → 参数列表

#### 1.1 第一个方法示例

下面定义了一个 `FindMax` 方法，接收两个整数并返回其中较大的一个：

### 实例

using System;

namespace MethodDemo
{
class NumberHelper
{
// 定义方法：接收两个 int，返回较大的那个
public int FindMax(int num1, int num2)
{
if (num1 > num2)
return num1;
else
return num2;
}

static void Main(string[] args)
{
// 创建对象并调用方法
NumberHelper helper = new NumberHelper();
int result = helper.FindMax(100, 200);

Console.WriteLine($"最大值是：{result}"); // 最大值是：200
}
}
}

方法的调用流程： 方法调用流程 调用方 FindMax(100, 200) 实参 方法接收 num1=100, num2=200 执行方法体 比较并返回结果 返回 返回值 result = 200 1. 调用方法，传递实参 → 2. 形参接收值 → 3. 执行方法体 → 4. 返回结果给调用方

跨类调用：只要方法声明为 `public`，就可以从其他类中通过创建实例来调用。这是面向对象编程中代码复用的基础。

### 2. 参数传递方式

C# 提供三种向方法传递参数的方式，它们的核心区别在于是否共享内存： 三种参数传递方式对比 按值传递（默认） 调用方 a = 100 复制 方法内 x = 100 各自独立的内存空间 修改 x 不影响 a swap(a, b) 后 a、b 不变 无需额外关键字 void Swap(int x, int y) 引用传递（ref） 调用方 a = 200 同一地址 方法内 ref x = 200 共享同一块内存 修改 x 会改变 a swap(ref a, ref b) 后交换成功 需要 ref 关键字 void Swap(ref int x, ref int y) 输出参数（out） 调用方 a = 5 ← 输出 方法内 out x = 5 方法必须为 out 参数赋值 调用前无需初始化 可返回多个值 需要 out 关键字 void Calc(out int x, out int y)

方式 关键字 行为 典型场景 按值传递 （无） 复制参数值，方法内修改不影响原变量 大多数方法调用（默认） 按引用传递 `ref` 传递变量的引用，方法内修改会改变原变量 需要修改外部变量（如交换值） 输出参数 `out` 类似 ref，但方法必须赋值，调用前无需初始化 方法需要返回多个值

#### 2.1 按值传递（默认）

按值传递是默认方式。实参的值会被复制给形参，方法内对形参的修改不会影响原始变量：

### 实例

using System;

namespace ParameterDemo
{
class Program
{
// 按值传递：x 和 y 是 a、b 的副本
public void Swap(int x, int y)
{
int temp = x;
x = y;
y = temp;
Console.WriteLine($" 方法内：x={x}, y={y}"); // 交换成功
}

static void Main(string[] args)
{
var p = new Program();
int a = 100, b = 200;

Console.WriteLine($"调用前：a={a}, b={b}");
p.Swap(a, b);
Console.WriteLine($"调用后：a={a}, b={b}"); // a、b 没有变化！
}
}
}

```
调用前：a=100, b=200
方法内：x=200, y=100
调用后：a=100, b=200
```

#### 2.2 按引用传递（ref）

使用 `ref` 关键字传递变量的引用（内存地址），方法内对形参的修改会直接影响原始变量。注意调用时也必须加上 `ref`：

### 实例

using System;

namespace ParameterDemo
{
class Program
{
// 引用传递：x 和 y 直接引用 a、b 的内存
public void Swap(ref int x, ref int y)
{
int temp = x;
x = y;
y = temp;
}

static void Main(string[] args)
{
var p = new Program();
int a = 100, b = 200;

Console.WriteLine($"调用前：a={a}, b={b}");
p.Swap(ref a, ref b); // 调用时也必须加 ref
Console.WriteLine($"调用后：a={a}, b={b}"); // 交换成功！
}
}
}

```
调用前：a=100, b=200
调用后：a=200, b=100
```

ref vs out 的区别：`ref` 要求变量在传递前必须已初始化；`out` 则不要求，但方法内部必须为 out 参数赋值后才能使用。

#### 2.3 输出参数（out）

`out` 参数让方法可以返回多个值。与 `ref` 不同的是，调用前变量无需初始化，但方法内部必须为其赋值：

### 实例

using System;

namespace ParameterDemo
{
class Program
{
// out 参数：方法负责赋值
public void GetValues(out int x, out int y)
{
Console.Write("请输入第一个值：");
x = Convert.ToInt32(Console.ReadLine());

Console.Write("请输入第二个值：");
y = Convert.ToInt32(Console.ReadLine());
// 注意：方法结束前必须为所有 out 参数赋值，否则编译错误
}

// 实用示例：TryParse 模式
public bool TryDivide(int a, int b, out double result)
{
if (b == 0)
{
result = 0;
return false; // 除数为零
}
result = (double)a / b;
return true;
}

static void Main(string[] args)
{
var p = new Program();

// out 参数无需初始化
int a, b;
p.GetValues(out a, out b);
Console.WriteLine($"a={a}, b={b}");

// TryParse 模式
if (p.TryDivide(10, 3, out double res))
Console.WriteLine($"10 / 3 = {res:F2}"); // 3.33
}
}
}

从 C# 7.0 起，可以在调用时内联声明 out 变量，使代码更简洁：

```
// C# 7.0+ 内联声明
if (int.TryParse("123", out int number))
Console.WriteLine($"解析成功：{number}");

```

### 3. 递归方法

递归（Recursion）是指方法调用自身。每个递归方法都需要两个要素：

- 基准条件（Base Case）：递归终止的条件，防止无限递归
- 递归步骤（Recursive Step）：将问题分解为更小的子问题

经典示例——计算阶乘 $$n! = n \times (n-1)!$$：

### 实例

using System;

namespace RecursionDemo
{
class Program
{
// 递归计算阶乘
public int Factorial(int n)
{
// 基准条件：0! = 1，1! = 1
if (n <= 1)
return 1;

// 递归步骤：n! = n × (n-1)!
return n * Factorial(n - 1);
}

static void Main(string[] args)
{
var p = new Program();
Console.WriteLine($"6! = {p.Factorial(6)}"); // 720
Console.WriteLine($"7! = {p.Factorial(7)}"); // 5040
Console.WriteLine($"8! = {p.Factorial(8)}"); // 40320
}
}
}

```
6! = 720
7! = 5040
8! = 40320
```

递归调用的展开过程：

```
Factorial(4)
= 4 × Factorial(3)
= 4 × 3 × Factorial(2)
= 4 × 3 × 2 × Factorial(1)
= 4 × 3 × 2 × 1
= 24
```

注意：递归深度过大时会导致 `StackOverflowException`。对于性能敏感的场景，考虑改用循环实现。

### 4. 方法的更多特性

#### 4.1 默认参数与命名参数

C# 允许为参数设置默认值，调用时可以省略这些参数。配合命名参数，还可以跳过某些参数只指定后面的：

### 实例

using System;

namespace MethodFeatures
{
class Program
{
// 默认参数：power 默认为 2
static double Power(double baseNum, int power = 2)
{
double result = 1;
for (int i = 0; i < power; i++)
result *= baseNum;
return result;
}

// 多个默认参数
static void PrintInfo(string name, int age = 18, string city = "未知")
{
Console.WriteLine($"{name}, {age}岁, {city}");
}

static void Main(string[] args)
{
// 使用默认参数
Console.WriteLine(Power(3)); // 9.0（使用默认 power=2）
Console.WriteLine(Power(3, 3)); // 27.0

// 命名参数：跳过 age，只指定 city
PrintInfo("小明", city: "北京");
// 输出：小明, 18岁, 北京

// 命名参数可以不按顺序
PrintInfo(city: "上海", name: "小红", age: 25);
// 输出：小红, 25岁, 上海
}
}
}

#### 4.2 表达式体方法

当方法体只有一条语句时，可以用 `=>` 简化写法（C# 6.0+）：

```
// 传统写法
public int Add(int a, int b)
{
return a + b;
}

// 表达式体写法（等价）
public int Add(int a, int b) => a + b;

```

#### 4.3 方法重载（Overloading）

同一个类中可以有多个同名方法，只要参数列表不同（参数类型、数量或顺序不同）：

### 实例

using System;

namespace MethodFeatures
{
class Calculator
{
// 重载 1：两个 int 相加
public int Add(int a, int b) => a + b;

// 重载 2：三个 int 相加
public int Add(int a, int b, int c) => a + b + c;

// 重载 3：两个 double 相加
public double Add(double a, double b) => a + b;

static void Main(string[] args)
{
var calc = new Calculator();
Console.WriteLine(calc.Add(1, 2)); // 3（调用重载 1）
Console.WriteLine(calc.Add(1, 2, 3)); // 6（调用重载 2）
Console.WriteLine(calc.Add(1.5, 2.5)); // 4.0（调用重载 3）
}
}
}

### 小结

特性 语法 说明 定义方法 `int Add(int a, int b) { ... }` 指定返回类型、方法名和参数 无返回值 `void DoSomething() { ... }` `void` 表示不返回任何值 按值传递 `void Foo(int x)` 默认方式，修改形参不影响实参 引用传递 `void Foo(ref int x)` 共享内存，修改会反映到实参 输出参数 `void Foo(out int x)` 方法内必须赋值，可返回多个值 默认参数 `void Foo(int x = 10)` 调用时可省略，使用默认值 命名参数 `Foo(city: "北京")` 按名称指定参数，可不按顺序 表达式体 `int Add(int a, int b) => a + b;` 单行方法的简写（C# 6.0+） 方法重载 同名不同参数 编译器根据参数列表自动选择

---

## C# 可空类型（Nullable）

Source: https://www.runoob.com/csharp/csharp-nullable.html

## C# 可空类型（Nullable）

在 C# 中，`int`、`float`、`bool`、`DateTime` 等都是值类型，它们默认必须有一个值，不能为 `null`：

```
int a = null; // 编译错误：无法将 null 转换为 int
```

可空类型（Nullable Types） 解决了这个问题——它让值类型额外拥有一个"没有值"的状态（即 `null`），在处理数据库字段、API 返回值等可能缺失数据的场景中非常实用。

声明方式很简单，在类型后面加一个 `?`：

```
int? a = null; // 合法
int? b = 42; // 合法
```

这里的 `int?` 是 `Nullable<int>` 的语法糖（简写形式），两者完全等价：

```
Nullable&lt;int&gt; a = null; // 完整写法
int? a = null; // 简写，效果相同
```

下图展示了可空类型在内存中的工作方式——它用一个额外的布尔标记来记录"是否有值"： Nullable<T> 的内存结构 int（普通值类型） Value: 42 4 bytes，始终有值 包装 Nullable<int>（可空类型） HasValue Value 1 byte 4 bytes 共 5+ bytes int? x = 42; true ✓ 42 HasValue = true，Value = 42 int? y = null; false ✗ 无 HasValue = false，访问 Value 会抛异常

### 单问号 ? 与双问号 ?? 的区别

C# 中与可空类型相关的两个运算符经常一起使用，但含义完全不同：

运算符 名称 用途 示例 `?` 可空类型修饰符 让值类型可以为 `null` `int? i = 3;` 等价于 `Nullable<int> i = new Nullable<int>(3);` `??` 空合并运算符（Null-Coalescing Operator） 当变量为 `null` 时提供默认值 `int result = i ?? 0;`

一个简单的对比：

```
int i; // 普通值类型，默认值为 0，永远不能是 null
int? ii; // 可空类型，默认值为 null
```

### 可空类型的声明与赋值

#### 声明语法

```
<data_type>? <variable_name> = null;
```

例如：

```
int? age = null;
double? temperature = 36.6;
bool? isActive = new bool?(); // 显式构造，默认为 null
DateTime? birthday = null;
```

`Nullable<T>` 可以表示其基础值类型的正常范围，再加上一个 `null` 值。例如 `Nullable<int>` 可以是 `-2,147,483,648` 到 `2,147,483,647` 之间的任意整数，或者 `null`。

#### 完整示例

### 实例

using System;

namespace NullableDemo
{
class Program
{
static void Main(string[] args)
{
// 声明不同类型的可空变量
int? num1 = null;
int? num2 = 45;
double? num3 = new double?(); // null
double? num4 = 3.14157;
bool? boolVal = new bool?(); // null

// 显示值（null 会显示为空字符串）
Console.WriteLine($"num1 = {num1 ?? 0}"); // 0
Console.WriteLine($"num2 = {num2 ?? 0}"); // 45
Console.WriteLine($"num3 = {num3 ?? 0.0}"); // 0
Console.WriteLine($"num4 = {num4 ?? 0.0}"); // 3.14157
Console.WriteLine($"boolVal = {boolVal}"); // （空）
}
}
}

输出结果：

```
num1 = 0
num2 = 45
num3 = 0
num4 = 3.14157
boolVal =
```

### Null 合并运算符（??）

`??` 运算符用于为可空类型或引用类型提供一个兜底默认值。当左侧不为 `null` 时返回左侧值，否则返回右侧值：

```
<表达式1> ?? <表达式2>
```

- 如果 `<表达式1>` 不为 `null`，返回 `<表达式1>`；
- 否则返回 `<表达式2>`。

从 C# 8.0 起，还可以使用 `??=`（空合并赋值运算符），仅当变量为 `null` 时才赋值：

```
int? x = null;
x ??= 10; // x 为 null，赋值为 10
x ??= 20; // x 已有值 10，不再赋值
```

### 实例

using System;

namespace NullableDemo
{
class Program
{
static void Main(string[] args)
{
double? num1 = null;
double? num2 = 3.14157;

// ?? 提供默认值
double result1 = num1 ?? 5.34; // num1 为 null → 返回 5.34
double result2 = num2 ?? 5.34; // num2 有值 → 返回 3.14157

Console.WriteLine($"num1 ?? 5.34 = {result1}"); // 5.34
Console.WriteLine($"num2 ?? 5.34 = {result2}"); // 3.14157

// 链式使用：提供多个备选值
int? a = null;
int? b = null;
int? c = 42;
int value = a ?? b ?? c ?? 0; // 依次尝试，最终得到 42
Console.WriteLine($"a ?? b ?? c ?? 0 = {value}"); // 42
}
}
}

输出结果：

```
num1 ?? 5.34 = 5.34
num2 ?? 5.34 = 3.14157
a ?? b ?? c ?? 0 = 42
```

### 可空类型的常用属性和方法

成员 说明 示例 `.HasValue` 判断变量是否有值（返回 `bool`） `if (num.HasValue) { ... }` `.Value` 获取实际值（若为 `null` 会抛 `InvalidOperationException`） `int x = num.Value;` `.GetValueOrDefault()` 安全获取值，若为 `null` 返回类型的默认值（如 `0`） `num.GetValueOrDefault()` `.GetValueOrDefault(T)` 安全获取值，若为 `null` 返回指定的默认值 `num.GetValueOrDefault(100)` `??` 空合并运算符（语法糖，等价于 `GetValueOrDefault`） `int result = num ?? 100;`

### 实例

using System;

namespace NullableDemo
{
class Program
{
static void Main(string[] args)
{
int? num = null;

// HasValue + Value（传统写法）
if (num.HasValue)
Console.WriteLine($"值为: {num.Value}");
else
Console.WriteLine("num 没有值");

// 安全获取默认值
Console.WriteLine(num.GetValueOrDefault()); // 0
Console.WriteLine(num.GetValueOrDefault(99)); // 99
Console.WriteLine(num ?? 99); // 99（等价写法）

// &#x26a0; 注意：直接访问 .Value 会抛异常
// int x = num.Value; // InvalidOperationException!
}
}
}

### 实际应用场景

可空类型在实际开发中非常常见，尤其是处理可能缺失的数据时：

#### 数据库字段映射

数据库中的字段允许为 `NULL`，用可空类型可以精确对应：

用户ID 年龄 是否激活 1 28 true 2 null false

```
int? age = GetUserAgeFromDB(userId: 2);

// 安全处理 null
string display = age.HasValue
? $"用户年龄：{age.Value}"
: "年龄未知";

// 或者更简洁的写法
string display2 = $"用户年龄：{age ?? 0}";

```

#### 可选参数与配置

```
// 配置项可能未设置
int? timeout = GetConfigValue("timeout");
int actualTimeout = timeout ?? 30; // 未设置时使用默认值 30 秒

```

#### 链式空值检查（C# 6.0+）

```
// 使用 ?. 运算符安全访问可能为 null 的对象
string? name = user?.Address?.City?.Name;
// 如果 user、Address、City 中任何一个为 null，结果就是 null，不会抛异常

```

### 小结

功能 示例 说明 定义可空类型 `int? x = null;` 等价于 `Nullable<int>` 判断是否有值 `x.HasValue` 返回 `true` 或 `false` 获取值（不安全） `x.Value` 若为 `null` 会抛 `InvalidOperationException` 获取默认值 `x ?? 0` 若为 `null` 返回右侧值 获取指定默认值 `x.GetValueOrDefault(10)` 若为 `null` 返回 10 空合并赋值 `x ??= 10` 仅当 `x` 为 `null` 时赋值（C# 8.0+） 安全导航访问 `user?.Name` 若 `user` 为 `null` 则返回 `null`（C# 6.0+）

### C# 8.0 的"可空引用类型"

从 C# 8.0 开始，引入了 可空引用类型（Nullable Reference Types），它与本文介绍的可空值类型是两套不同的机制：

对比项 可空值类型 可空引用类型 示例 `int?` `string?` 作用对象 值类型（`int`、`bool`、`struct` 等） 引用类型（`string`、`class`、数组等） 实现方式 运行时通过 `Nullable<T>` 结构体实现 编译器静态分析，不改变运行时行为 默认状态 始终存在（从 C# 1.0 起） 需在项目中启用 `<Nullable>enable</Nullable>` null 检查 用 `.HasValue` 判断 编译器发出警告（非错误）

可空引用类型是编译期的安全检查工具——它不会改变程序的运行时行为，但会在代码可能产生 `NullReferenceException` 的地方发出编译器警告，帮助你在开发阶段就发现潜在问题。

---

## C# 数组（Array）

Source: https://www.runoob.com/csharp/csharp-array.html

## C# 数组（Array）

数组是一个存储相同类型元素的固定大小的顺序集合。数组是用来存储数据的集合，通常认为数组是一个同一类型变量的集合。

声明数组变量并不是声明 number0、number1、...、number99 一个个单独的变量，而是声明一个就像 numbers 这样的变量，然后使用 numbers[0]、numbers[1]、...、numbers[99] 来表示一个个单独的变量。数组中某个指定的元素是通过索引来访问的。

所有的数组都是由连续的内存位置组成的。最低的地址对应第一个元素，最高的地址对应最后一个元素。

### 声明数组

在 C# 中声明一个数组，您可以使用下面的语法：

```

datatype[] arrayName;

```

其中，

- datatype 用于指定被存储在数组中的元素的类型。
- [ ] 指定数组的秩（维度）。秩指定数组的大小。
- arrayName 指定数组的名称。

例如：

```

double[] balance;

```

### 初始化数组

声明一个数组不会在内存中初始化数组。当初始化数组变量时，您可以赋值给数组。

数组是一个引用类型，所以您需要使用 new 关键字来创建数组的实例。

例如：

```

double[] balance = new double[10];

```

### 赋值给数组

您可以通过使用索引号赋值给一个单独的数组元素，比如：

```

double[] balance = new double[10];
balance[0] = 4500.0;

```

您可以在声明数组的同时给数组赋值，比如：

```

double[] balance = { 2340.0, 4523.69, 3421.0};

```

您也可以创建并初始化一个数组，比如：

```

int [] marks = new int[5] { 99, 98, 92, 97, 95};

```

在上述情况下，你也可以省略数组的大小，比如：

```

int [] marks = new int[] { 99, 98, 92, 97, 95};

```

您也可以赋值一个数组变量到另一个目标数组变量中。在这种情况下，目标和源会指向相同的内存位置：

```

int [] marks = new int[] { 99, 98, 92, 97, 95};
int[] score = marks;

```

当您创建一个数组时，C# 编译器会根据数组类型隐式初始化每个数组元素为一个默认值。例如，int 数组的所有元素都会被初始化为 0。

### 访问数组元素

元素是通过带索引的数组名称来访问的。这是通过把元素的索引放置在数组名称后的方括号中来实现的。例如：

```

double salary = balance[9];

```

下面是一个实例，使用上面提到的三个概念，即声明、赋值、访问数组：

### 实例

using System;
namespace ArrayApplication
{
class MyArray
{
static void Main(string[] args)
{
int [] n = new int[10]; /* n 是一个带有 10 个整数的数组 */
int i,j;

/* 初始化数组 n 中的元素 */
for ( i = 0; i < 10; i++ )
{
n[ i ] = i + 100;
}

/* 输出每个数组元素的值 */
for (j = 0; j < 10; j++ )
{
Console.WriteLine("Element[{0}] = {1}", j, n[j]);
}
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

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

### 使用 foreach 循环

在前面的实例中，我们使用一个 for 循环来访问每个数组元素，您也可以使用一个 foreach 语句来遍历数组。

以下实例我们使用 foreach 来遍历一个数组：

### 实例

using System;

namespace ArrayApplication
{
class MyArray
{
static void Main(string[] args)
{
int [] n = new int[10]; /* n 是一个带有 10 个整数的数组 */

/* 初始化数组 n 中的元素 */
for ( int i = 0; i < 10; i++ )
{
n[i] = i + 100;
}

/* 输出每个数组元素的值 */
foreach (int j in n )
{
int i = j-100;
Console.WriteLine("Element[{0}] = {1}", i, j);
}
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

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

### C# 数组细节

在 C# 中，数组是非常重要的，且需要了解更多的细节。下面列出了 C# 程序员必须清楚的一些与数组相关的重要概念：

概念描述 多维数组C# 支持多维数组。多维数组最简单的形式是二维数组。 交错数组C# 支持交错数组，即数组的数组。 传递数组给函数您可以通过指定不带索引的数组名称来给函数传递一个指向数组的指针。 参数数组这通常用于传递未知数量的参数给函数。 Array 类在 System 命名空间中定义，是所有数组的基类，并提供了各种用于数组的属性和方法。

---

## C# 字符串（String）

Source: https://www.runoob.com/csharp/csharp-string.html

## C# 字符串（String）

在 C# 中，您可以使用字符数组来表示字符串，但是，更常见的做法是使用 string 关键字来声明一个字符串变量。string 关键字是 System.String 类的别名。

### 创建 String 对象

您可以使用以下方法之一来创建 string 对象：

- 通过给 String 变量指定一个字符串
- 通过使用 String 类构造函数
- 通过使用字符串串联运算符（ + ）
- 通过检索属性或调用一个返回字符串的方法
- 通过格式化方法来转换一个值或对象为它的字符串表示形式

下面的实例演示了这点：

### 实例

using System;

namespace StringApplication
{
class Program
{
static void Main(string[] args)
{
//字符串，字符串连接
string fname, lname;
fname = "Rowan";
lname = "Atkinson";

string fullname = fname + lname;
Console.WriteLine("Full Name: {0}", fullname);

//通过使用 string 构造函数
char[] letters = { 'H', 'e', 'l', 'l','o' };
string greetings = new string(letters);
Console.WriteLine("Greetings: {0}", greetings);

//方法返回字符串
string[] sarray = { "Hello", "From", "Tutorials", "Point" };
string message = String.Join(" ", sarray);
Console.WriteLine("Message: {0}", message);

//用于转化值的格式化方法
DateTime waiting = new DateTime(2012, 10, 10, 17, 58, 1);
string chat = String.Format("Message sent at {0:t} on {0:D}",
waiting);
Console.WriteLine("Message: {0}", chat);
Console.ReadKey() ;
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Full Name: RowanAtkinson
Greetings: Hello
Message: Hello From Tutorials Point
Message: Message sent at 17:58 on Wednesday, 10 October 2012

```

### String 类的属性

String 类有以下两个属性：

序号属性名称 & 描述 1Chars
在当前 String 对象中获取 Char 对象的指定位置。 2Length
在当前的 String 对象中获取字符数。

### String 类的方法

String 类有许多方法用于 string 对象的操作。下面的表格提供了一些最常用的方法：

序号方法名称 & 描述 1public static int Compare( string strA, string strB )
比较两个指定的 string 对象，并返回一个表示它们在排列顺序中相对位置的整数。该方法区分大小写。 2public static int Compare( string strA, string strB, bool ignoreCase )
比较两个指定的 string 对象，并返回一个表示它们在排列顺序中相对位置的整数。但是，如果布尔参数为真时，该方法不区分大小写。 3public static string Concat( string str0, string str1 )
连接两个 string 对象。 4public static string Concat( string str0, string str1, string str2 )
连接三个 string 对象。 5public static string Concat( string str0, string str1, string str2, string str3 )
连接四个 string 对象。 6public bool Contains( string value )
返回一个表示指定 string 对象是否出现在字符串中的值。 7public static string Copy( string str )
创建一个与指定字符串具有相同值的新的 String 对象。 8public void CopyTo( int sourceIndex, char[] destination, int destinationIndex, int count )
从 string 对象的指定位置开始复制指定数量的字符到 Unicode 字符数组中的指定位置。 9public bool EndsWith( string value )
判断 string 对象的结尾是否匹配指定的字符串。 10public bool Equals( string value )
判断当前的 string 对象是否与指定的 string 对象具有相同的值。 11public static bool Equals( string a, string b )
判断两个指定的 string 对象是否具有相同的值。 12public static string Format( string format, Object arg0 )
把指定字符串中一个或多个格式项替换为指定对象的字符串表示形式。 13public int IndexOf( char value )
返回指定 Unicode 字符在当前字符串中第一次出现的索引，索引从 0 开始。 14public int IndexOf( string value )
返回指定字符串在该实例中第一次出现的索引，索引从 0 开始。 15public int IndexOf( char value, int startIndex )
返回指定 Unicode 字符从该字符串中指定字符位置开始搜索第一次出现的索引，索引从 0 开始。 16public int IndexOf( string value, int startIndex )
返回指定字符串从该实例中指定字符位置开始搜索第一次出现的索引，索引从 0 开始。 17public int IndexOfAny( char[] anyOf )
返回某一个指定的 Unicode 字符数组中任意字符在该实例中第一次出现的索引，索引从 0 开始。 18public int IndexOfAny( char[] anyOf, int startIndex )
返回某一个指定的 Unicode 字符数组中任意字符从该实例中指定字符位置开始搜索第一次出现的索引，索引从 0 开始。 19public string Insert( int startIndex, string value )
返回一个新的字符串，其中，指定的字符串被插入在当前 string 对象的指定索引位置。 20public static bool IsNullOrEmpty( string value )
指示指定的字符串是否为 null 或者是否为一个空的字符串。 21public static string Join( string separator, string[] value )
连接一个字符串数组中的所有元素，使用指定的分隔符分隔每个元素。 22public static string Join( string separator, string[] value, int startIndex, int count )
连接一个字符串数组中的指定位置开始的指定元素，使用指定的分隔符分隔每个元素。 23public int LastIndexOf( char value )
返回指定 Unicode 字符在当前 string 对象中最后一次出现的索引位置，索引从 0 开始。 24public int LastIndexOf( string value )
返回指定字符串在当前 string 对象中最后一次出现的索引位置，索引从 0 开始。 25public string Remove( int startIndex )
移除当前实例中的所有字符，从指定位置开始，一直到最后一个位置为止，并返回字符串。 26public string Remove( int startIndex, int count )
从当前字符串的指定位置开始移除指定数量的字符，并返回字符串。 27public string Replace( char oldChar, char newChar )
把当前 string 对象中，所有指定的 Unicode 字符替换为另一个指定的 Unicode 字符，并返回新的字符串。 28public string Replace( string oldValue, string newValue )
把当前 string 对象中，所有指定的字符串替换为另一个指定的字符串，并返回新的字符串。 29public string[] Split( params char[] separator )
返回一个字符串数组，包含当前的 string 对象中的子字符串，子字符串是使用指定的 Unicode 字符数组中的元素进行分隔的。 30public string[] Split( char[] separator, int count )
返回一个字符串数组，包含当前的 string 对象中的子字符串，子字符串是使用指定的 Unicode 字符数组中的元素进行分隔的。int 参数指定要返回的子字符串的最大数目。 31public bool StartsWith( string value )
判断字符串实例的开头是否匹配指定的字符串。 32public char[] ToCharArray()
返回一个带有当前 string 对象中所有字符的 Unicode 字符数组。 33public char[] ToCharArray( int startIndex, int length )
返回一个带有当前 string 对象中所有字符的 Unicode 字符数组，从指定的索引开始，直到指定的长度为止。 34public string ToLower()
把字符串转换为小写并返回。 35public string ToUpper()
把字符串转换为大写并返回。 36public string Trim()
移除当前 String 对象中的所有前导空白字符和后置空白字符。

上面的方法列表并不详尽，请访问 MSDN 库，查看完整的方法列表和 String 类构造函数。

### 实例

下面的实例演示了上面提到的一些方法：

比较字符串

### 实例

using System;

namespace StringApplication
{
class StringProg
{
static void Main(string[] args)
{
string str1 = "This is test";
string str2 = "This is text";

if (String.Compare(str1, str2) == 0)
{
Console.WriteLine(str1 + " and " + str2 + " are equal.");
}
else
{
Console.WriteLine(str1 + " and " + str2 + " are not equal.");
}
Console.ReadKey() ;
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

This is test and This is text are not equal.

```

字符串包含字符串：

### 实例

using System;

namespace StringApplication
{
class StringProg
{
static void Main(string[] args)
{
string str = "This is test";
if (str.Contains("test"))
{
Console.WriteLine("The sequence 'test' was found.");
}
Console.ReadKey() ;
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

The sequence 'test' was found.

```

获取子字符串：

### 实例

using System;
namespace StringApplication
{
class StringProg
{
static void Main(string[] args)
{
string str = "Last night I dreamt of San Pedro";
Console.WriteLine(str);
string substr = str.Substring(23);
Console.WriteLine(substr);
Console.ReadKey() ;
}
}
}

运行实例 »

当上面的代码被编译和执行时，它会产生下列结果：

```

Last night I dreamt of San Pedro
San Pedro

```

连接字符串：

### 实例

using System;

namespace StringApplication
{
class StringProg
{
static void Main(string[] args)
{
string[] starray = new string[]{"Down the way nights are dark",
"And the sun shines daily on the mountain top",
"I took a trip on a sailing ship",
"And when I reached Jamaica",
"I made a stop"};

string str = String.Join("\n", starray);
Console.WriteLine(str);
Console.ReadKey() ;
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Down the way nights are dark
And the sun shines daily on the mountain top
I took a trip on a sailing ship
And when I reached Jamaica
I made a stop

```

---

## C# 结构体（Struct）

Source: https://www.runoob.com/csharp/csharp-struct.html

## C# 结构体（Struct）

在 C# 中，结构体（struct）是一种值类型（value type），用于组织和存储相关数据。

在 C# 中，结构体是值类型数据结构，这样使得一个单一变量可以存储各种数据类型的相关数据。

struct 关键字用于创建结构体。

结构体是用来代表一个记录，假设您想跟踪图书馆中书的动态，您可能想跟踪每本书的以下属性：

- Title
- Author
- Subject
- Book ID

### 定义结构体

为了定义一个结构体，您必须使用 struct 语句。

struct 语句为程序定义了一个带有多个成员的新的数据类型。

例如，您可以按照如下的方式声明 Book 结构：

```

struct Books
{
public string title;
public string author;
public string subject;
public int book_id;
};

```

下面的程序演示了结构的用法：

### 实例

using System;
using System.Text;

struct Books
{
public string title;
public string author;
public string subject;
public int book_id;
};

public class testStructure
{
public static void Main(string[] args)
{

Books Book1; /* 声明 Book1，类型为 Books */
Books Book2; /* 声明 Book2，类型为 Books */

/* book 1 详述 */
Book1.title = "C Programming";
Book1.author = "Nuha Ali";
Book1.subject = "C Programming Tutorial";
Book1.book_id = 6495407;

/* book 2 详述 */
Book2.title = "Telecom Billing";
Book2.author = "Zara Ali";
Book2.subject = "Telecom Billing Tutorial";
Book2.book_id = 6495700;

/* 打印 Book1 信息 */
Console.WriteLine( "Book 1 title : {0}", Book1.title);
Console.WriteLine("Book 1 author : {0}", Book1.author);
Console.WriteLine("Book 1 subject : {0}", Book1.subject);
Console.WriteLine("Book 1 book_id :{0}", Book1.book_id);

/* 打印 Book2 信息 */
Console.WriteLine("Book 2 title : {0}", Book2.title);
Console.WriteLine("Book 2 author : {0}", Book2.author);
Console.WriteLine("Book 2 subject : {0}", Book2.subject);
Console.WriteLine("Book 2 book_id : {0}", Book2.book_id);

Console.ReadKey();

}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Book 1 title : C Programming
Book 1 author : Nuha Ali
Book 1 subject : C Programming Tutorial
Book 1 book_id : 6495407
Book 2 title : Telecom Billing
Book 2 author : Zara Ali
Book 2 subject : Telecom Billing Tutorial
Book 2 book_id : 6495700

```

### C# 结构的特点

结构提供了一种轻量级的数据类型，适用于表示简单的数据结构，具有较好的性能特性和值语义：

- 结构可带有方法、字段、索引、属性、运算符方法和事件，适用于表示轻量级数据的情况，如坐标、范围、日期、时间等。
- 结构可定义构造函数，但不能定义析构函数。但是，您不能为结构定义无参构造函数。无参构造函数(默认)是自动定义的，且不能被改变。
- 与类不同，结构不能继承其他的结构或类。
- 结构不能作为其他结构或类的基础结构。
- 结构可实现一个或多个接口。
- 结构成员不能指定为 abstract、virtual 或 protected。
- 当您使用 New 操作符创建一个结构对象时，会调用适当的构造函数来创建结构。与类不同，结构可以不使用 New 操作符即可被实例化。
- 如果不使用 New 操作符，只有在所有的字段都被初始化之后，字段才被赋值，对象才被使用。
- 结构变量通常分配在栈上，这使得它们的创建和销毁速度更快。但是，如果将结构用作类的字段，且这个类是引用类型，那么结构将存储在堆上。
- 结构默认情况下是可变的，这意味着你可以修改它们的字段。但是，如果结构定义为只读，那么它的字段将是不可变的。

### 类 vs 结构

类和结构在设计和使用时有不同的考虑因素，类适合表示复杂的对象和行为，支持继承和多态性，而结构则更适合表示轻量级数据和值类型，以提高性能并避免引用的管理开销。

类和结构有以下几个基本的不同点：

值类型 vs 引用类型：

- 结构是值类型（Value Type）： 结构是值类型，它们在栈上分配内存，而不是在堆上。当将结构实例传递给方法或赋值给另一个变量时，将复制整个结构的内容。
- 类是引用类型（Reference Type）： 类是引用类型，它们在堆上分配内存。当将类实例传递给方法或赋值给另一个变量时，实际上是传递引用（内存地址）而不是整个对象的副本。

继承和多态性：

- 结构不能继承： 结构不能继承其他结构或类，也不能作为其他结构或类的基类。
- 类支持继承： 类支持继承和多态性，可以通过派生新类来扩展现有类的功能。

默认构造函数：

- 结构不能有无参数的构造函数： 结构不能包含无参数的构造函数。
- 类可以有无参数的构造函数： 类可以包含无参数的构造函数，如果没有提供构造函数，系统会提供默认的无参数构造函数。

赋值行为：

- 类型为类的变量在赋值时存储的是引用，因此两个变量指向同一个对象。
- 结构变量在赋值时会复制整个结构，因此每个变量都有自己的独立副本。

传递方式：

- 类型为类的对象在方法调用时通过引用传递，这意味着在方法中对对象所做的更改会影响到原始对象。
- 结构对象通常通过值传递，这意味着传递的是结构的副本，而不是原始结构对象本身。因此，在方法中对结构所做的更改不会影响到原始对象。

可空性：

- 结构体是值类型，不能直接设置为 null：因为 null 是引用类型的默认值，而不是值类型的默认值。如果你需要表示结构体变量的缺失或无效状态，可以使用 Nullable<T> 或称为 T? 的可空类型。
- 类默认可为null： 类的实例默认可以为 `null`，因为它们是引用类型。

性能和内存分配：

- 结构通常更轻量： 由于结构是值类型且在栈上分配内存，它们通常比类更轻量，适用于简单的数据表示。
- 类可能有更多开销： 由于类是引用类型，可能涉及更多的内存开销和管理。

以下实例中，MyStruct 是一个结构，而 MyClass 是一个类。

注释部分演示了结构不能包含无参数的构造函数、不能继承以及结构的实例复制是复制整个结构的内容。与之相反，类可以包含无参数的构造函数，可以继承，并且实例复制是复制引用。

### 实例

using System;

// 结构声明
struct MyStruct
{
public int X;
public int Y;

// 结构不能有无参数的构造函数
// public MyStruct()
// {
// }

// 有参数的构造函数
public MyStruct(int x, int y)
{
X = x;
Y = y;
}

// 结构不能继承
// struct MyDerivedStruct : MyBaseStruct
// {
// }
}

// 类声明
class MyClass
{
public int X;
public int Y;

// 类可以有无参数的构造函数
public MyClass()
{
}

// 有参数的构造函数
public MyClass(int x, int y)
{
X = x;
Y = y;
}

// 类支持继承
// class MyDerivedClass : MyBaseClass
// {
// }
}

class Program
{
static void Main()
{
// 结构是值类型，分配在栈上
MyStruct structInstance1 = new MyStruct(1, 2);
MyStruct structInstance2 = structInstance1; // 复制整个结构

// 类是引用类型，分配在堆上
MyClass classInstance1 = new MyClass(3, 4);
MyClass classInstance2 = classInstance1; // 复制引用，指向同一个对象

// 修改结构实例不影响其他实例
structInstance1.X = 5;
Console.WriteLine($"Struct: {structInstance1.X}, {structInstance2.X}");

// 修改类实例会影响其他实例
classInstance1.X = 6;
Console.WriteLine($"Class: {classInstance1.X}, {classInstance2.X}");
}
}

针对上述讨论，让我们重写前面的实例：

### 实例

using System;
using System.Text;

struct Books
{
private string title;
private string author;
private string subject;
private int book_id;
public void setValues(string t, string a, string s, int id)
{
title = t;
author = a;
subject = s;
book_id =id;
}
public void display()
{
Console.WriteLine("Title : {0}", title);
Console.WriteLine("Author : {0}", author);
Console.WriteLine("Subject : {0}", subject);
Console.WriteLine("Book_id :{0}", book_id);
}

};

public class testStructure
{
public static void Main(string[] args)
{

Books Book1 = new Books(); /* 声明 Book1，类型为 Books */
Books Book2 = new Books(); /* 声明 Book2，类型为 Books */

/* book 1 详述 */
Book1.setValues("C Programming",
"Nuha Ali", "C Programming Tutorial",6495407);

/* book 2 详述 */
Book2.setValues("Telecom Billing",
"Zara Ali", "Telecom Billing Tutorial", 6495700);

/* 打印 Book1 信息 */
Book1.display();

/* 打印 Book2 信息 */
Book2.display();

Console.ReadKey();

}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Title : C Programming
Author : Nuha Ali
Subject : C Programming Tutorial
Book_id : 6495407
Title : Telecom Billing
Author : Zara Ali
Subject : Telecom Billing Tutorial
Book_id : 6495700

```

---

## C# 枚举（Enum）

Source: https://www.runoob.com/csharp/csharp-enum.html

## C# 枚举（Enum）

枚举是一组命名整型常量。枚举类型是使用 enum 关键字声明的。

C# 枚举是值类型。换句话说，枚举包含自己的值，且不能继承或传递继承。

### 声明 enum 变量

声明枚举的一般语法：

```

enum <enum_name>
{
enumeration list
};

```

其中，

- enum_name 指定枚举的类型名称。
- enumeration list 是一个用逗号分隔的标识符列表。

枚举列表中的每个符号代表一个整数值，一个比它前面的符号大的整数值。默认情况下，第一个枚举符号的值是 0.例如：

```

enum Days { Sun, Mon, tue, Wed, thu, Fri, Sat };

```

### 实例

下面的实例演示了枚举变量的用法：

### 实例

using System;

public class EnumTest
{
enum Day { Sun, Mon, Tue, Wed, Thu, Fri, Sat };

static void Main()
{
int x = (int)Day.Sun;
int y = (int)Day.Fri;
Console.WriteLine("Sun = {0}", x);
Console.WriteLine("Fri = {0}", y);
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Sun = 0
Fri = 5

```

---

## C# 类（Class）

Source: https://www.runoob.com/csharp/csharp-class.html

## C# 类（Class）

当你定义一个类时，你定义了一个数据类型的蓝图。这实际上并没有定义任何的数据，但它定义了类的名称意味着什么，也就是说，类的对象由什么组成及在这个对象上可执行什么操作。对象是类的实例。构成类的方法和变量称为类的成员。

### 类的定义

类的定义是以关键字 class 开始，后跟类的名称。类的主体，包含在一对花括号内。下面是类定义的一般形式：

<access specifier> class class_name
{
// member variables
<access specifier> <data type> variable1;
<access specifier> <data type> variable2;
...
<access specifier> <data type> variableN;
// member methods
<access specifier> <return type> method1(parameter_list)
{
// method body
}
<access specifier> <return type> method2(parameter_list)
{
// method body
}
...
<access specifier> <return type> methodN(parameter_list)
{
// method body
}
}

请注意：

- 访问标识符 <access specifier> 指定了对类及其成员的访问规则。如果没有指定，则使用默认的访问标识符。类的默认访问标识符是 internal，成员的默认访问标识符是 private。
- 数据类型 <data type> 指定了变量的类型，返回类型 <return type> 指定了返回的方法返回的数据类型。
- 如果要访问类的成员，你要使用点（.）运算符。
- 点运算符链接了对象的名称和成员的名称。

下面的实例说明了目前为止所讨论的概念：

### 实例

using System;
namespace BoxApplication
{
class Box
{
public double length; // 长度
public double breadth; // 宽度
public double height; // 高度
}
class Boxtester
{
static void Main(string[] args)
{
Box Box1 = new Box(); // 声明 Box1，类型为 Box
Box Box2 = new Box(); // 声明 Box2，类型为 Box
double volume = 0.0; // 体积

// Box1 详述
Box1.height = 5.0;
Box1.length = 6.0;
Box1.breadth = 7.0;

// Box2 详述
Box2.height = 10.0;
Box2.length = 12.0;
Box2.breadth = 13.0;

// Box1 的体积
volume = Box1.height * Box1.length * Box1.breadth;
Console.WriteLine("Box1 的体积： {0}", volume);

// Box2 的体积
volume = Box2.height * Box2.length * Box2.breadth;
Console.WriteLine("Box2 的体积： {0}", volume);
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Box1 的体积： 210
Box2 的体积： 1560

```

### 成员函数和封装

类的成员函数是一个在类定义中有它的定义或原型的函数，就像其他变量一样。作为类的一个成员，它能在类的任何对象上操作，且能访问该对象的类的所有成员。

成员变量是对象的属性（从设计角度），且它们保持私有来实现封装。这些变量只能使用公共成员函数来访问。

让我们使用上面的概念来设置和获取一个类中不同的类成员的值：

### 实例

using System;
namespace BoxApplication
{
class Box
{
private double length; // 长度
private double breadth; // 宽度
private double height; // 高度
public void setLength( double len )
{
length = len;
}

public void setBreadth( double bre )
{
breadth = bre;
}

public void setHeight( double hei )
{
height = hei;
}
public double getVolume()
{
return length * breadth * height;
}
}
class Boxtester
{
static void Main(string[] args)
{
Box Box1 = new Box(); // 声明 Box1，类型为 Box
Box Box2 = new Box(); // 声明 Box2，类型为 Box
double volume; // 体积

// Box1 详述
Box1.setLength(6.0);
Box1.setBreadth(7.0);
Box1.setHeight(5.0);

// Box2 详述
Box2.setLength(12.0);
Box2.setBreadth(13.0);
Box2.setHeight(10.0);

// Box1 的体积
volume = Box1.getVolume();
Console.WriteLine("Box1 的体积： {0}" ,volume);

// Box2 的体积
volume = Box2.getVolume();
Console.WriteLine("Box2 的体积： {0}", volume);

Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Box1 的体积： 210
Box2 的体积： 1560

```

### C# 中的构造函数

类的 构造函数 是类的一个特殊的成员函数，当创建类的新对象时执行。

构造函数的名称与类的名称完全相同，它没有任何返回类型。

下面的实例说明了构造函数的概念：

### 实例

using System;
namespace LineApplication
{
class Line
{
private double length; // 线条的长度
public Line()
{
Console.WriteLine("对象已创建");
}

public void setLength( double len )
{
length = len;
}
public double getLength()
{
return length;
}

static void Main(string[] args)
{
Line line = new Line();
// 设置线条长度
line.setLength(6.0);
Console.WriteLine("线条的长度： {0}", line.getLength());
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

对象已创建
线条的长度： 6

```

默认的构造函数没有任何参数。但是如果你需要一个带有参数的构造函数可以有参数，这种构造函数叫做参数化构造函数。这种技术可以帮助你在创建对象的同时给对象赋初始值，具体请看下面实例：

### 实例

using System;
namespace LineApplication
{
class Line
{
private double length; // 线条的长度
public Line(double len) // 参数化构造函数
{
Console.WriteLine("对象已创建，length = {0}", len);
length = len;
}

public void setLength( double len )
{
length = len;
}
public double getLength()
{
return length;
}

static void Main(string[] args)
{
Line line = new Line(10.0);
Console.WriteLine("线条的长度： {0}", line.getLength());
// 设置线条长度
line.setLength(6.0);
Console.WriteLine("线条的长度： {0}", line.getLength());
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

对象已创建，length = 10
线条的长度： 10
线条的长度： 6

```

### C# 中的析构函数

类的 析构函数 是类的一个特殊的成员函数，当类的对象超出范围时执行。

析构函数的名称是在类的名称前加上一个波浪形（~）作为前缀，它不返回值，也不带任何参数。

析构函数用于在结束程序（比如关闭文件、释放内存等）之前释放资源。析构函数不能继承或重载。

下面的实例说明了析构函数的概念：

### 实例

using System;
namespace LineApplication
{
class Line
{
private double length; // 线条的长度
public Line() // 构造函数
{
Console.WriteLine("对象已创建");
}
~Line() //析构函数
{
Console.WriteLine("对象已删除");
}

public void setLength( double len )
{
length = len;
}
public double getLength()
{
return length;
}

static void Main(string[] args)
{
Line line = new Line();
// 设置线条长度
line.setLength(6.0);
Console.WriteLine("线条的长度： {0}", line.getLength());
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

对象已创建
线条的长度： 6
对象已删除

```

### C# 类的静态成员

我们可以使用 static 关键字把类成员定义为静态的。当我们声明一个类成员为静态时，意味着无论有多少个类的对象被创建，只会有一个该静态成员的副本。

关键字 static 意味着类中只有一个该成员的实例。静态变量用于定义常量，因为它们的值可以通过直接调用类而不需要创建类的实例来获取。静态变量可在成员函数或类的定义外部进行初始化。你也可以在类的定义内部初始化静态变量。

下面的实例演示了静态变量的用法：

### 实例

using System;
namespace StaticVarApplication
{
class StaticVar
{
public static int num;
public void count()
{
num++;
}
public int getNum()
{
return num;
}
}
class StaticTester
{
static void Main(string[] args)
{
StaticVar s1 = new StaticVar();
StaticVar s2 = new StaticVar();
s1.count();
s1.count();
s1.count();
s2.count();
s2.count();
s2.count();
Console.WriteLine("s1 的变量 num： {0}", s1.getNum());
Console.WriteLine("s2 的变量 num： {0}", s2.getNum());
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

s1 的变量 num： 6
s2 的变量 num： 6

```

你也可以把一个成员函数声明为 static。这样的函数只能访问静态变量。静态函数在对象被创建之前就已经存在。下面的实例演示了静态函数的用法：

### 实例

using System;
namespace StaticVarApplication
{
class StaticVar
{
public static int num;
public void count()
{
num++;
}
public static int getNum()
{
return num;
}
}
class StaticTester
{
static void Main(string[] args)
{
StaticVar s = new StaticVar();
s.count();
s.count();
s.count();
Console.WriteLine("变量 num： {0}", StaticVar.getNum());
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

变量 num： 3

```

---

## C# 继承

Source: https://www.runoob.com/csharp/csharp-inheritance.html

## C# 继承

继承是面向对象程序设计中最重要的概念之一。继承允许我们根据一个类来定义另一个类，这使得创建和维护应用程序变得更容易。同时也有利于重用代码和节省开发时间。

当创建一个类时，程序员不需要完全重新编写新的数据成员和成员函数，只需要设计一个新的类，继承了已有的类的成员即可。这个已有的类被称为的基类，这个新的类被称为派生类。

继承的思想实现了 属于（IS-A） 关系。例如，哺乳动物 属于（IS-A） 动物，狗 属于（IS-A） 哺乳动物，因此狗 属于（IS-A） 动物。

### 基类和派生类

一个类可以继承自另一个类，被称为基类（父类）和派生类（子类）。

C# 不支持类的多重继承，但支持接口的多重继承，一个类可以实现多个接口。

概括来说：一个类可以继承多个接口，但只能继承自一个类。

C# 中创建派生类的语法如下：

<访问修饰符> class <基类>
{
...
}
class <派生类> : <基类>
{
...
}

派生类会继承基类的成员（字段、方法、属性等），除非它们被明确地标记为私有（private）。

派生类可以通过关键字base来调用基类的构造函数和方法。

### 实例

class BaseClass
{
public void SomeMethod()
{
// Method implementation
}
}

class DerivedClass : BaseClass
{
public void AnotherMethod()
{
// Accessing base class method
base.SomeMethod();

// Method implementation
}
}

假设，有一个基类 Shape，它的派生类是 Rectangle：

### 实例

using System;
namespace InheritanceApplication
{
class Shape
{
public void setWidth(int w)
{
width = w;
}
public void setHeight(int h)
{
height = h;
}
protected int width;
protected int height;
}

// 派生类
class Rectangle: Shape
{
public int getArea()
{
return (width * height);
}
}

class RectangleTester
{
static void Main(string[] args)
{
Rectangle Rect = new Rectangle();

Rect.setWidth(5);
Rect.setHeight(7);

// 打印对象的面积
Console.WriteLine("总面积： {0}", Rect.getArea());
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

总面积： 35

```

### 基类的初始化

派生类继承了基类的成员变量和成员方法。因此父类对象应在子类对象创建之前被创建。您可以在成员初始化列表中进行父类的初始化。

下面的程序演示了这点：

### 实例

using System;
namespace RectangleApplication
{
class Rectangle
{
// 成员变量
protected double length;
protected double width;
public Rectangle(double l, double w)
{
length = l;
width = w;
}
public double GetArea()
{
return length * width;
}
public void Display()
{
Console.WriteLine("长度： {0}", length);
Console.WriteLine("宽度： {0}", width);
Console.WriteLine("面积： {0}", GetArea());
}
}//end class Rectangle
class Tabletop : Rectangle
{
private double cost;
public Tabletop(double l, double w) : base(l, w)
{ }
public double GetCost()
{
double cost;
cost = GetArea() * 70;
return cost;
}
public void Display()
{
base.Display();
Console.WriteLine("成本： {0}", GetCost());
}
}
class ExecuteRectangle
{
static void Main(string[] args)
{
Tabletop t = new Tabletop(4.5, 7.5);
t.Display();
Console.ReadLine();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

长度： 4.5
宽度： 7.5
面积： 33.75
成本： 2362.5

```

### 继承接口（Interface Inheritance）

一个接口可以继承自一个或多个其他接口，派生接口继承了基接口的所有成员。

派生接口可以扩展基接口的成员列表，但不能改变它们的访问修饰符。

### 实例

interface IBaseInterface
{
void Method1();
}

interface IDerivedInterface : IBaseInterface
{
void Method2();
}

继承接口的实例可以通过以下方式来实现：

### 实例

using System;

// 定义一个基接口
interface IBaseInterface
{
void Method1();
}

// 定义一个派生接口，继承自基接口
interface IDerivedInterface : IBaseInterface
{
void Method2();
}

// 实现派生接口的类
class MyClass : IDerivedInterface
{
public void Method1()
{
Console.WriteLine("Method1 implementation");
}

public void Method2()
{
Console.WriteLine("Method2 implementation");
}
}

class Program
{
static void Main(string[] args)
{
// 创建 MyClass 类的实例
MyClass obj = new MyClass();

// 调用继承自基接口的方法
obj.Method1();

// 调用派生接口新增的方法
obj.Method2();
}
}

以上实例中 MyClass 类实现了 IDerivedInterface 接口，因此必须提供 IDerivedInterface 中定义的所有方法，包括从 IBaseInterface继承的 Method1() 方法。 在 Main 方法中，我们创建了 MyClass 的实例 obj 并调用了它的方法。

输出结果为：

```
Method1 implementation
Method2 implementation
```

### C# 多重继承

多重继承指的是一个类别可以同时从多于一个父类继承行为与特征的功能。与单一继承相对，单一继承指一个类别只可以继承自一个父类。

C# 不支持多重继承。但是，您可以使用接口来实现多重继承。下面的程序演示了这点：

### 实例

using System;
namespace InheritanceApplication
{
class Shape
{
public void setWidth(int w)
{
width = w;
}
public void setHeight(int h)
{
height = h;
}
protected int width;
protected int height;
}

// 基类 PaintCost
public interface PaintCost
{
int getCost(int area);

}
// 派生类
class Rectangle : Shape, PaintCost
{
public int getArea()
{
return (width * height);
}
public int getCost(int area)
{
return area * 70;
}
}
class RectangleTester
{
static void Main(string[] args)
{
Rectangle Rect = new Rectangle();
int area;
Rect.setWidth(5);
Rect.setHeight(7);
area = Rect.getArea();
// 打印对象的面积
Console.WriteLine("总面积： {0}", Rect.getArea());
Console.WriteLine("油漆总成本： ${0}" , Rect.getCost(area));
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

总面积： 35
油漆总成本： $2450

```

---

## C# 多态性

Source: https://www.runoob.com/csharp/csharp-polymorphism.html

## C# 多态性

多态是同一个行为具有多个不同表现形式或形态的能力。

多态性意味着有多重形式。在面向对象编程范式中，多态性往往表现为"一个接口，多个功能"。

多态性可以是静态的或动态的。在静态多态性中，函数的响应是在编译时发生的。在动态多态性中，函数的响应是在运行时发生的。

在 C# 中，每个类型都是多态的，因为包括用户定义类型在内的所有类型都继承自 Object。

多态就是同一个接口，使用不同的实例而执行不同操作，如图所示：

现实中，比如我们按下 F1 键这个动作：

- 如果当前在 Flash 界面下弹出的就是 AS 3 的帮助文档；
- 如果当前在 Word 下弹出的就是 Word 帮助；
- 在 Windows 下弹出的就是 Windows 帮助和支持。

同一个事件发生在不同的对象上会产生不同的结果。

### 静态多态性

在编译时，函数和对象的连接机制被称为早期绑定，也被称为静态绑定。C# 提供了两种技术来实现静态多态性。分别为：

- 函数重载
- 运算符重载

运算符重载将在下一章节讨论，接下来我们将讨论函数重载。

### 函数重载

您可以在同一个范围内对相同的函数名有多个定义。函数的定义必须彼此不同，可以是参数列表中的参数类型不同，也可以是参数个数不同。不同重载只有返回类型不同的函数声明。

下面的实例演示了几个相同的函数 Add()，用于对不同个数参数进行相加处理：

### 实例

using System;
namespace PolymorphismApplication
{
public class TestData
{
public int Add(int a, int b, int c)
{
return a + b + c;
}
public int Add(int a, int b)
{
return a + b;
}
}
class Program
{
static void Main(string[] args)
{
TestData dataClass = new TestData();
int add1 = dataClass.Add(1, 2);
int add2 = dataClass.Add(1, 2, 3);

Console.WriteLine("add1 :" + add1);
Console.WriteLine("add2 :" + add2);
}
}
}

下面的实例演示了几个相同的函数 print()，用于打印不同的数据类型：

### 实例

using System;
namespace PolymorphismApplication
{
class Printdata
{
void print(int i)
{
Console.WriteLine("输出整型: {0}", i );
}

void print(double f)
{
Console.WriteLine("输出浮点型: {0}" , f);
}

void print(string s)
{
Console.WriteLine("输出字符串: {0}", s);
}
static void Main(string[] args)
{
Printdata p = new Printdata();
// 调用 print 来打印整数
p.print(1);
// 调用 print 来打印浮点数
p.print(1.23);
// 调用 print 来打印字符串
p.print("Hello Runoob");
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

输出整型: 1
输出浮点型: 1.23
输出字符串: Hello Runoob

```

### 动态多态性

C# 允许您使用关键字 abstract 创建抽象类，用于提供接口的部分类的实现。当一个派生类继承自该抽象类时，实现即完成。抽象类包含抽象方法，抽象方法可被派生类实现。派生类具有更专业的功能。

请注意，下面是有关抽象类的一些规则：

- 您不能创建一个抽象类的实例。
- 您不能在一个抽象类外部声明一个抽象方法。
- 通过在类定义前面放置关键字 sealed，可以将类声明为密封类。当一个类被声明为 sealed 时，它不能被继承。抽象类不能被声明为 sealed。

下面的程序演示了一个抽象类：

### 实例

using System;
namespace PolymorphismApplication
{
abstract class Shape
{
abstract public int area();
}
class Rectangle: Shape
{
private int length;
private int width;
public Rectangle( int a=0, int b=0)
{
length = a;
width = b;
}
public override int area ()
{
Console.WriteLine("Rectangle 类的面积：");
return (width * length);
}
}

class RectangleTester
{
static void Main(string[] args)
{
Rectangle r = new Rectangle(10, 7);
double a = r.area();
Console.WriteLine("面积： {0}",a);
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Rectangle 类的面积：
面积： 70

```

当有一个定义在类中的函数需要在继承类中实现时，可以使用虚方法。

虚方法是使用关键字 virtual 声明的。

虚方法可以在不同的继承类中有不同的实现。

对虚方法的调用是在运行时发生的。

动态多态性是通过 抽象类 和 虚方法 实现的。

以下实例创建了 Shape 基类，并创建派生类 Circle、 Rectangle、Triangle， Shape 类提供一个名为 Draw 的虚拟方法，在每个派生类中重写该方法以绘制该类的指定形状。

### 实例

using System;
using System.Collections.Generic;

public class Shape
{
public int X { get; private set; }
public int Y { get; private set; }
public int Height { get; set; }
public int Width { get; set; }

// 虚方法
public virtual void Draw()
{
Console.WriteLine("执行基类的画图任务");
}
}

class Circle : Shape
{
public override void Draw()
{
Console.WriteLine("画一个圆形");
base.Draw();
}
}
class Rectangle : Shape
{
public override void Draw()
{
Console.WriteLine("画一个长方形");
base.Draw();
}
}
class Triangle : Shape
{
public override void Draw()
{
Console.WriteLine("画一个三角形");
base.Draw();
}
}

class Program
{
static void Main(string[] args)
{
// 创建一个 List<Shape> 对象，并向该对象添加 Circle、Triangle 和 Rectangle
var shapes = new List<Shape>
{
new Rectangle(),
new Triangle(),
new Circle()
};

// 使用 foreach 循环对该列表的派生类进行循环访问，并对其中的每个 Shape 对象调用 Draw 方法
foreach (var shape in shapes)
{
shape.Draw();
}

Console.WriteLine("按下任意键退出。");
Console.ReadKey();
}

}

当上面的代码被编译和执行时，它会产生下列结果：

```

画一个长方形
执行基类的画图任务
画一个三角形
执行基类的画图任务
画一个圆形
执行基类的画图任务
按下任意键退出。

```

下面的程序演示通过虚方法 area() 来计算不同形状图像的面积：

### 实例

using System;
namespace PolymorphismApplication
{
class Shape
{
protected int width, height;
public Shape( int a=0, int b=0)
{
width = a;
height = b;
}
public virtual int area()
{
Console.WriteLine("父类的面积：");
return 0;
}
}
class Rectangle: Shape
{
public Rectangle( int a=0, int b=0): base(a, b)
{

}
public override int area ()
{
Console.WriteLine("Rectangle 类的面积：");
return (width * height);
}
}
class Triangle: Shape
{
public Triangle(int a = 0, int b = 0): base(a, b)
{

}
public override int area()
{
Console.WriteLine("Triangle 类的面积：");
return (width * height / 2);
}
}
class Caller
{
public void CallArea(Shape sh)
{
int a;
a = sh.area();
Console.WriteLine("面积： {0}", a);
}
}
class Tester
{

static void Main(string[] args)
{
Caller c = new Caller();
Rectangle r = new Rectangle(10, 7);
Triangle t = new Triangle(10, 5);
c.CallArea(r);
c.CallArea(t);
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Rectangle 类的面积：
面积：70
Triangle 类的面积：
面积：25

```

---

## C# 运算符重载

Source: https://www.runoob.com/csharp/csharp-operator-overloading.html

## C# 运算符重载

您可以重定义或重载 C# 中内置的运算符。因此，程序员也可以使用用户自定义类型的运算符。重载运算符是具有特殊名称的函数，是通过关键字 operator 后跟运算符的符号来定义的。与其他函数一样，重载运算符有返回类型和参数列表。

例如，请看下面的函数：

public static Box operator+ (Box b, Box c)
{
Box box = new Box();
box.length = b.length + c.length;
box.breadth = b.breadth + c.breadth;
box.height = b.height + c.height;
return box;
}

上面的函数为用户自定义的类 Box 实现了加法运算符（+）。它把两个 Box 对象的属性相加，并返回相加后的 Box 对象。

### 运算符重载的实现

下面的程序演示了完整的实现：

### 实例

using System;

namespace OperatorOvlApplication
{
class Box
{
private double length; // 长度
private double breadth; // 宽度
private double height; // 高度

public double getVolume()
{
return length * breadth * height;
}
public void setLength( double len )
{
length = len;
}

public void setBreadth( double bre )
{
breadth = bre;
}

public void setHeight( double hei )
{
height = hei;
}
// 重载 + 运算符来把两个 Box 对象相加
public static Box operator+ (Box b, Box c)
{
Box box = new Box();
box.length = b.length + c.length;
box.breadth = b.breadth + c.breadth;
box.height = b.height + c.height;
return box;
}

}

class Tester
{
static void Main(string[] args)
{
Box Box1 = new Box(); // 声明 Box1，类型为 Box
Box Box2 = new Box(); // 声明 Box2，类型为 Box
Box Box3 = new Box(); // 声明 Box3，类型为 Box
double volume = 0.0; // 体积

// Box1 详述
Box1.setLength(6.0);
Box1.setBreadth(7.0);
Box1.setHeight(5.0);

// Box2 详述
Box2.setLength(12.0);
Box2.setBreadth(13.0);
Box2.setHeight(10.0);

// Box1 的体积
volume = Box1.getVolume();
Console.WriteLine("Box1 的体积： {0}", volume);

// Box2 的体积
volume = Box2.getVolume();
Console.WriteLine("Box2 的体积： {0}", volume);

// 把两个对象相加
Box3 = Box1 + Box2;

// Box3 的体积
volume = Box3.getVolume();
Console.WriteLine("Box3 的体积： {0}", volume);
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Box1 的体积： 210
Box2 的体积： 1560
Box3 的体积： 5400

```

### 可重载和不可重载运算符

下表描述了 C# 中运算符重载的能力：

运算符描述 +, -, !, ~, ++, --这些一元运算符只有一个操作数，且可以被重载。 +, -, *, /, %这些二元运算符带有两个操作数，且可以被重载。 ==, !=, <, >, <=, >=这些比较运算符可以被重载。 &&, ||这些条件逻辑运算符不能被直接重载。 +=, -=, *=, /=, %=这些赋值运算符不能被重载。 =, ., ?:, ->, new, is, sizeof, typeof这些运算符不能被重载。

### 实例

针对上述讨论，让我们扩展上面的实例，重载更多的运算符：

### 实例

using System;

namespace OperatorOvlApplication
{
class Box
{
private double length; // 长度
private double breadth; // 宽度
private double height; // 高度

public double getVolume()
{
return length * breadth * height;
}
public void setLength( double len )
{
length = len;
}

public void setBreadth( double bre )
{
breadth = bre;
}

public void setHeight( double hei )
{
height = hei;
}
// 重载 + 运算符来把两个 Box 对象相加
public static Box operator+ (Box b, Box c)
{
Box box = new Box();
box.length = b.length + c.length;
box.breadth = b.breadth + c.breadth;
box.height = b.height + c.height;
return box;
}

public static bool operator == (Box lhs, Box rhs)
{
bool status = false;
if (lhs.length == rhs.length && lhs.height == rhs.height
&& lhs.breadth == rhs.breadth)
{
status = true;
}
return status;
}
public static bool operator !=(Box lhs, Box rhs)
{
bool status = false;
if (lhs.length != rhs.length || lhs.height != rhs.height
|| lhs.breadth != rhs.breadth)
{
status = true;
}
return status;
}
public static bool operator <(Box lhs, Box rhs)
{
bool status = false;
if (lhs.length < rhs.length && lhs.height
< rhs.height && lhs.breadth < rhs.breadth)
{
status = true;
}
return status;
}

public static bool operator >(Box lhs, Box rhs)
{
bool status = false;
if (lhs.length > rhs.length && lhs.height
> rhs.height && lhs.breadth > rhs.breadth)
{
status = true;
}
return status;
}

public static bool operator <=(Box lhs, Box rhs)
{
bool status = false;
if (lhs.length <= rhs.length && lhs.height
<= rhs.height && lhs.breadth <= rhs.breadth)
{
status = true;
}
return status;
}

public static bool operator >=(Box lhs, Box rhs)
{
bool status = false;
if (lhs.length >= rhs.length && lhs.height
>= rhs.height && lhs.breadth >= rhs.breadth)
{
status = true;
}
return status;
}
public override string ToString()
{
return String.Format("({0}, {1}, {2})", length, breadth, height);
}

}

class Tester
{
static void Main(string[] args)
{
Box Box1 = new Box(); // 声明 Box1，类型为 Box
Box Box2 = new Box(); // 声明 Box2，类型为 Box
Box Box3 = new Box(); // 声明 Box3，类型为 Box
Box Box4 = new Box();
double volume = 0.0; // 体积

// Box1 详述
Box1.setLength(6.0);
Box1.setBreadth(7.0);
Box1.setHeight(5.0);

// Box2 详述
Box2.setLength(12.0);
Box2.setBreadth(13.0);
Box2.setHeight(10.0);

// 使用重载的 ToString() 显示两个盒子
Console.WriteLine("Box1： {0}", Box1.ToString());
Console.WriteLine("Box2： {0}", Box2.ToString());

// Box1 的体积
volume = Box1.getVolume();
Console.WriteLine("Box1 的体积： {0}", volume);

// Box2 的体积
volume = Box2.getVolume();
Console.WriteLine("Box2 的体积： {0}", volume);

// 把两个对象相加
Box3 = Box1 + Box2;
Console.WriteLine("Box3： {0}", Box3.ToString());
// Box3 的体积
volume = Box3.getVolume();
Console.WriteLine("Box3 的体积： {0}", volume);

//comparing the boxes
if (Box1 > Box2)
Console.WriteLine("Box1 大于 Box2");
else
Console.WriteLine("Box1 不大于 Box2");
if (Box1 < Box2)
Console.WriteLine("Box1 小于 Box2");
else
Console.WriteLine("Box1 不小于 Box2");
if (Box1 >= Box2)
Console.WriteLine("Box1 大于等于 Box2");
else
Console.WriteLine("Box1 不大于等于 Box2");
if (Box1 <= Box2)
Console.WriteLine("Box1 小于等于 Box2");
else
Console.WriteLine("Box1 不小于等于 Box2");
if (Box1 != Box2)
Console.WriteLine("Box1 不等于 Box2");
else
Console.WriteLine("Box1 等于 Box2");
Box4 = Box3;
if (Box3 == Box4)
Console.WriteLine("Box3 等于 Box4");
else
Console.WriteLine("Box3 不等于 Box4");

Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Box1： (6, 7, 5)
Box2： (12, 13, 10)
Box1 的体积： 210
Box2 的体积： 1560
Box3： (18, 20, 15)
Box3 的体积： 5400
Box1 不大于 Box2
Box1 小于 Box2
Box1 不大于等于 Box2
Box1 小于等于 Box2
Box1 不等于 Box2
Box3 等于 Box4

```

---

## C# 接口（Interface）

Source: https://www.runoob.com/csharp/csharp-interface.html

## C# 接口（Interface）

接口定义了所有类继承接口时应遵循的语法合同。接口定义了语法合同 "是什么" 部分，派生类定义了语法合同 "怎么做" 部分。

接口定义了属性、方法和事件，这些都是接口的成员。接口只包含了成员的声明。成员的定义是派生类的责任。接口提供了派生类应遵循的标准结构。

接口使得实现接口的类或结构在形式上保持一致。

抽象类在某种程度上与接口类似，但是，它们大多只是用在当只有少数方法由基类声明由派生类实现时。

接口本身并不实现任何功能，它只是和声明实现该接口的对象订立一个必须实现哪些行为的契约。

抽象类不能直接实例化，但允许派生出具体的，具有实际功能的类。

### 定义接口: MyInterface.cs

接口使用 interface 关键字声明，它与类的声明类似。接口声明默认是 public 的。下面是一个接口声明的实例：

interface IMyInterface
{
void MethodToImplement();
}

以上代码定义了接口 IMyInterface。通常接口命名以 I 字母开头，这个接口只有一个方法 MethodToImplement()，没有参数和返回值，当然我们可以按照需求设置参数和返回值。

值得注意的是，该方法并没有具体的实现。

#### 接下来我们来实现以上接口：InterfaceImplementer.cs

### 实例

using System;

interface IMyInterface
{
// 接口成员
void MethodToImplement();
}

class InterfaceImplementer : IMyInterface
{
static void Main()
{
InterfaceImplementer iImp = new InterfaceImplementer();
iImp.MethodToImplement();
}

public void MethodToImplement()
{
Console.WriteLine("MethodToImplement() called.");
}
}

InterfaceImplementer 类实现了 IMyInterface 接口，接口的实现与类的继承语法格式类似：

```
class InterfaceImplementer : IMyInterface
```

继承接口后，我们需要实现接口的方法 MethodToImplement() , 方法名必须与接口定义的方法名一致。

### 接口继承: InterfaceInheritance.cs

以下实例定义了两个接口 IMyInterface 和 IParentInterface。

如果一个接口继承其他接口，那么实现类或结构就需要实现所有接口的成员。

以下实例 IMyInterface 继承了 IParentInterface 接口，因此接口实现类必须实现 MethodToImplement() 和 ParentInterfaceMethod() 方法：

### 实例

using System;

interface IParentInterface
{
void ParentInterfaceMethod();
}

interface IMyInterface : IParentInterface
{
void MethodToImplement();
}

class InterfaceImplementer : IMyInterface
{
static void Main()
{
InterfaceImplementer iImp = new InterfaceImplementer();
iImp.MethodToImplement();
iImp.ParentInterfaceMethod();
}

public void MethodToImplement()
{
Console.WriteLine("MethodToImplement() called.");
}

public void ParentInterfaceMethod()
{
Console.WriteLine("ParentInterfaceMethod() called.");
}
}

实例输出结果为：

```

MethodToImplement() called.
ParentInterfaceMethod() called.

```

---

## C# 命名空间（Namespace）

Source: https://www.runoob.com/csharp/csharp-namespace.html

## C# 命名空间（Namespace）

命名空间的设计目的是提供一种让一组名称与其他名称分隔开的方式。在一个命名空间中声明的类的名称与另一个命名空间中声明的相同的类的名称不冲突。

我们举一个计算机系统中的例子，一个文件夹(目录)中可以包含多个文件夹，每个文件夹中不能有相同的文件名，但不同文件夹中的文件可以重名。

### 定义命名空间

命名空间的定义是以关键字 namespace 开始，后跟命名空间的名称，如下所示：

namespace namespace_name
{
// 代码声明
}

为了调用支持命名空间版本的函数或变量，会把命名空间的名称置于前面，如下所示：

```

namespace_name.item_name;

```

下面的程序演示了命名空间的用法：

### 实例

using System;
namespace first_space
{
class namespace_cl
{
public void func()
{
Console.WriteLine("Inside first_space");
}
}
}
namespace second_space
{
class namespace_cl
{
public void func()
{
Console.WriteLine("Inside second_space");
}
}
}
class TestClass
{
static void Main(string[] args)
{
first_space.namespace_cl fc = new first_space.namespace_cl();
second_space.namespace_cl sc = new second_space.namespace_cl();
fc.func();
sc.func();
Console.ReadKey();
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Inside first_space
Inside second_space

```

### using 关键字

using 关键字表明程序使用的是给定命名空间中的名称。例如，我们在程序中使用 System 命名空间，其中定义了类 Console。我们可以只写：

```

Console.WriteLine ("Hello there");

```

我们可以写完全限定名称，如下：

```

System.Console.WriteLine("Hello there");

```

您也可以使用 using 命名空间指令，这样在使用的时候就不用在前面加上命名空间名称。该指令告诉编译器随后的代码使用了指定命名空间中的名称。下面的代码演示了命名空间的应用。

让我们使用 using 指定重写上面的实例：

### 实例

using System;
using first_space;
using second_space;

namespace first_space
{
class abc
{
public void func()
{
Console.WriteLine("Inside first_space");
}
}
}
namespace second_space
{
class efg
{
public void func()
{
Console.WriteLine("Inside second_space");
}
}
}
class TestClass
{
static void Main(string[] args)
{
abc fc = new abc();
efg sc = new efg();
fc.func();
sc.func();
Console.ReadKey();
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Inside first_space
Inside second_space

```

### 嵌套命名空间

命名空间可以被嵌套，即您可以在一个命名空间内定义另一个命名空间，如下所示：

```

namespace namespace_name1
{
// 代码声明
namespace namespace_name2
{
// 代码声明
}
}

```

您可以使用点（.）运算符访问嵌套的命名空间的成员，如下所示：

### 实例

using System;
using SomeNameSpace;
using SomeNameSpace.Nested;

namespace SomeNameSpace
{
public class MyClass
{
static void Main()
{
Console.WriteLine("In SomeNameSpace");
Nested.NestedNameSpaceClass.SayHello();
}
}

// 内嵌命名空间
namespace Nested
{
public class NestedNameSpaceClass
{
public static void SayHello()
{
Console.WriteLine("In Nested");
}
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

In SomeNameSpace
In Nested

```

---

## C# 预处理器指令

Source: https://www.runoob.com/csharp/csharp-preprocessor-directives.html

## C# 预处理器指令

预处理器指令（Preprocessor Directives）指导编译器在实际编译开始之前对信息进行预处理。

通过这些指令，可以控制编译器如何编译文件或编译哪些部分。常见的预处理器指令包括条件编译、宏定义等。

所有的预处理器指令都是以 # 开始，且在一行上，只有空白字符可以出现在预处理器指令之前。

预处理器指令不是语句，所以它们不以分号 ; 结束。

C# 编译器没有一个单独的预处理器，但是，指令被处理时就像是有一个单独的预处理器一样。在 C# 中，预处理器指令用于在条件编译中起作用。与 C 和 C++ 不同的是，它们不是用来创建宏。一个预处理器指令必须是该行上的唯一指令。

### C# 预处理器指令列表

下表列出了 C# 中可用的预处理器指令：

指令描述`#define`定义一个符号，可以用于条件编译。`#undef`取消定义一个符号。`#if`开始一个条件编译块，如果符号被定义则包含代码块。`#elif`如果前面的 `#if` 或 `#elif` 条件不满足，且当前条件满足，则包含代码块。`#else`如果前面的 `#if` 或 `#elif` 条件不满足，则包含代码块。`#endif`结束一个条件编译块。`#warning`生成编译器警告信息。`#error`生成编译器错误信息。`#region`标记一段代码区域，可以在IDE中折叠和展开这段代码，便于代码的组织和阅读。`#endregion`结束一个代码区域。`#line`更改编译器输出中的行号和文件名，可以用于调试或生成工具的代码。`#pragma`用于给编译器发送特殊指令，例如禁用或恢复特定的警告。`#nullable`控制可空性上下文和注释，允许启用或禁用对可空引用类型的编译器检查。

### 实例

#define DEBUG

#if DEBUG
Console.WriteLine("Debug mode");
#elif RELEASE
Console.WriteLine("Release mode");
#else
Console.WriteLine("Other mode");
#endif

#warning This is a warning message
#error This is an error message

#region MyRegion
// Your code here
#endregion

#line 100 "MyFile.cs"
// The next line will be reported as line 100 in MyFile.cs
Console.WriteLine("This is line 100");
#line default
// Line numbering returns to normal

#pragma warning disable 414
private int unusedVariable;
#pragma warning restore 414

#nullable enable
string? nullableString = null;
#nullable disable

### #define 和 #undef 预处理器

#define 用于定义符号（通常用于条件编译），#undef 用于取消定义符号。

```
#define DEBUG

#undef RELEASE
```

#define 允许您定义一个符号，这样，通过使用符号作为传递给 #if 指令的表达式，表达式将返回 true。它的语法如下：

```

#define symbol

```

下面的程序说明了这点：

### 实例

#define PI
using System;
namespace PreprocessorDAppl
{
class Program
{
static void Main(string[] args)
{
#if (PI)
Console.WriteLine("PI is defined");
#else
Console.WriteLine("PI is not defined");
#endif
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

PI is defined

```

### 条件指令：#if, #elif, #else 和 #endif

您可以使用 #if 指令来创建一个条件指令。

条件指令用于测试符号是否为真。如果为真，编译器会执行 #if 和下一个指令之间的代码。

条件指令的语法：

```

#if symbol [operator symbol]...

```

其中，symbol 是要测试的符号名称。您也可以使用 true 和 false，或在符号前放置否定运算符。

常见运算符有：

- == (等于)
- != (不等于)
- && (与)
- || (或)

您也可以用括号把符号和运算符进行分组。条件指令用于在调试版本或编译指定配置时编译代码。一个以 #if 指令开始的条件指令，必须显示地以一个 #endif 指令终止。

```
#define DEBUG

#if DEBUG
Console.WriteLine("Debug mode");
#elif RELEASE
Console.WriteLine("Release mode");
#else
Console.WriteLine("Other mode");
#endif
```

下面的程序演示了条件指令的用法：

### 实例

#define DEBUG
#define VC_V10
using System;
public class TestClass
{
public static void Main()
{

#if (DEBUG && !VC_V10)
Console.WriteLine("DEBUG is defined");
#elif (!DEBUG && VC_V10)
Console.WriteLine("VC_V10 is defined");
#elif (DEBUG && VC_V10)
Console.WriteLine("DEBUG and VC_V10 are defined");
#else
Console.WriteLine("DEBUG and VC_V10 are not defined");
#endif
Console.ReadKey();
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

DEBUG and VC_V10 are defined

```

### #warning 和 #error

#warning 用于生成编译器警告，#error 用于生成编译器错误。

```
#warning This is a warning message
#error This is an error message
```

### #region 和 #endregion

用于代码折叠，使代码更加可读。

```
#region MyRegion
// Your code here
#endregion
```

### #line

用于更改文件行号和文件名的编译器输出。

```
#line 100 "MyFile.cs"
// The next line will be reported as line 100 in MyFile.cs
Console.WriteLine("This is line 100");
#line default
// Line numbering returns to normal
```

### #pragma

用于向编译器发送特殊指令。最常见的用法是禁用特定的警告。

```

#pragma warning disable 414
private int unusedVariable;
#pragma warning restore 414
```

### 使用预处理器指令的注意事项

- 提高代码可读性：使用`#region`可以帮助分隔代码块，提高代码的组织性。
- 条件编译：通过`#if`等指令可以在开发和生产环境中编译不同的代码，方便调试和发布。
- 警告和错误：通过`#warning`和`#error`可以在编译时提示开发人员注意特定问题。

通过正确使用这些预处理器指令，可以更好地控制代码的编译过程，提高代码的灵活性和可维护性。

---

## C# 正则表达式

Source: https://www.runoob.com/csharp/csharp-regular-expressions.html

## C# 正则表达式

正则表达式 是一种用来描述、匹配文本模式的搜索公式。简单来说，你可以把它理解成一套特殊的通配符语言，用来精确地找到、验证或替换字符串中的内容。

例如，想判断用户输入的是不是一个合法的邮箱地址、手机号，或者想从一段文本里提取所有日期——这些都是正则表达式的典型应用场景。

.NET 框架内置了功能完善的正则表达式引擎，通过 `System.Text.RegularExpressions` 命名空间提供支持。

一个正则表达式模式由一个或多个字符、运算符和结构组成，它们共同描述要匹配的文本规则。

如果你还不理解正则表达式可以先阅读我们的正则表达式 - 教程。

### 定义正则表达式

正则表达式由以下几类构建块组成，每类负责不同的匹配功能：

- 字符转义：让特殊字符被当作普通字符处理（如 `\.` 匹配真正的句点）
- 字符类：匹配"某一类"字符（如 `\d` 匹配任意数字）
- 定位点：指定匹配发生的"位置"（如 `^` 表示行首）
- 分组构造：将子模式组合在一起，方便捕获或引用
- 限定符：控制某个元素出现的次数（如 `+` 表示一次或多次）
- 反向引用构造：引用前面已捕获的内容进行再次匹配
- 备用构造：实现"或"逻辑（如 `cat|dog`）
- 替换：在替换操作中引用捕获的内容
- 杂项构造：内联选项、注释等辅助功能

#### 字符转义

正则表达式中，反斜杠字符（`\`）有两个作用：一是将其后的普通字符变成特殊含义（如 `\n` 表示换行），二是将特殊字符转义为原义字符（如 `\.` 匹配真实的点号而非"任意字符"）。

初学者常见误区：在 C# 字符串中，`\` 本身也是转义字符，所以正则中写 `\d`，在 C# 字符串里要写 `"\\d"`，或者使用逐字字符串 `@"\d"`（推荐）。

下表列出了常用的转义字符：

转义字符描述模式匹配 \a与报警 (bell) 符 \u0007 匹配。\a"Warning!" + '\u0007' 中的 "\u0007" \b在字符类中，与退格键 \u0008 匹配。（注意：在字符类外，\b 表示单词边界，见"定位点"部分）[\b]{3,}"\b\b\b\b" 中的 "\b\b\b\b" \t与制表符 \u0009 匹配。常用于匹配 Tab 分隔的文本。(\w+)\t"Name\tAddr\t" 中的 "Name\t" 和 "Addr\t" \r与回车符 \u000D 匹配。（\r 与换行符 \n 不是等效的。Windows 换行通常是 \r\n）\r\n(\w+)"\r\nHello\nWorld." 中的 "\r\nHello" \v与垂直制表符 \u000B 匹配。[\v]{2,}"\v\v\v" 中的 "\v\v\v" \f与换页符 \u000C 匹配。[\f]{2,}"\f\f\f" 中的 "\f\f\f" \n与换行符 \u000A 匹配。Unix/Linux 系统的行尾通常只有 \n。\r\n(\w+)"\r\nHello\nWorld." 中的 "\r\nHello" \e与转义符 \u001B 匹配。\e"\x001B" 中的 "\x001B" \ nnn使用八进制表示形式指定一个字符（nnn 由二到三位数字组成）。\w\040\w"a bc d" 中的 "a b" 和 "c d" \x nn使用十六进制表示形式指定字符（nn 恰好由两位数字组成）。\w\x20\w"a bc d" 中的 "a b" 和 "c d" \c X \c x 匹配 X 或 x 指定的 ASCII 控件字符，其中 X 或 x 是控件字符的字母。\cC"\x0003" 中的 "\x0003" (Ctrl-C) \u nnnn使用十六进制表示形式匹配一个 Unicode 字符（由 nnnn 表示的四位数）。\w\u0020\w"a bc d" 中的 "a b" 和 "c d" \在后面带有不识别的转义字符时，与该字符匹配。\d+[\+-x\*]\d+\d+[\+-x\*\d+"(2+2) * 3*9" 中的 "2+2" 和 "3*9"

#### 字符类

字符类用于匹配"某一类"字符中的任意一个。例如 `[aeiou]` 匹配任意一个元音字母，`\d` 匹配任意一个数字。这是正则表达式中最常用的基础功能之一。

下表列出了字符类：

字符类描述模式匹配 [character_group]匹配 character_group 中的任何单个字符。 默认情况下，匹配区分大小写。[mn]"mat" 中的 "m"，"moon" 中的 "m" 和 "n" [^character_group]非：与不在 character_group 中的任何单个字符匹配。 默认情况下，character_group 中的字符区分大小写。[^aei]"avail" 中的 "v" 和 "l" [ first - last ]字符范围：与从 first 到 last 的范围中的任何单个字符匹配。[b-d][b-d]irds 可以匹配 Birds、 Cirds、 Dirds .通配符：与除 \n 之外的任何单个字符匹配。
若要匹配原意句点字符（. 或 \u002E），您必须在该字符前面加上转义符 (\.)。a.e"have" 中的 "ave"， "mate" 中的 "ate" \p{ name }与 name 指定的 Unicode 通用类别或命名块中的任何单个字符匹配。\p{Lu} "City Lights" 中的 "C" 和 "L" \P{ name }与不在 name 指定的 Unicode 通用类别或命名块中的任何单个字符匹配。\P{Lu}"City" 中的 "i"、 "t" 和 "y" \w与任何单词字符匹配（字母、数字、下划线）。等价于 [a-zA-Z0-9_]（ASCII 范围内）。\w"Room#1" 中的 "R"、 "o"、 "m" 和 "1" \W与任何非单词字符匹配。是 \w 的反义。\W"Room#1" 中的 "#" \s与任何空白字符匹配（空格、制表符、换行等）。\w\s"ID A1.3" 中的 "D " \S与任何非空白字符匹配。是 \s 的反义。\s\S"int __ctr" 中的 " _" \d与任何十进制数字匹配。等价于 [0-9]。\d"4 = IV" 中的 "4" \D匹配不是十进制数的任意字符。是 \d 的反义。\D"4 = IV" 中的 " "、 "="、 " "、 "I" 和 "V"

#### 定位点

定位点（也叫"锚点"）不匹配任何具体字符，而是匹配字符串中的某个位置。它们是"零宽度"的，不消耗任何字符，只是断言当前位置满足某个条件。

例如 `^\d{3}` 表示"字符串开头的三个数字"，`\b` 表示"单词与非单词字符之间的边界"。

下表列出了定位点：

断言描述模式匹配 ^匹配必须从字符串或一行的开头开始。^\d{3}"567-777-" 中的 "567" $匹配必须出现在字符串的末尾或出现在行或字符串末尾的 \n 之前。-\d{4}$"8-12-2012" 中的 "-2012" \A匹配必须出现在字符串的开头（不受多行模式影响，始终是整个字符串的开头）。\A\w{4}"Code-007-" 中的 "Code" \Z匹配必须出现在字符串的末尾或出现在字符串末尾的 \n 之前。-\d{3}\Z"Bond-901-007" 中的 "-007" \z匹配必须出现在字符串的末尾（严格末尾，不允许末尾有 \n）。-\d{3}\z"-901-333" 中的 "-333" \G匹配必须出现在上一个匹配结束的地方。常用于连续匹配场景。\G\(\d\)"(1)(3)(5)[7](9)" 中的 "(1)"、 "(3)" 和 "(5)" \b匹配一个单词边界，也就是指单词和空格间的位置。er\b匹配"never"中的"er"，但不能匹配"verb"中的"er"。 \B匹配非单词边界。er\B匹配"verb"中的"er"，但不能匹配"never"中的"er"。

#### 分组构造

分组构造使用圆括号 `( )` 将正则表达式的一部分括起来，形成一个子表达式。分组有两个主要用途：

- 捕获：将匹配到的子字符串"保存"下来，方便后续提取或在替换中引用。
- 作用域限定：让限定符或备用构造只作用于组内的子表达式。

这一部分比较难于理解，可以阅读 正则表达式-选择 、正则表达式的先行断言(lookahead)和后行断言(lookbehind) 帮助理解。

下表列出了分组构造：

分组构造描述模式匹配 ( subexpression )捕获匹配的子表达式并将其分配到一个从零开始的序号中。(\w)\1"deep" 中的 "ee" (?< name >subexpression)将匹配的子表达式捕获到一个命名组中。命名组比数字编号更易读，推荐在复杂正则中使用。(?< double>\w)\k< double>"deep" 中的 "ee" (?< name1 -name2 >subexpression)定义平衡组定义。用于匹配嵌套结构（如括号配对），属于高级特性。(((?'Open'\()[^\(\)]*)+((?'Close-Open'\))[^\(\)]*)+)*(?(Open)(?!))$"3+2^((1-3)*(3-1))" 中的 "((1-3)*(3-1))" (?: subexpression)定义非捕获组。当只需要分组（限定范围）而不需要保存匹配内容时使用，性能略优于捕获组。Write(?:Line)?"Console.WriteLine()" 中的 "WriteLine" (?imnsx-imnsx:subexpression)应用或禁用 subexpression 中指定的选项。 A\d{2}(?i:\w+)\b"A12xl A12XL a12xl" 中的 "A12xl" 和 "A12XL" (?= subexpression)零宽度正预测先行断言（lookahead）。匹配后面紧跟 subexpression 的位置，但不消耗字符。\w+(?=\.)"He is. The dog ran. The sun is out." 中的 "is"、 "ran" 和 "out" (?! subexpression)零宽度负预测先行断言。匹配后面不跟 subexpression 的位置。\b(?!un)\w+\b"unsure sure unity used" 中的 "sure" 和 "used" (?<=subexpression)零宽度正回顾后发断言（lookbehind）。匹配前面紧接 subexpression 的位置。(?<=19)\d{2}\b"1851 1999 1950 1905 2003" 中的 "99"、"50"和 "05" (?<! subexpression)零宽度负回顾后发断言。匹配前面不接 subexpression 的位置。(?<!wo)man\b"Hi woman Hi man" 中的 "man" (?> subexpression)非回溯（原子组）子表达式。匹配成功后不允许回溯，可以提升某些场景的性能。[13579](?>A+B+)"1ABB 3ABBC 5AB 5AC" 中的 "1ABB"、 "3ABB" 和 "5AB"

### 实例

using System;
using System.Text.RegularExpressions;

public class Example
{
public static void Main()
{
string input = "1851 1999 1950 1905 2003";
string pattern = @"(?<=19)\d{2}\b";

foreach (Match match in Regex.Matches(input, pattern))
Console.WriteLine(match.Value);
}
}

运行实例 »

#### 限定符

限定符指定前面的元素（字符、字符类或分组）必须出现多少次才算匹配成功。

限定符默认是贪婪的，即尽可能多地匹配字符。在限定符后加 `?` 可变为懒惰（非贪婪）模式，即尽可能少地匹配字符。初学者可先掌握 `*`、`+`、`?`、`{n}` 这四种基础限定符。

下表列出了限定符：

限定符描述模式匹配 *匹配上一个元素零次或多次。（零次也算匹配）\d*\.\d".0"、 "19.9"、 "219.9" +匹配上一个元素一次或多次。（至少一次）"be+""been" 中的 "bee"， "bent" 中的 "be" ?匹配上一个元素零次或一次。（即该元素可选）"rai?n""ran"、 "rain" { n }匹配上一个元素恰好 n 次。",\d{3}""1,043.6" 中的 ",043"， "9,876,543,210" 中的 ",876"、 ",543" 和 ",210" { n ,}匹配上一个元素至少 n 次。"\d{2,}""166"、 "29"、 "1930" { n , m }匹配上一个元素至少 n 次，但不多于 m 次。"\d{3,5}""166"， "17668"， "193024" 中的 "19302" *?匹配上一个元素零次或多次，但次数尽可能少（懒惰模式）。\d*?\.\d".0"、 "19.9"、 "219.9" +?匹配上一个元素一次或多次，但次数尽可能少（懒惰模式）。"be+?""been" 中的 "be"， "bent" 中的 "be" ??匹配上一个元素零次或一次，但次数尽可能少（懒惰模式）。"rai??n""ran"、 "rain" { n }?匹配前导元素恰好 n 次。",\d{3}?""1,043.6" 中的 ",043"， "9,876,543,210" 中的 ",876"、 ",543" 和 ",210" { n ,}?匹配上一个元素至少 n 次，但次数尽可能少。"\d{2,}?""166"、 "29" 和 "1930" { n , m }?匹配上一个元素的次数介于 n 和 m 之间，但次数尽可能少。"\d{3,5}?""166"， "17668"， "193024" 中的 "193" 和 "024"

#### 反向引用构造

反向引用允许在同一正则表达式中，引用前面已捕获的分组内容进行再次匹配。例如，用 `(\w)\1` 可以找出像 "ee"、"ll" 这样的连续重复字符。

下表列出了反向引用构造：

反向引用构造描述模式匹配 \ number反向引用。 匹配编号子表达式的值。(\w)\1"seek" 中的 "ee" \k< name >命名反向引用。 匹配命名表达式的值。比数字引用更易读，推荐使用。(?< char>\w)\k< char>"seek" 中的 "ee"

#### 备用构造

备用构造使用竖线 `|` 实现"或"逻辑，让正则表达式可以匹配多个候选模式中的任意一个。类似于编程语言中的 `||` 运算符。

下表列出了备用构造：

备用构造描述模式匹配 |匹配以竖线 (|) 字符分隔的任何一个元素。th(e|is|at)"this is the day. " 中的 "the" 和 "this" (?( expression )yes | no )如果正则表达式模式由 expression 匹配指定，则匹配 yes；否则匹配可选的 no 部分。 expression 被解释为零宽度断言。(?(A)A\d{2}\b|\b\d{3}\b)"A10 C103 910" 中的 "A10" 和 "910" (?( name )yes | no )如果 name 或已命名或已编号的捕获组具有匹配，则匹配 yes；否则匹配可选的 no。(?< quoted>")?(?(quoted).+?"|\S+\s)"Dogs.jpg "Yiska playing.jpg"" 中的 Dogs.jpg 和 "Yiska playing.jpg"

#### 替换

替换语法用于 `Regex.Replace()` 方法的替换模式字符串中，可以通过 `$` 加编号或名称来引用之前捕获的分组内容，实现灵活的文本重组。

下表列出了用于替换的字符：

字符描述模式替换模式输入字符串结果字符串 $number替换按组 number 匹配的子字符串。\b(\w+)(\s)(\w+)\b$3$2$1"one two""two one" ${name}替换按命名组 name 匹配的子字符串。\b(?< word1>\w+)(\s)(?< word2>\w+)\b${word2} ${word1}"one two""two one" $$替换字符"$"。\b(\d+)\s?USD$$$1"103 USD""$103" $&替换整个匹配项的一个副本。(\$*(\d*(\.+\d+)?){1})**$&"$1.30""**$1.30" $`替换匹配前的输入字符串的所有文本。B+$`"AABBCC""AAAACC" $'替换匹配后的输入字符串的所有文本。B+$'"AABBCC""AACCCC" $+替换最后捕获的组。B+(C+)$+"AABBCCDD"AACCDD $_替换整个输入字符串。B+$_"AABBCC""AAAABBCCCC"

#### 杂项构造

下表列出了各种杂项构造：

构造描述实例 (?imnsx-imnsx)在模式中间对诸如不区分大小写这样的选项进行设置或禁用。\bA(?i)b\w+\b 匹配 "ABA Able Act" 中的 "ABA" 和 "Able" (?#注释)内联注释。该注释在第一个右括号处终止。\bA(?#匹配以A开头的单词)\w+\b # [行尾]该注释以非转义的 # 开头，并继续到行的结尾。(?x)\bA\w+\b#匹配以 A 开头的单词

### Regex 类

Regex 类是 .NET 中使用正则表达式的核心类，位于 `System.Text.RegularExpressions` 命名空间下。使用前需要在文件顶部添加：

```
using System.Text.RegularExpressions;
```

下表列出了 Regex 类中一些常用的方法：

序号方法 & 描述 1public bool IsMatch( string input )
判断输入字符串是否包含与正则模式匹配的内容。常用于表单验证，如验证手机号、邮箱格式。 2public bool IsMatch( string input, int startat )
从字符串中指定的起始位置开始进行匹配判断。 3public static bool IsMatch( string input, string pattern )
静态方法，无需先创建 Regex 对象，直接传入模式字符串进行匹配判断。适合一次性使用的场景。 4public MatchCollection Matches( string input )
在输入字符串中搜索所有匹配项，返回 MatchCollection 集合，可用 foreach 遍历每个匹配结果。 5public string Replace( string input, string replacement )
将输入字符串中所有匹配正则模式的内容替换为指定字符串。 6public string[] Split( string input )
按照正则模式定义的分隔符，将输入字符串分割为子字符串数组。比 string.Split() 更灵活，支持复杂分隔规则。

如需了解 Regex 类的完整的属性列表，请参阅微软的 C# 文档。

### 实例 1

下面的实例匹配了以 'S' 开头的单词：

解析：`\b` 是单词边界，`S` 匹配大写字母 S，`\S*` 匹配零个或多个非空白字符，组合起来就是"以 S 开头的完整单词"。

### 实例

using System;
using System.Text.RegularExpressions;

namespace RegExApplication
{
class Program
{
private static void showMatch(string text, string expr)
{
Console.WriteLine("The Expression: " + expr);
MatchCollection mc = Regex.Matches(text, expr);
foreach (Match m in mc)
{
Console.WriteLine(m);
}
}
static void Main(string[] args)
{
string str = "A Thousand Splendid Suns";

Console.WriteLine("Matching words that start with 'S': ");
showMatch(str, @"\bS\S*");
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Matching words that start with 'S':
The Expression: \bS\S*
Splendid
Suns

```

### 实例 2

下面的实例匹配了以 'm' 开头以 'e' 结尾的单词：

解析：`\bm` 匹配以 m 开头的单词，`\S*` 匹配中间的任意非空白字符，`e\b` 要求以 e 结尾且后面是单词边界，合起来就是找所有 m 开头 e 结尾的单词。

### 实例

using System;
using System.Text.RegularExpressions;

namespace RegExApplication
{
class Program
{
private static void showMatch(string text, string expr)
{
Console.WriteLine("The Expression: " + expr);
MatchCollection mc = Regex.Matches(text, expr);
foreach (Match m in mc)
{
Console.WriteLine(m);
}
}
static void Main(string[] args)
{
string str = "make maze and manage to measure it";

Console.WriteLine("Matching words start with 'm' and ends with 'e':");
showMatch(str, @"\bm\S*e\b");
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Matching words start with 'm' and ends with 'e':
The Expression: \bm\S*e\b
make
maze
manage
measure

```

### 实例 3

下面的实例替换掉多余的空格：

解析：`\\s+` 匹配一个或多个连续空白字符（空格、制表符等），将它们全部替换为单个空格，从而合并多余空白。

### 实例

using System;
using System.Text.RegularExpressions;

namespace RegExApplication
{
class Program
{
static void Main(string[] args)
{
string input = "Hello World ";
string pattern = "\\s+";
string replacement = " ";
Regex rgx = new Regex(pattern);
string result = rgx.Replace(input, replacement);

Console.WriteLine("Original String: {0}", input);
Console.WriteLine("Replacement String: {0}", result);
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Original String: Hello World
Replacement String: Hello World

```

### GeneratedRegex：源生成器优化（.NET 7+）

从 .NET 7 开始，C# 引入了 `[GeneratedRegex]` 特性（源生成器），这是对传统 `Regex` 类的重要升级，专门用于对性能敏感或反复使用的正则表达式场景。

#### 传统方式的问题

使用传统 `new Regex(pattern)` 时，正则表达式的解析和编译发生在运行时，每次创建都有开销。虽然可以用 `RegexOptions.Compiled` 提升运行速度，但编译仍在运行时进行，且会增加启动时间和内存占用。

### 传统方式

using System.Text.RegularExpressions;

// 方式一：每次调用都重新解析（最慢）
bool isMatch = Regex.IsMatch(input, @"\d{4}-\d{2}-\d{2}");

// 方式二：静态字段 + Compiled（常见优化做法）
private static readonly Regex DateRegex = new Regex(@"\d{4}-\d{2}-\d{2}", RegexOptions.Compiled);

#### GeneratedRegex 的用法

使用 `[GeneratedRegex]` 特性，编译器会在编译阶段将正则表达式直接生成为高效的 C# 代码，无需运行时解析或编译。

使用条件：需要 .NET 7 或更高版本，且该方法必须是 `partial` 方法，所在类也必须是 `partial` 类。

### GeneratedRegex 基本用法

using System;
using System.Text.RegularExpressions;

namespace RegExApplication
{
// 类必须声明为 partial
partial class Program
{
// 使用 [GeneratedRegex] 特性，编译器自动生成实现
[GeneratedRegex(@"\d{4}-\d{2}-\d{2}")]
private static partial Regex DateRegex();

// 也可以附加选项，例如忽略大小写
[GeneratedRegex(@"\bS\S*", RegexOptions.IgnoreCase)]
private static partial Regex StartsWithSRegex();

static void Main(string[] args)
{
string input = "Today is 2024-06-18, next event: 2025-01-01";

// 像调用普通方法一样使用，返回的是 Regex 实例
foreach (Match m in DateRegex().Matches(input))
{
Console.WriteLine("找到日期：" + m.Value);
}

// 验证是否匹配
Console.WriteLine(StartsWithSRegex().IsMatch("Splendid")); // True
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

找到日期：2024-06-18
找到日期：2025-01-01
True

```

#### GeneratedRegex 与传统 Regex 对比

下表从多个维度对比两种方式，帮助你选择适合的用法：

对比维度传统 RegexGeneratedRegex（源生成器）推荐 最低版本要求.NET Framework / .NET Core 全版本支持.NET 7 及以上— 编译时机运行时解析和编译编译阶段（build time）由 Roslyn 源生成器生成代码GeneratedRegex 执行性能普通：较慢；加 Compiled 选项：较快，但启动开销大最快，接近手写代码的性能，且无启动开销GeneratedRegex 内存占用Compiled 模式会生成 IL 代码，内存占用较高生成的是普通 C# 代码，内存更友好GeneratedRegex AOT 兼容性RegexOptions.Compiled 不兼容 AOT（提前编译）完全兼容 Native AOT，适合发布独立应用GeneratedRegex 代码可读性模式字符串直接写在代码中，较直观需要定义 partial 方法，代码结构略复杂传统 Regex（简单场景） 动态模式支持，模式可以是运行时的变量不支持，模式必须是编译时常量传统 Regex（动态场景） 调试支持普通生成的代码可直接查看和调试，更透明GeneratedRegex 适用场景一次性使用、动态模式、旧版 .NET 项目高频调用、性能敏感、AOT 发布、新版 .NET 项目视场景而定

#### 使用建议

- 如果你使用 .NET 7+ 且正则表达式模式是固定的，优先使用 `[GeneratedRegex]`。
- 如果正则模式需要在运行时动态构建（如根据用户输入生成），则只能使用传统 `new Regex(pattern)`。
- 如果项目需要支持 Native AOT 发布（如 .NET 8 的 AOT 模式），务必使用 `[GeneratedRegex]` 而非 `RegexOptions.Compiled`。
- 对于初学者，可以先用传统方式学习和练习，理解正则原理后，在实际项目中逐步迁移到 `[GeneratedRegex]`。

### 实例：用 GeneratedRegex 验证邮箱格式

using System;
using System.Text.RegularExpressions;

partial class EmailValidator
{
// 定义邮箱验证正则，忽略大小写
[GeneratedRegex(@"^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$", RegexOptions.IgnoreCase)]
private static partial Regex EmailRegex();

public static bool IsValidEmail(string email)
{
return EmailRegex().IsMatch(email);
}

static void Main(string[] args)
{
string[] emails = { "user@example.com", "bad-email", "hello@world.org", "missing@dot" };
foreach (var email in emails)
{
Console.WriteLine($"{email}: {(IsValidEmail(email) ? "有效" : "无效")}");
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

user@example.com: 有效
bad-email: 无效
hello@world.org: 有效
missing@dot: 无效

```

---

## C# 异常处理

Source: https://www.runoob.com/csharp/csharp-exception-handling.html

## C# 异常处理

异常是在程序执行期间出现的问题。C# 中的异常是对程序运行时出现的特殊情况的一种响应，比如尝试除以零。

异常提供了一种把程序控制权从某个部分转移到另一个部分的方式。C# 异常处理时建立在四个关键词之上的：try、catch、finally 和 throw。

- try：一个 try 块标识了一个将被激活的特定的异常的代码块。后跟一个或多个 catch 块。
- catch：程序通过异常处理程序捕获异常。catch 关键字表示异常的捕获。
- finally：finally 块用于执行给定的语句，不管异常是否被抛出都会执行。例如，如果您打开一个文件，不管是否出现异常文件都要被关闭。
- throw：当问题出现时，程序抛出一个异常。使用 throw 关键字来完成。

### 语法

假设一个块将出现异常，一个方法使用 try 和 catch 关键字捕获异常。try/catch 块内的代码为受保护的代码，使用 try/catch 语法如下所示：

try
{
// 引起异常的语句
}
catch( ExceptionName e1 )
{
// 错误处理代码
}
catch( ExceptionName e2 )
{
// 错误处理代码
}
catch( ExceptionName eN )
{
// 错误处理代码
}
finally
{
// 要执行的语句
}

您可以列出多个 catch 语句捕获不同类型的异常，以防 try 块在不同的情况下生成多个异常。

### C# 中的异常类

C# 异常是使用类来表示的。C# 中的异常类主要是直接或间接地派生于 System.Exception 类。System.ApplicationException 和 System.SystemException 类是派生于 System.Exception 类的异常类。

System.ApplicationException 类支持由应用程序生成的异常。所以程序员定义的异常都应派生自该类。

System.SystemException 类是所有预定义的系统异常的基类。

下表列出了一些派生自 System.SystemException 类的预定义的异常类：

异常类描述 System.IO.IOException处理 I/O 错误。 System.IndexOutOfRangeException处理当方法指向超出范围的数组索引时生成的错误。 System.ArrayTypeMismatchException处理当数组类型不匹配时生成的错误。 System.NullReferenceException处理当依从一个空对象时生成的错误。 System.DivideByZeroException处理当除以零时生成的错误。 System.InvalidCastException处理在类型转换期间生成的错误。 System.OutOfMemoryException处理空闲内存不足生成的错误。 System.StackOverflowException处理栈溢出生成的错误。

### 异常处理

C# 以 try 和 catch 块的形式提供了一种结构化的异常处理方案。使用这些块，把核心程序语句与错误处理语句分离开。

这些错误处理块是使用 try、catch 和 finally 关键字实现的。下面是一个当除以零时抛出异常的实例：

### 实例

using System;
namespace ErrorHandlingApplication
{
class DivNumbers
{
int result;
DivNumbers()
{
result = 0;
}
public void division(int num1, int num2)
{
try
{
result = num1 / num2;
}
catch (DivideByZeroException e)
{
Console.WriteLine("Exception caught: {0}", e);
}
finally
{
Console.WriteLine("Result: {0}", result);
}

}
static void Main(string[] args)
{
DivNumbers d = new DivNumbers();
d.division(25, 0);
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Exception caught: System.DivideByZeroException: Attempted to divide by zero.
at ...
Result: 0

```

### 创建用户自定义异常

您也可以定义自己的异常。用户自定义的异常类是派生自 ApplicationException 类。下面的实例演示了这点：

### 实例

using System;
namespace UserDefinedException
{
class TestTemperature
{
static void Main(string[] args)
{
Temperature temp = new Temperature();
try
{
temp.showTemp();
}
catch(TempIsZeroException e)
{
Console.WriteLine("TempIsZeroException: {0}", e.Message);
}
Console.ReadKey();
}
}
}
public class TempIsZeroException: ApplicationException
{
public TempIsZeroException(string message): base(message)
{
}
}
public class Temperature
{
int temperature = 0;
public void showTemp()
{
if(temperature == 0)
{
throw (new TempIsZeroException("Zero Temperature found"));
}
else
{
Console.WriteLine("Temperature: {0}", temperature);
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

TempIsZeroException: Zero Temperature found

```

### 抛出对象

如果异常是直接或间接派生自 System.Exception 类，您可以抛出一个对象。您可以在 catch 块中使用 throw 语句来抛出当前的对象，如下所示：

```

Catch(Exception e)
{
...
Throw e
}

```

---

## C# 文件的输入与输出

Source: https://www.runoob.com/csharp/csharp-file-io.html

## C# 文件的输入与输出

一个 文件 是一个存储在磁盘中带有指定名称和目录路径的数据集合。当打开文件进行读写时，它变成一个 流。

从根本上说，流是通过通信路径传递的字节序列。有两个主要的流：输入流 和 输出流。输入流用于从文件读取数据（读操作），输出流用于向文件写入数据（写操作）。

### C# I/O 类

System.IO 命名空间有各种不同的类，用于执行各种文件操作，如创建和删除文件、读取或写入文件，关闭文件等。

下表列出了一些 System.IO 命名空间中常用的非抽象类：

I/O 类描述 BinaryReader从二进制流读取原始数据。 BinaryWriter以二进制格式写入原始数据。 BufferedStream字节流的临时存储。 Directory有助于操作目录结构。 DirectoryInfo用于对目录执行操作。 DriveInfo提供驱动器的信息。 File有助于处理文件。 FileInfo用于对文件执行操作。 FileStream用于文件中任何位置的读写。 MemoryStream用于随机访问存储在内存中的数据流。 Path对路径信息执行操作。 StreamReader用于从字节流中读取字符。 StreamWriter用于向一个流中写入字符。 StringReader用于读取字符串缓冲区。 StringWriter用于写入字符串缓冲区。

### FileStream 类

System.IO 命名空间中的 FileStream 类有助于文件的读写与关闭。该类派生自抽象类 Stream。

您需要创建一个 FileStream 对象来创建一个新的文件，或打开一个已有的文件。创建 FileStream 对象的语法如下：

```

FileStream <object_name> = new FileStream( <file_name>,
<FileMode Enumerator>, <FileAccess Enumerator>, <FileShare Enumerator>);

```

例如，创建一个 FileStream 对象 F 来读取名为 sample.txt 的文件：

```

FileStream F = new FileStream("sample.txt", FileMode.Open, FileAccess.Read, FileShare.Read);

```

参数描述 FileMode

FileMode 枚举定义了各种打开文件的方法。FileMode 枚举的成员有：

- Append：打开一个已有的文件，并将光标放置在文件的末尾。如果文件不存在，则创建文件。
- Create：创建一个新的文件。如果文件已存在，则删除旧文件，然后创建新文件。
- CreateNew：指定操作系统应创建一个新的文件。如果文件已存在，则抛出异常。
- Open：打开一个已有的文件。如果文件不存在，则抛出异常。
- OpenOrCreate：指定操作系统应打开一个已有的文件。如果文件不存在，则用指定的名称创建一个新的文件打开。
- Truncate：打开一个已有的文件，文件一旦打开，就将被截断为零字节大小。然后我们可以向文件写入全新的数据，但是保留文件的初始创建日期。如果文件不存在，则抛出异常。 FileAccess

FileAccess 枚举的成员有：Read、ReadWrite 和 Write。 FileShare

FileShare 枚举的成员有：

- Inheritable：允许文件句柄可由子进程继承。Win32 不直接支持此功能。
- None：谢绝共享当前文件。文件关闭前，打开该文件的任何请求（由此进程或另一进程发出的请求）都将失败。
- Read：允许随后打开文件读取。如果未指定此标志，则文件关闭前，任何打开该文件以进行读取的请求（由此进程或另一进程发出的请求）都将失败。但是，即使指定了此标志，仍可能需要附加权限才能够访问该文件。
- ReadWrite：允许随后打开文件读取或写入。如果未指定此标志，则文件关闭前，任何打开该文件以进行读取或写入的请求（由此进程或另一进程发出）都将失败。但是，即使指定了此标志，仍可能需要附加权限才能够访问该文件。
- Write：允许随后打开文件写入。如果未指定此标志，则文件关闭前，任何打开该文件以进行写入的请求（由此进程或另一进过程发出的请求）都将失败。但是，即使指定了此标志，仍可能需要附加权限才能够访问该文件。
- Delete：允许随后删除文件。

### 实例

下面的程序演示了 FileStream 类的用法：

### 实例

using System;
using System.IO;

namespace FileIOApplication
{
class Program
{
static void Main(string[] args)
{
FileStream F = new FileStream("test.dat",
FileMode.OpenOrCreate, FileAccess.ReadWrite);

for (int i = 1; i <= 20; i++)
{
F.WriteByte((byte)i);
}

F.Position = 0;

for (int i = 0; i <= 20; i++)
{
Console.Write(F.ReadByte() + " ");
}
F.Close();
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 -1

```

### C# 高级文件操作

上面的实例演示了 C# 中简单的文件操作。但是，要充分利用 C# System.IO 类的强大功能，您需要知道这些类常用的属性和方法。

在下面的章节中，我们将讨论这些类和它们执行的操作。请单击链接详细了解各个部分的知识：

主题描述 文本文件的读写它涉及到文本文件的读写。StreamReader 和 StreamWriter 类有助于完成文本文件的读写。 二进制文件的读写它涉及到二进制文件的读写。BinaryReader 和 BinaryWriter 类有助于完成二进制文件的读写。 Windows 文件系统的操作它让 C# 程序员能够浏览并定位 Windows 文件和目录。

---

## C# 特性（Attribute）

Source: https://www.runoob.com/csharp/csharp-attribute.html

## C# 特性（Attribute）

特性（Attribute）是用于在运行时传递程序中各种元素（比如类、方法、结构、枚举、组件等）的行为信息的声明性标签。您可以通过使用特性向程序添加声明性信息。一个声明性标签是通过放置在它所应用的元素前面的方括号（[ ]）来描述的。

特性（Attribute）用于添加元数据，如编译器指令和注释、描述、方法、类等其他信息。.Net 框架提供了两种类型的特性：预定义特性和自定义特性。

### 规定特性（Attribute）

规定特性（Attribute）的语法如下：

```

[attribute(positional_parameters, name_parameter = value, ...)]
element

```

特性（Attribute）的名称和值是在方括号内规定的，放置在它所应用的元素之前。positional_parameters 规定必需的信息，name_parameter 规定可选的信息。

### 预定义特性（Attribute）

.Net 框架提供了三种预定义特性：

- AttributeUsage
- Conditional
- Obsolete

#### AttributeUsage

预定义特性 AttributeUsage 描述了如何使用一个自定义特性类。它规定了特性可应用到的项目的类型。

规定该特性的语法如下：

```

[AttributeUsage(
validon,
AllowMultiple=allowmultiple,
Inherited=inherited
)]

```

其中：

- 参数 validon 规定特性可被放置的语言元素。它是枚举器 AttributeTargets 的值的组合。默认值是 AttributeTargets.All。
- 参数 allowmultiple（可选的）为该特性的 AllowMultiple 属性（property）提供一个布尔值。如果为 true，则该特性是多用的。默认值是 false（单用的）。
- 参数 inherited（可选的）为该特性的 Inherited 属性（property）提供一个布尔值。如果为 true，则该特性可被派生类继承。默认值是 false（不被继承）。

例如：

```

[AttributeUsage(AttributeTargets.Class |
AttributeTargets.Constructor |
AttributeTargets.Field |
AttributeTargets.Method |
AttributeTargets.Property,
AllowMultiple = true)]

```

#### Conditional

这个预定义特性标记了一个条件方法，其执行依赖于指定的预处理标识符。

它会引起方法调用的条件编译，取决于指定的值，比如 Debug 或 Trace。例如，当调试代码时显示变量的值。

规定该特性的语法如下：

```

[Conditional(
conditionalSymbol
)]

```

例如：

```

[Conditional("DEBUG")]

```

下面的实例演示了该特性：

### 实例

#define DEBUG
using System;
using System.Diagnostics;
public class Myclass
{
[Conditional("DEBUG")]
public static void Message(string msg)
{
Console.WriteLine(msg);
}
}
class Test
{
static void function1()
{
Myclass.Message("In Function 1.");
function2();
}
static void function2()
{
Myclass.Message("In Function 2.");
}
public static void Main()
{
Myclass.Message("In Main function.");
function1();
Console.ReadKey();
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

In Main function.
In Function 1.
In Function 2.

```

#### Obsolete

这个预定义特性标记了不应被使用的程序实体。它可以让您通知编译器丢弃某个特定的目标元素。例如，当一个新方法被用在一个类中，但是您仍然想要保持类中的旧方法，您可以通过显示一个应该使用新方法，而不是旧方法的消息，来把它标记为 obsolete（过时的）。

规定该特性的语法如下：

```

[Obsolete(
message
)]
[Obsolete(
message,
iserror
)]

```

其中：

- 参数 message，是一个字符串，描述项目为什么过时以及该替代使用什么。
- 参数 iserror，是一个布尔值。如果该值为 true，编译器应把该项目的使用当作一个错误。默认值是 false（编译器生成一个警告）。

下面的实例演示了该特性：

### 实例

using System;
public class MyClass
{
[Obsolete("Don't use OldMethod, use NewMethod instead", true)]
static void OldMethod()
{
Console.WriteLine("It is the old method");
}
static void NewMethod()
{
Console.WriteLine("It is the new method");
}
public static void Main()
{
OldMethod();
}
}

当您尝试编译该程序时，编译器会给出一个错误消息说明：

```

Don't use OldMethod, use NewMethod instead

```

### 创建自定义特性（Attribute）

.Net 框架允许创建自定义特性，用于存储声明性的信息，且可在运行时被检索。该信息根据设计标准和应用程序需要，可与任何目标元素相关。

创建并使用自定义特性包含四个步骤：

- 声明自定义特性
- 构建自定义特性
- 在目标程序元素上应用自定义特性
- 通过反射访问特性

最后一个步骤包含编写一个简单的程序来读取元数据以便查找各种符号。元数据是用于描述其他数据的数据和信息。该程序应使用反射来在运行时访问特性。我们将在下一章详细讨论这点。

#### 声明自定义特性

一个新的自定义特性应派生自 System.Attribute 类。例如：

```

// 一个自定义特性 BugFix 被赋给类及其成员
[AttributeUsage(AttributeTargets.Class |
AttributeTargets.Constructor |
AttributeTargets.Field |
AttributeTargets.Method |
AttributeTargets.Property,
AllowMultiple = true)]

public class DeBugInfo : System.Attribute

```

在上面的代码中，我们已经声明了一个名为 DeBugInfo 的自定义特性。

#### 构建自定义特性

让我们构建一个名为 DeBugInfo 的自定义特性，该特性将存储调试程序获得的信息。它存储下面的信息：

- bug 的代码编号
- 辨认该 bug 的开发人员名字
- 最后一次审查该代码的日期
- 一个存储了开发人员标记的字符串消息

我们的 DeBugInfo 类将带有三个用于存储前三个信息的私有属性（property）和一个用于存储消息的公有属性（property）。所以 bug 编号、开发人员名字和审查日期将是 DeBugInfo 类的必需的定位（ positional）参数，消息将是一个可选的命名（named）参数。

每个特性必须至少有一个构造函数。必需的定位（ positional）参数应通过构造函数传递。下面的代码演示了 DeBugInfo 类：

### 实例

// 一个自定义特性 BugFix 被赋给类及其成员
[AttributeUsage(AttributeTargets.Class |
AttributeTargets.Constructor |
AttributeTargets.Field |
AttributeTargets.Method |
AttributeTargets.Property,
AllowMultiple = true)]

public class DeBugInfo : System.Attribute
{
private int bugNo;
private string developer;
private string lastReview;
public string message;

public DeBugInfo(int bg, string dev, string d)
{
this.bugNo = bg;
this.developer = dev;
this.lastReview = d;
}

public int BugNo
{
get
{
return bugNo;
}
}
public string Developer
{
get
{
return developer;
}
}
public string LastReview
{
get
{
return lastReview;
}
}
public string Message
{
get
{
return message;
}
set
{
message = value;
}
}
}

#### 应用自定义特性

通过把特性放置在紧接着它的目标之前，来应用该特性：

### 实例

[DeBugInfo(45, "Zara Ali", "12/8/2012", Message = "Return type mismatch")]
[DeBugInfo(49, "Nuha Ali", "10/10/2012", Message = "Unused variable")]
class Rectangle
{
// 成员变量
protected double length;
protected double width;
public Rectangle(double l, double w)
{
length = l;
width = w;
}
[DeBugInfo(55, "Zara Ali", "19/10/2012",
Message = "Return type mismatch")]
public double GetArea()
{
return length * width;
}
[DeBugInfo(56, "Zara Ali", "19/10/2012")]
public void Display()
{
Console.WriteLine("Length: {0}", length);
Console.WriteLine("Width: {0}", width);
Console.WriteLine("Area: {0}", GetArea());
}
}

在下一章中，我们将使用 Reflection 类对象来检索这些信息。

---

## C# 反射（Reflection）

Source: https://www.runoob.com/csharp/csharp-reflection.html

## C# 反射（Reflection）

反射指程序可以访问、检测和修改它本身状态或行为的一种能力。

程序集包含模块，而模块包含类型，类型又包含成员。反射则提供了封装程序集、模块和类型的对象。

您可以使用反射动态地创建类型的实例，将类型绑定到现有对象，或从现有对象中获取类型。然后，可以调用类型的方法或访问其字段和属性。

#### 优缺点

优点：

- 1、反射提高了程序的灵活性和扩展性。
- 2、降低耦合性，提高自适应能力。
- 3、它允许程序创建和控制任何类的对象，无需提前硬编码目标类。

缺点：

- 1、性能问题：使用反射基本上是一种解释操作，用于字段和方法接入时要远慢于直接代码。因此反射机制主要应用在对灵活性和拓展性要求很高的系统框架上，普通程序不建议使用。
- 2、使用反射会模糊程序内部逻辑；程序员希望在源代码中看到程序的逻辑，反射却绕过了源代码的技术，因而会带来维护的问题，反射代码比相应的直接代码更复杂。

### 反射（Reflection）的用途

反射（Reflection）有下列用途：

- 它允许在运行时查看特性（attribute）信息。
- 它允许审查集合中的各种类型，以及实例化这些类型。
- 它允许延迟绑定的方法和属性（property）。
- 它允许在运行时创建新类型，然后使用这些类型执行一些任务。

### 查看元数据

我们已经在上面的章节中提到过，使用反射（Reflection）可以查看特性（attribute）信息。

System.Reflection 类的 MemberInfo 对象需要被初始化，用于发现与类相关的特性（attribute）。为了做到这点，您可以定义目标类的一个对象，如下：

```

System.Reflection.MemberInfo info = typeof(MyClass);

```

下面的程序演示了这点：

### 实例

using System;

[AttributeUsage(AttributeTargets.All)]
public class HelpAttribute : System.Attribute
{
public readonly string Url;

public string Topic // Topic 是一个命名（named）参数
{
get
{
return topic;
}
set
{

topic = value;
}
}

public HelpAttribute(string url) // url 是一个定位（positional）参数
{
this.Url = url;
}

private string topic;
}
[HelpAttribute("Information on the class MyClass")]
class MyClass
{
}

namespace AttributeAppl
{
class Program
{
static void Main(string[] args)
{
System.Reflection.MemberInfo info = typeof(MyClass);
object[] attributes = info.GetCustomAttributes(true);
for (int i = 0; i < attributes.Length; i++)
{
System.Console.WriteLine(attributes[i]);
}
Console.ReadKey();

}
}
}

当上面的代码被编译和执行时，它会显示附加到类 MyClass 上的自定义特性：

```

HelpAttribute

```

### 实例

在本实例中，我们将使用在上一章中创建的 DeBugInfo 特性，并使用反射（Reflection）来读取 Rectangle 类中的元数据。

### 实例

using System;
using System.Reflection;
namespace BugFixApplication
{
// 一个自定义特性 BugFix 被赋给类及其成员
[AttributeUsage(AttributeTargets.Class |
AttributeTargets.Constructor |
AttributeTargets.Field |
AttributeTargets.Method |
AttributeTargets.Property,
AllowMultiple = true)]

public class DeBugInfo : System.Attribute
{
private int bugNo;
private string developer;
private string lastReview;
public string message;

public DeBugInfo(int bg, string dev, string d)
{
this.bugNo = bg;
this.developer = dev;
this.lastReview = d;
}

public int BugNo
{
get
{
return bugNo;
}
}
public string Developer
{
get
{
return developer;
}
}
public string LastReview
{
get
{
return lastReview;
}
}
public string Message
{
get
{
return message;
}
set
{
message = value;
}
}
}
[DeBugInfo(45, "Zara Ali", "12/8/2012",
Message = "Return type mismatch")]
[DeBugInfo(49, "Nuha Ali", "10/10/2012",
Message = "Unused variable")]
class Rectangle
{
// 成员变量
protected double length;
protected double width;
public Rectangle(double l, double w)
{
length = l;
width = w;
}
[DeBugInfo(55, "Zara Ali", "19/10/2012",
Message = "Return type mismatch")]
public double GetArea()
{
return length * width;
}
[DeBugInfo(56, "Zara Ali", "19/10/2012")]
public void Display()
{
Console.WriteLine("Length: {0}", length);
Console.WriteLine("Width: {0}", width);
Console.WriteLine("Area: {0}", GetArea());
}
}//end class Rectangle

class ExecuteRectangle
{
static void Main(string[] args)
{
Rectangle r = new Rectangle(4.5, 7.5);
r.Display();
Type type = typeof(Rectangle);
// 遍历 Rectangle 类的特性
foreach (Object attributes in type.GetCustomAttributes(false))
{
DeBugInfo dbi = (DeBugInfo)attributes;
if (null != dbi)
{
Console.WriteLine("Bug no: {0}", dbi.BugNo);
Console.WriteLine("Developer: {0}", dbi.Developer);
Console.WriteLine("Last Reviewed: {0}",
dbi.LastReview);
Console.WriteLine("Remarks: {0}", dbi.Message);
}
}

// 遍历方法特性
foreach (MethodInfo m in type.GetMethods())
{
foreach (Attribute a in m.GetCustomAttributes(true))
{
DeBugInfo dbi = (DeBugInfo)a;
if (null != dbi)
{
Console.WriteLine("Bug no: {0}, for Method: {1}",
dbi.BugNo, m.Name);
Console.WriteLine("Developer: {0}", dbi.Developer);
Console.WriteLine("Last Reviewed: {0}",
dbi.LastReview);
Console.WriteLine("Remarks: {0}", dbi.Message);
}
}
}
Console.ReadLine();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Length: 4.5
Width: 7.5
Area: 33.75
Bug No: 49
Developer: Nuha Ali
Last Reviewed: 10/10/2012
Remarks: Unused variable
Bug No: 45
Developer: Zara Ali
Last Reviewed: 12/8/2012
Remarks: Return type mismatch
Bug No: 55, for Method: GetArea
Developer: Zara Ali
Last Reviewed: 19/10/2012
Remarks: Return type mismatch
Bug No: 56, for Method: Display
Developer: Zara Ali
Last Reviewed: 19/10/2012
Remarks:

```

---

## C# 属性（Property）

Source: https://www.runoob.com/csharp/csharp-property.html

## C# 属性（Property）

C# 中的属性（Property）是类和结构体中用于封装数据的成员。它们提供了一种方式来定义类成员的访问和设置规则，通常用于隐藏字段（Fields）的内部实现细节，同时提供控制数据访问的机制。

属性可以看作是对字段的包装器，通常由 get 和 set 访问器组成。

属性（Property）不会确定存储位置。相反，它们具有可读写或计算它们值的 访问器（accessors）。

例如，有一个名为 Student 的类，带有 age、name 和 code 的私有域。我们不能在类的范围以外直接访问这些域，但是我们可以拥有访问这些私有域的属性。

#### 基本语法

```
public class Person
{
private string name;

public string Name
{
get { return name; }
set { name = value; }
}
}
```

以上代码中，`Name` 属性封装了私有字段 `name`。`get` 访问器用于获取字段值，而 `set` 访问器用于设置字段值。

#### 自动实现的属性

如果你只需要一个简单的属性，C# 允许使用自动实现的属性，这样你不需要显式地定义字段。

```
public class Person
{
public string Name { get; set; }
}
```

在这种情况下，编译器会自动为 Name 属性生成一个私有的匿名字段来存储值。

#### 只读属性

如果你只需要一个只读属性，可以省略 set 访问器。

```

public class Person
{
public string Name { get; }

public Person(string name)
{
Name = name;
}
}
```

#### 只写属性

类似地，如果你只需要一个只写属性，可以省略 get 访问器。

```
public class Person
{
private string name;

public string Name
{
set { name = value; }
}
}
```

#### 自定义逻辑

你可以在 get 和 set 访问器中包含自定义的逻辑。

```
public class Person
{
private string name;

public string Name
{
get { return name; }
set
{
if (string.IsNullOrWhiteSpace(value))
throw new ArgumentException("Name cannot be empty.");
name = value;
}
}
}
```

#### 计算属性

属性也可以是计算的，不依赖于字段。

```
public class Rectangle
{
public int Width { get; set; }
public int Height { get; set; }

public int Area
{
get { return Width * Height; }
}
}
```

### 访问器（Accessors）

属性（Property）的访问器（accessor）包含有助于获取（读取或计算）或设置（写入）属性的可执行语句。访问器（accessor）声明可包含一个 get 访问器、一个 set 访问器，或者同时包含二者。例如：

// 声明类型为 string 的 Code 属性
public string Code
{
get
{
return code;
}
set
{
code = value;
}
}

// 声明类型为 string 的 Name 属性
public string Name
{
get
{
return name;
}
set
{
name = value;
}
}

// 声明类型为 int 的 Age 属性
public int Age
{
get
{
return age;
}
set
{
age = value;
}
}

### 实例

下面的实例演示了属性（Property）的用法：

### 实例

using System;
namespace runoob
{
class Student
{

private string code = "N.A";
private string name = "not known";
private int age = 0;

// 声明类型为 string 的 Code 属性
public string Code
{
get
{
return code;
}
set
{
code = value;
}
}

// 声明类型为 string 的 Name 属性
public string Name
{
get
{
return name;
}
set
{
name = value;
}
}

// 声明类型为 int 的 Age 属性
public int Age
{
get
{
return age;
}
set
{
age = value;
}
}
public override string ToString()
{
return "Code = " + Code +", Name = " + Name + ", Age = " + Age;
}
}
class ExampleDemo
{
public static void Main()
{
// 创建一个新的 Student 对象
Student s = new Student();

// 设置 student 的 code、name 和 age
s.Code = "001";
s.Name = "Zara";
s.Age = 9;
Console.WriteLine("Student Info: {0}", s);
// 增加年龄
s.Age += 1;
Console.WriteLine("Student Info: {0}", s);
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Student Info: Code = 001, Name = Zara, Age = 9
Student Info: Code = 001, Name = Zara, Age = 10

```

### 抽象属性（Abstract Properties）

抽象类可拥有抽象属性，这些属性应在派生类中被实现。下面的程序说明了这点：

### 实例

using System;
namespace Runoob
{
public abstract class Person
{
public abstract string Name { get; set; }
public abstract int Age { get; set; }
}

class Student : Person
{
// 声明自动实现的属性
public string Code { get; set; } = "N.A";
public override string Name { get; set; } = "N.A";
public override int Age { get; set; } = 0;

public override string ToString()
{
return $"Code = {Code}, Name = {Name}, Age = {Age}";
}
}

class ExampleDemo
{
public static void Main()
{
// 创建一个新的 Student 对象
Student s = new Student
{
Code = "001",
Name = "Zara",
Age = 9
};

Console.WriteLine("Student Info:- {0}", s);

// 增加年龄
s.Age += 1;
Console.WriteLine("Student Info:- {0}", s);

Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Student Info: Code = 001, Name = Zara, Age = 9
Student Info: Code = 001, Name = Zara, Age = 10

```

---

## C# 索引器（Indexer）

Source: https://www.runoob.com/csharp/csharp-indexer.html

## C# 索引器（Indexer）

索引器（Indexer） 允许一个对象可以像数组一样使用下标的方式来访问。

当您为类定义一个索引器时，该类的行为就会像一个 虚拟数组（virtual array） 一样。您可以使用数组访问运算符 [ ] 来访问该类的的成员。

### 语法

一维索引器的语法如下：

element-type this[int index]
{
// get 访问器
get
{
// 返回 index 指定的值
}

// set 访问器
set
{
// 设置 index 指定的值
}
}

### 索引器（Indexer）的用途

索引器的行为的声明在某种程度上类似于属性（property）。就像属性（property），您可使用 get 和 set 访问器来定义索引器。但是，属性返回或设置一个特定的数据成员，而索引器返回或设置对象实例的一个特定值。换句话说，它把实例数据分为更小的部分，并索引每个部分，获取或设置每个部分。

定义一个属性（property）包括提供属性名称。索引器定义的时候不带有名称，但带有 this 关键字，它指向对象实例。下面的实例演示了这个概念：

### 实例

using System;
namespace IndexerApplication
{
class IndexedNames
{
private string[] namelist = new string[size];
static public int size = 10;
public IndexedNames()
{
for (int i = 0; i < size; i++)
namelist[i] = "N. A.";
}
public string this[int index]
{
get
{
string tmp;

if( index >= 0 && index <= size-1 )
{
tmp = namelist[index];
}
else
{
tmp = "";
}

return ( tmp );
}
set
{
if( index >= 0 && index <= size-1 )
{
namelist[index] = value;
}
}
}

static void Main(string[] args)
{
IndexedNames names = new IndexedNames();
names[0] = "Zara";
names[1] = "Riz";
names[2] = "Nuha";
names[3] = "Asif";
names[4] = "Davinder";
names[5] = "Sunil";
names[6] = "Rubic";
for ( int i = 0; i < IndexedNames.size; i++ )
{
Console.WriteLine(names[i]);
}
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Zara
Riz
Nuha
Asif
Davinder
Sunil
Rubic
N. A.
N. A.
N. A.

```

### 重载索引器（Indexer）

索引器（Indexer）可被重载。索引器声明的时候也可带有多个参数，且每个参数可以是不同的类型。没有必要让索引器必须是整型的。C# 允许索引器可以是其他类型，例如，字符串类型。

下面的实例演示了重载索引器：

### 实例

using System;
namespace IndexerApplication
{
class IndexedNames
{
private string[] namelist = new string[size];
static public int size = 10;
public IndexedNames()
{
for (int i = 0; i < size; i++)
{
namelist[i] = "N. A.";
}
}
public string this[int index]
{
get
{
string tmp;

if( index >= 0 && index <= size-1 )
{
tmp = namelist[index];
}
else
{
tmp = "";
}

return ( tmp );
}
set
{
if( index >= 0 && index <= size-1 )
{
namelist[index] = value;
}
}
}
public int this[string name]
{
get
{
int index = 0;
while(index < size)
{
if (namelist[index] == name)
{
return index;
}
index++;
}
return index;
}

}

static void Main(string[] args)
{
IndexedNames names = new IndexedNames();
names[0] = "Zara";
names[1] = "Riz";
names[2] = "Nuha";
names[3] = "Asif";
names[4] = "Davinder";
names[5] = "Sunil";
names[6] = "Rubic";
// 使用带有 int 参数的第一个索引器
for (int i = 0; i < IndexedNames.size; i++)
{
Console.WriteLine(names[i]);
}
// 使用带有 string 参数的第二个索引器
Console.WriteLine(names["Nuha"]);
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Zara
Riz
Nuha
Asif
Davinder
Sunil
Rubic
N. A.
N. A.
N. A.
2

```

---

## C# 委托（Delegate）

Source: https://www.runoob.com/csharp/csharp-delegate.html

## C# 委托（Delegate）

在 C# 中，委托（Delegate） 是一种类型安全的函数指针，它允许将方法作为参数传递给其他方法。

C# 中的委托（Delegate）类似于 C 或 C++ 中函数的指针。委托（Delegate） 是存有对某个方法的引用的一种引用类型变量，引用可在运行时被改变。

委托在 C# 中非常常见，用于事件处理、回调函数、LINQ 等操作。

所有的委托（Delegate）都派生自 System.Delegate 类。

### 声明委托（Delegate）

委托是一个引用类型，它定义了一个方法签名，可以用于存储指向该签名的方法。通过委托，你可以调用其他类中的方法。

委托声明决定了可由该委托引用的方法。委托可指向一个与其具有相同标签的方法。

声明委托的语法如下：

```

public delegate <return type> <delegate-name> <parameter list>

```

中文格式说明：

```
public delegate 返回类型 委托名(参数类型 参数名, ...);
```

例如以下代码，我们定义一个接受两个整数并返回一个整数的委托：

```

public delegate int MathOperation(int x, int y);
```

以下例子的委托可被用于引用任何一个带有一个单一的 string 参数的方法，并返回一个 int 类型变量。

```

public delegate int MyDelegate (string s);

```

### 实例化委托（Delegate）

一旦声明了委托类型，委托对象必须使用 new 关键字来创建，且与一个特定的方法有关。当创建委托时，传递到 new 语句的参数就像方法调用一样书写，但是不带有参数。例如：

public delegate void printString(string s);
...
printString ps1 = new printString(WriteToScreen);
printString ps2 = new printString(WriteToFile);

下面的实例演示了委托的声明、实例化和使用，该委托可用于引用带有一个整型参数的方法，并返回一个整型值。

### 实例

using System;

delegate int NumberChanger(int n);
namespace DelegateAppl
{
class TestDelegate
{
static int num = 10;
public static int AddNum(int p)
{
num += p;
return num;
}

public static int MultNum(int q)
{
num *= q;
return num;
}
public static int getNum()
{
return num;
}

static void Main(string[] args)
{
// 创建委托实例
NumberChanger nc1 = new NumberChanger(AddNum);
NumberChanger nc2 = new NumberChanger(MultNum);
// 使用委托对象调用方法
nc1(25);
Console.WriteLine("Value of Num: {0}", getNum());
nc2(5);
Console.WriteLine("Value of Num: {0}", getNum());
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Value of Num: 35
Value of Num: 175

```

### 委托的多播（Multicasting of a Delegate）

委托对象可使用 + 运算符进行合并。

一个合并委托调用它所合并的两个委托，只有相同类型的委托可被合并。

- 运算符可用于从合并的委托中移除组件委托。

使用委托的这个有用的特点，您可以创建一个委托被调用时要调用的方法的调用列表，这被称为委托的 多播（multicasting），也叫组播。

下面的程序演示了委托的多播：

### 实例

using System;

delegate int NumberChanger(int n);
namespace DelegateAppl
{
class TestDelegate
{
static int num = 10;
public static int AddNum(int p)
{
num += p;
return num;
}

public static int MultNum(int q)
{
num *= q;
return num;
}
public static int getNum()
{
return num;
}

static void Main(string[] args)
{
// 创建委托实例
NumberChanger nc;
NumberChanger nc1 = new NumberChanger(AddNum);
NumberChanger nc2 = new NumberChanger(MultNum);
nc = nc1;
nc += nc2;
// 调用多播
nc(5);
Console.WriteLine("Value of Num: {0}", getNum());
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Value of Num: 75

```

### 委托（Delegate）的用途

下面的实例演示了委托的用法。委托 printString 可用于引用带有一个字符串作为输入的方法，并不返回任何东西。

我们使用这个委托来调用两个方法，第一个把字符串打印到控制台，第二个把字符串打印到文件：

### 实例

using System;
using System.IO;

namespace DelegateAppl
{
class PrintString
{
static FileStream fs;
static StreamWriter sw;
// 委托声明
public delegate void printString(string s);

// 该方法打印到控制台
public static void WriteToScreen(string str)
{
Console.WriteLine("The String is: {0}", str);
}
// 该方法打印到文件
public static void WriteToFile(string s)
{
fs = new FileStream("c:\\message.txt", FileMode.Append, FileAccess.Write);
sw = new StreamWriter(fs);
sw.WriteLine(s);
sw.Flush();
sw.Close();
fs.Close();
}
// 该方法把委托作为参数，并使用它调用方法
public static void sendString(printString ps)
{
ps("Hello World");
}
static void Main(string[] args)
{
printString ps1 = new printString(WriteToScreen);
printString ps2 = new printString(WriteToFile);
sendString(ps1);
sendString(ps2);
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

The String is: Hello World

```

### 移除委托

如果你不再需要某个方法，可以通过 -= 运算符将该方法从委托链中移除。

### 实例

public class Program
{
public delegate void PrintMessage(string message);

public static void PrintUpperCase(string message)
{
Console.WriteLine(message.ToUpper());
}

public static void PrintLowerCase(string message)
{
Console.WriteLine(message.ToLower());
}

public static void Main()
{
PrintMessage print = PrintUpperCase;
print += PrintLowerCase;

// 移除 PrintLowerCase 方法
print -= PrintLowerCase;

// 调用委托（只会调用 PrintUpperCase）
print("Hello, C#"); // 输出: HELLO, C#
}
}

### 委托和事件

委托常常与事件（Event）一起使用，事件是一种特殊类型的委托，用于发布和订阅机制。

在 C# 中，事件本质上就是一个封装了委托的类型，它用于响应程序中的某些操作。

### 实例

public class Button
{
// 定义一个事件
public event EventHandler Click;

// 引发事件的方法
public void OnClick()
{
if (Click != null)
{
Click(this, EventArgs.Empty); // 调用事件
}
}
}

public class Program
{
public static void Main()
{
Button button = new Button();

// 订阅事件
button.Click += Button_Click;

// 引发事件
button.OnClick(); // 输出 "Button clicked!"
}

private static void Button_Click(object sender, EventArgs e)
{
Console.WriteLine("Button clicked!");
}
}

### 委托的类型

C# 提供了几种常见的委托类型：

#### 1、Action

Action：代表不返回值的方法。可以接受最多 16 个参数。

```

Action<string> printMessage = Console.WriteLine;
printMessage("Hello");
```

#### 2、Func

Func：代表有返回值的方法。最多接受 16 个参数，第一个参数是输入参数，最后一个参数是返回值类型。

```
Func<int, int, int> add = (x, y) => x + y;
Console.WriteLine(add(3, 4)); // 输出 7
```

#### 3、Predicate

Predicate：代表返回 bool 值的方法，通常用于条件判断。

```

Predicate<int> isEven = x => x % 2 == 0;
Console.WriteLine(isEven(4)); // 输出 True
```

委托的注意事项

类型安全：委托是类型安全的，这意味着只有签名匹配的方法才能赋值给委托。

匿名方法和 lambda 表达式：你可以使用匿名方法或 lambda 表达式来创建委托实例，简化代码。

```

Func<int, int, int> add = (x, y) => x + y;
Console.WriteLine(add(5, 3)); // 输出 8
```

异步调用：可以将委托与 BeginInvoke 和 EndInvoke 方法一起使用，进行异步调用。

委托是 C# 中一个非常强大和灵活的特性，可以帮助实现事件驱动的编程、回调机制和函数式编程风格。它不仅提供了代码重用的能力，还提高了程序的模块化程度。理解和掌握委托的使用对于 C# 编程是非常重要的。

---

## C# 事件（Event）

Source: https://www.runoob.com/csharp/csharp-event.html

## C# 事件（Event）

C# 事件（Event）是一种成员，用于将特定的事件通知发送给订阅者。事件通常用于实现观察者模式，它允许一个对象将状态的变化通知其他对象，而不需要知道这些对象的细节。

事件（Event） 基本上说是一个用户操作，如按键、点击、鼠标移动等等，或者是一些提示信息，如系统生成的通知。应用程序需要在事件发生时响应事件。例如，中断。

C# 中使用事件机制实现线程间的通信。

关键点：

- 声明委托：定义事件将使用的委托类型。委托是一个函数签名。
- 声明事件：使用 `event` 关键字声明一个事件。
- 触发事件：在适当的时候调用事件，通知所有订阅者。
- 订阅和取消订阅事件：其他类可以通过 `+=` 和 `-=` 运算符订阅和取消订阅事件。

### 通过事件使用委托

事件在类中声明且生成，且通过使用同一个类或其他类中的委托与事件处理程序关联。包含事件的类用于发布事件。这被称为 发布器（publisher） 类。其他接受该事件的类被称为 订阅器（subscriber） 类。事件使用 发布-订阅（publisher-subscriber） 模型。

发布器（publisher） 是一个包含事件和委托定义的对象。事件和委托之间的联系也定义在这个对象中。发布器（publisher）类的对象调用这个事件，并通知其他的对象。

订阅器（subscriber） 是一个接受事件并提供事件处理程序的对象。在发布器（publisher）类中的委托调用订阅器（subscriber）类中的方法（事件处理程序）。

### 声明事件（Event）

在类的内部声明事件，首先必须声明该事件的委托类型。例如：

```

public delegate void BoilerLogHandler(string status);

```

然后，声明事件本身，使用 event 关键字：

```

// 基于上面的委托定义事件
public event BoilerLogHandler BoilerEventLog;

```

上面的代码定义了一个名为 BoilerLogHandler 的委托和一个名为 BoilerEventLog 的事件，该事件在生成的时候会调用委托。

以下示例展示了如何在 C# 中使用事件：

### 实例

using System;

namespace EventDemo
{
// 定义一个委托类型，用于事件处理程序
public delegate void NotifyEventHandler(object sender, EventArgs e);

// 发布者类
public class ProcessBusinessLogic
{
// 声明事件
public event NotifyEventHandler ProcessCompleted;

// 触发事件的方法
protected virtual void OnProcessCompleted(EventArgs e)
{
ProcessCompleted?.Invoke(this, e);
}

// 模拟业务逻辑过程并触发事件
public void StartProcess()
{
Console.WriteLine("Process Started!");

// 这里可以加入实际的业务逻辑

// 业务逻辑完成，触发事件
OnProcessCompleted(EventArgs.Empty);
}
}

// 订阅者类
public class EventSubscriber
{
public void Subscribe(ProcessBusinessLogic process)
{
process.ProcessCompleted += Process_ProcessCompleted;
}

private void Process_ProcessCompleted(object sender, EventArgs e)
{
Console.WriteLine("Process Completed!");
}
}

class Program
{
static void Main(string[] args)
{
ProcessBusinessLogic process = new ProcessBusinessLogic();
EventSubscriber subscriber = new EventSubscriber();

// 订阅事件
subscriber.Subscribe(process);

// 启动过程
process.StartProcess();

Console.ReadLine();
}
}
}

#### 说明

1、定义委托类型：

```
public delegate void NotifyEventHandler(object sender, EventArgs e);
```

这是一个委托类型，它定义了事件处理程序的签名。通常使用 `EventHandler` 或 `EventHandler<TEventArgs>` 来替代自定义的委托。

2、声明事件：

```
public event NotifyEventHandler ProcessCompleted;
```

这是一个使用 NotifyEventHandler 委托类型的事件。

3、触发事件：

```
protected virtual void OnProcessCompleted(EventArgs e)
{
ProcessCompleted?.Invoke(this, e);
}
```

这是一个受保护的方法，用于触发事件。使用 `?.Invoke` 语法来确保只有在有订阅者时才调用事件。

4、订阅和取消订阅事件：

```
process.ProcessCompleted += Process_ProcessCompleted;
```

订阅者使用 `+=` 运算符订阅事件，并定义事件处理程序 `Process_ProcessCompleted`。

### 实例

### 实例 1

using System;
namespace SimpleEvent
{
using System;
/***********发布器类***********/
public class EventTest
{
private int value;

public delegate void NumManipulationHandler();

public event NumManipulationHandler ChangeNum;
protected virtual void OnNumChanged()
{
if ( ChangeNum != null )
{
ChangeNum(); /* 事件被触发 */
}else {
Console.WriteLine( "event not fire" );
Console.ReadKey(); /* 回车继续 */
}
}

public EventTest()
{
int n = 5;
SetValue( n );
}

public void SetValue( int n )
{
if ( value != n )
{
value = n;
OnNumChanged();
}
}
}

/***********订阅器类***********/

public class subscribEvent
{
public void printf()
{
Console.WriteLine( "event fire" );
Console.ReadKey(); /* 回车继续 */
}
}

/***********触发***********/
public class MainClass
{
public static void Main()
{
EventTest e = new EventTest(); /* 实例化对象,第一次没有触发事件 */
subscribEvent v = new subscribEvent(); /* 实例化对象 */
e.ChangeNum += new EventTest.NumManipulationHandler( v.printf ); /* 注册 */
e.SetValue( 7 );
e.SetValue( 11 );
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

event not fire
event fire
event fire


```

本实例提供一个简单的用于热水锅炉系统故障排除的应用程序。当维修工程师检查锅炉时，锅炉的温度和压力会随着维修工程师的备注自动记录到日志文件中。

### 实例 2

using System;
using System.IO;

namespace BoilerEventAppl
{
// Boiler 类
class Boiler
{
public int Temp { get; private set; }
public int Pressure { get; private set; }

public Boiler(int temp, int pressure)
{
Temp = temp;
Pressure = pressure;
}
}

// 事件发布器
class DelegateBoilerEvent
{
public delegate void BoilerLogHandler(string status);

// 基于上面的委托定义事件
public event BoilerLogHandler BoilerEventLog;

public void LogProcess()
{
string remarks = "O.K.";
Boiler boiler = new Boiler(100, 12);
int temp = boiler.Temp;
int pressure = boiler.Pressure;

if (temp > 150 || temp < 80 || pressure < 12 || pressure > 15)
{
remarks = "Need Maintenance";
}

OnBoilerEventLog($"Logging Info:\nTemperature: {temp}\nPressure: {pressure}\nMessage: {remarks}");
}

protected void OnBoilerEventLog(string message)
{
BoilerEventLog?.Invoke(message);
}
}

// 该类保留写入日志文件的条款
class BoilerInfoLogger : IDisposable
{
private readonly StreamWriter _streamWriter;

public BoilerInfoLogger(string filename)
{
_streamWriter = new StreamWriter(new FileStream(filename, FileMode.Append, FileAccess.Write));
}

public void Logger(string info)
{
_streamWriter.WriteLine(info);
}

public void Dispose()
{
_streamWriter?.Close();
}
}

// 事件订阅器
public class RecordBoilerInfo
{
static void Logger(string info)
{
Console.WriteLine(info);
}

static void Main(string[] args)
{
using (BoilerInfoLogger fileLogger = new BoilerInfoLogger("e:\\boiler.txt"))
{
DelegateBoilerEvent boilerEvent = new DelegateBoilerEvent();
boilerEvent.BoilerEventLog += Logger;
boilerEvent.BoilerEventLog += fileLogger.Logger;
boilerEvent.LogProcess();
}

Console.ReadLine();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Logging info:

Temperature 100
Pressure 12

Message: O. K

```

---

## C# 集合（Collection）

Source: https://www.runoob.com/csharp/csharp-collection.html

## C# 集合（Collection）

在 C# 中，集合（Collection）类是专门用于数据存储和检索的一类类。

集合让我们能够轻松地存放、管理和操作一组数据，比如：

- 通过索引访问列表中的项；
- 通过键访问字典中的值；
- 动态地添加或删除元素等。

集合类提供了对多种常见数据结构的支持，例如：

- 栈（Stack） —— 后进先出（LIFO）结构；
- 队列（Queue） —— 先进先出（FIFO）结构；
- 列表（List） —— 可动态扩展的有序序列；
- 哈希表（HashTable） —— 通过键值对快速查找数据。

此外，大多数集合类都实现了相同的接口，
这意味着它们具有类似的操作方式，例如添加、删除、遍历等方法。

在早期的 C# 版本中，集合中的元素通常以 `Object` 类型存储，我们可以往集合中放入任何类型的对象。这是因为在 C# 中，`Object` 是所有数据类型的基类——无论是 `int`、`string` 还是自定义类，最终都继承自 `Object`。

### 各种集合类和它们的用法

下面是各种常用的 System.Collection 命名空间的类。点击下面的链接查看细节。

类描述和用法 动态数组（ArrayList）它代表了可被单独索引的对象的有序集合。

它基本上可以替代一个数组。但是，与数组不同的是，您可以使用索引在指定的位置添加和移除项目，动态数组会自动重新调整它的大小。它也允许在列表中进行动态内存分配、增加、搜索、排序各项。 哈希表（Hashtable）它使用键来访问集合中的元素。

当您使用键访问元素时，则使用哈希表，而且您可以识别一个有用的键值。哈希表中的每一项都有一个键/值对。键用于访问集合中的项目。 排序列表（SortedList）它可以使用键和索引来访问列表中的项。

排序列表是数组和哈希表的组合。它包含一个可使用键或索引访问各项的列表。如果您使用索引访问各项，则它是一个动态数组（ArrayList），如果您使用键访问各项，则它是一个哈希表（Hashtable）。集合中的各项总是按键值排序。 堆栈（Stack）它代表了一个后进先出的对象集合。

当您需要对各项进行后进先出的访问时，则使用堆栈。当您在列表中添加一项，称为推入元素，当您从列表中移除一项时，称为弹出元素。 队列（Queue）它代表了一个先进先出的对象集合。

当您需要对各项进行先进先出的访问时，则使用队列。当您在列表中添加一项，称为入队，当您从列表中移除一项时，称为出队。 点阵列（BitArray）它代表了一个使用值 1 和 0 来表示的二进制数组。

当您需要存储位，但是事先不知道位数时，则使用点阵列。您可以使用整型索引从点阵列集合中访问各项，索引从零开始。

---

## C# 泛型（Generic）

Source: https://www.runoob.com/csharp/csharp-generic.html

## C# 泛型（Generic）

泛型（Generic） 允许您延迟编写类或方法中的编程元素的数据类型的规范，直到实际在程序中使用它的时候。换句话说，泛型允许您编写一个可以与任何数据类型一起工作的类或方法。

您可以通过数据类型的替代参数编写类或方法的规范。当编译器遇到类的构造函数或方法的函数调用时，它会生成代码来处理指定的数据类型。下面这个简单的实例将有助于您理解这个概念：

### 实例

using System;
using System.Collections.Generic;

namespace GenericApplication
{
public class MyGenericArray<T>
{
private T[] array;
public MyGenericArray(int size)
{
array = new T[size + 1];
}
public T getItem(int index)
{
return array[index];
}
public void setItem(int index, T value)
{
array[index] = value;
}
}

class Tester
{
static void Main(string[] args)
{
// 声明一个整型数组
MyGenericArray<int> intArray = new MyGenericArray<int>(5);
// 设置值
for (int c = 0; c < 5; c++)
{
intArray.setItem(c, c*5);
}
// 获取值
for (int c = 0; c < 5; c++)
{
Console.Write(intArray.getItem(c) + " ");
}
Console.WriteLine();
// 声明一个字符数组
MyGenericArray<char> charArray = new MyGenericArray<char>(5);
// 设置值
for (int c = 0; c < 5; c++)
{
charArray.setItem(c, (char)(c+97));
}
// 获取值
for (int c = 0; c < 5; c++)
{
Console.Write(charArray.getItem(c) + " ");
}
Console.WriteLine();
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

0 5 10 15 20
a b c d e

```

### 泛型（Generic）的特性

使用泛型是一种增强程序功能的技术，具体表现在以下几个方面：

- 它有助于您最大限度地重用代码、保护类型的安全以及提高性能。
- 您可以创建泛型集合类。.NET 框架类库在 System.Collections.Generic 命名空间中包含了一些新的泛型集合类。您可以使用这些泛型集合类来替代 System.Collections 中的集合类。
- 您可以创建自己的泛型接口、泛型类、泛型方法、泛型事件和泛型委托。
- 您可以对泛型类进行约束以访问特定数据类型的方法。
- 关于泛型数据类型中使用的类型的信息可在运行时通过使用反射获取。

### 泛型（Generic）方法

在上面的实例中，我们已经使用了泛型类，我们可以通过类型参数声明泛型方法。下面的程序说明了这个概念：

### 实例

using System;
using System.Collections.Generic;

namespace GenericMethodAppl
{
class Program
{
static void Swap<T>(ref T lhs, ref T rhs)
{
T temp;
temp = lhs;
lhs = rhs;
rhs = temp;
}
static void Main(string[] args)
{
int a, b;
char c, d;
a = 10;
b = 20;
c = 'I';
d = 'V';

// 在交换之前显示值
Console.WriteLine("Int values before calling swap:");
Console.WriteLine("a = {0}, b = {1}", a, b);
Console.WriteLine("Char values before calling swap:");
Console.WriteLine("c = {0}, d = {1}", c, d);

// 调用 swap
Swap<int>(ref a, ref b);
Swap<char>(ref c, ref d);

// 在交换之后显示值
Console.WriteLine("Int values after calling swap:");
Console.WriteLine("a = {0}, b = {1}", a, b);
Console.WriteLine("Char values after calling swap:");
Console.WriteLine("c = {0}, d = {1}", c, d);
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Int values before calling swap:
a = 10, b = 20
Char values before calling swap:
c = I, d = V
Int values after calling swap:
a = 20, b = 10
Char values after calling swap:
c = V, d = I

```

### 泛型（Generic）委托

您可以通过类型参数定义泛型委托。例如：

```

delegate T NumberChanger<T>(T n);

```

下面的实例演示了委托的使用：

### 实例

using System;
using System.Collections.Generic;

delegate T NumberChanger<T>(T n);
namespace GenericDelegateAppl
{
class TestDelegate
{
static int num = 10;
public static int AddNum(int p)
{
num += p;
return num;
}

public static int MultNum(int q)
{
num *= q;
return num;
}
public static int getNum()
{
return num;
}

static void Main(string[] args)
{
// 创建委托实例
NumberChanger<int> nc1 = new NumberChanger<int>(AddNum);
NumberChanger<int> nc2 = new NumberChanger<int>(MultNum);
// 使用委托对象调用方法
nc1(25);
Console.WriteLine("Value of Num: {0}", getNum());
nc2(5);
Console.WriteLine("Value of Num: {0}", getNum());
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Value of Num: 35
Value of Num: 175

```

---

## C# 匿名方法

Source: https://www.runoob.com/csharp/csharp-anonymous-methods.html

## C# 匿名方法

在 C# 中，匿名函数是一种没有名字的方法，可以在代码中定义和使用。

我们已经提到过，委托是用于引用与其具有相同标签的方法。换句话说，您可以使用委托对象调用可由委托引用的方法。

匿名方法（Anonymous methods） 提供了一种传递代码块作为委托参数的技术。

在匿名方法中您不需要指定返回类型，它是从方法主体内的 return 语句推断的。

### Lambda 表达式

Lambda 表达式是一个简洁的语法，用于创建匿名函数。它们通常用于 LINQ 查询和委托。

#### 语法

```
(parameters) => expression
// 或
(parameters) => { statement; }
```

### 实例

// 示例：使用 Lambda 表达式定义一个委托
Func<int, int, int> add = (a, b) => a + b;
Console.WriteLine(add(2, 3)); // 输出 5

// 示例：使用 Lambda 表达式过滤数组中的元素
int[] numbers = { 1, 2, 3, 4, 5 };
var evenNumbers = numbers.Where(n => n % 2 == 0);
foreach (var num in evenNumbers)
{
Console.WriteLine(num); // 输出 2 4
}

### 匿名方法

匿名方法是通过使用 delegate 关键字创建委托实例来声明的。

#### 语法

```

delegate(parameters) { statement; }
```

例如：

```

delegate void NumberChanger(int n);
...
NumberChanger nc = delegate(int x)
{
Console.WriteLine("Anonymous Method: {0}", x);
};


```

代码块 Console.WriteLine("Anonymous Method: {0}", x); 是匿名方法的主体。

委托可以通过匿名方法调用，也可以通过命名方法调用，即，通过向委托对象传递方法参数。

注意: 匿名方法的主体后面需要一个 ;。

例如：

```

nc(10);

```

### 实例

// 示例：使用匿名方法定义一个委托
Func<int, int, int> multiply = delegate (int a, int b)
{
return a * b;
};
Console.WriteLine(multiply(2, 3)); // 输出 6

// 示例：使用匿名方法作为事件处理程序
Button button = new Button();
button.Click += delegate (object sender, EventArgs e)
{
Console.WriteLine("Button clicked!");
};

### 实例

下面的实例演示了匿名方法的概念：

### 实例

using System;

delegate void NumberChanger(int n);
namespace DelegateAppl
{
class TestDelegate
{
static int num = 10;
public static void AddNum(int p)
{
num += p;
Console.WriteLine("Named Method: {0}", num);
}

public static void MultNum(int q)
{
num *= q;
Console.WriteLine("Named Method: {0}", num);
}

static void Main(string[] args)
{
// 使用匿名方法创建委托实例
NumberChanger nc = delegate(int x)
{
Console.WriteLine("Anonymous Method: {0}", x);
};

// 使用匿名方法调用委托
nc(10);

// 使用命名方法实例化委托
nc = new NumberChanger(AddNum);

// 使用命名方法调用委托
nc(5);

// 使用另一个命名方法实例化委托
nc = new NumberChanger(MultNum);

// 使用命名方法调用委托
nc(2);
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Anonymous Method: 10
Named Method: 15
Named Method: 30

```

在 C# 2.0 及更高版本中，引入了 lambda 表达式，它是一种更简洁的语法形式，用于编写匿名方法。

使用 lambda 表达式：

### 实例

using System;

delegate void NumberChanger(int n);

namespace DelegateAppl
{
class TestDelegate
{
static int num = 10;

public static void AddNum(int p)
{
num += p;
Console.WriteLine("Named Method: {0}", num);
}

public static void MultNum(int q)
{
num *= q;
Console.WriteLine("Named Method: {0}", num);
}

static void Main(string[] args)
{
// 使用 lambda 表达式创建委托实例
NumberChanger nc = x => Console.WriteLine($"Lambda Expression: {x}");

// 使用 lambda 表达式调用委托
nc(10);

// 使用命名方法实例化委托
nc = new NumberChanger(AddNum);

// 使用命名方法调用委托
nc(5);

// 使用另一个命名方法实例化委托
nc = new NumberChanger(MultNum);

// 使用命名方法调用委托
nc(2);

Console.ReadKey();
}
}
}

---

## C# 不安全代码

Source: https://www.runoob.com/csharp/csharp-unsafe-codes.html

## C# 不安全代码

当一个代码块使用 unsafe 修饰符标记时，C# 允许在函数中使用指针变量。不安全代码或非托管代码是指使用了指针变量的代码块。

### 指针变量

指针 是值为另一个变量的地址的变量，即，内存位置的直接地址。就像其他变量或常量，您必须在使用指针存储其他变量地址之前声明指针。

指针变量声明的一般形式为：

```
type* var-name;
```

下面是指针类型声明的实例：

实例 描述 `int* p` `p` 是指向整数的指针。 `double* p` `p` 是指向双精度数的指针。 `float* p` `p` 是指向浮点数的指针。 `int** p` `p` 是指向整数的指针的指针。 `int*[] p` `p` 是指向整数的指针的一维数组。 `char* p` `p` 是指向字符的指针。 `void* p` `p` 是指向未知类型的指针。

在同一个声明中声明多个指针时，星号 * 仅与基础类型一起写入；而不是用作每个指针名称的前缀。 例如:

```
int* p1, p2, p3; // 正确
int *p1, *p2, *p3; // 错误
```

下面的实例说明了 C# 中使用了 unsafe 修饰符时指针的使用：

### 实例

using System;
namespace UnsafeCodeApplication
{
class Program
{
static unsafe void Main(string[] args)
{
int var = 20;
int* p = &var;
Console.WriteLine("Data is: {0} ", var);
Console.WriteLine("Address is: {0}", (int)p);
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Data is: 20
Address is: 99215364

```

您也可以不用声明整个方法作为不安全代码，只需要声明方法的一部分作为不安全代码。下面的实例说明了这点。

### 使用指针检索数据值

您可以使用 ToString() 方法检索存储在指针变量所引用位置的数据。下面的实例演示了这点：

### 实例

using System;
namespace UnsafeCodeApplication
{
class Program
{
public static void Main()
{
unsafe
{
int var = 20;
int* p = &var;
Console.WriteLine("Data is: {0} " , var);
Console.WriteLine("Data is: {0} " , p->ToString());
Console.WriteLine("Address is: {0} " , (int)p);
}
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Data is: 20
Data is: 20
Address is: 77128984

```

### 传递指针作为方法的参数

您可以向方法传递指针变量作为方法的参数。下面的实例说明了这点：

### 实例

using System;
namespace UnsafeCodeApplication
{
class TestPointer
{
public unsafe void swap(int* p, int *q)
{
int temp = *p;
*p = *q;
*q = temp;
}

public unsafe static void Main()
{
TestPointer p = new TestPointer();
int var1 = 10;
int var2 = 20;
int* x = &var1;
int* y = &var2;

Console.WriteLine("Before Swap: var1:{0}, var2: {1}", var1, var2);
p.swap(x, y);

Console.WriteLine("After Swap: var1:{0}, var2: {1}", var1, var2);
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Before Swap: var1: 10, var2: 20
After Swap: var1: 20, var2: 10

```

### 使用指针访问数组元素

在 C# 中，数组名称和一个指向与数组数据具有相同数据类型的指针是不同的变量类型。例如，int *p 和 int[] p 是不同的类型。您可以增加指针变量 p，因为它在内存中不是固定的，但是数组地址在内存中是固定的，所以您不能增加数组 p。

因此，如果您需要使用指针变量访问数组数据，可以像我们通常在 C 或 C++ 中所做的那样，使用 fixed 关键字来固定指针。

下面的实例演示了这点：

### 实例

using System;
namespace UnsafeCodeApplication
{
class TestPointer
{
public unsafe static void Main()
{
int[] list = {10, 100, 200};
fixed(int *ptr = list)

/* 显示指针中数组地址 */
for ( int i = 0; i < 3; i++)
{
Console.WriteLine("Address of list[{0}]={1}",i,(int)(ptr + i));
Console.WriteLine("Value of list[{0}]={1}", i, *(ptr + i));
}
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

Address of list[0] = 31627168
Value of list[0] = 10
Address of list[1] = 31627172
Value of list[1] = 100
Address of list[2] = 31627176
Value of list[2] = 200

```

### 编译不安全代码

为了编译不安全代码，您必须切换到命令行编译器指定 /unsafe 命令行。

例如，为了编译包含不安全代码的名为 prog1.cs 的程序，需在命令行中输入命令：

```

csc /unsafe prog1.cs

```

如果您使用的是 Visual Studio IDE，那么您需要在项目属性中启用不安全代码。

步骤如下：

- 通过双击资源管理器（Solution Explorer）中的属性（properties）节点，打开项目属性（project properties）。
- 点击 Build 标签页。
- 选择选项"Allow unsafe code"。

---

## C# 多线程

Source: https://www.runoob.com/csharp/csharp-multithreading.html

## C# 多线程

线程 被定义为程序的执行路径。每个线程都定义了一个独特的控制流。如果您的应用程序涉及到复杂的和耗时的操作，那么设置不同的线程执行路径往往是有益的，每个线程执行特定的工作。

线程是轻量级进程。一个使用线程的常见实例是现代操作系统中并行编程的实现。使用线程节省了 CPU 周期的浪费，同时提高了应用程序的效率。

到目前为止我们编写的程序是一个单线程作为应用程序的运行实例的单一的过程运行的。但是，这样子应用程序同时只能执行一个任务。为了同时执行多个任务，它可以被划分为更小的线程。

### 线程生命周期

线程生命周期开始于 System.Threading.Thread 类的对象被创建时，结束于线程被终止或完成执行时。

下面列出了线程生命周期中的各种状态：

- 未启动状态：当线程实例被创建但 Start 方法未被调用时的状况。
- 就绪状态：当线程准备好运行并等待 CPU 周期时的状况。
- 不可运行状态：下面的几种情况下线程是不可运行的：

- 已经调用 Sleep 方法
- 已经调用 Wait 方法
- 通过 I/O 操作阻塞
- 死亡状态：当线程已完成执行或已中止时的状况。

### 主线程

在 C# 中，System.Threading.Thread 类用于线程的工作。它允许创建并访问多线程应用程序中的单个线程。进程中第一个被执行的线程称为主线程。

当 C# 程序开始执行时，主线程自动创建。使用 Thread 类创建的线程被主线程的子线程调用。您可以使用 Thread 类的 CurrentThread 属性访问线程。

下面的程序演示了主线程的执行：

### 实例

using System;
using System.Threading;

namespace MultithreadingApplication
{
class MainThreadProgram
{
static void Main(string[] args)
{
Thread th = Thread.CurrentThread;
th.Name = "MainThread";
Console.WriteLine("This is {0}", th.Name);
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

This is MainThread

```

### Thread 类常用的属性和方法

下表列出了 Thread 类的一些常用的 属性：

属性描述 CurrentContext获取线程正在其中执行的当前上下文。 CurrentCulture获取或设置当前线程的区域性。 CurrentPrincipal获取或设置线程的当前负责人（对基于角色的安全性而言）。 CurrentThread获取当前正在运行的线程。 CurrentUICulture获取或设置资源管理器使用的当前区域性以便在运行时查找区域性特定的资源。 ExecutionContext获取一个 ExecutionContext 对象，该对象包含有关当前线程的各种上下文的信息。 IsAlive获取一个值，该值指示当前线程的执行状态。 IsBackground获取或设置一个值，该值指示某个线程是否为后台线程。 IsThreadPoolThread获取一个值，该值指示线程是否属于托管线程池。 ManagedThreadId获取当前托管线程的唯一标识符。 Name获取或设置线程的名称。 Priority获取或设置一个值，该值指示线程的调度优先级。 ThreadState获取一个值，该值包含当前线程的状态。

下表列出了 Thread 类的一些常用的 方法：

序号方法名 & 描述 1public void Abort()
在调用此方法的线程上引发 ThreadAbortException，以开始终止此线程的过程。调用此方法通常会终止线程。 2public static LocalDataStoreSlot AllocateDataSlot()
在所有的线程上分配未命名的数据槽。为了获得更好的性能，请改用以 ThreadStaticAttribute 属性标记的字段。 3public static LocalDataStoreSlot AllocateNamedDataSlot( string name)
在所有线程上分配已命名的数据槽。为了获得更好的性能，请改用以 ThreadStaticAttribute 属性标记的字段。 4public static void BeginCriticalRegion()
通知主机执行将要进入一个代码区域，在该代码区域内线程中止或未经处理的异常的影响可能会危害应用程序域中的其他任务。 5public static void BeginThreadAffinity()
通知主机托管代码将要执行依赖于当前物理操作系统线程的标识的指令。 6public static void EndCriticalRegion()
通知主机执行将要进入一个代码区域，在该代码区域内线程中止或未经处理的异常仅影响当前任务。 7public static void EndThreadAffinity()
通知主机托管代码已执行完依赖于当前物理操作系统线程的标识的指令。 8public static void FreeNamedDataSlot(string name)
为进程中的所有线程消除名称与槽之间的关联。为了获得更好的性能，请改用以 ThreadStaticAttribute 属性标记的字段。 9public static Object GetData( LocalDataStoreSlot slot )
在当前线程的当前域中从当前线程上指定的槽中检索值。为了获得更好的性能，请改用以 ThreadStaticAttribute 属性标记的字段。 10public static AppDomain GetDomain()
返回当前线程正在其中运行的当前域。 11public static AppDomain GetDomainID()
返回唯一的应用程序域标识符。 12public static LocalDataStoreSlot GetNamedDataSlot( string name )
查找已命名的数据槽。为了获得更好的性能，请改用以 ThreadStaticAttribute 属性标记的字段。 13public void Interrupt()
中断处于 WaitSleepJoin 线程状态的线程。 14public void Join()
在继续执行标准的 COM 和 SendMessage 消息泵处理期间，阻塞调用线程，直到某个线程终止为止。此方法有不同的重载形式。 15public static void MemoryBarrier()
按如下方式同步内存存取：执行当前线程的处理器在对指令重新排序时，不能采用先执行 MemoryBarrier 调用之后的内存存取，再执行 MemoryBarrier 调用之前的内存存取的方式。 16public static void ResetAbort()
取消为当前线程请求的 Abort。 17public static void SetData( LocalDataStoreSlot slot, Object data )
在当前正在运行的线程上为此线程的当前域在指定槽中设置数据。为了获得更好的性能，请改用以 ThreadStaticAttribute 属性标记的字段。 18public void Start()
开始一个线程。 19public static void Sleep( int millisecondsTimeout )
让线程暂停一段时间。 20public static void SpinWait( int iterations )
导致线程等待由 iterations 参数定义的时间量。 21public static byte VolatileRead( ref byte address )
public static double VolatileRead( ref double address )
public static int VolatileRead( ref int address )
public static Object VolatileRead( ref Object address )
读取字段值。无论处理器的数目或处理器缓存的状态如何，该值都是由计算机的任何处理器写入的最新值。此方法有不同的重载形式。这里只给出了一些形式。 22public static void VolatileWrite( ref byte address, byte value )
public static void VolatileWrite( ref double address, double value )
public static void VolatileWrite( ref int address, int value )
public static void VolatileWrite( ref Object address, Object value )
立即向字段写入一个值，以使该值对计算机中的所有处理器都可见。此方法有不同的重载形式。这里只给出了一些形式。 23public static bool Yield()
导致调用线程执行准备好在当前处理器上运行的另一个线程。由操作系统选择要执行的线程。

### 创建线程

线程是通过扩展 Thread 类创建的。扩展的 Thread 类调用 Start() 方法来开始子线程的执行。

下面的程序演示了这个概念：

### 实例

using System;
using System.Threading;

namespace MultithreadingApplication
{
class ThreadCreationProgram
{
public static void CallToChildThread()
{
Console.WriteLine("Child thread starts");
}

static void Main(string[] args)
{
ThreadStart childref = new ThreadStart(CallToChildThread);
Console.WriteLine("In Main: Creating the Child thread");
Thread childThread = new Thread(childref);
childThread.Start();
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

In Main: Creating the Child thread
Child thread starts

```

### 管理线程

Thread 类提供了各种管理线程的方法。

下面的实例演示了 sleep() 方法的使用，用于在一个特定的时间暂停线程。

### 实例

using System;
using System.Threading;

namespace MultithreadingApplication
{
class ThreadCreationProgram
{
public static void CallToChildThread()
{
Console.WriteLine("Child thread starts");
// 线程暂停 5000 毫秒
int sleepfor = 5000;
Console.WriteLine("Child Thread Paused for {0} seconds",
sleepfor / 1000);
Thread.Sleep(sleepfor);
Console.WriteLine("Child thread resumes");
}

static void Main(string[] args)
{
ThreadStart childref = new ThreadStart(CallToChildThread);
Console.WriteLine("In Main: Creating the Child thread");
Thread childThread = new Thread(childref);
childThread.Start();
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

In Main: Creating the Child thread
Child thread starts
Child Thread Paused for 5 seconds
Child thread resumes

```

### 销毁线程

Abort() 方法用于销毁线程。

通过抛出 threadabortexception 在运行时中止线程。这个异常不能被捕获，如果有 finally 块，控制会被送至 finally 块。

下面的程序说明了这点：

### 实例

using System;
using System.Threading;

namespace MultithreadingApplication
{
class ThreadCreationProgram
{
public static void CallToChildThread()
{
try
{

Console.WriteLine("Child thread starts");
// 计数到 10
for (int counter = 0; counter <= 10; counter++)
{
Thread.Sleep(500);
Console.WriteLine(counter);
}
Console.WriteLine("Child Thread Completed");

}
catch (ThreadAbortException e)
{
Console.WriteLine("Thread Abort Exception");
}
finally
{
Console.WriteLine("Couldn't catch the Thread Exception");
}

}

static void Main(string[] args)
{
ThreadStart childref = new ThreadStart(CallToChildThread);
Console.WriteLine("In Main: Creating the Child thread");
Thread childThread = new Thread(childref);
childThread.Start();
// 停止主线程一段时间
Thread.Sleep(2000);
// 现在中止子线程
Console.WriteLine("In Main: Aborting the Child thread");
childThread.Abort();
Console.ReadKey();
}
}
}

当上面的代码被编译和执行时，它会产生下列结果：

```

In Main: Creating the Child thread
Child thread starts
0
1
2
In Main: Aborting the Child thread
Thread Abort Exception
Couldn't catch the Thread Exception

```

---

## C# 语言测验知识挑战

Source: https://www.runoob.com/csharp/csharp-quiz.html

## C# 语言测验

## 知识挑战

通过这些互动问题测试学习情况

### 测验完成！

0/0

重新开始

上一题 下一题

### 其他相关测试

1. C# 测验 1 2. C# 测验 2 3. C# 测验 3 4. C# 测验 4 5. C# 测验 5 6. C# 测验 6 7. C# 测验 7 8. C# 测验 8 9. C# 测验 9 10. C# 测验 10 11. C# 测验 11
