# Ecommerce Marketplace Platform

Welcome to the **Ecommerce Marketplace Platform** repository! This project is a fully functional online marketplace where users can buy, sell, and post products. Inspired by platforms like **HYPD** in India and global marketplaces such as **Etsy**, **Depop**, and **Poshmark**, this platform provides a seamless experience for both buyers and sellers. Built with **Go (Golang)** and the **Goa Framework**, it is scalable, efficient, and easy to deploy.

---

## Features

- **User Roles**:
  - **Buyers**: Browse products, add to cart, and place orders.
  - **Sellers**: Post products, manage listings, and track sales.
- **Product Listings**: Sellers can create, update, and delete product listings with details like price, description, and images.
- **Search and Filters**: Buyers can search for products and filter by category, price range, and more.
- **Shopping Cart**: Add, remove, and update items in the cart.
- **Order Management**: Track orders, view order history, and manage deliveries.
- **User Profiles**: Personalized profiles for buyers and sellers.
- **Authentication**: Secure login and registration using **JWT (JSON Web Tokens)**.
- **Database**: Robust data storage with **PostgreSQL**.
- **Responsive Design**: Optimized for desktop, tablet, and mobile devices.
- **Admin Dashboard**: Manage users, products, and orders (admin-only).

---

## Technologies Used

- **Backend**: Go (Golang)
- **Framework**: Goa Framework (DSL-based API design)
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **API Design**: Goa DSL for clean and scalable API definitions
- **Deployment**: Docker, or any cloud platform (e.g., AWS, GCP)

---

## Getting Started

Follow these steps to set up the project locally on your machine.

### Prerequisites

- Go (v1.20 or higher)
- PostgreSQL (local or cloud instance)
- Goa Framework (install via `go install goa.design/goa/v3/cmd/goa@v3`)

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/prth1234/Ecommerce.git
   cd Ecommerce
   ```

2. **Set up the database**:
   - Create a PostgreSQL database and update the connection string in the configuration file (e.g., `config.yaml` or `.env`).
   - Example connection string:
     ```env
     DB_HOST=localhost
     DB_PORT=5432
     DB_USER=your_db_user
     DB_PASSWORD=your_db_password
     DB_NAME=your_db_name
     ```

3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

4. **Build and run the application**:
   - Navigate to the `cmd/store` directory:
     ```bash
     cd cmd/store
     ```
   - Build the application:
     ```bash
     go build
     ```
   - Run the application:
     ```bash
     ./store
     ```
Shortcut : Once go module are downloaded, just run 
  ```bash
lsof -t -i :8000 | xargs -r kill -9 && go build && ./store 
  ```
5. **Access the application**:
   - The API server will start running at `http://localhost:8080`.
   - Use tools like **Postman** or **cURL** to interact with the API endpoints.

---

## Folder Structure

```
Ecommerce/
├── cmd/
│   └── store/               # Main application entry point
├── gen/                     # Generated code by Goa Framework
├── design/                  # API design definitions (Goa DSL)
├── internal/                # Application logic (services, models, etc.)
├── migrations/              # Database migration scripts
├── scripts/                 # Helper scripts for setup and deployment
├── config.yaml              # Configuration file for the application
├── go.mod                   # Go module dependencies
├── go.sum                   # Go dependency checksums
└── README.md                # Project documentation
```

---

## Contributing

We welcome contributions! If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeatureName`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/YourFeatureName`).
5. Open a pull request.

Please ensure your code follows the project's coding standards and includes appropriate tests.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- Inspired by platforms like **HYPD**, **Etsy**, **Depop**, and **Poshmark**.
- [Goa Framework Documentation](https://goa.design/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [JWT Introduction](https://jwt.io/introduction)

---

## Contact

For questions or feedback, feel free to reach out:

- **Author**: Parth  
- **Email**: dearparthsingh@gmail.com  
- **GitHub**: [prth1234](https://github.com/prth1234)  

---

Thank you for checking out this project! We hope you find it useful and welcome any suggestions for improvement. Happy coding! 🚀
```

---

### Key Highlights:
1. **Tech Stack**: Clearly mentions **Go (Golang)**, **Goa Framework**, **PostgreSQL**, **JWT**, and **DSL**.
2. **Running the Application**: Provides specific instructions for building and running the app using `go build` and `./store`.
3. **Folder Structure**: Explains the purpose of each folder in the project.
4. **Database Setup**: Includes steps for setting up PostgreSQL.

Let me know if you need further adjustments! 🚀
