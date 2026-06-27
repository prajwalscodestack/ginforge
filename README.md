# GinForge

GinForge is a developer-focused CLI tool for bootstrapping, scaffolding, and analyzing production-ready applications built with the Gin framework.

Instead of manually creating folders, modules, and architectural boilerplate, GinForge generates structured project layouts based on popular backend architecture patterns and provides developer tooling for route discovery and project analysis.

---

## Features

### Project Generation

* Generate new Gin projects
* Layered Architecture support
* Hexagonal Architecture support
* Architecture-aware project metadata

### Module Generation

* Generate modules within existing projects
* Architecture-specific module scaffolding
* Automatic architecture detection
* Module path resolution from `go.mod`

### Route Scanner

* Discover Gin routes across the project
* Support for route groups
* Support for nested route groups
* Extract handler names
* Detect inline route handlers
* Display source file information

### Export Formats

* Table output
* JSON output
* CSV output
* Markdown output

### Developer Experience

* Template-driven code generation
* Extensible architecture system
* Consistent project structure
* CLI-first workflow

---

## Installation

```bash
go install github.com/prajwalscodestack/ginforge@latest
```

---

## Create a New Project

### Layered Architecture

```bash
ginforge new bookstore -a layered
```

### Hexagonal Architecture

```bash
ginforge new bookstore -a hexagonal
```

---

## Supported Architectures

### Layered Architecture

```text
internal/
├── handler/
├── service/
├── repository/
├── model/
└── routes/
```

### Hexagonal Architecture

```text
internal/
├── domain/
├── application/
│   ├── ports/
│   └── services/
└── adapters/
    ├── http/
    └── persistence/
```

---

## Generate a Module

```bash
ginforge generate module user
```

GinForge automatically detects the project architecture and generates the appropriate module structure.

---

## Route Scanner

List all routes in the current project:

```bash
ginforge routes
```

Example output:

```text
METHOD   PATH                             HANDLER         FILE
GET      /health                          HealthHandler  internal/routes/health.go
GET      /api/v1/users                    GetUsers       internal/routes/user.go
POST     /api/v1/users                    CreateUser     internal/routes/user.go
```

---

## Export Routes as JSON

```bash
ginforge routes --json
```

Example:

```json
[
  {
    "method": "GET",
    "path": "/users",
    "handler": "GetUsers"
  }
]
```

---

## Export Routes as CSV

```bash
ginforge routes --csv
```

---

## Export Routes as Markdown

```bash
ginforge routes --md
```

---

## Roadmap

* Architecture validation (`ginforge doctor`)
* OpenAPI generation
* Resource generation
* Plugin system
* Custom architecture templates

---

## Why GinForge?

GinForge aims to go beyond project generation by providing architecture-aware tooling for Gin applications.

The goal is to help developers:

* Start faster
* Maintain consistent project structure
* Discover and document routes
* Enforce architectural conventions
* Scale Gin applications more confidently
