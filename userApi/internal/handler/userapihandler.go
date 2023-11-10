package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"userApi/internal/logic"
	"userApi/internal/svc"
	"userApi/internal/types"
)

func UserApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserApiLogic(r.Context(), svcCtx)
		resp, err := l.UserApi(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
