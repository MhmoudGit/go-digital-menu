package utils

import (
    "fmt"
    "time"
)

func CheckSubscriptions() {
    // Calculate the duration until the next midnight.
    now := time.Now()
    nextMidnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
    if now.After(nextMidnight) {
        nextMidnight = nextMidnight.Add(24 * time.Hour)
    }
    durationUntilMidnight := nextMidnight.Sub(now)

    // Create a timer that triggers at midnight and then repeats every 24 hours.
    timer := time.NewTimer(durationUntilMidnight)
    defer timer.Stop()

    for {
        <-timer.C
        // Your code to run at midnight goes here.
        fmt.Println("Running background task at midnight...")

        // Calculate the duration until the next midnight.
        nextMidnight = nextMidnight.Add(24 * time.Hour)
        durationUntilMidnight = nextMidnight.Sub(now)
        timer.Reset(durationUntilMidnight)
    }
}