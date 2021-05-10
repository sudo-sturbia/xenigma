# How to Configure?
`xenigma` reads configurations from ~/.config/xenigma/xenigma.conf, which is a JSON
representation of a machine.

See [test data](test-data) for examples of machines. Below is an explaination of each
component and how to specify it using JSON.

## Generating A Machine
`xenigma` provides two flags, `-generate`, and `-gen-w`, that can be used to generate
a full machine.

- `-generate` generates a machine, and encrypts a message using it, the machine itself
is not saved, so the message can't be decrypted or retrieved.

- `-gen-w` generates a machine, writes it to ~/.config/xenigma/xenigma.conf for later
usage, and uses it to encrypt a message. The machine is written before usage, so it
can be used for decryption.

## Components
### Rotors
`xenigma` allows a variable number of rotors. The number of rotors is the size of
"rotors" array.

Rotor's fields are: pathways, position, step, and cycle.

#### Pathways
Pathways are the electric connections between characters. They are represented
using a map-like 26 element array where an index and a character represent a
map pair. Keys are translated into their position in the english alphabet. For
example, if pathways[0]="c", then a is mapped to c. Arrays are chosen over maps
for pathways because ordering matters.

#### Position
Position is an integer representing the current position of the rotor, and must
be reachable from the starting position ("a").

#### Step
Step is the number of positions a rotor jumps when moving one step forward.
For example, if a rotor with position="a" and step="3" jumps once, the position
will change to "d". The default step is 1.

#### Cycle
Cycle is the number of steps needed to complete a full cycle, after which the
following rotor is shifted. For example, if a rotor with cycle=13, then it
needs to complete 13 steps for the next rotor to move one step. The default
cycle is 26.

### Reflector
Reflector is connections map, which must contain all characters in the english
alphabet, and must be symmetric. Symmetry means that if "a" is connected to "b",
then "b" must also be connected to "a".

### Plugboard
Plugboard is also a connections map similar to reflector. To keep a character
unconnected/unplugged, connect it to itself.

Run `xenigma -h` for other options.
