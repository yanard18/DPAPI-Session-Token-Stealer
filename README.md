# DPAPI-Session-Token-Stealer

A Golang-based tool for extracting and decrypting cookies from Chromium-based browsers using DPAPI. 

## Features
- **Cookie Extraction**: Retrieves session cookies from Chromium-based browsers (Chrome, Edge, Brave) using DPAPI.
- **Automatic and Manual Modes**: Can automatically extract from default locations or manually specified files.
- **Browser Process Termination**: Optionally kills browser processes before extraction.
- **Flexible Output Formats**: Saves extracted cookies as text or SQLite database.

## Prerequisites
- **Go** (latest version)
- Windows OS (tested on Windows 10/11)

## Installation
### 1. Clone the Repository
```sh
git clone https://github.com/yanard18/DPAPI-Session-Token-Stealer.git
cd DPAPI-Session-Token-Stealer
```

### 2. Install Dependencies
Ensure Go modules are initialized and dependencies are installed:
```sh
go mod tidy
```

### 3. Build the Project
```sh
make build
```
This will generate `bin/cookie.exe`.

## Usage
### Running the Executable
```sh
bin/cookie.exe [OPTIONS]
```

### Command-Line Arguments
- `--auto` : Automatically detects and extracts cookies from Chrome, Edge, and Brave.
- `--state-file <path>` : Specify the `Local State` file path.
- `--cookies-file <path>` : Specify the `Cookies` database file path.
- `--output <file>` : Save extracted cookies to a specified file.
- `--format <text/sql>` : Choose output format (default: text, alternative: SQLite database).
- `--kill-browsers` : Kills browser processes before extraction.

### Example Commands
#### Auto Mode (Extract from Default Locations)
```sh
bin/cookie.exe --auto --format sql --output cookies.sqlite3
```

#### Manual Mode (Specify File Paths)
```sh
bin/cookie.exe --state-file "C:\\Users\\User\\AppData\\Local\\Google\\Chrome\\User Data\\Local State" --cookies-file "C:\\Users\\User\\AppData\\Local\\Google\\Chrome\\User Data\\Default\\Network\\Cookies"
```

#### Kill Browsers and Extract Cookies
```sh
bin/cookie.exe --kill-browsers --auto
```

## Disclaimer
This project is for educational and research purposes only. Use at your own risk.
