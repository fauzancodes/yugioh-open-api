# Yu-Gi-Oh! Open API Documentation

Welcome to the **Yu-Gi-Oh! Open API** documentation. This API is built using the [Go](https://go.dev/) programming language with the [Echo](https://echo.labstack.com/) framework, leveraging [GORM](https://gorm.io/) as the ORM for database connections. [PostgreSQL](https://www.postgresql.org/) is used as the main database, ensuring high performance and scalability. The API is designed to efficiently and securely support various data management features, including authentication, card database, deck management, and other functionalities to meet player needs.

## Features

- **Secure Authentication**:  
  Effortless user access with login and register endpoints, powered by [JWT](https://jwt.io/) authentication and industry-standard [bcrypt](https://en.wikipedia.org/wiki/Bcrypt) hashing for a secure user experience.

- **Dynamic Card Database**:  
  Access the comprehensive Yu-Gi-Oh! card database through public endpoints featuring advanced filtering, searching, and sorting options to explore cards effortlessly.

- **Personalized Deck Builder**:  
  Build your dream decks with endpoints that support creating, updating, and deleting custom decks, empowering users to manage their strategies effectively.

- **Public Deck Sharing**:  
  Discover community-created decks through public deck endpoints, featuring filtering by user preferences.

- **Dynamic Data Retrieval**:  
  All GET endpoints feature dynamic pagination, filtering, searching, and sorting options, making data retrieval intuitive and efficient.

- **Robust Input Validation**:  
  Ensure data integrity with comprehensive validation on all POST endpoints, preventing errors and enhancing user experience.

- **SQL Injection Prevention**:  
  All endpoints are protected with query parameter sanitization, guarding against SQL injection attacks for enhanced security.

- **Secure API Key**:  
  Each request to the endpoints is protected with an API key generated using [HMAC](https://en.wikipedia.org/wiki/HMAC) and [SHA-256](https://en.wikipedia.org/wiki/SHA-2) methods, ensuring robust security and integrity for every transaction while safeguarding your data from unauthorized access.

- **Comprehensive Documentation**:  
  Navigate through the API with ease using clear, organized [Postman](https://www.postman.com/) documentation designed to ensure a smooth development experience for all developers.

---

© 2024 Yu-Gi-Oh! Open API Project. All rights reserved. By [fauzancodes](https://fauzancodes.id/).
