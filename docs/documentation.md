
## Components

### main.go

The entry point of the application, which runs the console interface.

### controllers/library_controller.go

Handles console input and invokes the appropriate service methods to interact with the library.

### models/book.go

Defines the `Book` struct with fields for ID, title, author, and status.

### models/member.go

Defines the `Member` struct with fields for ID, name, and a slice of borrowed books.

### services/library_service.go

Contains the implementation of the `LibraryManager` interface, including methods to add, remove, borrow, return, and list books.

## Functionality

- **Add a new book**: Adds a book to the library.
- **Remove an existing book**: Removes a book from the library by its ID.
- **Borrow a book**: Allows a member to borrow a book if it is available.
- **Return a book**: Allows a member to return a borrowed book.
- **List all available books**: Lists all books in the library that are currently available.
- **List all borrowed books by a member**: Lists all books borrowed by a specific member.

## Error Handling

The system includes error handling for scenarios such as:

- Book or member not found.
- Book is already borrowed or available.
- Member not found when returning a book.

## Running the Application

To run the application, execute the following command in the root directory:

```sh
go run main.go
