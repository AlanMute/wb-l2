Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
Так как мы передаем лишь копию слайса, при добавлении нового элемента в этот слайс, для него выделяется новая память и она уже будет указывать на нее. От чего исходный не изменится.
...
Output:
[3, 2, 3]

```
