package main

import (
"fmt"
)

// Tree\[T] ist eine wiederverwendbare, typunspezifische binäre Suchbaum-Datenstruktur.
// Verwende Generics, wenn du dieselbe Logik für verschiedene Typen ohne duplizierten Code
// einsetzen möchtest. Hier wird eine Vergleichsfunktion benötigt, um die Ordnung zu definieren.
//
// Beispiel:
//   - Numerische Werte: Vergleich durch Subtraktion
//   - Zeichenketten: Vergleich durch lexikografische Ordnung
//
// Ohne Generics müsste man für jeden Typ einen eigenen Baum implementieren oder interface{} nutzen,
// was Typsicherheit und Performance beeinträchtigt.
type Tree\[T any] struct {
// cmp vergleicht zwei Werte vom Typ T und gibt -1, 0 oder +1 zurück.
cmp  func(a, b T) int
root \*node\[T]
}

// node\[T] ist ein Knoten im binären Baum.
type node\[T any] struct {
left, right \*node\[T]
val         T
}

// find gibt einen Zeiger auf den Zeiger auf einen node\[T] zurück:
//  - zeigt auf einen existierenden Knoten mit val,
//  - oder auf die Position, wo ein neuer Knoten angehängt würde.
func (bt \*Tree\[T]) find(val T) \*\*node\[T] {
pl := \&bt.root
for \*pl != nil {
switch cmp := bt.cmp(val, (\*pl).val); {
case cmp < 0:
pl = &(\*pl).left
case cmp > 0:
pl = &(\*pl).right
default:
return pl // gefunden
}
}
return pl // nicht gefunden, Rückgabe der Einfügeposition
}

// Insert fügt val in den Baum ein, falls er noch nicht vorhanden ist.
// Gibt true zurück, falls eingefügt wurde, sonst false.
func (bt \*Tree\[T]) Insert(val T) bool {
pl := bt.find(val)
if \*pl != nil {
// Wert bereits vorhanden
return false
}
// Neuen Knoten erstellen und verlinken
\*pl = \&node\[T]{val: val}
return true
}

func main() {
// Beispiel 1: Baum mit Integern
intTree := Tree\[int]{cmp: func(a, b int) int {
// Subtraktion liefert negative, null oder positive Zahl
return a - b
}}
for \_, v := range \[]int{5, 3, 7, 3, 9, 1} {
inserted := intTree.Insert(v)
fmt.Printf("Insert %d: inserted=%v\n", v, inserted)
}

```
// Beispiel 2: Baum mit Strings
stringTree := Tree[string]{cmp: func(a, b string) int {
    if a < b {
        return -1
    } else if a > b {
        return 1
    }
    return 0
}}
for _, s := range []string{"go", "java", "cpp", "go", "rust"} {
    inserted := stringTree.Insert(s)
    fmt.Printf("Insert %q: inserted=%v\n", s, inserted)
}

// Hinweis: ohne Generics könnte man entweder:
//   1) interface{} und Type Assertion nutzen (unsicherer, langsamer),
//   2) für jeden konkreten Typ eine eigene Implementierung schreiben.
// Beide Ansätze sind unpraktisch im Vergleich zu Generics.
```

}

/\*
Nicht-generische Version (für int) würde so aussehen:

package main

import (
"fmt"
)

type IntTree struct {
root \*intNode
}

type intNode struct {
left, right \*intNode
val         int
}

// find und Insert hier speziell für int implementiert...

func main() {
tree := IntTree{}
// Wiederholung des Insert-Codes nur für int
}
\*/
