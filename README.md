# TheGoLerarnNote
learn from book &lt;TheGoProgrammingLanguage>

##2020-1-18
###学习了tag


tag做为工程的一个时间点，每个tag仅于一个commit号绑定。系统中自然不容许相同的tag。

tag不会像branch，一旦commit新的工程，tag不会更新。

tag也属于标识，不push只会留在本地。本例中v0.2不在github，只在本地有。

若目前tag分支与master脱离，以及 若此时的tag就是master分支但是
没有push master。都不会影响github上的master。v0.1 和v0.3 的readme.md与本文不符就是证明。

所以 1.请在需要tag时，commit后再tag出一个时间点（一个分支），push完tag后请立刻回到原来所在分支（branch）。2.或者 tag + commit号来指定项目时间点来打标签。
###学了书 hellworld

中括弧有规则


大写pubilc 小写 private

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
- new 是给 struct  make 是给 map slice chan\
