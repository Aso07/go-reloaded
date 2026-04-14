package main
import(
	"fmt"
	"os"
)
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		os.Exit(1)
	}

		inputFile := os.Args[1]
		outputFile := os.Args[2]

		content, err := os.ReadFile(inputFile)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		text := string(content)
		result := process(text)

		err = os.WriteFile(outputFile, []byte(result), 0644)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	
}