# xenigma
> A modified version of the enigma encryption machine created in Go.

## How to install?

```shell
go get github.com/sudo-sturbia/xenigma/cmd/xenigma
```

## How to use?

`xenigma` can be used as a command line tool or exported for usage as a package.
For documentation of the package check [godoc](https://godoc.org/github.com/sudo-sturbia/xenigma/pkg/machine).

For command line tool, the help message below is available

```shell
xenigma -help
```

```
Description
    xenigma is a modified version of the enigma encryption machine.

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

### Usage example

```shell
xenigma -gen-w 50 Hello, world! # Generate a 50-rotor machine, write generated config to
                                # ~/.config/xenigma.json, and use generated machine to 
                                # encrypt "Hello, world!".
```
```
onjjk, gqkdx!
```

## What's different?

All components of `xenigma` are customizable through `JSON`, allowing for more
complicated encryption.

The default machine configuration should exist at **~/.config/xenigma.json**.

A configuration file typically looks like the following:

```json
{
    "pathways": [
        ["j", "h", "s", "e", "y", "z", "r", "k", "p", "m", "x", "i", "w", "b", "v", "f", "d", "c", "a", "t", "l", "o", "n", "g", "u", "q"],
        ["n", "c", "v", "w", "q", "t", "h", "z", "o", "m", "a", "s", "x", "r", "g", "u", "d", "i", "f", "k", "j", "b", "e", "y", "p", "l"],
        ["t", "s", "h", "m", "c", "v", "n", "y", "r", "q", "p", "e", "i", "u", "k", "z", "w", "d", "j", "a", "f", "x", "g", "b", "o", "l"]
    ],
    "reflector": ["q", "y", "x", "n", "o", "r", "t", "w", "v", "p", "u", "z", "s", "d", "e", "j", "a", "f", "m", "g", "k", "i", "h", "c", "b", "l"],
    "plugboard": ["r", "n", "w", "q", "p", "u", "v", "o", "y", "x", "s", "t", "z", "b", "h", "e", "d", "a", "k", "l", "f", "g", "c", "j", "i", "m"],

    "rotorPositions": ["a", "b", "c"],
    "rotorStep": 1,
    "rotorCycle": 26
}
```
Some examples are in test/data directory.

### Rotors
`xenigma` allows for **any number of rotors**.
The number of rotors is the number of pathways' arrays and rotor positions (which should be equal).
*For example* in the above example the number of `pathways` (and `rotorPositions`) is 3 so a machine is created with 3 rotors.

`xenigma` also allows for configuration of rotors' cycle and step sizes.

`step` is the number of positions a rotor jumps when shifting.
*For example* if a rotor, with step size 2, is at position "a", then the rotor's position will jump to "c" when shifted once.

`cycle` size is the number of steps a rotor takes before completing a full cycle.
When a rotor completes a full cycle, the adjacent rotor is shifted.
*For example* in a 3-rotor machine if cycle size is 13 then the second rotor is shifted once every time the first rotor completes 13 steps,
the third rotor operates similarly but depends on second rotor's movement, etc.

In a normal enigma machine default values of `step`, and `cycle` are 1, and 26.

### Connections
Connections, such as `pathways`, `reflector`, and `plugboard` are specified through arrays of size 26
where elements' indices represents a character's position in the alphabet.
*For example* if element at index 0 of plugboard array is "c", then "a" (character at 0) is connected to "c".

Both reflector and plugboard arrays should be symmetric,
meaning that if "a" is connected to "b", "b" must also be connected to "a";
Otherwise connections are considered incorrect.

A configuration help message similar to the above is available from the command line.

```shell
xenigma -config-h # Generate configuration help message.
```

