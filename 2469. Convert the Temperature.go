// https://leetcode.com/problems/convert-the-temperature/description/

// You are given a non-negative floating point number rounded to two decimal places celsius, that denotes the temperature in Celsius.

// You should convert Celsius into Kelvin and Fahrenheit and return it as an array ans = [kelvin, fahrenheit].

// Return the array ans. Answers within 10-5 of the actual answer will be accepted.


func convertTemperature(celsius float64) []float64 {
    ans := []float64{}
    ans = append(ans, celsius + 273.15)
    ans = append(ans, celsius * 1.80 + 32.00)
    return ans
}
