package swift

import "log"

// Generate random Swift function
func GenFunction() (string, error) {
	function, id := chooseTypeFunction()
	log.Printf("Позиция, с которой можно начать вставку дальнейшего кода: %v", id)

	return function, nil
}
