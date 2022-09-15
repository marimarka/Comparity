package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/barbora", returnBarboraProduct)
	myRouter.HandleFunc("/rimi", returnRimiProduct)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnBarboraProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: barbora url")

	responseBarbora, err := http.Get("https://www.barbora.lv/produkti/baltmaize-zeltene-350-g")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseBarboraData, err := io.ReadAll(responseBarbora.Body)

	if err != nil {
		log.Fatal(err)
	}
	responseBarboraText := string(responseBarboraData)

	reader := strings.NewReader(responseBarboraText)
	tokenizer := html.NewTokenizer(reader)
	for {
		tt := tokenizer.Next()
		if tt == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				return
			}
			fmt.Printf("Error: %v", tokenizer.Err())
			return
		}
		for {
			attrKey, attrValue, moreAttr := tokenizer.TagAttr()
			if string(attrKey) == "itemprop" && string(attrValue) == "price" {
				if moreAttr {
					attrPriceKey, attrPriceValue, morePriceAttr := tokenizer.TagAttr()
					if string(attrPriceKey) == "content" {
						fmt.Println(string(attrPriceValue))
						fmt.Fprintf(w, "Zeltene, Barbora, cena: "+string(attrPriceValue))
					}
					if !morePriceAttr {
						break
					}
				}
				break
			}

			if !moreAttr {
				break
			}
		}
	}
}

func returnRimiProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Rimi url")

	responseRimi, err := http.Get("https://www.rimi.lv/e-veikals/lv/produkti/maize-un-konditoreja/maize/baltmaize/baltmaize-zeltene-salda-sagriezta-350g/p/285134")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseRimiData, err := io.ReadAll(responseRimi.Body)

	if err != nil {
		log.Fatal(err)
	}
	responseRimiText := string(responseRimiData)
	priceIndex := strings.Index(responseRimiText, "\"price\":") + 8
	stringLastPart := responseRimiText[priceIndex:]
	priceNumberIndex := strings.Index(stringLastPart, ",")
	price := stringLastPart[0:priceNumberIndex]
	fmt.Fprintf(w, "Zeltene, Rimi, cena: "+string(price))
}

func main() {

	fmt.Println("Rest API v1.0. Marina.")

	handleRequests()
}
