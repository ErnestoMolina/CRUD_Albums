package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Cat Black", Artist: "Felix Turner", Price: 64.99},
	{ID: "2", Title: "Happy Nation", Artist: "Monaha Kelux", Price: 89.99},
	{ID: "3", Title: "Country Cry post war", Artist: "Zara kavharaskelia", Price: 74.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumsById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	for _, a := range albums {
		if a.ID == newAlbum.ID {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No se ha creado el album, el consecutivo ya existe"})
			return
		}
	}
	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func putAlbumsById(c *gin.Context) {
	id := c.Param("id")
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	for _, a := range albums {
		if a.ID == newAlbum.ID {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No se ha modificado el album, el consecutivo ya existe"})
			return
		}
	}

	for i, a := range albums {
		if a.ID == id {
			albums[i] = newAlbum
			c.IndentedJSON(http.StatusOK, newAlbum)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No se ha modificado el album, el album no existe"})
}

func deteleAlbumsById(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			i2 := i + 1
			albums = append(albums[:i], albums[i2:]...)
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No se ha eliminado el album, el album no existe"})
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Mi Princesa Hermosa, Te amo mucho",
		})
	})

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsById)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", putAlbumsById)
	router.DELETE("/albums/:id", deteleAlbumsById)
	router.Run("localhost:8080")
}
