// Package space defines Age and Planet
package space

// Planet is a string
type Planet string

// Age defines an age for a given planet
func Age(seconds float64, planet Planet) float64 {
	earthSeconds := 31557600.0

	var conversion = map[Planet]float64{
		"Mercury": 0.2408467 * earthSeconds,
		"Earth":   1.0 * earthSeconds,
		"Venus":   0.61519726 * earthSeconds,
		"Mars":    1.8808158 * earthSeconds,
		"Jupiter": 11.862615 * earthSeconds,
		"Saturn":  29.447498 * earthSeconds,
		"Uranus":  84.016846 * earthSeconds,
		"Neptune": 164.79132 * earthSeconds,
	}

	return seconds / conversion[planet]
}
