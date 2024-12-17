
# md2html - Markdown to HTML Converter (Golang)

`md2html` is a simple command-line tool written in Go that converts Markdown files into HTML. It reads a Markdown file as input, processes it, and generates an equivalent HTML file as output.

---

## Features

- Convert Markdown files (`.md`) into HTML files.
- Command-line interface for easy usage.
- Lightweight and fast, built with Go.
- Uses the `gomarkdown/markdown` library for Markdown processing.

---

## Requirements

- **Go** version 1.18 or later.

---

## Installation

1. Clone the repository:

```bash
git clone https://github.com/Yamaguchi-Katsuya/md2html.git
cd md2html
```

2. Install dependencies:

```bash
go get github.com/gomarkdown/markdown
```

3. Build the executable:

```bash
go build -o md2html
```

This creates an executable named `md2html` in the project directory.

---

## Usage

### Command Syntax:

```bash
./md2html <input_path> <output_path>
```

- `<input_path>`: Path to the input Markdown file.
- `<output_path>`: Path to the output HTML file.

---

### Examples

#### Example 1: Basic Markdown to HTML Conversion

```bash
./md2html input.md output.html
```

- Converts `input.md` to `output.html`.

#### Example 2: Run with `go run` (no build)

```bash
go run md2html.go input.md output.html
```

---

## Example Input and Output

### Input File (`input.md`):

```markdown
# Hello, World!

This is a **Markdown** example.

- Bullet point 1
- Bullet point 2
```

### Output File (`output.html`):

```html
<h1>Hello, World!</h1>
<p>This is a <strong>Markdown</strong> example.</p>
<ul>
    <li>Bullet point 1</li>
    <li>Bullet point 2</li>
</ul>
```

---

## Error Handling

- If the input file does not exist, the program will output:
  ```
  File <input.md> does not exist
  ```
- If the incorrect number of arguments is provided, it will display:
  ```
  Usage: go run md2html.go <input_path> <output_path>
  ```

---

## Dependencies

- **gomarkdown/markdown**: A Go library for converting Markdown to HTML.  
  Install using:
  ```bash
  go get github.com/gomarkdown/markdown
  ```

---
