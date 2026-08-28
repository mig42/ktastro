package cli

import "fmt"

func Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("uso: nombre_repo <comando>")
	}
	fmt.Println("comando:", args[0])
	return nil
}
