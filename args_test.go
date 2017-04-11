package main

import (
	"testing"
	"regexp"
	"fmt"
	"strings"
)

func TestRegexp(t *testing.T)  {
	tReg,_:= regexp.Compile("(.+)hello")
	fmt.Println(tReg.FindStringSubmatch("abchello"))
	fmt.Println(tReg.ReplaceAllString("serrrhello","$1+"))
}

func TestArg(t *testing.T) {
	cmd := "libdeploy -no=123 -no=ert -no=aae hello.exe"
	cmdList := strings.Split(cmd," ")
	fmt.Println(parseArgs(cmdList))
}

