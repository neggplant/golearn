# Constants
```go
const (
		retryLimit = 4
		httpMethod = "GET"
	)
```
## 常量不能被修改
## const 声明的常量必须在编译时能够确定其值
```go
var a = math.Sqrt(4) // 允许，因为 var 声明的变量可以在运行时计算
fmt.Println(a)
const b = math.Sqrt(4) // 不允许，因为 const 声明的常量必须在编译时确定值
```
```go
// 使用反射，灵活构建函数
// reflect.TypeOf：获取变量的类型。
// reflect.ValueOf：获取变量的值。
func Add(a, b int) int {
	return a + b
}

func main() {
	fn := reflect.ValueOf(Add)
	args := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}

	result := fn.Call(args)
	fmt.Println("Result of Add(10, 20):", result[0].Int())
}
```
```go
func main() {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	// 定义标签 randloop
randloop:
	for {
		// 生成一个随机数 i
		switch i := rand.Intn(100); {
		// 检查 i 是否为偶数
		case i%2 == 0:
			fmt.Printf("Generated even number %d\n", i)
			// 使用带标签的 break 语句退出循环
			break randloop
		}
	}
}
```
```go
	// 满足条件继续判断下个条件
	switch num := number(); { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
```

# Arrays and Slices

```go
// cap是slice从数组开头到数据结尾的长度，len是slice的长度，len<=cap,

arr1 := [5]int{1, 2, 3, 4, 5}
// s1 := arr1[1:2] // 切片容量足够时不会分配新的数组，而是直接修改原数组。
s1 := arr1[4:5] // 切片 s1 不再引用原数组 arr1 的底层数组，而是引用新的底层数组。
s1 = append(s1, 7) // 追加一个元素，根据切片的容量，是否会分配一个新的底层数组。

// 当装不下，cap乘2扩大
// append一个一个元素，使用...添加。类似Python的**kwargs
append(s2[2:3], s2[4:]...)
make([]int, 4, 5)  // 4：表示切片的初始长度（length）,5：表示切片的容量（capacity）
```
# function
```go
// 可变参数s为一个[]string的slice
func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func main() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}
```
```go
	ticketPrice := 0
	if age := 10; age < 5 {
		ticketPrice = 0
	} else if age >= 5 && age <= 22 {
		ticketPrice = 10
	} else {
		ticketPrice = 15
	}
	fmt.Printf("Ticket price is $%d", ticketPrice)
```
# map
```go
make(map[string]string)
// Checking if a key exists
value, ok := map[key]
if currencyName, ok := currencyCode[cyCode]; ok {
		return
	}
// Iterate
for code, name := range currencyCode {
		fmt.Printf("Currency Name for currency code %s is %s\n", code, name)
	}
// del
delete(currencyCode, "EUR")
// Map of structs
type currency struct {
	name   string
	symbol string
}

func main() {
	curUSD := currency{
		name:   "US Dollar",
		symbol: "$",
	}

	currencyCode := map[string]currency{
		"USD": curUSD,
	}

	for cyCode, cyInfo := range currencyCode {
		fmt.Printf("Currency Code: %s, Name: %s, Symbol: %s\n", cyCode, cyInfo.name, cyInfo.symbol)
	}
```

```go

func dd(a map[string]int) {
	a["fddf"] = 3334
}

func main() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}
	fmt.Println("Original employee salary", employeeSalary)
	modified := employeeSalary
	modified["mike"] = 18000
	dd(employeeSalary)
	fmt.Println("Employee salary changed", employeeSalary)
// 不修改原map的值，只能生成副本 ******************
// 创建映射副本的函数
func copyMap(original map[string]int) map[string]int {
    copy := make(map[string]int)
    for key, value := range original {
        copy[key] = value
    }
    return copy
}
```
# string
```go
//  %x is the format specifier for hexadecimal
func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func main() {
	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printBytes(name)
}
// String: Hello World
// Bytes: 48 65 6c 6c 6f 20 57 6f 72 6c 64  
```

```go

```

```go

```

```go

```

```go

```

