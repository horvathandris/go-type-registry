# go-type-registry

Create a type registry of all exported struct types in your package in Go.

Currently, only supports single file parsing, but planning on implementing multi-file soon.

## Usage

You can either use the CLI or the parser directly.

### CLI

```
go-type-registry [ OPTIONS ]
```

#### Flags:

- `-i, --input` the input .go file, from which the registry is created **(required)**
- `-o, --output` the output .go file, where the registry is created **(required)**

### Parser

```
parser.Start(inputFilename, outputFilename)
```

## Using the type registry

You can create an instance from the registry using MakeInstance

```
instance := typeRegistry.MakeInstance(typeName)
```

