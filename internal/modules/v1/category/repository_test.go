package category

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"haircompany-shop-rest/internal/modules/v1/category/model"
	"haircompany-shop-rest/pkg/database"
	"testing"
)

func setupTestDB(t *testing.T) *database.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal("Failed to connect to test database:", err)
	}

	err = db.AutoMigrate(&model.Category{})
	if err != nil {
		t.Fatal("Failed to migrate test database:", err)
	}

	return &database.DB{DB: db}
}

func TestRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRepository(db)

	category := &model.Category{
		Name:        "Test Category",
		Description: "Test Description",
		Slug:        "test-category",
		IsActive:    true,
	}

	result, err := repo.Create(category)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.ID == 0 {
		t.Fatal("Expected category ID to be set")
	}

	if result.Name != "Test Category" {
		t.Errorf("Expected name 'Test Category', got %s", result.Name)
	}

	if result.Slug != "test-category" {
		t.Errorf("Expected slug 'test-category', got %s", result.Slug)
	}
}

func TestRepository_GetAll(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRepository(db)

	categories := []*model.Category{
		{Name: "Category 1", Slug: "category-1", IsActive: true},
		{Name: "Category 2", Slug: "category-2", IsActive: true},
		{Name: "Category 3", Slug: "category-3", IsActive: false},
	}

	for _, cat := range categories {
		_, err := repo.Create(cat)
		if err != nil {
			t.Fatalf("Failed to create test category: %v", err)
		}
	}

	result, err := repo.GetAll()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result) != 3 {
		t.Errorf("Expected 3 categories, got %d", len(result))
	}
}

func TestRepository_GetById(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRepository(db)

	category := &model.Category{
		Name:     "Test Category",
		Slug:     "test-category",
		IsActive: true,
	}

	created, err := repo.Create(category)
	if err != nil {
		t.Fatalf("Failed to create test category: %v", err)
	}

	result, err := repo.GetById(created.ID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.ID != created.ID {
		t.Errorf("Expected ID %d, got %d", created.ID, result.ID)
	}

	if result.Name != "Test Category" {
		t.Errorf("Expected name 'Test Category', got %s", result.Name)
	}

	// Тест с несуществующим ID
	_, err = repo.GetById(99999)
	if err == nil {
		t.Error("Expected error for non-existent ID")
	}
	// Для GORM проверяем, что это именно ошибка "record not found"
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("Expected gorm.ErrRecordNotFound, got %v", err)
	}
}

func TestRepository_Update(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRepository(db)

	category := &model.Category{
		Name:     "Original Name",
		Slug:     "original-slug",
		IsActive: true,
	}

	created, err := repo.Create(category)
	if err != nil {
		t.Fatalf("Failed to create test category: %v", err)
	}

	created.Name = "Updated Name"
	created.Description = "Updated Description"
	created.IsActive = false

	result, err := repo.Update(created)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.Name != "Updated Name" {
		t.Errorf("Expected name 'Updated Name', got %s", result.Name)
	}

	if result.Description != "Updated Description" {
		t.Errorf("Expected description 'Updated Description', got %s", result.Description)
	}

	if result.IsActive != false {
		t.Errorf("Expected IsActive to be false, got %t", result.IsActive)
	}
}

func TestRepository_Delete(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRepository(db)

	category := &model.Category{
		Name:     "Test Category",
		Slug:     "test-category",
		IsActive: true,
	}

	created, err := repo.Create(category)
	if err != nil {
		t.Fatalf("Failed to create test category: %v", err)
	}

	err = repo.Delete(created.ID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = repo.GetById(created.ID)
	if err == nil {
		t.Error("Expected error when getting deleted category")
	}
}

func TestRepository_GetByUniqueFields(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRepository(db)

	category := &model.Category{
		Name:     "Unique Name",
		Slug:     "unique-slug",
		IsActive: true,
	}

	_, err := repo.Create(category)
	if err != nil {
		t.Fatalf("Failed to create test category: %v", err)
	}

	result, err := repo.GetByUniqueFields("Unique Name", "")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected to find category by name")
	}

	if result.Name != "Unique Name" {
		t.Errorf("Expected name 'Unique Name', got %s", result.Name)
	}

	result, err = repo.GetByUniqueFields("", "unique-slug")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected to find category by slug")
	}

	if result.Slug != "unique-slug" {
		t.Errorf("Expected slug 'unique-slug', got %s", result.Slug)
	}

	result, err = repo.GetByUniqueFields("Non-existent", "non-existent")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result != nil {
		t.Error("Expected nil result for non-existent category")
	}
}

func TestRepository_CountChildrenByParentId(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRepository(db)

	parent := &model.Category{
		Name:     "Parent Category",
		Slug:     "parent-category",
		IsActive: true,
	}

	parentCreated, err := repo.Create(parent)
	if err != nil {
		t.Fatalf("Failed to create parent category: %v", err)
	}

	children := []*model.Category{
		{Name: "Child 1", Slug: "child-1", ParentID: &parentCreated.ID, IsActive: true},
		{Name: "Child 2", Slug: "child-2", ParentID: &parentCreated.ID, IsActive: true},
		{Name: "Child 3", Slug: "child-3", ParentID: &parentCreated.ID, IsActive: false},
	}

	for _, child := range children {
		_, err := repo.Create(child)
		if err != nil {
			t.Fatalf("Failed to create child category: %v", err)
		}
	}

	count, err := repo.CountChildrenByParentId(parentCreated.ID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if count != 3 {
		t.Errorf("Expected 3 children, got %d", count)
	}

	count, err = repo.CountChildrenByParentId(99999)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if count != 0 {
		t.Errorf("Expected 0 children for non-existent parent, got %d", count)
	}
}
