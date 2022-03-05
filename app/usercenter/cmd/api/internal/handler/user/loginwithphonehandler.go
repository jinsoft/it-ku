package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"it-ku/app/usercenter/cmd/api/internal/logic/user"
	"it-ku/app/usercenter/cmd/api/internal/svc"
	"it-ku/app/usercenter/cmd/api/internal/types"
)

func LoginWithPhoneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginWithPhoneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewLoginWithPhoneLogic(r.Context(), svcCtx)
		resp, err := l.LoginWithPhone(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
