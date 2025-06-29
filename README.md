# When to use Generics in Go

Seit 1.18 erlauben [Generics](https://go.dev/blog/intro-generics), Funktionen und Datentypen so zu schreiben, dass sie mit verschiedenen Typen arbeiten, ohne den Code jedes Mal neu zu implementieren.

>  **🔖 Faustregel:** <br>
> Nutze Generics erst dann, sobald du merkst, dass derselbe Code sich ausschließlich durch die Typen unterscheidet.<br>
-> „Code first, generics later“

> **❗Hinweis:** <br>
> Die nachfolgenden Aufzählungen sind Empfehlungen, keine starren Regeln. Wende sie je nach Projekt und Anforderung flexibel an.

## ✅ Typische Use-Cases
- 
    <details>
    <summary><strong>Container-Utilities:</strong><br> Helfer-Funktionen für beliebige Container-Strukturen. </summary><br>

    Container-Utilities sind kleine Helfer, die häufig wiederkehrende Operationen auf Slices, Maps oder Channels abdecken – z.B. Filtern, Extrahieren, Umwandeln, ...

    **Beispiel:**
    ```go
    ages := map[string]int{
        "Alice": 31,
        "Bob":   29,
    }

    todo := map[int]string{
        1: "Einkaufen",
        2: "Gassi gehen",
        3: "Lernen",
    }
    ```

    **Ohne Generics:**
    ```go
    func MapKeysStringInt(m map[string]int) []string {
        keys := make([]string, 0, len(m))
        for k := range m {
            keys = append(keys, k)
        }
        return keys
    }

    func MapKeysIntString(m map[int]string) []int {
        keys := make([]int, 0, len(m))
        for k := range m {
            keys = append(keys, k)
        }
        return keys
    }

    ageKeys := MapKeysStringInt(ages)       // -> []string{"Alice", "Bob"}
    todoKeys := MapKeysIntString(todo)      // -> []int{1, 2, 3}
    ```

    **Mit Generics:**
    ```go
    func MapKeys[K comparable, V any](m map[K]V) []K {
        keys := make([]K, 0, len(m))
        for k := range m {
            keys = append(keys, k)
        }
        return keys
    }  

    ageKeys := MapKeys[string, int](ages)       // -> []string{"Alice", "Bob"}
    todoKeys := MapKeys[int, string](todo)      // -> []int{1, 2, 3}
    ```
    </details><br>
- 
    <details>
    <summary><strong>Allgemeine Datenstrukturen:</strong><br> Universelle Strukturen mit  beliebige Typen.</summary><br>
    
    Eigene Datenstrukturen wie Stacks, Queues oder Bäume kommen in vielen Programmen vor. Ohne Generics müsste man sie für jeden Elementtyp neu schreiben.

    **Beispiel:** <br>

    **Ohne Generics:**
    ```go
    // Spezifisch für int
    type IntStack struct { items []int }

    func (s *IntStack) Push(v int) { 
        s.items = append(s.items, v) 
    }

    func (s *IntStack) Pop() int {
        n := len(s.items)
        v := s.items[n-1]
        s.items = s.items[:n-1] 
        return v
    }

    // Spezifisch für string
    type StringStack struct { items []string }

    func (s *StringStack) Push(v string) { 
        s.items = append(s.items, v) 
    }

    func (s *StringStack) Pop() string {
        n := len(s.items)
        v := s.items[n-1]
        s.items = s.items[:n-1] 
        return v
    }

    var intStack IntStack
    intStack.Push(42)
    intStack.Pop()

    var stringStack StringStack
    stringStack.Push("Hallo")
    stringStack.Pop()
    ```

    **Mit Generics:**
    ```go
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
    
    var gi Stack[int]
    gi.Push(42)
    gi.Pop()

    var gs Stack[string]
    gs.Push("Hello")
    gs.Pop()
    ```
    </details><br>
- 
    <details>
    <summary><strong>Identische Methoden:</strong><br> Wrapper, die exakt dieselbe Logik für verschiedene Typen bereitstellen.</summary><br>

    **Beispiel:** <br>
    **Ohne Generics:**
    ```go
    func IndexOfInt(slice []int, target int) (int, error) {
        for i, v := range slice {
            if v == target {
                return i, nil
            }
        }
        return -1, fmt.Errorf("int %v nicht gefunden", target)
    }

    func IndexOfString(slice []string, target string) (int, error) {
        for i, v := range slice {
            if v == target {
                return i, nil
            }
        }
        return -1, fmt.Errorf("string %q nicht gefunden", target)
    }
    ```

    **Mit Generics:**
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

    idx1, err1 := IndexOf([]int{1, 2, 3}, 2)
    // idx1 == 1, err1 == nil

    idx2, err2 := IndexOf([]string{"foo", "bar"}, "baz")
    // idx2 == -1, err2 == Error("baz nicht gefunden")
    ```
    </details><br>

## ❌ Wann Generics vermieden werden sollten 
- 
    <details>
    <summary><strong>Wenn Overgeneralization droht:</strong><br> Nicht jede Funktion muss generisch sein.</summary><br>

    Oft ist eine Funktion nur für einen bestimmte Datentypen gedacht. Dann bringt eine generische Signatur keinen echten Mehrwert, macht den Code sogar komplizierter und fehleranfälliger. 
    
    ```go
    // Eine generische Variante bringt hier nur Komplexität:

    func ToUpperCase[T any](s T) T {
        // Runtime Casting auf string 
        // –> erzeugt Panic, wenn s kein string ist
        upper := strings.ToUpper(s.(string)) 
        // Rück-Cast auf T 
        // –> zusätziger Overhead ohne Mehrwert
        return any(upper).(T)
    }
    ```

    ```go
    func ToUpperCase(s string) string {
        return strings.ToUpper(s)
    }
    ```
    </details><br>

- 
    <details>
    <summary><strong>Wenn ein Interface ausreicht:</strong><br> Verwende eine vorhandene Schnittstelle, sobald alle benötigten Methoden schon definiert sind.</summary><br>

    Go bietet [Interface-Types](https://go.dev/tour/methods/9) an. Sie erlauben ebenfalls generischen Code zu schreiben. Falls alles, was mit einem Wert eines Typs getan werden muss, das Aufrufen einer oder mehrerer Methoden auf diesem Wert ist, genügt ein Interface-Typ. Typparameter würden hier nur unnötig Komplexität hinzufügen. 
    
    ```go
    func Print[T fmt.Stringer](v T) {
        // v ruft Println() auf
        // -> T muss zwangsläufig fmt.Stringer sein
        fmt.Println(v.String())
    }
    ```

    ```go
    // Signatur ist einfacher zu lesen: 
    // fmt.Stringer übernimmt die Rolle eines "Typs"

    func Print(v fmt.Stringer) {
        fmt.Println(v.String())
    }
    ```

    </details><br>
- 
    <details>
    <summary><strong>Wenn die Logik sich pro Typ unterscheidet:</strong><br> Setze auf Interfaces oder konkrete Typen, wenn verschiedene Typen unterschiedliche Implementierungen benötigen. </summary><br>

    Wenn jede Typ-Variante ihre ganz eigene Funktionslogik erfordert, führt ein generischer Ansatz unweigerlich zu Type Switches, Type Casts oder Reflection und bricht die Typsicherheit.

    ```go
    type Circle struct{R float64}
    type Rectangle struct{W, H float64}

    func CalculateArea[T any](s T) float64 {
        switch v := any(s).(type) {
        case Circle:
            return math.Pi * v.R * v.R
        case Rectangle:
            return v.W * v.H
        default:
            panic("unsupported type")
        }
    }
    ```
    
    ```go
    type Circle struct{R float64}
    type Rectangle struct{W, H float64}

    // Einzigartige Implementierung für Circle
    func (c Circle) CalculateArea() float64 { 
        return math.Pi * c.R * c.R 
    }

    // Einzigartige Implementierung für Rectangle
    func (r Rectangle) CalculateArea() float64 { 
        return r.W * r.H 
    }
    ```
    </details><br>
- 
    <details>
    <summary><strong>Wenn Reflection die passendere Wahl ist:</strong><br> Greife auf Reflection zurück, wenn du hochdynamisch mit beliebigen Typen arbeiten musst und generische Constraints nicht ausreichen. </summary><br>

    Go bietet [Reflection](https://go.dev/blog/laws-of-reflection) an. Selbst mit Generics bleibt Reflection unvermeidbar, wenn das Programm seine Typen und Werte während der Laufzeit manipuliert.
    
    ```go
    func GetNumberOfFieldsInStruct[T any](v T) int {
        // Generics ändert hier nichts: reflect.TypeOf bleibt nötig
        return reflect.TypeOf(v).NumField()
    }
    ```
    ```go
    func ChangeNumberOfFieldsInStruct(v interface{}) int {
    // interface{} ist klarer Input für Reflection
    return reflect.TypeOf(v).NumField()
    ```
    </details><br>

## Weiterführende Ressourcen 
- [When To Use Generics](https://go.dev/blog/when-generics)  
- [Type Parameters Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)  
- [Generics in Go: Tips & Pitfalls](https://medium.com/@letsCodeDevelopers/generics-in-go-use-cases-tips-and-pitfalls-e25ec564c9a5)
- [15 Reasons Why You Should Use Generics in Go](https://medium.com/@jamal.kaksouri/15-reasons-why-you-should-use-generics-in-go-39601c3be6e0)
- [Generics in Go: A Comprehensive Guide with Code Examples](https://expertbeacon.com/generics-in-go-a-comprehensive-guide-with-code-examples/)