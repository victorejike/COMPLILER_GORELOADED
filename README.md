# Text Processing Utility (Go)

## Overview

This project is a Go-based text transformation utility that processes and modifies strings using a set of custom rules and inline directives. It is designed to parse human-readable text and apply transformations such as capitalization, case conversion, article correction, quote formatting, and numeric base conversion.

The core logic is implemented in the `helper` package, which exposes multiple functions that can be composed to build a full text-processing pipeline.

---

## Features

### 1. Case Transformation

* **Upper**: Converts the previous word(s) to uppercase using `(up)` or `(up, n)`
* **TransformText**: Converts the previous word(s) to lowercase using `(low)` or `(low, n)`
* **ToCap**: Capitalizes the previous word(s) using `(cap)` or `(cap, n)`

**Examples:**

```
hello world (up)        → hello WORLD
this is nice (up, 2)    → this IS NICE
HELLO WORLD (low)       → HELLO world
go is fun (cap)         → go is Fun
```

---

### 2. Article Correction (`a` → `an`)

* **FixAToAn**: Automatically replaces `"a"` with `"an"` when followed by a vowel sound.
* Handles exceptions such as:

  * Words that start with vowels but sound like consonants (`university`, `user`)
  * Words that start with silent consonants (`hour`, `honest`)

**Example:**

```
a apple     → an apple
a university → a university
an hour     → an hour
```

---

### 3. Quote Formatting

* **FixQuotes**: Cleans spacing around single quotes.
* Ensures:

  * Proper spacing before opening quotes
  * No extra spaces inside quotes

**Example:**

```
hello' world ' → hello 'world'
```

---

### 4. Number Conversion

* **ConvertNumbers**:

  * Converts hexadecimal numbers followed by `(hex)` into decimal

**Example:**

```
1A (hex) → 26
```

* **BinTodecimal**:

  * Converts binary numbers followed by `(bin)` into decimal

**Example:**

```
1010 (bin) → 10
```

---

### 5. Inline Directive System

The system uses inline markers to control transformations:

* `(up)` / `(up, n)` → uppercase
* `(low)` / `(low, n)` → lowercase
* `(cap)` / `(cap, n)` → capitalize
* `(hex)` → convert hexadecimal
* `(bin)` → convert binary

These markers are removed after processing.

---

## Package Structure

```
COMPLILER_GORELOADED/
│── helper
   |¬¬¬ helper.go   # Core implementation
```
# Text Processing Utility (Go)

## Overview

This project is a Go-based text transformation utility that processes and modifies strings using a set of custom rules and inline directives. It is designed to parse human-readable text and apply transformations such as capitalization, case conversion, article correction, quote formatting, and numeric base conversion.

The core logic is implemented in the `helper` package, which exposes multiple functions that can be composed to build a full text-processing pipeline.

---

## Features

### 1. Case Transformation

* **Upper**: Converts the previous word(s) to uppercase using `(up)` or `(up, n)`
* **TransformText**: Converts the previous word(s) to lowercase using `(low)` or `(low, n)`
* **ToCap**: Capitalizes the previous word(s) using `(cap)` or `(cap, n)`

**Examples:**

```
hello world (up)        → hello WORLD
this is nice (up, 2)    → this IS NICE
HELLO WORLD (low)       → HELLO world
go is fun (cap)         → go is Fun
```

---

### 2. Article Correction (`a` → `an`)

* **FixAToAn**: Automatically replaces `"a"` with `"an"` when followed by a vowel sound.
* Handles exceptions such as:

  * Words that start with vowels but sound like consonants (`university`, `user`)
  * Words that start with silent consonants (`hour`, `honest`)

**Example:**

```
a apple     → an apple
a university → a university
an hour     → an hour
```

---

### 3. Quote Formatting

* **FixQuotes**: Cleans spacing around single quotes.
* Ensures:

  * Proper spacing before opening quotes
  * No extra spaces inside quotes

**Example:**

```
hello' world ' → hello 'world'
```

---

### 4. Number Conversion

