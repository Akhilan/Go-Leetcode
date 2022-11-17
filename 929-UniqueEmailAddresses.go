func numUniqueEmails(emails []string) int {
    
    emailMap := make(map[string]bool)
    
    for _, email := range emails {
        
        // Split the email name in to local and domain names 
        split := strings.Split(email, "@") 
        
        //Split the local name to used part and ignored part
        localTemp := strings.Split(split[0], "+")
        
        //Removing the dot character
        localFinal := strings.Replace(localTemp[0], ".", "", -1)
        
        //Final email address by joining local and domain names
        emailAddr := string(localFinal) + "@" + split[1]
        
        
        //make the map value to true
        emailMap[emailAddr] = true
        
    }
    
    return len(emailMap)
         
}
