# Release

- Create a new draft release : https://github.com/c100k/ddot/releases/new
- Add the meaningful commit messages from `git log --oneline`
- Bump the version in `main.go`
- Bump the version below and execute the commands :
    ```sh
    version="0.1.0-beta.2"

    docker run --rm -v $(pwd):/app golang:1.23 bash -c "(cd /app && GOARCH=amd64 GOOS=darwin go build -o /app/dist/ddot-${version}-darwin-amd64 -v)"
    docker run --rm -v $(pwd):/app golang:1.23 bash -c "(cd /app && GOARCH=arm64 GOOS=darwin go build -o /app/dist/ddot-${version}-darwin-arm64 -v)"
    docker run --rm -v $(pwd):/app golang:1.23 bash -c "(cd /app && GOARCH=amd64 GOOS=linux go build -o /app/dist/ddot-${version}-linux-amd64 -v)"
    docker run --rm -v $(pwd):/app golang:1.23 bash -c "(cd /app && GOARCH=arm64 GOOS=linux go build -o /app/dist/ddot-${version}-linux-arm64 -v)"

    (cd dist && ./ddot-${version}-darwin-arm64 version)
    ```
- Upload the generated binaries to the release
- Publish the release
