# SE2-leaseEase-backend

## 📂 Project Setup

1. Clone the Repository, then

`cd LeaseEase`

2. Initialize the Project

`go mod init LeaseEase`

3. Install Dependencies
   Run the following commands to install the necessary dependencies:

## Docker Setup

1. Create a Docker Compose file  
Create a file named `docker-compose.yml` in the project root with content similar to the example below:
```yaml

```

2. Launch the Container  
Run the following command to build and start the container:
```bash
docker-compose up --build
```

## Environment Setup

Copy the example environment file and update your environment variables accordingly.

For Bash (Linux/macOS Terminal):
```bash
cp .env.example .env
```
For Windows PowerShell:
```powershell
Copy-Item .env.example .env
```

Then, open the `.env` file and configure your environment variables as needed.

## Run Application

To run the main application, execute the following command:

```bash
go run cmd/main.go
```