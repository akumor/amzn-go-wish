// Fetches Amazon wishlists based on user ID and stores the results.
package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

	// Define host port id outputfile flags
	apiHost := flag.String("host", "localhost", "API host to query for the wishlist.")
	apiPort := flag.String("port", "80", "API port to query for the wishlist.")
	id := flag.String("id", "DEFAULT", "ID of the Amazon wishlist.")
	outputFile := flag.String("file", "output.csv", "Path to output CSV file.")

	flag.Parse()

	// Create CSV file
	file, err := os.Create(*outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Display the value.
	fmt.Printf("Goroutine wishlistFetch ID %s\n", *id)

	// Build URL to query
	wishlistURL = fmt.Sprintf("http://%s:%s/wishlist.php?id=%s&reveal=all&sort=priority&format=json", *apiHost, *apiPort, *id)

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

	// Create slice of slices for writing to CSV
	records := [][]string{
		{"Num", "Name", "Link", "OldPrice", "NewPrice", "DateAdded", "Priority", "Rating", "TotalRatings", "Comment", "Picture", "Page", "ASIN", "LargeSslImage", "AffiliateURL"},
	}

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
		line := []string{
			strconv.Itoa(item.Num), item.Name, item.Link, item.OldPrice, item.NewPrice, item.DateAdded, item.Priority, item.Rating, item.TotalRatings, item.Comment, item.Picture, strconv.Itoa(item.Page), item.ASIN, item.LargeSslImage, item.AffiliateURL,
		}
		records = append(records, line)
	}

	err = writer.WriteAll(records)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
