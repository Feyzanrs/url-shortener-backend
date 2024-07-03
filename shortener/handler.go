package shortener

import (
	"encoding/json"
	"net/http"
	"strings"
)

// URLRequest yapılandırması, gelen URL'leri JSON formatında almak için kullanılır.
type URLRequest struct {
	URL string `json:"url"`
}

// HandleCreate, gelen bir URL için kısa bir URL oluşturur ve bunu JSON olarak geri döndürür.
func HandleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var request URLRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Error parsing JSON request body", http.StatusBadRequest)
		return
	}

	// Kısa URL oluştur, bu örnek sabit bir değer döndürüyor ama gerçek bir implementasyon eklenebilir.
	shortURL := StoreURL(request.URL)
	response := map[string]string{"shortUrl": shortURL}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// HandleRedirect, kısaltılmış URL'yi alır ve ilişkili uzun URL'ye yönlendirir.
func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	// "/r/" prefix'inden sonra gelen kısmı kısa URL olarak kabul eder.
	shortURL := strings.TrimPrefix(r.URL.Path, "/r/")
	longURL := RetrieveURL(shortURL)
	if longURL == "" {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Kullanıcıyı asıl URL'ye yönlendir.
	http.Redirect(w, r, longURL, http.StatusFound)
}
