# sw-go-template-server

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Welcome to the `sw-go-template-server` repository! This foundational Golang server template offers flexible HTTPS/HTTP support out-of-the-box. Designed for extensibility, it easily integrates hardware I/O, various databases, and other functionalities, making it ideal for a wide range of applications.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Flexible HTTP/HTTPS Support**: Configurable for secure (HTTPS) and regular (HTTP) communication.
- **Modular Architecture**: Easily integrate additional modules like databases, hardware I/O, and custom functionalities.
- **Scalability**: Efficiently handles both small and large-scale deployments.
- **Basic Authentication**: Implements basic authentication using JWT tokens to secure endpoints.

## Getting Started

To get a local copy up and running, follow these installation steps.

### Prerequisites

- [Go (Golang)](https://golang.org/dl/) - Tested with version `1.20`, but newer versions are likely compatible.

### Installation

1. **Clone the Repo**

    ```bash
    git clone https://github.com/tpasson/sw-go-template-server.git
    ```

2. **Navigate to the Directory**

    ```bash
    cd sw-go-template-server
    ```

3. **Choose the Environment in build.sh**

    Uncomment your choice:

    ```bash
    GOOS=linux GOARCH=amd64 go build -o $SCRIPT_DIR/$PACKAGE_NAME-amd64-linux $SCRIPT_DIR/main.go
    #GOOS=linux GOARCH=arm64 go build -o $SCRIPT_DIR/$PACKAGE_NAME-arm64-linux $SCRIPT_DIR/main.go
    #GOOS=darwin GOARCH=arm64 go build -o $SCRIPT_DIR/$PACKAGE_NAME-arm64-darwin $SCRIPT_DIR/main.go
    ```

4. **Build the Project**

    ```bash
    ./build
    ```

## Usage

After successfully building the project:

1. **Start the Server**

    For HTTP:

    ```bash
    ./sw-go-template-server -http="localhost:8080" -tlscert="" -tlskey=""
    ```

    For HTTPS:

    ```bash
    ./sw-go-template-server -http="localhost:8443" -tlscert="/path/to/fullchain.pem" -tlskey="/path/to/privkey.pem"
    ```

    Visit `http://localhost:8080` for HTTP or `https://localhost:8443` for HTTPS to access the server.

    **Note**: For HTTPS, configure SSL/TLS certificates appropriately. Follow [this guide](https://letsencrypt.org/getting-started/) for obtaining free certificates from Let's Encrypt.

2. **Testing**

    You can use `curl` to test the issuing of certificates (cookie tokens for authentication). Follow the steps below:

    **Login to Obtain Authentication Token**

    Use the following `curl` command to log in and save the authentication token in a cookie file (`cookies.txt`):

    ```bash
    curl -v -X POST http://localhost:8080/login -d '{"username":"your_username", "password":"your_password"}' -H "Content-Type: application/json" -c cookies.txt
    ```

    **Access Protected Content with Authentication Token**

    After logging in, use the saved cookie to access the protected content:

    ```bash
    curl -v -X GET http://localhost:8080/content -b cookies.txt
    ```

    **Verify Access Without Authentication**

    To verify that the server does not send the protected content without authentication, use the following command without specifying the cookie file:

    ```bash
    curl -v -X GET http://localhost:8080/content -b
    ```

    Replace `your_username` and `your_password` with your actual login credentials.

## Configuration

- The server can be configured to run in either HTTP or HTTPS mode.
- SSL/TLS certificates are required for HTTPS.
- Command-line flags allow dynamic configuration of server parameters like address and SSL certificate paths.

## Contributing

We welcome contributions! Here's how you can contribute:

1. **Fork the Project**: Use the 'Fork' button at the top right of this page.
2. **Clone Your Fork**: 

    ```bash
    git clone https://github.com/YOUR_USERNAME/sw-go-template-server.git
    ```

3. **Navigate to Your Clone**:

    ```bash
    cd sw-go-template-server
    ```

4. **Create a New Branch**: 

    ```bash
    git checkout -b new-feature
    ```

5. **Make Changes**: Add new features or fix bugs.
6. **Push**: 

    ```bash
    git push origin new-feature
    ```

7. **Open a Pull Request**: Go back to your fork on GitHub and click 'New pull request'.

## License

Distributed under the MIT License. See `LICENSE` for more information.

---

Crafted with pass(i)on by [tpasson](https://github.com/tpasson)!