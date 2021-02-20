package twitterdl

import (
	"fmt"
	"log"
	"os/exec"
)

func DownloadTwitterMedia(twitterURL string, fileName string) error {
	out, err := exec.Command("youtube-dl", twitterURL, "--output", fileName, "--max-filesize", "50m").Output()
	log.Println(string(out))
	if err != nil {
		return err
	}
	fmt.Println(twitterURL)
	return nil
}
