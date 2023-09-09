package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
)

func ScrapeGoQueryCostco(c *fiber.Ctx) error {
	url := "https://www.costco.com/grocery-household.html?currentPage=1&pageSize=24&sortBy=item_page_views+desc&deliveryFacetFlag=true&keyword=OFF"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error fetching URL:", err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error parsing HTML:", err)
	}
	fmt.Println("doc", doc)
	products := doc.Find(".product-list")
	if products.Length() == 0 {
		html, _ := doc.Html()
		fmt.Println("html", html)
		log.Fatal("No ProductGridContainer elements found")
	}

	products.Each(func(i int, s *goquery.Selection) {
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

// doesn't quite work with albertsons - most likely due to anit-scrapping from albertsons?
func ScrapeGoQueryAlbertsons(c *fiber.Ctx) error {
	url := "https://www.albertsons.com/home/snacks-on-the-go.html?sort=&price=0%20TO%204%2E99"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error fetching URL:", err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error parsing HTML:", err)
	}
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
		productName := s.Find(".product-title").Text()
		productPrice := s.Find("span.product-price__saleprice").Text()
		fmt.Printf("Product %d: %s\n", i+1, productName)
		fmt.Printf("Price %d: %s\n", i+1, productPrice)
		fmt.Println("-------------")
	})

	return c.Status(200).JSON(fiber.Map{"result": "testing"})
}

func ScrapeGoQueryHmart(c *fiber.Ctx) error {
	url := "https://www.hmart.com/weekly-super-sale"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error fetching URL:", err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error parsing HTML:", err)
	}

	fmt.Println("doc", doc)
	products := doc.Find(".product-items")

	products.Each(func(i int, s *goquery.Selection) {
		fmt.Println("S", s)
		productName := s.Find(".product-item-link").Text()
		productPrice := s.Find(".price").Text()
		fmt.Printf("Product %d: %s\n", i+1, productName)
		fmt.Printf("Price %d: %s\n", i+1, productPrice)
		fmt.Println("-------------")
	})

	return c.Status(200).JSON(fiber.Map{"result": "testing"})
}

func ScrapeGoQueryTarget(c *fiber.Ctx) error {
	url := "https://www.target.com/c/grocery-deals/-/N-k4uyq"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error fetching URL:", err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error parsing HTML:", err)
	}

	fmt.Println("doc", doc)
	products := doc.Find(".styles__BaseContainer-sc-1ealqwt-1")

	var results []interface{}

	products.Each(func(i int, s *goquery.Selection) {
		fmt.Println("S", s)
		data := s.Find("span").Text()
		results = append(results, data)
	})

	return c.Status(200).JSON(fiber.Map{"result": results})
}
