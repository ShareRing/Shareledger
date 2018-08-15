package types

//----------------------------------------

type Signature interface {
	Bytes() []byte
	IsZero() bool
	Equals(Signature) bool
}


//-------------------------------------

var _ Signature = SignatureSecp256k1{}

// Implements Signature
type SignatureSecp256k1 []byte

func (sig SignatureSecp256k1) Bytes() []byte {
	bz, err := cdc.MarshalBinaryBare(sig)
	if err != nil {
		panic(err)
	}
	return bz
}

func (sig SignatureSecp256k1) IsZero() bool { return len(sig) == 0 }

func (sig SignatureSecp256k1) String() string { return fmt.Sprintf("/%X.../", Fingerprint(sig[:])) }

func (sig SignatureSecp256k1) Equals(other Signature) bool {
	if otherSecp, ok := other.(SignatureSecp256k1); ok {
		return bytes.Equal(sig[:], otherSecp[:])
	} else {
		return false
	}
}
