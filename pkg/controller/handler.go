package controller

import (
	"fmt"
	"gigmile/pkg/database"
	"gigmile/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"log"
	"net/http"
	"strings"
	"time"
)

type NewHttp struct {
	Db    database.DB
	Route *gin.Engine
}

func New(model database.DB) *NewHttp {
	return &NewHttp{Db: model}
}

// Handler is the interface that wraps the basic Handle method.

func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func (n *NewHttp) Serve() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		switch path {
		case "/countries":
			n.GetCountries()
			break
		case "/countries/:id":
			if c.Request.Method == "GET" {
				n.GetCountry()
			} else if c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
				n.UpdateCountry()
			}
			break
		case "/country":
			n.CreateCountry()
		}
		c.Next()
	}
}

func (h *NewHttp) GetCountries() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Method not allowed"})
			return
		} else {

			result := h.Db.GetCountries()

			c.JSON(http.StatusOK, gin.H{"message": "here are the countries", "data": result})
		}

	}

}

func (h *NewHttp) GetCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Method not allowed"})
			return
		} else {
			id := strings.TrimSpace(c.Param("id"))

			result, err := h.Db.FindCountryByID(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "country found", "data": result})
		}
	}
}

func (h *NewHttp) CreateCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		var country model.Country
		if c.Request.Method != "POST" {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Method not allowed"})
			return
		} else {
			err := c.BindJSON(&country)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
				return
			}
			id, _ := uuid.NewV4()
			country.ID = id.String()
			country.Name = strings.TrimSpace(strings.ToUpper(country.Name))
			country.ShortName = strings.TrimSpace(strings.ToUpper(country.ShortName))
			country.Continent = strings.TrimSpace(strings.ToUpper(country.Continent))
			country.CreatedAt = time.Now().String()
			country.UpdatedAt = time.Now().String()

			result, err := h.Db.CreateCountry(&country)
			c.JSON(http.StatusOK, gin.H{"message": result})
		}
	}

}

func (h *NewHttp) UpdateCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "PATCH" {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Method not allowed1"})
			return
		} else {
			id := strings.TrimSpace(c.Param("id"))
			var country model.Country
			c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 100)

			country.Name = c.PostForm("name")
			country.ShortName = c.PostForm("short_name")
			country.Continent = c.PostForm("continent")
			country.IsOperational = c.GetBool("is_operational")

			err := c.BindJSON(&country)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "unable to bind json"})
				return
			}

			res, err := h.Db.FindCountryByID(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
				return
			}

			log.Println(country.IsOperational)

			country.ID = id
			if country.Name == "" {
				country.Name = res.Name
			} else {
				country.Name = strings.TrimSpace(strings.ToUpper(country.Name))
			}
			if country.ShortName == "" {
				country.ShortName = res.ShortName
			} else {
				country.ShortName = strings.TrimSpace(strings.ToUpper(country.ShortName))
			}
			if country.Continent == "" {
				country.Continent = res.Continent
			} else {
				country.Continent = strings.TrimSpace(strings.ToUpper(country.Continent))
			}
			country.CreatedAt = res.CreatedAt
			country.UpdatedAt = time.Now().String()

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "unable to decode json"})
				return
			}

			update := map[string]interface{}{
				"name":           country.Name,
				"short_name":     country.ShortName,
				"continent":      country.Continent,
				"is_operational": country.IsOperational,
				"created_at":     country.CreatedAt,
				"updated_at":     country.UpdatedAt,
			}

			err = h.Db.UpdateCountry(id, update)

			c.JSON(http.StatusOK, gin.H{"message": "Update Country"})
		}
	}
}
