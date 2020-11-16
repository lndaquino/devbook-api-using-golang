package repositories

import (
	"api/src/models"
	"database/sql"
)

// Posts represents a posts repository
type Posts struct {
	db *sql.DB
}

// NewPostsRepository creates a posts repository
func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

// Create saves a new post
func (repository Posts) Create(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare("insert into posts (title, content, userID) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.UserID)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

// SearchByID returns a post from database
func (repository Posts) SearchByID(postID uint64) (models.Post, error) {
	line, err := repository.db.Query(`
		select p.*, u.nick from posts p inner join users u on u.id = p.userID where p.id = ?
	`, postID)
	if err != nil {
		return models.Post{}, err
	}
	defer line.Close()

	var post models.Post
	if line.Next() {
		if err = line.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.UserID,
			&post.Likes,
			&post.CreatedAt,
			&post.UserNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

// SearchPosts returns user´s posts and also posts from all users he´s following
func (repository Posts) SearchPosts(userID uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(`
		select distinct p.*, u.nick from posts p
		inner join users u on u.id = p.userID
		inner join followers f on p.userID = f.userID
		where u.id = ? or f.followerID = ?
		order by 1 desc
	`, userID, userID)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post
	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.UserID,
			&post.Likes,
			&post.CreatedAt,
			&post.UserNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// Update updates a post´s data
func (repository Posts) Update(postID uint64, updatedPost models.Post) error {
	statement, err := repository.db.Prepare("update posts set title = ?, content = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(updatedPost.Title, updatedPost.Content, postID); err != nil {
		return err
	}

	return nil
}

// Delete deletes a post
func (repository Posts) Delete(postID uint64) error {
	statement, err := repository.db.Prepare("delete from posts where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

// SearchUserPosts return all user´s posts
func (repository Posts) SearchUserPosts(userID uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(`
		select p.*, u.nick from posts p
		join users u on u.id = p.userID
		where p.userID = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.UserID,
			&post.Likes,
			&post.CreatedAt,
			&post.UserNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil

}

// Like adds one like to a post
func (repository Posts) Like(postID uint64) error {
	statement, err := repository.db.Prepare("update posts set likes = likes + 1 where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

// Dislike subtracts one like to a post
func (repository Posts) Dislike(postID uint64) error {
	statement, err := repository.db.Prepare(`
		update posts set likes = 
		CASE
			WHEN likes > 0 THEN likes - 1
			ELSE likes
		END
		where id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}
