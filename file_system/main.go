package main


import (
	  "fmt"
	  "os"
)

func main(){
	 
	// reading the files
	file,err:=os.Open("file.txt")

	if err !=nil{
        panic(err)
	}

	file_info,err:=file.Stat()
   
	if err !=nil{
        panic(err)
	}
	fmt.Println(file_info.Name())
	fmt.Println(file_info.IsDir()) // ye check karta hai ki file hai or folder
	fmt.Println(file_info.Size())
	fmt.Println(file_info.Mode())
	fmt.Println(file_info.ModTime())

	
}