# When to use Generics in Go

Seit GO 1.18 erlauben [Generics](https://go.dev/blog/intro-generics), Funktionen und Datentypen so zu schreiben, dass sie mit verschiedenen Typen arbeiten, ohne den Code jedes Mal neu zu implementieren.

>  **🔖 Faustregel:** <br>
> Nutze Generics erst dann, sobald du merkst, dass derselbe Code sich ausschließlich durch die Typen unterscheidet.<br>
-> „Code first, generics later“

> **❗Hinweis:** <br>
> Die nachfolgenden Aufzählungen sind Empfehlungen, keine starren Regeln. Wende sie je nach Projekt und Anforderung flexibel an.

<br>

## ✅ Typische Use-Cases

- 
    <details>
    <summary><strong>Container-Utilities:</strong><br> Helfer-Funktionen für beliebige Container-Strukturen. <br><a>[Mehr anzeigen]</a></summary><br>

    Container-Utilities sind praktische Funktionen, die typische Operationen auf Datenstrukturen wie [Slices](https://go.dev/blog/slices-intro), [Maps](https://go.dev/blog/maps) oder [Channels](https://golangdocs.com/channels-in-golang) vereinfachen. Dazu gehören zum Beispiel das Extrahieren von Keys, Filtern von Elementen oder das Umwandeln von Werten. <br>

    ### **Beispiel: Alle Keys aus einer Maps extrahieren**
    ---  

    In Go gibt es keine eingebaute Funktion, um alle Schlüssel (Keys) einer Map direkt als Slice zu bekommen. Maps sind assoziative Datenstrukturen, die Schlüssel (Keys) auf Werte (Values) abbilden. <br> <br>

    **Ohne Generics:**
    ```go
    // Ziel: Aus Maps beliebiger Typen sollen die Keys extrahiert werden
    // Ohne Generics: Für JEDE Map-Kombination muss eine eigene Funktion geschrieben werden
    // Das führt zu Copy&Paste und doppeltem Code, sobald mehrere Map-Typen verwendet werden

    personAgeMap := map[string]int{
        "Alice": 31,
        "Bob":   29,
    }

    todoMap := map[int]string{
        1: "Einkaufen",
        2: "Gassi gehen",
        3: "Lernen",
    }

    // Funktion für map[string]int
    func GetKeysFromStringIntMap(stringIntMap map[string]int) []string {
        // Erstellt ein Slice für alle Keys (vom Typ string)
        keys := make([]string, 0, len(stringIntMap))
        for key := range stringIntMap {
            keys = append(keys, key)
        }
        return keys
    }

    // Funktion für map[int]string
    func GetKeysFromIntStringMap(intStringMap map[int]string) []int {
        // Gleiches Prinzip, aber diesmal sind die Keys vom Typ int
        keys := make([]int, 0, len(intStringMap))
        for key := range intStringMap {
            keys = append(keys, key)
        }
        return keys
    }

    // Anwendung:
    personNames := GetKeysFromStringIntMap(personAgeMap)    // -> []string{"Alice", "Bob"}
    todoIDs := GetKeysFromIntStringMap(todoMap)             // -> []int{1, 2, 3}

    // Nachteil: 
    // - Exakt dieselbe Logik wird mehrfach geschrieben
    // - Lediglich der Typ des Keys unterscheidet sich
    ```
    <br>

    **Mit Generics:**
    ```go
    // Mit Generics kann diese Logik verallgemeinert werden
    // Es genügt EINE Funktion für beliebige Map-Key- und Value-Typen

    personAgeMap := map[string]int{
        "Alice": 31,
        "Bob":   29,
    }

    todoMap := map[int]string{
        1: "Einkaufen",
        2: "Gassi gehen",
        3: "Lernen",
    }

    // KEY_TYPE: Typ des Keys (muss "comparable" sein, damit er als Key erlaubt ist)
    // VALUE_TYPE: Typ des Values (kann alles sein)

    func GetKeysFromMap[KEY_TYPE comparable, VALUE_TYPE any](inputMap map[KEY_TYPE]VALUE_TYPE) []KEY_TYPE {
        // Erstellt ein Slice aller Keys, unabhängig vom Typ
        keys := make([]KEY_TYPE, 0, len(inputMap))
        for key := range inputMap {
            keys = append(keys, key)
        }
        return keys
    }

    // Anwendung:
    personNames := GetKeysFromMap[string, int](personAgeMap)    // -> []string{"Alice", "Bob"}
    todoIDs := GetKeysFromMap[int, string](todoMap)             // -> []int{1, 2, 3}

    // Vorteil: 
    // - Kein Code-Duplikat mehr
    // - Typsicherheit bleibt erhalten
    ```
    </details><br>
- 
    <details>
    <summary><strong>Allgemeine Datenstrukturen:</strong><br> Universelle Strukturen mit  beliebige Typen.<br><a>[Mehr anzeigen]</a></summary><br>
    
    Eigene Datenstrukturen wie Stacks, Queues oder Bäume sind elementare Bausteine vieler Programme. Sie werden für unterschiedliche Typen gebraucht, etwa für Zahlen, Zeichenketten oder eigene Strukturen.

    ### **Beispiel: Werte in einem Stack speichern und abrufen**
    --- 
    Go bringt keine generische Stack-Implementierung mit.
    Ein Stack (LIFO) ist eine Datenstruktur, bei der Elemente immer „oben“ abgelegt und entnommen werden. Typische Methoden sind unter anderem:
    - **Push**: Fügt ein Element oben auf den Stapel hinzu
    - **Pop**: Entfernt und gibt das oberste Element zurück
    - (Peek, IsEmpty, Size) 
    <br> <br>


    **Ohne Generics:**
    ```go
    // Problem: Es wird ein Stack für verschiedene ELEMENT_TYPEn (z.B. int, string) benötigt
    // Ohne Generics muss für jeden Typ eine eigene Stack-Implementierung existieren

    // ------- Stack für int-Werte -------
    type IntStack struct {
        elements []int
    }

    func (stack *IntStack) Push(value int) {
        stack.elements = append(stack.elements, value)
    }

    func (stack *IntStack) Pop() int {
        // Entfernt das oberste Element und gibt es zurück
        if len(stack.elements) == 0 {
            panic("Pop() wurde auf leeren IntStack aufgerufen")
        }
        elementCount := len(stack.elements)
        value := stack.elements[elementCount-1]
        stack.elements = stack.elements[:elementCount-1]
        return value
    }

    // ------- Stack für string-Werte -------
    type StringStack struct {
        elements []string
    }

    func (stack *StringStack) Push(value string) {
        stack.elements = append(stack.elements, value)
    }

    func (stack *StringStack) Pop() string {
        if len(stack.elements) == 0 {
            panic("Pop() wurde auf leeren StringStack aufgerufen")
        }
        elementCount := len(stack.elements)
        value := stack.elements[elementCount-1]
        stack.elements = stack.elements[:elementCount-1]
        return value
    }

    // Anwendung:
    var numberStack IntStack
    numberStack.Push(42)
    numberStack.Pop()

    var wordStack StringStack
    wordStack.Push("Hallo")
    wordStack.Pop()

    // Nachteil: 
    // - Viel redundanter Code.
    ```
    <br>

    **Mit Generics:**
    ```go
    // Mit Generics kann ein Stack für beliebige Typen implementiert werden.
    type Stack[ELEMENT_TYPE any] struct {
        elements []ELEMENT_TYPE
    }

    func (stack *Stack[ELEMENT_TYPE]) Push(value ELEMENT_TYPE) {
        stack.elements = append(stack.elements, value)
    }

    func (stack *Stack[ELEMENT_TYPE]) Pop() ELEMENT_TYPE {
        if len(stack.elements) == 0 {
            panic("Pop() wurde auf leeren Stack aufgerufen")
        }
        elementCount := len(stack.elements)
        value := stack.elements[elementCount-1]
        stack.elements = stack.elements[:elementCount-1]
        return value
    }

    // Anwendung:
    var intStack Stack[int]
    intStack.Push(42)
    intStack.Pop()

    var stringStack Stack[string]
    stringStack.Push("Hallo")
    stringStack.Pop()

    // Vorteil:
    // - Es wird nur noch eine Datenstruktur benötigt.
    // - Typsicherheit bleibt erhalten
    // - Intuitive Nutzung: Stack[int], Stack[string], Stack[MeinTyp] usw.
    ```
    </details><br>
- 
    <details>
    <summary><strong>Identische Methoden:</strong><br> Wrapper, die exakt dieselbe Logik für verschiedene Typen bereitstellen.<br><a>[Mehr anzeigen]</a></summary> <br>

    Viele Programme benötigen Funktionen, die unabhängig vom Element-Typ dieselbe Aufgabe erledigen, zum Beispiel das Suchen oder Zählen von Werten in Listen. Ohne Generics muss für jeden Typ eine eigene Funktion geschrieben werden, obwohl die Logik identisch ist.


    ### **Beispiel:** Die Position eines Elements in einer Liste bestimmen
    --- 

    In Go gibt es keine Standardfunktion, die für beliebige Slice-Typen den Index eines gesuchten Elements liefert. Ein häufiger Stolperstein: Der Typ im Slice muss „vergleichbar“ sein. Zusätzlich gibt es in Go keine Exceptions wie in anderen Sprachen. Fehlerfälle (z.B. "Element nicht gefunden") werden hier stattdessen über die [error](https://go.dev/doc/tutorial/handle-errors)-Bibliothek behandelt.<br> <br>

    **Ohne Generics:**
    ```go
    // Ziel: Die Position (Index) eines Elements im Slice soll ermittelt werden
    // Ohne Generics ist für jeden Typ eine eigene Funktion erforderlich
    // Go hat keine Exceptions, wenn das Element nicht gefunden wird, wird -1 und ein Fehler zurückgegeben

    // Funktion für int-Slices
    func GetIndexOfInt(intSlice []int, searchValue int) (int, error) {
        for index, element := range intSlice {
            if element == searchValue {
                return index, nil
            }
        }
        return -1, fmt.Errorf("int %v nicht gefunden", searchValue)
    }

    // Funktion für string-Slices
    func GetIndexOfString(stringSlice []string, searchValue string) (int, error) {
        for index, element := range stringSlice {
            if element == searchValue {
                return index, nil
            }
        }
        return -1, fmt.Errorf("string %q nicht gefunden", searchValue)
    }

    // Anwendung:
    numbers := []int{1, 2, 3}
    words := []string{"foo", "bar"}

    idxNum, errNum := GetIndexOfInt(numbers, 2)          // idxNum == 1, errNum == nil
    idxWord, errWord := GetIndexOfString(words, "baz")   // idxWord == -1, errWord != nil

    // Nachteil: 
    // - Es entstehen Code-Duplikate für jede Typ-Variante
    ```
    <br>

    **Mit Generics:**
    ```go
    // Mit Generics kann die Logik für alle vergleichbaren Typen verwendet werden
    // ELEMENT_TYPE: beliebiger Typ, muss aber "comparable" sein (für ==)
    // Mit Generics kann die Fehlerbehandlung für alle vergleichbaren Typen einheitlich abgebildet werden

    func GetIndexOfElement[ELEMENT_TYPE comparable](inputSlice []ELEMENT_TYPE, searchValue ELEMENT_TYPE) (int, error) {
        for index, element := range inputSlice {
            if element == searchValue {
                return index, nil
            }
        }
        return -1, errors.New("element not found")
    }

    // Anwendung:
    numbers := []int{1, 2, 3}
    words := []string{"foo", "bar"}

    idxNum, errNum := GetIndexOfElement(numbers, 2)        // idxNum == 1, errNum == nil
    idxWord, errWord := GetIndexOfElement(words, "baz")    // idxWord == -1, errWord != nil

    // Vorteil: 
    // - Weniger Code, keine Duplikate
    // - Typsicherheit bleibt, da der Compiler prüft, ob == erlaubt ist
    ```
    </details><br>

## ❌ Wann Generics vermieden werden sollten 
- 
    <details>
    <summary><strong>Wenn Overgeneralization droht:</strong><br> Nicht jede Funktion muss generisch sein.<br><a>[Mehr anzeigen]</a></summary><br>

    Nicht jede wiederverwendbare Funktion profitiert von Generics. In einigen Fällen sind klassische Interfaces, konkrete Typen oder Reflection die bessere Wahl, um den Code übersichtlich und sicher zu halten.

    ### **Beispiel:** Zeichenketten in Großbuchstaben umwandeln
    --- 
    <br>

    
    **Mit Generics:**
    ```go
    // Diese Funktion ist eigentlich nur für strings gedacht, verwendet aber Generics
    // Das führt zu unsicherem Casten und unnötiger Komplexität

    func ToUpperCase[SOME_TYPE any](inputValue SOME_TYPE) SOME_TYPE {
        // Castet inputValue zur Laufzeit auf string. Führt zu einer Panic, falls kein string übergeben wird
        upper := strings.ToUpper(inputValue.(string)) 
        // Castet zurück auf SOME_TYPE. Mehr Overhead, keine echte Typsicherheit
        return any(upper).(SOME_TYPE)
    }

    // Nachteil: 
    // - Generics bringen hier keinen echten Nutzen, erhöhen aber das Risiko für Fehler zur Laufzeit
    ```
    <br>

    **Ohne Generics:**
    ```go
    // Hier ist es direkt ersichtlich, dass diese Funktion nur für strings vorgesehen ist
    // Der Compiler prüft den Typ automatisch, Laufzeitfehler werden so verhindert

    func ToUpperCase(inputString string) string {
        return strings.ToUpper(inputString)
    }

    // Vorteil: 
    // - Der Nutzen der Funktion ist einfacher und schneller erkennbar
    ```
    </details><br>

- 
    <details>
    <summary><strong>Wenn ein Interface ausreicht:</strong><br> Verwende eine vorhandene Schnittstelle, sobald alle benötigten Methoden schon definiert sind.<br><a>[Mehr anzeigen]</a></summary><br>

    Go bietet [Interface-Types](https://golangdocs.com/interfaces-in-golang) an. Sie erlauben ebenfalls generischen Code zu schreiben. Wenn alle benötigten Methoden über ein Interface abgedeckt werden können (z.B. [fmt.Stringer](https://pkg.go.dev/fmt#Stringer) für alles, was als String darstellbar ist), reicht das Interface völlig aus.

    ### **Beispiel:** Print-Methode für beliebige Typen, die fmt.Stringer implementieren
    --- 

    Interfaces in Go beschreiben Verhalten, nicht Typen. Das Interface fmt.Stringer garantiert, dass String() implementiert ist. <br> <br>
    
    **Mit Generics:**
    ```go
    // In diesem Beispiel werden Generics verwendet, aber gleichzeitig ein Interface als Typ-Constraint gefordert
    
    func Print[STRINGABLE_TYPE fmt.Stringer](value STRINGABLE_TYPE) {
        fmt.Println(value.String())
    }
    
    // Nachteil: 
    // - Der generische Ansatz macht die Signatur nur komplizierter
    // Generics bringen hier keinen Vorteil gegenüber der klassischen Interface-Schreibweise
    ```
    <br>

    **Ohne Generics:**
    ```go
    // Besser: Das Interface wird direkt als Typ verwendet
    // Jedes Objekt, das fmt.Stringer implementiert, kann unabhängig vom konkreten Typ genutzt werden
    
    func Print(value fmt.Stringer) {
        fmt.Println(value.String())
    }

    // Vorteil: 
    // - Klarer, lesbarer und idiomatischer Go-Code
    ```
    <br>

    </details><br>
- 
    <details>
    <summary><strong>Wenn die Logik sich pro Typ unterscheidet:</strong><br> Setze auf Interfaces oder konkrete Typen, wenn verschiedene Typen unterschiedliche Implementierungen benötigen. <br><a>[Mehr anzeigen]</a></summary><br>

    Wenn jede Typ-Variante ihre ganz eigene Funktionslogik erfordert, führt ein generischer Ansatz unweigerlich zu Type Switches, Type Casts oder Reflection und bricht die Typsicherheit. Hier sind Interfaces und Methoden die idiomatische Lösung in Go.

    ### **Beispiel:** Flächenberechnung für unterschiedliche Geometrie-Typen
    ---
    <br>

    **Mit Generics:**
    ```go
    // Mit Generics können zwar verschiedene Typen akzeptiert werden, allerdings benötigt jede Variante eine eigene Berechnung
    // Das führt zu Konstrukte wie "type switches" und ist fehleranfällig sowie schwer wartbar

    type Circle struct{ Radius float64 }
    type Rectangle struct{ Width, Height float64 }

    func CalculateArea[SHAPE_TYPE any](shape SHAPE_TYPE) float64 {
        switch typedShape := any(shape).(type) {
        case Circle:
            return math.Pi * typedShape.Radius * typedShape.Radius
        case Rectangle:
            return typedShape.Width * typedShape.Height
        default:
            panic("unsupported type")
        }
    }

    // Nachteil: 
    // - Kein echter Nutzen von Generics
    // - Die einzelnen Berechnungen sind nicht auf den ersten Blick ersichtlich
    ```
    <br>
    
    **Ohne Generics:**
    ```go
    // Jeder Typ implementiert die für ihn passende Berechnung selbst
    type Circle struct{ Radius float64 }
    type Rectangle struct{ Width, Height float64 }

    // Methode für Circle
    func (c Circle) Area() float64 {
        return math.Pi * c.Radius * c.Radius
    }

    // Methode für Rectangle
    func (r Rectangle) Area() float64 {
        return r.Width * r.Height
    }

    // Vorteil: 
    // - Implementierungen sind klar strukturiert und kompakt
    // - Neue Geometrie-Typen können einfach Area() implementieren und sofort eingesetzt werden
    ```
    </details><br>
- 
    <details>
    <summary><strong>Wenn Reflection die passendere Wahl ist:</strong><br> 

    Greife auf [Reflection](https://go.dev/blog/laws-of-reflection) zurück, wenn hochdynamisch mit beliebigen Typen gearbeitet werden muss und generische Constraints nicht ausreichen. <br><a>[Mehr anzeigen]</a></summary><br>
    Bei hochdynamischen Aufgaben, für die zur Laufzeit Informationen über beliebige Typen benötigt werden (z.B. Anzahl der Felder in einem Struct), kommen in Go üblicherweise Reflection zum Einsatz.

    ### **Beispiel:** Anzahl der Felder in einem Struct ermitteln
    ---
    <br>
    
    **Mit Generics:**
    ```go
    // Problem: Soll zur Laufzeit z.B. die Anzahl der Felder eines Structs ermittelt werden, ist hier Reflection nötig
    // Die Typsicherheit geht dabei verloren, weil alles zur Laufzeit geschieht

    func GetNumberOfFieldsInStruct[STRUCT_TYPE any](structValue STRUCT_TYPE) int {
        // reflect.TypeOf(structValue) gibt zur Laufzeit Infos zum Typ zurück.
        return reflect.TypeOf(structValue).NumField()
    }
    ```
    <br>

    **Ohne Generics:**
    ```go
    // Für Reflection wird in Go traditionell interface{} verwendet
    // Das ist in diesem Fall vollkommen ausreichend und nicht schlechter als die Generics-Variante
    
    func GetNumberOfFieldsInStruct(structValue interface{}) int {
        return reflect.TypeOf(structValue).NumField()
    }
    
    // Vorteil: 
    // - Die Funktion ist besser lesbar
    ```
    </details><br>

## Weiterführende Ressourcen 
- [When To Use Generics](https://go.dev/blog/when-generics)  
- [Type Parameters Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)  
- [Generics in Go: Tips & Pitfalls](https://medium.com/@letsCodeDevelopers/generics-in-go-use-cases-tips-and-pitfalls-e25ec564c9a5)
- [15 Reasons Why You Should Use Generics in Go](https://medium.com/@jamal.kaksouri/15-reasons-why-you-should-use-generics-in-go-39601c3be6e0)
- [Generics in Go: A Comprehensive Guide with Code Examples](https://expertbeacon.com/generics-in-go-a-comprehensive-guide-with-code-examples/)