package service

import "context"

func (m *service) Auth(ctx context.Context, uid, uuid int64, token, sdkVersion, deviceId, platform, model, system string) error {
	_, err := m.repo.SelectUserByToken(ctx, uid, token)
	if err != nil {
		return err
	}
	err = m.repo.UpdateUserById(ctx, uid, uuid)
	if err != nil {
		return err
	}
	return nil
}
