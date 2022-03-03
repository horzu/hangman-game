# Kazanımlar

1. Random sayı üretme

    dictionary slice'ı üzerinde rastgele bir elemanı seçip selectedWord değişkenine atamak için:
    - math/rand ve time modülleri import edilir.
    - rasgele bir sayı üretmek için öncelikle rand.Seed() kullanılır.
    - Her seferinde farklı random sayılar almak için yeni bir seed oluşturulur.
    - Yeni bir seed oluşturmak için genellikle rand.Seed(time.Now().Unix()) kullanılır.
    - Seed oluşturduktan sonra rand.Intn(sayı) 0 ile sayı-1 arasında bir sayı üretir.

    ```go
    import (
    "math/rand"
    "time"
    )

    rand.Seed(time.Now().Unix())
    selectedWord := fmt.Sprint(dictionary[rand.Intn(len(dictionary))])
    ```

2. Belli bir sayıda kendini tekrarlayan string üretme

    Bir stringi istenen sayı kadar tekrar ettirmek için strings.Repeat(string, sayı) kullanılır.

    ```go
        repeatedString := strings.Repeat("_", tekrar)
    ```

3. Bir string'in içinde aranan başka bir string'in olup olmadığını bulmak için strings.Contains(aranacakString, arananString) kullanılır.

    ```go
        mertBul := strings.Contains("MertŞakar", "Mert") // bool
    ```

4. Rune'ları string'e çevirmek için basitce string([]rune) kullanılır.

5. Rune'larla strings işlemleri yapılamaz. İşlem yapmak için yukarıdaki gibi öncelikle string'e çevrilmelidir.

6. Yeni bir string oluşturmak için fmt.Sprintf() kullanılır.

7. Dosyadan okuma işlemleri için os.ReadFile("dosyakonumu") kullanılır.

    ```go
    import "os"
    data, err := os.ReadFile("/istenen/dosya/konumu") // istenen dosyanın içeriğini alır.
    if err != nil {
        panic(err)
    }
    ```

8. Kullanıcıdan input almak için bufio paketi kullanılır.
    Kullanıcıdan sadece tek bir string almak isteniyorsa ReadRune() kullanılır. ReadRune() kullanılırsa, kullanıcı birden çok değer girse bile ilk değeri alır. Örn: asdfasdf girilirse "a" değeri rune olarak alınır.

    ```go
    var inputReader = bufio.NewReader(os.Stdin)
    letter, _ , err:= inputReader.ReadRune() // kullanıcı girdisine ait ilk rune'u input olarak alır.
    if err != nil {
        fmt.Println(err)
    }
    ```

9. rune değerlerinin harf olup olmadığını kontrol etmek için direkt ascii tablosundaki decimal değerlere göre kontrol edilebilir.
    Büyük harflere ait değerler => 65(a) - 90(z) arası
    Küçük harflere ait değerler => 97(A) - 122(Z) arası
