package main

//gravatar.com
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func getGravatarHash(email string) string {
	email = strings.TrimSpace(email)

	email = strings.ToLower(email)

	h := md5.New()
	io.WriteString(h, email)
	finalBytes := h.Sum(nil)
	finalString := hex.EncodeToString(finalBytes)
	return finalString

}
func main() {

	fmt.Fprint(os.Stderr, "Please enter your email ")
	var email string
	fmt.Scanln(&email)

	gravatarHash := getGravatarHash(email)
	fmt.Fprint(os.Stderr, "Enter your name")
	var name string
	fmt.Scanln(&name)
	fmt.Println(`<!DOCTYPE html>
  <html>
    <head>
      <script>
          console.log("Hello")
       </script>
    </head>
    <body>
     <img src="http://www.gravatar.com/avatar/` + gravatarHash + `?d=identicon">
     <br>
     <h1>` + name + `</h1>
  </body>
  </html>`)
}
