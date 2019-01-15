package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// IsExist returns whether a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// CreateDirs 循环创建目录
func CreateDirs(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Print(path + " Create Directory OK!")
	}
}

// GetGOPATHs returns all paths in GOPATH variable.
func GetGOPATHs() []string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" && strings.Compare(runtime.Version(), "go1.8") >= 0 {
		gopath = defaultGOPATH()
	}
	return filepath.SplitList(gopath)
}

// IsInGOPATH checks whether the path is inside of any GOPATH or not
func IsInGOPATH(thePath string) bool {
	for _, gopath := range GetGOPATHs() {
		if strings.Contains(thePath, filepath.Join(gopath, "src")) {
			return true
		}
	}
	return false
}

// 获取默认路径
func defaultGOPATH() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	if home := os.Getenv(env); home != "" {
		return filepath.Join(home, "go")
	}
	return ""
}
