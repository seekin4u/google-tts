package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go \"<text to synthesize>\"")
		os.Exit(1)
	}

	text := os.Args[1]

	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	req := &texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "ru",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, req)
	if err != nil {
		log.Fatalf("failed to synthesize speech: %v", err)
	}

	// Play audio directly using mpg123 (must be installed)
	cmd := exec.Command("mpg123", "-") // "-" = read from stdin
	cmd.Stdin = bytes.NewReader(resp.AudioContent)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("failed to play audio: %v", err)
	}
}
