package model

// Item is data that will be stored about a websocket connection in dynamoDB
type Item struct {
	ConnectionID string `json:"connectionId"`
}

// Message is a struct representing data sent to and from a client
type Message struct {
	Data string `json:"data"`
}
