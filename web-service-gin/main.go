package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"

	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"log"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

type RegisterRequest struct {
	NRIC          string `json:"nric"`
	WalletAddress string `json:"walletAddress"`
}

type RegisterResponse struct {
	Receipt string `json:"receipt"`
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

var db *sql.DB

func handleRegister(c *gin.Context) {
	fmt.Println("Inside handleRegister")
	// Parse request body
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Req:", req);

	// Check if NRIC is unique
	if !isNRICUnique(req.NRIC) {
		c.JSON(http.StatusConflict, gin.H{"error": "NRIC already exists"})
		return
	}
	fmt.Println("NRIC Unique")

	// Check if wallet address is already associated with another NRIC
	if !isWalletUnique(req.WalletAddress) {
		c.JSON(http.StatusConflict, gin.H{"error": "Wallet address already associated with another NRIC"})
		return
	}
	fmt.Println("Wallet Unique")

	// Generate a random receipt hash using crypto/rand
	receiptHash := make([]byte, 32)
	_, err := rand.Read(receiptHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate receipt hash"})
		return
	}
	receipt := hex.EncodeToString(receiptHash)

	// Generate receipt using SHA256 hash of request body

	// Insert record into database
	// _, err := db.Exec("INSERT INTO registrations (nric, wallet_address, receipt) VALUES ($1, $2, $3)", req.NRIC, req.WalletAddress, receipt)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// Return receipt in response
	res := RegisterResponse{Receipt: receipt}
	c.JSON(http.StatusOK, res)
}

type Registration struct {
    id   int    `json:"id"`
    nric string `json:"nric"`
    walletAddress string `json:"wallet_address"`
}

func isNRICUnique(nric string) bool {
	var count int
	fmt.Println("NRIC->", nric)
	err := db.QueryRow("SELECT COUNT(*) FROM registration where nric = ?", nric).Scan(&count)
	if err != nil {
		fmt.Println("Err at finding nric")
		log.Fatal(err)
	}
	fmt.Println("NRIC Count:", count)
	return count == 0
}

func isWalletUnique(walletAddress string) bool {
	var count int
	fmt.Println("WalletAddress:", walletAddress)
	err := db.QueryRow("SELECT COUNT(*) FROM registration where wallet_address = ?", walletAddress).Scan(&count)
	if err != nil {
		fmt.Println("Err at finding walletAddress")
		log.Fatal(err)
	}
	fmt.Println("WalletAddress Count:", count)
	return count == 0
}

type City struct {
    Id         int
    Name       string
    Population int
}

func initDb() {
	connection, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	db = connection
	fmt.Println("InitDb ", db)
}

func main() {

	initDb()

    res, err := db.Query("SELECT * FROM city")
	defer res.Close()
	if err != nil {
        log.Fatal(err)
    }

	for res.Next() {
        var city City
        err := res.Scan(&city.Id, &city.Name, &city.Population)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%v\n", city)
    }
	// defer db.Close()

	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.POST("/register", handleRegister)
	router.Run("localhost:8080")
}
