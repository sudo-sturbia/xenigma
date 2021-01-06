# How to Configure?
`xenigma` reads configurations from ~/.config/xenigma/xenigma.conf, which is a JSON
representation of a machine.

See [test data](test-data) for examples of machines. Below is an explaination of each
component and how to configure it using JSON.

## Generating A Machine
`xenigma` provides two flags, `-generate`, and `-gen-w`, that can be used to generate
a full machine.

`-generate` generates a machine, and encrypts a message using it, the machine itself
is not saved.

`-gen-w` generates a machine, writes it to ~/.config/xenigma/xenigma.conf for later
usage, and uses it to encrypt a message.

## Components
### Rotors
`xenigma` allows a variable number of rotors. The number of rotors is the size of
"rotors" array.

Rotor's fields are: pathways, position, step, and cycle.

#### Pathways
Pathways are represented using a 26 element array where indices represent characters
and array elements represent the character they are connected to. For example if
element at index 0 is "b", then "a" (character 0) is connected to "b".

#### Position
Position is an character representing the current position of the rotor, and must
be reachable from the starting position ("a").

#### Step
Step is the number of positions a rotor shifts when stepping once (the size of rotor's
jump.) For example, if a rotor at position "a", has step 3, then a jump will change
rotor's position to "d". The default step is 1.

#### Cycle
Cycle is the number of rotor steps considered a full cycle, after which the following
rotor is shifted. For example, if a rotor has cycle 13, then it needs to complete 13
steps for the following rotor to step once. The default cycle size is 26.

### Reflector
Reflector consists of a connections array similar to pathways with a condition that
it must be symmetric, meaning that if "a" is connected to "b", then "b" must also
be connected to "a".

### Plugboard
Plugboard, also, consists of a connections array with the same condition as a reflector.

Plugboard's connections are required to have 26 elements. In order to keep a character 
without a connection, connect it to itself.

Run `xenigma -h` for other options.
