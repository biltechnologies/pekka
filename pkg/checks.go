package pkg

import (
	"log"
	"os/exec"
)

//GetDockerComposePath returns the docker-compose path else prints an error and exits
func GetDockerComposePath() string {
	dockerComposePath, err := exec.LookPath("docker-compose")
	if err != nil {
		log.Fatalln("docker-compose could not be found")
	}
	return dockerComposePath
}
