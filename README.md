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

    dalamKalimat := masa.New(gotime).Convert("Kamaren pas bulan MMMM kan yak?")
    fmt.Println(dalamKalimat) // This will print: Kamaren pas bulan Desember kan yak?
}
```

# Available Formats
|         | Token | Output                                          |
| ------: | ----- | ----------------------------------------------- |
|    Hari | dd    | Mg Sn Sl Rb Km Jm Sb                            |
|         | ddd   | Min Sen Sel Rab Kam Jum Sab                     |
|         | dddd  | Minggu Senin Selasa Rabu Kamis Jumat Sabtu      |
|   Bulan | M     | 1 2 3 ... 11 12                                 |
|         | MM    | 01 02 03 ... 11 12                              |
|         | MMM   | Jan Feb Mar Apr Mei Jun Jul Agt Sep Okt Nov Des |
|         | MMMM  | Januari Februari Maret ... November Desember    |
