## 📄 URL Shortener - Design Document

### 📌 Overview

This document outlines the architecture, data modeling, and design decisions made in building a basic, scalable, and maintainable URL shortener system using Golang, Gin, and GORM with SQLite for persistent storage.

---

### 🧱 Architecture

#### Components

* **HTTP Server**: Built using the Gin web framework for routing and middleware.
* **Controller Layer**: Handles request validation, response formatting, and control flow.
* **Service Layer**: Encapsulates business logic such as short code generation and expiry logic.
* **Repository Layer**: Responsible for database interactions using GORM.
* **Database**: SQLite for initial development and prototyping.

#### Flow Diagram (Simplified)

```
Client -> Router -> Controller -> Service -> Repository -> SQLite DB
                             ↑                ↓
                      Redirect / JSON     Data Fetch / Store
```

---

### 📦 Technology Stack

| Component          | Technology               | Justification                                              |
| ------------------ | ------------------------ | ---------------------------------------------------------- |
| Language           | Golang                   | Efficient concurrency, fast performance, strong typing     |
| Framework          | Gin                      | Lightweight, performant, and middleware-friendly           |
| ORM                | GORM                     | Simplifies DB operations, migrations, and model management |
| Database           | SQLite                   | Suitable for lightweight and local development             |
| Environment Config | godotenv                 | Manages configuration using `.env` files                   |
| UUID Generator     | `github.com/google/uuid` | Reliable generation of unique short codes                  |

---

### 🧩 Data Model

```go
type Link struct {
  gorm.Model
  Url        string
  Validity   int
  ShortCode  string
  ClickCount int
  ExpireAt   time.Time
  ShortLink  string
}
```

#### Field Explanations

* **Url**: The original long URL provided by the user.
* **Validity**: Duration (in minutes) for which the link remains active.
* **ShortCode**: Unique 8-character identifier used in the short link.
* **ClickCount**: Tracks number of redirection requests.
* **ExpireAt**: Timestamp indicating when the short link expires.
* **ShortLink**: Final shortened URL (e.g., `http://localhost:8000/abc123`).

---

### ⚙️ Key Design Decisions

#### 1. **Layered Architecture**

* Separation of concerns allows easier maintenance and testing.
* Controller handles HTTP; services manage logic; repositories handle persistence.

#### 2. **Short Code Generation**

* Utilizes `uuid.New().String()[0:8]` to generate a unique short code.
* Ensures uniqueness by checking the database before accepting a code.

#### 3. **Extensible Expiry Logic**

* URLs are associated with expiration timestamps, allowing easy cleanup and link lifecycle control.

#### 4. **Routing Design**

* RESTful structure:

  * `POST /shorturls` – Create a new short URL
  * `GET /:shortCode` – Redirect to original URL
  * `GET /shorturls/:shortCode` – Get statistics
  * `GET /test` – Health check

#### 5. **Database Migration with GORM**

* Automatically creates tables if they don’t exist, reducing setup effort.
* Easy switch to PostgreSQL or MySQL in the future.

---

### 🧠 Assumptions

* The application is initially used in a local or small-scale environment (SQLite suffices).
* Short codes are 8 characters, which balances readability and uniqueness.
* Authentication and rate limiting are not required in the MVP.
* Time-based expiration is sufficient for link validity control.
* No frontend/UI layer is considered at this stage.

---

### 🚀 Scalability & Maintainability

* **Scalable**:

  * Easily switch to more robust databases like PostgreSQL.
  * Goroutines can be used for background cleanup or logging.
  * RESTful API supports integration with frontend or third-party clients.

* **Maintainable**:

  * Modular file structure encourages separation of logic.
  * Config-driven design with `.env` makes it easy to manage environments.
  * Logs errors for debugging without exposing internals to the client.

---

### 🔜 Future Enhancements

* Add authentication to manage user-specific short URLs.
* Track geo-location and device statistics for analytics.
* Use Redis for caching frequently accessed URLs.
* Build a UI dashboard for usage tracking and link management.
* Periodic job to clean expired links.
