package db

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ContactType represents the type of contact information
type ContactType string

// Define enum values as constants
const (
	ContactTypePhone    ContactType = "phone"
	ContactTypeEmail    ContactType = "email"
	ContactTypeFax      ContactType = "fax"
	ContactTypeAddress  ContactType = "address"
	ContactTypeLinkedIn ContactType = "linkedin"
	ContactTypeTwitter  ContactType = "twitter"
	ContactTypeOther    ContactType = "other"
)

// IsValid checks if the contact type is valid
func (ct ContactType) IsValid() bool {
	switch ct {
	case ContactTypePhone, ContactTypeEmail, ContactTypeFax,
		ContactTypeAddress, ContactTypeLinkedIn, ContactTypeTwitter, ContactTypeOther:
		return true
	}
	return false
}

// String returns the string representation
func (ct ContactType) String() string {
	return string(ct)
}

// GetAllContactTypes returns all valid contact types
func GetAllContactTypes() []ContactType {
	return []ContactType{
		ContactTypePhone,
		ContactTypeEmail,
		ContactTypeFax,
		ContactTypeAddress,
		ContactTypeLinkedIn,
		ContactTypeTwitter,
		ContactTypeOther,
	}
}

type ContactInfo struct {
	Model
	Type             ContactType `json:"type" gorm:"not null;type:varchar(20);check:type IN ('phone','email','fax','address','linkedin','twitter', 'other')"`
	Description      string      `json:"description"`
	Value            string      `json:"value" gorm:"not null"`
	DirectoryEntryID uuid.UUID   `json:"directoryEntryId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// BeforeCreate is a GORM hook that validates the contact type before creating
func (c *ContactInfo) BeforeCreate(tx *gorm.DB) error {
	if !c.Type.IsValid() {
		return fmt.Errorf("invalid contact type: %s", c.Type)
	}
	return c.validateContactValue()
}

// BeforeUpdate is a GORM hook that validates the contact type before updating
func (c *ContactInfo) BeforeUpdate(tx *gorm.DB) error {
	if !c.Type.IsValid() {
		return fmt.Errorf("invalid contact type: %s", c.Type)
	}
	return c.validateContactValue()
}

// validateContactValue performs basic validation on contact values
func (c *ContactInfo) validateContactValue() error {
	if c.Value == "" {
		return fmt.Errorf("contact value cannot be empty")
	}

	// Add specific validation based on contact type
	switch c.Type {
	case ContactTypeEmail:
		// Basic email validation (you might want to use a proper email validation library)
		if !strings.Contains(c.Value, "@") {
			return fmt.Errorf("invalid email format")
		}
	case ContactTypePhone:
		// Basic phone validation - you can make this more sophisticated
		if len(c.Value) < 7 {
			return fmt.Errorf("phone number too short")
		}
	}

	return nil
}
