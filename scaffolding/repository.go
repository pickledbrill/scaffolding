package scaffolding

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// RepositoryPost is the post body for create repository request.
type RepositoryPost struct {
	Name    string `json:"name"`
	Private bool   `json:"private"`
}

// RepositoryPostReturn stores the return body of repository creation request.
type RepositoryPostReturn struct {
	Name     string `json:"name"`
	CloneURL string `json:"clone_url"`
}

// SetupRepository setup the working directory for the repository.
func SetupRepository(dir string, repoFiles []string) (bool, error) {
	var setupErr error
	if _, setupErr = os.Stat(dir); os.IsNotExist(setupErr) {
		os.MkdirAll(dir, 0755)
	}
	for _, name := range repoFiles {
		s := strings.Split(name, "/")
		if len(s) == 1 {
			// single file
			_, setupErr = os.Create(fmt.Sprintf("%s/%s", dir, s[0]))
			if setupErr != nil {
				return false, setupErr
			}
		} else if len(s) > 1 {
			// folder
			dirName := fmt.Sprintf("%s/%s", dir, strings.Join(s[:len(s)-1], "/"))
			if _, setupErr = os.Stat(dirName); os.IsNotExist(setupErr) {
				os.MkdirAll(dirName, 0755)
			}
			if strings.TrimSpace(s[1]) != "" {
				_, setupErr = os.Create(fmt.Sprintf("%s/%s", dirName, s[len(s)-1]))
				if setupErr != nil {
					return false, setupErr
				}
			}
		}
	}
	return true, nil
}

// GitInitRepository git init the repository working directory
func GitInitRepository(dir string, repoInfo RepositoryPostReturn) {
	workDir := fmt.Sprintf("-workdir %s", dir)
	gitURL := fmt.Sprintf("-giturl %s", repoInfo.CloneURL)

	fmt.Println(workDir)
	fmt.Println(gitURL)

	fmt.Println(runtime.GOOS == "windows")

	if runtime.GOOS == "windows" {
		currentDir, _ := os.Getwd()
		fmt.Println(currentDir)
		cmd := exec.Command("powershell", "../../script/git-init.ps1", workDir, gitURL)
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}
