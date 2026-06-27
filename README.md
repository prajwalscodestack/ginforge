# GinForge

GinForge is a developer-focused CLI tool for bootstrapping and scaffolding production-ready applications built with the Gin framework.

Instead of manually creating folders, modules, and architectural boilerplate, GinForge generates a structured project layout based on popular backend architecture patterns.

## Features

- Generate new Gin projects
- Layered Architecture support
- Hexagonal Architecture support
- Module scaffolding
- Template-driven code generation
- Architecture-aware project metadata
- Extensible architecture system

## Supported Architectures

### Layered Architecture

```text
internal/
├── handler/
├── service/
├── repository/
├── model/
└── routes/
