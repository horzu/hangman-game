package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

var dictionary = []string{
	"masa",
	"sandalye",
	"elma",
	"almanya",
}
func main(){
	// oyun durumunu yazdır -> state0
	state, err := os.ReadFile("./states/hangman0")
	if err != nil {
		panic(err)
	}
	fmt.Println("Hangman durumu: ", string(state))


	// tahmin etmemiz gereken bir kelime türet (rand.Seed(time.Now().UnixNano()))
	rand.Seed(time.Now().Unix())
	selectedWord := fmt.Sprint(dictionary[rand.Intn(len(dictionary))])
	selectedWorDashes := strings.Repeat("_ ", utf8.RuneCountInString(selectedWord))
	fmt.Println("Tahmin etmeniz gereken kelime: ", selectedWorDashes, "( kelime uzunluğu:", len(selectedWord),")")


	// kullanıcı girdisini oku
	var inputReader = bufio.NewReader(os.Stdin)
	fmt.Print("Tahmin ettiğiniz harfi girip Enter'a basınız: ")
	letter, _ := inputReader.ReadString('\n')
	fmt.Printf("Tahmin edilen harf: %s\n", letter)


	// * tahmin ettiğiniz kelimeyi yazdırın


// * adam asmaca durumunu yazdır

// * validate (sadece harf olarak validate edilmeli)
// harf doğrusa tahmin olarak yerine yazacaksın ( m a _ a) // birden fazla harf varsa tamamlayacak.
// * doğruysa, tahmin edilen harfleri güncelleyin
// * yanlışsa, adam asmaca durumunu güncelleyin
// kelime tahmin edilirse -> kazanırsın
// adam asmaca tamamlandıysa -> kaybedersiniz
}
