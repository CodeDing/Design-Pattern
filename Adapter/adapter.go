package Adapter

import "fmt"

//NOT RECOMMEND: not straightforward

type MediaPlayer interface {
	Play(string, string)
}

type AdvancedMediaPlayer interface {
	PlayVLC(string)
	PlayMP4(string)
}

type VLCPlayer struct{}

func (VLCPlayer) PlayVLC(file string) {
	fmt.Println("play vlc file: " + file)
}

func (VLCPlayer) PlayMP4(file string) {
	return
}

type MP4Player struct{}

func (MP4Player) PlayVLC(file string) {
	return
}

func (MP4Player) PlayMP4(file string) {
	fmt.Println("play mp4 file: " + file)
}

type MediaAdapter struct {
	AdvancedMediaPlayer
}

func newMediaAdapter(audioType string) *MediaAdapter {
	switch audioType {
	case "vlc":
		return &MediaAdapter{VLCPlayer{}}
	case "mp4":
		return &MediaAdapter{MP4Player{}}
	}
	return nil
}

func (m MediaAdapter) Play(audioType string, file string) {
	switch audioType {
	case "vlc":
		m.AdvancedMediaPlayer.PlayVLC(file)
	case "mp4":
		m.AdvancedMediaPlayer.PlayMP4(file)
	default:
		panic("unsupported audio type: " + audioType)
	}
}

type AudioPlayer struct {
	MediaAdapter
}

//func newAudioPlayer(audioType string) *AudioPlayer {
//	return &AudioPlayer{MediaAdapter{newMediaAdapter(audioType)}}
//}
func NewAudioPlayer() *AudioPlayer {
	return &AudioPlayer{MediaAdapter{}}
}

func (p AudioPlayer) newAudioPlayer(audioType string) *MediaAdapter {
	return newMediaAdapter(audioType)
}

func (p AudioPlayer) Play(audioType, file string) {
	switch audioType {
	case "mp3":
		fmt.Println("play mp3: " + file)
	case "vlc":
		fallthrough
	case "mp4":
		p.newAudioPlayer(audioType).Play(audioType, file)
	}
}
