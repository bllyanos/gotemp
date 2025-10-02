# Gotemp

An opinionated Go template loading and rendering library that provides a clean way to organize and render HTML templates with layouts, partials, and content pages.

## Overview

Gotemp is a lightweight template engine built on top of Go's standard `html/template` package. **This library is opinionated** - it enforces a predefined directory structure for your templates, providing consistency and reducing configuration overhead.

Features include:

- **Template inheritance** with layouts and partials
- **Component-based design** with reusable partials
- **Content separation** with dedicated page templates
- **Type-safe rendering** with data injection
- **Error handling** for missing templates
- **Opinionated structure** - no configuration needed for template organization

## Installation

```bash
go get github.com/bllyanos/gotemp
```

## Quick Start

Gotemp requires a specific directory structure for your templates:

```
templates/
├── root.html           # Base HTML structure (required)
├── partials/           # Reusable components (optional)
│   └── _header.html
├── layouts/            # Page layouts (required)
│   └── app.html
└── pages/              # Page content (required)
    └── home/
        └── index.html
```

```go
package main

import (
    "os"
    "github.com/bllyanos/gotemp"
)

func main() {
    // Initialize the template engine with your templates directory
    // Note: The directory must contain the required structure (root.html, layouts/, pages/)
    g, err := gotemp.New("templates")
    if err != nil {
        panic(err)
    }

    // Render a page with layout and data
    data := map[string]string{
        "Title": "Welcome",
        "Content": "Hello, World!",
    }

    err = g.RenderPage(os.Stdout, "main_layout", "home/index.html", data)
    if err != nil {
        panic(err)
    }
}
```

## API Documentation

### `New(basePath string) (*Gotemp, error)`

Creates a new Gotemp instance with templates loaded from the specified base directory.

**Parameters:**
- `basePath`: Path to the directory containing your template files

**Returns:**
- `*Gotemp`: Template engine instance
- `error`: Error if template loading fails

### `RenderPage(w io.Writer, layout, page string, data any) error`

Renders a page template within a specified layout.

**Parameters:**
- `w`: Writer to output the rendered content
- `layout`: Name of the layout template to use
- `page`: Path to the page template relative to the pages directory
- `data`: Data to pass to the template (can be nil)

**Returns:**
- `error`: Error if rendering fails

## Directory Structure

**Gotemp is opinionated and enforces a specific directory structure** - you must organize your templates as follows:

```
templates/
├── root.html           # Base HTML structure
├── partials/           # Reusable components
│   └── _header.html
├── layouts/            # Page layouts
│   ├── app.html
│   └── auth.html
└── pages/              # Page content
    ├── home/
    │   └── index.html
    └── auth/
        └── sign_in.html
```

### Required Template Components

#### Root Template (`root.html`) - **Required**
Must be present in the base directory. Defines the base HTML structure with start and end blocks:

```html
{{ define "__start" }}
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{ .Title }}</title>
  </head>
  <body>
{{ end }}

{{ define "__end" }}
  </body>
</html>
{{ end }}
```

#### Partials (`partials/*.html`) - **Optional**
Reusable components that can be included in layouts. Files in this directory should be named with underscore prefix (e.g., `_header.html`):

```html
{{ define "_header" }}
<div class="navbar bg-base-200">
  <h1>Your APP!!</h1>
</div>
{{ end }}
```

#### Layouts (`layouts/*.html`) - **Required**
Templates that combine partials and define page structure. Each layout file defines a template name used in `RenderPage()`:

```html
{{ define "app_layout" }}
{{ template "__start" . }}

{{ template "_header" . }}

<div class="container mx-auto">
  {{ block "content" . }}
  {{ end }}
</div>

{{ template "__end" . }}
{{ end }}
```

#### Pages (`pages/*/*.html`) - **Required**
Content templates that define the main content blocks. **Must be organized in subdirectories** within the pages folder. The page path in `RenderPage()` should match the relative path from the pages directory:

```html
{{ define "content" }}
<h1>Homepage</h1>
<p>{{ .Content }}</p>
{{ end }}
```

## Important: Opinionated Design

