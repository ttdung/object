package crypto

import (
	"encoding"
	"encoding/base64"
	"os"
	"fmt"
	"crypto/rand"

	cpabe "github.com/cloudflare/circl/abe/cpabe/tkn20"
)

func GenParams() (cpabe.PublicKey, cpabe.SystemSecretKey, error) {
	return cpabe.Setup(rand.Reader)
}
var gmsk = "xAADAAIAURYDyWlxF1QBlZxupo0eskZeG8TjydG6Y9LCuWphMSsAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAARJo39Jk1acNJKyacpHYhsJHt4jinpDmKKP+hpB4/JgQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABBAEEAAIAQWESZbb53LZKz0CJlHv/7pd7BN5znbfMFM54dI3VVP4jUPo4oA4K7WudRtkYLyvpDeHokK4pv07N+lCUwncKOQAz7QaP8XhATCc8AzM2ayA6l8SWCrROp1xg+5iVAUDgYghlIp8w7Nm/ahxvEi7I03IBwJqCUvAnO9twVolEFe5PlBPXEeHMtBV2kQlukSI6XA1Sa9EhgDF0AwhaNm9t1iCPiqOWsme+a71mEhhYHrxUBFefefMPtuDNay9M3C26EV1wNWVVi4IU+mdqs1FnBGsSxw5/NrKM2+MTbx09q6cE2+EZmIuDZEH+GNBnlmK6WbfeAYszhxumUWmdrRZEPQQBBAACADkOSSx0yItSjcRzsvilSb0YEpJNOBsKYpXctuyUedpcOuT/XZJ0Ox+ChR1PEW1DD2LtXYIaHW2TfFC1AplEmdNEa5VcS2W8d1GLmWDi/TXCBW16ojU2bUqJVuXrD0GxdGkWcQ2X0rNloXQLSdfcuk+TFCLytiTOqJfY+KUhwtwaaiKtMXHoNGB53HWiwB3a8eiH/mBB0ZajPDyiDo2T9cRlIrdvFXzCydUTCu7OuyQ9X9+heLNk9b7nxgf2Q4vX6XIsqwm6VQ2FDk6Gij4wrLC9GqB445+lBodDoMZukKohcr+rM7bCXSCyqhYDZqWINA62bhG13T/3bppL7L3FdvQEAQQAAgADNGjcEenwLEv+GN/Jbe+UN5vm/VobqPeVmnFm+K/DHmn74FuvlKcNXb5UCPeUv0ye4WBo+uoUSpHn2u4P+9lUVTmCgBrlJx5DRGgdQHa9MYgLJ/fT4WMxAK9ufWlFuT5Kolrz6gHh6NEBxZGx0msaPd+PR3TvCN9lBhwsWWlmDTMc1TrNb0JNhhGJGMjDuZ70Yjb35VM6RImJ/rOEbBB8RiUpply+Uye8fnzMDtM97Jr7Pg2kxLhc3JCBJBHNeuBjzi+2gxI7nsBYjQOqRx3GFgxb7GPBtVuOwwk638HixmSf9nlTGCmCzsVLagPVylk/1luVEOUKF/D9g/5eNtZ7hAAEAAEABNdXedrwMDfNkFmc3QAEWf8ZE5Nz5dEE9TvrZDD60NxFCZdb+M3UGZT4ObLNLGkOi5YvmmdoUShUASlCnvRwdh6ub1FYGt7VOPWaDBj4RlgkHax7Owj8sU684SLWw1z5WovERlf9uss10UC0DWXGCkxSjIdv8zkPt2SBpD9/be8QANgvef5lcp+5/itfwhZQLsY="
var gmpk = "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"

func KeyGen(msk cpabe.SystemSecretKey, attr cpabe.Attributes) (cpabe.AttributeKey, error) {
	return msk.KeyGen(rand.Reader, attr)
}

func GenerateABEKey(in map[string]string) (string, error) {
	attr := cpabe.Attributes{}
	attr.FromMap(in)

	keybyte, err := base64.StdEncoding.DecodeString(gmsk)
	if err != nil {
		return "", err
	}

	msk := cpabe.SystemSecretKey{}
	err = msk.UnmarshalBinary(keybyte)

	abekey, err := msk.KeyGen(rand.Reader, attr)
	if err != nil {
		return "", err
	}

	key, err := MarshalUserKey(abekey)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(key), nil
}

func WriteToFile(name string, data []byte) {
	err := os.WriteFile(name, data, 0o644)
	if err != nil {
		panic(err)
	}
}

func ExportToFile(name string, m encoding.BinaryMarshaler) {
	data, err := m.MarshalBinary()
	if err != nil {
		panic(err)
	}
	WriteToFile(name, data)
}

