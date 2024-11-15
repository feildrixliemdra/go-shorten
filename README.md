## Getting Started
### Prerequisites
- Go (version 1.16 or higher)
- Docker (optional for containerization)
- Make

### Installation

1. Clone the repo
   ```sh
    git clone https://github.com/feildrixliemdra/go-shorten.git
    cd go-shorten
    ```
2. Install dependencies
   ```sh
   go mod tidy
   ```
3. Set up environment variables:
   ```sh
   make env 
   or 
   cp config/config.yaml.example config/config.yaml  #fill in necessary values.
   ```
4. Change git remote url to avoid accidental pushes to base project
   ```sh
   git remote set-url origin your_github_username/repo_name
   git remote -v # confirm the changes
   ```
5. Run the application:
    ```sh
    go run main.go serve-http
    ```
   Or Using Docker
    ```sh
   docker-compose up --build
    ```
   Access the application at http://localhost:<your_port>.