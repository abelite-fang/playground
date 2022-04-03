package responses

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlResponse struct {
	Username   string             `json:"userName"`
	UserId     primitive.ObjectID `json:"userId"`
	ShortId    primitive.ObjectID `json:"shortId"`
	ShortUrl   string             `json:"shortUrl"`
	ReqestDate time.Time          `json:"requestDate"`
	CreateDate time.Time          `json:"createDate"`
	ExpireDate time.Time          `json:"expireDate"`
}
