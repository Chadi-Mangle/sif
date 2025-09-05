package utils

import (
	"os"
	"strings"

	pkgutils "github.com/Chadi-Mangle/templ-hmr-setup/package/utils"
)

var domain string = os.Getenv("EMAIL_DOMAIN")

func GetEmailAddress(firstName string, lastName string) string {
	cleanFirstName := pkgutils.RemoveAccents(firstName)
	cleanLastName := pkgutils.RemoveAccents(lastName)

	cleanFirstName = strings.ReplaceAll(cleanFirstName, " ", "_")
	cleanLastName = strings.ReplaceAll(cleanLastName, " ", "_")

	cleanFirstName = strings.ToLower(cleanFirstName)
	cleanLastName = strings.ToLower(cleanLastName)

	email := cleanFirstName + "." + cleanLastName + "@" + domain

	return email
}
