package i18n

import (
	"fmt"
	"unicode"
)

func InvalidUUID(objectType ObjectType, value string) string {
	return fmt.Sprintf(`Ongeldige ID "%s" voor %s`, value, objectType)
}

func DeleteFailed(objectType ObjectType, name string, err error) string {
	if name == "" {
		return fmt.Sprintf(`Verwijderen van %s mislukt: %v`, objectType, err)
	}
	return fmt.Sprintf(`Verwijderen van %s "%s" mislukt: %v`, objectType, name, err)
}

func DeleteSuccessful(objectType ObjectType) string {
	return fmt.Sprintf(`%s succesvol verwijderd`, capitalFirst(string(objectType)))
}

func capitalFirst(input string) string {
	// Based on https://stackoverflow.com/a/70259366
	r := []rune(input)
	return string(append([]rune{unicode.ToUpper(r[0])}, r[1:]...))
}
