<p align="center">
  <a href="" rel="noopener">
 <img class="" width=200px height=200px src="https://files.oaiusercontent.com/file-3Htw9zJAcqhirPe6g9UgGWdj?se=2024-10-08T13%3A13%3A34Z&sp=r&sv=2024-08-04&sr=b&rscc=max-age%3D604800%2C%20immutable%2C%20private&rscd=attachment%3B%20filename%3D894318ed-629f-41aa-8a29-7dbcbd0ec537.webp&sig=lTF7sqftCXj3llsKptjjQaq/w91/WqZw1XpuhW7IFdk%3D" alt="Project logo"></a>
</p>

<h3 align="center">Go-Base</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/kylelobo/The-Documentation-Compendium.svg)](https://github.com/Tutuacs/gbase/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/kylelobo/The-Documentation-Compendium.svg)](https://github.com/Tutuacs/gbase/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align=""> Go-Base is a simple project that aims to provide a base for a Go project. It includes a simple web server, a database connection, and a simple API. You can generate new Handlers too, create new routes, and customize the defalt code.
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Deployment](#deployment)
- [Usage](#usage)
- [Built Using](#built_using)
- [Contributing](../CONTRIBUTING.md)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## 🧐 About <a name = "about"></a>

This project is a simple base for a Go project. It includes a simple web server, a database connection, and a simple API. You can generate new Handlers too, create new routes, customize the defalt code and also change the cli.

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them.

```bash
sudo apt-get install git
```
```bash
go -v 1.23.1
```

### Installing

A step by step series of examples that tell you how to get a development env running.

Run cd on the root

```bash
cd ~
```
Clone the repository

```bash
git clone https://github.com/Tutuacs/gbase.git
```
Install the [Go 1.23.1](https://go.dev/dl/)

```bash
linux: https://go.dev/dl/go1.23.1.src.tar.gz
```

After installing all you can add the export path on your .bashrc or .zshrc

```bash
nano ~/.bashrc || nano ~/.zshrc
```

Paste on the end of the file
```bash
export PATH=$HOME/gbase/bin:$PATH
```

## 🎈 Usage <a name="usage"></a>

All done, you can now create new projects with the command

```bash
gbase new [project_name] || gbase new .
```
Inside the path created you can run the *Generate* command

```bash
gbase g h [handler_name] || gbase generate handler [handler_name]
```

## ⛏️ Built Using <a name = "built_using"></a>

- [Cobra-cli](https://github.com/spf13/cobra) - Engine
- [Go](https://go.dev/) - Language

## ✍️ Authors <a name = "authors"></a>

- [@Arthur Silva](https://github.com/Tutuacs) - Idea & Initial work
