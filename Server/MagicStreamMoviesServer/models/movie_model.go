package models

import "go.mongodb.org/mongo-driver/v2/bson"

// Genre 表示电影的一个分类信息。
// 它会作为 Movie 结构体中的子结构出现，既用于 MongoDB 存储，也用于接口返回。
type Genre struct {
	GenreID   int    `bson:"genre_id" json:"genre_id" validate:"required"`
	GenreName string `bson:"genre_name" json:"genre_name" validate:"required,min=2,max=100"`
}

// Ranking 表示电影的评分信息。
// 一般包括评分值，以及评分名称或评分等级的描述。
type Ranking struct {
	RankingValue int    `bson:"ranking_value" json:"ranking_value" validate:"required"`
	RankingName  string `bson:"ranking_name" json:"ranking_name" validate:"required"`
}

// Movie 是项目中的核心领域模型，表示一部电影完整的数据结构。
//
// 这个结构体有两个主要作用：
// 1. 接收和承载 MongoDB 中 movies 集合里的文档数据。
// 2. 作为接口返回值，被序列化成 JSON 返回给前端或调用方。
type Movie struct {
	// ID 是 MongoDB 自动生成的主键字段。
	ID bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	// ImdbID 是电影在 IMDb 平台上的唯一标识。
	// 当前项目通过它来查询单部电影。
	ImdbID string `bson:"imdb_id" json:"imdb_id" validate:"required"`
	// Title 是电影标题。
	Title string `bson:"title" json:"title" validate:"required,min=2,max=500"`
	// PosterPath 是电影海报图片地址。
	PosterPath string `bson:"poster_path" json:"poster_path" validate:"required,url"`
	// YoutubeUID 是 YouTube 视频的唯一标识，通常可用于预告片或相关视频。
	YoutubeUID string `bson:"youtube_uid" json:"youtube_id" validate:"required"`
	// Genre 表示这部电影所属的分类列表，例如动作、喜剧、科幻等。
	Genre []Genre `bson:"genre" json:"genre" validate:"required,dive"`
	// AdminReview 表示后台管理员或运营人员填写的影评内容。
	AdminReview string `bson:"admin_review" json:"admin_review"`
	// Ranking 表示电影的评分汇总信息。
	Ranking Ranking `bson:"ranking" json:"ranking" validate:"required"`
}
