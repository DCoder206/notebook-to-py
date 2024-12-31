# Python Notebook to script converter
A tool to export Jupyter notebooks as Python files

## Description

This Go-based tool converts Jupyter Notebook files (`.ipynb`) into Python script files (`.py`). The script reads a `.ipynb` notebook file, extracts the code and markdown cells, and outputs the code in a `.py` file format. Markdown cells are converted into Python comments, making it easy to preserve documentation alongside the code.

### Features:
- Converts Jupyter notebook code cells into executable Python code.
- Converts markdown cells into Python comments for easier documentation.
- Works cross-platform (Windows, Linux, macOS).

## Installation

### 1. Downloading the Script

1. Navigate to the [releases page](https://github.com/DCoder206/notebook-to-py/releases/tag/v1.0.0).
2. Download the latest release file: `.tar.gz` or `.exe` (depending on your OS).

### 2. Building the Script by Itself

To build the Go binary from source, follow these steps:

#### Prerequisites
- Go 1.18+ installed on your machine. You can download and install Go from [here](https://go.dev/dl/).

#### Steps to Build

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/DCoder206/notebook-to-py.git
   cd notebook-to-py
   ```

2. Build the Go binary:

   ```bash
   go build -o ipyconverter.exe .\src\converter.go
   ```

3. The compiled binary (`ipyconverter`) will now be available in your current directory.

## Usage

1. Open a terminal or command prompt and navigate to the directory where the `ipyconverter` binary is located.
2. Run the tool:

   ```bash
   ./ipynb_to_py_converter
   ```

3. When prompted, enter the path to the `.ipynb` file you want to convert. Make sure the file has the `.ipynb` extension.

   ```bash
   Enter file path >>> /path/to/your/notebook.ipynb
   ```

4. The tool will create a `.py` file with the same name as the notebook, containing the converted Python code and comments.

   ```bash
   Python file created successfully
   ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
