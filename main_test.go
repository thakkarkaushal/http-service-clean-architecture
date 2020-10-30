package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	_ "github.com/lib/pq"
)

func performGet(url string){
	resp, err := http.Get(url)
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	
	fmt.Print(string(body))
}

func TestFindByID(t *testing.T) {
	performGet("http://localhost:8080/users/1")
}

func TestFirstPageUsers(t *testing.T){
	performGet("http://localhost:8080/users")
}

func TestSecondPagesUsers(t *testing.T){
	performGet("http://localhost:8080/users?page=2")
}

func TestWhenDataNotAvailableForPage(t *testing.T){
	performGet("http://localhost:8080/users?page=4")
}
