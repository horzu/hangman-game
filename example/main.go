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

var state int = 3 // Hangman state'i başlat. 3'ten başlattım çünkü en başta adamın asıldaığı direği görsel olarak görmek istedim.
var iteration int = 0 // kullanıcının tahmin sayısını tutmak için iteration değişkeni tanımlandı.
var wrongGuesses []string // yanlış tahminleri göstermek için bir slice oluştur.

func main(){
	// tahmin etmemiz gereken bir kelime türet (rand.Seed(time.Now().UnixNano())) ile.
	rand.Seed(time.Now().Unix())
	selectedWord := fmt.Sprint(dictionary[rand.Intn(len(dictionary))])

	selectedWorDashes := strings.Repeat("_", utf8.RuneCountInString(selectedWord)) // seçilen kelime uzunluğunda _ gösteren yeni bir string oluştur.
	var newWordDashes = strings.Split(selectedWorDashes, "") // seçilen kelime uzunluğunda oluşturulan _'lar için yeni bir slice oluştur.

	game(selectedWord, selectedWorDashes, newWordDashes) // oyunun oynandığı fonsiyonu çağır.

	for state < 10 { // Son state'e gelene kadar oyunun oynanması için while tarzı döngü yaratılır.
		game(selectedWord, selectedWorDashes, newWordDashes)
		fmt.Print("\n") // hamleleri görsel olarak ayırabilmek için 100 tane yıldız yazdır.
		fmt.Println(strings.Repeat("*",100)) // hamleleri görsel olarak ayırabilmek için 100 tane yıldız yazdır.
	}

	if state == 11 { // eğer state 11'se kazandığını göster. game fonksiyonu üzerinden oyun kazanıldıysa state 11'e eşitleniyor.
		fmt.Println("Tebrikler kazandınız! İstenen kelime: ", selectedWord)
		fmt.Println(iteration, "tahminde bildiniz!")
	}

	if state == 10 { // adam asmaca tamamlandıysa -> kaybedersiniz
		fmt.Println("Kaybettin! İstenen kelime: ", selectedWord)
	}

}

// game fonksiyonu oyun içi işlemlerin yapılacağı fonsiyondur. Seçilen kelime(selectedWord), seçilen kelimenin _ ile güncellenmiş hali (selectedWorDashes) ve selectedworDashes'in slice versiyonu(newWordDashes) parametrelerini alır)
func game(selectedWord string, selectedWorDashes string, newWordDashes []string) {
	fmt.Println(hangmanState(state)) 	// oyun durumunu yazdır -> state3
	fmt.Println("Tahmin etmeniz gereken kelime: ", strings.Join(newWordDashes, " "), "( kelime uzunluğu:", len(newWordDashes),")") 	// Tahmin edilmesi gereken kelimenin mevcut durumunu yazdır
	fmt.Println("Yanlış tahminler: ", strings.Join(wrongGuesses, ",")) // yanlış tahminleri yazdır.

	letter := getUserInput() 	// kullanıcı girdisini getUserInput fonksiyonu ile al

	isLetter := validateInput(letter) // girilen input harf olana kadar input iste
	for !isLetter {
		letter = getUserInput()
		isLetter = validateInput(letter)
	}

	fmt.Printf("Tahmin edilen harf: %s\n", strings.ToLower(string(letter))) // tahmin edilen harfi yazdır.
	
	if strings.Contains(selectedWord, strings.ToLower(string(letter))) { // tahmin edilen harf seçilen kelimenin içinde mi kontrol et
		for k, elem := range selectedWord{ // harf doğruysa tahmin olarak yerine yazacaksın ( m a _ a) // birden fazla harf varsa tamamlayacak.
			if string(elem) == strings.ToLower(string(letter)){
				newWordDashes[k] = strings.ToLower(string(letter)) // * doğruysa, tahmin edilen harfleri güncelleyin
			}
		}
		
	} else { // * yanlışsa, adam asmaca durumunu güncelleyin
		state++ 
		wrongGuesses = append(wrongGuesses, string(letter))
	}
	selectedWorDashes = strings.Join(newWordDashes, "") // seçilen kelime _ ile güncellenen kelime olarak düzenlenir.

	if selectedWorDashes == selectedWord { // kelime tahmin edilirse -> kazanırsın. Kazanmak için state 11'e eşitlenip main function'da karşılanıyor.
		state = 11
	}
	iteration++
}

// hangman state fonksiyonu adam asmaca durumunu gösterir.
func hangmanState(strings int) string{
	stringState := fmt.Sprintf("./states/hangman%d", strings) // Mevcut state'e göre istenen dosyayı belirler.
	state, err := os.ReadFile(stringState) // istenen state'e ait dosyayı alır.
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Hangman durumu: \n%s", string(state)) // mevcut state'e ait durumu gösterir.
}

// kullanıcı girdisini oku
func getUserInput() rune {
	var inputReader = bufio.NewReader(os.Stdin)
	fmt.Print("Tahmin ettiğiniz harfi giriniz: ")
	letter, _ , err:= inputReader.ReadRune() // kullanıcı girdisine ait ilk rune'u input olarak alır.
	if err != nil {
		fmt.Println(err)
	}
	return letter // ilgili rune'u döner.
}

// * validate (sadece harf olarak validate edilmeli)
func validateInput(letter rune) bool {
	if (letter > 64 && letter < 91) || (letter > 96 && letter < 123) { // girilen input rune'u harf mi kontrol eder.
		// * tahmin ettiğiniz harfi yazdırın
		return true
	} else {
		fmt.Printf("Lütfen geçerli bir harf giriniz! \n")
		return false
	}
}

