package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

type MyClaims struct {
	UserID   int64
	Username string
	jwt.StandardClaims
}

var jwtSecret = []byte("sakura")

func GenerateToken(userID int64, username string) (aToken string, rToken string, err error) {
	//设置token有效时间
	nowTime := time.Now()
	accessExpireTime := nowTime.Add(ATokenExistTime)
	refreshExpireTime := nowTime.Add(RTokenExistTime)

	accessClaims := MyClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: accessExpireTime.Unix(),
			// 指定token发行人
			Issuer: viper.GetString("name"),
		},
	}

	refreshClaims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: refreshExpireTime.Unix(),
			// 指定token发行人
			Issuer: viper.GetString("name"),
		},
	}

	accessTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	aToken, err = accessTokenClaims.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}
	rToken, err = refreshTokenClaims.SignedString(jwtSecret)
	return aToken, rToken, err
}

func ParseToken(token string) (*MyClaims, error) {
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func RefreshToken(userID int64) (string, error) {
	//设置token有效时间
	nowTime := time.Now()
	accessExpireTime := nowTime.Add(ATokenExistTime)
	accessClaims := MyClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: accessExpireTime.Unix(),
			// 指定token发行人
			Issuer: viper.GetString("name"),
		},
	}
	accessTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	aToken, err := accessTokenClaims.SignedString(jwtSecret)
	return aToken, err
}
