package main

import (
	"fmt"
	"vuka-api/pkg/config"
	"vuka-api/pkg/models/db"
)

func init() {
	config.LoadEnvVariables()
	config.Connect()
}

// Run to create the tables in the database
func main() {
	config.LoadEnvVariables()
	config.Connect()
	fmt.Println("Migrating database...")
	err := config.GetDB().AutoMigrate(
		&db.Source{},
		&db.Section{},
		&db.Category{},
		&db.Article{},
		&db.ArticleImage{},
		&db.Region{},
		&db.DirectoryCategory{},
		&db.DirectoryEntry{},
		&db.Permission{},
		&db.Role{},
		&db.User{},
		&db.RoleSectionPermission{},
		&db.UserDirectoryMeta{},
	)
	if err != nil {
		fmt.Printf("Migration failed: %v\n", err)
		return
	}
	fmt.Println("Migration completed successfully!")
}
