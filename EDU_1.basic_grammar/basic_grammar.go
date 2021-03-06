package GO_std

// Go 标记
// Go 程序可以由多个标记组成，可以是关键字，标识符，常量，字符串，符号。如以下 GO 语句由 6 个标记组成：
//1. fmt
//2. .
//3. Println
//4. (
//5. "Hello, World!"
//6. )

// 行分隔符
//在 Go 程序中，一行代表一个语句结束。每个语句不需要像 C 家族中的其它语言一样以分号 ; 结尾，因为这些工作都将由 Go 编译器自动完成。
//
//如果你打算将多个语句写在同一行，它们则必须使用 ; 人为区分，但在实际开发中我们并不鼓励这种做法。
//
//以下为两个语句：
//
//fmt.Println("Hello, World!")
//fmt.Println("菜鸟教程：runoob.com")

// 注释
//注释不会被编译，每一个包应该有相关注释。
//
//单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。多行注释也叫块注释，均已以 /* 开头，并以 */ 结尾。如：
// 单行注释
/*
 Author by 菜鸟教程
 我是多行注释
*/

// 标识符
// 标识符用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字。
//
//以下是有效的标识符：
//mahesh   kumar   abc   move_name   a_123
//myname50   _temp   j   a23b9   retVal

// 关键字
//break	default	func	interface	select
//case	defer	go	map	struct
//chan	else	goto	package	switch
//const	fallthrough	if	range	type
//continue	for	import	return	var
//除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：
//
//append	bool	byte	cap	close	complex	complex64	complex128	uint16
//copy	false	float32	float64	imag	int	int8	int16	uint32
//int32	int64	iota	len	make	new	nil	panic	uint64
//print	println	real	recover	string	true	uint	uint8	uintptr
// 程序一般由关键字、常量、变量、运算符、类型和函数组成。
//
//程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。
//程序中可能会使用到这些标点符号：.、,、;、: 和 …。

// Go 语言的空格
// Go 语言中变量的声明必须使用空格隔开，如：
// var age int;
//语句中适当使用空格能让程序更易阅读。
//无空格：
//fruit=apples+oranges;
//在变量与运算符间加入空格，程序看起来更加美观，如：
//fruit = apples + oranges;
