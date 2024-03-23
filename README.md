# js

string.slice() in go just like js/javascript!

# sliceutil

golang sliceutil, including methods like TakeWhile, Filter, ForEach and so on, which were inpired by js 

similar: [github.com/conbanwa/slice](https://github.com/conbanwa/slice)

# examples

```go
package main

import (
	"github.com/conbanwa/js"
)

func main() {
	println(js.SubString(`1234567890`, 2))       //34567890
	println(js.SubString(`1234567890`, -2))      //90
	println(js.SubString(`1234567890`, 0, -4))   //123456
	println(js.SubString(`1234567890`, 2, -4))   //3456
	println(js.SubString(`1234567890`, 6, 99))   //7890
	println(js.SubString(`1234567890`, -5, 4))   //empty
	println(js.SubString(`1234567890`, -5, 99))  //67890
	println(js.SubString(`1234567890`, -99, 99)) //1234567890

	println(js.Slice([]byte(`1234567890`), 2))       //34567890
	println(js.Slice([]byte(`1234567890`), -2))      //90
	println(js.Slice([]byte(`1234567890`), 0, -4))   //123456
	println(js.Slice([]byte(`1234567890`), 2, -4))   //3456
	println(js.Slice([]byte(`1234567890`), 6, 99))   //7890
	println(js.Slice([]byte(`1234567890`), -5, 4))   //empty
	println(js.Slice([]byte(`1234567890`), -5, 99))  //67890
	println(js.Slice([]byte(`1234567890`), -99, 99)) //1234567890
}
```

# install

import github.com/conbanwa/js
