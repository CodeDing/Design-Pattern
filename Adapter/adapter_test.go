package Adapter

import "testing"

func TestAdapter(t *testing.T) {
	audioPlayer := NewAudioPlayer()
	audioPlayer.Play("mp3", "file.mp3")
	audioPlayer.Play("vlc", "file.vlc")
	audioPlayer.Play("mp4", "file.mp4")
	//t.Logf("play mp3: %s")
}
