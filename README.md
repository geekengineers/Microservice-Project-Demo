# ☠️ Microservice Project Demo
This repository provides a well-structured example of building microservices with a focus on **learning purposes**. It demonstrates the use of RESTful APIs with **Protocol Buffers (Proto3)**, and containerization with **Docker** and **Docker Compose** and deployment with **Terraform**.

## Services Explored
The project consists of two microservices:

- Auth: Handles user authentication and authorization functionalities.  
- Blog: Manages blog posts, including creation, retrieval, and updates.
  
### Key Technologies
The project leverages several key technologies:

- RESTful APIs: Both services expose APIs that follow the REST architectural style for communication between services and clients.  
- Proto3: The project utilizes Proto3, a language-neutral mechanism for defining data structures and remote procedure calls (RPCs). This promotes communication efficiency and ensures consistency between services.  
- connectrpc.com (HTTP2): This project might be using [connectrpc.com](https://connectrpc.com), a library that allows implementing gRPC functionalities over HTTP/2. It could facilitate communication between services using the gRPC framework. However, this requires further confirmation from the actual codebase.  
- Docker: Containers are employed to package and isolate the microservices, ensuring each service runs in a self-contained environment.  
- Docker Compose: This tool simplifies the management of multi-container applications like this microservice project. It allows defining and running all dependent services with a single command.
- PostgreSQL: As the shared database between services (Auth, Blog)
- GORM: GoLang rich-feature ORM
- Hexagonal Architecture: Used as the main architecture of the services with a beautiful bootstrap tricks that facilitates port adapting and interaction with config package.

## Deployment Process
The project utilizes Docker and Docker Compose for deployment:

Also it includes a Terraform config to deploy easily at any server using by SSH connection! 

## 👽 Open To Contribute

How to Contribute:

1. Fork the Repository: Fork the repository to your own GitHub account.
1. Create a Branch: Create a new branch for your feature or bug fix.
1. Make Changes: Make your changes and commit them to your branch.
1. Submit a Pull Request: Open a pull request to merge your changes back into the main repository.
