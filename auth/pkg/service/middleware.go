package service

// Middleware describes a service middleware.
type Middleware func(AuthService) AuthService

// func NewParser(keyfunc jwt.Keyfunc) endpoint.Middleware {
// 	return func(next endpoint.Endpoint) endpoint.Endpoint {
// 		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

// 			tokenString, ok := ctx.Value("JWT").(string)
// 			if !ok {
// 				return nil, errors.New("ErrTokenContextMissing")
// 			}

// 			token, err := jwt.ParseWithClaims(tokenString, &model.Token{}, func(token *jwt.Token) (interface{}, error) {
// 				if token.Method != jwt.GetSigningMethod("HS256") {
// 					return nil, errors.New("ErrUnexpectedSigningMethod")
// 				}

// 				return keyfunc(token)
// 			})

// 			if err != nil {
// 				if e, ok := err.(*jwt.ValidationError); ok {
// 					switch {
// 					case e.Errors&jwt.ValidationErrorMalformed != 0:
// 						// Token is malformed
// 						return nil, errors.New("ErrTokenMalformed")
// 					case e.Errors&jwt.ValidationErrorExpired != 0:
// 						// Token is expired
// 						return nil, errors.New("ErrTokenExpired")
// 					case e.Errors&jwt.ValidationErrorNotValidYet != 0:
// 						// Token is not active yet
// 						return nil, errors.New("ErrTokenNotActive")
// 					case e.Inner != nil:
// 						// report e.Inner
// 						return nil, e.Inner
// 					}
// 					// We have a ValidationError but have no specific Go kit error for it.
// 					// Fall through to return original error.
// 				}
// 				return nil, err
// 			}

// 			if !token.Valid {
// 				return nil, errors.New("ErrTokenInvalid")
// 			}

// 			ctx = context.WithValue(ctx, "JWT", token.Claims)

// 			return next(ctx, request)
// 		}
// 	}
// }
