package main

import ("fmt")

fun main(){
	e := ebook{
		name : "The little fox"
		date: 6-6-1996,
		study: false,
	}
	 err := CreateEbook(e)
	 if err != nil{
		 log.Fatal(err)
	 }
	 fmt.Println("Created! :)")
}