func NewUserKey(keystr string) (cpabe.AttributeKey, error) {
	key := cpabe.AttributeKey{}

	keybyte, err := base64.StdEncoding.DecodeString(keystr)
	if err != nil {
		return key, err
	}

	return UnmarshalUserKey(keybyte)
}

func ImportUserKeyFromFile(name string) (cpabe.AttributeKey, error) {
	sk := cpabe.AttributeKey{}

	attributeKey, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Unable to read filename ", name)
		return sk, err
	}

	err = sk.UnmarshalBinary(attributeKey)

	return sk, err
}

func MarshalUserKey(attributeKey cpabe.AttributeKey) ([]byte, error) {
	return attributeKey.MarshalBinary()
}

func UnmarshalUserKey(attributeKey []byte) (cpabe.AttributeKey, error) {
	sk := cpabe.AttributeKey{}

	err := sk.UnmarshalBinary(attributeKey)

	return sk, err
}

func MarshalMPK(key cpabe.PublicKey) ([]byte, error) {
	return key.MarshalBinary()
}

func UnmarshalMPK(mpk []byte) (cpabe.PublicKey, error) {
	pk := cpabe.PublicKey{}
	err := pk.UnmarshalBinary(mpk)

	return pk, err
}

func NewMPK(keystr string) (cpabe.PublicKey, error) {
	key := cpabe.PublicKey{}

	keybyte, err := base64.StdEncoding.DecodeString(keystr)
	if err != nil {
		return key, err
	}

	return UnmarshalMPK(keybyte)
}

func MPKToString(key cpabe.PublicKey) (string, error) {
	kbyte, err := key.MarshalBinary()

	return base64.StdEncoding.EncodeToString(kbyte), err
}

func MarshalMSK(key cpabe.SystemSecretKey) ([]byte, error) {
	return key.MarshalBinary()
}

func UnmarshalMSK(msk []byte) (cpabe.SystemSecretKey, error) {
	sk := cpabe.SystemSecretKey{}
	err := sk.UnmarshalBinary(msk)
	
	return sk, err
}

func MSKToString(key cpabe.SystemSecretKey) (string, error) {
	kbyte, err := key.MarshalBinary()

	return base64.StdEncoding.EncodeToString(kbyte), err
}

func NewMSK(keystr string) (cpabe.SystemSecretKey, error) {
	key := cpabe.SystemSecretKey{}

	keybyte, err := base64.StdEncoding.DecodeString(keystr)
	if err != nil {
		return key, err
	}

	return UnmarshalMSK(keybyte)
}

func ImportMPKFromFile(name string) (cpabe.PublicKey, error) {
	pk := cpabe.PublicKey{}

	key, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Unable to read filename ", name)
		return pk, err
	}

	err = pk.UnmarshalBinary(key)

	return pk, err
}

func ImportMSKFromFile(name string) (cpabe.SystemSecretKey, error) {
	sk := cpabe.SystemSecretKey{}

	key, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Unable to read filename ", name)
		return sk, err
	}


	err = sk.UnmarshalBinary(key)
	if err != nil {
		fmt.Println("unable to parse public key")
		return sk, err
	}

	return sk, nil
}

func Encrypt(pk cpabe.PublicKey, policy cpabe.Policy, msg []byte) ([]byte, error){
	return pk.Encrypt(rand.Reader, policy, msg)
}

func Decrypt(key cpabe.AttributeKey, ct []byte) ([]byte, error){
	return key.Decrypt(ct)
}

func AbeEncrypt(mpkstr string, accesstree string, msg string) (string, error) {
	mpk, err := NewMPK(mpkstr)
	if err != nil {
		return "", err
	}

	policy := cpabe.Policy{}
	err = policy.FromString(accesstree)
	if err != nil {
		return "", err
	}

	msgbyte, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return "", err
	}

	ctbyte, err := Encrypt(mpk, policy, msgbyte)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(ctbyte), nil
}

func AbeDecrypt(mpkstr string, idx string, ekey string, ct string) (string, error) {
	key, err := NewUserKey(ekey)
	if err != nil {
		return "", err
	}

	ctbyte, err := base64.StdEncoding.DecodeString(ct)
	if err != nil {
		return "", err
	}

	ptbyte, err := Decrypt(key, ctbyte)

	return base64.StdEncoding.EncodeToString(ptbyte), nil
}

func mapToString(m map[string]string) string {
    var sb strings.Builder
    for key, value := range m {
        sb.WriteString(fmt.Sprintf("%s=%s ", key, value))
    }
    return sb.String()
}

func stringToMap(s string) map[string]string {
    m := make(map[string]string)
    pairs := strings.Fields(s)
    for _, pair := range pairs {
        kv := strings.SplitN(pair, "=", 2)
        if len(kv) == 2 {
            m[kv[0]] = kv[1]
        }
    }
    return m
}