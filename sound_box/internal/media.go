package internal

import (
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/sound_box"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/util/media"
	"github.com/zsy-cn/bms/util/pagination"
)

// SaveMediaFile 为已经保存到服务器的媒体文件创建存储记录
func SaveMediaFile(db *gorm.DB, log log.Logger, customerID uint64, name string, path string, duration float64, size uint64) (err error) {
	file := &sound_box.SoundBoxMedia{
		Name:       name,
		Path:       path,
		CustomerID: customerID,
		Duration:   duration,
		Size:       size,
	}
	err = db.Create(file).Error
	if err != nil {
		log.Errorf("create media file: %s failed in SaveMediaFile(): %s", name, err.Error())
		return
	}
	return
}

func GetSoundBoxMedias(db *gorm.DB, log log.Logger, req *protos.GetSoundBoxMediasRequest) (resp *protos.GetSoundBoxMediasResponse, err error) {
	resp = &protos.GetSoundBoxMediasResponse{
		List: []*protos.SoundBoxMedia{},
	}
	query := db.Model(&sound_box.SoundBoxMedia{})
	// 条件查询语句
	whereArgs := map[string]interface{}{
		"customer_id": req.CustomerID,
	}

	query = query.Where(whereArgs)
	var count uint64
	err = query.Count(&count).Error
	if err != nil {
		log.Errorf("get customer: %d's media count failed in GetSoundBoxMedias(): %s", req.CustomerID, err.Error())
		return
	}
	// 构建分页查询语句
	query = pagination.BuildPaginationQuery(query, req.Pagination)
	mediaRecords := []*sound_box.SoundBoxMedia{}
	err = query.Find(&mediaRecords).Error
	if err != nil {
		if err.Error() == "record not found" {
			err = nil
		} else {
			log.Errorf("get customer: %d's groups failed in GetSoundBoxMedias(): %s", req.CustomerID, err.Error())
			return
		}
	}

	for _, mediaRecord := range mediaRecords {
		mediaPB := &protos.SoundBoxMedia{}
		_ = mediaRecord2PB(mediaRecord, mediaPB)
		resp.List = append(resp.List, mediaPB)
	}
	resp.Count = count
	resp.TotalCount = count
	resp.CurrentPage = req.Pagination.Page
	resp.PageSize = req.Pagination.PageSize
	return
}

func UpdateSoundBoxMedia(db *gorm.DB, log log.Logger, req *protos.UpdateSoundBoxMediaRequest) (err error) {
	mediaRecord := &sound_box.SoundBoxMedia{}
	whereArgs := map[string]interface{}{
		"id":          req.ID,
		"customer_id": req.CustomerID,
	}

	err = db.Where(whereArgs).First(&mediaRecord).Error
	if err != nil {
		log.Errorf("get media record failed in UpdateSoundBoxMedia(): %s", err.Error())
		return
	}

	err = db.Model(mediaRecord).UpdateColumn("name", req.Name).Error
	if err != nil {
		log.Errorf("update media info failed in UpdateSoundBoxMedia(): %s", err.Error())
		return
	}
	return
}

func DeleteSoundBoxMedia(db *gorm.DB, log log.Logger, id uint64, customerID uint64) (err error) {
	mediaRecord := &sound_box.SoundBoxMedia{}
	whereArgs := map[string]interface{}{
		"id":          id,
		"customer_id": customerID,
	}
	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	err = tx.Where(whereArgs).First(mediaRecord).Error
	if err != nil {
		log.Errorf("find customer: %d's media failed in DeleteSoundBoxMedia(): %s", customerID, err.Error())
		return
	}
	mediaPath := mediaRecord.Path
	err = tx.Delete(mediaRecord).Error
	if err != nil {
		log.Errorf("delete media file: %s failed in DeleteSoundBoxMedia(): %s", mediaRecord.Name, err.Error())
		return
	}
	err = os.Remove(mediaPath)

	return
}

func mediaRecord2PB(record *sound_box.SoundBoxMedia, pb *protos.SoundBoxMedia) (err error) {
	pb.ID = record.ID
	pb.Name = record.Name
	pb.CustomerID = record.CustomerID
	pb.Duration = media.FormatDuration(record.Duration)
	pb.Size = record.Size
	pb.Path = record.Path

	loc, _ := time.LoadLocation("Asia/Shanghai")
	pb.CreatedAt = record.CreatedAt.In(loc).Format("2006-01-02")
	return
}
