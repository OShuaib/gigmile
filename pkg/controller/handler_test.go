package controller

import (
	"encoding/json"
	db "gigmile/pkg/mocks"
	"gigmile/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func Loader(mdb *db.MockDB) *gin.Engine {
	router := gin.Default()
	newhttp := &NewHttp{
		Db:    mdb,
		Route: router,
	}
	newhttp.Routes(router)
	return router
}

func TestNewHttp_CreateCountry(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mdb := db.NewMockDB(ctrl)
	router := Loader(mdb)
	form := model.Country{
		Name:          "nigeria",
		ShortName:     "ngn",
		Continent:     "africa",
		IsOperational: false,
		CreatedAt:     time.Now().String(),
		UpdatedAt:     time.Now().String(),
	}

	jsondata, err := json.Marshal(form)
	if err != nil {
		t.Fail()
		return
	}

	t.Run("testing no error", func(t *testing.T) {
		mdb.EXPECT().CreateCountry(gomock.Any()).Return(&form, nil)
		request, err := http.NewRequest(http.MethodPost, "/country", strings.NewReader(string(jsondata)))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}

	})
}

func TestNewHttp_GetCountries(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mdb := db.NewMockDB(ctrl)
	router := Loader(mdb)

	countries := []model.Country{
		{
			ID:            "1",
			Name:          "Qatar",
			ShortName:     "qa",
			Continent:     "Asia",
			IsOperational: true,
			CreatedAt:     time.Now().String(),
			UpdatedAt:     time.Now().String(),
		},
		{
			ID:            "2",
			Name:          "niger",
			ShortName:     "ni",
			Continent:     "africa",
			IsOperational: false,
			CreatedAt:     time.Now().String(),
			UpdatedAt:     time.Now().String(),
		},
	}

	t.Run("testing getting all countries success", func(t *testing.T) {
		mdb.EXPECT().GetCountries().Return(countries)
		request, err := http.NewRequest(http.MethodGet, "/countries", nil)
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}

	})
}

func TestNewHttp_GetCountry(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mdb := db.NewMockDB(ctrl)
	router := Loader(mdb)

	country := model.Country{
		ID:            "1",
		Name:          "nigeria",
		ShortName:     "ngn",
		IsOperational: false,
		CreatedAt:     time.Now().String(),
		UpdatedAt:     time.Now().String(),
	}

	mdb.EXPECT().FindCountryByID(gomock.Any()).Return(&country, nil)
	t.Run("testing getting country success", func(t *testing.T) {

		request, err := http.NewRequest(http.MethodGet, "/countries/1", nil)
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}

	})

}
