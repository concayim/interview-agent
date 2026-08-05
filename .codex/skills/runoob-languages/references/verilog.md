# Verilog - 菜鸟教程

Tutorial: https://www.runoob.com/w3cnote/verilog-tutorial.html

---

## Verilog 教程

Source: https://www.runoob.com/w3cnote/verilog-tutorial.html

Verilog HDL（简称 Verilog ）是一种硬件描述语言，用于数字电路的系统设计。可对算法级、门级、开关级等多种抽象设计层次进行建模。

Verilog 继承了 C 语言的多种操作符和结构，与另一种硬件描述语言 VHDL 相比，语法不是很严格，代码更加简洁，更容易上手。

Verilog 不仅定义了语法，还对语法结构都定义了清晰的仿真语义。因此，Verilog 编写的数字模型就能够使用 Verilog 仿真器进行验证。

### 谁适合阅读本教程

本教程主要针对 Verilog 初学者打造。

有一定 Verilog 基础的同学也可以对进阶篇、实例篇进行学习、交流。

### 阅读本教程前，你需要了解的知识

在学习本教程之前，你需要了解数字电路的一些基本信息。

如果你对 C 语言有一定的了解，有助于 Verilog 的快速上手。

#### 第一个 Verilog 设计

4 位宽 10 进制计数器：

### 实例

module counter10(
//端口定义
input rstn, //复位端，低有效
input clk, //输入时钟
output [3:0] cnt, //计数输出
output cout); //溢出位

reg [3:0] cnt_temp ; //计数器寄存器
always@(posedge clk or negedge rstn) begin
if(! rstn)begin //复位时，计时归0
cnt_temp <= 4'b0 ;
end
else if (cnt_temp==4'd9) begin //计时10个cycle时，计时归0
cnt_temp <=4'b000;
end
else begin //计时加1
cnt_temp <= cnt_temp + 1'b1 ;
end
end

assign cout = (cnt_temp==4'd9) ; //输出周期位
assign cnt = cnt_temp ; //输出实时计时器

endmodule

### Cat Me

本人从事过 FPGA 设计、 IC 设计。学生时代用 VHDL 语言设计比较多，目前一直用 Verilog 。为方便查询语法，也为其他学者提供便利的学习通道，特意写此教程。需要说明的是：

- (1) 教程内容是以自己曾经的学习角度进行撰写的，学习起来可能会容易些。其中有不妥之处还望指出，一起交流进步。
- (2) 当用 Verilog 设计完成数字模块后进行仿真时，需要在外部添加激励，激励文件叫 testbench。有时 testbench 设计可能比数字模块本身都复杂。所以前面在介绍 Verilog 基本语法时，几乎没有仿真。后面介绍行为级和时序级相关知识时，会多用仿真说明。

联系人：Think · In · Hardware

全篇教程都是本人手动搜集、整理、编写的，所有设计仿真都有原创或改进。如果您从中受益，您的赞赏或关注将是最不耍流氓的支持，鼓励我饥饿的灵魂去撰写饱满的篇章。
