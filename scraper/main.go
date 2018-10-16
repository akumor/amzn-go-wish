// Fetches Amazon wishlists based on user ID and stores the results.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// Wishlist is a slice of wishlist items.
type Wishlist struct {
	items []struct {
		num           int    `json:"num"`
		name          string `json:"name"`
		link          string `json:"link"`
		oldPrice      string `json:"old-price"`
		newPrice      string `json:"new-price"`
		dateAdded     string `json:"date-added"`
		priority      string `json:"priority"`
		rating        string `json:"rating"`
		totalRatings  string `json:"total-ratings"`
		comment       string `json:"comment"`
		picture       string `json:"picture"`
		page          int    `json:"page"`
		ASIN          string `json:"ASIN"`
		largeSSLimage string `json:"large-ssl-image"`
		affiliateURL  string `json:"affiliate-url"`
	}
}

func main() {
	start := time.Now()
	IDCh := make(chan string)
	itemsCh := make(chan Wishlist)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		wishlistFetch(IDCh, itemsCh) // start a goroutine for getting wishlists and putting items into a channel
		wg.Done()
	}()
	go func() {
		storeWishlistItems(itemsCh) // start a goroutine for storing wishlist items
		wg.Done()
	}()
	for _, amznID := range os.Args[1:] {
		IDCh <- amznID
	}
	close(IDCh)
	wg.Wait()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func wishlistFetch(IDCh chan string, itemsCh chan Wishlist) {
	var currentWishlist Wishlist
	var wishlistURL string

	// Loop over wishlist channel
	for {
		// TODO: Insert random delay greater than 1 second

		// Wait to receive a value
		id, ok := <-IDCh

		if !ok {
			// If the channel was closed, return.
			close(itemsCh)
			fmt.Printf("Goroutine wishlistFetch Down\n")
			return
		}

		// Display the value.
		fmt.Printf("Goroutine wishlistFetch ID %s\n", id)

		// Build URL to query
		wishlistURL = fmt.Sprintf("http://192.168.0.150:8080/wishlist.php?id=%s&reveal=all&sort=priority&format=json", id)

		//DEBUG
		fmt.Printf("DEBUG wishlistURL: %s\n", wishlistURL)

		// TODO: query amazon-wish-lister for JSON

		//For control over proxies,
		//TLS configuration, keep-alives,
		//compression, and other settings, create a Transport
		tr := &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		}

		//DEBUG
		fmt.Printf("DEBUG AFTER CREATING TRANSPORT\n")

		// For control over HTTP client headers,
		// redirect policy, and other settings,
		// create a Client
		// A Client is an HTTP client
		client := &http.Client{Transport: tr}

		//DEBUG
		fmt.Printf("DEBUG AFTER CREATING CLIENT\n")

		// Send the request via a client
		// Do sends an HTTP request and
		// returns an HTTP response
		resp, err := client.Get(wishlistURL)
		log.Println(resp)
		log.Println(err)
		if err != nil {
			log.Println("client.Get err was not nil!")
			log.Fatal(err)
		}

		fmt.Printf("DEBUG AFTER SENDING REQUEST\n")

		// Use json.Decode for reading streams of JSON data
		if err := json.NewDecoder(resp.Body).Decode(&currentWishlist); err != nil {
			log.Println(err)
		}

		// Callers should close resp.Body
		// when done reading from it
		resp.Body.Close()

		itemsCh <- currentWishlist
	}

}

func storeWishlistItems(itemCh chan Wishlist) {
	for {
		// Wait to receive a value
		wl, ok := <-itemCh

		if !ok {
			// If the channel was closed, return.
			fmt.Printf("Goroutine storeWishlistItems Down\n")
			return
		}

		// Display the value.
		fmt.Printf("Goroutine storeWishlistItems item %s\n", wl.items)

		// TODO: write goroutine to store item in Postgres database
	}
}
