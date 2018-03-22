读取用户的输入

我们如何读取用户的键盘（控制台）输入呢？从键盘和标准输入 os.Stdin 读取输入，最简单的办法是使用 fmt 包提供的 Scan 和 Sscan 开头的函数。请看以下程序：

示例 1：

// 从控制台读取输入:
package main
import "fmt"

var (
   firstName, lastName, s string
   i int
   f float32
   input = "56.12 / 5212 / Go"
   format = "%f / %d / %s"
)

func main() {
   fmt.Println("Please enter your full name: ")
   fmt.Scanln(&firstName, &lastName)
   // fmt.Scanf("%s %s", &firstName, &lastName)
   fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi liu pengjie!
   fmt.Sscanf(input, format, &f, &i, &s)
   fmt.Println("From the string we read: ", f, i, s)
    // 输出结果: From the string we read: 56.12 5212 Go
}
Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
Scanf 与其类似，除了 Scanf 的第一个参数用作格式字符串，用来决定如何读取。
Sscan 和以 Sscan 开头的函数则是从字符串读取，除此之外，与 Scanf 相同。如果这些函数读取到的结果与您预想的不同，您可以检查成功读入数据的个数和返回的错误。

您也可以使用 bufio 包提供的缓冲读取（buffered reader）来读取数据，正如以下例子所示：

示例 2：

package main
import (
    "fmt"
    "bufio"
    "os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
    inputReader = bufio.NewReader(os.Stdin)
    fmt.Println("Please enter some input: ")
    input, err = inputReader.ReadString('\n')
    if err == nil {
        fmt.Printf("The input was: %s\n", input)
    }
}
inputReader 是一个指向 bufio.Reader 的指针。inputReader := bufio.NewReader(os.Stdin) 这行代码，将会创建一个读取器，并将其与标准输入绑定。

bufio.NewReader() 构造函数的签名为：func NewReader(rd io.Reader) *Reader

该函数的实参可以是满足 io.Reader 接口的任意对象（任意包含有适当的 Read() 方法的对象，请参考章节11.8），函数返回一个新的带缓冲的 io.Reader 对象，它将从指定读取器（例如 os.Stdin）读取内容。

返回的读取器对象提供一个方法 ReadString(delim byte)，该方法从输入中读取内容，直到碰到 delim 指定的字符，然后将读取到的内容连同 delim 字符一起放到缓冲区。

ReadString 返回读取到的字符串，如果碰到错误则返回 nil。如果它一直读到文件结束，则返回读取到的字符串和 io.EOF。如果读取过程中没有碰到 delim 字符，将返回错误 err != nil。

在上面的例子中，我们会读取键盘输入，直到回车键（\n）被按下。

屏幕是标准输出 os.Stdout；os.Stderr 用于显示错误信息，大多数情况下等同于 os.Stdout。

一般情况下，我们会省略变量声明，而使用 :=，例如：

inputReader := bufio.NewReader(os.Stdin)
input, err := inputReader.ReadString('\n')
我们将从现在开始使用这种写法。

第二个例子从键盘读取输入，使用了 switch 语句：

示例 3：

package main
import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    inputReader := bufio.NewReader(os.Stdin)
    fmt.Println("Please enter your name:")
    input, err := inputReader.ReadString('\n')

    if err != nil {
        fmt.Println("There were errors reading, exiting program.")
        return
    }

    fmt.Printf("Your name is %s", input)
    // For Unix: test with delimiter "\n", for Windows: test with "\r\n"
    switch input {
    case "Philip\r\n":  fmt.Println("Welcome Philip!")
    case "Chris\r\n":   fmt.Println("Welcome Chris!")
    case "Ivo\r\n":     fmt.Println("Welcome Ivo!")
    default: fmt.Printf("You are not welcome here! Goodbye!\r\n")
    }

    // version 2:   
    switch input {
    case "Philip\r\n":  fallthrough
    case "Ivo\r\n":     fallthrough
    case "Chris\r\n":   fmt.Printf("Welcome %s\n", input)
    default: fmt.Printf("You are not welcome here! Goodbye!\n")
    }

    // version 3:
    switch input {
    case "Philip\r\n", "Ivo\r\n":   fmt.Printf("Welcome %s\n", input)
    default: fmt.Printf("You are not welcome here! Goodbye!\n")
    }
}
注意：Unix和Windows的行结束符是不同的！