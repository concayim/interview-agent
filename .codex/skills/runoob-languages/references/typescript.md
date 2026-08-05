# TypeScript - 菜鸟教程

Tutorial: https://www.runoob.com/typescript/ts-tutorial.html

---

## TypeScript 教程

Source: https://www.runoob.com/typescript/ts-tutorial.html

## TypeScript 教程

TypeScript 是 JavaScript 的一个超集，支持 ECMAScript 6 标准（ES6 教程）。

TypeScript 由微软开发的自由和开源的编程语言，在 JavaScript 的基础上增加了静态类型检查的超集。

TypeScript 设计目标是开发大型应用，它可以编译成纯 JavaScript，编译出来的 JavaScript 可以运行在任何浏览器上。

### 为什么选择 TypeScript

TypeScript 是由 Microsoft 开发的开源编程语言，它是 JavaScript 的超集。

TypeScript 通过添加可选的静态类型、接口、枚举等特性，大大增强了 JavaScript 的开发体验。

#### TypeScript 的主要优势

- 静态类型检查：在编译时发现类型错误，减少运行时错误
- 强大的 IDE 支持：智能代码补全、导航和重构
- 更好的代码可读性：类型本身就是最好的文档
- 现代 JavaScript 特性：支持 ES6+ 语法，如箭头函数、模块、类等
- 渐进式迁移：可以逐步将现有 JavaScript 项目迁移到 TypeScript

#### TypeScript 工作原理

TypeScript 不能直接在浏览器中运行，它需要先编译为 JavaScript。这个编译过程会检查类型错误，并将 TypeScript 特有的语法转换为纯 JavaScript。 .ts 文件 TypeScript 编译器 (tsc) .js 文件 浏览器/Node.js ✓

TypeScript 编译器 (tsc) 在编译过程中进行类型检查，如果发现类型错误会报错并阻止编译。编译成功后生成纯 JavaScript 代码，可以在任何浏览器或 Node.js 环境中运行。

### 语言特性

TypeScript 是对 JavaScript 的类型增强扩展，在保持原有语法兼容的基础上，引入了更强的类型系统和工程能力。

#### 核心增强能力（TypeScript 独有）

特性 说明 类型批注 & 编译时检查 在开发阶段发现类型错误，提高代码可靠性 类型推断 无需显式声明类型，编译器可自动推断 类型擦除 编译后移除类型信息，输出纯 JavaScript 接口（Interface） 用于定义对象结构，增强代码约束 枚举（Enum） 提供命名常量集合，提升可读性 泛型（Generics） 支持类型复用，增强函数和类的灵活性 Mixin 实现多重继承的组合模式 命名空间（Namespace） 组织代码结构（旧方案，现多用模块） 元组（Tuple） 固定长度、不同类型的数组 异步（Await） 提供更直观的异步流程控制

#### 来自 ES2015 的语法支持（向下兼容）

TypeScript 同时支持并向下编译 ECMAScript 2015 的核心特性：

特性 说明 类（Class） 面向对象编程支持 模块（Module） 支持模块化开发（import/export） 箭头函数（=>） 更简洁的函数写法，绑定词法作用域 可选参数 函数参数可以省略 默认参数 为参数提供默认值

### JavaScript 与 TypeScript 的区别

TypeScript 是 JavaScript 的超集，扩展了 JavaScript 的语法，因此现有的 JavaScript 代码可与 TypeScript 一起工作无需任何修改，TypeScript 通过类型注解提供编译时的静态类型检查。

TypeScript 可处理已有的 JavaScript 代码，并只对其中的 TypeScript 代码进行编译。

### 第一个 TypeScript 实例

以下实例我们使用 TypeScript 来输出 Hello World!:

### 实例

const hello : string = "Hello World!" console.log(hello)
尝试一下 »

---

## TypeScript 简介

Source: https://www.runoob.com/typescript/ts-intro.html

## TypeScript 简介

TypeScript 是由微软开发并开源的编程语言，它是 JavaScript 的超集，在完全兼容 JavaScript 语法的基础上，增加了可选的静态类型系统和基于类的面向对象编程能力。

TypeScript 代码最终会被编译为纯 JavaScript，可以运行在任何支持 JavaScript 的环境中，包括浏览器、Node.js 和移动端。

### 为什么需要 TypeScript

JavaScript 是一门动态类型语言，变量的类型在运行时才能确定。

这种灵活性在小型项目中表现出色，但随着项目规模扩大，问题逐渐暴露：

问题场景JavaScript 的困境TypeScript 的解决方式 函数参数传错类型运行时才报错，难以提前发现编译阶段即报错，IDE 实时提示 访问不存在的属性返回 `undefined`，行为难以预测编译器直接报错，拒绝通过 大型项目重构改一处，不知道哪里会崩类型系统自动追踪所有引用 团队协作函数接受什么参数、返回什么全靠注释或文档类型签名即文档，IDE 自动补全 代码可读性看函数定义无法直接知道数据结构接口和类型别名让数据结构一目了然

TypeScript 的本质目标不是取代 JavaScript，而是让大规模 JavaScript 工程变得可管理、可维护。它是 JavaScript 的"安全网"，而不是竞争对手。

看一个最直观的例子：

### 实例

// JavaScript：运行时才会报错
function greet(name) {
return "Hello, " + name.toUpperCase();
}

greet(123); // 运行时报错：name.toUpperCase is not a function

### 实例

// TypeScript：编译时就能发现错误
function greet(name: string): string {
return "Hello, " + name.toUpperCase();
}

greet(123); // 编译错误：Argument of type 'number' is not assignable to parameter of type 'string'
greet("runoob"); // 正确：Hello, RUNOOB

### TypeScript 与 JavaScript 的关系

TypeScript 和 JavaScript 之间的关系，可以用一句话概括：TypeScript 是 JavaScript 的超集，即所有合法的 JavaScript 代码，同时也是合法的 TypeScript 代码。

集合关系示意

JavaScript（所有 JS 代码）
⊂ TypeScript（JS + 类型系统 + 新特性）
⊂ 编译产物（标准 JavaScript，可在任意环境运行）

两者的核心差异对比：

对比项JavaScriptTypeScript 类型系统动态类型，运行时确定静态类型，编译时检查（可选） 运行方式直接在浏览器 / Node.js 中运行需先编译为 JS 再运行 错误发现时机运行时编译时（提前发现） IDE 支持基础补全强类型推断、精准补全、重构支持 学习曲线较平缓需额外学习类型系统 现有 JS 兼容性—完全兼容，可渐进式迁移 文件扩展名`.js``.ts` 或 `.tsx`（含 JSX）

TypeScript 支持"渐进式采用"：你不必一次性重写整个项目。可以先把 `.js` 改为 `.ts`，逐步为关键模块添加类型，现有代码照常运行。

### 核心特性

TypeScript 在 JavaScript 的基础上引入了一套完整的类型系统，以下是最常用的核心特性。

#### 基础类型注解

在变量、函数参数、返回值后用冒号声明类型，这是 TypeScript 最基础的写法。

### 实例

// 基础类型：number、string、boolean、null、undefined、symbol、bigint
let age: number = 25;
let username: string = "runoob";
let isActive: boolean = true;

// 数组类型，两种等价写法
let scores: number[] = [90, 85, 92];
let tags: Array<string> = ["typescript", "javascript"];

// 元组：固定长度和类型的数组
let point: [number, number] = [10, 20];
let entry: [string, number] = ["RUNOOB", 100];

// 函数：参数类型 + 返回值类型
function add(a: number, b: number): number {
return a + b;
}

// void：函数无返回值
function log(msg: string): void {
console.log(msg);
}

// 可选参数：参数名后加 ?
function greet(name: string, title?: string): string {
return title ? `${title} ${name}` : name;
}

console.log(greet("RUNOOB")); // 输出：RUNOOB
console.log(greet("runoob", "Mr.")); // 输出：Mr. runoob

#### 接口（Interface）

接口用来描述对象的"形状"（Shape），定义对象应该具备哪些属性和方法。

### 实例

// 定义接口：描述用户对象的结构
interface User {
id: number; // 必填
name: string; // 必填
email?: string; // 可选，使用 ? 标记
readonly role: string; // 只读，赋值后不能修改
}

// 函数参数使用接口类型
function printUser(user: User): void {
console.log(`ID: ${user.id}, Name: ${user.name}`);
if (user.email) {
console.log(`Email: ${user.email}`);
}
}

const admin: User = {
id: 1,
name: "RUNOOB",
email: "runoob@example.com",
role: "admin",
};

printUser(admin);
// 输出：ID: 1, Name: RUNOOB
// 输出：Email: runoob@example.com

// admin.role = "user"; // 错误：Cannot assign to 'role' because it is a read-only property.

// 接口继承
interface AdminUser extends User {
permissions: string[];
}

#### 类型别名（Type Alias）

类型别名用 type 关键字创建，可以为任何类型定义别名，包括联合类型、交叉类型等复杂结构。

### 实例

// 联合类型：变量可以是多种类型之一
type ID = string | number;

let userId: ID = "abc-123"; // 合法
userId = 456; // 也合法

// 字面量类型：限制变量只能取特定的值
type Direction = "up" | "down" | "left" | "right";
type Status = "pending" | "active" | "inactive";

function move(dir: Direction): void {
console.log(`Moving ${dir}`);
}

move("up"); // 正确
// move("diagonal"); // 错误：不在字面量范围内

// 交叉类型：合并多个类型
type WithTimestamp = {
createdAt: Date;
updatedAt: Date;
};

type UserRecord = User & WithTimestamp; // 同时具有 User 和 WithTimestamp 的所有属性

// 函数类型别名
type Transformer<T, U> = (input: T) => U;
const toNumber: Transformer<string, number> = (s) => parseInt(s, 10);

#### 泛型（Generics）

泛型允许编写可复用的组件，同时保持类型安全。可以把泛型理解为"类型的变量"——使用时再传入具体类型。

### 实例

// 没有泛型的写法：只能处理 number 类型
function firstNumber(arr: number[]): number {
return arr[0];
}

// 使用泛型：T 是类型参数，调用时确定具体类型
function first<T>(arr: T[]): T {
return arr[0];
}

const n = first<number>([1, 2, 3]); // n 的类型是 number
const s = first<string>(["a", "b"]); // s 的类型是 string
const inferred = first([true, false]); // TypeScript 自动推断 T 为 boolean

// 泛型接口
interface ApiResponse<T> {
data: T;
status: number;
message: string;
}

// 使用泛型接口
const userResponse: ApiResponse<User> = {
data: { id: 1, name: "RUNOOB", role: "admin" },
status: 200,
message: "success",
};

// 泛型约束：限制 T 必须有 id 属性
function getById<T extends { id: number }>(items: T[], id: number): T | undefined {
return items.find(item => item.id === id);
}

#### 枚举（Enum）

枚举用于定义一组命名常量，使代码中的"魔法数字"或"魔法字符串"变得有意义。

### 实例

// 数字枚举（默认从 0 开始自增）
enum Direction {
Up, // 0
Down, // 1
Left, // 2
Right, // 3
}

console.log(Direction.Up); // 输出：0
console.log(Direction[0]); // 输出："Up"（反向映射）

// 字符串枚举（更推荐，调试时值更可读）
enum Color {
Red = "RED",
Green = "GREEN",
Blue = "BLUE",
}

function paint(color: Color): void {
console.log(`Painting in ${color}`);
}

paint(Color.Red); // 输出：Painting in RED
// paint("red"); // 错误：字符串 "red" 不是 Color 类型

// const 枚举：编译后会被内联，减少运行时开销
const enum HttpStatus {
OK = 200,
NotFound = 404,
InternalError = 500,
}

const status: HttpStatus = HttpStatus.OK; // 编译后直接替换为数字 200

#### 类型推断

TypeScript 不需要为每个变量都手动标注类型，编译器会根据赋值自动推断。

### 实例

// TypeScript 自动推断变量类型，无需显式标注
let count = 0; // 推断为 number
let name = "RUNOOB"; // 推断为 string
let flag = true; // 推断为 boolean

// 推断数组元素类型
let numbers = [1, 2, 3]; // 推断为 number[]

// 推断函数返回值类型
function double(n: number) { // 返回值自动推断为 number
return n * 2;
}

// 对象字面量推断
const config = {
host: "localhost", // string
port: 3000, // number
debug: false, // boolean
};

// config.port = "3000"; // 错误：Type 'string' is not assignable to type 'number'.

一个实用建议：函数的参数通常需要显式标注类型（因为调用方无法从参数值推断类型），而函数体内的局部变量和返回值大多数情况下依赖推断即可，不必全部手动标注。

#### 类与访问修饰符

TypeScript 完整支持 ES6 的类语法，并新增了 public、private、protected、readonly 四种访问修饰符。

### 实例

class Animal {
readonly name: string; // 只读，初始化后不能修改
private age: number; // 私有，只能在类内部访问
protected species: string; // 受保护，子类可访问

constructor(name: string, age: number, species: string) {
this.name = name;
this.age = age;
this.species = species;
}

// public 方法（默认即为 public，可省略）
public introduce(): string {
return `I'm ${this.name}, a ${this.species}.`;
}

// getter：像访问属性一样调用方法
get info(): string {
return `${this.name} (${this.age} years old)`;
}
}

class Dog extends Animal {
private breed: string;

constructor(name: string, age: number, breed: string) {
super(name, age, "Canis lupus familiaris"); // 调用父类构造函数
this.breed = breed;
}

// 子类可以访问 protected 属性
describe(): string {
return `${this.name} is a ${this.species}, breed: ${this.breed}`;
}
}

const dog = new Dog("RUNOOB", 3, "Labrador");
console.log(dog.introduce()); // 输出：I'm RUNOOB, a Canis lupus familiaris.
console.log(dog.info); // 输出：RUNOOB (3 years old)
// console.log(dog.age); // 错误：私有属性无法从外部访问

#### 装饰器（Decorator）

装饰器是一种元编程语法，用于在类、方法、属性上附加额外行为，常见于 Angular、NestJS 等框架。

TypeScript 5.0 已将装饰器实现为正式标准（Stage 3 提案），不再需要开启 `experimentalDecorators` 标志。

### 实例

// 类装饰器：在类定义时执行
function sealed(constructor: Function) {
Object.seal(constructor); // 禁止添加新属性
Object.seal(constructor.prototype);
}

// 方法装饰器：可用于日志记录、权限校验等
function log(target: any, propertyKey: string, descriptor: PropertyDescriptor) {
const original = descriptor.value;
descriptor.value = function (...args: any[]) {
console.log(`Calling ${propertyKey} with args:`, args);
const result = original.apply(this, args);
console.log(`${propertyKey} returned:`, result);
return result;
};
}

@sealed
class Calculator {
@log
add(a: number, b: number): number {
return a + b;
}
}

const calc = new Calculator();
calc.add(1, 2);
// 输出：Calling add with args: [ 1, 2 ]
// 输出：add returned: 3

### 应用领域

TypeScript 凭借其对 JavaScript 的完全兼容性，几乎可以用在所有 JavaScript 能运行的地方，并且在大型项目中带来额外收益。

#### 前端 Web 开发

这是 TypeScript 使用最广泛的场景。三大主流前端框架均对 TypeScript 提供一级支持：

框架TypeScript 支持情况典型场景 Angular官方语言，默认使用 TypeScript，无法绕过企业级 SPA、后台管理系统 React官方提供 `@types/react`，Create React App 和 Vite 均内置 TypeScript 模板电商、内容平台、中后台应用 VueVue 3 核心代码用 TypeScript 重写，组合式 API 对 TypeScript 友好度显著提升中小型项目、渐进式迁移 Next.js / Nuxt内置 TypeScript 支持，新项目默认开启全栈 SSR/SSG 应用

#### 后端 Node.js 开发

TypeScript 在 Node.js 后端开发中同样普及。NestJS 是目前最流行的 TypeScript 后端框架，采用与 Angular 相近的模块化架构，内置依赖注入和装饰器支持。

此外，Deno 运行时从第一天起就原生支持 TypeScript，无需任何配置即可直接运行 `.ts` 文件。Bun 同样内置 TypeScript 支持。

#### 命令行工具与脚本

借助 `ts-node`、`tsx` 等工具，TypeScript 代码可以在不预编译的情况下直接执行，非常适合编写构建脚本和 CLI 工具。

#### 移动端与跨平台

React Native 完整支持 TypeScript，大量企业级移动应用（如 Microsoft Office 移动版、Shopify 等）均采用此技术栈。Expo 的新项目模板默认也以 TypeScript 为基础。

#### 游戏开发与图形

Babylon.js（微软出品的 3D 引擎）完全用 TypeScript 编写，并将 TypeScript 类型定义作为一等公民。Phaser 等 2D 游戏引擎同样提供完整的类型定义。

### 发展历史

TypeScript 的诞生和成长历程与 JavaScript 生态的演变密切相关，每一次重大版本发布背后都有清晰的工程需求驱动。

#### 起源（2010 - 2012）

TypeScript 的故事始于微软内部的一个工程挑战：如何让大型 JavaScript 项目变得可维护。

设计者 Anders Hejlsberg 是业界传奇——C# 的首席架构师，也是 Turbo Pascal 和 Delphi 的创造者。他在 2010 年左右开始设计这门语言，核心目标是在不破坏 JavaScript 兼容性的前提下，引入静态类型检查。

微软推出 TypeScript 的另一个背景是 Windows 8 的发布——应用程序可以使用 HTML + JavaScript 开发，微软希望吸引 .NET 程序员参与，因此 TypeScript 的许多语法有意与 C# 和 .NET 保持相近。

2012 年 10 月，TypeScript 0.8 公开发布，此前已在微软内部开发了约两年。

#### 早期版本（2013 - 2015）

时间版本重要事件 2013-060.9正式稳定版发布，性能大幅提升，IDE 整合更完善 2014-041.0第一个正式稳定版，引入类、接口、模块，支持编译为标准 JavaScript 2014-07—TypeScript 编译器代码开源，托管于 GitHub，接受社区贡献 2015-04—微软发布 Visual Studio Code，内置对 TypeScript 的深度支持，两者相互成就 2015-071.5引入 ES6 模块语法支持、装饰器（实验性）、命名空间

#### 高速成长（2016 - 2019）

这一阶段是 TypeScript 从小众工具走向主流的关键时期，核心推动力来自框架生态的选择。

时间版本 / 事件里程碑意义 2016-09TypeScript 2.0引入非空类型（`--strictNullChecks`），这是 TypeScript 类型安全最重要的特性之一；引入标记联合类型 2016Angular 2 发布Google Angular 团队宣布采用 TypeScript 作为官方开发语言，将 TypeScript 带入大众视野 2017TypeScript 2.x 系列条件类型、映射类型、infer 关键字相继推出，类型系统表达能力大幅提升 2018-07TypeScript 3.0项目引用（Project References），支持大型 monorepo；引入 unknown 类型（比 any 更安全的顶级类型） 2019DefinitelyTyped社区维护的类型定义仓库 DefinitelyTyped（`@types/*`）超过 7000 个包，基本覆盖所有主流 npm 库

#### 成熟与普及（2020 - 2022）

时间版本 / 事件重要内容 2020-08TypeScript 4.0可变元组类型、标记元组元素；编辑器体验大幅改善 2021TypeScript 4.x 系列模板字符串类型（Template Literal Types）、内置工具类型增强、`noImplicitOverride` 等严格模式增强 2022State of JS 调查TypeScript 使用率首次超越纯 JavaScript，成为 JS 生态中最受欢迎的超集语言 2022Vue 3 / Vite 普及Vue 3 核心用 TypeScript 重写，配合 Vite 让前端工程的 TypeScript 体验达到新高度

#### 现代阶段（2023 至今）

时间版本重要特性 2023-03TypeScript 5.0现代化装饰器（Stage 3 标准）、const 类型参数、所有枚举升级为联合枚举，性能显著提升 2023-08TypeScript 5.2引入 Explicit Resource Management（`using` 关键字），解决数据库连接、文件句柄等资源的确定性清理问题 2024-03TypeScript 5.4闭包中类型收窄改进，`NoInfer` 工具类型 2024-06TypeScript 5.5类型谓词推断，正则表达式语法检查，让类型系统对代码逻辑理解更精准 2024-11TypeScript 5.7支持 `--target es2024`，未初始化变量检测更可靠 2025-03TypeScript 5.8条件和索引访问类型的返回值检查增强；支持在 `--module nodenext` 下通过 `require()` 加载 ESM 模块（需 Node.js 22+）

#### 未来展望：用 Go 重写编译器

2025 年初，微软宣布了一项重大决定：以 Go 语言重写 TypeScript 编译器，项目代号 Corsa。

官方表示，重写后的编译器性能提升预计超过 10 倍，大型项目的构建时间将从分钟级降至秒级。

按照规划，Go 版本将在 TypeScript 7.x 正式发布，目前 5.x / 6.x 的 TypeScript 版本继续正常迭代和维护。预计 2025 年末完成重构工作。

这次重写并不会改变 TypeScript 语言本身的语法或类型系统，对开发者而言是透明的升级——你的代码不需要任何改动，只是编译速度大幅提升。

### TypeScript 与其他强类型语言的对比

如果你有其他强类型语言的背景，以下对比可以帮助你快速定位 TypeScript 的设计理念。

对比项TypeScriptJava / C#Go 类型系统结构化类型（鸭子类型）名义类型（必须显式声明继承关系）结构化类型（隐式接口） 空值安全开启 `strictNullChecks` 后支持Java 需借助注解或 Optional，C# 8+ 支持通过 error 值和 nil 检查 泛型支持，语法与 Java/C# 相近支持，Java 有类型擦除限制Go 1.18+ 支持，语法较简洁 编译产物JavaScript（在 JS 环境运行）字节码（在 JVM / CLR 上运行）原生机器码 运行时类型检查无（类型信息编译后抹除）有（反射机制）无 学习曲线对 JS 开发者友好，可渐进采用对初学者较陡相对简洁，学习曲线中等

TypeScript 使用的是结构化类型系统（Structural Typing）：只要两个类型的"形状"一致，就认为它们是兼容的，不要求显式声明继承关系。这与 Java/C# 的名义类型（Nominal Typing）有本质区别，也是 TypeScript 能与 JavaScript 无缝互操作的关键原因。

### TypeScript 的局限性

TypeScript 不是银弹，了解它的边界同样重要。

局限说明应对思路 运行时无类型类型信息在编译后全部抹除，运行时无法依赖类型做判断需要运行时校验时使用 `zod`、`io-ts` 等库 编译步骤相比纯 JS 多了一个编译环节，增加了工程复杂度现代构建工具（Vite、esbuild）对此已有良好优化 any 类型逃生舱滥用 `any` 会让类型检查形同虚设开启 `noImplicitAny` 和 `strict` 模式，配合 ESLint 规则 类型体操门槛复杂的条件类型和映射类型对初学者不友好大多数业务代码不需要高级类型，循序渐进即可 第三方库支持少数老旧库缺少类型定义，需要自行编写 .d.ts先查 `@types/*`，确实没有再手写声明文件

---

## TypeScript 安装

Source: https://www.runoob.com/typescript/ts-install.html

## TypeScript 安装

本文介绍 TypeScript 环境的安装。

我们需要使用到 npm 工具安装，如果你还不了解 npm，可以参考我们的NPM 使用介绍。

#### NPM 安装 TypeScript

如果你的本地环境已经安装了 npm 工具，可以使用以下命令来安装。

使用国内镜像：

```
npm config set registry https://registry.npmmirror.com
```

安装 typescript：

```
npm install -g typescript
```

安装完成后我们可以使用 tsc 命令来执行 TypeScript 的相关代码，以下是查看版本号：

```
$ tsc -v
Version 3.2.2
```

然后我们新建一个 app.ts 的文件，代码如下：

var message:string = "Hello World" console.log(message)

通常我们使用 .ts 作为 TypeScript 代码文件的扩展名。

然后执行以下命令将 TypeScript 转换为 JavaScript 代码：

```
tsc app.ts
```

这时候在当前目录下（与 app.ts 同一目录）就会生成一个 app.js 文件，代码如下：

var message = "Hello World"; console.log(message);

使用 node 命令来执行 app.js 文件：

```
$ node app.js
Hello World
```

TypeScript 转换为 JavaScript 过程如下图：

### VS Code 介绍

很多 IDE 都有支持 TypeScript 插件，如：VS Code，Sublime Text 2，WebStorm / PHPStorm，Eclipse 等。

本章节主要介绍 VS Code，VS Code 是一个可以运行于 Mac OS X、Windows 和 Linux 之上的，针对于编写现代 Web 和云应用的跨平台源代码编辑器，由 Microsoft 公司开发。

VS Code 教程：https://www.runoob.com/vscode/vscode-tutorial.html

另外国内阿里与字节也有基于 VS Code 开发的 AI IDE：

- 阿里 Qoder：https://qoder.com/
- 字节 Trae：https://www.trae.com.cn/

#### Windows 上安装 Visual Studio Code

1、下载 Visual Studio Code。

2、双击 VSCodeSetup.exe 图标 安装。

3、安装完成后，打开 Visual Studio Code 界面类似如下：

4、 我们可以在左侧窗口中点击当前编辑的代码文件，选择 open in command prompt（在终端中打开），这时候我们就可以在屏幕的右侧下半部分使用 tsc 命令来执行 TypeScript 文件代码了。

#### Mac OS X 安装 Visual Studio Code

Mac OS X 安装配置 Visual Studio Code 可以查看： https://code.visualstudio.com/Docs/editor/setup

#### Linux 安装 Visual Studio Code

Linux 安装配置 Visual Studio Code 可以查看： https://code.visualstudio.com/Docs/editor/setup

---

## 使用 VS Code 开发 TypeScript

Source: https://www.runoob.com/typescript/ts-vscode.html

## 使用 VS Code 开发 TypeScript

很多 IDE 都有支持 TypeScript 插件，如：VS Code，Sublime Text 2，WebStorm / PHPStorm，Eclipse 等。

本章节主要介绍 VS Code，VS Code 是一个可以运行于 Mac OS X、Windows 和 Linux 之上的，针对于编写现代 Web 和云应用的跨平台源代码编辑器，由 Microsoft 公司开发。

VS Code 原生内置 TypeScript 语言语法支持，但不含 TypeScript 编译工具 tsc，如需将 TS 源码转译为 JS（执行命令 tsc HelloWorld.ts），需要在全局或项目工作区单独安装 TypeScript 编译器。

VS Code 教程：https://www.runoob.com/vscode/vscode-tutorial.html

另外国内阿里与字节也有基于 VS Code 开发的 AI IDE：

#### 安装 TypeScript 编译器

最简安装方式是通过 Node.js 包管理器 npm 安装，已装好 npm 的前提下，执行下方命令即可全局（-g）安装：

```

npm install -g typescript
```

安装完成后可通过输出版本号验证：

```

tsc --version
```

#### Windows 上安装 Visual Studio Code

1、下载 Visual Studio Code。

2、双击 VSCodeSetup.exe 图标 安装。

3、安装完成后，打开 Visual Studio Code 界面类似如下：

4、 我们可以在左侧窗口中点击当前编辑的代码文件，选择 open in command prompt（在终端中打开），这时候我们就可以在屏幕的右侧下半部分使用 tsc 命令来执行 TypeScript 文件代码了。

#### Mac OS X 安装 Visual Studio Code

Mac OS X 安装配置 Visual Studio Code 可以查看： https://code.visualstudio.com/Docs/editor/setup

#### Linux 安装 Visual Studio Code

Linux 安装配置 Visual Studio Code 可以查看： https://code.visualstudio.com/Docs/editor/setup

### 入门示例：Hello World

我们以 Node.js 环境的极简 Hello World 项目为例，新建项目文件夹并打开 VS Code：

```

mkdir HelloWorld
cd HelloWorld
code .
```

在资源管理器中新建文件 helloworld.ts。

写入下述 TS 代码，能看到 TS 专属关键字 let 与字符串类型标注：

```

let message : string = "Hello World";
console.log(message);
```

打开【集成终端】（快捷键 kb(workbench.action.terminal.toggleTerminal)），输入 tsc helloworld.ts 编译代码，编译成功后会自动生成 JS 文件 helloworld.js。

本机已安装 Node.js 时，运行 node helloworld.js 即可执行代码。

打开编译后的 helloworld.js 可以发现：代码移除了类型声明，let 被降级为 ES5 的 var：

```

var message = "Hello World";
console.log(message);
```

### 智能提示（IntelliSense）

智能提示提供代码自动补全、悬浮文档、函数参数签名提示，提升编码效率与准确率。

VS Code 对单个 TS 文件、带 `tsconfig.json` 配置的 TS 项目均提供全套智能提示。

#### 悬浮文档提示

鼠标悬浮在任意 TS 标识符上，快速查看变量/函数的类型与说明文档；
快捷键 `kb(editor.action.showHover)` 可在光标位置手动唤起悬浮提示。

#### 函数签名提示

编写函数调用代码时，编辑器自动展示函数入参说明，并高亮当前正在填写的参数；
输入左括号 `(` 或逗号 `,` 自动触发签名提示，快捷键 `kb(editor.action.triggerParameterHints)` 手动唤起。

### 代码片段（Snippets）

除智能补全外，VS Code 内置基础 TS 代码片段，输入关键字时自动弹出快捷模板。

可安装第三方扩展扩充片段库，也能自定义 TS 代码片段，详情参考【用户自定义代码片段】文档。

小贴士：在配置文件中将 `editor.snippetSuggestions` 设为 `"none"` 可关闭片段提示；如需保留，可配置片段展示优先级：`top`（置顶）、`bottom`（置底）、`inline`（和普通提示混排、字母排序，默认配置）。

### 错误与警告提示

TS 语言服务实时校验代码，自动标记语法与逻辑问题：

- 编辑器状态栏汇总当前项目错误、警告总数；
- 点击状态栏统计数字，或快捷键 `kb(workbench.actions.view.problems)` 打开【问题】面板，罗列全部报错；
- 存在异常的代码行，行内标红提示，同时在右侧缩略标尺标注错误位置。

快捷键 `kb(editor.action.marker.nextInFiles)` / `kb(editor.action.marker.prevInFiles)` 跳转当前文件上下一处问题，弹窗展示详情与可用快速修复。

### 代码跳转导航

依托导航功能快速浏览大型 TS 项目源码：

- 跳转到定义 `kb(editor.action.revealDefinition)`：定位标识符源码定义处；
- 预览定义 `kb(editor.action.peekDefinition)`：弹窗预览源码，不切换当前文件；
- 查找所有引用 `kb(editor.action.goToReferences)`：列出项目内所有调用该标识符的位置；
- 跳转至类型定义：定位变量所属的类型源码（类实例会跳转至类本体，而非实例创建代码）；
- 跳转至实现 `kb(editor.action.goToImplementation)`：查找接口/抽象方法的具体实现代码。

通过命令面板（`kb(workbench.action.showCommands)`）的符号检索实现快速跳转：

- 当前文件内查找符号 `kb(editor.action.gotoSymbol)`
- 全工作区检索符号 `kb(workbench.action.showAllSymbols)`

### 代码格式化

VS Code 内置 TS 格式化工具，默认规则开箱即用。

通过 `js/ts.format.*` 系列配置自定义格式化规则（例如大括号单独换行）不需要内置格式化时，设置:

```
"js/ts.format.enable": false
```

关闭。

如需贴合团队自定义编码规范，可在插件市场安装专用格式化扩展。

### 代码重构

VS Code 自带常用 TS 重构能力，如提取函数、提取常量。选中待重构代码，点击编辑器侧边灯泡图标或快捷键 `kb(editor.action.quickFix)` 调出重构菜单。

更多重构规则、自定义重构快捷键参考【TypeScript代码重构】文档。

#### 统一重命名

选中变量/方法，快捷键 `kb(editor.action.rename)` 一键修改项目内所有同名标识符。

### 调试功能

VS Code 原生完善支持 TS 调试，兼容源码映射（SourceMap），支持断点、变量查看、调用栈追溯、调试控制台交互；详情查阅【TS调试指南】与通用调试文档。

#### 前端网页调试

借助内置 Chrome/Edge 调试插件，或 Firefox 调试扩展，调试浏览器端 TS 代码。

#### 后端Node调试

使用内置调试器调试 Node.js 服务，官方配套教程可快速上手 Express 项目调试。

### 代码校验工具（Linter）

Linter 用于检查代码隐患、规范编码风格。VS Code 未内置 TS 校验工具，可在插件市场安装对应扩展。

主流工具 ESLint 全面兼容 TS，安装 ESLint 插件后，编辑器实时标错，大量问题支持一键快速修复；`typescript-eslint` 文档指导 TS 项目 ESLint 环境配置。

---

## TypeScript 特性

Source: https://www.runoob.com/typescript/ts-features.html

## TypeScript 特性

相对 JavaScript，TypeScript 增加了许多关键功能，特别是围绕类型系统和代码结构的增强功能。

TypeScript 的一些关键特性：

- 静态类型检查：TypeScript 在编译时就会检查代码的类型是否匹配，能够发现很多潜在的错误。即使是简单的错误（例如拼写错误或类型不一致），也可以在编写代码时被捕获到。
- 类型推断：TypeScript 能够自动推断变量的类型。比如当你声明一个变量并赋值时，TypeScript 会根据赋值来推断这个变量的类型，不需要每次都显式声明类型。
- 接口和类型定义：TypeScript 提供了 `interface` 和 `type` 关键字，允许你定义复杂的数据结构。这对于项目中不同部分的代码协作和数据交互来说非常重要。
- 类和模块支持：TypeScript 支持面向对象编程中的类（class）概念，增加了构造函数、继承、访问控制修饰符（如 `public`、`private`、`protected`），并且支持 ES 模块化规范。
- 工具和编辑器支持：TypeScript 拥有良好的编辑器支持，特别是与 Visual Studio Code 集成时，能提供智能提示、自动补全、重构等工具，使开发过程更高效。
- 兼容 JavaScript：TypeScript 是 JavaScript 的超集，这意味着所有合法的 JavaScript 代码都是合法的 TypeScript 代码。这使得 JavaScript 项目可以逐步迁移到 TypeScript，而无需完全重写。

以下是 TypeScript 增加的主要功能：

#### 1. 静态类型

TypeScript 的最大特性就是增加了静态类型系统。在 TypeScript 中，开发者可以显式地声明变量、参数、返回值的类型，这样可以在编译时捕获很多潜在的类型错误。常见类型包括 `number`、`string`、`boolean`、`array`、`tuple`、`enum` 等，此外也支持自定义类型。

```

let name: string = "Alice";
let age: number = 25;

```

#### 2. 类型推断

TypeScript 可以自动推断变量类型，即使不显式声明类型，TypeScript 也会根据变量的赋值内容来推断类型，从而在大多数情况下减少类型注解的书写量。

```
let name = "Alice"; // 推断为 string

```

#### 3. 接口 (Interfaces)

TypeScript 提供了接口，允许定义复杂的对象结构。接口可以定义属性和方法，还可以通过 `implements` 关键字实现接口，或者通过 `extends` 进行扩展，便于定义复杂的数据类型。

```

interface Person {
name: string;
age: number;
greet(): void;
}

class Student implements Person {
constructor(public name: string, public age: number) {}

greet() {
console.log(`Hello, my name is ${this.name}`);
}
}

```

#### 4. 类型别名 (Type Aliases)

类型别名 (`type`) 可以为复杂的类型定义简短的别名，便于代码复用。

```
type StringOrNumber = string | number;
let value: StringOrNumber = 42;

```

#### 5. 枚举 (Enums)

TypeScript 引入了 `enum` 类型，用于定义一组命名的常量，提高代码的可读性。枚举在 JavaScript 中没有直接的对应。

```
enum Direction {
Up,
Down,
Left,
Right,
}
let dir: Direction = Direction.Up;

```

#### 6. 元组 (Tuples)

元组允许定义具有固定数量和类型的数组。它适用于需要固定数据结构的场景，比如坐标或 RGB 颜色值。

```
let point: [number, number] = [10, 20];

```

#### 7. 访问控制修饰符 (Access Modifiers)

TypeScript 在类中提供了 `public`、`private` 和 `protected` 修饰符，允许控制属性或方法的可见性，支持更好的封装。

```
class Person {
private name: string;
protected age: number;
public constructor(name: string, age: number) {
this.name = name;
this.age = age;
}
}

```

#### 8. 抽象类 (Abstract Classes)

TypeScript 支持抽象类，抽象类不能直接实例化，需要由子类实现。抽象类适用于定义通用行为和抽象方法的类层次结构。

```
abstract class Animal {
abstract makeSound(): void;
}

class Dog extends Animal {
makeSound() {
console.log("Woof!");
}
}

```

#### 9. 泛型 (Generics)

TypeScript 支持泛型，允许在类、接口和函数中使用参数化类型，使得代码可以适应不同的类型需求，同时保持类型安全。

```
function identity<T>(value: T): T {
return value;
}

let num = identity<number>(42);

```

#### 10. 模块和命名空间

TypeScript 提供了基于 ES6 的模块系统，使用 `import` 和 `export` 导入和导出模块。此外，TypeScript 还支持命名空间（Namespace），用于组织代码和避免命名冲突。

```
// math.ts
export function add(a: number, b: number): number {
return a + b;
}

// main.ts
import { add } from "./math";
console.log(add(2, 3));

```

#### 11. 类型守卫 (Type Guards)

TypeScript 提供了类型守卫，可以在代码中检查变量类型，帮助编译器推断更加具体的类型。这对于联合类型尤为重要。

```
function printId(id: string | number) {
if (typeof id === "string") {
console.log(id.toUpperCase());
} else {
console.log(id.toFixed(2));
}
}

```

#### 12. 可选链和空值合并运算符

TypeScript 增加了 JavaScript 的可选链 (`?.`) 和空值合并运算符 (`??`)，简化了代码中对可能为 `null` 或 `undefined` 值的处理。

```
let user = { name: "Alice", address: { city: "Wonderland" } };
console.log(user?.address?.city); // 如果 address 存在则输出 city，否则返回 undefined

let value = null;
console.log(value ?? "default"); // 如果 value 为 null 或 undefined，则返回 "default"

```

#### 13. 类型兼容性和工具类型

TypeScript 提供了一些工具类型，如 `Partial`、`Pick`、`Readonly`、`Record` 等，这些类型可以帮助生成新的类型，简化类型定义。

```
interface Todo {
title: string;
description: string;
}

let partialTodo: Partial<Todo> = { title: "Learn TypeScript" }; // 可选属性

```

#### 14. 编译期错误检查

TypeScript 提供的编译期错误检查可以捕获 JavaScript 中不易发现的错误，如拼写错误、类型不匹配等，帮助提升代码质量。

#### 15. ES 新特性支持

TypeScript 提前支持了一些还未在所有环境中普及的 ES 特性，如装饰器（Decorators）、异步迭代器等，且能够将其编译成兼容 JavaScript 版本。

通过这些特性，TypeScript 提供了更安全、更结构化的代码能力，在大型项目和多人协作中尤其具有优势。

---

## TypeScript 基础语法

Source: https://www.runoob.com/typescript/ts-basic-syntax.html

## TypeScript 基础语法

TypeScript 程序由以下几个部分组成：

- 模块
- 函数
- 变量
- 语句和表达式
- 注释

#### 第一个 TypeScript 程序

我们可以使用以下 TypeScript 程序来输出 "Hello World" ：

### Runoob.ts 文件代码：

const hello : string = "Hello World!" console.log(hello)
尝试一下 »

以上代码首先通过 tsc 命令编译：

```
tsc Runoob.ts
```

得到如下 js 代码：

### Runoob.js 文件代码：

var hello = "Hello World!"; console.log(hello);

最后我们使用 node 命令来执行该 js 代码。

```
$ node Runoob.js
Hello World
```

整个流程如下图所示：

我们可以同时编译多个 ts 文件：

```
tsc file1.ts file2.ts file3.ts
```

tsc 常用编译参数如下表所示：

序号 编译参数说明 1.

--help

显示帮助信息 2.

--module

载入扩展模块 3.

--target

设置 ECMA 版本 4.

--declaration

额外生成一个 .d.ts 扩展名的文件。

```

tsc ts-hw.ts --declaration

```

以上命令会生成 ts-hw.d.ts、ts-hw.js 两个文件。 5.

--removeComments

删除文件的注释 6.

--out

编译多个文件并合并到一个输出的文件 7.

--sourcemap

生成一个 sourcemap (.map) 文件。

sourcemap 是一个存储源代码与编译代码对应位置映射的信息文件。 8.

--module noImplicitAny

在表达式和声明上有隐含的 any 类型时报错 9.

--watch

在监视模式下运行编译器。会监视输出文件，在它们改变时重新编译。

### TypeScript 保留关键字

TypeScript 保留关键字如下表所示：

关键字说明`abstract`用于定义抽象类或抽象方法。`any`表示任意类型，禁用类型检查。`as`类型断言，用于将某种类型转换为另一种类型。`await`用于异步函数中，暂停代码执行直到 Promise 解决。`boolean`表示布尔类型。`break`退出循环或 switch 语句。`case`用于 switch 语句中的分支。`catch`用于捕获异常。`class`用于定义类。`const`定义常量变量。`continue`跳过当前循环，继续下一次循环。`debugger`启动调试器，暂停代码执行。`declare`声明一个变量或模块，通常用于类型声明文件。`default`定义 switch 语句的默认分支。`delete`删除对象的属性或数组的元素。`do`用于 `do...while` 循环。`else`定义条件语句中的 else 部分。`enum`定义枚举类型。`export`用于从模块中导出变量、函数或类。`extends`用于类的继承，表示类继承其他类。`false`布尔值 `false`。`finally`定义 `try...catch` 语句中的最终执行代码块。`for`用于 `for` 循环。`from`用于模块导入语句，指定模块的来源。`function`定义函数。`get`用于对象的 getter 方法。`if`用于条件判断。`implements`用于类实现接口。`import`用于从模块中导入内容。`in`用于检查对象中是否包含指定的属性，或用于 `for...in` 循环。`infer`用于条件类型中推断类型。`instanceof`检查对象是否是指定类的实例。`interface`用于定义接口。`let`定义块级作用域的变量。`module`定义模块（在较早的 TypeScript 版本中使用）。`namespace`定义命名空间（在较早的 TypeScript 版本中使用）。`new`创建类的实例。`null`表示空值。`number`表示数字类型。`object`表示非原始类型。`of`用于 `for...of` 循环。`package`用于模块系统，标识包。`private`用于类成员的访问修饰符，表示私有。`protected`用于类成员的访问修饰符，表示受保护的。`public`用于类成员的访问修饰符，表示公共的。`readonly`表示只读属性。`require`用于导入 CommonJS 模块。`return`退出函数并可返回值。`set`用于对象的 setter 方法。`string`表示字符串类型。`super`用于调用父类的方法或构造函数。`switch`用于 switch 语句。`symbol`表示符号类型。`this`引用当前类或对象的实例。`throw`抛出异常。`try`用于异常处理语句 `try...catch`。`true`布尔值 `true`。`type`用于定义类型别名。`typeof`获取变量或表达式的类型。`undefined`表示未定义的值。`unique`用于 `symbol` 类型的唯一标识符。`var`用于声明变量（已不推荐使用）。`void`表示没有返回值的类型。`while`用于 `while` 循环。`with`用于创建一个作用域，在该作用域内可以省略对象的引用（不推荐使用）。`yield`用于生成器函数中，暂停和恢复生成器的执行。

#### 空白和换行

TypeScript 会忽略程序中出现的空格、制表符和换行符。

空格、制表符通常用来缩进代码，使代码易于阅读和理解。

#### TypeScript 区分大小写

TypeScript 区分大写和小写字符。

#### 分号是可选的

每行指令都是一段语句，你可以使用分号或不使用， 分号在 TypeScript 中是可选的，建议使用。

以下代码都是合法的：

```
console.log("Runoob")
console.log("Google");
```

如果语句写在同一行则一定需要使用分号来分隔，否则会报错，如：

```
console.log("Runoob");console.log("Google");
```

#### TypeScript 注释

注释是一个良好的习惯，虽然很多程序员讨厌注释，但还是建议你在每段代码写上文字说明。

注释可以提高程序的可读性。

注释可以包含有关程序一些信息，如代码的作者，有关函数的说明等。

编译器会忽略注释。

#### TypeScript 支持两种类型的注释

- 单行注释 ( // ) − 在 // 后面的文字都是注释内容。
- 多行注释 (/* */) − 这种注释可以跨越多行。

注释实例：

```
// 这是一个单行注释

/*
这是一个多行注释
这是一个多行注释
这是一个多行注释
*/
```

### TypeScript 与面向对象

面向对象是一种对现实世界理解和抽象的方法。

TypeScript 是一种面向对象的编程语言。

面向对象主要有两个概念：对象和类。

- 对象：对象是类的一个实例（对象不是找个女朋友），有状态和行为。例如，一条狗是一个对象，它的状态有：颜色、名字、品种；行为有：摇尾巴、叫、吃等。
- 类：类是一个模板，它描述一类对象的行为和状态。
- 方法：方法是类的操作的实现步骤。

下图中 girl、boy 为类，而具体的每个人为该类的对象：

TypeScript 面向对象编程实例：

class Site { name():void { console.log("Runoob") } } var obj = new Site(); obj.name();

以上实例定义了一个类 Site，该类有一个方法 name()，该方法在终端上输出字符串 Runoob。

new 关键字创建类的对象，该对象调用方法 name()。

编译后生成的 JavaScript 代码如下：

var Site = /** @class */ (function () { function Site() { } Site.prototype.name = function () { console.log("Runoob"); }; return Site; }()); var obj = new Site(); obj.name();

执行以上 JavaScript 代码，输出结果如下:

```
Runoob
```

---

## TypeScript 基本结构

Source: https://www.runoob.com/typescript/typescript-basic-structure.html

## TypeScript 基本结构

TypeScript 程序的基本结构可以分为几个部分，每个部分都有特定的作用。

以下是 TypeScript 程序的常见组成部分：

- 声明部分：包括类型声明、接口声明等。
- 变量声明：包括 `let`, `const` 和 `var` 的使用。
- 函数声明：包括普通函数和箭头函数。
- 类声明：用于定义类及其成员。
- 接口与类型别名：描述类型的结构。
- 模块化：通过 `import` 和 `export` 组织代码。
- 类型断言：强制类型转换。
- 泛型：使代码具备更多的复用性。
- 注释：增加代码的可读性。
- 类型推断：自动推断类型。
- 类型守卫：缩小类型范围。
- 异步编程：支持 `async/await`。
- 错误处理：通过 `try/catch` 进行错误捕捉。

以上几个部分共同构成了 TypeScript 程序的基本结构，并为开发者提供了强大的类型检查、代码结构和可维护性。

### 1. 声明部分（Declarations）

类型声明：TypeScript 是一种静态类型的语言，可以通过类型声明来定义变量、函数、类等的类型。类型声明可以帮助代码更具可维护性和可读性。

### 实例

let name: string = "Alice";
let age: number = 30;

接口声明：用于定义对象的结构，包括对象的属性和方法。

### 实例

interface Person {
name: string;
age: number;
}

### 2. 变量声明（Variable Declarations）

在 TypeScript 中，可以使用 let, const, 和 var 来声明变量。推荐使用 let 和 const，var 用法不再推荐。

### 实例

let age: number = 25;
const pi: number = 3.14;

### 3. 函数声明（Function Declarations）

函数声明：TypeScript 允许声明带有类型注解的函数，包括参数类型和返回值类型。

### 实例

function greet(name: string): string {
return "Hello, " + name;
}

箭头函数：TypeScript 同样支持 ES6 的箭头函数，使用简洁的语法来声明函数。

### 实例

const greet = (name: string): string => "Hello, " + name;

### 4. 类声明（Class Declarations）

TypeScript 提供对面向对象编程的支持，允许定义类和类的方法、属性。

### 实例

class Person {
name: string;
age: number;

constructor(name: string, age: number) {
this.name = name;
this.age = age;
}

greet() {
return `Hello, my name is ${this.name}`;
}
}

### 5. 接口与类型别名（Interfaces & Type Aliases）

接口（Interface）：用于描述对象的形状，接口可以继承和扩展。

### 实例

interface Animal {
name: string;
sound: string;
makeSound(): void;
}

类型别名（Type Alias）：允许为对象类型、联合类型、交叉类型等定义别名。

### 实例

type ID = string | number;

### 6. 模块和导入导出（Modules & Imports/Exports）

TypeScript 支持模块化编程，可以使用 import 和 export 来组织代码。

导出：

### 实例

export class Person {
constructor(public name: string) {}
}

导入：

### 实例

import { Person } from './person';

### 7. 类型断言（Type Assertions）

在某些情况下，TypeScript 无法推断出一个变量的准确类型，开发者可以使用类型断言来强制指定类型。

### 实例

let value: any = "hello";
let strLength: number = (value as string).length;

### 8. 泛型（Generics）

泛型允许在定义函数、接口或类时不指定具体类型，而是使用占位符，让用户在使用时传入具体类型。泛型能够增加代码的复用性和类型安全性。

### 实例

function identity<T>(arg: T): T {
return arg;
}

### 9. 注释（Comments）

注释在 TypeScript 程序中用于解释代码的作用、思路等，增加代码的可读性。

单行注释：

### 实例

// 这是一个单行注释

多行注释：

### 实例

/* 这是一个
多行注释 */

### 10. 类型推断（Type Inference）

TypeScript 在某些情况下会自动推断变量的类型。例如，在声明变量并赋值时，TypeScript 会推断出该变量的类型。

### 实例

let num = 10; // TypeScript 推断 num 为 number 类型

### 11. 类型守卫（Type Guards）

TypeScript 提供了类型守卫（如 typeof 和 instanceof），用于在运行时缩小变量的类型范围。

### 实例

function isString(value: any): value is string {
return typeof value === 'string';
}

### 12. 异步编程（Asynchronous Programming）

TypeScript 完全支持异步编程，可以使用 async/await 语法来处理异步操作。

### 实例

async function fetchData(): Promise<string> {
const response = await fetch("https://example.com");
const data = await response.text();
return data;
}

### 13. 错误处理（Error Handling）

TypeScript 允许使用 try/catch 块进行错误处理，还可以使用类型来描述错误的类型。

### 实例

try {
throw new Error("Something went wrong");
} catch (error) {
if (error instanceof Error) {
console.error(error.message);
}
}

---

## TypeScript vs JavaScript 对比

Source: https://www.runoob.com/typescript/typescript-vs-javascript.html

## TypeScript vs JavaScript 对比

TypeScript 和 JavaScript 是前端开发中两个最重要的语言，理解它们之间的区别对于现代 Web 开发至关重要。

JavaScript 是 Web 的原生脚本语言，而 TypeScript 是 JavaScript 的超集，添加了类型系统和其他高级特性。

### 核心区别

特性 JavaScript TypeScript 类型系统 动态类型 静态类型（可选） 编译 解释执行 编译为 JavaScript 开发时类型检查 无 有 IDE 支持 基础 强大的智能提示 代码重构 困难 安全简单 TypeScript 是 JavaScript 的超集 JavaScript 动态类型 解释执行 无需编译 TypeScript 静态类型（可选） 编译为 JS 类型检查 智能提示 包含关系 所有 JS 代码 都是有效 TS

### 类型系统对比

JavaScript 是动态类型语言，变量类型在运行时确定：

### JavaScript 动态类型

// JavaScript - 变量类型可以随时改变
var message = "Hello";
message = 123; // 合法，不会报错
message = true; // 仍然合法

// 运行前不知道变量类型
function greet(name) {
return "Hello, " + name;
}

TypeScript 是静态类型语言，在编译时检查类型：

### TypeScript 静态类型

// TypeScript - 声明时指定类型
var message: string = "Hello";
// message = 123; // 编译错误：Type 'number' is not assignable to type 'string'

// 函数参数和返回值类型明确
function greet(name: string): string {
return "Hello, " + name;
}

// 类型错误会在编译时发现
greet(123); // 编译错误：Argument of type 'number' is not assignable to parameter of type 'string'

### 编译过程

TypeScript 需要编译为 JavaScript 才能在浏览器中运行：

### TypeScript 源码

// app.ts - TypeScript 源码
interface User {
name: string;
age: number;
}

function createUser(name: string, age: number): User {
return { name, age };
}

var user = createUser("Alice", 25);
console.log(user);

编译后的 JavaScript：

### 编译后的 JavaScript

// app.js - 编译后的 JavaScript
function createUser(name, age) {
return { name: name, age: age };
}

var user = createUser("Alice", 25);
console.log(user);

### 类型推断

TypeScript 具有强大的类型推断能力，即使不显式声明类型：

### 类型推断

// TypeScript 自动推断类型
var num = 10; // 推断为 number
var str = "hello"; // 推断为 string
var isActive = true; // 推断为 boolean

// 函数返回值类型也会推断
function add(a, b) {
return a + b; // 推断返回值为 number
}

var result = add(1, 2); // result 类型为 number

### 开发体验对比

#### 智能提示

TypeScript 为 IDE 提供丰富的类型信息，实现智能提示：

```

// 在 TypeScript 中，IDE 知道 user 对象有 name 和 age 属性
user. // 自动提示 .name 和 .age

// JavaScript 中 IDE 无法确定类型，提示有限
user. // 可能没有有用的提示

```

#### 重构支持

TypeScript 使得代码重构更安全：

- 重命名变量/函数时，所有引用自动更新
- 修改函数签名时，调用处会显示错误
- 提取代码时，类型自动保持

### 选择 TypeScript 的理由

为什么选择 TypeScript：

- 编译时错误检测，提前发现 bug
- 更好的代码可读性和可维护性
- 强大的 IDE 智能提示
- 安全的重构
- 现代前端框架的标配（React、Vue、Angular）

TypeScript 是 JavaScript 的超集，它在保持 JavaScript 灵活性的同时，增加了类型系统来提高代码质量。对于大型项目，TypeScript 的优势更加明显。

---

## TypeScript tsconfig.json 配置

Source: https://www.runoob.com/typescript/ts-tsconfig.html

## TypeScript tsconfig.json 配置

tsconfig.json 是 TypeScript 项目的配置文件，用于指定编译选项和项目设置。

### 基本配置

最基础的 tsconfig.json 文件。

### 实例

{
"compilerOptions": {
"target": "ES2020",
"module": "commonjs",
"strict": true,
"outDir": "./dist"
},
"include": ["src/**/*"],
"exclude": ["node_modules", "dist"]
}

配置说明：

- target：编译目标 JavaScript 版本
- module：使用的模块系统
- strict：启用所有严格类型检查
- outDir：输出目录

### 编译目标版本

使用 target 指定编译到的 JavaScript 版本。

### 实例

// tsconfig.json
{
"compilerOptions": {
// ES3, ES5, ES6/ES2015, ES2020, ESNext
"target": "ES2020"
}
}

不同目标的输出差异：

```

// target: ES5 - 使用 var
var greeting = "Hello";

// target: ES2020 - 使用 let/const
let greeting = "Hello";

```

### 模块系统

配置模块化方案。

### 实例

{
"compilerOptions": {
"module": "commonjs",
// 可选: none, commonjs, amd, umd, es6, es2020, esnext
}
}

### 严格模式

strict 选项启用所有严格类型检查。

### 实例

{
"compilerOptions": {
"strict": true,
// 等同于开启以下所有选项：
// "strictNullChecks": true,
// "noImplicitAny": true,
// "strictFunctionTypes": true,
// 等等
}
}

建议：始终启用 strict: true，这是最佳实践。

### 路径别名

配置路径别名简化导入。

### 实例

{
"compilerOptions": {
"baseUrl": ".",
"paths": {
"@/*": ["src/*"],
"@components/*": ["src/components/*"]
}
}
}

使用方式：

```

import Button from "@components/Button";
import { Header } from "@/components";

```

### 文件包含与排除

控制哪些文件包含在编译中。

### 实例

{
"include": ["src/**/*"],
"exclude": [
"node_modules",
"dist",
"**/*.test.ts",
"**/*.spec.ts"
]
}

### 常用编译选项一览

选项 说明 target 编译目标版本 module 模块系统 strict 严格模式 outDir 输出目录 rootDir 源码根目录 esModuleInterop 允许 ES 模块互操作 skipLibCheck 跳过库检查 declaration 生成 .d.ts 声明文件

### 总结

- tsconfig.json：TypeScript 项目配置文件
- compilerOptions：编译选项核心配置
- include/exclude：控制文件范围
- 路径别名：简化导入路径
- strict: true：始终启用严格模式

---

## TypeScript 编译选项

Source: https://www.runoob.com/typescript/ts-compiler-options.html

## TypeScript 编译选项

TypeScript 编译器（tsc）有众多编译选项，本教程详细介绍常用的编译选项及其作用。 TypeScript 编译过程与编译选项 TypeScript 源码 .ts / .tsx 包含类型注解 TypeScript 编译器 (tsc) 编译选项作用： • 类型检查 (strict) • 输出格式 (module) • 目标版本 (target) • 声明文件 (declaration) ...更多选项 编译输出 .js - JavaScript .d.ts - 类型声明 .map - Source Map 编译选项分类 类型检查 strict noImplicitAny strictNullChecks 输出控制 outDir declaration sourceMap 模块系统 module moduleResolution esModuleInterop 语言特性 target lib jsx

### 输出控制选项

### 实例

{
"compilerOptions": {
// 输出目录
"outDir": "./dist",

// 源码根目录
"rootDir": "./src",

// 生成 .d.ts 声明文件
"declaration": true,

// 声明文件输出目录
"declarationDir": "./types",

// 生成 source map
"sourceMap": true,

// 生成 .js.map 文件
"mapRoot": "./map"
}
}

### 类型检查选项

### 实例

{
"compilerOptions": {
// 严格模式（推荐始终开启）
"strict": true,

// 检查 null 和 undefined
"strictNullChecks": true,

// 检查 this 参数
"noImplicitThis": true,

// 严格函数类型
"strictFunctionTypes": true,

// 严格属性初始化
"strictPropertyInitialization": true,

// 不允许隐式 any
"noImplicitAny": true,

// 不允许返回 void
"noImplicitReturns": true,

// 开启所有严格检查
"strict": true
}
}

### 模块选项

### 实例

{
"compilerOptions": {
// 模块系统
"module": "commonjs",

// 模块解析策略
"moduleResolution": "node",

// 解析基础路径
"baseUrl": ".",

// 路径别名
"paths": {
"@/*": ["src/*"]
},

// ES 模块互操作
"esModuleInterop": true,

// 允许默认导入
"allowSyntheticDefaultImports": true,

// 隔离模块
"isolatedModules": true
}
}

### ES 特性选项

### 实例

{
"compilerOptions": {
// 编译目标
"target": "ES2020",

// 启用的库
"lib": ["ES2020", "DOM"],

// 允许未使用的局部变量
"noUnusedLocals": true,

// 允许未使用的参数
"noUnusedParameters": true,

// 代码降级
"downlevelIteration": true
}
}

### 实验性选项

### 实例

{
"compilerOptions": {
// 启用装饰器
"experimentalDecorators": true,

// 启用装饰器元数据
"emitDecoratorMetadata": true,

// 启用异步迭代器
"emitDecoratorMetadata": true,

// 跳过库检查
"skipLibCheck": true
}
}

### 常用编译选项组合

### 实例

// Node.js 项目推荐配置
{
"compilerOptions": {
"target": "ES2020",
"module": "commonjs",
"lib": ["ES2020"],
"outDir": "./dist",
"rootDir": "./src",
"strict": true,
"esModuleInterop": true,
"skipLibCheck": true,
"forceConsistentCasingInFileNames": true,
"moduleResolution": "node",
"declaration": true
}
}

### 编译选项表

类别 常用选项 输出 outDir, rootDir, declaration, sourceMap 类型检查 strict, strictNullChecks, noImplicitAny 模块 module, moduleResolution, paths, esModuleInterop ES 特性 target, lib, downlevelIteration 实验性 experimentalDecorators, emitDecoratorMetadata

### 总结

- 严格模式：始终启用 strict: true
- 目标版本：根据环境选择合适的 target
- 模块系统：根据部署环境选择 module
- 类型声明：库开发时启用 declaration

---

## TypeScript 基础类型

Source: https://www.runoob.com/typescript/ts-type.html

## TypeScript 基础类型

基础类型可以开发者更准确地描述数据的结构和意图。

TypeScript 包含的数据类型如下表:

类型描述示例`string`表示文本数据`let name: string = "Alice";``number`表示数字，包括整数和浮点数`let age: number = 30;``boolean`表示布尔值 `true` 或 `false``let isDone: boolean = true;``array`表示相同类型的元素数组`let list: number[] = [1, 2, 3];``tuple`表示已知类型和长度的数组`let person: [string, number] = ["Alice", 30];``enum`定义一组命名常量`enum Color { Red, Green, Blue };``any`任意类型，不进行类型检查`let value: any = 42;``void`无返回值（常用于函数）`function log(): void {}``null`表示空值`let empty: null = null;``undefined`表示未定义`let undef: undefined = undefined;``never`表示不会有返回值`function error(): never { throw new Error("error"); }``object`表示非原始类型`let obj: object = { name: "Alice" };``union`联合类型，表示可以是多种类型之一`let id: string`unknown`不确定类型，需类型检查后再使用`let value: unknown = "Hello";`

注意：TypeScript 和 JavaScript 没有整数类型。

#### 1、string 字符串

表示文本数据，只能存储字符串，通常用于描述文字信息。

```

let message: string = "Hello, TypeScript!";
```

模板字符串：TypeScript 支持 模板字符串，用反引号 `（记住不是单引号 '）来定义，允许在字符串中插入变量或表达式，非常适合多行文本和拼接变量。

```
let name: string = "Alice";
let greeting: string = `Hello, ${name}! Welcome to TypeScript.`;
console.log(greeting); // 输出：Hello, Alice! Welcome to TypeScript.
```

#### 2、number 数字

TypeScript 使用 number 表示所有数字，包括整数和浮点数。

```

let age: number = 25;
let temperature: number = 36.5;
```

#### 3、boolean 布尔值

表示逻辑值 true 或 false，用于条件判断。

```

let isCompleted: boolean = false;
```

#### 4、array 数组

可以表示一组相同类型的元素。可以使用 type[] 或 Array<type> 两种方式表示。

```

let numbers: number[] = [1, 2, 3];
let names: Array<string> = ["Alice", "Bob"];
```

#### 5、tuple 元组

表示已知数量和类型的数组。每个元素可以是不同的类型，适合表示固定结构的数据。

```

let person: [string, number] = ["Alice", 25];
```

#### 6、enum 枚举

用来定义一组命名常量。默认情况下枚举的值从 0 开始递增。

```
enum Color {
Red,
Green,
Blue,
}
let favoriteColor: Color = Color.Green;
```

#### 7、any 类型

以表示任何类型。适合不确定数据类型的情况，但使用时需谨慎，因为 any 会绕过类型检查。

```
let randomValue: any = 42;
randomValue = "hello";
```

任意值是 TypeScript 针对编程时类型不明确的变量使用的一种数据类型，它常用于以下三种情况。

1、变量的值会动态改变时，比如来自用户的输入，任意值类型可以让这些变量跳过编译阶段的类型检查，示例代码如下：

```

let x: any = 1; // 数字类型
x = 'I am who I am'; // 字符串类型
x = false; // 布尔类型
```

改写现有代码时，任意值允许在编译时可选择地包含或移除类型检查，示例代码如下：

```
let x: any = 4;
x.ifItExists(); // 正确，ifItExists方法在运行时可能存在，但这里并不会检查
x.toFixed(); // 正确
```

定义存储各种类型数据的数组时，示例代码如下：

```
let arrayList: any[] = [1, false, 'fine'];
arrayList[1] = 100;
```

#### 8、void 空类型

用于没有返回值的函数。声明变量时，类型 void 意味着只能赋值 null 或 undefined。

```

function logMessage(message: string): void {
console.log(message);
}
```

#### 9、null 和 undefined

null 和 undefined分别表示"空值"和"未定义"。在默认情况下，它们是所有类型的子类型，但可以通过设置 strictNullChecks 严格检查。

```
let empty: null = null;
let notAssigned: undefined = undefined;
```

null

在 JavaScript 中 null 表示 "什么都没有"。

null是一个只有一个值的特殊类型。表示一个空对象引用。

用 typeof 检测 null 返回是 object。

undefined

在 JavaScript 中, undefined 是一个没有设置值的变量。

typeof 一个没有值的变量会返回 undefined。

Null 和 Undefined 是其他任何类型（包括 void）的子类型，可以赋值给其它类型，如数字类型，此时，赋值后的类型会变成 null 或 undefined。而在TypeScript中启用严格的空校验（--strictNullChecks）特性，就可以使得null 和 undefined 只能被赋值给 void 或本身对应的类型，示例代码如下：

```
// 启用 --strictNullChecks
let x: number;
x = 1; // 编译正确
x = undefined; // 编译错误
x = null; // 编译错误
```

上面的例子中变量 x 只能是数字类型。如果一个类型可能出现 null 或 undefined， 可以用 | 来支持多种类型，示例代码如下：

```
// 启用 --strictNullChecks
let x: number | null | undefined;
x = 1; // 编译正确
x = undefined; // 编译正确
x = null; // 编译正确
```

更多内容可以查看：JavaScript typeof, null, 和 undefined

#### 10、never 类型

表示不会有返回值，通常用于抛出错误或进入无限循环的函数，表示该函数永远不会正常结束。

```
function throwError(message: string): never {
throw new Error(message);
}
```

never 是其它类型（包括 null 和 undefined）的子类型，代表从不会出现的值。这意味着声明为 never 类型的变量只能被 never 类型所赋值，在函数中它通常表现为抛出异常或无法执行到终止点（例如无限循环），示例代码如下：

```
let x: never;
let y: number;

// 编译错误，数字类型不能转为 never 类型
x = 123;

// 运行正确，never 类型可以赋值给 never类型
x = (()=>{ throw new Error('exception')})();

// 运行正确，never 类型可以赋值给 数字类型
y = (()=>{ throw new Error('exception')})();

// 返回值为 never 的函数可以是抛出异常的情况
function error(message: string): never {
throw new Error(message);
}

// 返回值为 never 的函数可以是无法被执行到的终止点的情况
function loop(): never {
while (true) {}
}
```

#### 11、object 对象类型

表示非原始类型的值，适用于复杂的对象结构。

```

let person: object = { name: "Alice", age: 30 };
```

#### 12、联合类型 (Union)

表示一个变量可以是多种类型之一。通过 | 符号实现。

```

let id: string | number;
id = "123";
id = 456;
```

#### 13、unknown 不确定的类型

与 any 类似，但更严格。必须经过类型检查后才能赋值给其他类型变量。

```

let value: unknown = "Hello";
if (typeof value === "string") {
let message: string = value;
}
```

#### 14、类型断言 (Type Assertions)

类型断言可以让开发者明确告诉编译器变量的类型，常用于无法推断的情况。可以使用 as 或尖括号语法。

```

let someValue: any = "this is a string";
let strLength: number = (someValue as string).length;
```

#### 15、字面量类型

字面量类型可以让变量只能拥有特定的值，用于结合联合类型定义变量的特定状态。

```

let direction: "up" | "down" | "left" | "right";
direction = "up";
```

通过这些类型，TypeScript 提供了更强的类型安全性和代码检查能力，使开发者能够更清晰、准确地表达数据和意图，减少运行时错误。

### 实例

以下实例展示了 TypeScript 中主要基础类型的定义和使用，模拟一个用户对象和相关的操作函数：

### 实例

// 定义枚举类型，用于表示用户的角色
enum Role {
Admin,
User,
Guest,
}

// 使用 interface 定义用户的结构
interface User {
id: number; // number 类型，用于唯一标识用户
username: string; // string 类型，表示用户名
isActive: boolean; // boolean 类型，表示用户是否激活
role: Role; // enum 类型，用于表示用户角色
hobbies: string[]; // array 类型，存储用户的兴趣爱好
contactInfo: [string, number]; // tuple 类型，包含电话号码的元组，格式为：[区域码, 电话号码]
}

// 创建用户对象，符合 User 接口的结构
const user: User = {
id: 1,
username: "Alice",
isActive: true,
role: Role.User,
hobbies: ["Reading", "Gaming"],
contactInfo: ["+1", 123456789],
};

// 定义一个返回字符串的函数来获取用户信息
function getUserInfo(user: User): string {
return `User ${user.username} is ${user.isActive ? "active" : "inactive"} with role ${Role[user.role]}`;
}

// 使用 void 类型定义一个函数，专门打印用户信息
function printUserInfo(user: User): void {
console.log(getUserInfo(user));
}

// 定义一个 union 类型的函数参数，接受用户 ID（number）或用户名（string）
function findUser(query: number | string): User | undefined {
// 使用 typeof 来判断 query 的类型
if (typeof query === "number") {
// 如果是数字，则根据 ID 查找用户
return query === user.id ? user : undefined;
} else if (typeof query === "string") {
// 如果是字符串，则根据用户名查找用户
return query === user.username ? user : undefined;
}
return undefined;
}

// 定义一个 never 类型的函数，用于处理程序的异常情况
function throwError(message: string): never {
throw new Error(message);
}

// 使用 any 类型处理未知类型的数据
let unknownData: any = "This is a string";
unknownData = 42; // 重新赋值为数字，类型为 any

// 使用 unknown 类型处理不确定的数据，更加安全
let someData: unknown = "Possible data";
if (typeof someData === "string") {
console.log(`Length of data: ${(someData as string).length}`);
}

// 调用各个函数并测试
printUserInfo(user); // 打印用户信息
console.log(findUser(1)); // 根据 ID 查找用户
console.log(findUser("Alice")); // 根据用户名查找用户

// 使用 null 和 undefined 类型的变量
let emptyValue: null = null;
let uninitializedValue: undefined = undefined;

说明：

- 枚举类型 (`Role`)：用于定义用户角色的命名常量 `Admin`、`User` 和 `Guest`。
- 接口 (`User`)：定义了 `User` 对象的结构，包括 `id`、`username`、`isActive` 等属性，展示了 `string`、`number`、`boolean`、`array` 和 `tuple` 类型的使用。
- 函数 `getUserInfo`：返回用户的描述信息，使用了字符串插值并结合 `enum`。
- 函数 `printUserInfo`：类型为 `void`，因为它仅输出信息，不返回任何值。
- 联合类型 (`number | string`)：用于 `findUser` 函数的参数，可以接受数字或字符串。
- `never` 类型：`throwError` 函数抛出错误，不会正常返回，因此使用 `never` 类型。
- `any` 类型：展示如何声明一个可以接收任何类型的变量。
- `unknown` 类型：类似 `any`，但更加安全，示例中在类型断言之前进行类型检查。
- `null` 和 `undefined`：演示空值的使用。

---

## TypeScript 变量声明

Source: https://www.runoob.com/typescript/ts-variables.html

## TypeScript 变量声明

变量是一种使用方便的占位符，用于引用计算机内存地址。

我们可以把变量看做存储数据的容器。

TypeScript 变量的命名规则：

- 变量名称可以包含数字和字母。
- 除了下划线 _ 和美元 $ 符号外，不能包含其他特殊字符，包括空格。
- 变量名不能以数字开头。

变量使用前必须先声明，我们可以使用 var 来声明变量。

我们可以使用以下四种方式来声明变量：

声明变量的类型及初始值：

```
var [变量名] : [类型] = 值;
```

例如：

```
var uname:string = "Runoob";
```

声明变量的类型，但没有初始值，变量值会设置为 undefined：

```
var [变量名] : [类型];
```

例如：

```
var uname:string;
```

声明变量并初始值，但不设置类型，该变量可以是任意类型：

```
var [变量名] = 值;
```

例如：

```
var uname = "Runoob";
```

声明变量没有设置类型和初始值，类型可以是任意类型，默认初始值为 undefined：

```
var [变量名];
```

例如：

```
var uname;
```

#### 实例

var uname:string = "Runoob";
var score1:number = 50;
var score2:number = 42.50
var sum = score1 + score2
console.log("名字: "+uname)
console.log("第一个科目成绩: "+score1)
console.log("第二个科目成绩: "+score2)
console.log("总成绩: "+sum)

注意：变量不要使用 name 否则会与 DOM 中的全局 window 对象下的 name 属性出现了重名。

使用 tsc 命令编译以上代码，得到如下 JavaScript 代码：

var uname = "Runoob"; var score1 = 50; var score2 = 42.50; var sum = score1 + score2; console.log("名字: " + uname); console.log("第一个科目成绩: " + score1); console.log("第二个科目成绩: " + score2); console.log("总成绩: " + sum);

执行该 JavaScript 代码输出结果为：

```
名字: Runoob
第一个科目成绩: 50
第二个科目成绩: 42.5
总成绩: 92.5
```

TypeScript 遵循强类型，如果将不同的类型赋值给变量会编译错误，如下实例：

```
var num:number = "hello" // 这个代码会编译错误
```

### 类型断言（Type Assertion）

类型断言可以用来手动指定一个值的类型，即允许变量从一种类型更改为另一种类型。

语法格式：

```
<类型>值
```

或:

```
值 as 类型
```

#### 实例

var str = '1' var str2:number = <number> <any> str //str、str2 是 string 类型 console.log(str2)

#### TypeScript 是怎么确定单个断言是否足够

当 S 类型是 T 类型的子集，或者 T 类型是 S 类型的子集时，S 能被成功断言成 T。这是为了在进行类型断言时提供额外的安全性，完全毫无根据的断言是危险的，如果你想这么做，你可以使用 any。

它之所以不被称为类型转换，是因为转换通常意味着某种运行时的支持。但是，类型断言纯粹是一个编译时语法，同时，它也是一种为编译器提供关于如何分析代码的方法。

编译后，以上代码会生成如下 JavaScript 代码：

var str = '1'; var str2 = str; //str、str2 是 string 类型 console.log(str2);

执行输出结果为：

```
1
```

### 类型推断

当类型没有给出时，TypeScript 编译器利用类型推断来推断类型。

如果由于缺乏声明而不能推断出类型，那么它的类型被视作默认的动态 any 类型。

var num = 2; // 类型推断为 number console.log("num 变量的值为 "+num); num = "12"; // 编译错误 console.log(num);

- 第一行代码声明了变量 num 并=设置初始值为 2。 注意变量声明没有指定类型。因此，程序使用类型推断来确定变量的数据类型，第一次赋值为 2，num 设置为 number 类型。
- 第三行代码，当我们再次为变量设置字符串类型的值时，这时编译会错误。因为变量已经设置为了 number 类型。

```
error TS2322: Type '"12"' is not assignable to type 'number'.
```

### 变量作用域

变量作用域指定了变量定义的位置。

程序中变量的可用性由变量作用域决定。

TypeScript 有以下几种作用域：

- 全局作用域 − 全局变量定义在程序结构的外部，它可以在你代码的任何位置使用。
- 类作用域 − 这个变量也可以称为 字段。类变量声明在一个类里头，但在类的方法外面。 该变量可以通过类的对象来访问。类变量也可以是静态的，静态的变量可以通过类名直接访问。
- 局部作用域 − 局部变量，局部变量只能在声明它的一个代码块（如：方法）中使用。

以下实例说明了三种作用域的使用：

var global_num = 12 // 全局变量 class Numbers { num_val = 13; // 实例变量 static sval = 10; // 静态变量 storeNum():void { var local_num = 14; // 局部变量 } } console.log("全局变量为: "+global_num) console.log(Numbers.sval) // 静态变量 var obj = new Numbers(); console.log("实例变量: "+obj.num_val)

以上代码使用 tsc 命令编译为 JavaScript 代码为：

var global_num = 12; // 全局变量 var Numbers = /** @class */ (function () { function Numbers() { this.num_val = 13; // 实例变量 } Numbers.prototype.storeNum = function () { var local_num = 14; // 局部变量 }; Numbers.sval = 10; // 静态变量 return Numbers; }()); console.log("全局变量为: " + global_num); console.log(Numbers.sval); // 静态变量 var obj = new Numbers(); console.log("实例变量: " + obj.num_val);

执行以上 JavaScript 代码，输出结果为：

```
全局变量为: 12
10
实例变量: 13
```

如果我们在方法外部调用局部变量 local_num，会报错：

```
error TS2322: Could not find symbol 'local_num'.
```

---

## TypeScript 运算符

Source: https://www.runoob.com/typescript/ts-operators.html

## TypeScript 运算符

运算符用于执行程序代码运算，会针对一个以上操作数项目来进行运算。

考虑以下计算：

```
7 + 5 = 12
```

以上实例中 7、5 和 12 是操作数。

运算符 + 用于加值。

运算符 = 用于赋值。

TypeScript 主要包含以下几种运算：

- 算术运算符
- 逻辑运算符
- 关系运算符
- 按位运算符
- 赋值运算符
- 三元/条件运算符
- 字符串运算符
- 类型运算符

### 算术运算符

假定 y=5，下面的表格解释了这些算术运算符的操作：

运算符 描述 例子 x 运算结果 y 运算结果 + 加法 x=y+2 7 5 - 减法 x=y-2 3 5 * 乘法 x=y*2 10 5 / 除法 x=y/2 2.5 5 % 取模（余数） x=y%2 1 5 ++ 自增 x=++y 6 6 x=y++ 5 6 -- 自减 x=--y 4 4 x=y-- 5 4

#### 实例

var num1:number = 10
var num2:number = 2
var res:number = 0

res = num1 + num2
console.log("加: "+res);

res = num1 - num2;
console.log("减: "+res)

res = num1*num2
console.log("乘: "+res)

res = num1/num2
console.log("除: "+res)

res = num1%num2
console.log("余数: "+res)

num1++
console.log("num1 自增运算: "+num1)

num2--
console.log("num2 自减运算: "+num2)

使用 tsc 命令编译以上代码得到如下 JavaScript 代码：

var num1 = 10; var num2 = 2; var res = 0; res = num1 + num2; console.log("加: " + res); res = num1 - num2; console.log("减: " + res); res = num1 * num2; console.log("乘: " + res); res = num1 / num2; console.log("除: " + res); res = num1 % num2; console.log("余数: " + res); num1++; console.log("num1 自增运算: " + num1); num2--; console.log("num2 自减运算: " + num2);

执行以上 JavaScript 代码，输出结果为：

```
加: 12
减: 8
乘: 20
除: 5
余数: 0
num1 自增运算: 11
num2 自减运算: 1
```

### 关系运算符

关系运算符用于计算结果是否为 true 或者 false。

x=5，下面的表格解释了关系运算符的操作：

运算符 描述 比较 返回值 == 等于 x==8 false x==5 true != 不等于 x!=8 true > 大于 x>8 false < 小于 x<8 true >= 大于或等于 x>=8 false <= 小于或等于 x<=8 true

#### 实例

var num1:number = 5; var num2:number = 9; console.log("num1 的值为: "+num1); console.log("num2 的值为:"+num2); var res = num1>num2 console.log("num1 大于n num2: "+res) res = num1<num2 console.log("num1 小于 num2: "+res) res = num1>=num2 console.log("num1 大于或等于 num2: "+res) res = num1<=num2 console.log("num1 小于或等于 num2: "+res) res = num1==num2 console.log("num1 等于 num2: "+res) res = num1!=num2 console.log("num1 不等于 num2: "+res)

使用 tsc 命令编译以上代码得到如下 JavaScript 代码：

var num1 = 5; var num2 = 9; console.log("num1 的值为: " + num1); console.log("num2 的值为:" + num2); var res = num1 > num2; console.log("num1 大于n num2: " + res); res = num1 < num2; console.log("num1 小于 num2: " + res); res = num1 >= num2; console.log("num1 大于或等于 num2: " + res); res = num1 <= num2; console.log("num1 小于或等于 num2: " + res); res = num1 == num2; console.log("num1 等于 num2: " + res); res = num1 != num2; console.log("num1 不等于 num2: " + res);

执行以上 JavaScript 代码，输出结果为：

```

num1 的值为: 5
num2 的值为:9
num1 大于n num2: false
num1 小于 num2: true
num1 大于或等于 num2: false
num1 小于或等于 num2: true
num1 等于 num2: false
num1 不等于 num2: true

```

### 逻辑运算符

逻辑运算符用于测定变量或值之间的逻辑。

给定 x=6 以及 y=3，下表解释了逻辑运算符：

运算符 描述 例子 && and (x < 10 && y > 1) 为 true || or (x==5 || y==5) 为 false ! not !(x==y) 为 true

#### 实例

var avg:number = 20; var percentage:number = 90; console.log("avg 值为: "+avg+" ,percentage 值为: "+percentage); var res:boolean = ((avg>50)&&(percentage>80)); console.log("(avg>50)&&(percentage>80): ",res); var res:boolean = ((avg>50)||(percentage>80)); console.log("(avg>50)||(percentage>80): ",res); var res:boolean=!((avg>50)&&(percentage>80)); console.log("!((avg>50)&&(percentage>80)): ",res);

使用 tsc 命令编译以上代码得到如下 JavaScript 代码：

var avg = 20; var percentage = 90; console.log("avg 值为: " + avg + " ,percentage 值为: " + percentage); var res = ((avg > 50) && (percentage > 80)); console.log("(avg>50)&&(percentage>80): ", res); var res = ((avg > 50) || (percentage > 80)); console.log("(avg>50)||(percentage>80): ", res); var res = !((avg > 50) && (percentage > 80)); console.log("!((avg>50)&&(percentage>80)): ", res);

执行以上 JavaScript 代码，输出结果为：

```

avg 值为: 20 ,percentage 值为: 90
(avg>50)&&(percentage>80): false
(avg>50)||(percentage>80): true
!((avg>50)&&(percentage>80)): true

```

#### 短路运算符(&& 与 ||)

&& 与 || 运算符可用于组合表达式。 && 运算符只有在左右两个表达式都为 true 时才返回 true。

考虑以下实例：

```
var a = 10
var result = ( a<10 && a>5)
```

以上实例中 a < 10 与 a > 5 是使用了 && 运算符的组合表达式，第一个表达式返回了 false，由于 && 运算需要两个表达式都为 true，所以如果第一个为 false，就不再执行后面的判断(a > 5 跳过计算)，直接返回 false。

|| 运算符只要其中一个表达式为 true ，则该组合表达式就会返回 true。

考虑以下实例：

```
var a = 10
var result = ( a>5 || a<10)
```

以上实例中 a > 5 与 a < 10 是使用了 || 运算符的组合表达式，第一个表达式返回了 true，由于 || 组合运算只需要一个表达式为 true，所以如果第一个为 true，就不再执行后面的判断(a < 10 跳过计算)，直接返回 true。

### 位运算符

位操作是程序设计中对位模式按位或二进制数的一元和二元操作。

运算符 描述 例子 类似于 结果 十进制 & AND，按位与处理两个长度相同的二进制数，两个相应的二进位都为 1，该位的结果值才为 1，否则为 0。 x = 5 & 1 0101 & 0001 0001 1 | OR，按位或处理两个长度相同的二进制数，两个相应的二进位中只要有一个为 1，该位的结果值为 1。 x = 5 | 1 0101 | 0001 0101 5 ~ 取反，取反是一元运算符，对一个二进制数的每一位执行逻辑反操作。使数字 1 成为 0，0 成为 1。 x = ~ 5 ~0101 1010 -6 ^ 异或，按位异或运算，对等长二进制模式按位或二进制数的每一位执行逻辑异按位或操作。操作的结果是如果某位不同则该位为 1，否则该位为 0。 x = 5 ^ 1 0101 ^ 0001 0100 4 << 左移，把 << 左边的运算数的各二进位全部左移若干位，由 << 右边的数指定移动的位数，高位丢弃，低位补 0。 x = 5 << 1 0101 << 1 1010 10 >> 右移，把 >> 左边的运算数的各二进位全部右移若干位，>> 右边的数指定移动的位数。 x = 5 >> 1 0101 >> 1 0010 2 >>> 无符号右移，与有符号右移位类似，除了左边一律使用0 补位。 x = 2 >>> 1 0010 >>> 1 0001 1

#### 实例

var a:number = 2; // 二进制 10 var b:number = 3; // 二进制 11 var result; result = (a & b); console.log("(a & b) => ",result) result = (a | b); console.log("(a | b) => ",result) result = (a ^ b); console.log("(a ^ b) => ",result); result = (~b); console.log("(~b) => ",result); result = (a << b); console.log("(a << b) => ",result); result = (a >> b); console.log("(a >> b) => ",result); result = (a >>> 1); console.log("(a >>> 1) => ",result);

使用 tsc 命令编译以上代码得到如下 JavaScript 代码：

var a = 2; // 二进制 10 var b = 3; // 二进制 11 var result; result = (a & b); console.log("(a & b) => ", result); result = (a | b); console.log("(a | b) => ", result); result = (a ^ b); console.log("(a ^ b) => ", result); result = (~b); console.log("(~b) => ", result); result = (a << b); console.log("(a << b) => ", result); result = (a >> b); console.log("(a >> b) => ", result); result = (a >>> 1); console.log("(a >>> 1) => ", result);

执行以上 JavaScript 代码，输出结果为：

```

(a & b) => 2
(a | b) => 3
(a ^ b) => 1
(~b) => -4
(a << b) => 16
(a >> b) => 0
(a >>> 1) => 1

```

### 赋值运算符

赋值运算符用于给变量赋值。

给定 x=10 和 y=5，下面的表格解释了赋值运算符：

运算符 例子 实例 x 值 = (赋值) x = y x = y x = 5 += (先进行加运算后赋值) x += y x = x + y x = 15 -= (先进行减运算后赋值) x -= y x = x - y x = 5 *= (先进行乘运算后赋值) x *= y x = x * y x = 50 /= (先进行除运算后赋值) x /= y x = x / y x = 2

类似的逻辑运算符也可以与赋值运算符联合使用：<<=, >>=, >>>=, &=, |= 与 ^=。

#### 实例

var a: number = 12 var b:number = 10 a = b console.log("a = b: "+a) a += b console.log("a+=b: "+a) a -= b console.log("a-=b: "+a) a *= b console.log("a*=b: "+a) a /= b console.log("a/=b: "+a) a %= b console.log("a%=b: "+a)

使用 tsc 命令编译以上代码得到如下 JavaScript 代码：

var a = 12; var b = 10; a = b; console.log("a = b: " + a); a += b; console.log("a+=b: " + a); a -= b; console.log("a-=b: " + a); a *= b; console.log("a*=b: " + a); a /= b; console.log("a/=b: " + a); a %= b; console.log("a%=b: " + a);

执行以上 JavaScript 代码，输出结果为：

```

a = b: 10
a+=b: 20
a-=b: 10
a*=b: 100
a/=b: 10
a%=b: 0

```

### 三元运算符 (?)

三元运算有 3 个操作数，并且需要判断布尔表达式的值。该运算符的主要是决定哪个值应该赋值给变量。

```
Test ? expr1 : expr2
```

- Test − 指定的条件语句
- expr1 − 如果条件语句 Test 返回 true 则返回该值
- expr2 − 如果条件语句 Test 返回 false 则返回该值

让我们看下以下实例：

var num:number = -2 var result = num > 0 ? "大于 0" : "小于 0，或等于 0" console.log(result)

实例中用于判断变量是否大于 0。

使用 tsc 命令编译以上代码得到如下 JavaScript 代码：

var num = -2; var result = num > 0 ? "大于 0" : "小于 0，或等于 0"; console.log(result);

以上实例输出结果如下：

```
小于 0，或等于 0
```

### 类型运算符

#### typeof 运算符

typeof 是一元运算符，返回操作数的数据类型。

查看以下实例:

var num = 12 console.log(typeof num); //输出结果: number

使用 tsc 命令编译以上代码得到如下 JavaScript 代码：

var num = 12; console.log(typeof num); //输出结果: number

以上实例输出结果如下：

```
number
```

#### instanceof

instanceof 运算符用于判断对象是否为指定的类型，后面章节我们会具体介绍它。

### 其他运算符

#### 负号运算符(-)

更改操作数的符号，查看以下实例：

var x:number = 4 var y = -x; console.log("x 值为: ",x); // 输出结果 4 console.log("y 值为: ",y); // 输出结果 -4

使用 tsc 命令编译以上代码得到如下 JavaScript 代码：

var x = 4; var y = -x; console.log("x 值为: ", x); // 输出结果 4 console.log("y 值为: ", y); // 输出结果 -4

以上实例输出结果如下：

```
x 值为: 4
y 值为: -4
```

#### 字符串运算符: 连接运算符 (+)

+ 运算符可以拼接两个字符串，查看以下实例：

var msg:string = "RUNOOB"+".COM" console.log(msg)

使用 tsc 命令编译以上代码得到如下 JavaScript 代码：

var msg = "RUNOOB" + ".COM"; console.log(msg);

以上实例输出结果如下：

```
RUNOOB.COM
```

---

## TypeScript 条件语句

Source: https://www.runoob.com/typescript/ts-if-statement.html

## TypeScript 条件语句

条件语句用于基于不同的条件来执行不同的动作。

TypeScript 条件语句是通过一条或多条语句的执行结果（True 或 False）来决定执行的代码块。

可以通过下图来简单了解条件语句的执行过程:

### 条件语句

通常在写代码时，您总是需要为不同的决定来执行不同的动作。您可以在代码中使用条件语句来完成该任务。

在 TypeScript 中，我们可使用以下条件语句：

- if 语句 - 只有当指定条件为 true 时，使用该语句来执行代码
- if...else 语句 - 当条件为 true 时执行代码，当条件为 false 时执行其他代码
- if...else if....else 语句- 使用该语句来选择多个代码块之一来执行
- switch 语句 - 使用该语句来选择多个代码块之一来执行

### if 语句

TypeScript if 语句由一个布尔表达式后跟一个或多个语句组成。

#### 语法

语法格式如下所示：

```

if(boolean_expression){
# 在布尔表达式 boolean_expression 为 true 执行
}

```

如果布尔表达式 boolean_expression为 true，则 if 语句内的代码块将被执行。如果布尔表达式为 false，则 if 语句结束后的第一组代码（闭括号后）将被执行。

#### 流程图

#### 实例

var num:number = 5 if (num > 0) { console.log("数字是正数") }

编译以上代码得到如下 JavaScript 代码：

var num = 5; if (num > 0) { console.log("数字是正数"); }

执行以上 JavaScript 代码，输出结果为：

```
数字是正数
```

### if...else 语句

一个 if 语句后可跟一个可选的 else 语句，else 语句在布尔表达式为 false 时执行。

#### 语法

语法格式如下所示：

```

if(boolean_expression){
# 在布尔表达式 boolean_expression 为 true 执行
}else{
# 在布尔表达式 boolean_expression 为 false 执行
}

```

如果布尔表达式 boolean_expression 为 true，则执行 if 块内的代码。如果布尔表达式为 false，则执行 else 块内的代码。

#### 流程图

#### 实例

### TypeScript

var num:number = 12; if (num % 2==0) { console.log("偶数"); } else { console.log("奇数"); }

编译以上代码得到如下 JavaScript 代码：

### JavaScript

var num = 12; if (num % 2 == 0) { console.log("偶数"); } else { console.log("奇数"); }

执行以上 JavaScript 代码，输出结果为：

```
偶数
```

### if...else if....else 语句

if...else if....else 语句在执行多个判断条件的时候很有用。

#### 语法

语法格式如下所示：

```

if(boolean_expression 1) {
# 在布尔表达式 boolean_expression 1 为 true 执行
} else if( boolean_expression 2) {
# 在布尔表达式 boolean_expression 2 为 true 执行
} else if( boolean_expression 3) {
# 在布尔表达式 boolean_expression 3 为 true 执行
} else {
# 布尔表达式的条件都为 false 时执行
}

```

需要注意以下几点：

- 一个 if 判断语句可以有 0 或 1 个 else 语句，她必需在 else..if 语句后面。
- 一个 if 判断语句可以有 0 或多个 else..if，这些语句必需在 else 之前。
- 一旦执行了 else..if 内的代码，后面的 else..if 或 else 将不再执行。

#### 实例

### TypeScript

var num:number = 2 if(num > 0) { console.log(num+" 是正数") } else if(num < 0) { console.log(num+" 是负数") } else { console.log(num+" 不是正数也不是负数") }

编译以上代码得到如下 JavaScript 代码：

### JavaScript

var num = 2; if (num > 0) { console.log(num + " 是正数"); } else if (num < 0) { console.log(num + " 是负数"); } else { console.log(num + " 不是正数也不是负数"); }

执行以上 JavaScript 代码，输出结果为：

```
2 是正数
```

### switch…case 语句

一个 switch 语句允许测试一个变量等于多个值时的情况。每个值称为一个 case，且被测试的变量会对每个 switch case 进行检查。

switch 语句的语法：

switch(expression){ case constant-expression : statement(s); break; /* 可选的 */ case constant-expression : statement(s); break; /* 可选的 */ /* 您可以有任意数量的 case 语句 */ default : /* 可选的 */ statement(s); }

switch 语句必须遵循下面的规则：

- switch 语句中的 expression 是一个要被比较的表达式，可以是任何类型，包括基本数据类型（如 number、string、boolean）、对象类型（如 object、Array、Map）以及自定义类型（如 class、interface、enum）等。
- 在一个 switch 中可以有任意数量的 case 语句。每个 case 后跟一个要比较的值和一个冒号。
- case 的 constant-expression 必须与 switch 中的变量 expression 具有相同或兼容的数据类型。
- 当被测试的变量等于 case 中的常量时，case 后跟的语句将被执行，直到遇到 break 语句为止。
- 当遇到 break 语句时，switch 终止，控制流将跳转到 switch 语句后的下一行。
- 不是每一个 case 都需要包含 break。如果 case 语句不包含 break，控制流将会 继续 后续的 case，直到遇到 break 为止。
- 一个 switch 语句可以有一个可选的 default case，出现在 switch 的结尾。default 关键字则表示当表达式的值与所有 case 值都不匹配时执行的代码块。default case 中的 break 语句不是必需的。

### 流程图

#### 实例

### TypeScript

var grade:string = "A"; switch(grade) { case "A": { console.log("优"); break; } case "B": { console.log("良"); break; } case "C": { console.log("及格"); break; } case "D": { console.log("不及格"); break; } default: { console.log("非法输入"); break; } }

编译以上代码得到如下 JavaScript 代码：

### JavaScript

var grade = "A"; switch (grade) { case "A": { console.log("优"); break; } case "B": { console.log("良"); break; } case "C": { console.log("及格"); break; } case "D": { console.log("不及格"); break; } default: { console.log("非法输入"); break; } }

执行以上 JavaScript 代码，输出结果为：

```
优
```

---

## TypeScript 循环

Source: https://www.runoob.com/typescript/ts-loop.html

## TypeScript 循环

有的时候，我们可能需要多次执行同一块代码。一般情况下，语句是按顺序执行的：函数中的第一个语句先执行，接着是第二个语句，依此类推。

编程语言提供了更为复杂执行路径的多种控制结构。

循环语句允许我们多次执行一个语句或语句组，下面是大多数编程语言中循环语句的流程图：

### for 循环

TypeScript for 循环用于多次执行一个语句序列，简化管理循环变量的代码。

#### 语法

语法格式如下所示：

```

for ( init; condition; increment ){
statement(s);
}

```

下面是 for 循环的控制流程解析：

- init 会首先被执行，且只会执行一次。这一步允许您声明并初始化任何循环控制变量。您也可以不在这里写任何语句，只要有一个分号出现即可。
- 接下来，会判断 condition。如果为 true，则执行循环主体。如果为 false，则不执行循环主体，且控制流会跳转到紧接着 for 循环的下一条语句。
- 在执行完 for 循环主体后，控制流会跳回上面的 increment 语句。该语句允许您更新循环控制变量。该语句可以留空，只要在条件后有一个分号出现即可。
- 条件再次被判断。如果为 true，则执行循环，这个过程会不断重复（循环主体，然后增加步值，再然后重新判断条件）。在条件变为 false 时，for 循环终止。

在这里，statement(s) 可以是一个单独的语句，也可以是几个语句组成的代码块。

condition 可以是任意的表达式，当条件为 true 时执行循环，当条件为 false 时，退出循环。

#### 流程图

#### 实例

以下实例计算 5 的阶乘， for 循环生成从 5 到 1 的数字，并计算每次循环数字的乘积。

### TypeScript

var num:number = 5; var i:number; var factorial = 1; for(i = num;i>=1;i--) { factorial *= i; } console.log(factorial)

编译以上代码得到如下 JavaScript 代码：

### JavaScript

var num = 5; var i; var factorial = 1; for (i = num; i >= 1; i--) { factorial *= i; } console.log(factorial);

执行以上 JavaScript 代码，输出结果为：

```

120

```

### for...in 循环

for...in 语句用于一组值的集合或列表进行迭代输出。

#### 语法

语法格式如下所示：

```

for (var val in list) {
//语句
}

```
val 需要为 string 或 any 类型。

#### 实例

### TypeScript

var j:any; var n:any = "a b c" for(j in n) { console.log(n[j]) }

编译以上代码得到如下 JavaScript 代码：

### JavaScript

var j; var n = "a b c"; for (j in n) { console.log(n[j]); }

执行以上 JavaScript 代码，输出结果为：

```

a

b

c

```

###

### for…of 、forEach、every 和 some 循环

此外，TypeScript 还支持 for…of 、forEach、every 和 some 循环。

for...of 语句创建一个循环来迭代可迭代的对象。在 ES6 中引入的 for...of 循环，以替代 for...in 和 forEach() ，并支持新的迭代协议。for...of 允许你遍历 Arrays（数组）, Strings（字符串）, Maps（映射）, Sets（集合）等可迭代的数据结构等。

### TypeScript for...of 循环

let someArray = [1, "string", false]; for (let entry of someArray) { console.log(entry); // 1, "string", false }

forEach、every 和 some 是 JavaScript 的循环语法，TypeScript 作为 JavaScript 的语法超集，当然默认也是支持的。

因为 forEach 在 iteration 中是无法返回的，所以可以使用 every 和 some 来取代 forEach。

### TypeScript forEach 循环

let list = [4, 5, 6]; list.forEach((val, idx, array) => { // val: 当前值 // idx：当前index // array: Array });

### TypeScript every 循环

let list = [4, 5, 6]; list.every((val, idx, array) => { // val: 当前值 // idx：当前index // array: Array return true; // Continues // Return false will quit the iteration });

### while 循环

while 语句在给定条件为 true 时，重复执行语句或语句组。循环主体执行之前会先测试条件。

#### 语法

语法格式如下所示：

```

while(condition)
{
statement(s);
}

```

在这里，statement(s) 可以是一个单独的语句，也可以是几个语句组成的代码块。

condition 可以是任意的表达式，当条件为 true 时执行循环。 当条件为 false 时，程序流将退出循环。

#### 流程图

图表中，while 循环的关键点是循环可能一次都不会执行。当条件为 false 时，会跳过循环主体，直接执行紧接着 while 循环的下一条语句。

#### 实例

### TypeScript

var num:number = 5; var factorial:number = 1; while(num >=1) { factorial = factorial * num; num--; } console.log("5 的阶乘为："+factorial);

编译以上代码得到如下 JavaScript 代码：

### JavaScript

var num = 5; var factorial = 1; while (num >= 1) { factorial = factorial * num; num--; } console.log("5 的阶乘为：" + factorial);

执行以上 JavaScript 代码，输出结果为：

```

5 的阶乘为：120

```

### do...while 循环

不像 for 和 while 循环，它们是在循环头部测试循环条件。do...while 循环是在循环的尾部检查它的条件。

#### 语法

语法格式如下所示：

```

do
{
statement(s);
}while( condition );

```

请注意，条件表达式出现在循环的尾部，所以循环中的 statement(s) 会在条件被测试之前至少执行一次。

如果条件为 true，控制流会跳转回上面的 do，然后重新执行循环中的 statement(s)。这个过程会不断重复，直到给定条件变为 false 为止。

#### 流程图

#### 实例

### TypeScript

var n:number = 10; do { console.log(n); n--; } while(n>=0);

编译以上代码得到如下 JavaScript 代码：

### JavaScript

var num = 5; var n = 10; do { console.log(n); n--; } while (n >= 0);

执行以上 JavaScript 代码，输出结果为：

```

10
9
8
7
6
5
4
3
2
1
0

```

### break 语句

break 语句有以下两种用法：

- 当 break 语句出现在一个循环内时，循环会立即终止，且程序流将继续执行紧接着循环的下一条语句。
- 它可用于终止 switch 语句中的一个 case。

如果您使用的是嵌套循环（即一个循环内嵌套另一个循环），break 语句会停止执行最内层的循环，然后开始执行该块之后的下一行代码。

#### 语法

语法格式如下所示：

```

break;

```

#### 流程图

#### 实例

### TypeScript

var i:number = 1 while(i<=10) { if (i % 5 == 0) { console.log ("在 1~10 之间第一个被 5 整除的数为 : "+i) break // 找到一个后退出循环 } i++ } // 输出 5 然后程序执行结束

编译以上代码得到如下 JavaScript 代码：

### JavaScript

var i = 1; while (i <= 10) { if (i % 5 == 0) { console.log("在 1~10 之间第一个被 5 整除的数为 : " + i); break; // 找到一个后退出循环 } i++; } // 输出 5 然后程序执行结束

执行以上 JavaScript 代码，输出结果为：

```

在 1~10 之间第一个被 5 整除的数为 : 5

```

### continue 语句

continue 语句有点像 break 语句。但它不是强制终止，continue 会跳过当前循环中的代码，强迫开始下一次循环。

对于 for 循环，continue 语句执行后自增语句仍然会执行。对于 while 和 do...while 循环，continue 语句重新执行条件判断语句。

#### 语法

语法格式如下所示：

```

continue;

```

#### 流程图

#### 实例

### TypeScript

var num:number = 0 var count:number = 0; for(num=0;num<=20;num++) { if (num % 2==0) { continue } count++ } console.log ("0 ~20 之间的奇数个数为: "+count) //输出10个偶数

编译以上代码得到如下 JavaScript 代码：

### JavaScript

var num = 0; var count = 0; for (num = 0; num <= 20; num++) { if (num % 2 == 0) { continue; } count++; } console.log("0 ~20 之间的奇数个数为: " + count); //输出 10

执行以上 JavaScript 代码，输出结果为：

```

0 ~20 之间的奇数个数为: 10

```

### 无限循环

无限循环就是一直在运行不会停止的循环。 for 和 while 循环都可以创建无限循环。

for 创建无限循环语法格式：

```
for(;;) {
// 语句
}
```

实例

```
for(;;) {
console.log("这段代码会不停的执行")
}
```

while 创建无限循环语法格式：

```
while(true) {
// 语句
}
```

实例

```
while(true) {
console.log("这段代码会不停的执行")
}
```

---

## TypeScript 函数

Source: https://www.runoob.com/typescript/ts-function.html

## TypeScript 函数

函数是一组一起执行一个任务的语句。

您可以把代码划分到不同的函数中。如何划分代码到不同的函数中是由您来决定的，但在逻辑上，划分通常是根据每个函数执行一个特定的任务来进行的。

函数声明告诉编译器函数的名称、返回类型和参数。函数定义提供了函数的实际主体。

### 函数定义

函数就是包裹在花括号中的代码块，前面使用了关键词 function：

语法格式如下所示：

```

function function_name()
{
// 执行代码
}

```

#### 实例

### TypeScript

function () { // 函数定义 console.log("调用函数") }

### 调用函数

函数只有通过调用才可以执行函数内的代码。

语法格式如下所示：

```

function_name()

```

#### 实例

### TypeScript

function test() { // 函数定义 console.log("调用函数") } test() // 调用函数

### 函数返回值

有时，我们会希望函数将执行的结果返回到调用它的地方。

通过使用 return 语句就可以实现。

在使用 return 语句时，函数会停止执行，并返回指定的值。

语法格式如下所示：

```

function function_name():return_type {
// 语句
return value;
}

```

- return_type 是返回值的类型。
- return 关键词后跟着要返回的结果。
- 一般情况下，一个函数只有一个 return 语句。
- 返回值的类型需要与函数定义的返回类型(return_type)一致。

#### 实例

### TypeScript

// 函数定义 function greet():string { // 返回一个字符串 return "Hello World" } function caller() { var msg = greet() // 调用 greet() 函数 console.log(msg) } // 调用函数 caller()

- 实例中定义了函数 greet()，返回值的类型为 string。
- greet() 函数通过 return 语句返回给调用它的地方，即变量 msg，之后输出该返回值。。

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

// 函数定义 function greet() { return "Hello World"; } function caller() { var msg = greet(); // 调用 greet() 函数 console.log(msg); } // 调用函数 caller();

### 带参数函数

在调用函数时，您可以向其传递值，这些值被称为参数。

这些参数可以在函数中使用。

您可以向函数发送多个参数，每个参数使用逗号 , 分隔：

语法格式如下所示：

```

function func_name( param1 [:datatype], param2 [:datatype]) {
}

```

- param1、param2 为参数名。
- datatype 为参数类型。

#### 实例

### TypeScript

function add(x: number, y: number): number { return x + y; } console.log(add(1,2))

- 实例中定义了函数 add()，返回值的类型为 number。
- add() 函数中定义了两个 number 类型的参数，函数内将两个参数相加并返回。

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

function add(x, y) { return x + y; } console.log(add(1, 2));

输出结果为：

```
3
```

### 可选参数和默认参数

#### 可选参数

在 TypeScript 函数里，如果我们定义了参数，则我们必须传入这些参数，除非将这些参数设置为可选，可选参数使用问号标识 ？。

实例

### TypeScript

function buildName(firstName: string, lastName: string) { return firstName + " " + lastName; } let result1 = buildName("Bob"); // 错误，缺少参数 let result2 = buildName("Bob", "Adams", "Sr."); // 错误，参数太多了 let result3 = buildName("Bob", "Adams"); // 正确

以下实例，我们将 lastName 设置为可选参数：

### TypeScript

function buildName(firstName: string, lastName?: string) { if (lastName) return firstName + " " + lastName; else return firstName; } let result1 = buildName("Bob"); // 正确 let result2 = buildName("Bob", "Adams", "Sr."); // 错误，参数太多了 let result3 = buildName("Bob", "Adams"); // 正确

可选参数必须跟在必需参数后面。 如果上例我们想让 firstName 是可选的，lastName 必选，那么就要调整它们的位置，把 firstName 放在后面。

如果都是可选参数就没关系。

#### 默认参数

我们也可以设置参数的默认值，这样在调用函数的时候，如果不传入该参数的值，则使用默认参数，语法格式为：

```
function function_name(param1[:type],param2[:type] = default_value) {
}
```
注意：参数不能同时设置为可选和默认。

实例

以下实例函数的参数 rate 设置了默认值为 0.50，调用该函数时如果未传入参数则使用该默认值：

### TypeScript

function calculate_discount(price:number,rate:number = 0.50) { var discount = price * rate; console.log("计算结果: ",discount); } calculate_discount(1000) calculate_discount(1000,0.30)

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

function calculate_discount(price, rate) { if (rate === void 0) { rate = 0.50; } var discount = price * rate; console.log("计算结果: ", discount); } calculate_discount(1000); calculate_discount(1000, 0.30);

输出结果为：

```
计算结果: 500
计算结果: 300
```

### 剩余参数

有一种情况，我们不知道要向函数传入多少个参数，这时候我们就可以使用剩余参数来定义。

剩余参数语法允许我们将一个不确定数量的参数作为一个数组传入。

### TypeScript

function buildName(firstName: string, ...restOfName: string[]) { return firstName + " " + restOfName.join(" "); } let employeeName = buildName("Joseph", "Samuel", "Lucas", "MacKinzie");

函数的最后一个命名参数 restOfName 以 ... 为前缀，它将成为一个由剩余参数组成的数组，索引值从0（包括）到 restOfName.length（不包括）。

### TypeScript

function addNumbers(...nums:number[]) { var i; var sum:number = 0; for(i = 0;i<nums.length;i++) { sum = sum + nums[i]; } console.log("和为：",sum) } addNumbers(1,2,3) addNumbers(10,10,10,10,10)

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

function addNumbers() { var nums = []; for (var _i = 0; _i < arguments.length; _i++) { nums[_i] = arguments[_i]; } var i; var sum = 0; for (i = 0; i < nums.length; i++) { sum = sum + nums[i]; } console.log("和为：", sum); } addNumbers(1, 2, 3); addNumbers(10, 10, 10, 10, 10);

输出结果为：

```
和为： 6
和为： 50
```

### 匿名函数

匿名函数是一个没有函数名的函数。

匿名函数在程序运行时动态声明，除了没有函数名外，其他的与标准函数一样。

我们可以将匿名函数赋值给一个变量，这种表达式就成为函数表达式。

语法格式如下：

```
var res = function( [arguments] ) { ... }
```

#### 实例

不带参数匿名函数：

### TypeScript

var msg = function() { return "hello world"; } console.log(msg())

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var msg = function () { return "hello world"; }; console.log(msg());

输出结果为：

```
hello world
```

带参数匿名函数：

### TypeScript

var res = function(a:number,b:number) { return a*b; }; console.log(res(12,2))

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var res = function (a, b) { return a * b; }; console.log(res(12, 2));

输出结果为：

```
24
```

#### 匿名函数自调用 匿名函数自调用在函数后使用 () 即可：

### TypeScript

(function () { var x = "Hello!!"; console.log(x) })()

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

(function () { var x = "Hello!!"; console.log(x) })()

输出结果为：

```
Hello!!
```

### 构造函数

TypeScript 也支持使用 JavaScript 内置的构造函数 Function() 来定义函数：

语法格式如下：

```
var res = new Function ([arg1[, arg2[, ...argN]],] functionBody)
```

参数说明：

- arg1, arg2, ... argN：参数列表。
- functionBody：一个含有包括函数定义的 JavaScript 语句的字符串。

#### 实例

### TypeScript

var myFunction = new Function("a", "b", "return a * b"); var x = myFunction(4, 3); console.log(x);

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var myFunction = new Function("a", "b", "return a * b"); var x = myFunction(4, 3); console.log(x);

输出结果为：

```
12
```

### 递归函数 递归函数即在函数内调用函数本身。

举个例子：
从前有座山，山里有座庙，庙里有个老和尚，正在给小和尚讲故事呢！故事是什么呢？"从前有座山，山里有座庙，庙里有个老和尚，正在给小和尚讲故事呢！故事是什么呢？'从前有座山，山里有座庙，庙里有个老和尚，正在给小和尚讲故事呢！故事是什么呢？……'"

#### 实例

### TypeScript

function factorial(number) { if (number <= 0) { // 停止执行 return 1; } else { return (number * factorial(number - 1)); // 调用自身 } }; console.log(factorial(6)); // 输出 720

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

function factorial(number) { if (number <= 0) { // 停止执行 return 1; } else { return (number * factorial(number - 1)); // 调用自身 } } ; console.log(factorial(6)); // 输出 720

输出结果为：

```
720
```

### Lambda 函数

Lambda 函数也称之为箭头函数。

箭头函数表达式的语法比函数表达式更短。

函数只有一行语句：

```
( [param1, param2,…param n] )=>statement;
```

#### 实例 以下实例声明了 lambda 表达式函数，函数返回两个数的和：

### TypeScript

var foo = (x:number)=>10 + x console.log(foo(100)) //输出结果为 110

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var foo = function (x) { return 10 + x; }; console.log(foo(100)); //输出结果为 110

输出结果为：

```
110
```

函数是一个语句块：

```
( [param1, param2,…param n] )=> {

// 代码块
}
```

#### 实例

以下实例声明了 lambda 表达式函数，函数返回两个数的和：

### TypeScript

var foo = (x:number)=> { x = 10 + x console.log(x) } foo(100)

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var foo = function (x) { x = 10 + x; console.log(x); }; foo(100);

输出结果为：

```
110
```
我们可以不指定函数的参数类型，通过函数内来推断参数类型:

### TypeScript

var func = (x)=> { if(typeof x=="number") { console.log(x+" 是一个数字") } else if(typeof x=="string") { console.log(x+" 是一个字符串") } } func(12) func("Tom")

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var func = function (x) { if (typeof x == "number") { console.log(x + " 是一个数字"); } else if (typeof x == "string") { console.log(x + " 是一个字符串"); } }; func(12); func("Tom");

输出结果为：

```

12 是一个数字
Tom 是一个字符串

```

单个参数 () 是可选的：

### TypeScript

var display = x => { console.log("输出为 "+x) } display(12)

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var display = function (x) { console.log("输出为 " + x); }; display(12);

输出结果为：

```

输出为 12

```

无参数时可以设置空括号：

### TypeScript

var disp =()=> { console.log("Function invoked"); } disp();

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var disp = function () { console.log("调用函数"); }; disp();

输出结果为：

```

调用函数

```

### 函数重载

重载是方法名字相同，而参数不同，返回类型可以相同也可以不同。

每个重载的方法（或者构造函数）都必须有一个独一无二的参数类型列表。

参数类型不同：

```
function disp(string):void;
function disp(number):void;
```

参数数量不同：

```
function disp(n1:number):void;
function disp(x:number,y:number):void;
```

参数类型顺序不同：

```
function disp(n1:number,s1:string):void;
function disp(s:string,n:number):void;
```

如果参数类型不同，则参数类型应设置为 any。

参数数量不同你可以将不同的参数设置为可选。

#### 实例

以下实例定义了参数类型与参数数量不同：

### TypeScript

function disp(s1:string):void; function disp(n1:number,s1:string):void; function disp(x:any,y?:any):void { console.log(x); console.log(y); } disp("abc") disp(1,"xyz");

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

function disp(x, y) { console.log(x); console.log(y); } disp("abc"); disp(1, "xyz");

输出结果为：

```

abc
undefined
1
xyz

```

---

## TypeScript 函数重载

Source: https://www.runoob.com/typescript/ts-function-overload.html

## TypeScript 函数重载

函数重载（Function Overloading）允许为一个函数定义多个签名，编译器会根据传入的参数类型选择正确的实现。 函数重载工作原理 函数重载签名（声明） // 签名 1 function add(a: number, b: number): number; // 签名 2 function add(a: string, b: string): string; 编译器 匹配 函数实现（实际代码） function add(a: any, b: any): any { return a + b; } 必须兼容所有签名 调用时类型推断 add(1, 2) → number add("a", "b") → string add(true, false) → any

### 基本语法

先声明多个函数签名，然后实现一个统一函数。

### 实例

// 函数重载签名
function add(a: number, b: number): number;
function add(a: string, b: string): string;
function add(a: any, b: any): any {
return a + b;
}

console.log("数字相加: " + add(1, 2));
console.log("字符串相加: " + add("Hello, ", "World"));

运行结果：

```
数字相加: 3
字符串相加: Hello, World

```

### 多参数重载

可以定义多个参数的不同组合。

### 实例

// 多种重载签名
function greet(name: string): string;
function greet(name: string, greeting: string): string;

// 实现
function greet(name: any, greeting?: any): any {
if (greeting) {
return greeting + ", " + name + "!";
}
return "Hello, " + name + "!";
}

console.log(greet("Alice"));
console.log(greet("Bob", "Hi"));

运行结果：

```
Hello, Alice!
Hi, Bob!

```

### 方法重载

类中的方法也可以使用重载。

### 实例

class Calculator {
// 重载签名
add(a: number, b: number): number;
add(a: string, b: string): string;
add(a: number, b: string): string;
add(a: any, b: any): any {
return a + b;
}
}

var calc = new Calculator();
console.log("数字: " + calc.add(1, 2));
console.log("字符串: " + calc.add("Hello", "World"));
console.log("混合: " + calc.add(5, " apples"));

运行结果：

```
数字: 3
字符串: HelloWorld
混合: 5 apples

```

### 构造函数重载

构造函数同样可以重载。

### 实例

class User {
name: string;
age: number;

// 构造函数重载
constructor(name: string);
constructor(name: string, age: number);
constructor(name: any, age?: any) {
this.name = name;
this.age = age || 0;
}
}

var user1 = new User("Alice");
var user2 = new User("Bob", 25);

console.log("用户1: " + JSON.stringify(user1));
console.log("用户2: " + JSON.stringify(user2));

运行结果：

```
用户1: {"name":"Alice","age":0}
用户2: {"name":"Bob","age":25}

```

### 重载与联合类型

使用重载而不是联合类型可以获得更精确的类型推断。

### 实例

// 推荐：使用重载
function process(value: number): number;
function process(value: string): string;
function process(value: any): any {
if (typeof value === "number") {
return value * 2;
}
return value.toUpperCase();
}

// TypeScript 知道返回类型
var numResult: number = process(10); // number
var strResult: string = process("hello"); // string

console.log("数字结果: " + numResult);
console.log("字符串结果: " + strResult);

运行结果：

```
数字结果: 20
字符串结果: HELLO

```

### 注意事项

- 重载签名必须放在实现签名之前
- 实现签名必须兼容所有重载签名
- 重载签名只是类型声明，不生成实际代码

### 总结

- 函数重载：定义多个签名，编译器选择匹配的实现
- 方法重载：类中同样适用
- 构造函数重载：提供多种初始化方式
- 优于联合类型：返回类型更精确

---

## TypeScript 箭头函数与 this

Source: https://www.runoob.com/typescript/ts-arrow-function.html

## TypeScript 箭头函数与 this

箭头函数（Arrow Functions）是 ES6 引入的重要特性，也是 TypeScript 中的常用语法。

与普通函数最大的区别是，箭头函数不绑定自己的 this，而是捕获定义时所在上下文的 this。

这解决了 JavaScript 中常见的 this 指向问题。 箭头函数 this 绑定机制 普通函数 setTimeout(function() { console.log(this.name); }, 100); this 指向 window/undefined vs 箭头函数 setTimeout(() => { console.log(this.name); }, 100); this 指向定义时的上下文 箭头函数语法 多参数: (a, b) => a + b 单参数: n => n * 2 无参数: () => 42

### 为什么需要箭头函数

在 JavaScript 中，this 的指向常常令人困惑。

普通函数中的 this 取决于函数如何调用，而不是如何定义。

这导致在回调函数、事件处理等场景下，this 的指向常常出错。

箭头函数的出现解决了这个问题，它让 this 的行为更加可预测。

概念说明：箭头函数使用 `=>` 语法定义，它不绑定自己的 this，而是继承外层作用域的 this。

### 箭头函数基础

箭头函数提供更简洁的函数定义语法。

它可以省略大括号和 return 关键字（当函数体是单个表达式时）。

### 实例

// 传统函数定义
var add1 = function(a: number, b: number): number {
return a + b;
};

// 箭头函数：单行函数可以省略大括号和 return
var add2 = (a: number, b: number): number => a + b;

// 单参数可以省略括号
var double = (n: number): number => n * 2;

// 无参数函数
var getRandom = (): number => Math.random();

console.log("add1: " + add1(1, 2));
console.log("add2: " + add2(3, 4));
console.log("double: " + double(5));
console.log("random: " + getRandom().toFixed(2));

运行结果：

```
add1: 3
add2: 7
double: 10
random: 0.xx

```

语法说明：当只有单个参数时，括号可以省略；但没有参数或多于一个参数时，必须使用括号。

### 箭头函数与 this

箭头函数最核心的特性是不绑定自己的 this。

它会捕获定义时所在外层作用域的 this，并保持不变。

这解决了普通函数中 this 指向混乱的问题。

### 实例

// 使用普通函数
function Person1() {
this.name = "Alice";

// 普通函数会创建自己的 this
// 在 setTimeout 回调中，this 指向 window（浏览器）或 undefined（严格模式）
setTimeout(function() {
console.log("普通函数: " + this.name); // this.name 为 undefined
}, 100);
}

// 使用箭头函数
function Person2() {
this.name = "Bob";

// 箭头函数不创建自己的 this
// 它捕获外层的 this，所以能正确访问到 name
setTimeout(() => {
console.log("箭头函数: " + this.name); // this.name 为 "Bob"
}, 100);
}

// 测试
new Person1();
new Person2();

运行结果：

```
普通函数: undefined
箭头函数: Bob

```

关键区别：普通函数的 this 在调用时确定，箭头函数的 this 在定义时确定。这是两者最核心的区别。

### 类中的箭头函数

在 TypeScript 类中，可以使用箭头函数作为类的方法或属性。

这样可以确保方法被传递或作为回调使用时，this 仍然指向类的实例。

### 实例

// 定义计数器类
class Counter {
// 计数器的当前值
count: number = 0;

// 使用箭头函数作为类属性
// 每次创建实例时，都会创建一个新的函数
// this 指向实例
increment = () => {
this.count++;
console.log("当前计数: " + this.count);
};

// 普通方法
decrement() {
this.count--;
console.log("当前计数: " + this.count);
}
}

// 创建计数器实例
var counter = new Counter();

// 调用箭头函数方法
counter.increment();
counter.increment();

// 调用普通方法
counter.decrement();

运行结果：

```
当前计数: 1
当前计数: 2
当前计数: 1

```

权衡：箭头函数属性会在每个实例中创建新函数，可能增加内存开销。但当方法需要作为回调传递时，这是最好的选择。

### 回调函数中的 this

箭头函数在数组方法（map、filter、reduce 等）的回调中特别有用。

它确保回调内部可以正确访问外层的 this。

### 实例

// 定义处理器对象
var handler = {
// 处理器名称
name: "Handler",
// 数字数组
numbers: [1, 2, 3],

// 处理方法
processAll: function() {
// 使用箭头函数的回调
// 箭头函数捕获外层的 this，所以可以正确访问 this.name
this.numbers.forEach((n) => {
console.log(this.name + ": " + n);
});
}
};

// 调用处理方法
handler.processAll();

运行结果：

```
Handler: 1
Handler: 2
Handler: 3

```

最佳实践：在类的回调方法、数组方法的回调、事件处理函数中，优先使用箭头函数以避免 this 问题。

### 箭头函数的类型

TypeScript 中箭头函数的类型注解使用不同的语法。

使用 `=>` 而不是冒号来定义函数类型。

### 实例

// 直接定义箭头函数类型
// (a: number, b: number) => number 表示接受两个 number 参数，返回 number
var add: (a: number, b: number) => number = (a, b) => a + b;

// 使用接口定义箭头函数类型
// 这种方式更适合在接口或类型别名中复用
interface MathOperation {
// 定义函数签名
(a: number, b: number): number;
}

// 使用接口类型
var multiply: MathOperation = (a, b) => a * b;

console.log("加法: " + add(2, 3));
console.log("乘法: " + multiply(4, 5));

运行结果：

```
加法: 5
乘法: 20

```

注意：箭头函数类型的语法是 `(params) => returnType`，不是传统的 `(params): returnType`。

### 何时使用箭头函数

箭头函数虽然简洁，但并非所有场景都适用。

了解何时使用箭头函数可以写出更好的代码。

- 需要保持 this 上下文时：如回调函数、事件处理、数组方法
- 简单的一行函数：如 map、filter、reduce 的回调
- 类方法需要传递时：作为回调传递给其他函数

警告：不要在需要动态 this 的场景（如事件处理函数中需要获取事件目标）使用箭头函数，因为此时 this 已经被固定。

### 注意事项

- 不绑定 arguments：箭头函数不绑定自己的 arguments 对象
- 不能用作构造函数：不能使用 new 关键字调用箭头函数
- 不能用作方法：在对象字面量中作为方法时，this 可能不符合预期
- 适合回调：在需要保持 this 上下文的场景优先使用

最佳实践：根据场景选择：需要 this 绑定时用箭头函数，需要动态 this 时用普通函数。

### 总结

箭头函数是现代 JavaScript/TypeScript 开发中不可或缺的特性。

- 语法简洁：使用 `=>` 语法
- 不绑定 this：捕获定义时的上下文
- 适合回调：数组方法、事件处理等场景
- 类中使用：箭头属性方法解决传递问题
- 类型注解：使用 `=>` 语法定义类型

建议：在 TypeScript 开发中，充分利用箭头函数的 this 绑定特性，可以写出更安全、更易维护的代码。

---

## TypeScript 迭代器与生成器

Source: https://www.runoob.com/typescript/ts-iterator-generator.html

## TypeScript 迭代器与生成器

迭代器和生成器是 JavaScript/TypeScript 中处理集合的重要模式。

它们提供了一种统一的遍历数据的方式，让处理大数据流、无限序列等变得更加简单。 迭代器与生成器工作流程 迭代器协议 实现 Symbol.iterator 返回 next() 方法 { done, value } for...of 循环 自动调用 next() 遍历所有元素 done=true 时停止 生成器 function* yield 暂停 惰性计算 生成器特性 惰性求值 - 按需生成 状态保持 - 暂停位置 可组合 - yield* 委托

### 为什么需要迭代器和生成器

在处理集合数据时，我们经常需要遍历数组、对象等数据结构。

迭代器提供了一种统一的、可自定义的遍历接口，让任何对象都可以被遍历。

生成器是创建迭代器的简洁方式，它允许你使用函数来暂停和恢复执行，非常适合处理大数据流或无限序列。

概念说明：迭代器是一个对象，它提供 next() 方法用于遍历数据。生成器是一种特殊的函数，可以在执行过程中暂停并返回一个值。

### 可迭代协议

实现 Symbol.iterator 方法的对象可以被 for...of 循环遍历。

### 实例

// 数组默认可迭代
var arr = [1, 2, 3];
for (var _i = 0, arr_1 = arr; _i < arr_1.length; _i++) {
var item = arr_1[_i];
console.log("数组元素: " + item);
}

// 字符串默认可迭代
var str = "hello";
for (var _i = 0, str_1 = str; _i < str_1.length; _i++) {
var char = str_1[_i];
console.log("字符: " + char);
}

运行结果：

```
数组元素: 1
数组元素: 2
数组元素: 3
字符: h
字符: e
字符: l
字符: l
字符: o

```

说明：数组和字符串都内置实现了 Symbol.iterator 方法，所以可以直接使用 for...of 遍历。

### 自定义可迭代对象

让普通对象实现 Symbol.iterator 接口，使其可被遍历。

### 实例

// 创建自定义可迭代对象：范围
var range = {
from: 1,
to: 5,
// 实现 Symbol.iterator 方法
[Symbol.iterator]: function() {
return {
current: this.from,
last: this.to,
// next 方法返回迭代结果
next: function() {
if (this.current <= this.last) {
// 未完成，返回当前值并递增
return { done: false, value: this.current++ };
}
// 已完成
return { done: true, value: undefined };
}
};
}
};

// 使用 for...of 遍历
for (var _i = 0, range_1 = range; _i < range_1.length; _i++) {
var num = range_1[_i];
console.log("范围: " + num);
}

迭代器协议：迭代器必须有一个 next() 方法，返回 { done: boolean, value: any } 格式的对象。

### 生成器函数

使用 function* 语法创建生成器，使用 yield 暂停执行并返回值。

### 实例

// 生成器函数：使用 function* 语法
function* numberGenerator() {
yield 1; // 暂停并返回 1
yield 2; // 暂停并返回 2
yield 3; // 暂停并返回 3
}

// 创建生成器实例
var gen = numberGenerator();

// 每次调用 next() 都会执行到下一个 yield
console.log("第一个: " + gen.next().value);
console.log("第二个: " + gen.next().value);
console.log("第三个: " + gen.next().value);
console.log("完成: " + gen.next().done);

运行结果：

```
第一个: 1
第二个: 2
第三个: 3
完成: true

```

生成器：生成器函数会返回一个迭代器，每次调用 next() 都会执行到下一个 yield 语句。

### 无限生成器

生成器可以产生无限序列，由于是惰性求值，不会占用无限内存。

### 实例

// 无限数字生成器
// 每次调用只生成一个数字，不会一次性生成所有数字
function* infiniteNumbers() {
var n = 1;
while (true) { // 无限循环
yield n++; // 暂停并返回当前值，然后递增
}
}

var gen = infiniteNumbers();
console.log("第1个: " + gen.next().value);
console.log("第2个: " + gen.next().value);
console.log("第3个: " + gen.next().value);

// 只获取前5个数字
var nums = [];
var iter = infiniteNumbers();
for (var i = 0; i < 5; i++) {
nums.push(iter.next().value);
}
console.log("前5个: " + nums);

运行结果：

```
第1个: 1
第2个: 2
第3个: 3
前5个: 1,2,3,4,5

```

惰性求值：生成器的最大优势是惰性求值，只有在调用 next() 时才会计算下一个值，非常适合处理无限序列。

### 委托生成器

使用 yield* 委托另一个生成器或可迭代对象。

### 实例

// 第一个生成器
function* gen1() {
yield 1;
yield 2;
}

// 第二个生成器
function* gen2() {
yield 3;
yield 4;
}

// 组合生成器：使用 yield* 委托
function* combined() {
yield* gen1(); // 委托给 gen1
yield* gen2(); // 委托给 gen2
}

// 遍历组合生成器
for (var _i = 0, combined_1 = combined(); _i < combined_1.length; _i++) {
var num = combined_1[_i];
console.log("值: " + num);
}

运行结果：

```
值: 1
值: 2
值: 3
值: 4

```

yield*：委托生成器可以组合多个生成器或可迭代对象，非常适合构建可复用的数据流。

### TypeScript 生成器类型

生成器的类型注解使用 Generator 类型。

### 实例

// 生成器类型：Generator<yield类型, return类型, next参数类型>
function* idGenerator(): Generator<number, void, unknown> {
var i = 1;
while (i <= 3) {
yield i++; // yield number 类型
}
// return void
}

var gen = idGenerator();
console.log(Array.from(gen));

类型说明：Generator<T, R, N> 表示：T 是 yield 的类型，R 是最终返回的类型，N 是 next() 参数的类型。

### 注意事项

- 迭代器协议：实现 Symbol.iterator 返回带 next() 方法的对象
- 生成器语法：使用 function* 而非 function
- yield 关键字：暂停执行并返回值
- 惰性计算：生成器按需计算，不会一次性生成所有值

最佳实践：处理大数据流、无限序列或需要暂停/恢复的场景时，使用生成器。

### 总结

迭代器和生成器是 TypeScript 中强大的数据处理工具。

- 可迭代对象：实现 Symbol.iterator 接口
- 生成器：使用 function* 和 yield 创建
- yield：暂停执行并返回值
- 委托：使用 yield* 组合多个生成器

建议：在需要遍历自定义对象、处理数据流或创建无限序列时，使用迭代器和生成器。

---

## TypeScript async/await 异步编程

Source: https://www.runoob.com/typescript/ts-async-await.html

## TypeScript async/await 异步编程

async/await 是 ES2017 引入的异步编程语法糖，让异步代码看起来像同步代码。 async/await 执行流程 Promise 方式 fetchData() .then(result => { console.log(result); }) async/await 方式 async function main() { const result = await fetchData(); console.log(result); } await 执行顺序 主线程 await 暂停 Promise 后台执行 恢复 完成 优势对比 ✓ 代码更简洁 ✓ 同步风格 ✓ 更好的错误堆栈 ✓ 易于调试 ✓ try/catch 处理

上图展示了 async/await 相比传统 Promise 的优势：代码更简洁，执行流程更清晰。

### Promise 基础

Promise 代表一个异步操作的最终结果。

### 实例

// 创建 Promise
var promise = new Promise(function(resolve, reject) {
var success = true;
if (success) {
resolve("操作成功");
} else {
reject("操作失败");
}
});

promise.then(function(result) {
console.log("成功: " + result);
})["catch"](function(error) {
console.log("失败: " + error);
});

运行结果：

```
成功: 操作成功

```

### async 函数

使用 async 关键字声明异步函数。

### 实例

// async 函数自动返回 Promise
async function greet(): Promise<string> {
return "Hello, World!";
}

greet().then(function(result) {
console.log("结果: " + result);
});

// 异步函数返回 Promise
async function getData() {
return { name: "Alice", age: 25 };
}

getData().then(function(data) {
console.log("数据: " + JSON.stringify(data));
});

运行结果：

```
结果: Hello, World!
数据: {"name":"Alice","age":25}

```

### await 关键字

await 等待 Promise 完成并获取结果。

### 实例

// 模拟异步操作
function delay(ms: number): Promise<string> {
return new Promise(function(resolve) {
setTimeout(function() {
resolve("完成!");
}, ms);
});
}

async function main() {
console.log("开始...");
var result = await delay(100);
console.log("结果: " + result);
console.log("结束");
}

main();

运行结果：

```
开始...
结果: 完成!
结束

```

### 错误处理

使用 try/catch 处理异步错误。

### 实例

function mayFail(shouldFail: boolean): Promise<string> {
return new Promise(function(resolve, reject) {
if (shouldFail) {
reject(new Error("操作失败"));
} else {
resolve("操作成功");
}
});
}

async function handleError() {
try {
var result = await mayFail(true);
console.log("结果: " + result);
} catch (error) {
console.log("捕获错误: " + error.message);
}
}

handleError();

运行结果：

```
捕获错误: 操作失败

```

### 并行执行

使用 Promise.all 并行执行多个异步操作。

### 实例

function fetchUser(id: number): Promise<{ id: number; name: string }> {
return Promise.resolve({ id: id, name: "User" + id });
}

async function main() {
// 串行执行
console.time("串行");
var user1 = await fetchUser(1);
var user2 = await fetchUser(2);
console.log("串行完成: " + user1.name + ", " + user2.name);
console.timeEnd("串行");

// 并行执行
console.time("并行");
var results = await Promise.all([fetchUser(1), fetchUser(2)]);
console.log("并行完成: " + results[0].name + ", " + results[1].name);
console.timeEnd("并行");
}

main();

运行结果：

```
串行完成: User1, User2
并行完成: User1, User2

```

### async/await 相比 Promise 的优势

- 代码更简洁、更易读
- 同步代码风格
- 更好的错误堆栈
- 易于调试

### 总结

- async：声明异步函数
- await：等待 Promise
- 错误处理：try/catch
- 并行：Promise.all

---

## TypeScript Promise 详解

Source: https://www.runoob.com/typescript/ts-promise.html

## TypeScript Promise 详解

Promise 是 JavaScript 异步编程的基础，TypeScript 对 Promise 有完整的类型支持。

通过泛型参数，可以精确地指定 Promise 解决值和拒绝值的类型。 Promise 工作流程 创建 Promise new Promise((resolve, reject) => {...}) pending 等待中 fulfilled ✓ rejected ✗ then() · catch() · finally()

### 为什么需要 Promise

在 JavaScript 中，很多操作是异步的，如网络请求、文件读取、定时器等。

Promise 提供了统一的异步编程接口，让异步代码更容易编写和管理。

TypeScript 通过泛型支持，让 Promise 的类型安全得到了保障。

概念说明：Promise 是一个对象，表示一个异步操作的最终结果。它有三种状态：pending（进行中）、fulfilled（已成功）、rejected（已失败）。

### 创建 Promise

使用 Promise 构造函数创建 Promise，传入执行器函数。

### 实例

// 创建 Promise，使用泛型指定解决值的类型
// Promise<string> 表示成功时返回字符串
var promise = new Promise<string>(function(resolve, reject) {
var success = true;
if (success) {
// 调用 resolve 表示操作成功，传入结果值
resolve("成功!");
} else {
// 调用 reject 表示操作失败，传入错误
reject(new Error("失败"));
}
});

// 使用 then 处理成功情况
promise.then(function(value) {
console.log("完成: " + value);
})["catch"](function(error) {
// 使用 catch 处理失败情况
console.log("错误: " + error.message);
});

运行结果：

```
完成: 成功!

```

泛型说明：`Promise<T>` 中的 T 是 Promise 成功解决时的值的类型。这让 TypeScript 能够推断返回值的类型。

### Promise 链式调用

then 和 catch 方法返回新的 Promise，可以链式调用。

### 实例

// 链式调用：每个 then 返回新的值，被下一个 then 接收
var promise = Promise.resolve(1)
.then(function(n) {
// 第一个 then，n = 1
return n * 2; // 返回 2
})
.then(function(n) {
// 第二个 then，n = 2
return n + 10; // 返回 12
})
.then(function(n) {
// 第三个 then，n = 12
console.log("最终结果: " + n);
return n;
});

console.log("Promise 链: " + promise);

运行结果：

```
最终结果: 12
Promise 链: [object Promise]

```

链式调用：每个 then 会返回一个新的 Promise，这允许我们按顺序执行多个异步操作。

### Promise.all

Promise.all 等待所有 Promise 完成，返回一个包含所有结果的数组。

### 实例

// 创建三个 Promise
var p1 = Promise.resolve(1);
var p2 = Promise.resolve(2);
var p3 = Promise.resolve(3);

// Promise.all 等待所有 Promise 完成
// 返回一个数组，包含所有 Promise 的结果
Promise.all([p1, p2, p3]).then(function(results) {
console.log("全部完成: " + results);
// 计算总和
console.log("总和: " + results.reduce(function(a, b) { return a + b; }, 0));
});

运行结果：

```
全部完成: 1,2,3
总和: 6

```

注意：如果任何一个 Promise 失败，Promise.all 会立即 rejection，不会等待其他 Promise 完成。

### Promise.race

Promise.race 返回最先完成（无论成功或失败）的 Promise 的结果。

### 实例

// 创建三个不同延迟的 Promise
var p1 = new Promise(function(resolve) {
setTimeout(function() { resolve("p1"); }, 100);
});
var p2 = new Promise(function(resolve) {
setTimeout(function() { resolve("p2"); }, 50);
});
var p3 = new Promise(function(resolve) {
setTimeout(function() { resolve("p3"); }, 30);
});

// Promise.race 返回最先完成的 Promise 的结果
Promise.race([p1, p2, p3]).then(function(value) {
console.log("最先完成: " + value);
});

运行结果：

```
最先完成: p3

```

应用场景：Promise.race 常用于实现超时功能：把一个长时间操作的 Promise 和一个超时 Promise 竞争。

### Promise.allSettled

Promise.allSettled 等待所有 Promise 结束（无论成功或失败），返回每个 Promise 的状态和结果。

### 实例

// 创建三个 Promise，其中一个会失败
var p1 = Promise.resolve("成功");
var p2 = Promise.reject(new Error("失败"));
var p3 = Promise.resolve("完成");

// Promise.allSettled 等待所有 Promise 结束
// 返回每个 Promise 的状态和值/reason
Promise.allSettled([p1, p2, p3]).then(function(results) {
results.forEach(function(result, index) {
if (result.status === "fulfilled") {
console.log("Promise " + index + ": " + result.value);
} else {
console.log("Promise " + index + ": " + result.reason.message);
}
});
});

运行结果：

```
Promise 0: 成功
Promise 1: 失败
Promise 2: 完成

```

区别：Promise.all 会在第一个失败时立即停止；Promise.allSettled 会等待所有 Promise 结束。

### Promise 类型注解

TypeScript 的泛型支持让 Promise 的类型声明变得精确。

### 实例

// 定义返回 Promise 的函数
// Promise<{ name: string; age: number }> 指定了返回的用户对象类型
function getUser(): Promise<{ name: string; age: number }> {
return Promise.resolve({ name: "Alice", age: 25 });
}

// async 函数：隐式返回 Promise
async function main() {
// await 会自动推断 user 的类型
var user = await getUser();
console.log("用户: " + JSON.stringify(user));
}

main();

类型推断：TypeScript 会根据泛型参数自动推断 Promise 的返回类型，这让我们在 async/await 中也能获得完整的类型提示。

### 注意事项

- 泛型参数：始终为 Promise 指定泛型参数，明确返回类型
- 错误处理：记得使用 catch 处理 Promise 失败的情况
- all vs allSettled：需要全部结果时用 allSettled，需要快速失败时用 all
- async/await：现代代码推荐使用 async/await，语法更简洁

最佳实践：优先使用 async/await 语法，它本质上还是基于 Promise，但写起来像同步代码。

### 总结

Promise 是 TypeScript 异步编程的核心。

- Promise：异步操作容器，有 pending/fulfilled/rejected 三种状态
- then/catch：链式处理异步结果
- Promise.all：等待全部完成，任一失败则整体失败
- Promise.race：返回最先完成的结果
- Promise.allSettled：等待全部结束，返回每个的状态

建议：使用 async/await 语法配合 Promise，让异步代码既类型安全又易于阅读。 x

---

## TypeScript Number

Source: https://www.runoob.com/typescript/ts-number.html

## TypeScript Number

TypeScript 与 JavaScript 类似，支持 Number 对象。 在 TypeScript 中，Number 对象用于包装数值类型。

Number 对象是原始数值的包装对象。

类似于 String 对象，Number 对象是引用类型，与基本的 number 类型有所不同。

尽管 Number 对象提供了一些额外的属性和方法，但在 TypeScript 中更推荐直接使用基本的 number 类型，因为 Number 对象会带来性能开销和类型混淆。

#### 语法

```
var num = new Number(value);
```

需要注意的是，这会创建一个引用类型的对象，而非基本的 number 类型。

注意： 如果一个参数值不能转换为一个数字将返回 NaN (非数字值)。

#### Number 对象与基本 number 类型的区别

- 基本类型 `number`：原始数据类型，用于存储数值。
- `Number` 对象：引用类型，是一个包装对象，用于包装基本数值。

### 实例

let numLiteral: number = 42;
let numObject: Number = new Number(42);

console.log(typeof numLiteral); // 输出："number"
console.log(typeof numObject); // 输出："object"

#### Number 对象属性

下表列出了 Number 对象支持的属性：

序号 属性 & 描述 1.

MAX_VALUE

可表示的最大的数，MAX_VALUE 属性值接近于 1.79E+308。大于 MAX_VALUE 的值代表 "Infinity"。 2.

MIN_VALUE

可表示的最小的数，即最接近 0 的正数 (实际上不会变成 0)。最大的负数是 -MIN_VALUE，MIN_VALUE 的值约为 5e-324。小于 MIN_VALUE ("underflow values") 的值将会转换为 0。 3.

NaN

非数字值（Not-A-Number）。 4.

NEGATIVE_INFINITY

负无穷大，溢出时返回该值。该值小于 MIN_VALUE。 5.

POSITIVE_INFINITY

正无穷大，溢出时返回该值。该值大于 MAX_VALUE。 6.

prototype

Number 对象的静态属性。使您有能力向对象添加属性和方法。 7.

constructor

返回对创建此对象的 Number 函数的引用。

### TypeScript

console.log("TypeScript Number 属性: "); console.log("最大值为: " + Number.MAX_VALUE); console.log("最小值为: " + Number.MIN_VALUE); console.log("负无穷大: " + Number.NEGATIVE_INFINITY); console.log("正无穷大:" + Number.POSITIVE_INFINITY);

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

console.log("TypeScript Number 属性: "); console.log("最大值为: " + Number.MAX_VALUE); console.log("最小值为: " + Number.MIN_VALUE); console.log("负无穷大: " + Number.NEGATIVE_INFINITY); console.log("正无穷大:" + Number.POSITIVE_INFINITY);

输出结果为：

```

TypeScript Number 属性:
最大值为: 1.7976931348623157e+308
最小值为: 5e-324
负无穷大: -Infinity
正无穷大:Infinity

```

#### NaN 实例

### TypeScript

var month = 0 if( month<=0 || month >12) { month = Number.NaN console.log("月份是："+ month) } else { console.log("输入月份数值正确。") }

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var month = 0; if (month <= 0 || month > 12) { month = Number.NaN; console.log("月份是：" + month); } else { console.log("输入月份数值正确。"); }

输出结果为：

```

月份是：NaN

```

#### prototype 实例

### TypeScript

function employee(id:number,name:string) { this.id = id this.name = name } var emp = new employee(123,"admin") employee.prototype.email = "admin@runoob.com" console.log("员工号: "+emp.id) console.log("员工姓名: "+emp.name) console.log("员工邮箱: "+emp.email)

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

function employee(id, name) { this.id = id; this.name = name; } var emp = new employee(123, "admin"); employee.prototype.email = "admin@runoob.com"; console.log("员工号: " + emp.id); console.log("员工姓名: " + emp.name); console.log("员工邮箱: " + emp.email);

输出结果为：

```

员工号: 123
员工姓名: admin
员工邮箱: admin@runoob.com

```

### Number 对象方法

Number对象 支持以下方法：

序号 方法 & 描述 实例 1. toExponential()

把对象的值转换为指数计数法。

```

//toExponential()
var num1 = 1225.30
var val = num1.toExponential();
console.log(val) // 输出： 1.2253e+3

```
2. toFixed()

把数字转换为字符串，并对小数点指定位数。

```

var num3 = 177.234
console.log("num3.toFixed() 为 "+num3.toFixed()) // 输出：177
console.log("num3.toFixed(2) 为 "+num3.toFixed(2)) // 输出：177.23
console.log("num3.toFixed(6) 为 "+num3.toFixed(6)) // 输出：177.234000

```
3. toLocaleString()

把数字转换为字符串，使用本地数字格式顺序。

```

var num = new Number(177.1234);
console.log( num.toLocaleString()); // 输出：177.1234

```
4. toPrecision()

把数字格式化为指定的长度。

```

var num = new Number(7.123456);
console.log(num.toPrecision()); // 输出：7.123456
console.log(num.toPrecision(1)); // 输出：7
console.log(num.toPrecision(2)); // 输出：7.1

```
5. toString()

把数字转换为字符串，使用指定的基数。数字的基数是 2 ~ 36 之间的整数。若省略该参数，则使用基数 10。

```

var num = new Number(10);
console.log(num.toString()); // 输出10进制：10
console.log(num.toString(2)); // 输出2进制：1010
console.log(num.toString(8)); // 输出8进制：12

```
6. valueOf()

返回一个 Number 对象的原始数字值。

```

var num = new Number(10);
console.log(num.valueOf()); // 输出：10

```

#### Number 对象的使用建议

在 TypeScript 中，通常更推荐使用基本的 number 类型，而不是 Number 对象。原因如下：

- 性能：基本类型更轻量，性能更好。
- 类型一致性：TypeScript 的类型系统更倾向于基本类型，使用 `Number` 对象可能导致意外的类型不匹配。
- 最佳实践：基本类型的 `number` 更符合 TypeScript 的最佳实践，避免了对象包装带来的不必要复杂性。

如果确实需要使用 Number 对象的特定方法，可以通过 valueOf() 方法将 Number 对象转换为基本的 number 类型。

总之，TypeScript 更推荐使用基本类型 number 而不是 Number 对象，以保持代码的简洁、高效和一致性。

### 实例

let numLiteral: number = 123.456;
let numObject: Number = new Number(123.456);

console.log(numLiteral.toFixed(2)); // 输出："123.46"
console.log(numObject.valueOf()); // 输出：123.456

---

## TypeScript String（字符串）

Source: https://www.runoob.com/typescript/ts-string.html

## TypeScript String（字符串）

在 TypeScript 中，字符串（String）是用于表示文本数据的基本数据类型，它继承并扩展了 JavaScript 的字符串特性，同时增加了静态类型校验。

#### 语法

```

var txt = new String("string");
```

或者更简单方式：

```

var txt = "string";

```

#### 两种创建方式及核心区别

TypeScript 中创建字符串有两种方式，但二者在类型、性能和使用场景上有本质区别：

方式 1：字符串字面量（推荐）

这是 TypeScript/JavaScript 中最常用、性能最优的方式，创建的是原始字符串类型（string），也是 TypeScript 类型系统中默认的字符串类型。

```

// 原始字符串（推荐使用）
const txt1: string = "Hello TypeScript"; // 显式指定类型
const txt2 = "Hello JavaScript"; // 类型推导为 string
```

方式 2：String 对象（不推荐）

通过 new String() 创建的是包装对象类型（String），本质是一个对象，而非原始值，会带来类型混淆和性能损耗。

```

// String 对象（不推荐）
const txtObj: String = new String("Hello Object"); // 类型为 String（对象）
```

String 对象和字符串字面量在类型上是不同的：

- 字符串字面量是基本数据类型 `string`，用于直接存储字符串值。
- `String` 对象是 `String` 类型，实际上是一个对象，而非原始的字符串值。

### 实例

let strLiteral: string = "Hello";
let strObject: String = new String("Hello");

console.log(typeof strLiteral); // 输出："string"
console.log(typeof strObject); // 输出："object"

特性 字符串字面量（string） String 对象（String） 类型本质 原始值 引用类型（对象） 性能 高效，无额外内存开销 低效，创建对象实例 类型校验（TypeScript） 符合 TS 基础类型规范 类型不匹配（如 string 类型变量无法赋值 String 对象） 比较方式 直接比较值 比较引用地址（需用 valueOf() 取原始值）

#### 字符串字面量和 String 对象的类型兼容性

在 TypeScript 中，string 字面量类型和 String 对象类型不完全兼容。

例如，string 类型的变量无法直接使用 String 对象的方法，反之亦然。因此，通常情况下不需要使用 String 对象。

### 实例

let strLiteral: string = "Test";
let strObject: String = new String("Test");

console.log(strLiteral === strObject); // 输出：false，内容相同，类型不同
console.log(strLiteral == strObject); // 输出：true，内容相同
console.log(strLiteral === strObject.valueOf()); // 输出：true，将对象转为原始字符串后比较

strLiteral 是原始字符串类型（string），而 strObject 是 String 对象类型（String）。这意味着它们的类型不同。

以上代码转为 JavaScript 代码为：

### 实例

var strLiteral = "Test";
var strObject = new String("Test");
console.log(strLiteral === strObject); // 输出：false，内容相同，类型不同
console.log(strLiteral == strObject); // 输出：true，内容相同
console.log(strLiteral === strObject.valueOf()); // 输出：true，将对象转为原始字符串后比较

#### String 对象属性

下表列出了 String 对象支持的属性：

序号 属性 & 描述 实例 1. constructor

对创建该对象的函数的引用。

```
var str = new String( "This is string" );
console.log("str.constructor is:" + str.constructor)
```

输出结果：

```
str.constructor is:function String() { [native code] }
```
2. length

返回字符串的长度。

```
var uname = new String("Hello World")
console.log("Length "+uname.length) // 输出 11
```
3. prototype

允许您向对象添加属性和方法。

```
function employee(id:number,name:string) {
this.id = id
this.name = name
}
var emp = new employee(123,"admin")
employee.prototype.email="admin@runoob.com" // 添加属性 email
console.log("员工号: "+emp.id)
console.log("员工姓名: "+emp.name)
console.log("员工邮箱: "+emp.email)
```

#### String 方法

下表列出了 String 对象支持的方法：

序号 方法 & 描述 实例 1. charAt()

返回在指定位置的字符。

```

var str = new String("RUNOOB");
console.log("str.charAt(0) 为:" + str.charAt(0)); // R
console.log("str.charAt(1) 为:" + str.charAt(1)); // U
console.log("str.charAt(2) 为:" + str.charAt(2)); // N
console.log("str.charAt(3) 为:" + str.charAt(3)); // O
console.log("str.charAt(4) 为:" + str.charAt(4)); // O
console.log("str.charAt(5) 为:" + str.charAt(5)); // B

```
2. charCodeAt()

返回在指定的位置的字符的 Unicode 编码。

```

var str = new String("RUNOOB");
console.log("str.charCodeAt(0) 为:" + str.charCodeAt(0)); // 82
console.log("str.charCodeAt(1) 为:" + str.charCodeAt(1)); // 85
console.log("str.charCodeAt(2) 为:" + str.charCodeAt(2)); // 78
console.log("str.charCodeAt(3) 为:" + str.charCodeAt(3)); // 79
console.log("str.charCodeAt(4) 为:" + str.charCodeAt(4)); // 79
console.log("str.charCodeAt(5) 为:" + str.charCodeAt(5)); // 66

```
3. concat()

连接两个或更多字符串，并返回新的字符串。

```

var str1 = new String( "RUNOOB" );
var str2 = new String( "GOOGLE" );
var str3 = str1.concat( str2 );
console.log("str1 + str2 : "+str3) // RUNOOBGOOGLE

```
4. indexOf()

返回某个指定的字符串值在字符串中首次出现的位置。

```

var str1 = new String( "RUNOOB" );

var index = str1.indexOf( "OO" );
console.log("查找的字符串位置 :" + index ); // 3

```
5. lastIndexOf()

从后向前搜索字符串，并从起始位置（0）开始计算返回字符串最后出现的位置。

```

var str1 = new String( "This is string one and again string" );
var index = str1.lastIndexOf( "string" );
console.log("lastIndexOf 查找到的最后字符串位置 :" + index ); // 29

index = str1.lastIndexOf( "one" );
console.log("lastIndexOf 查找到的最后字符串位置 :" + index ); // 15

```
6. localeCompare()

用本地特定的顺序来比较两个字符串。

```

var str1 = new String( "This is beautiful string" );

var index = str1.localeCompare( "This is beautiful string");

console.log("localeCompare first :" + index ); // 0

```
7.

match()

查找找到一个或多个正则表达式的匹配。

```

var str="The rain in SPAIN stays mainly in the plain";
var n=str.match(/ain/g); // ain,ain,ain

```
8. replace()

替换与正则表达式匹配的子串

```

var re = /(\w+)\s(\w+)/;
var str = "zara ali";
var newstr = str.replace(re, "$2, $1");
console.log(newstr); // ali, zara

```
9. search()

检索与正则表达式相匹配的值

```

var re = /apples/gi;
var str = "Apples are round, and apples are juicy.";
if (str.search(re) == -1 ) {
console.log("Does not contain Apples" );
} else {
console.log("Contains Apples" );
}

```
10. slice()

提取字符串的片断，并在新的字符串中返回被提取的部分。 11. split()

把字符串分割为子字符串数组。

```

var str = "Apples are round, and apples are juicy.";
var splitted = str.split(" ", 3);
console.log(splitted) // [ 'Apples', 'are', 'round,' ]

```
12. substr()

从起始索引号提取字符串中指定数目的字符。 13. substring()

提取字符串中两个指定的索引号之间的字符。

```

var str = "RUNOOB GOOGLE TAOBAO FACEBOOK";
console.log("(1,2): " + str.substring(1,2)); // U
console.log("(0,10): " + str.substring(0, 10)); // RUNOOB GOO
console.log("(5): " + str.substring(5)); // B GOOGLE TAOBAO FACEBOOK

```
14. toLocaleLowerCase()

根据主机的语言环境把字符串转换为小写，只有几种语言（如土耳其语）具有地方特有的大小写映射。

```

var str = "Runoob Google";
console.log(str.toLocaleLowerCase( )); // runoob google

```
15. toLocaleUpperCase()

据主机的语言环境把字符串转换为大写，只有几种语言（如土耳其语）具有地方特有的大小写映射。

```

var str = "Runoob Google";
console.log(str.toLocaleUpperCase( )); // RUNOOB GOOGLE

```
16. toLowerCase()

把字符串转换为小写。

```

var str = "Runoob Google";
console.log(str.toLowerCase( )); // runoob google

```
17. toString()

返回字符串。

```

var str = "Runoob";
console.log(str.toString( )); // Runoob

```
18. toUpperCase()

把字符串转换为大写。

```

var str = "Runoob Google";
console.log(str.toUpperCase( )); // RUNOOB GOOGLE

```
19. valueOf()

返回指定字符串对象的原始值。

```

var str = new String("Runoob");
console.log(str.valueOf( )); // Runoob

```

#### String 对象的使用建议

在 TypeScript 中，使用 String 对象通常是不必要的，直接使用 string 字面量会更高效且符合 TypeScript 的最佳实践：

- 性能：`String` 对象是一个引用类型，会占用更多内存，且每次创建一个新对象性能开销更大。
- 类型安全：TypeScript 更鼓励使用 `string` 字面量类型，保持代码的简洁和一致性。

如果确实需要使用 String 对象的方法，可以通过 valueOf() 方法将对象转为原始字符串，然后继续处理。

通常情况下，TypeScript 推荐直接使用 string 字面量类型，以简化代码，提高性能，避免不必要的类型转换和复杂性。

### 实例

let strLiteral: string = "Use string literals whenever possible!";
let strObject: String = new String("Avoid using String objects.");

console.log(strLiteral); // 输出："Use string literals whenever possible!"
console.log(strObject.valueOf()); // 输出："Avoid using String objects."

---

## TypeScript Array(数组)

Source: https://www.runoob.com/typescript/ts-array.html

## TypeScript Array(数组)

数组对象是使用单独的变量名来存储一系列的值。

数组非常常用。

假如你有一组数据（例如：网站名字），存在单独变量如下所示：

var site1="Google"; var site2="Runoob"; var site3="Taobao";

如果有 10 个、100 个这种方式就变的很不实用，这时我们可以使用数组来解决：

var sites:string[]; sites = ["Google","Runoob","Taobao"]

这样看起来就简洁多了。

TypeScript 声明数组的语法格式如下所示：

```

var array_name[:datatype]; //声明
array_name = [val1,val2,valn..] //初始化

```

或者直接在声明时初始化：

```
var array_name[:datatype] = [val1,val2…valn]
```

如果数组声明时未设置类型，则会被认为是 any 类型，在初始化时根据第一个元素的类型来推断数组的类型。

#### 实例

创建一个 number 类型的数组：

```
var numlist:number[] = [2,4,6,8]
```

整个数组结构如下所示：

索引值第一个为 0，我们可以根据索引值来访问数组元素：

### TypeScript

var sites:string[]; sites = ["Google","Runoob","Taobao"] console.log(sites[0]); console.log(sites[1]);

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var sites; sites = ["Google", "Runoob", "Taobao"]; console.log(sites[0]); console.log(sites[1]);

输出结果为：

```
Google
Runoob
```

以下实例我们在声明时直接初始化：

### TypeScript

var nums:number[] = [1,2,3,4] console.log(nums[0]); console.log(nums[1]); console.log(nums[2]); console.log(nums[3]);

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var nums = [1, 2, 3, 4]; console.log(nums[0]); console.log(nums[1]); console.log(nums[2]); console.log(nums[3]);

输出结果为：

```
1
2
3
4
```

### Array 对象

我们也可以使用 Array 对象创建数组。

Array 对象的构造函数接受以下两种值：

- 表示数组大小的数值。
- 初始化的数组列表，元素使用逗号分隔值。

#### 实例

指定数组初始化大小：

### TypeScript

var arr_names:number[] = new Array(4) for(var i = 0; i<arr_names.length; i++) { arr_names[i] = i * 2 console.log(arr_names[i]) }

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var arr_names = new Array(4); for (var i = 0; i < arr_names.length; i++) { arr_names[i] = i * 2; console.log(arr_names[i]); }

输出结果为：

```
0
2
4
6
```

以下实例我们直接初始化数组元素：

### TypeScript

var sites:string[] = new Array("Google","Runoob","Taobao","Facebook") for(var i = 0;i<sites.length;i++) { console.log(sites[i]) }

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var sites = new Array("Google", "Runoob", "Taobao", "Facebook"); for (var i = 0; i < sites.length; i++) { console.log(sites[i]); }

输出结果为：

```
Google
Runoob
Taobao
Facebook
```

### 数组解构

我们也可以把数组元素赋值给变量，如下所示：

### TypeScript

var arr:number[] = [12,13] var[x,y] = arr // 将数组的两个元素赋值给变量 x 和 y console.log(x) console.log(y)

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var arr = [12, 13]; var x = arr[0], y = arr[1]; // 将数组的两个元素赋值给变量 x 和 y console.log(x); console.log(y);

输出结果为：

```
12
13
```

### 数组迭代

我们可以使用 for 语句来循环输出数组的各个元素：

### TypeScript

var j:any; var nums:number[] = [1001,1002,1003,1004] for(j in nums) { console.log(nums[j]) }

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var j; var nums = [1001, 1002, 1003, 1004]; for (j in nums) { console.log(nums[j]); }

输出结果为：

```
1001
1002
1003
1004
```

### 多维数组

一个数组的元素可以是另外一个数组，这样就构成了多维数组（Multi-dimensional Array）。

最简单的多维数组是二维数组，定义方式如下：

```
var arr_name:datatype[][]=[ [val1,val2,val3],[v1,v2,v3] ]
```

#### 实例

定义一个二维数组，每一个维度的数组有三个元素。

### TypeScript

var multi:number[][] = [[1,2,3],[23,24,25]] console.log(multi[0][0]) console.log(multi[0][1]) console.log(multi[0][2]) console.log(multi[1][0]) console.log(multi[1][1]) console.log(multi[1][2])

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var multi = [[1, 2, 3], [23, 24, 25]]; console.log(multi[0][0]); console.log(multi[0][1]); console.log(multi[0][2]); console.log(multi[1][0]); console.log(multi[1][1]); console.log(multi[1][2]);

输出结果为：

```
1
2
3
23
24
25
```

### 数组在函数中的使用

#### 作为参数传递给函数

### TypeScript

var sites:string[] = new Array("Google","Runoob","Taobao","Facebook") function disp(arr_sites:string[]) { for(var i = 0;i<arr_sites.length;i++) { console.log(arr_sites[i]) } } disp(sites);

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var sites = new Array("Google", "Runoob", "Taobao", "Facebook"); function disp(arr_sites) { for (var i = 0; i < arr_sites.length; i++) { console.log(arr_sites[i]); } } disp(sites);

输出结果为：

```
Google
Runoob
Taobao
Facebook
```

#### 作为函数的返回值

### TypeScript

function disp():string[] { return new Array("Google", "Runoob", "Taobao", "Facebook"); } var sites:string[] = disp() for(var i in sites) { console.log(sites[i]) }

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

function disp() { return new Array("Google", "Runoob", "Taobao", "Facebook"); } var sites = disp(); for (var i in sites) { console.log(sites[i]); }

输出结果为：

```
Google
Runoob
Taobao
Facebook
```

### 数组方法

下表列出了一些常用的数组方法：

序号 方法 & 描述 实例 1. concat()

连接两个或更多的数组，并返回结果。

```

var alpha = ["a", "b", "c"];
var numeric = [1, 2, 3];

var alphaNumeric = alpha.concat(numeric);
console.log("alphaNumeric : " + alphaNumeric ); // a,b,c,1,2,3

```
2. every()

检测数值元素的每个元素是否都符合条件。

```

function isBigEnough(element, index, array) {
return (element >= 10);
}

var passed = [12, 5, 8, 130, 44].every(isBigEnough);
console.log("Test Value : " + passed ); // false

```
3. filter()

检测数值元素，并返回符合条件所有元素的数组。

```

function isBigEnough(element, index, array) {
return (element >= 10);
}

var passed = [12, 5, 8, 130, 44].filter(isBigEnough);
console.log("Test Value : " + passed ); // 12,130,44

```
4. forEach()

数组每个元素都执行一次回调函数。

```

let num = [7, 8, 9];
num.forEach(function (value) {
console.log(value);
});

```

编译成 JavaScript 代码：

```

var num = [7, 8, 9];
num.forEach(function (value) {
console.log(value); // 7 8 9
});
```
5. indexOf()

搜索数组中的元素，并返回它所在的位置。

如果搜索不到，返回值 -1，代表没有此项。

```

var index = [12, 5, 8, 130, 44].indexOf(8);
console.log("index is : " + index ); // 2

```
6. join()

把数组的所有元素放入一个字符串。

```

var arr = new Array("Google","Runoob","Taobao");

var str = arr.join();
console.log("str : " + str ); // Google,Runoob,Taobao

var str = arr.join(", ");
console.log("str : " + str ); // Google, Runoob, Taobao

var str = arr.join(" + ");
console.log("str : " + str ); // Google + Runoob + Taobao

```
7. lastIndexOf()

返回一个指定的字符串值最后出现的位置，在一个字符串中的指定位置从后向前搜索。

```

var index = [12, 5, 8, 130, 44].lastIndexOf(8);
console.log("index is : " + index ); // 2

```
8. map()

通过指定函数处理数组的每个元素，并返回处理后的数组。

```

var numbers = [1, 4, 9];
var roots = numbers.map(Math.sqrt);
console.log("roots is : " + roots ); // 1,2,3

```
9. pop()

删除数组的最后一个元素并返回删除的元素。

```

var numbers = [1, 4, 9];

var element = numbers.pop();
console.log("element is : " + element ); // 9

var element = numbers.pop();
console.log("element is : " + element ); // 4

```
10. push()

向数组的末尾添加一个或更多元素，并返回新的长度。

```

var numbers = new Array(1, 4, 9);
var length = numbers.push(10);
console.log("new numbers is : " + numbers ); // 1,4,9,10
length = numbers.push(20);
console.log("new numbers is : " + numbers ); // 1,4,9,10,20

```
11. reduce()

将数组元素计算为一个值（从左到右）。

```

var total = [0, 1, 2, 3].reduce(function(a, b){ return a + b; });
console.log("total is : " + total ); // 6

```
12. reduceRight()

将数组元素计算为一个值（从右到左）。

```

var total = [0, 1, 2, 3].reduceRight(function(a, b){ return a + b; });
console.log("total is : " + total ); // 6

```
13. reverse()

反转数组的元素顺序。

```

var arr = [0, 1, 2, 3].reverse();
console.log("Reversed array is : " + arr ); // 3,2,1,0

```
14. shift()

删除并返回数组的第一个元素。

```

var arr = [10, 1, 2, 3].shift();
console.log("Shifted value is : " + arr ); // 10

```
15. slice()

选取数组的的一部分，并返回一个新数组。

```

var arr = ["orange", "mango", "banana", "sugar", "tea"];
console.log("arr.slice( 1, 2) : " + arr.slice( 1, 2) ); // mango
console.log("arr.slice( 1, 3) : " + arr.slice( 1, 3) ); // mango,banana

```
16. some()

检测数组元素中是否有元素符合指定条件。

```

function isBigEnough(element, index, array) {
return (element >= 10);

}

var retval = [2, 5, 8, 1, 4].some(isBigEnough);
console.log("Returned value is : " + retval ); // false

var retval = [12, 5, 8, 1, 4].some(isBigEnough);
console.log("Returned value is : " + retval ); // true

```
17. sort()

对数组的元素进行排序。

```

var arr = new Array("orange", "mango", "banana", "sugar");
var sorted = arr.sort();
console.log("Returned string is : " + sorted ); // banana,mango,orange,sugar

```
18. splice()

从数组中添加或删除元素。

```

var arr = ["orange", "mango", "banana", "sugar", "tea"];
var removed = arr.splice(2, 0, "water");
console.log("After adding 1: " + arr ); // orange,mango,water,banana,sugar,tea
console.log("removed is: " + removed);

removed = arr.splice(3, 1);
console.log("After removing 1: " + arr ); // orange,mango,water,sugar,tea
console.log("removed is: " + removed); // banana

```
19. toString()

把数组转换为字符串，并返回结果。

```

var arr = new Array("orange", "mango", "banana", "sugar");
var str = arr.toString();
console.log("Returned string is : " + str ); // orange,mango,banana,sugar

```
20. unshift()

向数组的开头添加一个或更多元素，并返回新的长度。

```

var arr = new Array("orange", "mango", "banana", "sugar");
var length = arr.unshift("water");
console.log("Returned array is : " + arr ); // water,orange,mango,banana,sugar
console.log("Length of the array is : " + length ); // 5

```

---

## TypeScript Map 对象

Source: https://www.runoob.com/typescript/ts-map.html

## TypeScript Map 对象

Map 对象保存键值对，并且能够记住键的原始插入顺序。

任何值(对象或者原始值) 都可以作为一个键或一个值。

Map 是 ES6 中引入的一种新的数据结构，可以参考 ES6 Map 与 Set。

### 创建 Map

TypeScript 使用 Map 类型和 new 关键字来创建 Map：

```
let myMap = new Map();
```

初始化 Map，可以以数组的格式来传入键值对：

```
let myMap = new Map([
["key1", "value1"],
["key2", "value2"]
]);
```

Map 相关的函数与属性：

- map.clear() – 移除 Map 对象的所有键/值对 。
- map.set() – 设置键值对，返回该 Map 对象。
- map.get() – 返回键对应的值，如果不存在，则返回 undefined。
- map.has() – 返回一个布尔值，用于判断 Map 中是否包含键对应的值。
- map.delete() – 删除 Map 中的元素，删除成功返回 true，失败返回 false。
- map.size – 返回 Map 对象键/值对的数量。
- map.keys() - 返回一个 Iterator 对象， 包含了 Map 对象中每个元素的键 。
- map.values() – 返回一个新的Iterator对象，包含了Map对象中每个元素的值 。
- map.entries() – 返回一个包含 Map 中所有键值对的迭代器 。

#### 常用函数

set(key: K, value: V): this - 向 Map 中添加或更新键值对。

```
map.set('key1', 'value1');

```

get(key: K): V | undefined - 根据键获取值，如果键不存在则返回 undefined。

```

const value = map.get('key1');
```

has(key: K): boolean - 检查 Map 中是否存在指定的键。

```

const exists = map.has('key1');
```

delete(key: K): boolean - 删除指定键的键值对，成功删除返回 true，否则返回 false。

```
const removed = map.delete('key1');
```

clear(): void - 清空 Map 中的所有键值对。

```

map.clear();
```

size: number - 返回 Map 中键值对的数量。

```

const size = map.size;
```

#### 迭代方法 keys(): IterableIterator<K> - 返回一个包含 Map 中所有键的迭代器。

```

for (const key of map.keys()) {
console.log(key);
}
```

values(): IterableIterator<V> - 返回一个包含 Map 中所有值的迭代器。

```

for (const value of map.values()) {
console.log(value);
}
```

entries(): IterableIterator<[K, V]> - 返回一个包含 Map 中所有键值对的迭代器，每个元素是一个 [key, value] 数组。

```

for (const [key, value] of map.entries()) {
console.log(key, value);
}

```

forEach(callbackfn: (value: V, key: K, map: Map<K, V>) => void, thisArg?: any): void - 对 Map 中的每个键值对执行一次提供的回调函数。

```

map.forEach((value, key) => {
console.log(key, value);
});
```

#### 实例

### 实例

const map = new Map<string, number>();

map.set('one', 1);
map.set('two', 2);

console.log(map.get('one')); // 输出: 1

console.log(map.has('two')); // 输出: true

map.delete('one');

console.log(map.size); // 输出: 1

map.forEach((value, key) => {
console.log(key, value); // 输出: two 2
});

map.clear();

console.log(map.size); // 输出: 0

### 实例 - test.ts 文件

let nameSiteMapping = new Map(); // 设置 Map 对象 nameSiteMapping.set("Google", 1); nameSiteMapping.set("Runoob", 2); nameSiteMapping.set("Taobao", 3); // 获取键对应的值 console.log(nameSiteMapping.get("Runoob")); // 2 // 判断 Map 中是否包含键对应的值 console.log(nameSiteMapping.has("Taobao")); // true console.log(nameSiteMapping.has("Zhihu")); // false // 返回 Map 对象键/值对的数量 console.log(nameSiteMapping.size); // 3 // 删除 Runoob console.log(nameSiteMapping.delete("Runoob")); // true console.log(nameSiteMapping); // 移除 Map 对象的所有键/值对 nameSiteMapping.clear(); // 清除 Map console.log(nameSiteMapping);

使用 es6 编译：

```
tsc --target es6 test.ts
```

编译以上代码得到如下 JavaScript 代码：

### 实例 - test.js 文件

let nameSiteMapping = new Map(); // 设置 Map 对象 nameSiteMapping.set("Google", 1); nameSiteMapping.set("Runoob", 2); nameSiteMapping.set("Taobao", 3); // 获取键对应的值 console.log(nameSiteMapping.get("Runoob")); //40 // 判断 Map 中是否包含键对应的值 console.log(nameSiteMapping.has("Taobao")); //true console.log(nameSiteMapping.has("Zhihu")); //false // 返回 Map 对象键/值对的数量 console.log(nameSiteMapping.size); //3 // 删除 Runoob console.log(nameSiteMapping.delete("Runoob")); // true console.log(nameSiteMapping); // 移除 Map 对象的所有键/值对 nameSiteMapping.clear(); //清除 Map console.log(nameSiteMapping);

执行以上 JavaScript 代码，输出结果为：

```
2
true
false
3
true
Map { 'Google' => 1, 'Taobao' => 3 }
Map {}
```

#### 迭代 Map

Map 对象中的元素是按顺序插入的，我们可以迭代 Map 对象，每一次迭代返回 [key, value] 数组。

TypeScript使用 for...of 来实现迭代：

### 实例 -test.ts 文件

let nameSiteMapping = new Map(); nameSiteMapping.set("Google", 1); nameSiteMapping.set("Runoob", 2); nameSiteMapping.set("Taobao", 3); // 迭代 Map 中的 key for (let key of nameSiteMapping.keys()) { console.log(key); } // 迭代 Map 中的 value for (let value of nameSiteMapping.values()) { console.log(value); } // 迭代 Map 中的 key => value for (let entry of nameSiteMapping.entries()) { console.log(entry[0], entry[1]); } // 使用对象解析 for (let [key, value] of nameSiteMapping) { console.log(key, value); }

使用 es6 编译：

```
tsc --target es6 test.ts
```

编译以上代码得到如下 JavaScript 代码：

### 实例

let nameSiteMapping = new Map(); nameSiteMapping.set("Google", 1); nameSiteMapping.set("Runoob", 2); nameSiteMapping.set("Taobao", 3); // 迭代 Map 中的 key for (let key of nameSiteMapping.keys()) { console.log(key); } // 迭代 Map 中的 value for (let value of nameSiteMapping.values()) { console.log(value); } // 迭代 Map 中的 key => value for (let entry of nameSiteMapping.entries()) { console.log(entry[0], entry[1]); } // 使用对象解析 for (let [key, value] of nameSiteMapping) { console.log(key, value); }

执行以上 JavaScript 代码，输出结果为：

```
Google
Runoob
Taobao
1
2
3
Google 1
Runoob 2
Taobao 3
Google 1
Runoob 2
Taobao 3

```

---

## TypeScript 枚举（Enum）

Source: https://www.runoob.com/typescript/ts-enum.html

## TypeScript 枚举（Enum）

枚举（Enum）是 TypeScript 中非常有用的特性，它允许我们定义一组命名常量。枚举可以使得代码更易读、更易维护，可以用有意义的名称替代"魔术数字"。

### 数字枚举

默认情况下，枚举从 0 开始编号。

### 实例

enum Direction {
Up, // 0
Down, // 1
Left, // 2
Right // 3
}

var dir: Direction = Direction.Up;
console.log("方向: " + dir);
console.log("方向名称: " + Direction[0]);

运行结果：

```
方向: 0
方向名称: Up

```

### 手动赋值

可以手动为枚举成员指定值。

### 实例

enum Status {
Success = 1,
Error = 2,
Pending = 3
}

console.log("状态: " + Status.Success);
console.log("状态名称: " + Status[1]);

运行结果：运

```
状态: 1
状态名称: Success

```

### 字符串枚举

字符串枚举每个成员都必须有字符串字面量值。

### 实例

enum Message {
Success = "SUCCESS",
Error = "ERROR",
Warning = "WARNING"
}

console.log("消息: " + Message.Success);

运行结果：

```
消息: SUCCESS

```

### 常量枚举

使用 `const` 修饰的枚举会在编译时内联，生成更优化的代码。

### 实例

const enum Color {
Red = "RED",
Green = "GREEN",
Blue = "BLUE"
}

var favoriteColor: Color = Color.Red;
console.log("喜欢的颜色: " + favoriteColor);

### 异构枚举

枚举可以混合数字和字符串值，但不推荐使用。

### 实例

enum BooleanLikeHeterogeneousEnum {
No = 0,
Yes = "YES"
}

console.log("值: " + BooleanLikeHeterogeneousEnum.No);
console.log("字符串值: " + BooleanLikeHeterogeneousEnum.Yes);

### 枚举成员类型

当枚举成员都是字面量值时，成员类型可以作为类型使用。

### 实例

enum ShapeKind {
Circle = "circle",
Square = "square"
}

interface Circle {
kind: ShapeKind.Circle;
radius: number;
}

interface Square {
kind: ShapeKind.Square;
sideLength: number;
}

var c: Circle = {
kind: ShapeKind.Circle,
radius: 10
};

console.log("圆形: " + JSON.stringify(c));

### 运行时常量枚举

普通枚举在运行时保留为真实对象。

### 实例

enum FileAccess {
Read = 1 << 1,
Write = 1 << 2,
ReadWrite = Read | Write
}

console.log("文件访问: " + FileAccess.ReadWrite);

运行结果：

```
文件访问: 6

```

### 总结

- 数字枚举：默认从 0 开始，可手动赋值
- 字符串枚举：每个成员必须是字符串字面量
- 常量枚举：使用 const，编译时内联
- 异构枚举：混合数字和字符串，不推荐
- 成员类型：字面量枚举成员可用作类型

---

## TypeScript 类型推断

Source: https://www.runoob.com/typescript/ts-inference.html

## TypeScript 类型推断

类型推断（Type Inference）是 TypeScript 最强大和便捷的特性之一。

它允许编译器自动分析代码的上下文，并根据变量的值、函数的返回值等自动推断出变量的类型。

这意味着开发者无需为每个变量显式声明类型，可以编写更简洁、更易于维护的代码。 TypeScript 类型推断过程 源代码 // 变量声明 var num = 10; var str = "hello"; function add(a, b) { return a + b; } 推断 推断结果 var num: number = 10; var str: string = "hello"; function add(a: number, b: number): number { return a + b; } 类型推断场景 基础类型推断 根据初始值推断类型 返回类型推断 根据 return 语句推断 上下文推断 根据使用位置推断

### 为什么需要类型推断

在 JavaScript 开发中，我们需要手动为每个变量声明类型。

这不仅增加了代码的冗余度，也降低了开发效率。

TypeScript 的类型推断功能可以在大多数情况下自动推断出变量的类型。

开发者只需要在类型复杂或不明确的情况下显式声明类型。

概念说明：类型推断是 TypeScript 编译器分析代码上下文，自动确定变量类型的过程。它基于初始值、函数返回值、参数类型等多个维度进行推断。

### 基础类型推断

当声明变量并初始化时，TypeScript 会根据初始值自动推断变量的类型。

这是最常见的类型推断场景。

### 实例

// 声明变量并初始化时，TypeScript 自动推断类型
// 初始值为数字 10，类型推断为 number
var num = 10;

// 初始值为字符串，类型推断为 string
var str = "hello";

// 初始值为布尔值，类型推断为 boolean
var isActive = true;

// 使用 typeof 验证推断结果
console.log("num 类型: " + typeof num);
console.log("str 类型: " + typeof str);
console.log("isActive 类型: " + typeof isActive);

运行结果：

```
num 类型: number
str 类型: string
isActive 类型: boolean

```

说明：TypeScript 使用 `var` 声明变量时，会根据初始值推断类型。如果初始值是数字，类型就是 number；如果是字符串，类型就是 string。

### 函数返回类型推断

TypeScript 会根据函数的 return 语句自动推断返回类型。

这使得函数调用时能够获得正确的类型提示。

### 实例

// TypeScript 推断返回类型为 number
// 因为返回的是 a + b 的结果
function add(a: number, b: number) {
return a + b;
}

// TypeScript 推断返回类型为 string
// 因为返回的是字符串拼接结果
function greet(name: string) {
return "Hello, " + name;
}

// 调用函数，返回值类型已被正确推断
var result = add(1, 2);
var message = greet("TypeScript");

console.log("加法结果: " + result);
console.log("问候语: " + message);

运行结果：

```
加法结果: 3
问候语: Hello, TypeScript

```

提示：虽然 TypeScript 可以推断返回类型，但在实际项目中，建议显式声明函数返回类型以提高代码可读性。

### 上下文类型推断

类型推断不仅基于变量本身，还会考虑变量使用的上下文环境。

例如，数组的 map 方法的回调函数参数类型会根据数组元素的类型自动推断。

### 实例

// 定义数字数组
var numbers = [1, 2, 3, 4, 5];

// 使用 map 方法
// 回调函数的参数 n 的类型会根据 numbers 数组的元素类型自动推断为 number
var doubled = numbers.map(function(n) {
// n 自动推断为 number 类型
return n * 2;
});

console.log("翻倍数组: " + doubled);

运行结果：

```
翻倍数组: 2,4,6,8,10

```

上下文推断：当变量在特定上下文中使用时（如回调函数、事件处理等），TypeScript 会根据上下文自动推断类型。这种推断叫做"上下文类型推断"。

### 最佳通用类型推断

当有多个候选类型时，TypeScript 会选择能够容纳所有候选类型的"最佳通用类型"。

这确保了类型的兼容性。

### 实例

// 混合数组包含 number 和 string
// TypeScript 推断为 (number | string)[]（联合类型）
var mixed = [1, "two", 3, "four"];

// 访问不同类型的元素
console.log("混合数组: " + mixed);
console.log("类型: " + typeof mixed[0] + ", " + typeof mixed[1]);

联合类型：当数组包含多种类型时，TypeScript 会推断为联合类型。例如 `(number | string)[]` 表示数组元素可以是 number 或 string。

### 类型推断的限制

虽然 TypeScript 的类型推断很强大，但在某些情况下可能无法正确推断类型。

这时需要显式声明类型。

### 实例

// 声明变量时不赋值，也没有初始值
// TypeScript 无法推断类型，会推断为 any
var unknown;

// 可以赋任何类型的值
unknown = "hello";
unknown = 123;

// 这种行为容易导致类型错误
console.log("未指定类型: " + unknown);

// 建议：显式指定类型以获得更好的类型安全
var fixedNumber: number = 42;

console.log("指定类型: " + fixedNumber);

运行结果：

```
未指定类型: 123
指定类型: 42

```

警告：避免声明变量时不赋值。最好在声明时提供初始值，或者显式声明类型，以充分利用 TypeScript 的类型检查功能。

### 泛型函数推断

泛型函数的类型参数会根据传入的参数自动推断。

这使得泛型函数既灵活又保持类型安全。

### 实例

// 泛型函数：接受任意类型参数，返回同类型值
// T 会根据传入的参数自动推断
function identity<T>(arg: T): T {
return arg;
}

// 传入字符串，T 推断为 string
var str = identity("hello");

// 传入数字，T 推断为 number
var num = identity(42);

// 传入对象，T 推断为对象类型
var obj = identity({ name: "TypeScript" });

console.log("字符串: " + str);
console.log("数字: " + num);
console.log("对象: " + JSON.stringify(obj));

泛型推断：TypeScript 会根据调用时传入的参数类型自动推断泛型参数的类型。这使得同一个函数可以处理不同类型的输入，同时保持类型安全。

### 注意事项

- 初始值的重要性：声明变量时务必提供初始值，以便 TypeScript 正确推断类型
- 显式声明：当类型复杂或不明确时，建议显式声明类型
- any 类型的风险：没有初始值的变量会被推断为 any，这会失去类型检查的保护
- strict 模式：建议启用 strict 模式，避免 any 类型的滥用

最佳实践：充分利用类型推断可以减少代码冗余，但在关键位置（如函数参数、复杂类型）显式声明类型可以提高代码可读性和可维护性。

### 总结

类型推断是 TypeScript 提高开发效率的核心特性。

- 基础类型推断：根据初始值推断变量类型
- 返回类型推断：根据 return 语句推断函数返回类型
- 上下文推断：根据使用位置推断类型
- 最佳通用类型：从多个候选类型中选择最合适的类型
- 泛型推断：根据传入参数自动推断泛型类型
- 显式声明：必要时可显式指定类型

建议：善用类型推断可以写出既简洁又类型安全的代码。

---

## TypeScript 类型断言

Source: https://www.runoob.com/typescript/ts-assertion.html

## TypeScript 类型断言

类型断言（Type Assertion）是一种告诉编译器"我比你更清楚这个值的类型"的机制。它允许开发者手动覆盖 TypeScript 的类型推断结果，类似于其他语言的类型转换，但仅在编译阶段生效，不产生任何运行时代码。

### 两种语法形式

类型断言有两种写法，推荐统一使用 `as` 语法，尖括号语法在 JSX 文件中会与标签语法冲突。

### 实例

// ① 尖括号语法（不推荐，在 .tsx 文件中无法使用）
const str1: any = "hello";
const len1: number = (<string>str1).length;

// ② as 语法（推荐）
const str2: any = "world";
const len2: number = (str2 as string).length;

console.log(len1); // 5
console.log(len2); // 5

运行结果：

```
5
5

```

注意：在 React 的 `.tsx` 文件中，尖括号写法会与 JSX 标签冲突，必须使用 `as` 语法。

### 常见使用场景

#### 1. 处理 any / unknown 类型

从外部接口、JSON 解析等场景获取的数据通常为 `any`，通过断言可以恢复类型提示。

### 实例

const response: any = { name: "Alice", age: 25 };

// 断言为具体类型，获得类型提示
const user = response as { name: string; age: number };

console.log(user.name); // Alice
console.log(user.age); // 25

运行结果：

```
Alice
25

```

#### 2. 收窄联合类型

当已知联合类型中的具体分支，但编译器无法自动判断时，可使用断言明确类型。

### 实例

type Shape =
| { kind: "circle"; radius: number }
| { kind: "rect"; width: number; height: number };

function getArea(shape: Shape): number {
if (shape.kind === "circle") {
return Math.PI * (shape as { kind: "circle"; radius: number }).radius ** 2;
}
const rect = shape as { kind: "rect"; width: number; height: number };
return rect.width * rect.height;
}

console.log(getArea({ kind: "circle", radius: 5 }).toFixed(2)); // 78.54
console.log(getArea({ kind: "rect", width: 4, height: 6 })); // 24

运行结果：

```
78.54
24

```

提示：优先使用类型守卫（`typeof` / `instanceof` / 判别属性）收窄类型，只有在守卫无法覆盖的场景才使用断言。

#### 3. 操作 DOM 元素

`document.querySelector` 返回的是 `Element | null`，断言为具体元素类型后才能访问其专属属性。

### 实例

// querySelector 返回 Element | null
const input = document.querySelector("#username") as HTMLInputElement;

// 断言后可访问 HTMLInputElement 的专属属性
// console.log(input.value);

// 更安全的写法：先判空再断言
const btn = document.querySelector("#submit");
if (btn) {
(btn as HTMLButtonElement).disabled = true;
}

### 非空断言 !

在值后加 `!`，告诉编译器该值不为 `null` 或 `undefined`，适用于你能确保非空但编译器无法推断的场景。

### 实例

function printLength(str?: string) {
// 用 ! 断言 str 一定有值
console.log(str!.length);
}

// 运行正常
printLength("hello"); // 5

// 运行时报错：Cannot read properties of undefined
// printLength();

运行结果：

```
5

```

警告：非空断言会绕过编译器的空值检查，若实际值为 `null` / `undefined`，运行时仍会抛出错误。在条件允许时优先使用可选链 `?.` 代替。

### 常量断言 as const

`as const` 将字面量推断为最窄的只读字面量类型，常用于定义枚举值、配置常量和元组。

### 实例

// 普通声明：推断为 string[]，元素可修改
const colors1 = ["red", "green", "blue"];
colors1.push("yellow"); // 允许

// as const：推断为 readonly ["red", "green", "blue"]，元素和长度均不可变
const colors2 = ["red", "green", "blue"] as const;
// colors2.push("yellow"); // 错误：只读数组不能 push

// 对象同理：所有属性变为 readonly 字面量类型
const config = {
host: "localhost",
port: 3000
} as const;
// config.port = 8080; // 错误：只读属性不能修改

console.log(colors2); // ['red', 'green', 'blue']
console.log(config.host); // localhost
console.log(config.port); // 3000

运行结果：

```
['red', 'green', 'blue']
localhost
3000

```

### 断言不是类型转换

类型断言只影响编译器的类型检查，编译后不生成任何转换代码。将 `"42"` 断言为 `number`，运行时它仍然是字符串。

### 实例

const strNum: any = "42";

// 断言为 number，但运行时仍是 string
const wrongNum = strNum as number;
console.log(typeof wrongNum); // string（不是 number！）
console.log(wrongNum + 1); // 421（字符串拼接，不是数字相加）

// 真正的类型转换需要使用转换函数
const realNum = Number(strNum);
console.log(typeof realNum); // number
console.log(realNum + 1); // 43

运行结果：

```
string
421
number
43

```

### 双重断言

当两个类型之间没有重叠关系，TypeScript 会拒绝直接断言。此时可借助 `unknown` 作为中转，但这是危险操作，应尽量避免。

### 实例

const num = 42;

// 直接断言：number 与 string 无重叠，编译器报错
// const str = num as string;

// 双重断言：先断言为 unknown，再断言为目标类型
const str = num as unknown as string;

console.log(str); // 42
console.log(typeof str); // number（运行时类型未变，断言只在编译期生效）

运行结果：

```
42
number

```

警告：双重断言会完全绕过 TypeScript 的类型系统，极易引发运行时错误。仅在处理第三方库类型不准确等特殊情况下使用，并在代码中注明原因。

### 断言方式对比

断言方式 语法 典型场景 风险 as 断言`value as Type`any 转具体类型、DOM 操作低（有重叠检查） 非空断言`value!`确认非 null/undefined中（跳过空值检查） 常量断言`value as const`字面量常量、枚举值无 双重断言`value as unknown as T`类型无重叠时强制转换高（完全绕过检查）

### 总结

- 语法选择：统一使用 `as` 语法，避免在 JSX 中使用尖括号写法
- 断言 vs 转换：断言仅影响编译期类型，不产生运行时代码，真正的值转换需使用 `Number()`、`String()` 等函数
- 非空断言：使用 `!` 前确保值真的非空，否则运行时会报错，条件允许时优先用 `?.`
- 常量断言：`as const` 是安全且实用的断言，推荐用于配置对象和常量数组
- 双重断言：是最后手段，使用时必须注释说明原因

---

## TypeScript null 和 undefined

Source: https://www.runoob.com/typescript/ts-null-undefined.html

## TypeScript null 和 undefined

在 TypeScript 中，null 和 undefined 有特殊的处理方式。

启用 `strictNullChecks` 后，需要显式处理这些值，这有助于编写更安全的代码。 null 和 undefined 处理方式 null 表示"空值" 需要显式声明 string | null vs undefined 表示"未定义" 可选属性自动包含 name?: string 安全 安全处理方式 可选链: obj?.prop 空值合并: value ?? default 类型守卫: if (val !== null) 处理场景 联合类型声明 可选参数和属性 空值合并运算符

### 为什么需要处理 null 和 undefined

JavaScript 中 null 和 undefined 是常见的错误来源，很多运行时错误都与之相关。

TypeScript 的 strictNullChecks 选项要求开发者显式处理可能为空的值。

这虽然增加了编码工作量，但能大幅减少空值导致的运行时错误，提高代码可靠性。

概念说明：`null` 表示"空值"，即这里没有值；`undefined` 表示"未定义"，即这里还没有被赋值。

### null 和 undefined 基础

null 表示"空值"，undefined 表示"未定义"。两者在 TypeScript 中是独立的类型。

### 实例

// 声明 null 类型的变量
// 只能赋值为 null
var empty: null = null;

// 声明 undefined 类型的变量
// 只能赋值为 undefined
var notDefined: undefined = undefined;

console.log("null: " + empty);
console.log("undefined: " + notDefined);

运行结果：

```
null: null
undefined: undefined

```

说明：在不启用 strictNullChecks 的情况下，这些类型可以相互赋值。启用后需要显式声明。

### 联合类型处理 null

启用 strictNullChecks 后，需要使用联合类型显式声明可能为 null 的值。

### 实例

// 联合类型：可以是字符串或 null
// 显式声明值可能为空
var name: string | null = "Alice";
name = null; // 正确：可以赋值为 null

// 访问可能为 null 的值需要先检查
// 函数参数可能是字符串或 null
function getLength(str: string | null): number {
// 使用条件检查是否为 null
if (str === null) {
return 0; // null 时返回 0
}
// TypeScript 会推断 str 不是 null
return str.length;
}

console.log("长度: " + getLength("hello"));
console.log("长度: " + getLength(null));

运行结果：

```
长度: 5
长度: 0

```

最佳实践：使用类型守卫（if 检查）来收窄类型，让 TypeScript 知道具体是什么类型。

### 可选参数和属性

可选参数和可选属性自动包含 undefined，不需要显式声明 null。

### 实例

// 可选参数：使用 ? 标记
// 参数可能是 string 或 undefined
function greet(name?: string): string {
// 检查是否为 undefined
if (name === undefined) {
return "Hello, stranger!";
}
return "Hello, " + name;
}

console.log(greet("Alice"));
console.log(greet());

// 可选属性：使用 ? 标记
// age 属性是可选的，可能不存在
interface User {
name: string;
age?: number; // 可选属性
}

var user: User = { name: "Bob" };
console.log("用户: " + JSON.stringify(user));

运行结果：

```
Hello, Alice!
Hello, stranger!
用户: {"name":"Bob"}

```

说明：可选参数 `name?: string` 等同于 `string | undefined`。

### 非空断言运算符

使用 `!` 告诉编译器值不为 null 或 undefined。这应该谨慎使用。

### 实例

function getLength(str: string | null): number {
// 非空断言：告诉编译器 str 不为 null
// 这是一个危险的写法，如果 str 真的为 null 会报错
return str!.length;
}

console.log("长度: " + getLength("hello"));

警告：非空断言会绕过 TypeScript 的类型检查，如果值实际为 null，运行时仍会报错。尽量避免使用。

### 空值合并运算符

使用 `??` 提供默认值，只有当值为 null 或 undefined 时才使用默认值。

### 实例

// 初始值为 null
var name: string | null = null;

// ?? 运算符：左侧为 null/undefined 时使用右侧值
// 如果 name 为 null，使用 "Guest"
var displayName = name ?? "Guest";

console.log("显示名称: " + displayName);

// 对比 || 运算符（会把 0 和空字符串视为 falsy）
var num: number | null = 0;
var result1 = num ?? 100; // 0（正确：0 不是 null/undefined）
var result2 = num || 100; // 100（错误：0 被视为 falsy）

console.log("?? 结果: " + result1);
console.log("|| 结果: " + result2);

运行结果：

```
显示名称: Guest
?? 结果: 0
|| 结果: 100

```

建议：优先使用 `??` 而不是 `||`，因为 `??` 只会处理 null 和 undefined，不会错误地将 0 或空字符串视为 falsy。

### 可选链

使用 `?.` 安全访问可能不存在的嵌套属性，避免多层嵌套时的空值检查。

### 实例

// 定义人员类型，地址是可选的
interface Person {
name: string;
address?: {
city: string;
};
}

// 创建人员对象，没有地址
var person: Person = { name: "Alice" };

// 可选链：安全访问可能不存在的属性
// 如果 address 不存在，返回 undefined 而不会报错
var city = person.address?.city;

// 结合空值合并运算符
console.log("城市: " + city);
console.log("城市: " + (person.address?.city ?? "未知"));

运行结果：

```
城市: undefined
城市: 未知

```

最佳实践：可选链 `?.` 是处理嵌套可选属性的最佳方式，配合 `??` 可以提供默认值的兜底。

### 注意事项

- 严格模式：建议启用 strictNullChecks 以获得更好的类型安全
- 避免断言：尽量不要使用非空断言 `!`
- ?? vs ||：优先使用 `??` 避免误判 0 和空字符串
- 可选链：处理嵌套可选属性时使用 `?.`

最佳实践：使用类型守卫、可选链和空值合并运算符来处理 null 和 undefined，让代码更安全、更易读。

### 总结

正确处理 null 和 undefined 是编写安全 TypeScript 代码的关键。

- 严格模式：启用 strictNullChecks 后需显式处理
- 联合类型：使用 `string | null` 声明
- 可选参数/属性：自动包含 undefined
- 非空断言：使用 `!`（谨慎使用）
- 空值合并：使用 `??` 提供默认值
- 可选链：使用 `?.` 安全访问

建议：养成处理空值的习惯，使用可选链和空值合并运算符让代码更健壮。

---

## TypeScript Symbol

Source: https://www.runoob.com/typescript/ts-symbol.html

## TypeScript Symbol

Symbol 是 ES6 引入的原始数据类型，表示唯一的标识符。

在 TypeScript 中，Symbol 可以用作对象的属性键，确保属性的唯一性。

这在需要创建私有属性、避免属性名冲突等场景非常有用。 Symbol 特性 Symbol() 创建 var sym1 = Symbol("key") var sym2 = Symbol("key") // sym1 !== sym2 每次创建都唯一 用作对象属性键 var obj = { [sym]: "value" } // 唯一属性名 Symbol.for() 全局注册表 相同 key 相等 常见内置 Symbol Symbol.iterator - 迭代器 Symbol.toStringTag - 对象描述 Symbol.hasInstance - instanceof

### 为什么需要 Symbol

在 JavaScript 中，对象的属性名都是字符串，有时可能会发生冲突。

Symbol 提供了创建唯一标识符的方式，每次调用 Symbol() 都会创建一个新的、唯一的值。

这在需要创建私有属性、定义唯一常量、实现迭代器等场景非常有用。

概念说明：Symbol 是 JavaScript 的原始数据类型之一，通过 Symbol() 函数创建。每个 Symbol 值都是唯一的，即使描述相同也不相等。

### 创建 Symbol

使用 Symbol() 函数创建唯一的 Symbol 值。

### 实例

// 创建 Symbol，传入描述字符串（可选）
var sym1 = Symbol("description");
var sym2 = Symbol("description");

// 每次创建的 Symbol 都是唯一的，即使描述相同
console.log("sym1 === sym2: " + (sym1 === sym2));
console.log("sym1: " + sym1.toString());

运行结果：

```
sym1 === sym2: false
sym1: Symbol(description)

```

唯一性：这是 Symbol 最重要的特性，每次调用 Symbol() 都会创建一个新值，与任何其他值都不相等。

### Symbol 作为对象属性

Symbol 可以用作对象的属性键，创建唯一的属性名。

### 实例

// 创建 Symbol 作为属性键
var sym = Symbol("key");

// 使用 Symbol 作为对象的属性键
var obj = {
name: "Alice", // 普通字符串属性
[sym]: "secret value" // Symbol 属性，计算属性名
};

// 访问普通属性
console.log("普通属性: " + obj.name);
// 访问 Symbol 属性
console.log("Symbol 属性: " + obj[sym]);
// Symbol 属性不会出现在 JSON 中
console.log("对象: " + JSON.stringify(obj));

运行结果：

```
普通属性: Alice
Symbol 属性: secret value
对象: {"name":"Alice"}

```

隐私：Symbol 属性不会出现在 JSON 序列化中，也不会被 for...in 遍历到，可以用来创建"私有"属性。

### 全局 Symbol 注册表

使用 Symbol.for() 访问全局注册表中的 Symbol，相同 key 会返回相同的 Symbol。

### 实例

// 使用 Symbol.for 方法创建/获取全局 Symbol
// 如果 key 不存在，会创建新的；如果已存在，返回已有的
var globalSym1 = Symbol.for("global");
var globalSym2 = Symbol.for("global");

// 相同 key 的 Symbol 是相等的
console.log("全局 Symbol 相等: " + (globalSym1 === globalSym2));

// 获取 Symbol 的 key
console.log("Symbol key: " + Symbol.keyFor(globalSym1));

运行结果：

```
全局 Symbol 相等: true
Symbol key: global

```

区别：Symbol() 每次创建新的值，Symbol.for() 在全局注册表中查找或创建。

### 内置 Symbol

JavaScript 定义了一些内置的 Symbol 值，用于自定义语言行为。

### 实例

// Symbol.iterator 用于定义对象的默认迭代器
var arr = [1, 2, 3];
// 获取数组的迭代器
var iterator = arr[Symbol.iterator]();

// 使用迭代器遍历
console.log("第一个元素: " + iterator.next().value);
console.log("第二个元素: " + iterator.next().value);

// Symbol.toStringTag 自定义对象的 toString() 返回值
var obj = {
[Symbol.toStringTag]: "MyObject"
};
console.log("对象类型: " + obj.toString());

运行结果：

```
第一个元素: 1
第二个元素: 2
对象类型: [object MyObject]

```

重要：内置 Symbol 用于自定义语言行为，如迭代器、类型转换等，是 JavaScript 高级特性的基础。

### Symbol 类型注解

在 TypeScript 中使用 Symbol 类型进行类型注解。

### 实例

// Symbol 类型注解
var sym: symbol = Symbol("key");

// 对象的键类型为 symbol，值为 string
var obj: { [key: symbol]: string } = {};

// 使用 Symbol 作为键
obj[sym] = "value";
console.log("Symbol 属性值: " + obj[sym]);

类型：TypeScript 中使用 `symbol` 类型注解 Symbol 值。

### 注意事项

- 唯一性：每次 Symbol() 调用的值都不同
- 不可枚举：Symbol 属性不会出现在 for...in 循环中
- JSON 忽略：Symbol 属性不会被 JSON 序列化
- 全局注册：Symbol.for() 在全局注册表中共享

应用场景：Symbol 常用于创建对象的私有属性、定义唯一常量、避免属性名冲突等。

### 总结

Symbol 是 TypeScript 中重要的原始类型。

- 唯一性：每次创建的 Symbol 都不相等
- 属性键：可用作对象的唯一属性键
- 全局注册：Symbol.for() 创建/获取全局 Symbol
- 内置 Symbol：Symbol.iterator、Symbol.toStringTag 等
- 类型注解：使用 symbol 类型

建议：在需要创建唯一标识符、避免属性名冲突、或需要"私有"属性时，使用 Symbol。

---

## TypeScript 特殊类型：never、void、unknown、any

Source: https://www.runoob.com/typescript/ts-special-types.html

## TypeScript 特殊类型：never、void、unknown、any

TypeScript 有四个特殊的类型：never、void、unknown 和 any。

它们在类型系统中扮演重要角色，理解它们的区别对于编写类型安全的代码至关重要。 四种特殊类型对比 any 任意类型 绕过类型检查 无安全性 unknown 安全任意类型 使用需检查 有安全性 void 无返回值 用于函数 可返回 undefined never 永不返回 抛出异常/循环 所有类型的子类型 类型安全等级 any（最低） 完全绕过类型检查 unknown 需类型检查后使用 void 函数无返回值 never（最高） 最严格，最安全

### 为什么需要特殊类型

TypeScript 的类型系统非常强大，除了常规的类型外，还提供了四个特殊类型来处理特定的场景。

理解这四种特殊类型的区别，能够帮助我们编写更安全、更准确的类型代码。

概念说明：never、void、unknown、any 是 TypeScript 的特殊类型，它们各自有不同的含义和使用场景。

### never 类型

never 表示永不返回的类型。通常用于抛出异常或无限循环的函数。

### 实例

// 抛出异常的函数
// 函数永远不会正常返回，总会抛出错误
function throwError(message: string): never {
throw new Error(message);
}

// 无限循环的函数
// 函数永远不会返回，程序会一直运行
function infiniteLoop(): never {
while (true) {
console.log("运行中...");
}
}

// never 是所有类型的子类型
// 这意味着 never 可以赋值给任何类型
var neverValue: never;
var num: number = neverValue; // 正确：never 是 number 的子类型
console.log("never 赋值给 number: " + num);

运行结果：

```
never 赋值给 number: undefined

```

重要：never 是所有类型的子类型，这意味着 never 可以赋值给任何类型，但任何类型都不能赋值给 never（除了 never 本身）。

### void 类型

void 表示没有返回值，通常用于声明没有 return 语句的函数。

### 实例

// 无返回值函数
// 函数执行后没有返回值
function logMessage(message: string): void {
console.log("日志: " + message);
// 没有 return 语句，或 return undefined
}

logMessage("Hello");

// void 变量（很少使用）
// 只能赋值为 undefined
var empty: void = undefined;
console.log("void 变量: " + empty);

运行结果：

```
日志: Hello
void 变量: undefined

```

说明：void 实际上和 undefined 很相似，主要用于函数的返回类型声明，表示这个函数没有返回值。

### unknown 类型

unknown 是类型安全的 any。使用 unknown 时，必须先进行类型检查才能使用。

### 实例

// unknown 接受任意类型
// 这是"安全的" any
var value: unknown = "hello";
value = 42;
value = true;

// 未知类型不能直接赋值给其他类型
// 如果取消注释下一行，编译器会报错
// var str: string = value;

// 需要进行类型检查后，才能赋值
if (typeof value === "string") {
// TypeScript 知道 value 是 string 类型
var str: string = value;
console.log("字符串长度: " + str.length);
}

运行结果：

```
字符串长度: 5

```

安全：unknown 强制要求在使用前进行类型检查，这保护了类型安全。处理未知来源的数据时，应该使用 unknown 而不是 any。

### any 类型

any 完全绕过类型检查，可以赋值为任意类型。这是 TypeScript 中最不安全的类型。

### 实例

// any 可以是任意类型
// 完全绕过类型检查
var anything: any = "hello";
anything = 42;
anything = true;

// any 可以赋值给任意类型
// 这会破坏类型安全
var str: string = anything;
var num: number = anything;

console.log("字符串: " + str);
console.log("数字: " + num);

// any 类型的对象可以调用任意方法
// 编译器不会报错，但运行时可能出错
var obj: any = {};
obj.foo(); // 不会报错，但实际上 foo 方法不存在
obj.bar = "value";

警告：any 类型会完全绕过 TypeScript 的类型检查，应该尽量避免使用。在必须使用的情况下，要确保值的类型是正确的。

### 对比 any vs unknown

any 和 unknown 都可以接受任意类型，但安全性完全不同。

特性 any unknown 接受任意类型 ✓ 是 ✓ 是 直接赋值给其他类型 ✓ 是 ✗ 否（需检查） 直接调用方法 ✓ 是 ✗ 否（需检查） 类型安全 ✗ 无 ✓ 有

建议：处理未知类型的数据时，优先使用 unknown 而不是 any，这样能强制进行类型检查。

### never vs void

never 和 void 都用于函数返回类型，但含义完全不同。

特性 never void 含义 永不返回 没有返回值 用于函数 抛出异常/无限循环 普通无返回值函数 可赋值给其他类型 ✓ 是（子类型） 只能赋值给 void/any

关键区别：never 表示函数永远不会正常返回（要么抛异常，要么死循环）；void 表示函数正常执行完毕，但没有返回值。

### 实际应用：Exhaustive Check

never 类型的一个重要用途是进行穷举检查，确保所有可能的情况都被处理。

### 实例

// 定义形状的联合类型
type Shape = { kind: "circle", radius: number }
| { kind: "square", side: number };

// 计算形状面积
function area(shape: Shape): number {
switch (shape.kind) {
case "circle":
return Math.PI * shape.radius ** 2;
case "square":
return shape.side ** 2;
default:
// default 分支应该处理所有未预料到的情况
// 如果添加新的 Shape 类型但忘记处理，编译器会报错
var _exhaustive: never = shape;
return _exhaustive;
}
}

var circle = { kind: "circle" as const, radius: 5 };
var square = { kind: "square" as const, side: 4 };

console.log("圆形面积: " + area(circle).toFixed(2));
console.log("正方形面积: " + area(square));

运行结果：

```
圆形面积: 78.54
正方形面积: 16

```

Exhaustive Check：这种模式确保当添加新的类型分支时，如果忘记在 switch 中处理，编译器会报错，因为新增的类型无法赋值给 never。

### 注意事项

- never 是子类型：never 可以赋值给任何类型，但任何类型都不能赋值给 never
- void vs undefined：在 JavaScript 中，void 和 undefined 基本等同
- unknown 需要检查：使用 unknown 类型的值前，必须进行类型检查
- 避免 any：尽量不要使用 any，优先使用 unknown

最佳实践：从高到低的类型安全程度：never > void > unknown > any。尽量使用更安全的类型。

### 总结

理解这四种特殊类型对于编写高质量的 TypeScript 代码非常重要。

- never：永不返回，用于 exhaustive check，是所有类型的子类型
- void：无返回值，用于普通没有返回值的函数
- unknown：安全的任意类型，使用前必须进行类型检查
- any：绕过所有类型检查，应该尽量避免使用

建议：编写代码时，优先使用更严格的类型。处理未知数据使用 unknown，函数无返回值使用 void，需要穷举检查时使用 never。

---

## TypeScript Set 和 WeakMap

Source: https://www.runoob.com/typescript/ts-set-weakmap.html

## TypeScript Set 和 WeakMap

TypeScript 继承自 JavaScript 的 Set 和 WeakMap 数据结构，提供了更强大的类型支持。

这些数据结构在处理唯一值集合、键值对映射、缓存等场景非常有用。 Set 和 Map 数据结构 Set 值的集合 值唯一，不重复 可遍历 add/has/delete WeakSet 对象弱引用 不影响 GC 不可遍历 add/has/delete Map 键值对集合 键可以是任意类型 可遍历 set/get/has WeakMap 键弱引用 键必须是对象 不影响 GC 不可遍历 应用场景 Set: 去重、唯一值集合 Map: 键值映射、缓存 WeakMap/WeakSet: 内存优化

### 为什么需要 Set 和 WeakMap

在开发中，我们经常需要处理唯一值集合和键值对映射。

Set 提供了自动去重的集合功能，比数组更方便处理唯一值。

WeakSet 和 WeakMap 使用弱引用，不会阻止垃圾回收，适用于需要避免内存泄漏的场景，如缓存 DOM 节点。

概念说明：Set 是值的集合，值唯一；Map 是键值对集合，键可以是任意类型。WeakSet 和 WeakMap 使用弱引用，不影响垃圾回收。

### Set

Set 是值的集合，值唯一，不允许重复。

### 实例

// 创建 Set，指定元素类型为 number
var numbers = new Set<number>();

// 添加元素
numbers.add(1);
numbers.add(2);
numbers.add(3);
numbers.add(1); // 重复值会被忽略，不会添加

// 检查大小和包含
console.log("Set 大小: " + numbers.size);
console.log("是否包含 2: " + numbers.has(2));

// 遍历 Set
numbers.forEach(function(value) {
console.log("值: " + value);
});

// 转换为数组
var arr = Array.from(numbers);
console.log("转换为数组: " + arr);

运行结果：

```
Set 大小: 3
是否包含 2: true
值: 1
值: 2
值: 3
转换为数组: 1,2,3

```

去重：Set 会自动忽略重复的值，这使得它非常适合用于数组去重。

### Set 类型注解

可以显式指定 Set 中值的类型。

### 实例

// 字符串 Set
// 只能添加字符串类型的值
var stringSet: Set<string> = new Set();
stringSet.add("a");
stringSet.add("b");

// 对象 Set
// 定义 Person 接口
interface Person {
name: string;
}
// 创建存储 Person 对象的 Set
var personSet: Set<Person> = new Set();
personSet.add({ name: "Alice" });
personSet.add({ name: "Bob" });

console.log("字符串 Set: " + Array.from(stringSet));
console.log("对象 Set 大小: " + personSet.size);

泛型：使用 `Set<T>` 语法指定 Set 中元素的类型。

### WeakSet

WeakSet 存储对象引用，引用为弱引用（不影响垃圾回收）。

### 实例

// WeakSet 只能存储对象，不能存储原始值
var weakSet = new WeakSet();

// 创建对象
var obj1 = { name: "Alice" };
var obj2 = { name: "Bob" };

// 添加对象到 WeakSet
weakSet.add(obj1);
weakSet.add(obj2);

// 检查是否包含
console.log("是否包含 obj1: " + weakSet.has(obj1));

// 移除引用后，对象可能被垃圾回收
weakSet.delete(obj1);
console.log("删除后是否包含 obj1: " + weakSet.has(obj1));

注意：WeakSet 不能遍历，类型注解只能是 `object`。这使得 WeakSet 适合存储需要被垃圾回收的对象。

### Map

Map 是键值对集合，键可以是任意类型。

### 实例

// 创建 Map，键类型为 string，值类型为 number
var map = new Map<string, number>();

// 设置键值对
map.set("one", 1);
map.set("two", 2);
map.set("three", 3);

// 获取值
console.log("获取 two: " + map.get("two"));
console.log("Map 大小: " + map.size);
console.log("是否包含 three: " + map.has("three"));

// 遍历 Map
map.forEach(function(value, key) {
console.log(key + ": " + value);
});

// 转换为数组
console.log("转换为数组: " + Array.from(map.entries()));

运行结果：

```
获取 two: 2
Map 大小: 3
是否包含 three: true
one: 1
two: 2
three: 3
转换为数组: one,1,two,2,three,3

```

优势：Map 的键可以是任意类型（对象、函数等），这比使用对象作为键更灵活。

### WeakMap

WeakMap 的键是弱引用，不影响垃圾回收。

### 实例

// WeakMap 的键必须是对象
// 键类型为 object，值类型为 string
var weakMap = new WeakMap<object, string>();

// 创建对象作为键
var keyObj = { id: 1 };
// 设置键值对
weakMap.set(keyObj, "value1");

// 获取值
console.log("获取值: " + weakMap.get(keyObj));
console.log("是否包含: " + weakMap.has(keyObj));

// 删除键值对
weakMap.delete(keyObj);
console.log("删除后: " + weakMap.has(keyObj));

应用场景：WeakMap 常用于缓存 DOM 节点数据，当 DOM 节点被移除时，缓存数据也会被自动清理，避免内存泄漏。

### 实际应用场景

使用 Map 统计数组元素出现次数。

### 实例

// 使用 Map 统计数组中每个元素的出现次数
function countElements(arr: string[]): Map<string, number> {
// 创建 Map，键是字符串，值是数字
var counts = new Map<string, number>();

// 遍历数组
for (var _i = 0, arr_1 = arr; _i < arr_1.length; _i++) {
var item = arr_1[_i];
// 获取当前计数，如果没有则返回 0
var currentCount = counts.get(item) || 0;
// 更新计数
counts.set(item, currentCount + 1);
}

return counts;
}

// 测试
var fruits = ["apple", "banana", "apple", "orange", "banana", "apple"];
var result = countElements(fruits);

// 遍历结果
result.forEach(function(count, fruit) {
console.log(fruit + ": " + count);
});

运行结果：

```
apple: 3
banana: 2
orange: 1

```

实用：Map 是实现缓存、统计、索引等功能的理想选择。

### 注意事项

- Set 唯一性：Set 自动忽略重复值
- WeakSet/WeakMap：键必须是对象，不能遍历
- Map 键类型：Map 的键可以是任意类型
- 内存管理：WeakSet/WeakMap 不阻止垃圾回收

选择建议：需要唯一值集合用 Set，需要键值映射用 Map，需要避免内存泄漏用 WeakSet/WeakMap。

### 总结

Set 和 Map 是 TypeScript 中非常有用的数据结构。

- Set：值的集合，值唯一，自动去重
- WeakSet：对象弱引用集合，不可遍历，适用于需要被垃圾回收的场景
- Map：键值对集合，键可以是任意类型
- WeakMap：键弱引用，不可遍历，适用于缓存和私有数据

建议：根据具体需求选择合适的数据结构：需要去重用 Set，需要映射用 Map，需要内存优化用 Weak 版本。

---

## TypeScript 元组

Source: https://www.runoob.com/typescript/ts-tuple.html

## TypeScript 元组

我们知道数组中元素的数据类型都一般是相同的（any[] 类型的数组可以不同），如果存储的元素数据类型不同，则需要使用元组。

TypeScript 中的元组（Tuple）是一种特殊类型的数组，它允许在数组中存储不同类型的元素，与普通数组不同，元组中的每个元素都有明确的类型和位置。元组可以在很多场景下用于表示固定长度、且各元素类型已知的数据结构。

创建元组的语法格式如下：

```
let tuple: [类型1, 类型2, 类型3, ...];
```

#### 实例

声明一个元组并初始化：

```

let mytuple: [number, string];
mytuple = [42,"Runoob"];
```

在上面的例子中，mytuple 是一个元组，它包含一个 number 类型和一个 string 类型的元素。

#### 访问元组

元组中元素使用索引来访问，第一个元素的索引值为 0，第二个为 1，以此类推第 n 个为 n-1，语法格式如下:

```
tuple_name[index]
```

#### 实例

以下实例定义了元组，包含了数字和字符串两种类型的元素：

### TypeScript

let mytuple: [number, string, boolean] = [42, "Runoob", true]; // 访问元组中的元素 let num = mytuple[0]; // 访问第一个元素，值为 42，类型为 number let str = mytuple[1]; // 访问第二个元素，值为 "Runoob"，类型为 string let bool = mytuple[2]; // 访问第三个元素，值为 true，类型为 boolean console.log(num); // 输出: 42 console.log(str); // 输出: Runoob console.log(bool); // 输出: true

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var mytuple = [42, "Runoob", true]; // 访问元组中的元素 var num = mytuple[0]; // 访问第一个元素，值为 42，类型为 number var str = mytuple[1]; // 访问第二个元素，值为 "Runoob"，类型为 string var bool = mytuple[2]; // 访问第三个元素，值为 true，类型为 boolean console.log(num); // 输出: 42 console.log(str); // 输出: Runoob console.log(bool); // 输出: true

输出结果为：

```
42
Runoob
true

```

### 元组运算

我们可以使用以下两个函数向元组添加新元素或者删除元素：

- push() 向元组添加元素，添加在最后面。
- pop() 从元组中移除元素（最后一个），并返回移除的元素。

push 方法可以向元组的末尾添加一个元素，类型必须符合元组定义中的类型约束。如果超出元组的类型约束，TypeScript 会报错。

### TypeScript

var tuple = [42, "Hello"]; // 添加符合类型的元素 tuple.push("World"); // 合法，因为元组定义了可选的 string 类型 console.log(tuple); // 输出: [42, "Hello", "World"]

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var tuple = [42, "Hello"]; // 添加符合类型的元素 tuple.push("World"); // 合法，因为元组定义了可选的 string 类型 console.log(tuple); // 输出: [42, "Hello", "World"]

输出结果为：

```
[ 42, 'Hello', 'World' ]
```

pop 方法从元组的末尾移除一个元素，并返回该元素。返回的元素类型将根据元组的定义类型推断。

### 实例

let tuple: [number, string, boolean] = [42, "Hello", true];

// 移除最后一个元素
let lastElement = tuple.pop();

console.log(lastElement); // 输出: true
console.log(tuple); // 输出: [42, "Hello"]

### 更新元组

元组是可变的，这意味着我们可以对元组进行更新操作：

### TypeScript

var mytuple = [42, "Runoob", "Taobao", "Google"]; // 创建一个元组 console.log("元组的第一个元素为：" + mytuple[0]) // 更新元组元素 mytuple[0] = 121 console.log("元组中的第一个元素更新为："+ mytuple[0])

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var mytuple = [42, "Runoob", "Taobao", "Google"]; // 创建一个元组 console.log("元组的第一个元素为：" + mytuple[0]); // 更新元组元素 mytuple[0] = 121; console.log("元组中的第一个元素更新为：" + mytuple[0]);

输出结果为：

```
元组的第一个元素为：42
元组中的第一个元素更新为：121
```

### 解构元组

我们也可以把元组元素赋值给变量，如下所示：

### TypeScript

let a: [number, string, boolean] = [42, "Hello", true];// 创建一个元组 var [b,c] = a console.log( b ) console.log( c )

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var a = [42, "Hello", true]; // 创建一个元组 var b = a[0], c = a[1]; console.log(b); console.log(c);

输出结果为：

```
42
Hello
```

### 使用标签元组

TypeScript 还允许为元组中的每个元素添加标签，这使得元组的含义更加清晰：

```

let tuple: [id: number, name: string] = [1, "John"];
```

在这个例子中，id 和 name 是元组的标签，可以让代码更加可读。

### 元组的实际应用

元组常用于函数返回多个值的场景，或者表示某些固定结构的数据，比如：

### 实例

function getUserInfo(): [number, string] {
return [1, "John Doe"];
}

const [userId, userName] = getUserInfo();
console.log(userId); // 输出: 1
console.log(userName); // 输出: John Doe

### 元组的类型推断

TypeScript 可以根据数组的元素自动推断出元组的类型：

```

let tuple = [42, "Hello"] as const; // 元组类型：[42, "Hello"]
```

通过 as const 断言，TypeScript 会将该元组视为一个不可变的常量元组。

### 连接元组

元组可以使用数组的 concat 方法进行连接，但需要注意连接后的结果是一个普通的数组，而不是元组。

### 实例

let tuple1: [number, string] = [42, "Hello"];
let tuple2: [boolean, number] = [true, 100];

let result = tuple1.concat(tuple2); // 结果是一个数组: [42, "Hello", true, 100]
console.log(result); // 输出: [42, "Hello", true, 100]

### 切片元组

你可以使用数组的 slice 方法对元组进行切片操作，返回一个新的数组。

### 实例

let tuple: [number, string, boolean] = [42, "Hello", true];

let sliced = tuple.slice(1); // 从索引 1 开始切片
console.log(sliced); // 输出: ["Hello", true]

### 遍历元组

你可以使用 for...of 循环或者 forEach 方法遍历元组中的元素。

### 实例

let tuple: [number, string, boolean] = [42, "Hello", true];

// 使用 for...of 循环
for (let item of tuple) {
console.log(item);
}

// 使用 forEach 方法
tuple.forEach(item => console.log(item));

### 转换为普通数组

虽然元组是一个固定长度、固定类型的数组，但可以通过 Array.prototype 的方法将其转换为普通数组进行进一步处理。

### 实例

let tuple: [number, string, boolean] = [42, "Hello", true];

// 转换为数组并使用数组方法
let array = Array.from(tuple);
array.push("New Element");

console.log(array); // 输出: [42, "Hello", true, "New Element"]

### 扩展元组

使用剩余运算符可以轻松地将多个元组合并成一个新的元组：

### 实例

let tuple1: [number, string] = [42, "Hello"];
let tuple2: [boolean] = [true];

let extendedTuple: [number, string, ...typeof tuple2] = [42, "Hello", ...tuple2];

console.log(extendedTuple); // 输出: [42, "Hello", true]

---

## TypeScript 联合类型

Source: https://www.runoob.com/typescript/ts-union.html

## TypeScript 联合类型

联合类型（Union Types）可以通过管道(|)将变量设置多种类型，赋值时可以根据设置的类型来赋值。

注意：只能赋值指定的类型，如果赋值其它类型就会报错。

创建联合类型的语法格式如下：

```
Type1|Type2|Type3
```

#### 实例

声明一个联合类型：

### TypeScript

var val:string|number val = 12 console.log("数字为 "+ val) val = "Runoob" console.log("字符串为 " + val)

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var val; val = 12; console.log("数字为 " + val); val = "Runoob"; console.log("字符串为 " + val);

输出结果为：

```
数字为 12
字符串为 Runoob
```

如果赋值其它类型就会报错：

```
var val:string|number
val = true
```

也可以将联合类型作为函数参数使用：

### TypeScript

function disp(name:string|string[]) { if(typeof name == "string") { console.log(name) } else { var i; for(i = 0;i<name.length;i++) { console.log(name[i]) } } } disp("Runoob") console.log("输出数组....") disp(["Runoob","Google","Taobao","Facebook"])

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

function disp(name) { if (typeof name == "string") { console.log(name); } else { var i; for (i = 0; i < name.length; i++) { console.log(name[i]); } } } disp("Runoob"); console.log("输出数组...."); disp(["Runoob", "Google", "Taobao", "Facebook"]);

输出结果为：

```
Runoob
输出数组....
Runoob
Google
Taobao
Facebook
```

### 联合类型数组

我们也可以将数组声明为联合类型：

### TypeScript

var arr:number[]|string[]; var i:number; arr = [1,2,4] console.log("**数字数组**") for(i = 0;i<arr.length;i++) { console.log(arr[i]) } arr = ["Runoob","Google","Taobao"] console.log("**字符串数组**") for(i = 0;i<arr.length;i++) { console.log(arr[i]) }

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var arr; var i; arr = [1, 2, 4]; console.log("**数字数组**"); for (i = 0; i < arr.length; i++) { console.log(arr[i]); } arr = ["Runoob", "Google", "Taobao"]; console.log("**字符串数组**"); for (i = 0; i < arr.length; i++) { console.log(arr[i]); }

输出结果为：

```
**数字数组**
1
2
4
**字符串数组**
Runoob
Google
Taobao
```

---

## TypeScript 接口

Source: https://www.runoob.com/typescript/ts-interface.html

## TypeScript 接口

接口是一系列抽象方法的声明，是一些方法特征的集合，这些方法都应该是抽象的，需要由具体的类去实现，然后第三方就可以通过这组抽象方法调用，让具体的类执行具体的方法。

TypeScript 接口定义如下：

```
interface interface_name {
}
```

#### 实例

以下实例中，我们定义了一个接口 IPerson，接着定义了一个变量 customer，它的类型是 IPerson。

customer 实现了接口 IPerson 的属性和方法。

### TypeScript

interface IPerson { firstName:string, lastName:string, sayHi: ()=>string } var customer:IPerson = { firstName:"Tom", lastName:"Hanks", sayHi: ():string =>{return "Hi there"} } console.log("Customer 对象 ") console.log(customer.firstName) console.log(customer.lastName) console.log(customer.sayHi()) var employee:IPerson = { firstName:"Jim", lastName:"Blakes", sayHi: ():string =>{return "Hello!!!"} } console.log("Employee 对象 ") console.log(employee.firstName) console.log(employee.lastName)

需要注意接口不能转换为 JavaScript。 它只是 TypeScript 的一部分。

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var customer = { firstName: "Tom", lastName: "Hanks", sayHi: function () { return "Hi there"; } }; console.log("Customer 对象 "); console.log(customer.firstName); console.log(customer.lastName); console.log(customer.sayHi()); var employee = { firstName: "Jim", lastName: "Blakes", sayHi: function () { return "Hello!!!"; } }; console.log("Employee 对象 "); console.log(employee.firstName); console.log(employee.lastName);

输出结果为：

```
Customer 对象
Tom
Hanks
Hi there
Employee 对象
Jim
Blakes
```

### 联合类型和接口

以下实例演示了如何在接口中使用联合类型：

### TypeScript

interface RunOptions { program:string; commandline:string[]|string|(()=>string); } // commandline 是字符串 var options:RunOptions = {program:"test1",commandline:"Hello"}; console.log(options.commandline) // commandline 是字符串数组 options = {program:"test1",commandline:["Hello","World"]}; console.log(options.commandline[0]); console.log(options.commandline[1]); // commandline 是一个函数表达式 options = {program:"test1",commandline:()=>{return "**Hello World**";}}; var fn:any = options.commandline; console.log(fn());

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

// commandline 是字符串 var options = { program: "test1", commandline: "Hello" }; console.log(options.commandline); // commandline 是字符串数组 options = { program: "test1", commandline: ["Hello", "World"] }; console.log(options.commandline[0]); console.log(options.commandline[1]); // commandline 是一个函数表达式 options = { program: "test1", commandline: function () { return "**Hello World**"; } }; var fn = options.commandline; console.log(fn());

输出结果为：

```
Hello
Hello
World
**Hello World**
```

### 接口和数组

接口中我们可以将数组的索引值和元素设置为不同类型，索引值可以是数字或字符串。

设置元素为字符串类型：

### 实例

interface namelist { [index:number]:string } // 类型一致，正确 var list2:namelist = ["Google","Runoob","Taobao"] // 错误元素 1 不是 string 类型 // var list2:namelist = ["Runoob",1,"Taobao"]

如果使用了其他类型会报错：

### 实例

interface namelist { [index:number]:string } // 类型一致，正确 // var list2:namelist = ["Google","Runoob","Taobao"] // 错误元素 1 不是 string 类型 var list2:namelist = ["John",1,"Bran"] 执行后报错如下，显示类型不一致：

```
test.ts:8:30 - error TS2322: Type 'number' is not assignable to type 'string'.

8 var list2:namelist = ["John",1,"Bran"]
~

test.ts:2:4
2 [index:number]:string
~~~~~~~~~~~~~~~~~~~~~
The expected type comes from this index signature.


Found 1 error.
```

### TypeScript

interface ages { [index:string]:number } var agelist:ages; // 类型正确 agelist["runoob"] = 15 // 类型错误，输出 error TS2322: Type '"google"' is not assignable to type 'number'. // agelist[2] = "google"

### 接口继承

接口继承就是说接口可以通过其他接口来扩展自己。

Typescript 允许接口继承多个接口。

继承使用关键字 extends。

单接口继承语法格式：

```
Child_interface_name extends super_interface_name
```

多接口继承语法格式：

```
Child_interface_name extends super_interface1_name, super_interface2_name,…,super_interfaceN_name
```

继承的各个接口使用逗号 , 分隔。

#### 单继承实例

### TypeScript

interface Person { age:number } interface Musician extends Person { instrument:string } var drummer = <Musician>{}; drummer.age = 27 drummer.instrument = "Drums" console.log("年龄: "+drummer.age) console.log("喜欢的乐器: "+drummer.instrument)

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var drummer = {}; drummer.age = 27; drummer.instrument = "Drums"; console.log("年龄: " + drummer.age); console.log("喜欢的乐器: " + drummer.instrument);

输出结果为：

```
年龄: 27
喜欢的乐器: Drums
```

#### 多继承实例

### TypeScript

interface IParent1 { v1:number } interface IParent2 { v2:number } interface Child extends IParent1, IParent2 { } var Iobj:Child = { v1:12, v2:23} console.log("value 1: "+Iobj.v1+" value 2: "+Iobj.v2)

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var Iobj = { v1: 12, v2: 23 }; console.log("value 1: " + Iobj.v1 + " value 2: " + Iobj.v2);

输出结果为：

```
value 1: 12 value 2: 23
```

---

## TypeScript 类

Source: https://www.runoob.com/typescript/ts-class.html

## TypeScript 类

类是面向对象编程（OOP）的核心概念，它是一种模板或蓝图，用于创建具有相同属性和方法的对象。TypeScript 完全支持面向对象编程，提供了类、继承、访问修饰符等特性。

类封装了数据（属性）和行为（方法），使得代码更加模块化、可复用和易维护。通过类，我们可以创建多个具有相同结构的对象，这些对象称为类的实例。

### 类的定义

TypeScript 使用 `class` 关键字定义类。一个类可以包含以下成员：

- 字段（Field）：类中声明的变量，表示对象的属性
- 构造函数（Constructor）：类实例化时调用的特殊方法，用于初始化对象
- 方法（Method）：类中定义的函数，表示对象的行为

#### 语法格式

```

class class_name {
// 字段声明
field1: type;
field2: type;

// 构造函数
constructor(parameters) {
// 初始化代码
}

// 方法
methodName(): return_type {
// 方法实现
}
}

```

#### 实例：创建简单的类

### 实例

// 定义一个空的 Person 类
class Person {
}

编译后的 JavaScript：

### 实例

var Person = /** @class */ (function () {
function Person() {
}
return Person;
}());

### 类的字段和构造函数

类的字段是存储对象数据的地方，构造函数在对象创建时自动调用，用于初始化字段。

### 实例

// 定义 Car 类
class Car {
// 字段：描述汽车的属性
engine: string;

// 构造函数：在创建对象时初始化 engine
constructor(engine: string) {
this.engine = engine;
}

// 方法：显示发动机信息
disp(): void {
console.log("发动机型号: " + this.engine);
}
}

// 创建类的实例
var car = new Car("V8 发动机");

// 访问字段
console.log("读取发动机: " + car.engine);

// 调用方法
car.disp();

运行结果：

```
读取发动机: V8 发动机
发动机型号: V8 发动机

```

说明：

- `this` 关键字指向当前类的实例
- 构造函数的参数名可以与字段名相同，通过 `this.field` 区分
- 使用 `new` 关键字创建类的实例

注意：TypeScript 的类在编译后会转换为 JavaScript 的构造函数原型模式，接口不会出现在编译结果中。

### 访问控制修饰符

TypeScript 提供了三种访问修饰符来控制类成员的可访问性：

修饰符 说明 `public` 公有成员，可以在任何地方访问（默认） `private` 私有成员，只能在类内部访问 `protected` 受保护成员，可以在类内部和子类中访问

#### public（默认）

### 实例

class Person {
public name: string; // 公有属性
public age: number; // 公有属性

constructor(name: string, age: number) {
this.name = name;
this.age = age;
}

public introduce(): void {
console.log("我是 " + this.name + "，今年 " + this.age + " 岁");
}
}

var person = new Person("Alice", 25);
console.log("姓名: " + person.name); // 可以访问
person.introduce(); // 可以访问

#### private（私有）

### 实例

class Person {
public name: string;
private secret: string; // 私有属性，外部无法直接访问

constructor(name: string, secret: string) {
this.name = name;
this.secret = secret;
}

// 公有方法可以访问私有属性
public revealSecret(): void {
console.log("秘密: " + this.secret);
}
}

var person = new Person("Alice", "我喜欢编程");

console.log("姓名: " + person.name); // 可以访问
// console.log(person.secret); // 错误：'secret' 是私有属性

person.revealSecret(); // 通过公有方法访问私有属性

#### protected（受保护）

### 实例

class Person {
protected name: string; // 受保护属性

constructor(name: string) {
this.name = name;
}

protected sayHello(): void {
console.log("你好，我是 " + this.name);
}
}

class Student extends Person {
private grade: string;

constructor(name: string, grade: string) {
super(name);
this.grade = grade;
}

public introduce(): void {
// 子类可以访问受保护的属性和方法
console.log("我是 " + this.name + "，年级: " + this.grade);
this.sayHello();
}
}

var student = new Student("Bob", "高三");
student.introduce(); // 可以访问

// console.log(student.name); // 错误：'name' 是受保护属性

运行结果：

```
我是 Bob，年级: 高三
你好，我是 Bob

```

### 类的继承

继承允许创建一个类（子类）从另一个类（父类）获取属性和方法。子类可以复用父类的代码，还可以扩展或重写父类的行为。

#### 基本语法

```

class child_class extends parent_class {
// 子类新增的属性和方法
}

```

#### 单继承

### 实例

// 父类：形状
class Shape {
area: number;

constructor(a: number) {
this.area = a;
}
}

// 子类：圆，继承自 Shape
class Circle extends Shape {
disp(): void {
console.log("圆的面积: " + this.area);
}
}

var circle = new Circle(223);
circle.disp();

运行结果：

```
圆的面积: 223

```

#### 多重继承

TypeScript 不支持多继承（一个类继承多个类），但支持多层继承（A 继承 B，B 继承 C）：

### 实例

// 根类
class Root {
str: string;
}

// 子类：继承 Root
class Child extends Root {
}

// 叶子类：继承 Child（多重继承）
class Leaf extends Child {
}

var leaf = new Leaf();
leaf.str = "hello";
console.log("str 值: " + leaf.str);

运行结果：

```
str 值: hello

```

### 方法重写（Override）

子类可以重写（Override）父类的方法，即在子类中定义与父类同名的方法，实现自己的行为。

使用 `super` 关键字可以调用父类的方法。

### 实例

// 父类
class PrinterClass {
doPrint(): void {
console.log("父类的 doPrint() 方法");
}
}

// 子类：重写父类方法
class StringPrinter extends PrinterClass {
doPrint(): void {
// 调用父类的方法
super.doPrint();

// 子类自己的逻辑
console.log("子类的 doPrint() 方法");
}
}

var obj = new StringPrinter();
obj.doPrint();

运行结果：

```
父类的 doPrint() 方法
子类的 doPrint() 方法

```

### 静态成员

使用 `static` 关键字定义的成员属于类本身，而不是类的实例。可以直接通过类名访问，不需要创建实例。

### 实例

class StaticMem {
// 静态属性
static num: number;

// 静态方法
static disp(): void {
console.log("num 值为 " + StaticMem.num);
}
}

// 直接通过类名访问静态成员
StaticMem.num = 12;
StaticMem.disp();

运行结果：

```
num 值为 12

```

应用场景：静态成员常用于定义类的常量、工具方法或单例模式。

### instanceof 运算符

`instanceof` 用于判断对象是否是某个类的实例。

### 实例

class Person {
}

var obj = new Person();
var isPerson = obj instanceof Person;

console.log("obj 是 Person 类的实例吗？ " + isPerson);

运行结果：

```
obj 是 Person 类的实例吗？ true

```

### 类实现接口

类可以使用 `implements` 关键字实现接口，确保类符合接口定义的契约。

### 实例

// 定义接口
interface ILoan {
interest: number; // 利率
}

// 类实现接口
class AgriLoan implements ILoan {
interest: number;
rebate: number; // 回扣

constructor(interest: number, rebate: number) {
this.interest = interest;
this.rebate = rebate;
}
}

var loan = new AgriLoan(10, 1);
console.log("利率: " + loan.interest + "%，回扣: " + loan.rebate);

运行结果：

```
利率: 10%，回扣: 1

```

### 抽象类

抽象类不能被实例化，只能作为基类供子类继承。抽象类可以包含抽象方法（没有实现的占位方法），子类必须实现这些方法。

### 实例

// 抽象类
abstract class Animal {
abstract makeSound(): void; // 抽象方法，子类必须实现

move(): void {
console.log("动物在移动");
}
}

// 具体类：继承抽象类
class Dog extends Animal {
makeSound(): void {
console.log("汪汪汪!");
}
}

var dog = new Dog();
dog.move();
dog.makeSound();

运行结果：

```
动物在移动
汪汪汪!

```

### 综合实例

综合运用类的各种特性：

### 实例

// 接口定义
interface I Printable {
print(): void;
}

// 抽象类
abstract class Item {
protected name: string;
protected price: number;

constructor(name: string, price: number) {
this.name = name;
this.price = price;
}

abstract getDetails(): string;
}

// 具体类
class Product extends Item implements I Printable {
private category: string;

constructor(name: string, price: number, category: string) {
super(name, price);
this.category = category;
}

// 实现抽象方法
getDetails(): string {
return `产品: ${this.name}, 价格: ¥${this.price}, 类别: ${this.category}`;
}

// 实现接口方法
print(): void {
console.log(this.getDetails());
}

// 静态方法
static create(name: string, price: number): Product {
return new Product(name, price, "默认类别");
}
}

// 使用
var product = Product.create("笔记本电脑", 5999);
product.print();

// 折扣方法
class DiscountedProduct extends Product {
private discount: number;

constructor(name: string, price: number, category: string, discount: number) {
super(name, price, category);
this.discount = discount;
}

getDetails(): string {
var discountedPrice = this.price * (1 - this.discount / 100);
return `产品: ${this.name}, 原价: ¥${this.price}, 折扣: ${this.discount}%, 现价: ¥${discountedPrice.toFixed(2)}`;
}
}

var discountedProduct = new DiscountedProduct("手机", 2999, "电子产品", 20);
discountedProduct.print();

运行结果：

```
产品: 笔记本电脑, 价格: ¥5999, 类别: 默认类别
产品: 手机, 原价: ¥2999, 折扣: 20%, 现价: ¥2399.20

```

### 总结

TypeScript 的类提供了完整的面向对象编程支持：

- 类的定义：使用 class 关键字，包含字段、构造函数和方法
- 访问修饰符：public、private、protected 控制成员访问权限
- 继承：使用 extends 实现类继承，支持多层继承
- 方法重写：子类可以重写父类方法，使用 super 调用父类
- 静态成员：使用 static 关键字，属于类本身而非实例
- 接口实现：使用 implements 关键字实现接口
- 抽象类：使用 abstract 关键字，不能实例化，作为基类

类是 TypeScript 面向对象编程的基础，合理使用类可以使代码更加结构化和可维护。

---

## TypeScript 对象

Source: https://www.runoob.com/typescript/ts-object.html

## TypeScript 对象

对象是包含一组键值对的实例。 值可以是标量、函数、数组、对象等，如下实例：

var object_name = { key1: "value1", // 标量 key2: "value", key3: function() { // 函数 }, key4:["content1", "content2"] //集合 }

以上对象包含了标量，函数，集合(数组或元组)。

#### 对象实例

### TypeScript

var sites = { site1:"Runoob", site2:"Google" }; // 访问对象的值 console.log(sites.site1) console.log(sites.site2)

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var sites = { site1:"Runoob", site2:"Google" }; // 访问对象的值 console.log(sites.site1) console.log(sites.site2)

输出结果为：

```
Runoob
Google
```

### TypeScript 类型模板

假如我们在 JavaScript 定义了一个对象：

var sites = { site1:"Runoob", site2:"Google" };

这时如果我们想在对象中添加方法，可以做以下修改：

```
sites.sayHello = function(){ return "hello";}
```

如果在 TypeScript 中使用以上方式则会出现编译错误，因为Typescript 中的对象必须是特定类型的实例。

### TypeScript

var sites = { site1: "Runoob", site2: "Google", sayHello: function () { } // 类型模板 }; sites.sayHello = function () { console.log("hello " + sites.site1); }; sites.sayHello();

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var sites = { site1: "Runoob", site2: "Google", sayHello: function () { } // 类型模板 }; sites.sayHello = function () { console.log("hello " + sites.site1); }; sites.sayHello();

输出结果为：

```
hello Runoob
```

此外对象也可以作为一个参数传递给函数，如下实例：

### TypeScript

var sites = { site1:"Runoob", site2:"Google", }; var invokesites = function(obj: { site1:string, site2 :string }) { console.log("site1 :"+obj.site1) console.log("site2 :"+obj.site2) } invokesites(sites)

编译以上代码，得到以下 JavaScript 代码：

### JavaScript

var sites = { site1: "Runoob", site2: "Google" }; var invokesites = function (obj) { console.log("site1 :" + obj.site1); console.log("site2 :" + obj.site2); }; invokesites(sites);

输出结果为：

```
site1 :Runoob
site2 :Google
```

### 鸭子类型(Duck Typing)

鸭子类型（英语：duck typing）是动态类型的一种风格，是多态(polymorphism)的一种形式。

在这种风格中，一个对象有效的语义，不是由继承自特定的类或实现特定的接口，而是由"当前方法和属性的集合"决定。

可以这样表述：

"当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子。"

在鸭子类型中，关注点在于对象的行为能做什么，而不是关注对象所属的类型。例如，在不使用鸭子类型的语言中，我们可以编写一个函数，它接受一个类型为"鸭子"的对象，并调用它的"走"和"叫"方法。在使用鸭子类型的语言中，这样的一个函数可以接受一个任意类型的对象，并调用它的"走"和"叫"方法。如果这些需要被调用的方法不存在，那么将引发一个运行时错误。任何拥有这样的正确的"走"和"叫"方法的对象都可被函数接受的这种行为引出了以上表述，这种决定类型的方式因此得名。

interface IPoint { x:number y:number } function addPoints(p1:IPoint,p2:IPoint):IPoint { var x = p1.x + p2.x var y = p1.y + p2.y return {x:x,y:y} } // 正确 var newPoint = addPoints({x:3,y:4},{x:5,y:1}) // 错误 var newPoint2 = addPoints({x:1},{x:4,y:3})

---

## TypeScript 泛型

Source: https://www.runoob.com/typescript/ts-generics.html

## TypeScript 泛型

泛型（Generics）是一种编程语言特性，允许在定义函数、类、接口等时使用占位符来表示类型，而不是具体的类型。

泛型是一种在编写可重用、灵活且类型安全的代码时非常有用的功能。

使用泛型的主要目的是为了处理不特定类型的数据，使得代码可以适用于多种数据类型而不失去类型检查。

泛型的优势包括：

- 代码重用： 可以编写与特定类型无关的通用代码，提高代码的复用性。
- 类型安全： 在编译时进行类型检查，避免在运行时出现类型错误。
- 抽象性： 允许编写更抽象和通用的代码，适应不同的数据类型和数据结构。

#### 泛型标识符

在泛型中，通常使用一些约定俗成的标识符，比如常见的 `T`（表示 Type）、`U`、`V` 等，但实际上你可以使用任何标识符。

T: 代表 "Type"，是最常见的泛型类型参数名。

```
function identity<T>(arg: T): T {
return arg;
}
```

K, V: 用于表示键（Key）和值（Value）的泛型类型参数。

```

interface KeyValuePair<K, V> {
key: K;
value: V;
}
```

E: 用于表示数组元素的泛型类型参数。

```

function printArray<E>(arr: E[]): void {
arr.forEach(item => console.log(item));
}
```

R: 用于表示函数返回值的泛型类型参数。

```
function getResult<R>(value: R): R {
return value;
}
```

U, V: 通常用于表示第二、第三个泛型类型参数。

```

function combine<U, V>(first: U, second: V): string {
return `${first} ${second}`;
}
```

这些标识符是约定俗成的，实际上你可以选择任何符合标识符规范的名称。关键是使得代码易读和易于理解，所以建议在泛型类型参数上使用描述性的名称，以便于理解其用途。

#### 泛型函数（Generic Functions）

使用泛型来创建一个可以处理不同类型的函数：

### 实例

function identity<T>(arg: T): T {
return arg;
}

// 使用泛型函数
let result = identity<string>("Hello");
console.log(result); // 输出: Hello

let numberResult = identity<number>(42);
console.log(numberResult); // 输出: 42

解析： 以上例子中，`identity` 是一个泛型函数，使用 `<T>` 表示泛型类型。它接受一个参数 `arg` 和返回值都是泛型类型 `T`。在使用时，可以通过尖括号 `<>` 明确指定泛型类型。第一个调用指定了 `string` 类型，第二个调用指定了 `number` 类型。

#### 2. 泛型接口（Generic Interfaces）

可以使用泛型来定义接口，使接口的成员能够使用任意类型：

### 实例

// 基本语法
interface Pair<T, U> {
first: T;
second: U;
}

// 使用泛型接口
let pair: Pair<string, number> = { first: "hello", second: 42 };
console.log(pair); // 输出: { first: 'hello', second: 42 }

解析： 这里定义了一个泛型接口 `Pair`，它有两个类型参数 `T` 和 `U`。然后，使用这个泛型接口创建了一个对象 `pair`，其中 `first` 是字符串类型，`second` 是数字类型。

#### 3. 泛型类（Generic Classes） 泛型也可以应用于类的实例变量和方法：

### 实例

// 基本语法
class Box<T> {
private value: T;

constructor(value: T) {
this.value = value;
}

getValue(): T {
return this.value;
}
}

// 使用泛型类
let stringBox = new Box<string>("TypeScript");
console.log(stringBox.getValue()); // 输出: TypeScript

解析： 在这个例子中，`Box` 是一个泛型类，使用 `<T>` 表示泛型类型。构造函数和方法都可以使用泛型类型 `T`。通过实例化 `Box<string>`，我们创建了一个存储字符串的 `Box` 实例，并通过 `getValue` 方法获取了存储的值。

#### 4. 泛型约束（Generic Constraints）

有时候你想限制泛型的类型范围，可以使用泛型约束：

### 实例

// 基本语法
interface Lengthwise {
length: number;
}

function logLength<T extends Lengthwise>(arg: T): void {
console.log(arg.length);
}

// 正确的使用
logLength("hello"); // 输出: 5

// 错误的使用，因为数字没有 length 属性
logLength(42); // 错误

解析： 在这个例子中，定义了一个泛型函数 `logLength`，它接受一个类型为 `T` 的参数，但有一个约束条件，即 `T` 必须实现 `Lengthwise` 接口，该接口要求有 `length` 属性。因此，可以正确调用 `logLength("hello")`，但不能调用 `logLength(42)`，因为数字没有 `length` 属性。

#### 5. 泛型与默认值

可以给泛型设置默认值，使得在不指定类型参数时能够使用默认类型：

### 实例

// 基本语法
function defaultValue<T = string>(arg: T): T {
return arg;
}

// 使用带默认值的泛型函数
let result1 = defaultValue("hello"); // 推断为 string 类型
let result2 = defaultValue(42); // 推断为 number 类型

说明： 这个例子展示了带有默认值的泛型函数。函数 `defaultValue` 接受一个泛型参数 `T`，并给它设置了默认类型为 `string`。在使用时，如果没有显式指定类型，会使用默认类型。在例子中，第一个调用中 `result1` 推断为 `string` 类型，第二个调用中 `result2` 推断为 `number` 类型。

---

## TypeScript 命名空间

Source: https://www.runoob.com/typescript/ts-namespace.html

## TypeScript 命名空间

命名空间一个最明确的目的就是解决重名问题。

假设这样一种情况，当一个班上有两个名叫小明的学生时，为了明确区分它们，我们在使用名字之外，不得不使用一些额外的信息，比如他们的姓（王小明，李小明），或者他们父母的名字等等。

命名空间定义了标识符的可见范围，一个标识符可在多个命名空间中定义，它在不同命名空间中的含义是互不相干的。这样，在一个新的命名空间中可定义任何标识符，它们不会与任何已有的标识符发生冲突，因为已有的定义都处于其他命名空间中。

TypeScript 中命名空间使用 namespace 来定义，语法格式如下：

namespace SomeNameSpaceName { export interface ISomeInterfaceName { } export class SomeClassName { } }

以上定义了一个命名空间 SomeNameSpaceName，如果我们需要在外部可以调用 SomeNameSpaceName 中的类和接口，则需要在类和接口添加 export 关键字。

要在另外一个命名空间调用语法格式为：

```
SomeNameSpaceName.SomeClassName;
```

如果一个命名空间在一个单独的 TypeScript 文件中，则应使用三斜杠 /// 引用它，语法格式如下：

```
/// <reference path = "SomeFileName.ts" />
```

以下实例演示了命名空间的使用，定义在不同文件中：

### IShape.ts 文件代码：

namespace Drawing { export interface IShape { draw(); } }

### Circle.ts 文件代码：

/// <reference path = "IShape.ts" /> namespace Drawing { export class Circle implements IShape { public draw() { console.log("Circle is drawn"); } } }

### Triangle.ts 文件代码：

/// <reference path = "IShape.ts" /> namespace Drawing { export class Triangle implements IShape { public draw() { console.log("Triangle is drawn"); } } }

### TestShape.ts 文件代码：

/// <reference path = "IShape.ts" /> /// <reference path = "Circle.ts" /> /// <reference path = "Triangle.ts" /> function drawAllShapes(shape:Drawing.IShape) { shape.draw(); } drawAllShapes(new Drawing.Circle()); drawAllShapes(new Drawing.Triangle());

使用 tsc 命令编译以上代码：

```
tsc --out app.js TestShape.ts
```

得到以下 JavaScript 代码：

### JavaScript

/// <reference path = "IShape.ts" /> var Drawing; (function (Drawing) { var Circle = /** @class */ (function () { function Circle() { } Circle.prototype.draw = function () { console.log("Circle is drawn"); }; return Circle; }()); Drawing.Circle = Circle; })(Drawing || (Drawing = {})); /// <reference path = "IShape.ts" /> var Drawing; (function (Drawing) { var Triangle = /** @class */ (function () { function Triangle() { } Triangle.prototype.draw = function () { console.log("Triangle is drawn"); }; return Triangle; }()); Drawing.Triangle = Triangle; })(Drawing || (Drawing = {})); /// <reference path = "IShape.ts" /> /// <reference path = "Circle.ts" /> /// <reference path = "Triangle.ts" /> function drawAllShapes(shape) { shape.draw(); } drawAllShapes(new Drawing.Circle()); drawAllShapes(new Drawing.Triangle());

使用 node 命令查看输出结果为：

```
$ node app.js
Circle is drawn
Triangle is drawn
```

### 嵌套命名空间

命名空间支持嵌套，即你可以将命名空间定义在另外一个命名空间里头。

namespace namespace_name1 { export namespace namespace_name2 { export class class_name { } } }

成员的访问使用点号 . 来实现，如下实例：

### Invoice.ts 文件代码：

namespace Runoob { export namespace invoiceApp { export class Invoice { public calculateDiscount(price: number) { return price * .40; } } } }

### InvoiceTest.ts 文件代码：

/// <reference path = "Invoice.ts" /> var invoice = new Runoob.invoiceApp.Invoice(); console.log(invoice.calculateDiscount(500));

使用 tsc 命令编译以上代码：

```
tsc --out app.js InvoiceTest.ts
```

得到以下 JavaScript 代码：

### JavaScript

var Runoob; (function (Runoob) { var invoiceApp; (function (invoiceApp) { var Invoice = /** @class */ (function () { function Invoice() { } Invoice.prototype.calculateDiscount = function (price) { return price * .40; }; return Invoice; }()); invoiceApp.Invoice = Invoice; })(invoiceApp = Runoob.invoiceApp || (Runoob.invoiceApp = {})); })(Runoob || (Runoob = {})); /// <reference path = "Invoice.ts" /> var invoice = new Runoob.invoiceApp.Invoice(); console.log(invoice.calculateDiscount(500));

使用 node 命令查看输出结果为：

```
$ node app.js
200
```

---

## TypeScript 模块系统

Source: https://www.runoob.com/typescript/ts-module.html

## TypeScript 模块系统

模块系统是现代 TypeScript 开发的基础。

TypeScript 完全支持 ES Module 语法，并提供了丰富的模块解析策略。

通过模块系统，可以将代码分割成可重用的单元，实现代码组织和复用。 模块导出导入流程 user.ts export const name = "Alice" export class User { ... } export interface Config { ... import main.ts import { name, User } from "./user"; import User from "./user" use 运行结果 JavaScript 执行代码 导出导入方式 命名导出/导入 默认导出/导入 重新导出 动态导入

### 为什么需要模块系统

随着项目规模增长，代码量会越来越大。

将代码分散到多个文件中，通过模块组织，可以提高代码的可维护性和可复用性。

模块系统让每个文件都有自己的作用域，避免全局变量污染。

概念说明：模块是包含导出和导入语句的 TypeScript 文件。通过 export 导出内容，通过 import 导入内容。

### 模块导出

使用 export 关键字可以将变量、函数、类、接口等导出供其他模块使用。

### user.ts 模块

// 导出变量
export var name = "Alice";
export const age = 25;

// 导出函数
export function greet(message: string): string {
return "Hello, " + message;
}

// 导出类
export class User {
// 构造函数参数属性
constructor(public name: string) {}

// 自我介绍方法
introduce(): string {
return "I am " + this.name;
}
}

// 导出接口（接口在编译后会消失，仅用于类型检查）
export interface Config {
// 配置主机
host: string;
// 配置端口
port: number;
}

// 批量导出：重命名导出
export { name as userName, age as userAge };

注意：接口和类型在编译后的 JavaScript 中不会产生实际代码，它们仅用于 TypeScript 的类型检查。

### 模块导入

使用 import 关键字从其他模块导入导出的内容。

### main.ts 导入模块

// 命名导入：从模块中导入指定的内容
import { name, age, greet } from "./user";

// 默认导入：导入模块的默认导出
import User from "./user";

// 全部导入：将模块所有导出放入一个对象
import * as UserModule from "./user";

// 重命名导入：避免命名冲突
import { greet as sayHello } from "./user";

// 使用导入的内容
console.log(greet("World"));
console.log(sayHello("TypeScript"));

运行结果：

```
Hello, World
Hello, TypeScript

```

路径说明：导入路径可以是相对路径（如 `./user`）或绝对路径（如 `@/utils`）。

### 默认导出

每个模块可以有一个默认导出。

默认导出在导入时不需要使用花括号，且可以取任意名字。

### math.ts 模块

// 默认导出：一个模块只能有一个默认导出
export default function add(a: number, b: number): number {
return a + b;
}

// 可以和其他导出混合使用
export function multiply(a: number, b: number): number {
return a * b;
}

### main.ts 导入

// 导入默认导出：可以取任意名字
import add from "./math";

// 导入命名导出：需要使用花括号
import { multiply } from "./math";

console.log("加法: " + add(2, 3));
console.log("乘法: " + multiply(4, 5));

运行结果：

```
加法: 5
乘法: 20

```

建议：对于工具函数、类等主要导出内容使用默认导出，对于辅助函数、接口等使用命名导出。

### 重新导出

重新导出（Re-export）用于聚合多个模块的内容，或将一个模块的导出暴露给另一个模块。

### index.ts 聚合模块

// 从其他模块重新导出指定内容
export { name, age } from "./user";

// 重新导出默认导出（需要重命名）
export { default as User } from "./user";

// 重新导出所有内容
export * from "./math";

应用场景：使用 index.ts 作为入口文件，集中导出子模块的内容，方便统一导入。

### 模块解析策略

TypeScript 提供了多种模块解析策略，用于查找导入的模块。

可以在 tsconfig.json 中配置。

### tsconfig.json 配置

{
"compilerOptions": {
// Node 解析策略
// 遵循 Node.js 的模块解析规则
"moduleResolution": "node",

// 经典解析策略
// TypeScript 早期版本使用的策略
"moduleResolution": "classic",

// base URL：设置基础路径
// 所有非相对路径导入都基于此路径解析
"baseUrl": "./src",

// 路径映射：为导入路径设置别名
"paths": {
// @ 开头的导入映射到 src 目录
"@/*": ["./*"],
// @components 开头的导入映射到 components 目录
"@components/*": ["./components/*"]
}
}
}

配置建议：新项目推荐使用 Node 解析策略，它是目前最常用的方式。

### 动态导入

动态导入（Dynamic Import）使用 import() 语法，可以在运行时按需加载模块。

这对于代码分割、懒加载非常有用。

### 实例

// 动态导入 - 懒加载
// import() 返回一个 Promise
async function loadMath() {
// 动态导入模块，只有执行到这里才会加载
var math = await import("./math");

// math.default 是默认导出的函数
console.log("动态加法: " + math.default(1, 2));
}

// 调用懒加载函数
loadMath();

// 条件导入：根据条件动态加载不同模块
async function loadFeature(enable: boolean) {
if (enable) {
// 只有条件满足时才加载模块
var feature = await import("./feature");
feature.run();
}
}

// 根据条件加载
loadFeature(true);

运行结果：

```
动态加法: 3

```

性能优化：动态导入可以实现代码分割，只在需要时加载额外的代码，减少初始加载时间。

### 注意事项

- 相对路径：使用相对路径导入本地模块（./、../）
- 模块扩展名：TypeScript 编译时会自动处理扩展名
- 默认 vs 命名：每个模块一个默认导出，多个命名导出
- esModuleInterop：启用此选项可以更方便地导入 CommonJS 模块

最佳实践：保持导入路径一致，使用路径别名简化长路径，建立清晰的模块组织结构。

### 总结

模块系统是 TypeScript 项目组织的核心。

- export：导出变量、函数、类、接口
- import：导入已导出的内容
- 默认导出：每个模块一个，使用灵活
- 重新导出：聚合模块，创建入口文件
- 动态导入：懒加载，优化性能
- 模块解析：配置路径别名和解析策略

建议：合理组织模块结构，使用路径别名简化导入，建立清晰的导出导入规范。

---

## TypeScript 声明文件

Source: https://www.runoob.com/typescript/ts-ambient.html

## TypeScript 声明文件

声明文件是 TypeScript 与 JavaScript 库之间的翻译官，它告诉 TypeScript 一个 JavaScript 库暴露了哪些功能、参数是什么类型、返回值是什么类型。

### 先从一个问题说起

假设你正在用 TypeScript 写项目，需要引入一个第三方的 JavaScript 库（比如 jQuery），你可能会写出这样的代码：

```

$('#foo');
// 或
jQuery('#foo');

```

这段代码在纯 JavaScript 项目中完全没问题，但在 TypeScript 文件中会直接报错：

```

jQuery('#foo');

// index.ts(1,1): error TS2304: Cannot find name 'jQuery'.

```

为什么会报错？因为 TypeScript 不认识 $ 和 jQuery 是什么。

TypeScript 的核心功能就是类型检查——它在编译阶段就要知道每个变量、每个函数的类型。但 jQuery 是一个纯 JavaScript 库，没有任何类型信息，TypeScript 自然看不懂它。

### 快速修复：declare 关键字

最简单的方式是用 declare 关键字手动告诉 TypeScript：「这个变量存在，它的类型是这样的」：

```

declare var jQuery: (selector: string) => any;

jQuery('#foo');

```

这行代码的含义是：声明一个变量 jQuery，它是一个函数，接收一个 string 类型的参数，返回值可以是任意类型（any）。

现在 TypeScript 就不会报错了，因为编译器已经知道 jQuery 是什么类型。

declare 关键字声明的类型只在编译阶段起作用，编译后的 JavaScript 代码中会被完全删除，不会影响运行时的行为。

上例编译后的 JavaScript 代码为：

```

jQuery('#foo');

```

可以看到，declare 语句消失了，只剩下了实际的调用代码。

但 declare 只能解决单个文件的临时问题。如果一个库有很多方法、很多类，在每个文件里都手写 declare 显然不现实。这就是声明文件的用武之地。

### 声明文件：一劳永逸的方案

声明文件就是把所有 declare 声明集中放到一个单独的文件里，项目中的任何 TypeScript 文件都可以引用它。

#### 文件命名规范

声明文件统一以 .d.ts 为后缀，d 代表 declaration（声明）。例如：

```

runoob.d.ts

```

#### 基本语法

声明一个模块的语法如下：

```

declare module Module_Name {
}

```

在 TypeScript 文件中通过三斜线指令引入声明文件：

```

/// <reference path = "runoob.d.ts" />

```

三斜线指令是 TypeScript 特有的语法，用于告诉编译器在编译时需要包含指定的声明文件。

很多流行第三方库（如 jQuery、Lodash）的声明文件已经由社区维护好了，存放在 DefinitelyTyped 项目中。你只需要通过 npm 安装对应的 @types/xxx 包即可使用，无需手动编写。

### 完整实例：从零创建声明文件

下面通过一个完整的例子，演示声明文件的创建和使用全流程。

整个流程涉及以下文件：

文件作用 CalcThirdPartyJsLib.js第三方 JavaScript 库（纯 JS，无类型信息） Calc.d.ts声明文件（手动编写，描述库的类型） CalcTest.tsTypeScript 业务代码（引用声明文件，调用库） CalcTest.js编译产物（tsc 编译后的 JS 文件） runoob.html浏览器中运行的最终页面

#### 第一步：创建第三方 JavaScript 库

假设我们有一个第三方库，提供了累加求和的功能。它使用命名空间 Runoob 来组织代码，避免变量名冲突：

### CalcThirdPartyJsLib.js 文件代码：

// 文件路径：CalcThirdPartyJsLib.js
// 这是一个模拟的第三方 JavaScript 库

// 声明 Runoob 命名空间变量（如果不存在则创建空对象）
var Runoob;

// 使用立即执行函数（IIFE）封装代码，避免内部变量泄漏到全局作用域
(function(Runoob) {
// Calc 构造函数，用于创建计算器对象
var Calc = (function () {
function Calc() {
// 当前无需初始化参数，保留构造函数以便后续扩展
}
})

// doSum 方法：计算从 0 到 limit 的所有整数之和
// limit（必填）：累加的上限值，包含该值本身
// 示例：limit=10 时，计算 0+1+2+...+10 = 55
Calc.prototype.doSum = function (limit) {
var sum = 0;

for (var i = 0; i <= limit; i++) {
sum = sum + i;
}
return sum;
}

// 将 Calc 构造函数挂载到 Runoob 命名空间下
// 这样外部就可以通过 new Runoob.Calc() 来创建实例
Runoob.Calc = Calc;
return Calc;
})(Runoob || (Runoob = {}));

// 库内部自测代码
var test = new Runoob.Calc();

这个库用到了一个常见模式：立即执行函数（IIFE）。通俗理解，就是把代码包在一个函数里立即运行，这样函数里定义的变量不会跑到外面去，避免和其他代码的变量名冲突。

#### 第二步：编写声明文件

现在我们有了 JS 库，但它没有任何类型信息。需要创建一个声明文件，只描述库的「形状」——有哪些类、有哪些方法、参数和返回值各是什么类型——但不包含任何实际代码逻辑：

### Calc.d.ts 文件代码：

// 文件路径：Calc.d.ts
// 这是声明文件，只包含类型信息，不包含任何可执行代码

// 声明 Runoob 模块，与 JS 库中的 Runoob 命名空间对应
declare module Runoob {
// 声明 Calc 类，告诉 TypeScript 这个类可以被 new 实例化
export class Calc {
// 声明 doSum 方法：接收一个 number 参数，返回一个 number
// 注意：这里只声明了方法的签名，没有方法体（大括号）
doSum(limit: number): number;
}
}

声明文件和普通 .ts 文件的最大区别：声明文件只有类型签名，没有实现代码。它只回答「这个函数长什么样」，不回答「这个函数做了什么」。

#### 第三步：在 TypeScript 代码中使用

现在编写业务代码，引用声明文件后，就可以像使用原生 TypeScript 库一样使用这个第三方 JS 库了：

### CalcTest.ts 文件代码：

// 文件路径：CalcTest.ts
// 三斜线指令：告诉 TypeScript 编译器引入声明文件
/// <reference path = "Calc.d.ts" />

// 创建 Calc 实例，TypeScript 现在能正确识别 obj 的类型
var obj = new Runoob.Calc();

// obj.doSum("Hello"); // 编译错误！"Hello" 是字符串，而 doSum 要求传入 number
console.log(obj.doSum(10)); // 正确调用：传入 10，期望得到 55

注意被注释掉的那一行：

```

obj.doSum("Hello");

```

如果取消这行的注释，TypeScript 编译时会直接报错，因为声明文件中已经规定了 doSum 只接受 number 类型参数。这就是类型检查在保护你——错误在写代码时就暴露出来，而不是等到浏览器运行时才崩溃。

#### 第四步：编译 TypeScript

使用 TypeScript 自带的 tsc 命令编译：

```

tsc CalcTest.ts

```

编译后会生成 CalcTest.js 文件，内容如下：

### CalcTest.js 编译产物：

// 文件路径：CalcTest.js（tsc 命令自动生成）
/// <reference path = "Calc.d.ts" />
var obj = new Runoob.Calc();
//obj.doSum("Hello"); // 编译错误
console.log(obj.doSum(10)); // 运行结果：在控制台输出 55

可以看到，三斜线指令和声明文件中的类型信息在编译产物中都不见了，只剩下纯粹的 JavaScript 运行代码。

#### 第五步：在浏览器中运行

最后，创建一个 HTML 页面把所有 JS 文件串联起来：

### 实例

<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>菜鸟教程(runoob.com)</title>
<!-- 第 1 步：加载第三方库，让 Runoob 命名空间在全局可用 -->
<script src = "CalcThirdPartyJsLib.js"></script>
<!-- 第 2 步：加载编译后的业务代码 -->
<script src = "CalcTest.js"></script>
</head>
<body>
<h1>声明文件测试</h1>
<p>菜鸟测试一下。</p>
</body>
</html>

用浏览器打开这个 HTML 文件，打开控制台（F12），可以看到输出结果：

```

55

```

这就是 doSum(10) 的返回值：0+1+2+...+10 = 55。

### 小结

声明文件本质上是一份「类型说明书」，让 TypeScript 能够理解和检查纯 JavaScript 库。

整个工作流程可以概括为三步：

1. 拿到一个 JS 库，分析它暴露了哪些 API。

2. 编写 .d.ts 声明文件，描述这些 API 的类型签名。

3. 在 TypeScript 代码中引用声明文件，即可享受完整的类型检查保护。

对于日常开发中常用的第三方库，绝大多数已经有现成的声明文件（通过 npm install @types/xxx 安装）。只有当你用到非常冷门的库或内部私有库时，才需要手动编写声明文件。

---

## TypeScript 测验知识挑战

Source: https://www.runoob.com/typescript/ts-quiz.html

## TypeScript 测验

## 知识挑战

通过这些互动问题测试学习情况

### 测验完成！

0/0

重新开始

上一题 下一题

### 其他相关测试

1. TypeScript 测验一 2. TypeScript 测验二 3. TypeScript 测验三 4. TypeScript 测验四 5. TypeScript 测验五 6. TypeScript 测验六 7. TypeScript 测验七 8. TypeScript 测验八 9. TypeScript 测验九 10. TypeScript 测验十 11. TypeScript 测验十一

---

## TypeScript type 别名

Source: https://www.runoob.com/typescript/ts-type-alias.html

## TypeScript type 别名

type 别名（Type Alias）用于为现有类型创建别名，让代码更简洁、更易读。

通过类型别名，可以为复杂类型定义一个简短的名字，提高代码的可维护性。 type 别名工作原理 原始复杂类型 type User = { name: string; age: number; email: string; } 定义 type 别名 type UserID = User; // 或 type ID = string | number; 使用 代码中使用 var id: UserID; var uid: ID; type 别名应用场景 基础类型别名 联合类型别名 函数类型别名 泛型类型别名

### 为什么需要 type 别名

当代码中多次使用同一个复杂类型时，每次都写完整类型会很冗长。

type 别名可以为复杂类型定义一个简洁的名字，让代码更易读、更易维护。

特别是在使用联合类型、函数类型、元组等复杂类型时，type 别名非常有用。

概念说明：type 别名使用 `type` 关键字定义，它只是为现有类型起了一个新名字，不会创建新的类型。

### 基本用法

使用 type 关键字为类型定义别名。

### 实例

// 类型别名：为联合类型定义别名
// ID 可以是字符串或数字
type ID = string | number;

// 类型别名：为对象类型定义别名
// Point 表示一个坐标点
type Point = { x: number; y: number };

// 使用类型别名
var userId: ID = "123";
var productId: ID = 456;

// 使用 Point 类型别名
var point: Point = { x: 10, y: 20 };

console.log("用户ID: " + userId);
console.log("产品ID: " + productId);
console.log("坐标: " + JSON.stringify(point));

运行结果：

```
用户ID: 123
产品ID: 456
坐标: {"x":10,"y":20}

```

说明：类型别名只是给类型起了个新名字，编译后不会产生实际代码，它完全用于开发时的类型检查。

### 接口 vs 类型别名

类型别名和接口非常相似，都可以用来定义对象类型，但有一些细微区别。

### 实例

// 使用 type 别名定义对象类型
// type 可以定义任何类型，不仅仅是对象
type PersonType = {
name: string;
age: number;
};

// 使用接口定义对象类型
// 接口可以声明合并，可以被类实现
interface PersonInterface {
name: string;
age: number;
}

// 两者都可以用来声明变量类型
var person1: PersonType = { name: "Alice", age: 25 };
var person2: PersonInterface = { name: "Bob", age: 30 };

console.log("PersonType: " + JSON.stringify(person1));
console.log("PersonInterface: " + JSON.stringify(person2));

区别：type 可以定义任何类型（联合类型、元组、函数类型等），而接口主要用于定义对象类型。接口支持声明合并，type 不支持。

### 类型别名与联合类型

类型别名非常适合定义联合类型，让代码更清晰。

### 实例

// 联合类型别名：定义状态的可能值
// Status 只能是这三个字符串字面量之一
type Status = "pending" | "success" | "error";

// 联合类型别名：定义多种可能的返回类型
// Result 可以是字符串、数字或布尔值
type Result = string | number | boolean;

// 使用联合类型别名
function getStatus(status: Status): void {
console.log("状态: " + status);
}

getStatus("success");
getStatus("error");

// 使用 Result 类型
var result: Result = "hello";
result = 42;
console.log("结果: " + result);

运行结果：

```
状态: success
状态: error
结果: 42

```

应用场景：联合类型别名常用于定义状态码、错误类型、API 返回值等场景。

### 类型别名与元组

类型别名也可以用于定义元组类型。

### 实例

// 元组类型别名：定义坐标
// Coordinate 是一个包含两个数字的元组
type Coordinate = [number, number];

// 元组类型别名：定义姓名和年龄
// NameAge 是一个字符串和数字的元组
type NameAge = [string, number];

// 使用元组类型别名
var coord: Coordinate = [10, 20];
var person: NameAge = ["Alice", 25];

console.log("坐标: " + coord);
console.log("信息: " + person[0] + ", " + person[1]);

说明：元组类型别名让元组的使用更加清晰，避免了每次都需要写出完整元组类型的问题。

### 类型别名与函数

类型别名可以简化函数类型的声明。

### 实例

// 函数类型别名：定义回调函数类型
// Callback 接受一个字符串参数，返回 void
type Callback = (result: string) => void;

// 函数类型别名：定义数学运算函数类型
// MathOperation 接受两个数字参数，返回数字
type MathOperation = (a: number, b: number) => number;

// 使用函数类型别名
var add: MathOperation = function(a, b) { return a + b; };
var multiply: MathOperation = function(a, b) { return a * b; };

console.log("加法: " + add(2, 3));
console.log("乘法: " + multiply(4, 5));

运行结果：

```
加法: 5
乘法: 20

```

优势：函数类型别名让函数类型声明更简洁，特别是在需要多次使用相同函数签名时。

### 类型别名与泛型

类型别名可以配合泛型使用，创建可复用的类型定义。

### 实例

// 泛型类型别名：定义结果类型
// Result<T> 是一个泛型类型，T 是成功时的数据类型
type Result<T> = { success: boolean; data?: T; error?: string };

// 泛型类型别名：定义键值对类型
// Pair<K, V> 有两个类型参数
type Pair<K, V> = { key: K; value: V };

// 使用泛型类型别名
var result: Result<string> = { success: true, data: "Hello" };
var pair: Pair<string, number> = { key: "age", value: 25 };

console.log("结果: " + JSON.stringify(result));
console.log("键值对: " + JSON.stringify(pair));

运行结果：

```
结果: {"success":true,"data":"Hello"}
键值对: {"key":"age","value":25}

```

泛型优势：泛型类型别名可以适应不同的数据类型，提高代码的复用性。

### 类型别名与映射类型

类型别名可以与映射类型结合，创建强大的类型转换。

### 实例

// 映射类型：将所有属性变为只读
// Readonly<T> 遍历 T 的所有属性并添加 readonly
type Readonly<T> = { readonly [P in keyof T]: T[P] };

// 映射类型：将所有属性变为可选
// Partial<T> 遍历 T 的所有属性并添加 ?
type Partial<T> = { [P in keyof T]?: T[P] };

// 定义用户接口
interface User {
name: string;
age: number;
}

// 使用映射类型别名
var readonlyUser: Readonly<User> = { name: "Alice", age: 25 };
// readonlyUser.name = "Bob"; // 错误：只读属性不能修改

var partialUser: Partial<User> = { name: "Bob" };

console.log("只读用户: " + JSON.stringify(readonlyUser));
console.log("部分用户: " + JSON.stringify(partialUser));

提示：映射类型是 TypeScript 非常强大的特性，可以基于现有类型创建新的类型变化。

### 注意事项

- 不能重复定义：同一个 type 别名不能重复定义（接口可以声明合并）
- 仅编译时有效：type 别名在编译后不产生任何代码
- 可读性优先：不要滥用类型别名，简洁明了的代码更好

最佳实践：当类型需要复用，或者类型名称过长时使用 type 别名。对于简单的对象类型，可以根据团队习惯选择 type 或接口。

### 总结

type 别名是 TypeScript 中非常有用的特性。

- 基本用法：使用 type 关键字为复杂类型定义别名
- 联合类型：定义多个可能类型的组合
- 函数类型：简化函数类型声明
- 泛型：创建可复用的泛型类型
- 映射类型：基于现有类型创建新类型

建议：合理使用 type 别名可以让代码更清晰，但不要过度使用，保持代码简洁易读最重要。

---

## TypeScript 字面量类型

Source: https://www.runoob.com/typescript/ts-literal-types.html

## TypeScript 字面量类型

字面量类型（Literal Types）是 TypeScript 中一种表示特定值而不是通用类型的特性。

它允许我们将变量类型限制为具体的值，而不是宽泛的 string、number 等类型。

这提供了更精确的类型控制，使代码更加类型安全。 字面量类型示例 通用类型 var direction: string; 可以赋任何字符串 "up", "down", "abc"... 字面量类型 var direction: "up" | "down" | "left" | "right" 只能赋这四个值之一 类型安全 赋值错误值时 编译错误！ direction = "upup" 字面量类型种类 字符串字面量 "up" | "down" | "left" 限制特定字符串 数字字面量 1 | 2 | 3 | 100 | 200 限制特定数字 布尔字面量 true | false true 或 false 模板字面量 `on${string}` 动态生成

### 为什么需要字面量类型

在 TypeScript 中，默认的类型如 string 和 number 过于宽泛。

有时我们需要更精确的类型控制。例如，一个变量只能取特定的几个值。

字面量类型正是为了解决这类问题而生的。

概念说明：字面量类型是泛型类型（string、number）的具体值。例如 `"up"` 是 `string` 的子类型，只能赋值给它的值是特定的字符串。

### 字符串字面量类型

字符串字面量类型将变量的取值限制为特定的字符串。

这在需要限制变量只能取某些固定值时非常有用。

### 实例

// 定义字符串字面量类型
// direction 只能是四个特定值之一
var direction: "up" | "down" | "left" | "right";

// 正确：赋值为字面量类型中的值
direction = "up";

// 错误：赋值不在列表中的值
// direction = "upup"; // 编译错误！

console.log("方向: " + direction);

// 使用类型别名创建可复用的字面量类型
// 定义状态类型：只能是三个固定值之一
type Status = "pending" | "active" | "completed";

// 使用类型别名
var currentStatus: Status = "active";
console.log("状态: " + currentStatus);

运行结果：

```
方向: up
状态: active

```

应用场景：字符串字面量常用于定义枚举值、状态码、配置选项等需要限制为特定值的场景。

### 数字字面量类型

数字字面量类型将变量的取值限制为特定的数字。

这在需要使用特定数值（如状态码、错误码）时非常有用。

### 实例

// 定义数字字面量类型
// code 只能是三个特定值之一
var code: 200 | 404 | 500;

// 正确：赋值 200
code = 200;

// 错误：赋值不在列表中的值
// code = 301; // 编译错误！

console.log("状态码: " + code);

// 使用数字字面量模拟枚举
// 定义一周的天数
type Weekday = 1 | 2 | 3 | 4 | 5 | 6 | 7;

// 只能赋值 1-7 之间的数字
var today: Weekday = 1;
console.log("今天是星期: " + today);

说明：数字字面量可以替代简单的枚举（enum）使用，尤其在只需要几个固定数值时。

### 布尔字面量类型

布尔字面量类型将变量的取值限制为 true 或 false。

实际上，TypeScript 中的 boolean 类型就是 true | false 的别名。

### 实例

// 定义布尔字面量类型
// isActive 只能是 true 或 false
var isActive: true | false;

// 赋值 true
isActive = true;

// 实际上 boolean 就是 true | false 的联合类型
// 所以这两种写法是等价的
var flag: boolean = true;

console.log("是否激活: " + isActive);

提示：虽然 boolean 本身已经足够使用，但在需要明确区分 true 和 false 的类型时，可以使用 true 或 false 作为类型。

### 对象字面量类型

对象字面量类型可以定义对象的结构，并可以使用 readonly 修饰符将属性设为只读。

这在需要创建不可变对象时非常有用。

### 实例

// 定义对象字面量类型
// 指定对象的结构和属性类型
type Point = {
// 点的 x 坐标
x: number;
// 点的 y 坐标
y: number;
};

// 使用对象字面量类型
var p: Point = { x: 10, y: 20 };
console.log("点: " + JSON.stringify(p));

// 定义只读对象类型
// 使用 readonly 修饰符将属性设为只读
type ReadonlyPoint = {
// 只读的 x 坐标
readonly x: number;
// 只读的 y 坐标
readonly y: number;
};

// 使用只读对象类型
var rp: ReadonlyPoint = { x: 1, y: 2 };

// 尝试修改只读属性会报错
// rp.x = 3; // 编译错误：只读属性不能修改

console.log("只读点: " + JSON.stringify(rp));

运行结果：

```
点: {"x":10,"y":20}
只读点: {"x":1,"y":2}

```

应用场景：只读对象常用于定义配置对象、常量对象等不希望被修改的数据。

### 字面量类型与类型推断

TypeScript 会根据变量的声明方式自动推断为字面量类型。

使用 const 声明的变量会被推断为具体的字面量类型，而不是宽泛的类型。

### 实例

// 使用 var 声明并赋值
// TypeScript 会推断为宽泛类型
var name = "Alice"; // 类型: string（因为 var 可以重新赋值）
var age = 25; // 类型: number
var enabled = true; // 类型: boolean

console.log("名字: " + name + ", 年龄: " + age + ", 启用: " + enabled);

// 使用 as const 创建深度只读字面量
// 数组变为只读元组，值为字面量类型
var colors = ["red", "green", "blue"] as const;

// colors 的类型变为：
// readonly ["red", "green", "blue"]

// 访问数组元素
console.log("颜色: " + colors[0]);

// 尝试修改会报错
// colors[0] = "yellow"; // 编译错误：只读

说明：`as const` 会将类型推断为最具体的字面量类型，并添加 readonly 修饰符。

### 模板字面量类型

模板字面量类型使用 JavaScript 模板字符串的语法来创建新的类型。

它允许我们动态地生成字符串类型。

### 实例

// 使用模板字面量类型定义事件名称
// `on${string}` 表示以 "on" 开头的任意字符串
type EventName = `on${string}`;

// 使用模板字面量类型定义处理器名称
// Capitalize 将字符串首字母大写
type Handler = `handle${Capitalize<string>}`;

// 使用模板字面量类型
var event: EventName = "onClick";
var handler: Handler = "handleSubmit";

console.log("事件: " + event);
console.log("处理器: " + handler);

// 错误：不符合模板格式
// var e: EventName = "click"; // 编译错误：不是以 "on" 开头

运行结果：

```
事件: onClick
处理器: handleSubmit

```

内置工具类型：模板字面量类型可以配合 Uppercase、Lowercase、Capitalize、Uncapitalize 等内置工具类型使用。

### 实际应用：Redux Action

字面量类型在实际开发中有广泛的应用。

下面展示如何使用字面量类型实现 Redux 风格的行为（Action）。

### 实例

// 定义 Redux 风格的 Action 类型
// 使用联合类型定义不同类型的 Action
type Action =
// 增加类型的 Action，包含数字类型的 payload
| { type: "increment"; payload: number }
// 减少类型的 Action，包含数字类型的 payload
| { type: "decrement"; payload: number }
// 重置类型的 Action，没有 payload
| { type: "reset" };

// 使用 action 的 reducer 函数
function reducer(action: Action): void {
// TypeScript 会根据 action.type 的值自动推断 action 的具体类型
switch (action.type) {
case "increment":
// 在这个分支中，action 被推断为 { type: "increment"; payload: number }
console.log("增加: " + action.payload);
break;
case "decrement":
// 在这个分支中，action 被推断为 { type: "decrement"; payload: number }
console.log("减少: " + action.payload);
break;
case "reset":
// 在这个分支中，action 被推断为 { type: "reset" }
console.log("重置");
break;
}
}

// 派发不同类型的 Action
reducer({ type: "increment", payload: 5 });
reducer({ type: "reset" });

运行结果：

```
增加: 5
重置

```

可辨识联合：这种模式叫做"可辨识联合"（Discriminated Union），是 TypeScript 中处理多种相关类型的最佳实践。

### 注意事项

- 类型收窄：使用 switch 或 if 判断时，TypeScript 会自动收窄字面量类型
- 类型别名：建议使用 type 别名为常用的字面量类型创建别名
- const 推断：const 声明的变量会自动推断为字面量类型
- as const：需要深度只读时使用 as const

最佳实践：在需要限制变量只能取特定值时，优先使用字面量类型而不是枚举。

### 总结

字面量类型是 TypeScript 实现精确类型控制的重要特性。

- 字符串字面量：限制为特定字符串，如 `"up" | "down"`
- 数字字面量：限制为特定数字，如 `1 | 2 | 3`
- 布尔字面量：限制为 true 或 false
- 对象字面量：定义对象结构，可使用 readonly
- as const：创建深度只读字面量类型
- 模板字面量：使用模板字符串动态生成类型

进阶：字面量类型与可辨识联合模式结合，可以实现强大的类型安全。

---

## TypeScript 抽象类

Source: https://www.runoob.com/typescript/ts-abstract-class.html

## TypeScript 抽象类

抽象类（Abstract Class）是 TypeScript 面向对象编程中的重要概念。

它是一种不能被直接实例化的类，只能作为基类供其他类继承。

抽象类主要用于定义子类的公共属性和方法，为子类提供一个统一的结构和行为模板。 抽象类继承层次结构 abstract class Animal 不能直接实例化 class Dog extends Animal class Cat extends Animal 抽象类特性 abstract 方法：子类必须实现 具象方法：子类可以直接使用 可作为类型：接收子类实例

### 为什么需要抽象类

在面向对象编程中，我们经常需要定义一些基类。

这些基类描述了子类共同的属性和方法，但有些方法的实现细节需要由子类自己决定。

抽象类就是用来解决这个问题的：它允许我们定义方法的签名（声明），而把具体实现交给子类。

概念说明：抽象类是一种介于普通类和接口之间的类型。它既有接口的特征（定义方法签名），又有类的特征（可以有具体的方法实现和构造函数）。

### 抽象类基础

使用 `abstract` 关键字声明抽象类和抽象方法。

抽象类可以包含抽象方法和具象方法（已有实现的方法）。

### 实例

// 使用 abstract 关键字声明抽象类
// 抽象类不能直接实例化，只能作为基类
abstract class Animal {
// 动物的名字属性
name: string;

// 构造函数
constructor(name: string) {
this.name = name;
}

// 抽象方法：使用 abstract 修饰，没有方法体
// 子类必须实现这个方法
abstract speak(): void;

// 具象方法：有具体实现的方法
// 子类可以直接继承使用，不需要重写
move(): void {
console.log(this.name + " 在移动");
}
}

// 尝试实例化抽象类会报错
// var animal = new Animal("动物"); // 错误：不能实例化抽象类

// 定义 Dog 类继承 Animal
class Dog extends Animal {
// 子类必须实现父类的抽象方法 speak()
speak(): void {
console.log(this.name + " 汪汪汪!");
}
}

// 创建 Dog 实例
var dog = new Dog("旺财");
dog.speak(); // 调用子类实现的方法
dog.move(); // 继承父类的具象方法

运行结果：

```
旺财 汪汪汪!
旺财 在移动

```

说明：抽象方法没有方法体（只有方法签名），子类必须实现该方法。具象方法有具体实现，子类可以直接继承使用。

### 抽象方法

抽象方法是抽象类中只声明但不实现的方法。

子类继承抽象类后，必须实现所有抽象方法，否则会报错。

### 实例

// 定义抽象类 Shape（图形）
abstract class Shape {
// 抽象方法：只有声明，没有实现
// 子类必须实现这两个方法
abstract area(): number;
abstract perimeter(): number;

// 具象方法：可以使用抽象方法
// 这个方法调用了抽象方法，子类继承后可以正常使用
describe(): void {
console.log("面积: " + this.area().toFixed(2));
}
}

// 定义矩形类继承 Shape
class Rectangle extends Shape {
// 矩形的宽度
width: number;
// 矩形的高度
height: number;

// 构造函数
constructor(width: number, height: number) {
super(); // 调用父类构造函数
this.width = width;
this.height = height;
}

// 实现抽象方法：计算面积
area(): number {
return this.width * this.height;
}

// 实现抽象方法：计算周长
perimeter(): number {
return 2 * (this.width + this.height);
}
}

// 创建矩形实例
var rect = new Rectangle(4, 5);

// 调用继承的 describe 方法，内部会调用子类的 area 方法
rect.describe();
console.log("周长: " + rect.perimeter());

运行结果：

```
面积: 20.00
周长: 18

```

提示：具象方法可以调用抽象方法。这是因为当具象方法被调用时，子类已经实现了抽象方法，所以可以正常执行。

### 抽象类作为类型

抽象类可以作为参数类型使用。

这意味着函数可以接受任何抽象类的子类实例，实现了多态。

### 实例

// 定义抽象类 Animal
abstract class Animal {
// 抽象方法：所有动物都会发出声音
abstract speak(): void;
}

// 定义 Cat 类继承 Animal
class Cat extends Animal {
// 实现抽象方法
speak(): void {
console.log("喵喵喵!");
}
}

// 定义 Dog 类继承 Animal
class Dog extends Animal {
// 实现抽象方法
speak(): void {
console.log("汪汪汪!");
}
}

// 定义函数参数类型为抽象类 Animal
// 这个函数可以接受任何 Animal 的子类实例
function makeSpeak(animal: Animal): void {
animal.speak();
}

// 传入不同的子类实例，实现不同的行为（多态）
makeSpeak(new Cat());
makeSpeak(new Dog());

// 抽象类类型数组：可以存储不同子类实例
var animals: Animal[] = [new Cat(), new Dog()];

运行结果：

```
喵喵喵!
汪汪汪!

```

多态：抽象类作为类型时，实际执行的是子类的方法实现。这就是面向对象的多态特性。

### 抽象类与接口的区别

抽象类和接口都用于定义类型规范，但有一些关键区别。

理解它们的差异有助于在实际开发中做出正确的选择。

特性 抽象类 接口 实例化 不能直接实例化 不能直接实例化 方法实现 可以有具体实现 只能有声明（TypeScript 3.6+ 可有默认实现） 成员修饰符 可以添加 public、protected、private 只能有 readonly 继承 单继承（只能 extends 一个类） 多实现（可以 implements 多个接口） 构造函数 可以有构造函数 不能有构造函数

选择建议：需要共享代码使用（用抽象类）；需要定义规范/契约（用接口）；需要多继承（用接口）。

### 完整示例：支付系统

下面是一个使用抽象类实现的支付系统示例。

这个示例展示了抽象类在实际项目中的应用。

### 实例

// 定义抽象支付类
abstract class Payment {
// 抽象方法：处理支付，子类必须实现
abstract process(amount: number): boolean;

// 具象方法：验证支付信息
// 所有子类都可以使用这个方法
validate(): void {
console.log("验证支付信息");
}
}

// 信用卡支付类
class CreditCardPayment extends Payment {
// 信用卡号码
cardNumber: string;

// 构造函数
constructor(cardNumber: string) {
super(); // 调用父类构造函数
this.cardNumber = cardNumber;
}

// 实现抽象方法：处理信用卡支付
process(amount: number): boolean {
console.log("处理信用卡支付: " + amount);
return true;
}
}

// PayPal 支付类
class PayPalPayment extends Payment {
// PayPal 邮箱
email: string;

// 构造函数
constructor(email: string) {
super();
this.email = email;
}

// 实现抽象方法：处理 PayPal 支付
process(amount: number): boolean {
console.log("处理 PayPal 支付: " + amount);
return true;
}
}

// 使用多态处理不同的支付方式
var payments: Payment[] = [
new CreditCardPayment("1234"),
new PayPalPayment("test@example.com")
];

// 遍历处理每种支付方式
for (var _i = 0, payments_1 = payments; _i < payments_1.length; _i++) {
var payment = payments_1[_i];
// 调用继承的验证方法
payment.validate();
// 调用子类的处理方法
payment.process(100);
}

运行结果：

```
验证支付信息
处理信用卡支付: 100
验证支付信息
处理 PayPal 支付: 100

```

应用场景：抽象类常用于框架设计和业务逻辑分层。例如支付处理、用户认证、数据持久化等场景。

### 注意事项

- 不能实例化：抽象类不能直接使用 new 创建实例，必须通过子类继承
- 必须实现抽象方法：子类如果不实现所有抽象方法，会报错
- 可以有构造函数：抽象类可以定义构造函数，子类需要调用 super()
- 单继承：一个类只能继承一个抽象类（单继承）
- 访问修饰符：抽象方法可以使用 public、protected 修饰符

最佳实践：当多个类需要共享代码和逻辑时，使用抽象类。当只需要定义规范和契约时，使用接口。

### 总结

抽象类是 TypeScript 面向对象编程的重要组成部分。

- abstract 关键字：用于声明抽象类和方法
- 不能实例化：只能通过子类继承使用
- 抽象方法：子类必须实现的方法
- 具象方法：子类可以直接继承使用的有实现的方法
- 模板方法模式：抽象类定义骨架，子类提供具体实现

建议：在设计类层次结构时，优先考虑使用抽象类来共享代码和定义规范。

---

## TypeScript 访问修饰符

Source: https://www.runoob.com/typescript/ts-access-modifiers.html

## TypeScript 访问修饰符

访问修饰符（Access Modifiers）是 TypeScript 面向对象编程的核心特性之一。

它们用于控制类成员（属性、方法、构造函数）的可见性。

通过访问修饰符，可以实现封装，保护类的内部实现细节。 访问修饰符作用域 public 类内部 ✓ 子类 ✓ 外部 ✓ protected 类内部 ✓ 子类 ✓ 外部 ✗ private 类内部 ✓ 子类 ✗ 外部 ✗ 使用场景 public 公开方法/属性 如：用户姓名、登录方法 protected 子类可见方法 如：计算公式、模板方法 private 内部实现细节 如：数据库连接、内部计算

### 为什么需要访问修饰符

面向对象编程的三大原则之一是封装。

封装意味着将数据和操作数据的方法隐藏起来，对外只暴露必要的接口。

访问修饰符就是实现封装的手段，它控制类成员的可见范围。

概念说明：访问修饰符决定了类成员可以在哪里被访问。TypeScript 提供三种访问修饰符：public、protected、private。

### public 修饰符

public 是默认的访问修饰符，表示成员可以在任何地方被访问。

无论是类内部、子类还是类的外部，都可以访问 public 成员。

### 实例

// 定义动物类
class Animal {
// 使用 public 修饰 name 属性（可以省略，默认就是 public）
public name: string;

// 构造函数
public constructor(name: string) {
this.name = name;
}

// 公开的说话方法
public speak(): void {
console.log(this.name + " 发出声音");
}
}

// 创建实例
var animal = new Animal("动物");

// 在类外部访问 public 属性
console.log(animal.name);

// 在类外部调用 public 方法
animal.speak();

运行结果：

```
动物
动物 发出声音

```

说明：如果不写访问修饰符，TypeScript 会默认使用 public。因此 public 可以省略，但为了代码清晰，建议明确声明。

### private 修饰符

private 表示私有成员，只能在定义它的类内部访问。

子类和类的外部都不能访问 private 成员。

这常用于隐藏类的内部实现细节，保护数据不被意外修改。

### 实例

// 定义银行账户类
class BankAccount {
// 使用 private 修饰余额，只能在类内部访问
private balance: number;

// 构造函数
constructor(initialBalance: number) {
this.balance = initialBalance;
}

// 存款方法
public deposit(amount: number): void {
if (amount > 0) {
this.balance += amount; // 类内部可以访问 private 属性
console.log("存款成功，当前余额: " + this.balance);
}
}

// 获取余额
public getBalance(): number {
return this.balance; // 类内部可以访问 private 属性
}
}

// 创建账户实例
var account = new BankAccount(1000);

// 存款
account.deposit(500);

// 通过公共方法获取余额
console.log("余额: " + account.getBalance());

// 错误：在类外部不能直接访问 private 属性
// console.log(account.balance); // 编译错误！

运行结果：

```
存款成功，当前余额: 1500
余额: 1500

```

最佳实践：将类的内部状态设为 private，通过 public 方法提供受控的访问途径，这是实现封装的标准做法。

### protected 修饰符

protected 表示受保护成员，可以在类内部和子类中访问。

类的外部不能直接访问 protected 成员。

这在需要让子类继承父类的某些功能，同时隐藏实现细节时非常有用。

### 实例

// 定义人员基类
class Person {
// 使用 protected 修饰 name，子类可以访问
protected name: string;

constructor(name: string) {
this.name = name;
}
}

// 定义员工类，继承 Person
class Employee extends Person {
// 部门是私有的
private department: string;

constructor(name: string, department: string) {
super(name); // 调用父类构造函数
this.department = department;
}

// 自我介绍方法
public introduce(): string {
// 子类可以访问 protected 成员 name
return "我是 " + this.name + "，在 " + this.department + " 工作";
}
}

// 创建员工实例
var emp = new Employee("Alice", "技术部");

console.log(emp.introduce());

// 错误：在类外部不能访问 protected 成员
// console.log(emp.name); // 编译错误！

运行结果：

```
我是 Alice，在 技术部 工作

```

应用场景：protected 常用于定义子类需要使用的属性，但不应该暴露给外部的值。

### readonly 修饰符

readonly 用于将属性设置为只读。

只能在声明时或构造函数中赋值，之后不能修改。

这在定义常量或标识符时非常有用。

### 实例

// 定义用户类
class User {
// 使用 readonly 修饰的属性只能在初始化时赋值
// 用户 ID
readonly id: number;
// 用户名
readonly name: string;

constructor(id: number, name: string) {
this.id = id;
this.name = name;
}
}

// 创建用户实例
var user = new User(1, "Alice");

console.log("用户: " + user.id + ", " + user.name);

// 错误：不能修改 readonly 属性
// user.id = 2; // 编译错误！
// user.name = "Bob"; // 编译错误！

运行结果：

```
用户: 1, Alice

```

注意：readonly 和 private 不冲突。可以同时使用 readonly 和 private，既不能修改也不能从外部访问。

### 参数属性

TypeScript 提供了参数属性（Parameter Properties）的简写语法。

可以在构造函数的参数上直接使用访问修饰符，自动创建并初始化属性。

### 实例

// 定义点类
class Point {
// 在构造函数参数上直接使用修饰符
// 相当于同时声明属性并赋值
constructor(
// public 修饰：创建公开属性 x
public x: number,
// public 修饰：创建公开属性 y
public y: number,
// private 修饰：创建私有属性 z
private z: number
) {
// 构造函数体可以为空，属性已自动创建
}

// 计算三维度总和
public sum(): number {
// 可以在类内部访问所有属性
return this.x + this.y + this.z;
}
}

// 创建点实例
var point = new Point(1, 2, 3);

// 公开属性可以从外部访问
console.log("x: " + point.x);
console.log("y: " + point.y);

// 调用方法
console.log("总和: " + point.sum());

// 错误：私有属性不能从外部访问
// console.log(point.z); // 编译错误！

运行结果：

```
x: 1
y: 2
总和: 6

```

简化代码：参数属性可以大幅简化类的定义，避免重复的声明和赋值代码。

### 访问修饰符对比

下面是三种访问修饰符的作用范围对比。

修饰符 类内部 子类 类外部（实例） public ✓ 可以访问 ✓ 可以访问 ✓ 可以访问 protected ✓ 可以访问 ✓ 可以访问 ✗ 不能访问 private ✓ 可以访问 ✗ 不能访问 ✗ 不能访问

选择建议：默认情况下使用 private，需要子类访问时使用 protected，需要完全公开时使用 public。

### 注意事项

- 默认修饰符：不写修饰符时默认为 public
- 构造函数：构造函数也可以使用访问修饰符，控制实例化权限
- readonly 组合：readonly 可以与 public、protected、private 组合使用
- 编译影响：访问修饰符仅在编译时检查，运行时无效

最佳实践：尽可能使用最严格的访问修饰符。只暴露必要的公开接口，将内部实现设为 private 或 protected。

### 总结

访问修饰符是 TypeScript 实现封装的核心工具。

- public：公开访问，是默认值
- private：私有，仅类内部可见
- protected：保护，类内部和子类可见
- readonly：只读，初始化后不可修改
- 参数属性：简化属性声明的语法

建议：养成使用访问修饰符的习惯，这是编写高质量 TypeScript 代码的基础。

---

## TypeScript 混入（Mixin）

Source: https://www.runoob.com/typescript/ts-mixin.html

## TypeScript 混入（Mixin）

混入（Mixin）是一种代码复用模式，用于将多个独立的功能模块混入到一个类中。

TypeScript 通过泛型函数返回扩展类的方式实现 Mixin，弥补了单继承无法复用多个来源行为的不足。 Mixin 组合过程 User 类 SerializableMixin + serialize() TimestampedMixin + timestamp 扩展后的类 Mixin 工作原理 1. Mixin 是函数 接收一个类，返回扩展后的新类 2. 继承而非修改 通过 extends 继承，不修改原类 3. 可叠加组合 多个 Mixin 可以嵌套组合使用

### 基本概念

Mixin 的核心是一个接收基类、返回扩展类的泛型函数。`Constructor` 类型约束用于描述"任意可被继承的类"。

### 实例

// Constructor 类型：描述任意可实例化的类
type Constructor<T = {}> = new (...args: any[]) => T;

// Mixin 函数：接收基类，返回扩展后的新类
function Timestamped<TBase extends Constructor>(Base: TBase) {
return class extends Base {
createdAt = new Date();
};
}

// 基类
class User {
constructor(public name: string) {}
}

// 混入后得到新类
const TimestampedUser = Timestamped(User);
const user = new TimestampedUser("Alice");

console.log(user.name); // Alice
console.log(user.createdAt instanceof Date); // true

运行结果：

```
Alice
true

```

注意：Mixin 不修改原始类，而是返回一个全新的扩展类，原始的 `User` 类不受影响。

### 组合多个 Mixin

将多个 Mixin 函数嵌套调用，即可将多份能力依次叠加到同一个类上。

### 实例

type Constructor<T = {}> = new (...args: any[]) => T;

// Mixin 1：添加时间戳
function Timestamped<TBase extends Constructor>(Base: TBase) {
return class extends Base {
createdAt = new Date();
};
}

// Mixin 2：添加序列化
function Serializable<TBase extends Constructor>(Base: TBase) {
return class extends Base {
serialize(): string {
return JSON.stringify(this);
}
};
}

// Mixin 3：添加日志（依赖 serialize 方法）
function Loggable<TBase extends Constructor<{ serialize(): string }>>(Base: TBase) {
return class extends Base {
log(): void {
console.log("[LOG]", this.serialize());
}
};
}

class Product {
constructor(public name: string, public price: number) {}
}

// 依次叠加三个 Mixin
const AdvancedProduct = Loggable(Serializable(Timestamped(Product)));

const p = new AdvancedProduct("Phone", 999);
p.log();
console.log(p.createdAt instanceof Date); // true

运行结果：

```
[LOG] {"name":"Phone","price":999,"createdAt":"2026-..."}
true

```

### Mixin 与接口结合

Mixin 返回的类可以声明实现某个接口，使消费方通过接口类型操作混入后的对象，而不依赖具体实现。

### 实例

type Constructor<T = {}> = new (...args: any[]) => T;

// 定义能力接口
interface ISerializable {
serialize(): string;
}

interface ICloneable<T> {
clone(): T;
}

// Mixin 实现接口
function Serializable<TBase extends Constructor>(Base: TBase) {
return class extends Base implements ISerializable {
serialize(): string {
return JSON.stringify(this);
}
};
}

function Cloneable<TBase extends Constructor>(Base: TBase) {
return class extends Base {
clone() {
return Object.assign(
Object.create(Object.getPrototypeOf(this)),
this
);
}
};
}

class Article {
constructor(public title: string, public content: string) {}
}

const RichArticle = Cloneable(Serializable(Article));

const a1 = new RichArticle("TypeScript 入门", "正文内容...");
const a2 = a1.clone();
a2.title = "TypeScript 进阶";

console.log(a1.serialize());
// {"title":"TypeScript 入门","content":"正文内容..."}

console.log(a2.serialize());
// {"title":"TypeScript 进阶","content":"正文内容..."}

console.log(a1.title === a2.title); // false（克隆后独立修改）

运行结果：

```
{"title":"TypeScript 入门","content":"正文内容..."}
{"title":"TypeScript 进阶","content":"正文内容..."}
false

```

### 带约束的 Mixin

通过泛型约束，可以限定 Mixin 只能混入满足特定条件的类，避免运行时缺少必要属性。

### 实例

type Constructor<T = {}> = new (...args: any[]) => T;

// 约束：基类必须有 id 和 name 属性
type WithIdAndName = Constructor<{ id: number; name: string }>;

function Printable<TBase extends WithIdAndName>(Base: TBase) {
return class extends Base {
print(): void {
console.log(`[${this.id}] ${this.name}`);
}
};
}

class Item {
constructor(public id: number, public name: string) {}
}

// 正确：Item 满足约束
const PrintableItem = Printable(Item);
const item = new PrintableItem(42, "Keyboard");
item.print(); // [42] Keyboard

// 错误示例（编译器会阻止）：
// class NoId { constructor(public name: string) {} }
// const Bad = Printable(NoId); // 错误：NoId 缺少 id 属性

运行结果：

```
[42] Keyboard

```

### Mixin 与继承的对比

维度 继承（extends） Mixin 来源数量只能继承一个父类可叠加任意数量 耦合程度子类与父类强耦合每个 Mixin 独立，低耦合 复用粒度复用整个类的能力按需复用单一功能 类型安全原生支持需借助泛型约束保证 适用场景强 "is-a" 关系横切关注点（日志、序列化、缓存等）

### 总结

- 核心模式：Mixin 是接收基类、返回扩展类的泛型函数，`Constructor<T>` 是标准约束类型
- 能力叠加：嵌套调用多个 Mixin 函数即可将多份能力组合到同一个类
- 接口结合：Mixin 返回的类可以实现接口，消费方只需依赖接口而非具体类
- 泛型约束：通过约束 `TBase` 可以限定 Mixin 的适用范围，在编译期阻止错误使用
- 适用场景：日志、序列化、克隆、时间戳等横切关注点，优于多层继承

---

## TypeScript 装饰器

Source: https://www.runoob.com/typescript/ts-decorators.html

## TypeScript 装饰器

装饰器（Decorators）是 TypeScript 的一项实验性功能。

它允许开发者在不修改原类的情况下，为类、方法、属性或参数添加额外的功能。

装饰器本质上是一个函数，它可以在运行时被调用，以修改目标对象的行为。 装饰器类型与应用位置 @ClassDecorator 类装饰器 属性装饰器 @propertyDecorator 方法装饰器 @methodDecorator 参数装饰器 @paramDecorator 访问器装饰器 @getterDecorator 装饰器工厂（Decorator Factory） 装饰器工厂是返回装饰器函数的函数，通过参数实现配置化 例如：function color(code: string) { return function(target) { ... }; } → @color("34")

### 装饰器简介

装饰器采用 `@` 符号作为语法糖，可以附加在类、方法、访问器、属性或参数上。

这种模式常用于框架开发，如 Angular、TypeORM 等都大量使用装饰器来实现依赖注入、数据验证等功能。

注意：装饰器目前是实验性功能，需要在 tsconfig.json 中显式启用。在生产环境中使用时请确认项目对实验性特性的支持程度。

### 配置启用装饰器

在使用装饰器之前，需要在 TypeScript 配置文件 tsconfig.json 中启用相关编译选项。

### tsconfig.json 配置

{
"compilerOptions": {
// 启用装饰器语法
"experimentalDecorators": true,

// 启用装饰器元数据（用于依赖注入框架）
"emitDecoratorMetadata": true
}
}

参数说明：

- experimentalDecorators：启用装饰器语法支持，这是使用装饰器的前提条件
- emitDecoratorMetadata：在编译后的 JavaScript 中生成装饰器的元数据，供依赖注入框架使用

### 类装饰器

类装饰器应用于类的构造函数，可以修改类的定义或添加额外的处理逻辑。

类装饰器接收一个参数，即目标类的构造函数。

### 类装饰器基本用法

// 定义一个类装饰器函数
// 参数 target 就是被装饰的类构造函数
function sealed(target: Function) {
// 打印装饰器被应用到的类名
console.log("装饰器 applied to: " + target.name);

// 使用 Object.seal 锁定构造函数和原型
// 防止在运行时添加或删除属性
Object.seal(target);
Object.seal(target.prototype);
}

// 使用 @ 语法将装饰器应用到类上
@sealed
class Person {
name: string;

constructor(name: string) {
this.name = name;
}
}

// 创建实例测试
var person = new Person("RUNOOB");
console.log("创建: " + person.name);

// 尝试添加新属性（会被阻止，因为类被 seal 了）
// person.age = 25; // 静默失败

运行结果：

```
装饰器 applied to: Person
创建: RUNOOB

```

说明：

类装饰器在类定义时就会执行，通常用于修改类行为、添加元数据或实现 AOP（面向切面编程）。

### 方法装饰器

方法装饰器应用于类的方法，可以修改方法的属性描述符（Property Descriptor）。

方法装饰器接收三个参数：目标对象、属性名称和属性描述符。

### 方法装饰器基本用法

// 定义方法装饰器工厂
// 返回一个装饰器函数
function enumerable(value: boolean) {
// 返回装饰器函数，接收三个参数
return function (
target: any, // 所属类的原型对象
propertyKey: string, // 方法名称
descriptor: PropertyDescriptor // 属性描述符
) {
// 修改属性的 enumerable 特性
// false 表示该方法不可遍历
descriptor.enumerable = value;
};
}

class Greeter {
greeting: string;

constructor(message: string) {
this.greeting = message;
}

// 应用装饰器，设置该方法为不可枚举
@enumerable(false)
greet() {
return "Hello, " + this.greeting;
}
}

var g = new Greeter("World");

// 检查 greet 方法是否可枚举
console.log("方法可枚举: " + g.propertyIsEnumerable("greet"));

// 遍历对象的属性
for (var key in g) {
console.log("属性: " + key);
}

运行结果：

```
方法可枚举: false

```

提示：PropertyDescriptor 包含可枚举（enumerable）、可配置（configurable）、可写（writable）和值（value）等属性，可以根据需要修改。

### 访问器装饰器

访问器装饰器应用于类的 getter 和 setter 方法。

与方法装饰器类似，访问器装饰器也可以修改属性描述符。

### 访问器装饰器用法

// 访问器装饰器工厂
function configurable(value: boolean) {
return function (
target: any,
propertyKey: string,
descriptor: PropertyDescriptor
) {
// 修改属性的 configurable 特性
// false 表示该访问器不可被重新配置或删除
descriptor.configurable = value;
};
}

class Point {
private _x: number = 0;
private _y: number = 0;

// 使用装饰器锁定 getter
@configurable(false)
get x() {
return this._x;
}

@configurable(false)
get y() {
return this._y;
}

set x(value: number) {
this._x = value;
}

set y(value: number) {
this._y = value;
}
}

var point = new Point();
point.x = 10;
point.y = 20;
console.log("坐标: (" + point.x + ", " + point.y + ")");

说明：

访问器装饰器不能同时应用于同一个属性的 getter 和 setter，只能选择其中一个。

### 属性装饰器

属性装饰器应用于类的属性定义。

属性装饰器接收两个参数：目标对象和属性名称。

### 属性装饰器用法

// 属性装饰器工厂
function format(formatString: string) {
return function (
target: any, // 类的原型对象
propertyKey: string // 属性名称
) {
// 在目标对象上存储元数据
// 使用 propertyKey + "_format" 作为键名避免冲突
Object.defineProperty(target, propertyKey + "_format", {
value: formatString,
writable: false,
enumerable: false,
configurable: true
});
};
}

class User {
// 应用属性装饰器，指定日期格式
@format("YYYY-MM-DD")
birthDate: string;

constructor(birthDate: string) {
this.birthDate = birthDate;
}
}

var user = new User("1990-01-01");
console.log("出生日期: " + user.birthDate);

// 访问存储的元数据
console.log("日期格式: " + (user as any).birthDate_format);

运行结果：

```
出生日期: 1990-01-01
日期格式: YYYY-MM-DD

```

### 参数装饰器

参数装饰器应用于类方法的参数，可以为参数添加元数据或标记。

参数装饰器接收三个参数：目标对象、方法名称和参数在函数中的索引。

### 参数装饰器用法

// 参数装饰器
// 用于记录参数信息或进行验证
function logParameter(
target: any, // 类的原型对象
propertyKey: string, // 方法名称
parameterIndex: number // 参数在函数中的索引位置（从 0 开始）
) {
console.log("参数装饰器: " + propertyKey +
" 第 " + (parameterIndex + 1) + " 个参数");
}

class Greeter {
greeting: string;

constructor(greeting: string) {
this.greeting = greeting;
}

// 在参数前使用 @ 语法应用装饰器
greet(@logParameter name: string) {
return this.greeting + ", " + name;
}
}

var greeter = new Greeter("Hello");
greeter.greet("RUNOOB");

运行结果：

```
参数装饰器: greet 第 1 个参数

```

### 装饰器工厂

装饰器工厂是返回装饰器函数的函数。

通过装饰器工厂，可以在应用装饰器时传入自定义参数，实现更灵活的配置。

### 装饰器工厂实现带颜色的日志

// 装饰器工厂：接收配置参数，返回装饰器函数
function color(colorCode: string) {
// colorCode 是 ANSI 转义序列的颜色代码
// 例如：34 = 蓝色，31 = 红色，32 = 绿色
return function (
target: any,
propertyKey: string,
descriptor: PropertyDescriptor
) {
// 保存原始方法
var originalMethod = descriptor.value;

// 重写方法，添加颜色
descriptor.value = function (...args: any[]) {
// 调用原始方法获取返回值
var result = originalMethod.apply(this, args);

// 如果在终端环境，给输出添加颜色
// ANSI 转义序列格式：\x1b[颜色码m 内容 \x1b[0m
return "\x1b[" + colorCode + "m" + result + "\x1b[0m";
};
};
}

class Logger {
// 使用装饰器工厂，传入蓝色代码 34
@color("34")
log(message: string): string {
return message;
}

@color("31")
error(message: string): string {
return message;
}

@color("32")
success(message: string): string {
return message;
}
}

var logger = new Logger();
console.log(logger.log("这是蓝色日志"));
console.log(logger.error("这是红色错误"));
console.log(logger.success("这是绿色成功"));

运行结果：

```
这是蓝色日志（终端显示为蓝色）
这是红色错误（终端显示为红色）
这是绿色成功（终端显示为绿色）

```

说明：装饰器工厂是实际开发中最常用的形式，它允许在应用装饰器时传递参数，实现配置化。

### 装饰器执行顺序

当一个类上有多个装饰器时，执行顺序遵循特定的规则。

- 装饰器从下往上应用
- 同一类型的多个装饰器从右到左执行
- 参数装饰器先于方法装饰器执行

### 装饰器执行顺序示例

// 多个装饰器叠加使用
function first() {
console.log("first 装饰器");
return function (target: any) {
console.log("first 装饰器函数");
};
}

function second() {
console.log("second 装饰器");
return function (target: any) {
console.log("second 装饰器函数");
};
}

@first()
@second()
class MyClass {
name: string;
}

var obj = new MyClass();

运行结果：

```
second 装饰器
first 装饰器
second 装饰器函数
first 装饰器函数

```

说明：

装饰器函数先执行定义（console.log），然后按照从下往上的顺序执行装饰器函数。

### 实际应用场景

装饰器在实际项目中有广泛的应用场景。

#### 日志记录

自动记录方法调用日志。

### 实例

// 日志装饰器
function log(target: any, propertyKey: string, descriptor: PropertyDescriptor) {
var originalMethod = descriptor.value;

descriptor.value = function (...args: any[]) {
console.log("调用方法: " + propertyKey + "，参数: " + JSON.stringify(args));
var result = originalMethod.apply(this, args);
console.log("方法返回: " + JSON.stringify(result));
return result;
};
}

class MathService {
@log
add(a: number, b: number): number {
return a + b;
}

@log
multiply(a: number, b: number): number {
return a * b;
}
}

var math = new MathService();
console.log("计算结果: " + math.add(5, 3));

运行结果：

```
调用方法: add，参数: [5,3]
方法返回: 8
计算结果: 8

```

#### 权限验证

实现方法级别的权限检查。

### 实例

// 模拟当前用户角色
var currentUser = { role: "admin" };

// 权限装饰器
function requireRole(role: string) {
return function (target: any, propertyKey: string, descriptor: PropertyDescriptor) {
var originalMethod = descriptor.value;

descriptor.value = function (...args: any[]) {
if (currentUser.role !== role) {
console.log("权限不足，无法执行 " + propertyKey);
return null;
}
return originalMethod.apply(this, args);
};
};
}

class AdminService {
@requireRole("admin")
deleteUser(id: number): string {
return "删除用户 " + id + " 成功";
}

@requireRole("admin")
viewUser(id: number): string {
return "查看用户 " + id;
}
}

var admin = new AdminService();
console.log(admin.viewUser(1));
console.log(admin.deleteUser(1));

// 模拟普通用户
currentUser = { role: "user" };
console.log(admin.deleteUser(2));

运行结果：

```
查看用户 1
删除用户 1 成功
权限不足，无法执行 deleteUser

```

### 注意事项

- 实验性功能：装饰器是 ECMAScript 的_stage 3_提案，当前仍是实验性功能
- 编译选项：必须启用 experimentalDecorators
- 类型定义：需要较新版本的 TypeScript 以获得完整的类型支持
- 调试注意：装饰器在编译时执行，某些调试工具可能无法正确映射源码位置

建议：在项目中使用装饰器时，建议创建专门的装饰器工具类或函数库，统一管理装饰器的定义和使用。

### 总结

TypeScript 装饰器提供了一种强大的元编程能力。

- 类装饰器：修改类本身，可用于添加元数据、锁定类等
- 方法装饰器：修改方法属性，可用于日志、验证等
- 访问器装饰器：修改 getter/setter，控制属性的可配置性
- 属性装饰器：为属性添加元数据
- 参数装饰器：标记或验证方法参数
- 装饰器工厂：通过参数化实现更灵活的装饰器配置

---

## TypeScript 类型守卫

Source: https://www.runoob.com/typescript/ts-type-guards.html

## TypeScript 类型守卫

类型守卫（Type Guards）是 TypeScript 中非常重要的一种类型缩小机制。

它允许开发者在运行时通过特定的条件检查，让 TypeScript 编译器能够准确推断出变量的具体类型。

通过类型守卫，我们可以安全地访问联合类型变量中特定类型的属性和方法。

### 为什么需要类型守卫

在 TypeScript 中，一个变量可能被声明为多种类型的联合。

当我们需要根据不同类型执行不同操作时，编译器无法自动判断当前的具体类型。

类型守卫就是解决这个问题的关键机制。

概念说明：类型守卫的核心原理是"类型缩窄"（Type Narrowing）。通过条件判断，TypeScript 会自动将联合类型缩小为具体的类型。 类型守卫 - 类型缩窄流程 原始联合类型 string | number | boolean typeof typeof === "string" 缩窄为 string 分支 类型已确定 可访问 .length 类型守卫方式 typeof typeof x === "string" 原始类型 instanceof x instanceof Array 类实例 自定义守卫 x is String value is Type in "prop" in x 属性检查

### typeof 类型守卫

typeof 是最常用的类型守卫，用于检查原始类型（string、number、boolean 等）。

它返回一个字符串，表示值的类型。

### typeof 基本用法

// 定义一个接收联合类型的函数
// 参数 value 可能是字符串或数字
function printValue(value: string | number): void {
// 使用 typeof 检查类型
// 在 if 条件为 true 时，TypeScript 会自动将 value 缩小为 string 类型
if (typeof value === "string") {
// 此时 TypeScript 知道 value 是 string
// 可以安全地访问字符串的 length 属性
console.log("字符串长度: " + value.length);
} else {
// 进入 else 分支时，TypeScript 知道 value 不是 string
// 只能是 number 类型
// 可以安全地进行数学运算
console.log("数字翻倍: " + (value * 2));
}
}

// 测试调用
printValue("hello"); // 传入字符串
printValue(42); // 传入数字

运行结果：

```
字符串长度: 5
数字翻倍: 84

```

typeof 支持的类型：

- "string" - 字符串类型
- "number" - 数字类型（包括 NaN 和 Infinity）
- "boolean" - 布尔类型
- "undefined" - 未定义类型
- "object" - 对象类型（注意：数组和 null 也会被识别为 "object"）
- "function" - 函数类型

注意：typeof 对于数组和 null 都会返回 "object"。如果需要精确区分数组和对象，需要使用其他方式。

### instanceof 类型守卫

instanceof 用于检查对象是否是某个类的实例。

它通过检查对象的原型链来判断类型。

### instanceof 基本用法

// 定义 Dog 类
class Dog {
// 狗的叫声方法
bark(): void {
console.log("汪汪汪!");
}
}

// 定义 Cat 类
class Cat {
// 猫的叫声方法
meow(): void {
console.log("喵喵喵!");
}
}

// 接收联合类型的函数
function makeSound(animal: Dog | Cat): void {
// 使用 instanceof 检查 animal 是 Dog 还是 Cat
// 在 if 条件为 true 时，TypeScript 将 animal 缩小为 Dog 类型
if (animal instanceof Dog) {
// 此时可以调用 Dog 特有的方法
animal.bark();
} else {
// else 分支中，TypeScript 将 animal 缩小为 Cat 类型
animal.meow();
}
}

// 测试调用
makeSound(new Dog()); // 创建 Dog 实例并调用
makeSound(new Cat()); // 创建 Cat 实例并调用

运行结果：

```
汪汪汪!
喵喵喵!

```

说明：

instanceof 检查的是对象的原型链，因此它只能用于类实例，不能用于接口和类型别名。

### 自定义类型守卫

当内置的 typeof 和 instanceof 无法满足需求时，可以创建自定义类型守卫函数。

自定义类型守卫使用 `value is Type` 的返回类型语法。

### 自定义守卫函数

// 定义一个自定义类型守卫函数
// 返回类型使用 "value is Type" 格式
// 这告诉 TypeScript：当函数返回 true 时，参数类型就是 string
function isString(value: any): value is string {
// 使用 typeof 检查是否为字符串
return typeof value === "string";
}

// 另一个自定义守卫：检查是否为数字
function isNumber(value: any): value is number {
return typeof value === "number";
}

// 定义一个数组类型守卫
function isArray(value: any): value is any[] {
return Array.isArray(value);
}

// 处理值的函数
function processValue(value: string | number | any[]): void {
// 使用自定义守卫进行类型检查
if (isString(value)) {
// TypeScript 知道 value 是 string 类型
// 可以调用 toUpperCase() 方法
console.log("字符串转大写: " + value.toUpperCase());
} else if (isNumber(value)) {
// TypeScript 知道 value 是 number 类型
// 可以调用 toFixed() 方法
console.log("数字格式化: " + value.toFixed(2));
} else if (isArray(value)) {
// TypeScript 知道 value 是数组类型
console.log("数组长度: " + value.length);
}
}

// 测试调用
processValue("hello");
processValue(3.14159);
processValue([1, 2, 3, 4, 5]);

运行结果：

```
字符串转大写: HELLO
数字格式化: 3.14
数组长度: 5

```

提示：自定义类型守卫的关键是返回类型 `value is Type`，这是 TypeScript 识别类型守卫的标志。

### in 操作符类型守卫

in 操作符可以检查对象是否包含某个属性。

在条件判断中使用 in，TypeScript 会自动缩小对象的类型范围。

### in 操作符用法

// 定义两个接口，它们有不同的属性
interface A {
a: string; // 只有属性 a
}

interface B {
b: number; // 只有属性 b
}

// 接收联合类型的函数
function process(obj: A | B): void {
// 使用 in 检查对象是否包含属性 "a"
if ("a" in obj) {
// 在 if 分支中，TypeScript 知道 obj 包含属性 a
// 因此 obj 的类型被缩小为 A
console.log("A 的属性 a: " + obj.a);
} else {
// else 分支中，obj 不包含属性 a
// TypeScript 知道 obj 只能是 B 类型
// 因此可以安全访问属性 b
console.log("B 的属性 b: " + obj.b);
}
}

// 测试调用
process({ a: "hello" }); // 传入包含属性 a 的对象
process({ b: 42 }); // 传入包含属性 b 的对象

运行结果：

```
A 的属性 a: hello
B 的属性 b: 42

```

### 可辨识联合与类型守卫

可辨识联合是一种强大的模式，它通过一个公共的"标识"属性来区分联合类型成员。

结合 switch 语句或 if 判断，可以实现完整的类型守卫。

### 可辨识联合实现计算器

// 定义圆形接口，使用 kind 属性作为标识
interface Circle {
kind: "circle"; // 标识字段：值为 "circle"
radius: number; // 半径
}

// 定义矩形接口
interface Rectangle {
kind: "rectangle"; // 标识字段：值为 "rectangle"
width: number; // 宽度
height: number; // 高度
}

// 定义三角形接口
interface Triangle {
kind: "triangle"; // 标识字段：值为 "triangle"
base: number; // 底边
height: number; // 高度
}

// 定义联合类型
type Shape = Circle | Rectangle | Triangle;

// 计算面积的函数
function getArea(shape: Shape): number {
// 使用 switch 语句进行类型守卫
// 根据 kind 属性的值，TypeScript 会自动缩小类型
switch (shape.kind) {
case "circle":
// shape 被缩小为 Circle 类型
// 可以访问 radius 属性
return Math.PI * shape.radius ** 2;

case "rectangle":
// shape 被缩小为 Rectangle 类型
// 可以访问 width 和 height 属性
return shape.width * shape.height;

case "triangle":
// shape 被缩小为 Triangle 类型
return 0.5 * shape.base * shape.height;
}
}

// 测试调用
var circle = { kind: "circle" as const, radius: 5 };
var rectangle = { kind: "rectangle" as const, width: 4, height: 6 };
var triangle = { kind: "triangle" as const, base: 3, height: 4 };

console.log("圆形面积: " + getArea(circle).toFixed(2));
console.log("矩形面积: " + getArea(rectangle));
console.log("三角形面积: " + getArea(triangle));

运行结果：

```
圆形面积: 78.54
矩形面积: 24
三角形面积: 6

```

说明：可辨识联合是 TypeScript 中最推荐使用的模式之一。它通过一个公共的字面量属性（通常是 kind 或 type）来区分不同的类型成员，使代码既类型安全又易于维护。

### null 和 undefined 检查

处理可能为 null 或 undefined 的值时，直接的相等性检查也是有效的类型守卫。

### null 检查

// 定义一个可能为 null 的函数参数
function getLength(str: string | null): number {
// 直接检查 str 不等于 null
// 在条件为 true 时，TypeScript 知道 str 不是 null
// 此时可以安全地访问 str 的属性
if (str !== null) {
return str.length;
}

// null 情况的处理
return 0;
}

// 调用测试
console.log(getLength("hello")); // 正常字符串
console.log(getLength(null)); // 传入 null

运行结果：

```
5
0

```

提示：在启用 strictNullChecks 后，建议始终进行 null 检查。可以使用可选链（?.）和空值合并（??）来简化代码。

### 真值缩小

除了显式的类型检查，TypeScript 还会通过真值断言（Truthiness）来缩小类型范围。

### 真值缩小

// 可能为 undefined 的字符串
function greet(name?: string): string {
// 使用短路运算符：如果 name 为 undefined 或空字符串，使用默认值
// 在 && 后的代码块中，TypeScript 知道 name 一定有值
return name && "Hello, " + name;
}

// 测试
console.log(greet("RUNOOB"));
console.log(greet());

运行结果：

```
Hello, RUNOOB
Hello, undefined

```

### 注意事项

- 类型守卫必须在条件分支中使用：只有在使用类型守卫进行条件判断后，TypeScript 才会进行类型缩小
- 返回类型必须是类型谓词：自定义类型守卫的返回类型必须是 `value is Type` 格式
- 可辨识联合是最佳实践：对于复杂的联合类型，建议使用可辨识联合模式
- 注意类型收窄的完整性：使用 switch 语句时，建议处理所有可能的分支

建议：在处理联合类型时，优先考虑使用可辨识联合模式。它不仅代码更清晰，还能充分利用 TypeScript 的类型推断能力。

### 总结

类型守卫是 TypeScript 类型系统的重要组成部分。

- typeof：最常用的方式，适用于原始类型的检查
- instanceof：检查对象是否是特定类的实例
- 自定义守卫：通过 `value is Type` 语法实现灵活的类检查
- in：检查对象是否包含特定属性
- 可辨识联合：推荐的最佳实践模式，通过标识字段区分类型
- 真值缩小：利用 JavaScript 的真值判断进行类型收窄

---

## TypeScript 可选链

Source: https://www.runoob.com/typescript/ts-optional-chaining.html

## TypeScript 可选链

可选链（Optional Chaining）是 TypeScript 和 JavaScript 中一种安全的属性访问方式。

它允许开发者以链式调用的方式安全地访问嵌套对象属性。当访问路径中的任意一个属性为 null 或 undefined 时，整个表达式会短路返回 undefined，而不会抛出错误。

这极大地简化了深层嵌套对象的属性访问代码。 可选链工作原理 传统方式（冗长） // 需要多层检查 var city = user && user.address && user.address.city; 容易出错，代码冗余 简化 可选链方式（简洁） // 一行搞定 var city = user ?.address ?.city; 安全简洁，自动短路 可选链的三种形式 obj?.prop —— 属性访问 arr?.[0] —— 数组访问 obj?.method() —— 方法调用

### 为什么需要可选链

在 JavaScript/TypeScript 开发中，经常需要访问深层嵌套的对象属性。

传统的写法需要逐层检查属性是否存在。这种方式不仅代码冗长，而且容易遗漏检查导致运行时错误。

概念说明：可选链的核心是"短路求值"。当链中某个属性的值为 null 或 undefined 时，整个表达式的结果立即返回 undefined，而不会继续访问后续属性。

### 基本语法

使用 `?.` 运算符安全访问可能不存在的属性。

与传统的 && 链式检查相比，可选链语法更加简洁直观。

### 实例

// 定义一个嵌套的用户对象
// 包含姓名和地址信息，地址中有城市
var user = {
name: "RUNOOB",
address: {
city: "Beijing"
}
};

// 传统方式：使用 && 逐层检查
// 这种方式代码冗长，容易遗漏
var city1 = user && user.address && user.address.city;

// 可选链方式：使用 ?. 运算符
// 如果任意一层为 null 或 undefined，直接返回 undefined
var city2 = user?.address?.city;

console.log("传统方式: " + city1);
console.log("可选链: " + city2);

运行结果：

```
传统方式: Beijing
可选链: Beijing

```

提示：当对象属性存在时，两种方式的结果相同。但可选链的代码更简洁，更易读。

### 处理不存在的属性

当访问路径中的属性不存在时，可选链会安全地返回 undefined，而不会抛出错误。

这对于处理来自 API 的数据或用户表单输入特别有用。

### 实例

// 定义一个不完整用户对象
// 只有 name 属性，没有 address 属性
var user = {
name: "RUNOOB"
// address 属性不存在
};

// 使用可选链访问深层属性
// user.address 为 undefined，所以 city 也是 undefined
var city = user?.address?.city;

// 访问更深的嵌套属性
// 即使 country 也不存在，仍然返回 undefined 而不报错
var country = user?.address?.country?.name;

console.log("城市: " + city);
console.log("国家: " + country);

运行结果：

```
城市: undefined
国家: undefined

```

注意：可选链只会在属性访问时返回 undefined，不会创建新的对象或属性。

### 可选链与数组结合

可选链可以与数组下标访问结合使用，安全地访问数组中的元素。

使用 `?.[index]` 语法，可以在数组元素不存在时返回 undefined。

### 实例

// 定义一个用户数组
var users = [
{ name: "Alice" },
{ name: "Bob" }
];

// 安全访问数组第一个元素的名字
// users?.[0] 存在，返回 "Alice"
var firstUser = users?.[0]?.name;

// 安全访问数组中不存在的元素
// users?.[9] 不存在（数组只有2个元素），返回 undefined
var tenthUser = users?.[9]?.name;

console.log("第一个用户: " + firstUser);
console.log("第十个用户: " + tenthUser);

运行结果：

```
第一个用户: Alice
第十个用户: undefined

```

说明：`?.[index]` 与 `?.` 的区别在于：前者用于数组，后者用于对象属性。

### 可选链与方法调用

可选链可以用于安全地调用可能不存在的方法。

使用 `?.()` 语法，如果方法不存在则返回 undefined，而不会抛出错误。

### 实例

// 定义一个用户对象，包含姓名和 greet 方法
var user = {
name: "Alice",
// 定义一个打招呼方法
greet: function() {
return "Hello, " + this.name;
}
};

// 安全调用存在的方法
// user.greet 存在，正常调用并返回结果
var message1 = user.greet?.();

// 安全调用不存在的方法
// user.sayHello 不存在，返回 undefined 而不报错
var message2 = user.sayHello?.();

console.log("greet: " + message1);
console.log("sayHello: " + message2);

运行结果：

```
greet: Hello, Alice
sayHello: undefined

```

应用场景：这在处理可选的回调函数或事件处理程序时特别有用。

### 可选链赋值

可选链也可以用于赋值操作，但需要注意其行为。

可选链赋值只会修改已存在的路径，不会自动创建中间对象。

### 实例

// 定义一个简单的用户对象
var user = {
name: "Alice"
};

// 尝试使用可选链赋值
// user?.address 不存在，所以整个赋值被跳过
// user 对象保持不变
user?.address?.city = "Beijing";

console.log("用户: " + JSON.stringify(user));

运行结果：

```
用户: {"name":"Alice"}
```

注意：可选链赋值不能用于创建新属性。如果需要创建嵌套对象，应该使用传统方式先创建中间对象。

### 空值合并与可选链

可选链经常与空值合并运算符 `??` 结合使用。

这种组合可以在属性不存在时提供默认值，使代码更加健壮。

### 实例

// 定义一个不完整的用户对象
var user = {
name: "Alice"
// address 不存在
};

// 可选链 + 空值合并：当 city 为 null 或 undefined 时使用默认值
var city = user?.address?.city ?? "未知城市";

console.log("城市: " + city);

// 对比：传统方式的复杂写法
var country = user && user.address && user.address.country
? user.address.country
: "未知国家";

console.log("国家: " + country);

运行结果：

```
城市: 未知城市
国家: 未知国家

```

说明：`??` 只在值为 null 或 undefined 时使用默认值，而 `||` 会在值为 falsy（0、""、false）时也使用默认值。在处理数字或字符串时，应优先使用 `??`。

### 注意事项

- 短路求值：可选链遇到 null 或 undefined 会立即返回，不会继续访问后续属性
- 不能创建属性：可选链赋值不会自动创建中间对象
- 与空值合并配合：建议始终使用 `??` 提供默认值，而少用 `||`
- 性能考虑：虽然可选链更安全，但在属性一定存在的情况下，直接访问性能更好

建议：在实际开发中，对于来自外部数据（如 API 响应、用户输入）的属性，优先使用可选链进行安全访问。

### 总结

可选链是现代 JavaScript/TypeScript 中处理嵌套对象的重要特性。

- `?.`：安全访问对象属性
- `?.[index]`：安全访问数组元素
- `?.()`：安全调用可能不存在的方法
- `??`：提供默认值，处理 undefined 情况

最佳实践：对于可能不存在的深层属性访问，始终使用可选链。结合空值合并运算符，可以写出既安全又简洁的代码。

---

## TypeScript 工具类型

Source: https://www.runoob.com/typescript/ts-utility-types.html

## TypeScript 工具类型

工具类型（Utility Types）是 TypeScript 内置的一系列高级类型。

它们可以帮助开发者快速创建和转换类型，提高代码的可复用性和类型安全性。

工具类型本质上是泛型类型，通过映射类型和条件类型实现。 TypeScript 工具类型 原始类型 interface User { id: number; name: string; email: string; } 工具类型 Partial<T> - 所有可选 Required<T> - 所有必填 Readonly<T> - 所有只读 Pick<T,K> - 选择属性 Omit<T,K> - 排除属性 结果类型 type Result = { id?: number; name?: string; email?: string; } 常用工具类型 Partial<T> 属性全部可选 { a?: T } Required<T> 属性全部必填 { -? } Readonly<T> 属性全部只读 { readonly } Pick<T,K> 选择指定属性 { pick } Record<K,T> 构造对象类型 Omit<T,K> 排除指定属性

### 为什么需要工具类型

在实际的 TypeScript 开发中，我们经常需要基于现有类型创建新的类型。

手动创建这些类型不仅繁琐，而且容易出错。

工具类型提供了一种声明式的方式来转换和创建类型，大大提高了开发效率。

概念说明：工具类型是 TypeScript 预定义的一系列泛型类型。它们使用映射类型和条件类型来实现类型的转换和生成。

### Partial<T> - 可选属性

Partial 将类型 T 的所有属性设置为可选。

这在创建部分更新对象或处理表单数据时非常有用。

### 实例

// 定义用户接口，包含必填属性
interface User {
// 用户 ID
id: number;
// 用户名
name: string;
// 用户邮箱
email: string;
}

// Partial：将所有属性变为可选
// 转换后的类型所有属性都是可选的
type PartialUser = Partial<User>;

// 使用 PartialUser 类型
// 可以只提供部分属性，不需要全部提供
var user: PartialUser = { name: "Alice" };

console.log("部分用户: " + JSON.stringify(user));

运行结果：

```
部分用户: {"name":"Alice"}

```

应用场景：Partial 常用于更新对象数据。例如：只更新用户的一部分信息时，不需要传入完整的用户对象。

### Required<T> - 必填属性

Required 与 Partial 相反，将所有可选属性设置为必填。

当需要确保对象包含所有属性时使用。

### 实例

// 定义配置接口，属性都是可选的
interface Config {
// 服务器地址
host?: string;
// 端口号
port?: number;
}

// Required：将所有可选属性变为必填
// 转换后的类型所有属性都是必填的
type RequiredConfig = Required<Config>;

// 使用 RequiredConfig 类型
// 必须提供所有属性
var config: RequiredConfig = { host: "localhost", port: 8080 };

console.log("配置: " + JSON.stringify(config));

运行结果：

```
配置: {"host":"localhost","port":8080}

```

说明：Required 不仅移除可选属性（?），还会移除 readonly 修饰符。

### Readonly<T> - 只读属性

Readonly 将所有属性设置为只读。

创建后不允许修改的 对象时非常有用。

### 实例

// 定义用户接口
interface User {
// 用户名
name: string;
// 用户年龄
age: number;
}

// Readonly：将所有属性变为只读
// 转换后的类型所有属性都不能修改
type ReadonlyUser = Readonly<User>;

// 创建只读用户对象
var user: ReadonlyUser = { name: "Alice", age: 25 };

// 尝试修改只读属性会报错
// user.name = "Bob"; // 错误：只读属性不能修改

console.log("只读用户: " + JSON.stringify(user));

应用场景：Readonly 常用于定义配置对象、枚举值映射等不希望被修改的数据。

### Pick<T, K> - 选择属性

Pick 从类型 T 中选择指定的属性 K 组成新类型。

当你只需要某个类型的部分属性时使用。

### 实例

// 定义完整的用户接口
interface User {
// 用户 ID
id: number;
// 用户名
name: string;
// 用户邮箱
email: string;
// 用户密码
password: string;
}

// Pick：选择指定的属性组成新类型
// 从 User 中选择 id 和 name 属性
type UserBasicInfo = Pick<User, "id" | "name">;

// 使用选择后的类型
var user: UserBasicInfo = { id: 1, name: "Alice" };

console.log("用户基本信息: " + JSON.stringify(user));

运行结果：

```
用户基本信息: {"id":1,"name":"Alice"}

```

对比：Pick 和 Partial 结合使用，可以创建只包含部分可选属性的类型。

### Omit<T, K> - 排除属性

Omit 从类型 T 中排除指定的属性 K，返回剩余属性组成的新类型。

与 Pick 相反，用于删除不需要的属性。

### 实例

// 定义完整的用户接口
interface User {
// 用户 ID
id: number;
// 用户名
name: string;
// 用户邮箱
email: string;
// 用户密码
password: string;
}

// Omit：排除指定的属性
// 从 User 中排除 password 属性
type UserWithoutPassword = Omit<User, "password">;

// 使用排除后的类型
var user: UserWithoutPassword = { id: 1, name: "Alice", email: "a@b.com" };

console.log("无密码用户: " + JSON.stringify(user));

运行结果：

```
无密码用户: {"id":1,"name":"Alice","email":"a@b.com"}

```

提示：Omit 是 Pick 的反向操作。当你需要排除少数属性时用 Omit，需要选择少数属性时用 Pick。

### Record<K, T> - 构造对象类型

Record 构造一个对象类型，键的类型为 K，值的类型为 T。

常用于创建键值对映射、字典类型等。

### 实例

// 定义角色类型
type Role = "admin" | "user" | "guest";

// Record：构造对象类型
// 键为 Role 类型，值为字符串数组类型
type RolePermissions = Record<Role, string[]>;

// 使用 Record 创建权限映射
var permissions: RolePermissions = {
// 管理员拥有所有权限
admin: ["read", "write", "delete"],
// 普通用户拥有读写权限
user: ["read", "write"],
// 访客只有读权限
guest: ["read"]
};

console.log("管理员权限: " + permissions.admin);
console.log("访客权限: " + permissions.guest);

运行结果：

```
管理员权限: read,write,delete
访客权限: read

```

应用场景：Record 常用于创建配置映射、状态映射、权限映射等需要根据键快速查找值的场景。

### Exclude<T, U> - 排除类型

Exclude 从类型 T 中排除可以赋值给类型 U 的类型。

主要用于联合类型，排除特定的类型成员。

### 实例

// 定义联合类型，包含 a、b、c、d
type T = "a" | "b" | "c" | "d";

// Exclude：从 T 中排除指定类型
// 排除 "a"、"b"、"c"，只保留 "d"
type NonABC = Exclude<T, "a" | "b" | "c">;

// 使用排除后的类型，只能赋值 "d"
var value: NonABC = "d";

console.log("值: " + value);

运行结果：

```
值: d

```

说明：Exclude 的原理是：如果 T 中的某个类型可以赋值给 U，就从 T 中移除它。

### Extract<T, U> - 提取类型

Extract 与 Exclude 相反，从类型 T 中提取可以赋值给类型 U 的类型。

用于筛选联合类型中的特定成员。

### 实例

// 定义混合联合类型，包含字符串和数字
type T = "a" | "b" | "c" | 1 | 2 | 3;

// Extract：从 T 中提取指定类型
// 提取所有字符串类型："a"、"b"、"c"
type Letters = Extract<T, string>;

// 使用提取后的类型
var letter: Letters = "a";

console.log("字母: " + letter);

运行结果：

```
字母: a

```

对比：Extract 和 Exclude 是互补的。Extract(T, U) 等价于 Exclude(T, Exclude<T, U>)。

### NonNullable<T> - 排除空值

NonNullable 从类型 T 中排除 null 和 undefined。

确保类型不包含空值时使用。

### 实例

// 定义混合类型，包含字符串、null、undefined、数字
type T = string | null | undefined | number;

// NonNullable：排除 null 和 undefined
// 转换后只保留 string 和 number
type NotNull = NonNullable<T>;

// 使用非空类型
var value: NotNull = "hello";
value = 42;

// 尝试赋值 null 会报错
// value = null; // 错误：不能赋值 null

console.log("值: " + value);

说明：NonNullable<T> 等价于 Exclude<T, null | undefined>。

### ReturnType<T> - 获取返回类型

ReturnType 获取函数类型 T 的返回类型。

常用于从已存在的函数中提取返回类型。

### 实例

// 定义获取用户的函数
function getUser() {
return { name: "Alice", age: 25 };
}

// 定义获取配置的函数
function getConfig() {
return { host: "localhost", port: 8080 };
}

// ReturnType：获取函数的返回类型
// 提取 getUser 函数的返回类型
type UserType = ReturnType<typeof getUser>;
// 提取 getConfig 函数的返回类型
type ConfigType = ReturnType<typeof getConfig>;

// 使用提取的返回类型创建对象
var user: UserType = { name: "Bob", age: 30 };
var config: ConfigType = { host: "example.com", port: 3000 };

console.log("用户: " + JSON.stringify(user));
console.log("配置: " + JSON.stringify(config));

运行结果：

```
用户: {"name":"Bob","age":30}
配置: {"host":"example.com","port":3000}

```

提示：ReturnType<T> 中的 T 必须是函数类型。可以使用 typeof 获取函数的类型。

### 注意事项

- 工具类型都是泛型：使用时应传入具体的类型参数
- 只读和可选：工具类型可以组合使用，如 Partial<Readonly<T>>
- 内置类型：这些工具类型都是 TypeScript 内置的，无需安装
- 自定义工具类型：可以基于映射类型和条件类型创建自定义工具类型

进阶：如果内置工具类型不满足需求，可以参考 TypeScript 源码自行实现自定义工具类型。

### 总结

TypeScript 工具类型是类型系统的重要组成部分。

- Partial<T>：将所有属性变为可选
- Required<T>：将所有属性变为必填
- Readonly<T>：将所有属性变为只读
- Pick<T,K>：选择指定属性
- Omit<T,K>：排除指定属性
- Record<K,T>：构造对象类型
- Exclude/Extract：类型过滤
- NonNullable：排除 null 和 undefined
- ReturnType：获取函数返回类型

最佳实践：善用工具类型可以使代码更加类型安全，减少重复的类型定义，提高代码可维护性。

---

## TypeScript 条件类型

Source: https://www.runoob.com/typescript/ts-conditional-types.html

## TypeScript 条件类型

条件类型（Conditional Types）是 TypeScript 类型系统中最强大的特性之一。

它允许根据条件动态地选择类型，类似于编程语言中的三元表达式。

条件类型使得类型定义更加灵活，是实现高级工具类型的基础。 条件类型工作原理 条件类型语法 T extends U ? X : Y 执行流程 检查 T 能否赋值给 U → 是 → X | 否 → Y 示例解析 示例 1 IsString<string> string extends string? ✓ → true 示例 2 IsString<number> number extends string? ✗ → false 示例 3 ReturnType<fn> fn extends (...args) → R (推导返回类型)

### 为什么需要条件类型

在实际的 TypeScript 开发中，我们经常需要根据不同的输入类型返回不同的类型。

例如，一个函数可能接受字符串或数字参数，我们需要根据参数类型返回不同的结果类型。

条件类型提供了一种在类型级别进行逻辑判断的能力，使得类型定义更加灵活和强大。

概念说明：条件类型的语法是 `T extends U ? X : Y`。如果类型 T 可以赋值给类型 U，则返回类型 X，否则返回类型 Y。

### 基本语法

条件类型使用三元表达式的语法，在类型层面进行条件判断。

这使得我们可以根据输入类型动态地计算返回类型。

### 实例

// 条件类型语法：T extends U ? X : Y
// 如果 T 是字符串类型，返回 true，否则返回 false
type IsString<T> = T extends string ? true : false;

// 使用条件类型
// string extends string 为 true，所以 A 类型是 true
type A = IsString<string>;
// number extends string 为 false，所以 B 类型是 false
type B = IsString<number>;

// 使用这些类型
var a: A = true;
var b: B = false;

console.log("string 是字符串?: " + a);
console.log("number 是字符串?: " + b);

运行结果：

```
string 是字符串?: true
number 是字符串?: false

```

说明：条件类型会在类型检查时自动进行求值，生成具体的类型。

### 实际应用：类型过滤

条件类型最常见的应用之一是过滤类型。

例如，可以创建一个类型来排除 null 和 undefined。

### 实例

// 使用条件类型实现 NonNullable
// 如果 T 是 null 或 undefined，返回 never（空类型），否则返回 T 本身
type NonNullable<T> = T extends null | undefined ? never : T;

// 使用 NonNullable 类型
// string 不是 null/undefined，所以类型是 string
type A = NonNullable<string>;
// null 是 null/undefined，所以类型是 never
type B = NonNullable<null>;
// undefined 是 null/undefined，所以类型是 never
type C = NonNullable<undefined>;

// 验证类型
var a: A = "hello";
console.log("非空: " + a);

运行结果：

```
非空: hello

```

never 类型：never 表示永不存在的类型。当条件不满足时，TypeScript 使用 never 来表示"不可用"的类型。

### 类型推导：infer 关键字

infer 关键字是条件类型中最强大的特性之一。

它允许从类型中"提取"或"推导"出特定的类型部分。

### 实例

// 使用 infer 推导函数的返回类型
// 如果 T 是函数类型，返回推断出的返回类型 R，否则返回 never
type ReturnType<T> = T extends (...args: any[]) => infer R ? R : never;

// 定义一个返回用户对象的函数
function getUser() {
return { name: "Alice" };
}

// 定义一个返回数字的函数
function getNumber() {
return 42;
}

// 使用 ReturnType 获取函数返回类型
// 推导为 { name: string }
type R1 = ReturnType<typeof getUser>;
// 推导为 number
type R2 = ReturnType<typeof getNumber>;

// 使用推导出的类型
var r1: R1 = { name: "Bob" };
var r2: R2 = 100;

console.log("用户: " + JSON.stringify(r1));
console.log("数字: " + r2);

运行结果：

```
用户: {"name":"Bob"}
数字: 100

```

infer 的作用：infer 就像类型系统中的"变量"，它可以捕获类型中的特定部分并在结果类型中使用。

### 分布条件类型

当条件类型的泛型参数是联合类型时，会自动进行"分布"处理。

也就是说，条件会对联合类型中的每个成员分别执行，然后合并结果。

### 实例

// ToArray 会将类型 T 转换为数组类型
// 当 T 是联合类型时，会自动分布处理每个类型
type ToArray<T> = T extends any ? T[] : never;

// 联合类型会自动分布
// string | number 会分布为：ToArray<string> | ToArray<number>
// 即：string[] | number[]
type StrOrNum = ToArray<string | number>;

// 可以赋值 string[] 或 number[]
var arr: StrOrNum = ["hello"];
// 也可以赋值 number
arr = 42;

console.log("数组: " + arr);

运行结果：

```
数组: 42

```

分布机制：条件类型的分布特性是自动开启的。如果想禁用分布，可以使用方括号：`[T] extends U`。

### 条件类型与映射类型结合

条件类型可以与映射类型结合，创建强大的类型转换工具。

这种组合是实现 TypeScript 内置工具类型的基础。

### 实例

// 定义用户接口
interface User {
// 用户 ID
id: number;
// 用户名
name: string;
// 用户邮箱
email: string;
}

// 使用映射类型和条件类型实现 Partial
// 遍历 T 的所有属性，添加可选修饰符 ?
type Partial<T> = {
[P in keyof T]?: T[P];
};

// 使用映射类型和条件类型实现 Required
// 遍历 T 的所有属性，移除可选修饰符 -?
type Required<T> = {
[P in keyof T]-?: T[P];
};

// 使用 Partial：所有属性变为可选
var partial: Partial<User> = { name: "Alice" };

// 使用 Required：将 Partial<User> 的属性变为必填
// 需要先有 Partial<User> 类型
type RequiredUser = Required<Partial<User>>;
var required: RequiredUser = { name: "Bob", id: 1 };

console.log("可选: " + JSON.stringify(partial));
console.log("必填: " + JSON.stringify(required));

运行结果：

```
可选: {"name":"Alice"}
必填: {"name":"Bob","id":1}

```

组合使用：条件类型和映射类型的组合可以实现各种复杂的类型转换，如 Partial、Required、Readonly 等内置工具类型。

### 高级示例：类型检查

条件类型还可以用于实现更复杂的类型检查。

下面是一些实际开发中有用的类型检查示例。

### 实例

// 检查类型是否为 any
// any 与任何类型交叉都会得到 any，0 extends any 为 true
type IsAny<T> = 0 extends (1 & T) ? true : false;

// 测试 IsAny
// any 是特殊类型，与任何类型交叉都返回 any
type A = IsAny<any>; // true
// string 不是 any
type B = IsAny<string>; // false

console.log("any 是 any?: " + A);
console.log("string 是 any?: " + B);

// 检查类型是否可以赋值
// 如果 T 可以赋值给 U，返回 true，否则返回 false
type IsAssignableTo<T, U> = T extends U ? true : false;

// 测试类型可赋值性
// string 可以赋值给 any，所以返回 true
type CanAssign = IsAssignableTo<string, any>;
console.log("string 赋值给 any?: " + CanAssign);

运行结果：

```
any 是 any?: true
string 是 any?: true

```

注意：any 是 TypeScript 中最"宽松"的类型，它可以赋值给任何类型，也可以接收任何类型的赋值。

### 注意事项

- 延迟求值：条件类型是延迟求值的，只有在使用具体类型时才会进行计算
- 分布特性：联合类型会自动触发条件类型的分布机制
- infer 只能用在 extends 条件中：这是 infer 唯一可以出现的位置
- 配合映射类型：条件类型与映射类型结合可以创建强大的工具类型

进阶：许多 TypeScript 内置工具类型（如 Partial、Required、Extract）都是用条件类型实现的。

### 总结

条件类型是 TypeScript 类型系统中最强大的特性之一。

- 基本语法：`T extends U ? X : Y`
- infer 关键字：从类型中推导特定部分
- 分布特性：联合类型自动分布处理
- 应用场景：类型过滤、类型推导、类型检查
- 工具类型：大多数内置工具类型基于条件类型实现

最佳实践：熟练掌握条件类型可以让你写出更加灵活和类型安全的代码。

---

## TypeScript 映射类型

Source: https://www.runoob.com/typescript/ts-mapped-types.html

## TypeScript 映射类型

映射类型（Mapped Types）是 TypeScript 中一种基于已有类型创建新类型的强大特性。

它允许开发者批量修改属性的特性，如将所有属性变为可选、将所有属性变为只读等。

映射类型是实现 TypeScript 内置工具类型的核心技术。 映射类型工作原理 原始类型 interface User { id: number; name: string; } 映射语法 { [P in keyof T]?: T[P] // 添加可选修饰符 } 结果类型 type PartialUser = { id?: number; name?: string; } 映射类型修饰符 ? 前缀 添加可选修饰符 type Partial<T> readonly 前缀 添加只读修饰符 type Readonly<T> -? 前缀 移除可选修饰符 type Required<T> as 关键字 重映射键名 as `${P}`

### 为什么需要映射类型

在实际的 TypeScript 开发中，我们经常需要基于现有类型创建变体。

例如，需要一个所有属性都是可选的版本，或者所有属性都是只读的版本。

传统的方式是手动定义这些类型，既繁琐又容易出错。

映射类型提供了一种声明式的方式来自动生成这些类型变体。

概念说明：映射类型使用 `keyof` 和 `in` 关键字遍历已有类型的所有键，然后对每个键应用相同的类型转换。

### 基础映射类型

基础映射类型通过遍历原始类型的所有键来创建新类型。

这是实现 Partial、Readonly 等工具类型的基础。

### 实例

// 定义用户接口
interface User {
// 用户 ID
id: number;
// 用户名
name: string;
// 用户邮箱
email: string;
}

// 实现 Partial 工具类型
// 遍历 T 的所有键，添加可选修饰符 ?
type Partial<T> = {
// P 遍历 keyof T 返回的所有键
// T[P] 获取原类型中对应键的值类型
[P in keyof T]?: T[P];
};

// 使用 Partial 类型
type PartialUser = Partial<User>;

// PartialUser 类型等同于：
// { id?: number; name?: string; email?: string }

// 可以只提供部分属性
var user: PartialUser = { name: "Alice" };

console.log("部分用户: " + JSON.stringify(user));

运行结果：

```
部分用户: {"name":"Alice"}

```

语法解释：`[P in keyof T]` 表示遍历 T 类型的所有键。`?` 修饰符将属性设为可选。

### 属性修饰符

映射类型支持多种属性修饰符来修改属性的特性。

这些修饰符可以组合使用，实现不同的类型转换需求。

### 实例

// 定义用户接口
interface User {
// 用户名
name: string;
// 用户年龄
age: number;
}

// 使用 readonly 映射：将所有属性变为只读
// 添加 readonly 修饰符
type Readonly<T> = {
readonly [P in keyof T]: T[P];
};

// 使用可选映射：将所有属性变为可选（基础版）
type Optional<T> = {
[P in keyof T]?: T[P];
};

// 使用 -? 映射：移除可选修饰符（变为必填）
// -? 会移除原有的 ? 修饰符
type Required<T> = {
[P in keyof T]-?: T[P];
};

// 测试只读类型
var readonlyUser: Readonly<User> = { name: "Alice", age: 25 };
// readonlyUser.age = 30; // 错误：只读属性不能修改

// 测试可选类型
var optionalUser: Optional<User> = { name: "Bob" };

console.log("只读: " + JSON.stringify(readonlyUser));
console.log("可选: " + JSON.stringify(optionalUser));

运行结果：

```
只读: {"name":"Alice","age":25}
可选: {"name":"Bob"}

```

修饰符说明：`?` 添加可选，`-?` 移除可选；`readonly` 添加只读，`-readonly` 移除只读。

### 键名映射

映射类型还可以使用 `as` 关键字来重映射键名。

这在需要统一修改键名格式时非常有用。

### 实例

// 定义用户接口
interface User {
// 用户 ID
id: number;
// 用户名
name: string;
// 用户年龄
age: number;
}

// 使用 as 关键字重映射键名
// 为所有键添加前缀
type WithPrefix<T, Prefix extends string> = {
// 使用模板字面量类型重命名键
// Capitalize 将首字母大写
[P in keyof T as `${Prefix}${Capitalize<string & P>}`]: T[P];
};

// 使用 WithPrefix 添加 "user" 前缀
type PrefixedUser = WithPrefix<User, "user">;

// 转换后的类型：
// { userId: number; userName: string; userAge: number }

// 使用带前缀的类型
var user: PrefixedUser = { userId: 1, userName: "Alice", userAge: 25 };

console.log("带前缀: " + JSON.stringify(user));

运行结果：

```
带前缀: {"userId":1,"userName":"Alice","userAge":25}

```

模板字面量类型：使用 ``${Prefix}${Capitalize}`` 可以动态生成新的键名。

### 键过滤

通过条件类型和映射类型的组合，可以实现键的过滤。

这在实现 Omit 等工具类型时非常有用。

### 实例

// 定义用户接口
interface User {
// 用户 ID
id: number;
// 用户名
name: string;
// 用户密码
password: string;
// 用户邮箱
email: string;
}

// 实现 Omit：排除指定键
// 使用条件类型过滤键
type Omit<T, K extends keyof T> = {
// P 遍历 T 的所有键
// 如果 P 可以赋值给 K（即在排除列表中），返回 never（不包含）
// 否则返回 P（保留该键）
[P in keyof T as P extends K ? never : P]: T[P];
};

// 使用 Omit 排除 password 键
type UserWithoutPassword = Omit<User, "password">;

// 转换后的类型：
// { id: number; name: string; email: string }

// 使用排除 password 后的类型
var user: UserWithoutPassword = { id: 1, name: "Alice", email: "a@b.com" };

console.log("无密码: " + JSON.stringify(user));

运行结果：

```
无密码: {"id":1,"name":"Alice","email":"a@b.com"}

```

never 类型：在映射类型中使用 never 作为属性类型，该属性会被完全移除。

### 条件映射

映射类型可以与条件类型结合，根据属性类型的不同应用不同的转换。

这使得类型转换更加灵活和智能。

### 实例

// 定义 API 响应接口
interface APIResponse {
// 响应数据
data: string;
// 错误信息
error: string;
// 是否加载中
isLoading: boolean;
// 时间戳
timestamp: number;
}

// 将函数类型转换为 () => void
// 遍历所有属性，根据属性类型进行条件转换
type FunctionToVoid<T> = {
// 如果 T[P] 是函数类型，转换为 () => void
// 否则保持原类型不变
[P in keyof T]: T[P] extends (...args: any[]) => any
? () => void
: T[P];
};

// 使用条件映射
var response: FunctionToVoid<APIResponse> = {
data: "hello",
error: "",
isLoading: false,
timestamp: Date.now()
};

console.log("响应: " + JSON.stringify(response));

运行结果：

```
响应: {"data":"hello","error":"","isLoading":false,"timestamp":...}

```

应用场景：条件映射常用于处理 API 响应、清理配置对象等需要根据类型做不同处理的场景。

### 内置映射类型

TypeScript 内置了许多基于映射类型实现的工具类型。

这些工具类型可以满足大多数日常开发需求。

### 实例

// Partial - 将所有属性变为可选
type P1 = Partial<{ a: string; b: number }>;
// 结果：{ a?: string; b?: number }

// Required - 将所有可选属性变为必填
type R1 = Required<{ a?: string; b?: number }>;
// 结果：{ a: string; b: number }

// Readonly - 将所有属性变为只读
type RO1 = Readonly<{ a: string; b: number }>;
// 结果：{ readonly a: string; readonly b: number }

// Pick - 选择指定的属性
type PK = Pick<{ a: string; b: number; c: boolean }, "a" | "b">;
// 结果：{ a: string; b: number }

// Omit - 排除指定的属性
type OM = Omit<{ a: string; b: number; c: boolean }, "c">;
// 结果：{ a: string; b: number }

// 测试 Partial
console.log("Partial: " + JSON.stringify({} as P1));

// 测试 Pick
console.log("Pick: " + JSON.stringify({ a: "x" } as PK));

工具类型组合：这些内置工具类型都是基于映射类型和条件类型实现的。了解其原理可以更好地使用它们。

### 注意事项

- keyof 关键字：用于获取类型的所有键组成的联合类型
- in 关键字：用于遍历键名联合类型
- 修饰符位置：`?` 和 `readonly` 在属性名前，表示添加修饰符
- 减号修饰符：`-?` 和 `-readonly` 用于移除修饰符
- as 关键字：用于重映射键名，必须返回字符串或数字字面量类型

进阶：映射类型可以与条件类型、模板字面量类型组合，实现复杂的类型转换。

### 总结

映射类型是 TypeScript 类型系统中最强大的特性之一。

- keyof：获取类型的所有键
- in：遍历键名进行映射
- `?`：添加可选修饰符
- readonly：添加只读修饰符
- -?：移除可选修饰符
- as：重映射键名

最佳实践：善用映射类型可以大幅减少重复的类型定义，提高代码的可维护性。

---

## TypeScript 类继承与多态

Source: https://www.runoob.com/typescript/ts-class-inheritance.html

## TypeScript 类继承与多态

TypeScript 支持面向对象编程的继承和多态特性。 类继承层次结构 Animal（父类） Dog（子类） extends Animal Cat（子类） extends Animal Dog实例 Cat实例 继承特性 • extends 继承 • super() 调用 • 方法重写 • 多态 • instanceof

### 类的继承

使用 extends 关键字实现继承。

### 实例

class Animal {
name: string;

constructor(name: string) {
this.name = name;
}

speak(): void {
console.log(this.name + " 发出声音");
}
}

class Dog extends Animal {
breed: string;

constructor(name: string, breed: string) {
super(name); // 调用父类构造函数
this.breed = breed;
}

speak(): void {
console.log(this.name + " 汪汪汪!");
}
}

var dog = new Dog("旺财", "金毛");
dog.speak();

运行结果：

```
旺财 汪汪汪!

```

### super 关键字

super 用于调用父类的方法和构造函数。

### 实例

class Shape {
color: string;

constructor(color: string) {
this.color = color;
}

describe(): string {
return "这是一个 " + this.color + " 的图形";
}
}

class Circle extends Shape {
radius: number;

constructor(color: string, radius: number) {
super(color);
this.radius = radius;
}

// 重写父类方法
describe(): string {
// 调用父类方法并扩展
return super.describe() + "，半径是 " + this.radius;
}

area(): number {
return Math.PI * this.radius * this.radius;
}
}

var circle = new Circle("红色", 5);
console.log(circle.describe());
console.log("面积: " + circle.area().toFixed(2));

运行结果：

```
这是一个 红色的图形，半径是 5
面积: 78.54

```

### 多态

子类的实例可以赋值给父类类型。

### 实例

class Animal {
name: string;
constructor(name: string) { this.name = name; }
speak(): void {
console.log(this.name + " 发出声音");
}
}

class Cat extends Animal {
speak(): void {
console.log(this.name + " 喵喵喵!");
}
}

class Dog extends Animal {
speak(): void {
console.log(this.name + " 汪汪汪!");
}
}

// 多态：数组中存储不同子类的实例
var animals: Animal[] = [
new Cat("小白"),
new Dog("旺财"),
new Animal("动物")
];

// 调用同一方法，不同子类有不同实现
for (var _i = 0, animals_1 = animals; _i < animals_1.length; _i++) {
var animal = animals_1[_i];
animal.speak();
}

运行结果：

```
小白 喵喵喵!
旺财 汪汪汪!
动物 发出声音

```

### instanceof 检查

使用 instanceof 检查实例类型。

### 实例

class Rectangle {
width: number;
height: number;
constructor(width: number, height: number) {
this.width = width;
this.height = height;
}
area(): number {
return this.width * this.height;
}
}

class Circle {
radius: number;
constructor(radius: number) {
this.radius = radius;
}
area(): number {
return Math.PI * this.radius ** 2;
}
}

var shapes = [new Rectangle(4, 5), new Circle(3)];

for (var _i = 0, shapes_1 = shapes; _i < shapes_1.length; _i++) {
var shape = shapes_1[_i];
if (shape instanceof Rectangle) {
console.log("矩形面积: " + shape.area());
} else if (shape instanceof Circle) {
console.log("圆形面积: " + shape.area().toFixed(2));
}
}

运行结果：

```
矩形面积: 20
圆形面积: 28.27

```

### protected 成员

protected 成员在子类中可见。

### 实例

class Person {
protected name: string;

constructor(name: string) {
this.name = name;
}
}

class Employee extends Person {
private department: string;

constructor(name: string, department: string) {
super(name);
this.department = department;
}

public introduce(): string {
// 可以访问 protected 成员
return "我是 " + this.name + "，在 " + this.department + " 工作";
}
}

var emp = new Employee("Alice", "技术部");
console.log(emp.introduce());

// console.log(emp.name); // 错误：protected 外部不可访问

运行结果：

```
我是 Alice，在 技术部 工作

```

### 总结

- 继承：extends 关键字
- super：调用父类
- 多态：同一接口不同实现
- protected：子类可见

---

## TypeScript 错误处理

Source: https://www.runoob.com/typescript/ts-error-handling.html

## TypeScript 错误处理

错误处理是保证程序健壮性的重要环节。

TypeScript 提供了强大的类型系统，可以帮助开发者更好地处理和预防错误。

本文介绍 TypeScript 中常见的错误处理模式和最佳实践。 错误处理方式 try-catch try { // 可能出错的代码 } catch (e) { // 处理错误 } 捕获并处理异常 Result 类型 type Result = | { ok: true; value: T } | { ok: false; error: E } 返回结果而非抛异常 自定义错误 class AppError extends Error { code: string; } 结构化错误信息 最佳实践 明确错误类型 集中错误处理 避免异常泄漏

### 为什么需要良好的错误处理

任何程序都可能遇到错误情况，如网络请求失败、文件不存在、用户输入错误等。

良好的错误处理可以防止程序崩溃，提供友好的错误提示，并帮助开发者定位问题。

TypeScript 的类型系统可以在编译阶段发现潜在问题，减少运行时错误。

概念说明：错误处理有两种主要方式：异常处理（try-catch）和返回值处理（Result 类型）。前者使用抛出异常表示错误，后者使用返回值携带错误信息。

### 自定义错误类型

通过扩展 Error 类创建自定义错误类型，可以携带更多错误信息。

这使得错误处理更加精确和结构化。

### 实例

// 定义应用程序错误类
// 扩展内置 Error 类，添加错误码
class AppError extends Error {
// 错误码，用于程序化处理错误
code: string;

// 构造函数
constructor(message: string, code: string) {
super(message); // 调用父类构造函数
this.name = "AppError"; // 设置错误名称
this.code = code; // 保存错误码
}
}

// 安全的除法函数
function divide(a: number, b: number): number {
// 检查除数是否为零
if (b === 0) {
// 抛出自定义错误
throw new AppError("Cannot divide by zero", "DIVIDE_BY_ZERO");
}
return a / b;
}

// 使用 try-catch 捕获错误
try {
var result = divide(10, 0);
} catch (error) {
// 检查错误类型
if (error instanceof AppError) {
console.log("应用错误: " + error.message + ", 代码: " + error.code);
} else {
console.log("未知错误: " + error);
}
}

运行结果：

```
应用错误: Cannot divide by zero, 代码: DIVIDE_BY_ZERO

```

错误码：为错误添加代码可以帮助程序更精确地处理不同类型的错误。

### Result 类型避免异常

另一种错误处理方式是使用 Result 类型。

它通过返回值携带错误信息，而不是抛出异常。这种方式在函数式编程中很常见。

### 实例

// 定义 Result 类型，使用联合类型
// 成功时包含 ok: true 和值，失败时包含 ok: false 和错误
type Result<T, E = Error> =
| { ok: true; value: T }
| { ok: false; error: E };

// 使用 Result 类型的除法函数
function safeDivide(a: number, b: number): Result<number, string> {
// 检查除数是否为零
if (b === 0) {
// 返回错误结果
return { ok: false, error: "Cannot divide by zero" };
}
// 返回成功结果
return { ok: true, value: a / b };
}

// 调用函数并处理结果
var result = safeDivide(10, 2);

// 根据结果类型进行处理
if (result.ok) {
console.log("结果: " + result.value);
} else {
console.log("错误: " + result.error);
}

运行结果：

```
结果: 5

```

优势：Result 类型让错误处理变得显式，调用者必须处理可能的错误，而不会忽略它。

### Async 函数错误处理

在异步函数中，错误处理尤为重要。

可以使用 try-catch 或 Result 类型来处理异步操作中的错误。

### 实例

// 定义用户接口
interface User {
id: number;
name: string;
}

// 模拟获取用户的异步函数
async function fetchUser(id: number): Promise<Result<User, Error>> {
try {
// 模拟网络请求
var response = await fetch("/api/users/" + id);
var user = await response.json();
// 返回成功结果
return { ok: true, value: user };
} catch (error) {
// 返回错误结果
return { ok: false, error: error as Error };
}
}

// 主函数
async function main() {
// 调用异步函数
var result = await fetchUser(1);

// 处理结果
if (result.ok) {
console.log("用户: " + JSON.stringify(result.value));
} else {
console.log("错误: " + result.error.message);
}
}

// 执行主函数
main();

运行结果：

```
用户: {"id":1,"name":"Alice"}

```

提示：异步函数中的 try-catch 会捕获 await 表达式抛出的任何错误。

### 通用错误处理封装

可以创建一个通用的错误处理函数，简化异步代码的错误处理。

### 实例

// 通用错误处理包装函数
// 接受一个异步函数，返回 Result 类型
async function withErrorHandling<T>(
fn: () => Promise<T>
): Promise<Result<T, Error>> {
try {
// 执行传入的异步函数
var data = await fn();
// 返回成功结果
return { ok: true, value: data };
} catch (error) {
// 返回错误结果
return { ok: false, error: error as Error };
}
}

// 使用通用错误处理
// 模拟获取数据
var result = await withErrorHandling(async function() {
var response = await fetch("/api/data");
return response.json();
});

// 根据结果处理
if (result.ok) {
console.log("数据: " + JSON.stringify(result.value));
} else {
console.error("错误:", result.error);
}

最佳实践：封装通用的错误处理逻辑可以减少代码重复，提高代码可维护性。

### 注意事项

- 不要忽略错误：不要使用空的 catch 块捕获并忽略错误
- 明确错误类型：尽量使用具体的错误类型，而不是通用的 Error
- 错误边界：在应用中建立统一的错误处理机制
- 不要过度使用异常：对于可预期的错误，优先使用返回值而非抛异常

建议：根据场景选择错误处理方式：程序错误用异常，业务错误用 Result。

### 总结

良好的错误处理是构建健壮应用的基础。

- 自定义错误：扩展 Error 类，添加错误码等信息
- Result 类型：通过返回值处理错误，避免异常
- async/await：使用 try-catch 处理异步错误
- 错误封装：创建通用错误处理函数
- 错误边界：建立统一的错误处理机制

最佳实践：根据具体场景选择合适的错误处理方式，平衡代码可读性和健壮性。

---

## TypeScript 交叉类型

Source: https://www.runoob.com/typescript/ts-intersection-types.html

## TypeScript 交叉类型

交叉类型（Intersection Types）将多个类型合并成一个新类型，新类型包含所有成员的类型。

这类似于面向对象中的多重继承，让一个类型可以拥有多个类型的特性。 交叉类型工作原理 类型 A (Person) name: string age: number & 类型 B (Worker) company: string salary: number 合并 A & B (Employee) name: string age: number company: string salary: number 交叉类型使用场景 类型组合 合并多个类型特性 Mixin 模式 组合多个类的功能 替代接口继承 更简洁的类型组合

### 为什么需要交叉类型

在开发中，一个类型往往需要拥有多个类型的特性。

例如，一个员工既是一个人（Person），也是一个工作者（Worker），需要同时具有两者的属性。

交叉类型让我们可以将多个类型合并成一个，满足这种需求。

概念说明：交叉类型使用 `&` 符号连接多个类型，表示新类型包含所有类型的成员。这类似于多重继承的概念。

### 基本语法

使用 `&` 符号组合多个类型。

### 实例

// 定义人员类型
// 包含姓名和年龄
interface Person {
name: string;
age: number;
}

// 定义工作者类型
// 包含公司名称和薪资
interface Worker {
company: string;
salary: number;
}

// 使用交叉类型合并两个接口
// Employee 类型同时具有 Person 和 Worker 的所有属性
type Employee = Person & Worker;

// 创建同时具有两个类型特性的对象
var employee: Employee = {
name: "Alice",
age: 25,
company: "Google",
salary: 100000
};

console.log("员工: " + JSON.stringify(employee));

运行结果：

```
员工: {"name":"Alice","age":25,"company":"Google","salary":100000}

```

说明：交叉类型 `A & B` 意味着新类型同时具有 A 和 B 的所有属性，缺一不可。

### 交叉类型与接口继承

交叉类型可以替代接口的多重继承。

### 实例

// 定义类型 A
interface A {
a: string;
}

// 定义类型 B
interface B {
b: number;
}

// 使用接口继承多个接口
// 需要使用 extends 继承多个接口
interface AB extends A, B {
c: boolean;
}

// 使用交叉类型（更简洁）
// 直接使用 & 符号组合类型
type ABType = A & B & { c: boolean };

// 两种方式都能创建包含所有属性的类型
var obj: ABType = { a: "hello", b: 42, c: true };
console.log("对象: " + JSON.stringify(obj));

运行结果：

```
对象: {"a":"hello","b":42,"c":true}

```

对比：交叉类型比接口继承更简洁，特别是当需要继承多个类型且需要添加额外属性时。

### 类型混合（Mixin 模式）

使用交叉类型实现 Mixin 模式，可以动态组合类的功能。

### 实例

// 定义构造函数类型
// 接受任意参数，返回一个对象
type Constructor = new (...args: any[]) => {};

// Mixin：添加时间戳功能
// 返回一个扩展了 Base 的新类
function Timestamped<T extends Constructor>(Base: T) {
return class extends Base {
timestamp = Date.now();
};
}

// Mixin：添加序列化功能
// 返回一个扩展了 Base 的新类，包含 serialize 方法
function Serializable<T extends Constructor>(Base: T) {
return class extends Base {
serialize() {
return JSON.stringify(this);
}
};
}

// 基础用户类
class User {
name: string;
constructor(name: string) {
this.name = name;
}
}

// 组合 Mixin
// 创建具有时间戳功能的用户类
var TimestampedUser = Timestamped(User);
// 创建具有序列化功能的用户类
var SerializableUser = Serializable(User);
// 组合两个 Mixin
var FullUser = Serializable(Timestamped(User));

// 创建实例并测试
var user = new FullUser("Alice");
console.log("时间戳: " + user.timestamp);
console.log("序列化: " + user.serialize());

运行结果：

```
时间戳: 17134...
序列化: {"name":"Alice","timestamp":17134...}

```

Mixin 模式：这是一种强大的模式，可以在不修改原有类的情况下为其添加新功能。

### 交叉类型与联合类型

交叉类型和联合类型结合时需要特别注意优先级。

### 实例

// 联合类型：可以是字符串或数字
type StringOrNumber = string | number;

// 交叉类型：不兼容类型的交叉
// string & number = never（没有类型同时是字符串和数字）
type Both = string & number;

// 定义三个类型
type A = { a: string };
type B = { b: number };
type C = { c: boolean };

// 联合类型与交叉类型结合
// (A | B) & C 会将联合类型的每个分支都与 C 交叉
type Combined = (A | B) & C;

// 实际结果是：{ a: string; c: boolean } | { b: number; c: boolean }
// 即要么是 A + C，要么是 B + C
var obj: Combined = { a: "hello", c: true };
console.log("组合: " + JSON.stringify(obj));

运行结果：

```
组合: {"a":"hello","c":true}

```

重要：不兼容的类型进行交叉会得到 `never` 类型。例如 `string & number` 是无效的。

### 实用交叉类型

交叉类型常用于创建工具类型，实现类型的转换。

### 实例

// 映射类型：将所有属性变为可选
// 遍历 T 的所有属性，添加 ? 使其可选
type Partial<T> = { [P in keyof T]?: T[P] };

// 映射类型：将所有属性变为必填
// 遍历 T 的所有属性，移除 ? 使其必填
type Required<T> = { [P in keyof T]-?: T[P] };

// 映射类型：将所有属性变为只读
// 遍历 T 的所有属性，添加 readonly
type Readonly<T> = { readonly [P in keyof T]: T[P] };

// 定义配置接口
interface Config {
host: string;
port: number;
}

// 使用工具类型
var partialConfig: Partial<Config> = { host: "localhost" };
var requiredConfig: Required<Config> = { host: "localhost", port: 8080 };
var readonlyConfig: Readonly<Config> = { host: "localhost", port: 8080 };

console.log("部分: " + JSON.stringify(partialConfig));
console.log("必填: " + JSON.stringify(requiredConfig));
console.log("只读: " + JSON.stringify(readonlyConfig));

运行结果：

```
部分: {"host":"localhost"}
必填: {"host":"localhost","port":8080}
只读: {"host":"localhost","port":8080}

```

工具类型：TypeScript 标准库中提供了许多基于交叉类型和映射类型的工具类型，如 Partial、Required、Readonly 等。

### 注意事项

- 不兼容类型：交叉不兼容的类型会得到 never
- 优先级：联合类型的优先级高于交叉类型
- 方法冲突：如果两个类型有同名的方法，需要手动处理冲突

最佳实践：交叉类型适合组合多个接口或类型，当需要多重继承时，优先使用交叉类型而非接口继承。

### 总结

交叉类型是 TypeScript 中强大的类型组合工具。

- 语法：使用 `&` 符号连接多个类型
- 合并：新类型包含所有类型的成员
- Mixin：可以用于实现类的功能组合
- never：不兼容类型交叉会得到 never 类型

建议：合理使用交叉类型可以创建灵活的类型组合，但要注意避免不必要类型交叉导致的 never 类型。

---

## TypeScript 模板字面量类型

Source: https://www.runoob.com/typescript/ts-template-literal.html

## TypeScript 模板字面量类型

模板字面量类型基于字符串字面量类型构建，支持通过插值生成新的字符串类型。

这让 TypeScript 能够对字符串进行更精确的类型检查，适用于事件名、路径、类名等场景。 模板字面量类型工作原理 输入类型 type Event = "click" | "focus" type Method = "get" | "post" 模板 模板字面量类型 type Name = `on${Capitalize}` type Path = `${Method}:/${string}` 输出类型 "onClick" "onFocus" "get:/..." 内置工具类型 Uppercase "hello" → "HELLO" Lowercase "HELLO" → "hello" Capitalize "hello" → "Hello" Uncapitalize "Hello" → "hello"

### 为什么需要模板字面量类型

在开发中，我们经常需要处理格式化的字符串，如事件名（onClick）、API 路径（get:/users）、CSS 类名（btn-primary-md）等。

使用普通的 string 类型无法精确描述这些格式，而模板字面量类型让我们可以精确地定义这些字符串的类型。

这大大增强了 TypeScript 的类型安全性，减少了运行时错误。

概念说明：模板字面量类型使用反引号（`` ` ``）和 `${}` 插值语法来定义字符串类型，类似于 JavaScript 的模板字符串，但用在类型层面。

### 基本语法

模板字面量类型使用反引号和插值来定义类型。

### 实例

// 定义基础字符串字面量类型
type World = "world";

// 使用模板字面量类型
// `Hello ${World}` 相当于 "Hello world"
type Greeting = `Hello ${World}`;

// 只能赋值符合类型定义的字符串
var greeting: Greeting = "Hello world";
console.log("问候: " + greeting);

运行结果：

```
问候: Hello world

```

说明：模板字面量类型会将插值的位置替换为实际的字符串，生成新的字面量类型。

### 内置工具类型

TypeScript 提供了四个内置的工具类型来处理字符串大小写。

### 实例

// Uppercase：将字符串转为大写
type UpperHello = Uppercase<"hello">; // "HELLO"

// Lowercase：将字符串转为小写
type LowerHELLO = Lowercase<"HELLO">; // "hello"

// Capitalize：将字符串首字母大写
type CapitalizedHello = Capitalize<"hello">; // "Hello"

// Uncapitalize：将字符串首字母小写
type UncapitalizedHello = Uncapitalize<"Hello">; // "hello"

console.log("Uppercase: " + UpperHello);
console.log("Lowercase: " + LowerHELLO);
console.log("Capitalize: " + CapitalizedHello);
console.log("Uncapitalize: " + UncapitalizedHello);

运行结果：

```
Uppercase: HELLO
Lowercase: hello
Capitalize: Hello
Uncapitalize: hello

```

应用场景：这些工具类型在处理事件名、方法名等需要统一格式的场景非常有用。

### 事件类型

使用模板字面量类型可以精确地定义事件名称的类型。

### 实例

// 构建事件名类型
// `on${Capitalize<string>}` 生成以 "on" 开头，首字母大写的字符串
type EventName = `on${Capitalize<string>}`;
// `handle${Capitalize<string>}` 生成以 "handle" 开头，首字母大写的字符串
type Handler = `handle${Capitalize<string>}`;

// 只能赋值符合格式的字符串
var clickEvent: EventName = "onClick";
var focusEvent: EventName = "onFocus";
var handler: Handler = "handleSubmit";

console.log("事件: " + clickEvent);
console.log("处理器: " + handler);

运行结果：

```
事件: onClick
处理器: handleSubmit

```

优势：使用模板字面量类型后，像 "onclick"（小写）这样的错误格式会被 TypeScript 拒绝。

### 路径类型

使用模板字面量类型可以精确地定义 API 路径的类型。

### 实例

// 定义 HTTP 方法类型
type HttpMethod = "get" | "post" | "put" | "delete";

// 定义路径格式
type ApiEndpoint = `/${string}`; // 以斜杠开头的字符串

// 组合成完整的 API 路径类型
type ApiPath = `${HttpMethod}${ApiEndpoint}`;

// 只能赋值符合格式的路径
var getUsers: ApiPath = "/get/users";
var createUser: ApiPath = "/post/users";

console.log("路径: " + getUsers);
console.log("路径: " + createUser);

运行结果：

```
路径: /get/users
路径: /post/users

```

类型安全：像 "/users"（没有方法前缀）这样的路径会被 TypeScript 报错。

### 复杂示例

模板字面量类型可以组合多个联合类型，生成所有可能的组合。

### 实例

// 带数字的模板类型
// ${number} 匹配任意数字
type Row = `row${number}`;
type Row10 = Row; // row0, row1, row2... 直到 row9...

// 组合多个类型
type Variant = "primary" | "secondary";
type Size = "sm" | "md" | "lg";
// 这会生成 6 种组合：btn-primary-sm, btn-primary-md, btn-primary-lg...
type ClassName = `btn-${Variant}-${Size}`;

// 只能赋值生成的 6 种组合之一
var className: ClassName = "btn-primary-md";
console.log("类名: " + className);

运行结果：

```
类名: btn-primary-md

```

组合爆炸：模板字面量类型会自动展开所有组合。如果联合类型有很多选项，生成的类型可能会非常庞大。

### 自定义工具类型

可以创建自己的模板字面量工具类型。

### 实例

// 添加前缀的工具类型
// T 是任意字符串，P 是要添加的前缀
type Prefix<T extends string, P extends string> = `${P}${Capitalize<T>}`;

// 添加后缀的工具类型
// T 是任意字符串，S 是要添加的后缀
type Suffix<T extends string, S extends string> = `${Capitalize<T>}${S}`;

// 使用自定义工具类型
type HandlerName = Prefix<"click", "on">;
type ButtonId = Suffix<"submit", "Btn">;

var handler: HandlerName = "onClick";
var id: ButtonId = "SubmitBtn";

console.log("处理器: " + handler);
console.log("ID: " + id);

运行结果：

```
处理器: onClick
ID: SubmitBtn

```

泛型模板：模板字面量类型可以与泛型结合，创建可复用的工具类型。

### 注意事项

- 插值类型：模板中的 `${}` 可以是具体字符串、联合类型、string、number 等
- 组合数量：组合多个联合类型时，生成的类型可能非常大
- 大小写处理：使用内置工具类型处理字符串大小写

最佳实践：使用模板字面量类型处理事件名、路径、类名等有固定格式的字符串，可以获得更好的类型安全。

### 总结

模板字面量类型是 TypeScript 强大的类型系统的一部分。

- 模板语法：使用 `${T}` 插值构建类型
- 内置工具：Uppercase、Lowercase、Capitalize、Uncapitalize
- 应用场景：事件名、API 路径、CSS 类名等
- 自定义：创建可复用的工具类型

建议：在需要格式化字符串的场景，优先使用模板字面量类型来获得编译期的类型检查。

---

## TypeScript 从 JavaScript 迁移

Source: https://www.runoob.com/typescript/ts-migration.html

## TypeScript 从 JavaScript 迁移

将现有 JavaScript 项目逐步迁移到 TypeScript。

### 迁移策略

- 添加 tsconfig.json
- 重命名 .js 为 .ts
- 逐步添加类型注解
- 启用严格模式

### 配置 tsconfig.json

### tsconfig.json

{
"compilerOptions": {
// 初始阶段：宽松配置
"target": "ES2020",
"module": "commonjs",
"strict": false,
"noImplicitAny": false,
"strictNullChecks": false,
"skipLibCheck": true,

// 允许 JS 文件
"allowJs": true,
"checkJs": false,

// 输出目录
"outDir": "./dist",
"rootDir": "./src"
},
"include": ["src/**/*"],
"exclude": ["node_modules", "dist"]
}

### 逐步启用严格检查

### 分阶段启用

// 阶段 1: 基础迁移
{
"compilerOptions": {
"strict": false,
"noImplicitAny": false
}
}

// 阶段 2: 启用类型检查
{
"compilerOptions": {
"strict": true,
"noImplicitAny": true,
"strictNullChecks": true
}
}

// 阶段 3: 完全严格
{
"compilerOptions": {
"strict": true,
"noImplicitAny": true,
"strictNullChecks": true,
"strictFunctionTypes": true,
"strictPropertyInitialization": true
}
}

### JSDoc 类型注释

在 JavaScript 中使用 JSDoc 添加类型。

### utils.js

/**
* @param {number} a
* @param {number} b
* @returns {number}
*/
function add(a, b) {
return a + b;
}

/**
* @typedef {Object} User
* @property {number} id
* @property {string} name
* @property {string} email
*/

/**
* @param {number} id
* @returns {Promise<User>}
*/
function getUser(id) {
return fetch(`/api/users/${id}`).then(r => r.json());
}

运行结果：

```
JSDoc 注释添加成功

```

### 类型声明文件

为没有类型定义的模块创建声明。

### src/types/my-module.d.ts

declare module "my-module" {
export function doSomething(param: string): void;
export class MyClass {
constructor(options: { name: string });
name: string;
}
}

### declare 关键字

### 实例

// 声明全局变量
declare var GLOBAL_CONFIG: {
apiUrl: string;
version: string;
};

// 声明全局函数
declare function myFunction(param: string): void;

// 声明命名空间
declare namespace MyNamespace {
function doSomething(): void;
}

// 使用
console.log(GLOBAL_CONFIG.apiUrl);
myFunction("hello");
MyNamespace.doSomething();

运行结果：

```
声明成功

```

### 迁移工具

- tsc --allowJs：编译 JS 文件
- checkJs：检查 JS 类型
- // @ts-check：单文件类型检查
- // @ts-ignore：忽略错误

### legacy.js

// @ts-check
// @ts-ignore
var result = someLegacyFunction();

### 最佳实践

- 从关键模块开始迁移
- 添加单元测试
- 逐步启用严格模式
- 使用 JSDoc 注释
- 创建类型声明文件

### 总结

- 渐进式：逐步迁移
- JSDoc：类型注释
- 声明文件：.d.ts
- 严格模式：分阶段启用

---

## TypeScript 单元测试

Source: https://www.runoob.com/typescript/ts-unit-testing.html

## TypeScript 单元测试

TypeScript 项目中的单元测试实践，确保代码质量。

单元测试可以验证代码的正确性，TypeScript 的类型系统与测试框架完美结合，可以编写类型安全的测试代码。 TypeScript 单元测试流程 源代码 src/utils/calculator.ts src/services/userService.ts 完整类型注解 测试代码 *.test.ts describe() / it() expect() 断言 测试结果 ✓ 通过 / ✗ 失败 覆盖率报告 错误堆栈跟踪 单元测试原则 每个测试只验证一件事 Arrange-Act-Assert 结构 测试应该相互独立

### 为什么需要单元测试

单元测试是保证代码质量的重要手段，它可以验证代码的正确性，防止 bug 出现。

TypeScript 项目中，测试代码同样受益于类型系统：类型错误会在编译时被发现，IDE 提供智能提示，测试代码更可靠。

此外，测试代码也是最好的文档，可以通过测试了解函数/类的预期行为。

质量保障：单元测试可以快速发现回归问题，确保代码改动不会破坏现有功能。

### 测试框架配置

Jest 是 TypeScript 项目最流行的测试框架。

### 安装 Jest

# 安装 Jest 和相关依赖
# - ts-jest: 让 Jest 能够运行 TypeScript
# - @types/jest: Jest 的类型定义
npm install --save-dev jest ts-jest @types/jest

# 初始化 Jest 配置
npx ts-jest config:init

ts-jest：这是一个预处理器，让 Jest 能够直接运行 TypeScript 文件，无需手动编译。

### 配置 jest.config.js

配置 Jest 测试环境。

### jest.config.js

module.exports = {
// 使用 ts-jest 预设
preset: 'ts-jest',
// 测试环境：node 或 browser
testEnvironment: 'node',
// 测试文件目录
roots: ['<rootDir>/src'],
// 测试文件匹配模式
testMatch: ['**/__tests__/**/*.ts'],
// 支持的文件扩展名
moduleFileExtensions: ['ts', 'js', 'json'],
// 收集覆盖率的文件
collectCoverageFrom: [
'src/**/*.ts',
'!src/**/*.d.ts' // 排除类型声明文件
]
}

配置说明：测试文件通常放在 __tests__ 目录或以 .test.ts 结尾。

### 测试函数

首先编写需要测试的业务代码。

### src/utils/calculator.ts

// 计算器类
export class Calculator {
// 加法
add(a: number, b: number): number {
return a + b;
}

// 减法
subtract(a: number, b: number): number {
return a - b;
}

// 乘法
multiply(a: number, b: number): number {
return a * b;
}

// 除法
divide(a: number, b: number): number {
if (b === 0) {
throw new Error("Cannot divide by zero");
}
return a / b;
}
}

然后编写对应的测试代码。

### src/utils/calculator.test.ts

import { Calculator } from "./calculator";

// 测试套件：Calculator 类的测试
describe("Calculator", () => {
let calculator: Calculator;

// 每个测试前创建新的 Calculator 实例
beforeEach(() => {
calculator = new Calculator();
});

// 加法测试
describe("add", () => {
it("should add two numbers", () => {
expect(calculator.add(2, 3)).toBe(5);
});

it("should handle negative numbers", () => {
expect(calculator.add(-1, 1)).toBe(0);
});
});

// 除法测试
describe("divide", () => {
it("should divide two numbers", () => {
expect(calculator.divide(10, 2)).toBe(5);
});

it("should throw error when dividing by zero", () => {
// 期望抛出错误
expect(() => calculator.divide(10, 0)).toThrow();
});
});
});

运行结果：

```
Calculator
add
✓ should add two numbers
✓ should handle negative numbers
divide
✓ should divide two numbers
✓ should throw error when dividing by zero

```

describe/it：describe 用于分组测试，it（或 test）用于定义单个测试用例。

### 测试 Service

测试 Service 层的业务逻辑。

### src/services/userService.ts

// 用户类型
export interface User {
id: number;
name: string;
}

// 用户服务类
export class UserService {
private users: User[] = [];
private nextId = 1;

// 创建用户
createUser(name: string): User {
const user = { id: this.nextId++, name };
this.users.push(user);
return user;
}

// 获取用户
getUser(id: number): User | undefined {
return this.users.find(u => u.id === id);
}

// 获取所有用户
getAllUsers(): User[] {
return [...this.users];
}
}

### src/services/userService.test.ts

import { UserService } from "./userService";

describe("UserService", () => {
let service: UserService;

beforeEach(() => {
service = new UserService();
});

describe("createUser", () => {
it("should create a user with id", () => {
const user = service.createUser("Alice");
expect(user.id).toBe(1);
expect(user.name).toBe("Alice");
});

it("should increment id for each user", () => {
const user1 = service.createUser("Alice");
const user2 = service.createUser("Bob");
expect(user2.id).toBe(user1.id + 1);
});
});

describe("getUser", () => {
it("should return user by id", () => {
const created = service.createUser("Alice");
const found = service.getUser(created.id);
// 使用可选链和 toBe
expect(found?.name).toBe("Alice");
});

it("should return undefined for non-existent id", () => {
const found = service.getUser(999);
expect(found).toBeUndefined();
});
});
});

运行结果：

```
UserService
createUser
✓ should create a user with id
✓ should increment id for each user
getUser
✓ should return user by id
✓ should return undefined for non-existent id

```

测试隔离：每个测试用例应该独立，使用 beforeEach 确保每个测试都有干净的状态。

### Mock

使用 Mock 模拟依赖，如外部 API、数据库等。

### 实例

// Mock 函数：创建模拟函数
const mockCallback = jest.fn(x => x * 2);

// 使用模拟函数
[1, 2, 3].forEach(mockCallback);

// 验证函数被调用了 3 次
expect(mockCallback).toHaveBeenCalledTimes(3);
// 验证函数被调用时的参数
expect(mockCallback).toHaveBeenCalledWith(2);

// Mock 模块：模拟整个模块
jest.mock("./api", () => ({
fetchUser: jest.fn(() => Promise.resolve({ id: 1, name: "Alice" }))
}));

运行结果：

```
✓ mock 函数被调用 3 次

```

Mock 用途：当测试的代码依赖外部系统时，使用 Mock 可以隔离依赖，只测试目标代码的逻辑。

### 注意事项

- 测试文件位置：放在 __tests__ 目录或使用 .test.ts 后缀
- 测试命名：使用描述性的测试名称，说明预期行为
- 独立测试：每个测试应该独立运行，不依赖其他测试
- 覆盖率：关注核心业务逻辑的测试覆盖率

最佳实践：测试应该快速、可靠、相互独立。遵循 AAA 原则：Arrange（准备）、Act（执行）、Assert（断言）。

### 总结

单元测试是保证 TypeScript 代码质量的重要手段。

- Jest：最流行的 TypeScript 测试框架
- describe：用于分组相关测试
- it/test：定义单个测试用例
- expect：断言测试结果
- Mock：模拟外部依赖

建议：为关键业务逻辑编写测试，确保代码改动不会引入 bug。

---

## TypeScript infer 关键字

Source: https://www.runoob.com/typescript/ts-infer.html

## TypeScript infer 关键字

infer 是 TypeScript 条件类型中的关键字，用于从类型中推断出新的类型。

它允许在泛型条件类型中提取和推导类型，实现强大的类型操作。 infer 关键字工作原理 原始类型 type Promise<T> = { value: T } infer 推导 提取出的类型 type ValueOf<T> = T extends Promise<infer V> ? V : never 示例 ValueOf<Promise<string>> → string infer 常见应用场景 提取 Promise 值类型 提取数组元素类型 提取函数返回类型

### 为什么需要 infer 关键字

在泛型编程中，我们经常需要从复杂的类型中提取出部分类型。

例如，从 `Promise<string>` 中提取 `string`，从 `Array<User>` 中提取 `User`。

infer 关键字让这种类型提取变得优雅和类型安全，它可以在条件类型中声明一个类型变量，然后在 true 分支中使用。

概念说明：infer 是 "infer" 的缩写，意为"推断"。它只能在条件类型的 extends 子句中使用，用于声明一个待推断的类型变量。

### 基本用法

在条件类型中使用 infer 推断类型。

### 实例

// 从 Promise 中提取值类型
// 如果 T 是 Promise<V>，则返回 V；否则返回 never
type ValueOf<T> = T extends Promise<infer V> ? V : never;

// 测试
type Str = ValueOf<Promise<string>>; // string
type Num = ValueOf<Promise<number>>; // number
type NotPrm = ValueOf<string>; // never

// 使用示例
var promise: Promise<string> = Promise.resolve("hello");
var value: ValueOf<typeof promise> = "world";

console.log("字符串提取: " + (typeof value === "string" ? "成功" : "失败"));

运行结果：

```
字符串提取: 成功

```

说明：`infer V` 声明了一个类型变量 V，TypeScript 会自动推断 V 的类型。

### 提取数组元素类型

从数组类型中提取元素类型。

### 实例

// 提取数组的元素类型
// 如果 T 是数组 V[]，则返回 V
type ArrayElement<T> = T extends (infer V)[] ? V : never;

// 测试
type User = { name: string };
type Users = User[];

type E1 = ArrayElement<string[]>; // string
type E2 = ArrayElement<Users>; // User
type E3 = ArrayElement<number>; // never（非数组）

// 使用示例
var users: Users = [{ name: "Alice" }];
var element: ArrayElement<Users> = users[0];

console.log("元素类型: " + element.name);

应用：这种模式常用于从联合类型中提取特定类型的元素。

### 提取函数返回类型

从函数类型中提取返回类型。

### 实例

// 提取函数的返回类型
type ReturnType<T> = T extends (...args: any[]) => infer R ? R : any;

// 测试
function getData() {
return { id: 1, name: "Alice" };
}
function fetchUser(id: number): Promise<User> {
return Promise.resolve({ id, name: "Bob" });
}

type R1 = ReturnType<typeof getData>; // { id: number; name: string }
type R2 = ReturnType<typeof fetchUser>; // Promise<User>

// 使用示例
var result: R1 = { id: 1, name: "Test" };
console.log("返回类型: " + JSON.stringify(result));

内置类型：TypeScript 内置的 `ReturnType<T>` 工具类型就是使用 infer 实现的。

### 提取函数参数类型

从函数类型中提取参数类型。

< h2 class="example">实例

// 提取函数的第一参数类型
type FirstParameter<T> = T extends (first: infer P, ...rest: any[]) => any ? P : never;

// 测试
function createUser(name: string, age: number): User {
return { id: 1, name };
}
function logMessage(msg: string): void {
console.log(msg);
}

type P1 = FirstParameter<typeof createUser>; // string
type P2 = FirstParameter<typeof logMessage>; // string
type P3 = FirstParameter<() => void>; // never

// 使用示例
var nameParam: P1 = "Alice";
console.log("参数类型: " + nameParam);

注意：可以使用 `...rest: any[]` 来匹配任意数量的剩余参数。

### 多个 infer

在一个条件类型中使用多个 infer。

### 实例

// 提取元组的前两个元素类型
type FirstTwo<T> = T extends [infer A, infer B, ...rest: any[]] ? [A, B] : never;

// 测试
type Tuple = [string, number, boolean];
type FirstTwoTypes = FirstTwo<Tuple>; // [string, number]

// 提取对象属性
type ObjectValue<T> = T extends { value: infer V } ? V : never;

type WithValue = { value: string; name: string };
type ExtractedValue = ObjectValue<WithValue>; // string

// 使用示例
var tupleResult: FirstTwoTypes = ["hello", 123];
console.log("元组: " + JSON.stringify(tupleResult));

元组：使用 `...rest: any[]` 可以匹配元组的剩余元素。

### 在递归类型中使用 infer

在递归类型工具中使用 infer 实现复杂类型操作。

### 实例

// 深度只读类型 - 递归应用 infer
type DeepReadonly<T> = T extends Function
? T
: T extends object
? { readonly [P in keyof T]: DeepReadonly<T[P]> }
: T;

// 测试
interface User {
name: string;
address: {
city: string;
zip: string;
};
}

type ReadonlyUser = DeepReadonly<User>;

// 扁平化 Promise 类型
type FlattenPromise<T> = T extends Promise<infer U>
? U extends Promise<any>
? FlattenPromise<U>
: U
: T;

type Nested = Promise<Promise<string>>;
type Flat = FlattenPromise<Nested>; // string

// 使用示例
var user: ReadonlyUser = {
name: "Alice",
address: { city: "Beijing", zip: "100000" }
};
// user.name = "Bob"; // 错误：只读
console.log("深度只读: " + user.name);

递归 infer：在递归类型中使用 infer 可以实现深度类型转换，如深度只读、深度可选等。

### 注意事项

- 只能用在 extends 右侧：infer 只能在条件类型的 extends 子句中使用
- 声明类型变量：使用 `infer V` 声明待推断的类型变量
- 可以多个使用：一个条件类型中可以使用多个 infer
- 推断失败返回 never：如果推断失败，条件类型返回 never

最佳实践：infer 是实现自定义工具类型的核心，掌握它可以实现强大的类型操作。

### 总结

infer 是 TypeScript 类型系统中最强大的特性之一。

- 类型提取：从复杂类型中提取部分类型
- 条件推断：在条件类型中推断类型
- 函数相关：提取函数参数、返回类型
- 递归工具：实现深度类型转换

建议：深入理解 infer 可以让你编写更高级的类型工具，提升类型安全性和开发效率。

---

## TypeScript 索引类型与 keyof 关键字

Source: https://www.runoob.com/typescript/ts-indexed-types.html

## TypeScript 索引类型与 keyof 关键字

索引类型和 keyof 是 TypeScript 中操作对象类型的强大工具。

它们允许我们动态地访问对象的属性，并创建灵活的类型映射。 索引类型与 keyof 工作原理 keyof 操作符 type Keys = keyof User // "id" | "name" | "email" 索引访问类型 type Value = User["name"] // string 映射类型 type Partial<T> {[P in keyof T]?: T[P]} 索引类型应用场景 动态属性访问 类型安全的对象操作 创建通用工具类型

### 为什么需要索引类型

在 JavaScript 中，我们经常需要动态访问对象的属性。

例如，一个函数可能需要获取对象的所有键，或者根据键访问对应的值类型。

索引类型和 keyof 让我们能够在类型系统中表达这种动态性，同时保持类型安全。

概念：索引类型是一类允许动态访问对象属性的类型，keyof 用于获取对象类型的所有键组成的联合类型。

### keyof 操作符

keyof 操作符用于获取对象类型的所有键组成的联合类型。

### 实例

// 定义用户类型
interface User {
id: number; // 用户ID
name: string; // 用户名
email: string; // 邮箱
age?: number; // 年龄（可选）
}

// 使用 keyof 获取所有键的联合类型
// 结果: "id" | "name" | "email" | "age"
type UserKeys = keyof User;

// 测试 keyof
function getProperty<T, K extends keyof T>(obj: T, key: K): T[K] {
return obj[key];
}

const user: User = {
id: 1,
name: "Alice",
email: "alice@example.com"
};

// 获取 name 属性
const userName: string = getProperty(user, "name");
console.log("用户名: " + userName);

运行结果：

```
用户名: Alice

```

说明：keyof 返回的是键名的字面量联合类型，而不是字符串类型。

### 索引访问类型

使用索引访问类型（Index Access Type）可以获取对象属性的类型。

### 实例

// 定义用户类型
interface User {
id: number;
name: string;
email: string;
}

// 使用索引访问类型获取属性类型
type UserId = User["id"]; // number
type UserName = User["name"]; // string

// 可以使用联合类型进行多个键的访问
type UserIdAndName = User["id" | "name"]; // number | string

// 使用 keyof 获取所有属性类型的联合
type AllUserValues = User[keyof User]; // number | string

// 实际使用示例
function getValue<T, K extends keyof T>(obj: T, key: K): T[K] {
return obj[key];
}

const user: User = { id: 1, name: "Bob", email: "bob@test.com" };

// 获取 id 类型
const idValue: number = getValue(user, "id");
console.log("ID: " + idValue);

// 获取 name 类型
const nameValue: string = getValue(user, "name");
console.log("Name: " + nameValue);

索引访问：使用方括号 [] 可以访问对象类型的具体属性类型，非常强大且灵活。

### 映射类型基础

映射类型允许基于现有类型创建新类型，遍历其所有键。

### 实例

// 定义用户类型
interface User {
id: number;
name: string;
email: string;
age: number;
}

// 将所有属性变为可选
type PartialUser = Partial<User>;

// 将所有属性变为只读
type ReadonlyUser = Readonly<User>;

// 自定义映射类型：将所有属性变为可选且字符串化
type Stringify<T> = {
[P in keyof T]: string;
};

type StringifiedUser = Stringify<User>;

// 实际使用示例
const partialUser: PartialUser = {
id: 1,
name: "Alice"
// email 和 age 可选
};

const readonlyUser: ReadonlyUser = {
id: 1,
name: "Bob",
email: "bob@test.com",
age: 25
};
// readonlyUser.name = "Charlie"; // 错误：只读

console.log("部分用户: " + partialUser.name);
console.log("只读用户: " + readonlyUser.name);

映射类型：使用 [P in keyof T] 语法遍历类型的所有键，是创建工具类型的基础。

### 约束键的类型

使用 keyof 和泛型约束来限制函数接受的键。

### 实例

// 定义配置类型
interface Config {
apiUrl: string;
timeout: number;
retry: boolean;
}

// 只能获取存在的键
function getConfigValue<T, K extends keyof T>(
config: T,
key: K
): T[K] {
return config[key];
}

// 定义配置
const config: Config = {
apiUrl: "https://api.example.com",
timeout: 5000,
retry: true
};

// 正确：键存在
const url: string = getConfigValue(config, "apiUrl");
const timeoutVal: number = getConfigValue(config, "timeout");

// 错误：键不存在（TypeScript 会报错）
// const invalid = getConfigValue(config, "unknown");

console.log("API URL: " + url);
console.log("Timeout: " + timeoutVal);

约束：K extends keyof T 确保传入的键一定是对象上存在的键，防止运行时错误。

### 只获取特定类型的属性

从对象类型中提取特定类型的属性键。

### 实例

// 定义混合类型
interface Mixed {
id: number;
name: string;
age: number;
email: string;
active: boolean;
}

// 提取所有字符串类型的键
type StringKeys<T> = {
[K in keyof T]: T[K] extends string ? K : never;
}[keyof T];

// 提取所有数字类型的键
type NumberKeys<T> = {
[K in keyof T]: T[K] extends number ? K : never;
}[keyof T];

// 测试
type StringProps = StringKeys<Mixed>; // "name" | "email"
type NumberProps = NumberKeys<Mixed>; // "id" | "age"

// 实际应用：获取字符串属性的值
function getStringProps<T, K extends StringKeys<T>>(
obj: T,
keys: K[]
): T[K][] {
return keys.map(key => obj[key]);
}

const mixed: Mixed = {
id: 1,
name: "Alice",
age: 25,
email: "alice@test.com",
active: true
};

const strings = getStringProps(mixed, ["name", "email"]);
console.log("字符串属性: " + strings.join(", "));

条件类型：结合条件类型和映射类型，可以实现复杂的类型过滤和提取。

### 遍历数组类型

使用索引类型操作数组和元组。

### 实例

// 定义元组类型
type Tuple = [string, number, boolean];

// 获取元组元素类型
type First = Tuple[0]; // string
type Second = Tuple[1]; // number
type Third = Tuple[2]; // boolean

// 使用 number 获取所有元素类型
type AllElements = Tuple[number]; // string | number | boolean

// 获取数组元素类型
type StringArray = string[];
type ArrayElement = StringArray[number]; // string

// 实际应用：函数重载
function getElement<T extends any[]>(
arr: T,
index: number
): T[indexof T] | undefined {
return index < arr.length ? arr[index] : undefined;
}

const tuple: Tuple = ["hello", 123, true];
const arr: string[] = ["a", "b", "c"];

console.log("元组元素[0]: " + getElement(tuple, 0));
console.log("数组元素[1]: " + getElement(arr, 1));

元组索引：元组可以使用数字索引，每个索引对应特定元素类型，这是数组做不到的。

### 注意事项

- keyof 返回联合类型：keyof 返回键名的字面量联合类型
- 索引访问安全：确保访问的键存在于目标类型中
- 泛型约束：使用 extends keyof 约束泛型参数
- 映射类型需要 in 关键字：映射类型使用 [P in keyof T] 语法

最佳实践：索引类型是创建通用工具类型的基石，熟练掌握可以大幅提升类型操作能力。

### 总结

索引类型和 keyof 是 TypeScript 类型系统的重要组成部分。

- keyof：获取对象类型的所有键
- 索引访问：通过键获取属性类型
- 映射类型：基于现有类型创建新类型
- 类型安全：确保动态属性访问的类型安全

建议：在需要操作对象属性时，优先使用索引类型来保持类型安全。

---

## TypeScript 递归类型

Source: https://www.runoob.com/typescript/ts-recursive-types.html

## TypeScript 递归类型

递归类型是一种引用自身的类型，在处理树结构、嵌套数据时非常有用。

TypeScript 支持递归类型定义，可以表达无限深度的数据结构。 递归类型工作原理 基础节点 interface TreeNode { value: string } 递归引用 递归类型 interface TreeNode { value: string children?: TreeNode[] } 实际结构 root ├── child1 └── child2 递归类型应用场景 树形结构 嵌套对象 深度类型转换

### 为什么需要递归类型

在现实世界中，数据结构往往是嵌套的。

例如，文件系统有文件夹和子文件夹，组织架构有部门和子部门，JSON 数据可以无限嵌套。

递归类型允许我们表达这种无限嵌套的结构，是处理树形数据的基石。

概念：递归类型是指在类型定义中引用自身的类型，可以表达任意深度的嵌套结构。

### 树形结构

递归类型最常见的应用是表示树形结构。

### 实例

// 定义树节点类型，children 引用自身
interface TreeNode {
id: number; // 节点ID
name: string; // 节点名称
children?: TreeNode[]; // 子节点数组，递归引用
}

// 创建树形结构
const fileSystem: TreeNode = {
id: 1,
name: "根目录",
children: [
{
id: 2,
name: "文件夹1",
children: [
{ id: 5, name: "文件A.txt" },
{ id: 6, name: "文件B.txt" }
]
},
{
id: 3,
name: "文件夹2",
children: [
{ id: 7, name: "文件C.txt" }
]
},
{
id: 4,
name: "文件.txt"
}
]
};

// 遍历树的函数
function traverse(node: TreeNode, depth: number = 0): void {
const indent = " ".repeat(depth);
console.log(indent + "&#x1f4c1; " + node.name);

if (node.children) {
for (const child of node.children) {
traverse(child, depth + 1);
}
}
}

traverse(fileSystem);

运行结果：

```
&#x1f4c1; 根目录
&#x1f4c1; 文件夹1
&#x1f4c1; 文件A.txt
&#x1f4c1; 文件B.txt
&#x1f4c1; 文件夹2
&#x1f4c1; 文件C.txt
&#x1f4c1; 文件.txt

```

文件系统：树形结构是递归类型的经典应用，可以表示目录树、组织结构等。

### 嵌套列表

递归类型也可以表示嵌套的列表结构。

### 实例

// 定义嵌套列表类型
type NestedList<T> = T | NestedList<T>[];

// 定义任务类型
interface Task {
id: number;
title: string;
completed: boolean;
}

// 创建嵌套任务列表
const tasks: NestedList<Task> = [
{ id: 1, title: "项目A", completed: false },
[
{ id: 2, title: "子任务1", completed: true },
{ id: 3, title: "子任务2", completed: false }
],
{ id: 4, title: "项目B", completed: false }
];

// 计算嵌套列表深度
function getDepth<T>(list: NestedList<T>, depth: number = 0): number {
if (Array.isArray(list)) {
let maxDepth = depth + 1;
for (const item of list) {
maxDepth = Math.max(maxDepth, getDepth(item, depth + 1));
}
return maxDepth;
}
return depth;
}

console.log("列表深度: " + getDepth(tasks));

联合类型：使用 T | NestedList[] 可以同时处理单个元素和数组。

### 深度只读类型

使用递归类型实现深度只读转换。

### 实例

// 深度只读类型 - 递归应用
type DeepReadonly<T> = T extends Function
? T // 函数保持原样
: T extends object
? { readonly [P in keyof T]: DeepReadonly<T[P]> }
: T;

// 用户类型
interface User {
name: string;
profile: {
email: string;
address: {
city: string;
zip: string;
};
};
friends: User[];
}

// 创建深度只读用户
const user: DeepReadonly<User> = {
name: "Alice",
profile: {
email: "alice@test.com",
address: {
city: "Beijing",
zip: "100000"
}
},
friends: []
};

// 尝试修改会报错
// user.name = "Bob"; // 错误：name 是只读的
// user.profile.address.city = "Shanghai"; // 错误：深层也是只读的

console.log("用户: " + user.name);
console.log("城市: " + user.profile.address.city);

递归转换：DeepReadonly 会递归地将所有嵌套对象属性转换为只读。

### 深度可选类型

使用递归类型实现深度可选转换。

### 实例

// 深度可选类型 - 递归应用
type DeepPartial<T> = T extends object
? { [P in keyof T]?: DeepPartial<T[P]> }
: T;

// 配置类型
interface AppConfig {
database: {
host: string;
port: number;
credentials: {
username: string;
password: string;
};
};
server: {
port: number;
ssl: boolean;
};
}

// 使用深度可选，可以只提供部分配置
const partialConfig: DeepPartial<AppConfig> = {
database: {
host: "localhost"
// port 和 credentials 可选
}
// server 可选
};

console.log("数据库主机: " + partialConfig.database?.host);

可选嵌套：DeepPartial 递归地将所有属性变为可选，便于处理部分配置。

### 链式数据结构

递归类型可以表示链表等链式数据结构。

### 实例

// 链表节点类型
interface ListNode<T> {
value: T; // 当前节点的值
next?: ListNode<T>; // 下一个节点，递归引用
}

// 创建链表
const linkedList: ListNode<number> = {
value: 1,
next: {
value: 2,
next: {
value: 3,
next: {
value: 4,
next: undefined
}
}
}
};

// 遍历链表
function traverseList<T>(node: ListNode<T>): void {
let current: ListNode<T> | undefined = node;
const values: T[] = [];

while (current) {
values.push(current.value);
current = current.next;
}

console.log("链表值: " + values.join(" -> "));
}

traverseList(linkedList);

// 计算链表长度
function getLength<T>(node: ListNode<T>): number {
let length = 0;
let current: ListNode<T> | undefined = node;

while (current) {
length++;
current = current.next;
}

return length;
}

console.log("链表长度: " + getLength(linkedList));

链表：ListNode 通过 next 引用自身，形成链式结构，是递归类型的经典应用。

### 联合类型的递归

使用递归类型处理 JSON 数据的联合类型。

JSON 类型：递归类型可以精确表达 JSON 的所有可能类型。

### 实例

// JSON 值的递归类型定义
type JSONValue = string | number | boolean | null | JSONValue[] | { [key: string]: JSONValue };

// 定义配置对象
const config: JSONValue = {
"name": "my-app",
"version": "1.0.0",
"enabled": true,
"settings": {
"debug": false,
"ports": [3000, 8080],
"metadata": {
"author": "Alice",
"tags": ["web", "typescript"]
}
}
};

// 获取 JSON 值的函数
function getValue(obj: JSONValue, path: string): JSONValue | undefined {
const keys = path.split(".");
let current: JSONValue | undefined = obj;

for (const key of keys) {
if (current && typeof current === "object" && !Array.isArray(current)) {
current = (current as { [key: string]: JSONValue })[key];
} else {
return undefined;
}
}

return current;
}

console.log("版本: " + getValue(config, "version"));
console.log("端口: " + getValue(config, "settings.ports"));
console.log("作者: " + getValue(config, "settings.metadata.author"));

### 注意事项

- 递归基例：确保递归类型有终止条件，避免无限递归
- 条件类型：递归通常与条件类型结合使用
- 深度限制：TypeScript 编译器对递归深度有限制
- 性能考虑：深度递归可能影响类型检查性能

最佳实践：递归类型是处理树形和嵌套数据的利器，熟练掌握可以解决很多复杂的类型问题。

### 总结

递归类型是 TypeScript 类型系统中的高级特性。

- 自引用：类型定义中引用自身
- 树形结构：表达无限嵌套的数据
- 深度转换：实现深度只读、深度可选等工具类型
- 链式结构：表示链表等线性递归结构

建议：在处理嵌套数据时，优先考虑使用递归类型来保证类型安全。

---

## TypeScript 协变与逆变

Source: https://www.runoob.com/typescript/ts-covariance.html

## TypeScript 协变与逆变

协变与逆变是 TypeScript 类型系统中的重要概念，理解它们有助于编写类型安全的代码。

它们描述了泛型类型在父子类型关系中的行为。 协变与逆变概念 协变 (Covariant) 输出类型：安全 Dog → Animal Provider<Dog> → Provider<Animal> 不变 (Invariant) 输入/输出类型：严格 不能相互赋值 Consumer<Dog> ≠ Consumer<Animal> 逆变 (Contravariant) 输入类型：反转 Animal → Dog Consumer<Animal> → Consumer<Dog> TypeScript 默认行为 返回值：协变 参数：逆变 (strictFunctionTypes) 属性：协变

### 为什么需要协变与逆变

当我们使用泛型类或函数时，类型参数的行为并不像你想象的那么简单。

将 Dog 赋值给 Animal 是安全的，但将处理 Animal 的函数赋值给处理 Dog 的函数可能不安全。

协变与逆变规则帮助 TypeScript 捕获这些潜在的类型错误。

概念：协变允许子类型向父类型转换，逆变允许父类型向子类型转换，不变则不允许任何方向转换。

### 协变 (Covariant)

协变是指子类型可以赋值给父类型。对于输出类型（如函数返回值），这是安全的。

### 实例

// 定义动物和狗的类型
class Animal {
name: string = "动物";
}

class Dog extends Animal {
breed: string = "田园犬";
}

// 协变：输出类型可以向更宽泛的类型转换
// 返回 Dog 的函数可以赋值给返回 Animal 的函数
type AnimalGetter = () => Animal;
type DogGetter = () => Dog;

// Dog 是 Animal 的子类，所以 DogGetter 可以赋值给 AnimalGetter
const getDog: DogGetter = () => new Dog();
const getAnimal: AnimalGetter = getDog; // 协变：安全

// 运行
const animal: Animal = getAnimal();
console.log("动物名称: " + animal.name);

运行结果：

```
动物名称: 动物

```

返回值协变：函数返回更具体的子类型是安全的，因为返回的对象必然符合父类型的要求。

### 逆变 (Contravariant)

逆变是指对于输入类型（如函数参数），父类型可以赋值给子类型。

### 实例

// 定义类型
class Animal {
name: string = "动物";
}

class Dog extends Animal {
breed: string = "田园犬";
}

// 逆变：输入类型可以向更具体的类型转换
// 接受 Animal 的函数可以赋值给接受 Dog 的函数
type DogConsumer = (dog: Dog) => void;
type AnimalConsumer = (animal: Animal) => void;

// 如果接受更宽泛类型的函数可以赋值给更具体类型的函数
// 那么当我们传入 Dog 时，函数可能会处理不了（缺少 Dog 特有的属性）
const consumeAnimal: AnimalConsumer = (animal) => {
console.log("处理动物: " + animal.name);
};

const consumeDog: DogConsumer = consumeAnimal; // 逆变：安全

// 运行
const dog = new Dog();
dog.breed = "哈士奇";
consumeDog(dog);

参数逆变：函数参数使用逆变，因为接受更具体类型的函数无法处理更宽泛的类型。

### 启用严格函数类型

TypeScript 默认对函数参数进行逆变检查。启用 strictFunctionTypes 会强制执行此规则。

### 实例

// 定义类型
interface Animal {
readonly name: string;
}

interface Dog extends Animal {
readonly breed: string;
}

// 定义函数类型
type GetName = (animal: Animal) => string;
type GetDogBreed = (dog: Dog) => string;

// 正确的赋值
const getDogBreed: GetDogBreed = (dog) => dog.breed;

// 尝试赋值 - 在 strictFunctionTypes 下会报错
// 因为 AnimalConsumer (参数更宽泛) 不能赋值给 DogConsumer (参数更具体)
// 这是因为参数是逆变的

function printAnimalName(animal: Animal): string {
return animal.name;
}

// 尝试将接受更宽泛类型的函数赋值给更具体类型
// const getSpecific: GetDogBreed = printAnimalName; // 错误！

console.log("犬种: " + getDogBreed({ name: "旺财", breed: "哈士奇" }));

strictFunctionTypes：在 tsconfig.json 中启用此选项可以获得更严格的类型检查。

### 泛型类的协变

泛型类的属性默认是协变的。

### 实例

// 定义类型
class Animal {
name: string = "动物";
}

class Dog extends Animal {
breed: string = "狗";
}

// 泛型容器类
class Cage<T> {
animal: T;

constructor(animal: T) {
this.animal = animal;
}
}

// 协变：可以子类型容器赋值给父类型容器
const dogCage = new Cage(new Dog());
const animalCage: Cage<Animal> = dogCage; // 协变：安全

// animalCage 现在可以安全地当作包含动物的笼子使用
console.log("动物名称: " + animalCage.animal.name);

属性协变：对象的属性是协变的，子类型属性可以赋值给父类型属性。

### 数组的协变

TypeScript 中数组是协变的，但需要注意可变性带来的问题。

### 实例

// 定义类型
class Animal {
name: string = "动物";
}

class Dog extends Animal {
breed: string = "狗";
}

// 数组协变
const dogs: Dog[] = [
{ name: "旺财", breed: "哈士奇" },
{ name: "小白", breed: "萨摩耶" }
];

// Dog[] 可以赋值给 Animal[]
const animals: Animal[] = dogs; // 协变：安全

// 问题：虽然类型上安全，但实际上可以添加其他动物
// animals.push({ name: "猫咪", breed: "猫" }); // 运行时可能出问题！

console.log("动物数量: " + animals.length);

数组可变性：协变赋值后修改数组可能导致运行时错误，需要注意。

### 使用 extends 实现安全赋值

了解协变与逆变后，可以安全地设计泛型接口。

### 实例

// 定义类型
interface Producer<T> {
// 生产方法：返回值是协变的
produce(): T;
}

interface Consumer<T> {
// 消费方法：参数是逆变的
consume(value: T): void;
}

// 具体实现
class DogProducer implements Producer<Dog> {
produce(): Dog {
return { name: "旺财", breed: "哈士奇" };
}
}

class AnimalConsumer implements Consumer<Animal> {
consume(animal: Animal): void {
console.log("消费动物: " + animal.name);
}
}

// Producer<Dog> 可以赋值给 Producer<Animal>（协变）
const animalProducer: Producer<Animal> = new DogProducer();

// Consumer<Animal> 可以赋值给 Consumer<Dog>（逆变）
const dogConsumer: Consumer<Dog> = new AnimalConsumer();

// 测试
const animal = animalProducer.produce();
console.log("生产: " + animal.name);

dogConsumer.consume({ name: "旺财", breed: "哈士奇" });

设计原则：根据方法的用途选择合适的类型方向，提高 API 的类型安全性。

### 注意事项

- 返回值协变：函数返回子类型是安全的
- 参数逆变：函数参数使用父类型是安全的
- 启用严格模式：使用 strictFunctionTypes 获得更严格检查
- 数组协变：注意可变性带来的潜在问题

最佳实践：理解协变与逆变可以帮助设计更类型安全的 API，避免运行时错误。

### 总结

协变与逆变是 TypeScript 类型系统的核心概念。

- 协变：子类型 → 父类型，用于输出类型
- 逆变：父类型 → 子类型，用于输入类型
- 不变：不能相互赋值
- strictFunctionTypes：启用严格函数类型检查

建议：在设计泛型 API 时考虑协变与逆变，编写更安全的类型代码。

---

## TypeScript 路径映射 paths

Source: https://www.runoob.com/typescript/ts-paths.html

## TypeScript 路径映射 paths

paths 是 tsconfig.json 中的配置选项，用于配置模块路径别名。

它可以让导入路径更简洁，同时保持代码的组织结构清晰。 路径映射 paths 工作原理 原始导入路径 import Button from "../../components/Button" 使用别名导入 import Button from "@/components/Button" tsconfig paths 配置 paths 配置项 baseUrl 基准路径 paths 路径别名 rootDirs 根目录

### 为什么需要路径映射

随着项目规模增长，文件目录会越来越深。

使用相对路径导入文件（如 ../../../components/Button）会让代码难以阅读和维护。

路径映射允许我们使用别名（如 @/components/Button）来替代冗长的相对路径。

别名：路径别名让导入路径更简洁，同时便于调整项目结构。

### 基本配置

在 tsconfig.json 中配置 paths 和 baseUrl。

### tsconfig.json

{
"compilerOptions": {
// 基准路径：所有路径别名的根目录
"baseUrl": ".",

// 路径映射配置
"paths": {
// @/ 开头的别名，指向 src 目录
"@/*": ["src/*"],

// @components 开头的别名
"@components/*": ["src/components/*"],

// @utils 开头的别名
"@utils/*": ["src/utils/*"],

// @services 开头的别名
"@services/*": ["src/services/*"],

// @assets 开头的别名
"@assets/*": ["src/assets/*"],

// @types 开头的别名
"@types/*": ["src/types/*"]
}
}
}

baseUrl：设置 baseUrl 后，paths 中的路径将相对于此目录解析。

### 使用路径别名

配置完成后，可以在代码中使用别名导入模块。

### src/components/Button.tsx

// 使用路径别名导入
import { Button } from '@/components/Button';
import { User } from '@types/user';
import { fetchUser } from '@services/userApi';
import { formatDate } from '@utils/date';

// 导入样式
import styles from '@/components/Button.module.css';

// 导入图片
import logo from '@assets/logo.png';

// 使用导入的模块
const handleClick = () => {
console.log('按钮点击');
};

const userButton = new Button({
text: '用户',
onClick: handleClick
});

console.log('组件加载成功');

语法：在 paths 配置中使用 * 作为通配符，导入时用 * 匹配实际路径。

### Webpack 别名配置

如果使用 Webpack，还需要配置 resolve.alias 以支持运行时。

### webpack.config.js

// 导入 webpack 模块
const path = require('path');

// Webpack 配置
module.exports = {
// 其他配置...
resolve: {
// 路径别名配置，需要与 tsconfig.json 保持一致
alias: {
// @ 开头的别名
'@': path.resolve(__dirname, 'src'),

// 其他别名
'@components': path.resolve(__dirname, 'src/components'),
'@utils': path.resolve(__dirname, 'src/utils'),
'@services': path.resolve(__dirname, 'src/services'),
'@types': path.resolve(__dirname, 'src/types'),
'@assets': path.resolve(__dirname, 'src/assets')
},

// 文件扩展名
extensions: ['.ts', '.tsx', '.js', '.jsx', '.json']
}
};

运行时要一致：Webpack 的 alias 配置必须与 TypeScript 的 paths 配置保持一致，否则运行时会出现模块找不到的错误。

### Vite 配置

使用 Vite 时，需要在 tsconfig.json 和 vite.config.ts 中同时配置。

### vite.config.ts

// 导入 vite 模块
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import path from 'path';

// Vite 配置
export default defineConfig({
plugins: [react()],

// 路径别名配置
resolve: {
alias: {
'@': path.resolve(__dirname, './src'),
'@components': path.resolve(__dirname, './src/components'),
'@utils': path.resolve(__dirname, './src/utils'),
'@services': path.resolve(__dirname, './src/services'),
'@types': path.resolve(__dirname, './src/types'),
'@assets': path.resolve(__dirname, './src/assets')
}
}
});

Vite：Vite 使用 Rollup 作为打包引擎，需要在 resolve.alias 中配置别名。

### 多项目路径配置

在 Monorepo 项目中，可以配置跨包的路径别名。

### tsconfig.json (Monorepo)

{
"compilerOptions": {
// 基准路径
"baseUrl": ".",

// 路径映射，支持跨包引用
"paths": {
// 引用当前包的模块
"@/*": ["src/*"],

// 引用其他包的模块
"@my-ui/button": ["packages/ui-button/src/index.ts"],
"@my-ui/modal": ["packages/ui-modal/src/index.ts"],
"@my-utils/date": ["packages/utils-date/src/index.ts"],
"@my-hooks/useFetch": ["packages/hooks-use-fetch/src/index.ts"]
}
}
}

Monorepo：在 Monorepo 项目中，paths 可以引用其他包的源码路径。

### 注意事项

- baseUrl 必需：paths 需要与 baseUrl 配合使用
- 通配符匹配：使用 * 匹配任意路径
- 保持一致：Webpack/Vite 配置必须与 tsconfig 保持一致
- 相对路径：paths 中的路径是相对于 baseUrl 的

最佳实践：使用路径别名可以让代码更简洁，但不要滥用，建议统一命名规范。

### 总结

路径映射是 TypeScript 项目中管理导入路径的重要工具。

- paths：配置路径别名
- baseUrl：设置基准路径
- 通配符：使用 * 匹配任意字符
- 构建工具：Webpack/Vite 需要同步配置

建议：为常用目录设置路径别名，提高代码可读性和可维护性。

---

## TypeScript 项目引用 References

Source: https://www.runoob.com/typescript/ts-references.html

## TypeScript 项目引用 References

项目引用是 TypeScript 提供的组织大型项目的功能，允许将 TypeScript 项目拆分为更小的部分。

它可以实现增量构建、更好的代码组织和更快的编译速度。 项目引用 References 原理 主项目 (app) package.json tsconfig.json references: [...] 组件库 (ui) Button, Modal, Input 工具库 (utils) formatDate, request 类型定义 (types) User, API Response

### 为什么需要项目引用

随着项目增长，单一的 tsconfig.json 会导致编译速度变慢。

项目引用允许将项目拆分为独立的子项目，每个子项目可以独立编译。

这不仅提高了编译速度，还提供了更好的代码组织方式。

概念：项目引用允许一个 TypeScript 项目引用其他项目，实现增量编译和更好的代码组织。

### 创建引用项目

首先创建被引用的子项目。

### packages/utils/tsconfig.json

{
// 继承基础配置
"extends": "../../tsconfig.base.json",

// 编译器选项
"compilerOptions": {
// 输出目录
"outDir": "./dist",

// 声明文件目录
"declarationDir": "./dist/types",

// 生成声明文件
"declaration": true,

// 生成 sourcemap
"sourceMap": true,

// 是否为库
"composite": true
},

// 包含的文件
"include": ["src/**/*"],

// 排除的文件
"exclude": ["node_modules", "dist", "**/*.test.ts"]
}

composite：设置为 true 启用项目引用功能，这是被引用的项目必须设置的选项。

### 主项目配置

在主项目的 tsconfig.json 中配置 references。

### tsconfig.json (主项目)

{
// 继承基础配置
"extends": "./tsconfig.base.json",

// 编译器选项
"compilerOptions": {
// 输出目录
"outDir": "./dist",

// 声明文件目录
"declarationDir": "./dist/types",

// 生成声明文件
"declaration": true,

// 生成 sourcemap
"sourceMap": true
},

// 项目引用配置
"references": [
// 引用 utils 项目
{ "path": "./packages/utils" },

// 引用 ui 项目
{ "path": "./packages/ui" },

// 引用 types 项目
{ "path": "./packages/types" }
],

// 包含的文件
"include": ["src/**/*"],

// 依赖包含的节点模块
// "files": []
}

references：数组中的每个对象指定一个引用的项目路径，path 相对于当前项目。

### 类型引用

在代码中使用被引用项目的类型。

### src/index.ts

// 导入工具函数（来自 utils 包）
import { formatDate, formatCurrency } from '@my-utils/format';

// 导入组件（来自 ui 包）
import { Button, Modal, Input } from '@my-ui/core';

// 导入类型（来自 types 包）
import { User, ApiResponse } from '@my-types/common';

// 定义用户数据
const user: User = {
id: 1,
name: "Alice",
email: "alice@example.com"
};

// 格式化日期
const dateStr = formatDate(new Date(), "YYYY-MM-DD");
console.log("日期: " + dateStr);

// 格式化货币
const price = formatCurrency(999);
console.log("价格: " + price);

// 创建用户界面
const button = new Button({
text: "提交",
variant: "primary"
});

console.log("应用初始化成功");

导入方式：被引用项目的导出可以直接导入，TypeScript 会自动解析类型。

### 增量构建

项目引用支持增量构建，只编译修改的部分。

### 构建命令

# 构建整个项目（包括所有引用）
npm run build

# 只构建主项目（不重新构建依赖）
npm run build -- --build

# 清理并重建
npm run clean
npm run build

# 增量构建（推荐）
# 修改某个包后，只重新编译该包和依赖它的包
npx tsc -b packages/utils
npx tsc -b packages/ui
npx tsc -b .

增量构建：使用 -b (build) 选项时，TypeScript 会自动检测哪些项目需要重新编译。

### 注意事项

- composite 选项：被引用的项目必须设置 composite: true
- 输出目录：每个项目需要独立的输出目录
- 声明文件：被引用项目需要生成声明文件
- 构建顺序：依赖的项目需要先构建

最佳实践：将公共代码拆分为独立包，使用项目引用进行管理，提高编译效率。

### 总结

项目引用是 TypeScript 管理大型项目的核心功能。

- references：配置项目引用关系
- composite：启用项目引用
- 增量构建：只编译修改的部分
- 代码组织：拆分为独立模块

建议：在大型项目中，使用项目引用来组织代码，提高开发效率。

---

## TypeScript Monorepo 配置

Source: https://www.runoob.com/typescript/ts-monorepo.html

## TypeScript Monorepo 配置

Monorepo 是一种将多个项目放在同一个代码仓库中的开发模式。

TypeScript 通过项目引用和工具链支持 Monorepo，可以高效管理多个包。 Monorepo 项目结构 根目录 (root) package.json | tsconfig.json | lerna.json | turbo.json packages/ 目录 utils ui-components hooks app 管理工具：npm workspaces | yarn workspaces | pnpm | lerna | turbo

### 为什么需要 Monorepo

当项目包含多个包（如工具库、组件库、应用程序）时，传统方式需要维护多个代码仓库。

Monorepo 将所有包放在同一个仓库中，共享代码更方便，版本管理更统一。

TypeScript 的项目引用功能让 Monorepo 项目的类型检查和构建更高效。

概念：Monorepo（单一仓库）将多个相关项目放在同一个代码仓库中，便于代码共享和协调开发。

### pnpm Workspace

pnpm 是现代的包管理器，原生支持 Workspace 功能。

### 根目录 package.json

{
"name": "my-monorepo",
"version": "1.0.0",
"private": true,

// 启用 pnpm workspace
"packages": [
"packages/*"
],

// 开发依赖
"devDependencies": {
"typescript": "^5.0.0"
},

// 脚本
"scripts": {
"build": "pnpm -r run build",
"clean": "pnpm -r run clean",
"type-check": "pnpm -r run type-check"
}
}

pnpm：pnpm 的 Workspace 功能可以自动将 packages 目录下的包链接在一起。

### 项目结构

创建 Monorepo 项目的目录结构。

### 目录结构

my-monorepo/
├── packages/
│ ├── utils/ # 工具包
│ │ ├── src/
│ │ │ └── index.ts
│ │ ├── package.json
│ │ └── tsconfig.json
│ │
│ ├── ui-components/ # UI 组件包
│ │ ├── src/
│ │ │ ├── Button.tsx
│ │ │ └── index.ts
│ │ ├── package.json
│ │ └── tsconfig.json
│ │
│ └── app/ # 应用程序
│ ├── src/
│ │ └── index.tsx
│ ├── package.json
│ └── tsconfig.json
│
├── package.json # 根目录配置
├── tsconfig.base.json # 基础配置
└── pnpm-workspace.yaml # pnpm 配置

packages：所有包都放在 packages 目录下，每个包有独立的 package.json 和 tsconfig.json。

### 基础 TypeScript 配置

创建基础配置供所有包使用。

### tsconfig.base.json

{
// 编译器选项
"compilerOptions": {
// 目标版本
"target": "ES2020",

// 模块系统
"module": "ESNext",

// 严格模式
"strict": true,

// 跳过库类型检查
"skipLibCheck": true,

// 启用 ES 模块互操作
"esModuleInterop": true,

// 强制文件名大小写一致
"forceConsistentCasingInFileNames": true,

// 模块解析
"moduleResolution": "bundler",

// 解析 JSON 模块
"resolveJsonModule": true,

// 隔离模块
"isolatedModules": true,

// 不生成输出文件
"noEmit": true
}
}

继承：各包的 tsconfig.json 继承基础配置，只覆盖需要自定义的选项。

### 工具包配置

为工具包配置 TypeScript。

### packages/utils/tsconfig.json

{
// 继承基础配置
"extends": "../../tsconfig.base.json",

// 编译器选项
"compilerOptions": {
// 输出目录
"outDir": "./dist",

// 声明文件目录
"declarationDir": "./dist/types",

// 生成声明文件
"declaration": true,

// 生成入口文件的声明
"declarationMap": true,

// 模块导出方式
"module": "ESNext",

// 启用项目引用
"composite": true
},

// 包含的文件
"include": ["src/**/*"],

// 排除的文件
"exclude": ["node_modules", "dist", "**/*.test.ts"]
}

composite：启用项目引用后，TypeScript 可以增量编译此包。

### 应用程序配置

为主应用程序配置 TypeScript 并引用其他包。

### packages/app/tsconfig.json

{
// 继承基础配置
"extends": "../../tsconfig.base.json",

// 编译器选项
"compilerOptions": {
// 输出目录
"outDir": "./dist",

// JSX 配置
"jsx": "react-jsx",

// 路径别名
"baseUrl": ".",
"paths": {
"@my-utils/*": ["../utils/src/*"],
"@my-ui/*": ["../ui-components/src/*"]
}
},

// 项目引用
"references": [
{ "path": "../utils" },
{ "path": "../ui-components" }
],

// 包含的文件
"include": ["src/**/*"],

// 排除的文件
"exclude": ["node_modules", "dist"]
}

路径别名：可以配置路径别名直接引用同仓库的其他包。

### 包之间的依赖

在 package.json 中声明对同仓库包的依赖。

### packages/app/package.json

{
"name": "@my-org/app",
"version": "1.0.0",
"private": true,

"dependencies": {
// 引用同仓库的包
"@my-org/utils": "workspace:*",
"@my-org/ui-components": "workspace:*",

// 外部依赖
"react": "^18.2.0",
"react-dom": "^18.2.0"
},

"devDependencies": {
// 开发依赖
"@types/react": "^18.2.0",
"typescript": "^5.0.0"
}
}

workspace：使用 workspace:* 指向同仓库的其他包，pnpm 会自动解析。

### 构建脚本

为 Monorepo 创建统一的构建脚本。

### 根目录 package.json

{
"scripts": {
// 构建所有包
"build": "pnpm -r run build",

// 清理所有构建产物
"clean": "pnpm -r run clean",

// 类型检查
"type-check": "pnpm -r run type-check",

// 测试所有包
"test": "pnpm -r run test",

// 增量构建
"build:watch": "pnpm -r --parallel run build:watch",

// 启动应用
"dev": "pnpm --filter @my-org/app run dev"
}
}

pnpm -r：递归执行所有包的同名脚本。

### 注意事项

- 包命名规范：使用 @org-name/package 格式
- 独立版本：每个包可以独立版本管理
- workspace 协议：使用 workspace:* 引用同仓库包
- 构建顺序：被依赖的包需要先构建

最佳实践：Monorepo 适合管理多个相关项目，可以显著提高代码复用和开发效率。

### 总结

Monorepo 是现代前端项目管理的推荐模式。

- pnpm Workspace：原生支持 Monorepo
- 项目引用：实现增量编译
- 路径别名：便捷引用同仓库包
- 统一管理：共享配置和依赖

建议：当项目包含多个相关包时，优先考虑 Monorepo 方案。

---

## TypeScript 设计模式

Source: https://www.runoob.com/typescript/ts-design-patterns.html

## TypeScript 设计模式

设计模式是软件开发中经过验证的解决方案，可以帮助我们编写可维护、可扩展的代码。

TypeScript 的类型系统让许多经典设计模式得以用类型安全的方式实现。 设计模式分类 创建型模式 Factory | Singleton | Builder 结构型模式 Decorator | Adapter | Proxy 行为型模式 Observer | Strategy | Command TypeScript 特有模式 类型安全的依赖注入 | 泛型工厂 | 条件类型选择 TypeScript 优势：编译期类型检查 + IDE 智能提示 + 运行时性能

### 为什么需要设计模式

设计模式是前人总结的代码组织经验，可以解决常见的软件设计问题。

使用设计模式让代码更易理解、更易维护，同时也便于团队协作。

TypeScript 的类型系统让这些模式更加健壮，错误可以在编译期被发现。

概念：设计模式是软件设计中常见问题的可重用解决方案，是代码设计经验的总结。

### 单例模式 (Singleton)

确保一个类只有一个实例，并提供全局访问点。

### 实例

// 单例模式：确保只有一个实例
class Singleton {
// 存储单例实例
private static instance: Singleton;
private static _data: string = "";

// 私有构造函数，防止外部实例化
private constructor() {}

// 获取单例实例的静态方法
public static getInstance(): Singleton {
if (!Singleton.instance) {
Singleton.instance = new Singleton();
}
return Singleton.instance;
}

// 设置数据
public setData(data: string): void {
Singleton._data = data;
}

// 获取数据
public getData(): string {
return Singleton._data;
}
}

// 测试单例模式
const instance1 = Singleton.getInstance();
const instance2 = Singleton.getInstance();

// 验证是同一个实例
console.log("是同一实例: " + (instance1 === instance2));

instance1.setData("Hello Singleton");
console.log("数据: " + instance2.getData());

运行结果：

```
是同一实例: true
数据: Hello Singleton

```

私有构造函数：通过将构造函数设为 private，防止外部使用 new 创建实例。

### 工厂模式 (Factory)

使用泛型工厂创建类型安全的对象实例。

### 实例

// 定义产品接口
interface Product {
name: string;
price: number;
getDescription(): string;
}

// 具体产品：电子产品
class ElectronicProduct implements Product {
constructor(
public name: string,
public price: number,
public warranty: number
) {}

getDescription(): string {
return `${this.name} - ¥${this.price} (保修${this.warranty}年)`;
}
}

// 具体产品：服装
class ClothingProduct implements Product {
constructor(
public name: string,
public price: number,
public size: string
) {}

getDescription(): string {
return `${this.name} - ¥${this.price} (尺码: ${this.size})`;
}
}

// 工厂类
class ProductFactory {
// 泛型工厂方法
static create<T extends Product>(
type: new (...args: any[]) => T,
...args: any[]
): T {
return new type(...args);
}
}

// 使用工厂创建产品
const laptop = ProductFactory.create(ElectronicProduct, "笔记本电脑", 5999, 2);
const shirt = ProductFactory.create(ClothingProduct, "T恤", 199, "L");

console.log(laptop.getDescription());
console.log(shirt.getDescription());

泛型工厂：使用泛型约束确保返回具体的产品类型。

### 装饰器模式 (Decorator)

使用装饰器为对象动态添加功能。

### 实例

// 基础咖啡接口
interface Coffee {
getCost(): number;
getDescription(): string;
}

// 基础咖啡实现
class SimpleCoffee implements Coffee {
getCost(): number {
return 10;
}

getDescription(): string {
return "咖啡";
}
}

// 装饰器基类
abstract class CoffeeDecorator implements Coffee {
constructor(protected coffee: Coffee) {}

getCost(): number {
return this.coffee.getCost();
}

getDescription(): string {
return this.coffee.getDescription();
}
}

// 牛奶装饰器
class MilkDecorator extends CoffeeDecorator {
getCost(): number {
return this.coffee.getCost() + 2;
}

getDescription(): string {
return this.coffee.getDescription() + ", 牛奶";
}
}

// 糖装饰器
class SugarDecorator extends CoffeeDecorator {
getCost(): number {
return this.coffee.getCost() + 1;
}

getDescription(): string {
return this.coffee.getDescription() + ", 糖";
}
}

// 使用装饰器
let coffee: Coffee = new SimpleCoffee();
console.log(coffee.getDescription() + " - ¥" + coffee.getCost());

coffee = new MilkDecorator(coffee);
console.log(coffee.getDescription() + " - ¥" + coffee.getCost());

coffee = new SugarDecorator(coffee);
console.log(coffee.getDescription() + " - ¥" + coffee.getCost());

装饰器：可以在不修改原类的情况下，动态添加新功能，是 TypeScript 实验性功能。

### 观察者模式 (Observer)

定义对象间的一对多依赖关系，当对象状态改变时，所有依赖者都会收到通知。

### 实例

// 观察者接口
interface Observer {
update(message: string): void;
}

// 主题接口
interface Subject {
attach(observer: Observer): void;
detach(observer: Observer): void;
notify(): void;
}

// 具体主题：消息中心
class MessageCenter implements Subject {
private observers: Observer[] = [];
private message: string = "";

// 添加观察者
attach(observer: Observer): void {
this.observers.push(observer);
}

// 移除观察者
detach(observer: Observer): void {
const index = this.observers.indexOf(observer);
if (index > -1) {
this.observers.splice(index, 1);
}
}

// 通知所有观察者
notify(): void {
for (const observer of this.observers) {
observer.update(this.message);
}
}

// 发布消息
publish(message: string): void {
this.message = message;
console.log("发布消息: " + message);
this.notify();
}
}

// 具体观察者：用户
class UserObserver implements Observer {
constructor(public name: string) {}

update(message: string): void {
console.log(`[${this.name}] 收到消息: ${message}`);
}
}

// 使用观察者模式
const center = new MessageCenter();

const user1 = new UserObserver("用户A");
const user2 = new UserObserver("用户B");

center.attach(user1);
center.attach(user2);

center.publish("新功能上线了！");

解耦：观察者模式实现了主题和观察者之间的松耦合。

### 策略模式 (Strategy)

定义一系列算法，把它们一个个封装起来，使它们可以相互替换。

### 实例

// 支付策略接口
interface PaymentStrategy {
pay(amount: number): void;
}

// 微信支付策略
class WechatPayStrategy implements PaymentStrategy {
pay(amount: number): void {
console.log(`使用微信支付 ¥${amount}`);
}
}

// 支付宝策略
class AlipayStrategy implements PaymentStrategy {
pay(amount: number): void {
console.log(`使用支付宝支付 ¥${amount}`);
}
}

// 银行卡策略
class CardPayStrategy implements PaymentStrategy {
pay(amount: number): void {
console.log(`使用银行卡支付 ¥${amount}`);
}
}

// 支付上下文
class PaymentContext {
private strategy: PaymentStrategy;

constructor(strategy: PaymentStrategy) {
this.strategy = strategy;
}

// 设置支付策略
setStrategy(strategy: PaymentStrategy): void {
this.strategy = strategy;
}

// 执行支付
pay(amount: number): void {
this.strategy.pay(amount);
}
}

// 使用策略模式
const context = new PaymentContext(new WechatPayStrategy());
context.pay(100);

context.setStrategy(new AlipayStrategy());
context.pay(200);

context.setStrategy(new CardPayStrategy());
context.pay(300);

算法切换：策略模式可以在运行时切换算法，提供了很大的灵活性。

### 依赖注入 (Dependency Injection)

通过构造函数注入依赖，是 TypeScript 中最常用的模式之一。

### 实例

// 定义服务接口
interface Logger {
log(message: string): void;
}

interface Storage {
save(key: string, data: any): void;
}

// 具体服务实现
class ConsoleLogger implements Logger {
log(message: string): void {
console.log("[日志]: " + message);
}
}

class LocalStorage implements Storage {
save(key: string, data: any): void {
console.log(`保存 ${key}: ${JSON.stringify(data)}`);
localStorage.setItem(key, JSON.stringify(data));
}
}

// 使用依赖注入的服务
class UserService {
constructor(
private logger: Logger,
private storage: Storage
) {}

createUser(name: string): void {
const user = { name, createdAt: new Date() };
this.logger.log("创建用户: " + name);
this.storage.save("user", user);
}
}

// 注入依赖
const logger = new ConsoleLogger();
const storage = new LocalStorage();
const userService = new UserService(logger, storage);

userService.createUser("Alice");

依赖倒置：依赖注入实现了高层模块不依赖低层模块，而是依赖抽象接口。

### 建造者模式 (Builder)

将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。

### 实例

// 建造者接口
interface Builder<T> {
build(): T;
}

// 复杂对象：用户配置
interface UserConfig {
name: string;
email: string;
age?: number;
role?: string;
theme?: string;
}

// 用户配置建造者
class UserConfigBuilder implements Builder<UserConfig> {
private config: Partial<UserConfig> = {};

setName(name: string): this {
this.config.name = name;
return this;
}

setEmail(email: string): this {
this.config.email = email;
return this;
}

setAge(age: number): this {
this.config.age = age;
return this;
}

setRole(role: string): this {
this.config.role = role;
return this;
}

setTheme(theme: string): this {
this.config.theme = theme;
return this;
}

build(): UserConfig {
if (!this.config.name || !this.config.email) {
throw new Error("Name and email are required");
}
return this.config as UserConfig;
}
}

// 使用建造者
const builder = new UserConfigBuilder();
const config = builder
.setName("Alice")
.setEmail("alice@example.com")
.setAge(25)
.setRole("admin")
.setTheme("dark")
.build();

console.log("用户配置:", JSON.stringify(config, null, 2));

链式调用：建造者模式支持链式调用，使代码更简洁易读。

### 注意事项

- 类型安全：利用 TypeScript 的类型系统确保模式实现的安全
- 不要过度：只在真正需要时才使用设计模式
- 保持简单：优先使用简单的解决方案
- 团队共识：在团队内部统一设计模式的使用

最佳实践：设计模式是工具，不是教条。选择适合项目实际情况的模式。

### 总结

设计模式是代码组织的重要经验。

- 单例模式：确保类只有一个实例
- 工厂模式：封装对象创建过程
- 装饰器模式：动态添加功能
- 观察者模式：一对多依赖关系
- 策略模式：算法可替换
- 依赖注入：解耦依赖关系

建议：理解并在实践中应用设计模式，可以写出更优雅、更易维护的代码。

---

## TypeScript 性能优化

Source: https://www.runoob.com/typescript/ts-performance.html

## TypeScript 性能优化

TypeScript 项目的性能优化涉及编译速度、运行时性能和代码体积等多个方面。

本教程介绍常见的性能优化技巧，帮助构建更高效的 TypeScript 应用。 TypeScript 性能优化维度 编译速度 增量构建 | skipLibCheck 运行时性能 类型推断 | 避免 any 包体积 Tree Shaking | 懒加载 关键优化策略 skipLibCheck 跳过库类型检查 incremental 增量编译 Project References 项目引用

### 为什么需要性能优化

TypeScript 虽然提供了强大的类型系统，但使用不当会影响编译速度和运行性能。

大型项目的编译可能需要数分钟，严重影响开发体验。

本教程介绍的配置和技巧可以显著提升 TypeScript 项目的性能。

优化目标：更快的编译速度、更小的打包体积、更高的运行时性能。

### 编译配置优化

通过优化 tsconfig.json 配置来提升编译速度。

### tsconfig.json 优化配置

{
"compilerOptions": {
// 启用增量编译，保存上次编译信息
"incremental": true,

// 跳过 node_modules 的类型检查
// 大幅提升编译速度
"skipLibCheck": true,

// 跳过声明文件的生成（仅在最终构建时生成）
"noEmit": true,

// 启用快速增量构建
"assumeChangesOnlyAffectDirectDependencies": true,

// 并行执行
"parallel": true,

// 启用缓存
"tsBuildInfoFile": ".tsbuildinfo",

// 不需要解析的文件
"exclude": [
"node_modules",
"dist",
"build",
"**/*.test.ts"
]
}
}

skipLibCheck：这是最重要的优化选项，可以将编译时间减少 50% 以上。

### 项目引用优化

使用项目引用将大型项目拆分为小模块，实现增量编译。

### packages/utils/tsconfig.json

{
// 继承基础配置
"extends": "../../tsconfig.base.json",

"compilerOptions": {
"composite": true,
"outDir": "./dist",
"declaration": true,
"declarationMap": true
},

"include": ["src/**/*"],
"exclude": ["node_modules", "dist"]
}

composite：启用后，TypeScript 会生成 .tsbuildinfo 文件来加速后续编译。

### 类型推断优化

充分利用 TypeScript 的类型推断，避免过度标注。

< h2 class="example">实例

// 不好的写法：过度标注类型
const name: string = "Alice";
const age: number = 25;
const isActive: boolean = true;

// 好的写法：利用类型推断
const name = "Alice";
const age = 25;
const isActive = true;

// 函数返回值类型可以省略（TypeScript 会自动推断）
function add(a: number, b: number) {
return a + b;
}

// 复杂对象使用类型推断
const user = {
id: 1,
name: "Bob",
email: "bob@example.com"
};
// TypeScript 会推断出:
// { id: number; name: string; email: string }

// 只有在类型推断不准确时才需要显式标注
const elements: HTMLElement[] = [];

减少标注：TypeScript 的类型推断已经很智能，大部分情况下不需要显式标注类型。

### 避免使用 any

使用 any 会失去类型检查的优势，影响运行时性能。

### 实例

// 不好的写法：使用 any
function processData(data: any): any {
return data.value;
}

// 好的写法：使用 unknown 或具体类型
function processData<T extends { value: string }>(data: T): string {
return data.value;
}

// 如果真的不知道类型，使用 unknown
function parseJSON(json: string): unknown {
return JSON.parse(json);
}

// 使用时进行类型检查
const data = parseJSON('{"key": "value"}');
if (typeof data === "object" && data !== null) {
const obj = data as { key: string };
console.log(obj.key);
}

// 更好的做法：使用泛型
function identity<T>(value: T): T {
return value;
}

const result = identity("hello");
console.log("结果: " + result);

unknown vs any：unknown 是类型安全的 any，使用前需要进行类型检查。

### 接口 vs 类型别名

根据场景选择合适的类型定义方式。

### 实例

// 接口：适合定义对象类型，支持声明合并
interface User {
id: number;
name: string;
}

// 扩展接口
interface User {
email: string;
}

// 类型别名：适合联合类型、元组、函数类型
type ID = string | number;
type Status = "pending" | "success" | "error";
type Callback = (data: string) => void;

// 工具类型通常使用 type
type PartialUser = Partial<User>;
type ReadonlyUser = Readonly<User>;

// 性能考虑：接口的编译速度通常比类型别名快
// 对于简单的对象类型，可以使用 interface
interface Point {
x: number;
y: number;
}

// 对于联合类型，使用 type
type Shape = Circle | Square | Triangle;

选择建议：对象类型使用接口，联合类型使用 type，功能类型使用 type。

### 构建工具优化

配置构建工具以获得最佳性能。

### vite.config.ts

// 导入 vite
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// Vite 配置优化
export default defineConfig({
plugins: [react()],

// 构建优化
build: {
// 启用代码拆分
rollupOptions: {
output: {
// 手动分包
manualChunks: {
'vendor': ['react', 'react-dom'],
'utils': ['lodash', 'axios']
}
}
},

// 压缩代码
minify: 'terser',

// 生成 sourcemap
sourcemap: false,

// 块大小限制
chunkSizeWarningLimit: 500
},

// 开发服务器优化
server: {
// 启用热更新
hmr: {
overlay: true
}
},

// 优化依赖解析
optimizeDeps: {
include: ['react', 'react-dom'],
exclude: ['some-large-library']
}
});

代码分割：使用 manualChunks 将大型库分离，减少主包体积。

### Tree Shaking

配置模块以支持 Tree Shaking，消除未使用的代码。

### src/utils/index.ts

// 使用 ES 模块导出，支持 Tree Shaking
export function add(a: number, b: number): number {
return a + b;
}

export function subtract(a: number, b: number): number {
return a - b;
}

export function multiply(a: number, b: number): number {
return a * b;
}

export function divide(a: number, b: number): number {
if (b === 0) {
throw new Error("Cannot divide by zero");
}
return a / b;
}

// 具名导出比默认导出更有利于 Tree Shaking
// 错误：export default 会阻止 Tree Shaking
// export default { add, subtract, multiply, divide };

// 正确：使用具名导出
console.log("工具模块加载");

具名导出：使用具名导出而不是默认导出，可以让构建工具更好地进行 Tree Shaking。

### 注意事项

- skipLibCheck：生产环境必须开启
- 增量编译：开发环境建议开启
- 避免 any：使用 unknown 代替
- Tree Shaking：使用 ES 模块和具名导出

最佳实践：在开发初期就设置好优化配置，避免后期重构。

### 总结

TypeScript 性能优化涉及多个方面。

- 编译优化：skipLibCheck、incremental、project references
- 类型优化：利用推断、避免 any、选择合适的类型定义
- 构建优化：代码分割、Tree Shaking、依赖优化
- 运行时优化：类型安全、泛型约束

建议：定期检查编译时间和包体积，持续优化项目性能。

---

## TypeScript 综合项目实战

Source: https://www.runoob.com/typescript/ts-comprehensive-project.html

## TypeScript 综合项目实战

本教程通过一个完整的项目案例，综合运用 TypeScript 的各种特性。

从项目搭建到实际开发，完整展示 TypeScript 在实际项目中的应用。 项目实战：任务管理系统 前端 (React + TS) Vite + React 18 组件 + Hooks 类型安全的 UI 后端 (Node.js) Express + TS RESTful API 类型定义 数据存储 In-Memory (演示) 项目功能特性 任务 CRUD 操作 类型安全的 API 状态管理

### 为什么需要综合项目实战

学习 TypeScript 语法后，需要通过实际项目来巩固知识。

本教程展示一个完整的任务管理系统，涵盖前端、后端和类型定义。

通过这个项目，可以掌握 TypeScript 在实际开发中的最佳实践。

项目目标：创建任务管理系统，包含任务创建、查询、更新、删除功能。

### 项目结构

使用 Monorepo 风格组织项目结构。

### 目录结构

task-manager/
├── src/
│ ├── types/ # 类型定义
│ │ ├── task.ts # 任务类型
│ │ ├── api.ts # API 类型
│ │ └── index.ts # 类型导出
│ │
│ ├── services/ # 服务层
│ │ ├── taskService.ts # 任务服务
│ │ └── index.ts
│ │
│ ├── components/ # React 组件
│ │ ├── TaskList.tsx # 任务列表
│ │ ├── TaskItem.tsx # 任务项
│ │ ├── TaskForm.tsx # 任务表单
│ │ └── index.ts
│ │
│ ├── hooks/ # 自定义 Hooks
│ │ ├── useTasks.ts # 任务状态管理
│ │ └── index.ts
│ │
│ ├── App.tsx # 主应用
│ ├── App.css # 样式
│ └── main.tsx # 入口文件
│
├── index.html
├── package.json
├── tsconfig.json
└── vite.config.ts

目录划分：按功能划分目录，类型定义、服务层、组件分离。

### 类型定义

首先定义项目的核心类型。

### src/types/task.ts

// 任务状态枚举
export type TaskStatus = "pending" | "in-progress" | "completed";

// 任务优先级枚举
export type TaskPriority = "low" | "medium" | "high";

// 任务接口定义
export interface Task {
id: string; // 任务ID
title: string; // 任务标题
description?: string; // 任务描述（可选）
status: TaskStatus; // 任务状态
priority: TaskPriority; // 任务优先级
createdAt: string; // 创建时间
updatedAt: string; // 更新时间
dueDate?: string; // 截止日期（可选）
tags?: string[]; // 标签（可选）
}

// 创建任务的输入类型
export interface CreateTaskInput {
title: string;
description?: string;
priority: TaskPriority;
dueDate?: string;
tags?: string[];
}

// 更新任务的输入类型
export interface UpdateTaskInput {
title?: string;
description?: string;
status?: TaskStatus;
priority?: TaskPriority;
dueDate?: string;
tags?: string[];
}

// 任务过滤选项
export interface TaskFilter {
status?: TaskStatus;
priority?: TaskPriority;
search?: string;
}

类型分层：将输入类型、输出类型、过滤类型分开定义，便于维护。

### API 类型定义

定义 API 相关的类型。

< h2 class="example">src/types/api.ts

// 通用 API 响应类型
export interface ApiResponse<T> {
success: boolean;
data?: T;
error?: string;
message?: string;
}

// 分页元数据
export interface PaginationMeta {
total: number;
page: number;
pageSize: number;
totalPages: number;
}

// 分页响应类型
export interface PaginatedResponse<T> {
items: T[];
meta: PaginationMeta;
}

// 请求错误类型
export interface ApiError {
code: string;
message: string;
details?: Record<string, string>;
}

// HTTP 方法类型
export type HttpMethod = "GET" | "POST" | "PUT" | "PATCH" | "DELETE";

// 任务相关的 API 端点
export interface TaskEndpoints {
getAll: "/api/tasks";
getById: "/api/tasks/:id";
create: "/api/tasks";
update: "/api/tasks/:id";
delete: "/api/tasks/:id";
}

API 类型：统一的响应格式和错误处理类型，便于前后端对接。

### 任务服务层

实现任务管理的业务逻辑。

### src/services/taskService.ts

// 导入类型定义
import {
Task,
CreateTaskInput,
UpdateTaskInput,
TaskFilter,
TaskStatus,
TaskPriority
} from "../types/task";

// 生成唯一ID
function generateId(): string {
return Date.now().toString(36) + Math.random().toString(36).substr(2);
}

// 模拟数据库（内存存储）
let tasks: Task[] = [
{
id: "1",
title: "学习 TypeScript",
description: "掌握 TypeScript 基础和高级特性",
status: "completed",
priority: "high",
createdAt: new Date().toISOString(),
updatedAt: new Date().toISOString(),
tags: ["学习", "TypeScript"]
},
{
id: "2",
title: "开发任务管理系统",
description: "使用 React + TypeScript 开发",
status: "in-progress",
priority: "high",
createdAt: new Date().toISOString(),
updatedAt: new Date().toISOString(),
tags: ["项目", "实战"]
}
];

// 任务服务类
class TaskService {
// 获取所有任务
getAll(filter?: TaskFilter): Task[] {
let result = [...tasks];

if (filter) {
if (filter.status) {
result = result.filter(t => t.status === filter.status);
}
if (filter.priority) {
result = result.filter(t => t.priority === filter.priority);
}
if (filter.search) {
const search = filter.search.toLowerCase();
result = result.filter(t =>
t.title.toLowerCase().includes(search) ||
t.description?.toLowerCase().includes(search)
);
}
}

return result;
}

// 根据ID获取任务
getById(id: string): Task | undefined {
return tasks.find(t => t.id === id);
}

// 创建任务
create(input: CreateTaskInput): Task {
const now = new Date().toISOString();
const task: Task = {
id: generateId(),
title: input.title,
description: input.description,
status: "pending",
priority: input.priority,
createdAt: now,
updatedAt: now,
dueDate: input.dueDate,
tags: input.tags
};

tasks.push(task);
return task;
}

// 更新任务
update(id: string, input: UpdateTaskInput): Task | null {
const index = tasks.findIndex(t => t.id === id);
if (index === -1) return null;

const task = tasks[index];
const updated: Task = {
...task,
...input,
updatedAt: new Date().toISOString()
};

tasks[index] = updated;
return updated;
}

// 删除任务
delete(id: string): boolean {
const index = tasks.findIndex(t => t.id === id);
if (index === -1) return false;

tasks.splice(index, 1);
return true;
}

// 更新任务状态
updateStatus(id: string, status: TaskStatus): Task | null {
return this.update(id, { status });
}
}

// 导出服务实例
export const taskService = new TaskService();

服务层：业务逻辑集中在服务层，便于测试和维护。

### 自定义 Hook

使用 Hook 管理任务状态。

< h2 class="example">src/hooks/useTasks.ts

// 导入 React Hooks 和类型
import { useState, useEffect, useCallback } from "react";
import {
Task,
CreateTaskInput,
UpdateTaskInput,
TaskFilter,
TaskStatus,
TaskPriority
} from "../types/task";
import { taskService } from "../services/taskService";

// Hook 返回的状态类型
interface UseTasksReturn {
tasks: Task[];
loading: boolean;
error: string | null;
filter: TaskFilter;
// 操作方法
createTask: (input: CreateTaskInput) => Promise<void>;
updateTask: (id: string, input: UpdateTaskInput) => Promise<void>;
deleteTask: (id: string) => Promise<void>;
updateStatus: (id: string, status: TaskStatus) => Promise<void>;
setFilter: (filter: TaskFilter) => void;
refresh: () => void;
}

// 初始化默认过滤器
const defaultFilter: TaskFilter = {};

export function useTasks(): UseTasksReturn {
// 任务列表状态
const [tasks, setTasks] = useState<Task[]>([]);
// 加载状态
const [loading, setLoading] = useState(true);
// 错误状态
const [error, setError] = useState<string | null>(null);
// 过滤条件
const [filter, setFilter] = useState<TaskFilter>(defaultFilter);

// 加载任务列表
const loadTasks = useCallback(() => {
setLoading(true);
setError(null);

try {
const data = taskService.getAll(filter);
setTasks(data);
} catch (err) {
setError(err instanceof Error ? err.message : "加载失败");
} finally {
setLoading(false);
}
}, [filter]);

// 初始加载和过滤器变化时重新加载
useEffect(() => {
loadTasks();
}, [loadTasks]);

// 创建任务
const createTask = useCallback(async (input: CreateTaskInput) => {
try {
taskService.create(input);
loadTasks();
} catch (err) {
setError(err instanceof Error ? err.message : "创建失败");
}
}, [loadTasks]);

// 更新任务
const updateTask = useCallback(async (id: string, input: UpdateTaskInput) => {
try {
taskService.update(id, input);
loadTasks();
} catch (err) {
setError(err instanceof Error ? err.message : "更新失败");
}
}, [loadTasks]);

// 删除任务
const deleteTask = useCallback(async (id: string) => {
try {
taskService.delete(id);
loadTasks();
} catch (err) {
setError(err instanceof Error ? err.message : "删除失败");
}
}, [loadTasks]);

// 更新任务状态
const updateStatus = useCallback(async (id: string, status: TaskStatus) => {
try {
taskService.updateStatus(id, status);
loadTasks();
} catch (err) {
setError(err instanceof Error ? err.message : "状态更新失败");
}
}, [loadTasks]);

// 刷新任务列表
const refresh = useCallback(() => {
loadTasks();
}, [loadTasks]);

return {
tasks,
loading,
error,
filter,
createTask,
updateTask,
deleteTask,
updateStatus,
setFilter,
refresh
};
}

自定义 Hook：将状态管理和业务逻辑封装在 Hook 中，组件更简洁。

### React 组件

使用 Hook 创建任务列表组件。

### src/components/TaskList.tsx

// 导入 React 和自定义 Hook
import React from "react";
import { useTasks } from "../hooks/useTasks";
import { Task, TaskStatus, TaskPriority } from "../types/task";

// 任务项组件 Props
interface TaskItemProps {
task: Task;
onStatusChange: (id: string, status: TaskStatus) => void;
onDelete: (id: string) => void;
}

// 任务项组件
const TaskItem: React.FC<TaskItemProps> = ({
task,
onStatusChange,
onDelete
}) => {
// 状态样式映射
const statusStyles: Record<TaskStatus, string> = {
"pending": "status-pending",
"in-progress": "status-progress",
"completed": "status-completed"
};

// 优先级样式映射
const priorityLabels: Record<TaskPriority, string> = {
"low": "低",
"medium": "中",
"high": "高"
};

return (
<div className={`task-item ${statusStyles[task.status]}`}>
<div className="task-content">
<h3 className="task-title">{task.title}</h3>
{task.description && (
<p className="task-description">{task.description}</p>
)}
<div className="task-meta">
<span className={`priority priority-${task.priority}`}>
{priorityLabels[task.priority]}
</span>
<span className="task-date">
{new Date(task.createdAt).toLocaleDateString()}
</span>
</div>
</div>
<div className="task-actions">
<select
value={task.status}
onChange={(e) => onStatusChange(task.id, e.target.value as TaskStatus)}
className="status-select"
>
<option value="pending">待处理</option>
<option value="in-progress">进行中</option>
<option value="completed">已完成</option>
</select>
<button
onClick={() => onDelete(task.id)}
className="delete-btn"
>
删除
</button>
</div>
</div>
);
};

// 任务列表组件
export const TaskList: React.FC = () => {
// 使用自定义 Hook 获取任务状态
const {
tasks,
loading,
error,
updateStatus,
deleteTask
} = useTasks();

// 渲染加载状态
if (loading) {
return <div className="loading">加载中...</div>;
}

// 渲染错误状态
if (error) {
return <div className="error">错误: {error}</div>;
}

// 渲染空状态
if (tasks.length === 0) {
return (
<div className="empty">
<p>暂无任务</p>
<p>点击上方按钮创建新任务</p>
</div>
);
}

// 渲染任务列表
return (
<div className="task-list">
{tasks.map(task => (
<TaskItem
key={task.id}
task={task}
onStatusChange={updateStatus}
onDelete={deleteTask}
/>
))}
</div>
);
};

export default TaskList;

组件拆分：将 TaskItem 拆分为独立组件，代码更清晰。

### 主应用组件

整合所有组件，组成完整的应用。

### src/App.tsx

// 导入 React 和类型
import React, { useState } from "react";
import { TaskList } from "./components/TaskList";
import { useTasks } from "./hooks/useTasks";
import { CreateTaskInput, TaskPriority } from "./types/task";
import "./App.css";

// 任务表单组件 Props
interface TaskFormProps {
onSubmit: (input: CreateTaskInput) => void;
}

// 任务表单组件
const TaskForm: React.FC<TaskFormProps> = ({ onSubmit }) => {
// 表单状态
const [title, setTitle] = useState("");
const [description, setDescription] = useState("");
const [priority, setPriority] = useState<TaskPriority>("medium");

// 提交处理
const handleSubmit = (e: React.FormEvent) => {
e.preventDefault();

if (!title.trim()) {
alert("请输入任务标题");
return;
}

onSubmit({
title: title.trim(),
description: description.trim() || undefined,
priority
});

// 重置表单
setTitle("");
setDescription("");
setPriority("medium");
};

return (
<form onSubmit={handleSubmit} className="task-form">
<input
type="text"
value={title}
onChange={(e) => setTitle(e.target.value)}
placeholder="输入任务标题"
className="form-input"
/>
<input
type="text"
value={description}
onChange={(e) => setDescription(e.target.value)}
placeholder="输入任务描述（可选）"
className="form-input"
/>
<select
value={priority}
onChange={(e) => setPriority(e.target.value as TaskPriority)}
className="form-select"
>
<option value="low">低优先级</option>
<option value="medium">中优先级</option>
<option value="high">高优先级</option>
</select>
<button type="submit" className="submit-btn">
添加任务
</button>
</form>
);
};

// 主应用组件
const App: React.FC = () => {
// 使用自定义 Hook
const { createTask, error } = useTasks();

// 渲染错误提示
const renderError = () => {
if (!error) return null;
return <div className="app-error">{error}</div>;
};

return (
<div className="app">
<header className="app-header">
<h1>TypeScript 任务管理系统</h1>
<p>综合实战项目</p>
</header>

{renderError()}

<main className="app-main">
<TaskForm onSubmit={createTask} />
<TaskList />
</main>

<footer className="app-footer">
<p>Powered by TypeScript + React</p>
</footer>
</div>
);
};

export default App;

组件组合：通过组合简单的组件构建复杂的 UI。

### 项目配置

项目的 TypeScript 和构建配置。

### tsconfig.json

{
"compilerOptions": {
"target": "ES2020",
"useDefineForClassFields": true,
"lib": ["ES2020", "DOM", "DOM.Iterable"],
"module": "ESNext",
"skipLibCheck": true,

"moduleResolution": "bundler",
"allowImportingTsExtensions": true,
"resolveJsonModule": true,
"isolatedModules": true,
"noEmit": true,
"jsx": "react-jsx",

"strict": true,
"noUnusedLocals": true,
"noUnusedParameters": true,
"noFallthroughCasesInSwitch": true
},
"include": ["src"],
"references": [{ "path": "./tsconfig.node.json" }]
}

严格模式：启用 strict 获得最完整的类型检查。

### 注意事项

- 类型优先：先定义类型，再编写实现代码
- 接口 vs 类型：对象类型用接口，联合类型用类型别名
- 分层架构：类型、服务、组件分层组织
- Hook 封装：业务逻辑封装在 Hook 中

最佳实践：类型定义是 TypeScript 项目的基础，要认真设计。

### 总结

通过这个综合项目，我们实践了 TypeScript 的核心概念。

- 类型定义：Task、CreateTaskInput、TaskFilter 等接口
- 服务层：TaskService 封装业务逻辑
- 自定义 Hook：useTasks 管理状态
- React 组件：类型安全的组件开发
- 项目配置：严格的 TypeScript 配置

建议：多参与实际项目，在实践中加深对 TypeScript 的理解。

---

## TypeScript + Node.js 实战

Source: https://www.runoob.com/typescript/ts-nodejs.html

## TypeScript + Node.js 实战

TypeScript 在 Node.js 后端开发中广泛应用，本教程介绍 Node.js 项目的 TypeScript 配置和使用。

使用 TypeScript 可以让 Node.js 代码更安全、更易维护，特别适合中大型后端项目。

Node.js 教程查看：https://www.runoob.com/nodejs/nodejs-tutorial.html Node.js + TypeScript 项目结构 项目配置 package.json tsconfig.json ts-node 源代码 (src/) types/ - 类型定义 services/ - 业务逻辑 routes/ - 路由处理 编译输出 (dist/) .js - JavaScript .d.ts - 类型声明 .map - 源码映射 Node.js 运行时 执行 JS 开发流程 npm install 安装依赖 tsc 编译 TypeScript node 运行 JavaScript API 请求测试

### 为什么需要在 Node.js 中使用 TypeScript

Node.js 项目通常涉及复杂的业务逻辑和数据处理，代码行数会快速增长。

使用 TypeScript 可以：提供编译期类型检查，减少运行时错误；利用 IDE 智能提示，提高开发效率；代码更易读和维护。

许多大型 Node.js 项目（如NestJS）都推荐或默认使用 TypeScript。

优势：TypeScript 的静态类型检查可以在开发时发现潜在问题，比运行时调试更高效。

### 项目初始化

首先初始化 Node.js 项目并安装 TypeScript 相关依赖。

### 初始化项目

# 初始化 npm
npm init -y

# 安装 TypeScript 和 Node.js 类型定义
# -D 表示安装为开发依赖
npm install -D typescript @types/node ts-node nodemon

# 初始化 tsconfig.json
npx tsc --init

说明：`@types/node` 提供了 Node.js API 的类型定义，`ts-node` 可以直接运行 TypeScript 文件，`nodemon` 监听文件变化自动重启。

### tsconfig.json 配置

配置 TypeScript 编译器选项，针对 Node.js 项目进行优化。

### tsconfig.json

{
"compilerOptions": {
// 编译目标：ES2020
"target": "ES2020",

// 模块系统：CommonJS（Node.js 原生支持）
"module": "commonjs",

// 启用的库
"lib": ["ES2020"],

// 输出目录
"outDir": "./dist",

// 源码根目录
"rootDir": "./src",

// 严格模式（始终开启）
"strict": true,

// ES 模块互操作
"esModuleInterop": true,

// 跳过库检查
"skipLibCheck": true,

// 强制文件名大小写一致
"forceConsistentCasingInFileNames": true,

// 模块解析策略
"moduleResolution": "node",

// 生成声明文件
"declaration": true
},
// 要编译的文件
"include": ["src/**/*"],
// 排除的文件
"exclude": ["node_modules", "dist"]
}

推荐配置：Node.js 项目推荐使用 `module: "commonjs"`，这是 Node.js 原生支持的模块系统。

### 定义类型

在 types 目录中定义项目的类型接口。

### src/types/index.ts

// 用户类型定义
export interface User {
id: number; // 用户 ID
name: string; // 用户名
email: string; // 邮箱
createdAt: Date; // 创建时间
}

// 创建用户的数据传输对象
export interface CreateUserDTO {
name: string; // 用户名（必填）
email: string; // 邮箱（必填）
password: string; // 密码（必填）
}

// API 响应类型（泛型）
export interface ApiResponse<T> {
success: boolean; // 是否成功
data?: T; // 成功时的数据
error?: string; // 失败时的错误信息
}

类型定义：将类型定义集中放在 types 目录，便于维护和复用。

### 实现服务

在 services 目录中实现业务逻辑，使用定义好的类型。

### src/services/userService.ts

// 导入类型定义
import { User, CreateUserDTO, ApiResponse } from "../types";

// 用户服务类
class UserService {
// 用户列表（内存存储）
private users: User[] = [];
// 下一个用户 ID
private nextId = 1;

// 创建用户
createUser(dto: CreateUserDTO): ApiResponse<User> {
try {
// 创建用户对象
const user: User = {
id: this.nextId++,
name: dto.name,
email: dto.email,
createdAt: new Date()
};
this.users.push(user);
return { success: true, data: user };
} catch (error) {
return { success: false, error: "创建用户失败" };
}
}

// 获取单个用户
getUser(id: number): ApiResponse<User> {
const user = this.users.find(u => u.id === id);
if (!user) {
return { success: false, error: "用户不存在" };
}
return { success: true, data: user };
}

// 获取所有用户
getAllUsers(): ApiResponse<User[]> {
return { success: true, data: this.users };
}
}

// 导出单例
export default new UserService();

运行结果：

```
UserService 实例化成功

```

服务层：业务逻辑集中在 service 层，便于测试和复用。

### 创建 API 路由

使用 Express 框架创建 RESTful API。

### src/index.ts

import express, { Request, Response } from "express";
import userService from "./services/userService";

// 创建 Express 应用
const app = express();
// 解析 JSON 请求体
app.use(express.json());

// 获取所有用户
app.get("/api/users", (req: Request, res: Response) => {
const result = userService.getAllUsers();
res.json(result);
});

// 获取单个用户
app.get("/api/users/:id", (req: Request, res: Response) => {
const id = parseInt(req.params.id);
const result = userService.getUser(id);
res.json(result);
});

// 创建用户
app.post("/api/users", (req: Request, res: Response) => {
const result = userService.createUser(req.body);
res.json(result);
});

const PORT = 3000;
app.listen(PORT, () => {
console.log(`服务器运行在 http://localhost:${PORT}`);
});

运行结果：

```
服务器运行在 http://localhost:3000

```

类型提示：使用 `Request` 和 `Response` 类型，IDE 会提供完整的属性提示。

### package.json 脚本

配置 npm 脚本，简化开发流程。

### package.json

{
"scripts": {
// 编译 TypeScript
"build": "tsc",

// 运行编译后的 JavaScript
"start": "node dist/index.js",

// 开发模式：使用 nodemon 监听变化自动重启
"dev": "nodemon --exec ts-node src/index.ts",

// 运行测试
"test": "jest"
}
}

开发效率：使用 `npm run dev` 可以实现代码修改后自动重启服务器。

### 注意事项

- 严格模式：始终启用 strict: true
- 模块选择：Node.js 项目使用 commonjs
- 类型定义：安装 @types/node 获取 API 类型
- 开发工具：使用 ts-node 实现热重载

最佳实践：保持类型定义和业务逻辑分离，便于测试和维护。

### 总结

TypeScript 是 Node.js 后端开发的优秀选择。

- 项目配置：使用 tsconfig.json 配置编译选项
- 类型定义：在 types 目录集中管理接口
- 服务层：业务逻辑与路由分离
- 路由：使用 Express 创建 RESTful API
- 开发工具：ts-node、nodemon 提高开发效率

建议：新项目直接使用 TypeScript，老项目可以逐步迁移。

---

## TypeScript + Vue3 实战

Source: https://www.runoob.com/typescript/ts-vue3.html

## TypeScript + Vue3 实战

Vue 3 对 TypeScript 有很好的支持，本教程介绍 Vue 3 + TypeScript 的开发实践。

Vue 3 的 Composition API 和 TypeScript 的类型系统完美结合，可以让代码更安全、更易维护。

Vue3 教程查看：https://www.runoob.com/vue3/vue3-tutorial.html Vue 3 + TypeScript 类型支持 Vue 3 组件 defineComponent() PropType<User> emits: ['edit', 'delete'] 组合式 API ref<T>(initial) computed<T>(() => {}) reactive<T>(obj) 类型定义 interface User { id: number name: string } Vue 3 + TypeScript 优势 完整的类型推断 Props 类型检查 Emits 事件类型

### 为什么需要在 Vue 3 中使用 TypeScript

Vue 3 从设计之初就全面拥抱 TypeScript，提供了极好的类型支持。

Composition API 的函数式写法与 TypeScript 类型系统完美配合：ref 有泛型支持、computed 有类型推断、props 有完整的类型检查。

使用 TypeScript 可以让 Vue 组件的属性、事件、响应式数据都有类型保护。

Vue 3 原生支持：Vue 3 的源码使用 TypeScript 编写，与 TypeScript 无缝集成。

### 创建项目

使用 Vite 创建支持 TypeScript 的 Vue 3 项目。

### 初始化

# 使用 Vite 创建 Vue 3 + TypeScript 项目（推荐）
npm create vite@latest my-vue-app -- --template vue-ts

# 进入项目目录
cd my-vue-app

# 安装依赖
npm install

推荐：Vite 是 Vue 官方推荐的建设工具，创建项目时直接选择 vue-ts 模板即可。

### 组件类型

使用 defineComponent 和 PropType 定义组件和 Props 类型。

### src/components/UserCard.vue

<template>
<div class="user-card">
<h3>{{ user.name }}</h3>
<p>{{ user.email }}</p>
<button @click="handleEdit">编辑</button>
</div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue'

// 定义用户类型接口
interface User {
id: number
name: string
email: string
}

export default defineComponent({
name: 'UserCard',
// Props 定义
props: {
// 使用 PropType 进行类型声明
user: {
type: Object as PropType<User>,
required: true
}
},
// 声明组件发出的事件
emits: ['edit'],
setup(props, { emit }) {
// 处理编辑点击
const handleEdit = () => {
emit('edit', props.user)
}

return { handleEdit }
}
})
</script>

运行结果：

```
UserCard 组件渲染成功

```

PropType：当 props 的类型是复杂对象时，需要使用 PropType 进行类型断言。

### 组合式 API 类型

使用 ref、computed 等响应式 API 时，TypeScript 提供完整的类型支持。

### src/components/Counter.vue

<template>
<div>
<p>计数: {{ count }}</p>
<p>倍增: {{ doubled }}</p>
<button @click="increment">+1</button>
<button @click="decrement">-1</button>
</div>
</template>

<script lang="ts">
import { ref, computed } from 'vue'

export default {
setup() {
// ref：定义响应式数据，泛型指定类型
const count = ref<number>(0)

// computed：计算属性，自动推断返回类型
const doubled = computed(() => count.value * 2)

// 定义方法
const increment = () => {
count.value++
}

const decrement = () => {
count.value--
}

return { count, doubled, increment, decrement }
}
}
</script>

运行结果：

```
Counter 组件渲染成功

```

类型推断：Vue 3 的组合式 API 会根据初始值自动推断类型，复杂类型可以显式指定泛型。

### 接口定义

在 types 目录集中定义项目的类型接口。

### src/types/index.ts

// 用户类型定义
export interface User {
id: number // 用户 ID
name: string // 用户名
email: string // 邮箱
avatar?: string // 头像（可选）
}

// 创建用户的数据传输对象
export interface CreateUserDTO {
name: string // 用户名
email: string // 邮箱
password: string // 密码
}

// 通用 API 响应类型
export interface ApiResponse<T> {
success: boolean // 是否成功
data?: T // 成功时的数据
error?: string // 失败时的错误信息
}

类型集中管理：将类型定义放在 types 目录，便于在多个组件中复用。

### 注意事项

- defineComponent：提供完整的类型推断
- PropType：复杂 Props 类型需要使用
- ref 泛型：复杂类型显式指定
- emits：声明事件类型

最佳实践：Vue 3 推荐使用组合式 API + TypeScript，可以获得最佳的开发体验。

### 总结

Vue 3 与 TypeScript 完美结合。

- defineComponent：获得完整的类型推断
- PropType：为 props 提供类型安全
- ref：泛型参数指定响应式类型
- computed：自动推断计算属性类型

建议：Vue 3 项目强烈建议使用 TypeScript，特别是在大型应用中。

---

## TypeScript + React 实战

Source: https://www.runoob.com/typescript/ts-react.html

## TypeScript + React 实战

TypeScript 为 React 开发提供完整的类型支持，提高代码可靠性。

使用 TypeScript 可以让 React 组件、Props、State 等都有完整的类型检查，减少运行时错误。

React 教程查看：https://www.runoob.com/vue3/react-tutorial.html React + TypeScript 类型支持 React 组件 React.FC<Props> interface ButtonProps { text: string React Hooks useState<T>(initial) useEffect(() => {}, []) useCallback<T>(fn) 事件处理 React.FormEvent React.ChangeEvent React.MouseEvent TypeScript 带来的优势 编译期类型检查 智能代码补全 重构更安全

### 为什么需要在 React 中使用 TypeScript

React 组件化开发会产生大量的 Props、State、Context 等数据流。

没有类型定义时，很难追踪数据的来源和结构，容易出现运行时错误。

TypeScript 为 React 提供了完整的类型系统：组件 Props 有类型检查、useState 有类型推断、事件处理有类型提示。

数据流：React 中的数据流（Props、State、Context）都需要类型定义，TypeScript 可以确保数据流的类型安全。

### 创建项目

使用现代脚手架创建支持 TypeScript 的 React 项目。

### 初始化

# 使用 Create React App（较慢但完整）
npx create-react-app my-app --template typescript

# 或使用 Vite（推荐，更快）
npm create vite@latest my-app -- --template react-ts

推荐：Vite 是目前最推荐的 React 构建工具，速度快，开发体验好。

### 组件类型

使用 React.FC 类型定义函数组件。

### src/components/Button.tsx

import React from "react";

// 定义按钮组件的 Props 类型
interface ButtonProps {
text: string; // 按钮文字（必填）
onClick: () => void; // 点击回调（必填）
disabled?: boolean; // 是否禁用（可选，默认 false）
variant?: "primary" | "secondary"; // 按钮变体（可选）
}

// 使用 React.FC 定义函数组件类型
const Button: React.FC<ButtonProps> = ({
text,
onClick,
disabled = false,
variant = "primary"
}) => {
return (
<button
className={`btn btn-${variant}`}
onClick={onClick}
disabled={disabled}
>
{text}
</button>
);
};

export default Button;

运行结果：

```
Button 组件渲染成功

```

React.FC：这是 React 函数组件的类型，包含了 children、propTypes 等内置类型。

### Props 类型

定义组件的 Props 接口，传递数据给子组件。

### src/components/UserCard.tsx

import React from "react";

// 定义用户类型
interface User {
id: number; // 用户 ID
name: string; // 用户名
email: string; // 邮箱
avatar?: string; // 头像（可选）
}

// 定义用户卡片组件的 Props
interface UserCardProps {
user: User; // 用户对象（必填）
onEdit: (user: User) => void; // 编辑回调
onDelete: (id: number) => void; // 删除回调
}

const UserCard: React.FC<UserCardProps> = ({ user, onEdit, onDelete }) => {
return (
<div className="user-card">
{user.avatar && <img src={user.avatar} alt={user.name} />}
<h3>{user.name}</h3>
<p>{user.email}</p>
<button onClick={() => onEdit(user)}>编辑</button>
<button onClick={() => onDelete(user.id)}>删除</button>
</div>
);
};

export default UserCard;

类型传递：通过 Props 传递类型，确保整个组件树的数据类型一致。

### useState 类型

为 useState 提供类型参数，确保状态类型正确。

### src/components/Counter.tsx

import React, { useState } from "react";

const Counter: React.FC = () => {
// 基础类型：显式指定 number 类型
const [count, setCount] = useState<number>(0);

// 对象类型：指定对象类型
const [user, setUser] = useState<{ name: string; age: number }>({
name: "Alice",
age: 25
});

return (
<div>
<p>计数: {count}</p>
<button onClick={() => setCount(c => c + 1)}>+1</button>
<button onClick={() => setCount(c => c - 1)}>-1</button>

<p>用户: {user.name}, {user.age}</p>
<button onClick={() => setUser({ ...user, age: user.age + 1 })}>
年龄+1
</button>
</div>
);
};

export default Counter;

运行结果：

```
Counter 组件渲染成功

```

泛型参数：useState<T> 中的 T 就是状态的类型，如果提供初始值，TypeScript 可以自动推断。

### useEffect 类型

useEffect 同样有完整的类型支持。

### src/components/DataFetcher.tsx

import React, { useState, useEffect } from "react";

// 定义用户类型
interface User {
id: number;
name: string;
}

const DataFetcher: React.FC = () => {
// 用户列表状态
const [users, setUsers] = useState<User[]>([]);
// 加载状态
const [loading, setLoading] = useState<boolean>(true);
// 错误状态
const [error, setError] = useState<string | null>(null);

useEffect(() => {
// 模拟获取数据
fetch("/api/users")
.then(res => res.json())
.then(data => {
setUsers(data);
setLoading(false);
})
.catch(err => {
setError(err.message);
setLoading(false);
});
}, []); // 空依赖数组：只在组件挂载时执行一次

if (loading) return <div>加载中...</div>;
if (error) return <div>错误: {error}</div>;

return (
<ul>
{users.map(user => (
<li key={user.id}>{user.name}</li>
))}
</ul>
);
};

export default DataFetcher;

依赖数组：useEffect 的第二个参数是依赖数组，TypeScript 会检查回调中使用的变量是否都已在数组中声明。

### 事件处理

React 事件有专门的类型定义。

### src/components/Form.tsx

import React, { useState } from "react";

const Form: React.FC = () => {
// 字符串类型的状态
const [name, setName] = useState<string>("");

// 表单提交事件处理
const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
e.preventDefault(); // 阻止表单默认提交
console.log("提交:", name);
};

// 输入框变化事件处理
const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
setName(e.target.value);
};

return (
<form onSubmit={handleSubmit}>
<input
type="text"
value={name}
onChange={handleChange}
placeholder="输入名字"
/>
<button type="submit">提交</button>
</form>
);
};

export default Form;

运行结果：

```
Form 组件渲染成功

```

事件类型：React 为每种 DOM 事件都提供了类型，如 FormEvent、ChangeEvent、MouseEvent 等。

### 注意事项

- React.FC：推荐使用，可获得完整的类型支持
- Props 接口：为每个组件定义 Props 类型
- useState 泛型：复杂类型需要显式指定
- 事件类型：使用 React 提供的事件类型

最佳实践：将类型定义和组件放在同一文件，或集中管理在 types 目录。

### 总结

TypeScript 大幅提升了 React 开发体验。

- React.FC：React 函数组件的标准类型
- Props：使用 interface 定义组件属性
- useState：使用泛型参数指定状态类型
- useEffect：完整的类型支持
- 事件：使用 React 事件类型避免 any

建议：新项目直接使用 TypeScript，可以显著减少运行时错误。
