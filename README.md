# TheGoLerarnNote
learn from book &lt;TheGoProgrammingLanguage>

##2020-1-18
###学习了tag

>git tag -a v0.1.0 -m "release 0.1.0 version"


tag做为工程的一个时间点，每个tag仅于一个commit号绑定。系统中自然不容许相同的tag。

tag不会像branch，一旦commit新的工程，tag不会更新。

tag也属于标识，不push只会留在本地。本例中v0.2不在github，只在本地有。

若目前tag分支与master脱离，以及 若此时的tag就是master分支但是
没有push master。都不会影响github上的master。v0.1 和v0.3 的readme.md与本文不符就是证明。

所以 1.请在需要tag时，commit后再tag出一个时间点（一个分支），push完tag后请立刻回到原来所在分支（branch）。2.或者 tag + commit号来指定项目时间点来打标签。
###学了书 hellworld

中括弧有规则


大写pubilc 小写 private

##2020-1-25
###banch
>git checkout -b XXX 新建 \
>git checkout xxx 切换 \
>git branch -d xxx 删除 \
>git merge xxx 当前HEAD向（与）xxx处合并

[分支与合并的参考链接](https://git-scm.com/book/zh/v2/Git-%E5%88%86%E6%94%AF-%E5%88%86%E6%94%AF%E7%9A%84%E6%96%B0%E5%BB%BA%E4%B8%8E%E5%90%88%E5%B9%B6 "Git 分支 - 分支的新建与合并")
- 在合并时，如果当前HEAD为xxx的父快照，则一般会
>Fast-forward \
>![Fast-forward](../notepic/Mmergeiss4 Fastforward.png)

一般是自动合并成功的，因为快照层的操作，也不需要再commit

- 合并时 ，若HEAD与xxx快照分叉，且对同一文件有修改，则会冲突
>CONFLICT  \
>![CONFLICT]( ../notepic/Mmergeiss3 CONFLICT.png)

显然此时自动合并失败，需要手动调整。可以从git status观察到冲突文件，并在文件中观察到
> <<<<<<< HEAD:Readme.md
> 
> =======
> 
> &gt;&gt;&gt;&gt;&gt;&gt;&gt; iss3:Readme.md

你可以手动去修改这些文件，比如去掉你不要的内容 以及 <<<<<<< 、&gt;&gt;&gt;&gt;&gt;&gt;&gt; 、=======
或者借助工具修改这些文件 \
但是都意味着需要重新add与commit,以下是修改后并add后的git status
>Clear 还需要commit哟
>![Clear]( ../notepic/Mmergeiss3 Clear.png)


### func与 数组/切片

在go中
> [5]int 为声明数组
>> 数组属于变量，赋值（=）及函数传参时复制副本。故若a = b,则a与b之后操作互不干涉。

> []int 为声明切片
>> 数组切片名属于指针，赋值及函数传参时复制地址。故若 a = b ,则ab之后同名 \
>> 此外，在go自己内存回收机制中，似乎允许函数外指针指向函数内声明的数组（变量），这在C中函数结束，函数内声明变量销毁，会出现函数外指针指向空的错误。



就像使用master一样，commit&&push注意，master就不会再动了。你现在commit将更新iss3

### ... 、interface{}
>interface{} 一个没有实现要求的接口
>> 其性质为指针，能够接受一切类型，就像C#中的object C++的空指针 是一切的父类

>... 实际上是切片的特殊语法 也可以归类至不定参数
>> 当函数func funca(arge []int ){} 则输入应为 funca([]{1,2,3,4}) \
>> 当函数func funca(arge ...int ){} 则输入应为 funca(1,2,3,4)、funca(arge...)  就是方便

##2020-1-27
###branch -b -iss4
今天我们就来建立iss4
并准备合并 iss3 iss4 master
### type 类方法 类
####type  
定义 同名 同define
####类方法 
首先不仅仅可以给struct 也可以给内置类型定义方法，以及指针不行！！
>func (a structname) xxxx(input) output 

a 在此处的身份很像== 必须显式表达的this \


#### 类
当然 go没有类 class, 因为任何类型都可以声明自己的方法，而go有又禁止继承，或者说由接口更加通用地代替继承。

所以go只有struct 用法就是结构体
> type xxxx struct{}
#####初始化
> 以 type Rect struct{}为例
> - rect1 := new(Rect) \
> - rect2 := &Rect{} \
> - func NewRect() *Rect{}

### 注意
- 1 [成员方法中，"this"对象是否使用指针 ](https://bbs.csdn.net/topics/390974889?page=1 "定义成员方法时，带*与否二者区别在哪").
声明类方法时，
>func (a structname) xxxx(input) output 

   如果希望a在运行函数后能够改表，则需要让a为指针 即 a *structname \
否则的话 a 依旧是个副本（当然若变量本身是指针，则是特殊情况）
- 2 [go 中没有 ->](https://bbs.csdn.net/topics/390974889?page=1 "定义成员方法时，带*与否二者区别在哪"). \
在go中 指针没有字段（成员、方法）,【注：指针是不允许有成员方法的】所以 以下两端代码将会没有歧义而等价：
>设 a *structname 
>> (*a).Area() <==> a.Area()

所以自然就没有 -> 的必要了。 \
但需要注意，只有涉及字段（成员）的结构体struct ，可以因上文原因同义。本质上
>>  *a 与 a 并不同义
- 3 new 是给 struct用  make 是给 map slice chan\
都得到的是指针（&struct）
##2020-1-29
###接口
#### 定义与赋值
定义
> 注意同样会存在  (*a).xxx 自动<==> a.xxx \
> 以下代码
>> type LessAddr interface{ \
>>    Less (b Integer) bool \
>>    Add (b Integer)  \
>>}
>
>若a为指针则自动生成 func (a *Integer) less (b Integer){ ....a to a\* }
>即以下合法
>>type Integer int //这句必须存在
>> var a Integer =1; 
>>func Less (b Integer) bool
>>var b LessAddr = &a;
>> 此时 Less(b)将自动变为  less (b *)
####接口查询 类查询
> if _,ok:= b.(LessAddr) ok{ \
>b 满足接口吗  当然！

> if _, ok:= c.(*structname) ; ok { \
 >c 的指向对象是类structname么？ 请注意c本身可能是接口或类。\
 
####类型查询
>如果是问是不是内置类型 则用 c.(type)
>switch v.(type) 请和switch配合使用


### 注意
- 1 成员方法中，"this" 类型必须是type定义的，且不能是接口 否则报错 \
你不能直接！为内置类型添加方法
>type miao struct{ //...  \
>func (this miao) xxx()(){ //...
- 2 成员方法的另一种定义方法就是，struct添加函数数据成员（类似函数指针）
>type miao struct{ \
>do func ()() \
>//...  } \
> func xxx()(){ //... 
>> 使用方法如下： \
>> nya:=&miao{do:xxx } \
>> nya.do()
- 3 接口查询时，括弧不可省去 以及结构是返回的第二个参数

##2020-1-30
###[test](https://www.cnblogs.com/52php/p/6985411.html)
>Golang 单元测试对文件名和方法名，参数都有很严格的要求。例如：
>- 1、文件名必须以 _test.go 结尾
>-  2、方法名必须是 Test 开头
>-  3、方法参数必须是 t *testing.T 或 b *testing.B
##2020-2-1
###channel

###select
select 对case选择是无序的、乱序的、以及如果case两个以上满足条件，就都执行的 ！
### Once
全局首次操作，sync.Once once.do保证全局唯一性操作 （就是初始化）
保证该函数无论开出多少次协程，只会运行一次
>var once sync.Once 
>func xxx(){ \
>once.do(yyy) \
>}
###注意
- fatal error: all goroutines are asleep - deadlock!
主程序堵塞，且没有其他协程运行（就是没有其他协程能帮（maingoroutine）主协程解堵塞） \
一种可能是声明channel时没有开缓冲空间，此时只剩（maingoroutine）主协程在运行时，往channel中读、写都会报错 
>make(chan int);

另一种情况就是在主程序读操作阻塞，且此时又只剩主协程，常见现象为询问channel关闭时，channel没数据也没关闭
> _,x :=ch 

第一种情况 \
如果声明channel没有给缓冲空间， 那对此channel写时，因为没有缓存空间只能阻塞，直到有协程能读时才会解除阻塞。
在此条件下，本协程的程序一旦阻塞（尽管本协程下文有读操作，但是因为堵塞永远也不会等到其执行），只能靠其他同时运行的协程解决。 \
如果发生这种阻塞且程序只有正在阻塞的main goroutine在运行意味着阻塞永远不会解除，系统会报错


第二种情况 \
close对于一个channel来说也是个写数据，但比较特殊，其确认方法即读数据。在channel有数据时，正常返回数据。在关闭时第二参数正常返回false
没数据、没关闭、又只剩main goroutine,还是那句话 永远也等不到下文！


> fatal error: all goroutines are asleep - deadlock!
>>解决方法: 
>>- 1 更改逻辑，添加goroutine 使channel完成一写一读 
>>- 2 为channel 添加缓存 make(chan int, 1024);
>> -3 尝试让问题channel在主程序之外执行

- select 的选择是无序的、乱序的、以及如果满足两个以上就都执行的 ！
##2020-2-2
###TCP
测一下go与C++相连状况
- C++socket编程的数据格式与go不同
>C++ 的send/recv主要使用 \0来结束字符串的传输（当然要小于参数len） \
>在go中需要使用\000来代表 C++的 \0


>go中 使用io.EOF来判定结束字符串的传输，C++完全无法输出该符号 \
>因为C++无法给出对方go终端来自报错模块的专有符号， C++发完信息后，go中，虽然会结束一个conn.Read 但是不会
>触发err == EOF,而是err == nil 代表没发送完。并接收到C++send参数中len长的字符串，在C++中\0后面的部分也会发送过来！！ 


- string(result) 遇到\0会停止喵
- bytes 似乎可以无限长喵 用write来输入
##2020-2-3
###HTTP
- [http](https://www.jianshu.com/p/4e8cdf3b2f88) 注册路由中需要提供url与func的对应
- [url与http](https://www.cnblogs.com/zhuanzhuruyi/p/6508565.html)  函数handler应该从虚拟目录部分开始吧？
- [Content-Type post](https://www.cnblogs.com/tugenhua0707/p/8975121.html)
- [重定向](https://www.cnblogs.com/bq-med/p/8602629.html) 本服务器无法处理该请求 返回3XX 与对应新的location给客户端
- [cookie](https://www.cnblogs.com/bq-med/p/8603664.html) http验证身份用
- [TCP+RPC](https://blog.csdn.net/a1158375969/article/details/79516112)
