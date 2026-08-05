# Assembly - 菜鸟教程

Tutorial: https://www.runoob.com/assembly/assembly-tutorial.html

---

## 汇编语言 - 教程

Source: https://www.runoob.com/assembly/assembly-tutorial.html

## 汇编语言 - 教程

汇编语言（Assembly Language）是一种面向特定硬件的低级语言。

汇编语言用于电子计算机、微处理器或微控制器编程。

汇编语言与机器指令集一一对应，不可跨平台移植。

与 C、Python 等高级语言不同，汇编语言直接操作 寄存器、内存地址 和 CPU 指令，没有任何抽象层的遮挡。

### 为什么要学汇编

学习汇编语言能让你真正理解计算机的底层工作原理。

以下是学习汇编的几个核心价值：

学习目标说明 理解计算机底层掌握 CPU、内存、寄存器如何协同工作 提升调试能力能够阅读反汇编代码，定位底层 bug 性能优化理解编译器生成的代码，写出更高效的高级语言程序 安全研究逆向工程、漏洞分析、shellcode 编写的基础 嵌入式开发资源受限设备上直接控制硬件

### 学习本教程前需要了解

学习本教程前，建议具备以下基础：

- 了解基本的计算机操作（文件管理、命令行使用）
- 了解任意一门高级编程语言的基本概念（变量、循环、函数）
- 了解二进制、十六进制的基本概念（非必需，教程中会讲解）

### 教程使用的工具

工具用途版本建议 NASM汇编器，将汇编源码转为目标文件2.16 或更高 GCC 或 LD链接器，将目标文件链接为可执行文件任意版本 GDB调试器，单步执行和检查寄存器/内存任意版本

本教程所有示例代码均在 Linux 环境下使用 NASM 汇编器编写和测试。

如果你使用 Windows，可以安装 WSL 或使用虚拟机来搭建 Linux 环境。

### 学习建议

逐一阅读每个章节，动手敲打并运行每个代码示例。

汇编语言的学习重在实践，只看不写是无法掌握的。

遇到不懂的概念可以放慢节奏，配合网上搜索加深理解。

汇编语言学习曲线较陡，但一旦理解，你对计算机的认知将发生质变。坚持下去，你会看到不一样的风景。

---

## 汇编语言 - 简介

Source: https://www.runoob.com/assembly/assembly-intro.html

## 汇编语言 - 简介

汇编语言（Assembly Language）是计算机编程语言家族中最贴近硬件的一层，它是人类可读的机器指令助记符表示。

理解汇编语言的本质，是深入学习计算机系统的第一步。

### 什么是汇编语言

计算机的 CPU 只能理解和执行由 0 和 1 组成的二进制机器码。

例如，在 x86 架构中，机器码 `B8 2A 00 00 00` 表示将数值 42 移动到 EAX 寄存器。

直接编写二进制机器码对人类来说极其困难且容易出错，于是汇编语言应运而生。

汇编语言将这些二进制指令用人类可读的 助记符（Mnemonic） 来表示，例如上面的机器码在汇编中写作：

### 实例

mov eax, 42 ; 将数值 42 移动到 EAX 寄存器

汇编器（Assembler）负责将这些助记符翻译回 CPU 可以执行的机器码。

上面例子中，`mov` 就是助记符，`eax` 是目标操作数，`42` 是源操作数。

汇编语言用助记符替代机器语言操作，支持标签、符号指代地址与常量，避免硬编码，汇编语言与对应机器语言指令集基本一一对应。

### 汇编语言与机器码的关系

汇编指令和机器指令之间存在 一一对应 的关系。

每一条汇编指令都精确映射到一条 CPU 机器指令，不存在中间抽象层。

形式示例说明 机器码（二进制）10111000 00101010 00000000 00000000 00000000CPU 直接执行的指令 机器码（十六进制）B8 2A 00 00 00方便人类阅读的机器码表示 汇编语言mov eax, 42人类可读的助记符表示

汇编语言和机器码是一一对应的，但不同架构的 CPU 有不同的指令集。例如 x86 和 ARM 的汇编语法完全不同。

### 汇编与高级语言的区别

下图展示了高级语言和汇编语言从源码到执行的不同路径：

特性汇编语言高级语言（如 Python） 抽象层级最低级，直接操作硬件高级，操作系统和运行时封装细节 代码量简单功能需要大量代码少量代码实现复杂功能 可读性难以阅读和维护人类友好，易于理解 执行效率最高，无运行时开销较低，有解释或编译开销 硬件控制完全控制 CPU 和内存通过抽象接口间接访问 可移植性差，不同 CPU 需要重写好，同一代码可在多平台运行

### 汇编语言的应用场景

汇编语言虽然不是日常开发的主流选择，但在特定领域仍然不可替代：

领域具体应用 操作系统内核启动引导（Bootloader）、上下文切换、中断处理等必须用汇编编写 嵌入式系统资源极端受限的设备，需要精确控制硬件 逆向工程分析恶意软件、破解软件保护、理解闭源程序逻辑 性能敏感代码加密算法、音视频编解码等对性能要求极高的核心函数 安全研究编写 shellcode、漏洞利用代码 编译器开发理解目标代码生成，编写或优化编译器后端

### x86 架构简介

x86 是 Intel 公司推出的微处理器架构，从 1978 年的 8086 处理器开始，经历了 80286、80386（i386）、80486、Pentium 等代发展。

x86 是最广泛使用的桌面和服务器 CPU 架构，也是学习汇编语言的理想平台。

本教程使用 32 位 x86 架构（也叫 IA-32 或 i386）作为教学基础，因为它的寄存器模型简洁清晰，适合入门。

32 位 x86 汇编是学习汇编语言的最佳起点。掌握了 32 位后，过渡到 64 位（x86-64）非常自然，因为后者是前者的扩展。

### NASM 汇编器简介

NASM（Netwide Assembler） 是一个开源的 x86 汇编器，支持多种输出格式和所有主流操作系统。

与 MASM（微软汇编器）和 GAS（GNU 汇编器）相比，NASM 的语法更加清晰直观，是学习汇编语言的最佳选择。

汇编器语法风格平台特点 NASMIntel 语法跨平台开源、语法清晰、文档完善 MASMIntel 语法Windows微软官方，功能强大但平台受限 GASAT&T 语法跨平台GNU 工具链默认，语法反直觉

本教程全部使用 NASM 汇编器，语法采用 Intel 风格。

---

## 汇编语言 - 环境搭建

Source: https://www.runoob.com/assembly/assembly-env-setup.html

## 汇编语言 - 环境搭建

本章将指导你在 Linux 系统上搭建 NASM 汇编开发环境，并完成你的第一个汇编程序。

### Linux 环境准备

如果你使用的是 Windows 系统，建议安装 WSL（Windows Subsystem for Linux） 或使用虚拟机来获得 Linux 环境。

macOS 用户可以在虚拟机中安装 Linux 或使用 Docker 容器。

本教程所有示例基于 Ubuntu/Debian 系统，其他发行版请自行调整包管理器命令。

### 安装 NASM 汇编器

在 Ubuntu/Debian 系统上，使用 apt 包管理器安装 NASM：

```

$ sudo apt update
$ sudo apt install nasm

```

安装完成后，验证 NASM 是否安装成功：

```

$ nasm -version
NASM version 2.16.01

```

如果你看到版本号输出，说明 NASM 已成功安装。

### 安装链接器

汇编器生成的是 目标文件（.o 文件），不能直接运行。

你需要链接器将目标文件转换为可执行文件。我们使用 GCC 自带的链接器：

```

$ sudo apt install gcc

```

除此之外，你也可以直接使用 GNU 链接器 ld：

```

$ ld --version

```

在本教程中，我们主要使用 gcc 来链接，因为它会自动处理一些底层细节，更适合初学者。

### 安装调试工具 GDB

GDB（GNU Debugger） 是 Linux 下最常用的调试器，可以单步执行程序、查看寄存器状态、检查内存内容。

```

$ sudo apt install gdb

```

验证 GDB 安装：

```

$ gdb --version
GNU gdb (Ubuntu 12.1-0ubuntu1) 12.1

```

### 第一个汇编程序：Hello, runoob

创建一个新文件 `hello.asm`，输入以下内容：

### 实例

; 文件路径：hello.asm
; 第一个汇编程序：打印 Hello, runoob!
; NASM 语法，Linux 32位

section .data ; 数据段：存放已初始化的数据
msg db 'Hello, runoob!', 0xA ; 定义字符串 msg，0xA 是换行符
len equ $ - msg ; 计算字符串长度：当前地址 - msg 起始地址

section .text ; 代码段：存放可执行指令
global _start ; 声明 _start 为程序入口点

_start: ; 程序入口
; 系统调用 sys_write (4)：向 stdout 输出字符串
mov eax, 4 ; 系统调用号 4 = sys_write
mov ebx, 1 ; 文件描述符 1 = stdout（标准输出）
mov ecx, msg ; 要输出的字符串地址
mov edx, len ; 字符串长度
int 0x80 ; 触发中断，执行系统调用

; 系统调用 sys_exit (1)：正常退出程序
mov eax, 1 ; 系统调用号 1 = sys_exit
mov ebx, 0 ; 返回值 0 = 正常退出
int 0x80 ; 触发中断，执行系统调用

### 编译和运行

整个编译运行流程如图所示：

使用以下命令编译和运行这个程序：

```

$ nasm -f elf32 hello.asm -o hello.o # 汇编：将 .asm 转为 .o 目标文件
$ ld -m elf_i386 hello.o -o hello # 链接：将 .o 转为可执行文件
$ ./hello # 运行程序
Hello, runoob!

```

上面三条命令各自的作用：

步骤命令说明 1. 汇编`nasm -f elf32 hello.asm -o hello.o`-f elf32 指定输出格式为 32 位 ELF，-o 指定输出文件名 2. 链接`ld -m elf_i386 hello.o -o hello`-m elf_i386 指定 32 位链接模式，生成可执行文件 hello 3. 运行`./hello`在当前目录执行程序

### 使用 GCC 链接（推荐方式）

如果你的环境中的 ld 配置比较复杂，可以使用 gcc 来链接，它会自动处理 C 运行时等细节：

```

$ nasm -f elf32 hello.asm -o hello.o
$ gcc -m32 hello.o -o hello -nostartfiles
$ ./hello
Hello, runoob!

```

`-nostartfiles` 选项告诉 GCC 不要添加默认的启动代码，因为我们自己定义了 `_start` 入口点。

### 代码结构解析

让我们逐段理解上面的程序：

部分代码说明 数据段声明`section .data`声明数据段，存放变量和常量 字符串定义`msg db 'Hello, runoob!', 0xA`db 定义字节序列，0xA 是换行符 ASCII 码 长度计算`len equ $ - msg`$ 代表当前地址，减去 msg 地址得到字符串长度 代码段声明`section .text`声明代码段，存放指令 入口声明`global _start`将 _start 导出为全局符号，链接器需要它 入口标签`_start:`程序开始执行的位置

`int 0x80` 是 Linux 32 位系统调用指令。它触发一个软件中断，让内核执行我们通过 eax 指定的系统调用。这是用户程序和操作系统内核之间的桥梁。

### 使用 GDB 调试

编译时添加调试信息，然后用 GDB 逐步执行：

```

$ nasm -f elf32 -g hello.asm -o hello.o # -g 添加调试信息
$ ld -m elf_i386 hello.o -o hello
$ gdb ./hello
(gdb) break _start # 在 _start 处设置断点
(gdb) run # 运行程序
(gdb) stepi # 单步执行一条指令
(gdb) info registers # 查看所有寄存器
(gdb) x/s msg # 查看 msg 字符串内容
(gdb) quit # 退出 GDB

```

GDB 命令功能 `break _start`在 _start 标签处设置断点 `run`运行程序，遇断点停止 `stepi`单步执行一条机器指令 `info registers`显示所有寄存器当前值 `x/s 地址`以字符串格式查看指定地址内容

### 常见问题

如果你在 64 位系统上编译 32 位汇编，需要安装 32 位兼容库：`sudo apt install gcc-multilib`

如果 ld 链接时报错 "cannot find entry symbol _start"，检查你的代码中是否写对了 `global _start` 以及标签名是否完全一致（注意下划线）。

---

## 汇编语言 - 基础语法

Source: https://www.runoob.com/assembly/assembly-basic-syntax.html

## 汇编语言 - 基础语法

本章介绍 NASM 汇编程序的基本结构、语法规则和书写规范，帮助你理解汇编代码的骨架。

### 汇编程序的基本结构

一个完整的 NASM 汇编程序通常由以下几个部分组成：

### 实例

; 文件路径：structure.asm
; NASM 程序基本结构示例

section .data ; 数据段：存放已初始化的数据
; 这里定义变量和常量
msg db 'Hello, RUNOOB!', 0xA
len equ $ - msg

section .bss ; BSS 段：存放未初始化的数据
; 这里预留内存空间
buffer resb 64 ; 预留 64 字节的缓冲区

section .text ; 代码段：存放可执行指令
global _start

_start:
; 这里写程序逻辑
mov eax, 4
mov ebx, 1
mov ecx, msg
mov edx, len
int 0x80

mov eax, 1
mov ebx, 0
int 0x80

段（Section）用途特点 `.data`存放已初始化的全局变量和常量编译时确定大小和内容，存入可执行文件 `.bss`存放未初始化的全局变量只在运行时分配空间，不占用可执行文件大小 `.text`存放可执行的机器指令只读，包含程序的全部逻辑代码

至少需要 `.text` 段才能构成一个有效的汇编程序。如果没有数据，可以省略 `.data` 和 `.bss` 段。

### 汇编语句格式

每条汇编语句的通用格式为：

```

[标签:] 指令助记符 [操作数1 [, 操作数2 [, 操作数3]]] [; 注释]

```

各部分说明：

部分是否必须说明 标签（Label）可选代表一个内存地址的符号名，以冒号结尾 指令助记符必须如 mov、add、sub 等，告诉 CPU 要做什么 操作数可选（部分指令无操作数）指令操作的数据对象，可为寄存器、内存地址、立即数 注释可选以分号开头，一直延续到行尾

### 实例

; 各种语句格式示例

; 只有指令，无操作数
ret ; 从子程序返回

; 指令 + 单个操作数
push eax ; 将 eax 的值压入栈中
inc ecx ; ecx 加 1

; 指令 + 两个操作数（最常见）
mov eax, 42 ; 将 42 复制到 eax 寄存器
add ebx, ecx ; ebx = ebx + ecx

; 带有标签
loop_start: ; 标签：标记循环开始位置
dec ecx ; ecx 减 1
jnz loop_start ; 如果 ecx 不为 0，跳回 loop_start

### 注释规范

NASM 使用 分号（;） 表示注释，从分号到行尾的内容都会被汇编器忽略。

### 实例

; 整行注释：说明下面代码块的用途
; 计算两个数的和并输出结果

mov eax, 10 ; 行内注释：将 10 放入 eax
add eax, 20 ; 行内注释：将 20 加到 eax，现在 eax = 30

汇编代码中注释极其重要。没有注释的汇编代码过几周连作者自己都可能看不懂。养成每条指令都写注释的习惯。

### 标识符命名规则

标识符（标签、变量名、常量名等）须遵循以下规则：

规则说明 组成字符字母、数字、下划线 _、点 .、问号 ?、@、$、# 等 起始字符必须以字母、下划线、点或问号开头，不能以数字开头 大小写默认区分大小写（可通过编译选项修改） 保留字不能和指令助记符、寄存器名或 NASM 关键字重名

### 实例

; 合法的标识符
my_variable: ; 字母开头 + 下划线
.loop_start: ; 点开头（局部标签）
?error_handler: ; 问号开头
counter2: ; 字母 + 数字

; 不合法的标识符（仅供参考，不要使用）
; 1st_value: ; 错误：不能以数字开头
; mov: ; 错误：mov 是保留字
; my-variable: ; 错误：减号不是合法字符

### 伪指令（Directives）

伪指令 是给汇编器的命令，不是给 CPU 的指令，用于控制汇编过程和定义数据结构。

伪指令用途示例 `db`定义字节（1 字节）`byte_val db 0x55` `dw`定义字（2 字节）`word_val dw 0x1234` `dd`定义双字（4 字节）`dword_val dd 0x12345678` `equ`定义常量`MAX_SIZE equ 256` `resb`预留字节空间`buffer resb 128` `resw`预留字空间`wbuf resw 64` `resd`预留双字空间`dbuf resd 32` `%define`宏定义常量`%define COUNT 10`

### 大小写规范

NASM 默认对标签和标识符 区分大小写：

### 实例

; 大小写敏感的示例

section .data
msg db 'RUNOOB', 0 ; 定义变量 msg

section .text
global _start

_start:
mov eax, MSG ; 错误：MSG 和 msg 不同（除非开启忽略大小写）
mov eax, msg ; 正确：msg 与定义完全一致

MOV EAX, 42 ; 语法正确：指令助记符不区分大小写
mov eax, 42 ; 推荐写法：用小写，可读性更好

指令助记符和寄存器名不区分大小写（`MOV`、`Mov`、`mov` 效果相同），但推荐的风格是统一小写。

### 数值表示方式

NASM 支持多种进制的数值表示：

### 实例

; NASM 中不同进制的表示方式

mov eax, 42 ; 十进制：直接写数字
mov eax, 0x2A ; 十六进制：0x 前缀（推荐写法）
mov eax, 2Ah ; 十六进制：h 后缀
mov eax, 0o52 ; 八进制：0o 前缀
mov eax, 52o ; 八进制：o 后缀
mov eax, 101010b ; 二进制：b 后缀
mov eax, 0b101010 ; 二进制：0b 前缀

推荐使用 `0x` 前缀表示十六进制（如 `0x2A`），这样不容易和标签混淆。

### 一个完整的语法示例

下面程序综合运用以上语法元素，计算 1 到 10 的和并输出：

### 实例

; 文件路径：sum.asm
; 计算 1+2+...+10 并输出结果字符

section .data
result db 0 ; 存放计算结果（1字节）
newline db 0xA ; 换行符

section .text
global _start

_start:
; 初始化寄存器和变量
mov ecx, 10 ; 循环计数器：从 10 开始倒数
mov eax, 0 ; eax 存放累加和，初始为 0

sum_loop: ; 循环开始标签
add eax, ecx ; eax = eax + ecx
dec ecx ; ecx 减 1
jnz sum_loop ; 如果 ecx != 0，继续循环

; 此时 eax = 55（10+9+...+1）
add eax, '0' ; 将数字转为 ASCII 字符（'0'=48，55+48=103='g'，不对）
; 实际演示需要更复杂的转换，见后续章节

; 这里只输出 result（简化演示）
mov [result], al ; 将累加结果存入 result

; 退出程序
mov eax, 1
mov ebx, 0
int 0x80

注意：上面的示例中，直接加 '0' 只在数字是 0-9 范围内正确。处理两位数及以上的数字转换，将在后续章节详细讲解。

---

## 汇编语言 - 内存分段

Source: https://www.runoob.com/assembly/assembly-memory-segment.html

## 汇编语言 - 内存分段

内存分段（Memory Segmentation）是 x86 架构中的一个核心概念，它决定了程序在内存中如何组织代码、数据和栈空间。

### 为什么需要分段

在 x86 实模式下，CPU 使用 16 位寄存器，但需要访问 1MB 的内存空间（20 位地址）。

16 位寄存器只能表示 64KB 范围（`2^16 = 65536`），远远不够。

Intel 的解决方案是将内存划分为 段（Segment），每个段有基地址和偏移量，通过 段地址 + 偏移量 的方式组合出完整的物理地址。

物理地址计算公式：

```

物理地址 = 段地址 × 16 + 偏移地址

```

虽然在保护模式（32 位）下，分段机制更多地用于内存保护和权限控制，但理解分段对于理解汇编程序结构仍然至关重要。

### 三种基本段

一个典型的汇编程序使用三个主要段：

段名称英文名用途对应的 Section 代码段Code Segment存放可执行的机器指令`section .text` 数据段Data Segment存放已初始化的全局变量`section .data` 栈段Stack Segment存放函数调用信息、局部变量操作系统自动管理

### 代码段（.text）

代码段存放程序的全部机器指令，是 只读、可执行 的内存区域。

操作系统在加载程序时，将代码段映射到只读内存页，防止程序意外修改指令。

### 实例

; 文件路径：code_segment.asm
; 演示代码段使用

section .text ; 开始代码段
global _start

_start:
mov eax, 1 ; 这些指令都存放在代码段中
mov ebx, 0
int 0x80

; 以下是一个辅助函数，也存放在代码段
my_function:
mov eax, 42
ret

### 数据段（.data）

数据段存放程序中 已初始化的全局变量。

数据段是可读可写的，程序运行时可以修改其中的内容。

### 实例

; 文件路径：data_segment.asm
; 演示数据段使用

section .data
; 定义各种类型的已初始化变量
msg db 'Hello, runoob!', 0 ; 字符串变量（字节序列）
count db 100 ; 字节变量，初始值 100
pi dd 314159 ; 双字变量，存放 π 的近似值 × 100000
array db 1, 2, 3, 4, 5 ; 字节数组

section .text
global _start

_start:
; 读取数据段中的变量
mov al, [count] ; 将 count 的值（100）加载到 al 寄存器
; ...
mov eax, 1
mov ebx, 0
int 0x80

### BSS 段（.bss）

BSS（Block Started by Symbol）段存放 未初始化的全局变量。

