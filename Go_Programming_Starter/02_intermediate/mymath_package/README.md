# Module 02: Go Modules and Advanced Organization

Moving beyond the fundamentals, this section explores structural organization, package management, and identifier visibility in Go.

---

## Technical Concepts

### 1. Go Modules (`go.mod`)
A **Module** is a collection of related Go packages that are managed together. It is defined by the `go.mod` file at the root of the project.
- The module name (identity) provides the base path for importing internal packages.
- It enables dependency management and versioning.

### 2. Custom Packages
Packages are used to group related functionality in a dedicated directory.
- For example, the `calculator` package contains specific mathematical logic.
- Importing internal packages is done using the absolute path from the module root: `import "mymath/calculator"`.

### 3. Identifier Visibility (The Capitalization Rule)
Go uses the capitalization of the first letter of an identifier to determine its scope and visibility.
- **Exported Identifiers** (Public): Start with a **Capital Letter** (e.g., `Add`). These are accessible from other packages.
- **Unexported Identifiers** (Private): Start with a **Lowercase Letter** (e.g., `subtract`). These are only accessible within the same package.

---

## JavaScript to Go: Modules and Packages

The way Go organizes code is a major departure from JavaScript's file-based modularity.

### The Scope Shift

| Feature | JavaScript (ES6) | Go (Golang) |
| :--- | :--- | :--- |
| **Unit of Code** | A single **File**. | A **Directory** (Package). |
| **Exporting** | `export const myVar = ...` | Capitalize the first letter (`MyVar`). |
| **Importing** | `import { x } from './file.js'` | `import "module/package"` |
| **Visibility** | Everything is private by default. | Shared within the same directory. |

**The Big Secret**: In JS, if you have `main.js` and `helper.js`, they are two separate worlds. In Go, if `main.go` and `helper.go` are in the same folder, they are the **same world**. They share all variables and functions automatically.

### Dependency Management: `npm` vs `go mod`

- **`go.mod`** is your `package.json`. It defines the project name and its external dependencies.
- **`go.sum`** is your `package-lock.json`. It ensures every developer has the exact same version of every dependency.
- **No `node_modules`**: Go stores dependencies in a central cache on your machine, not inside your project folder. This makes your project folders much smaller and cleaner.

---

## Professional Project Layout

As your project grows, putting everything in the root directory becomes messy. Here is the industry-standard "Standard Go Layout":

```text
my_project/
├── go.mod            (The project identity/package.json)
├── cmd/              (Entry points / Main files)
│   └── app/
│       └── main.go   (The "package main" file)
├── pkg/              (Public library code / Shared logic)
│   └── auth/
│       └── login.go  (The "package auth" file)
├── internal/         (Private code forbidden to other modules)
│   └── db/
│       └── connect.go
└── README.md
```

- **`cmd/`**: Use this for the code that starts your app. 
- **`pkg/`**: Use this for logic that you want others (even other projects) to import and use.
- **`internal/`**: This is a special Go keyword. Any code inside an `internal` folder can **only** be imported by files within the same parent folder. It's the ultimate "Private" flag for your architecture.

---

## Compilation and Execution

1.  **Singular Package Names**: Package names should be short, concise, and typically singular (e.g., `calculator` instead of `calculators`).
2.  **Explicit Scope Management**: Use identifier visibility to keep your internal implementation private while exposing a clean public API.
3.  **Module Initialization**: Always initialize your project with `go mod init <module_path>` to ensure proper dependency tracking.

---

## Compilation and Execution
To compile the module and execute the resulting binary:
```bash
# Compile and output a named binary
go build -o mymath_bin main.go

# Execute the binary
./mymath_bin
```
