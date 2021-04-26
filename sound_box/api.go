package sound_box

import (
	"github.com/zsy-cn/bms/protos"
)

type SoundBoxService interface {
	GetDevice(customerID uint64, deviceSN string) (device *protos.SoundBoxDevice, err error)
	GetDevices(req *protos.GetDevicesRequestForCustomer) (devices *protos.SoundBoxDeviceList, err error)
	GetSoundBoxDeviceStatus(customerID uint64, deviceSN string) (status int8, err error)
	GetSoundBoxDeviceGroups(req *protos.GetDevicesRequestForCustomer) (resp *protos.SoundBoxDeviceGroupList, err error)

	SaveMediaFile(customerID uint64, filename string, path string, duration float64, size uint64) (err error)
	GetSoundBoxMedias(req *protos.GetSoundBoxMediasRequest) (resp *protos.GetSoundBoxMediasResponse, err error)
	UpdateSoundBoxMedia(req *protos.UpdateSoundBoxMediaRequest) (err error)
	DeleteSoundBoxMedia(id, customerID uint64) (err error)
}