与数据段不同，BSS 段在可执行文件中不占用实际空间，只在程序加载时由操作系统分配内存并清零。

### 实例

; 文件路径：bss_segment.asm
; 演示 BSS 段使用

section .bss
buffer resb 256 ; 预留 256 字节的缓冲区（未初始化）
num_array resd 100 ; 预留 100 个双字（400 字节）

section .data
; 已初始化数据

section .text
global _start

_start:
; 使用 BSS 段的缓冲区
mov byte [buffer], 'A' ; 向缓冲区写入字符 'A'
; ...
mov eax, 1
mov ebx, 0
int 0x80

将未初始化的数据放在 BSS 段而不是 .data 段中，可以减小可执行文件的体积。例如，一个 10KB 的未初始化缓冲区在 BSS 段中不占用文件大小，但如果放在 .data 段中则会让文件增大 10KB。

### 段寄存器

x86 架构有 6 个段寄存器，用于跟踪当前正在使用的段：

段寄存器英文全称用途 CSCode Segment指向代码段，存放当前执行指令所在的段基址 DSData Segment指向数据段，存放大部分数据访问的段基址 SSStack Segment指向栈段，存放栈所在的段基址 ESExtra Segment附加段寄存器，用于字符串操作等额外数据段 FSGeneral Purpose通用段寄存器，常用于线程局部存储 GSGeneral Purpose通用段寄存器，常用于线程局部存储

在 32 位保护模式下，程序员通常不需要手动设置段寄存器，操作系统和链接器会处理好。

### 内存分段示意图

一个运行中的程序在内存中的分布如下：

一个运行中的程序在内存中的分布大致如下：

```

高地址
+-------------------+
| 栈 (Stack) | <-- SS:ESP 指向栈顶
| 向下增长 |
+-------------------+
| |
| 空闲内存 |
| |
+-------------------+
| BSS 段 (.bss) | 未初始化的全局变量
+-------------------+
| 数据段 (.data) | 已初始化的全局变量
+-------------------+
| 代码段 (.text) | 程序指令（只读）
+-------------------+
| 保留区域 | 操作系统使用
+-------------------+
低地址

```

### 实模式 vs 保护模式

特性实模式（16 位）保护模式（32 位） 内存寻址段地址 × 16 + 偏移段选择子查描述符表 + 偏移 最大内存1MB4GB（32 位） 内存保护无保护有页级和段级保护 多任务不支持支持 寻址方式segment:offsetselector:offset（通过描述符表）

本教程主要介绍 32 位保护模式下的汇编编程。在这种模式下，段的概念更多地是对内存区域的逻辑划分，物理地址由分页机制管理。你可以把 section .data/.bss/.text 简单地理解为程序不同部分在内存中的标记。

---

## 汇编语言 - 寄存器

Source: https://www.runoob.com/assembly/assembly-registers.html

## 汇编语言 - 寄存器

寄存器（Register）是 CPU 内部的高速存储单元，是汇编编程中最频繁操作的对象。理解寄存器是学好汇编语言的关键第一步。

### 什么是寄存器

寄存器是 CPU 芯片内部集成的 超高速小型存储器，用于暂存指令、数据和地址。

与内存不同，寄存器就嵌在 CPU 内部，CPU 访问寄存器几乎零延迟，而访问内存则需要几十到几百个时钟周期。

在汇编语言中，绝大多数运算都是围绕寄存器展开的——数据从内存加载到寄存器，在寄存器中完成运算，再将结果存回内存。

可以把寄存器理解为 CPU 的"工作台"。工作台上的工具随时可用，而内存则像是仓库，需要走过去取放。

### x86 32 位寄存器分类

x86 32 位架构提供了多种类型的寄存器，各有不同的用途。下图展示了完整的寄存器分类：

下面分别详细介绍各类寄存器：

### 通用寄存器

通用寄存器（General Purpose Registers） 是最常用的寄存器，用于存放运算数据和临时结果。

x86 提供了 8 个 32 位通用寄存器：

32位16位低8位高8位（低16位）主要用途 EAXAXALAH累加器，存放函数返回值、算术运算结果 EBXBXBLBH基址寄存器，常用于存放内存基地址 ECXCXCLCH计数器，常用于循环计数和移位 EDXDXDLDH数据寄存器，存放乘除法的高位结果 ESISISIL-源变址寄存器，字符串操作的源地址 EDIDIDIL-目的变址寄存器，字符串操作的目标地址 EBPBPBPL-基址指针，指向当前栈帧的底部 ESPSPSPL-栈指针，始终指向栈顶

名称规律：E 前缀表示 Extended（扩展到 32 位），X 后缀表示可拆分为高低字节。

### 实例

; 文件路径：register_parts.asm
; 演示寄存器的各部分访问

section .text
global _start

_start:
mov eax, 0x12345678 ; 完整的 32 位寄存器
; 此时：EAX = 0x12345678
; AX = 0x5678 (低 16 位)
; AH = 0x56 (高 8 位，指 AX 的高 8 位)
; AL = 0x78 (低 8 位)

mov ax, 0xAABB ; 修改 AX（低 16 位）
; 此时：EAX = 0x1234AABB (高 16 位保持不变！)
; AX = 0xAABB
; AL = 0xBB

mov al, 0xCC ; 修改 AL（最低 8 位）
; 此时：EAX = 0x1234AACC (只有低 8 位变了)
; AX = 0xAACC
; AL = 0xCC

mov eax, 1
mov ebx, 0
int 0x80

修改 32 位寄存器的低 16 位（如 AX）时，高 16 位保持不变。但将 32 位寄存器作为目标操作数时，会覆盖整个 32 位。这是初学者容易出错的地方。

### 段寄存器

段寄存器用于指定当前使用的内存段：

寄存器名称用途 CS代码段寄存器指向当前指令所在的段 DS数据段寄存器指向数据所在的段 SS栈段寄存器指向栈所在的段 ES附加段寄存器额外的数据段 FS附加段寄存器通用，常用于线程局部存储 GS附加段寄存器通用，常用于线程局部存储

在 32 位保护模式下编程时，操作系统已经设置好了段寄存器，你通常不需要手动修改它们。

### 指针和变址寄存器

这些寄存器主要用于访问内存，存放内存地址：

寄存器全称用途 EIP指令指针（Instruction Pointer）指向 CPU 下一条要执行的指令地址（不可直接访问） ESP栈指针（Stack Pointer）指向栈顶，PUSH/POP 指令自动调整它 EBP基址指针（Base Pointer）指向当前函数栈帧的底部，用于访问函数参数和局部变量 ESI源变址（Source Index）字符串/内存操作的源地址 EDI目的变址（Destination Index）字符串/内存操作的目标地址

ESP 和 EBP 不能当作普通通用寄存器随意使用。ESP 指向栈顶，push/pop/call/ret 都会改变它；EBP 是访问函数参数的关键。随意修改它们会导致程序崩溃。

### 标志寄存器（EFLAGS）

EFLAGS 是一个 32 位寄存器，每个比特位代表一个状态标志（Flag）。

你不能直接读写整个 EFLAGS，但 CPU 会根据运算结果自动更新这些标志位，条件跳转指令则根据标志位决定是否跳转。

标志位名称含义 CF进位标志（Carry Flag）无符号运算产生进位/借位时置 1 PF奇偶标志（Parity Flag）结果低 8 位中 1 的个数为偶数时置 1 AF辅助进位标志低 4 位向高 4 位进位/借位时置 1 ZF零标志（Zero Flag）运算结果为 0 时置 1 SF符号标志（Sign Flag）运算结果为负数时置 1（等于结果的最高位） OF溢出标志（Overflow Flag）有符号运算溢出时置 1

### 实例

; 文件路径：flags_demo.asm
; 演示运算对标志位的影响

section .text
global _start

_start:
mov eax, 10
sub eax, 10 ; 10 - 10 = 0
; ZF = 1（结果为零）
; SF = 0（结果非负）
; CF = 0（无借位）

mov eax, 0xFFFFFFFF
add eax, 1 ; 0xFFFFFFFF + 1 = 0x100000000（超出了 32 位）
; ZF = 1（32位结果为0）
; CF = 1（产生进位）
; OF = 0（有符号角度看无溢出）

mov eax, 1
mov ebx, 0
int 0x80

### 寄存器使用约定

在实际编程中，一些寄存器有约定俗成的用法——称为 调用约定（Calling Convention）：

寄存器调用约定中的用途 EAX存放函数返回值 ECX计数器（循环计数） EDX存放除法的高位结果、扩展 EAX EBX、ESI、EDI、EBP被调用函数必须保存和恢复（callee-saved） EAX、ECX、EDX调用者负责保存（caller-saved）

在编写自己的汇编程序时，不一定严格遵循调用约定。但如果你要和 C 语言混合编程，就必须遵守 cdecl 等调用约定。

---

## 汇编语言 - 系统调用

Source: https://www.runoob.com/assembly/assembly-syscall.html

## 汇编语言 - 系统调用

系统调用（System Call）是用户程序与操作系统内核交互的唯一途径。通过系统调用，汇编程序可以读写文件、分配内存、创建进程等。

### 什么是系统调用

用户程序运行在受限的 用户态（User Mode），无法直接访问硬件或执行特权操作。

当程序需要执行 IO 操作、内存分配等内核级功能时，必须通过 系统调用 向操作系统内核请求服务。

系统调用的流程是：用户程序发起系统调用 → CPU 切换到内核态 → 内核执行对应服务 → 返回用户态继续执行。

可以把系统调用理解为：用户程序给操作系统打个电话，请求帮忙做一件自己没有权限做的事情。

### Linux 系统调用机制

在 Linux 32 位系统中，系统调用通过以下步骤完成：

- 将 系统调用号 放入 EAX 寄存器
- 将参数放入 EBX、ECX、EDX、ESI、EDI 等寄存器
- 执行 `int 0x80` 指令触发软件中断
- CPU 切换到内核态，内核根据 EAX 中的调用号执行对应服务
- 内核将返回值放入 EAX，切换回用户态继续执行

### 常用系统调用一览

系统调用调用号（EAX）功能参数 sys_exit1退出程序EBX = 返回值（exit code） sys_fork2创建子进程无参数 sys_read3读取文件/输入EBX=fd, ECX=buf, EDX=count sys_write4写入文件/输出EBX=fd, ECX=buf, EDX=count sys_open5打开文件EBX=filename, ECX=flags, EDX=mode sys_close6关闭文件EBX=fd sys_creat8创建文件EBX=filename, ECX=mode sys_lseek19移动文件指针EBX=fd, ECX=offset, EDX=whence sys_brk45调整数据段大小（内存分配）EBX=新地址

完整的系统调用列表可在 `/usr/include/asm/unistd_32.h` 中查看。注意 32 位和 64 位的系统调用号不同。

### sys_write - 输出字符串

最常用的系统调用，用于向屏幕打印内容：

### 实例

; 文件路径：write_syscall.asm
; 使用 sys_write 输出字符串

section .data
msg db 'Hello, runoob! Welcome to assembly.', 0xA
len equ $ - msg

section .text
global _start

_start:
; 调用 sys_write (4)
mov eax, 4 ; 系统调用号 4 = sys_write
mov ebx, 1 ; 文件描述符 1 = stdout（标准输出）
mov ecx, msg ; 要输出的数据地址
mov edx, len ; 数据长度（字节数）
int 0x80 ; 触发系统调用

; 成功时 EAX 返回实际写入的字节数

; 退出程序
mov eax, 1
mov ebx, 0
int 0x80

运行结果：

```

$ nasm -f elf32 write_syscall.asm -o write_syscall.o
$ ld -m elf_i386 write_syscall.o -o write_syscall
$ ./write_syscall
Hello, runoob! Welcome to assembly.

```

### sys_read - 读取用户输入

从标准输入（键盘）读取数据：

### 实例

; 文件路径：read_syscall.asm
; 使用 sys_read 读取用户输入并回显

section .bss
buffer resb 64 ; 预留 64 字节输入缓冲区

section .data
prompt db 'Please enter your name: '
prompt_len equ $ - prompt
output_msg db 'Hello, '
output_msg_len equ $ - output_msg

section .text
global _start

_start:
; 输出提示信息
mov eax, 4
mov ebx, 1
mov ecx, prompt
mov edx, prompt_len
int 0x80

; 读取用户输入
mov eax, 3 ; 系统调用号 3 = sys_read
mov ebx, 0 ; 文件描述符 0 = stdin（标准输入）
mov ecx, buffer ; 输入缓冲区地址
mov edx, 64 ; 最多读取 64 字节
int 0x80

; EAX 返回实际读取的字节数（包含末尾的换行符）
mov esi, eax ; 保存实际输入长度到 esi

; 输出 "Hello, "
mov eax, 4
mov ebx, 1
mov ecx, output_msg
mov edx, output_msg_len
int 0x80

; 输出用户输入的内容
mov eax, 4
mov ebx, 1
mov ecx, buffer
mov edx, esi ; 使用实际输入的字节数
int 0x80

; 退出程序
mov eax, 1
mov ebx, 0
int 0x80

运行结果：

```

$ nasm -f elf32 read_syscall.asm -o read_syscall.o
$ ld -m elf_i386 read_syscall.o -o read_syscall
$ ./read_syscall
Please enter your name: runoob
Hello, runoob

```

### sys_exit - 退出程序

每个程序都必须调用 sys_exit 来正常退出，否则 CPU 会继续执行后面的垃圾数据导致段错误。

### 实例

; 使用不同退出码退出

mov eax, 1 ; 系统调用号 1 = sys_exit
mov ebx, 0 ; 退出码 0 = 正常退出（也可以是非零值）
int 0x80

退出码可以在 shell 中通过 `$?` 查看：

```

$ ./program
$ echo $?
0

```

### 系统调用通用模板

编写系统调用时，可以遵循以下通用模板：

### 实例

; 系统调用通用模板
; 适用于 Linux 32 位系统

; 步骤1：将系统调用号放入 EAX
mov eax, 系统调用号

; 步骤2：按顺序放入参数
mov ebx, 第1个参数
mov ecx, 第2个参数
mov edx, 第3个参数
mov esi, 第4个参数
mov edi, 第5个参数

; 步骤3：触发系统调用
int 0x80

; 步骤4：检查返回值（在 EAX 中）
; 通常负值表示错误
cmp eax, 0
jl error_handler ; 如果 EAX < 0，跳转到错误处理

`int 0x80` 触发后，EAX 存放返回值。大多数系统调用成功时返回非负值，失败时返回负的错误码（如 -1 表示 EPERM）。

---

## 汇编语言 - 寻址方式

Source: https://www.runoob.com/assembly/assembly-addressing.html

## 汇编语言 - 寻址方式

寻址方式（Addressing Mode）决定了 CPU 如何定位指令的操作数——数据从哪里来、结果存到哪里去。

### 什么是寻址方式

每条汇编指令的操作数可以是立即数、寄存器中的值或内存中的数据。

寻址方式 就是告诉 CPU 如何计算操作数的实际地址或直接给出操作数值。

x86 架构提供了多种灵活的寻址方式，理解每种方式的使用场景是写出高效汇编代码的基础。

### 立即寻址（Immediate Addressing）

操作数直接包含在指令中，是一个常量值。

源操作数是立即数，CPU 直接从指令中读取，不需要访问内存或寄存器。

### 实例

; 立即寻址示例

mov eax, 42 ; 42 是立即数，直接编码在指令中
add ebx, 100 ; 100 是立即数
mov ecx, 0x2A ; 十六进制立即数
mov edx, 'A' ; 字符 'A' = 0x41，也是立即数

立即寻址是最快的"寻址方式"，因为数据就在指令流中，CPU 取指令的同时就拿到了数据。但立即数只能作为源操作数，不能作为目标操作数。不能写 `mov 42, eax`。

### 寄存器寻址（Register Addressing）

操作数存放在寄存器中，CPU 直接操作寄存器。

这同样是最快的操作方式之一，因为没有内存访问开销。

### 实例

; 寄存器寻址示例

mov eax, ebx ; 将 ebx 的值复制到 eax（两个操作数都是寄存器寻址）
add ecx, edx ; ecx = ecx + edx
push eax ; 将 eax 的值压入栈
inc ebx ; ebx = ebx + 1

寄存器寻址速度最快，在写汇编代码时优先使用寄存器存放频繁访问的数据。但寄存器数量有限，不能把所有数据都塞在寄存器中。

### 直接寻址（Direct Addressing）

操作数是一个内存地址，地址直接写在指令中（以变量标签的形式）。

CPU 需要访问一次内存来读取或写入数据。

### 实例

; 文件路径：direct_addr.asm
; 直接寻址示例

section .data
value dd 12345678 ; 在内存中定义一个双字变量
name db 'runoob', 0

section .text
global _start

_start:
; 直接寻址读取内存
mov eax, [value] ; 从内存地址 value 读取 4 字节到 eax
; eax 现在是 12345678

; 直接寻址写入内存
mov dword [value], 98765 ; 将 98765 写入 value 所在的内存地址

; 直接寻址读取字节
mov al, [name] ; 读取 name 的第一个字节 'r' = 0x72
mov bl, [name + 1] ; 读取 name 的第二个字节 'u' = 0x75

mov eax, 1
mov ebx, 0
int 0x80

### 寄存器间接寻址（Register Indirect Addressing）

寄存器中存放的是一个内存地址，CPU 用该地址去访问内存。

方括号内的寄存器被当作指针使用。

### 实例

; 文件路径：indirect_addr.asm
; 寄存器间接寻址示例

section .data
msg db 'Hello, RUNOOB!', 0xA
len equ $ - msg

section .text
global _start

_start:
mov eax, msg ; 将 msg 的地址（指针）放入 eax
mov al, [eax] ; 间接寻址：读取 eax 指向的内存字节
; 现在 al = 'H' = 0x48

; 遍历字符串，将小写字母转为大写
mov esi, msg ; esi 指向字符串起始位置
mov ecx, len ; ecx 存放字符串长度

convert_loop:
mov al, [esi] ; 间接寻址：读取 esi 指向的字符
cmp al, 'a' ; 是否大于等于 'a'
jb next_char ; 否，跳过
cmp al, 'z' ; 是否小于等于 'z'
ja next_char ; 否，跳过
sub al, 32 ; 转为大写（ASCII 表中小写-大写=32）
mov [esi], al ; 间接寻址：写回 esi 指向的位置

next_char:
inc esi ; 指针移动到下一个字符
loop convert_loop ; 继续循环直到全部处理完

; 输出转换后的字符串
mov eax, 4
mov ebx, 1
mov ecx, msg
mov edx, len
int 0x80

mov eax, 1
mov ebx, 0
int 0x80

间接寻址是数组遍历、字符串操作和数据结构访问的基础。ESI 和 EDI 是专门设计用来配合间接寻址的寄存器，配合 `inc esi` 可以轻松遍历连续内存。

### 基址寻址（Base Addressing）

有效地址 = 基址寄存器的值 + 偏移量（位移量）。

基址寄存器可以是 EBX、EBP、ESI、EDI 等。

### 实例

; 基址寻址示例：访问结构体成员

section .data
; 模拟一个简单的结构体：{id, age, score}
; id = 2 字节
; age = 2 字节
; score = 4 字节
student db 0x01, 0x00 ; id = 1
db 0x14, 0x00 ; age = 20
dd 95 ; score = 95

section .text
global _start

_start:
mov ebx, student ; ebx 存放结构体基址

; 基址 + 偏移量访问各成员
mov ax, [ebx] ; 读取 id（偏移 0）
mov ax, [ebx + 2] ; 读取 age（偏移 2）
mov eax, [ebx + 4] ; 读取 score（偏移 4）

; 修改 age
mov word [ebx + 2], 21 ; age = 21

; 修改 score
mov dword [ebx + 4], 98 ; score = 98

mov eax, 1
mov ebx, 0
int 0x80

### 变址寻址（Indexed Addressing）

使用变址寄存器（ESI 或 EDI）加上偏移量来访问数组元素。

### 实例

; 变址寻址示例：遍历数组

section .data
array dd 10, 20, 30, 40, 50 ; 5 个双字元素的数组
array_len equ ($ - array) / 4 ; 元素个数 = 总字节数 / 4

section .text
global _start

_start:
mov ecx, array_len ; 循环计数器
mov esi, 0 ; 变址（下标从 0 开始）
mov ebx, 0 ; 累加和

sum_loop:
mov eax, [array + esi * 4] ; 变址寻址：array + 下标 * 元素大小
; esi * 4 因为每个元素是 4 字节
add ebx, eax ; 累加到 ebx
inc esi ; 下标加 1
loop sum_loop
; ebx = 10+20+30+40+50 = 150

; 变址 + 偏移：访问第二个元素
; array + 2*4 = array + 8，即 30
mov eax, [array + 2*4] ; eax = 30

mov eax, 1
mov ebx, 0
int 0x80

### 基址变址寻址（Base-Indexed Addressing）

有效地址 = 基址寄存器 + 变址寄存器 × 比例因子 + 偏移量。

这是 x86 中最强大的寻址方式，可以在一条指令中完成地址计算。

注意：x86 仅支持一个变址寄存器 × 比例因子，不支持多个带比例因子的寄存器。

### 实例

; 基址变址寻址示例：二维数组访问

section .data
; 3 行 4 列的二维数组
matrix dd 1, 2, 3, 4
dd 5, 6, 7, 8
dd 9, 10, 11, 12

section .text
global _start

_start:
; 访问 matrix[1][2]（第 2 行第 3 列，下标从 0 开始）
; 地址 = matrix + 行*每行字节数 + 列*每元素字节数
; = matrix + 1*16 + 2*4
; = matrix + 24

