# Portmonetka Authorization & User Service

Portmonetka Authorization & User Service is a backend service that handles user authentication, authorization, and basic user management. The service provides a set of RESTful API endpoints to facilitate user login, creation, deletion, and profile updates. It is designed for integration with other applications requiring secure user authentication and management.

## API Endpoints Overview

The following is a brief overview of the key API endpoints offered by the service:

- **Login Endpoint** (`/login`): Allows users to authenticate using their username and password. Successful login returns a token for subsequent requests.
- **User Creation Endpoint** (`/users`): Enables the creation of new users with the required details like username and password.
- **User Deletion Endpoint** (`/users/{userId}`): Deletes a user by their unique ID.
- **Update User Password Endpoint** (`/users/{userId}/password`): Updates the password of a user by providing the user ID and new password.
- **Update Username Endpoint** (`/users/{userId}/username`): Updates the username for a given user by their user ID.

## License

This project is licensed under the Apache 2.0 License. For more details, see the [Apache 2.0 License](http://www.apache.org/licenses/LICENSE-2.0.html).

## Contact

For questions, comments, or contributions, please contact the project maintainers or open an issue on GitHub.
