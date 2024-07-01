// this program was created by ChatGpt and released under the GNU General Public License.
//
//

package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if a filename was provided
	if len(os.Args) < 2 {
		fmt.Println(
			"Usage: touchgo filename ( file name will be prefixed with .go) \n---------------------------------------------------------------",
		)
		os.Exit(1)
	}

	// Get the filename from the command-line arguments
	filename := os.Args[1] + ".go"

	// Check if the file already exists
	if _, err := os.Stat(filename); err == nil {
		fmt.Printf("File %s already exists.\n", filename)
		os.Exit(1)
	}

	// Define the basic Go source code template
	template := `package main

import (
  "fmt"
)

func main() {
    // add you code here :-)
    fmt.Println("template was created by touchgo")
}
`

	// Write the template to the specified file
	err := os.WriteFile(filename, []byte(template), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}

	fmt.Printf("Created %s with a basic Go source code template.\n", filename)
}
