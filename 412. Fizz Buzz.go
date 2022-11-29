//https://leetcode.com/problems/fizz-buzz/

//Given an integer n, return a string array answer (1-indexed) where:

//    answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
//    answer[i] == "Fizz" if i is divisible by 3.
//    answer[i] == "Buzz" if i is divisible by 5.
//    answer[i] == i (as a string) if none of the above conditions are true.


func fizzBuzz(n int) []string {
    
    
    s := make([]string, 0, n)
    
    
    for i:=1; i <=n ; i++{
        var st string
        if i%3 == 0 && i%5 == 0 {
            st = "FizzBuzz"
        } else if i%3 == 0 {
            st = "Fizz"
        } else if i%5 == 0 {
            st = "Buzz"
        } else {
            st = strconv.Itoa(i)
        }
        
        s = append(s, st)
    }
    
    return s
    
}
