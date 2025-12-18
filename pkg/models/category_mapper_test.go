package models

import (
	"testing"
)

func TestCategoryMapper_MapCategories(t *testing.T) {
	mapper := NewCategoryMapper()

	tests := []struct {
		name       string
		categories []string
		expected   []string
	}{
		{
			name:       "Single sports category",
			categories: []string{"Football"},
			expected:   []string{"Sports"},
		},
		{
			name:       "Multiple sports keywords",
			categories: []string{"Soccer", "Basketball"},
			expected:   []string{"Sports"},
		},
		{
			name:       "Mixed categories",
			categories: []string{"Politics", "Sports", "Technology"},
			expected:   []string{"Politics", "Sports", "Technology"},
		},
		{
			name:       "Case insensitive matching",
			categories: []string{"FOOTBALL", "politics", "BuSiNeSs"},
			expected:   []string{"Sports", "Politics", "Business"},
		},
		{
			name:       "Partial keyword matching",
			categories: []string{"Political News", "Financial Markets"},
			expected:   []string{"Politics", "Business"},
		},
		{
			name:       "No matching categories",
			categories: []string{"Random", "Unknown", "Other"},
			expected:   []string{},
		},
		{
			name:       "Empty categories",
			categories: []string{},
			expected:   []string{},
		},
		{
			name:       "Categories with whitespace",
			categories: []string{"  Football  ", "  Politics  "},
			expected:   []string{"Sports", "Politics"},
		},
		{
			name:       "Entertainment keywords",
			categories: []string{"Movie", "Music", "Celebrity"},
			expected:   []string{"Entertainment"},
		},
		{
			name:       "Health and wellness",
			categories: []string{"Medical News", "Fitness Tips"},
			expected:   []string{"Health"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mapper.MapCategories(tt.categories)

			// Check length
			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d categories, got %d. Expected: %v, Got: %v",
					len(tt.expected), len(result), tt.expected, result)
				return
			}

			// Check that all expected categories are in the result
			resultMap := make(map[string]bool)
			for _, cat := range result {
				resultMap[cat] = true
			}

			for _, expected := range tt.expected {
				if !resultMap[expected] {
					t.Errorf("Expected category '%s' not found in result: %v", expected, result)
				}
			}
		})
	}
}

func TestCategoryMapper_CustomGroups(t *testing.T) {
	customGroups := []CategoryGroup{
		{
			Name:     "African News",
			Keywords: []string{"africa", "african", "south africa", "nigeria", "kenya"},
		},
		{
			Name:     "Local News",
			Keywords: []string{"local", "community", "regional"},
		},
	}

	mapper := NewCategoryMapperWithGroups(customGroups)

	tests := []struct {
		name       string
		categories []string
		expected   []string
	}{
		{
			name:       "African news category",
			categories: []string{"South Africa News", "African Politics"},
			expected:   []string{"African News"},
		},
		{
			name:       "Local news category",
			categories: []string{"Community Events"},
			expected:   []string{"Local News"},
		},
		{
			name:       "Mixed custom categories",
			categories: []string{"Kenya Updates", "Regional News"},
			expected:   []string{"African News", "Local News"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mapper.MapCategories(tt.categories)

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d categories, got %d. Expected: %v, Got: %v",
					len(tt.expected), len(result), tt.expected, result)
				return
			}

			resultMap := make(map[string]bool)
			for _, cat := range result {
				resultMap[cat] = true
			}

			for _, expected := range tt.expected {
				if !resultMap[expected] {
					t.Errorf("Expected category '%s' not found in result: %v", expected, result)
				}
			}
		})
	}
}

func TestCategoryMapper_AddGroup(t *testing.T) {
	mapper := NewCategoryMapper()
	
	newGroup := CategoryGroup{
		Name:     "Science",
		Keywords: []string{"science", "research", "study", "discovery"},
	}
	
	mapper.AddGroup(newGroup)

	categories := []string{"Science News", "Research"}
	result := mapper.MapCategories(categories)

	if len(result) != 1 || result[0] != "Science" {
		t.Errorf("Expected ['Science'], got %v", result)
	}
}