**Gotemp follows convention over configuration** - the library strictly enforces the directory structure and template organization. This approach provides:

- **Consistency** across projects using Gotemp
- **Zero configuration** - just point to your templates directory
- **Predictable structure** - developers know exactly where to find templates
- **Reduced cognitive load** - no need to decide on template organization

**If you need flexible template organization, this library may not be suitable for your project.**

## Usage Examples

### Basic Page Rendering

```go
g, err := gotemp.New("templates")
if err != nil {
    return err
}

// Render home page without data
err = g.RenderPage(os.Stdout, "app_layout", "home/index.html", nil)
```

### Page with Data

```go
data := map[string]interface{}{
    "Title": "User Profile",
    "Name": "John Doe",
    "Email": "john@example.com",
}

err = g.RenderPage(os.Stdout, "app_layout", "user/profile.html", data)
```

### Different Layouts

```go
// Use app layout for main pages
err = g.RenderPage(os.Stdout, "app_layout", "home/index.html", data)

// Use auth layout for authentication pages
err = g.RenderPage(os.Stdout, "auth_layout", "auth/login.html", data)
```

### Error Handling

```go
g, err := gotemp.New("templates")
if err != nil {
    log.Fatalf("Failed to initialize templates: %v", err)
}

err = g.RenderPage(&buf, "app_layout", "home/index.html", data)
if err != nil {
    if strings.Contains(err.Error(), "page template not found") {
        log.Printf("Page template not found: %v", err)
    } else {
        log.Printf("Rendering error: %v", err)
    }
    return err
}
```

## Testing

Run the test suite:

```bash
go test
```

Run tests with verbose output:

```bash
go test -v
```

The test suite covers:
- Template initialization
- Page rendering with and without data
- Error handling for missing templates
- Multiple page rendering
- Output to different writers

## License

This project is open source. See the LICENSE file for details.

## LLM Quick Setup Guide

<details>
<summary>📖 Complete LLM Setup and Usage Guide</summary>

For a comprehensive guide on installing, setting up, and using this package with detailed examples and best practices, see [llm.txt](./llm.txt).

### Quick Start for LLM Users

1. **Install the package:**
   ```bash
   go get github.com/bllyanos/gotemp
   ```

2. **Create required template structure:**
   ```bash
   mkdir -p templates/{partials,layouts,pages/home}
   ```

3. **Create basic templates:**
   ```bash
   # root.html (required)
   cat > templates/root.html << 'EOF'
   {{ define "__start" }}
   <!DOCTYPE html>
   <html lang="en">
     <head>
       <meta charset="UTF-8">
       <title>{{ .Title }}</title>
     </head>
     <body>
   {{ end }}
   {{ define "__end" }}
     </body>
   </html>
   {{ end }}
   EOF

   # layouts/app.html (required)
   cat > templates/layouts/app.html << 'EOF'
   {{ define "app_layout" }}
   {{ template "__start" . }}
   <div class="container">
     {{ block "content" . }}{{ end }}
   </div>
   {{ template "__end" . }}
   {{ end }}
   EOF

   # pages/home/index.html (required)
   cat > templates/pages/home/index.html << 'EOF'
   {{ define "content" }}
   <h1>Welcome!</h1>
   <p>{{ .Message }}</p>
   {{ end }}
   EOF
   ```

4. **Basic usage:**
   ```go
   package main

   import (
       "os"
       "github.com/bllyanos/gotemp"
   )

   func main() {
       g, err := gotemp.New("templates")
       if err != nil {
           panic(err)
       }

       data := map[string]string{
           "Title":   "My Site",
           "Message": "Hello from Gotemp!",
       }

       err = g.RenderPage(os.Stdout, "app_layout", "home/index.html", data)
       if err != nil {
           panic(err)
       }
   }
   ```

### Advanced Features

- **Multiple layouts**: Create different layouts for auth, admin, etc.
- **Partials**: Reusable components like headers, footers
- **Complex data**: Use Go structs for type-safe data injection
- **Error handling**: Comprehensive error management
- **Web server integration**: HTTP server examples

For detailed explanations, troubleshooting, and advanced patterns, see the complete [llm.txt guide](./llm.txt).

</details>

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.