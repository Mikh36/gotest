package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	log.Println("Read config")
	var cfg map[string]string = map[string]string{"port": "8080"} // Обьявлем и иницилизируем переменную с дефолтным портом.
	cleanenv.ReadConfig("config.yaml", &cfg)                      // Если есть файл с конфином берем значение порта оттуда.

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if pair[0] == "APP_PORT" {
			log.Printf(fmt.Sprintf("Set the application port from environment variables in the value: %s", pair[1]))
			cfg["port"] = pair[1] // Если есть переменная окружения берем порт их окружения
		}
	}

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Printf(fmt.Sprintf("We use port: %s", cfg["port"]), router)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg["port"]), router))
}
