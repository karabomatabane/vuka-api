package models

import (
	_ "embed"
	"testing"
	"vuka-api/pkg/models/db"
)

//go:embed fixtures/sample.html
var sampleHTML string

func TestExtractImagesFromHTML(t *testing.T) {
	tests := []struct {
		name           string
		htmlContent    string
		expectedCount  int
		expectedImages []db.ArticleImage
		expectError    bool
	}{
		{
			name: "Single image with alt text",
			htmlContent: `<div>
				<p>Some text</p>
				<img src="https://example.com/image1.jpg" alt="Test image">
			</div>`,
			expectedCount: 1,
			expectedImages: []db.ArticleImage{
				{
					URL:     "https://example.com/image1.jpg",
					AltText: "Test image",
					IsMain:  true,
				},
			},
			expectError: false,
		},
		{
			name: "Image without alt text",
			htmlContent: `<div>
				<img src="https://example.com/image1.jpg">
			</div>`,
			expectedCount: 1,
			expectedImages: []db.ArticleImage{
				{
					URL:     "https://example.com/image1.jpg",
					AltText: "",
					IsMain:  true,
				},
			},
			expectError: false,
		},
		{
			name: "Multiple images",
			htmlContent: `<div>
				<img src="https://example.com/image1.jpg" alt="First image">
				<p>Some text</p>
				<img src="https://example.com/image2.jpg" alt="Second image">
				<img src="https://example.com/image3.jpg" alt="Third image">
			</div>`,
			expectedCount: 3,
			expectedImages: []db.ArticleImage{
				{
					URL:     "https://example.com/image1.jpg",
					AltText: "First image",
					IsMain:  true,
				},
				{
					URL:     "https://example.com/image2.jpg",
					AltText: "Second image",
					IsMain:  false,
				},
				{
					URL:     "https://example.com/image3.jpg",
					AltText: "Third image",
					IsMain:  false,
				},
			},
			expectError: false,
		},
		{
			name:           "No images",
			htmlContent:    `<div><p>Just some text without images</p></div>`,
			expectedCount:  0,
			expectedImages: []db.ArticleImage{},
			expectError:    false,
		},
		{
			name:           "Empty HTML",
			htmlContent:    "",
			expectedCount:  0,
			expectedImages: []db.ArticleImage{},
			expectError:    false,
		},
		{
			name:          "Image with only src attribute",
			htmlContent:   `<img src="https://example.com/only-src.jpg">`,
			expectedCount: 1,
			expectedImages: []db.ArticleImage{
				{
					URL:     "https://example.com/only-src.jpg",
					AltText: "",
					IsMain:  true,
				},
			},
			expectError: false,
		},
		{
			name: "Image without src attribute is ignored",
			htmlContent: `<div>
				<img alt="No source">
				<img src="https://example.com/valid.jpg" alt="Valid">
			</div>`,
			expectedCount: 1,
			expectedImages: []db.ArticleImage{
				{
					URL:     "https://example.com/valid.jpg",
					AltText: "Valid",
					IsMain:  true,
				},
			},
			expectError: false,
		},
		{
			name: "Nested images",
			htmlContent: `<div>
				<div>
					<span>
						<img src="https://example.com/nested1.jpg" alt="Nested 1">
					</span>
				</div>
				<div>
					<img src="https://example.com/nested2.jpg" alt="Nested 2">
				</div>
			</div>`,
			expectedCount: 2,
			expectedImages: []db.ArticleImage{
				{
					URL:     "https://example.com/nested1.jpg",
					AltText: "Nested 1",
					IsMain:  true,
				},
				{
					URL:     "https://example.com/nested2.jpg",
					AltText: "Nested 2",
					IsMain:  false,
				},
			},
			expectError: false,
		},
		{
			name:          "Image with additional attributes",
			htmlContent:   `<img src="https://example.com/image.jpg" alt="Test" class="responsive" width="500" height="300">`,
			expectedCount: 1,
			expectedImages: []db.ArticleImage{
				{
					URL:     "https://example.com/image.jpg",
					AltText: "Test",
					IsMain:  true,
				},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			images, err := extractImagesFromHTML(tt.htmlContent)

			if tt.expectError && err == nil {
				t.Errorf("Expected an error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Did not expect an error but got: %v", err)
			}

			if len(images) != tt.expectedCount {
				t.Errorf("Expected %d images, got %d", tt.expectedCount, len(images))
			}

			for i, expectedImg := range tt.expectedImages {
				if i >= len(images) {
					t.Errorf("Expected image at index %d but got fewer images", i)
					break
				}
				actualImg := images[i]

				if actualImg.URL != expectedImg.URL {
					t.Errorf("Image %d: expected URL %s, got %s", i, expectedImg.URL, actualImg.URL)
				}
				if actualImg.AltText != expectedImg.AltText {
					t.Errorf("Image %d: expected AltText %s, got %s", i, expectedImg.AltText, actualImg.AltText)
				}
				if actualImg.IsMain != expectedImg.IsMain {
					t.Errorf("Image %d: expected IsMain %t, got %t", i, expectedImg.IsMain, actualImg.IsMain)
				}
			}
		})
	}
}

func TestExtractImagesFromRealWorldHTML(t *testing.T) {
	// Use the embedded sample.html file
	images, err := extractImagesFromHTML(sampleHTML)
	if err != nil {
		t.Fatalf("extractImagesFromHTML returned an error: %v", err)
	}

	// Expected images from the sample.html file (5 images with srcset attributes)
	expectedImages := []struct {
		url     string
		altText string
		isMain  bool
	}{
		{
			url:     "https://justassociates.org/wp-content/uploads/2025/03/IMG_3195-300x225.jpg",
			altText: "",
			isMain:  true,
		},
		{
			url:     "https://justassociates.org/wp-content/uploads/2025/03/IMG_3194-300x225.jpg",
			altText: "",
			isMain:  false,
		},
		{
			url:     "https://justassociates.org/wp-content/uploads/2025/03/IMG_3191-300x200.jpg",
			altText: "",
			isMain:  false,
		},
		{
			url:     "https://justassociates.org/wp-content/uploads/2025/03/IMG_3192-300x200.jpg",
			altText: "",
			isMain:  false,
		},
		{
			url:     "https://justassociates.org/wp-content/uploads/2025/03/IMG_3193-300x225.jpg",
			altText: "",
			isMain:  false,
		},
	}

	// Verify we extracted the expected number of images
	if len(images) != len(expectedImages) {
		t.Errorf("Expected %d images, got %d", len(expectedImages), len(images))
	}

	// Verify each image
	for i, expected := range expectedImages {
		if i >= len(images) {
			t.Errorf("Expected image at index %d but got fewer images", i)
			break
		}

		actualImg := images[i]

		if actualImg.URL != expected.url {
			t.Errorf("Image %d: expected URL %s, got %s", i, expected.url, actualImg.URL)
		}
		if actualImg.AltText != expected.altText {
			t.Errorf("Image %d: expected AltText '%s', got '%s'", i, expected.altText, actualImg.AltText)
		}
		if actualImg.IsMain != expected.isMain {
			t.Errorf("Image %d: expected IsMain %t, got %t", i, expected.isMain, actualImg.IsMain)
		}
	}
}
