# Go Sleeping Barber

This project simulates the Sleeping Barber problem using Go. It models a barbershop where customers arrive randomly, and a barber cuts their hair. If no customers are present, the barber sleeps until a new customer arrives.

## Project Structure

- `internal/customer/customer.go`: Defines the `Customer` struct and a function to create new customers.
- `internal/shop/barbershop.go`: Contains the `BarberShop` struct and methods to manage the barbershop operations.

## Dependencies

- `github.com/google/uuid`: For generating unique customer IDs.
- `go.uber.org/zap`: For logging.

## Getting Started

### Prerequisites

- Go 1.23.2 or later

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/LeonardoBerlatto/go-sleeping-barber.git
    cd go-sleeping-barber
    ```

2. Download the dependencies:
    ```sh
    go mod download
    ```

### Running the Simulation

To run the barbershop simulation:
```sh
go run ./...