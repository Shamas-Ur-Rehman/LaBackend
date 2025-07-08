package controllers

import (
	"Laorgaincs/models"
	"Laorgaincs/services"
	"Laorgaincs/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	file, fileHeader, _ := c.Request.FormFile("image")
	if file != nil {
		url, err := utils.UploadToCloudinary(file, fileHeader)
		if err == nil {
			product.ImageURL = url
		}
	}

	product.Name = c.PostForm("name")
	product.Description = c.PostForm("description")
	product.Category = c.PostForm("category")
	product.THC = c.PostForm("thc")
	product.CBD = c.PostForm("cbd")
	product.Strain = c.PostForm("strain")
	product.Badge = c.PostForm("badge")
	product.Price = parseFloat(c.PostForm("price"))
	product.Inventory = parseInt(c.PostForm("inventory"))

	effects := c.PostForm("effects")
	if effects != "" {
		product.Effects = strings.Split(effects, ",")
	}

	id, err := services.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "message": "Product created successfully", "name": product.Name,
		"description": product.Description,
		"category":    product.Category,
		"thc":         product.THC,
		"cbd":         product.CBD,
		"strain":      product.Strain,
		"badge":       product.Badge,
		"price":       product.Price,
		"inventory":   product.Inventory,
		"effects":     product.Effects,
		"image_url":   product.ImageURL,
	})
}

func GetAllProducts(c *gin.Context) {
	products, err := services.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := services.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	var update models.Product
	update.Name = c.PostForm("name")
	update.Description = c.PostForm("description")
	update.Category = c.PostForm("category")
	update.THC = c.PostForm("thc")
	update.CBD = c.PostForm("cbd")
	update.Strain = c.PostForm("strain")
	update.Badge = c.PostForm("badge")
	if inv, err := strconv.Atoi(c.PostForm("inventory")); err == nil {
		update.Inventory = inv
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory"})
		return
	}
	if price, err := strconv.ParseFloat(c.PostForm("price"), 64); err == nil {
		update.Price = price
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
		return
	}
	effects := c.PostForm("effects")
	if effects != "" {
		update.Effects = strings.Split(effects, ",")
	}
	file, fileHeader, err := c.Request.FormFile("image")
	if err == nil {
		defer file.Close()
		cloudURL, err := utils.UploadToCloudinary(file, fileHeader)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
			return
		}
		update.ImageURL = cloudURL
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	err = services.UpdateProduct(objID.Hex(), update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated"})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func parseFloat(val string) float64 {
	var f float64
	fmt.Sscanf(val, "%f", &f)
	return f
}

func parseInt(val string) int {
	var i int
	fmt.Sscanf(val, "%d", &i)
	return i
}