mov ebx, matrix ; 基址寄存器
mov esi, 24 ; 计算好的总偏移量

; 正确格式：基址 + 偏移量
mov eax, [ebx + esi] ; eax = 7

; 标准基址变址寻址格式：[基址 + 变址*比例因子 + 偏移量]
; 直接访问第 6 个元素（matrix[6] = 7）
mov edi, 6
mov eax, [matrix + edi*4] ; eax = 7

; 程序退出
mov eax, 1
mov ebx, 0
int 0x80

### 寻址方式总览

下图完整展示了 x86 的 7 种寻址方式及其工作原理：

寻址方式语法格式有效地址/值典型用途 立即寻址`mov eax, 42`42（常数值）初始化、常量运算 寄存器寻址`mov eax, ebx`寄存器 ebx 的值寄存器间数据传递 直接寻址`mov eax, [var]`内存地址 var 处的值访问全局变量 间接寻址`mov eax, [ebx]`地址 = ebx 的值指针操作、遍历内存 基址寻址`mov eax, [ebx+8]`地址 = ebx + 8结构体成员访问 变址寻址`mov eax, [arr+esi*4]`地址 = arr + esi × 4一维数组访问 基址变址`mov eax, [ebx+esi*4+8]`地址 = ebx + esi × 4 + 8二维数组、复杂结构体

在 32 位保护模式下，所有的通用寄存器都可以作为基址或变址寄存器。这与16位实模式有很大不同（16位模式下只有 BX、BP、SI、DI 可用于寻址）。32 位的灵活性让寻址变得更加方便。

---

## 汇编语言 - 变量

Source: https://www.runoob.com/assembly/assembly-variables.html

## 汇编语言 - 变量

变量是程序中存储数据的基本单元。在汇编语言中定义和使用变量比高级语言更接近底层，你需要直接控制变量的大小、类型和内存布局。

### 汇编中的变量概念

在汇编语言中，变量 本质上是内存中一个有名字的存储位置。

与高级语言不同，汇编变量没有自动的类型检查——一个标签只是内存地址的别名，你可以用任何大小去读写它。

NASM 提供了三种段来存放不同类型的变量：`.data`（已初始化）、`.bss`（未初始化）和 `.rodata`（只读）。

### 在 .data 段定义已初始化变量

使用 DB、DW、DD、DQ、DT 伪指令定义不同大小的变量：

伪指令数据大小含义示例 DB1 字节Define Byte`flag db 1` DW2 字节Define Word`count dw 1000` DD4 字节Define Doubleword`price dd 9999` DQ8 字节Define Quadword`big_val dq 0x1234567890ABCDEF` DT10 字节Define Ten Bytes`ext_val dt 3.14`

### 实例

; 文件路径：variables_data.asm
; 演示 .data 段中各种变量的定义

section .data
; 字节变量
status db 1 ; 1 字节，值为 1
grade db 'A' ; 1 字节，字符 'A' = 0x41

; 字变量（2 字节）
year dw 2026 ; 2 字节，值为 2026

; 双字变量（4 字节）
salary dd 50000 ; 4 字节，值为 50000

; 多字节序列
msg db 'Hello, runoob!', 0 ; 字符串以 0 结尾（C 风格）
hex_bytes db 0x55, 0xAA, 0x00, 0xFF ; 十六进制字节序列

; 重复值
stars db 10 dup('*') ; 10 个星号字符

; 带名称的多个变量
x dd 10
y dd 20
z dd 30

section .text
global _start

_start:
; 读取变量值到寄存器
mov al, [status] ; al = 1（读取 1 字节）
mov ax, [year] ; ax = 2026（读取 2 字节）
mov eax, [salary] ; eax = 50000（读取 4 字节）

; 修改变量值
mov byte [status], 0 ; status = 0
mov dword [x], 100 ; x = 100

mov eax, 1
mov ebx, 0
int 0x80

使用 `[变量名]` 访问变量时，务必确保读取/写入的字节数与变量定义时的字节数匹配。用 `mov al, [status]` 读 1 字节，用 `mov eax, [salary]` 读 4 字节。NASM 会检查操作数大小是否一致。

### 在 .bss 段预留未初始化空间

使用 RESB、RESW、RESD 等伪指令预留空间（不初始化）：

伪指令预留大小示例 RESB字节`buffer resb 256` RESW字（2字节）`wbuf resw 100` RESD双字（4字节）`dbuf resd 50` RESQ四字（8字节）`qbuf resq 25`

### 实例

; 文件路径：variables_bss.asm
; 演示 .bss 段中预留空间

section .bss
input_buf resb 128 ; 预留 128 字节的输入缓冲区
numbers resd 100 ; 预留 100 个双字（400 字节）
temp resb 1 ; 预留 1 字节的临时变量

section .data
prompt db 'Enter a number: '
prompt_len equ $ - prompt

section .text
global _start

_start:
; 输出提示
mov eax, 4
mov ebx, 1
mov ecx, prompt
mov edx, prompt_len
int 0x80

; 读取输入到 .bss 段的缓冲区
mov eax, 3
mov ebx, 0
mov ecx, input_buf ; 使用 .bss 段预留的空间
mov edx, 128
int 0x80

; 向 numbers 数组填入数据
mov dword [numbers], 42 ; numbers[0] = 42
mov dword [numbers + 4], 100 ; numbers[1] = 100

mov eax, 1
mov ebx, 0
int 0x80

### 字节、字、双字的内存布局

x86 使用 小端序（Little Endian）——低字节存放在低地址。

例如，当定义 `value dd 0x12345678` 时，内存中的布局为：

```

地址 ：[value] [value+1] [value+2] [value+3]
内容 ：0x78 0x56 0x34 0x12

```

### 实例

; 文件路径：endianness.asm
; 演示小端序存储

section .data
value dd 0x12345678 ; 双字，4 字节

section .text
global _start

_start:
; 按不同大小读取同一个变量
mov eax, [value] ; 读 4 字节：eax = 0x12345678
mov ax, [value] ; 读 2 字节：ax = 0x5678（低2字节）
mov al, [value] ; 读 1 字节：al = 0x78（最低1字节）

; 验证小端序：低地址存放低字节
mov al, [value] ; al = 0x78
mov bl, [value + 1] ; bl = 0x56
mov cl, [value + 2] ; cl = 0x34
mov dl, [value + 3] ; dl = 0x12

mov eax, 1
mov ebx, 0
int 0x80

### 变量的初始化与访问

汇编中变量的操作遵循"加载-运算-存储"三步模式：

### 实例

; 文件路径：var_operations.asm
; 演示变量操作的三步模式

section .data
a dd 100
b dd 200
result dd 0

section .text
global _start

_start:
; result = a + b
; 第1步：加载
mov eax, [a] ; 加载 a 到 eax
; 第2步：运算
add eax, [b] ; eax = eax + b
; 第3步：存储
mov [result], eax ; 存储结果到 result

; result = a * 2 - b
mov eax, [a]
imul eax, 2 ; eax = a * 2
sub eax, [b] ; eax = eax - b
mov [result], eax

mov eax, 1
mov ebx, 0
int 0x80

汇编不支持 `mov [result], [a] + [b]` 这种高级语法。x86 的 mov 指令不能同时用两个内存操作数，必须先经过寄存器中转。

### 变量大小操作符

当汇编器无法推断操作数大小时，需要显式指定：

### 实例

; 使用大小操作符明确数据大小

section .data
var dd 0

section .text
global _start

_start:
; 大小操作符：byte, word, dword, qword
mov byte [var], 1 ; 写入 1 字节
mov word [var], 1000 ; 写入 2 字节
mov dword [var], 999999 ; 写入 4 字节

; 当目标大小明确时（由寄存器决定），可以省略
mov al, [var] ; al 是 1 字节，自动按 byte 读取
mov ax, [var] ; ax 是 2 字节，自动按 word 读取
mov eax, [var] ; eax 是 4 字节，自动按 dword 读取

; 但下面这行会报错（大小不明确）：
; mov [var], 1 ; 错误！1 可以是字节/字/双字
; 必须写成：
mov dword [var], 1 ; 明确指定 4 字节

mov eax, 1
mov ebx, 0
int 0x80

---

## 汇编语言 - 常量

Source: https://www.runoob.com/assembly/assembly-constants.html

## 汇编语言 - 常量

常量是在编译时确定且程序运行期间不可改变的值。在汇编中使用常量可以让代码更易读、更易维护。

### 什么是汇编中的常量

在 NASM 中，常量是通过 伪指令 定义的符号，汇编器在编译时将所有常量引用替换为实际值。

与变量不同，常量不占用数据段内存，它们在 编译时 就被替换，相当于 C 语言中的 `#define`。

### EQU - 等值常量

`EQU` 是最常用的定义常量方式，定义后不可重新定义：

### 实例

; EQU 常量定义示例

; 数值常量
MAX_SIZE equ 256 ; 缓冲区最大大小
DEFAULT_PORT equ 8080 ; 默认端口号
TRUE equ 1 ; 真值
FALSE equ 0 ; 假值

; 字符常量
LF equ 0xA ; 换行符
CR equ 0xD ; 回车符
NULL equ 0 ; 空字符

; 表达式常量
ARRAY_SIZE equ 100 * 4 ; 100 个双字的字节数
BUFFER_SIZE equ MAX_SIZE * 2 ; 引用另一个常量
TOTAL equ 10 + 20 + 30 ; 计算结果

section .data
; 使用常量
buffer db MAX_SIZE dup(0) ; 定义 MAX_SIZE 字节的缓冲区

section .text
global _start

_start:
mov eax, MAX_SIZE ; eax = 256
mov ebx, BUFFER_SIZE ; ebx = 512
mov ecx, ARRAY_SIZE ; ecx = 400

mov eax, 1
mov ebx, 0
int 0x80

`EQU` 定义的常量不能在程序中修改，也不能重新定义。如果需要可以重新定义的常量（类似宏变量），使用 `%define`。

### %define - 宏定义常量

`%define` 比 EQU 更灵活，支持重新定义和作用域控制：

### 实例

; %define 常量定义与重定义示例

%define VERSION 1 ; 定义版本号
%define APP_NAME 'runoob' ; 定义应用名称

; 使用 %define 定义的常量
mov eax, VERSION ; eax = 1

; 可以重新定义（EQU 不允许这样做）
%define VERSION 2 ; 版本升级
mov eax, VERSION ; 现在 eax = 2

; %undef 取消定义
%undef VERSION ; VERSION 不再有效
; mov eax, VERSION ; 错误：VERSION 未定义

; %define 可以定义多行（带括号的宏）
%define sum(a, b) (a + b)
mov eax, sum(10, 20) ; mov eax, (10+20) -> eax = 30

特性EQU%define 可否重新定义否是 可取消定义否是（%undef） 支持参数否是（宏函数） 使用场景真正的常量（不变值）可变的配置、宏替换

### %assign - 可赋值的常量

`%assign` 用于定义可以在编译时多次修改的数值常量：

### 实例

; %assign 编译时可变常量

%assign counter 0 ; 初始为 0

section .data
; 配合 %rep 重复块使用
; 生成 table_0, table_1, ..., table_9
%rep 10
db counter ; 使用当前 counter 值
%assign counter counter + 1
%endrep

### 常量在程序中的实际应用

### 实例

; 文件路径：constants_practice.asm
; 综合使用常量提高代码可读性

; 系统调用号常量
SYS_EXIT equ 1
SYS_FORK equ 2
SYS_READ equ 3
SYS_WRITE equ 4
SYS_OPEN equ 5
SYS_CLOSE equ 6

; 文件描述符常量
STDIN equ 0
STDOUT equ 1
STDERR equ 2

; 程序常量
MAX_INPUT equ 256
EXIT_SUCCESS equ 0
EXIT_FAILURE equ 1

section .data
msg db 'Hello, RUNOOB!', 0xA
msg_len equ $ - msg

section .bss
buffer resb MAX_INPUT ; 使用常量定义缓冲区大小

section .text
global _start

_start:
; 使用常量名代替数字，代码意图一目了然
mov eax, SYS_WRITE ; 清晰：这是写操作
mov ebx, STDOUT ; 清晰：输出到屏幕
mov ecx, msg
mov edx, msg_len
int 0x80

mov eax, SYS_EXIT ; 清晰：退出程序
mov ebx, EXIT_SUCCESS ; 清晰：正常退出
int 0x80

使用常量名代替"魔法数字"是一个好习惯。三个月后，`mov eax, SYS_WRITE` 比 `mov eax, 4` 更容易理解。

### 字符串常量

NASM 支持几种定义字符串常量的方式：

### 实例

; 字符串常量的定义方式

section .data
; 直接定义（可用 EQU 表示长度）
s1 db 'hello', 0
s1_len equ $ - s1

; 带转义字符
s2 db 'Line1', 0xA, 'Line2', 0 ; 0xA 是换行符

; 使用反引号支持 C 风格转义
s3 db `hello\nworld\n`, 0 ; \n 自动替换为 0xA

; 多行字符串
s4 db 'first line', 0xA
db 'second line', 0xA
db 'third line', 0

NASM 的反引号字符串支持 `\n`（换行）、`\t`（制表符）、`\\`（反斜杠）、`\'`（单引号）等标准 C 转义序列。

---

## 汇编语言 - 算术指令

Source: https://www.runoob.com/assembly/assembly-arithmetic.html

## 汇编语言 - 算术指令

算术指令是 CPU 执行数学运算的基础。本章详细介绍 x86 架构中的加法、减法、乘法、除法以及相关进位运算指令。

### ADD - 加法指令

`ADD` 将源操作数和目标操作数相加，结果存入目标操作数。

### 实例

; 文件路径：add_demo.asm
; ADD 指令示例：计算 100 + 200 + 300

section .data
a dd 100
b dd 200
c dd 300
sum dd 0

section .text
global _start

_start:
; 方式1：寄存器 += 立即数
mov eax, [a] ; eax = 100
add eax, 200 ; eax = 100 + 200 = 300

; 方式2：寄存器 += 寄存器
mov ebx, [b] ; ebx = 200
add eax, ebx ; eax = 300 + 200 = 500

; 方式3：寄存器 += 内存
add eax, [c] ; eax = 500 + 300 = 800

; 存储结果
mov [sum], eax ; sum = 800

; 加法对标志位的影响
mov eax, 0xFFFFFFFF ; eax = 最大的 32 位无符号数
add eax, 1 ; eax = 0（溢出回绕）
; CF = 1（产生进位）
; ZF = 1（结果为 0）
; OF = 0（有符号视角无溢出）

mov eax, 1
mov ebx, 0
int 0x80

### SUB - 减法指令

`SUB` 从目标操作数中减去源操作数，结果存入目标操作数。

### 实例

; 文件路径：sub_demo.asm
; SUB 指令示例

section .data
x dd 1000
y dd 300

section .text
global _start

_start:
; 基本减法
mov eax, [x] ; eax = 1000
sub eax, [y] ; eax = 1000 - 300 = 700

; 减法对标志位的影响
mov eax, 10
sub eax, 20 ; eax = -10（即 0xFFFFFFF6）
; CF = 1（产生借位：10 < 20）
; SF = 1（结果为负）
; ZF = 0（结果非零）
; OF = 0（无符号溢出）

; 减自身：常用于清零
mov eax, 12345
sub eax, eax ; eax = 0
; ZF = 1, CF = 0
; 这是一条将寄存器清零的经典方式（比 mov eax, 0 效率高）

mov eax, 1
mov ebx, 0
int 0x80

`sub eax, eax` 是将寄存器清零的经典技巧，它只占用 2 字节，而 `mov eax, 0` 占用 5 字节。在需要极致优化体积的场景（如 shellcode）中常用。

### INC / DEC - 自增/自减指令

比 ADD/SUB 更简洁的加1和减1指令：

### 实例

; INC 和 DEC 示例

section .data
counter dd 0

section .text
global _start

_start:
mov dword [counter], 0 ; counter = 0
inc dword [counter] ; counter = 1（内存操作数）
inc dword [counter] ; counter = 2

mov ecx, 10
dec ecx ; ecx = 9
dec ecx ; ecx = 8

; INC/DEC 不影响 CF 标志位（这是与 ADD/SUB 的重要区别）
; 但会影响 ZF、SF、OF、PF

; 循环中使用 INC
mov ecx, 5
mov eax, 0
loop_inc:
inc eax ; eax 每次加 1
loop loop_inc ; 重复 5 次，eax 最终 = 5

mov ebx, eax ; 返回值 = 5
mov eax, 1
int 0x80

### MUL - 无符号乘法

`MUL` 执行无符号乘法。乘法的规则比较特殊：

操作数大小乘数被乘数（隐式）结果存放 1 字节任何 8 位寄存器或内存ALAX = AL × 操作数 2 字节任何 16 位寄存器或内存AXDX:AX = AX × 操作数 4 字节任何 32 位寄存器或内存EAXEDX:EAX = EAX × 操作数

### 实例

; 文件路径：mul_demo.asm
; MUL 无符号乘法示例

section .data
val1 dd 1000
val2 dd 2000
result_low dd 0
result_high dd 0

section .text
global _start

_start:
; 32 位乘法：EDX:EAX = EAX × 操作数
mov eax, [val1] ; eax = 1000
mul dword [val2] ; edx:eax = 1000 × 2000 = 2,000,000
; 结果 = 2,000,000 = 0x001E8480
; EAX = 0x001E8480（低 32 位）
; EDX = 0x00000000（高 32 位，因为结果没超过 32 位）

; 大数乘法演示（结果超过 32 位）
mov eax, 0xFFFFFFFF ; eax = 4,294,967,295
mov ebx, 2 ; ebx = 2
mul ebx ; edx:eax = 0xFFFFFFFF × 2
; EAX = 0xFFFFFFFE
; EDX = 0x00000001（高 32 位）
; CF = 1（结果超出 32 位）

; 保存结果
mov [result_low], eax
mov [result_high], edx

mov eax, 1
mov ebx, 0
int 0x80

### IMUL - 有符号乘法

`IMUL` 用于有符号数的乘法，有三种形式：

### 实例

; 文件路径：imul_demo.asm
; IMUL 有符号乘法示例

section .data
a dd -100
b dd 3

section .text
global _start

_start:
; 单操作数形式：同 MUL
mov eax, [a] ; eax = -100
imul dword [b] ; edx:eax = -100 × 3 = -300
; EAX = -300（0xFFFFFED4）
; EDX = 0xFFFFFFFF（符号扩展）

; 双操作数形式：reg = reg × 操作数
mov ebx, [a] ; ebx = -100
imul ebx, [b] ; ebx = -100 × 3 = -300
; 结果必须是 32 位能容纳的

; 三操作数形式：reg = 操作数1 × 立即数
imul ecx, [a], 5 ; ecx = -100 × 5 = -500

mov eax, 1
mov ebx, 0
int 0x80

`MUL` 和 `IMUL` 的主要区别是对符号的处理：`MUL` 解释为无符号数，`IMUL` 解释为有符号数。比如 0xFFFFFFFF 在 MUL 中是 4294967295，在 IMUL 中是 -1。

### DIV - 无符号除法

`DIV` 执行无符号除法，规则与 MUL 对称：

除数大小被除数（隐式）商余数 1 字节AXALAH 2 字节DX:AXAXDX 4 字节EDX:EAXEAXEDX

### 实例

; 文件路径：div_demo.asm
; DIV 无符号除法示例

section .data
dividend dd 1000
divisor dd 7
quotient dd 0
remainder dd 0

section .text
global _start

_start:
; 32 位除法：edx:eax / 除数
; 被除数要先扩展到 EDX:EAX
mov eax, [dividend] ; eax = 1000
mov edx, 0 ; edx = 0（高 32 位清零）
div dword [divisor] ; edx:eax / 7
; EAX = 142（商）
; EDX = 6（余数：1000 = 142×7 + 6）

mov [quotient], eax ; quotient = 142
mov [remainder], edx ; remainder = 6

; 字节除法示例：55 / 4
mov ax, 55 ; 被除数
mov bl, 4 ; 除数
div bl ; al = 13（商）, ah = 3（余数）

mov eax, 1
mov ebx, 0
int 0x80

做除法前务必用 `mov edx, 0` 或 `xor edx, edx` 清零 EDX！如果 EDX 中有旧数据，被除数的值就不对了。这是初学者最常见的除法 bug。

### IDIV - 有符号除法

`IDIV` 用于有符号除法。做除法前需要用 `CDQ` 指令将 EAX 符号扩展到 EDX:EAX：

### 实例

; IDIV 有符号除法示例

; 计算 -100 / 3
mov eax, -100 ; eax = -100
cdq ; 符号扩展：edx:eax = -100
; cdq 将 eax 的符号位复制到 edx 的所有位
mov ebx, 3 ; 除数
idiv ebx ; eax = -33（商）, edx = -1（余数）
; 验证：-100 = -33 × 3 + (-1)

### ADC / SBB - 带进位/借位运算

用于 大数运算（超过 32 位的数据）：

### 实例

; 文件路径：adc_sbb_demo.asm
; 64 位加法：两个 64 位数相加

section .data
; 64 位数 a = 0x00000001 FFFFFFFF
a_low dd 0xFFFFFFFF
a_high dd 0x00000001

; 64 位数 b = 0x00000000 00000005
b_low dd 0x00000005
b_high dd 0x00000000

; 结果
result_low dd 0
result_high dd 0

section .text
global _start

