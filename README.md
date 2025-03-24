![](assets/mamba.png)

# MambaPy

MambaPy is a lightweight and fast CLI tool for managing Python projects. It allows for quick project structure generation, dependency installation, and
execution of essential development commands.

## 📌 Key Features

- 🔧 **Initialize a new project** – automatically creates the project structure.
- 📦 **Dependency management** – integrates with `poetry` for package handling.
- 🛠️ **Built-in development tools** – testing, formatting, and linting.
- 🚀 **Easy setup** – minimal steps to get started.

## 🚀 Installation

To install MambaPy, simply run:

```bash
go install github.com/wokacz/mamba/cmd/mamba@latest
```

## 🎯 Usage

After installation, you can use MambaPy by running:

```sh
mamba init
```

Available commands:

```sh
mamba init         # Creates a new project
mamba install      # Installs dependencies
mamba doctor       # Checks if all dependencies are correctly installed
```

## 🔗 Requirements

- Python 3.13+
- Poetry (optional but recommended)

## 🤝 Contributing

Want to help improve MambaPy? Go ahead! Fork the repository and submit PRs.

## 📜 License

MambaPy is available under the MIT license. Feel free to use, modify, and enhance it! 🚀

