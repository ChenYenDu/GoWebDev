# JSON (JavaScript Object Notation)

Json is a text format for the serialization of structured data.

It is derived from the object literals of JavaScript.

Json can represent four __primitive types__ (strings, numbers, booleans, and null) and two __structured types__ (objects and arrays)

## Primitive JSON

Here are four small JSON texts containing only values:

```json
"Hello world!"
42
true
null
```

## Object JSON

An object structure is represented as a pair of curly brackets surrounding zero or more name/value pairs (or members).

An object is an unordered collection of zero or more name/value pairs

A __name__ is a string

A __value__ is a string, number, boolean, null, object, or array.

Declare properties using __name:value__ pairings separated by commas

Enclose names in curly braces

There is no trailing comma

This is a JSON object:

```json
{
    "Image": {
        "Width":  800,
        "Height": 600,
        "Title":  "View from 15th Floor",
        "Thumbnail": {
            "Url":    "http://www.example.com/image/481989943",
            "Height": 125,
            "Width":  100
        },
        "Animated" : false,
        "IDs": [116, 943, 234, 38793]
    }
}
```

## Array JSON

An array structure is represented as square brackets surrounding zero or more values (or elements).

Elements are separated by commas.

A __value__ must be an:

- object
- array
- number
- string
- three literal names
  - true
  - false
  - null

This is a JSON array containing two objects:

```json
[
        {
           "precision": "zip",
           "Latitude":  37.7668,
           "Longitude": -122.3959,
           "Address":   "",
           "City":      "SAN FRANCISCO",
           "State":     "CA",
           "Zip":       "94107",
           "Country":   "US"
        },
        {
           "precision": "zip",
           "Latitude":  37.371991,
           "Longitude": -122.026020,
           "Address":   "",
           "City":      "SUNNYVALE",
           "State":     "CA",
           "Zip":       "94085",
           "Country":   "US"
        }
    ]
```

## Number

The representation of numbers is similar to that used in most programming languages.  A number is represented in base 10 using decimal digits.  It contains an integer component that may be prefixed with an optional minus sign, which may be followed by a fraction part and/or an exponent part.  Leading zeros are not allowed. A fraction part is a decimal point followed by one or more digits.

## Strings

The representation of strings is similar to conventions used in the C family of programming languages.  A string begins and ends with quotation marks.

[The Internet Engineering Task Force (IETF)](https://tools.ietf.org/html/rfc7159)
