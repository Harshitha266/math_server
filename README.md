
Math Server:

This is a simple Go application that provides two APIs for performing mathematical operations: addition and subtraction. The server listens on port 8080 and exposes the following endpoints:

/math/add: Takes two numbers as query parameters (num1 and num2) and returns their sum.
/math/subtract: Takes two numbers as query parameters (num1 and num2) and returns their difference.
Getting Started
To run the math server, follow these steps:

Make sure you have Go installed on your machine. If not, you can download it from the official website.
Clone this repository to your local machine:
"git clone https://github.com/your-username/math-server.git"
Navigate to the project directory:
"cd math-server:
Run the server:
"go run main.go"
This will start the server on port 8080.
Usage:
Addition API:
To perform addition, make a GET request to /math/add with the num1 and num2 query parameters. For example:

'curl "http://localhost:8080/math/add?num1=10&num2=20" '
This will return the sum of num1 and num2.

Subtraction API:
To perform subtraction, make a GET request to /math/subtract with the num1 and num2 query parameters. For example:

' curl "http://localhost:8080/math/subtract?num1=20&num2=10" '
This will return the difference of num1 and num2.