_start:
; 低 32 位相加
mov eax, [a_low]
add eax, [b_low] ; eax = 0xFFFFFFFF + 5 = 0x00000004
; CF = 1（产生进位！）
mov [result_low], eax ; result_low = 0x00000004

; 高 32 位加进位
mov eax, [a_high]
adc eax, [b_high] ; adc = add + CF
; eax = 1 + 0 + 1(进位) = 2
mov [result_high], eax ; result_high = 2

; 最终 64 位结果：0x00000002 00000004
; 验证：0x1FFFFFFF + 5 = 0x200000004

; SBB 减法同理（带借位）
mov eax, [a_low]
sub eax, [b_low] ; 低 32 位相减
mov [result_low], eax

mov eax, [a_high]
sbb eax, [b_high] ; sbb = sub - CF

mov eax, 1
mov ebx, 0
int 0x80

### 算术指令速查表

指令格式功能影响的标志位 ADD`add dest, src`dest = dest + srcCF, ZF, SF, OF, PF SUB`sub dest, src`dest = dest - srcCF, ZF, SF, OF, PF INC`inc dest`dest = dest + 1ZF, SF, OF, PF（不影响 CF） DEC`dec dest`dest = dest - 1ZF, SF, OF, PF（不影响 CF） MUL`mul src`无符号乘法CF, OF IMUL`imul src`有符号乘法CF, OF DIV`div src`无符号除法无定义（不定） IDIV`idiv src`有符号除法无定义（不定） ADC`adc dest, src`dest = dest + src + CFCF, ZF, SF, OF, PF SBB`sbb dest, src`dest = dest - src - CFCF, ZF, SF, OF, PF NEG`neg dest`dest = -dest（取负）CF, ZF, SF, OF, PF

---

## 汇编语言 - 逻辑指令

Source: https://www.runoob.com/assembly/assembly-logic.html

## 汇编语言 - 逻辑指令

逻辑指令在汇编编程中用于位级操作，广泛用于掩码运算、位提取、标志位设置和高效乘除运算。

### AND - 按位与

`AND` 对两个操作数逐位执行"与"运算（都为 1 时才为 1）。

### 实例

; 文件路径：and_demo.asm
; AND 指令示例

section .text
global _start

_start:
; 基本按位与
mov eax, 0x0F0F ; 0000 1111 0000 1111
and eax, 0x00FF ; 0000 0000 1111 1111
; 结果：eax = 0x000F ; 0000 0000 0000 1111

; 常用技巧1：屏蔽低 4 位（保留低 4 位）
mov eax, 0xAB ; 1010 1011
and eax, 0x0F ; 0000 1111
; eax = 0x0B ; 0000 1011

; 常用技巧2：判断奇偶性（和 1 做 AND）
mov eax, 42 ; 偶数
and eax, 1 ; eax = 0（偶数）
; 42 的二进制末尾是 0，42 & 1 = 0

mov eax, 43 ; 奇数
and eax, 1 ; eax = 1（奇数）
; 43 的二进制末尾是 1，43 & 1 = 1

; 常用技巧3：将寄存器清零
xor eax, eax ; 等同于 mov eax, 0，但更快更短

; AND 会设置标志位：CF=0, OF=0, ZF 和 SF 根据结果
mov eax, 1
mov ebx, 0
int 0x80

### OR - 按位或

`OR` 对两个操作数逐位执行"或"运算（只要有一个为 1 就是 1）。

### 实例

; OR 指令示例

; 合并标志位
mov eax, 0x0F00 ; 0000 1111 0000 0000
or eax, 0x00FF ; 0000 0000 1111 1111
; eax = 0x0FFF ; 0000 1111 1111 1111

; 设置某个位（把第 3 位设为 1）
mov eax, 0 ; 0000 0000
or eax, 0x08 ; 0000 1000
; eax = 0x8 ; 0000 1000

; 大小写转换：大写转小写
mov al, 'A' ; al = 0x41 (0100 0001)
or al, 0x20 ; 0x20 = 0010 0000
; al = 0x61 = 'a' ; 0110 0001

### NOT - 按位取反

`NOT` 将操作数的每个位取反（0 变 1，1 变 0）。

### 实例

; NOT 指令示例

mov eax, 0x0F0F0F0F ; 0000 1111 0000 1111 ...
not eax ; 1111 0000 1111 0000 ...
; eax = 0xF0F0F0F0

; NOT 不影响任何标志位（与 AND/OR/XOR 不同）

### XOR - 按位异或

`XOR` 对两个操作数逐位执行"异或"运算（不同为 1，相同为 0）。

### 实例

; 文件路径：xor_demo.asm
; XOR 指令示例

section .text
global _start

_start:
; 基本异或
mov eax, 0x0F0F ; 0000 1111 0000 1111
xor eax, 0x00FF ; 0000 0000 1111 1111
; 结果：eax = 0x0FF0 ; 0000 1111 1111 0000

; 经典技巧1：寄存器清零（比 mov reg, 0 高效）
xor eax, eax ; eax = 0，只占用 2 字节
xor ebx, ebx
xor ecx, ecx

; 经典技巧2：交换两个寄存器（不需要第三个临时寄存器）
mov eax, 100 ; eax = 100
mov ebx, 200 ; ebx = 200
xor eax, ebx ; eax = 100 xor 200
xor ebx, eax ; ebx = 200 xor (100 xor 200) = 100
xor eax, ebx ; eax = (100 xor 200) xor 100 = 200
; 现在 eax = 200, ebx = 100（交换完成！）

; 经典技巧3：简单的加密/解密
mov al, 'A' ; 原始字符 'A' = 0x41
xor al, 0x55 ; 加密：al = 'A' xor 0x55
; al 现在是某个乱码值
xor al, 0x55 ; 解密：再 xor 0x55 恢复
; al = 'A' 回来了！

mov eax, 1
mov ebx, 0
int 0x80

XOR 交换技巧虽然很酷，但在现代 CPU 上效率不如使用 `XCHG` 指令或临时寄存器。了解即可，不必执着使用。

### 移位指令

移位指令将二进制位向左或向右移动，常用于高效乘除运算（乘2、除2等）。

指令功能示例 SHL逻辑左移（左边出去，右边补0）`shl eax, 1`（乘以 2） SHR逻辑右移（右边出去，左边补0）`shr eax, 1`（无符号除以 2） SAL算术左移（同 SHL）`sal eax, 1` SAR算术右移（右边出去，左边保持符号位）`sar eax, 1`（有符号除以 2）

### 实例

; 文件路径：shift_demo.asm
; 移位指令示例

section .text
global _start

_start:
; SHL：逻辑左移 = 乘以 2 的幂
mov eax, 10 ; eax = 10 (1010)
shl eax, 1 ; eax = 20 (10100)，即 10×2
shl eax, 2 ; eax = 80 (1010000)，即 20×4

; SHR：逻辑右移 = 无符号除以 2 的幂
mov eax, 80 ; eax = 80
shr eax, 3 ; eax = 10，即 80÷8

; SAR：算术右移 = 有符号除以 2 的幂（保留符号位）
mov eax, -16 ; eax = -16 (0xFFFFFFF0)
sar eax, 2 ; eax = -4 (0xFFFFFFFC)，即 -16÷4
; SAR 对比 SHR：
; 如果 eax = 0xFFFFFFF0 (-16)，SHR 会得到 0x3FFFFFFC (很大的正数)
; 而 SAR 会得到 0xFFFFFFFC (-4)，保留了符号位

; 移出的最后一位会进入 CF 标志位
mov eax, 5 ; 0101
shr eax, 1 ; eax = 2, CF = 1（末尾的1被移出）
jc carry_was_set ; 如果 CF=1，说明原数是奇数

carry_was_set:
mov eax, 1
mov ebx, 0
int 0x80

### 循环移位指令

循环移位将移出的位重新填入另一端，形成循环：

指令功能 ROL循环左移（绕过 CF） ROR循环右移（绕过 CF） RCL带进位的循环左移（通过 CF） RCR带进位的循环右移（通过 CF）

### 实例

; 循环移位示例

; ROL：循环左移
mov al, 0x85 ; 1000 0101
rol al, 1 ; 左移 1 位：0000 1011
; CF = 1（最高位移出到 CF）
; 同时 CF 原本的 1 被移入最低位，形成循环

; ROL 实际效果：向左移 1 位，最高位同时进入 CF 和最低位
; 1000 0101 -> ROL 1 -> 0000 1011

; RCL：带进位循环左移（CF 参与循环）
clc ; 清除 CF = 0
mov al, 0x85 ; 1000 0101
rcl al, 1 ; CF 参与：CF bit7...bit0 -> CF
; 结果：0000 1010, CF = 1

; 循环移位在加密算法和位操作中常用

### 逻辑运算的实际用途

一个综合示例——使用逻辑运算实现简单的位标志系统：

### 实例

; 文件路径：bit_flags.asm
; 使用逻辑运算实现位标志系统

section .data
flags db 0 ; 8 个标志位，初始全为 0
; bit0: 是否激活
; bit1: 是否可见
; bit2: 是否需要保存
; bit3: 是否已修改

FLAG_ACTIVE equ 1 ; 0000 0001
FLAG_VISIBLE equ 2 ; 0000 0010
FLAG_NEED_SAVE equ 4 ; 0000 0100
FLAG_MODIFIED equ 8 ; 0000 1000

section .text
global _start

_start:
; 设置标志位：OR
mov al, [flags]
or al, FLAG_ACTIVE ; 设置 bit0
or al, FLAG_VISIBLE ; 设置 bit1
; al = 0000 0011 = 3
mov [flags], al

; 检查标志位：AND + TEST
test byte [flags], FLAG_ACTIVE ; 测试 bit0 是否为 1
jnz is_active ; ZF=0 表示该位是 1

is_active:
; 清除标志位：AND + NOT
mov al, [flags]
and al, ~FLAG_VISIBLE ; 清除 bit1（保留其他位不变）
; al = 0000 0001 = 1
mov [flags], al

; 切换标志位：XOR
mov al, [flags]
xor al, FLAG_MODIFIED ; 切换 bit3
; 如果原来是 0 变成 1，原来是 1 变成 0

mov eax, 1
mov ebx, 0
int 0x80

位标志操作在系统编程中极为常见。操作系统内核、设备驱动和嵌入式系统中广泛使用这种技巧来管理状态。

---

## 汇编语言 - 条件判断

Source: https://www.runoob.com/assembly/assembly-condition.html

## 汇编语言 - 条件判断

条件判断是程序控制流的基石。汇编语言通过比较指令和条件跳转指令来实现 if-else、switch 等判断逻辑。

### CMP - 比较指令

`CMP` 对两个操作数做减法运算（目标 - 源），但不保存结果，只更新标志位。

本质上 `CMP dest, src` 等同于 `SUB dest, src` 但丢弃计算结果。

### 实例

; CMP 比较指令示例

section .text
global _start

_start:
mov eax, 10
cmp eax, 10 ; eax == 10?
; ZF = 1（相等，结果为零）
; CF = 0（无借位）
; SF = 0（结果非负）

mov eax, 5
cmp eax, 10 ; eax == 10?
; ZF = 0（不等，结果非零）
; CF = 1（有借位：5 < 10）

mov eax, 20
cmp eax, 10 ; eax == 10?
; ZF = 0（不等）
; CF = 0（无借位：20 >= 10）
; SF = 0（结果非负：10 > 0）

mov eax, 1
mov ebx, 0
int 0x80

### 条件跳转指令

条件跳转指令根据标志位的状态决定是否跳转。如果条件成立，跳转到目标标签；否则继续执行下一条指令。

指令含义检查的标志位适用场景 JE / JZ相等 / 为零时跳转ZF = 1`cmp a, b` 后检查 a == b JNE / JNZ不相等 / 不为零时跳转ZF = 0`cmp a, b` 后检查 a != b JG / JNLE大于时跳转（有符号）ZF=0 且 SF=OF有符号数 a > b JGE / JNL大于等于时跳转（有符号）SF = OF有符号数 a >= b JL / JNGE小于时跳转（有符号）SF != OF有符号数 a < b JLE / JNG小于等于时跳转（有符号）ZF=1 或 SF!=OF有符号数 a <= b JA / JNBE大于时跳转（无符号）CF=0 且 ZF=0无符号数 a > b JAE / JNB大于等于时跳转（无符号）CF = 0无符号数 a >= b JB / JNAE小于时跳转（无符号）CF = 1无符号数 a < b JBE / JNA小于等于时跳转（无符号）CF=1 或 ZF=1无符号数 a <= b

很多指令有两个别名（如 JE 和 JZ），它们在机器码层面完全一样。使用哪个取决于语境：比较后用 JE/JNE，运算结果检查后用 JZ/JNZ。这样代码可读性更好。

### 单分支 IF 结构

### 实例

; 文件路径：if_demo.asm
; 实现：if (x > 10) x = 10;

section .data
x dd 15 ; 测试值
limit dd 10

section .text
global _start

_start:
mov eax, [x] ; 加载 x 到 eax
cmp eax, [limit] ; x > 10 ?
jle skip_update ; 如果 x <= 10，跳过更新

mov dword [x], 10 ; x = 10

skip_update:
; 程序继续...（这里 x 已经是 10 了）

mov eax, 1
mov ebx, 0
int 0x80

### IF-ELSE 双分支结构

### 实例

; 文件路径：if_else_demo.asm
; 实现：if (score >= 60) grade = 'P' else grade = 'F'

section .data
score dd 75 ; 考试分数
grade db 0 ; 成绩等级
PASS_SCORE equ 60

section .text
global _start

_start:
mov eax, [score] ; 加载分数
cmp eax, PASS_SCORE ; score >= 60 ?
jge pass_label ; 如果 >= 60，跳到 pass

; else 分支：不及格
mov byte [grade], 'F' ; grade = 'F'
jmp end_if ; 跳过 if 分支

pass_label:
; if 分支：及格
mov byte [grade], 'P' ; grade = 'P'

end_if:
; grade 变量现在已设置好

; 输出成绩等级
mov eax, 4
mov ebx, 1
mov ecx, grade
mov edx, 1
int 0x80

mov eax, 1
mov ebx, 0
int 0x80

运行结果：

```

$ nasm -f elf32 if_else_demo.asm -o if_else_demo.o
$ ld -m elf_i386 if_else_demo.o -o if_else_demo
$ ./if_else_demo
P

```

### IF-ELSE IF-ELSE 多分支结构

下面以分数评级系统为例，展示多分支判断的流程：

### 实例

; 文件路径：multi_branch.asm
; 分数评级：>=90 -> A, >=80 -> B, >=70 -> C, >=60 -> D, <60 -> F

section .data
score dd 85
result db 0
newline db 0xA

section .text
global _start

_start:
mov eax, [score] ; 加载分数

; 检查 >= 90
cmp eax, 90
jl check_80 ; 如果 < 90，继续检查
mov byte [result], 'A'
jmp print_result

check_80:
cmp eax, 80
jl check_70
mov byte [result], 'B'
jmp print_result

check_70:
cmp eax, 70
jl check_60
mov byte [result], 'C'
jmp print_result

check_60:
cmp eax, 60
jl fail_label
mov byte [result], 'D'
jmp print_result

fail_label:
mov byte [result], 'F'

print_result:
; 输出评级
mov eax, 4
mov ebx, 1
mov ecx, result
mov edx, 1
int 0x80

; 输出换行
mov eax, 4
mov ebx, 1
mov ecx, newline
mov edx, 1
int 0x80

mov eax, 1
mov ebx, 0
int 0x80

### TEST - 非破坏性测试指令

`TEST` 执行按位 AND 但不保存结果，只更新标志位。

`TEST` 常用于检查特定位是否为 1，或者检查寄存器是否为 0。

### 实例

; TEST 指令示例

; 检查 eax 是否为 0（比 CMP eax, 0 更高效）
test eax, eax ; eax & eax = eax，只更新 ZF
jz eax_is_zero ; 若 ZF=1，说明 eax=0

; 检查特定位
test al, 0x01 ; 检查 bit0 是否为 1
jnz bit0_is_set ; 若 bit0=1，跳转

; 检查多个位
test al, 0x03 ; 检查 bit0 和 bit1 是否有至少一个为 1
jnz some_bit_set

### SETcc - 条件设置指令

不跳转，而是根据条件将目标字节设为 1 或 0：

### 实例

; SETcc 条件设置示例

; 将 a > b 的结果存入 al
mov eax, 10
cmp eax, 5 ; 10 > 5 ?
setg al ; al = 1（大于成立）
; 如果 eax = 3，则 al = 0

; 其他 SETcc 指令：
; sete / setz -> 相等/为零
; setne / setnz -> 不等/不为零
; setl -> 小于（有符号）
; setb -> 小于（无符号）
; setg -> 大于（有符号）
; seta -> 大于（无符号）

### 完整示例：判断奇偶和正负

### 实例

; 文件路径：number_check.asm
; 判断数字的奇偶、正负

section .data
number dd -42
msg_even db 'Even', 0xA
msg_even_len equ $ - msg_even
msg_odd db 'Odd', 0xA
msg_odd_len equ $ - msg_odd
msg_pos db 'Positive', 0xA
msg_pos_len equ $ - msg_pos
msg_neg db 'Negative', 0xA
msg_neg_len equ $ - msg_neg
msg_zero db 'Zero', 0xA
msg_zero_len equ $ - msg_zero

section .text
global _start

_start:
mov eax, [number]

; 判断是否为 0
cmp eax, 0
jne check_sign
; 为零
mov eax, 4
mov ebx, 1
mov ecx, msg_zero
mov edx, msg_zero_len
int 0x80
jmp exit

check_sign:
; 判断正负
cmp eax, 0
jg is_positive ; eax > 0

; 负数
mov eax, 4
mov ebx, 1
mov ecx, msg_neg
mov edx, msg_neg_len
int 0x80
jmp check_parity

is_positive:
; 正数
mov eax, 4
mov ebx, 1
mov ecx, msg_pos
mov edx, msg_pos_len
int 0x80

check_parity:
; 判断奇偶（只需看最低位）
mov eax, [number]
test eax, 1 ; 检查 bit0
jnz is_odd

; 偶数
mov eax, 4
mov ebx, 1
mov ecx, msg_even
mov edx, msg_even_len
int 0x80
jmp exit

is_odd:
; 奇数
mov eax, 4
mov ebx, 1
mov ecx, msg_odd
mov edx, msg_odd_len
int 0x80

exit:
mov eax, 1
mov ebx, 0
int 0x80

条件跳转指令只能跳转到当前代码段内的标签，跳转范围通常受限于约 128 字节（短跳转）到 2GB（近跳转）。NASM 会自动选择合适的跳转编码。

---

## 汇编语言 - 循环结构

Source: https://www.runoob.com/assembly/assembly-loop.html

## 汇编语言 - 循环结构

循环是程序中最常见的控制结构之一。在汇编语言中，你可以用 LOOP 指令或条件跳转来实现各种循环逻辑。

### LOOP 指令

`LOOP` 是 x86 专门为循环设计的指令，它使用 ECX 作为计数器：

每次执行 LOOP 时，ECX 自动减 1，如果 ECX 不为 0 就跳转到目标标签。

### 实例

; 文件路径：loop_basic.asm
; LOOP 指令基本用法：重复 5 次输出

section .data
msg db 'runoob', 0xA
len equ $ - msg

section .text
global _start

_start:
mov ecx, 5 ; 循环计数器 = 5（循环 5 次）

repeat:
; 保存 ecx（系统调用可能修改它）
push ecx

; 输出 msg
mov eax, 4
mov ebx, 1
mov ecx, msg
mov edx, len
int 0x80

; 恢复 ecx
pop ecx
loop repeat ; ecx--; if ecx != 0 跳转到 repeat

mov eax, 1
mov ebx, 0
int 0x80

运行结果：

```

$ nasm -f elf32 loop_basic.asm -o loop_basic.o
$ ld -m elf_i386 loop_basic.o -o loop_basic
$ ./loop_basic
runoob
runoob
runoob
runoob
runoob

```

`loop` 默认使用 ECX（32位）作为计数器。在 16 位模式下使用 CX，64 位模式下使用 RCX。务必在 loop 之前正确设置 ECX 的值。

### LOOP 的变体

指令跳转条件说明 LOOPECX != 0标准循环，先减 ECX 再判断 LOOPE / LOOPZECX != 0 且 ZF = 1相等时继续循环 LOOPNE / LOOPNZECX != 0 且 ZF = 0不等时继续循环

### 实例

; LOOPE 示例：在数组中查找第一个非零值

section .data
array db 0, 0, 0, 5, 0, 0 ; 前三个是 0，第四个是 5

section .text
global _start

_start:
mov ecx, 6 ; 数组长度
mov esi, array - 1 ; 指向数组前一个位置

find_nonzero:
inc esi ; 移动到下一个元素
cmp byte [esi], 0 ; 当前元素 == 0 ?
loope find_nonzero ; 如果是 0 且 ecx > 0，继续循环
; 循环结束：找到了第一个非零元素，或遍历结束

; esi 现在指向第一个非零元素的地址
; ecx 剩余未比较的元素个数

mov eax, 1
mov ebx, 0
int 0x80

### 条件循环（WHILE 循环）

用条件跳转实现 while 循环：先判断条件，再执行循环体。

### 实例

; 文件路径：while_loop.asm
; 实现：while (x < 100) x = x * 2;

section .data
x dd 1
limit dd 100

section .text
global _start

_start:
; while 循环
while_start:
mov eax, [x] ; 加载 x
cmp eax, [limit] ; x < 100 ?
jge while_end ; 如果 x >= 100，结束循环

