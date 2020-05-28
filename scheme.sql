create table `user` (
	`room_id` int not null,
	`winner_player` nvarchar(20) not null,
	`winner_fighter` varchar(20) not null,
	`winner_rate` int(4),
	`loser_player` nvarchar(20) not null,
	`loser_fighter` nvarchar(20) not null,
	`loser_rate` int(4),
	primary key (`room_id`)
)

