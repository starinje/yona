# Yona

<div align="center">

![A bear](images/yona.png)

</div>

## Overview

Yona is a comprehensive SaaS boilerplate designed for building scalable, production-ready SaaSplatforms. It provides developers with a straightforward starting point for creating enterprise-grade applications.

Often with new SaaS projects, a significant amount of time is often consumed in selecting the right tech stack and setting up boilerplate code. This process can be tedious and overwhelming, especially with the vast array of technologies available today.

Yona addresses this challenge by covering everything you need from the outset, including local environments, CI/CD deployments, hosting environments, database migrations, and more. With Yona, you can bypass the lengthy setup process and focus on building your application specific business logic.

## 🚀 Technologies

- **Frontend**: React + TypeScript
- **Backend**: Go + Gin
- **Database**: PostgreSQL, Prisma, Neon
- **Local Dev**: Nix, Docker, Docker Compose
- **Infrastructure-As-Code**: Terraform
- **Cloud**: AWS
- **Authentication**: Auth0
- **Observabilty**: Datadog

## 🛠 Local Development

### 1. Clone the Repository

```bash
git clone https://github.com/starinje/yona.git
cd yona
```

### 2. Install Nix

This project uses Nix to ensure that developers' local environments are all the same.

To install Nix, run the following command at the root level of the `yona` directory:

```bash
make install-nix
```

### 3. Start Nix

To enter the Nix shell and set up your development environment, run the following command in your terminal:

```bash
nix-shell
```

This command creates an isolated environment based on your `shell.nix` file, allowing you to access all specified tools and dependencies. You can exit the Nix shell by typing `exit` or pressing `Ctrl + D`. For more information about Nix, visit the [Nix website](https://nixos.org/nix/).

### 4. Start the Application

To start the application, run:

```bash
make start
```

This command will launch the application in the background.

### 5. Stop the Application

To stop the application, use:

```bash
make stop
```

This command will terminate the running application.

### 6. Run Backend Tests

To run the tests for the application, execute:

```bash
make test-backend
```

This command will run all the defined tests to ensure everything is functioning correctly.

### More to come, if you see something that should be here, please add it.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch from the development branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request against the development branch

## 📄 License

MIT License

## 🌟 Acknowledgements

- [What is a Yona](https://nativehistoryassociation.org/tutor_tsalagi2_study.php)
