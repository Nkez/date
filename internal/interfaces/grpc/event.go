package grpc

import (
	"context"
	date_protobuf "github.com/Nkez/date-protobuf"
	"github.com/Nkez/date/internal/model"
	"github.com/Nkez/date/internal/repositories"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventApiStruct struct {
	date_protobuf.UnimplementedEventServiceServer
	repository repositories.Event
}

func NewEventApiStruct(repository repositories.Event) *EventApiStruct {
	return &EventApiStruct{repository: repository}
}

func (s EventApiStruct) Get(_ context.Context, id *date_protobuf.GetEvent) (*date_protobuf.Event, error) {
	event, err := s.repository.Get(id.GetId())
	if err != nil {
		return nil, err
	}
	protoEvent := &date_protobuf.Event{
		Id:          event.ID,
		TypeRequest: event.TypeRequest,
		Browser:     event.Browser,
		Os:          event.OS,
		Device:      event.Device,
		City:        event.City,
		Country:     event.Country,
		CreatedAt:   timestamppb.New(event.CreatedAt),
	}
	return protoEvent, nil
}

func (s EventApiStruct) List(_ context.Context, input *date_protobuf.FilterEvent) (*date_protobuf.EventList, error) {
	filter := encodeQRFilter(input)
	events, err := s.repository.List(filter)
	if err != nil {
		return nil, err
	}
	response := &date_protobuf.EventList{
		Event: make([]*date_protobuf.Event, 0, len(events)),
	}
	for _, e := range events {
		response.Event = append(response.Event, decodeEvent(e))
	}
	return response, nil
}

func decodeEvent(event *model.Event) *date_protobuf.Event {
	return &date_protobuf.Event{
		Id:          event.ID,
		TypeRequest: event.TypeRequest,
		Browser:     event.Browser,
		Os:          event.OS,
		Device:      event.Device,
		City:        event.City,
		Country:     event.Country,
		CreatedAt:   timestamppb.New(event.CreatedAt),
	}
}

func encodeQRFilter(input *date_protobuf.FilterEvent) *model.Filter {
	filter := &model.Filter{
		Page: nil,
		Size: nil,
	}
	if input.GetPageNumber() != nil {
		filter.Page = &input.GetPageNumber().Value
	}
	if input.GetPageSize() != nil {
		filter.Page = &input.GetPageSize().Value
	}
	return filter
}
