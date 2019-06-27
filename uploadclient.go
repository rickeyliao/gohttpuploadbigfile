package main

import (
	"io"
	"mime/multipart"

	"os"
	"net/http"
	"sync"
	"fmt"
	"io/ioutil"
)

func main()  {

	wg := &sync.WaitGroup{}
	name:=os.Args[1]

	piper,pipew:=io.Pipe()
	defer piper.Close()

	m:=multipart.NewWriter(pipew)
	//fmt.Println(m.FormDataContentType())
	//defer m.Close()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer pipew.Close()

		part,err:=m.CreateFormFile("filename","foo.txt")
		if err!=nil{
			fmt.Println(err)
			return
		}
		defer m.Close()

		file, err := os.Open(name)
		if err != nil {
			return
		}
		defer file.Close()

		if _, err = io.Copy(part, file); err != nil {
			return
		}

	}()

	resp,err:=http.Post("http://localhost:50810/upload", m.FormDataContentType(), piper)
	wg.Wait()

	if err!=nil{
		fmt.Println(err)
	}

	if resp!=nil && resp.Body != nil{
		ret,_:=ioutil.ReadAll(resp.Body)
		fmt.Println(string(ret))
		resp.Body.Close()
	}


}
