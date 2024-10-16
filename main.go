package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf/v2"
)

func main() {
	// Array of image URLs
	imageUrls := []string{
		"https://play-lh.googleusercontent.com/oFf8seyX0_gbEM0UFltiDkmiDqwy6VsD2rwMq1CR8wUSIgQRL4VdU_944gFGF0PiDpcy",
		"https://play-lh.googleusercontent.com/a3uxiTZu-JRRncw8_tGurtwkmyxYEroykNXYReM5WPbVzyxX-et4eNiqjg0XK8iuvLYX",
		"https://play-lh.googleusercontent.com/MwDsU-goAWUoF99n7zAgEkWOmB5yUuUefLpu8Luo6tPPi_Hbqu2tgZGNUbpgXVm0_ZcX",
		// Add more URLs as needed...
	}

	outputPDF := "output.pdf"
	CreatePDFWithImages(imageUrls, outputPDF)
}

// CreatePDFWithImages creates a PDF file with images from the given URLs.
func CreatePDFWithImages(imageUrls []string, outputPDF string) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	for _, url := range imageUrls {
		pdf.AddPage()
		filePath, err := DownloadAndConvertImage(url)
		if err != nil {
			fmt.Printf("Failed to download or convert image from %s: %v\n", url, err)
			continue
		}

		// Get image dimensions for scaling
		imgFile, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Failed to open downloaded image %s: %v\n", filePath, err)
			continue
		}
		img, _, err := image.Decode(imgFile)
		imgFile.Close()
		if err != nil {
			fmt.Printf("Failed to decode image %s: %v\n", filePath, err)
			continue
		}

		width := float64(img.Bounds().Dx()) * 0.264583 // Convert px to mm (assuming 96 DPI)
		height := float64(img.Bounds().Dy()) * 0.264583

		// Adjust dimensions to fit within A4 page size if needed
		maxWidth, maxHeight := 190.0, 277.0 // A4 size in mm minus margins
		if width > maxWidth {
			scale := maxWidth / width
			width *= scale
			height *= scale
		}
		if height > maxHeight {
			scale := maxHeight / height
			width *= scale
			height *= scale
		}

		// Add the image to the PDF
		pdf.Image(filePath, (210-width)/2, (297-height)/2, width, height, false, "", 0, "")
		os.Remove(filePath) // Remove the downloaded image file
	}

	err := pdf.OutputFileAndClose(outputPDF)
	if err != nil {
		fmt.Printf("Failed to save PDF: %v\n", err)
		return
	}

	fmt.Printf("PDF saved as %s\n", outputPDF)
}

// DownloadAndConvertImage downloads an image from a URL, converts it to JPEG if necessary, and returns the local file path.
func DownloadAndConvertImage(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Extract image file name from the URL
	parts := strings.Split(url, "/")
	fileName := parts[len(parts)-1] + ".jpg"
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Decode the image to ensure it is in a valid format
	img, format, err := image.Decode(response.Body)
	if err != nil {
		return "", err
	}

	// Convert PNG to JPEG if necessary
	if format == "png" {
		err = jpeg.Encode(file, img, nil)
		if err != nil {
			return "", err
		}
	} else if format == "jpeg" {
		_, err = io.Copy(file, response.Body)
		if err != nil {
			return "", err
		}
	} else {
		// For unsupported formats, return an error
		return "", fmt.Errorf("unsupported image format: %s", format)
	}

	return fileName, nil
}
