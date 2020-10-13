# go-exception
golang error异常处理。提供类似java的`try...catch...finally`和`try...with...resource`机制。

# Usage

`try...catch...finally`需要显式调用Throw方法抛出异常才能被外层捕获。

```
tr := exception.NewTrier()
tr.Try(func() {
        // 函数体
	n1, err := strconv.Atoi("123")
	tr.Throw(err)
	n2, err := strconv.Atoi("qwe")
	tr.Throw(err)
	res := n1 / n2
	fmt.Println(res)
}).Catch(func(e Exception) {
        // 捕获异常和异常处理
	fmt.Println("exception:", e)
}).Finally(func() {
        // 事后处理
	fmt.Println("finally")
})
```

`try...with...resource`对象必须实现io.Closer接口，处理结束后实现io自动关闭。

```
f, _ := os.Open("./exception_test.go")
trs := exception.NewTryResource(f)
trs.Try(func() {
        // 函数体
	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		break
	}
}).Catch(func(e Exception) {
        // 捕获异常和异常处理
	fmt.Println("exception:", e)
})
// 处理结束后自动关闭io流
```
