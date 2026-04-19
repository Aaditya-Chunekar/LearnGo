
//Every Go program is made up of packages.
// comment convention is same as cpp
//Programs start running in package main.

package main
// below is what we call a factored import
import (
	"fmt"
	"math/rand"
)
// can rand be accessed without importing math/rand? 
// - nope, math/rand is reqd.
func main() {
	fmt.Println("mi favorito numero", rand.Intn(10))
}