* **ConvertNumbers**:

  * Converts hexadecimal numbers followed by `(hex)` into decimal

**Example:**

```
1A (hex) → 26
```

* **BinTodecimal**:

  * Converts binary numbers followed by `(bin)` into decimal

**Example:**

```
1010 (bin) → 10
```

---

### 5. Inline Directive System

The system uses inline markers to control transformations:

* `(up)` / `(up, n)` → uppercase
* `(low)` / `(low, n)` → lowercase
* `(cap)` / `(cap, n)` → capitalize
* `(hex)` → convert hexadecimal
* `(bin)` → convert binary

These markers are removed after processing.

---

## Package Structure

```
helper/
│── helper.go   # Core implementation
```

---

## Usage

Import the package:

```go
import "yourmodule/helper"
```

Example:

```go
text := "hello world (up)"
result := helper.Upper(text)
fmt.Println(result) // hello WORLD
```

You can chain multiple transformations:

```go
text := "a hour 1010 (bin) hello (cap)"
text = helper.FixAToAn(text)
text = helper.BinTodecimal(text)
text = helper.ToCap(text)

fmt.Println(text) // an hour 10 Hello
```

---

## Design Notes

* Uses `strings.Fields` for tokenization
* Relies on index-based backward transformations
* Regular expressions are used for pattern-based numeric conversion
* Designed to be modular and composable

---

## Limitations

* Assumes well-formed input (e.g., `(up, 2)` must be correctly spaced)
* Minimal error handling for malformed directives
* Does not support punctuation-aware tokenization beyond spaces

---

## Future Improvements

* Support punctuation-aware parsing
* Add more linguistic rules (e.g., pluralization)
* Improve directive parsing robustness
* CLI interface for file input/output

---

## Contributors software developers

* Multiple contributors contributed to different functions:

  * Case transformations    | by emeyanga / cjohn / AYUBA D'SWAY
  * Quote handling          | by bonyeke 
  * Numeric conversions     | by abraham zion(zabraham) / christain patrick (cpatrick)
  * punctuation handling    | by victor ejike(vejike) /
  * main function           | by joel obiabo(jobiabo) / roland elaigwu(relaigwu)
  * fix A to An             | by  Mekomcu (Melchizedek Ekondu)

---

## Summary

This project demonstrates practical string manipulation in Go, combining parsing, transformation logic, and modular design. It is especially useful for learning how to process structured text with inline commands.

---

## Usage

Import the package:

```go
import "yourmodule/helper"
```

Example:

```go
text := "hello world (up)"
result := helper.Upper(text)
fmt.Println(result) // hello WORLD
```

You can chain multiple transformations:

```go
text := "a hour 1010 (bin) hello (cap)"
text = helper.FixAToAn(text)
text = helper.BinTodecimal(text)
text = helper.ToCap(text)

fmt.Println(text) // an hour 10 Hello
```

---

## Design Notes

* Uses `strings.Fields` for tokenization
* Relies on index-based backward transformations
* Regular expressions are used for pattern-based numeric conversion
* Designed to be modular and composable

---

## Limitations

* Assumes well-formed input (e.g., `(up, 2)` must be correctly spaced)
* Minimal error handling for malformed directives
* Does not support punctuation-aware tokenization beyond spaces

---

## Future Improvements

* Support punctuation-aware parsing
* Add more linguistic rules (e.g., pluralization)
* Improve directive parsing robustness
* CLI interface for file input/output

---

## Contributors
* Multiple contributors contributed to different functions:

  * Case transformations    | by emeyanga / cjohn / AYUBA D'SWAY
  * Quote handling          | by bonyeke 
  * Numeric conversions     | by abraham zion(zabraham) / christain patrick (cpatrick)
  * punctuation handling    | by victor ejike(vejike) /
  * main function           | by joel obiabo(jobiabo) / roland elaigwu(relaigwu)
  * fix A to An             | by  Mekomcu (Melchizedek Ekondu)

---

## Summary

This project demonstrates practical string manipulation in Go, combining parsing, transformation logic, and modular design. It is especially useful for learning how to process structured text with inline commands.
