package repositories

import (
	"BookShopGoProject/internal/models"
	"BookShopGoProject/internal/testhelpers"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestBookRepo_GetBooks(t *testing.T) {
	updateTime := time.Now()
	testcases := []struct {
		name        string
		expected    []models.Book
		expectedErr error
		mockDB      *gorm.DB
		preparePath string
	}{
		//{
		//	name:        "Get Books failed with Erros",
		//	expected:    nil,
		//	expectedErr: errors.New("pq: password authentication failed for user \"postgrespassword=1235\""),
		//	mockDB:      testhelpers.ConnectDBFailed(),
		//	preparePath: "../testhelpers/preparedata/data",
		//},
		{
			name: "Get Books succeeds",
			expected: []models.Book{
				{
					BookId:      1,
					BookName:    "Harry Potter and the Sorcerer s Stone",
					Description: "",
					Price:       19.99,
					Genre:       "Fantasy",
					Status:      false,
					Authors: []models.Author{
						{
							AuthorId:   1,
							AuthorName: "J.K. Rowling",
							Status:     false,
							Books:      nil,
						},
					},
					Orders: nil,
					Stocks: []models.Stock{
						{
							StockId:  1,
							Quantity: 50,
							UpdateAt: updateTime,
							BookId:   1, // Set to match the top-level BookId
						},
					},
				},
				{
					BookId:      2,
					BookName:    "1984",
					Description: "",
					Price:       14.99,
					Genre:       "Dystopian",
					Status:      false,
					Authors: []models.Author{
						{
							AuthorId:   2,
							AuthorName: "George Orwell",
							Status:     false,
							Books:      nil,
						},
					},
					Orders: nil,
					Stocks: []models.Stock{
						{
							StockId:  2,
							Quantity: 30,
							UpdateAt: updateTime,
							BookId:   2, // Set to match the top-level BookId
						},
					},
				},
				{
					BookId:      3,
					BookName:    "The Hobbit",
					Description: "",
					Price:       15.99,
					Genre:       "Fantasy",
					Status:      false,
					Authors: []models.Author{
						{
							AuthorId:   3,
							AuthorName: "J.R.R. Tolkien",
							Status:     false,
							Books:      nil,
						},
					},
					Orders: nil,
					Stocks: []models.Stock{
						{
							StockId:  3,
							Quantity: 40,
							UpdateAt: updateTime,
							BookId:   3, // Set to match the top-level BookId
						},
					},
				},
				{
					BookId:      4,
					BookName:    "The Da Vinci Code",
					Description: "",
					Price:       18.99,
					Genre:       "Thriller",
					Status:      false,
					Authors: []models.Author{
						{
							AuthorId:   5,
							AuthorName: "Dan Brown",
							Status:     false,
							Books:      nil,
						},
					},
					Orders: nil,
					Stocks: []models.Stock{
						{
							StockId:  4,
							Quantity: 60,
							UpdateAt: updateTime,
							BookId:   4, // Set to match the top-level BookId
						},
					},
				},
			},
			expectedErr: nil,
			mockDB:      testhelpers.ConnectDB(),
			preparePath: "../testhelpers/preparedata/data",
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			testhelpers.PrepareDataForDb(testcase.mockDB, testcase.preparePath)

			asserts := assert.New(t)

			bookRepo := BookRepo{
				DB: testcase.mockDB,
			}

			result, gotErr := bookRepo.GetBooks()

			assertBooksEqualIgnoringTime(t, testcase.expected, result)

			if gotErr != nil {
				asserts.EqualError(gotErr, testcase.expectedErr.Error())
			} else {
				asserts.NoError(gotErr)
			}
		})
	}

}

