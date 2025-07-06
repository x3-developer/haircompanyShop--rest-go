package category

import (
	"bytes"
	"encoding/json"
	"fmt"
	"haircompany-shop-rest/internal/modules/v1/category/dto"
	"haircompany-shop-rest/pkg/response"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockService struct {
	categories                     map[uint]*dto.ResponseDTO
	nextId                         uint
	validationErrors               []response.ErrorField
	shouldReturnError              bool
	deleteReturnsCategoryWithError bool
}

func newMockService() *mockService {
	return &mockService{
		categories: make(map[uint]*dto.ResponseDTO),
		nextId:     1,
	}
}

func (m *mockService) Create(createDto dto.CreateDTO) (*dto.ResponseDTO, []response.ErrorField, error) {
	if m.shouldReturnError {
		return nil, nil, fmt.Errorf("service error")
	}

	if len(m.validationErrors) > 0 {
		return nil, m.validationErrors, nil
	}

	category := &dto.ResponseDTO{
		Id:              m.nextId,
		Name:            createDto.Name,
		Description:     createDto.Description,
		Image:           createDto.Image,
		HeaderImage:     createDto.HeaderImage,
		Slug:            createDto.Slug,
		ParentID:        createDto.ParentID,
		SortIndex:       createDto.SortIndex,
		SeoTitle:        createDto.SeoTitle,
		SeoDescription:  createDto.SeoDescription,
		SeoKeys:         createDto.SeoKeys,
		IsActive:        createDto.IsActive,
		IsShade:         createDto.IsShade,
		IsVisibleInMenu: createDto.IsVisibleInMenu,
		IsVisibleOnMain: createDto.IsVisibleOnMain,
	}

	m.categories[m.nextId] = category
	m.nextId++
	return category, nil, nil
}

func (m *mockService) GetAll() ([]*dto.ResponseDTO, error) {
	if m.shouldReturnError {
		return nil, fmt.Errorf("service error")
	}

	var categories []*dto.ResponseDTO
	for _, cat := range m.categories {
		categories = append(categories, cat)
	}
	return categories, nil
}

func (m *mockService) GetById(id uint) (*dto.ResponseDTO, error) {
	if m.shouldReturnError {
		return nil, fmt.Errorf("service error")
	}

	if category, exists := m.categories[id]; exists {
		return category, nil
	}
	return nil, fmt.Errorf("category not found")
}

func (m *mockService) Update(id uint, updateDto dto.UpdateDTO) (*dto.ResponseDTO, []response.ErrorField, error) {
	if m.shouldReturnError {
		return nil, nil, fmt.Errorf("service error")
	}

	if len(m.validationErrors) > 0 {
		return nil, m.validationErrors, nil
	}

	category, exists := m.categories[id]
	if !exists {
		return nil, nil, fmt.Errorf("category not found")
	}

	if updateDto.Name != nil {
		category.Name = *updateDto.Name
	}
	if updateDto.Description != nil {
		category.Description = *updateDto.Description
	}
	if updateDto.Image != nil {
		category.Image = *updateDto.Image
	}
	if updateDto.HeaderImage != nil {
		category.HeaderImage = *updateDto.HeaderImage
	}
	if updateDto.Slug != nil {
		category.Slug = *updateDto.Slug
	}
	if updateDto.ParentID != nil {
		category.ParentID = updateDto.ParentID
	}
	if updateDto.SortIndex != nil {
		category.SortIndex = *updateDto.SortIndex
	}
	if updateDto.SeoTitle != nil {
		category.SeoTitle = *updateDto.SeoTitle
	}
	if updateDto.SeoDescription != nil {
		category.SeoDescription = *updateDto.SeoDescription
	}
	if updateDto.SeoKeys != nil {
		category.SeoKeys = *updateDto.SeoKeys
	}
	if updateDto.IsActive != nil {
		category.IsActive = *updateDto.IsActive
	}
	if updateDto.IsShade != nil {
		category.IsShade = *updateDto.IsShade
	}
	if updateDto.IsVisibleInMenu != nil {
		category.IsVisibleInMenu = *updateDto.IsVisibleInMenu
	}
	if updateDto.IsVisibleOnMain != nil {
		category.IsVisibleOnMain = *updateDto.IsVisibleOnMain
	}

	return category, nil, nil
}

func (m *mockService) Delete(id uint) (*dto.ResponseDTO, int64, error) {
	if m.shouldReturnError {
		return nil, 0, fmt.Errorf("service error")
	}

	category, exists := m.categories[id]
	if !exists {
		return nil, 0, fmt.Errorf("category not found")
	}

	if m.deleteReturnsCategoryWithError {
		return category, 0, fmt.Errorf("service error")
	}

	delete(m.categories, id)
	return category, 0, nil
}

func setupTestHandler() (*Handler, *mockService) {
	mockSvc := newMockService()
	handler := NewHandler(mockSvc)
	return handler, mockSvc
}

func TestHandler_Create_Success(t *testing.T) {
	handler, _ := setupTestHandler()

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

	jsonData, err := json.Marshal(createDto)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/category/create", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Create(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rr.Code)
		t.Logf("Response body: %s", rr.Body.String())
		return
	}

	var res map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if res["isSuccess"] != true {
		t.Error("Expected isSuccess to be true")
		t.Logf("Response: %+v", res)
		return
	}

	data := res["data"].(map[string]interface{})
	if data["name"] != "Test Category" {
		t.Errorf("Expected name 'Test Category', got %s", data["name"])
	}
}

