package main

import (

     "encoding/csv"
     "fmt"
     "io"
     "log"
     "os"
     "strconv"
)

const (
     Id int = 0
     Name int = 1
     Red int = 2
     Green int =3
     Blue int = 4
)
type Colormap struct {
  id int
  r  uint8
  g  uint8
  b  uint8  
}

func string2uint8(str string)uint8 {
     n, _ := strconv.Atoi(str)
     return uint8(n)
}

func string2uint(str string) int {
     n, _ := strconv.Atoi(str)
     return n
}

func show_colormap(colormap map[string]Colormap) {
     fmt.Println("####Show Colomap###")
     for key, value := range colormap {
     	 fmt.Println(key, value.id, value.r, value.g, value.b)
     }
}

func search_name_from_id(colormap map[string]Colormap, id int) string{
     name := ""
     for key, value := range colormap {
     	 if value.id == id {
	    name = key
	 }
     }
     return name
}


func search_id_from_rgb(colormap map[string]Colormap, r,g,b uint8) int{
     id := -1
     for _, value := range colormap {
     	 if value.r == r && value.g == g && value.b == b {
	    id = value.id
	 }
     }
     return id
}

func get_colorlist_from_csv(csvfile string) map[string]Colormap{
     file, err := os.Open(csvfile)
     if err != nil {
	log.Fatal(err)
     }	
     defer file.Close()

     reader := csv.NewReader(file)
     //skip header
     record, _ := reader.Read() // CSV を 1 レコードずつ読み込み
     fmt.Println(record)
     colormap := map[string]Colormap{}
     for {
		record, err := reader.Read() // CSV を 1 レコードずつ読み込み
		if err == io.EOF {
			break // 最後まで読み出した
		}
		if err != nil {
			log.Fatal(err) // 読み出し時にエラー発生
		}

		// 1 レコード分のデータを出力（record は []string 型）
		fmt.Println(record[Id], record[Name], record[Red], record[Green], record[Blue])
		c := Colormap{string2uint(record[Id]), string2uint8(record[Red]), string2uint8(record[Green]), string2uint8(record[Blue])}
		colormap[record[Name]] = c
	}
     return colormap
}
