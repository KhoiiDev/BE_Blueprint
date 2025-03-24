package upload

import (
	"os"
)

const VideoPath = "upload/files/videos/"

func GetVideoFullPath() string {
	if _, err := os.Stat(VideoPath); os.IsNotExist(err) {
		os.MkdirAll(VideoPath, os.ModePerm)
	}
	return VideoPath
}

// CheckVideoSavePath checks if the video save path exists
func CheckVideoSavePath() bool {
	_, err := os.Stat(GetVideoFullPath())
	return !os.IsNotExist(err)
}

func MkdirVideoSavePath() error {
	path := GetVideoFullPath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}
