package models

type Group struct {
	JsonEmbeddable
	ID string `json:"id" validate:"-" sql:"id"`
	GroupName string `json:"group_name" validate:"required" sql:"group_name"`
	Bio string `json:"group_bio" validate:"required" sql:"group_bio"`
	NumGroupMembers int32 `json:"num_group_members" validate:"required" sql:"num_group_members"`
}