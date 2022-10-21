package main

import (
	// "encoding/json"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

// ProductItem represents product document in mongodb
type ProductItem struct {
	ID        string    `json:"id"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Product   *Product  `json:"product"`
}

// Product represents a single product data
type Product struct {
	Name        string `json:"name"`
	ImageURL    string `json:"imageURL"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Reviews     int    `json:"totalReviews"`
}

// amzonScraper scarpes the url and fetchs the product details
func amzonScraper(ctx *gin.Context) {
	cly := colly.NewCollector()
	Product := Product{}
	cly.OnHTML("div[id=centerCol]", func(h *colly.HTMLElement) {
		Product.Name = string(h.ChildText("span[id=productTitle]"))
		Reviews := string(h.ChildText("span[id=acrCustomerReviewText]"))
		// Reviews will be cleaned to get only numbers
		totalReviewStr := regexp.MustCompile("\\D").ReplaceAllString(Reviews, "")
		TotalReviewsInt, err := strconv.Atoi(totalReviewStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}
		Product.Reviews = TotalReviewsInt
		h.ForEach("table.a-lineitem.a-align-top tr", func(_ int, el *colly.HTMLElement) {
			if el.ChildText("td.a-color-secondary.a-size-base.a-text-right.a-nowrap") == "Price:" {
				Product.Price = string(el.ChildText("span.a-price.a-text-price.a-size-medium.apexPriceToPay > span.a-offscreen"))
			}
		})
		//Hack around way to get the price for special cases when Price text is not defined
		if Product.Price == "" {
			priceSymbol := string(h.ChildText("span.a-price-symbol"))
			str := string(h.ChildText("span.a-offscreen"))
			price := strings.Split(str, "$")[1]
			Product.Price = priceSymbol + price
		}
	})
	cly.OnHTML("div[id=imageBlock]", func(h *colly.HTMLElement) {
		Product.ImageURL = string(h.ChildAttr("img.a-dynamic-image", "src"))
	})

	cly.OnHTML("div[id=feature-bullets]", func(h *colly.HTMLElement) {
		Product.Description = string(h.ChildText("span.a-list-item"))
	})

	url := ctx.Query("url")
	if url == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
	}
	cly.Visit(url)
	requestBody, err := json.Marshal(ProductItem{Url: url, CreatedAt: time.Now(), UpdatedAt: time.Now(), Product: &Product})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	//Calls the api to load data into the database
	response, err := http.Post("http://loader-service:3001/products", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	respObj := ProductItem{}
	json.Unmarshal(responseData, &respObj)
	ctx.JSON(http.StatusCreated,
		gin.H{
			"status_code": http.StatusCreated,
			"result":      respObj,
		})
}

func main() {
	r := gin.Default()
	r.POST("/products", amzonScraper)
	r.Run("0.0.0.0:3001")
}
