# Module 00: Go Programming Starter

This section is dedicated to mastering the Go programming language from the ground up. Before diving into complex backend engineering, it's essential to understand Go's unique features, syntax, and organizational patterns.

## Curriculum Overview

### [01 Basics](./01_basics)
- **Hello Go**: Setting up your first executable `package main`.
- **Variables and Constant Types**: Exploring Go's strict typing system (int, float64, string, bool).
- **Zero Values**: Understanding Go's default initialization (no more `undefined`!).
- **JS-to-Go Mapping**: Direct technical comparisons for JavaScript/ES6 developers.

### [02 Intermediate](./02_intermediate/mymath_package)
- **Go Modules**: Initializing and managing your project with `go.mod`.
- **Custom Packages**: Organizing logic into specialized directories.
- **Identifier Visibility**: Mastering the "Capitalization Rule" (Exported vs. Unexported).
- **Professional Layout**: Industry-standard project structures for scalable engineering.

---

## Technical Standards

1.  **Strict Typing**: Every variable must have a type, and Go will not perform "magic" (implicit) conversions. 
2.  **Package-Level Scope**: Files within the same directory share logic automatically—no complex `export` required for internal helpers.
3.  **Explicit Errors**: Go encourages a "Fail Early" approach by making errors values you must check.

---
*Ready to master the language? Start with [01 Basics](./01_basics)*
