# Python Notebook to script converter
A tool to export Jupyter notebooks as Python files

## Description

This tool converts Jupyter Notebook files (`.ipynb`) into Python script files (`.py`). The script reads a `.ipynb` notebook file, extracts the code and markdown cells, and outputs the code in a `.py` file format. Markdown cells are converted into Python comments, making it easy to preserve documentation alongside the code.

### Features:
- Converts Jupyter notebook code cells into executable Python code.
- Converts markdown cells into Python comments for easier documentation.
- Works cross-platform (Windows, Linux, macOS).

## Installation

### 1. Downloading the Script

You have two options to get the tool: either by downloading the precompiled `.exe` file or by building it from source.

#### Option 1: Download Precompiled Executable

1. Navigate to the [releases page](https://github.com/DCoder206/notebook-to-py/releases/tag/v1.0.0).
2. Download the appropriate file for your operating system:
   - For **Windows**: Download the `.exe` file.
   - For **macOS/Linux**: Download the `.tar.gz` or `.zip` file.
   
3. Once downloaded, **extract** the `.tar.gz` or `.zip` file (if applicable), or simply locate the `.exe` file for Windows.

#### Option 2: Building the script from source code

If you prefer to build the Go binary from source, follow these steps:

##### Prerequisites
- Go 1.18+ installed on your machine. You can download and install Go from [here](https://go.dev/dl/).

##### Steps to Build

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/DCoder206/notebook-to-py.git
   cd notebook-to-py
   ```

2. Build the Go binary:

   ```bash
   go build -o ipyconverter.exe .\src\converter.go
   ```

3. The compiled binary (`ipyconverter.exe` for Windows, or `ipyconverter` for macOS/Linux) will now be available in your current directory.

## Usage

1. Open a terminal or command prompt and navigate to the directory where the `ipyconverter` binary is located.

2. Run the tool by specifying one or more `.ipynb` files to convert. You can provide the files either by using the `-file` flag or as positional arguments.

   **Using the `-file` flag**:
   ```bash
   ./ipynb_to_py_converter -file /path/to/your/notebook1.ipynb -file /path/to/your/notebook2.ipynb

3. The tool will create a `.py` file with the same name as the notebook, containing the converted Python code and comments.

   ```bash
   Python file notebook1.py created
   ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
