package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-geojson-svg"
	"log"
	"os"
	"io/ioutil"
)

func main() {

	height := flag.Float64("height", 640., "")
	width := flag.Float64("width", 1024., "")	
	mercator := flag.Bool("mercator", false, "")
	
	flag.Parse()

	s := svg.New()
	s.Mercator = *mercator

	features := flag.Args()

	for _, path := range features {

		fh, err := os.Open(path)

		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(fh)

		if err != nil {
			log.Fatal(err)
		}

		err = s.AddFeature(string(body))

		if err != nil {
			log.Fatal(err)
		}		
	}

	out := s.Draw(*width, *height,
		svg.WithAttribute("xmlns", "http://www.w3.org/2000/svg"),
		svg.WithAttribute("viewBox", fmt.Sprintf("0 0 %d %d", int(*width), int(*height))),
	);
	
	fmt.Println(out)
}
