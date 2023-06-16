package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/kyosu-1/ims/gen/api"
)

func (h *Handlers) GetMe(ec echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handlers) PostSignin(ec echo.Context) error {
	req := &api.UserSignin{}
	if err := ec.Bind(req); err != nil {
		return err
	}
	ctx := ec.Request().Context()

	// リポジトリからユーザーを取得
	user, err := h.userRepository.FindUserByID(ctx, req.Id)
	if err != nil {
		return err
	}

	// パスワードが一致するか確認
	if !user.IsSamePassword(req.Password) {
		return echo.ErrUnauthorized
	}

	// storeにsessionを保存
	sess, err := h.sessionStore.Get(ec.Request(), "session")
	if err != nil {
		return err
	}

	sess.Values["user_id"] = user.ID
	if err := sess.Save(ec.Request(), ec.Response()); err != nil {
		return err
	}

	return ec.NoContent(204)
}

func (h *Handlers) PostSignout(ec echo.Context) error {
	// TODO: implement
	return nil
}
