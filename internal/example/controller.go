package example

import (
    "github.com/gin-gonic/gin"
    "net/http"
		"time"
)

func index(c *gin.Context) {
		example, err := GetExample()
		if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
						"status":  false,
						"message": "Failed to get example",
				})
				return
		}

    c.JSON(http.StatusOK, gin.H{
        "status":  true,
				"data": example,
				"message": "Berhasil mengambil data",
    })
}

func store(c *gin.Context){
	var request struct {
		Example string `json:"example" binding:"required,max=255"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Data yang dikirim tidak valid",
		})
		return
	}

	var data = Example{
		Example: request.Example,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, err := data.SaveExample()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to store example",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"data": id,
			"message": "Berhasil menyimpan data",
	})
}

func update(c *gin.Context){
	var request struct {
		Example string `json:"example" binding:"required,max=255"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Data yang dikirim tidak valid",
		})
		return
	}

	var data = Example{
		Example: request.Example,
		UpdatedAt: time.Now(),
	}

	if err := data.UpdateExampleByID(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to update example",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data": c.Param("id"),
		"message": "Berhasil mengubah data",
	})
}

func destroy(c *gin.Context){
	if err := DeleteExampleByID(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to delete example",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data": c.Param("id"),
		"message": "Berhasil menghapus data",
	})
}