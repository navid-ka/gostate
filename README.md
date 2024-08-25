# State Package

The `state` package provides a simple and flexible way to manage the state within a Go application. It allows you to initialize, switch, and retrieve states using a map-based approach.

## Installation

To install the package, use the following command:

```bash
go get github.com/navid-ka/gostate@latest
```

## Usage

```
import "github.com/navid-ka/gostate"
```

##Creating a New State

To create a new state instance, use the NewState function:

```
s := gostate.NewState()
```

## Initializing States

The InitState method allows you to initialize the state with a map of keys (state names) and their corresponding boolean values.

```
initialStates := map[string]bool{
    "INIT": true,
    "GAME": false,
    "EXIT": false,
}
s.InitState(initialStates)
```

## Switching States

```
s.SwitchState("GAME")
```

## Getting the Current State

```
isGameActive := s.Get("GAME")
if isGameActive {
    // Perform actions related to the GAME state
}
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