func TestHandler_Create_InvalidJSON(t *testing.T) {
	handler, _ := setupTestHandler()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/category/create", bytes.NewReader([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Create(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestHandler_Create_ValidationErrors(t *testing.T) {
	handler, mockSvc := setupTestHandler()

	mockSvc.validationErrors = []response.ErrorField{
		{Field: "name", ErrorCode: string(response.BadRequest)},
	}

	createDto := dto.CreateDTO{
		Name:            "Test Category",
		Description:     "Test Description",
		Image:           "test.jpg",
		HeaderImage:     "header.jpg",
		Slug:            "test-category",
		SortIndex:       100,
		IsActive:        true,
		IsShade:         false,
		IsVisibleInMenu: true,
		IsVisibleOnMain: false,
	}

	jsonData, err := json.Marshal(createDto)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/category/create", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Create(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestHandler_GetAll_Success(t *testing.T) {
	handler, mockSvc := setupTestHandler()

	categories := []dto.CreateDTO{
		{Name: "Category 1", Slug: "category-1", Image: "img1.jpg", HeaderImage: "header1.jpg", SortIndex: 100, IsActive: true},
		{Name: "Category 2", Slug: "category-2", Image: "img2.jpg", HeaderImage: "header2.jpg", SortIndex: 200, IsActive: true},
	}

	for _, cat := range categories {
		_, _, err := mockSvc.Create(cat)
		if err != nil {
			return
		}
	}

	rr := httptest.NewRecorder()
	handler.GetAll(rr)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var res map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if res["isSuccess"] != true {
		t.Error("Expected isSuccess to be true")
	}

	data := res["data"].([]interface{})
	if len(data) != 2 {
		t.Errorf("Expected 2 categories, got %d", len(data))
	}
}

func TestHandler_GetAll_ServiceError(t *testing.T) {
	handler, mockSvc := setupTestHandler()

	mockSvc.shouldReturnError = true

	rr := httptest.NewRecorder()
	handler.GetAll(rr)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rr.Code)
	}
}

func TestHandler_GetById_Success(t *testing.T) {
	handler, mockSvc := setupTestHandler()

	createDto := dto.CreateDTO{
		Name:        "Test Category",
		Slug:        "test-category",
		Image:       "test.jpg",
		HeaderImage: "header.jpg",
		SortIndex:   100,
		IsActive:    true,
	}

	category, _, _ := mockSvc.Create(createDto)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/category/%d", category.Id), nil)
	req.SetPathValue("id", fmt.Sprintf("%d", category.Id))

	rr := httptest.NewRecorder()
	handler.GetById(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var res map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if res["isSuccess"] != true {
		t.Error("Expected isSuccess to be true")
	}

	data := res["data"].(map[string]interface{})
	if data["name"] != "Test Category" {
		t.Errorf("Expected name 'Test Category', got %s", data["name"])
	}
}

func TestHandler_GetById_NotFound(t *testing.T) {
	handler, _ := setupTestHandler()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/category/999", nil)
	req.SetPathValue("id", "999")

	rr := httptest.NewRecorder()
	handler.GetById(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, rr.Code)
	}
}