shl eax, 1 ; x = x * 2
mov [x], eax ; 存回 x

jmp while_start ; 回到循环开始，重新判断

while_end:
; x 现在是 128（1->2->4->8->16->32->64->128，128 >= 100 停止）

mov eax, 1
mov ebx, [x] ; 返回 x 作为退出码
int 0x80

### DO-WHILE 循环

先执行循环体，再判断条件：

### 实例

; DO-WHILE 循环：至少执行一次

section .data
counter dd 0

section .text
global _start

_start:
mov dword [counter], 0

do_loop:
; 循环体：counter++（至少执行一次）
inc dword [counter]

; 判断条件
cmp dword [counter], 5
jl do_loop ; counter < 5 时继续
; counter 最终 = 5

mov eax, 1
mov ebx, 0
int 0x80

### 嵌套循环

嵌套循环需要注意外层计数器的保存：

### 实例

; 文件路径：nested_loop.asm
; 嵌套循环示例：打印乘法表（3×3）

section .data
space db ' '
newline db 0xA
; 用于存放转换后的数字字符（2位 + 空格 + null）
num_str db ' ', 0

section .text
global _start

_start:
mov ecx, 3 ; 外层循环计数器（行数）

outer_loop:
push ecx ; ★ 保存外层计数器！（重要）
mov ecx, 3 ; 内层循环计数器（列数）

inner_loop:
push ecx ; 保存内层计数器

; 计算乘积 = 外层行号(4-当前ecx) × 内层列号
; 此处省略具体数字转换输出，仅演示结构
; 实际输出：行号×列号

pop ecx ; 恢复内层计数器
loop inner_loop ; 内层循环

; 输出换行
mov eax, 4
mov ebx, 1
mov ecx, newline
mov edx, 1
int 0x80

pop ecx ; ★ 恢复外层计数器
loop outer_loop ; 外层循环

mov eax, 1
mov ebx, 0
int 0x80

嵌套循环中最常见的错误是忘记用栈保存外层计数器的值。LOOP 会修改 ECX，当你在内层循环中重新设置 ECX 时，外层的计数值就丢失了。解决方法是每次进入内层循环前 `push ecx`，离开后 `pop ecx`。

### 循环优化技巧

一些提高循环效率的方法：

### 实例

; 循环优化对比

; 方式A：使用 LOOP 指令（较慢，但不明显）
mov ecx, 1000
loop_a:
add eax, [ebx]
add ebx, 4
loop loop_a

; 方式B：使用条件跳转（手工计数）（通常更快）
mov ecx, 1000
loop_b:
add eax, [ebx]
add ebx, 4
dec ecx
jnz loop_b

; 方式C：循环展开（最快，但代码体积大）
; 将 1000 次循环展开为每次处理 4 个元素
mov ecx, 250 ; 1000 / 4 = 250
loop_unrolled:
add eax, [ebx]
add eax, [ebx + 4]
add eax, [ebx + 8]
add eax, [ebx + 12]
add ebx, 16
dec ecx
jnz loop_unrolled

; 方式D：倒计数（减少比较指令）
mov ecx, 1000
lea ebx, [array + 1000*4 - 4] ; 从数组末尾开始
loop_reverse:
add eax, [ebx]
sub ebx, 4
dec ecx
jnz loop_reverse

### 完整示例：计算 1 到 100 的和

### 实例

; 文件路径：sum_1_to_100.asm
; 计算 1+2+...+100 = 5050

section .data
msg db 'Sum of 1 to 100 is: '
msg_len equ $ - msg
newline db 0xA

section .bss
result_str resb 12 ; 存放转换后的数字字符串

section .text
global _start

_start:
; 计算累加和
mov ecx, 100 ; 循环 100 次
mov eax, 0 ; 累加和初始为 0
mov ebx, 1 ; 当前要加的数

sum_loop:
add eax, ebx ; eax += ebx
inc ebx ; 下一个数
loop sum_loop ; 继续循环
; eax = 5050

; 输出提示文字
push eax ; 保存结果
mov eax, 4
mov ebx, 1
mov ecx, msg
mov edx, msg_len
int 0x80
pop eax ; 恢复结果

; 将数字转为字符串（简单版：直接输出整数值）
; 这里简化处理，将 eax 的低字节转为字符
add al, '0' ; 不准确，仅演示
mov [result_str], al

; 输出结果
mov eax, 4
mov ebx, 1
mov ecx, result_str
mov edx, 12
int 0x80

mov eax, 4
mov ebx, 1
mov ecx, newline
mov edx, 1
int 0x80

mov eax, 1
mov ebx, 0
int 0x80

数字转字符串是汇编中的一个常见问题。上面的简单方法只在数字是 0-9 时有效。完整的数字转字符串算法将在"数字处理"章节详细讲解。

---

## 汇编语言 - 数字处理

Source: https://www.runoob.com/assembly/assembly-numbers.html

## 汇编语言 - 数字处理

在汇编语言中处理数字涉及到进制转换、ASCII 转换和算术运算的结合。本章详细讲解如何在程序中输入和输出各种进制的数字。

### 数字的表示方式

计算机内部所有数字都以二进制存储，但程序中可以用不同进制表示：

进制基数NASM 写法数字范围（32位） 二进制2`0b10101010` 或 `10101010b`0 和 1 八进制8`0o52` 或 `52o`0-7 十进制10`42`0-9 十六进制16`0x2A` 或 `2Ah`0-9, A-F

### ASCII 与数字的关系

在屏幕上输入和输出的都是 ASCII 字符，而非二进制数值。

理解字符与数值的转换是汇编数字处理的核心：

字符ASCII 值对应的数值 '0'0x30（48）0 '1'0x31（49）1 '9'0x39（57）9 'A'0x41（65）10（十六进制） 'F'0x46（70）15（十六进制） 'a'0x61（97）10（十六进制）

字符转数字：`数值 = ASCII码 - '0'`（即 `ASCII码 - 0x30`）

数字转字符：`ASCII码 = 数值 + '0'`（即 `数值 + 0x30`）

### 输入数字（字符串转数值）

用户输入的是字符序列，需要转换为二进制数值才能用于计算：

### 实例

; 文件路径：string_to_int.asm
; 将十进制数字字符串转换为整数值

section .data
input_str db '12345', 0 ; 输入字符串 "12345"
input_len dd 5 ; 字符串长度

section .bss
result_val resd 1 ; 存放转换后的数值

section .text
global _start

_start:
; 转换算法：result = 0
; result = result * 10 + (当前字符 - '0')
mov esi, input_str ; esi 指向输入字符串
mov ecx, [input_len] ; ecx = 循环次数
mov eax, 0 ; eax = 累加结果

convert_loop:
mov ebx, 10
mul ebx ; eax = eax * 10
; mul ebx: edx:eax = eax * ebx

mov bl, [esi] ; 读取当前字符
sub bl, '0' ; 转为数值：字符 - '0'
movzx ebx, bl ; 零扩展 ebx（bl -> ebx）
add eax, ebx ; eax = eax + 当前位数值

inc esi ; 移动到下一个字符
loop convert_loop
; eax 现在 = 12345

mov [result_val], eax ; 保存结果

mov eax, 1
mov ebx, 0
int 0x80

算法解析：以 "12345" 为例：
第1轮：result = 0×10 + 1 = 1
第2轮：result = 1×10 + 2 = 12
第3轮：result = 12×10 + 3 = 123
...最终得到 12345。

### 输出数字（数值转字符串）

将二进制数值转为可显示的十进制字符串：

### 实例

; 文件路径：int_to_string.asm
; 将整数值转换为十进制字符串输出

section .data
number dd 12345 ; 要输出的数字
newline db 0xA

section .bss
output_buf resb 12 ; 输出缓冲区（最大 32 位整数是 10 位 + 符号 + null）

section .text
global _start

_start:
mov eax, [number] ; 加载数字
mov edi, output_buf + 11 ; edi 指向缓冲区末尾
mov byte [edi], 0 ; null 终止符（方便调试）
dec edi

mov ebx, 10 ; 除数 = 10
mov ecx, 0 ; 位数计数器

convert_digit:
mov edx, 0 ; 清零 edx（div 需要 edx:eax）
div ebx ; eax = eax/10, edx = eax%10
add dl, '0' ; 余数转字符
mov [edi], dl ; 存入缓冲区（从后往前）
dec edi ; 缓冲区指针前移
inc ecx ; 位数 +1
cmp eax, 0 ; 商是否为 0？
jne convert_digit ; 否，继续循环

; 现在 edi+1 指向第一个有效数字字符
inc edi ; 修正指针，指向字符串开头

; 计算有效字符长度
; output_buf + 11 - edi 即有效长度

; 输出数字
mov eax, 4
mov ebx, 1
mov ecx, edi ; 指向转换后的字符串
; 计算长度
mov edx, output_buf + 11
sub edx, edi ; edx = 缓冲区末尾 - 字符串开头
int 0x80

; 输出换行
mov eax, 4
mov ebx, 1
mov ecx, newline
mov edx, 1
int 0x80

mov eax, 1
mov ebx, 0
int 0x80

运行结果：

```

$ nasm -f elf32 int_to_string.asm -o int_to_string.o
$ ld -m elf_i386 int_to_string.o -o int_to_string
$ ./int_to_string
12345

```

### 十六进制输入输出

处理十六进制需要同时处理 0-9 和 A-F/a-f 的转换：

### 实例

; 文件路径：hex_output.asm
; 将数值以十六进制格式输出

section .data
number dd 0x1A2B3C4D ; 要输出的值
hex_prefix db '0x'
hex_prefix_len equ $ - hex_prefix
newline db 0xA

section .bss
hex_buf resb 10 ; 8 个十六进制位 + 前缀 + null

section .text
global _start

_start:
; 输出前缀 "0x"
mov eax, 4
mov ebx, 1
mov ecx, hex_prefix
mov edx, hex_prefix_len
int 0x80

mov eax, [number] ; 要转换的数值
mov edi, hex_buf + 8 ; 缓冲区末尾（8 个十六进制位）
mov ecx, 8 ; 循环 8 次（32 位 = 8 个十六进制位）

hex_loop:
mov edx, 0 ; 准备除法
mov ebx, 16 ; 除以 16
div ebx ; 现在：eax 除以 ebx
; 注意：这里 div ebx 实际是 edx:eax / ebx
; 余数在 edx（0-15），商在 eax
mov ebx, eax ; 保存商

; 将余数转为十六进制字符
mov al, dl
cmp al, 10
jl digit_09 ; 0-9
add al, 'A' - 10 ; A-F
jmp store_char

digit_09:
add al, '0' ; 0-9

store_char:
dec edi
mov [edi], al
mov eax, ebx ; 恢复商
loop hex_loop

; 移除了loop，直接使用递减方式

; 输出十六进制数字（8 位）
mov eax, 4
mov ebx, 1
mov ecx, hex_buf
mov edx, 8
int 0x80

; 输出换行
mov eax, 4
mov ebx, 1
mov ecx, newline
mov edx, 1
int 0x80

mov eax, 1
mov ebx, 0
int 0x80

### 完整示例：简单的加法计算器

从键盘输入两个一位数，计算它们的和并输出：

### 实例

; 文件路径：simple_calc.asm
; 输入两个一位数（0-9），计算和并输出

section .data
prompt1 db 'Enter first number (0-9): '
prompt1_len equ $ - prompt1
prompt2 db 'Enter second number (0-9): '
prompt2_len equ $ - prompt2
result_msg db 'Sum = '
result_msg_len equ $ - result_msg
newline db 0xA

section .bss
num1 resb 2 ; 第一个数的输入（1位+换行）
num2 resb 2 ; 第二个数的输入
result_str resb 3 ; 结果字符串（最多2位+换行）

section .text
global _start

_start:
; 提示输入第一个数
mov eax, 4
mov ebx, 1
mov ecx, prompt1
mov edx, prompt1_len
int 0x80

; 读取第一个数
mov eax, 3
mov ebx, 0
mov ecx, num1
mov edx, 2
int 0x80

; 提示输入第二个数
mov eax, 4
mov ebx, 1
mov ecx, prompt2
mov edx, prompt2_len
int 0x80

; 读取第二个数
mov eax, 3
mov ebx, 0
mov ecx, num2
mov edx, 2
int 0x80

; 计算和
mov al, [num1] ; 读取第一个数（ASCII 字符）
sub al, '0' ; 转为数值
mov bl, [num2] ; 读取第二个数
sub bl, '0' ; 转为数值
add al, bl ; 求和

; 将结果转为字符串
mov ah, 0 ; 准备除法
mov bl, 10 ; 除以 10
div bl ; al = 十位数, ah = 个位数

; 保存结果
add al, '0' ; 十位转字符
mov [result_str], al ; 存十位
add ah, '0' ; 个位转字符
mov [result_str + 1], ah ; 存个位
mov byte [result_str + 2], 0xA ; 换行

; 输出结果信息
mov eax, 4
mov ebx, 1
mov ecx, result_msg
mov edx, result_msg_len
int 0x80

; 输出计算结果（可能 1 或 2 位）
mov eax, 4
mov ebx, 1
mov ecx, result_str
cmp byte [result_str], '0'
je output_one_digit ; 如果十位是 '0'，只输出个位
mov edx, 3 ; 2 位 + 换行
jmp do_output

output_one_digit:
inc ecx ; 跳过十位的 '0'
mov edx, 2 ; 1 位 + 换行

do_output:
int 0x80

mov eax, 1
mov ebx, 0
int 0x80

运行结果：

```

$ nasm -f elf32 simple_calc.asm -o simple_calc.o
$ ld -m elf_i386 simple_calc.o -o simple_calc
$ ./simple_calc
Enter first number (0-9): 7
Enter second number (0-9): 8
Sum = 15

```

在汇编中处理数字输入输出需要对 ASCII 编码和进制转换有清晰的理解。数字转字符串（int to string）和字符串转数字（string to int）是最基本的两个工具函数，建议写好后保存为模板，后续可直接复用。

---

## 汇编语言 - 字符串处理

Source: https://www.runoob.com/assembly/assembly-string.html

## 汇编语言 - 字符串处理

字符串处理是编程中常见的需求。

在汇编中处理字符串意味着直接操作内存中的字节序列，虽然更底层但也更灵活。

### 字符串的定义和存储

在 NASM 中，字符串本质上是字节序列，可以用 DB 伪指令定义：

### 实例

; 字符串定义方式

section .data
; 方式1：标准字符串
str1 db 'Hello, runoob!', 0 ; C 风格：以 null 结尾

; 方式2：带换行符
str2 db 'Line1', 0xA, 'Line2', 0xA

; 方式3：反引号支持转义
str3 db `hello\nworld\n`, 0 ; 自动转换转义序列

; 方式4：逐个字符定义
str4 db 'A', 'B', 'C', 'D', 0

; 方式5：使用 dup 生成重复字符
border db 40 dup('-') ; 40 个减号

; 字符串长度计算（编译时）
str1_len equ $ - str1 ; 包含结尾 null

### 计算字符串长度

有两种方式获取字符串长度：编译时计算（EQU）和运行时计算（遍历查找 null）：

### 实例

; 文件路径：strlen_demo.asm
; 两种计算字符串长度的方式

section .data
; 方式A：编译时计算（适用于常量字符串）
msg1 db 'Hello, RUNOOB!', 0xA
msg1_len equ $ - msg1 ; 编译时自动计算

; 方式B：null 结尾的字符串（运行时计算）
msg2 db 'Find my length', 0 ; 以 0 结尾

section .text
global _start

_start:
; 方式A：直接使用编译时计算的长度
mov eax, 4
mov ebx, 1
mov ecx, msg1
mov edx, msg1_len ; 直接使用常量
int 0x80

; 方式B：运行时计算 null 结尾字符串的长度
mov esi, msg2 ; esi 指向字符串起始
mov ecx, 0 ; 计数器

strlen_loop:
cmp byte [esi], 0 ; 当前字节是 null？
je strlen_done ; 是，结束
inc ecx ; 计数 +1
inc esi ; 指针 +1
jmp strlen_loop

strlen_done:
; ecx 现在存放字符串长度（不含 null）
mov eax, 4
mov ebx, 1
mov ecx, msg2
mov edx, ecx ; 使用计算出的长度
; 这里有个问题：ecx 被覆盖了，应该先保存
; 正确做法：push ecx 再 pop 到 edx

mov eax, 1
mov ebx, 0
int 0x80

### 字符串复制

使用循环逐字节复制，或使用 x86 的字符串操作指令 `MOVSB`：

### 实例

; 文件路径：strcpy_demo.asm
; 字符串复制两种实现方式

section .data
src db 'runoob source string', 0
src_len equ $ - src

section .bss
dest_manual resb 64 ; 手动复制目标
dest_fast resb 64 ; 快速复制目标

section .text
global _start

_start:
; 方式A：手动逐字节复制
mov esi, src ; 源地址
mov edi, dest_manual ; 目标地址
mov ecx, src_len ; 字节数

copy_loop:
mov al, [esi] ; 读取一字节
mov [edi], al ; 写入一字节
inc esi ; 源指针++
inc edi ; 目标指针++
loop copy_loop

; 方式B：使用字符串操作指令（更快）
cld ; 清除方向标志（DF=0，正向复制）
mov esi, src ; 源地址（ESI）
mov edi, dest_fast ; 目标地址（EDI）
mov ecx, src_len ; 字节数（ECX）
rep movsb ; 重复执行 movsb ecx 次
; rep movsb：while(ecx>0) { [edi]=[esi]; esi++; edi++; ecx--; }

; 验证：输出两个复制结果
mov eax, 4
mov ebx, 1
mov ecx, dest_manual
mov edx, src_len
int 0x80

mov eax, 4
mov ebx, 1
mov ecx, dest_fast
mov edx, src_len
int 0x80

mov eax, 1
mov ebx, 0
int 0x80

`REP MOVSB` 是 x86 最经典的字符串复制方式。但在现代 CPU 上，手动复制循环配合循环展开可能比 REP MOVSB 更快，因为现代 CPU 对简单操作有更好的流水线优化。

### 字符串比较

使用 `CMPSB` 配合 `REPE`（重复直到不等）逐字节比较：

### 实例

; 文件路径：strcmp_demo.asm
; 比较两个字符串是否相等

section .data
str_a db 'runoob', 0
str_b db 'runoob', 0
str_c db 'RUNOOB', 0
eq_msg db 'Strings are equal', 0xA
eq_len equ $ - eq_msg
ne_msg db 'Strings are NOT equal', 0xA
ne_len equ $ - ne_msg

section .text
global _start

_start:
cld ; 正向比较
mov esi, str_a ; 第一个字符串
mov edi, str_b ; 第二个字符串
mov ecx, 7 ; 最多比较 7 字节（含 null）

repe cmpsb ; 重复比较直到不等或 ecx=0
; repe: 如果 ZF=1（相等）且 ecx>0 则继续
; cmpsb: 比较 [esi] 和 [edi]，然后 esi++, edi++

je strings_equal ; 如果 ZF=1，所有字节都相等

; 不等
mov eax, 4
mov ebx, 1
mov ecx, ne_msg
mov edx, ne_len
int 0x80
jmp compare_next

strings_equal:
mov eax, 4
mov ebx, 1
mov ecx, eq_msg
mov edx, eq_len
int 0x80

compare_next:
; 比较 str_a 和 str_c（大小写不同）
cld
mov esi, str_a
mov edi, str_c
mov ecx, 7
repe cmpsb
jne not_equal_2

mov eax, 4
mov ebx, 1
mov ecx, eq_msg
mov edx, eq_len
int 0x80
jmp exit

not_equal_2:
mov eax, 4
mov ebx, 1
mov ecx, ne_msg
mov edx, ne_len
int 0x80

exit:
mov eax, 1
mov ebx, 0
int 0x80

### 字符串操作指令汇总

指令功能使用的寄存器 MOVSB复制字节：[EDI] = [ESI]ESI=源, EDI=目标, ECX=次数, DF=方向 MOVSW复制字（2 字节）同上 MOVSD复制双字（4 字节）同上 STOSB存字节：[EDI] = ALEDI=目标, AL=值, ECX=次数 LODSB取字节：AL = [ESI]ESI=源 CMPSB比较字节：[ESI] - [EDI]ESI=源, EDI=目标, ECX=次数 SCASB扫描字节：AL - [EDI]EDI=目标, AL=查找值, ECX=次数

### 大小写转换示例

### 实例

; 文件路径：case_convert.asm
; 将字符串中的小写字母转为大写

section .data
msg db 'Hello, runoob! Welcome to Assembly.', 0xA
len equ $ - msg

section .text
global _start

_start:
; 输出原始字符串
mov eax, 4
mov ebx, 1
mov ecx, msg
mov edx, len
int 0x80

; 转换：遍历字符串，小写 -> 大写
mov esi, msg ; 指向字符串开头
mov ecx, len ; 循环次数

convert_loop:
mov al, [esi] ; 读取字符
cmp al, 'a' ; 是否 >= 'a'
jb next_char ; 小于 'a'，跳过
cmp al, 'z' ; 是否 <= 'z'
ja next_char ; 大于 'z'，跳过

; 小写转大写：'a'(97) - 'A'(65) = 32
sub al, 32 ; 转为大写
mov [esi], al ; 写回

next_char:
inc esi
loop convert_loop

; 输出转换后的字符串
mov eax, 4
mov ebx, 1
mov ecx, msg
mov edx, len
int 0x80

