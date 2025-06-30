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

    **Beispiel:** Schlüssel aus einer unspezifischen Map extrahieren
    --- 

    In Go sind Maps assoziative Datensammlungen (map[key]value). Der Key muss immer comparable sein (damit Go Hashing & Vergleiche machen kann). Oft will man alle Keys unabhängig vom Wert-Typ haben. <br> <br>
    
    **Eingabebeispiel:**
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
    // Mithilfe einfacher Funktionen muss für jede Map 'map[typ1]typ2' eine eigene Funktion mit derselben Struktur ...
 
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
    // Eine Funktion für alle Key/Value-Kombinationen!
    // K: Map-Key
    // V: Map-Value
    func MapKeys[K comparable, V any](m map[K]V) []K {
        keys := make([]K, 0, len(m))
        for k := range m {
            keys = append(keys, k)
        }
        return keys
    }  

    // Klar lesbare, wiederverwendbare Aufrufe:
    ageKeys := MapKeys[string, int](ages)       // -> []string{"Alice", "Bob"}
    todoKeys := MapKeys[int, string](todo)      // -> []int{1, 2, 3}
    
    // -> Vorteil: nur noch eine Codebasis!
    ```
    </details><br>
- 
    <details>
    <summary><strong>Allgemeine Datenstrukturen:</strong><br> Universelle Strukturen mit  beliebige Typen.</summary><br>
    
    Eigene Datenstrukturen wie Stacks, Queues oder Bäume kommen in vielen Programmen vor. Ohne Generics müsste man sie für jeden Elementtyp neu schreiben.

    **Beispiel:** Stack-Datenstruktur für verschiedene Typen
    --- 
    Ein Stack (LIFO) ist eine Datenstruktur, die zwei grundlegende Methoden hat:
    - Push: Element oben drauflegen
    - Pop: oberstes Element entfernen und zurückgeben <br> <br>


    **Ohne Generics:**
    ```go
    // Stack für int-Elemente
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

    // Stack für string-Elemente
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

    // Zwei unterschiedliche Stack-Typen notwendig:

    var intStack IntStack
    intStack.Push(42)
    intStack.Pop()

    var stringStack StringStack
    stringStack.Push("Hallo")
    stringStack.Pop()
    ```

    **Mit Generics:**
    ```go
    // Stack kann hier Werte jeden Typs halten:
    // Allerdings: Innerhalb eines Stacks nur ein Typ
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
    <summary><strong>Identische Methoden:</strong><br> Wrapper, die exakt dieselbe Logik für verschiedene Typen bereitstellen.</summary> 

    Wenn exakt die gleiche Logik für unterschiedliche Typen gebraucht wird (z.B. Suchen in Listen), kann eine generische Funktion helfen, redundanten Code zu vermeiden.


    **Beispiel:** Index eines gesuchten Elements im Slice finden
    --- 

    Wenn man in Go wissen will, an welcher Stelle ein Wert in einem Slice steht, braucht man oft eine Schleife, die mit '==Ä vergleicht. Der Typ im Slice muss daher comparable sein, sonst gibt’s einen Compiler-Fehler.<br> <br>

    **Ohne Generics:**
    ```go
    // Für int-Elemente
    func IndexOfInt(slice []int, target int) (int, error) {
        for i, v := range slice {
            if v == target {
                return i, nil
            }
        }
        return -1, fmt.Errorf("int %v nicht gefunden", target)
    }

    // Für string-Elemente:
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
    // T muss vergleichbar sein, damit wir == verwenden dürfen.
    // Einfache, universelle Funktion.
    func indexOf[T comparable](s []T, e T) (int, error) {
        for i, v := range s {
            if v == e {
                return i, nil
            }
        }
        return -1, errors.New("element not found")
    }

    idx1, err1 := IndexOf([]int{1, 2, 3}, 2)                // idx1 == 1, err1 == nil

    idx2, err2 := IndexOf([]string{"foo", "bar"}, "baz")    // idx2 == -1, err2 == Error("baz nicht gefunden")
    ```
    </details><br>

