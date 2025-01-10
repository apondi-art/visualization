# Groupie Trackers

**Groupie Trackers** is a web-based platform designed to receive, manipulate, and display data from a given API, providing comprehensive information about bands and artists. The platform visualizes key details such as band names, images, founding years, concert locations, and dates in an interactive and user-friendly manner.


## Project Overview

This project focuses on retrieving and displaying the following key elements from an external RESTful API:

- **Artists**: Information about various bands and artists, including their name(s), image, start year, first album release date, and members.
- **Locations**: A list of past and upcoming concert locations.
- **Dates**: Concert dates, both past and upcoming.
- **Relation**: Data that links the artists, concert dates, and locations together.

## Key Features

- **Data Visualization**: Information is displayed using blocks, cards, tables, lists, pages, graphs, and other visual formats.
- **Client-Server Communication**: Event-driven features allow client-server requests, ensuring that data updates dynamically. For example, you can trigger events such as fetching new concert dates or updating artist information.
- **Error Handling**: The website handles errors gracefully, including issues related to API requests or invalid data, ensuring smooth operation at all times.
- **Backend in Go**: The backend is written entirely in Go, leveraging its concurrency and performance strengths.
- **Good Code Practices**: The code follows best practices, including unit tests for critical functionality to ensure reliability and maintainability.

## Getting Started

### Prerequisites

- Go (1.16 or later)
- HTML/CSS/JavaScript for frontend development

### Installation

1. Clone the repository:

   ```bash
   git clone https://learn.zone01kisumu.ke/git/quochieng/groupie-tracker.git
   ```

2. Navigate to the project directory:

   ```bash
   cd groupie-trackers
   ```

3. Run the Go server:

   ```bash
   cd cmd/server
   go run .
   ```

4. Open the website in your browser:

   ```
   http://localhost:8080


   ## Usage

Once the server is running, the homepage will display a list of bands and artists. Each band can be clicked to view detailed information such as:

- Band Members
- Start Year and First Album Release Date
- Upcoming and Past Concert Locations
- Concert Dates

You can also trigger specific events (such as fetching the latest data) that send client requests to the server and update the information in real-time.

## API

The application communicates with an external RESTful API that provides the data for artists, locations, and concert dates. The relationship between these entities is maintained by the API.

## Contributing

Contributions are welcome! If you would like to contribute:

1. Fork the project.
2. Create your feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

## Authors

- Quinter Ochieng
- Mokwa Moffat


## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

Special thanks to all the contributors and the open-source community.

---