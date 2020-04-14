package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const liverampSidecarHost = "http://localhost"
const liverampSidecarPort = 3000
const liverampTestEnvelope = "AjfowUURXDJnQmc_HNeuswelMv4ZHZQJFM8TpiUnYEyA81Vdgg"

func main() {

	// health endpoint
	fmt.Printf("> GET /health HTTP/1.1\n> Host: %s:%d\n>\n", liverampSidecarHost, liverampSidecarPort)
	resp, err := http.Get(liverampSidecarHost + ":" + strconv.Itoa(liverampSidecarPort) + "/health")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("< HTTP/1.1 %s\n< Content-Length: %d\n\n", string(resp.Status), resp.ContentLength)
	}

	// map endpoint
	fmt.Printf("> GET /map?env=%s HTTP/1.1\n> Host: %s:%d\n>\n", liverampTestEnvelope, liverampSidecarHost, liverampSidecarPort)
	resp, err = http.Get(liverampSidecarHost + ":" + strconv.Itoa(liverampSidecarPort) + "/map?env=" + liverampTestEnvelope)
	if err != nil {
		fmt.Println(err)
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("< HTTP/1.1 %s\n< Content-Type: %s\n<\n%s", string(resp.Status), resp.Header.Get("Content-type"), string(body))
		}
	}

}
