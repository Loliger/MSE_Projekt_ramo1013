# MSE_Projekt_ramo1013

## Ist es möglich beide Übersetzungsmethoden für jedes Beispiel anzuwenden?
### 1. Monomorphization Methode
Die *monomorphization Methode* um Generics zu übersetzen funktioniert für alle Beispiele, da wir für alle Datentypen, welche durch die Generics abgedeckt sind eigene Methoden, structs usw. definieren können. So können wir zB jeweils eine Methode für int und eine für float32 Werte wie im "sum" Beispiel definieren. 

    func SumMMInt(xs []int) int {
	    var x int
	    x = 0
	    for _, v := range xs {

		    x = x + v
	    }

	    return x
    }
    func SumMMFloat(xs []float32) float32 {
	    var x float32
	    x = 0
	    for _, v := range xs {

		    x = x + v
	    }
	    return x
    }

Dies erfordert natürlich einen deutlich größeren Aufwand, umso mehr Datentypen wir mit den Generics abdecken.

### 2. Generic translation Methode
Die *generic translation* Methode funktioniert bei dem "node" Beispiel, da wir hier nur maximal den Wert des Knotens lesen und sonst nichts. Dies funktioniert, da *interface{}* das leere Interface darstellt, welches keine Methoden hat. Für structural subtyping müssen die alle Methoden des Interface für den Datentyp implementiert sein. Da alle Datentype mindestens keine Methode implementieren erfüllen sie alle die Vorraussetzung und sind somit ein struktural subtyp von *interface{}*.

Für das Beispiel "sum" funktioniert das ganze auch. Zunächst müssen die Eingabe Arrays umgewandelt werden in Arrays vom Typ *interface{}*. Dann gibt es einen Fehler, da auf *interface{}* die Addition nicht definiert ist und der Datentyp der Variable *interface{}* ist und nicht der Typ, welcher übergeben wurde. An dieser Stelle müssen wir *Type Assertions* verwenden:

    switch xs[0].(type) {
	case int:
        x = 0
		for _, v := range xs {
			x = x.(int) + v.(int)
		}
    case float32:
        x = float32(0)
		for _, v := range xs {
			x = x.(float32) + v.(float32)
		}
    }

Mit Hilfe der *Type Assertions* können wir nun verschiedene Fallunterscheidungen vornehmen, je nachdem welcher Datentyp vorliegt. Wir müssen ebenfalls im Fall *float32* die Variable x speziell mit der 0 vom Typ float32 initialisieren, da der Compiler x sonst als Int oder bei 0.0 als float64 interpretiert.

Für das dritte Beispiel "swap" funktioniert das ganze, man muss hier nur darauf achten, dass man die Variablen, welche man mit dem Pointer übergibt als empty interface initialisiert. Da die Methode einen Pointer auf ein Objekt vom Typ *interface{}* erwartet und nicht etwa einen string oder int Pointer. Hier ist zu beachten, dass wir an dieser Stelle Variablen von unterschiedlichen Typen tauschen können. Es werden an dieser Stelle ebenfalls type Assertions benötigt, um sicherzustellen, dass die Typen gleich sind.

    func SwapGT(x *interface{}, y *interface{}) {
        switch (*x).(type) {
        case int:
            _, ok := (*y).(int)
            if ok {
                tmp := *x
                *x = *y
                *y = tmp
            }
        case bool:
            _, ok := (*y).(bool)
            if ok {
                tmp := *x
                *x = *y
                *y = tmp
            }
        case string:
            _, ok := (*y).(string)
            if ok {
                tmp := *x
                *x = *y
                *y = tmp
            }
        }
    }

## Welche Programmiersprachen Features werden benötigt?
### 1. Monomorphization Methode

Für die *momorphization*-Methode brauchen wir keine besonderen Features, da wir hier alle Implementierungen für jeden Datentyp machen. Dabei bennenen wir die Methoden einfach etwas anderst und rufen dann die spezifische Methoden auf.

### 2. Generic translation Methode

Damit die *generic translation* Methode funktioniert wird structural subtyping benötigt, da wir durch structural subtyping jeden Datentyp über das leere Interface abbilden können. Ebenfalls kann es sein, dass wir für bestimmte Implementierungen type assertions benötigen. 