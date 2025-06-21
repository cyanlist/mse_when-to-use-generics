# When to use Generics in Go

> "The Go 1.18 release adds support for generics. Generics are the biggest change we’ve made to Go since the first open source release." (Quelle: The Go Blog, „Type Parameters in Go 1.18")

Generics ermöglichen es, Funktionen und Datenstrukturen zu definieren, die unabhängig von konkreten Typen arbeiten. Dadurch kann wiederverwendbaren, typsicheren und klaren Code geschrieben werden, ohne auf Reflection oder Interface{} zurückzugreifen.

**Vorteile:**

- Höhere Typsicherheit durch Compilezeit-Typprüfung
- Reduziert Boilerplate-Code
- Verbessert Wartbarkeit und Lesbarkeit 
- Potenzielle Performance-Optimierung durch Compilezeit-Monomorphisierung

**Nachteile:**

- Begrenzte Ausdrucksstärke der Typconstraints (komplexe Beziehungen nicht ausdrückbar)
- Keine generischen Methoden (keine eigenständige Typparameter definierbar)
- Eingeschränkte Unterstützung durch die Standardbibliothek

## Allgemeine Empfehlungen

_Die folgenden Punkte sind keine in Stein gemeißelten Regeln, sondern Vorschläge, die im jeweiligen Kontext mit gesundem Menschenverstand bewertet werden sollten._

>  **Faustregel:**
> Vermeide Generics, bis man denselben Code mehrmals schreiben muss.

- **Klarheit & Wartbarkeit:** Nutze Generics nur, wenn sie echten Mehrwert bieten.
- **Lesbarkeit priorisieren:** Der generische Code sollte intuitiv und verständlich bleiben.
- **Einfachheit bewahren:** Bei schmalen oder trivialen Use-Cases lieber auf konkrete Typen setzen.

## Generics vs. Interface{}

Vor Go 1.18 wurde oft interface{} verwendet, um generelle Funktionen zu implementieren. 
Mit Generics steht jetzt eine bessere Alternative bereit, um typisierte, effiziente und wartbare Lösungen zu schaffen.

|  | Generics | Interface{} |
| -- | -- | -- |
| Typsicherheit | Vollständig | Nicht vorhanden |
| Performance | Höher (keine Laufzeitüberprüfung) | Niedriger (Reflection)
| Lesbarkeit | Hoch | Niedrig
| Wartbarkeit | Gut | Mittel bis niedrig

-> Generics sollten grundsätzlich bevorzugt und Interface{} vermieden werden.

## Generics vs. Interfaces

Generics und Interfaces werden beide verwendet, um Abstraktion und Wiederverwendbarkeit in Go zu erzielen. 
Allerdings unterscheiden sie sich in der Art und Weise, wie sie Typflexibilität und Polymorphismus realisieren:
|  | Generics | Interfaces |
|--|--| -- |
| **Typflexibilität**| Zur Compilezeit festgelegt| Dynamisch zur Laufzeit |
| **Performance**| Meist besser | Etwas langsamer
| **Polymorphismus**| Statisch (kompilierungsbasiert) | Dynamisch (Laufzeit)
| **Code-Wiederverwendung** | Hoch (allgemeine Logik) | Hoch (gemeinsame Methodik)
  
 -> Interfaces für dynamisches Verhalten zur Laufzeit, Generics für statisch sichere Typverallgemeinerungen.

## Implementierungs-Use-Cases: Wann Generics einsetzen vs. vermeiden
  
### ✅ Sinnvolle Einsatzgebiete
-  **Container-Funktionen verallgemeinern:** Operationen für beliebige Datenansammlungen (Maps, Slices, ...)
-  **Wiederverwendbare, typunspezifische Datenstrukturen:** Generische Bäume, Listen, Stacks, ...
-  **Funktionale Helfer:** Transformationen/Auswertungen generischer Datenstrukturen
-  **Einheitliche Methoden-Implementierung:** Gemeinsame Logik für alle Typen
  
### 🚫 Wann du besser darauf verzichten solltest
-  **Einzelne Methodenaufrufe:** Bei einmalige, typgebundenen Operationen
-  **Stark dynamische Typen:** Fällen, in denen Reflection ohnehin nötig ist
-  **Heterogene Implementierungen:** Unterschiedliche Logik pro Typvariante

## Weiterführende Ressourcen
- [An Introduction To Generics](https://go.dev/blog/intro-generics)
- [When To Use Generics - The Go Programming Language](https://go.dev/blog/when-generics)
[Generics in Go: Use Cases, Tips, and Pitfalls 🧰🐹 | by Let's code | Medium](https://medium.com/@letsCodeDevelopers/generics-in-go-use-cases-tips-and-pitfalls-e25ec564c9a5)
- [Type Parameters Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)(https://medium.com/@letsCodeDevelopers/generics-in-go-use-cases-tips-and-pitfalls-e25ec564c9a5)
- [Why Go’s Generics Might Be Worse Than No Generics at All | by Leapcell | Apr, 2025 | Medium](https://leapcell.medium.com/why-gos-generics-might-be-worse-than-no-generics-at-all-7b2373ce99f0)
- <!-- https://stackedit.io/app# -->