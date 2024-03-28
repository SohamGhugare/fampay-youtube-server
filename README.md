# Fampay Backend Assignment

### Problem Statement

To make an API to fetch latest videos sorted in reverse chronological order of their publishing date-time from YouTube for a given tag/search query in a paginated response.

### Tech Stack Used

- Gin Gonic (Golang): API
- Gorm (Golang): ORM (for database operations)
- YouTube v3 API

### Running the server
I've containerized the whole project using Docker, so running it is pretty simple. <br />

**1. Building the image** <br />
```docker build -t fampay-backend .``` <br /> <br />
**2. Running the image** <br />
```API_KEY="your api key here" docker run -it -p 8080:8080 fampay-backend```
