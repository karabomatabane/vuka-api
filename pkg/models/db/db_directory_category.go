package db

type DirectoryCategory struct {
	Model
	Name        string           `json:"name"`
	Directories []DirectoryEntry `json:"entries" gorm:"foreignKey:CategoryID"`
}
