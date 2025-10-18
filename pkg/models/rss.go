package models

import (
	"encoding/xml"
	"golang.org/x/net/html"
	"regexp"
	"strings"
	"time"
	"vuka-api/pkg/models/db"
)

// The root of the RSS feed
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

// The channel element containing feed metadata and items
type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Language    string   `xml:"language"`
	Items       []Item   `xml:"item"`
}

// An item represents a single article or post in the feed
type Item struct {
	XMLName        xml.Name `xml:"item"`
	Title          string   `xml:"title"`
	Link           string   `xml:"link"`
	Description    string   `xml:"description"`
	GUID           string   `xml:"guid"`
	Author         string   `xml:"author"`
	PubDate        string   `xml:"pubDate"` // You might want to parse this into a time.Time later
	Categories     []string `xml:"category"`
	ContentEncoded string   `xml:"http://purl.org/rss/1.0/modules/content/ encoded"`
}

func (feed *Item) ToArticle(language string) (*db.Article, error) {
	publishedAt, err := time.Parse(time.RFC1123Z, feed.PubDate)
	if err != nil {
		return nil, err
	}

	images, _ := extractImagesFromHTML(feed.Description)

	// Remove img tags from description
	re := regexp.MustCompile(`<img[^>]*>`)
	summary := re.ReplaceAllString(feed.Description, "")

	article := &db.Article{
		Title:       feed.Title,
		Language:    language,
		OriginalUrl: feed.Link,
		Summary:     summary, // Use cleaned summary
		ContentBody: feed.ContentEncoded,
		PublishedAt: publishedAt,
		IsFeatured:  false,
		Images:      images,
	}

	return article, nil
}

func extractImagesFromHTML(htmlContent string) ([]db.ArticleImage, error) {
	var images []db.ArticleImage
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			var src, alt string
			for _, a := range n.Attr {
				if a.Key == "src" {
					src = a.Val
				}
				if a.Key == "alt" {
					alt = a.Val
				}
			}
			if src != "" {
				isMain := len(images) == 0
				images = append(images, db.ArticleImage{
					URL:     src,
					AltText: alt,
					IsMain:  isMain,
				})
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return images, nil
}