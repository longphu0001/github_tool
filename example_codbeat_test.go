package example
import (
    "testing"
)
 
func TestFib(t *testing.T) {
    n := 10
    f := Fib(n)
    if f != 55 {
        t.Error("Fib() failed. Got", f, "Expected 55 ")
    }
}

func TestFib2(t *testing.T) {
    n := 10
    f := Fib2(n)
    if f != 55 {
        t.Error("Fib2() failed. Got", f, "Expected 55 ")
    }
}

func TestFibRec(t *testing.T) {
    n := 10
    f := FibRec(n)
    if f != 55 {
        t.Error("FibRec() failed. Got", f, "Expected 55 ")
    }
}

func TestFibTail(t *testing.T) {
    n := 10
    f := FibTail(n)
    if f != 55 {
        t.Error("FibTail() failed. Got", f, "Expected 55")
    }
}

func BenchmarkFib(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Fib(1000)
    }
}

func BenchmarkFib2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Fib2(1000)
    }
}

func BenchmarkFibTail(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FibTail(1000)
    }
}

