package i18n

import (
	"fmt"
	"unicode"

	"github.com/JenswBE/go-commerce/entities"
)

func InvalidUUID(objectType ObjectType, value string) string {
	return fmt.Sprintf(`Ongeldige ID "%s" voor %s`, value, objectType)
}

func DeleteFailed(objectType ObjectType, objectID entities.ID, err error) string {
	return fmt.Sprintf(`Verwijderen van %s "%s" mislukt: %v`, objectType, objectID, err)
}

func DeleteSuccessful(objectType ObjectType) string {
	return fmt.Sprintf(`%s succesvol verwijderd`, capitalFirst(string(objectType)))
}

func capitalFirst(input string) string {
	// Based on https://stackoverflow.com/a/70259366
	r := []rune(input)
	return string(append([]rune{unicode.ToUpper(r[0])}, r[1:]...))
}
