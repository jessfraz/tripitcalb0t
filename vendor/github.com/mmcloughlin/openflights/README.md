# openflights

[OpenFlights airport and airline data](https://openflights.org/data.html) encapsulated in a Golang package.

This package contains a list `Airports` with the following fields.

```go
// Airport represents an airport.
type Airport struct {
    // Airport name.
    Name string
    // Main city served by airport.
    City string
    // Country or territory where airport is located.
    Country string
    // 3-letter IATA code. Empty if not assigned/unknown.
    IATA string
    // 4-letter ICAO code.
    ICAO string
    // Latitude in decimal degrees.
    Latitude float64
    // Longitude in decimal degrees.
    Longitude float64
    // Altitude in feet.
    Altitude float64
}
```

Also offers `Airlines` with the following fields.

```go
// Airline represents an airline.
type Airline struct {
    // Name of the airline.
    Name string
    // Alias of the airline. For example, All Nippon Airways is commonly known as "ANA".
    Alias string
    // 2-letter IATA code, if available.
    IATA string
    // 3-letter ICAO code, if available.
    ICAO string
    // Airline callsign.
    Callsign string
    // Country or territory where airline is incorporated.
    Country string
}
```
