package service

import (
	"web_complier/core"
	"web_complier/internal/appserver/store/mysql"
	"web_complier/internal/pkg/response"
	"web_complier/pkg"
)

var ComplierService = &complierService{}

type complierService struct {
}

// Create 创建
func (s *complierService) Create(code, lang string) (gid string, err error) {
	share := &mysql.Share{
		GID:  pkg.GUID(),
		Code: code,
		Type: lang,
	}
	err = mysql.DB().Create(share).Error
	if err != nil {
		core.ZLogger.Sugar().Errorf(err.Error())
		return "", err
	}
	return share.GID, nil
}

func (s *complierService) FindOne(gid string) (rs *response.CodeResponse, err error) {
	share := &mysql.Share{}
	if err := mysql.DB().First(share, "gid = ?", gid).Error; err != nil {
		core.ZLogger.Sugar().Errorf(err.Error())
		return nil, err
	}
	shareRes := &response.CodeResponse{
		Code: share.Code,
		Lang: share.Type,
	}
	return shareRes, nil

}
