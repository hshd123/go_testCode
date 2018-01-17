package main

import (
	"bytes"
	"io"
	"os/exec"
	"log"
	"os"
	"os/signal"
	"syscall"
)
import "fmt"

func main() {
	execTest5()
}

func execTest1()  {
	cmd := exec.Command("ls","-lah")
	out , err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func execTest2()  {
	fmt.Println("test2")
	cmd := exec.Command("ls","-lah")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run err %s\n",err)
	}
	outStr , errStr := string(stdout.Bytes()) , string(stderr.Bytes())
	fmt.Println(outStr , errStr)
}

func execTest3()  {
	cmd := exec.Command("ls" , "-lah")
	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	cmd.Start()
	go func() {
		stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
	}()
	go func() {
		stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)
	}()
	err := cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatalf("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdout), string(stderr)
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}

func copyAndCapture(w io.Writer , r io.Reader)([]byte , error)  {
	var out []byte
	buf := make([]byte,1024,1024)
	for {
		n,err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out,d...)
		}
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return out , err
		}
	}
	panic(true)
	return nil , nil
}

func execTest4() {
	lsPath , err := exec.LookPath("ls")
	if err != nil {
		fmt.Println(err)
	}
	args := []string{"ls" ,"-lah"}
	env := os.Environ()
	err2 := syscall.Exec(lsPath,args,env)
	if err != nil {
		fmt.Println(err2)
	}
	fmt.Println(lsPath)
}

func execTest5()  {
	sigs := make(chan os.Signal , 1)
	done := make(chan bool , 1)
	signal.Notify(sigs , syscall.SIGINT,syscall.SIGTERM )

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println("signal -- " , sig)
		done <- true
	}()
	
	fmt.Println("awaiting signal")
	<- done
	fmt.Println("exiting")
	
}