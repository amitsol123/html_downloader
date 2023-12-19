# HTML Downloader

HTML Downloader is a small Go program designed to download HTML files from a list of URLs concurrently using Go's concurrency features.

## Table of Contents
- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [TearDown](#tear-down)

## Overview

This program utilizes Go's concurrency features (goroutines, channels, and wait groups) to download HTML files from a provided list of URLs concurrently. It limits the number of concurrent downloads using a specified number of workers.

## Installation

1. **Clone the repository:**
    ```bash
    git clone https://github.com/amitsol123/html_downloader.git
    cd html_downloader/
    ```
   
2. **Build the project:**
    ```bash
    go build -o html_downloader
    ```

## Usage

**Run the program:**
```bash
# Run with default number of workers (5)
./html_downloader

# Run with a specific number of workers (e.g., 8)
./html_downloader 8
```


   Optionally, adjust the number of concurrent workers by modifying the `maxWorkers` constant in the `main.go` file.

## Configuration

- `maxWorkers`: The number of concurrent workers can be controlled by passing it as a command-line argument while running the program.

## Tear Down

After downloading the HTML files, you might want to clean up the downloaded files or directories. To do so:

1. **Remove Downloaded Files:**
   If you want to remove the downloaded HTML files, execute the following command:
    ```bash
    rm -rf downloaded_html/
    ```

2. **Remove Executable:**
   If you no longer need the executable, remove it by running:
    ```bash
    rm html_downloader  # Replace 'html_downloader' with your executable name
    ```

Make sure to use these commands carefully, as they will permanently delete the specified files or directories. Adjust these instructions according to any specific cleanup or tear down actions required by your program.