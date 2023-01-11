package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBlock_Serialize(t *testing.T) {
	type fields struct {
		Hash     []byte
		Data     []byte
		PrevHash []byte
		Nonce    int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test",
			fields: fields{
				Hash:     []byte("hash"),
				Data:     []byte("data"),
				PrevHash: []byte("prevhash"),
				Nonce:    1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Block{
				Hash:     tt.fields.Hash,
				Data:     tt.fields.Data,
				PrevHash: tt.fields.PrevHash,
				Nonce:    tt.fields.Nonce,
			}
			// Serialize
			data, err := b.Serialize()
			require.NoError(t, err)
			// Deserialize
			newBlock := &Block{}
			err = newBlock.Deserialize(data)
			require.NoError(t, err)
			// Compare
			require.Equal(t, b, newBlock)
		})
	}
}
