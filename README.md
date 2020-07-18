# go-exception
golang error异常处理，提供类似java的`try...catch...finally`和`try...with...resource‘机制。

# Usage

`try...catch...finally`提供链式调用，需要显式调用Throw方法抛出异常才能被外层捕获。

```golang
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
```

`try...with...resource‘对象必须实现io.Closer接口，处理结束后实现io自动关闭，隐藏细节。

```golang
    f, _ := os.Open("./exception_test.go")
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
```
