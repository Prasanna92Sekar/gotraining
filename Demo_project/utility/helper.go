// Contains helper function for various functions
package utility


type str_map map[string]interface{} 

// Function to return map with user message and data entered by user.
func GetPrompt(prompt string , data ...interface{} ) (str_map){
	promptMap := make(str_map)

	promptMap["message"] = prompt

	if len(data) > 0 { 
		promptMap["data"] = data
	}

	return promptMap
}


