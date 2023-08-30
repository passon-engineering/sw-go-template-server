# sw-go-template-server

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Welcome to the `sw-go-template-server` repository! This is a foundational Golang server template that offers HTTPS/HTTP support out-of-the-box. It is designed with extensibility in mind, allowing you to easily integrate hardware I/O, various databases, and other functionalities as you see fit.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Basic HTTP/HTTPS Support**: Ready to launch with security in mind.
- **Modular Architecture**: Designed for seamless integration of additional modules such as databases, hardware I/O, and other custom functionality.
- **Scalability**: Efficiently designed for both small and large-scale deployments.

## Getting Started

To get a local copy up and running, follow the installation steps.

### Prerequisites

- [Go (Golang)](https://golang.org/dl/) - Tested with version `1.x`, but newer versions are likely compatible.

### Installation

1. **Clone the Repo**

    ```bash
    git clone https://github.com/passon-engineering/sw-go-template-server.git
    ```

2. **Navigate to the Directory**

    ```bash
    cd sw-go-template-server
    ```

3. **Build the Project**

    ```bash
    go build
    ```

## Usage

After successfully building the project:

1. **Start the Server**

    ```bash
    ./sw-go-template-server
    ```

Visit `http://localhost:8080` (or `https://localhost:8443` for HTTPS) to access the server.

**Note**: You might need to configure SSL/TLS certificates for HTTPS. Follow [this guide](https://letsencrypt.org/getting-started/) for obtaining free certificates from Let's Encrypt.

## Contributing

We welcome contributions! If you'd like to contribute:

1. **Fork the Project**: Use the 'Fork' button at the top right of this page.
2. **Clone Your Fork**: `git clone https://github.com/YOUR_USERNAME/sw-go-template-server.git`
3. **Navigate to Your Clone**: `cd sw-go-template-server`
4. **Create a New Branch**: `git checkout -b new-feature`
5. **Make Changes**: Add your new features or bug fixes.
6. **Push**: `git push origin new-feature`
7. **Open a Pull Request**: Navigate back to your fork on GitHub and click 'New pull request'.

## License

Distributed under the MIT License. See `LICENSE` for more information.


---

Crafted with pass(i)on by the [passon-engineering](https://github.com/passon-engineering) team!