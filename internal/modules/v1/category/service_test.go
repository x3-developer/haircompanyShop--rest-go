package category

import (
	"context"
	"errors"
	"haircompany-shop-rest/internal/modules/v1/category/dto"
	"haircompany-shop-rest/internal/modules/v1/category/model"
	"haircompany-shop-rest/pkg/response"
	"mime/multipart"
	"sync"
	"testing"
	"time"
)

type mockRepository struct {
	categories map[uint]*model.Category
	nextID     uint
}

func newMockRepository() *mockRepository {
	return &mockRepository{
		categories: make(map[uint]*model.Category),
		nextID:     1,
	}
}

func (m *mockRepository) Create(category *model.Category) (*model.Category, error) {
	if category == nil {
		return nil, errors.New("category is nil")
	}
	category.ID = m.nextID
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	m.categories[m.nextID] = category
	m.nextID++
	return category, nil
}

func (m *mockRepository) GetAll() ([]*model.Category, error) {
	var categories []*model.Category
	for _, cat := range m.categories {
		categories = append(categories, cat)
	}
	return categories, nil
}

func (m *mockRepository) GetById(id uint) (*model.Category, error) {
	if category, exists := m.categories[id]; exists {
		return category, nil
	}
	return nil, errors.New("category not found")
}

func (m *mockRepository) Update(category *model.Category) (*model.Category, error) {
	if category == nil {
		return nil, errors.New("category is nil")
	}
	if _, exists := m.categories[category.ID]; !exists {
		return nil, errors.New("category not found")
	}
	category.UpdatedAt = time.Now()
	m.categories[category.ID] = category
	return category, nil
}

func (m *mockRepository) Delete(id uint) error {
	if _, exists := m.categories[id]; !exists {
		return errors.New("category not found")
	}
	delete(m.categories, id)
	return nil
}

func (m *mockRepository) GetByUniqueFields(name, slug string) (*model.Category, error) {
	for _, cat := range m.categories {
		if (name != "" && cat.Name == name) || (slug != "" && cat.Slug == slug) {
			return cat, nil
		}
	}
	return nil, nil
}

func (m *mockRepository) CountChildrenByParentId(parentId uint) (int64, error) {
	count := int64(0)
	for _, cat := range m.categories {
		if cat.ParentID != nil && *cat.ParentID == parentId {
			count++
		}
	}
	return count, nil
}

// Мок файлового сервиса
type mockFileService struct {
	moveToPermCalled bool
	moveToPermError  error
}

func newMockFileService() *mockFileService {
	return &mockFileService{}
}

func (m *mockFileService) MoveToPermanent(filenames []string, directory string) error {
	m.moveToPermCalled = true
	return m.moveToPermError
}

func (m *mockFileService) Delete(filenames []string, folder string) error {
	return nil
}

func (m *mockFileService) SaveToTemp(file multipart.File, filename string) (string, error) {
	return "temp_" + filename, nil
}

func (m *mockFileService) CleanTemp() {
}

func setupTestService() (Service, *mockRepository, *mockFileService) {
	mockRepo := newMockRepository()
	mockFS := newMockFileService()
	ctx := context.Background()
	wg := &sync.WaitGroup{}

	service := NewService(mockRepo, mockFS, ctx, wg)
	return service, mockRepo, mockFS
}

func TestService_Create_Success(t *testing.T) {
	service, mockRepo, _ := setupTestService()

	createDto := dto.CreateDTO{
		Name:            "Test Category",
		Description:     "Test Description",
		Image:           "test.jpg",
		HeaderImage:     "header.jpg",
		Slug:            "test-category",
		SortIndex:       100,
		SeoTitle:        "Test SEO Title",
		SeoDescription:  "Test SEO Description",
		SeoKeys:         "test, seo, keys",
		IsActive:        true,
		IsShade:         false,
		IsVisibleInMenu: true,
		IsVisibleOnMain: false,
	}

	result, validationErrors, err := service.Create(createDto)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(validationErrors) > 0 {
		t.Fatalf("Expected no validation errors, got %d", len(validationErrors))
	}

	if result == nil {
		t.Fatal("Expected result to be not nil")
	}

	if result.Name != "Test Category" {
		t.Errorf("Expected name 'Test Category', got %s", result.Name)
	}

	if result.Slug != "test-category" {
		t.Errorf("Expected slug 'test-category', got %s", result.Slug)
	}

	// Проверяем, что категория была создана в репозитории
	if len(mockRepo.categories) != 1 {
		t.Errorf("Expected 1 category in repository, got %d", len(mockRepo.categories))
	}
}

