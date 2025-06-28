# When to use Generics in Go

Seit 1.18 erlauben [Generics](https://go.dev/blog/intro-generics), Funktionen und Datentypen so zu schreiben, dass sie mit verschiedenen Typen arbeiten, ohne den Code jedes Mal neu zu implementieren.

>  **🔖 Faustregel:**
> Nutze Generics erst dann, sobald du merkst, dass derselbe Code sich ausschließlich durch die Typen unterscheidet.<br>
-> „Code first, generics later“

> **❗Hinweis:**
> Diese Aufzählungen sind als Empfehlungen zu verstehen und keine starren Regeln. Wende sie je nach Projekt und Anforderung flexibel an.

## ✅ Typische Use-Cases
- 
    <details>
    <summary><strong>Container-Utilities:</strong><br> Helfer-Funktionen wie `Map`, `Filter`, `Reduce` oder `MapKeys` für beliebige Slices oder Maps. </summary>
    
    ```go
    // MapKeys gibt alle Keys einer Map zurück.
    func MapKeys[K comparable, V any](m map[K]V) []K {
        keys := make([]K, 0, len(m))
        for k := range m {
            keys = append(keys, k)
        }
        return keys
    }  
    ```
    </details><br>
- 
    <details>
    <summary><strong>Allgemeine Datenstrukturen:</strong><br> Universelle Strukturen (z.B. 'Tree', 'LinkedList') für beliebige Typen implementieren, ohne für jeden Typ eine eigene Version zu schreiben.</summary>
    
    ```go
    // Stack ist ein generischer LIFO-Stapel.
    type Stack[T any] struct {
	    items []T
    }

    func (s *Stack[T]) Push(v T) {
        s.items = append(s.items, v)
    }+

    func (s *Stack[T]) Pop() T {
        l := len(s.items)
        val := s.items[l-1]
        .items = s.items[:l-1]
    return val
    }
    ```
    </details><br>
- 
    <details>
    <summary><strong>Identische Methoden:</strong><br> Identische Methoden: Wrapper, die exakt dieselbe Logik für verschiedene Typen bereitstellen.</summary>
    
    ```go
    // SortFn sortiert ein Slice anhand der übergebenen Less-Funktion.
    func SortFn[T any](s []T, less func(T, T) bool) {
        sort.Slice(s, func(i, j int) bool {
            return less(s[i], s[j])
        })
    }
    ```
    </details><br>

    <details>
    <summary><strong>Error-Handling:</strong><br> Präzisere Fehlermeldungen mit generischen Funktionen</summary>
    
    ```go
    // IndexOf gibt den Index von target in slice zurück.
    func indexOf[T comparable](s []T, e T) (int, error) {
        for i, v := range s {
            if v == e {
                return i, nil
            }
        }
        return -1, errors.New("element not found")
    }
    ```
    - einheitliches Fehlverhalten für beliebige Typen
    </details><br>

## ❌ Wann Generics vermieden werden sollten 
- 
    <details>
    <summary><strong>Wenn Overgeneralization droht:</strong><br> Nicht jede Funktion muss generisch sein.</summary>
    
    ```go
    func PrintAnything[T any](x T) {
	    fmt.Println(x)
    }
    ```
    </details><br>

    <details>
    <summary><strong>Wenn ein Interface ausreicht:</strong><br> Verwende eine vorhandene Schnittstelle, sobald alle benötigten Methoden schon definiert sind.</summary>
    
    ```go
    // empfohlen:
    func ReadSome(r io.Reader) ([]byte, error)

    // nicht empfohlen:
    func ReadSome[T io.Reader](r T) ([]byte, error)
    ```
    </details><br>
- 
    <details>
    <summary><strong>Wenn die Logik sich pro Typ unterscheidet:</strong><br> Setze auf Interfaces oder konkrete Typen, wenn verschiedene Typen unterschiedliche Implementierungen benötigen. </summary>
    
    ```go
    type Circle struct{ R float64 }
    func (c Circle) Area() float64 { return math.Pi * c.R * c.R }

    type Rectangle struct{ W, H float64 }
    func (r Rectangle) Area() float64 { return r.W * r.H }
    ```
    </details><br>
- 
    <details>
    <summary><strong>Wenn Reflection die passendere Wahl ist:</strong><br> Greife auf Reflection zurück, wenn du hochdynamisch mit beliebigen Typen arbeiten musst und generische Constraints nicht ausreichen. </summary>
    
    ```go
    // Inspect zeigt Felder und Werte eines Structs via Reflection.
    func Inspect(v interface{}) {
        val := reflect.ValueOf(v).Elem()
        for i := 0; i < val.NumField(); i++ {
            fmt.Println(val.Type().Field(i).Name, "=", val.Field(i).Interface())
        }
    }
    ```
    </details><br>

## Weiterführende Ressourcen 
- [When To Use Generics](https://go.dev/blog/when-generics)  
- [Type Parameters Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)  
- [Generics in Go: Tips & Pitfalls](https://medium.com/@letsCodeDevelopers/generics-in-go-use-cases-tips-and-pitfalls-e25ec564c9a5)
- [15 Reasons Why You Should Use Generics in Go](https://medium.com/@jamal.kaksouri/15-reasons-why-you-should-use-generics-in-go-39601c3be6e0)
- [Generics in Go: A Comprehensive Guide with Code Examples](https://expertbeacon.com/generics-in-go-a-comprehensive-guide-with-code-examples/)