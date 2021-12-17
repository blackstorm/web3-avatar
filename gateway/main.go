package main

import (
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

var web3 *Proxy

func init() {
	err := godotenv.Load()
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}

	w, err := NewProxy(os.Getenv("INFURA_TOKEN"))
	if err != nil {
		panic(err)
	}

	web3 = w
}

func handle(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	addr := params.ByName("address")
	name := params.ByName("name")
	if name == "" {
		name = "default"
	}

	avatar, err := web3.GetAvatar(addr, name)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	if avatar == "" {
		w.WriteHeader(404)
		return
	}

	w.Header().Add("Location", "https://cloudflare-ipfs.com/ipfs/"+avatar)
	w.WriteHeader(302)
}

func main() {
	r := httprouter.New()
	r.GET("/v1/:address", handle)
	r.GET("/v1/:address/:name", handle)

	log.Fatalln(http.ListenAndServe(":8080", r))
}
