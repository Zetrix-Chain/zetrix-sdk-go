package signature

import (
	"testing"

	"github.com/Zetrix-Chain/zetrix-sdk-go/src/crypto/keypair"
)

var message = []byte("zetrix signature test")

func newSignedMessage(t *testing.T) (publicKey string, sign string) {
	t.Helper()
	publicKey, privateKey, _, err := keypair.Create()
	if err != nil {
		t.Fatalf("keypair.Create: %v", err)
	}
	sign, err = Sign(privateKey, message)
	if err != nil {
		t.Fatalf("Sign: %v", err)
	}
	if len(sign) != 2*SignatureSize {
		t.Fatalf("Sign returned %d hex chars, want %d", len(sign), 2*SignatureSize)
	}
	return publicKey, sign
}

// verifyNoPanic calls Verify and fails the test if it panics.
func verifyNoPanic(t *testing.T, publicKey string, msg []byte, sign string) (ok bool) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Verify panicked: %v", r)
		}
	}()
	return Verify(publicKey, msg, sign)
}

func TestVerifyAcceptsValidSignature(t *testing.T) {
	publicKey, sign := newSignedMessage(t)
	if !verifyNoPanic(t, publicKey, message, sign) {
		t.Fatal("Verify rejected a valid signature")
	}
}

func TestVerifyRejectsSignatureForOtherMessage(t *testing.T) {
	publicKey, sign := newSignedMessage(t)
	if verifyNoPanic(t, publicKey, []byte("another message"), sign) {
		t.Fatal("Verify accepted a signature for a different message")
	}
}

// A signature that does not decode to exactly 64 bytes must return false, not panic.
func TestVerifyRejectsWrongLengthSignatureWithoutPanic(t *testing.T) {
	publicKey, sign := newSignedMessage(t)
	cases := []struct {
		name string
		sign string
	}{
		{"empty", ""},
		{"one byte", "00"},
		{"32 bytes", sign[:64]},
		{"63 bytes", sign[:126]},
		{"65 bytes (valid signature plus one byte)", sign + "00"},
		{"128 bytes (valid signature twice)", sign + sign},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if verifyNoPanic(t, publicKey, message, c.sign) {
				t.Fatalf("Verify accepted a %d-hex-char signature", len(c.sign))
			}
		})
	}
}

func TestVerifyRejectsInvalidHexSignature(t *testing.T) {
	publicKey, sign := newSignedMessage(t)
	for _, bad := range []string{sign[:127], "zz" + sign[2:]} {
		if verifyNoPanic(t, publicKey, message, bad) {
			t.Fatalf("Verify accepted invalid hex %q", bad)
		}
	}
}
