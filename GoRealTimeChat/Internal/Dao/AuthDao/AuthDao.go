package AuthDao

import (
	"GoRealTimeChat/DataModels/Internal/DataModels/IMUser"
	"GoRealTimeChat/DataModels/packages/Model"
	"GoRealTimeChat/Internal/Dao/SessionDao"
	"GoRealTimeChat/Internal/Utils"
	"GoRealTimeChat/Packages/Date"
	"GoRealTimeChat/Packages/Hash"
	"GoRealTimeChat/Packages/SetAvatar"
	"fmt"
	"go.uber.org/zap"
)

type AUTHDAO struct {
}

func (*AUTHDAO) CreateUser(email string, password string, name string) int64 {
	createdAt := Date.NewDate()
	hashedPassword, saltCryptoHashPasswordError := Hash.SaltCryptoHashPassword(password)
	zap.S().Errorf("加密密码出错，:%#v", saltCryptoHashPasswordError)
	users := IMUser.ImUsers{
		Email:         email,
		Password:      hashedPassword,
		Name:          name,
		CreatedAt:     createdAt,
		UpdatedAt:     createdAt,
		Avatar:        SetAvatar.GetAvatarBase64Png(fmt.Sprintf("https://api.multiavatar.com/%s.svg", name)),
		LastLoginTime: createdAt,
		Uid:           Utils.GetUuid(),
		UserJson:      "{这名用户还没有任何简介哦}",
		UserType:      1,
	}
	Model.MYSQLDB.Table("im_users").Create(&users)
	var sessionDao SessionDao.SessionDAO
	sessionDao.CreateSession(users.ID, 1, 1)
	sessionDao.CreateSession(1, users.ID, 1)
	return users.ID
}
