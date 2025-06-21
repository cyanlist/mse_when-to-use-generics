package main

import "fmt"

// Das Beispiel demonstriert, wie Generics den Boilerplate-Aufwand in Go drastisch reduzieren, insbesondere bei der Extraktion von Map-Schlüsseln, 
// und vergleicht eine generische Implementierung mit spezialisierten Funktionen ohne Generics.

// MapKeys extrahiert alle Schlüssel aus einer beliebigen Map.
//   - Key:   Typ der Map-Schlüssel (muss comparable sein)
//   - Value: Typ der Map-Werte (beliebig)
// Gibt ein Slice mit allen gefundenen Schlüsseln in keiner bestimmten Reihenfolge zurück.
func MapKeys[Key comparable, Value any](inputMap map[Key]Value) []Key {
	// Wir reservieren sofort genug Kapazität für alle Schlüssel,
	// um mehrfache Allokationen während des Appendens zu vermeiden.
	collectedKeys := make([]Key, 0, len(inputMap))
	for key := range inputMap {
		collectedKeys = append(collectedKeys, key)
	}
	return collectedKeys
}

// ExtractStringIntMapKeys demonstriert, wie umständlich es ohne Generics wird.
// Spezialisierte Funktion nur für map[string]int.
func ExtractStringIntMapKeys(stringIntMap map[string]int) []string {
	sliceOfKeys := make([]string, 0, len(stringIntMap))
	for key := range stringIntMap {
		sliceOfKeys = append(sliceOfKeys, key)
	}
	return sliceOfKeys
}

// ExtractIntBoolMapKeys demonstriert eine weitere spezialisierte Variante.
// Spezialisierte Funktion nur für map[int]bool.
// Diese Funktion unterscheidet sich im Körper nur anhand der Datentypen in make(...)
func ExtractIntBoolMapKeys(intBoolMap map[int]bool) []int {
	sliceOfKeys := make([]int, 0, len(intBoolMap))
	for key := range intBoolMap {
		sliceOfKeys = append(sliceOfKeys, key)
	}
	return sliceOfKeys
}

func main() {
	// --- Beispiel 1: map[string]int ---
	mapOfFruitCounts := map[string]int{
		"apple":  5,
		"banana": 3,
		"cherry": 7,
	}

	// Generischer Aufruf: Schlüssel extrahieren
	fruitCountsViaGenericFunction := MapKeys(mapOfFruitCounts)
	fmt.Println("Keys of fruit counts via Generic MapKeys:", fruitCountsViaGenericFunction) // Ergebnis z.B.: [banana cherry apple]

	// Spezialfunktion (ohne Generics) für denselben Einsatz
	fruitCountsViaSpecializedFunction := ExtractStringIntMapKeys(mapOfFruitCounts)
	fmt.Println("Keys of fruit counts via Specialized ExtractStringIntMapKeys:", fruitCountsViaSpecializedFunction) // Ergebnis z.B.: [apple banana cherry]

	// --- Beispiel 2: map[int]bool ---
	mapOfFeatureFlags := map[int]bool{
		10: true,
		20: false,
		30: true,
	}

	featureFlagsViaGenericFunction := MapKeys(mapOfFeatureFlags)
	fmt.Println("Keys of feature flags via Generic MapKeys:", featureFlagsViaGenericFunction) // Ergebnis z.B.: [10 20 30]

	// ...
	// featureFlagsViaSpecializedFunction := ExtractStringIntMapKeys(mapOfFeatureFlags) funktioniert nicht!
	featureFlagsViaSpecializedFunction := ExtractIntBoolMapKeys(mapOfFeatureFlags)
	fmt.Println("Keys of feature flags via Specialized ExtractIntBoolMapKeys", featureFlagsViaSpecializedFunction) // Ergebnis z.B.: [30 10 20]
}