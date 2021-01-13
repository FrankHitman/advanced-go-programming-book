package main

import (
	"runtime"
)

// func CopyFile(dstName, srcName string) (written int64, err error) {
// 	src, err := os.Open(srcName)
// 	if err != nil {
// 		return
// 	}
// 	defer src.Close()
// 	// defer 语句可以保证即使 io.copy 发生了异常，文件依然可以安全地关闭
//
// 	dst, err := os.Create(dstName)
// 	if err != nil {
// 		return
// 	}
// 	defer dst.Close()
//
// 	written, err = io.Copy(dst, src)
// 	return
// }

// func MyRecover() interface{} {
// 	return recover()
// }
//
// func main() {
// 	// 可以正常捕获异常
// 	defer MyRecover()
// 	panic(1)
// }
// func main() {
// 	// 无法捕获异常
// 	defer recover()
// 	panic(1)
// }
// “必须要和有异常的栈帧只隔一个栈帧，recover函数才能正常捕获异常。
// 换言之，recover函数捕获的是祖父一级调用函数栈帧的异常（刚好可以跨越一层defer函数）！”

func main() {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case runtime.Error:
				// 这是运行时错误类型异常
			case error:
				// 普通错误类型异常
			default:
				// 其他类型异常
			}
		}
	}()

	// ...
}
