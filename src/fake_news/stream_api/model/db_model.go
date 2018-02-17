package model

//create table training(
//id int PRIMARY KEY AUTO_INCREMENT,
//content varchar(65535) not NULL ,
//answer int
//);
type TrainingModel struct{
	Id 			int 	`json:"id"`
	Content 	string 	`json:"content"`
	Answer	 	int 	`json:"answer"`
}

type TrainingDTO struct{
	Content 	string	`json:"content"`
	Answer 		int 	`json:"answer"`
}