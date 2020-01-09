package repository

import (
	"l-semi-chat/pkg/domain"
)

type tagRepository struct {
	SQLHandler SQLHandler
}

type TagRepository interface {
	StoreTag(tagID string, tag string, categoryID string) error
	FindTagByTagID(tagID string) (domain.Tag, error)
	FindTags() (domain.Tags, error)
}

func NewTagRepository(sh SQLHandler) TagRepository {
	return &tagRepository{
		SQLHandler: sh,
	}
}

func (tr *tagRepository) StoreTag(tagID, tag, categoryID string) error {
	_, err := tr.SQLHandler.Execute(
		"INSERT INTO tags(id, tag, category_id) VALUES (?,?,?)",
		tagID,
		tag,
		categoryID,
	)

	if err != nil {
		return domain.InternalServerError(err)
	}
	return nil
}

func (tr *tagRepository) FindTagByTagID(tagID string) (tag domain.Tag, err error) {
	row := tr.SQLHandler.QueryRow("SELECT id, tag, category_id FROM tags WHERE id=?", tagID)
	// TODO: カテゴリも一緒にselectで取得したい
	if err = row.Scan(&tag.ID, &tag.Tag, &tag.Category.ID); err != nil {
		return tag, domain.InternalServerError(err)
	}
	return
}

func (tr *tagRepository) FindTags() (tags domain.Tags, err error) {
	// TODO: categoryへの対処
	rows, err := tr.SQLHandler.Query("SELECT id, tag, category_id FROM tags")
	if err != nil {
		return tags, domain.InternalServerError(err)
	}
	for rows.Next() {
		var tag domain.Tag
		if err = rows.Scan(&tag.ID, &tag.Tag, &tag.Category.ID); err != nil {
			continue
		}
		tags = append(tags, tag)
	}
	return
}