func TestHandler_GetById_InvalidID(t *testing.T) {
	handler, _ := setupTestHandler()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/category/invalid", nil)
	req.SetPathValue("id", "invalid")

	rr := httptest.NewRecorder()
	handler.GetById(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestHandler_GetById_MissingID(t *testing.T) {
	handler, _ := setupTestHandler()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/category/", nil)

	rr := httptest.NewRecorder()
	handler.GetById(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestHandler_Update_Success(t *testing.T) {
	handler, mockSvc := setupTestHandler()

	createDto := dto.CreateDTO{
		Name:        "Original Category",
		Slug:        "original-category",
		Image:       "original.jpg",
		HeaderImage: "original-header.jpg",
		SortIndex:   100,
		IsActive:    true,
	}

	category, _, _ := mockSvc.Create(createDto)

	updatedName := "Updated Category"
	updatedDescription := "Updated Description"
	updateDto := dto.UpdateDTO{
		Name:        &updatedName,
		Description: &updatedDescription,
	}

	jsonData, err := json.Marshal(updateDto)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/category/%d", category.Id), bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("id", fmt.Sprintf("%d", category.Id))

	rr := httptest.NewRecorder()
	handler.Update(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
		t.Logf("Response body: %s", rr.Body.String())
		return
	}

	var res map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if res["isSuccess"] != true {
		t.Error("Expected isSuccess to be true")
		t.Logf("Response: %+v", res)
		return
	}

	data := res["data"].(map[string]interface{})
	if data["name"] != "Updated Category" {
		t.Errorf("Expected name 'Updated Category', got %s", data["name"])
	}
	if data["description"] != "Updated Description" {
		t.Errorf("Expected description 'Updated Description', got %s", data["description"])
	}
}

func TestHandler_Update_NotFound(t *testing.T) {
	handler, _ := setupTestHandler()

	updatedName := "Updated Category"
	updateDto := dto.UpdateDTO{
		Name: &updatedName,
	}

	jsonData, err := json.Marshal(updateDto)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	req := httptest.NewRequest(http.MethodPut, "/api/v1/category/999", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("id", "999")

	rr := httptest.NewRecorder()
	handler.Update(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rr.Code)
	}
}

func TestHandler_Update_InvalidID(t *testing.T) {
	handler, _ := setupTestHandler()

	updatedName := "Updated Category"
	updateDto := dto.UpdateDTO{
		Name: &updatedName,
	}

	jsonData, err := json.Marshal(updateDto)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	req := httptest.NewRequest(http.MethodPut, "/api/v1/category/invalid", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("id", "invalid")

	rr := httptest.NewRecorder()
	handler.Update(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestHandler_Update_InvalidJSON(t *testing.T) {
	handler, mockSvc := setupTestHandler()

	createDto := dto.CreateDTO{
		Name:        "Original Category",
		Slug:        "original-category",
		Image:       "original.jpg",
		HeaderImage: "original-header.jpg",
		SortIndex:   100,
		IsActive:    true,
	}

	category, _, _ := mockSvc.Create(createDto)

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/category/%d", category.Id), bytes.NewReader([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("id", fmt.Sprintf("%d", category.Id))

	rr := httptest.NewRecorder()
	handler.Update(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestHandler_Update_ValidationErrors(t *testing.T) {
	handler, mockSvc := setupTestHandler()

	createDto := dto.CreateDTO{
		Name:        "Original Category",
		Slug:        "original-category",
		Image:       "original.jpg",
		HeaderImage: "original-header.jpg",
		SortIndex:   100,
		IsActive:    true,
	}

	category, _, _ := mockSvc.Create(createDto)

	mockSvc.validationErrors = []response.ErrorField{
		{Field: "name", ErrorCode: string(response.BadRequest)},
	}

	updatedName := "Updated Category"
	updateDto := dto.UpdateDTO{
		Name: &updatedName,
	}

	jsonData, err := json.Marshal(updateDto)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/category/%d", category.Id), bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("id", fmt.Sprintf("%d", category.Id))

	rr := httptest.NewRecorder()
	handler.Update(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestHandler_Delete_Success(t *testing.T) {
	handler, mockSvc := setupTestHandler()

	createDto := dto.CreateDTO{
		Name:        "Category to Delete",
		Slug:        "category-to-delete",
		Image:       "delete.jpg",
		HeaderImage: "delete-header.jpg",
		SortIndex:   100,
		IsActive:    true,
	}

	category, _, _ := mockSvc.Create(createDto)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/category/%d", category.Id), nil)
	req.SetPathValue("id", fmt.Sprintf("%d", category.Id))

	rr := httptest.NewRecorder()
	handler.Delete(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
		t.Logf("Response body: %s", rr.Body.String())
		return
	}

	var res map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if res["isSuccess"] != true {
		t.Error("Expected isSuccess to be true")
		t.Logf("Response: %+v", res)
		return
	}

	_, exists := mockSvc.categories[category.Id]
	if exists {
		t.Error("Category should have been deleted")
	}
}

func TestHandler_Delete_NotFound(t *testing.T) {
	handler, _ := setupTestHandler()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/category/999", nil)
	req.SetPathValue("id", "999")

	rr := httptest.NewRecorder()
	handler.Delete(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, rr.Code)
	}
}

func TestHandler_Delete_InvalidID(t *testing.T) {
	handler, _ := setupTestHandler()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/category/invalid", nil)
	req.SetPathValue("id", "invalid")

	rr := httptest.NewRecorder()
	handler.Delete(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestHandler_Delete_MissingID(t *testing.T) {
	handler, _ := setupTestHandler()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/category/", nil)

	rr := httptest.NewRecorder()
	handler.Delete(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestHandler_Delete_ServiceError(t *testing.T) {
	handler, mockSvc := setupTestHandler()

	createDto := dto.CreateDTO{
		Name:        "Category to Delete",
		Slug:        "category-to-delete",
		Image:       "delete.jpg",
		HeaderImage: "delete-header.jpg",
		SortIndex:   100,
		IsActive:    true,
	}

	category, _, _ := mockSvc.Create(createDto)

	mockSvc.deleteReturnsCategoryWithError = true

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/category/%d", category.Id), nil)
	req.SetPathValue("id", fmt.Sprintf("%d", category.Id))

	rr := httptest.NewRecorder()
	handler.Delete(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rr.Code)
	}
}
