package services

import (
	"BookShopGoProject/internal/models"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestBookService_GetBookById(t *testing.T) {
	updateTime := time.Now()
	testcases := []struct {
		name           string
		input          uint
		expected       models.Book
		expectedErr    error
		mockRepoInput  uint
		mockRepoOutput models.Book
		mockRepoError  error
	}{
		{
			name:  "Get book successfully",
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
			expectedErr:   nil,
			mockRepoInput: 1,
			mockRepoOutput: models.Book{
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
			mockRepoError: nil,
		},
		{
			name:           "Get book fail",
			input:          0,
			expected:       models.Book{},
			expectedErr:    nil,
			mockRepoInput:  0,
			mockRepoOutput: models.Book{},
			mockRepoError:  nil,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			bookRepoMock := new(BookRepoMock)

			bookRepoMock.On("GetBookById", testcase.input).Return(
				testcase.mockRepoOutput, testcase.mockRepoError)

			bookService := BookService{
				IBookRepo: bookRepoMock,
			}

			result, err := bookService.GetBookById(testcase.input)

			require.Equal(t, result, testcase.expected)
			if err != nil {
				require.Equal(t, testcase.expectedErr, err)
			}
		})
	}
}

func TestBookService_GetBookByName(t *testing.T) {
	updateTime := time.Now()
	testcases := []struct {
		name           string
		input          string
		expected       *models.Book
		expectedErr    error
		mockRepoInput  string
		mockRepoOutput *models.Book
		mockRepoError  error
	}{
		{
			name:  "Get book successfully",
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
			expectedErr:   nil,
			mockRepoInput: "Harry Potter and the Sorcerer s Stone",
			mockRepoOutput: &models.Book{
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
		},
		{
			name:           "Get book fail",
			input:          "",
			expected:       &models.Book{},
			expectedErr:    nil,
			mockRepoInput:  "",
			mockRepoOutput: &models.Book{},
			mockRepoError:  nil,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			bookRepoMock := new(BookRepoMock)
			bookService := BookService{
				IBookRepo: bookRepoMock,
			}
			bookRepoMock.On("GetBookByName", testcase.mockRepoInput).Return(testcase.mockRepoOutput, testcase.mockRepoError)

			result, err := bookService.GetBookByName(testcase.input)
			require.Equal(t, result, testcase.expected)
			if err != nil {
				require.EqualError(t, err, testcase.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}

		})
	}
}

func TestBookService_GetAllBooks(t *testing.T) {
	updateTime := time.Now()
	testcases := []struct {
		name           string
		expected       []models.Book
		expectedErr    error
		mockRepoOutput []models.Book
		mockRepoError  error
	}{
		{
			name: "Get all books successfully",
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
			mockRepoOutput: []models.Book{
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
			mockRepoError: nil,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			bookRepoMock := new(BookRepoMock)

			bookRepoMock.On("GetBooks").Return(testcase.mockRepoOutput, testcase.mockRepoError)

			bookService := BookService{
				IBookRepo: bookRepoMock,
			}
			result, err := bookService.GetAllBooks()
			require.Equal(t, result, testcase.expected)
			if err != nil {
				require.EqualError(t, err, testcase.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBookService_CreateBook(t *testing.T) {
	testcases := []struct {
		name           string
		input          models.BookRepoInput
		expectedErr    error
		mockRepoInput  models.BookRepoInput
		mockRepoOutput *models.Book
		mockRepoError  error
	}{
		{
			name: "success",
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
			mockRepoInput: models.BookRepoInput{
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
			mockRepoOutput: &models.Book{
				BookId: 0,
			},
			mockRepoError: nil,
		},
		{
			name: "book already exists",
			input: models.BookRepoInput{
				Book: &models.Book{
					BookName:    "Harry Potter and the Sorcerer s Stone",
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
			expectedErr: errors.New("book already exists"),
			mockRepoInput: models.BookRepoInput{
				Book: &models.Book{
					BookName:    "Harry Potter and the Sorcerer s Stone",
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
			mockRepoOutput: &models.Book{
				BookId: 1,
			},
			mockRepoError: nil,
		},
		{
			name: "failed to create book",
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
			expectedErr: errors.New("failed to create book"),
			mockRepoInput: models.BookRepoInput{
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
			mockRepoOutput: &models.Book{
				BookId: 0,
			},
			mockRepoError: errors.New("failed to create book"),
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			bookRepoMock := new(BookRepoMock)

			bookRepoMock.On("CreateBook", testcase.mockRepoInput.Book, testcase.mockRepoInput.Quantity).Return(testcase.mockRepoError)
			bookRepoMock.On("GetBookByName", testcase.mockRepoInput.Book.BookName).Return(testcase.mockRepoOutput, testcase.mockRepoError)
			bookService := BookService{
				IBookRepo: bookRepoMock,
			}

			err := bookService.CreateBook(testcase.input.Book, testcase.input.Quantity)
			if err != nil {
				require.EqualError(t, err, testcase.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}

}
