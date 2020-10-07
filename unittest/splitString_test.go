package unittest

import (
	"fmt"
	"testing"
)

type ExeFunc func(...interface{}) []interface{}

var caseCurId = 0
var TestCases []TestCase

type TestCase struct {
	name   string
	params []interface{}
	f      ExeFunc
	result []interface{}
}

func AddCase(name string, f ExeFunc, params []interface{}) {
	caseCurId++
	name = fmt.Sprintf("case%0d: %s\n", caseCurId, name)
	TestCases = append(TestCases, TestCase{name: name, params: params, f: f, result: nil})
}

func RunTest(id int) {
	defer func() {
		if err := recover(); err != nil {
			errStr := fmt.Sprintf("No.%0d Test Err:%s\n", id, err)
			fmt.Println(errStr)
		}
	}()
	testCase := TestCases[id-1]
	testCase.result = testCase.f(testCase.params...)
	//fmt.Printf("No.%0d Test Result:\n", id)
	//for _, v := range testCase.result {
	//	fmt.Printf("%+v", v)
	//	fmt.Print(" ")
	//}
}

func RunAll() {
	for i := 1; i <= caseCurId; i++ {
		RunTest(i)
		fmt.Println()
	}
}

func setup() {
	f := func(ps ...interface{}) (ress []interface{}) {
		res := SplitString(ps[0].(string), ps[1].(string))
		ress = append(ress, res)
		return ress
	}
	AddCase("SplitByChar", f, []interface{}{"Hello! Can you split me ?", " "})
	AddCase("SplitByStr", f, []interface{}{"Hello!aaCanaayouaasplitaameaa?", "aa"})
}

// 单元测试
func TestSplitString(t *testing.T) {
	setup()
	RunAll()
}

/*
	PASS
	coverage: 100.0% of statements
	ok      example/unittest        0.146s

*/

// 基准测试
func BenchmarkSplitString(b *testing.B) {
	setup()
	for i := 0; i < b.N; i++ {
		RunTest(1)
	}
}

/*
	goos: windows
	goarch: amd64
	pkg: example/unittest
	BenchmarkSplitString
	BenchmarkSplitString-8   	 3485949	       322 ns/op
	PASS
*/
