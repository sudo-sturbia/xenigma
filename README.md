# xenigma
A modified version of the enigma machine.

## What's different?
`xenigma` adds another layer of complexity to the original enigma machine.

- Allows an unlimited number of rotors,
- Fully configurable, and generatable,
- Handles most user errors,
- Many different configuration options are provided.

## How to Install?

```shell
go get github.com/sudo-sturbia/xenigma/v4/cmd/xenigma
```

## How to Use?

`xenigma` can be used as a command line tool or exported for usage as a package.
For package documentation see [pkg.go.dev](https://pkg.go.dev/github.com/sudo-sturbia/xenigma/v4).

For command line tool, the help message below is available
```
xenigma is a modified version of the enigma encryption machine.

Usage
  xenigma [options] <message>

Options
  -help                Print this help message.

  -config-h            Print a help guide explaining xenigma.conf.

  -verify <path>       Verify the correctness of the the machine at
                       path.

  -generate <count>    Generate a machine with given number of rotors
                       use it for encryption.

  -gen-w <count>       Generate a machine with given number of rotors,
                       write it to ~/.config/xenigma/xenigma.conf, and
                       use it for encryption.

  -gen-backup <count>  Generate a new machine, if
                       ~/.config/xenigma/xenigma.conf is invalid.

  -load <path>         Use machine at path instead of
                       ~/.config/xenigma/xenigma.conf.

  -read <path>         Read message from path. If both -read is invoked and
                       a message is given as argument, both are encrypted
                       and printed seperated by a new line.

  -update              Save updated machine to ~/.config/xenigma/xenigma.conf
                       before exiting.

  -defaults            Use default values for rotor-related fields.
                       Default values are "a"'s for rotor positions,
                       1 for step size, and 26 for cycle size.

See github.com/sudo-sturbia/xenigma for source code.
```

### Usage Example

```shell
xenigma -gen-w 50 Hello, world! # Generate a 50-rotor machine, write generated config to
                                # ~/.config/xenigma/xenigma.conf, and use generated machine to
                                # encrypt "Hello, world!".
```
```shell
onjjk, gqkdx!                   # "Hello, world!" encrypted.
```

## How to Configure?
See [How To Configure?](config.md).
