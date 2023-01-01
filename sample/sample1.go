package sample

import (
	"fmt"
	"strings"
)

func main() {
	keyArr := []string{"asitha", "asdfasdf", "asdfasdf"}
	keyArrString := strings.Join(keyArr, "','")
	keyArrString = "('" + keyArrString + "')"
	query := fmt.Sprintf(`SELECT data FROM %q WHERE data->'payload'->>'stream_id'=$1 and data->'payload'->>'status' NOT IN %s and deleted=$2;`, "table", keyArrString)
	fmt.Println(query)
}
