# Wallet Balance Disbursement Service

A microservice that facilitates the disbursement of user balances from an application wallet. Built with Go 1.23, Echo framework, and fully documented using Swagger.

## 🚀 Features

- **Secure Disbursement:** Ensures safe and verified balance disbursement.
- **CSV Data Handling:** Reads and processes user balance data from CSV files.
- **RESTful API:** Fast, scalable, and lightweight API powered by Echo.
- **Comprehensive Documentation:** API endpoints documented with Swagger for easy integration.

## 🛠️ Tech Stack

- **Language:** Go 1.23
- **Framework:** Echo
- **Documentation:** Swagger
- **Data Source:** CSV files

## 📊 Sample Data (CSV Format)

The service processes balance data sourced from CSV files. Here's a sample of the data:

| _id                                   | name           | currency | scale | balance | createdBy | createdDate          | updatedBy                       | updatedDate          | Version | IsDeleted |
|--------------------------------------|----------------|----------|-------|---------|-----------|----------------------|---------------------------------|----------------------|---------|-----------|
| 38690cf6-4efb-454b-b641-5d4b4f71f5b3 | Muhammad Rezki | IDR      | 2     | 1500000 | system    | 2025-02-01T10:00:00Z | muhammad.rezki.ananda@gmail.com | 2025-03-01T09:40:00Z | 32     | 0        |


## ⚙️ Getting Started

### Prerequisites

- Go 1.23 installed
- CSV file with the required structure

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/mrrezkii/PAPER-WALLET-CORE-SERVICE.git
   cd PAPER-WALLET-CORE-SERVICE
   ```

2. Install dependencies:
   ```bash
   make deps
   ```

3. Run the service:
   ```bash
   make run
   ```

## 📌 API Documentation

Access the API documentation at:
```
http://localhost:8080/swagger/index.html
```

## 📂 Project Structure

```
PAPER-WALLET-CORE-SERVICE/
├── bin/
├── cmd/main.go
├── config/config.go
├── data/users.csv
├── docs/* (Swagger files)
├── internal/* (Clean Architecture files)
└── shared/*
```

## 🤝 Contributing

Contributions are welcome! Please fork the repository, create a feature branch, and submit a pull request.

## 🙏 Acknowledgements

- Thanks to the Echo framework and Swagger communities for their excellent tools and documentation.