mov eax, 1
mov ebx, 0
int 0x80

运行结果：

```

$ nasm -f elf32 case_convert.asm -o case_convert.o
$ ld -m elf_i386 case_convert.o -o case_convert
$ ./case_convert
Hello, runoob! Welcome to Assembly.
HELLO, RUNOOB! WELCOME TO ASSEMBLY.

```

方向标志 DF 决定了字符串操作的方向：DF=0 时 ESI/EDI 递增（正向），DF=1 时递减（反向）。用 `CLD` 清零 DF，用 `STD` 置位 DF。在调用 C 函数或系统调用前应该确保 DF=0（cdecl 要求）。

---

## 汇编语言 - 数组

Source: https://www.runoob.com/assembly/assembly-array.html

## 汇编语言 - 数组

数组是连续存放的同类型数据集合。在汇编中，数组就是内存中一段连续的空间，通过基址加偏移的方式来访问每个元素。

### 一维数组定义

### 实例

; 数组的定义方式

section .data
; 方式1：逐个列出元素
arr1 dd 10, 20, 30, 40, 50 ; 5 个双字元素的数组

; 方式2：用 dup 初始化
arr2 dd 10 dup(0) ; 10 个双字，全为 0

; 方式3：字节数组
bytes db 1, 2, 3, 4, 5, 6 ; 6 个字节

; 方式4：字符数组（字符串）
chars db 'runoob', 0 ; 7 字节

; 数组长度计算（编译时）
arr1_len equ ($ - arr1) / 4 ; 双字数组的元素个数
arr2_len equ ($ - arr2) / 4
bytes_len equ ($ - bytes) ; 字节数组，不需除法

section .bss
; 未初始化的数组
buffer resd 100 ; 预留 100 个双字的空间

### 数组元素访问

访问数组元素使用 基址 + 索引 × 元素大小 的寻址方式：

### 实例

; 文件路径：array_access.asm
; 数组元素的读取和写入

section .data
nums dd 100, 200, 300, 400, 500 ; 5 个元素
nums_len equ ($ - nums) / 4

section .text
global _start

_start:
; 访问第 0 个元素（下标 0）
mov eax, [nums] ; eax = 100

; 访问第 2 个元素（下标 2）
mov eax, [nums + 2 * 4] ; eax = nums[2] = 300
; 2 * 4 = 8，从 nums 偏移 8 字节

; 使用寄存器作为下标
mov esi, 3 ; 下标 = 3
mov eax, [nums + esi * 4] ; eax = nums[3] = 400

; 修改数组元素
mov dword [nums + 4], 250 ; nums[1] = 250
; 数组中现在：100, 250, 300, 400, 500

; 使用 EBX 作为基址寄存器
mov ebx, nums ; ebx = 数组基址
mov eax, [ebx + 4 * 4] ; eax = nums[4] = 500

mov eax, 1
mov ebx, 0
int 0x80

注意：`[nums + esi * 4]` 中的乘以 4 是因为每个元素是 4 字节（双字）。如果是字节数组（db），直接写成 `[arr + esi]`；字数组（dw），写成 `[arr + esi * 2]`。

### 数组遍历

### 实例

; 文件路径：array_traverse.asm
; 遍历数组并计算总和

section .data
numbers dd 5, 10, 15, 20, 25, 30, 35, 40, 45, 50
count equ ($ - numbers) / 4 ; 元素个数

section .text
global _start

_start:
mov ecx, count ; 循环次数
mov esi, 0 ; 下标（从 0 开始）
mov eax, 0 ; 累加和

sum_loop:
add eax, [numbers + esi * 4] ; 累加数组元素
inc esi ; 下标 +1
loop sum_loop
; eax = 5+10+15+...+50 = 275

; 另一种遍历方式：使用指针
mov ecx, count
mov ebx, numbers ; ebx 指向数组起始
mov eax, 0 ; 累加和

sum_loop2:
add eax, [ebx] ; 累加当前元素
add ebx, 4 ; 指针移动到下一个元素
loop sum_loop2
; eax = 275

mov ebx, eax ; 退出码 = 累加和
mov eax, 1
int 0x80

### 数组查找

### 实例

; 文件路径：array_search.asm
; 在数组中查找指定值

section .data
data dd 12, 45, 67, 23, 89, 34, 56, 78, 90, 11
data_len equ ($ - data) / 4
target dd 23 ; 要查找的值
found_msg db 'Found at index: ', 0
found_len equ $ - found_msg
notfound_msg db 'Not found (runoob)', 0xA
notfound_len equ $ - notfound_msg
newline db 0xA

section .text
global _start

_start:
mov ecx, data_len ; 循环次数
mov esi, 0 ; 当前下标

search_loop:
mov eax, [data + esi * 4] ; 加载当前元素
cmp eax, [target] ; 是否等于目标值
je found ; 找到了

inc esi ; 下标 +1
loop search_loop

; 没找到
mov eax, 4
mov ebx, 1
mov ecx, notfound_msg
mov edx, notfound_len
int 0x80
jmp exit

found:
; 找到了（esi 是下标）
mov eax, 4
mov ebx, 1
mov ecx, found_msg
mov edx, found_len
int 0x80

; 将下标转为 ASCII 字符输出
add esi, '0' ; 单数字下标转字符
push esi ; 压栈作为临时存储
mov eax, 4
mov ebx, 1
mov ecx, esp ; 栈顶地址
mov edx, 1
int 0x80
pop esi

; 输出换行
mov eax, 4
mov ebx, 1
mov ecx, newline
mov edx, 1
int 0x80

exit:
mov eax, 1
mov ebx, 0
int 0x80

### 二维数组

二维数组在内存中按行展开为一维存储。访问 `arr[i][j]` 的地址公式为：

```

地址 = 基址 + (i*列数 + j) * 元素大小

```

### 实例

; 文件路径：2d_array.asm
; 二维数组的定义和访问

section .data
; 3 行 4 列的二维数组
matrix dd 1, 2, 3, 4
dd 5, 6, 7, 8
dd 9, 10, 11, 12

rows equ 3
cols equ 4
elem_size equ 4 ; 双字 = 4 字节

section .text
global _start

_start:
; 访问 matrix[1][2] = 7（第2行第3列）
; 地址 = matrix + (1*4 + 2) * 4 = matrix + 24

mov eax, [matrix + (1 * cols + 2) * elem_size]
; eax = 7

; 使用寄存器动态计算（假设 i=2, j=1）
; matrix[2][1] = 10（第3行第2列）
mov esi, 2 ; 行 i = 2
mov edi, 1 ; 列 j = 1

mov eax, cols ; 列数
mul esi ; eax = i * cols = 2*4 = 8
add eax, edi ; eax = i*cols + j = 8+1 = 9
; eax = eax * elem_size
; 注意：MUL 的结果在 EAX 中，这里直接使用
mov eax, [matrix + eax * elem_size]
; eax = 10

; 遍历二维数组所有元素
mov ecx, rows * cols ; 总元素数 = 12
mov esi, 0 ; 下标
mov ebx, 0 ; 累加和

traverse:
add ebx, [matrix + esi * elem_size]
inc esi
loop traverse
; ebx = 1+2+3+...+12 = 78

mov eax, 1
int 0x80

### 冒泡排序完整示例

### 实例

; 文件路径：bubble_sort.asm
; 冒泡排序算法

section .data
array dd 64, 34, 25, 12, 22, 11, 90, 78
array_len equ ($ - array) / 4

section .text
global _start

_start:
mov ecx, array_len ; 外层循环：n 次
dec ecx ; 外层只需 n-1 次

outer_loop:
push ecx ; 保存外层计数器

mov esi, 0 ; 内层下标从 0 开始
mov ecx, array_len - 1 ; 内层循环次数

inner_loop:
mov eax, [array + esi * 4] ; a[j]
mov ebx, [array + esi * 4 + 4] ; a[j+1]
cmp eax, ebx ; a[j] > a[j+1] ?
jle no_swap ; 否，不交换

; 交换 a[j] 和 a[j+1]
mov [array + esi * 4], ebx
mov [array + esi * 4 + 4], eax

no_swap:
inc esi
loop inner_loop

pop ecx
loop outer_loop
; 数组现在已排序：11, 12, 22, 25, 34, 64, 78, 90

mov eax, 1
mov ebx, 0
int 0x80

汇编语言中的数组没有任何边界检查。访问越界的索引不会报错，而是会静默地读写邻近内存中的数据，这可能导致难以调试的 bug。务必自己确保索引在合法范围内。

---

## 汇编语言 - 过程（子程序）

Source: https://www.runoob.com/assembly/assembly-procedure.html

## 汇编语言 - 过程（子程序）

过程（Procedure）是汇编语言中实现函数/子程序复用的机制，它能让代码结构化、可重用。理解过程和栈的关系是汇编编程的重要里程碑。

### 什么是过程

过程 是一段可以被多次调用的代码块，类似于高级语言中的函数（Function）或子程序（Subroutine）。

在汇编中，过程通过 `CALL` 指令调用，通过 `RET` 指令返回。`CALL` 会将返回地址压入栈中，`RET` 从栈中弹出返回地址并跳转回去。

### 实例

; 文件路径：simple_proc.asm
; 最简单的过程调用示例

section .data
msg db 'Hello, runoob!', 0xA
len equ $ - msg

section .text
global _start

_start:
call print_message ; 调用过程（将下条指令地址压栈）
call print_message ; 再次调用

mov eax, 1
mov ebx, 0
int 0x80

; 过程：打印消息
print_message:
push eax ; 保存要修改的寄存器
push ebx
push ecx
push edx

mov eax, 4
mov ebx, 1
mov ecx, msg
mov edx, len
int 0x80

pop edx ; 恢复寄存器（按相反顺序）
pop ecx
pop ebx
pop eax
ret ; 返回调用处（从栈中弹出返回地址）

在过程内部，应当保存和恢复所有被修改的寄存器（除了 EAX 当返回值时），这是汇编编程的基本礼仪。不这样做会导致调用方的寄存器值被意外修改。

### CALL 和 RET 原理

`CALL` 和 `RET` 配对工作，依赖栈来管理返回地址：

指令实际执行的操作 `CALL label``push eip`（保存下一条指令地址）+ `jmp label` `RET``pop eip`（从栈中弹出地址并跳转）

### 通过寄存器传递参数

在调用过程前将参数放入约定好的寄存器：

### 实例

; 文件路径：proc_params_reg.asm
; 通过寄存器传递参数

section .data
newline db 0xA

section .text
global _start

_start:
; 调用 add_two 过程：计算 10 + 20
mov eax, 10 ; 第一个参数放在 eax
mov ebx, 20 ; 第二个参数放在 ebx
call add_two ; 调用过程
; 返回值在 eax 中 = 30

mov ebx, eax ; 退出码 = 30
mov eax, 1
int 0x80

; 过程：计算 eax + ebx，结果返回在 eax 中
add_two:
add eax, ebx ; eax = eax + ebx
ret

### 通过栈传递参数

栈传递参数更灵活，是 C 语言等高级语言的标准方式：

### 实例

; 文件路径：proc_params_stack.asm
; 通过栈传递参数（cdecl 风格）

section .data
msg db 'Sum is: '
msg_len equ $ - msg
newline db 0xA

section .bss
result_buf resb 4

section .text
global _start

_start:
; 调用 sum 过程：计算 100 + 200
push dword 200 ; 压入第 2 个参数
push dword 100 ; 压入第 1 个参数
call sum ; 调用过程
add esp, 8 ; ★ 调用者清理栈（cdecl 约定）
; 返回值在 eax 中 = 300

mov ebx, eax
mov eax, 1
int 0x80

; 过程：sum(a, b) = a + b
sum:
push ebp ; ★ 保存旧的基址指针
mov ebp, esp ; ★ 设置新的栈帧基址

; 栈布局（从高地址到低地址）：
; [ebp+12] = 参数 b（200）
; [ebp+8] = 参数 a（100）
; [ebp+4] = 返回地址
; [ebp] = 旧的 ebp
; [ebp-4] = 局部变量空间

mov eax, [ebp + 8] ; 获取第 1 个参数（a）
add eax, [ebp + 12] ; 加上第 2 个参数（b）

pop ebp ; ★ 恢复旧的基址指针
ret ; 返回

栈帧结构图解：

下面是对应的文本示意：

```

高地址
+------------------+
| 参数2 (200) | <-- [ebp + 12]
+------------------+
| 参数1 (100) | <-- [ebp + 8]
+------------------+
| 返回地址 | <-- [ebp + 4]
+------------------+
| 旧的 EBP | <-- [ebp] (当前 EBP 指向这里)
+------------------+
| 局部变量区 | <-- [ebp - 4], [ebp - 8], ...
+------------------+ <-- ESP 指向这里
低地址

```

函数序言（Prologue）`push ebp; mov ebp, esp` 和尾声（Epilogue）`pop ebp; ret`（或 `leave; ret`）是标准的栈帧管理模板，几乎所有汇编函数都以这个结构开始和结束。

### 过程的返回值

返回值通常存放在 EAX 寄存器中（32 位值）或 EDX:EAX（64 位值）：

### 实例

; 返回值示例

; 返回 int
call get_answer
; eax = 42

; 返回 64 位值（如 long long）
call get_big_value
; edx:eax = 64 位结果

get_answer:
mov eax, 42 ; 返回值放在 eax 中
ret

get_big_value:
mov eax, 0x00000001 ; 低 32 位
mov edx, 0x00000000 ; 高 32 位
ret

### 局部变量

局部变量使用栈空间，在过程序言中通过减小 ESP 来分配：

### 实例

; 文件路径：local_vars.asm
; 在过程中使用局部变量

section .text
global _start

_start:
push dword 10
push dword 20
call multiply_add
add esp, 8
; eax = 10 * 20 + 10 + 20 = 230

mov ebx, eax
mov eax, 1
int 0x80

; 过程：multiply_add(x, y) = x*y + x + y
multiply_add:
push ebp
mov ebp, esp
sub esp, 8 ; ★ 在栈上分配 8 字节局部变量空间

; 现在 [ebp-4] 和 [ebp-8] 都是可用的局部变量

mov eax, [ebp + 8] ; x
mov ebx, [ebp + 12] ; y

; [ebp-4] = x * y
imul eax, ebx
mov [ebp - 4], eax ; 局部变量1 = x * y

; [ebp-8] = x + y
mov eax, [ebp + 8]
add eax, [ebp + 12]
mov [ebp - 8], eax ; 局部变量2 = x + y

; 返回值 = [ebp-4] + [ebp-8]
mov eax, [ebp - 4]
add eax, [ebp - 8]

; ★ 清理局部变量空间并恢复
mov esp, ebp ; 等价于 add esp, 8
pop ebp
ret

局部变量在栈帧中分配，过程返回后自动释放。使用 `leave` 指令可以替代 `mov esp, ebp; pop ebp`，效果等价。

### 过程调用约定对比

约定参数传递栈清理者寄存器保存返回值 cdecl栈（右到左压入）调用者EAX,ECX,EDX 由调用者保存EAX stdcall栈（右到左压入）被调用者EAX,ECX,EDX 由调用者保存EAX fastcallECX, EDX + 栈被调用者EAX,ECX,EDX 由调用者保存EAX

---

## 汇编语言 - 递归

Source: https://www.runoob.com/assembly/assembly-recursion.html

## 汇编语言 - 递归

递归（Recursion）是函数调用自身的技术。在汇编中实现递归需要正确管理栈帧，每次调用都有独立的参数和局部变量副本。

### 递归原理与栈帧

递归的核心在于 每次调用创建新的栈帧，将当前状态保存在栈上。

每次递归调用都会把参数、返回地址和局部变量压入栈中。当递归终止时，栈帧逐层弹出，每层获得其子调用的结果并完成计算。

递归要素说明汇编实现 终止条件防止无限递归条件判断 + je/jg 跳转 递归调用函数调用自身call 指令 状态保存每次调用独立状态栈帧（push ebp; mov ebp, esp） 参数传递每次不同参数push 参数到栈

递归虽然代码优雅，但在汇编中每次调用都有不小的栈开销。递归深度过大可能导致栈溢出（Stack Overflow）。对性能敏感的场景，考虑将递归改写为循环。

### 阶乘递归实现

### 实例

; 文件路径：factorial.asm
; 递归计算阶乘：n! = n * (n-1)!

section .data
n dd 5 ; 计算 5!
result dd 0
output_msg db 'Factorial result: '
output_len equ $ - output_msg
newline db 0xA

section .bss
result_str resb 12

section .text
global _start

_start:
; 调用 factorial(5)
push dword [n] ; 压入参数 n
call factorial
add esp, 4 ; 清理参数
; eax = 120 (5!)

mov [result], eax

; 输出结果
mov eax, 4
mov ebx, 1
mov ecx, output_msg
mov edx, output_len
int 0x80

; 简单输出数字（仅用于个位数演示）
mov eax, [result]
; 这里简化处理
mov ebx, eax
mov eax, 1
int 0x80

; 过程：factorial(n)
; 输入：[ebp+8] = n
; 输出：eax = n!
factorial:
push ebp ; 保存旧栈帧
mov ebp, esp ; 建立新栈帧

mov eax, [ebp + 8] ; 获取参数 n
cmp eax, 1 ; n <= 1 ?
jg recurse ; n > 1，继续递归

; 终止条件：n <= 1，返回 1
mov eax, 1
jmp factorial_end

recurse:
; 保存当前 n 的值（寄存器会被递归调用破坏）
push eax ; 保存 n

; 递归调用 factorial(n-1)
dec eax ; n - 1
push eax ; 参数：n-1
call factorial ; 递归调用
add esp, 4 ; 清理参数

; 恢复 n
pop ebx ; ebx = n

; eax = n * factorial(n-1)
mul ebx ; edx:eax = eax * ebx
; 对于较小的 n，edx 为 0，结果全在 eax 中

factorial_end:
pop ebp ; 恢复旧栈帧
ret

递归计算 factorial(5) 的完整调用栈与返回过程：

下面是文本版调用过程：

```

factorial(5)
-> 5 * factorial(4)
-> 4 * factorial(3)
-> 3 * factorial(2)
-> 2 * factorial(1)
-> 1 (终止条件)
<- 1 * 2 = 2
<- 2 * 3 = 6
<- 6 * 4 = 24
<- 24 * 5 = 120

```

### 斐波那契递归实现

### 实例

; 文件路径：fibonacci.asm
; 递归计算斐波那契数列：fib(n) = fib(n-1) + fib(n-2)

section .data
n dd 10 ; 计算 fib(10)

section .text
global _start

_start:
; 调用 fib(10)
push dword [n]
call fibonacci
add esp, 4
; eax = fib(10) = 55

mov ebx, eax ; 退出码 = 55
mov eax, 1
int 0x80

; 过程：fibonacci(n)
; 输入：[ebp+8] = n
; 输出：eax = fib(n)
fibonacci:
push ebp
mov ebp, esp

mov eax, [ebp + 8] ; 获取参数 n

; 终止条件：n <= 1 返回 n
cmp eax, 1
jg fib_recurse ; n > 1，需要递归
; n <= 1：fib(0)=0, fib(1)=1
jmp fib_end

fib_recurse:
; ★ 注意：递归有两个分支，需要仔细管理栈
push ebx ; 保存 ebx（将被用于保存中间结果）

; 计算 fib(n-1)
mov eax, [ebp + 8]
dec eax ; n - 1
push eax
call fibonacci
add esp, 4
mov ebx, eax ; ebx = fib(n-1)（保存结果！）

; 计算 fib(n-2)
mov eax, [ebp + 8]
sub eax, 2 ; n - 2
push eax
call fibonacci
add esp, 4
; eax = fib(n-2)

; eax = fib(n-1) + fib(n-2)
add eax, ebx ; eax = fib(n-2) + fib(n-1)

pop ebx ; 恢复 ebx

fib_end:
pop ebp
ret

递归版本的斐波那契存在大量重复计算。fib(10) 会调用 fib(9)、fib(8)、fib(8) 被两次调用...导致时间复杂度为 O(2^n)。实际工程中建议使用循环迭代版本。

### 斐波那契迭代版本（对比）

### 实例

; 文件路径：fibonacci_iter.asm
; 迭代版本：更高效，无递归开销

section .data
n dd 10

section .text
global _start

_start:
mov ecx, [n] ; ecx = n
dec ecx ; 循环 n-1 次（从第 2 项开始）

mov eax, 0 ; fib(0) = 0
mov ebx, 1 ; fib(1) = 1

cmp ecx, 0
jle fib_done ; n <= 1，直接返回

fib_loop:
mov edx, ebx ; 保存旧的 fib(n-1)
add ebx, eax ; ebx = fib(n-1) + fib(n-2)
mov eax, edx ; 更新 fib(n-2)
loop fib_loop

fib_done:
; 对于 n=10，最后 ebx = 55
mov eax, 1
mov ebx, ebx ; 返回值 = fib(n)
int 0x80

### 递归与迭代对比

特性递归迭代 代码量更短，逻辑清晰较长，但性能好 栈使用每次调用占一个栈帧，深度大时可能溢出固定少量空间 性能有 call/ret 开销无函数调用开销 可读性对数学归纳类问题直观需要手动维护循环状态 适用场景树遍历、分治算法、数学归纳简单重复、性能敏感场景

汇编中每次递归调用都会消耗栈空间。Linux 默认栈大小约 8MB，对于深度为 100000 的递归可能就会栈溢出。判断何时用递归：问题的递归深度可控（如平衡树深度 O(log n)），且递归逻辑比迭代清晰很多时。

