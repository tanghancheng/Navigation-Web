package dto

type NoteQueryDto struct {
	PageInfo
	EditStaus uint `json:"edit_status"`
}

func NewNoteNoteQueryDto() *NoteQueryDto {
	noteQueryDto := NoteQueryDto{
		PageInfo: *NewPageInfo(),
	}
	return &noteQueryDto
}