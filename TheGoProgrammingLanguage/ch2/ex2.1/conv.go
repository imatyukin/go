// Упражнение 2.1.
// Добавьте в пакет tempconv типы, константы и функции для работы с температурой по шкале Кельвина, в которой нуль
// градусов соответствует температуре -273.15°С, а разница температур в 1К имеет ту же величину, что и 1°С.

//!+

package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k + ZeroK) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

//!-