func TestBookRepo_GetBookById(t *testing.T) {
	updateTime := time.Now()
	testcases := []struct {
		name        string
		input       uint
		expected    models.Book
		expectedErr error
		mockDB      *gorm.DB
		preparePath string
	}{
		{
			name:  "Get Books succeeds",
			input: 1,
			expected: models.Book{

				BookId:      1,
				BookName:    "Harry Potter and the Sorcerer s Stone",
				Description: "",
				Price:       19.99,
				Genre:       "Fantasy",
				Status:      false,
				Authors: []models.Author{
					{
						AuthorId:   1,
						AuthorName: "J.K. Rowling",
						Status:     false,
						Books:      nil,
					},
				},
				Orders: nil,
				Stocks: []models.Stock{
					{
						StockId:  1,
						Quantity: 50,
						UpdateAt: updateTime,
						BookId:   1, // Set to match the top-level BookId
					},
				},
			},
			expectedErr: nil,
			mockDB:      testhelpers.ConnectDB(),
			preparePath: "../testhelpers/preparedata/data",
		},
		{
			name:        "Get Books not found",
			input:       0,
			expected:    models.Book{},
			expectedErr: gorm.ErrRecordNotFound,
			mockDB:      testhelpers.ConnectDB(),
			preparePath: "../testhelpers/preparedata/data",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {

			asserts := assert.New(t)

			bookRepo := BookRepo{
				DB: testcase.mockDB,
			}

			result, gotErr := bookRepo.GetBookById(testcase.input)

			assertBooksEqualIgnoringTime(t, []models.Book{testcase.expected}, []models.Book{result})

			if gotErr != nil {
				asserts.EqualError(gotErr, testcase.expectedErr.Error())
			} else {
				asserts.NoError(gotErr)
			}
		})
	}
}
func TestBookRepo_CreateBook(t *testing.T) {
	testcases := []struct {
		name        string
		input       models.BookRepoInput
		expectedErr error
		mockDB      *gorm.DB
		preparePath string
	}{
		{
			name: "Create Book succeeds",
			input: models.BookRepoInput{
				Book: &models.Book{
					BookName:    "Test Book",
					Description: "Test Description",
					Price:       15.99,
					Genre:       "Fantasy",
					Status:      true,
					Authors: []models.Author{
						{
							AuthorId: 1,
						},
						{
							AuthorId: 2,
						},
					},
				},
				Quantity: 100,
			},
			expectedErr: nil,
			mockDB:      testhelpers.ConnectDB(),
			preparePath: "../testhelpers/preparedata/data",
		},
	}
	fmt.Println(testcases)
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			asserts := assert.New(t)

			bookRepo := BookRepo{
				DB: testcase.mockDB,
			}

			result := bookRepo.CreateBook(testcase.input.Book, testcase.input.Quantity)

			if result != nil {
				asserts.EqualError(result, testcase.expectedErr.Error())

			} else {
				asserts.NoError(result)
			}

		})
	}
}

func TestBookRepo_GetBookByName(t *testing.T) {
	updateTime := time.Now()
	testcases := []struct {
		name        string
		input       string
		expected    *models.Book
		expectedErr error
		mockDB      *gorm.DB
		preparePath string
	}{
		{
			name:  "Get Books succeeds",
			input: "Harry Potter and the Sorcerer s Stone",
			expected: &models.Book{
				BookId:      1,
				BookName:    "Harry Potter and the Sorcerer s Stone",
				Description: "",
				Price:       19.99,
				Genre:       "Fantasy",
				Status:      false,
				Authors: []models.Author{
					{
						AuthorId:   1,
						AuthorName: "J.K. Rowling",
						Status:     false,
						Books:      nil,
					},
				},
				Orders: nil,
				Stocks: []models.Stock{
					{
						StockId:  1,
						Quantity: 50,
						UpdateAt: updateTime,
						BookId:   1,
					},
				},
			},
			expectedErr: nil,
			mockDB:      testhelpers.ConnectDB(),
			preparePath: "../testhelpers/preparedata/data",
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			requires := assert.New(t)

			bookRepo := BookRepo{
				DB: testcase.mockDB,
			}

			result, err := bookRepo.GetBookByName(testcase.input)
			assertBooksEqualIgnoringTime(t, []models.Book{*testcase.expected}, []models.Book{*result})
			if err != nil {
				requires.EqualError(err, testcase.expectedErr.Error())
			} else {
				requires.NoError(err)
			}

		})
	}
}

func assertBooksEqualIgnoringTime(t *testing.T, expected, actual []models.Book) {
	for i := range expected {
		for j := range expected[i].Stocks {
			expected[i].Stocks[j].UpdateAt = time.Time{} // Set to zero value for comparison
			actual[i].Stocks[j].UpdateAt = time.Time{}   // Set to zero value for comparison
		}
	}
	assert.Equal(t, expected, actual)
}
