package main

import (
	"fmt"
	"net/http"
)

const (
	googleUrl = "https://google.com"
	yandexUrl = "https://ya.ru"
)

func main() {
	//incorrectVariant()
	//firstCorrectVariant()
	//secondCorrectVariant()
}

func incorrectVariant() {
	var res *http.Response
	flag := true

	if flag {
		res, err := http.Get(googleUrl)
		if err != nil {
			fmt.Printf("failed to get url: %v", err)
		}
		fmt.Println("Google res:", res.Status)
	} else {
		res, err := http.Get(yandexUrl)
		if err != nil {
			fmt.Printf("failed to get url: %v", err)
		}
		fmt.Println("Yandex res:", res.Status)
	}

	fmt.Println("final:", res.Status)
}

func firstCorrectVariant() {
	var result *http.Response
	flag := true

	if flag {
		res, err := http.Get(googleUrl)
		if err != nil {
			fmt.Printf("failed to get url: %v", err)
		}

		fmt.Println("Google res:", res.Status)
		result = res
	} else {
		res, err := http.Get(yandexUrl)
		if err != nil {
			fmt.Printf("failed to get url: %v", err)
		}

		fmt.Println("Yandex res:", res.Status)
		result = res
	}

	fmt.Println("final:", result.Status)
}

func secondCorrectVariant() {
	var res *http.Response
	var err error
	flag := true

	if flag {
		res, err = http.Get(googleUrl)
	} else {
		res, err = http.Get(yandexUrl)
	}
	if err != nil {
		fmt.Printf("failed to get url: %v", err)
	}

	fmt.Println("final:", res.Status)
}
