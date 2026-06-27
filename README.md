# GinForge

GinForge is an architecture-aware CLI toolkit for building, analyzing, and maintaining production-ready applications built with the Gin framework.

Instead of manually creating project structures, modules, and architectural boilerplate, GinForge provides project scaffolding, module generation, route discovery, and architecture validation for Gin applications.

---

## Features

### Project Scaffolding

Generate production-ready Gin projects using popular architecture patterns.

```bash
ginforge new myapp --architecture layered
```

```bash
ginforge new myapp --architecture hexagonal
```

### Module Generation

Generate modules with architecture-specific boilerplate.

```bash
ginforge generate module user
```

### Route Discovery

Analyze Gin applications and discover routes automatically using Go AST.

```bash
ginforge routes
```

Supports:

* Route groups
* Nested route groups
* Named handlers
* Method handlers
* Inline handlers

### Route Export

Export discovered routes in multiple formats.

```bash
ginforge routes --json
ginforge routes --csv
ginforge routes --md
```

### Architecture Validation

Validate project structure and architecture rules.

```bash
ginforge doctor
```

Checks include:

* Architecture detection
* Structure validation
* Duplicate route detection
* Layered architecture dependency validation
* Hexagonal architecture dependency validation

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
└── adapters/
    ├── http/
    ├── persistence/
    └── external/
```

---

## Installation

### Using Go Install

```bash
go install github.com/prajwalscodestack/ginforge@latest
```

### Verify Installation

```bash
ginforge --help
```

---

## Commands

### Create Project

```bash
ginforge new myapp --architecture layered
```

```bash
ginforge new myapp --architecture hexagonal
```

### Generate Module

```bash
ginforge generate module user
```

### Discover Routes

```bash
ginforge routes
```

### Export Routes

```bash
ginforge routes --json
```

```bash
ginforge routes --csv
```

```bash
ginforge routes --md
```

### Validate Project

```bash
ginforge doctor
```

---

## Doctor Validation Rules

### Layered Architecture

Allowed:

```text
handler -> service
service -> repository
```

Forbidden:

```text
service -> handler
repository -> service
repository -> handler
model -> service
model -> repository
```

### Hexagonal Architecture

Forbidden:

```text
domain -> gin
domain -> database/sql
application -> adapters
```

---

## Example

### Generate Project

```bash
ginforge new ecommerce-api --architecture layered
```

### Generate Module

```bash
ginforge generate module product
```

### Discover Routes

```bash
ginforge routes
```

### Validate Architecture

```bash
ginforge doctor
```

---

## Roadmap

* OpenAPI generation
* Swagger generation
* CI integration
* Custom architecture templates
* Plugin system
* Resource generators

---

## Contributing

Contributions, issues, and feature requests are welcome.

If you find a bug or have an idea for improvement, feel free to open an issue or submit a pull request.

---

## License

MIT License
