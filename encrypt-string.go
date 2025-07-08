package eliotlibs

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
)

// Función para encriptar un mensaje con AES-256-GCM
func Encrypt(secret, base64Key string) (string, error) {
	// Decodificar la clave base64
	key, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		return "", err
	}

	// Crear un nuevo bloque AES
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Crear un nuevo cifrador GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Crear un nonce aleatorio de 12 bytes
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	// Cifrar el mensaje
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(secret), nil)

	// Convertir a base64 para almacenar
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

type InvalidEncryptData struct{}

func (InvalidEncryptData) Error() string {
	return "INVALID_ENCRYPT_DATA"
}

// Función para desencriptar un mensaje con AES-256-GCM
func Decrypt(encryptedData, base64Key string) (string, error) {
	// Decodificar la clave base64
	key, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		return "", err
	}

	// Decodificar los datos cifrados base64
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	// Crear un nuevo bloque AES
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Crear un nuevo cifrador GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Extraer el nonce (primeros 12 bytes)
	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", InvalidEncryptData{}
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Descifrar los datos
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