---

## 汇编语言 - 宏

Source: https://www.runoob.com/assembly/assembly-macro.html

## 汇编语言 - 宏

宏（Macro）是 NASM 提供的一种代码复用机制，它允许你在编译时展开代码模板，减少重复编写相似代码的需求。

### 什么是宏

宏（Macro） 是一种编译时文本替换机制，由汇编器的预处理器在编译前展开。

与过程（Procedure）不同，宏不会有 CALL/RET 开销——每次使用宏时，编译器直接将宏的内容复制到使用位置。

宏不是函数调用，没有栈帧开销；但可执行文件体积会更大，因为每次展开都会复制一份代码。

### 单行宏：%define

`%define` 是单行宏，语法简单：

### 实例

; 单行宏示例

; 定义常量
%define MAX_SIZE 256
%define APP_NAME 'runoob'

; 定义带参数的宏（宏函数）
%define mul_by_2(x) (x * 2)
%define sum3(a, b, c) ((a) + (b) + (c))

section .text
global _start

_start:
mov eax, MAX_SIZE ; eax = 256
mov eax, mul_by_2(10) ; eax = (10 * 2) = 20
mov eax, sum3(1, 2, 3) ; eax = ((1)+(2)+(3)) = 6

; 重新定义（%define 可以改变）
%define MAX_SIZE 512
mov eax, MAX_SIZE ; eax = 512

; 取消定义
%undef MAX_SIZE
; mov eax, MAX_SIZE ; 错误：MAX_SIZE 未定义

mov eax, 1
mov ebx, 0
int 0x80

`%define` 宏展开时是纯文本替换。例如 `mul_by_2(2+3)` 展开为 `(2+3 * 2)`，由于运算符优先级，结果不是 10 而是 8。这就是宏参数一定要加括号的原因。

### 多行宏：%macro / %endmacro

`%macro` 用于定义包含多条指令的复杂宏：

### 实例

; 文件路径：macro_multi.asm
; 多行宏示例

; 宏定义：退出程序
; 参数 argc：宏接受多少个参数
%macro exit_program 1
mov eax, 1 ; sys_exit
mov ebx, %1 ; 第 1 个参数作为退出码
int 0x80
%endmacro

; 宏定义：打印字符串
; 接受 2 个参数：字符串地址、长度
%macro print_string 2
push eax
push ebx
push ecx
push edx
mov eax, 4 ; sys_write
mov ebx, 1 ; stdout
mov ecx, %1 ; 字符串地址
mov edx, %2 ; 字符串长度
int 0x80
pop edx
pop ecx
pop ebx
pop eax
%endmacro

section .data
msg db 'Hello, RUNOOB!', 0xA
msg_len equ $ - msg

section .text
global _start

_start:
print_string msg, msg_len ; 使用宏打印消息
print_string msg, msg_len ; 再次调用
exit_program 0 ; 退出程序

### 宏参数的高级用法

NASM 宏支持默认参数、参数计数和条件展开：

### 实例

; 高级宏参数示例

; 带默认参数的宏（参数范围 2-3）
%macro debug_print 2-3 1 ; 至少 2 个参数，最多 3 个，第 3 个默认 = 1
%if %3 = 1 ; 如果第 3 个参数 = 1（debug 模式开启）
push eax
push ebx
push ecx
push edx
mov eax, 4
mov ebx, 1
mov ecx, %1
mov edx, %2
int 0x80
pop edx
pop ecx
pop ebx
pop eax
%endif
%endmacro

; 不定数量参数的宏
%macro push_registers 1-* ; 1 到任意多个参数
%rep %0 ; %0 是参数个数
push %1 ; 展开第1个
%rotate 1 ; 向左旋转参数列表
%endrep
%endmacro

section .text
global _start

_start:
; 使用 push_registers 保存多个寄存器
push_registers eax, ebx, ecx, edx
; 展开为：
; push eax
; push ebx
; push ecx
; push edx

; 对应地弹出
pop edx
pop ecx
pop ebx
pop eax

mov eax, 1
mov ebx, 0
int 0x80

### 宏中的局部标签

宏中使用 `%%label` 定义局部标签，避免多次展开时标签冲突：

### 实例

; 宏中的局部标签

; 比较两个值并设置最小值
%macro min_val 2
mov eax, %1
mov ebx, %2
cmp eax, ebx
jle %%skip ; ★ 局部标签，每次展开会生成唯一名
mov eax, ebx
%%skip:
%endmacro

; 如果不使用局部标签，展开两次后会出现重复的 skip 标签
section .text
global _start

_start:
min_val 10, 5 ; eax = 5
min_val eax, 3 ; eax = 3

mov eax, 1
mov ebx, 0
int 0x80

普通标签在宏中展开两次会导致标签重名错误。局部标签 `%%label` 让 NASM 每次展开都生成唯一的标签名（如 `..@0001.skip`），避免冲突。

### 宏与过程的对比

特性宏（Macro）过程（Procedure） 实现方式编译时文本展开运行时 CALL/RET 执行开销无调用开销（直接内联）有 CALL/RET 开销（约几个时钟周期） 代码大小每次展开增加体积一份代码多次调用 参数类型任意文本（寄存器、立即数、内存）运行时值 调试难（展开后无痕迹）易（有函数调用栈） 适用场景短小频繁调用、类型灵活的代码复杂逻辑、代码量大的函数

### 条件编译：%if / %elif / %else / %endif

宏预处理器支持条件编译，可以根据符号定义选择生成不同的代码：

### 实例

; 文件路径：cond_compile.asm
; 条件编译示例：调试和发布模式

%define DEBUG 1 ; 1=调试模式, 0=发布模式

section .data
msg db 'Program running...', 0xA
msg_len equ $ - msg
debug_msg db '[DEBUG] Entering function', 0xA
debug_len equ $ - debug_msg

section .text
global _start

_start:
%if DEBUG = 1
; 调试模式：输出调试信息
mov eax, 4
mov ebx, 1
mov ecx, debug_msg
mov edx, debug_len
int 0x80
%endif

; 正常业务逻辑
mov eax, 4
mov ebx, 1
mov ecx, msg
mov edx, msg_len
int 0x80

mov eax, 1
mov ebx, 0
int 0x80

条件编译常用场景：

场景示例 调试/发布切换`%if DEBUG` / `%else` / `%endif` 平台适配`%ifdef LINUX` / `%ifdef WINDOWS` 功能开关通过符号存在与否控制特性

### %rep 重复块

`%rep` 用于生成重复的代码或数据：

### 实例

; %rep 重复块示例

section .data
; 生成 256 字节的查找表
; 0, 1, 4, 9, 16, 25, ...（平方数表）
%assign i 0
; 这里不能用 %rep 生成平方表来进行复杂计算
; 简单示例：生成 0-9 的 ASCII 表
digits: db '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'

section .text
global _start

_start:
; %rep 在代码中重复指令
mov eax, 0
%rep 5 ; 重复 5 次
inc eax ; eax 每次加 1
%endrep
; eax = 5

mov ebx, eax
mov eax, 1
int 0x80

过度使用宏会让代码难以阅读和调试。当宏体超过 10 行时，考虑改为过程调用。在性能敏感的热路径中，短小宏的内联效果才有明显好处。

---

## 汇编语言 - 文件管理

Source: https://www.runoob.com/assembly/assembly-file.html

## 汇编语言 - 文件管理

文件管理是程序与外部存储交互的基础。

通过系统调用，汇编程序可以打开、读取、写入和关闭文件，与高级语言一样灵活地处理文件 IO。

### 文件操作流程总览

文件操作遵循打开→读写→关闭的基本流程，以及三个标准文件描述符：

### 文件操作相关的系统调用

系统调用调用号用途主要参数 sys_open5打开/创建文件EBX=文件名, ECX=标志, EDX=权限 sys_read3读取文件EBX=fd, ECX=缓冲区, EDX=字节数 sys_write4写入文件EBX=fd, ECX=缓冲区, EDX=字节数 sys_close6关闭文件EBX=fd sys_creat8创建文件（旧方式）EBX=文件名, ECX=权限 sys_lseek19移动文件指针EBX=fd, ECX=偏移, EDX=起始位置 sys_unlink10删除文件EBX=文件名

### 文件打开标志

常量值说明 O_RDONLY0只读 O_WRONLY1只写 O_RDWR2读写 O_CREAT64（0x40）若文件不存在则创建 O_TRUNC512（0x200）打开时清空文件内容 O_APPEND1024（0x400）追加到文件末尾

多个标志用 OR 组合，例如 `O_WRONLY | O_CREAT | O_TRUNC = 1 | 64 | 512 = 577`

### 创建并写入文件

### 实例

; 文件路径：file_write.asm
; 创建文件并写入内容

section .data
filename db 'runoob_output.txt', 0 ; 文件名（null 结尾）
content db 'Hello, RUNOOB!', 0xA ; 要写入的内容
content_len equ $ - content
create_msg db 'File created successfully!', 0xA
create_msg_len equ $ - create_msg

section .bss
fd resd 1 ; 文件描述符

section .text
global _start

_start:
; 1. 创建文件
mov eax, 8 ; sys_creat
mov ebx, filename ; 文件名
mov ecx, 0o644 ; 权限：rw-r--r--
; 0o644 = 所有者读写、组只读、其他只读
int 0x80

mov [fd], eax ; 保存文件描述符

; 2. 写入内容到文件
mov eax, 4 ; sys_write
mov ebx, [fd] ; 文件描述符
mov ecx, content ; 数据地址
mov edx, content_len ; 数据长度
int 0x80

; 3. 关闭文件
mov eax, 6 ; sys_close
mov ebx, [fd] ; 文件描述符
int 0x80

; 4. 提示成功
mov eax, 4
mov ebx, 1
mov ecx, create_msg
mov edx, create_msg_len
int 0x80

mov eax, 1
mov ebx, 0
int 0x80

运行结果：

```

$ nasm -f elf32 file_write.asm -o file_write.o
$ ld -m elf_i386 file_write.o -o file_write
$ ./file_write
File created successfully!
$ cat runoob_output.txt
Hello, RUNOOB!

```

### 读取文件内容

### 实例

; 文件路径：file_read.asm
; 打开并读取文件内容

section .data
filename db 'runoob_output.txt', 0
open_error db 'Error: cannot open file', 0xA
open_error_len equ $ - open_error
read_success db 'File content:', 0xA
read_success_len equ $ - read_success

section .bss
fd resd 1
buffer resb 1024 ; 读取缓冲区

section .text
global _start

_start:
; 1. 打开文件（只读）
mov eax, 5 ; sys_open
mov ebx, filename ; 文件名
mov ecx, 0 ; 只读模式
mov edx, 0 ; 权限（只读时忽略）
int 0x80

cmp eax, 0 ; 文件描述符 < 0 表示错误
jl open_failed

mov [fd], eax ; 保存文件描述符

; 2. 读取文件内容
mov eax, 3 ; sys_read
mov ebx, [fd] ; 文件描述符
mov ecx, buffer ; 缓冲区
mov edx, 1024 ; 最多读取 1024 字节
int 0x80
; 返回值 eax = 实际读取的字节数
mov esi, eax ; 保存读取字节数

; 3. 关闭文件
mov eax, 6 ; sys_close
mov ebx, [fd]
int 0x80

; 4. 输出提示信息
mov eax, 4
mov ebx, 1
mov ecx, read_success
mov edx, read_success_len
int 0x80

; 5. 输出文件内容到屏幕
mov eax, 4
mov ebx, 1
mov ecx, buffer
mov edx, esi ; 使用实际读取的字节数
int 0x80
jmp exit

open_failed:
mov eax, 4
mov ebx, 1
mov ecx, open_error
mov edx, open_error_len
int 0x80

exit:
mov eax, 1
mov ebx, 0
int 0x80

运行结果：

```

$ nasm -f elf32 file_read.asm -o file_read.o
$ ld -m elf_i386 file_read.o -o file_read
$ ./file_read
File content:
Hello, RUNOOB!

```

### 追加写入文件

### 实例

; 文件路径：file_append.asm
; 以追加模式打开文件并写入

section .data
filename db 'runoob_log.txt', 0
log_entry db '[INFO] Program executed successfully.', 0xA
log_len equ $ - log_entry

section .bss
fd resd 1

section .text
global _start

_start:
; 打开文件（创建 + 追加模式）
mov eax, 5 ; sys_open
mov ebx, filename ; 文件名
mov ecx, 0x441 ; O_WRONLY | O_CREAT | O_APPEND
; O_WRONLY=1, O_CREAT=0x40, O_APPEND=0x400
; 组合：1 | 0x40 | 0x400 = 0x441
mov edx, 0o644 ; 创建时的权限
int 0x80

mov [fd], eax

; 追加写入
mov eax, 4 ; sys_write
mov ebx, [fd] ; 文件描述符
mov ecx, log_entry ; 日志内容
mov edx, log_len
int 0x80

; 关闭文件
mov eax, 6
mov ebx, [fd]
int 0x80

mov eax, 1
mov ebx, 0
int 0x80

### 复制文件

一个综合示例：将一个文件的内容复制到另一个文件：

### 实例

; 文件路径：file_copy.asm
; 复制文件：将 runoob_input.txt 复制到 runoob_copy.txt

section .data
src_file db 'runoob_input.txt', 0
dst_file db 'runoob_copy.txt', 0
success_msg db 'File copied successfully!', 0xA
success_len equ $ - success_msg
error_msg db 'Error during file copy', 0xA
error_len equ $ - error_msg

section .bss
src_fd resd 1
dst_fd resd 1
buffer resb 4096 ; 4KB 复制缓冲区

section .text
global _start

_start:
; 1. 打开源文件（只读）
mov eax, 5
mov ebx, src_file
mov ecx, 0 ; O_RDONLY
mov edx, 0
int 0x80
cmp eax, 0
jl error_exit
mov [src_fd], eax

; 2. 创建目标文件
mov eax, 8 ; sys_creat
mov ebx, dst_file
mov ecx, 0o644
int 0x80
cmp eax, 0
jl error_exit
mov [dst_fd], eax

copy_loop:
; 3. 从源文件读取
mov eax, 3 ; sys_read
mov ebx, [src_fd]
mov ecx, buffer
mov edx, 4096
int 0x80
; eax = 实际读取的字节数
cmp eax, 0 ; 读到 0 字节？
jle copy_done ; 是，文件结束

; 4. 写入目标文件
mov esi, eax ; 保存读取的字节数
mov eax, 4 ; sys_write
mov ebx, [dst_fd]
mov ecx, buffer
mov edx, esi ; 写入实际读取的字节数
int 0x80
jmp copy_loop ; 继续循环

copy_done:
; 5. 关闭文件
mov eax, 6 ; sys_close
mov ebx, [src_fd]
int 0x80

mov eax, 6
mov ebx, [dst_fd]
int 0x80

; 6. 输出成功信息
mov eax, 4
mov ebx, 1
mov ecx, success_msg
mov edx, success_len
int 0x80
jmp exit

error_exit:
mov eax, 4
mov ebx, 1
mov ecx, error_msg
mov edx, error_len
int 0x80

exit:
mov eax, 1
mov ebx, 0
int 0x80

文件操作前后务必检查系统调用的返回值（在 EAX 中）。负值表示错误（如 -2 是 ENOENT 文件不存在，-13 是 EACCES 权限不足），忽略错误返回值是文件操作 bug 的最大来源。

---

## 汇编语言 - 内存管理

Source: https://www.runoob.com/assembly/assembly-memory-mgmt.html

## 汇编语言 - 内存管理

内存管理是系统编程的重要部分。在汇编语言中，你可以通过系统调用动态分配和释放内存，直接操作内存地址，实现高效的内存管理。

### 程序的内存布局

一个运行中的程序在内存中的分布如下：

一个 Linux 进程在内存中的布局如下：

```

高地址 (0xFFFFFFFF)
+----------------------+
| 内核空间 | 用户程序不可访问
+----------------------+ (0xC0000000)
| 栈 (Stack) | <-- ESP 指向栈顶
| 向下增长 |
+----------------------+
| |
| 内存映射区域 | mmap 分配的区域
| |
+----------------------+
| 堆 (Heap) | <-- brk/sbrk 管理
| 向上增长 |
+----------------------+
| BSS 段 (.bss) | 未初始化全局变量
+----------------------+
| 数据段 (.data) | 已初始化全局变量
+----------------------+
| 代码段 (.text) | 程序指令（只读）
+----------------------+
低地址 (0x08048000)

```

### 动态内存分配：brk 系统调用

`sys_brk`（系统调用号 45）通过调整程序的 数据段边界（program break） 来分配/释放内存。

program break 是数据段（包括 .data、.bss 和堆）的结束位置，在其之上的内存程序不能访问。

调用方式说明返回值 `ebx = 0`获取当前 break 地址EAX = 当前 break 地址 `ebx = 新地址`设置新的 break 地址EAX = 新的 break 地址（失败返回原地址）

### 实例

; 文件路径：brk_demo.asm
; 使用 sys_brk 动态分配内存

section .data
alloc_msg db 'Memory allocated successfully at: 0x'
alloc_len equ $ - alloc_msg
newline db 0xA

section .bss
orig_break resd 1 ; 保存原始 break 地址

section .text
global _start

_start:
; 1. 获取当前 program break
mov eax, 45 ; sys_brk
mov ebx, 0 ; ebx=0 表示查询当前 break
int 0x80
mov [orig_break], eax ; 保存原始 break 地址

; 2. 分配 4096 字节（4KB）内存
mov ebx, eax ; 当前 break 地址
add ebx, 4096 ; 增加 4096 字节
mov eax, 45 ; sys_brk
int 0x80
; EAX = 新的 break 地址（即分配后的区域末尾）

; 3. 新分配的内存在 orig_break 到 eax-1 之间
; 可以安全地读写这片内存
mov ebx, [orig_break] ; 获取分配区域的起始地址
mov dword [ebx], 42 ; 在新内存中写入 42
mov eax, [ebx] ; 读取回来

; 4. 释放内存：将 break 恢复到原位
mov eax, 45 ; sys_brk
mov ebx, [orig_break] ; 恢复到原始 break
int 0x80

mov eax, 1
mov ebx, 0
int 0x80

sys_brk 只能调整连续的数据段边界，不能释放中间的内存。要释放某块已分配的内存，必须先释放之后分配的所有内存。这就是为什么现代程序更多使用 mmap 或 malloc 库函数。

### 使用 mmap 分配内存（推荐）

`sys_mmap2`（系统调用号 192）更灵活，可以独立分配和释放多块内存：

### 实例

; 文件路径：mmap_demo.asm
; 使用 mmap 分配可独立释放的内存

section .data
; mmap 相关常量
PROT_READ equ 1
PROT_WRITE equ 2
MAP_PRIVATE equ 2
MAP_ANONYMOUS equ 0x20

section .bss
mem_block resd 1 ; 保存分配的内存地址

section .text
global _start

_start:
; mmap 调用：分配 4096 字节的匿名内存
mov eax, 192 ; sys_mmap2
mov ebx, 0 ; 让内核选择地址
mov ecx, 4096 ; 分配大小：4KB
mov edx, PROT_READ | PROT_WRITE ; 可读可写
mov esi, MAP_PRIVATE | MAP_ANONYMOUS ; 私有匿名映射
mov edi, -1 ; 文件描述符（匿名映射用 -1）
mov ebp, 0 ; 偏移量（匿名映射用 0）
int 0x80
; EAX = 分配的内存地址（失败返回负值）

cmp eax, -4096 ; 检查是否失败
ja mmap_failed ; 如果在 -4095 ~ -1 之间，失败

mov [mem_block], eax ; 保存内存地址

; 使用分配的内存：写入数据
mov ebx, [mem_block]
mov dword [ebx], 0x12345678 ; 写入 4 字节
mov dword [ebx + 4], 'runo' ; 写入 "runo"
mov dword [ebx + 8], 'ob!!' ; 写入 "ob!!"

; 释放内存：munmap
mov eax, 91 ; sys_munmap
mov ebx, [mem_block] ; 内存地址
mov ecx, 4096 ; 释放大小
int 0x80

jmp exit

mmap_failed:
; 处理错误...

exit:
mov eax, 1
mov ebx, 0
int 0x80

### 内存读写操作

汇编语言中，所有内存访问通过 `mov` 指令配合方括号完成：

### 实例

; 内存读写基本操作

section .data
var1 db 0x55 ; 1 字节
var2 dw 0x1234 ; 2 字节
var3 dd 0x12345678 ; 4 字节

section .text
global _start

_start:
; 读取不同大小的内存
mov al, [var1] ; 读取 1 字节：al = 0x55
mov ax, [var2] ; 读取 2 字节：ax = 0x1234
mov eax, [var3] ; 读取 4 字节：eax = 0x12345678

; 写入不同大小的内存
mov byte [var1], 0xAA ; 写入 1 字节
mov word [var2], 0xABCD ; 写入 2 字节
mov dword [var3], 0xDEADBEEF ; 写入 4 字节

; 通过指针操作内存
mov ebx, var3 ; ebx 指向 var3
mov eax, [ebx] ; 间接读取 var3 的值
add dword [ebx], 1 ; var3 = var3 + 1

mov eax, 1
mov ebx, 0
int 0x80

### 内存操作的安全规范

在汇编中操作内存没有"安全网"，以下规范需要牢记：

