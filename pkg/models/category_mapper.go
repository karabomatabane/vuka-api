package models

import "strings"

// CategoryGroup represents a group of related category keywords
type CategoryGroup struct {
	Name     string   // The name of the category group (e.g., "Sports", "Politics")
	Keywords []string // Keywords that map to this category (case-insensitive)
}

// CategoryMapper handles mapping article categories to grouped categories
type CategoryMapper struct {
	groups []CategoryGroup
}

// NewCategoryMapper creates a new CategoryMapper with predefined groups
func NewCategoryMapper() *CategoryMapper {
	return &CategoryMapper{
		groups: []CategoryGroup{
			{
				Name: "Sports",
				Keywords: []string{
					"sports", "sport", "football", "soccer", "basketball",
					"rugby", "cricket", "tennis", "athletics", "olympics",
				},
			},
			{
				Name: "Politics",
				Keywords: []string{
					"politics", "political", "government", "parliament",
					"election", "minister", "president", "democracy",
				},
			},
			{
				Name: "Business",
				Keywords: []string{
					"business", "economy", "finance", "financial",
					"trade", "market", "stock", "investment", "banking",
				},
			},
			{
				Name: "Technology",
				Keywords: []string{
					"technology", "tech", "digital", "software",
					"hardware", "ai", "artificial intelligence", "computing",
				},
			},
			{
				Name: "Entertainment",
				Keywords: []string{
					"entertainment", "celebrity", "movie", "film",
					"music", "television", "tv", "show", "arts", "culture",
				},
			},
			{
				Name: "Health",
				Keywords: []string{
					"health", "medical", "medicine", "healthcare",
					"wellness", "fitness", "hospital", "doctor",
				},
			},
			{
				Name: "Education",
				Keywords: []string{
					"education", "school", "university", "college",
					"student", "learning", "academic", "teaching",
				},
			},
		},
	}
}

// NewCategoryMapperWithGroups creates a CategoryMapper with custom groups
func NewCategoryMapperWithGroups(groups []CategoryGroup) *CategoryMapper {
	return &CategoryMapper{
		groups: groups,
	}
}

// AddGroup adds a new category group to the mapper
func (cm *CategoryMapper) AddGroup(group CategoryGroup) {
	cm.groups = append(cm.groups, group)
}

// MapCategories maps an array of category strings to their grouped category names
// Returns unique category group names that match the input categories
func (cm *CategoryMapper) MapCategories(categories []string) []string {
	if len(categories) == 0 {
		return []string{}
	}

	matchedGroups := make(map[string]bool)

	for _, category := range categories {
		categoryLower := strings.ToLower(strings.TrimSpace(category))

		// Check each group for a match
		for _, group := range cm.groups {
			for _, keyword := range group.Keywords {
				if strings.Contains(categoryLower, strings.ToLower(keyword)) ||
					strings.Contains(strings.ToLower(keyword), categoryLower) {
					matchedGroups[group.Name] = true
					break
				}
			}
		}
	}

	// Convert map to slice
	result := make([]string, 0, len(matchedGroups))
	for groupName := range matchedGroups {
		result = append(result, groupName)
	}

	return result
}
