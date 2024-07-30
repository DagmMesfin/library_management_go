# Library Management System

## Overview
The Library Management System is a console application developed in Go. It showcases the use of structs, interfaces, and various Go features such as methods, slices, and maps. The application allows for managing books and members in a library, including adding, removing, borrowing, and returning books.

## Features
- **Book Management**: Add, remove, and list books in the library.
- **Member Management**: Manage library members and track the books they borrow.
- **Borrowing and Returning**: Handle the borrowing and returning of books, ensuring accurate tracking of book statuses.

## Installation

### Prerequisites
- [Go](https://golang.org/dl/) 1.16 or later

### Steps
1. Clone the repository:
   \`\`\`bash
   git clone https://github.com/yourusername/library_management.git
   \`\`\`
2. Navigate to the project directory:
   \`\`\`bash
   cd library_management
   \`\`\`
3. Install the dependencies:
   \`\`\`bash
   go mod tidy
   \`\`\`

## Usage

### Running the Application
To run the application, execute the following command:
\`\`\`bash
go run main.go
\`\`\`

### Application Operations
Once the application is running, you can perform various operations from the console:

- **Add a Book**: Add a new book to the library by entering the book details.
- **Remove a Book**: Remove an existing book from the library by its ID.
- **Borrow a Book**: Borrow a book by providing the book ID and the member ID.
- **Return a Book**: Return a borrowed book by providing the book ID and the member ID.
- **List Available Books**: Display all available books in the library.
- **List Borrowed Books**: Display all books borrowed by a specific member.

## Project Structure
\`\`\`
library_management/
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   └── book.go
│   └── member.go
├── services/
│   └── library_service.go
├── docs/
│   └── api_documentation.md
└── go.mod
\`\`\`
- **main.go**: Entry point of the application.
- **controllers/library_controller.go**: Handles the main logic and user interactions in the console application.
- **models/book.go**: Defines the `Book` struct.
- **models/member.go**: Defines the `Member` struct.
- **services/library_service.go**: Contains business logic and data manipulation functions.
- **docs/api_documentation.md**: Contains additional documentation.

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments
Special thanks to the Go community for their excellent documentation and resources.
