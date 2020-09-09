package filter

import (
	"github.com/dgrijalva/jwt-go"
	"gocherry-api-gateway/proxy/enum"
)

type JwtFilter struct {
	Filter
}

func (f *JwtFilter) Init(proxyContext *ProxyContext) {
}

func (f *JwtFilter) Name(proxyContext *ProxyContext) string {
	return Jwt
}

func (f *JwtFilter) Pre(proxyContext *ProxyContext) (statusCode int, err string) {
	if proxyContext.Api.JwtAuth == false {
		return enum.STATUS_CODE_OK, ""
	}

	jwtToken := proxyContext.RequestContext.GetHeader("jwt_token")
	token, errJwt := jwt.Parse(jwtToken, secretFunc(proxyContext.AppConfig.Common.JwtSecret))
	if errJwt != nil {
		return enum.STATUS_CODE_FAILED, "Jwt授权失败"
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		proxyContext.UserInfo.UserId = claims["user_id"].(string)
		proxyContext.UserInfo.UserName = claims["user_name"].(string)
		return enum.STATUS_CODE_OK, ""
	} else {
		return enum.STATUS_CODE_FAILED, "获取用户信息失败"
	}
}

func (f *JwtFilter) AfterDo(proxyContext *ProxyContext) (statusCode int, err string) {
	return enum.STATUS_CODE_OK, ""
}

//解析
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}

/**
生成jwt token的例子 应该在其他服务或接口生成对应用户的token
*/
func CreateToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   "sgg",
		"user_name": "jjjjj",
	})
	//可以从配置里读取
	tokenString, _ := token.SignedString([]byte("chris"))
	return tokenString
}
