package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

const TokenExpireTime = time.Hour * 2

var mySecret = []byte("Hello,Piwriw")

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID int64, username string) (string, error) {
	// 声明一个JWT
	c := MyClaims{
		userID,
		"username", // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(), // 过期时间
			Issuer:    "my-project",                                                                      // 签发人
		},
	}
	// 使用指定的签名方法创建签发对象
	token := jwt.NewWithClaims(jwt.SigningMethodES256,  c)
	// 使用指定secret签名获得完整编码之后的字符串token
	return token.SignedString(mySecret)

}

// ParseToken : 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// GenTokenWithRefreshToken : 生成Access token 和 RefreshToken
func GenTokenWithRefreshToken(userID int64) (acToken, reToken string, err error) {
	c := MyClaims{
		userID,
		"username",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireTime).Unix(), //签发日期
			Issuer:    "Piwirw",                               //签发人
		},
	}
	// 加密并获得编码Token
	acToken, err = jwt.NewWithClaims(jwt.SigningMethodES256, c).SignedString(mySecret)

	// refresh token 不需要存储子自定义数据
	reToken, err = jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(), // 过期时间
		Issuer:    "Piwriw",                                //签发人
	}).SignedString(mySecret)
	// 使用自定secret签名并获得完整后的Token
	return
}

// ParseToken2 : 解析Access Token
func ParseToken2(tokenString string)(claims *MyClaims,err error)  {
	// 解析tokenn
	var token *jwt.Token
	claims=new(MyClaims)
	token,err=jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return
	}
	if !token.Valid {
		err=errors.New("Invalid token")
	}
	return
}

// RefreshToken : 刷新RefreshToken
func RefreshToken(acToken, reToken string) (newToken, newReToken string, err error) {
	// refresh token is invalid returnn
	_,err=jwt.Parse(reToken,func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err!=nil{
		return
	}
	// 从 older access token 解析出claims
	var claims MyClaims
	_,err=jwt.ParseWithClaims(acToken,&claims,func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	v,_:=err.(*jwt.ValidationError)
	// 当access token is 过期错误并且 refresh token 没有过期 ->>>create a  new access token
	if v.Errors==jwt.ValidationErrorExpired{
		return GenTokenWithRefreshToken(claims.UserID)
	}
	return
}