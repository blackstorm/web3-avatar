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

	gatewayParam := r.URL.Query().Get("gateway")
	gateway := DEFAULT
	switch gatewayParam {
	case "cloudflare":
		gateway = CLOUDFLARE
	case "ipfs":
		gateway = IPFS_IO
	case "dweb":
		gateway = DWEB_LINK
	case "infura":
		gateway = INFURA_IO
	default:
		gateway = DEFAULT
	}

	accessControlHeaders(w)
	w.Header().Add("Location", gateway+avatar)
	w.WriteHeader(302)
}

func main() {
	r := httprouter.New()
	r.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		accessControlHeaders(w)
	})
	r.GET("/v1/:address", handle)
	r.GET("/v1/:address/:name", handle)

	log.Fatalln(http.ListenAndServe(":8080", r))
}

func accessControlHeaders(w http.ResponseWriter) {
	header := w.Header()
	header.Set("Access-Control-Allow-Methods", "GET")
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Allow-Headers", "*")
}
