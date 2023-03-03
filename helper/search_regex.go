package helper

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSerchRegex(keys []string, pattern string) (res []bson.M) {
	for _, val := range keys {
		each := bson.M{
			val: bson.M{
				"$regex": primitive.Regex{
					Pattern: pattern,
					Options: "i",
				},
			},
		}

		res = append(res, each)

	}

	return
}
