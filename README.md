# Docs 

## [intregers](#) 


[![License](https://img.shields.io/badge/status-on%20going-yellowgreen.svg)](#)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/bxcodec/intregers/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/bxcodec/intregers?status.svg)](https://godoc.org/github.com/bxcodec/intregers)

Go (Golang) Math Library for Integer Operations

Any Math operations that not support integer, can used this library

## Index

* [Support](#support)
* [Getting Started](#getting-started)	
* [Example](#example)	


## Support


You can also email <iman.tumorang@gmail.com> or file an [Issue](https://github.com/bxcodec/intregers/issues/new).



## Getting Started

#### Download

```shell
go get -u github.com/bxcodec/intregers
```
## Example

```go

package main

import (
	"fmt"
	"github.com/bxcodec/intregers"
)

func main() {
	 	
    
   arr:=[] int {2,1,3,5,6}
   var x int 
   x= intregers.Max(arr...) 

   fmt.Println(x) // 6

   x= intregers.Min(arr...) 

   fmt.Println(x) // 1

   x= intregers.Sum(arr...) 

   fmt.Println(x) // 17

   x = intregers.Sum(4,5)
   fmt.Println(x) // 9
}


```