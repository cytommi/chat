package config

import _ "embed"

//go:embed public-key.pem
var JwtPublicPEM []byte

//go:embed private-key.pem
var JwtPrivateKeyPEM []byte
