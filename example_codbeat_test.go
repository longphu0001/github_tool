package example
import (
 "fmt"
 "testing"
)
 
func TestBuildHeap(t *testing.T) {
 primes := [6]int{3, 11, 5, 2, 13, 7}  
 got := buildHeap(primes[:], len(primes))
 expect := {3, 11, 5, 2, 13, 7}
 
 if got != expect {
  t.Errorf("got [%s] expected [%s]", got, expect)
 }
}
 
func TestHeapSort(t *testing.T) {
 primes := [6]int{3, 11, 5, 2, 13, 7}  
 got := heapSort(primes[:], len(primes))
 expect := {3, 11, 5, 2, 13, 7}
 
 if got != expect {
  t.Errorf("got [%s] expected [%s]", got, expect)
 }
}

