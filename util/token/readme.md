# reference link
- https://dev.to/techschoolguru/how-to-create-and-verify-jwt-paseto-token-in-golang-1l5j
- https://pkg.go.dev/github.com/golang-jwt/jwt#section-readme
# What the heck is a JWT?
  it's a signed JSON object that does something useful (for example, authentication).

  A token is made of three parts, separated by .'s. The first two parts are JSON objects, that have been base64url encoded. The last part is the signature, encoded the same way.

  The first part is called the header. It contains the necessary information for verifying the last part, the signature. For example, which encryption method was used for signing and what key was used.

  The part in the middle is the interesting bit. It's called the Claims and contains the actual stuff you care about. Refer to RFC 7519 for information about reserved keys and the proper way to add your own.