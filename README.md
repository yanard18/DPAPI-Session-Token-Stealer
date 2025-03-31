# DPAPI-Session-Token-Stealer

A Golang-based tool for extracting and decrypting cookies from Chromium-based browsers using DPAPI. 

## **Learned Key Points**

### **1. Where Encrypted Cookies Are Stored**
Encrypted cookies for Chromium-based browsers are typically stored in the following locations:

- **Google Chrome**:  
  - `C:\Users\%USERNAME%\AppData\Local\Google\Chrome\User Data\Default\Network\Cookies`  
  - `C:\Users\%USERNAME%\AppData\Local\Google\Chrome\User Data\Local State` (contains encryption key)  

- **Microsoft Edge**:  
  - `C:\Users\%USERNAME%\AppData\Local\Microsoft\Edge\User Data\Default\Network\Cookies`  
  - `C:\Users\%USERNAME%\AppData\Local\Microsoft\Edge\User Data\Local State`  

- **Brave Browser**:  
  - `C:\Users\%USERNAME%\AppData\Local\BraveSoftware\Brave-Browser\User Data\Default\Network\Cookies`  
  - `C:\Users\%USERNAME%\AppData\Local\BraveSoftware\Brave-Browser\User Data\Local State`  

The `Cookies` file is an SQLite database containing encrypted cookie data, while the `Local State` file holds the encryption key.

### **2. Why Windows DPAPI Encryption Is Used**
Windows **Data Protection API (DPAPI)** is a built-in Windows encryption mechanism that allows applications to store sensitive data securely. Chromium-based browsers use DPAPI to encrypt stored cookies for security reasons:

- **Tied to User Profile**: DPAPI encryption is linked to the Windows user account, making decryption difficult without access to the same profile.  
- **Automatic Key Management**: The Windows OS manages encryption keys internally, preventing the need for manual key storage.  
- **Secure Against External Access**: Even if an attacker gains access to the cookie database, they cannot decrypt the data without retrieving the DPAPI key from the `Local State` file under the same Windows profile.

This project bypasses this security by extracting and using the stored DPAPI key to decrypt the cookies.

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