## ❌ Wann Generics vermieden werden sollten 
- 
    <details>
    <summary><strong>Wenn Overgeneralization droht:</strong><br> Nicht jede Funktion muss generisch sein.</summary><br>

    Oft ist eine Funktion nur für einen bestimmte Datentypen gedacht. Dann bringt eine generische Signatur keinen echten Mehrwert, macht den Code sogar komplizierter und fehleranfälliger.

    **Beispiel:** Print-Methode für beliebige Typen, die fmt.Stringer implementieren
    --- 

    Manche Funktionen sind einfach nur für bestimmte Typen sinnvoll, z.B. Strings in Großbuchstaben umwandeln. <br> <br>

    
    **Mit Generics:**
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

    **Ohne Generics:**
    ```go
    // Das ist viel klarer: Funktioniert nur für Strings, was auch der Sinn ist.
    func ToUpperCase(s string) string {
        return strings.ToUpper(s)
    }
    ```
    </details><br>

- 
    <details>
    <summary><strong>Wenn ein Interface ausreicht:</strong><br> Verwende eine vorhandene Schnittstelle, sobald alle benötigten Methoden schon definiert sind.</summary><br>

    Go bietet [Interface-Types](https://go.dev/tour/methods/9) an. Sie erlauben ebenfalls generischen Code zu schreiben. Falls alles, was mit einem Wert eines Typs getan werden muss, das Aufrufen einer oder mehrerer Methoden auf diesem Wert ist, genügt ein Interface-Typ. Typparameter würden hier nur unnötig Komplexität hinzufügen. 

    **Beispiel:** Print-Methode für beliebige Typen, die fmt.Stringer implementieren
    --- 

    Interfaces in Go beschreiben Verhalten, nicht Typen. Das Interface fmt.Stringer garantiert, dass String() implementiert ist. Damit kann jede Funktion, die einen "druckbaren" Wert haben will, einfach fmt.Stringer nehmen – egal, was für ein Typ dahintersteckt. <br> <br>
    
    **Mit Generics:**
    ```go
    // Generics hier sind unnötig, weil das Interface alles abdeckt.
    func Print[T fmt.Stringer](v T) {
        // v ruft Println() auf
        // -> T muss zwangsläufig fmt.Stringer sein
        fmt.Println(v.String())
    }
    ```

    **Ohne Generics:**
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

    **Beispiel:** Flächenberechnung für verschiedene Geometrie-Typen
    ---

    Verschiedene Geometrie-Typen (wie Circle, Rectangle) haben zwar alle eine Fläche, aber die Berechnung ist unterschiedlich! <br> <br>

    **Mit Generics:**
    ```go
    type Circle struct{R float64}
    type Rectangle struct{W, H float64}

    // Generics verleiten zu unsauberem "type switch" und Panic.

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
    
    **Ohne Generics:**
    ```go
    type Circle struct{R float64}
    type Rectangle struct{W, H float64}

    // Besser: Jeder Typ weiß selbst, wie Fläche berechnet wird!

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

    **Beispiel:** Anzahl der Felder in einem Struct ermitteln
    ---
    
    Reflection ist ein Feature in Go, mit dem man zur Laufzeit Infos über Typen und Werte herausfinden kann (reflect.TypeOf()). <br> <br>
    
    **Mit Generics:**
    ```go
    func GetNumberOfFieldsInStruct[T any](v T) int {
        // Generics ändert hier nichts: reflect.TypeOf bleibt nötig
        return reflect.TypeOf(v).NumField()
    }
    ```

    **Ohne Generics:**
    ```go
    // interface{} ist für Reflection der Standard und reicht vollkommen aus
    func ChangeNumberOfFieldsInStruct(v interface{}) int {
    return reflect.TypeOf(v).NumField()
    ```
    </details><br>

## Weiterführende Ressourcen 
- [When To Use Generics](https://go.dev/blog/when-generics)  
- [Type Parameters Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)  
- [Generics in Go: Tips & Pitfalls](https://medium.com/@letsCodeDevelopers/generics-in-go-use-cases-tips-and-pitfalls-e25ec564c9a5)
- [15 Reasons Why You Should Use Generics in Go](https://medium.com/@jamal.kaksouri/15-reasons-why-you-should-use-generics-in-go-39601c3be6e0)
- [Generics in Go: A Comprehensive Guide with Code Examples](https://expertbeacon.com/generics-in-go-a-comprehensive-guide-with-code-examples/)