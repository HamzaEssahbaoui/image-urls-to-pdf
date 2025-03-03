# Images to PDF Converter

A simple Go program that downloads images from specified URLs and compiles them into a single PDF document, optimizing them to fit properly on standard A4 pages.

## Features

- Downloads images from provided URLs
- Automatically converts various image formats to JPEG
- Scales images to fit A4 page dimensions while preserving aspect ratio
- Centers images on each page
- Cleans up temporary image files after PDF creation

## Requirements

- Go 1.13 or higher
- External package: github.com/jung-kurt/gofpdf/v2

## Installation

```bash
git clone https://github.com/your-username/images-to-pdf
cd images-to-pdf
go mod download
```

## Usage

1. Edit the `imageUrls` slice in `main.go` to include your image URLs:

```go
imageUrls := []string{
    "https://example.com/image1.png",
    "https://example.com/image2.jpg",
    // Add more URLs as needed...
}
```

2. Run the program:

```bash
go run main.go
```

3. The program will generate `output.pdf` in the current directory.

## How It Works

1. The program creates a new PDF document with A4 page size
2. For each image URL, it:
   - Downloads the image
   - Converts it to JPEG format if necessary
   - Calculates optimal dimensions to fit on an A4 page
   - Centers and places the image on a new page
   - Removes the temporary image file
3. Saves the final PDF to disk

## Limitations

- Currently supports PNG and JPEG formats only
- Requires active internet connection to download images

## License

[MIT License](LICENSE)
