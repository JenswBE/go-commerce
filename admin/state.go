package admin

import (
	"fmt"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// handleStatefulBoolFlag manages a stateful boolean flag.
// The boolean can be set using a query parameter.
// If the query parameter is missing, the state is fetched from the session.
// Session is automatically updated on changed state.
func handleStatefulBoolFlag(c *gin.Context, stateFlag string) (bool, error) {
	// Get session
	session := sessions.Default(c)

	// Fetch flag state from query
	queryString := c.Query(stateFlag)
	queryState := strings.EqualFold(queryString, "true")

	// Fetch session state
	sessionStateFlag := "state_" + stateFlag
	sessionRaw := session.Get(sessionStateFlag)
	sessionState, ok := sessionRaw.(bool)

	// Return session state if query is empty
	if queryString == "" {
		if !ok {
			return false, nil
		}
		return sessionState, nil
	}

	// Persist state if was changed or incorrectly set
	if sessionState != queryState || !ok {
		session.Set(sessionStateFlag, queryState)
		err := session.Save()
		if err != nil {
			err = fmt.Errorf("failed to save session after setting bool %s to %t: %w", sessionStateFlag, queryState, err)
			log.Error().Err(err).Str("key", sessionStateFlag).Bool("value", queryState).Msg("Failed to save session after setting bool")
			return false, fmt.Errorf("failed to save session after setting bool %s to %t: %w", sessionStateFlag, queryState, err)
		}
	}
	return queryState, nil
}
