package blog

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"
)

var posts = []Post{
	{1, "Hello World!", "This is our first post.", time.Now()},
	{2, "Another World!", "This is our second post.", time.Now()},
}

// GetPostsHandler returns all blog posts
func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	randomFile, _ := os.Open("/dev/random")
	randomReader := bufio.NewReader(randomFile)
	randomNumber, _ := rand.Int(randomReader, big.NewInt(10))
	iterCount := int(100 + randomNumber.Int64())

	file, _ := os.OpenFile("/dev/null", 0, 644)

	for i := 0; i < iterCount; i++ {
		for j := 0; j < iterCount; j++ {
			mul := []byte(strconv.Itoa(i * j))
			file.Write(mul)
		}
	}

	file.Close()

	json.NewEncoder(w).Encode(posts)
}
