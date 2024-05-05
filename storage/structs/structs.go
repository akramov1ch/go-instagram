package structs

type User struct{
	User_id int
	Username string
	Firstname string
	Lastname string
	Email string
	Password string
	Phone_number string
	Followers int
	Following int
	Followers_list []int
	Following_list []int
}

type Post struct{
	Post_id int
	User_id int
	Caption string
	Likes int
	Comments int
}

type Like struct{
	Like_id int
	User_id int
	Post_id int
}

type Comment struct{
	Comment_id int
	User_id int
	Caption string
	Post_id int
	Comment string
}

