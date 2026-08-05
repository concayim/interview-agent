# Julia - 菜鸟教程

Tutorial: https://www.runoob.com/julia/julia-tutorial.html

---

## Julia 教程

Source: https://www.runoob.com/julia/julia-tutorial.html

## Julia 教程

Julia 是一个开源的编程语言，采用 MIT 许可证，每个人都可以免费使用。

Julia 是一个面向科学计算的高性能动态高级程序设计语言。

Julia 最初是为了满足高性能数值分析和计算科学的需要而设计的，不需要解释器，速度快。

Julia 于 2012 年首次发行，支持各种平台：macOS、Windows、Linux、FreeBSD、Android。

### Julia 语言特点

- 核心语言非常小，标准库用的是 Julia 语言本身写的。
- 调用许多其它成熟的高性能基础代码，如线性代数、随机数生成、快速傅里叶变换、字符串处理。
- 丰富的用于创建或描述对象的类型语法。
- 高性能，接近于静态编译型语言，包括用户自定义类型等。
- 为并行计算和分布式计算而设计。
- 轻量级协程。
- 优雅的可扩展的类型转换/提升。
- 支持 Unicode，包括但不限于 UTF-8。
- 可直接调用 C 函数（不需要包装或是借助特殊的 API）。
- 有类似 shell 的进程管理能力。
- 有类似 Lisp 的宏以及其它元编程工具。
- 可与 Jupyter notebook 一起使用。
- ### Julia 语言用途

Julia 主要功能是用于数值计算。

### 第一个 Julia 程序

接下来我们来编写第一个 Julia 程序 hello.jl（Julia 文件扩展名 .jl），代码如下：

### hello.jl 文件

println("Hello World!")

要执行 Julia 语言代码可以使用 julia hello.jl 命令。

执行以上代码输出:

```

$ julia hello.jl
Hello, World!

```

#### 参考链接

Julia 官网：https://julialang.org/

Julia 中文手册：https://docs.juliacn.com/latest/

---

## Julia 语言环境安装

Source: https://www.runoob.com/julia/julia-environment.html

## Julia 语言环境安装

Julia 语言支持以下系统：

- Linux
- FreeBSD
- macOS
- Windows
- Android

Julia 安装包下载地址为：https://julialang.org/downloads/。

Github 源码地址：https://github.com/JuliaLang/julia。

国内镜像地址：https://mirrors.tuna.tsinghua.edu.cn/julia-releases/bin/

各个系统对应的包名：

操作系统包名 Windowsjulia-1.7.2-win64.exe Linuxx86_64.tar.gz Macjulia-1.7.2-mac64.dmg 安装包或 julia-1.7.2-mac64.tar.gz 二进制文件 FreeBSDjulia-1.7.2-freebsd-x86_64.tar.gz

注：CPU 是 X86 还是 ARM 可以通过 uname -m 命令查看：

```
$ uname -m
x86_64
```

### Windows 系统下安装

从 https://julialang.org/downloads/ 下载 Windows Julia 安装程序。

注意：32 位 Julia 二进制文件可在 32 位和 64 位 Windows（x86 和 x86_64）上运行，但 64 位 Julia 二进制文件只能在 64 位 Windows (x86_64) 上运行，现在电脑也基本上是 64 位的。

运行安装程序，安装过程点击 Next 就好了。

勾选 Add Julia To PATH 自动将 Julia 添加到环境变量。

这样我们就可以在终端执行 Julia 命令了。

Julia 默认安装目录应该类似于 C:\Users\RUNOOB\AppData\Local\Programs\Julia 1.7.2。

### Linux/FreeBSD 安装

以下介绍了在 Linux/FreeBSD 系统下使用二进制包安装方法，下载二进制包：

```
wget https://julialang-s3.julialang.org/bin/linux/x64/1.7/julia-1.7.2-linux-x86_64.tar.gz
```

以上是官网提供的下载地址，如果速度慢，可以采用国内镜像地址下载：

```
wget https://mirrors.tuna.tsinghua.edu.cn/julia-releases/bin/linux/x86/1.7/julia-1.7.2-linux-i686.tar.gz --no-check-certificate
```

解压：

```

tar zxvf julia-1.7.2-linux-i686.tar.gz
```
Linux/FreeBSD 二进制文件不需要安装，但需要我们的系统可以找到 julia 可执行文件，这就需要将 julia 的目录添加到系统环境中。

解压完成后将 julia 的解压目录移动到 /usr/local 目录下：

```
mv julia-1.7.2 /usr/local/
```

移动完成后我们就可以使用 julia 的完整目录执行 Julia 命令：

```
# /usr/local/julia-1.7.2/bin/julia -v
julia version 1.7.2
```

julia -v 命令用于查看版本号。

julia 使用完整路径调用可执行文件：/usr/local/julia-1.7.2/bin/julia -v

也可以将 julia 命令添加到您的系统 PATH 环境变量中，编辑 ~/.bashrc（或 ~/.bash_profile）文件，在最后一行添加以下代码：

```
export PATH="$PATH:/usr/local/julia-1.7.2/bin/"
```

添加后执行以下命令，让环境变量立即生效：

```
source ~/.bashrc
或
source ~/.bash_profile
```
这样我们就可以直接执行 julia 命令而不需要添加完整路径：

```
# julia -v
julia version 1.7.2
```

### macOS 安装

以下介绍了在 macOS 系统下使用二进制包安装方法，下载二进制包：

```
wget https://julialang-s3.julialang.org/bin/mac/x64/1.7/julia-1.7.2-mac64.tar.gz
```

以上是官网提供的下载地址，如果速度慢，可以采用国内镜像地址下载：

```
wget https://mirrors.tuna.tsinghua.edu.cn/julia-releases/bin/mac/x64/1.7/julia-1.7.2-mac64.tar.gz
```

解压：

```

tar zxvf julia-1.7.2-mac64.tar.gz
```
macOS 二进制文件不需要安装，但需要我们的系统可以找到 julia 可执行文件，这就需要将 julia 的目录添加到系统环境中。

解压完成后，可以先把解压包改名为 julia-1.7.2:

```
mv julia-bf53498635 julia-1.7.2
```

完成后将 julia-1.7.2目录移动到 /usr/local 目录下：

```
sudo mv julia-1.7.2 /usr/local/
```

移动完成后我们就可以使用 julia 的完整目录执行 Julia 命令：

```
$ /usr/local/julia-1.7.2/bin/julia -v
julia version 1.7.2
```

julia -v 命令用于查看版本号。

julia 使用完整路径调用可执行文件：/usr/local/julia-1.7.2/bin/julia -v

也可以将 julia 命令添加到您的系统 PATH 环境变量中，编辑 ~/.bash_profile文件，在最后一行添加以下代码：

```
export PATH="$PATH:/usr/local/julia-1.7.2/bin/"
```

添加后执行以下命令，让环境变量立即生效：

```

source ~/.bash_profile
```
这样我们就可以直接执行 julia 命令而不需要添加完整路径：

```
$ julia -v
julia version 1.7.2
```

---

## Julia 交互式命令窗口

Source: https://www.runoob.com/julia/julia-repl.html

## Julia 交互式命令窗口

执行 julia 命令可以直接进入交互式命令窗口：

