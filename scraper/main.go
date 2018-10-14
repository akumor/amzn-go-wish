// Fetches Amazon wishlists based on user ID and stores the results.
package main

import (
	"fmt"
	"os"
	"time"
)

// WishlistItem to represent item retrieved from Amazon wishlist
type WishlistItem struct {
	amznID        string
	num           int
	name          string
	link          string
	oldPrice      string
	newPrice      string
	dateAdded     string
	priority      string
	rating        string
	totalRatings  string
	comment       string
	picture       string
	page          int
	ASIN          string
	largeSSLimage string
	affiliateURL  string
}

func main() {
	start := time.Now()
	IDCh := make(chan string)
	itemsCh := make(chan WishlistItem)
	go wishlistFetch(IDCh, itemsCh) // start a goroutine for getting wishlists and putting items into a channel
	go storeWishlistItems(itemsCh)  // start a goroutine for storing wishlist items
	for _, amznID := range os.Args[1:] {
		IDCh <- amznID
	}
	close(IDCh)
	close(itemsCh)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func wishlistFetch(IDCh chan string, itemsCh chan WishlistItem) {
	var exampleID WishlistItem

	// Loop over wishlist channel
	for {
		// TODO: Insert random delay greater than 1 second

		// Wait to receive a value
		id, ok := <-IDCh

		if !ok {
			// If the channel was closed, return.
			fmt.Printf("Goroutine wishlistFetch Down\n")
			return
		}

		// Display the value.
		fmt.Printf("Goroutine wishlistFetch ID %s\n", id)

		// TODO: query amazon-wish-lister for JSON
		exampleID = WishlistItem{
			id,
			1,
			"exampleName",
			"exampleLink",
			"exampleOldPrice",
			"exampleNewPrice",
			"exampleDateAdded",
			"examplePriority",
			"exampleRating",
			"exampleTotalRatings",
			"exampleComment",
			"examplePicture",
			1,
			"exampleASIN",
			"exampleLargeSSLimage",
			"exampleAffiliateURL",
		}

		itemsCh <- exampleID
	}

}

func storeWishlistItems(itemCh chan WishlistItem) {
	for {
		// Wait to receive a value
		item, ok := <-itemCh

		if !ok {
			// If the channel was closed, return.
			fmt.Printf("Goroutine storeWishlistItems Down\n")
			return
		}

		// Display the value.
		fmt.Printf("Goroutine storeWishlistItems item %s\n", item.amznID)

		// TODO: write goroutine to store item in Postgres database
	}
}
