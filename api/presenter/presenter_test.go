package presenter_test

import (
	"testing"

	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/generics"
	"github.com/JenswBE/go-commerce/utils/sanitizer"
	"github.com/JenswBE/go-commerce/utils/shortid"
	"github.com/stretchr/testify/require"
)

func Test_ParseID_ShortID_Success(t *testing.T) {
	// Get presenter
	p := presenter.New(shortid.NewBase58Service(), sanitizer.NewBluemondayService())

	// Call function
	id, err := p.ParseID("KCrsG1RUKHLuc5hS2GZvJa")

	// Assert results
	require.NoError(t, err)
	require.Equal(t, generics.Must(entities.NewIDFromString("936c689f-f746-4076-a5bc-ece83f55353f")), id)
}

func Test_ParseID_UUID_Success(t *testing.T) {
	// Get presenter
	p := presenter.New(shortid.NewBase58Service(), sanitizer.NewBluemondayService())

	// Call function
	id, err := p.ParseID("619add22-faf2-4d54-a662-8a0206b967c4")

	// Assert results
	require.NoError(t, err)
	require.Equal(t, generics.Must(entities.NewIDFromString("619add22-faf2-4d54-a662-8a0206b967c4")), id)
}

func Test_ParseID_InvalidID_Failure(t *testing.T) {
	// Get presenter
	p := presenter.New(shortid.NewBase58Service(), sanitizer.NewBluemondayService())

	// Call function
	id, err := p.ParseID("invalid")

	// Assert results
	require.Error(t, err)
	require.True(t, id.IsNil(), "ID must be nil")
}
