
# YouTube Video Fetcher

YouTube Video Fetcher is a GoLang-based backend service that fetches videos from the YouTube API using multiple API keys and stores them in a PostgreSQL database. The service rotates through a list of API keys to ensure that the quota is not exceeded, and supports fetching videos based on a single search query.

## Project Overview

### What It Does

- **Fetches YouTube Videos:** The service uses the YouTube Data API to fetch the latest videos based on a specified search query.
- **API Key Rotation:** It rotates through multiple API keys to avoid hitting the quota limit of any single key.
- **Stores Data:** The fetched video data (title, description, publication date, and thumbnail URL) is stored in a PostgreSQL database.
- **APIs for Video Retrieval:** The project exposes RESTful APIs to retrieve videos from the database.

### Technologies Used

- **GoLang**: The primary language used to develop the backend service.
- **PostgreSQL**: The relational database used to store video data.
- **GORM**: The ORM (Object Relational Mapping) library used to interact with the PostgreSQL database.
- **YouTube Data API**: The external API used to fetch video data.
- **Docker**: Used for containerizing the application.

## Project Structure

```bash
├── config
│   └── config.yaml  
├── internal
│   ├── db           
        └── databse.go 
│   ├── handlers      
        └── handlers.go 
│   └── yt            
        └── youtube.go 
│   └── models 
        └── models.go           
├── main.go           
└── Dockerfile       
```
## To run the project
Clone the repository and configure the docker and config.yaml files accordingly
**Run Without Docker**
- Install Postgres and Go
- Run 
- `go mod tidy`
- `go build -o main`
- `./main`

**With Docker**
- Install docker and docker compose
-  Run 
- `docker-compose build`
- `docker-compose up`
- *Side Note- If it doesn't run in the first try, retry after ''Ctrl+C*

  **Hosted**
  - This has been hosted to collect data at render.
    ![Hosted and Connected][https://github.com/AbhAsg09/20MIS0158_Backend/blob/134e1cde906472c687049ca0a4dbb26cbf31d7c7/Images/Deployed.png]
  - To connect to render psql database run-
  - `PGPASSWORD=fJomBjck74yooAnOeqdZF7hUbhAOXndl psql -h dpg-cqru6ndumphs73cqo98g-a.oregon-postgres.render.com -U abhishek youtube_videos_lf49`
