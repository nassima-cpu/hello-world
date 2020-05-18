// estimateurPI_test
package estimateurPI

import (
	//	"fmt"
	"testing"
)

func TestGlideWindow(t *testing.T) {
	total := glideWindow([]byte("0126011"))
	if total != 3 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 3)
	}
	total = glideWindow([]byte("010101"))
	t.Logf("count %d", total)
	if total != 3 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 3)
	}

}

func TestNewWindowsDecimal(t *testing.T) {
	/*
		total := NewWindowsDecimal(5, "012601101260110126011")
		res := resultWindow{0, 2, 2}
		out, _ := total.gliderWindow()
		if !Equal(res, out) {
			t.Errorf("gliderWindow was incorrect, got: %+v, want: %+v.", out, res)
		}
	*/
	//fmt.Println(out)

	//t.Logf("count %+v\n", total)
	//t.Logf("count %+v\n", out)
}

func TestNewWindowsDecimal2(t *testing.T) {

}

func TestNewWindowsDecimal4(t *testing.T) {

}

func TestNewWindowsDecimal3(t *testing.T) {

}

func TestEqual(t *testing.T) {
	a := resultWindow{1, 5, 1, 1, 2, 1, 2}
	b := resultWindow{1, 5, 1, 1, 2, 1, 2}
	if !Equal(a, b) {
		t.Errorf("function Equal is not working - different value case \n")
	}
}

func TestComputeCount(t *testing.T) {
	a := []int{1, 5, 1, 0, 1, 2, 1, 2}
	aResult := resultWindow{1, 4, 2, 0, 0, 1}
	aCompute := computeCount(a)
	t.Logf("count %+v\n", aCompute)
	if !Equal(aCompute, aResult) {
		t.Errorf("Error in output %+v\n", aCompute)
	}
}

var result int

//run a fixed size
func BenchmarkGlideWindow(b *testing.B) {
	// run the glideWindow function b.N times
	var total int
	for n := 0; n < b.N; n++ {
		total = glideWindow([]byte("010101"))
	}
	result = total
}

var slice []byte

func benchmarkGlideWindow2(i int, b *testing.B) {
	// run the glideWindow function b.N times
	for n := 0; n < b.N; n++ {
		//slice := make([]byte, i)
		glideWindow(slice)
	}
}

func BenchmarkGlideWindow3(b *testing.B) {
	slice = make([]byte, 2000)
	benchmarkGlideWindow2(2000, b)
}
