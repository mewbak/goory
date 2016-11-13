package goory_test

import (
	"testing"

	"github.com/bongo227/goory"
)

func TestExampleFpadd(t *testing.T) {
	double := goory.DoubleType()
	module := goory.NewModule("test")

	// Function: fpadd
	{
		function := module.NewFunction("fpadd", double)
		a := function.AddArgument(double, "a")
		b := function.AddArgument(double, "b")

		block := function.Entry()

		result := block.Fadd(a, b)
		block.Ret(result)
	}

	got := module.LLVM()
	expected := `define double @fpadd(double %a, double %b){
	entry:
		%temp0 = fadd double %a, %b
		ret double %temp0
}`

	if got != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, got)
	}
}

func TestExampleMulAdd(t *testing.T) {
	integer := goory.IntType(32)
	module := goory.NewModule("test")

	// Function: mul_add
	{
		function := module.NewFunction("mul_add", integer)
		x := function.AddArgument(integer, "x")
		y := function.AddArgument(integer, "y")
		z := function.AddArgument(integer, "z")

		block := function.Entry()

		result := block.Mul(x, y)
		result = block.Add(result, z)
		block.Ret(result)
	}

	got := module.LLVM()
	expected := `define i32 @mul_add(i32 %x, i32 %y, i32 %z){
	entry:
		%temp0 = mul i32 %x, %y
		%temp1 = add i32 %temp0, %z
		ret i32 %temp1
}`

	if got != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, got)
	}
}

func TestExampleGcd(t *testing.T) {
	integer := goory.IntType(32)
	module := goory.NewModule("test")

	// Function: gcd
	{
		function := module.NewFunction("gcd", integer)
		x := function.AddArgument(integer, "x")
		y := function.AddArgument(integer, "y")

		block := function.Entry()

		xyEqual := block.Icmp(goory.IntEq, x, y)

		trueBlock := function.AddBlock()
		trueBlock.Ret(x)

		falseBlock := function.AddBlock()
		xyLess := falseBlock.Icmp(goory.IntUlt, x, y)

		elseIfBlock := function.AddBlock()
		call1 := elseIfBlock.Call(function, x, elseIfBlock.Sub(y, x))
		elseIfBlock.Ret(call1)

		elseBlock := function.AddBlock()
		call2 := elseBlock.Call(function, elseBlock.Sub(x, y), y)
		elseBlock.Ret(call2)

		block.CondBr(xyEqual, trueBlock, falseBlock)
		falseBlock.CondBr(xyLess, elseIfBlock, elseBlock)
	}

	got := module.LLVM()
	expected := `define i32 @gcd(i32 %x, i32 %y){
	entry:
		%temp0 = icmp eq bool %x, %y
		br i1 %temp0, label %temp1, label %temp3
	temp1:
		ret i32 %x
	temp3:
		%temp4 = icmp ult bool %x, %y
		br i1 %temp4, label %temp5, label %temp9
	temp5:
		%temp6 = sub i32 %y, %x
		%temp7 = call i32 @gcd(i32 %x, i32 %temp6)
		ret i32 %temp7
	temp9:
		%temp10 = sub i32 %x, %y
		%temp11 = call i32 @gcd(i32 %temp10, i32 %y)
		ret i32 %temp11
}`

	if got != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, got)
	}

}

func TestExampleFibonaci(t *testing.T) {
	integer := goory.IntType(32)
	module := goory.NewModule("test")

	// Function: fibonaci
	{
		function := module.NewFunction("fib", integer)
		n := function.AddArgument(integer, "n")

		block := function.Entry()

		trueBlock := function.AddBlock()
		nEqualsZero := block.Icmp(goory.IntEq, n, goory.Constant(integer, 0))
		nEqualsOne := block.Icmp(goory.IntEq, n, goory.Constant(integer, 1))
		nOneOrZero := block.Or(nEqualsZero, nEqualsOne)
		trueBlock.Ret(n)

		falseBlock := function.AddBlock()
		fib1 := falseBlock.Call(function, falseBlock.Sub(n, goory.Constant(integer, 1)))
		fib2 := falseBlock.Call(function, falseBlock.Sub(n, goory.Constant(integer, 2)))
		falseBlock.Ret(falseBlock.Add(fib1, fib2))

		block.CondBr(nOneOrZero, trueBlock, falseBlock)
	}

	got := module.LLVM()
	expected := `define i32 @fib(i32 %n){
	entry:
		%temp1 = icmp eq bool %n, 0
		%temp2 = icmp eq bool %n, 1
		%temp3 = or bool %temp1 %temp2
		br i1 %temp3, label %temp0, label %temp5
	temp0:
		ret i32 %n
	temp5:
		%temp6 = sub i32 %n, 1
		%temp7 = call i32 @fib(i32 %temp6)
		%temp8 = sub i32 %n, 2
		%temp9 = call i32 @fib(i32 %temp8)
		%temp10 = add i32 %temp7, %temp9
		ret i32 %temp10
}`

	if got != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, got)
	}
}
