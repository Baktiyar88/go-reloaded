# Go-Reloaded

## Project Description
Go-Reloaded is a powerful text completion, editing, and auto-correction tool written in Go. The program processes text files and applies various modifications including case transformations, number conversions, punctuation formatting, quote correction, and article adjustments.

## Features
- **Number System Conversions**: Converts hexadecimal and binary numbers to decimal
- **Case Transformations**: Supports uppercase, lowercase, and capitalization
- **Punctuation Formatting**: Automatically formats punctuation marks with proper spacing
- **Quote Correction**: Fixes single and double quote positioning
- **Article Correction**: Automatically changes "a" to "an" before vowels and silent 'h'
- **Batch Processing**: Handles multiple modifications in sequence

## Installation
```bash
git clone https://github.com/yourusername/go-reloaded.git
cd go-reloaded
```

## Usage
```bash
go run . <input_file.txt> [output_file.txt]
```

### Arguments:
- `input_file.txt`: Path to the input text file (required)
- `output_file.txt`: Path to the output file (optional, defaults to "result.txt")

### Examples:
```bash
# Basic usage with default output file
go run . sample.txt

# Specify custom output file
go run . sample.txt result.txt
```

## Modification Rules

### Number Conversions
- `(hex)`: Converts the preceding hexadecimal number to decimal
  - Example: `"1E (hex) files"` → `"30 files"`
- `(bin)`: Converts the preceding binary number to decimal
  - Example: `"10 (bin) years"` → `"2 years"`

### Case Transformations
- `(up)`: Converts the preceding word to uppercase
  - Example: `"go (up)!"` → `"GO!"`
- `(low)`: Converts the preceding word to lowercase
  - Example: `"SHOUTING (low)"` → `"shouting"`
- `(cap)`: Capitalizes the preceding word
  - Example: `"bridge (cap)"` → `"Bridge"`

### Multiple Word Transformations
- `(up, n)`, `(low, n)`, `(cap, n)`: Applies transformation to the preceding n words
  - Example: `"exciting (up, 2)"` → `"SO EXCITING"`

### Punctuation Formatting
- Automatically formats punctuation marks (`.`, `,`, `!`, `?`, `:`, `;`) to be close to the preceding word with space after
- Handles punctuation groups like `...` or `!?` without spaces
- Example: `"there ,and then !!"` → `"there, and then!!"`

### Quote Correction
- Single quotes (`'`) are positioned around words without spaces
- Double quotes (`"`) are properly spaced and positioned
- Handles contractions correctly (e.g., `don't`, `we'll`)
- Example: `"' awesome '"` → `"'awesome'"`

### Article Correction
- Changes `a` to `an` before words starting with vowels or silent 'h'
- Example: `"a amazing"` → `"an amazing"`
- Example: `"a hour"` → `"an hour"`

## File Structure
```
go-reloaded/
├── main.go
├── proces/
│   ├── punctuation.go    # Punctuation and quote processing
│   ├── transform.go      # Case transformations
│   ├── numbers.go        # Number system conversions
│   └── articles.go       # Article corrections
└── README.md
```

## Example Usage

### Input (`sample.txt`):
```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
```

### Output (`result.txt`):
```
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

### Number Conversion Example:
```
# Input
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.

# Output
Simply add 66 and 2 and you will see the result is 68.
```

### Article Correction Example:
```
# Input
There is no greater agony than bearing a untold story inside you.

# Output
There is no greater agony than bearing an untold story inside you.
```

## Error Handling
- Validates file extensions (must be `.txt`)
- Prevents overwriting input file
- Handles file read/write errors gracefully
- Maintains original text structure and line breaks

## Requirements
- Go programming language
- Only standard Go packages are used
- Input and output files must have `.txt` extension

## Testing
The project includes comprehensive processing functions that can be unit tested. It's recommended to create test files for validation.

## Technical Details
- Preserves original line structure
- Processes each line independently
- Uses regular expressions for pattern matching
- Implements state machines for quote processing
- Handles edge cases for contractions and punctuation groups

## Author
Baktiyar

---

*This project demonstrates advanced string manipulation, regular expressions, and file I/O operations in Go.*