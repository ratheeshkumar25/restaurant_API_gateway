***Restaurant Management API Gateway***
This is the version 1 of the Restaurant Management project, where three services have been converted into microservices using the API Gateway pattern. Each service has its own database to ensure scalability. The application is designed so that it can be scaled at any time based on business requirements.

The application workflow includes user sign-up via mobile number, OTP verification for account creation, and admin login with predefined credentials provided in the Admin service.

***Microservices Included:***
>>Restaurant User Service
>>Restaurant Menu Service
>>Restaurant Admin Service
Setup and Workflow:
Clone the Repositories:
Clone all four repositories (three services and the API gateway) to your local machine.

Database Configuration:
Set up the necessary database configurations for each service as per the individual service requirements.

Redis and Twilio Configuration:
Configure Redis for temporary data storage and Twilio for OTP verification in the User service.

Run the Services:
Start all four services (User Service, Menu Service, Admin Service, and API Gateway).

API Testing with Postman:
You can test the API endpoints using Postman. Below are the sample API requests:

User Service:

Sign-up:
POST http://localhost:8081/api/user/signup
OTP Verification:
POST http://localhost:8081/api/user/verify-otp
Login:
POST http://localhost:8081/api/user/login
Admin Service:

Before Login:
Set up the Admin credentials in the database:
Username: AdminRERA
Password: admin@123
Login:
POST http://localhost:8081/api/admin/login
Add Menu:
POST http://localhost:8081/api/admin/auth/menu
List All Menus:
GET http://localhost:8081/api/admin/auth/menus
List Menu by ID:
GET http://localhost:8081/api/admin/auth/menu?id=3
List Menu by Name:
GET http://localhost:8081/api/admin/auth/menu?name=Hyderbad%20Biryani
