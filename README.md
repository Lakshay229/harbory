# Harbory

Harbory is a cross-platform Docker management tool built with Go and React. It provides a web interface to manage your Docker containers, images, and system information.

## Features

- 📦 **Container Management**: View, start, stop, and remove containers
- 🖼️ **Image Management**: Browse, pull, and delete Docker images
- 📊 **System Information**: Monitor Docker system resources and performance
- 🔒 **Secure API**: Communication between frontend and backend using RESTful API
- 🌐 **Web Interface**: Modern and responsive UI built with React

## Architecture

Harbory consists of three main components:

1. **Backend**: Go API server that interacts with the Docker daemon
2. **Frontend**: React application that provides the user interface
3. **Nginx**: Web server that serves as a reverse proxy

## Prerequisites

- [Docker](https://www.docker.com/get-started) (v20.10.0+)
- [Docker Compose](https://docs.docker.com/compose/install/) (v2.0.0+)

## Getting Started

### Installation

1. Clone the repository:

```bash
git clone https://github.com/preetindersinghbadesha/harbory.git
cd harbory
```

2. Start the application:

```bash
docker-compose up -d
```

3. Access the web interface at [http://localhost:3000](http://localhost:3000)

### Development Setup

#### Backend (Go)

```bash
cd harbory-backend
go mod tidy
go run main.go
```

The backend server will be available at [http://localhost:8080](http://localhost:8080)

#### Frontend (React)

```bash
cd harbory-frontend
npm install
npm run dev
```

The development server will be available at [http://localhost:3000](http://localhost:3000)

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | `/api/containers` | List all containers |
| GET    | `/api/images` | List all images |
| GET    | `/api/system/info` | Get system information |

## Technologies Used

- **Backend**:
  - [Go](https://golang.org/)
  - [Docker Engine API](https://docs.docker.com/engine/api/)
  - [Gorilla Mux](https://github.com/gorilla/mux)

- **Frontend**:
  - [React](https://reactjs.org/)
  - [Vite](https://vitejs.dev/)
  - [Nginx](https://nginx.org/)

- **DevOps**:
  - [Docker](https://www.docker.com/)
  - [Docker Compose](https://docs.docker.com/compose/)

## Project Structure

```
harbory/
├── docker-compose.yml        # Docker Compose configuration
├── harbory-backend/          # Go backend service
│   ├── Dockerfile            # Backend Docker configuration
│   ├── main.go               # Entry point
│   ├── handlers/             # API route handlers
│   ├── router/               # API route definitions
│   └── utils/                # Utility functions
├── harbory-frontend/         # React frontend application
│   ├── Dockerfile            # Frontend Docker configuration
│   ├── public/               # Static assets
│   └── src/                  # React components and logic
└── nginx/                    # Nginx reverse proxy configuration
    ├── Dockerfile            # Nginx Docker configuration
    └── default.conf          # Nginx configuration file
```

## Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.
