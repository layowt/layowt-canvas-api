## Dockerfile

# Run

docker build -t [image-name] .

# Then

docker run -p 8080:8080 --rm -v $(pwd):/app -v /app/tmp --name layowt-canvas-api-air [image-name]
