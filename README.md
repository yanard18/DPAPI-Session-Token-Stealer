# DPAPI Session Token Stealer

A Go-based tool that extracts and decrypts session tokens (cookies) from Chromium-based browsers on Windows using the Data Protection API (DPAPI).

## 1. Features

- **DPAPI Decryption:**  
  Utilizes Windows DPAPI to decrypt the encryption key from the Local State file and then decrypts the session tokens stored in the browser’s SQLite cookie database.

- **SQLite Cookie Parsing:**  
  Reads and parses cookie entries from an SQLite database, extracting details such as host, name, value, path, and metadata.

- **Multiple Output Options:**  
  Supports logging the decrypted cookies in plain text and saving them into an SQLite database for further analysis.

- **Modular Design:**  
  Organized with clear separation of functionality:
  - Cookie extraction
  - Decryption logic
  - Output formatting (text and SQL)

## 2. Prerequisites

- **Go:**  
  Go 1.15+ is required to build the project.

- **SQLite:**  
  The cookies file is an SQLite database; ensure your environment supports SQLite.

- **Windows OS:**  
  This tool leverages Windows DPAPI for decryption and is designed for Windows systems.

## 3. Installation Guide

### 3.1 System Dependencies (Windows)

- Make sure your Windows system supports DPAPI (Data Protection API).  
- No additional system packages are needed.

### 3.2 Cloning and Building

Clone the repository:

```bash
git clone https://github.com/yanard18/DPAPI-Session-Token-Stealer.git
cd DPAPI-Session-Token-Stealer
```

Build the binary:

```bash
go build -o dpapi-token-stealer
```

## 4. Usage

Run the binary with the paths to your Local State file and Cookies SQLite file:

```bash
dpapi-token-stealer -state="path/to/Local State" -cookies="path/to/Cookies"
```

This command will extract and decrypt the session tokens, displaying the details in the console. Use the available output modules to log in text format or save to an SQLite database.

## 5. Output Options

- **Plain Text Logging:**  
  Uses the built-in text output module (see `textout.go`) to print decrypted cookie details on the console.

- **SQLite Database Output:**  
  Saves decrypted cookies to a new SQLite database using the SQL output module (see `sqlout.go`) for further inspection.


## 6. Code Structure

- **cookiemonster.go:**  
  Contains the main function to retrieve and decrypt cookies.

- **cookies.go:**  
  Handles parsing of the cookies database file and mapping of cookie attributes.

- **localstate.go:**  
  Parses the Local State file to extract the DPAPI encrypted key, then strips the DPAPI prefix.

- **sqlout.go:**  
  Provides functionality to save decrypted cookies into an SQLite database.

- **textout.go:**  
  Logs decrypted cookie information in a human-readable text format.

- **internal/decryption:**  
  Implements the DPAPI decryption logic used to decrypt both the encryption key and cookie values.

## 7. Disclaimer

> **Warning:**  
> This tool is intended **for educational and research purposes only**. Unauthorized extraction of session tokens or unauthorized access to data is illegal. Use responsibly and only on systems you have explicit permission to test.

## 8. License

This project is released under the [LICENSE NAME] License.  
*(Replace `[LICENSE NAME]` with the appropriate license if available.)*
