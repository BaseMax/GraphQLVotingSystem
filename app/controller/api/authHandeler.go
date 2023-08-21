package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Login User
func LoginUser(c *gin.Context) {

	user := c.Param("email")
	pass := c.Param("password")
	if (user == nil) && (pass == nil) {
		fmt.Println("Time for work")
	}
		// Check if username and password are provided
		if user == "" || pass == "" {
			fmt.Println("Please provide both username and password.")
			return
		}
	
		// Establish a database connection
		db, err := sql.Open("mysql", "username:password@tcp(hostname:port)/database")
		if err != nil {
			fmt.Println("Failed to connect to the database:", err)
			return
		}
		defer db.Close()
	
		// Query the database to check the credentials
		query := "SELECT COUNT(*) FROM users WHERE username = ? AND password = ?"
		var count int
		err = db.QueryRow(query, user, pass).Scan(&count)
		if err != nil {
			fmt.Println("Error executing query:", err)
			return
		}
	
		// Check if the username and password are valid
		if count > 0 {
			fmt.Println("Login successful!")
		} else {
			fmt.Println("Invalid username or password.")
		}
	

}

// Register User
func RegisterUser(c *gin.Context) {

}
