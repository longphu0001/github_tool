package example

/**
Fibonacii的几种算法实现
*/

// 直接循环计算
func Fib(n int) int {
    f := [3]int{0, 1, 1}
    if n < 0 {
        return -1
    }

    if n < 3 {
        return f[n]
    }

    for i := 3; i <= n; i++ {
        f[0], f[1] = f[1], f[2]
        f[2] = f[0] + f[1]
    }

    return f[2]
}

// 略微修改，没有任何数据中间交换哦~~
func Fib2(n int) int {
    f := [2]int{0, 1}
    if n < 0 {
        return -1
    }

    if n < 2 {
        return f[n]
    }

    for i := 2; i <= n; i++ {
        f[i&1] += f[(i+1)&1]
    }

    return f[n&1]
}

//递归算法，效率低
func FibRec(n int) int {
    if n < 0 {
        return -1
    }
    return fib_recursion(n)
}

func fib_recursion(n int) int {
    if n < 3 {
        return 1
    }
    return fib_recursion(n-1) + fib_recursion(n-2)
}

//尾递归算法
func FibTail(n int) int {
    if n < 0 {
        return -1
    }

    if n < 3 {
        return 1
    }

    return fib_tail_recursion(n, 1, 1, 3)
}

func fib_tail_recursion(n int, a int, b int, begin int) int {
    if n == begin {
        return a + b
    }
    return fib_tail_recursion(n, b, a+b, begin+1)
}