drop DATABASE if EXISTS hophacks;
create DATABASE IF NOT EXISTS hophacks;
use hophacks;

create table training(
  id int PRIMARY KEY AUTO_INCREMENT,
  content varchar(60000) not NULL ,
  answer int
);