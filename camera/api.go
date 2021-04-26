package camera

type CameraService interface {
	GetAccessToken() (accessToken string, err error)
	GetMainScreen() (cameraID string, err error)
	SetMainScreen(cameraID string) (err error)
	RequestAccessToken()(accessToken string, err error)
}
