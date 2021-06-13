package model

import "strconv"

type Comments struct {
	Comment Tweet
}
type Author struct {
	Name      string
	TwitterID string
}
type Tweet struct {
	ID       int
	Author   Author
	DateTime string
	Content  string
	Likes    int
	Retweets int
	Comments []Comments
}

func (t Tweet) Populate() []Tweet {
	comments := []Comments{}
	tweets := []Tweet{}
	author := Author{Name: "Christopher Nolan", TwitterID: "@chrisnolan"}

	for i := 2; i < 4; i++ {

		commentAuthor := Author{Name: "jon don" + strconv.Itoa(i), TwitterID: "@jonny" + strconv.Itoa(i+3)}
		comment := Tweet{ID: i, Author: commentAuthor, DateTime: "08/06/2021 06:23", Content: "big fan of you nolan :)", Likes: i, Retweets: 0}
		comments = append(comments, Comments{Comment: comment})

	}

	t.ID = 1
	t.Author = author
	t.DateTime = "08/06/2021 06:22"
	t.Content = "Hello folks"
	t.Likes = 10
	t.Retweets = 2
	t.Comments = comments
	tweets = append(tweets, t)
	return tweets
}
