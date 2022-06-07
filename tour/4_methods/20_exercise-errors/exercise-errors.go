/* 练习：错误
从之前的练习中复制 Sqrt 函数，修改它使其返回 error 值。

Sqrt 接受到一个负数时，应当返回一个非 nil 的错误值。复数同样也不被支持。

创建一个新的类型

type ErrNegativeSqrt float64
并为其实现

func (e ErrNegativeSqrt) Error() string
方法使其拥有 error 值，通过 ErrNegativeSqrt(-2).Error() 调用该方法应返回 "cannot Sqrt negative number: -2"。

注意: 在 Error 方法内调用 fmt.Sprint(e) 会让程序陷入死循环。可以通过先转换 e 来避免这个问题：fmt.Sprint(float64(e))。这是为什么呢？

修改 Sqrt 函数，使其接受一个负数时，返回 ErrNegativeSqrt 值。 */




package main

import (
	"fmt";
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number:%g", e)
}

func Sqrt(x float64) (float64, error) {
	if x<0 {
		return 0,ErrNegativeSqrt(x)
	}
	// 设置可允许的误差
	//err:=0.001;
	err:=1e-15
	z:=1.0
	for ;math.Abs(z*z - x) >err;{
		z = (z*z+x)/(2*z)
	}
	return z,nil
}

func main() {
	fmt.Println(Sqrt(2))	// 1.414213562373095 <nil>
	fmt.Println(Sqrt(-2))	// 0 cannot Sqrt negative number:-2
}