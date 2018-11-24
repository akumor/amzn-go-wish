// Fetches Amazon wishlists based on user ID and stores the results.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Wishlist is a slice of wishlist items.
type Wishlist []struct {
	Num           int    `json:"num"`
	Name          string `json:"name"`
	Link          string `json:"link"`
	OldPrice      string `json:"old-price"`
	NewPrice      string `json:"new-price"`
	DateAdded     string `json:"date-added"`
	Priority      string `json:"priority"`
	Rating        string `json:"rating"`
	TotalRatings  string `json:"total-ratings"`
	Comment       string `json:"comment"`
	Picture       string `json:"picture"`
	Page          int    `json:"page"`
	ASIN          string `json:"ASIN"`
	LargeSslImage string `json:"large-ssl-image"`
	AffiliateURL  string `json:"affiliate-url"`
}

func main() {
	var currentWishlist Wishlist
	var wishlistURL string

	start := time.Now()

	// Get the first command line argument for the ID
	id := os.Args[1]

	// Display the value.
	fmt.Printf("Goroutine wishlistFetch ID %s\n", id)

	// Build URL to query
	wishlistURL = fmt.Sprintf("http://192.168.0.150:8080/wishlist.php?id=%s&reveal=all&sort=priority&format=json", id)

	//DEBUG
	fmt.Printf("DEBUG wishlistURL: %s\n", wishlistURL)

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

	// Print contents of the wishlist
	for _, item := range currentWishlist {
		fmt.Print("##########################################\n")
		fmt.Printf("Num: %d\n", item.Num)
		fmt.Printf("Name: %s\n", item.Name)
		fmt.Printf("Link: %s\n", item.Link)
		fmt.Printf("OldPrice: %s\n", item.OldPrice)
		fmt.Printf("NewPrice: %s\n", item.NewPrice)
		fmt.Printf("DateAdded: %s\n", item.DateAdded)
		fmt.Printf("Priority: %s\n", item.Priority)
		fmt.Printf("Rating: %s\n", item.Rating)
		fmt.Printf("TotalRatings: %s\n", item.TotalRatings)
		fmt.Printf("Comment: %s\n", item.Comment)
		fmt.Printf("Picture: %s\n", item.Picture)
		fmt.Printf("Page: %d\n", item.Page)
		fmt.Printf("ASIN: %s\n", item.ASIN)
		fmt.Printf("LargeSslImage: %s\n", item.LargeSslImage)
		fmt.Printf("AffiliateURL: %s\n", item.AffiliateURL)
		fmt.Print("##########################################\n")
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
