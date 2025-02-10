Below is an example of a complete README file for the Book Shop Go Project. You can adjust the endpoints and payloads as needed to reflect your actual implementation.

---

# Book Shop Go Project

:robot: A helpful project to manage and organize your bookstore smoothly.

## Building the App

:seedling: To run this project on your local machine, use the following Docker commands:

```docker
docker-compose build
docker-compose up
```

## APIs

:hammer_and_wrench: Below is a list of the APIs available in this project.

### Create Book

**Endpoint:** `POST /books`

**Description:**  
Creates a new book along with its stock quantity.  
The request checks if the book already exists and, if not, creates a new entry.

**Request Body:**

```json
{
  "book": {
    "BookName": "Test Book",
    "Description": "Test Description",
    "Price": 15.99,
    "Genre": "Fantasy",
    "Status": true,
    "Authors": [
      { "AuthorId": 1 },
      { "AuthorId": 2 }
    ]
  },
  "quantity": 100
}
```

**Response Body (Success):**

```json
{
  "success": true,
  "book": {
    "BookId": 1,
    "BookName": "Test Book",
    "Description": "Test Description",
    "Price": 15.99,
    "Genre": "Fantasy",
    "Status": true,
    "Authors": [
      { "AuthorId": 1, "AuthorName": "Author Name 1", "Status": false },
      { "AuthorId": 2, "AuthorName": "Author Name 2", "Status": false }
    ],
    "Stocks": [
      {
        "StockId": 1,
        "Quantity": 100,
        "UpdateAt": "2025-02-09T12:00:00Z",
        "BookId": 1
      }
    ]
  }
}
```

**Response Body (Failure):**

```json
{
  "success": false,
  "error": "failed to create book"
}
```

---

### Get Book by ID

**Endpoint:** `GET /books/{id}`

**Description:**  
Retrieves the details of a book given its unique identifier.

**Response Body (Success):**

```json
{
  "success": true,
  "book": {
    "BookId": 1,
    "BookName": "Harry Potter and the Sorcerer s Stone",
    "Description": "",
    "Price": 19.99,
    "Genre": "Fantasy",
    "Status": false,
    "Authors": [
      {
        "AuthorId": 1,
        "AuthorName": "J.K. Rowling",
        "Status": false
      }
    ],
    "Stocks": [
      {
        "StockId": 1,
        "Quantity": 50,
        "UpdateAt": "2025-02-09T12:00:00Z",
        "BookId": 1
      }
    ]
  }
}
```

If the book is not found, the API will return an appropriate error message.

---

### Get Book by Name

**Endpoint:** `GET /books`

**Query Parameter:**  
`name` (e.g., `/books?name=Harry Potter and the Sorcerer s Stone`)

**Description:**  
Retrieves the details of a book by its name.

**Response Body (Success):**

```json
{
  "success": true,
  "book": {
    "BookId": 1,
    "BookName": "Harry Potter and the Sorcerer s Stone",
    "Description": "",
    "Price": 19.99,
    "Genre": "Fantasy",
    "Status": false,
    "Authors": [
      {
        "AuthorId": 1,
        "AuthorName": "J.K. Rowling",
        "Status": false
      }
    ],
    "Stocks": [
      {
        "StockId": 1,
        "Quantity": 50,
        "UpdateAt": "2025-02-09T12:00:00Z",
        "BookId": 1
      }
    ]
  }
}
```

On failure, the API will respond with an error message indicating the issue.

---

### Get All Books

**Endpoint:** `GET /books`

**Description:**  
Retrieves a list of all books available in the bookstore along with a count.

**Response Body (Success):**

```json
{
  "success": true,
  "books": [
    {
      "BookId": 1,
      "BookName": "Harry Potter and the Sorcerer s Stone",
      "Description": "",
      "Price": 19.99,
      "Genre": "Fantasy",
      "Status": false,
      "Authors": [
        {
          "AuthorId": 1,
          "AuthorName": "J.K. Rowling",
          "Status": false
        }
      ],
      "Stocks": [
        {
          "StockId": 1,
          "Quantity": 50,
          "UpdateAt": "2025-02-09T12:00:00Z",
          "BookId": 1
        }
      ]
    },
    {
      "BookId": 2,
      "BookName": "1984",
      "Description": "",
      "Price": 14.99,
      "Genre": "Dystopian",
      "Status": false,
      "Authors": [
        {
          "AuthorId": 2,
          "AuthorName": "George Orwell",
          "Status": false
        }
      ],
      "Stocks": [
        {
          "StockId": 2,
          "Quantity": 30,
          "UpdateAt": "2025-02-09T12:00:00Z",
          "BookId": 2
        }
      ]
    },
    {
      "BookId": 3,
      "BookName": "The Hobbit",
      "Description": "",
      "Price": 15.99,
      "Genre": "Fantasy",
      "Status": false,
      "Authors": [
        {
          "AuthorId": 3,
          "AuthorName": "J.R.R. Tolkien",
          "Status": false
        }
      ],
      "Stocks": [
        {
          "StockId": 3,
          "Quantity": 40,
          "UpdateAt": "2025-02-09T12:00:00Z",
          "BookId": 3
        }
      ]
    },
    {
      "BookId": 4,
      "BookName": "The Da Vinci Code",
      "Description": "",
      "Price": 18.99,
      "Genre": "Thriller",
      "Status": false,
      "Authors": [
        {
          "AuthorId": 5,
          "AuthorName": "Dan Brown",
          "Status": false
        }
      ],
      "Stocks": [
        {
          "StockId": 4,
          "Quantity": 60,
          "UpdateAt": "2025-02-09T12:00:00Z",
          "BookId": 4
        }
      ]
    }
  ],
  "count": 4
}
```

---

## Project Architecture

This project is built using a layered architecture to separate concerns:

- **Handlers:**  
  Receive HTTP requests, perform decoding and validation, and then invoke the appropriate service functions.

- **Services:**  
  Contain the business logic of the application and interact with the repository layer.

- **Repositories:**  
  Handle all data access operations, interacting directly with the database.

---

## Running Tests

:microscope: To run the unit tests for the project, execute the following command in the project root:

```bash
go test ./...
```

The tests cover key Book Service operations such as retrieving books by ID or name, listing all books, and creating new books. The tests are implemented using the [Testify](https://github.com/stretchr/testify) package.

---

## Contributing

Contributions are welcome! If you have ideas or improvements, please open an issue or submit a pull request.

---

## License

This project is licensed under the MIT License.

---

Feel free to update the endpoints, request/response examples, and other sections to accurately reflect your implementation details and project structure.
