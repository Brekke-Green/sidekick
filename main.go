package main

import (
    "log"
    "os"
    "time"
    "fmt"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/gopxl/beep"
    "github.com/gopxl/beep/speaker"
    "github.com/gopxl/beep/mp3"
)

func main() {
    f, err := os.Open("random.mp3")
    if err != nil {
        log.Fatal(err)
    }

    streamer, format, err := mp3.Decode(f)
    if err != nil {
        log.Fatal(err)
    }

    defer streamer.Close()

    speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

    done := make(chan bool)
    speaker.Play(beep.Seq(streamer, beep.Callback(func() {
        done <- true
    })))

    <-done
}
