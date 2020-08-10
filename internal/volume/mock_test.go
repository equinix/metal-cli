package volume

import "github.com/packethost/packngo"

type MockVolumeService struct {
	ListFn   func(string, *packngo.ListOptions) ([]packngo.Volume, *packngo.Response, error)
	GetFn    func(string, *packngo.GetOptions) (*packngo.Volume, *packngo.Response, error)
	UpdateFn func(string, *packngo.VolumeUpdateRequest) (*packngo.Volume, *packngo.Response, error)
	DeleteFn func(string) (*packngo.Response, error)
	CreateFn func(*packngo.VolumeCreateRequest, string) (*packngo.Volume, *packngo.Response, error)
	LockFn   func(string) (*packngo.Response, error)
	UnlockFn func(string) (*packngo.Response, error)
}

type MockVolumeAttachmentService struct {
	GetFn    func(string, *packngo.GetOptions) (*packngo.VolumeAttachment, *packngo.Response, error)
	CreateFn func(string, string) (*packngo.VolumeAttachment, *packngo.Response, error)
	DeleteFn func(string) (*packngo.Response, error)
}

func (mock *MockVolumeService) List(projectID string, opts *packngo.ListOptions) ([]packngo.Volume, *packngo.Response, error) {
	return mock.ListFn(projectID, opts)
}
func (mock *MockVolumeService) Get(volumeID string, opts *packngo.GetOptions) (*packngo.Volume, *packngo.Response, error) {
	return mock.GetFn(volumeID, opts)
}
func (mock *MockVolumeService) Update(volumeID string, req *packngo.VolumeUpdateRequest) (*packngo.Volume, *packngo.Response, error) {
	return mock.UpdateFn(volumeID, req)
}
func (mock *MockVolumeService) Delete(volumeID string) (*packngo.Response, error) {
	return mock.DeleteFn(volumeID)
}
func (mock *MockVolumeService) Create(req *packngo.VolumeCreateRequest, projectID string) (*packngo.Volume, *packngo.Response, error) {
	return mock.CreateFn(req, projectID)
}

func (mock *MockVolumeService) Lock(volumeID string) (*packngo.Response, error) {
	return mock.LockFn(volumeID)
}
func (mock *MockVolumeService) Unlock(volumeID string) (*packngo.Response, error) {
	return mock.UnlockFn(volumeID)
}

func (mock *MockVolumeAttachmentService) Get(attachmentID string, opts *packngo.GetOptions) (*packngo.VolumeAttachment, *packngo.Response, error) {
	return mock.GetFn(attachmentID, opts)
}
func (mock *MockVolumeAttachmentService) Delete(attachmentID string) (*packngo.Response, error) {
	return mock.DeleteFn(attachmentID)
}
func (mock *MockVolumeAttachmentService) Create(volumeID, deviceID string) (*packngo.VolumeAttachment, *packngo.Response, error) {
	return mock.CreateFn(volumeID, deviceID)
}

var (
	_ packngo.VolumeService           = &MockVolumeService{}
	_ packngo.VolumeAttachmentService = &MockVolumeAttachmentService{}
)
