package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
)

type pageInfo struct {
	StatusCode int
	Links      map[string]int
}

type NewScrape struct {
	Url string `json:"url" bson:"url"`
}

type RalphsProduct struct {
	Url   string `json:"url"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ScrapeGoQueryCostco(c *fiber.Ctx) error {
	url := "https://www.costco.com/grocery-household.html?currentPage=1&pageSize=24&sortBy=item_page_views+desc&deliveryFacetFlag=true&keyword=OFF"

	// Create a new HTTP client and set a custom User-Agent header
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36")

	// Make the request with the custom User-Agent header
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error fetching URL:", err)
	}
	defer resp.Body.Close()

	// Parse the HTML content
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error parsing HTML:", err)
	}

	// Make an HTTP GET request to the URL
	// doc, err := goquery.NewDocument(url)
	// if err != nil {
	// 	log.Fatal("Error fetching URL:", err)
	// }
	fmt.Println("doc", doc)
	products := doc.Find(".product-list")
	if products.Length() == 0 {
		html, _ := doc.Html()
		fmt.Println("html", html)
		log.Fatal("No ProductGridContainer elements found")
	}

	products.Each(func(i int, s *goquery.Selection) {
		// Extract the information from each element
		productName := s.Find(".description").Text()
		productPrice := s.Find(".price").Text()
		promo := s.Find(".promo").Text()

		fmt.Printf("Product %d: %s\n", i+1, productName)
		fmt.Printf("Price %d: %s\n", i+1, productPrice)
		fmt.Println("promo", promo)
		fmt.Println("-------------")
	})

	return c.Status(200).JSON(fiber.Map{"result": "testing"})
}

func ScrapeGoQueryAlbertsons(c *fiber.Ctx) error {
	url := "https://www.albertsons.com/home/snacks-on-the-go.html?sort=&price=0%20TO%204%2E99"

	// Create a new HTTP client and set a custom User-Agent header
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36")

	// Make the request with the custom User-Agent header
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error fetching URL:", err)
	}
	defer resp.Body.Close()

	// Parse the HTML content
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error parsing HTML:", err)
	}

	// Make an HTTP GET request to the URL
	// doc, err := goquery.NewDocument(url)
	// if err != nil {
	// 	log.Fatal("Error fetching URL:", err)
	// }
	fmt.Println("doc", doc)
	products := doc.Find(".prd-itm")
	html, _ := doc.Html()
	fmt.Println("html", html)
	if products.Length() == 0 {
		html, _ := doc.Html()
		fmt.Println("html", html)
		log.Fatal("No ProductGridContainer elements found")
	}

	products.Each(func(i int, s *goquery.Selection) {
		fmt.Println("S", s)
		// Extract the information from each element
		productName := s.Find(".product-title").Text()
		productPrice := s.Find("span.product-price__saleprice").Text()
		//product-price__saleprice product-price__discounted-price

		fmt.Printf("Product %d: %s\n", i+1, productName)
		fmt.Printf("Price %d: %s\n", i+1, productPrice)
		fmt.Println("-------------")
	})

	return c.Status(200).JSON(fiber.Map{"result": "testing"})
}

func ScrapeGoQueryHmart(c *fiber.Ctx) error {
	url := "https://www.hmart.com/weekly-super-sale"

	// Create a new HTTP client and set a custom User-Agent header
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36")

	// Make the request with the custom User-Agent header
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error fetching URL:", err)
	}
	defer resp.Body.Close()

	// Parse the HTML content
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error parsing HTML:", err)
	}

	// Make an HTTP GET request to the URL
	// doc, err := goquery.NewDocument(url)
	// if err != nil {
	// 	log.Fatal("Error fetching URL:", err)
	// }
	fmt.Println("doc", doc)
	products := doc.Find(".product-items")
	//html, _ := doc.Html()
	//fmt.Println("html", html)
	// if products.Length() == 0 {
	// 	html, _ := doc.Html()
	// 	fmt.Println("html", html)
	// 	log.Fatal("No ProductGridContainer elements found")
	// }

	products.Each(func(i int, s *goquery.Selection) {
		fmt.Println("S", s)
		// Extract the information from each element
		productName := s.Find(".product-item-link").Text()
		productPrice := s.Find(".price").Text()
		//product-price__saleprice product-price__discounted-price

		fmt.Printf("Product %d: %s\n", i+1, productName)
		fmt.Printf("Price %d: %s\n", i+1, productPrice)
		fmt.Println("-------------")
	})

	return c.Status(200).JSON(fiber.Map{"result": "testing"})
}

func ScrapeGoQueryTarget(c *fiber.Ctx) error {
	url := "https://www.target.com/c/grocery-deals/-/N-k4uyq"

	// Create a new HTTP client and set a custom User-Agent header
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36")

	// Make the request with the custom User-Agent header
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error fetching URL:", err)
	}
	defer resp.Body.Close()

	// Parse the HTML content
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error parsing HTML:", err)
	}

	fmt.Println("doc", doc)
	products := doc.Find(".styles__BaseContainer-sc-1ealqwt-1")

	var results []interface{}

	products.Each(func(i int, s *goquery.Selection) {
		fmt.Println("S", s)
		// Extract the information from each element
		data := s.Find("span").Text()
		results = append(results, data)
	})

	return c.Status(200).JSON(fiber.Map{"result": results})
}
