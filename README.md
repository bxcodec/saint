# Docs 

## [saint](#) 


[![License](https://img.shields.io/badge/status-on%20going-yellowgreen.svg)](#)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/bxcodec/saint/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/bxcodec/saint?status.svg)](https://godoc.org/github.com/bxcodec/saint)

Go (Golang) Math Library for Integer Operations

Any Math operations that not support integer, can used this library

## Index

* [Support](#support)
* [Getting Started](#getting-started)	
* [Example](#example)	


## Support


You can also email <iman.tumorang@gmail.com> or file an [Issue](https://github.com/bxcodec/saint/issues/new).
See documentation in [Godoc](https://godoc.org/github.com/bxcodec/saint)




## Getting Started

#### Download

```shell
go get -u github.com/bxcodec/saint
```
## Example

```go

package main

import (
	"fmt"
	"github.com/bxcodec/saint"
)

func main() {
	 	
    
   arr:=[] int {2,1,3,5,6}
   var x int 

   x= saint.Max(arr...) 
   fmt.Println(x) // 6

   x= saint.Max(4,3,1,5,7) 
   fmt.Println(x) // 7

   x= saint.Min(arr...) 
   fmt.Println(x) // 1

   x= saint.Min(4,3,5,5,7) 
   fmt.Println(x) // 3

   x= saint.Sum(arr...) 
   fmt.Println(x) // 17

   x = saint.Sum(4,5)
   fmt.Println(x) // 9
}


```