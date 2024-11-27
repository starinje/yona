# Yona

<div align="center">

![A bear](images/yona.png)

</div>

## Overview

Yona is a comprehensive boilerplate designed for building scalable, production-ready SaaS platforms. It provides developers with a straightforward starting point for creating enterprise-grade applications.

Often with new SAAS projects, a significant amount of time is often consumed in selecting the right tech stack and setting up boilerplate code. This process can be tedious and overwhelming, especially with the vast array of technologies available today. Developers frequently find themselves debating the merits of various frameworks, libraries, and tools, which can delay the actual development of business logic.

Yona addresses this challenge by covering everything you need from the outset, including local environments, CI/CD deployments, hosting environments, database migrations, and more. With Yona, you can bypass the lengthy setup process and focus on what truly matters—building your application. Simply clone the repository, and you’ll be ready to start developing your business logic right away.

## 🚀 Technologies

- **Frontend**: React + TypeScript
- **Backend**: Go + Gin
- **Database**: PostgreSQL, Prisma, Neon
- **Infrastructure**: Docker, Docker Compose, Terraform
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
make install-prerequisites
```

### 3. Start Nix

To enter the Nix shell and set up your development environment, run the following command in your terminal:

```bash
nix-shell

```

This command creates an isolated environment based on your `shell.nix` file, allowing you to access all specified tools and dependencies. You can exit the Nix shell by typing `exit` or pressing `Ctrl + D`. For more information about Nix, visit the [Nix website](https://nixos.org/nix/).

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
