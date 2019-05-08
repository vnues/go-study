```go
// Print 将参数列表 a 中的各个参数转换为字符串并写入到标准输出中。
// 非字符串参数之间会添加空格，返回写入的字节数。
func Print(a ...interface{}) (n int, err error)

// Println 功能类似 Print，只不过最后会添加一个换行符。
// 所有参数之间会添加空格，返回写入的字节数。
func Println(a ...interface{}) (n int, err error)

// Printf 将参数列表 a 填写到格式字符串 format 的占位符中。
// 填写后的结果写入到标准输出中，返回写入的字节数。
func Printf(format string, a ...interface{}) (n int, err error)

------------------------------

// 功能同上面三个函数，只不过将转换结果写入到 w 中。
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)

------------------------------

// 功能同上面三个函数，只不过将转换结果以字符串形式返回。
func Sprint(a ...interface{}) string
func Sprintln(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string

------------------------------

// 功能同 Sprintf，只不过结果字符串被包装成了 error 类型。
func Errorf(format string, a ...interface{}) error

------------------------------

// 示例
func main() {
	fmt.Print("a", "b", 1, 2, 3, "c", "d", "\n")
	fmt.Println("a", "b", 1, 2, 3, "c", "d")
	fmt.Printf("ab %d %d %d cd\n", 1, 2, 3)
	// ab1 2 3cd
	// a b 1 2 3 c d
	// ab 1 2 3 cd

	if err := percent(30, 70, 90, 160); err != nil {
		fmt.Println(err)
	}
	// 30%
	// 70%
	// 90%
	// 数值 160 超出范围（100）
}

func percent(i ...int) error {
	for _, n := range i {
		if n > 100 {
			return fmt.Errorf("数值 %d 超出范围（100）", n)
		}
		fmt.Print(n, "%\n")
	}
	return nil
}

------------------------------------------------------------

// Formatter 由自定义类型实现，用于实现该类型的自定义格式化过程。
// 当格式化器需要格式化该类型的变量时，会调用其 Format 方法。
type Formatter interface {
	// f 用于获取占位符的旗标、宽度、精度等信息，也用于输出格式化的结果
	// c 是占位符中的动词
	Format(f State, c rune)
}

// 由格式化器（Print 之类的函数）实现，用于给自定义格式化过程提供信息
type State interface {
	// Formatter 通过 Write 方法将格式化结果写入格式化器中，以便输出。
	Write(b []byte) (ret int, err error)
	// Formatter 通过 Width 方法获取占位符中的宽度信息及其是否被设置。
	Width() (wid int, ok bool)
	// Formatter 通过 Precision 方法获取占位符中的精度信息及其是否被设置。
	Precision() (prec int, ok bool)
	// Formatter 通过 Flag 方法获取占位符中的旗标[+- 0#]是否被设置。
	Flag(c int) bool
}

// Stringer 由自定义类型实现，用于实现该类型的自定义格式化过程。
// 当格式化器需要输出该类型的字符串格式时就会调用其 String 方法。
type Stringer interface {
	String() string
}

// Stringer 由自定义类型实现，用于实现该类型的自定义格式化过程。
// 当格式化器需要输出该类型的 Go 语法字符串（%#v）时就会调用其 String 方法。
type GoStringer interface {
	GoString() string
}

------------------------------

// 示例
type Ustr string

func (us Ustr) String() string {
	return strings.ToUpper(string(us))
}

func (us Ustr) GoString() string {
	return `"` + strings.ToUpper(string(us)) + `"`
}

func (u Ustr) Format(f fmt.State, c rune) {
	write := func(s string) {
		f.Write([]byte(s))
	}
	switch c {
	case 'm', 'M':
		write("旗标：[")
		for s := "+- 0#"; len(s) > 0; s = s[1:] {
			if f.Flag(int(s[0])) {
				write(s[:1])
			}
		}
		write("]")
		if v, ok := f.Width(); ok {
			write(" | 宽度：" + strconv.FormatInt(int64(v), 10))
		}
		if v, ok := f.Precision(); ok {
			write(" | 精度：" + strconv.FormatInt(int64(v), 10))
		}
	case 's', 'v': // 如果使用 Format 函数，则必须自己处理所有格式，包括 %#v
		if c == 'v' && f.Flag('#') {
			write(u.GoString())
		} else {
			write(u.String())
		}
	default: // 如果使用 Format 函数，则必须自己处理默认输出
		write("无效格式：" + string(c))
	}
}

func main() {
	u := Ustr("Hello World!")
	// "-" 标记和 "0" 标记不能同时存在
	fmt.Printf("%-+ 0#8.5m\n", u) // 旗标：[+- #] | 宽度：8 | 精度：5
	fmt.Printf("%+ 0#8.5M\n", u)  // 旗标：[+ 0#] | 宽度：8 | 精度：5
	fmt.Println(u)                // HELLO WORLD!
	fmt.Printf("%s\n", u)         // HELLO WORLD!
	fmt.Printf("%#v\n", u)        // "HELLO WORLD!"
	fmt.Printf("%d\n", u)         // 无效格式：d
}

------------------------------------------------------------

// Scan 从标准输入中读取数据，并将数据用空白分割并解析后存入 a 提供
// 的变量中（换行符会被当作空白处理），变量必须以指针传入。
// 当读到 EOF 或所有变量都填写完毕则停止扫描。
// 返回成功解析的参数数量。
func Scan(a ...interface{}) (n int, err error)

// Scanln 和 Scan 类似，只不过遇到换行符就停止扫描。
func Scanln(a ...interface{}) (n int, err error)

// Scanf 从标准输入中读取数据，并根据格式字符串 format 对数据进行解析，
// 将解析结果存入参数 a 所提供的变量中，变量必须以指针传入。
// 输入端的换行符必须和 format 中的换行符相对应（如果格式字符串中有换行
// 符，则输入端必须输入相应的换行符）。
// 占位符 %c 总是匹配下一个字符，包括空白，比如空格符、制表符、换行符。
// 返回成功解析的参数数量。
func Scanf(format string, a ...interface{}) (n int, err error)

// 功能同上面三个函数，只不过从 r 中读取数据。
func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)

// 功能同上面三个函数，只不过从 str 中读取数据。
func Sscan(str string, a ...interface{}) (n int, err error)
func Sscanln(str string, a ...interface{}) (n int, err error)
func Sscanf(str string, format string, a ...interface{}) (n int, err error)

------------------------------

// 示例

// 对于 Scan 而言，回车视为空白
func main() {
	a, b, c := "", 0, false
	fmt.Scan(&a, &b, &c)
	fmt.Println(a, b, c)
	// 在终端执行后，输入 abc 1 回车 true 回车
	// 结果 abc 1 true
}

// 对于 Scanln 而言，回车结束扫描
func main() {
	a, b, c := "", 0, false
	fmt.Scanln(&a, &b, &c)
	fmt.Println(a, b, c)
	// 在终端执行后，输入 abc 1 true 回车
	// 结果 abc 1 true
}

// 格式字符串可以指定宽度
func main() {
	a, b, c := "", 0, false
	fmt.Scanf("%4s%d%t", &a, &b, &c)
	fmt.Println(a, b, c)
	// 在终端执行后，输入 1234567true 回车
	// 结果 1234 567 true
}

------------------------------------------------------------

// Scanner 由自定义类型实现，用于实现该类型的自定义扫描过程。
// 当扫描器需要解析该类型的数据时，会调用其 Scan 方法。
type Scanner interface {
	// state 用于获取占位符中的宽度信息，也用于从扫描器中读取数据进行解析。
	// verb 是占位符中的动词
	Scan(state ScanState, verb rune) error
}

// 由扫描器（Scan 之类的函数）实现，用于给自定义扫描过程提供数据和信息。
type ScanState interface {
	// ReadRune 从扫描器中读取一个字符，如果用在 Scanln 类的扫描器中，
	// 则该方法会在读到第一个换行符之后或读到指定宽度之后返回 EOF。
	// 返回“读取的字符”和“字符编码所占用的字节数”
	ReadRune() (r rune, size int, err error)
	// UnreadRune 撤消最后一次的 ReadRune 操作，
	// 使下次的 ReadRune 操作得到与前一次 ReadRune 相同的结果。
	UnreadRune() error
	// SkipSpace 为 Scan 方法提供跳过开头空白的能力。
	// 根据扫描器的不同（Scan 或 Scanln）决定是否跳过换行符。
	SkipSpace()
	// Token 用于从扫描器中读取符合要求的字符串，
	// Token 从扫描器中读取连续的符合 f(c) 的字符 c，准备解析。
	// 如果 f 为 nil，则使用 !unicode.IsSpace(c) 代替 f(c)。
	// skipSpace：是否跳过开头的连续空白。返回读取到的数据。
	// 注意：token 指向共享的数据，下次的 Token 操作可能会覆盖本次的结果。
	Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
	// Width 返回占位符中的宽度值以及宽度值是否被设置
	Width() (wid int, ok bool)
	// 因为上面实现了 ReadRune 方法，所以 Read 方法永远不应该被调用。
	// 一个好的 ScanState 应该让 Read 直接返回相应的错误信息。
	Read(buf []byte) (n int, err error)
}

------------------------------

// 示例
type Ustr string

func (u *Ustr) Scan(state fmt.ScanState, verb rune) (err error) {
	var s []byte
	switch verb {
	case 'S':
		s, err = state.Token(true, func(c rune) bool { return 'A' <= c && c <= 'Z' })
		if err != nil {
			return
		}
	case 's', 'v':
		s, err = state.Token(true, func(c rune) bool { return 'a' <= c && c <= 'z' })
		if err != nil {
			return
		}
	default:
		return fmt.Errorf("无效格式：%c", verb)
	}
	*u = Ustr(s)
	return nil
}

func main() {
	var a, b, c, d, e Ustr
	n, err := fmt.Scanf("%3S%S%3s%2v%x", &a, &b, &c, &d, &e)
	fmt.Println(a, b, c, d, e)
	fmt.Println(n, err)
	// 在终端执行后，输入 ABCDEFGabcdefg 回车
	// 结果：
	// ABC DEFG abc de
	// 4 无效格式：x
}
```
