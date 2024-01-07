# Codebase Structure of templ-ui

The templ-ui project is structured to facilitate ease of navigation and development. Below is an outline of the project's directory structure and a description of each top-level directory and its contents.

## Overview

```plaintext
.
├── app                     # Application code for the documentation website
│   ├── handlers            # Handlers for routing and rendering views
│   │   ├── index.go
│   │   ├── meta.go
│   │   └── renderer.go
│   ├── src                 # Source files for client-side scripting or styles
│   │   └── css
│   │       └── input.css
│   ├── static              # Static files served directly to the client
│   │   └── css
│   │       └── styles.css
│   ├── views               # Views and layouts for the UI
│   │   ├── layouts
│   │   │   ├── base.templ
│   │   │   └── base_templ.go
│   │   └── pages
│   │       ├── home.templ
│   │       └── home_templ.go
│   └── websocket.js        # Client-side scripts for websocket communication
├── cmd                     # Command-line interface related files
│   └── server.go
├── components              # UI components directory
│   ├── accordion           # Accordion component
│   │   ├── accordion.templ
│   │   └── accordion_templ.go
│   └── button              # Button component
│       ├── button.templ
│       └── button_templ.go
├── CODE_OF_CONDUCT.md      # Code of conduct for contributors
├── CONTRIBUTING.md         # Guidelines for contributing to the project
├── go.mod                  # Go module file
├── go.sum                  # Go checksum file
├── LICENSE                 # The license for the project
├── package-lock.json       # NPM package lock file
├── package.json            # NPM package configuration
├── README.md               # The top-level README for developers using this project
├── tailwind.config.js      # Configuration for Tailwind CSS
└── tmp                     # Temporary files
    └── main                # Main executable
```

## Directory Descriptions

### `app/`

The `app` directory contains the code for the documentation website for templ-ui. This includes handlers for routing, views for the user interface, and static assets like stylesheets.

### `cmd/`

The `cmd` directory contains entry points for the command-line interface, which is used to run the server for the documentation site.

### `components/`

The `components` directory is where the UI components for templ-ui are developed. Each component has its own subdirectory to keep related files organized and avoid namespace collisions.

### `CODE_OF_CONDUCT.md`

This file contains the code of conduct for contributors, outlining the expectations for behavior and participation.

### `CONTRIBUTING.md`

The `CONTRIBUTING.md` file provides guidelines for how developers can contribute to the project, including how to submit bug reports, feature requests, and pull requests.

### `go.mod` and `go.sum`

These files are used by the Go programming language to manage project dependencies.

### `LICENSE`

The `LICENSE` file contains the licensing information for templ-ui, explaining how others may use the project's code.

### `package-lock.json` and `package.json`

These files are used by NPM to manage JavaScript dependencies for the project.

### `README.md`

The `README.md` file is the first point of reference for developers interested in using or contributing to templ-ui. It provides an overview of the project and how to get started.

### `tailwind.config.js`

The `tailwind.config.js` file contains configuration settings for Tailwind CSS used in the project.

### `tmp/`

The `tmp` directory is used to store temporary files, including the main executable for the documentation site when running locally.

## Conclusion

Understanding the structure of templ-ui is crucial for both navigating the project and contributing effectively. If you have any questions or need further clarification, please refer to the `CONTRIBUTING.md` file or contact the project maintainers.
