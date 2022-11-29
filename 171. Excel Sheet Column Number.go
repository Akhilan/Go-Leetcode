// https://leetcode.com/problems/excel-sheet-column-number/


// Given a string columnTitle that represents the column title as appears in an Excel sheet, return its corresponding column number.


func titleToNumber(columnTitle string) int {
    
    
    var colNum int

	for _, char := range columnTitle {
        
		 colNum  = colNum*26 + int(char-'A') + 1
	}

	return colNum
    
}
