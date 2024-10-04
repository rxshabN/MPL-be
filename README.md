<div align="center">
  <h3 align="center">MPL Backend</h3>

  <p align="center">
    A robust backend API for MPL (Multi-Purpose League) platform
    <br />
    <a href="https://github.com/SIAM-VIT/MPL-be"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/SIAM-VIT/MPL-be/issues">Report Bug</a>
    ·
    <a href="https://github.com/SIAM-VIT/MPL-be/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
## Table of Contents

- [About The Project](#about-the-project)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

<!-- ABOUT THE PROJECT -->
## About The Project

This is the backend API for the Multi-Purpose League (MPL) platform, which provides a service for managing leagues, games, and teams. The platform handles user registration, game scheduling, and match result management.

Key Features:
- User authentication and authorization
- League and team management
- Match scheduling and result tracking
- Real-time game updates

### Built With

This project is built using the following technologies and frameworks:

- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)

<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running, follow these simple steps.

### Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/doc/install)
- PostgreSQL 13 or higher
- Docker (optional but recommended)

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/SIAM-VIT/MPL-be.git

2. Change into the project directory
   ```sh
   cd MPL-be

3. Install dependencies
   ```sh
   go mod download

4. Set up PostgreSQL and configure .env with your database credentials:
   ```sh
    DB_HOST=localhost
    DB_USER=postgres
    DB_PASSWORD=yourpassword
    DB_NAME=mpl
    JWT_SECRET_KEY=your_jwt_secret_key

6. Start the application
   ```sh
   go run cmd/main.go

<!-- USAGE EXAMPLES -->
## Usage
Once your application is running, you can use the following API endpoints to interact with the MPL platform:

POST /api/v1/auth/register - Register a new user
POST /api/v1/auth/login - Login user
GET /api/v1/leagues - Get all leagues
POST /api/v1/matches - Schedule a new match
For more examples, please refer to the Documentation

<!-- CONTRIBUTING -->
## Contributing
Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are greatly appreciated.

## Fork the Project
Create your Feature Branch (git checkout -b feature/your-feature)
Commit your Changes (git commit -m 'Add some feature')
Push to the Branch (git push origin feature/your-feature)
Open a Pull Request
<!-- LICENSE -->
## License
Distributed under the MIT License. See LICENSE for more information.


<p align="center">
	Made with :heart: by SIAM-VIT
</p>