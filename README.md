# Bunny PV Demo

A simple Go web application with a Tailwind-styled interface for editing and saving text content.

## Features

- Clean, modern UI with Tailwind CSS
- Textarea for content editing
- Save functionality that persists content to `/test_pv1/content.txt`
- Automatic content loading on page refresh
- Responsive design

## How to Run

### Option 1: Run with Go directly

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

### Option 2: Run with Docker

1. Build the Docker image:
   ```bash
   docker build -t bunny-pv-demo .
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 -v /test_pv1:/test_pv1 bunny-pv-demo
   ```

3. Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

**Note:** The `-v /test_pv1:/test_pv1` flag mounts the host's `/test_pv1` directory to the container, so your saved content will persist even when the container is stopped.

## How it Works

- The application creates a `/test_pv1/` directory if it doesn't exist
- Content is saved to `/test_pv1/content.txt` when the "Save Content" button is clicked
- When the page loads, it automatically reads and displays the saved content
- The application uses the Gorilla Mux router for HTTP routing

## File Structure

```
bunny-pv-demo/
├── main.go              # Main application code
├── go.mod              # Go module file
├── templates/
│   └── index.html      # HTML template with Tailwind CSS
└── README.md           # This file
```

## Docker Commands

### Build the image:
```bash
docker build -t bunny-pv-demo .
```

### Run the container:
```bash
docker run -p 8080:8080 -v /test_pv1:/test_pv1 bunny-pv-demo
```

### Run in detached mode (background):
```bash
docker run -d -p 8080:8080 -v /test_pv1:/test_pv1 --name bunny-app bunny-pv-demo
```

### Stop the container:
```bash
docker stop bunny-app
```

### Remove the container:
```bash
docker rm bunny-app
```

### View container logs:
```bash
docker logs bunny-app
```

## Dependencies

- `github.com/gorilla/mux` - HTTP router and URL matcher
- Tailwind CSS (loaded via CDN)
