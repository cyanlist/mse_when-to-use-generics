# When to use Generics in Go

Generics ermöglichen das Schreiben von Funktionen und Datenstrukturen, die für verschiedene Typen arbeiten, ohne auf `interface{}`-Tricks oder Reflection zurückzugreifen.

---

## Generics vs. Generic Code mit Interfaces

Es gibt in Go drei grundlegende Wege, generischen Code zu schreiben:

1. **Generics-Feature** (Go 1.18+):  
   – Typparameter an Funktionen oder Datentypen  
   – Compile-Time-Checks und Monomorphisierung  

2. **Leeres Interface** (`interface{}`):  
   – Fängt **jeden** Typ ein  
   – Kein richtiger Polymorphismus, Type Assertions/Reflection nötig  

3. **Custom Interfaces**:  
   – Definieren eine Methodensignatur (`io.Reader`, `http.Handler`…)  
   – Subtyp-Polymorphismus über dynamischen Dispatch

| Aspekt               | Generics-Feature                                  | `interface{}` (leeres Interface)                  | Custom Interfaces                                 |
|----------------------|---------------------------------------------------|---------------------------------------------------|---------------------------------------------------|
| **Polymorphismus**   | Parametrischer Polymorphismus (Compile-Time)      | kein echter—alle Typen erlaubt                    | Subtyp-Polymorphismus (Laufzeit)                  |
| **Typsicherheit**    | Vollständig—Compiler prüft Constraints            | keine—Type Assertions/Reflection nötig            | Methoden-Level-Typsicherheit                      |
| **Flexibilität**     | einmal festgelegter Typparameter                  | maximal—jeder Typ passt                           | auf definierte Methoden beschränkt                |
| **Performance**      | kein Laufzeit-Overhead                            | hoher Overhead bei Assertions/Reflection          | leichter Overhead durch Interface-Table           |
| **Komplexität**      | komplexe Signaturen bei vielen Constraints        | sehr einfach deklarierbar, aber Nutzung aufwändig | lesbar, solange Interface schlank bleibt          |
| **Fehlerdiagnose**   | Compile-Time-Fehler sichtbar                      | viele Laufzeit-Fehler                             | Methodennamen bekannt, Laufzeit-Fehler bei Aufruf |

---

## Einsatzgebiete generischen Codes

> **Wann generell generischer Code?**  
> Wenn Algorithmen oder Datenstrukturen _einmal_ implementiert und für _viele_ Typen wiederverwendet werden sollen.

...

---

## Konkrete Use-Cases

### ✅ Sinnvolle Use-Cases
- **Container-Utilities:** `Map`, `Filter`, `Reduce` für beliebige Slices. [(mapkeys.go)](examples/mapkeys.go)
- **Datenstrukturen:** Stacks, Queues, Trees mit Generics. [(tree.go)](examples/tree.go)
- **Algorithmen-Bibliotheken:** Sortieren, Suchen, Vergleichen ohne Duplikate. [(slicefn.go)](examples/slicefn.go)

---

## Weiterführende Ressourcen

- [An Introduction To Generics (Go Blog)](https://go.dev/blog/intro-generics)  
- [When To Use Generics (Go Blog)](https://go.dev/blog/when-generics)  
- [Type Parameters Proposal (Go-Proposal)](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)  
- [Generics in Go: Tips & Pitfalls (Medium)](https://medium.com/@letsCodeDevelopers/generics-in-go-use-cases-tips-and-pitfalls-e25ec564c9a5)  
