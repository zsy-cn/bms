package app_test

import (
	"os"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/theplant/gofixtures"
	"github.com/theplant/testingutils"

	"github.com/zsy-cn/bms/protos"
)

var mediaSQLStr = `
insert into sound_box_media(id, created_at, name, customer_id, duration, size, path)
select 			 1, '2019-01-01 12:00:00'::timestamp, '音乐01', 1, 120, 3869451, '/upload/1/music01.mp3'
union all select 2, '2019-01-01 12:00:00'::timestamp, '音乐02', 1, 180, 6235265, '/upload/1/music02.mp3'
;
`

var mediaData = gofixtures.Data(gofixtures.Sql(mediaSQLStr, []string{"sound_box_media"}))

func TestGetSoundBoxMedias(t *testing.T) {
	mediaData.TruncatePut(thedb)

	expected := &protos.GetSoundBoxMediasResponse{
		List: []*protos.SoundBoxMedia{
			{
				ID:         2,
				Name:       "音乐02",
				CustomerID: 1,
				Duration:   "3分0秒",
				Size:       6235265,
				Path:       "/upload/1/music02.mp3",
				CreatedAt:  "2019-01-01",
			},
			{
				ID:         1,
				Name:       "音乐01",
				CustomerID: 1,
				Duration:   "2分0秒",
				Size:       3869451,
				Path:       "/upload/1/music01.mp3",
				CreatedAt:  "2019-01-01",
			},
		},
		Count:       2,
		CurrentPage: 1,
		PageSize:    10,
		TotalCount:  2,
	}

	req := &protos.GetSoundBoxMediasRequest{
		Pagination: &protos.Pagination{
			Page:     1,
			PageSize: 10,
		},
		CustomerID: 1,
	}

	actual, err := ss.GetSoundBoxMedias(req)
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestSaveMediaFile(t *testing.T) {
	mediaData.TruncatePut(thedb)

	var customerID uint64 = 1
	var name string = "music03"
	var path string = "/upload/1/music03.mp3"
	var duration float64 = 210
	var size uint64 = 233333
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fakeNow, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-01-01 12:00:00", loc)
	timenowPatcher := monkey.Patch(time.Now, func() time.Time {
		return fakeNow
	})
	defer timenowPatcher.Unpatch()

	err := ss.SaveMediaFile(customerID, name, path, duration, size)
	if err != nil {
		t.Error(err)
	}

	expected := &protos.GetSoundBoxMediasResponse{
		List: []*protos.SoundBoxMedia{
			{
				Name:       "music03",
				CustomerID: 1,
				Duration:   "3分30秒",
				Size:       233333,
				Path:       "/upload/1/music03.mp3",
				CreatedAt:  "2019-01-01",
			},
		},
		Count:       3,
		CurrentPage: 1,
		PageSize:    1,
		TotalCount:  3,
	}

	req := &protos.GetSoundBoxMediasRequest{
		Pagination: &protos.Pagination{
			Page:     1,
			PageSize: 1,
		},
		CustomerID: 1,
	}

	actual, err := ss.GetSoundBoxMedias(req)
	if err != nil {
		t.Error(err)
	}
	for _, item := range actual.List {
		item.ID = 0
	}
	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestUpdateSoundBoxMedia(t *testing.T) {
	mediaData.TruncatePut(thedb)

	_req := &protos.UpdateSoundBoxMediaRequest{
		ID:         1,
		Name:       "音乐11",
		CustomerID: 1,
	}
	err := ss.UpdateSoundBoxMedia(_req)
	if err != nil {
		t.Error(err)
	}

	expected := &protos.GetSoundBoxMediasResponse{
		List: []*protos.SoundBoxMedia{
			{
				ID:         2,
				Name:       "音乐02",
				CustomerID: 1,
				Duration:   "3分0秒",
				Size:       6235265,
				Path:       "/upload/1/music02.mp3",
				CreatedAt:  "2019-01-01",
			},
			{
				ID:         1,
				Name:       "音乐11",
				CustomerID: 1,
				Duration:   "2分0秒",
				Size:       3869451,
				Path:       "/upload/1/music01.mp3",
				CreatedAt:  "2019-01-01",
			},
		},
		Count:       2,
		CurrentPage: 1,
		PageSize:    10,
		TotalCount:  2,
	}

	req := &protos.GetSoundBoxMediasRequest{
		Pagination: &protos.Pagination{
			Page:     1,
			PageSize: 10,
		},
		CustomerID: 1,
	}

	actual, err := ss.GetSoundBoxMedias(req)
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestDeleteSoundBoxMedia(t *testing.T) {
	mediaData.TruncatePut(thedb)
	osRemove := monkey.Patch(os.Remove, func(name string) error {
		return nil
	})
	defer osRemove.Unpatch()

	err := ss.DeleteSoundBoxMedia(1, 1)
	if err != nil {
		t.Error(err)
	}

	expected := &protos.GetSoundBoxMediasResponse{
		List: []*protos.SoundBoxMedia{
			{
				ID:         2,
				Name:       "音乐02",
				CustomerID: 1,
				Duration:   "3分0秒",
				Size:       6235265,
				Path:       "/upload/1/music02.mp3",
				CreatedAt:  "2019-01-01",
			},
		},
		Count:       1,
		CurrentPage: 1,
		PageSize:    10,
		TotalCount:  1,
	}

	req := &protos.GetSoundBoxMediasRequest{
		Pagination: &protos.Pagination{
			Page:     1,
			PageSize: 10,
		},
		CustomerID: 1,
	}

	actual, err := ss.GetSoundBoxMedias(req)
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}
