# How to Configure?

`xenigma` allows for configuration of all machine's componenets using 
file **~/.config/xenigma.conf**, xenigma.conf is parsed as a normal `JSON` file.

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

## Rotors

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

As said before, `xenigma` allows for **any number of rotors**.
The number of rotors is the size of *"rotors"* array in **~/.config/xenigma.conf**

Rotor's fields are *pathways*, *position*, *step*, and *cycle*.

### Pathways
Pathways are the electric connections between characters.
Pathways are represented using a 26 element array where indices represent
characters and array elements represent the character they are connected to.

*For example*, if element at index 0 is "b", then "a" (character 0) is connected
to "b".

### Position
Position is an integer which represents the current position of the rotor.
The given position must be reachable from the starting position *("a")*.

### Step
Step is the number of positions a rotor shifts when stepping once (the size of
rotor's jump.)

*For example*, if a rotor at position *"a"*, with *step = 3*, steps once,
then rotor's position changes to *"d"*. The default step size is 1.

### Cycle
Cycle is the number of rotor steps considered a full cycle, after which the
following rotor steps (is shifted.)

*For example*, if a rotor has a *cycle = 13*, then the rotor needs to complete
13 steps in order for the following rotor to step once. The default cycle size is 26.

To avoid position collisions "step \* cycle" must divide 26. Given step-cycle
combinations that don't satisfy that relation are considered wrong.

## Reflector

```json
"reflector": {
    "connections": ["q", "y", "x", "n", "o", "r", "t", "w", "v", "p", "u", "z", "s", "d", "e", "j", "a", "f", "m", "g", "k", "i", "h", "c", "b", "l"]
}
```

Reflector consists of a connections array similar to pathways with a condition
that it must be *symmetric*, meaning that if *"a"* is connected to *"b"*, then
*"b"* must also be connected to *"a"*.

## Plugboard

```json
"plugboard": {
    "connections": ["r", "n", "w", "q", "p", "u", "v", "o", "y", "x", "s", "t", "z", "b", "h", "e", "d", "a", "k", "l", "f", "g", "c", "j", "i", "m"]
}
```

Plugboard, also, consists of a connections array exactly the same as a reflector.

Plugboard's connections are required to have 26 elements,
so characters not connected to anything should be connected to themselves
(in order to not be transformed).

