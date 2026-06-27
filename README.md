# GinForge

GinForge is an architecture-aware CLI toolkit for building, analyzing, and maintaining production-ready applications built with the Gin framework.

Instead of manually creating project structures, modules, and architectural boilerplate, GinForge provides project scaffolding, module generation, route discovery, and architecture validation for Gin applications.

---

## 🚀 Features

### ⚡ Project Scaffolding

Generate production-ready Gin projects using popular architecture patterns with interactive support.

```bash
ginforge new myapp --architecture layered
ginforge new myapp --architecture hexagonal
ginforge new
```

---

### 🧩 Module Generation

Generate modules with architecture-aware boilerplate inside existing projects.

```bash
ginforge generate module user
ginforge generate module
```

---

### 🔍 Route Discovery

Analyze Gin applications and discover routes automatically using Go AST parsing.

```bash
ginforge routes
```

Supports:

- Route groups
- Nested route groups
- Named handlers
- Method handlers
- Inline handlers

---

### 📤 Route Export

Export discovered routes in multiple formats for debugging and documentation.

```bash
ginforge routes --json
ginforge routes --csv
ginforge routes --md
```

---

### 🧪 Architecture Validation (Doctor)

Validate project structure, architecture rules, and detect issues in your Gin project.

```bash
ginforge doctor
ginforge doctor --strict
```

Checks include:

- Architecture detection
- Structure validation
- Duplicate route detection (with file-level details)
- Layered architecture dependency validation
- Hexagonal architecture dependency validation

---

## 🏗 Supported Architectures

### Layered Architecture

```
internal/
├── handler/
├── service/
├── repository/
├── model/
└── routes/
```

### Hexagonal Architecture

```
internal/
├── domain/
├── application/
└── adapters/
    ├── http/
    ├── persistence/
    └── external/
```

---

## 📦 Installation

### Using Go Install

```bash
go install github.com/prajwalscodestack/ginforge@latest
```

### Verify Installation

```bash
ginforge --help
```

---

## 🛠 Commands

### Create Project

```bash
ginforge new myapp --architecture layered
ginforge new myapp --architecture hexagonal
ginforge new
```

---

### Generate Module

```bash
ginforge generate module user
ginforge generate module
```

---

### Discover Routes

```bash
ginforge routes
```

---

### Export Routes

```bash
ginforge routes --json
ginforge routes --csv
ginforge routes --md
```

---

### Validate Project

```bash
ginforge doctor
ginforge doctor --strict
```

---

## 🧠 Doctor Validation Rules

### Layered Architecture

Allowed:

```
handler -> service
service -> repository
```

Forbidden:

```
service -> handler
repository -> service
repository -> handler
model -> service
model -> repository
```

---

### Hexagonal Architecture

Forbidden:

```
domain -> gin
domain -> database/sql
application -> adapters
```

---

## 📊 Example Workflow

```bash
ginforge new ecommerce-api --architecture layered

cd ecommerce-api

ginforge generate module user

ginforge routes

ginforge doctor --strict
```

---

## 🗺 Roadmap

- OpenAPI generation from routes
- Swagger integration
- CI pipeline integration
- Custom architecture templates
- Plugin system for custom checks
- Doctor auto-fix suggestions

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome.

If you find a bug or have an idea for improvement, feel free to open an issue or submit a pull request.

---

## 📄 License

MIT License