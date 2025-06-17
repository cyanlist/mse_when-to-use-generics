# When to use Generics in Go
_Kurze Zusammenfassung, wann und warum man Generics in Go einsetzen sollte._

**Kurz & Knapp:** Generics (Typ-Parameter) ermöglichen typsichere Wiederverwendung von Funktionen, Typen und Datenstrukturen ohne Code-Duplikation oder `interface{}`-Nutzung.

## Allgemeine Empfehlungen

_Die folgenden Punkte sind keine in Stein gemeißelten Regeln, sondern Vorschläge, die im jeweiligen Kontext mit gesundem Menschenverstand bewertet werden sollten._

> **Faustregel:**  
> Vermeide Generics, bis du denselben Code mehrmals schreiben musst.  

- **Klarheit & Wartbarkeit:** Nutze Generics nur, wenn sie echten Mehrwert bieten.
- **Einfachheit bewahren:** Bei schmalen oder trivialen Use-Cases lieber auf konkrete Typen setzen.  

## Implementierungs-Use-Cases: Wann Generics einsetzen vs. vermeiden

### ✅ Sinnvolle Einsatzgebiete
- **Container-Funktionen verallgemeinern:** Operationen für beliebige Datenansammlungen (Maps, Slices, ...)
- **Wiederverwendbare, typunspezifische Datenstrukturen:** Generische Bäume, Listen, Stacks, ...
- **Einheitliche Methoden-Implementierung:** Gemeinsame Logik für alle Typen 
- **Funktionale Helfer:** Transformationen/Auswertungen generischer Datenstrukturen

### 🚫 Wann du besser darauf verzichten solltest
- **Einzelner Methodenaufruf:** Bei einmalige, typgebundenen Operationen
- **Heterogene Implementierungen:** Unterschiedliche Logik pro Typvariante
- **Stark dynamische Typen:** Fällen, in denen Reflection ohnehin nötig ist

## TODO

 - [ ] "Einleitung" verbessern
 - [ ] "Klarheit & Wartbarkeit" ergänzen -> Overhead minimieren. Falls benötigt, kann man Generics nachträglich leicht ergänzen
 - [ ] Generics für Tests?

## Quellen
- [When To Use Generics - The Go Programming Language](https://go.dev/blog/when-generics)
- [5 Practical Go Generics Examples to Level Up Your Code - DEV Community](https://dev.to/shrsv/5-practical-go-generics-examples-to-level-up-your-code-3m96#:~:text=Go%20generics,%20introduced%20in%20Go%201.18,%20let%20you,cases%20that%20show%20their%20power%20in%20real-world%20scenarios.)
- [Generics in Go: Use Cases, Tips, and Pitfalls 🧰🐹 | by Let's code | Medium](https://medium.com/@letsCodeDevelopers/generics-in-go-use-cases-tips-and-pitfalls-e25ec564c9a5)
- [GitHub - akutz/go-generics-the-hard-way: A hands-on approach to getting started with Go generics.](https://github.com/akutz/go-generics-the-hard-way)
- [Type Parameters Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)

- <!-- https://stackedit.io/app# -->