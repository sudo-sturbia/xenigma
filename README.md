# xenigma
A modified version of the enigma machine created in Go.

## What's different?
`xenigma` adds another layer of complexity to the original enigma machine.

- Allows for **any number of rotors**.
- All of machine's components are **fully configurable**.
- **Handles most user errors**. Configurations are verified throughout every step.
- All used components can be **randomly generated**.
- Many different configuration options are provided,
such as rotor's step, cycle and more.

## How to Install?

```shell
go get github.com/sudo-sturbia/xenigma/cmd/xenigma
```

## How to Use?

`xenigma` can be used as a command line tool or exported for usage as a package.
For documentation of the package check [godoc](https://godoc.org/github.com/sudo-sturbia/xenigma/pkg/machine).

For command line tool, the help message below is available

```shell
xenigma -help
```

```
Usage
    xenigma [options] <message>

Options
    -h                           Print a short help message.

    -help                        Print a detailed help message.

    -config-h                    Print a help message specifying how to
                                 configure a machine.

    -verify <path>               Verify the correctness of the configuration
                                 at the given path.

    -generate <numberofrotors>   Generate a machine with specified number
                                 of rotors and use it for encryption.

    -gen-w <numberofrotors>      Generate a machine with specified number
                                 of rotors, use it for encryption, and write
                                 generated configs to ~/.config/xenigma.json

    -correct <numberofrotors>    Load ~/.config/xenigma.json, generate a new
                                 machine if configs are incorrect.

    -load <path>                 Load and use config at given path instead
                                 of ~/.config/xenigma.json

    -read <path>                 Read and encrypt contents of file at given
                                 path. If both -read is invoked and a message
                                 is given as argument, both are encrypted and
                                 and printed seperated by a new line.

    -write <path>                Write encrypted message to file at given path.

    -update                      Save updated config to ~/.config/xenigma.json
                                 before exiting. Updated config is config at
                                 ~/.config/xenigma.json after rotor shifting.

    -default-rotors              Use default values for rotor-related fields.
                                 Default values are "a"'s for rotor positions,
                                 1 for step size, and 26 for cycle size.

xenigma is licensed under MIT license.
```

### Usage Example

```shell
xenigma -gen-w 50 Hello, world! # Generate a 50-rotor machine, write generated config to
                                # ~/.config/xenigma.json, and use generated machine to
                                # encrypt "Hello, world!".
```
```shell
onjjk, gqkdx! # "Hello, world!" encrypted.
```

## How to Configure?

`xenigma` allows for configuration of all machine's componenets through `JSON`.
Configurations file should be located at **~/.config/xenigma.json**

A configuration file typically looks like the following

```json
{
    "rotors": [
        {
            "pathways": ["j", "h", "s", "e", "y", "z", "r", "k", "p", "m", "x", "i", "w", "b", "v", "f", "d", "c", "a", "t", "l", "o", "n", "g", "u", "q"],
            "position": "a",
            "step": 1,
            "cycle": 26
        },
        {
            "pathways": ["n", "c", "v", "w", "q", "t", "h", "z", "o", "m", "a", "s", "x", "r", "g", "u", "d", "i", "f", "k", "j", "b", "e", "y", "p", "l"],
            "position": "b",
            "step": 1,
            "cycle": 26
        },
        {
            "pathways": ["t", "s", "h", "m", "c", "v", "n", "y", "r", "q", "p", "e", "i", "u", "k", "z", "w", "d", "j", "a", "f", "x", "g", "b", "o", "l"],
            "position": "c",
            "step": 1,
            "cycle": 26
        }
    ],

    "reflector": {
        "connections": ["q", "y", "x", "n", "o", "r", "t", "w", "v", "p", "u", "z", "s", "d", "e", "j", "a", "f", "m", "g", "k", "i", "h", "c", "b", "l"]
    },

    "plugboard": {
        "connections": ["r", "n", "w", "q", "p", "u", "v", "o", "y", "x", "s", "t", "z", "b", "h", "e", "d", "a", "k", "l", "f", "g", "c", "j", "i", "m"]
    }
}
```

### Rotors

```json
"rotors": [
    {
        "pathways": ["j", "h", "s", "e", "y", "z", "r", "k", "p", "m", "x", "i", "w", "b", "v", "f", "d", "c", "a", "t", "l", "o", "n", "g", "u", "q"],
        "position": "a",
        "step": 1,
        "cycle": 26
    }
]
```

As said before `xenigma` allows for **any number of rotors**.
The number of rotors is the size of *"rotors"* array in **~/.config/xenigma.json**

Rotor's fields are *pathways*, *position*, *step*, and *cycle*.

#### Pathways
Pathways are the electric connections between characters.
Pathways are represented using a 26 element array where indices represent
characters and array elements represent the character they are connected to.

*For example*, if element at index 0 is "b", then "a" (character 0) is connected
to "b".

#### Position
Position is an integer which represents the current position of the rotor.
The given position must be reachable from the starting position *("a")*.

#### Step
Step is the number of positions a rotor shifts when stepping once (the size of
rotor's jump.)

*For example*, if a rotor at position *"a"*, with *step = 3*, steps once,
then rotor's position changes to *"d"*. The default step size is 1.

#### Cycle
Cycle is the number of rotor steps considered a full cycle, after which the
following rotor steps (is shifted.)

*For example*, if a rotor has a *cycle = 13*, then the rotor needs to complete
13 steps in order for the following rotor to step once. The default cycle size is 26.

To avoid position collisions "step \* cycle" must divide 26. Given step-cycle
combinations that don't satisfy that relation are considered wrong.

### Reflector

```json
"reflector": {
    "connections": ["q", "y", "x", "n", "o", "r", "t", "w", "v", "p", "u", "z", "s", "d", "e", "j", "a", "f", "m", "g", "k", "i", "h", "c", "b", "l"]
}
```

Reflector consists of a connections array similar to pathways with a condition
that it must be *symmetric*, meaning that if *"a"* is connected to *"b"*, then
*"b"* must also be connected to *"a"*.

### Plugboard

```json
"plugboard": {
    "connections": ["r", "n", "w", "q", "p", "u", "v", "o", "y", "x", "s", "t", "z", "b", "h", "e", "d", "a", "k", "l", "f", "g", "c", "j", "i", "m"]
}
```

Plugboard, also, consists of a connections array exactly the same as a reflector.

Plugboard's connections are required to have 26 elements,
so characters not connected to anything should be connected to themselves
(in order to not be transformed).

