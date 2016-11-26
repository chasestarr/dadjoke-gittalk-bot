package main

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/http"
)

func inboundRequest(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var inbound InboundPayload
	decoder.Decode(&inbound)

	s := inbound.Text[:8]
	if s == "/dadjoke" {
		outboundRequest(inbound)
	}
	req.Body.Close()
}

func outboundRequest(inbound InboundPayload) {
	var joke Joke
	gorequest.New().Get("http://tambal.azurewebsites.net/joke/random").EndStruct(&joke)

	avatar := "http://assets.nydailynews.com/polopoly_fs/1.1353156!/img/httpImage/image.jpg_gen/derivatives/article_970/25309-324-1-jpg.jpg"
	message := OutboundPayload{"ea2bb9207d881724e2e258b64925c2ed", "message", inbound.Room, Action{Text: joke.Joke, Image: "", Avatar: avatar}}
	gorequest.New().Post("http://localhost:8000/apps").Send(message).End()
}

func main() {
	port := ":9000"
	http.HandleFunc("/", inboundRequest)
	fmt.Println("server running on port " + port)
	http.ListenAndServe(port, nil)
}

type Action struct {
	Text   string `json:"text"`
	Image  string `json:"image"`
	Avatar string `json:"avatar"`
}

type OutboundPayload struct {
	ApiKey string `json:"apiKey"`
	Method string `json:"method"`
	Room   string `json:"room"`
	Action Action `json:"action"`
}

type InboundPayload struct {
	User  string `json:"user"`
	Text  string `json:"text"`
	Image string `json:"image"`
	Room  string `json:"room"`
}

type Joke struct {
	Joke string `json:"joke"`
}
