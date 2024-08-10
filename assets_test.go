package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

// MockFetchJSON is a mock of the FetchJSON function
type MockFetchJSON struct {
	mock.Mock
}

func (m *MockFetchJSON) FetchJSON(url string, headers map[string]string) ([]byte, error) {
	args := m.Called(url, headers)
	return args.Get(0).([]byte), args.Error(1)
}

func TestFetchChampionAssetURL(t *testing.T) {
	// mockFetchJSON := new(MockFetchJSON)

	// // Mock environment variable
	// os.Setenv("LOLVERSION", "1.0.0")
	// defer os.Unsetenv("LOLVERSION")

	// // Sample JSON response
	// championData := ChampionData{
	// 	Data: map[string]Champion{
	// 		"1": {Key: "1", Image: Image{ImageUrl: "url1"}},
	// 		"2": {Key: "2", Image: Image{ImageUrl: "url2"}},
	// 	},
	// }
	// jsonData, _ := json.Marshal(championData)

	// // Set up expectations
	// mockFetchJSON.On("FetchJSON", "https://example.com/1.0.0/data", mock.Anything).Return(jsonData, nil)

	// // Call the function
	// result, err := fetchChampionAssetURL()

	// // Assert expectations
	// assert.NoError(t, err)
	// expected := map[int]string{
	// 	1: "url1",
	// 	2: "url2",
	// }
	// assert.Equal(t, expected, result)

	// // Ensure all expectations are met
	// mockFetchJSON.AssertExpectations(t)
}
