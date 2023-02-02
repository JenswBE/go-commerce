package shortid_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/generics"
	"github.com/JenswBE/go-commerce/utils/shortid"
)

func Test_Base58_RoundTrip_Success(t *testing.T) {
	// Setup test
	s := shortid.NewBase58Service()
	id := entities.NewID()

	// Encode
	shortID := s.Encode(id)
	require.NotEmpty(t, shortID)

	// Decode
	result, err := s.Decode(shortID)
	require.NoError(t, err)
	require.Equal(t, id, result)
}

func Test_Base58_Encode_Success(t *testing.T) {
	id := generics.Must(entities.NewIDFromString("936c689f-f746-4076-a5bc-ece83f55353f"))
	result := shortid.NewBase58Service().Encode(id)
	require.Equal(t, "KCrsG1RUKHLuc5hS2GZvJa", result)
}

func Test_Base58_Decode_Success(t *testing.T) {
	result, err := shortid.NewBase58Service().Decode("KCrsG1RUKHLuc5hS2GZvJa")
	require.NoError(t, err)
	require.Equal(t, generics.Must(entities.NewIDFromString("936c689f-f746-4076-a5bc-ece83f55353f")), result)
}

func Test_Base58_Decode_InvalidID_Failure(t *testing.T) {
	result, err := shortid.NewBase58Service().Decode("invalid")
	require.Error(t, err)
	require.True(t, result.IsNil(), "Result must be a nil UUID")
}
