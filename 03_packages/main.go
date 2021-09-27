package main

// wrap with parenthesis to import more than 1
import (
	"fmt"
	"math"

	"github.com/novelworm777/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("olleh"))
}
