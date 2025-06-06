package nanoid

import nanoid "github.com/matoous/go-nanoid/v2"

func Generate() string {
	id, _ := nanoid.New()
	return id
}

func GenerateId() string {
	return nanoid.MustGenerate("abcdefghijklmnopqrstuvwxyz0123456789", 21)
}
