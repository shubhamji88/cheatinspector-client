package ably

import (
    "github.com/ably/ably-go/ably"
    "fmt"
)

var ClientChannel *ably.RealtimeChannel 

func Init(machineID string) {
    fmt.Println("Ably Client Init called")
    var err error
    AblyClient, err := ably.NewRealtime(ably.WithKey("wUARoQ.XNweeg:zu-adaPGzIzaTY1Ko7TW8e1rhXV-5LkQen0CYn0t_ck"))
    if err != nil {
        panic(err)
    }

    AblyClient.Connect()
    ClientChannel = AblyClient.Channels.Get("History_" + machineID)
}
