# 💵💰💳 Paper Wallet Core Service 💻⚙️

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

3. Update swagger documentation
   ```bash
   make swag-init
   ```

4. Run the service:
   ```bash
   make run
   ```

## 📌 API Documentation

Access the API documentation at:
```
{host}/swagger/index.html
```

### Disburse Balance Endpoint

**POST** `{host}/paper-wallet-core-service/wallet/withdraw`

#### Example cURL Request
```
curl -X 'POST' \
  'http://localhost:8080/paper-wallet-core-service/wallet/withdraw' \
  -H 'accept: application/json' \
  -H 'X-Channel-Id: iOS' \
  -H 'X-Request-Id: cb41cafb-8b35-439a-aa05-30e022a4f323' \
  -H 'X-Service-Id: gateway' \
  -H 'X-Username: muhammad.rezki.ananda@gmail.com' \
  -H 'Accept-Language: en' \
  -H 'User-Agent: User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_2 like Mac OS X' \
  -H 'Authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=' \
  -H 'X-App-Version: 1.2.3-4' \
  -H 'Content-Type: application/json' \
  -d '{
  "amount": 1,
  "userId": "38690cf6-4efb-454b-b641-5d4b4f71f5b3"
}'
```
#### Request Headers

- **accept:** application/json
- **X-Channel-Id:** Channel identifier
- **X-Request-Id:** Unique request identifier
- **X-Service-Id:** Service identifier (e.g., gateway)
- **X-Username:** User identifier (e.g., muhammad.rezki.ananda@gmail.com)
- **Accept-Language:** Preferred language (e.g., id & en)
- **User-Agent:** Client information
- **Authorization:** Authentication
- **X-App-Version:** Application information
- **Content-Type:** application/json

#### Request Body

```json
{
  "amount": "number", 
  "userId": "string"
}
```

#### Response

English Language (default)
```json
{
   "code": "SUCCESS",
   "message": "SUCCESS",
   "data": {
      "message": "Success! You have successfully requested to disburse `IDR 1`. Your previous balance was `IDR 1,499,999`, and after the disbursement, your new balance is `IDR 1,499,998`",
      "detail": {
         "id": "38690cf6-4efb-454b-b641-5d4b4f71f5b3",
         "name": "Muhammad Rezki",
         "currency": "IDR",
         "scale": 2,
         "balance": "1499998"
      }
   },
   "errors": null,
   "serviceTime": 1738446802
}
```

Bahasa Indonesia
```json
{
   "code": "SUCCESS",
   "message": "SUCCESS",
   "data": {
      "message": "Sukses! Anda berhasil meminta untuk mencairkan `IDR 1`. Saldo Anda sebelumnya adalah `IDR 1,500,000`, dan setelah pencairan, saldo baru Anda menjadi `IDR 1,499,999`",
      "detail": {
         "id": "38690cf6-4efb-454b-b641-5d4b4f71f5b3",
         "name": "Muhammad Rezki",
         "currency": "IDR",
         "scale": 2,
         "balance": "1499999"
      }
   },
   "errors": null,
   "serviceTime": 1738446784
}
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

