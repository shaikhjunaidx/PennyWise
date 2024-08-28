
# PennyWise

PennyWise is a personal finance management application designed to help users track income, expenses, and savings goals. The application provides insights into spending patterns, generates financial reports, and offers budget management tools to optimize spending.

## **Backend Overview**

The backend of PennyWise is built using Go (Golang) with MySQL as the relational database. The backend architecture follows a clean and modular design, ensuring scalability, maintainability, and ease of testing.

### **Key Components:**

1. **Project Structure:**
   - The backend is organized into different packages (`internal`, `pkg`, `db`, `models`, etc.) to separate concerns and keep the codebase modular.
   - `internal/user`: Handles user management, including authentication, profile updates, and password resets.
   - `db`: Manages database connections and migrations using Gorm.

2. **Database Integration:**
   - MySQL is used as the relational database, with Gorm as the ORM for database interactions.
   - Automatic database creation and schema migrations are set up to ensure the database is always in sync with the data models.

3. **Authentication:**
   - JWT (JSON Web Token) is planned to manage user sessions securely, allowing for stateless authentication.
   - Middleware will be implemented to protect routes, ensuring that only authenticated users can access certain endpoints.

4. **Service Layer:**
   - The service layer encapsulates the business logic for user management, including signing up, logging in, and updating user profiles.
   - Utility functions are provided for common tasks like password hashing and token generation.

5. **Testing:**
   - A comprehensive test suite is in place, covering unit and integration tests for the service and repository layers.
   - A separate test database is used to ensure that tests do not interfere with the development environment.

### **Setting Up the Environment**

To run the PennyWise backend locally, you need to set up environment variables that configure database connections and other settings. Here's how to set up the required environment files:

1. **Create `.env.dev` and `.env.test` Files:**

   - **`.env.dev`:** Used for local development.
   - **`.env.test`:** Used for running tests.

   **Example `.env.dev` file:**
   ```plaintext
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=pennywise_dev
   DB_HOST=localhost
   DB_PORT=3306
   ```

   **Example `.env.test` file:**
   ```plaintext
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=pennywise_test
   DB_HOST=localhost
   DB_PORT=3306
   ```

2. **Load Environment Variables:**

   The application automatically loads environment variables from the appropriate `.env` file based on the environment. Ensure that you have `APP_ENV` set to either `dev` or `test`:

   **Setting Environment for Development:**
   ```bash
   export APP_ENV=dev
   ```

   **Setting Environment for Testing:**
   ```bash
   export APP_ENV=test
   ```

3. **Running the Backend:**

   Once your environment is set up, you can start the backend using:

    ```bash
    go run cmd/main.go
    ```

4. **Running Tests:**

   To run the test suite, make sure you're using the test environment and run:

   ```bash
   go test ./...
   ```

## **Future Work**

- Implement JWT-based authentication and route protection.
- Add user profile management features for updating user information.
- Enhance error handling and validation mechanisms.
