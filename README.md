# Barcodez
A GUI tool that generates barcodes in multiple formats and outputs them as PDF.

Homepage:

![homepage](./examples/homepage.png)

Output:

![output](./examples/output.png)

## Motivation
I once used an application that generated barcodes, but some of the barcodes did not turn out right due to special characters used or had character limitations. That led me to build barcodez which allowed any characters to be used.

## Quick Start

### Pre-requisites
Runs on linux/amd64.\
You should have a working Go environment, if not please see [this page](https://go.dev/doc/install) first.

### Clone this repository

To clone this repo:

```bash
git clone https://github.com/AhGr3y/barcode-label-pdf.git
```

### Usage

Just run the following command to open the GUI:

```bash
go run .
```

## Running Using a Windows Executable (.exe)

1. Generate the executable `.exe`.
    - This will generate a `.zip` file which contains the `.exe` file in the `./fyne-cross/dist/` directory.
    - You can use any name other than 'com.example.barcode-pdf.generator' as long as it's a unique name throughout your system.
```bash
fyne-cross windows -app-id com.example.barcode-pdf-generator
```
2. Create a folder anywhere in your Windows system, then copy the `/output` and `/templates` directory into the folder.
    - The application needs the files inside these directories to run.
3. Run the application!




