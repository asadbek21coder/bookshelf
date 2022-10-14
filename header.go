package bookshelf

type Header struct {
	Key  *string `header:"key" binding:"required"`
	Sign *string `header:"sign" binding:"required"`
}
