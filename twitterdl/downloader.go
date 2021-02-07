package twitterdl

import (
	"fmt"
	"log"
	"os/exec"
)

func DownloadTwitterMedia(twitterURL string, fileName string) {
	output := fmt.Sprintf("%s.mp4", fileName)
	cmd := exec.Command("youtube-dl", twitterURL, "--output", output)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(twitterURL)
}
