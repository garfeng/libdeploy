package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"path/filepath"
	"regexp"
)

func runLdd(filename string) ([]string, error) {
	reg, _ := regexp.Compile(`=> (.+\.`+EXT+`.*) `)
	cmd := exec.Command("ldd", filename)
	cmd.Stderr = os.Stderr
	buff := bytes.NewBuffer(nil)
	cmd.Stdout = buff
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	bufStr := buff.String()
	match := reg.FindAllStringSubmatch(bufStr, -1)
	bufList := make([]string, 0)
	if match == nil {
		return nil, nil
	}
	for _, value := range match {
		bufList = append(bufList, value[1])
	}
	//bufList := strings.Split(bufStr, "\n")
	return bufList, err
}

func parseLibs(list []string, a *arg) {
	path, _ := filepath.Split(a.File)
	if path == "" {
		path = "./"
	}
	fmt.Println("Path is", path)
OUT:
	for _, dll := range list {
		for _, reg := range a.No {
			if reg.MatchString(dll) {
				fmt.Println("skip", dll)
				continue OUT
			}
		}
		_, dllName := filepath.Split(dll)
		thisDirDll := filepath.Join(path, dllName)
		fmt.Println("copy", dll, "=>", thisDirDll)
		/*
			if !mfile.Exist(dll) {
				dll = diskReg.ReplaceAllString(dll,"$1:/")
			}
			if err := mfile.Copy(dll, thisDirDll);err != nil {
				fmt.Println(err)
			}
		*/
		strCmd := fmt.Sprintln("cp", dll, path, "-f")
		cmd := exec.Command("sh", "-c", strCmd)
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}
	a, err := getArgs()
	if err != nil {
		fmt.Println(err)
		printHelp()
		return
	}

	list, err := runLdd(a.File)
	if err != nil {
		fmt.Println(err)
		return
	}
	if list == nil {
		return
	}
	parseLibs(list, a)
}
