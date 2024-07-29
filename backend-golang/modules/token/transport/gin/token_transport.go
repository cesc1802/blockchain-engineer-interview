package tokengin

import (
	"errors"
	"net/http"

	"backend-golang/common"
	"backend-golang/component/appctx"
	tokenbizs "backend-golang/modules/token/biz"
	tokenmodel "backend-golang/modules/token/model"

	"github.com/gin-gonic/gin"
)

func CreateToken(appCtx appctx.AppContext, tokenBiz tokenbizs.TokenBiz) gin.HandlerFunc {
	return func(c *gin.Context) {
		saveSessionDTO := tokenmodel.SaveSessionDTO{}
		err := c.ShouldBind(&saveSessionDTO)
		if err != nil {
			panic(common.ErrInternal(err))
		}
		if saveSessionDTO.DocID == nil {
			panic(common.ErrInternal(errors.New("no doc_id")))
		}

		err = tokenBiz.UploadData(c.Request.Context(), *saveSessionDTO.DocID)
		if err != nil {
			panic(common.ErrInternal(err))
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func Confirm(appCtx appctx.AppContext, tokenBiz tokenbizs.TokenBiz) gin.HandlerFunc {
	return func(c *gin.Context) {
		var sessionDTO tokenmodel.SaveSessionDTO
		err := c.ShouldBind(&sessionDTO)
		if err != nil {
			panic(common.ErrInternal(err))
		}

		err = tokenBiz.Confirm(c.Request.Context(), *sessionDTO.DocID, sessionDTO.ContentHash, sessionDTO.Proof, sessionDTO.SessionId, sessionDTO.RiskScore)
		if err != nil {
			panic(common.ErrInternal(err))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func GetListUploadSessions(appCtx appctx.AppContext, tokenBiz tokenbizs.TokenBiz) gin.HandlerFunc {
	return func(c *gin.Context) {
		uploadSessions := tokenBiz.GetListSession(c.Request.Context())
		c.JSON(http.StatusOK, common.NewSuccessResponse(uploadSessions, nil, nil))
	}
}
