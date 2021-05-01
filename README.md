# masa ![GO](https://github.com/mhdiiilham/masa/actions/workflows/go.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/mhdiiilham/masa)](https://goreportcard.com/report/github.com/mhdiiilham/masa)

*This package was inspired by* **[masa](https://github.com/armedi/masa)**

Ini merupakan package go sederhana yang berfungsi untuk membantu memformat type `time.Time` ke dalam Bahasa Indonesia.

# How To Use
```
go get github.com/mhdiiilham/masa
```

```go
package main

import (
    "time.Time"
    "github.com/mhdiiilham/masa"
)

func main() {
    goTime, _ := time.Parse("2006-01-02", "2000-12-15")
    hari := masa.New(goTime).Convert(tc.arg).String()
    fmt.Println(hari) // This will print: Jumat
}
```

# Available Formats
|         | Token | Output                                          |
| ------: | ----- | ----------------------------------------------- |
|    Hari | dd    | Mg Sn Sl Rb Km Jm Sb                            |
|         | ddd   | Min Sen Sel Rab Kam Jum Sab                     |
|         | dddd  | Minggu Senin Selasa Rabu Kamis Jumat Sabtu      |
