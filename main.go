package main

import (
     "fmt"
     "flag"
     "io/ioutil"
     "path/filepath"
     "os"
)

func do_parallel(src, dst string, colormap map[string]Colormap, para int) {
     _ = os.Mkdir(dst, 0777)
     files, err := ioutil.ReadDir(src)
     if err != nil {
        panic(err)
     }
     ch := make(chan int, para)     
     for p := 0 ; p < para ; p++ {
     	 go func(p int) {
	    count := 0
	    for i := p; i < len(files); i += para {
	    	name := files[i].Name()
		filename := filepath.Join(src, name)
//		fmt.Println(filename, len(files), p)
		img := read_image(filename)
		gray := color2gray(img, colormap, filename)
		dstPath := filepath.Join(dst, name)
		write_image(dstPath, gray)
	    }
	    ch <- count
	 }(p)
     }
     for p := 0 ; p < para ; p++ {
      	 <-ch
     }
}


func main() {
    var (	    
	input = flag.String("i", "sample", "string flag")
	output = flag.String("o", "dst", "string flag")	
	csv = flag.String("csv", "BDD100k_lane.csv", "string flag")

    )

    flag.Parse()
    colormap :=  get_colorlist_from_csv(*csv)
    show_colormap(colormap)
    fmt.Println(search_name_from_id(colormap, 3))
    fmt.Println(search_name_from_id(colormap, 7))
    fmt.Println(search_id_from_rgb(colormap, 219, 94, 86))
    fmt.Println(search_id_from_rgb(colormap, 86, 219, 127))
    fmt.Println(search_id_from_rgb(colormap, 127, 127, 127))                

    fmt.Println(*input)
    do_parallel(*input, *output, colormap, 12)
}