规范说明违反后果 不越界访问读写不超过变量/缓冲区的大小数据损坏、段错误 不读未初始化内存.bss 段的值不确定不可预测的结果 不在无效地址上读写确保指针指向有效内存段错误（Segmentation Fault） 不写只读内存.text 段和常量区不可写段错误 地址对齐字访问用偶地址，双字用 4 的倍数性能下降或总线错误

### 实例

; 常见内存错误示例（警告：会导致崩溃！）

section .data
small_buf db 0, 0, 0, 0 ; 只有 4 字节

section .text
global _start

_start:
; 危险操作1：越界写入
; mov dword [small_buf + 3], 0x12345678
; 这会覆盖 small_buf 之后 3 字节的数据！

; 危险操作2：写入代码段（只读）
; mov dword [_start], 0x90 ; 尝试修改代码，会导致段错误

; 危险操作3：空指针/无效地址
; mov eax, [0] ; 读取地址 0，会导致段错误
; mov [0], eax ; 写入地址 0，会导致段错误

; 安全做法：始终在分配的内存范围内操作
mov dword [small_buf], 'runo' ; 正确：在 4 字节内操作

mov eax, 1
mov ebx, 0
int 0x80

### 栈内存管理

栈是另一种重要的内存区域，由 PUSH/POP 指令和 ESP 寄存器管理：

### 实例

; 栈操作全面示例

section .data
original_esp dd 0

section .text
global _start

_start:
mov [original_esp], esp ; 保存原始栈指针

; PUSH：压栈（ESP 减 4，写入数据）
push dword 100 ; 压入 100
; ESP = ESP - 4, [ESP] = 100

push dword 200 ; 压入 200
push dword 300 ; 压入 300

; 栈布局：
; ESP+0: 300
; ESP+4: 200
; ESP+8: 100

; POP：弹栈（读取数据，ESP 加 4）
pop eax ; eax = 300, ESP += 4
pop ebx ; ebx = 200, ESP += 4
pop ecx ; ecx = 100, ESP += 4

; 通过栈分配局部空间
sub esp, 256 ; 在栈上分配 256 字节
; 现在 ESP 到 ESP+255 这 256 字节可以安全使用
mov dword [esp], 42 ; 在局部空间写入 42
add esp, 256 ; 释放局部空间

; 确保 ESP 回到原始位置！
mov eax, 1
mov ebx, 0
int 0x80

栈操作最重要的规则：PUSH 和 POP 必须配对。函数返回前必须确保 ESP 回到正确的位置，否则 RET 会弹出错误的返回地址，导致程序崩溃或执行任意代码。这是汇编编程中最致命的 bug 之一。

### 内存管理方案对比

方案灵活性复杂度适用场景 全局变量（.data/.bss）低（编译时固定）最低固定大小的数据 栈分配（sub esp）中（函数内动态）低函数局部变量 brk/sbrk中（只能增长/收缩）中连续内存区域 mmap高（任意分配释放）高需要灵活管理内存时

---

## 汇编指令速查表

Source: https://www.runoob.com/assembly/x86-assembly-instructions-html.html

## 汇编指令速查表

本文以最常见的 x86 架构汇编语言为对象，将常用指令助记符按功能分类整理，方便学习和快速查阅。

指令助记符在不同汇编器（MASM、NASM、GAS）中基本一致，仅在操作数书写格式上略有差异。

九大类指令概览：

分类 指令数量 主要用途 数据传送指令 15 条 在寄存器、内存、I/O 端口之间搬运数据 算术运算指令 18 条 加减乘除、比较及数制调整 逻辑运算与移位指令 14 条 按位与或非异或、移位和循环移位 串操作指令 8 条 处理连续内存区域，配合重复前缀批量操作 控制转移指令 30+ 条 跳转、循环、过程调用与中断 条件设置指令 16 条 根据标志位设置字节值，用于消除分支 输入输出指令 4 条 读写 I/O 端口 处理器控制指令 14 条 控制标志位、同步、系统管理 浮点指令（x87 FPU） 14 条 传统 x87 浮点栈运算

### 数据传送指令

数据传送指令是最基本的指令类型，负责在寄存器、内存和 I/O 端口之间搬运数据。

大部分传送指令不影响标志位（SAHF、POPF 除外）。

助记符 功能简述 `MOV` 把源操作数复制到目的操作数，支持寄存器与寄存器、寄存器与内存、立即数到寄存器或内存 `MOVSX` 带符号扩展的传送，将小尺寸操作数符号扩展为大尺寸（如字节扩展为双字，高位填充符号位） `MOVZX` 带零扩展的传送，高位补 0（用于无符号数扩展） `PUSH` 将操作数压入堆栈，栈指针 ESP/RSP 自动减小（先减后存） `POP` 从堆栈弹出数据到目的操作数，栈指针自动增加（先取后加） `PUSHA / PUSHAD` 将所有通用寄存器压栈（16 位用 PUSHA，32 位用 PUSHAD），常用于保存现场 `POPA / POPAD` 从堆栈恢复所有通用寄存器，与 PUSHA/PUSHAD 对称使用 `XCHG` 交换两个操作数的内容，至少有一个操作数是寄存器 `LEA` 取有效地址（Load Effective Address），将源操作数的地址而非内存内容存入目的寄存器，常用于指针计算 `LDS / LES / LFS / LGS / LSS` 将远指针装入段寄存器与通用寄存器（现代平坦内存模型中较少使用） `LAHF` 将标志寄存器的低 8 位装入 AH 寄存器 `SAHF` 将 AH 寄存器的内容存回标志寄存器的低 8 位 `PUSHF / PUSHFD` 标志寄存器压栈（16 位用 PUSHF，32 位用 PUSHFD） `POPF / POPFD` 从堆栈恢复标志寄存器 `XLAT / XLATB` 查表翻译指令，将 AL 中的索引值对应内存字节送入 AL，BX/EBX 指向表基址

`LEA` 与 `MOV` 的关键区别：`LEA REG, [ADDR]` 装入的是地址值本身，而 `MOV REG, [ADDR]` 装入的是该地址处的内存内容。在进行地址计算（如 `LEA EAX, [EBX+4*ECX+8]`）时 LEA 非常高效。

### 算术运算指令

算术运算指令完成加、减、乘、除以及比较等操作，执行结果会直接影响标志寄存器中的相关标志位。

程序员通常根据运算后的零标志（ZF）、进位标志（CF）、符号标志（SF）和溢出标志（OF）来判断计算结果。

助记符 功能简述 `ADD` 加法，目的 = 目的 + 源，影响 OF、SF、ZF、CF 等标志位 `ADC` 带进位加法，目的 = 目的 + 源 + CF，用于多精度算术运算 `SUB` 减法，目的 = 目的 - 源 `SBB` 带借位减法，目的 = 目的 - 源 - CF，配合 ADC 用于多精度运算 `INC` 操作数加 1，不影响 CF 标志 `DEC` 操作数减 1，不影响 CF 标志 `MUL` 无符号乘法，被乘数在 AL/AX/EAX，乘积存入 AX/DX:AX/EDX:EAX `IMUL` 带符号乘法，支持单操作数、双操作数和三操作数三种形式 `DIV` 无符号除法，被除数在 AX/DX:AX/EDX:EAX，商和余数存入指定寄存器 `IDIV` 带符号除法，规则与 DIV 类似但处理符号 `CMP` 比较两个操作数，执行减法但不保存结果，仅影响标志位（通常后跟条件跳转） `NEG` 求补，对操作数取反再加 1，等价于 0 减去操作数 `DAA` 加法后十进制调整，将 AL 调整为压缩 BCD 码格式 `DAS` 减法后十进制调整 `AAA` 加法后 ASCII 调整，用于非压缩 BCD 码 `AAS` 减法后 ASCII 调整 `AAM` 乘法后 ASCII 调整 `AAD` 除法前 ASCII 调整 `CBW / CWDE` 符号扩展：CBW 将 AL 符号扩展为 AX，CWDE 将 AX 符号扩展为 EAX `CWD / CDQ` 符号扩展：CWD 将 AX 扩展为 DX:AX，CDQ 将 EAX 扩展为 EDX:EAX（常用于除法前准备被除数）

有符号运算使用 `IMUL` / `IDIV`，无符号运算使用 `MUL` / `DIV`。选错指令会导致结果错误——比如同一个二进制值在有符号和无符号解释下代表不同的数值。

### 逻辑运算与移位指令

这类指令按位进行与、或、非、异或等逻辑操作，以及移位和循环移位。

移位操作常用于快速乘除 2 的幂、位域提取和位掩码操作。

助记符 功能简述 `AND` 按位与，常用于掩码操作（将特定位清零） `OR` 按位或，常用于将特定位置 1 `XOR` 按位异或，相同位为 0、不同为 1。XOR REG, REG 是清零寄存器的经典高效写法 `NOT` 按位取反，所有位翻转 `TEST` 测试位，执行 AND 运算但不保存结果，仅影响标志位（常用于检测某位是否为 0） `SHL / SAL` 逻辑左移 / 算术左移，两者功能完全相同，低位补 0，高位移入 CF `SHR` 逻辑右移，高位补 0，低位移入 CF（用于无符号数） `SAR` 算术右移，高位用原符号位填充，保留数值符号（用于有符号数） `ROL` 循环左移，移出的位补回另一侧 `ROR` 循环右移，移出的位补回另一侧 `RCL` 带进位循环左移，CF 参与循环（CF 作为临时存储位） `RCR` 带进位循环右移，CF 参与循环 `SHLD` 双精度左移，将目的和源操作数联合左移，结果存入目的 `SHRD` 双精度右移，将目的和源操作数联合右移，结果存入目的

`XOR REG, REG` 是清零寄存器的最优方式——生成的机器码比 `MOV REG, 0` 更短，执行速度也更快。

### 串操作指令

串操作指令专门处理连续内存区域，搭配重复前缀使用，常用于数组复制、内存比较和缓冲区搜索等批量操作。

这些指令的源地址默认由 DS:ESI 指向，目的地址由 ES:EDI 指向。

助记符 功能简述 `MOVS` 串传送，将 [DS:ESI] 复制到 [ES:EDI] 并自动更新 ESI、EDI。按操作数大小分为 MOVSB（字节）、MOVSW（字）、MOVSD（双字） `STOS` 串存储，将 AL/AX/EAX 存入 [ES:EDI] 并更新 EDI。常用于初始化内存区域为同一值 `LODS` 串加载，从 [DS:ESI] 读入 AL/AX/EAX 并更新 ESI。较少使用，因为一次只加载一个值到累加器 `CMPS` 串比较，比较 [DS:ESI] 与 [ES:EDI] 并影响标志位。通常配合 REPE/REPNE 搜索匹配或不匹配的位置 `SCAS` 串扫描，比较 AL/AX/EAX 与 [ES:EDI] 并更新 EDI。常用于在内存中搜索特定值 `REP` 重复前缀，当 ECX ≠ 0 时重复执行后续串指令，每次执行后 ECX 自动减 1 `REPE / REPZ` 相等/为零时重复，ECX ≠ 0 且 ZF = 1 时继续（通常与 CMPS、SCAS 联用搜索匹配项） `REPNE / REPNZ` 不等/不为零时重复，ECX ≠ 0 且 ZF = 0 时继续（用于搜索不匹配项）

方向标志 DF 控制串操作指针的移动方向：执行 `CLD` 后 DF=0，指针递增（正向处理）；执行 `STD` 后 DF=1，指针递减（反向处理）。在使用串操作前应显式设置 DF，不要依赖默认值。

### 控制转移指令

控制转移指令改变程序执行流向，包括无条件跳转、条件跳转、循环以及过程调用与返回。

这是实现程序逻辑分支、循环和函数调用的基础。

#### 无条件跳转与调用

无条件跳转不检查任何条件，直接改变 EIP/RIP 的值，让 CPU 从新地址继续执行。

助记符 功能简述 `JMP` 无条件跳转，支持短跳转（-128 到 +127 字节）、近跳转（段内）和远跳转（跨段） `CALL` 调用过程，先将返回地址压入堆栈，然后跳转到目标地址 `RET / RETF` 从过程返回。RET 用于近返回（段内），RETF 用于远返回（跨段）。可带立即数参数同时释放栈上的参数空间

#### 条件跳转

条件跳转根据标志寄存器中特定标志位的值决定是否跳转。

通常紧跟在 `CMP` 或 `TEST` 指令之后使用。

助记符 跳转条件 典型用途 `JE / JZ` ZF = 1（相等 / 为零） CMP 后相等时跳转 `JNE / JNZ` ZF = 0（不等 / 不为零） CMP 后不等时跳转 `JS` SF = 1（结果为负） 运算结果为负数时跳转 `JNS` SF = 0（结果为正） 运算结果为非负数时跳转 `JC` CF = 1（有进位/借位） 无符号比较中小于时跳转，等同于 JB `JNC` CF = 0（无进位/借位） 无符号比较中大于等于时跳转，等同于 JAE `JO` OF = 1（有溢出） 有符号运算溢出时跳转 `JNO` OF = 0（无溢出） 有符号运算无溢出时跳转 `JP / JPE` PF = 1（低 8 位中 1 的个数为偶数） 奇偶校验为偶时跳转 `JNP / JPO` PF = 0（1 的个数为奇数） 奇偶校验为奇时跳转 `JA / JNBE` CF = 0 且 ZF = 0 无符号大于（Above） `JAE / JNB` CF = 0 无符号大于等于（Above or Equal） `JB / JNAE` CF = 1 无符号小于（Below） `JBE / JNA` CF = 1 或 ZF = 1 无符号小于等于（Below or Equal） `JG / JNLE` ZF = 0 且 SF = OF 有符号大于（Greater） `JGE / JNL` SF = OF 有符号大于等于（Greater or Equal） `JL / JNGE` SF ≠ OF 有符号小于（Less） `JLE / JNG` ZF = 1 或 SF ≠ OF 有符号小于等于（Less or Equal） `JCXZ` CX = 0 时跳转 16 位模式下判断 CX 是否为零 `JECXZ` ECX = 0 时跳转 32 位模式下判断 ECX 是否为零

无符号比较和有符号比较使用不同的跳转指令，这是初学者最容易出错的地方。例如，比较两个数时 `JA` 检查无符号大于，`JG` 检查有符号大于——同一个 CMP 结果在这两种判断下可能不同。

#### 循环指令

循环指令使用 ECX 作为计数器，每次迭代自动递减并判断是否继续循环。

助记符 功能简述 `LOOP` ECX 减 1，若 ECX ≠ 0 则跳转到目标地址 `LOOPE / LOOPZ` ECX 减 1，若 ECX ≠ 0 且 ZF = 1 则跳转（等于时继续循环） `LOOPNE / LOOPNZ` ECX 减 1，若 ECX ≠ 0 且 ZF = 0 则跳转（不等时继续循环）

#### 中断与返回

助记符 功能简述 `INT n` 软件中断，调用中断向量号为 n 的中断服务程序（如 DOS 下 INT 21h 为系统功能调用） `INTO` 溢出中断，当 OF = 1 时调用 INT 4（用于捕获算术溢出异常） `IRET / IRETD` 从中断服务程序返回，恢复被中断时的 EFLAGS 和返回地址

### 条件设置指令

条件设置指令（386 及以上处理器支持）根据标志位状态将目的字节置为 1 或 0。

这类指令常用于消除分支跳转，提升代码执行效率——用条件设置替代短分支可以避免分支预测失败带来的性能损失。

助记符 设置条件 等价语义 `SETZ / SETE` ZF = 1 相等 / 结果为零时设置 `SETNZ / SETNE` ZF = 0 不等 / 结果不为零时设置 `SETC` CF = 1 有进位/借位时设置 `SETNC` CF = 0 无进位/借位时设置 `SETO` OF = 1 有溢出时设置 `SETNO` OF = 0 无溢出时设置 `SETS` SF = 1 结果为负时设置 `SETNS` SF = 0 结果为正时设置 `SETG / SETNLE` ZF=0 且 SF=OF 有符号大于时设置 `SETGE / SETNL` SF = OF 有符号大于等于时设置 `SETL / SETNGE` SF ≠ OF 有符号小于时设置 `SETLE / SETNG` ZF=1 或 SF≠OF 有符号小于等于时设置 `SETA / SETNBE` CF=0 且 ZF=0 无符号大于时设置 `SETAE / SETNB` CF = 0 无符号大于等于时设置 `SETB / SETNAE` CF = 1 无符号小于时设置 `SETBE / SETNA` CF=1 或 ZF=1 无符号小于等于时设置

`CMP EAX, EBX` 后跟 `SETG AL` 的效果是：如果 EAX 有符号大于 EBX，则 AL 被设为 1，否则为 0。这比 `JG label; MOV AL, 1; JMP done; label: MOV AL, 0; done:` 更简洁高效。

### 输入输出指令

I/O 指令用于 CPU 与外部设备端口之间的数据交换。

在 x86 架构中，端口地址空间与内存地址空间相互独立，需要通过专门的 I/O 指令来访问。

助记符 功能简述 `IN` 从指定端口读入数据到 AL/AX/EAX。端口地址在 0-255 时可直接写立即数，超过 255 时必须通过 DX 寄存器指定 `OUT` 将 AL/AX/EAX 中的数据写入指定端口。端口地址规则与 IN 相同 `INS` 从端口读入字节/字/双字到 ES:EDI 指向的内存，并自动更新 EDI（类似 STOS 但从 I/O 读取） `OUTS` 将 DS:ESI 指向的内存字节/字/双字输出到端口，并自动更新 ESI

在现代操作系统中（Windows、Linux），用户态程序通常无法直接执行 I/O 指令——这些指令被操作系统限制为仅内核态可用。在 DOS 环境或裸机编程中则可以自由使用。

### 处理器控制与杂项指令

这类指令用于控制处理器的运行状态，包括标志位操作、同步控制和系统信息获取。

助记符 功能简述 `CLC` 清进位标志，CF 置为 0 `STC` 置进位标志，CF 置为 1 `CMC` 进位标志取反，CF 变为其相反值 `CLD` 清方向标志，DF 置为 0，串操作指针自动递增 `STD` 置方向标志，DF 置为 1，串操作指针自动递减 `CLI` 清中断标志，IF 置为 0，禁止可屏蔽硬件中断 `STI` 置中断标志，IF 置为 1，允许可屏蔽硬件中断 `NOP` 空操作，机器码为 0x90，不执行任何有效操作。常用于指令对齐、延时填充或预留补丁空间 `HLT` 停机，处理器暂停执行直到收到外部中断或复位信号 `WAIT / FWAIT` 等待 FPU 完成当前操作，主要用于与 x87 协处理器同步 `LOCK` 总线锁定前缀，使紧随其后的指令以原子方式执行。经典用法：LOCK XCHG 实现自旋锁、LOCK CMPXCHG 实现 CAS 操作 `CPUID` 获取 CPU 特征信息，包括厂商 ID、支持的指令集扩展等。调用前需在 EAX 中设置功能号 `RDTSC` 读取时间戳计数器，将 64 位计数值存入 EDX:EAX，用于高精度性能测量

### 浮点指令

以下为传统 x87 FPU 浮点指令，适用于使用浮点栈进行科学计算的场景。

在现代编程中，SSE/SSE2 等 SIMD 指令集已逐步取代 x87，但在理解遗留代码或进行简单浮点运算时，x87 仍被广泛使用。

助记符 功能简述 `FLD` 将浮点数从内存加载到 FPU 栈顶 ST(0)，栈指针上移 `FST / FSTP` 将栈顶数据存入内存。FST 不弹出栈，FSTP 弹出栈（栈指针下移） `FADD` 浮点加法，将 ST(0) 与指定操作数相加，结果存入 ST(0) `FSUB` 浮点减法，ST(0) 减去指定操作数 `FMUL` 浮点乘法 `FDIV` 浮点除法，ST(0) 除以指定操作数 `FCOM / FCOMP` 比较 ST(0) 与指定操作数，设置 FPU 状态字中的条件码。FCOMP 比较后弹出栈 `FCHS` 改变 ST(0) 中数值的符号（正变负、负变正） `FSQRT` 计算 ST(0) 的平方根，结果存入 ST(0) `FSIN / FCOS` 计算 ST(0) 的正弦 / 余弦值（参数以弧度为单位），结果存入 ST(0) `FPTAN` 计算 ST(0) 的正切值 `FLD1` 将常数 1.0 压入 FPU 栈顶 `FLDZ` 将常数 0.0 压入 FPU 栈顶 `FLDPI` 将常数 π（3.14159...）压入 FPU 栈顶

x87 FPU 使用浮点寄存器栈（ST(0) - ST(7) 共 8 个 80 位寄存器）进行操作。使用前需注意栈的深度限制——压入超过 8 个值会导致栈溢出异常。

### 注意事项

- 区分有符号和无符号指令。MUL/IMUL、DIV/IDIV、条件跳转中的 JG/JA 系列对应不同的数值解释。汇编器不会帮你检查类型——用错指令会静默产生错误结果。
- 串操作前务必设置 DF 标志。使用 CLD（递增）或 STD（递减）显式指定方向，不要假设默认值。不同编译器、不同调用约定的默认 DF 状态可能不同。
- 现代操作系统限制特权指令。CLI/STI、HLT、IN/OUT 等指令在 Ring 3（用户态）下执行会导致异常。这些指令通常仅在操作系统内核或裸机环境中使用。
- 优先使用现代指令集。在新代码中，条件设置指令（SETcc）优于条件跳转（消除分支预测开销），SSE/AVX 指令优于 x87（性能更好、编程模型更简单）。本表涵盖的 x87 和传统指令主要用于理解遗留代码。
