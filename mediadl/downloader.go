package mediadl

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	shortuuid "github.com/lithammer/shortuuid/v3"
)

func DownloadMedia(URL string) ([]byte, error) {
	fileName := shortuuid.New()
	defer os.Remove(fileName)
	cmd := exec.Command("youtube-dl", URL, "--output", fileName, "--max-filesize", "50m")
	log.Println(cmd.String())
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return content, nil
}
