# projinit

**projinit** is a CLI tool for initializing Git projects with essential configuration files. It automates the setup of a new project with LICENSE, README.md, and .gitignore files, providing a quick start for developers.

## Features

- Initializes Git projects with default files.
- Supports multiple platforms: Linux, Windows, and macOS.
- Creates packages for Debian (.deb) and RPM-based systems (.rpm).

## Installation

1. **Build the binaries:**

   ```bash
   make build-all
   ```

2. **Package the binaries:**

   - For Debian (.deb):

     ```bash
     make package-deb
     ```

   - For RPM-based systems (.rpm):

     ```bash
     make package-rpm
     ```

   - To package for all supported architectures:

     ```bash
     make package-all
     ```

## Usage

After building, you can find the binaries in the `build/bin` directory. The `.deb` and `.rpm` packages will be located in the respective `build/deb` and `build/rpm` directories.

## Commands

- `make build-linux`: Build binaries for Linux.
- `make build-windows`: Build binaries for Windows.
- `make build-macos`: Build binaries for macOS.
- `make clean`: Clean build artifacts and package directories.
- `make all`: Build all binaries, clean, and package for Debian and RPM.

## Contributing

Feel free to open issues or submit pull requests to contribute to the project. For any questions or feedback, please contact Mohammad Reza Fadaei at [mohrezfadaei@gmail.com](mailto:mohrezfadaei@gmail.com).