func TestService_Create_DuplicateName(t *testing.T) {
	service, mockRepo, _ := setupTestService()

	// Создаем категорию с дублирующимся именем
	existingCategory := &model.Category{
		Name:     "Existing Category",
		Slug:     "existing-category",
		IsActive: true,
	}
	_, err := mockRepo.Create(existingCategory)
	if err != nil {
		return
	}

	createDto := dto.CreateDTO{
		Name:            "Existing Category",
		Description:     "Test Description",
		Image:           "test.jpg",
		HeaderImage:     "header.jpg",
		Slug:            "new-slug",
		SortIndex:       100,
		IsActive:        true,
		IsShade:         false,
		IsVisibleInMenu: true,
		IsVisibleOnMain: false,
	}

	result, validationErrors, err := service.Create(createDto)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result != nil {
		t.Error("Expected result to be nil due to validation error")
	}

	if len(validationErrors) == 0 {
		t.Fatal("Expected validation errors for duplicate name")
	}

	found := false
	for _, validationError := range validationErrors {
		if validationError.Field == "name" && validationError.ErrorCode == string(response.NotUnique) {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation error for duplicate name")
	}
}

func TestService_Create_InvalidParentId(t *testing.T) {
	service, _, _ := setupTestService()

	parentId := uint(999)
	createDto := dto.CreateDTO{
		Name:            "Test Category",
		Description:     "Test Description",
		Image:           "test.jpg",
		HeaderImage:     "header.jpg",
		Slug:            "test-category",
		ParentID:        &parentId,
		SortIndex:       100,
		IsActive:        true,
		IsShade:         false,
		IsVisibleInMenu: true,
		IsVisibleOnMain: false,
	}

	result, validationErrors, err := service.Create(createDto)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result != nil {
		t.Error("Expected result to be nil due to validation error")
	}

	if len(validationErrors) == 0 {
		t.Fatal("Expected validation errors for invalid parent ID")
	}

	found := false
	for _, validationError := range validationErrors {
		if validationError.Field == "parentId" && validationError.ErrorCode == string(response.NotFound) {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation error for invalid parent ID")
	}
}

func TestService_GetAll_Success(t *testing.T) {
	service, mockRepo, _ := setupTestService()

	// Создаем тестовые категории
	categories := []*model.Category{
		{Name: "Category 1", Slug: "category-1", IsActive: true},
		{Name: "Category 2", Slug: "category-2", IsActive: true},
		{Name: "Category 3", Slug: "category-3", IsActive: false},
	}

	for _, cat := range categories {
		_, err := mockRepo.Create(cat)
		if err != nil {
			return
		}
	}

	result, err := service.GetAll()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result) != 3 {
		t.Errorf("Expected 3 categories, got %d", len(result))
	}
}

func TestService_GetById_Success(t *testing.T) {
	service, mockRepo, _ := setupTestService()

	// Создаем тестовую категорию
	category := &model.Category{
		Name:     "Test Category",
		Slug:     "test-category",
		IsActive: true,
	}
	created, _ := mockRepo.Create(category)

	result, err := service.GetById(created.ID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result to be not nil")
	}

	if result.Name != "Test Category" {
		t.Errorf("Expected name 'Test Category', got %s", result.Name)
	}
}

func TestService_GetById_NotFound(t *testing.T) {
	service, _, _ := setupTestService()

	result, err := service.GetById(999)
	if err == nil {
		t.Error("Expected error for non-existent category")
	}

	if result != nil {
		t.Error("Expected result to be nil for non-existent category")
	}
}

func TestService_Update_Success(t *testing.T) {
	service, mockRepo, _ := setupTestService()

	// Создаем тестовую категорию
	category := &model.Category{
		Name:     "Original Name",
		Slug:     "original-slug",
		IsActive: true,
	}
	created, _ := mockRepo.Create(category)

	name := "Updated Name"
	description := "Updated Description"
	image := "updated.jpg"
	headerImage := "updated_header.jpg"
	slug := "updated-slug"
	sortIndex := 200
	seoTitle := "Updated SEO Title"
	seoDescription := "Updated SEO Description"
	seoKeys := "updated, seo, keys"
	isActive := false
	isShade := true
	isVisibleInMenu := false
	isVisibleOnMain := true

	updateDto := dto.UpdateDTO{
		Name:            &name,
		Description:     &description,
		Image:           &image,
		HeaderImage:     &headerImage,
		Slug:            &slug,
		SortIndex:       &sortIndex,
		SeoTitle:        &seoTitle,
		SeoDescription:  &seoDescription,
		SeoKeys:         &seoKeys,
		IsActive:        &isActive,
		IsShade:         &isShade,
		IsVisibleInMenu: &isVisibleInMenu,
		IsVisibleOnMain: &isVisibleOnMain,
	}

	result, validationErrors, err := service.Update(created.ID, updateDto)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(validationErrors) > 0 {
		t.Fatalf("Expected no validation errors, got %d", len(validationErrors))
	}

	if result == nil {
		t.Fatal("Expected result to be not nil")
	}

	if result.Name != "Updated Name" {
		t.Errorf("Expected name 'Updated Name', got %s", result.Name)
	}

	if result.Description != "Updated Description" {
		t.Errorf("Expected description 'Updated Description', got %s", result.Description)
	}
}

func TestService_Delete_Success(t *testing.T) {
	service, mockRepo, _ := setupTestService()

	category := &model.Category{
		Name:     "Test Category",
		Slug:     "test-category",
		IsActive: true,
	}
	created, _ := mockRepo.Create(category)

	result, childrenCount, err := service.Delete(created.ID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result to be not nil")
	}

	if childrenCount != 0 {
		t.Errorf("Expected 0 children count, got %d", childrenCount)
	}

	if len(mockRepo.categories) != 0 {
		t.Errorf("Expected 0 categories in repository, got %d", len(mockRepo.categories))
	}
}

func TestService_Delete_NotFound(t *testing.T) {
	service, _, _ := setupTestService()

	result, _, err := service.Delete(999)
	if err == nil {
		t.Error("Expected error for non-existent category")
	}

	if result != nil {
		t.Error("Expected result to be nil for non-existent category")
	}
}