```
$ julia
_
_ _ _(_)_ | Documentation: https://docs.julialang.org
(_) | (_) (_) |
_ _ _| |_ __ _ | Type "?" for help, "]?" for Pkg help.
| | | | | | |/ _` | |
| | |_| | | | (_| | | Version 1.7.2 (2022-02-06)
_/ |\__'_|_|_|\__'_| | release-1.7/bf53498635 (fork: 461 commits, 259 days)
|__/ |

julia>

```

执行 exit() 退出交互式命令窗口，也可以通过输入 CTRL-D（同时按 Ctrl 键和 d 键）退出。

当然我们也可以执行一个 Julia 的代码文件，文件名以 .jl 结尾。

以下是一个名为 runoob_test.jl 的文件：

### runoob_test.jl 文件

println("Hello World!")
println("RUNOOB")
println(1+1)

要执行 Julia 语言代码可以使用 julia runoob_test.jl 命令。

执行以上代码输出:

```

$ julia runoob_test.jl
Hello World!
RUNOOB
2

```

---

## Julia 基本语法

Source: https://www.runoob.com/julia/julia-basic-syntax.html

## Julia 基本语法

### 变量

变量来源于数学，是计算机语言中能储存计算结果或能表示值抽象概念。

变量可以通过变量名访问。

Julia 语言变量名由字母、数字、下划线 _ 组成，其中首个字符不能为数字。

变量名是大小写敏感的。

Julia 变量使用很简单，只需赋值即可，如下实例：

### 实例

# 将 10 赋值给变量 x
x = 10

# 使用 x 的值做计算
x + 1
11

# 将字符串赋值给变量 site_name
site_name = "RUNOOB"

# 浮点型数据
marks_math = 9.5

交互式命令下输出结果如下：

从实例中我们看到，与其他编程语言 C++、Java 等不同的是 Julia 不需要指定变量类型，它可以自动推断等号右侧的对象类型。

#### 命名规范 建议遵循以下这些命名规范：

- 变量名采用小写。
- 使用下划线 _ 来分变量名中的单词。
- 类型 Type 和模块 Module 的名称使用大写字母开头，并且用大写字母而不是用下划线分隔单词。
- 函数（`function`）和宏（`macro`）的名字使用小写，不使用下划线。
- 会对输入参数进行更改的函数要使用 `!` 结尾。这些函数有时叫做 "mutating" 或 "in-place" 函数，因为它们在被调用后会修改他们的输入参数的内容而不仅仅只是返回一个值。

### 注释

Julia 中的注释有单行注释和多行注释。

Julia 中单行注释以 # 开头，例如：

### 实例

# 这是一行注释
# 这是另外一行注释
println("Hello World!")

多行注释用 #= 与 =# 将注释括起来，例如:

### 实例

#=
1、这是一行注释
2、这是另外一行注释
=#
println("Hello World!")

---

## Julia 数组

Source: https://www.runoob.com/julia/julia-array.html

## Julia 数组

数组，就是相同数据类型的元素按一定顺序排列的集合，可以是一维数组和多维数组。

Julia 支持数组数据结构，它可以存储一个大小不是固定的，类型可以相同也可以不同的顺序集合。

Julia 数组是可变类型集合，用于列表、向量、表格和矩阵。

Julia 数组的索引键值可以使用整数表示，数组的大小不是固定的。

Julia 提供了很多函数帮助我们来操作数组，比如数组添加元素，合并数组等。

Julia 数组用方括号 [...] 指定，多个元素使用逗号 , 分隔。

创建一个一维数组(即一个向量)语法为:

```
[A, B, C, ...]
```

### 创建一维数组

下面实例创建了一个简单的一维数组：

### 实例

julia> arr = [1,2,3]
3-element Vector{Int64}:
1
2
3

上面的实例中我们创建了一个包含 3 个元素的一维数组，每个元素都是一个 64 位整数，这个一维数组绑定到变量 arr 中。

数组元素的类型也可以不一样：

### 实例

julia> arr =[1, "RUNOOB", 2.5, pi]
4-element Vector{Any}:
1
"RUNOOB"
2.5
π = 3.1415926535897...

上面的实例中我们创建了一个包含 4 个元素不同类型的一维数组， pi 是常量 π，每个元素都是一个 64 位整数，这个一维数组绑定到变量 arr 中。

当然也可以强制指定类型：

### 实例

julia> arr = Int64[1,2,3]
3-element Vector{Int64}:
1
2
3
julia> arr2 = String["Taobao","RUNOOB","GOOGLE"]
3-element Vector{String}:
"Taobao"
"RUNOOB"
"GOOGLE"

以上实例数组 arr 限制只能输入整数，arr2 限制只能输入字符串。

我们也可以创建一个空数组：

### 实例

julia> arr = Int64[]
Int64[]

julia> arr2 = String[]
String[]

创建的数组可以直接使用索引值来访问，第一个值的索引为 1（不是 0），第二个值索引为 2，以此类推，最后一个可以使用 end 表示：

### 实例

julia> arr = Int64[1,2,3]
3-element Vector{Int64}:
1
2
3

julia> arr[2]
2
julia> arr2 = String["Taobao","RUNOOB","GOOGLE"]
3-element Vector{String}:
"Taobao"
"RUNOOB"
"GOOGLE"

julia> arr2[1]
"Taobao"

julia> arr2[end]
"GOOGLE"

### 指定数组类型及维度

我们还可以使用以下语法指定数组的类型和维度:

```
Array{type}(undef, dims...)
```

undef 表示数组未初始化。

dims... 可以是维度的单多元组，也可以是维度作为可变参数时的一组值。

dims... 数字表示元素个数，多个维度使用逗号 , 分隔。

### 实例

julia> array = Array{Int64}(undef, 3) # 表示一维数组，数组有 3 个元素
3-element Vector{Int64}:
4834342704
4377305096
0

julia> array = Array{Int64}(undef, 3, 3, 3) # 表示 3 维数组，每个维度数组有 3 个元素
3×3×3 Array{Int64, 3}:
[:, :, 1] =
4562265712 0 0
1 0 0
0 0 0

[:, :, 2] =
0 0 0
0 0 0
0 0 0

[:, :, 3] =
0 0 0
0 0 0
0 0 0

以上实例中，数组的类型我们放在花括号中 {}， undef 用于设置数组未初始化为任何已知值，这就是我们在输出中得到随机数的原因。

### 创建二维数组和矩阵

我们可以将数组元素中的逗号 , 省略掉或者使用两个冒号 ;;，这样就可以创建一个二维数组了，如下实例：

### 实例

julia> [1 2 3 4]
1×4 Matrix{Int64}:
1 2 3 4
julia> [1;; 2;; 3;; 4]
1×4 Matrix{Int64}:
1 2 3 4

注意：第一行输出的 1×4 Matrix{Int64}:，1x4 表示一行四列的矩阵。

虽然只有一行，也是二维数组，因为 Julia 只认可列向量，而不认可所谓的行向量。

要添加另一行，只需添加分号 ;，看以下实例：

### 实例

julia> [1 2; 3 4]
2×2 Matrix{Int64}:
1 2
3 4

也可以使用冒号 : 和空格 来实现，看以下实例：

### 实例

julia> [1:2 3:4]
2×2 Matrix{Int64}:
1 3
2 4

注意：第一行输出的 2×2 Matrix{Int64}:，2×2 表示两行两列的矩阵。

我们也可以在方括号 [] 中嵌入多个长度相同的一维数组，并用空格分隔来创建二维数组：

### 实例

julia> [[1,2] [3,4] [5,6]]
2×3 Matrix{Int64}:
1 3 5
2 4 6

2x3 表示两行三列的数组。

下面我们通过灵活运用分号 ; 和空格 创建一个两行三列和三行两列的二维数组：

### 实例

julia> [[1;2] [3;4] [5;6]]
2×3 Matrix{Int64}:
1 3 5
2 4 6

julia> [[1 2]; [3 4]; [5 6]]
3×2 Matrix{Int64}:
1 2
3 4
5 6

### 使用范围函数来创建数组

#### 省略号 ...

可以使用省略号 ... 来创建一个数组，实例如下：

### 实例

julia> [0:10...]
11-element Vector{Int64}:
0
1
2
3
4
5
6
7
8
9
10

#### collect() 函数

collect() 函数语法格式如下：

```
collect(start:step:stop)
```

start 为开始值，step 为步长，stop 为结束值。

该函数返回数组。

以下实例值为 1，步长为 2，结束值为 13：

### 实例

julia> collect(1:2:13)
7-element Vector{Int64}:
1
3
5
7
9
11
13

collect() 函数也可以指定类型，语法格式如下：

```
collect(element_type, start:step:stop)
```

以下实例创建一个浮点型数组：

### 实例

julia> collect(Float64, 1:2:5)
3-element Vector{Float64}:
1.0
3.0
5.0

#### range() 函数

range() 函数可以生存一个区间范围并指定步长，可以方便 collect() 函数 调用。

range() 函数语法格式如下： range(start, stop, length) range(start, stop; length, step) range(start; length, stop, step) range(;start, length, stop, step)

start 为开始值，step 为步长，stop 为结束值，length 为长度。

### 实例

julia> range(1, length=100)
1:100

julia> range(1, stop=100)
1:100

julia> range(1, step=5, length=100)
1:5:496

julia> range(1, step=5, stop=100)
1:5:96

julia> range(1, 10, length=101)
1.0:0.09:10.0

julia> range(1, 100, step=5)
1:5:96

julia> range(stop=10, length=5)
6:10

julia> range(stop=10, step=1, length=5)
6:1:10

julia> range(start=1, step=1, stop=10)
1:1:10

如果未指定长度 length，且 stop - start 不是 step 的整数倍，则将生成在 stop 之前结束的范围。

```

julia> range(1, 3.5, step=2)
1.0:2.0:3.0
```

使用 range() 和 collect() 创建数组：

### 实例

julia> collect(range(1,stop=10))
10-element Vector{Int64}:
1
2
3
4
5
6
7
8
9
10
julia> collect(range(1, length=15, stop=150))
15-element Vector{Float64}:
1.0
11.642857142857142
22.285714285714285
32.92857142857143
43.57142857142857
54.214285714285715
64.85714285714286
75.5
86.14285714285714
96.78571428571429
107.42857142857143
118.07142857142857
128.71428571428572
139.35714285714286
150.0

#### 使用推导式和生成器创建数组

创建数组的另一种有用方法是使用推导。

数组推导式语法格式如下：

```
A = [ F(x,y,...) for x=rx, y=ry, ... ]
```

F(x,y,...) 取其给定列表中变量 x，y 等的每个值进行计算。值可以指定为任何可迭代对象，但通常是 1:n 或 2:(n-1) 之类的范围，或者像 [1.2, 3.4, 5.7] 这样的显式数组值。结果是一个 N 维密集数组，将变量范围 rx，ry 等的维数拼接起来得到其维数，并且每次 F(x,y,...) 计算返回一个标量。

### 实例

julia> [n^2 for n in 1:10]
10-element Vector{Int64}:
1
4
9
16
25
36
49
64
81
100

创建二维数组：

### 实例

julia> [n*m for n in 1:10, m in 1:10]
10×10 Matrix{Int64}:
1 2 3 4 5 6 7 8 9 10
2 4 6 8 10 12 14 16 18 20
3 6 9 12 15 18 21 24 27 30
4 8 12 16 20 24 28 32 36 40
5 10 15 20 25 30 35 40 45 50
6 12 18 24 30 36 42 48 54 60
7 14 21 28 35 42 49 56 63 70
8 16 24 32 40 48 56 64 72 80
9 18 27 36 45 54 63 72 81 90
10 20 30 40 50 60 70 80 90 100

也可以在没有方括号的情况下编写（数组）推导，从而产生称为生成器的对象。

以下实例创建一个数组：

### 实例

julia> collect(n^2 for n in 1:5)
5-element Vector{Int64}:
1
4
9
16
25

以下表达式在不分配内存的情况下对一个序列进行求和：

### 实例

julia> sum(1/n^2 for n=1:1000)
1.6439345666815615

### Julia 数组基本函数

函数 描述 `eltype(A)` `A` 中元素的类型 `length(A)` `A` 中元素的数量 `ndims(A)` `A` 的维数 `size(A)` 一个包含 `A` 各个维度上元素数量的元组 `size(A,n)` `A` 第 `n` 维中的元素数量 `axes(A)` 一个包含 `A` 有效索引的元组 `axes(A,n)` 第 `n` 维有效索引的范围 `eachindex(A)` 一个访问 `A` 中每一个位置的高效迭代器 `stride(A,k)` 在第 `k` 维上的间隔（stride）（相邻元素间的线性索引距离） `strides(A)` 包含每一维上的间隔（stride）的元组

### Julia构造和初始化

Julia 提供了许多用于构造和初始化数组的函数。在下列函数中，参数 dims ... 可以是一个元组 tuple 来表示维数，也可以是一个可变长度的整数值作为维数。大部分函数的第一个参数都表示数组的元素类型 T 。如果类型 T 被省略，那么将默认为 Float64。

函数 描述 `Array{T}(undef, dims...)` 一个没有初始化的密集 `Array` `zeros(T, dims...)` 一个全零 `Array` `ones(T, dims...)` 一个元素均为 1 的 `Array` `trues(dims...)` 一个每个元素都为 `true` 的 `BitArray` `falses(dims...)` 一个每个元素都为 `false` 的 `BitArray` `reshape(A, dims...)` 一个包含跟 `A` 相同数据但维数不同的数组 `copy(A)` 拷贝 `A` `deepcopy(A)` 深拷贝，即拷贝 `A`，并递归地拷贝其元素 `similar(A, T, dims...)` 一个与`A`具有相同类型（这里指的是密集，稀疏等）的未初始化数组，但具有指定的元素类型和维数。第二个和第三个参数都是可选的，如果省略则默认为元素类型和 `A` 的维数。 `reinterpret(T, A)` 与 `A` 具有相同二进制数据的数组，但元素类型为 `T` `rand(T, dims...)` 一个随机 `Array`，元素值是 [ 0 , 1 ) [0, 1) [0,1) 半开区间中的均匀分布且服从一阶独立同分布 [1] `randn(T, dims...)` 一个随机 `Array`，元素为标准正态分布，服从独立同分布 `Matrix{T}(I, m, n)` `m` 行 `n` 列的单位矩阵 （需要先执行 `using LinearAlgebra` 来才能使用 `I`） `range(start, stop=stop, length=n)` 从 `start` 到 `stop` 的带有 `n` 个线性间隔元素的范围 `fill!(A, x)` 用值 `x` 填充数组 `A` `fill(x, dims...)` 一个被值 `x` 填充的 `Array`

zeros() 创建数组实例，元素初始值 都是 0:

### 实例

julia> zeros(Int8, 2, 3)
2×3 Matrix{Int8}:
0 0 0
0 0 0

julia> zeros(Int8, (2, 3))
2×3 Matrix{Int8}:
0 0 0
0 0 0

julia> zeros((2, 3))
2×3 Matrix{Float64}:
0.0 0.0 0.0
0.0 0.0 0.0

---

## Julia 元组

Source: https://www.runoob.com/julia/julia-tuples.html

## Julia 元组

Julia 的元组与数组类似，都是有序的元素集合，不同之处在于元组的元素不能修改。

另外元组使用小括号 (...)，数组使用方括号 [...]。

元组创建很简单，只需要在括号中添加元素，并使用逗号隔开即可，数组中的很多函数也可以在元组中使用。

如下实例：

### 实例

julia> tupl=(5,10,15,20,25,30) # 创建一个元组
(5, 10, 15, 20, 25, 30)

julia> tupl
(5, 10, 15, 20, 25, 30)

julia> tupl[3:end] # 输出第三个到最后一个元素的元组
(15, 20, 25, 30)

julia> tupl = ((1,2),(3,4)) # 创建二维元组
((1, 2), (3, 4))

julia> tupl[1] # 访问二维元组元素，输出第一维元组
(1, 2)

julia> tupl[1][2] # 访问二维元组元素，输出第一维元组的第二个元素
2

元组的元素是不能修改，如果我们尝试修改它就回报错：

### 实例

julia> tupl2=(1,2,3,4)
(1, 2, 3, 4)

julia> tupl2[2]=0
ERROR: MethodError: no method matching setindex!(::NTuple{4, Int64}, ::Int64, ::Int64)
Stacktrace:
[1] top-level scope
@ REPL[8]:1

### 元组命名

我们可以为元组命名，从而可以更方便的访问它。

以下列出了几种不同元组的命名方式。

#### 1、元组中的键(key)和值(value)分开命名

元组中的键(key)和值(value)可以分开独立命名，实例如下：

### 实例

julia> names_shape = (:corner1, :corner2)
(:corner1, :corner2)

julia> values_shape = ((100, 100), (200, 200))
((100, 100), (200, 200))

julia> shape_item2 = NamedTuple{names_shape}(values_shape)
(corner1 = (100, 100), corner2 = (200, 200))

我们可以使用 . 点号来访问元组：

### 实例

julia> shape_item2.corner1
(100, 100)

julia> shape_item2.corner2
(200, 200)

#### 2、键(key)和值(value)同时在一个元组中

键(key)和值(value)可以同时在一个元组中，实例如下：

### 实例

julia> shape_item = (corner1 = (1, 1), corner2 = (-1, -1), center = (0, 0))
(corner1 = (1, 1), corner2 = (-1, -1), center = (0, 0))

我们可以使用 . 点号来访问元组：

### 实例

julia> shape_item.corner1
(1, 1)

julia> shape_item.corner2
(-1, -1)

julia> shape_item.center
(0, 0)

julia> (shape_item.center,shape_item.corner2)
((0, 0), (-1, -1))

我们还可以像使用普通元组一样访问所有值，如下所示：

### 实例

julia> c1, c2, center = shape_item
(corner1 = (1, 1), corner2 = (-1, -1), center = (0, 0))

julia> c1
(1, 1)

#### 3、合并两个已命名的元组

我们可以使用 merge() 函数来合并两个已命名的元组，实例如下：

### 实例

julia> colors_shape = (top = "red", bottom = "green")
(top = "red", bottom = "green")

julia> shape_item = (corner1 = (1, 1), corner2 = (-1, -1), center = (0, 0))
(corner1 = (1, 1), corner2 = (-1, -1), center = (0, 0))

julia> merge(shape_item, colors_shape)
(corner1 = (1, 1), corner2 = (-1, -1), center = (0, 0), top = "red", bottom = "green")

### 元组作为函数参数

以下实例我们创建一个 testFunc 函数，并将元组 options 作为参数传入：

### 实例：test.jl 文件代码

# 创建函数
function testFunc(x, y, z; a=10, b=20, c=30)
println("x = $x, y = $y, z = $z; a = $a, b = $b, c = $c")
end

# 创建元组
options = (b = 200, c = 300)

# 执行函数，元组作为参数传入
testFunc(1, 2, 3; options...)

使用 julia 命令执行以上文件，输出结果为：

```
$ julia test.jl
x = 1, y = 2, z = 3; a = 10, b = 200, c = 300
```

如果指定的参数在元组后面，则会覆盖元组中已有的参数：

### 实例

# 创建函数
function testFunc(x, y, z; a=10, b=20, c=30)
println("x = $x, y = $y, z = $z; a = $a, b = $b, c = $c")
end

# 创建元组
options = (b = 200, c = 300)

# 执行函数，元组作为参数传入，指定参数在元组前，不会覆盖
testFunc(1, 2, 3; b = 1000_000, options...)

# 执行函数，元组作为参数传入，指定参数在元组后，会覆盖
testFunc(1, 2, 3; options..., b= 1000_000)

使用 julia 命令执行以上文件，输出结果为：

```
$ julia test.jl
x = 1, y = 2, z = 3; a = 10, b = 200, c = 300
x = 1, y = 2, z = 3; a = 10, b = 1000000, c = 300
```

---

## Julia 数据类型

Source: https://www.runoob.com/julia/julia-data-type.html

## Julia 数据类型 在编程语言中，都有基本的数学运算和科学计算，它们常用的数据类型为整数和浮点数。

另外还有一个"字面量"的术语，字面量（literal）用于表达源代码中一个固定值的表示法（notation），整数、浮点数以及字符串等等都是字面量。

例如： a=1 // a 是变量，1 是整型字面量 b=1.0 // b 是变量，1.0 是浮点型字面量

Julia 提供了很丰富的原始数值类型，并基于它们定义了一整套算术运算操作，另外还提供按位运算符以及一些标准数学函数。

### 整数类型

下表列出来 Julia 支持的整数类型：

类型 带符号？ 比特数 最小值 最大值 Int8 ✓ 8 -2^7 2^7 – 1 UInt8 8 0 2^8 – 1 Int16 ✓ 16 -2^15 2^15 – 1 UInt16 16 0 2^16 – 1 Int32 ✓ 32 -2^31 2^31 – 1 UInt32 32 0 2^32 – 1 Int64 ✓ 64 -2^63 2^63 – 1 UInt64 64 0 2^64 – 1 Int128 ✓ 128 -2^127 2^127 – 1 UInt128 128 0 2^128 – 1 Bool N/A 8 false (0) true (1)

整数字面量形式：

### 实例

julia> 1
1

julia> 1234
1234

整型字面量的默认类型取决于目标系统是 32 位还是 64 位架构（目前大部分系统都是 64 位）：

### 实例

# 32 位系统：
julia> typeof(1)
Int32

# 64 位系统：
julia> typeof(1)
Int64

Julia 的内置变量 Sys.WORD_SIZE 表明了目标系统是 32 位还是 64 位架构：

### 实例

# 32 位系统：
julia> Sys.WORD_SIZE
32

# 64 位系统：
julia> Sys.WORD_SIZE
64

Julia 也定义了 Int 与 UInt 类型，它们分别是系统有符号和无符号的原生整数类型的别名。

### 实例

# 32 位系统：
julia> Int
Int32
julia> UInt
UInt32

# 64 位系统：
julia> Int
Int64
julia> UInt
UInt64

#### 溢出行为

在 Julia 里，超出一个类型可表示的最大值会导致环绕 (wraparound) 行为：

### 实例

julia> x = typemax(Int64)
9223372036854775807

julia> x + 1
-9223372036854775808

julia> x + 1 == typemin(Int64)
true

因此，Julia 的整数算术实际上是模算数的一种形式，它反映了现代计算机实现底层算术的特点。在可能有溢出产生的程序中，对最值边界出现循环进行显式检查是必要的。否则，推荐使用任意精度算术中的 BigInt 类型作为替代。

下面是溢出行为的一个例子以及如何解决溢出：

### 实例

julia> 10^19
-8446744073709551616

julia> big(10)^19
10000000000000000000

#### 除法错误 在以下两种例外情况下，整数除法会触发 DivideError 错误：

- 除以零
- 除以最小的负数

rem 取余函数和 mod 取模函数在除零时抛出 DivideError 错误，实例如下：

### 实例

julia> mod(1, 0)
ERROR: DivideError: integer division error
Stacktrace:
[1] div at .\int.jl:260 [inlined]
[2] div at .\div.jl:217 [inlined]
[3] div at .\div.jl:262 [inlined]
[4] fld at .\div.jl:228 [inlined]
[5] mod(::Int64, ::Int64) at .\int.jl:252
[6] top-level scope at REPL[52]:1

julia> rem(1, 0)
ERROR: DivideError: integer division error
Stacktrace:
[1] rem(::Int64, ::Int64) at .\int.jl:261
[2] top-level scope at REPL[54]:1

### 浮点类型

下表列出来 Julia 支持的浮点类型:

类型 精度 比特数 Float16 半精度 16 Float32 单精度 32 Float64 双精度 64

此外，对复数和有理数的完整支持是在这些原始数据类型之上建立起来的。

浮点数字面量格式表示如下，必要时可使用 E 来表示。

### 实例

julia> 1.0
1.0

julia> 1.
1.0

julia> 0.5
0.5

julia> .5
0.5

julia> -1.23
-1.23

julia> 1e10
1.0e10

julia> 2.5e-4
0.00025

注：在科学计数法中，为了使公式简便，可以用带 E 的格式表示。例如 1.03乘10的8次方，可简写为 "1.03E+08" 的形式，其中 "E" 是 exponent(指数) 的缩写。

上面的结果都是 Float64 类型的值。使用 f 替代 e 可以得到 Float32 类型的字面量：

### 实例

julia> x = 0.5f0
0.5f0

julia> typeof(x)
Float32

julia> 2.5f-4
0.00025f0
数值可以很容易地转换为 Float32 类型：

julia> x = Float32(-1.5)
-1.5f0

julia> typeof(x)
Float32

也存在十六进制的浮点数字面量，但只适用于 Float64 类型的值。一般使用 p 前缀及以 2 为底的指数来表示：

### 实例

julia> 0x1p0
1.0

julia> 0x1.8p3
12.0

julia> x = 0x.4p-1
0.125

julia> typeof(x)
Float64
Julia 也支持半精度浮点数（Float16），但它们是使用 Float32 进行软件模拟实现的。

julia> sizeof(Float16(4.))
2

julia> 2*Float16(4.)
Float16(8.0)

下划线 _ 可用作数字分隔符：

### 实例

julia> 10_000, 0.000_000_005, 0xdead_beef, 0b1011_0010
(10000, 5.0e-9, 0xdeadbeef, 0xb2)

#### 浮点数中的零

浮点数有两种零，正零和负零。它们相互相等但有着不同的二进制表示，可以使用 bitstring 函数来查看：

### 实例

julia> 0.0 == -0.0
true

julia> bitstring(0.0)
"0000000000000000000000000000000000000000000000000000000000000000"

julia> bitstring(-0.0)
"1000000000000000000000000000000000000000000000000000000000000000"

#### 特殊的浮点值

有三种特定的标准浮点值不和实数轴上任何一点对应：

Float16 Float32 Float64 名称 描述 Inf16 Inf32 Inf 正无穷 一个大于所有有限浮点数的数 -Inf16 -Inf32 -Inf 负无穷 一个小于所有有限浮点数的数 NaN16 NaN32 NaN 不是一个数 一个不和任何浮点值（包括自己）相等（==）的值

以下列举了一些浮点数的运算实例：

### 实例

julia> 1/Inf
0.0

julia> 1/0
Inf

julia> -5/0
-Inf

julia> 0.000001/0
Inf

julia> 0/0
NaN

julia> 500 + Inf
Inf

julia> 500 - Inf
-Inf

julia> Inf + Inf
Inf

julia> Inf - Inf
NaN

julia> Inf * Inf
Inf

julia> Inf / Inf
NaN

julia> 0 * Inf
NaN

julia> NaN == NaN
false

julia> NaN != NaN
true

julia> NaN < NaN
false

julia> NaN > NaN
false

我们还可以使用 typemin 和 typemax 函数：

### 实例

julia> (typemin(Float16),typemax(Float16))
(-Inf16, Inf16)

julia> (typemin(Float32),typemax(Float32))
(-Inf32, Inf32)

julia> (typemin(Float64),typemax(Float64))
(-Inf, Inf)

#### 机器精度

大多数实数都无法用浮点数准确地表示，因此有必要知道两个相邻可表示的浮点数间的距离，它通常被叫做机器精度。

Julia 提供了 eps 函数，它可以给出 1.0 与下一个 Julia 能表示的浮点数之间的差值：

### 实例

julia> eps(Float32)
1.1920929f-7

julia> eps(Float64)
2.220446049250313e-16

julia> eps() # 与 eps(Float64) 相同
2.220446049250313e-16

这些值分别是 Float32 中的 2.0^-23 和 Float64 中的 2.0^-52。eps 函数也可以接受一个浮点值作为参数，然后给出这个值与下一个可表示的浮点数值之间的绝对差。也就是说，eps(x) 产生一个和 x 类型相同的值，并且 x + eps(x) 恰好是比 x 更大的下一个可表示的浮点值：

### 实例

julia> eps(1.0)
2.220446049250313e-16

julia> eps(1000.)
1.1368683772161603e-13

julia> eps(1e-27)
1.793662034335766e-43

julia> eps(0.0)
5.0e-324

两个相邻可表示的浮点数之间的距离并不是常数，数值越小，间距越小，数值越大，间距越大。换句话说，可表示的浮点数在实数轴上的零点附近最稠密，并沿着远离零点的方向以指数型的速度变得越来越稀疏。根据定义，eps(1.0) 与 eps(Float64) 相等，因为 1.0 是个 64 位浮点值。

Julia 也提供了 nextfloat 和 prevfloat 两个函数分别返回基于参数的下一个更大或更小的可表示的浮点数：

### 实例

julia> x = 1.25f0
1.25f0

julia> nextfloat(x)
1.2500001f0

julia> prevfloat(x)
1.2499999f0

julia> bitstring(prevfloat(x))
"00111111100111111111111111111111"

julia> bitstring(x)
"00111111101000000000000000000000"

julia> bitstring(nextfloat(x))
"00111111101000000000000000000001"

这个例子体现了一般原则，即相邻可表示的浮点数也有着相邻的二进制整数表示。

### 舍入模式

一个数如果没有精确的浮点表示，就必须被舍入到一个合适的可表示的值。

Julia 所使用的默认模式总是 RoundNearest，指舍入到最接近的可表示的值，这个被舍入的值会使用尽量少的有效位数。

### 实例

julia> BigFloat("1.510564889",2,RoundNearest)
1.5

julia> BigFloat("1.550564889",2,RoundNearest)
1.5

julia> BigFloat("1.560564889",2,RoundNearest)
1.5

### 0 和 1 的字面量

Julia 提供了 0 和 1 的字面量函数，可以返回特定类型或所给变量的类型。

函数描述 zero(x) x 类型或变量 x 的类型的零字面量 one(x) x 类型或变量 x 的类型的一字面量

这些函数在数值比较中可以用来避免不必要的类型转换带来的开销。

例如：

### 实例

julia> zero(Float32)
0.0f0

julia> zero(1.0)
0.0

julia> one(Int32)
1

julia> one(BigFloat)
1.0

### 类型转换

类型转换是把变量从一种类型转换为另一种数据类型。例如，如果您想存储一个 float 类型的值到一个简单的整型中，您需要把 float 类型强制转换为 int 类型。您可以使用强制类型转换运算符来把值显式地从一种类型转换为另一种类型，如下所示： Julia 支持三种数值转换，它们在处理不精确转换上有所不同。

第一种：

```
T(x)
或
convert(T,x)
```

以上都会把 x 转换为 T 类型。

- 如果 T 是浮点类型，转换的结果就是最近的可表示值， 可能会是正负无穷大。
- 如果 T 为整数类型，当 x 不能由 T 类型表示时，会抛出 InexactError。

第二种：

x % T 也可以将整数 x 转换为整型 T，与 x 模 2^n 的结果一致，其中 n 是 T 的位数。

第三种：

舍入函数接收一个 T 类型的可选参数。比如，round(Int,x) 是 Int(round(x)) 的简写版。

### 实例

julia> Int8(127)
127

julia> Int8(128)
ERROR: InexactError: trunc(Int8, 128)
Stacktrace:
[...]

julia> Int8(127.0)
127

julia> Int8(3.14)
ERROR: InexactError: Int8(3.14)
Stacktrace:
[...]

julia> Int8(128.0)
ERROR: InexactError: Int8(128.0)
Stacktrace:
[...]

julia> 127 % Int8
127

julia> 128 % Int8
-128

julia> round(Int8,127.4)
127

julia> round(Int8,127.6)
ERROR: InexactError: trunc(Int8, 128.0)
Stacktrace:
[...]

---

## Julia 复数和有理数

Source: https://www.runoob.com/julia/julia-complex-and-rational-numbers.html

## Julia 复数和有理数

本章节我们主要要来学习 Julia 的复数和有理数。

Julia 语言包含了预定义的复数和有理数类型，并且支持它们的各种标准数学运算和初等函数。

#### 复数

复数，为实数的延伸，它使任一多项式方程都有根。

我们把形如 z=a+bi（a、b均为实数）的数称为复数。其中，a 称为实部，b 称为虚部，i 称为虚数单位，它有着性质。当 z 的虚部 b＝0 时，则 z 为实数；当 z 的虚部 b≠0 时，实部 a＝0 时，常称 z 为纯虚数。

全局常量 im 被绑定到复数 i，表示 -1 的主平方根。

由于 Julia 允许数值字面量作为数值字面量系数，这种绑定就足以为复数提供很方便的语法，类似于传统的数学记法：

### 实例

julia> 1+2im
1 + 2im

我们也可以对复数进行各种算术操作：

### 实例

julia> (1 + 2im)*(2 - 3im)
8 + 1im

julia> (1 + 2im)/(1 - 2im)
-0.6 + 0.8im

julia> (1 + 2im) + (1 - 2im)
2 + 0im

julia> (-3 + 2im) - (5 - 1im)
-8 + 3im

julia> (-1 + 2im)^2
-3 - 4im

julia> (-1 + 2im)^2.5
2.729624464784009 - 6.9606644595719im

julia> (-1 + 2im)^(1 + 1im)
-0.27910381075826657 + 0.08708053414102428im

julia> 3(2 - 5im)
6 - 15im

julia> 3(2 - 5im)^2
-63 - 60im

julia> 3(2 - 5im)^-1.0
0.20689655172413796 + 0.5172413793103449im

类型提升机制也确保你可以使用不同类型的操作数的组合：

### 实例

julia> 2(1 - 1im)
2 - 2im

julia> (2 + 3im) - 1
1 + 3im

julia> (1 + 2im) + 0.5
1.5 + 2.0im

julia> (2 + 3im) - 0.5im
2.0 + 2.5im

julia> 0.75(1 + 2im)
0.75 + 1.5im

julia> (2 + 3im) / 2
1.0 + 1.5im

julia> (1 - 3im) / (2 + 2im)
-0.5 - 1.0im

julia> 2im^2
-2 + 0im

julia> 1 + 3/4im
1.0 - 0.75im

注意 3/4im == 3/(4*im) == -(3/4*im)，因为系数比除法的优先级更高。

Julia 提供了一些操作复数的标准函数：

### 实例

julia> z = 1 + 2im
1 + 2im

julia> real(1 + 2im) # z 的实部
1

julia> imag(1 + 2im) # z 的虚部
2

julia> conj(1 + 2im) # z 的复共轭
1 - 2im

julia> abs(1 + 2im) # z 的绝对值
2.23606797749979

julia> abs2(1 + 2im) # 取平方后的绝对值
5

julia> angle(1 + 2im) # 以弧度为单位的相位角
1.1071487177940904

按照惯例，复数的绝对值（abs）是从零点到它的距离。abs2 给出绝对值的平方，作用于复数上时非常有用，因为它避免了取平方根。angle 返回以弧度为单位的相位角（也被称为辐角函数）。所有其它的初等函数在复数上也都有完整的定义：

### 实例

julia> sqrt(1im)
0.7071067811865476 + 0.7071067811865475im

julia> sqrt(1 + 2im)
1.272019649514069 + 0.7861513777574233im

julia> cos(1 + 2im)
2.0327230070196656 - 3.0518977991517997im

julia> exp(1 + 2im)
-1.1312043837568135 + 2.4717266720048188im

julia> sinh(1 + 2im)
-0.4890562590412937 + 1.4031192506220405im

注意数学函数通常应用于实数就返回实数值，应用于复数就返回复数值。例如，当 sqrt 应用于 -1 与 -1 + 0im 会有不同的表现，虽然 -1 == -1 + 0im：

### 实例

julia> sqrt(-1)
ERROR: DomainError with -1.0:
sqrt will only return a complex result if called with a complex argument. Try sqrt(Complex(x)).
Stacktrace:
[...]

julia> sqrt(-1 + 0im)
0.0 + 1.0im

从变量构建复数时，文本型数值系数记法不再适用。相反地，乘法必须显式地写出：

### 实例

julia> a = 1; b = 2; a + b*im
1 + 2im

然而，我们并不推荐这样做，而应改为使用更高效的 complex 函数直接通过实部与虚部构建一个复数值：

### 实例

julia> a = 1; b = 2; complex(a, b)
1 + 2im

这种构建避免了乘法和加法操作。

Inf 和 NaN 可能出现在复数的实部和虚部，正如特殊的浮点值章节所描述的：

### 实例

julia> 1 + Inf*im
1.0 + Inf*im

julia> 1 + NaN*im
1.0 + NaN*im

#### 有理数

有理数是整数（正整数、0、负整数）和分数的统称，是整数和分数的集合。 数学上，可以表达为两个整数比的数(, )被定义为有理数，例如，0.75(可被表达为)。整数和分数统称为有理数。与有理数相对的是无理数，如无法用整数比表示。

Julia 有一个用于表示整数精确比值的分数类型。分数通过 // 运算符构建：

### 实例

julia> 2//3
2//3

如果一个分数的分子和分母含有公因子，它们会被约分到最简形式且分母非负：

### 实例

julia> 6//9
2//3

julia> -4//8
-1//2

julia> 5//-15
-1//3

julia> -4//-12
1//3

整数比值的这种标准化形式是唯一的，所以分数值的相等性可由校验分子与分母都相等来测试。分数值的标准化分子和分母可以使用 numerator 和 denominator 函数得到：

### 实例

julia> numerator(2//3)
2

julia> denominator(2//3)
3

分子和分母的直接比较通常是不必要的，因为标准算术和比较操作对分数值也有定义：

### 实例

julia> 2//3 == 6//9
true

julia> 2//3 == 9//27
false

julia> 3//7 < 1//2
true

julia> 3//4 > 2//3
true

julia> 2//4 + 1//6
2//3

julia> 5//12 - 1//4
1//6

julia> 5//8 * 3//12
5//32

julia> 6//5 / 10//7
21//25

分数可以很容易地转换成浮点数：

### 实例

julia> float(3//4)
0.75

对任意整数值 a 和 b（除了 a == 0 且 b == 0 时），从分数到浮点数的转换遵从以下的一致性：

### 实例

julia> a = 1; b = 2;

julia> isequal(float(a//b), a/b)
true

Julia接受构建无穷分数值：

### 实例

julia> 5//0
1//0

julia> x = -3//0
-1//0

julia> typeof(x)
Rational{Int64}

但不接受试图构建一个 NaN 分数值：

### 实例

julia> 0//0
ERROR: ArgumentError: invalid rational: zero(Int64)//zero(Int64)
Stacktrace:
[...]

像往常一样，类型提升系统使得分数可以轻松地同其它数值类型进行交互：

### 实例

julia> 3//5 + 1
8//5

julia> 3//5 - 0.5
0.09999999999999998

julia> 2//7 * (1 + 2im)
2//7 + 4//7*im

julia> 2//7 * (1.5 + 2im)
0.42857142857142855 + 0.5714285714285714im

julia> 3//2 / (1 + 2im)
3//10 - 3//5*im

julia> 1//2 + 2im
1//2 + 2//1*im

julia> 1 + 2//3im
1//1 - 2//3*im

julia> 0.5 == 1//2
true

julia> 0.33 == 1//3
false

julia> 0.33 < 1//3
true

julia> 1//3 - 0.33
0.0033333333333332993

---

## Julia 基本运算符

Source: https://www.runoob.com/julia/julia-basic-operators.html

## Julia 基本运算符

运算符是一种告诉编译器执行特定的数学或逻辑操作的符号，如: 3+2=5。

Julia 语言内置了丰富的运算符，支持的运算有：

- 算术运算符
- 逻辑运算符
- 关系运算符
- 位运算符
- 赋值运算符
- 向量化 "点" 运算符

### 算术运算符

下表显示了 Julia 的基本算术运算符，适用于所有的基本数值类型：

表达式 名称 描述 `+x` 一元加法运算符 全等操作 `-x` 一元减法运算符 将值变为其相反数 `x + y` 二元加法运算符 两数相加 `x - y` 二元减法运算符 两数相减 `x * y` 乘法运算符 两数相乘 `x / y` 除法运算符 两数相除 `x ÷ y` 整除 取 x / y 的整数部分 `x \ y` 反向除法 等价于 `y / x` `x ^ y` 幂操作符 `x` 的 `y` 次幂 `x % y` 取余 等价于 `rem(x,y)`

### 实例

julia> 1 + 2 + 3
6
julia> 1 - 2
-1
julia> 3*2/12
0.5
julia> 2+20-5
17
julia> 50*2/10
10.0
julia> 23%2
1
julia> 2^4
16

### 布尔运算符

下表显示了 Julia 的布尔运算符：

表达式 名称 `!x` 否定 `x && y` 短路与，在表达式 x && y 中，子表达式 y 仅当 x 为 true 的时候才会被执行。 `x || y` 短路或，在表达式 x || y 中，子表达式 y 仅在 x 为 false 的时候才会被执行。

### 实例

julia> !true
false

julia> !false
true

julia> true && (x = (1, 2, 3))
(1, 2, 3)

julia> false && (x = (1, 2, 3))
false

julia> false || (x = (1, 2, 3))
(1, 2, 3)

### 关系运算符

下表显示了 Julia 的关系运算符：

操作符 名称 `==` 相等 `!=`, `≠` 不等 `<` 小于 `<=`, `≤` 小于等于 `>` 大于 `>=`, `≥` 大于等于

### 实例

julia> 100 == 100
true

julia> 100 == 101
false

julia> 100 != 101
true

julia> 100 == 100.0
true

julia> 100 < 500
true

julia> 100 > 500
false

julia> 100 >= 100.0
true

julia> -100 <= 100
true

julia> -100 <= -100
true

julia> -100 <= -500
false

julia> 100 < -10.0
false

#### 链式比较

链式比较在写数值代码时特别方便，它使用 && 运算符比较标量，数组则用 & 进行按元素比较。比如，0 .< A .< 1 会得到一个 boolean 数组，如果 A 的元素都在 0 和 1 之间则数组元素就都是 true。

Julia 允许链式比较：

### 实例

julia> 1 < 2 <= 2 < 3 == 3 > 2 >= 1 == 1 < 3 != 5
true

注意链式比较的执行顺序：

### 实例

julia> M(a) = (println(a); a)
M (generic function with 1 method)
julia> M(1) < M(2) <= M(3)
2
1
3
true
julia> M(1) > M(2) <= M(3)
2
1
false

### 位运算符

下表显示了 Julia 的位运算符：

表达式 名称 `~x` 按位取反 `x & y` 按位与 `x | y` 按位或 `x ⊻ y` 按位异或（逻辑异或） `x ⊼ y` 按位与（非与） `x ⊽ y` 按位或（非或） `x >>> y` 逻辑右移 `x >> y` 算术右移 `x << y` 逻辑/算术左移

### 实例

julia> ~123
-124

julia> 123 & 234
106

julia> 123 | 234
251

julia> 123 ⊻ 234
145

julia> xor(123, 234)
145

julia> nand(123, 123)
-124

julia> 123 ⊼ 123
-124

julia> nor(123, 124)
-128

julia> 123 ⊽ 124
-128

julia> ~UInt32(123)
0xffffff84

julia> ~UInt8(123)
0x84

### 赋值运算符 每一个二元运算符和位运算符都可以给左操作数复合赋值，方法是把 = 直接放在二元运算符后面。比如，x += 3 等价于 x = x + 3 。

下表显示了 Julia 的赋值运算符：

运算符描述实例 =简单的赋值运算符，把右边操作数的值赋给左边操作数 C = A + B 将把 A + B 的值赋给 C +=加且赋值运算符，把右边操作数加上左边操作数的结果赋值给左边操作数 C += A 相当于 C = C + A -=减且赋值运算符，把左边操作数减去右边操作数的结果赋值给左边操作数 C -= A 相当于 C = C - A *=乘且赋值运算符，把右边操作数乘以左边操作数的结果赋值给左边操作数 C *= A 相当于 C = C * A /=除且赋值运算符，把左边操作数除以右边操作数的结果赋值给左边操作数 C /= A 相当于 C = C / A %=求模且赋值运算符，求两个操作数的模赋值给左边操作数 C %= A 相当于 C = C % A <<=左移且赋值运算符 C <<= 2 等同于 C = C << 2 >>=右移且赋值运算符 C >>= 2 等同于 C = C >> 2 &=按位与且赋值运算符 C &= 2 等同于 C = C & 2 ^=按位异或且赋值运算符 C ^= 2 等同于 C = C ^ 2 |=按位或且赋值运算符 C |= 2 等同于 C = C | 2 >>>=左移且逻辑运算符 C >>>= 2 等同于 C = C >>>2 ⊻=异或（逻辑异或）赋值运算符 C ⊻= 2 等同于 C = C ⊻= 2

### 实例

julia> x = 1
1

julia> x += 3
4

julia> x
4

### 向量化 "点" 运算符

Julia 中，每个二元运算符都有一个 "点" 运算符与之对应，例如 ^ 就有对应的 .^ 存在。这个对应的 .^ 被 Julia 自动地定义为逐元素地执行 ^ 运算。比如 [1,2,3] ^ 3 是非法的，因为数学上没有给（长宽不一样的）数组的立方下过定义。但是 [1,2,3] .^ 3 在 Julia 里是合法的，它会逐元素地执行 ^ 运算（或称向量化运算），得到 [1^3, 2^3, 3^3]。类似地，! 或 √ 这样的一元运算符，也都有一个对应的 .√ 用于执行逐元素运算。

### 实例

julia> [1,2,3] .^ 3
3-element Vector{Int64}:
1
8
27

除了点运算符，我们还有逐点赋值运算符，类似 a .+= b（或者 @. a += b）会被解析成 a .= a .+ b。

### 运算符的优先级与结合性

运算符的优先级确定表达式中项的组合。这会影响到一个表达式如何计算。某些运算符比其他运算符有更高的优先级，例如，乘除运算符具有比加减运算符更高的优先级。

例如 x = 7 + 3 * 2，在这里，x 被赋值为 13，而不是 20，因为运算符 * 具有比 + 更高的优先级，所以首先计算乘法 3*2，然后再加上 7。

下表将按运算符优先级从高到低列出各个运算符，具有较高优先级的运算符出现在表格的上面，具有较低优先级的运算符出现在表格的下面。在表达式中，较高优先级的运算符会优先被计算。

分类 运算符 结合性 语法 `.` followed by `::` 左结合 幂运算 `^` 右结合 一元运算符 `+ - √` 右结合 移位运算 `<< >> >>>` 左结合 除法 `//` 左结合 乘法 `* / % & \ ÷` 左结合 加法 `+ - | ⊻` 左结合 语法 `: ..` 左结合 语法 `|>` 左结合 语法 `<|` 右结合 比较 `> < >= <= == === != !== <:` 无结合性 流程控制 `&&` followed by `||` followed by `?` 右结合 Pair 操作 `=>` 右结合 赋值 `= += -= *= /= //= \= ^= ÷= %= |= &= ⊻= <<= >>= >>>=` 右结合

我们也可以通过内置函数 Base.operator_precedence 查看任何给定运算符的优先级数值，数值越大优先级越高：

### 实例

julia> Base.operator_precedence(:+), Base.operator_precedence(:*), Base.operator_precedence(:.)
(11, 12, 17)

julia> Base.operator_precedence(:sin), Base.operator_precedence(:+=), Base.operator_precedence(:(=)) # (注意：等号前后必须有括号 `:(=)`)
(0, 1, 1)

另外，内置函数 Base.operator_associativity 可以返回运算符结合性的符号表示：

### 实例

julia> Base.operator_associativity(:-), Base.operator_associativity(:+), Base.operator_associativity(:^)
(:left, :none, :right)

julia> Base.operator_associativity(:⊗), Base.operator_associativity(:sin), Base.operator_associativity(:→)
(:left, :none, :right)

---

## Julia 数学函数

Source: https://www.runoob.com/julia/julia-mathematical-functions.html

## Julia 数学函数

Julia 提供了一套高效、可移植的标准数学函数。

### 数值比较

下表列出了用于数值比较的函数：

函数 测试是否满足如下性质 `isequal(x, y)` `x` 与 `y` 值与类型是否完全相同 `isfinite(x)` `x` 是否是有限大的数字 `isinf(x)` `x` 是否是（正/负）无穷大 `isnan(x)` `x` 是否是 `NaN`

isequal 认为 NaN 之间是相等的：

### 实例

julia> isequal(NaN, NaN)
true

julia> isequal([1 NaN], [1 NaN])
true

julia> isequal(NaN, NaN32)
true

isequal 也能用来区分带符号的零：

### 实例

julia> -0.0 == 0.0
true

julia> isequal(-0.0, 0.0)
false

其他函数实例：

### 实例

julia> isfinite(5)
true

julia> isfinite(NaN32)
false

### 舍入函数

下表列出了 Julia 支持的舍入函数：

函数 描述 返回类型 `round(x)` `x` 舍到最接近的整数 `typeof(x)` `round(T, x)` `x` 舍到最接近的整数 `T` `floor(x)` `x` 向 `-Inf` 舍入 `typeof(x)` `floor(T, x)` `x` 向 `-Inf` 舍入 `T` `ceil(x)` `x` 向 `+Inf` 方向取整 `typeof(x)` `ceil(T, x)` `x` 向 `+Inf` 方向取整 `T` `trunc(x)` `x` 向 0 取整 `typeof(x)` `trunc(T, x)` `x` 向 0 取整 `T`

### 实例

julia> round(3.8)
4.0
julia> round(Int, 3.8)
4
julia> floor(3.8)
3.0
julia> floor(Int, 3.8)
3
julia> ceil(3.8)
4.0
julia> ceil(Int, 3.8)
4
julia> trunc(3.8)
3.0
julia> trunc(Int, 3.8)
3

### 除法函数

下表列出了 Julia 支持的除法函数：

函数 描述 `div(x,y)`, `x÷y` 截断除法，无论任何类型相除的结果都会省略小数部分，剩下整数部分，商向零近似。 `fld(x,y)` 向下取整除法；商向 `-Inf` 近似 `cld(x,y)` 向上取整除法；商向 `+Inf` 近似 `rem(x,y)` 取余；满足 `x == div(x,y)*y + rem(x,y)`；符号与 `x` 一致 `mod(x,y)` 取模；满足 `x == fld(x,y)*y + mod(x,y)`；符号与 `y` 一致 `mod1(x,y)` 偏移 1 的 `mod`；若 `y>0`，则返回 `r∈(0,y]`，若 `y<0`，则 `r∈[y,0)` 且满足 `mod(r, y) == mod(x, y)` `mod2pi(x)` 对 2pi 取模；`0 <= mod2pi(x) < 2pi` `divrem(x,y)` 返回 `(div(x,y),rem(x,y))` `fldmod(x,y)` 返回 `(fld(x,y),mod(x,y))` `gcd(x,y...)` `x`, `y`,... 的最大公约数 `lcm(x,y...)` `x`, `y`,... 的最小公倍数

### 实例

julia> div(11, 4)
2

julia> div(7, 4)
1

julia> fld(11, 4)
2

julia> fld(-5,3)
-2

julia> fld(7.5,3.3)
2.0

julia> cld(7.5,3.3)
3.0

julia> mod(5, 0:2)
2

julia> mod(3, 0:2)
0

julia> mod(8.9,2)
0.9000000000000004

julia> rem(8,4)
0

julia> rem(9,4)
1

julia> mod2pi(7*pi/5)
4.39822971502571

julia> divrem(8,3)
(2, 2)

julia> fldmod(12,4)
(3, 0)

julia> fldmod(13,4)
(3, 1)

julia> mod1(5,4)
1

julia> gcd(6,0)
6

julia> gcd(1//3,2//3)
1//3

julia> lcm(1//3,2//3)
2//3

### 符号和绝对值函数

下表列出了 Julia 支持的符号和绝对值函数：

函数 描述 `abs(x)` `x` 的模 `abs2(x)` `x` 的模的平方 `sign(x)` 表示 `x` 的符号，返回 -1，0，或 +1 `signbit(x)` 表示符号位是 true 或 false `copysign(x,y)` 返回一个数，其值等于 `x` 的模，符号与 `y` 一致 `flipsign(x,y)` 返回一个数，其值等于 `x` 的模，符号与 `x*y` 一致

### 实例

julia> abs(-7)
7

julia> abs(5+3im)
5.830951894845301

julia> abs2(-7)
49

julia> abs2(5+3im)
34

julia> copysign(5,-10)
-5

julia> copysign(-5,10)
5

julia> sign(5)
1

julia> sign(-5)
-1

julia> signbit(-5)
true

julia> signbit(5)
false

julia> flipsign(5,10)
5

julia> flipsign(5,-10)
-5

### 符号和绝对值函数

下表列出了 Julia 支持的符号和绝对值函数：

函数 描述 `sqrt(x)`, `√x` `x` 的平方根 `cbrt(x)`, `∛x` `x` 的立方根 `hypot(x,y)` 当直角边的长度为 `x` 和 `y`时，直角三角形斜边的长度 `exp(x)` 自然指数函数在 `x` 处的值 `expm1(x)` 当 `x` 接近 0 时的 `exp(x)-1` 的精确值 `ldexp(x,n)` `x*2^n` 的高效算法，`n` 为整数 `log(x)` `x` 的自然对数 `log(b,x)` 以 `b` 为底 `x` 的对数 `log2(x)` 以 2 为底 `x` 的对数 `log10(x)` 以 10 为底 `x` 的对数 `log1p(x)` 当 `x`接近 0 时的 `log(1+x)` 的精确值 `exponent(x)` `x` 的二进制指数 `significand(x)` 浮点数 `x` 的二进制有效数（也就是尾数）

### 实例

julia> sqrt(49)
7.0

julia> sqrt(-49)
ERROR: DomainError with -49.0:
sqrt will only return a complex result if called with a complex argument. Try sqrt(Complex(x)).
Stacktrace:
[1] throw_complex_domainerror(::Symbol, ::Float64) at .\math.jl:33
[2] sqrt at .\math.jl:573 [inlined]
[3] sqrt(::Int64) at .\math.jl:599
[4] top-level scope at REPL[43]:1

julia> cbrt(8)
2.0

julia> cbrt(-8)
-2.0

julia> a = Int64(5)^10;

julia> hypot(a, a)
1.3810679320049757e7

julia> exp(5.0)
148.4131591025766

julia> expm1(10)
22025.465794806718

julia> expm1(1.0)
1.718281828459045

julia> ldexp(4.0, 2)
16.0

julia> log(5,2)
0.43067655807339306

julia> log(4,2)
0.5

julia> log(4)
1.3862943611198906

julia> log2(4)
2.0

julia> log10(4)
0.6020599913279624

julia> log1p(4)
1.6094379124341003

julia> log1p(-2)
ERROR: DomainError with -2.0:
log1p will only return a complex result if called with a complex argument. Try log1p(Complex(x)).
Stacktrace:
[1] throw_complex_domainerror(::Symbol, ::Float64) at .\math.jl:33
[2] log1p(::Float64) at .\special\log.jl:356
[3] log1p(::Int64) at .\special\log.jl:395
[4] top-level scope at REPL[65]:1
julia> exponent(6.8)
2

julia> significand(15.2)/10.2
0.18627450980392157

julia> significand(15.2)*8
15.2

### 三角和双曲函数

Julia 也提供了所有标准的三角和双曲函数：

```
sin cos tan cot sec csc
sinh cosh tanh coth sech csch
asin acos atan acot asec acsc
asinh acosh atanh acoth asech acsch
sinc cosc
```

下图中以弧度为单位的角度对应于单位圆上的一个点，其坐标定义了角度的正弦和余弦。

### 实例

julia> pi
π = 3.1415926535897...

julia> sin(0)
0.0

julia> sin(pi/6)
0.49999999999999994

julia> sin(pi/4)
0.7071067811865475

julia> cos(0)
1.0

julia> cos(pi/6)
0.8660254037844387

julia> cos(pi/3)
0.5000000000000001

以上提供的函数都是单参数函数，不过 atan 也可以接收两个参数 来表示传统的 atan2 函数。

```
atan(y)
atan(y, x)
```

分别计算 y 或 y/x 的反正切。

### 实例

julia> theta = 3pi/4
2.356194490192345

julia> x,y = (cos(theta), sin(theta))
(-0.7071067811865475, 0.7071067811865476)

julia> atan(y/x)
-0.7853981633974484

julia> atan(y, x)
2.356194490192345

另外，sinpi(x) 和 cospi(x) 分别用来对 sin(pi*x) 和 cos(pi*x) 进行更精确的计算。

要计算角度而非弧度的三角函数，以 d 做后缀。 比如，sind(x) 计算 x 的 sine 值，其中 x 是一个角度值。 下面是角度变量的三角函数完整列表：

```
sind cosd tand cotd secd cscd
asind acosd atand acotd asecd acscd
```

### 实例

julia> cos(56)
0.853220107722584

julia> cosd(56)
0.5591929034707468

---

## Julia 字符串

Source: https://www.runoob.com/julia/julia-strings.html

## Julia 字符串

字符串（英语：string），是由零个或多个字符组成的有限序列。它是编程语言中表示文本的数据类型。

Julia 通常使用单引号 ' 创建单个字符，双引号 " 或三个引号 """ 创建字符串。例如：

```

c = 'x'

str = "RUNOOB"

runoob = """菜鸟教程 "RUNOOB"，包含了单个引号"""

```

Julia 字符串类型特性：

- Julia 中用于字符串（和字符串字面量）的内置具体类型是 String。
- Julia 的字符串类型都是抽象类型 AbstractString 的子类型。
- Julia 有优秀的表示单字符的类型，即 AbstractChar。Char 是 AbstractChar 的内置子类型，它能表示任何 Unicode 字符的 32 位原始类型（基于 UTF-8 编码）。
- Julia 字符串是不可更改的——任何 AbstractString 对象的值不可改变。

### 字符

单个字符用 Char 值表示。

Char 是 32 位原始类型，可以转换为其对应的整数值，即 Unicode 代码：

### 实例

julia> c = 'x'
'x': ASCII/Unicode U+0078 (category Ll: Letter, lowercase)

julia> typeof(c)
Char

julia> c = Int('x')
120

julia> typeof(c)
Int64

我们也可以将一个整数值转换为 Char：

### 实例

julia> Char(97)
'a': ASCII/Unicode U+0061 (category Ll: Letter, lowercase)

julia> Char(120)
'x': ASCII/Unicode U+0078 (category Ll: Letter, lowercase)

我们可以对 Char 的值进行比较和有限的算术运算：

### 实例

julia> 'A' < 'a'
true

julia> 'A' <= 'a' <= 'Z'
false

julia> 'A' <= 'X' <= 'Z'
true

julia> 'x' - 'a'
23

julia> 'A' + 1
'B': ASCII/Unicode U+0042 (category Lu: Letter, uppercase)

### 字符串

Julia 中的字符串可以使用双引号 " 或三个双引号 """ 来声明，如果您需要在字符串中的某个部分使用引号，就可以使用三个双引号来执行此操作，如下所示：

### 实例

julia> str = "RUNOOB"
"RUNOOB"

julia> runoob = """菜鸟教程 "RUNOOB"，包含了单个引号"""
"菜鸟教程 \"RUNOOB\"，包含了单个引号"

如果字符串太长了，我们可以使用反斜杠 \ 来拆分：

### 实例

julia> "This is a long \
line"
"This is a long line"
Julia 字符串也可以跟数组一样使用索引来读取指定位置的字符或者截取指定位置的字符串，索引的起时位置为 1 或者 begin，结束位置为 end ：

### 实例

julia> str = "RUNOOB"
"RUNOOB"

julia> str[begin]
'R': ASCII/Unicode U+0052 (category Lu: Letter, uppercase)

julia> str[1]
'R': ASCII/Unicode U+0052 (category Lu: Letter, uppercase)

julia> str[2]
'U': ASCII/Unicode U+0055 (category Lu: Letter, uppercase)

julia> str[end]
'B': ASCII/Unicode U+0042 (category Lu: Letter, uppercase)

julia> str[end-1]
'O': ASCII/Unicode U+004F (category Lu: Letter, uppercase)

我们可以用范围索引来提取子字符串：

### 实例

julia> str = "RUNOOB"
"RUNOOB"

julia> str[begin:end]
"RUNOOB"

julia> str[begin:end-1]
"RUNOO"

julia> str[2:5]
"UNOO"

另外，表达式 str[k] 和 str[k:k] 给出的结果是不同的，前者是通过索引获取指定位置的字符，类型是 Char 类型，后者是通过给定的范围读取字符串，只不过它只是包含了一个字符的字符串：

### 实例

julia> str[6]
'B': ASCII/Unicode U+0042 (category Lu: Letter, uppercase)

julia> str[6:6]
"B"

范围截取也可以使用 SubString 方法来实现， 例如：

### 实例

julia> str = "long string"
"long string"

julia> substr = SubString(str, 1, 4)
"long"

julia> typeof(substr)
SubString{String}

### 字符串拼接 我们可以使用 string() 方法将多个字符串拼接起来：

### 实例

julia> greet = "Hello"
"Hello"

julia> whom = "world"
"world"

julia> string(greet, ", ", whom, ".\n")
"Hello, world.\n"

### 插值

拼接构造字符串的方式有时有些麻烦。为了减少对于 string 的冗余调用或者重复地做乘法，Julia 允许像 Perl 中一样使用 $ 对字符串字面量进行插值：

### 实例

julia> "$greet, $whom.\n"
"Hello, world.\n"

这更易读更方便，而且等效于上面的字符串拼接——系统把这个显然一行的字符串字面量重写成带参数的字符串字面量拼接 string(greet, ", ", whom, ".\n")。

在 $ 之后最短的完整表达式被视为插入其值于字符串中的表达式。因此，你可以用括号向字符串中插入任何表达式：

### 实例

julia> "1 + 2 = $(1 + 2)"
"1 + 2 = 3"

拼接和插值都调用 string 以转换对象为字符串形式。 然而，string 实际上仅仅返回了 print 的输出，因此，新的类型应该添加 print 或 show 方法，而不是 string 方法。

多数非 AbstractString 对象被转换为和它们作为文本表达式输入的方式密切对应的字符串：

### 实例

julia> v = [1,2,3]
3-element Vector{Int64}:
1
2
3

julia> "v: $v"
"v: [1, 2, 3]"

string 是 AbstractString 和 AbstractChar 值的标识，所以它们作为自身被插入字符串，无需引用，无需转义：

### 实例

julia> c = 'x'
'x': ASCII/Unicode U+0078 (category Ll: Letter, lowercase)

julia> "hi, $c"
"hi, x"

若要在字符串字面量中包含文本 $，就用反斜杠转义：

### 实例

julia> print("I have \$100 in my account.\n")
I have $100 in my account.

### Unicode 和 UTF-8

Julia 完全支持 Unicode 字符和字符串。

在字符字面量中，Unicode 代码可以用 Unicode \u 和 \U 转义序列表示，也可以用所有标准 C 转义序列表示。这些同样可以用来写字符串字面量：

### 实例

julia> s = "\u2200 x \u2203 y"
"∀ x ∃ y"

这些 Unicode 字符是作为转义还是特殊字符显示，取决于你终端的语言环境设置以及它对 Unicode 的支持。字符串字面量用 UTF-8 编码。UTF-8 是一种可变长度的编码，也就是说并非所有字符都以相同的字节数（code units）编码。在 UTF-8 中，ASCII 字符（小于 0x80(128) 的那些）如它们在 ASCII 中一样使用单字节编码；而 0x80 及以上的字符使用最多 4 个字节编码。

### 三引号字符串字面量

三引号 """...""" 为我们创建更长更复杂的字符串提供了便利，三引号内部可以方便的使用换行、引号及缩进等，不需要特别处理。

### 实例

julia> str = """
Hello,
world.
"""
" Hello,\n world.\n"

### 字符串比较

我们可以使用比较操作符按照字典顺序比较字符串：

### 实例

julia> "abracadabra" < "xylophone"
true

julia> "abracadabra" == "xylophone"
false

julia> "Hello, world." != "Goodbye, world."
true

julia> "1 + 2 = 3" == "1 + 2 = $(1 + 2)"
true

你可以使用 findfirst 与 findlast 函数搜索特定字符的索引：

### 实例

julia> findfirst(isequal('o'), "xylophone")
4

julia> findlast(isequal('o'), "xylophone")
7

julia> findfirst(isequal('z'), "xylophone")

你可以带上第三个参数，用 findnext 与 findprev 函数来在给定偏移量处搜索字符：

### 实例

julia> findnext(isequal('o'), "xylophone", 1)
4

julia> findnext(isequal('o'), "xylophone", 5)
7

julia> findprev(isequal('o'), "xylophone", 5)
4

julia> findnext(isequal('o'), "xylophone", 8)

你可以用 occursin 函数检查在字符串中某子字符串可否找到。

### 实例

julia> occursin("world", "Hello, world.")
true

julia> occursin("o", "Xylophon")
true

julia> occursin("a", "Xylophon")
false

julia> occursin('o', "Xylophon")
true

最后那个例子表明 occursin 也可用于搜寻字符字面量。

另外还有两个方便的字符串函数 repeat 和 join：

### 实例

julia> repeat(".:Z:.", 10)
".:Z:..:Z:..:Z:..:Z:..:Z:..:Z:..:Z:..:Z:..:Z:..:Z:."

julia> join(["apples", "bananas", "pineapples"], ", ", " and ")
"apples, bananas and pineapples"

其它有用的函数还包括：

- firstindex(str) - 给出可用来索引到 str 的最小（字节）索引（对字符串来说这总是 1，对于别的容器来说却不一定如此）。
- lastindex(str) - 给出可用来索引到 str 的最大（字节）索引。
- length(str) - str 中的字符个数。
- length(str, i, j) - str 中从 i 到 j 的有效字符索引个数。
- ncodeunits(str) - 字符串中代码单元（码元）的数目。
- codeunit(str, i) - 给出在字符串 str 中索引为 i 的代码单元值。
- thisind(str, i) - 给定一个字符串的任意索引，查找索引点所在的首个索引。
- nextind(str, i, n=1) - 查找在索引 i 之后第 n 个字符的开头。
- prevind(str, i, n=1) - 查找在索引 i 之前第 n 个字符的开始。

### 原始字符串字面量

无插值和非转义的原始字符串可用 raw"..." 形式的非标准字符串字面量表示。原始字符串字面量生成普通的 String 对象，它无需插值和非转义地包含和输入完全一样的封闭式内容。这对于包含其他语言中使用 " 或 \" 作为特殊字符的代码或标记的字符串很有用。

例外的是，引号仍必须转义，例如 raw"\"" 等效于 "\""。为了能够表达所有字符串，反斜杠也必须转义，不过只是当它刚好出现在引号前面时。

```
julia> println(raw"\\ \\\"")
\\ \"
```

请注意，前两个反斜杠在输出中逐字显示，这是因为它们不是在引号前面。然而，接下来的一个反斜杠字符转义了后面的一个反斜杠；又由于这些反斜杠出现在引号前面，最后一个反斜杠转义了一个引号。

---

## Julia 正则表达式

Source: https://www.runoob.com/julia/julia-regexes.html

## Julia 正则表达式

正则表达式(regular expression)描述了一种字符串匹配的模式，可以用来检查一个串是否含有某种子串、将匹配的子串做替换或者从某个串中取出符合某个条件的子串等。

Julia 具有与 Perl 兼容的正则表达式 (regexes)。

Julia 的正则表达式的三种形式，分别是匹配，替换和转化:

- 匹配：m//（还可以简写为//，略去m）
- 替换：s///
- 转化：tr///

这三种形式一般都和 =~ 或 !~ 搭配使用， =~ 表示相匹配，!~ 表示不匹配。

Julia 中正则表达式的输入使用了前缀各类以 r 开头：

### 实例

julia> re = r"^\s*(?:#|$)"
r"^\s*(?:#|$)"

julia> typeof(re)
Regex

若要检查正则表达式是否匹配某字符串，就用 occursin：

### 实例

julia> occursin(r"^\s*(?:#|$)", "not a comment")
false

julia> occursin(r"^\s*(?:#|$)", "# a comment")
true

可以看到，occursin 只返回正确或错误，表明给定正则表达式是否在该字符串中出现。然而，通常我们不只想知道字符串是否匹配，更想了解它是如何匹配的。要捕获匹配的信息，可以改用 match 函数：

### 实例

julia> match(r"^\s*(?:#|$)", "not a comment")

julia> match(r"^\s*(?:#|$)", "# a comment")
RegexMatch("#")

若正则表达式与给定字符串不匹配，match 返回 nothing——在交互式提示框中不打印任何东西的特殊值。除了不打印，它是一个完全正常的值，这可以用程序来测试：

### 实例

m = match(r"^\s*(?:#|$)", line)
if m === nothing
println("not a comment")
else
println("blank or comment")
end

如果正则表达式匹配，match 的返回值是一个 RegexMatch 对象。这些对象记录了表达式是如何匹配的，包括该模式匹配的子字符串和任何可能被捕获的子字符串。上面的例子仅仅捕获了匹配的部分子字符串，但也许我们想要捕获的是注释字符后面的任何非空文本。我们可以这样做：

### 实例

julia> m = match(r"^\s*(?:#\s*(.*?)\s*$|$)", "# a comment ")
RegexMatch("# a comment ", 1="a comment")

当调用 match 时，你可以选择指定开始搜索的索引。例如：

### 实例

julia> m = match(r"[0-9]","aaaa1aaaa2aaaa3",1)
RegexMatch("1")

julia> m = match(r"[0-9]","aaaa1aaaa2aaaa3",6)
RegexMatch("2")

julia> m = match(r"[0-9]","aaaa1aaaa2aaaa3",11)
RegexMatch("3")

你可以从 RegexMatch 对象中提取如下信息：

- 匹配的整个子字符串：m.match
- 作为字符串数组捕获的子字符串：m.captures
- 整个匹配开始处的偏移：m.offset
- 作为向量的捕获子字符串的偏移：m.offsets

当捕获不匹配时，m.captures 在该处不再包含一个子字符串，而是 什么也不 包含；此外，m.offsets 的偏移量为 0（回想一下，Julia 的索引是从 1 开始的，因此字符串的零偏移是无效的）。下面是两个有些牵强的例子：

### 实例

julia> m = match(r"(a|b)(c)?(d)", "acd")
RegexMatch("acd", 1="a", 2="c", 3="d")

julia> m.match
"acd"

julia> m.captures
3-element Vector{Union{Nothing, SubString{String}}}:
"a"
"c"
"d"

julia> m.offset
1

julia> m.offsets
3-element Vector{Int64}:
1
2
3

julia> m = match(r"(a|b)(c)?(d)", "ad")
RegexMatch("ad", 1="a", 2=nothing, 3="d")

julia> m.match
"ad"

julia> m.captures
3-element Vector{Union{Nothing, SubString{String}}}:
"a"
nothing
"d"

julia> m.offset
1

julia> m.offsets
3-element Vector{Int64}:
1
0
2

让捕获作为数组返回是很方便的，这样就可以用解构语法把它们和局域变量绑定起来。为了方便，RegexMatch 对象实现了传递到 captures 字段的迭代器方法，因此您可以直接解构匹配对象：

### 实例

julia> first, second, third = m; first
"a"

通过使用捕获组的编号或名称对 RegexMatch 对象进行索引，也可实现对捕获的访问：

### 实例

julia> m=match(r"(?<hour>\d+):(?<minute>\d+)","12:45")
RegexMatch("12:45", hour="12", minute="45")

julia> m[:minute]
"45"

julia> m[2]
"45"

使用 replace 时利用 \n 引用第 n 个捕获组和给替换字符串加上 s 的前缀，可以实现替换字符串中对捕获的引用。捕获组 0 指的是整个匹配对象。可在替换中用 \g 对命名捕获组进行引用。例如：

```
julia> replace("first second", r"(\w+) (?<agroup>\w+)" => s"\g<agroup> \1")
"second first"
```

为明确起见，编号捕获组也可用 \g 进行引用，例如：

```

julia> replace("a", r"." => s"\g<0>1")
"a1"
```

你可以在后双引号的后面加上 i, m, s 和 x 等标志对正则表达式进行修改。

更多正则表达式内容可以参考：正则表达式 - 教程

---

## Julia 函数

Source: https://www.runoob.com/julia/julia-functions.html

## Julia 函数

函数是一组一起执行一个任务的语句。

在 Julia 里，函数是将参数值组成的元组映射到返回值的一个对象。

Julia 中使用 function 定义函数，基本语法是：

```
function functionname(args)
expression
expression
expression
...
expression
end
```

默认情况下，函数返回的值是最后计算的表达式的值，所以我们看到上面是没有 return 语句的，当然，如果使用 return 关键字，函数就会立即返回：。

### 实例

julia> function f(x,y)
x + y
end
f (generic function with 1 method)
julia> f(2,3)
5

julia> function bills(money)
if money < 0
return false
else
return true
end
end
bills (generic function with 1 method)

julia> bills(50)
true

julia> bills(-50)
false

如果函数要返回多个值，可以使用元组：

### 实例

julia> function mul(x,y)
x+y, x*y
end
mul (generic function with 1 method)

julia> mul(5, 10)
(15, 50)

当函数中只有一个表达式时，您可以省略 function 关键字，等号左侧设置函数名与参数，右侧设置表达式，类似赋值形式：

```
julia> f(x,y) = x + y
f (generic function with 1 method)
```

没有括号时，表达式 f 指的是函数对象，可以像任何值一样被传递：

```
julia> g = f;
julia> g(2,3)
5
```

### 实例

julia> f(a) = a * a
f (generic function with 1 method)

julia> f(5)
25

julia> func(x, y) = sqrt(x^2 + y^2)
func (generic function with 1 method)

julia> func(5, 4)
6.4031242374328485

和变量名一样，Unicode 字符也可以用作函数名：

```
julia> ∑(x,y) = x + y
∑ (generic function with 1 method)

julia> ∑(2, 3)
5
```

#### 返回类型

我们可以使用 :: 运算符来指定函数的返回类型。

### 实例

julia> function g(x, y)::Int8
return x * y
end;

julia> typeof(g(1, 2))
Int8

以上函数实例将忽略 x 和 y 的类型，返回 Int8 类型的值。

#### 可选参数

在函数中我们可以设置参数默认值，这样在没有提供该参数的时候，就可以使用默认值来计算：

以下实例定义函数 pos，设置三个参数，其中参数 cz 设置默认值为 0，在函数调用时，可以不提供该参数：

### 实例

julia> function pos(ax, by, cz=0)
println("$ax, $by, $cz")
end
pos (generic function with 2 methods)

julia> pos(10, 30)
10, 30, 0

julia> pos(10, 30, 50)
10, 30, 50

#### 关键字参数

有时候我们定义的一些函数需要大量参数，但调用这些函数可能很麻烦，因为我们可能会忘记提供参数的顺序。 例如：

function foo(a, b, c, d, e, f)
...
end

我们可能会忘记参数的顺序，发生以下调用函数的情况：

```
foo("25", -5.6987, "hello", 56, good, 'ABC')
或
foo("hello", 56, "25", -5.6987, 'ABC', good)
```

这样看起来就非常混乱。

Julia 关键字参数允许通过名称而不是仅通过位置来识别参数，使得这些复杂函数易于使用和扩展。

使用关键字来标记参数，需要在函数的未标记参数之后使用分号 ;，并在其后跟一个或多个关键值对 key=value，如下所示:

### 实例

julia> function foo(a, b ; c = 10, d = "hi")
println("a is $a")
println("b is $b")
return "c => $c, d => $d"
end
foo (generic function with 1 method)

julia> foo(100,20)
a is 100
b is 20
"c => 10, d => hi"

julia> foo("Hello", "Runoob", c=pi, d=22//7)
a is Hello
b is Runoob
"c => π, d => 22//7"

使用关键字参数，参数的位置也不太重要了，以上函数我们也可以这样调用：

### 实例

julia> foo(c=pi, d =22/7, "Hello", "Runoob")
a is Hello
b is Runoob
"c => π, d => 3.142857142857143"

### 匿名函数

匿名函数是一个没有函数名的函数。

匿名函数在程序运行时动态声明，除了没有函数名外，其他的与标准函数一样。

在 Julia 中，匿名函数可用于许多地方，例如 map() 和 列表推导。

使用匿名函数后，我们的代码变得更简洁了。

匿名函数的语法使用符号 ->。

### 实例

julia> x -> x^2 + 2x - 1
#1 (generic function with 1 method)

julia> function (x)
x^2 + 2x - 1
end
#3 (generic function with 1 method)

以上实例创建了一个接受一个参数 x 并返回当前值的多项式 x^2+2x-1 的函数。

匿名函数最主要的用法是传递给接收函数作为参数的函数。一个经典的例子是 map ，为数组的每个元素应用一次函数，然后返回一个包含结果值的新数组：

### 实例

julia> map(round, [1.2, 3.5, 1.7])
3-element Vector{Float64}:
1.0
4.0
2.0

如果做为第一个参数传递给 map 的转换函数已经存在，那直接使用函数名称是没问题的。但是通常要使用的函数还没有定义好，这样使用匿名函数就更加方便：

### 实例

julia> map(x -> x^2 + 2x - 1, [1, 3, -1])
3-element Vector{Int64}:
2
14
-2

接受多个参数的匿名函数写法可以使用语法 (x,y,z)->2x+y-z，而无参匿名函数写作 ()->3 。无参函数的这种写法看起来可能有些奇怪，不过它对于延迟计算很有必要。这种用法会把代码块包进一个无参函数中，后续把它当做 f 调用。

例如，考虑对 get 的调用：

### 实例

get(dict, key) do
# default value calculated here
time()
end

上面的代码等效于使用包含代码的匿名函数调用 get。 被包围在 do 和 end 之间，如下所示:

```
get(()->time(), dict, key)
```

这里对 time 的调用，被包裹了它的一个无参数的匿名函数延迟了。该匿名函数只当 dict 缺少被请求的键时，才被调用。

### 函数嵌套与递归

在 Julia 中，函数可以嵌套使用。

以下实例 add() 函数内嵌套一个 add1 实例：

### 实例

julia> function add(x)
Y = x * 2
function add1(Y)
Y += 1
end
add1(Y)
end
add (generic function with 1 method)

julia> d = 10
10

julia> add(d)
21

同样，Julia 中的函数也可以是递归的。

递归指的是在函数的定义中使用函数自身的方法。

举个例子：
从前有座山，山里有座庙，庙里有个老和尚，正在给小和尚讲故事呢！故事是什么呢？"从前有座山，山里有座庙，庙里有个老和尚，正在给小和尚讲故事呢！故事是什么呢？'从前有座山，山里有座庙，庙里有个老和尚，正在给小和尚讲故事呢！故事是什么呢？……'"

以下我们使用三元运算符来测试递归，三元运算符操作三个操作对象 expr ? a : b，如果 expr 为 true ，值为 a 的计算结果 ，否则为 b 的计算结果。

### 实例

julia> sum(x) = x > 1 ? sum(x-1) + x : x
sum (generic function with 1 method)

julia> sum(10)
55

以上实例用于计算某个整数之前所有数的总和，直到并包括某个数字。 在这个递归中，因为有一个基本情况，即当 x 为 1 时，这个值被返回。

递归最著名的例子是计算第 n 个斐波那契数，斐波那契数列指的是这样一个数列 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233，377，610，987，1597，2584，4181，6765，10946，17711，28657，46368........

这个数列从第 3 项开始，每一项都等于前两项之和。

### 实例

julia> fib(x) = x < 2 ? x : fib(x-1) + fib(x-2)
fib (generic function with 1 method)

julia> fib(10)
55

julia> fib(20)
6765

### Map

Map 定义格式如下：

```
map(func, coll)
```

这里，func 是一个函数，它依次应用于集合 coll 的每个元素。 Map 一般包含匿名函数并返回一个新的集合。

### 实例

julia> map(A -> A^3 + 3A - 3, [10,3,-2])
3-element Array{Int64,1}:
1027
33
-17

### Filter

Filter 定义格式如下：

```

filter(function, collection)
```

filter 函数返回集合的副本，并删除通过调用该函数结果为 false 的元素。

### 实例

julia> array = Int[1,2,3]
3-element Array{Int64,1}:
1
2
3

julia> filter(x -> x % 2 == 0, array)
1-element Array{Int64,1}:
2

---

## Julia 流程控制

Source: https://www.runoob.com/julia/julia-flow-control.html

## Julia 流程控制

流程控制语句通过程序设定一个或多个条件语句来实现。在条件为 true 时执行指定程序代码，在条件为 false 时执行其他指定代码。

Julia 提供了大量的流程控制语句：

- 复合表达式：begin 和 ;。
- 条件表达式：if-elseif-else 和 ?: (三元运算符)。
- 短路运算：逻辑运算符 &&（与）和 ||（或），以及链式比较。
- 循环语句：循环：while 和 for。
- 异常处理：try-catch、error 和 throw。
- Task（协程）：yieldto。

### 复合表达式

begin ... end 表达式可以按顺序计算若干子表达式，并返回最后一个子表达式的值：

### 实例

julia> z = begin
x = 1
y = 2
x + y
end
3

因为这些是非常简短的表达式，它们可以简单地被放到一行里，这也是 ; 链的由来：

### 实例

julia> z = (x = 1; y = 2; x + y)
3

实际使用过程并不要求 begin 代码块是多行的，或者 ; 链是单行的：

### 实例

julia> begin x = 1; y = 2; x + y end
3

julia> (x = 1;
y = 2;
x + y)
3

### 条件表达式

条件表达式可以根据布尔表达式的值来决定执行哪一个代码块。

if-elseif-else 语法：

if boolean_expression 1
/* 当布尔表达式 1 为真时执行 */
elseif boolean_expression 2
/* 当布尔表达式 2 为真时执行 */
elseif boolean_expression 3
/* 当布尔表达式 3 为真时执行 */
else
/* 当上面条件都不为真时执行 */

一个 if 语句后可跟一个可选的 else if...else 语句，这可用于测试多种条件。

当使用 if...else if...else 语句时，以下几点需要注意：

- 一个 if 后可跟零个或一个 else，else 必须在所有 else if 之后。
- 一个 if 后可跟零个或多个 else if，else if 必须在 else 之前。
- 一旦某个 else if 匹配成功，其他的 else if 或 else 将不会被测试。

下面是对 if-elseif-else 条件语法的分析：

### 实例

if x < y
println("x is less than y")
elseif x > y
println("x is greater than y")
else
println("x is equal to y")
end

如果表达式 x < y 是 true，那么对应的代码块会被执行；否则判断条件表达式 x > y，如果它是 true，则执行对应的代码块；如果没有表达式是 true，则执行 else 代码块。

### 实例

julia> function test(x, y)
if x < y
println("x is less than y")
elseif x > y
println("x is greater than y")
else
println("x is equal to y")
end
end
test (generic function with 1 method)

julia> test(1, 2)
x is less than y

julia> test(2, 1)
x is greater than y

julia> test(1, 1)
x is equal to y

#### 三元运算符

三元运算符 ?: 类似 if-elseif-else 语法: a ? b : c

在 ? 之前的表达式 a, 是一个条件表达式，如果条件 a 是 true，三元运算符计算在 : 之前的表达式 b；如果条件 a 是 false，则执行 : 后面的表达式 c。

注意：? 和 : 旁边的空格是强制的，像 a?b:c 这种表达式不是一个有效的三元表达式（但在? 和 : 之后的换行是允许的）。

### 实例

julia> x = 1; y = 2;

julia> println(x < y ? "less than" : "not less than")
less than

julia> x = 1; y = 0;

julia> println(x < y ? "less than" : "not less than")
not less than

如果表达式 x < y 为真，整个三元运算符会执行字符串 "less than"，否则执行字符串 "not less than"。

链式嵌套使用三元运算符：

### 实例

julia> test(x, y) = println(x < y ? "x is less than y" :
x > y ? "x is greater than y" : "x is equal to y")
test (generic function with 1 method)

julia> test(1, 2)
x is less than y

julia> test(2, 1)
x is greater than y

julia> test(1, 1)
x is equal to y

为了方便链式传值，运算符从右到左连接到一起。

与 if-elseif-else 类似，: 之前和之后的表达式只有在条件表达式为 true 或者 false 时才会被相应地执行：

### 实例

julia> v(x) = (println(x); x)
v (generic function with 1 method)

julia> 1 < 2 ? v("yes") : v("no")
yes
"yes"

julia> 1 > 2 ? v("yes") : v("no")
no
"no"

### 短路运算

Julia 中的 && 和 || 运算符分别对应于逻辑 "与" 和 "或" 操作。

- 在表达式 `a && b` 中，子表达式 `b` 仅当 `a` 为 `true` 的时候才会被执行，如果 a 为 false 则直接返回 false。
- 在表达式 `a || b` 中，子表达式 `b` 仅在 `a` 为 `false` 的时候才会被执行，如果 a 为 true ，直接返回 a。

### && 实例

julia> isodd(3) && @warn("An odd Number!")
┌ Warning: An odd Number!
└ @ Main REPL[5]:1

julia> isodd(4) && @warn("An odd Number!")
false

### || 实例

julia> isodd(3) || @warn("An odd Number!")
true

julia> isodd(4) || @warn("An odd Number!")
┌ Warning: An odd Number!
└ @ Main REPL[8]:1

&& 和 || 都依赖于右边，但是 && 比 || 有更高的优先级，查看实例：

### 实例

julia> t(x) = (println(x); true)
t (generic function with 1 method)

julia> f(x) = (println(x); false)
f (generic function with 1 method)

julia> t(1) && t(2)
1
2
true

julia> t(1) && f(2)
1
2
false

julia> f(1) && t(2)
1
false

julia> f(1) && f(2)
1
false

julia> t(1) || t(2)
1
true

julia> t(1) || f(2)
1
true

julia> f(1) || t(2)
1
2
true

julia> f(1) || f(2)
1
2
false

### 循环语句

循环语句使用 while 和 for 两个关键字来实现。

下面是一个 while 循环的例子：

### 实例

julia> i = 1;

julia> while i <= 5
println(i)
global i += 1
end
1
2
3
4
5

while 循环语句通过会执行条件表达式（ i <= 5），只要它为 true，就一直执行 while 循环的主体部分。如果条件表达式为 false，也就是 i=6 的时候，那么循环就结束了。

for 循环使用起来会更加方便，以上实例使用 for 循环实现如下：

### 实例

julia> for i = 1:5
println(i)
end
1
2
3
4
5

这里的 1:5 是一个范围对象，代表数字 1, 2, 3, 4, 5 的序列。

for 循环在这些值之中迭代，对每一个变量 i 进行赋值。

for 循环与之前 while 循环的一个非常重要区别是作用域，即变量的可见性。如果变量 i 没有在另一个作用域里引入，在 for 循环内，它就只在 for 循环内部可见，在外部和后面均不可见。你需要一个新的交互式会话实例或者一个新的变量名来测试这个特性：

### 实例

julia> for j = 1:5
println(j)
end
1
2
3
4
5

julia> j
ERROR: UndefVarError: j not defined

一般来说，for 循环组件可以用于迭代任一个容器。在这种情况下，相比 =，另外的（但完全相同）关键字 in 或者 ∈ 则更常用，因为它使得代码更清晰：

### 实例

julia> for i in [1,4,0]
println(i)
end
1
4
0

julia> for s ∈ ["foo","bar","baz"]
println(s)
end
foo
bar
baz

为了方便，我们可能会在测试条件不成立之前终止一个 while 循环，或者在访问到迭代对象的结尾之前停止一个 for 循环，这可以用关键字 break 来完成：

### 实例

julia> i = 1;

julia> while true
println(i)
if i >= 5
break
end
global i += 1
end
1
2
3
4
5

julia> for j = 1:1000
println(j)
if j >= 5
break
end
end
1
2
3
4
5

没有关键字 break 的话，上面的 while 循环永远不会自己结束，而 for 循环会迭代到 1000，这些循环都可以使用 break 来提前结束。

在某些场景下，需要直接结束此次迭代，并立刻进入下次迭代，continue 关键字可以用来完成此功能：

### 实例

julia> for i = 1:10
if i % 3 != 0
continue
end
println(i)
end
3
6
9

这是一个有点做作的例子，因为我们可以通过否定这个条件，把 println 调用放到 if 代码块里来更简洁的实现同样的功能。在实际应用中，在 continue 后面还会有更多的代码要运行，并且调用 continue 的地方可能会有多个。

多个嵌套的 for 循环可以合并到一个外部循环，可以用来创建其迭代对象的笛卡尔积：

### 实例

julia> for i = 1:2, j = 3:4
println((i, j))
end
(1, 3)
(1, 4)
(2, 3)
(2, 4)

有了这个语法，迭代变量依然可以正常使用循环变量来进行索引，例如 for i = 1:n, j = 1:i 是合法的，但是在一个循环里面使用 break 语句则会跳出整个嵌套循环，不仅仅是内层循环。每次内层循环运行的时候，变量（i 和 j）会被赋值为他们当前的迭代变量值。所以对 i 的赋值对于接下来的迭代是不可见的：

### 实例

julia> for i = 1:2, j = 3:4
println((i, j))
i = 0
end
(1, 3)
(1, 4)
(2, 3)
(2, 4)

如果这个例子给每个变量一个关键字 for 来重写，那么输出会不一样：第二个和第四个变量包含 0。

可以使用 zip 在单个 for 循环中同时迭代多个容器：

### 实例

julia> for (j, k) in zip([1 2 3], [4 5 6 7])
println((j,k))
end
(1, 4)
(2, 5)
(3, 6)

julia> for x in zip(0:15, 100:110, 200:210)
println(x)
end
(0, 100, 200)
(1, 101, 201)
(2, 102, 202)
(3, 103, 203)
(4, 104, 204)
(5, 105, 205)
(6, 106, 206)
(7, 107, 207)
(8, 108, 208)
(9, 109, 209)
(10, 110, 210)

// 处理元素数量不一致
julia> for x in zip(0:10, 100:115, 200:210)
println(x)
end
(0, 100, 200)
(1, 101, 201)
(2, 102, 202)
(3, 103, 203)
(4, 104, 204)
(5, 105, 205)
(6, 106, 206)
(7, 107, 207)
(8, 108, 208)
(9, 109, 209)
(10, 110, 210)

使用 zip 将创建一个迭代器，它是一个包含传递给它的容器的子迭代器的元组。 zip 迭代器将按顺序迭代所有子迭代器，在 for 循环的第 ii 次迭代中选择每个子迭代器的第 ii 个元素。 一旦任何子迭代器用完，for 循环就会停止。

for 语句可以嵌套多个循环条件，使用逗号 , 分隔：

### 实例

julia> for n in 1:5, m in 1:5
@show (n, m)
end
(n, m) = (1, 1)
(n, m) = (1, 2)
(n, m) = (1, 3)
(n, m) = (1, 4)
(n, m) = (1, 5)
(n, m) = (2, 1)
(n, m) = (2, 2)
(n, m) = (2, 3)
(n, m) = (2, 4)
(n, m) = (2, 5)
(n, m) = (3, 1)
(n, m) = (3, 2)
(n, m) = (3, 3)
(n, m) = (3, 4)
(n, m) = (3, 5)
(n, m) = (4, 1)
(n, m) = (4, 2)
(n, m) = (4, 3)
(n, m) = (4, 4)
(n, m) = (4, 5)
(n, m) = (5, 1)
(n, m) = (5, 2)
(n, m) = (5, 3)
(n, m) = (5, 4)
(n, m) = (5, 5)

#### 推导式 推导式是一种独特的数据处理方式，可以从一个数据序列构建另一个新的数据序列的结构体。

格式如下：

[表达式 for 变量 in 列表]
[out_exp_res for out_exp in input_list]

或者

[表达式 for 变量 in 列表 if 条件]
[out_exp_res for out_exp in input_list if condition]

### 实例

julia> [X^2 for X in 1:5]
5-element Array{Int64,1}:
1
4
9
16
25

我们还可以指定想要生成的元素类型：

### 实例

julia> Complex[X^2 for X in 1:5]
5-element Array{Complex,1}:
1 + 0im
4 + 0im
9 + 0im
16 + 0im
25 + 0im

#### 遍历数组

有时我们希望遍历数组的每个元素，包含元素的索引号。

Julia 提供了 enumerate(iter) 函数，参数 iter 为可迭代对象，该函数将生成索引号以及每个索引号对应的值。

### 实例

julia> a = ["a", "b", "c"];

julia> for (index, value) in enumerate(a)
println("$index $value")
end
1 a
2 b
3 c

julia> arr = rand(0:9, 4, 4)
4×4 Array{Int64,2}:
7 6 5 8
8 6 9 4
6 3 0 7
2 3 2 4

julia> [x for x in enumerate(arr)]
4×4 Array{Tuple{Int64,Int64},2}:
(1, 7) (5, 6) (9, 5) (13, 8)
(2, 8) (6, 6) (10, 9) (14, 4)
(3, 6) (7, 3) (11, 0) (15, 7)
(4, 2) (8, 3) (12, 2) (16, 4)

### 异常处理

程序执行过程中，如果发生意外条件，一个函数可能无法向调用者返回一个合理的值。在这种情况下，最好让意外条件终止程序并打印出调试的错误信息，这样可以方便开发者处理它们的代码。

通过 try / catch 语句，可以方便处理异常情况。

例如， 在下面的代码中，平方根函数 sqrt 会引发异常。 通过 try / catch 我们可以准确输出异常信息。

### 实例

julia> try
sqrt("ten")
catch e
println("你需要输入一个数字")
end
你需要输入一个数字

#### finally 子句

在进行状态改变或者使用类似文件的资源的编程时，经常需要在代码结束的时候进行必要的清理工作（比如关闭文件）。由于异常会使得部分代码块在正常结束之前退出，所以可能会让上述工作变得复杂。finally 关键字提供了一种方式，无论代码块是如何退出的，都能够让代码块在退出时运行某段代码。

这里是一个确保一个打开的文件被关闭的例子：

### 实例

f = open("file")
try
# operate on file f
finally
close(f)
end

当控制流离开 try 代码块（例如，遇到 return，或者正常结束），close(f) 就会被执行。如果 try 代码块由于异常退出，这个异常会继续传递。catch 代码块可以和 try 还有 finally 配合使用。这时 finally 代码块会在 catch 处理错误之后才运行。

#### throw 函数

我们可以用 throw 显式地创建异常。

例如，若一个函数只对非负数有定义，当输入参数是负数的时候，可以用 throw 抛出一个 DomainError。

### 实例

julia> f(x) = x>=0 ? exp(-x) : throw(DomainError(x, "argument must be nonnegative"))
f (generic function with 1 method)

julia> f(1)
0.36787944117144233

julia> f(-1)
ERROR: DomainError with -1:
argument must be nonnegative
Stacktrace:
[1] f(::Int64) at ./none:1

注意 DomainError 后面不接括号的话不是一个异常，而是一个异常类型。我们需要调用它来获得一个 Exception 对象：

### 实例

julia> typeof(DomainError(nothing)) <: Exception
true

julia> typeof(DomainError) <: Exception
false

另外，一些异常类型会接受一个或多个参数来进行错误报告：

### 实例

julia> throw(UndefVarError(:x))
ERROR: UndefVarError: x not defined

我们可以仿照 UndefVarError 的写法，用自定义异常类型来轻松实现这个机制：

### 实例

julia> struct MyUndefVarError <: Exception
var::Symbol
end

julia> Base.showerror(io::IO, e::MyUndefVarError) = print(io, e.var, " not defined")

### Task（协程）

Task（协程）提供了非局部的流程控制，这使得在暂时挂起的计算任务之间进行切换成为可能。在后面的异步编程章节我们会详细介绍。

---

## Julia 字典和集合

Source: https://www.runoob.com/julia/julia-dictionaries-sets.html

## Julia 字典和集合

前面几个章节我们学到了 Julia 数组和 julia 元组。

数组是一种集合，此外 Julia 也有其他类型的集合，比如字典和 set（无序集合列表）。

### 字典

字典是一种可变容器模型，且可存储任意类型对象。

字典的每个键值 key=>value 对用 => 分割，每个键值对之间用逗号 , 分割，整个字典包括在花括号 {} 中 ,格式如下所示：

#### 创建字典 创建字典的语法格式如下：

```
Dict("key1" => value1, "key2" => value2,,…, "keyn" => valuen)
```

以下实例创建一个简单的字典，键 A 对应的值为 1，键 B 对应的值为 2：

```
Dict("A"=>1, "B"=>2)
```

### 实例

julia> D=Dict("A"=>1, "B"=>2)
Dict{String, Int64} with 2 entries:
"B" => 2
"A" => 1

julia>

使用 for 来创建一个字典：

### 实例

julia> first_dict = Dict(string(x) => sind(x) for x = 0:5:360)
Dict{String, Float64} with 73 entries:
"285" => -0.965926
"310" => -0.766044
"245" => -0.906308
"320" => -0.642788
"350" => -0.173648
"20" => 0.34202
"65" => 0.906308
"325" => -0.573576
"155" => 0.422618
"80" => 0.984808
"335" => -0.422618
"125" => 0.819152
"360" => 0.0
"75" => 0.965926
"110" => 0.939693
"185" => -0.0871557
"70" => 0.939693
"50" => 0.766044
"190" => -0.173648
⋮ => ⋮

#### 键（Key）

字典中的键是唯一的， 如果我们为一个已经存在的键分配一个值，我们不会创建一个新的，而是修改现有的键。

查找 key

我们可以使用 haskey() 函数来检查字典是否包含指定的 key：

### 实例

julia> D=Dict("A"=>1, "B"=>2)
Dict{String, Int64} with 2 entries:
"B" => 2
"A" => 1

julia> haskey(first_dict, "A")
false

julia> haskey(D, "A")
true

julia> haskey(D, "Z")
false

也可以使用 in() 函数来检查字典是否包含键/值对：

### 实例

julia> D=Dict("A"=>1, "B"=>2)
Dict{String, Int64} with 2 entries:
"B" => 2
"A" => 1

julia> in(("A" => 1), D)
true

julia> in(("X" => 220), first_dict)
false

添加 key/value 对

我们可以在已存在的字典中添加一个新的 key/value 对，如下所示：

### 实例

julia> D=Dict("A"=>1, "B"=>2)
Dict{String, Int64} with 2 entries:
"B" => 2
"A" => 1

julia> D["C"] = 3
3

julia> D
Dict{String, Int64} with 3 entries:
"B" => 2
"A" => 1
"C" => 3

删除 key/value 对

我们可以使用 delete!() 函数删除已存在字典的 key：

### 实例

julia> delete!(D, "C")
Dict{String, Int64} with 2 entries:
"B" => 2
"A" => 1

获取字典中所有的 key

我们可以使用 keys() 函数获取字典中所有的 key：

### 实例

julia> keys(D)
KeySet for a Dict{String, Int64} with 2 entries. Keys:
"B"
"A"

julia>

#### 值（Value）

字典中的每个键都有一个对应的值。

查看字典所有值

我们可以使用 values() 查看字典所有值：

### 实例

julia> D=Dict("A"=>1, "B"=>2)
Dict{String, Int64} with 2 entries:
"B" => 2
"A" => 1

julia> values(D)
ValueIterator for a Dict{String, Int64} with 2 entries. Values:
2
1

julia>

#### 字典作为可迭代对象

我们可以将字典作为可迭代对象来查看键/值对：

### 实例

julia> D=Dict("A"=>1, "B"=>2)
Dict{String, Int64} with 2 entries:
"B" => 2
"A" => 1

julia> for kv in D
println(kv)
end
"B" => 2
"A" => 1

实例中 kv 是一个包含每个键/值对的元组。

#### 字典排序

字典是无序的，但我们可以使用 sort() 函数来对字典进行排序：

### 实例

julia> runoob_dict = Dict("R" => 100, "S" => 220, "T" => 350, "U" => 400, "V" => 575, "W" => 670)
Dict{String, Int64} with 6 entries:
"S" => 220
"U" => 400
"T" => 350
"W" => 670
"V" => 575
"R" => 100

julia> for key in sort(collect(keys(runoob_dict)))
println("$key => $(runoob_dict[key])")
end
R => 100
S => 220
T => 350
U => 400
V => 575
W => 670

我们可以使用 DataStructures.ji 包中的 SortedDict 数据类型让字典始终保持排序状态。

使用 DataStructures 包需要先安装它，可以在 REPL 的 Pkg 模式中，使用 add 命令添加 SortedDict。

在 REPL 中输入符号 ] ，进入 pkg 模式。

### 进入 pkg 模式

julia> ] # 输入 ] 就进入 pkg 模式

添加包预防语法格式：

```
add 包名
```

以下我们添加 DataStructures 包后，后面的实例就可以正常运行了：

```
(@v1.7) pkg> add DataStructures
```
未注册的包，可以直接指定 url：

```
add https://github.com/fredrikekre/ImportMacros.jl
```

本地包：

```
add 本地路径/包名.jl
```

### 实例

julia> import DataStructures

julia> runoob_dict = DataStructures.SortedDict("S" => 220, "T" => 350, "U" => 400, "V" => 575, "W" => 670)
DataStructures.SortedDict{String, Int64, Base.Order.ForwardOrdering} with 5 entries:
"S" => 220
"T" => 350
"U" => 400
"V" => 575
"W" => 670

julia> runoob_dict["R"] = 100
100

julia> runoob_dict
DataStructures.SortedDict{String, Int64, Base.Order.ForwardOrdering} with 6 entries:
"R" => 100
"S" => 220
"T" => 350
"U" => 400
"V" => 575
"W" => 670

### Set(集合) Julia Set(集合)是没有重复的对象数据集，所有的元素都是唯一的。

以下是 set 和其他类型的集合之间的区别:

- set 中的元素是唯一的
- set 中元素的顺序不重要

set 用于创建不重复列表。

#### 创建 Set 集合

借助 Set 构造函数，我们可以创建如下集合：

### 实例

julia> var_site = Set()
Set{Any}()

julia> num_primes = Set{Int64}()
Set{Int64}()

julia> var_site = Set{String}(["Google","Runoob","Taobao"])
Set{String} with 3 elements:
"Google"
"Taobao"
"Runoob"
Alternatively we can also use push!() function, as arrays, to add elements in sets as follows −

我们可以使用 push!() 函数添加集合元素，如下所示：

### 实例

julia> push!(var_site, "Wiki")
Set{String} with 4 elements:
"Google"
"Wiki"
"Taobao"
"Runoob"

我们可以使用 in() 函数查看元素是否存在于集合中：

### 实例

julia> in("Runoob", var_site)
true

julia> in("Zhihu", var_site)
false

#### 常用操作

并集、交集和差集是我们可以对集合常用的一些操作， 这些操作对应的函数是 union()、intersect() 和 setdiff()。

并集 两个集合 A，B，把他们所有的元素合并在一起组成的集合，叫做集合 A 与集合 B 的并集。

### 实例

julia> A = Set{String}(["red","green","blue", "black"])
Set{String} with 4 elements:
"blue"
"green"
"black"
"red"

julia> B = Set(["red","orange","yellow","green","blue","indigo","violet"])
Set{String} with 7 elements:
"indigo"
"yellow"
"orange"
"blue"
"violet"
"green"
"red"

julia> union(A, B)
Set{String} with 8 elements:
"indigo"
"green"
"black"
"yellow"
"orange"
"blue"
"violet"
"red"

交集

集合 A 和 B 的交集是含有所有既属 A 又属于 B 的元素，而没有其他元素的集合。

### 实例

julia> intersect(A, B)
Set{String} with 3 elements:
"blue"
"green"
"red"

差集

集合 A 和 B 的差集是含有所有属 A 但不属于 B 的元素，即去除 B 与 A 重叠的元素。

### 实例

julia> setdiff(A, B)
Set{String} with 1 element:
"black"

### 字典与集合常用函数实例

在下面的实例中，演示了字典中常用的函数，在集合中也同样适用：

创建两个字典 dict1 和 dict2：

### 实例

julia> dict1 = Dict(100=>"X", 220 => "Y")
Dict{Int64,String} with 2 entries:
100 => "X"
220 => "Y"

julia> dict2 = Dict(220 => "Y", 300 => "Z", 450 => "W")
Dict{Int64,String} with 3 entries:
450 => "W"
220 => "Y"
300 => "Z"

字典并集：

### 实例

julia> union(dict1, dict2)
4-element Array{Pair{Int64,String},1}:
100 => "X"
220 => "Y"
450 => "W"
300 => "Z"
Intersect
julia> intersect(dict1, dict2)
1-element Array{Pair{Int64,String},1}:
220 => "Y"

字典差集：

### 实例

julia> setdiff(dict1, dict2)
1-element Array{Pair{Int64,String},1}:
100 => "X"

合并字典：

### 实例

julia> merge(dict1, dict2)
Dict{Int64,String} with 4 entries:
100 => "X"
450 => "W"
220 => "Y"
300 => "Z"

查看字典中的最小值：

### 实例

julia> dict1
Dict{Int64,String} with 2 entries:
100 => "X"
220 => "Y"

julia> findmin(dict1)
("X", 100)

---

## Julia 日期和时间

Source: https://www.runoob.com/julia/julia-date-time.html

## Julia 日期和时间

Julia 通过 Dates 模块提供了以下三个函数来处理日期和时间：

- Date：表示日期，精确到日，只显示日期。
- DateTime：表示日期和时间，精确到毫秒。
- DateTime：表示日时间，精确到纳秒，代表一天 24 小时中的特定时刻。

使用前，我们需要先导入 Dates 模块：

```
import Dates
```

Date 和 DateTime 类型可以通过整数或 Period 类型解析。

Period 基于日期值，表示年、月、日等：

```
Period
Year
Quarter
Month
Week
Day
Hour
Minute
Second
Millisecond
Microsecond
Nanosecond
```

Date 和 DateTime 都是抽象类型 TimeType 的子类型。

下图展示了日期类型间的关系，点击图片可以放大查看：

输出日期的时间：

### 实例

julia> import Dates

julia> rightnow = Dates.Time(Dates.now()) # 时间
08:41:15.917

julia> theday = Dates.Date(2022,5,6) # 日期
2022-05-06

julia> today_date = Dates.today()
2022-05-11

julia> Dates.now(Dates.UTC)
2022-05-11T00:44:20.136

# 格式化时间
julia> Dates.DateTime("20220429 120000", "yyyymmdd HHMMSS")
2022-04-29T12:00:00

julia> Dates.DateTime("19/04/2022 17:42", "dd/mm/yyyy HH:MM")
2022-04-19T17:42:00

下表给出了日期格式代码，通过它们可以格式化我们的日期：

字符 日期/时间 元素 Y 表示年，例如： yyyy => 1984, yy => 84 m 表示年，例如： m => 7 or 07 u 表示月份简写名称，例如： Jun U 表示月份完整名称，例如： January e 表示简写星期几，例如： Mon E 表示完整星期几，例如： Monday d 表示日，例如： 1 or 01 H 表示小时，例如： HH => 00 M 表示分钟，例如： MM => 00 S 表示秒，例如： S => 00 s 表示毫秒，例如： .000

### 实例

julia> Dates.Date("Sun, 27 Sep 2022", "e, d u y")
2022-09-27

julia> Dates.DateTime("Sun, 27 Sep 2022 10:25:10", "e, d u y H:M:S")
2022-09-27T10:25:10

通过上面的实例我们常见了一些日期时间对象，接下来我们就可以使用这些对象来获取数据（包含年、月、日、分、秒、时等）：

### 实例

julia> theday = Dates.Date(2022,5,6) # 创建日期对象
2022-05-06

# 接下来获取 theday 中的数据
julia> Dates.year(theday)
2022

julia> Dates.month(theday)
5

# 获取当前时间的数据
julia> rightnow = Dates.now()
2022-05-11T08:51:45.342

julia> Dates.minute(rightnow)
51

julia> Dates.hour(rightnow)
8

julia> Dates.second(rightnow)
45

# 获取月份及星期几
julia> Dates.dayofweek(theday)
5

julia> Dates.dayname(theday)
"Friday"

julia> Dates.yearmonthday(theday)
(2022, 5, 6)

julia> Dates.dayofweekofmonth(theday)
1

#### 日期运算

我们可以对日期对象进行算术运算。

比如我们计算两个日期相差几天：

### 实例

julia> day1 = Dates.Date(2022,1,17)
2022-01-17

julia> day2 = Dates.Date(2022,3,23)
2022-03-23

julia> day2 - day1
65 days

# 使用不同时间单位
julia> Dates.canonicalize(Dates.CompoundPeriod(day2 - day1))
9 weeks, 2 days

我们也可以对日期相加，比如计算 2 年 6个月后的日期：

### 实例

julia> rightnow = Dates.now()
2022-05-11T09:01:07.946

julia> rightnow + Dates.Year(20) + Dates.Month(6)
2042-11-11T09:01:07.946

# 6 天后的日期
julia> rightnow + Dates.Day(6)
2022-05-17T09:01:07.946

#### 日期范围

Julia 可以通过可迭代的 range(区间范围)对象来创建指定区间的日期。 在下面给出的示例中，我们将创建一个生成每个月的第一天的迭代器。

### 实例

date_range = Dates.Date(2011,1,1):Dates.Month(1):Dates.Date(2022,1,1)
Dates.Date("2011-01-01"):Dates.Month(1):Dates.Date("2022-01-01")

在上的区间范围对象中，我们可以找出其中哪些属于工作日，这里需要为 filter() 创建一个匿名函数，它将根据给定的日期判断是否为工作日：

### 实例

julia> weekdaysfromrange = filter(dy -> Dates.dayname(dy) != "Saturday" && Dates.dayname(dy) != "Sunday" , date_range)
94-element Vector{Dates.Date}:
2011-02-01
2011-03-01
2011-04-01
2011-06-01
2011-07-01
2011-08-01
2011-09-01
2011-11-01
2011-12-01
2012-02-01
⋮
2021-02-01
2021-03-01
2021-04-01
2021-06-01
2021-07-01
2021-09-01
2021-10-01
2021-11-01
2021-12-01

#### 四舍五入日期和时间

我们常用 round()、floor() 和 ceil() 函数来对参数进行向上或向下舍入，同样这些函数也可用于对日期进行四舍五入，以便及时向前或向后调整日期。

### 实例

julia> Dates.now()
2022-05-11T09:17:49.824

julia> Dates.format(round(Dates.DateTime(Dates.now()), Dates.Minute(15)), Dates.RFC1123Format)
"Wed, 11 May 2022 09:15:00"

ceil() 函数将向前调整日期/时间，如下所示：

### 实例

julia> theday = Dates.Date(2022,5,6)
2022-05-06

# 下个月
julia> ceil(theday, Dates.Month)
2022-06-01

#明年
julia> ceil(theday, Dates.Year)
2023-01-01

#下周
julia> ceil(theday, Dates.Week)
2022-05-09

#### 重复日期

我们可以查找一个区间范围内的重复日期，比如每个周日：

### 实例

# 创建区间日期
julia> date_range = Dates.Date(2011,1,1):Dates.Month(1):Dates.Date(2022,1,1)
Dates.Date("2011-01-01"):Dates.Month(1):Dates.Date("2022-01-01")
# 使用 filter() 函数过滤出每个周日
julia> filter(d -> Dates.dayname(d) == "Sunday", date_range)
20-element Vector{Dates.Date}:
2011-05-01
2012-01-01
2012-04-01
2012-07-01
2013-09-01
2013-12-01
2014-06-01
2015-02-01
2015-03-01
2015-11-01
2016-05-01
2017-01-01
2017-10-01
2018-04-01
2018-07-01
2019-09-01
2019-12-01
2020-03-01
2020-11-01
2021-08-01
Unix 时间

UNIX 时间，或称 POSIX 时间是 UNIX 或类 UNIX 系统使用的时间表示方式：从UTC 1970 年 1 月 1 日 0 时 0 分 0 秒起至现在的总秒数，不考虑闰秒。

time() 函数返回 Unix 值：

### 实例

julia> time()
1.652232489777e9

unix2datetime() 函数将 Unix 时间转化为日期/时间对象：

### 实例

julia> Dates.unix2datetime(time())
2022-05-11T01:28:23.493

#### 当下时刻

DateTimes 以毫秒为单位，我们可以使用 Dates.value 函数获取以毫秒计的时间：

### 实例

julia> moment=Dates.now()
2022-05-11T09:31:31.037

julia> Dates.value(moment)
63787944691037

julia> moment.instant
Dates.UTInstant{Millisecond}(Millisecond(63787944691037))

#### 时间和监控

Julia 为我们提供了 @elapsed 宏，它将返回表达式执行所需的时间（秒数）。

计算以下代码需要执行的时间：

### 实例

julia> function foo(n)
for i in 1:n
x = sin(rand())
end
end
foo (generic function with 1 method)

julia> @elapsed foo(100000000)
1.360567967

julia> @time foo(100000000)
1.363258 seconds

#### 更多实例

### 实例

julia> DateTime(2013)
2013-01-01T00:00:00

julia> DateTime(2013,7)
2013-07-01T00:00:00

julia> DateTime(2013,7,1)
2013-07-01T00:00:00

julia> DateTime(2013,7,1,12)
2013-07-01T12:00:00

julia> DateTime(2013,7,1,12,30)
2013-07-01T12:30:00

julia> DateTime(2013,7,1,12,30,59)
2013-07-01T12:30:59

julia> DateTime(2013,7,1,12,30,59,1)
2013-07-01T12:30:59.001

julia> Date(2013)
2013-01-01

julia> Date(2013,7)
2013-07-01

julia> Date(2013,7,1)
2013-07-01

julia> Date(Dates.Year(2013),Dates.Month(7),Dates.Day(1))
2013-07-01

julia> Date(Dates.Month(7),Dates.Year(2013))
2013-07-01

---

## Julia 文件(File)读写

Source: https://www.runoob.com/julia/julia-file.html

## Julia 文件(File)读写

Julia 提供了一些基本的函数来处理文件：

- open() - 打开文件
- read() - 读取文件内容
- close() - 关闭文件

从文件读取或者写入数据需要使用文件句柄。

文件句柄其实就是一个指针，指针就是指向文件中的某个位置。

从一个文件读取数据，应用程序首先要调用操作系统函数并传送文件名，并选一个到该文件的路径来打开文件，打开文件的函数取回一个顺序号，即文件句柄（file handle），该文件句柄对于打开的文件是唯一的识别依据。

#### open() 函数

Julia 可以使用 open() 函数打开一个文件，该函数返回文件句柄：

语法格式：

```
open(filename, mode)
```

filename 为文件名，mode 为读写模式，可以是以下值：

mode(模式) 描述 关键词 `r` read none `w` 写入、创建、截断 `write = true` `a` 写入、创建、追加 `append = true` `r+` 读取, 写入 `read = true, write = true` `w+` 读取, 写入、创建、截断 `truncate = true, read = true` `a+` 读取, 写入、创建、写入、创建、追加 `append = true, read = true`

```

# Windows 打开文件
foo = open("D://runoob-test//julia//runoob-file.txt")

# Linux 或 Mac 打开文件
foo = open("./runoob-test/julia/runoob-file.txt")

```

#### close() 函数

在对文件操作完成后，我们需要对文件进行关闭操作，不然会造成资源泄露等问题，关闭文件使用 close() 函数：

```

# 关闭上面代码中打开文件返回的 foo 句柄
close(foo)

```

在 Julia 中，我们建议任何文件处理函数包装在 do 语句块中，将文件处理函数包装在 do 语句块中的优点是，当该语句块执行完成时，打开的文件将自动关闭，就是会自动调用 close() 函数，如下所示：

### 实例

open("./runoob-test/julia/runoob-file.txt") do file
# 执行文件操作
end

以上代码不需要调用 close() 函数，因为它会自动 close() 函数来关闭文件。

接下来我们创建一个 test.jl 文件，代码如下：

### test.jl 文件代码：

# 打开文件，如果文件不存在就创建 runoob-file.txt
io = open("runoob-file.txt", "w");

# 写入文件内容
write(io, "Hello world!\nRunoob Julia Test");

# 关闭文件
close(io);

执行以上代码：

```
$ julia test.jl
```

执行成功后，我们就在当前代码文件的目录下创建了一个 runoob-file.txt 文件，写入内容如下：

```
Hello world!
Runoob Julia Test
```

我们可以使用 do 语句，这样就不需要使用 close() 函数：

### test.jl 文件代码：

open("runoob-file.txt", "w") do io
write(io, "Hello world!\nRunoob Julia Test\n使用 do 语句")
end;

执行以上代码：

```
$ julia test.jl
```

执行成功后，我们就在当前代码文件的目录下创建了一个 runoob-file.txt 文件，写入内容如下：

```
Hello world!
Runoob Julia Test
使用 do 语句
```

### 实例

# 读取文件内容
txt = open("runoob-file.txt") do file
read(file, String)
end;
# 输出
println(txt)

#### read() 函数

使用 read() 函数，我们可以读取文件的全部内容，例如：

```
txt = read(foo, String)
```

### 实例

# 打开文件
file = open("runoob-file.txt")

# 读取文件内容
txt = read(file, String)

# 关闭文件
close(file)

# 输出
println(txt)

执行以上代码，输出结果如下：

```
$ Julia test.jl
Hello world!
Runoob Julia Test

```

do 语句中使用 read() 函数：

```
open(filename) do file
read(file, String)
end
```

### 实例

# 读取文件内容
txt = open("runoob-file.txt") do file
read(file, String)
end;
# 输出
println(txt)

执行以上代码，输出结果如下：

```
$ julia test.jl
Hello world!
Runoob Julia Test

```

我们可以使用 readlines() 函数将文件内容按行放到数组中，语法如下：

```
readlines(io::IO=stdin; keep::Bool=false)
readlines(filename::AbstractString; keep::Bool=false)
```

### 实例

# 打开文件
file = open("runoob-file.txt")

# 读取文件内容
txt=readlines(file, keep=true)

# 关闭文件
close(file)

# 输出
println(txt)

执行以上代码，输出结果如下：

```
$ julia test.jl
["Hello world!\n", "Runoob Julia Test\n", "使用 do 语句"]

```

我们也可以使用 eachline() 函数逐行处理文件：

### 实例

# 打开文件
open("runoob-file.txt") do file
# 逐行读取文件内容
for ln in eachline(file)
# 输出字符串长度与字符串
println("$(length(ln)), $(ln)")
end
end

执行以上代码，输出结果如下：

```
$ julia test.jl
12, Hello world!
17, Runoob Julia Test
8, 使用 do 语句

```

我们也可以获取当前的行号：

### 实例

open("runoob-file.txt") do f
# 设置行号
line = 1
while !eof(f)
x = readline(f)
println("$line $x")
# 逐行递增
line += 1
end
end

执行以上代码，输出结果如下：

```
$ julia test.jl
1 Hello world!
2 Runoob Julia Test
3 使用 do 语句

```

#### stat() 函数

stat() 函数用于获取文件的信息，格式如下：

```
stat(pathname)
```

### 实例

for n in fieldnames(typeof(stat("runoob-file.txt")))
println(n, ": ", getfield(stat("runoob-file.txt"),n))
end

执行以上代码，输出结果如下：

```
$ julia test.jl
desc: runoob-file.txt
device: 16777222
inode: 35345094
mode: 33188
nlink: 1
uid: 501
gid: 20
rdev: 0
size: 47
blksize: 4096
blocks: 8
mtime: 1.6524042534674733e9
ctime: 1.6524042534674733e9

```

#### abspath() 函数

abspath() 函数用于获取文件的绝对路径，可以映射列表中：

```
stat(pathname)
```

### 实例

println(map(abspath, readdir()))

执行以上代码，输出结果如下：

```

["/Users/RUNOOB/runoob-test/.Rhistory", "/Users/RUNOOB/runoob-test/.vscode",
...
```

#### 更多函数

下表列出来处理文件的其他相关函数：

序号 函数及描述 1

cd(path)

切换当前目录 2

pwd()

获取当前目录 3

readdir(path)

返回当前目录的文件与目录列表。 4

abspath(path)

将当前目录的文件名生成绝对路径。 5

joinpath(str, str, ...)

从参数中组装路径名。 6

isdir(path)

判断提供的路径参数 path 是否是一个目录。 7

splitdir(path)

将路径拆分为目录名和文件名的元组。 8

splitdrive(path)

Windows 上将路径拆分为驱动器号部分和路径部分，在 Unix 系统上，第一个组件始终是空字符串。 9

splitext(path)

如果路径的最后一个组件包含一个点，则将路径拆分为点之前的所有内容以及包括点和点之后的所有内容。 否则，返回一个未修改的参数和空字符串的元组。 10

expanduser(path)

将路径开头的波浪字符 ～ 替换为当前用户的主目录。 11

normpath(path)

规范化路径，删除 "." 和 ".." 目录 12

realpath(path)

如果符号链接来规范化路径，并删除 "." 和 ".." 目录 13

homedir()

获取当前用户的主目录。 14

dirname(path)

获取路径参数 path 的目录部分 15

basename(path)

获取路径参数 path 的文件名部分。

---

## Julia 元编程

Source: https://www.runoob.com/julia/julia-metaprogramming.html

## Julia 元编程

Julia 把自己的代码表示为语言中的数据结构，这样我们就可以编写操纵程序的程序。

元编程也可以简单理解为编写可以生成代码的代码。

元编程（英语：Metaprogramming），是指某类计算机程序的编写，这类计算机程序编写或者操纵其它程序（或者自身）作为它们的资料，或者在编译时完成部分本应在运行时完成的工作。多数情况下，与手工编写全部代码相比，程序员可以获得更高的工作效率，或者给与程序更大的灵活度去处理新的情形而无需重新编译。

编写元程序的语言称之为元语言。被操纵的程序的语言称之为"目标语言"。一门编程语言同时也是自身的元语言的能力称之为"反射"或者"自反"。

-- 维基百科

#### Julia 源代码执行阶段

1、解析原始 Julia 代码: Julia 解析器首先将解析字符串以获得抽象语法树(AST)，AST 是一种结构，它以易于操作的格式包含所有代码。

2、执行已解析的 Julia 代码:: 在这个阶段，执行已解析的 Julia 代码。

当我们在交互式编程环境(REPL)中输入代码并按回车键时，会执行以上两个阶段。

使用元编程工具，我们就可以访问这两个阶段之间的 Julia 代码，即在源代码解析之后，但在执行之前。

### 程序表示

Julia 提供了一个 Meta 类，其中可以用 Meta.parse(str) 来对一个字符串进行解析，用 typeof(e1) 返回 Expr：

### 实例

julia> prog = "1 + 1"
"1 + 1"
julia> ex1 = Meta.parse(prog)
:(1 + 1)

julia> typeof(ex1)
Expr

返回 :(1 + 1) , 这个返回的值由一个冒号和后面的表达式组成，用 typeof(ex1) 返回 Expr。

Expr 对象包含两个部分（ex1 包含了 head 和 args 属性）：

一个标识表达式类型的符号对象。

### 实例

julia> ex1.head
:call

另一个是表达式的参数，可能是符号、其他表达式或字面量：

### 实例

julia> ex1.args
3-element Vector{Any}:
:+
1
1

表达式也可能直接用 Expr 构造：

### 实例

julia> ex2 = Expr(:call, :+, 1, 1)
:(1 + 1)

上面构造的两个表达式：一个通过解析构造，一个通过直接构造，两个是等价的：

### 实例

julia> ex1 == ex2
true

Expr 对象也可以嵌套：

### 实例

julia> ex3 = Meta.parse("(4 + 4) / 2")
:((4 + 4) / 2)

我们也可以使用 Meta.show_sexpr 查看表达式，以下实例展示了嵌套的 Expr：

### 实例

julia> Meta.show_sexpr(ex3)
(:call, :/, (:call, :+, 4, 4), 2)

#### 符号

我们可以通过冒号 : 前缀运算符存储一个未计算但已解析的表达式。

### 实例

julia> ABC = 100
100

julia> :ABC
:ABC

引用下面的整个表达式：

### 实例

julia> :(100-50)
:(100 - 50)

引用算数表达式：

### 实例

julia> ex = :(a+b*c+1)
:(a + b * c + 1)

julia> typeof(ex)
Expr

注意等价的表达式也可以使用 Meta.parse 或者直接用 Expr 构造：

### 实例

julia> :(a + b*c + 1) ==
Meta.parse("a + b*c + 1") ==
Expr(:call, :+, :a, Expr(:call, :*, :b, :c), 1)
true

引用多个表达式也可以在 quote ... end 中包含代码块。

### 实例

julia> ex = quote
x = 1
y = 2
x + y
end
quote
#= none:2 =#
x = 1
#= none:3 =#
y = 2
#= none:4 =#
x + y
end

julia> typeof(ex)
Expr

### 执行表达式

表达式解析后，我们可以使用 eval() 函数来执行：

### 实例

julia> ex1 = :(1 + 2)
:(1 + 2)

julia> eval(ex1)
3

julia> ex = :(a + b)
:(a + b)

julia> eval(ex)
ERROR: UndefVarError: b not defined
[...]

julia> a = 1; b = 2;

julia> eval(ex)
3

### 抽象语法树 (AST)

抽象语法树 (AST) 是一种结构，是源代码语法结构的一种抽象表示。

它以树状的形式表现编程语言的语法结构，树上的每个节点都表示源代码中的一种结构。

我们可以在 dump() 函数查看表达式的层次结构：

### 实例

julia> dump(:(1 * cos(pi/2)))
Expr
head: Symbol call
args: Array{Any}((3,))
1: Symbol *
2: Int64 1
3: Expr
head: Symbol call
args: Array{Any}((2,))
1: Symbol cos
2: Expr
head: Symbol call
args: Array{Any}((3,))
1: Symbol /
2: Symbol pi
3: Int64 2

### 插值

使用值参数直接构造 Expr 对象虽然很强大，但与 Julia 语法相比，Expr 构造函数可能让人觉得乏味。作为替代方法，Julia 允许将字面量或表达式插入到被引用的表达式中。表达式插值由前缀 $ 表示。

在此示例中，插入了变量 a 的值：

### 实例

julia> a = 1;

julia> ex = :($a + b)
:(1 + b)

对未被引用的表达式进行插值是不支持的，这会导致编译期错误：

```
julia> $a + b
ERROR: syntax: "$" expression outside quote
```

在此示例中，元组 (1,2,3) 作为表达式插入到条件测试中：

```
julia> ex = :(a in $:((1,2,3)) )
:(a in (1, 2, 3))
```

在表达式插值中使用 $ 是有意让人联想到字符串插值和命令插值。表达式插值使得复杂 Julia 表达式的程序化构造变得方便和易读。

### 宏

宏提供了一种机制，可以将生成的代码包含在程序的最终主体中。 宏将一组参数映射到返回的 表达式，并且生成的表达式被直接编译，而不需要运行时 eval 调用。 宏参数可能包括表达式、字面量和符号。

这是一个非常简单的宏：

### 实例

julia> macro sayhello()
return :( println("Hello, world!") )
end
@sayhello (macro with 1 method)

宏在Julia的语法中有一个专门的字符 @ (at-sign)，紧接着是其使用 macro NAME ... end 形式来声明的唯一的宏名。在这个例子中，编译器会把所有的 @sayhello 替换成：

```
:( println("Hello, world!") )
```

当 @sayhello 在REPL中被输入时，解释器立即执行，因此我们只会看到计算后的结果：

```
julia> @sayhello()
Hello, world!
```

现在，考虑一个稍微复杂一点的宏：

### 实例

julia> macro sayhello(name)
return :( println("Hello, ", $name) )
end
@sayhello (macro with 1 method)

这个宏接受一个参数 name。当遇到 @sayhello 时，quoted 表达式会被展开并将参数中的值插入到最终的表达式中：

### 实例

julia> @sayhello("human")
Hello, human

我们可以使用函数 macroexpand 查看引用的返回表达式：

### 实例

julia> ex = macroexpand(Main, :(@sayhello("human")) )
:(Main.println("Hello, ", "human"))

julia> typeof(ex)
Expr

我们可以看到 "human" 字面量已被插入到表达式中了。

还有一个宏 @ macroexpand，它可能比 macroexpand 函数更方便：

### 实例

julia> @macroexpand @sayhello "human"
:(println("Hello, ", "human"))
