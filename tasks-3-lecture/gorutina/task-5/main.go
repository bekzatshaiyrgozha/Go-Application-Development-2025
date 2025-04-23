package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Уақыт шектеулі функция
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 3 секундтық контекст жасаймыз
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		select {
		case <-time.After(2 * time.Second): // Егер 2 сек-та орындалса
			fmt.Fprintln(w, "Response OK ✅")
		case <-ctx.Done(): // Егер 3 сек-та орындалмаса
			http.Error(w, "Request Timeout ⏳❌", http.StatusGatewayTimeout)
		}
	}()

	wg.Wait() // Барлық горутиналардың аяқталғанын күтеміз
}

func main() {
	http.HandleFunc("/process", handleRequest)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
