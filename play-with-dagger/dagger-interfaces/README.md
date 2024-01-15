# Interface

## Requirement

⚠️ You need at least Dagger v0.9.6 (I personnally tried with it and also a version compiled from main)

## Repro

```shell
cd car

dagger call race
```

Output:

```
✘ dagger call race ERROR [1.21s]
┃ Error: response from query: input:1: race.race failed to get function output directory: process "/runtime" did not com
┃ plete successfully: exit code: 2                                                                                      
✘ race ERROR [1.21s]
✘ exec /runtime ERROR [1.19s]
┃ input:1: car.withCar failed to get function output directory: process "/runtime" did not complete successfully: exit c
┃ e: 2                                                                                                                  
✘ withCar(car: "moddata:sha256:8fcd3c7ecfc9b5ffa3c4cc72b094a6e0d79df912acd1695bf84dfa06652eae78:Audi:eyJDb2xvciI6InJlZCIsIk1vZGVsIjoiQTQiLCJTcGVlZCI6MTAwfQ==") ERROR [0.14s]
✘ exec /runtime ERROR [0.14s]
┃ json: error calling MarshalJSON for type main.ICar: input:1: Cannot query field "loadCarICarFromID" on type "Query". D
┃  you mean "loadCarIcarFromID", "loadServiceFromID", "loadCarFromID", "loadContainerFromID", or "loadDirectoryFromID"? 
• Engine: 84734f0d7c6f (version v0.9.6)
⧗ 8.80s ✔ 223 ∅ 8 ✘ 5
```

## Explanation

This project design a `ICar` interface with the following methods:

```go
type ICar interface {
	WithModel(model string) ICar
	WithSpeed(speed int) ICar
	WithColor(color string) ICar
	Drive(ctx context.Context) error
}
```

```shell
├── audi # Car implementation
├── bmw  # Car implementation
├── car  # Car interface
└── race # Global usage
```

The implementation looks good, the generation seems okay but the convertion into ID is failed (see repro).