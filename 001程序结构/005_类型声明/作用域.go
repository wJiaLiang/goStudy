package main


/*
一个声明语句将程序中的实体和一个名字关联，比如一个函数或一个变量。声明语句的作用域是指源代码中可以有效使用这个名字的范围。
句法块是由花括弧所包含的一系列语句，就像函数体或循环体花括弧包裹的内容一样。句法块内部声明的名字是无法被外部块访问的。

声明在代码中并未显式地使用花括号包裹起来，我们称之为词法块:
	1、对全局的源代码来说，存在一个整体的词法块，称为全局词法块
	2、对于每个包；每个for、if和switch语句，也都有对应词法块
	3、每个switch或select的分支也有独立的词法块；
	4、也包括显式书写的词法块（花括弧包含的语句）

1、对于内置的类型、函数和常量，比如int、len和true等是在全局作用域
2、任何在函数外部（也就是包级语法域）声明的名字可以在同一个包的任何源文件中访问的
3、控制流标号，就是break、continue或goto语句后面跟着的那种标号，则是函数级的作用域。
4、在函数中词法域可以深度嵌套，因此内部的一个声明可能屏蔽外部的声明


*/

func main()  {




}
