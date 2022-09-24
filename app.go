package main

import (
	"context"
	"fmt"
	"github.com/kkdai/youtube/v2"
	_ "image/jpeg"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) DownloadVideo(url string) (map[string]string, error) {
	client := youtube.Client{}
	url, err := youtube.ExtractVideoID(url)
	video, err := client.GetVideo(url)
	if err != nil {
		fmt.Println(err)
		return make(map[string]string), err
	}
	fmt.Println(video.Title)
	for i, thumbnail := range video.Thumbnails {
		fmt.Printf("Thumbnail %d: %s", i, thumbnail.URL)
	}
	out := make(map[string]string)
	out["title"] = video.Title
	out["thumbnail"] = video.Thumbnails[1].URL
	return out, nil

	//formats := video.Formats.WithAudioChannels()
	//stream, size, err := client.GetStream(video, &formats[0])
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//return video.Title, nil
}
