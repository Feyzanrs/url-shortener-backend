package shortener

var urlMap = make(map[string]string)

// StoreURL bir URL'yi kaydeder ve kısaltılmış versiyonunu döndürür
func StoreURL(longURL string) string {
	shortURL := "someShortURL" // Gerçek bir kısaltma mantığı burada olacak
	urlMap[shortURL] = longURL
	return shortURL
}

// RetrieveURL kısaltılmış bir URL için asıl URL'yi döndürür
func RetrieveURL(shortURL string) string {
	return urlMap[shortURL]
}
