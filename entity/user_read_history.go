package entity

type UserReadHistory struct {
	/**
	userId
	*/
	Id int64 `json:"id,omitempty"`

	/**
	newsId
	*/
	LinkId int `json:"linkId,omitempty"`
}

func (this *UserReadHistory) GetId() int64 {
	return this.Id
}

func (this *UserReadHistory) getLinkId() int {
	return this.LinkId
}
