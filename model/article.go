package model

import (
	"encoding/json"
	"errors"
	dm "git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
	"time"
	"os"
	. "git.vdo.space/berk/consts"
)

type ArticleStatus int32

const (
	ArticleStatusNULL ArticleStatus = 0
	ArticleStatusUnpublished ArticleStatus = 1
	ArticleStatusPublished ArticleStatus = 2
)

//GetStrByArticlePictureStatus 获取状态的string
func GetStrByArticleStatus(status ArticleStatus) string {

	switch status {

	case ArticleStatusPublished:

		return "ArticleStatusPublished"

		break

	case ArticleStatusUnpublished:

		return "ArticleStatusUnpublished"

		break
	}

	return "ArticleStatusNULL"

}

//GetArticleStatusByStr 根据string获取状态
func GetArticleStatusByStr(status string) ArticleStatus {

	switch status {

	case "ArticleStatusPublished":

		return ArticleStatusPublished

		break

	case "ArticleStatusUnpublished":

		return ArticleStatusUnpublished

		break
	}

	return ArticleStatusNULL

}

type ArticleRequestQueryParams struct {
	dm.RequestQueryParams
	Where ArticleParams `json:"where"`
}

type ArticleParams struct {
	Article
	Status string `json:"article_status"`
}

type Article struct {
	ID          *bson.ObjectId   `bson:"_id" json:"id"`
	UserID      *bson.ObjectId   `bson:"user_id" json:"user_id"`
	Title       string           `bson:"title" json:"title"`
	KeyWords    string           `bson:"key_words" json:"key_words"`
	Description string           `bson:"description" json:"description"`
	Content     string           `bson:"content" json:"content"`
	TagIDs      []*bson.ObjectId `bson:"tag_ids" json:"tag_ids"`
	Tags        []*Tag           `json:"tags"`
	Cover       string          `bson:"cover" json:"cover"`
	Status      ArticleStatus    `bson:"article_status" json:"article_status"`
	CreatedAt   *time.Time       `bson:"created_at" json:"created_at"`
	UpdatedAt   *time.Time       `bson:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time       `bson:"deleted_at" json:"deleted_at"`
}

type Avatar struct {
	AttachmentID *bson.ObjectId `bson:"attachment_id" json:"attachment_id"`
	Default      bool           `bson:"default" json:"default"`
}

type ContactInfo struct {
	Name     string `bson:"attachment_id" json:"attachment_id"`
	Img_Name string `bson:"default" json:"default"`
	Detail   string `bson:"detail" json:"detail"`
	Url      string `bson:"url" json:"url"`
}

func (v Article) GetCollectionName() string {

	return strings.ToLower(reflect.TypeOf(v).Name())
}

func ArticleToData(md *Article) (dm.Dungeons, error) {
	d := make(dm.Dungeons)
	d[`id`] = md.ID
	d[`title`] = md.Title
	d[`key_words`] = md.KeyWords
	d[`description`] = md.Description
	d[`content`] = md.Content
	if len(md.Cover) > 0 {
		d["cover"] = os.Getenv(ResourcesAddr) + "articles/" + md.UserID.Hex() + "/" + md.Cover
	}

	d[`tags`] = TagsToDatas(md.Tags)
	d[`article_status`] = GetStrByArticleStatus(md.Status)
	d[`created_at`] = md.CreatedAt
	d[`updated_at`] = md.UpdatedAt
	return d, nil
}

func ArticlesToDatas(mds []*Article) ([]dm.Dungeons, error) {
	dl := []dm.Dungeons{}
	for _, x := range mds {
		d, err := ArticleToData(x)
		if err != nil {
			return nil, err
		}
		dl = append(dl, d)
	}

	return dl, nil
}

//SetTagsForArticle 给article对象设置标签
func SetTagsForArticle(art []*Article) []*Article {

	tgs := []*bson.ObjectId{}
	tagMap := make(dm.Dungeons)

	for _, item := range art {
		tgs = append(tgs, item.TagIDs...)
	}

	tags, err := GetTagByIDs(tgs, nil)
	if err != nil {
		panic(dm.NewError(err))
	}

	for _, item := range tags {
		tagMap[item.ID.Hex()] = item
	}

	for i, item := range art {
		art[i].Tags = []*Tag{}
		for _, tagid := range item.TagIDs {
			if tagMap[tagid.Hex()] != nil {
				art[i].Tags = append(art[i].Tags, tagMap[tagid.Hex()].(*Tag))
			}
		}
	}

	return art
}

//NewArticles 批量创建Article
func NewArticles(ts []*Article) ([]*Article, error) {

	var err error
	tss := make([]interface{}, len(ts))

	for i, t := range ts {
		now := time.Now()
		bid := bson.NewObjectId()
		t.DeletedAt = &time.Time{}
		t.ID = &bid
		t.UpdatedAt = &now

		if t.CreatedAt == nil {
			t.CreatedAt = &now
		}

		if t.Status == ArticleStatusNULL {
			t.Status = ArticleStatusUnpublished
		}

		if len(strings.TrimSpace(t.Title)) == 0 {

			return nil, errors.New("The Title name does not a null string,try agin!")
		}
		if len(strings.TrimSpace(t.Description)) == 0 {

			return nil, errors.New("The Description does not a null string,try agin!")
		}
		if len(strings.TrimSpace(t.KeyWords)) == 0 {

			return nil, errors.New("The KeyWords does not a null string,try agin!")
		}
		tss[i] = t

	}

	err = CArticle.Insert(tss...)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

//DeleteArticlesByIDs 根据数据的id数组删除数据
func DeleteArticlesByIDs(ids []*bson.ObjectId, userID *bson.ObjectId) bool {

	_, err := CArticle.UpdateAll(bson.M{
		"_id": bson.M{
			"$in": ids,
		},
		"user_id": userID,
	}, bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
		},
	})

	return err == nil
}

//GetArticleByIDs 根据数据的id数组查询数据
func GetArticleByIDs(ids []*bson.ObjectId, userID *bson.ObjectId, artStatus ArticleStatus) ([]*Article, error) {
	var ts []*Article
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["_id"] = bson.M{
		"$in": ids,
	}
	if userID != nil {
		where[`user_id`] = userID
	}
	if artStatus != ArticleStatusNULL {
		where[`article_status`] = artStatus
	}
	err := CArticle.Find(where).All(&ts)
	if err != nil {

		return nil, err
	}
	return ts, nil
}

//UpdateArticlesByID 批量更新Article
func UpdateArticlesByID(ts []*Article) ([]*Article, error) {

	var err error
	tss := []*Article{}

	for _, m := range ts {

		if m.ID == nil || !m.ID.Valid() {
			return nil, errors.New("The id is invalid,try agin!")
		}

		now := time.Now()
		m.UpdatedAt = &now

		tss = append(tss, m)
		update := make(bson.M)
		col := make(bson.M)

		update["$set"] = col

		if len(strings.TrimSpace(m.Title)) > 0 {

			col["title"] = m.Title
		}

		if len(strings.TrimSpace(m.Title)) > 0 {

			col[`key_words`] = m.Title

		}
		if len(strings.TrimSpace(m.Description)) > 0 {

			col[`description`] = m.Description

		}
		if len(strings.TrimSpace(m.Content)) > 0 {

			col[`content`] = m.Content
		}
		if len(strings.TrimSpace(m.Cover)) > 0 {

			col[`cover`] = m.Cover
		}
		if len(m.TagIDs) > 0 {

			col[`tag_ids`] = m.TagIDs
		}
		if m.Status != ArticleStatusNULL {
			col[`article_status`] = m.Status
		}
		if m.CreatedAt != nil {
			col[`created_at`]=m.CreatedAt
		}

		col[`deleted_at`] = &time.Time{}
		col[`updated_at`] = m.UpdatedAt

		err = CArticle.Update(bson.D{{"_id", m.ID}, {"user_id", m.UserID}}, update)

		if err != nil {
			return nil, err
		}

	}

	return tss, nil
}

//GetArticleWhereMap 根据byte数组获取map数据
func GetArticleWhereMap(data []byte, where bson.M) bson.M {

	m := &Article{}
	json.Unmarshal(data, m)

	if len(strings.TrimSpace(m.Title)) > 0 {

		where["title"] = m.Title
	}

	if m.ID != nil && m.ID.Valid() {

		where[`_id`] = m.ID

	}
	if len(m.TagIDs) > 0 {

		where[`tag_ids`] = bson.M{
			"$all": m.TagIDs,
		}

	}
	if m.UserID != nil && m.UserID.Valid() {

		where[`user_id`] = m.UserID

	}
	if m.Status != ArticleStatusNULL {
		where[`article_status`] = m.Status
	}

	return where
}

//FindArticlesByParams 根据查询参数获取Article的分页数据
func FindArticlesByParams(req *proto.RequestParams) ([]*Article, int64, error) {

	var ts []*Article

	offsetHead := req.OffsetHead
	OffsetTail := req.OffsetTail

	if OffsetTail - offsetHead <= 0 || OffsetTail - offsetHead > 200 {

		return nil, 0, errors.New("OffsetTail must be uper than offsetHead and  OffsetTail-offsetHead must lower 200!")
	}

	where := bson.M{}
	createT := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}

	if len(strings.TrimSpace(req.UserUUID)) > 0 {

		where["user_id"] = bson.ObjectIdHex(req.UserUUID)
	}

	if len(strings.TrimSpace(req.StartTime)) > 0 {
		st, _ := time.Parse(time.RFC3339, req.StartTime)
		createT["$gte"] = st
		where["created_at"] = createT

	}

	if len(strings.TrimSpace(req.EndTime)) > 0 {
		et, _ := time.Parse(time.RFC3339, req.EndTime)
		createT["$lte"] = et
		where["created_at"] = createT

	}

	if len(req.Where) > 0 {

		where = GetArticleWhereMap(req.Where, where)
	}

	if len(strings.TrimSpace(req.Keywords)) > 0 {
		reg := bson.M{"$regex": bson.RegEx{req.Keywords, "i"}}
		or := []bson.M{bson.M{"title": reg}, bson.M{"key_words": reg}, bson.M{"description": reg}, bson.M{"content": reg}}
		where["$or"] = or

	}

	query := CArticle.Find(where)

	if len(strings.TrimSpace(req.OrderBy)) > 0 {
		if req.IsDesc {
			query = query.Sort("-" + req.OrderBy)
		} else {
			query = query.Sort(req.OrderBy)
		}
	} else {
		query = query.Sort("-created_at")
	}

	err := query.Skip(int(offsetHead)).Limit(int(OffsetTail - offsetHead)).All(&ts)

	if err != nil {

		return nil, 0, err
	}

	//totalNum mongodb中所有Article的总数，而非查询返回的条数
	totalNum, _ := CArticle.Find(where).Count()

	return ts, int64(totalNum), nil
}

//GetArticleByID 根据ArticleID查询数据
func GetArticleByID(id, userID bson.ObjectId) (u *Article, err error) {
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["_id"] = id
	where[`user_id`] = userID
	err = CArticle.Find(where).One(&u)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}
