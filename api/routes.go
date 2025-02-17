package api

import (
	"github.com/gin-gonic/gin"
)

// das kann auch in einer eigenen Datei als helper
// hinzugefügt werden!
type StructBoilerplate struct {
	Field int64 `db:"field"`
}

func CreateRoutes() {
	r := gin.Default()

	r.GET("/api/v1/get", getData)

	r.Run(":8080")
}

func getData(c *gin.Context) {
	StructBoilerplate := []StructBoilerplate{}

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

	err := db.Select(&StructBoilerplate, "SELECT * FROM table")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, StructBoilerplate)
}
