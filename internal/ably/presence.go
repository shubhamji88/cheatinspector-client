package ably

import (
    "context"
    "fmt"
)

func UserOnlinePresence(machineID string) {
    fmt.Println("Entering Online presence for machineID:", machineID)
    err := ClientChannel.Presence.EnterClient(context.Background(), machineID, "Online")
    if err != nil {
        fmt.Println("Error entering presence:", err)
    }
}

func UserWritingCodePresence(machineID string) {
    fmt.Println("Updating presence to Coding")
    err := ClientChannel.Presence.UpdateClient(context.Background(),machineID, "Working")
    if err != nil {
        fmt.Println("Error updating presence:", err)
    }
}

func UserLeavePresence() {
    fmt.Println("Leaving Offline presence")
    err := ClientChannel.Presence.Leave(context.Background(), "Offline")
    if err != nil {
        fmt.Println("Error leaving presence:", err)
    }
}
