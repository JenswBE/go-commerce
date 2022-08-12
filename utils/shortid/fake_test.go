package shortid_test

import (
	"testing"

	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/shortid"
	"github.com/stretchr/testify/require"
)

func Test_Fake_RoundTrip_Success(t *testing.T) {
	// Setup test
	s := shortid.NewFakeService()
	id := entities.NewID()

	// Encode
	shortID := s.Encode(id)
	require.NotEmpty(t, shortID)

	// Decode
	result, err := s.Decode(shortID)
	require.NoError(t, err)
	require.Equal(t, id, result)
}
