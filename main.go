package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

const HELLO_D = `import std.stdio;

void main()
{
    writeln("Hello, world!");
}
`

type SuiteLstItm struct {
	Name string
	Func func() error
}

var st_lst = []SuiteLstItm{
	SuiteLstItm{"gcc gdc", try_GCC},
	SuiteLstItm{"llvm ldc", try_LLVM},
	SuiteLstItm{"dmd", try_DMD},
}

func main() {
	err := ioutil.WriteFile("hello.d", []byte(HELLO_D), 0o777)
	if err != nil {
		log.Fatalln(err)
	}

	for _, i := range st_lst {
		fmt.Printf("%20s:", i.Name)
		r := i.Func()
		res := "Ok"
		if r != nil {
			res = "Fail"
		}
		fmt.Println(res)
	}

	return
}

func try_GCC() error {
	cmd := []string{"gdc", "./hello.d"}
	c := exec.Command(cmd[0], cmd[1:]...)
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func try_LLVM() error {
	cmd := []string{"ldmd2", "./hello.d"}
	c := exec.Command(cmd[0], cmd[1:]...)
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func try_DMD() error {
	cmd := []string{"dmd", "./hello.d"}
	c := exec.Command(cmd[0], cmd[1:]...)
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}
