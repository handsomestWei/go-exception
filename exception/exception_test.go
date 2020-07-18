package exception

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

// test 1st throw
func TestTrier_Try_1(t *testing.T) {
	tr := NewTrier()
	tr.Try(func() {
		n1, err := strconv.Atoi("qwe")
		tr.Throw(err)
		n2, err := strconv.Atoi("0")
		tr.Throw(err)
		res := n1 / n2
		fmt.Println(res)
	}).Catch(func(e Exception) {
		fmt.Println("exception:", e)
	}).Finally(func() {
		fmt.Println("finally")
	})
}

// test 2nd throw
func TestTrier_Try_2(t *testing.T) {
	tr := NewTrier()
	tr.Try(func() {
		n1, err := strconv.Atoi("123")
		tr.Throw(err)
		n2, err := strconv.Atoi("qwe")
		tr.Throw(err)
		res := n1 / n2
		fmt.Println(res)
	}).Catch(func(e Exception) {
		fmt.Println("exception:", e)
	}).Finally(func() {
		fmt.Println("finally")
	})
}

// test panic
func TestTrier_Try_3(t *testing.T) {
	tr := NewTrier()
	tr.Try(func() {
		n1, err := strconv.Atoi("123")
		tr.Throw(err)
		panic("test panic")
		n2, err := strconv.Atoi("1")
		tr.Throw(err)
		res := n1 / n2
		fmt.Println(res)
	}).Catch(func(e Exception) {
		fmt.Println("exception:", e)
	})
}

// test pass
func TestTrier_Try_4(t *testing.T) {
	tr := NewTrier()
	tr.Try(func() {
		n1, err := strconv.Atoi("123")
		tr.Throw(err)
		n2, err := strconv.Atoi("1")
		tr.Throw(err)
		res := n1 / n2
		fmt.Println(res)
	}).Catch(func(e Exception) {
		fmt.Println("exception:", e)
	})
}

// test auto closed
func TestTryResource_Try_1(t *testing.T) {
	f, err := os.Open("./exception_test.go")
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	var scanner *bufio.Scanner

	trs := NewTryResource(f)
	trs.Try(func() {
		scanner = bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			break
		}
	}).Catch(func(e Exception) {
		fmt.Println("exception:", e)
	})

	_, err = f.Read(make([]byte, 1024))
	if err != nil {
		t.Error(err)
	}
}

// test throw error
func TestTryResource_Try_2(t *testing.T) {
	f, err := os.Open("./exception_test.go")
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	var scanner *bufio.Scanner

	trs := NewTryResource(f)
	trs.Try(func() {
		scanner = bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			_, err := strconv.Atoi("123a")
			trs.Throw(err)
		}
	}).Catch(func(e Exception) {
		fmt.Println("exception:", e)
	})

	_, err = f.Read(make([]byte, 1024))
	if err != nil {
		t.Error(err)
	}
}

// test not auto closed
func TestTryResource_Try_3(t *testing.T) {
	f, err := os.Open("./exception_test.go")
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	var scanner *bufio.Scanner

	trs := NewTryResource(nil)
	trs.Try(func() {
		scanner = bufio.NewScanner(f)
	}).Catch(func(e Exception) {
		fmt.Println("exception:", e)
	})

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		break
	}
	f.Close()
	_, err = f.Read(make([]byte, 1024))
	if err != nil {
		t.Error(err)
	}
}
