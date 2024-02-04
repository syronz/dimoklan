package main

import (
	"fmt"
	"os"
)

func generateHexColors() []string {
	var colors []string

	// Generate all possible combinations of two hexadecimal digits
	for i := 0; i <= 0xFF; i++ {
		for j := 0; j <= 0xFF; j++ {
			for k := 0; k <= 0xFF; k++ {
				color := fmt.Sprintf("%02X%02X%02X", i, j, k)
				colors = append(colors, color)
			}
		}
	}

	return colors
}

func saveToFile(colors []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, color := range colors {
		_, err := file.WriteString(color + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	colors := generateHexColors()

	err := saveToFile(colors, "hex_colors.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Hex colors saved to hex_colors.txt")
	}
}
