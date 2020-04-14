# Go Starter Kit

**This project is WIP**

![Go](https://img.shields.io/badge/Go-1.14.2-blue.svg)

TODO:
- Add GitHub Actions shields
- Add CircleCI shields

## What is this?

This is starter kit for development of API server in Go.

We can prepare the development environment quickly by using this starter kit.


## Packages

I show list of using packages and why I select it.

If you know better one, please tell me :)
→ [Twitter](https://twitter.com/yyh_gl)


### ▼ Web Application Framework

**null**

I prefer no framework.

#### Routing

[**httprouter**](https://github.com/julienschmidt/httprouter)

It's simple.

### ▼ ORM

[**xsql**](https://github.com/jmoiron/sqlx)

I just want to write raw SQL and map result to go struct.

### ▼ Live reload

[**air**](https://github.com/cosmtrek/air)

it's simple and practical. 

Actually I want to use [realize](https://github.com/oxequa/realize),
but realize repository is unmanaged.

### ▼ CI/CD

- **GitHub Actions**
- **CircleCI**

### ▼ Linter

[**GolangCI-Lint**](https://github.com/golangci/golangci-lint)

GolangCI-Lint realizes centralized management of linter for Go.

### ▼ Development Environment

- [**Docker**](https://www.docker.com/)
- [**Docker Compose**](https://docs.docker.com/compose/)

I use conntainer because:
- To realize the same environment between local and production environment.
- To start development quickly.
- I suppose that system on production environment is worked on container. 
  - For example, [GKE](https://cloud.google.com/kubernetes-engine), [ECS](https://aws.amazon.com/ecs/) and [AKS](https://azure.microsoft.com/en-us/services/kubernetes-service/).

All third-party tools that use in project are installed in docker image.

### ▼ IDE/Editor

`.gitignore` has below rules for IDE and editors.

- **JetBrains series**
- **Emacs**
- **Vim**
- **Visual Studio Code**
