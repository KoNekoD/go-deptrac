package main

import (
	"flag"
	"log"
	"net/http"
)

//func main() {
//	//color.New(color.FgRed, color.BgGreen).Println("Hello, World!")
//	//
//	//ordered := orderedmap.NewOrderedMap[int, string]()
//	//
//	//ordered.Set(-2, "CCC")
//	//ordered.Set(-1, "AAA")
//	//ordered.Set(10, "BBB")
//	//
//	//keys := ordered.Keys()
//	//slices.Sort(keys)
//	//
//	//for _, key := range keys {
//	//	color.New(color.FgRed, color.BgBlack).Printf("key: %v\n", key)
//	//	v, _ := ordered.Get(key)
//	//	color.New(color.FgRed, color.BgBlack).Printf("value: %v\n", v)
//	//}
//}

func main() {
	port := flag.String("p", "55555", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
