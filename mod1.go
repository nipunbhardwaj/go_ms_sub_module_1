package mod1

import (
	"fmt"

	"github.com/nipunbhardwaj/go_ms_sub_module_1/api"
)

func SecretProcess() {
	fmt.Println("Running the secret process!")
	item := api.Item{
		Label: "Inside Item",
	}
	item.PrintLabel()
}
