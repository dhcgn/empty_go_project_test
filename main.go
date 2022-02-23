// package main holds only imports of another product
package main

import (
	"fmt"

	_ "code.cloudfoundry.org/clock"

	_ "github.com/denisbrodbeck/machineid"

	_ "github.com/dhcgn/dockerdetector"

	_ "github.com/fatih/color"

	_ "github.com/gofrs/uuid"

	_ "github.com/gorilla/mux"

	_ "github.com/joho/godotenv"

	_ "github.com/julienschmidt/httprouter"

	_ "github.com/mattn/go-colorable"

	_ "github.com/microsoft/ApplicationInsights-Go/appinsights"

	_ "github.com/patrickmn/go-cache"

	_ "github.com/sirupsen/logrus"

	_ "github.com/stretchr/testify/assert"

	_ "github.com/urfave/cli/v2"

	_ "golang.org/x/crypto/ed25519"

	_ "golang.org/x/sys/cpu"

	_ "golang.org/x/text"

	_ "gopkg.in/check.v1"

	_ "gopkg.in/yaml.v3"

	_ "software.sslmate.com/src/go-pkcs12"
)

func main() {
	fmt.Println("hello world")
}
