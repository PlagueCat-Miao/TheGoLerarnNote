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
>git checkout -b iss3

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

 

