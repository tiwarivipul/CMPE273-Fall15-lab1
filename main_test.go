package main
import "testing"

type testpair struct {
values int 
fibo int
}
var tests = []testpair{
{ 3,2 },
{ 5,5 },
}
func TestFib(t *testing.T){
for _, pair :=range tests{
	v:=fib(pair.values)
	if v!= pair.fibo {
	t.Error(
	"expected ", pair.fibo, " got",v,
	)
	}
	
	}
}