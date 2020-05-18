// estimateurPI project estimateurPI.go
package estimateurPI

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type resultWindow []int

type windowDecimal struct {
	gliderWindow func() (resultWindow, error)
	windowSize   int
	targetData   []byte
	loaderData   func(a string) error
}

func (obj *windowDecimal) WriteGlideWindow() error {
	val, err := obj.GlideWindow()
	if err == nil {
		w := csv.NewWriter(os.Stdout)
		w.Comma = ';'
		for _, record := range val {

			if err := w.Write([]string{ /*strconv.Itoa(i),*/ strconv.Itoa(record)}); err != nil {
				return err
			}
		}
		// Write any buffered data to the underlying writer (standard output).
		w.Flush()
		if err := w.Error(); err != nil {
			return err
		}
	}
	return err
}

func (obj *windowDecimal) loadData(a string) error {
	if obj.targetData == nil {
		var err error
		obj.targetData, err = ioutil.ReadFile(a)
		if err != nil {
			panic(fmt.Sprintf("error %v", err))
		}
		return nil
	}
	return errors.New("initialization error")
}

//the function take a slice and iterate over it to count the number of potential A..Z
func glideWindow(a []byte) int {
	sizeString := len(a) - 2
	var resultValue int
	for i := 0; i <= sizeString; i = i + 2 {
		memo := 10*(a[i]-48) + 1*(a[i+1]-48)
		//fmt.Printf("memo %s %d ", string(memo+64), i)
		if (memo > 0) && (memo < 27) { //27
			resultValue++
		}
	}
	//fmt.Printf("result %d \n", resultValue)
	return resultValue
}

func (obj *windowDecimal) GlideWindow() (resultWindow, error) {
	if obj.gliderWindow != nil {

		return obj.gliderWindow()
	} else {
		return nil, errors.New("initialization error")
	}

}

//get a list of value 1 0 2 2 3 2 and compute the numer of seen value 1 1 3 2
func computeCount(a []int) resultWindow {
	p := make(resultWindow, 0)
	for _, value := range a {
		if len(p) > value {
			p[value]++
		} else {
			j := value - len(p) + 1
			p = append(p, make(resultWindow, j)...)
			p[value]++
		}
	}
	return p
}

func Equal(a resultWindow, b resultWindow) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func NewWindowsDecimal(a int, data string) windowDecimal {
	d := windowDecimal{windowSize: a, targetData: []byte(data)}
	p := make([]int, 0)
	d.gliderWindow = func() (resultWindow, error) {
		sizeString := len(d.targetData) - a
		var countWindow int
		for i := 0; i < sizeString; i = i + a {
			//fmt.Printf("count %v %v ", i, d.targetData[i:i+a])
			countWindow = glideWindow(d.targetData[i : i+a])
			p = append(p, countWindow)
		}
		k := computeCount(p)
		return k, nil
	}
	//d.loaderData = d.loadData
	return d
}

func NewWindowsDecimal2(a int, data string) windowDecimal {
	d := windowDecimal{windowSize: a}
	d.loaderData = d.loadData
	_ = d.loaderData(data)
	p := make([]int, 0)
	d.gliderWindow = func() (resultWindow, error) {
		sizeString := len(d.targetData) - a
		var countWindow int
		for i := 0; i < sizeString; i = i + a {
			countWindow = glideWindow(d.targetData[i : i+a])
			p = append(p, countWindow)
		}
		k := computeCount(p)
		return k, nil
	}

	return d
}
