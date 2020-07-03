package entity

type UserReadHistory struct {
	/**
	用户id
	*/
	Id int `json:"id,omitempty"`

	/**
	新闻id
	*/
	LinkId int `json:"linkId,omitempty"`
}

func (this *UserReadHistory) GetId() int {
	return this.Id
}

func (this *UserReadHistory) getLinkId() int {
	return this.LinkId
}
