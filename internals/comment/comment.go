package comment

import (
	"gorm.io/gorm"
)

// the struct for our comment service
type Service struct {
	DB *gorm.DB
}

// the db model for our Comment service
type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Auther string
}

// the interface for our comment service
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

// return a new comment service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment

	if res := s.DB.First(&comment, ID); res.Error != nil {
		return Comment{}, res.Error
	}
	return comment, nil
}

func (s *Service) GetCommentBySlug(slug string) ([]Comment, error) {
	var comments []Comment

	if res := s.DB.Find(&comments).Where("slug = ?", slug); res.Error != nil {
		return []Comment{}, res.Error
	}
	return comments, nil
}

func (s *Service) PostComment(comment Comment) (Comment, error) {

	if res := s.DB.Save(&comment); res.Error != nil {
		return Comment{}, res.Error
	}
	return comment, nil
}

func (s *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return Comment{}, err
	}

	if res := s.DB.Model(&comment).Updates(newComment); res.Error != nil {
		return Comment{}, res.Error
	}
	return comment, nil
}

func (s *Service) DeleteComment(ID uint) error {

	if res := s.DB.Delete(&Comment{}, ID); res.Error != nil {
		return res.Error
	}
	return nil
}

func (s *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment
	if res := s.DB.Find(&comments); res.Error != nil {
		return []Comment{}, res.Error
	}
	return comments, nil
}
