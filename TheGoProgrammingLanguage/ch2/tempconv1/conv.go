// Упражнение 2.2.
// Напишите программу общего назначения для преобразования единиц, аналогичную cf, которая считывает числа из
// аргументов командной строки (или из стандартного ввода, если аргументы командной строки отсутствуют) и преобразует
// каждое число в другие единицы, как температуру — в градусы Цельсия и Фаренгейта, длину — в футы и метры, вес — в
// фунты и килограммы и т.д.

//!+

package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToM converts a Feet length to Meter.
func FToM(f Feet) Meter { return Meter(f * 0.3048) }

// MToF converts a Meter length to Feet.
func MToF(m Meter) Feet { return Feet(m / 0.3048) }

// PToK converts a Pound weight to Kilogram.
func PToK(p Pound) Kilogram { return Kilogram(p * 0.453592) }

// KToP converts a Kilogram weight to Pound.
func KToP(k Kilogram) Pound { return Pound(k / 0.453592) }

//!-