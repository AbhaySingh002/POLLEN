package orchestrator

import (
    "context"
    "encoding/xml"
    // "fmt"
    "io"
    "log"
    "strings"
    "sync"

    "github.com/AbhaySingh002/Pollen/internal/ai"
    "github.com/AbhaySingh002/Pollen/internal/parser"
    "github.com/AbhaySingh002/Pollen/internal/types"
)


func Orchestrator(prompt string, isCoding bool) {
    if !isCoding {
        log.Print("🧠 Clarifying the Prompt :")
    } else {
        log.Print("🧠 Vibe coding your app : ")
    }

	// taskChan := make(chan types.Task, 10)
    var wg sync.WaitGroup
    ctx := context.Background()
    streamchan := make(chan string)
    isIntent := !isCoding
    var thoughtChan chan string
    if isCoding {
        thoughtChan = make(chan string)
        wg.Add(1)
        go func() {
            defer wg.Done()
            parser.ParseThoughts(thoughtChan)
        }()
    }

    wg.Add(1)
    go func() {
        defer wg.Done()
        ai.AiCaller(ctx, prompt, streamchan, thoughtChan, isIntent)
    }()

    if isCoding {
        pr, pw := io.Pipe()

        // Goroutine: Consume streamchan, print, buffer, and write to pipe
        wg.Add(1)
        go func() {
            defer wg.Done()
            defer pw.Close()
            for msg := range streamchan {
                // log.Printf("🔄 Stream chunk received: %d bytes", len(msg))
                // if len(msg) > 0 && len(msg) < 200 {
                //     log.Printf("📨 Content preview: %s", msg)
                // }
                if _, err := pw.Write([]byte(msg)); err != nil {
                    log.Printf("Pipe write error: %v", err)
                    return
                }
            }
        }()
		wg.Add(1)
        go func() {
            defer wg.Done()
            parser.ParseStreamFromReader(pr)
        }()

        wg.Wait()

    } else {
        
        var xmlBuffer strings.Builder
        for msg := range streamchan {
            xmlBuffer.WriteString(msg)
        }

        var intent types.IntentCheckingResponse // Assuming this type exists in types
        if xmlBuffer.Len() > 0 {
            if err := xml.Unmarshal([]byte(xmlBuffer.String()), &intent); err != nil {
                log.Printf("Intent XML unmarshal error: %v", err)
                return
            }
            intent.IsDone = true 
        }
        if intent.IsSafe { 
            log.Println("☑️ Passed the Intent checking")
            Orchestrator(prompt, true) // Recursive call for coding
        } else {
            log.Printf("❌ %s", intent.Reason)
        }
    }

    // Final wait if needed (for intent recursion)
    if !isCoding {
        wg.Wait()
    }
}