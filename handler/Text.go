package handler

import (
	"context"
	"ogm-config/model"

	"github.com/asim/go-micro/v3/logger"
	proto "github.com/xtech-cloud/ogm-msp-config/proto/config"
)

type Text struct{}

func (this *Text) Write(_ctx context.Context, _req *proto.TextWriteRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Text.Write, req is %v", _req)
	_rsp.Status = &proto.Status{}

	if "" == _req.Path {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "path is required"
		return nil
	}

	if "" == _req.Content {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "content is required"
		return nil
	}

	dao := model.NewTextDAO(nil)
	entity := &model.Text{
		UUID:    model.ToUUID(_req.Path),
		Path:    _req.Path,
		Content: _req.Content,
	}

	err := dao.Upsert(entity)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Uuid = entity.UUID
	return nil
}

func (this *Text) Read(_ctx context.Context, _req *proto.TextReadRequest, _rsp *proto.TextReadResponse) error {
	logger.Infof("Received Text.Read, req is %v", _req)
	_rsp.Status = &proto.Status{}

	if "" == _req.Path {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "path is required"
		return nil
	}

	dao := model.NewTextDAO(nil)
	entity, err := dao.FindByPath(_req.Path)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Entity = &proto.TextEntity{
		Uuid:    entity.UUID,
		Path:    entity.Path,
		Content: entity.Content,
	}

	return nil
}

func (this *Text) Delete(_ctx context.Context, _req *proto.DeleteRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Text.Delete, req is %v", _req)
	_rsp.Status = &proto.Status{}

	if "" == _req.Uuid {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "uuid is required"
		return nil
	}

	dao := model.NewTextDAO(nil)
	err := dao.Delete(_req.Uuid)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Uuid = _req.Uuid

	return nil
}

func (this *Text) Get(_ctx context.Context, _req *proto.GetRequest, _rsp *proto.TextGetResponse) error {
	logger.Infof("Received Text.Get, req is %v", _req)
	_rsp.Status = &proto.Status{}

	if "" == _req.Uuid {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "uuid is required"
		return nil
	}

	dao := model.NewTextDAO(nil)
	entity, err := dao.Get(_req.Uuid)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Entity = &proto.TextEntity{
		Uuid:    entity.UUID,
		Path:    entity.Path,
		Content: entity.Content,
	}

	return nil
}

func (this *Text) List(_ctx context.Context, _req *proto.ListRequest, _rsp *proto.TextListResponse) error {
	logger.Infof("Received Text.List, req is %v", _req)
	_rsp.Status = &proto.Status{}

	offset := int64(0)
	if _req.Offset > 0 {
		offset = _req.Offset
	}

	count := int64(0)
	if _req.Count > 0 {
		count = _req.Count
	}

	dao := model.NewTextDAO(nil)
	total, entity, err := dao.List(offset, count)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Total = total
	_rsp.Entity = make([]*proto.TextEntity, len(entity))
	for i, e := range entity {
		_rsp.Entity[i] = &proto.TextEntity{
			Uuid:    e.UUID,
			Path:    e.Path,
			Content: e.Content,
		}
	}

	return nil
}

func (this *Text) Search(_ctx context.Context, _req *proto.TextSearchRequest, _rsp *proto.TextSearchResponse) error {
	logger.Infof("Received Text.Search, req is %v", _req)
	_rsp.Status = &proto.Status{}

	offset := int64(0)
	if _req.Offset > 0 {
		offset = _req.Offset
	}

	count := int64(0)
	if _req.Count> 0 {
		count = _req.Count
	}

	dao := model.NewTextDAO(nil)
	total, entity, err := dao.Search(offset, count, _req.Path)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Total = total
	_rsp.Entity = make([]*proto.TextEntity, len(entity))
	for i, e := range entity {
		_rsp.Entity[i] = &proto.TextEntity{
			Uuid:    e.UUID,
			Path:    e.Path,
			Content: e.Content,
		}
	}

	return nil
}
