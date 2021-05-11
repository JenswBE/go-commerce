package presenter_test

import (
	"testing"

	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/pkg/shortid"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_ParseID_UUID_Success(t *testing.T) {
	// Get presenter
	p := presenter.New(shortid.NewBase58Service())

	// Call function
	id, err := p.ParseID("936c689f-f746-4076-a5bc-ece83f55353f")

	// Assert results
	require.NoError(t, err)
	require.Equal(t, uuid.MustParse("936c689f-f746-4076-a5bc-ece83f55353f"), id)
}

func Test_ParseID_ShortID_Success(t *testing.T) {
	// Get presenter
	p := presenter.New(shortid.NewBase58Service())

	// Call function
	id, err := p.ParseID("KCrsG1RUKHLuc5hS2GZvJa")

	// Assert results
	require.NoError(t, err)
	require.Equal(t, uuid.MustParse("936c689f-f746-4076-a5bc-ece83f55353f"), id)
}

func Test_ParseID_InvalidID_Failure(t *testing.T) {
	// Get presenter
	p := presenter.New(shortid.NewBase58Service())

	// Call function
	id, err := p.ParseID("invalid")

	// Assert results
	require.Error(t, err)
	require.Equal(t, uuid.Nil, id)
}
