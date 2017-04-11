package main

import (
	"regexp"
	"os"
	"errors"
)

type arg struct {
	No []*regexp.Regexp
	File string
}

var (
	argReg,_ = regexp.Compile("-no=(.+)")
)

func parseArgs(a []string) (*arg,error) {
	var err  error
	argNew := new(arg)
	for i := 1; i < len(a); i++ {
		s := a[i]
		if find := argReg.FindStringSubmatch(s);find != nil {
			newReg,err := regexp.Compile(".*"+find[1]+".*")
			if err != nil {
				return nil,err
			}
			argNew.No = append(argNew.No,newReg)
		} else {
			argNew.File = s
		}
	}

	if argNew.File =="" {
		err = errors.New("filename not input")
	}
	return argNew,err
}

func getArgs() (*arg,error) {
	return parseArgs(os.Args)
}
