create table `test_table` (
	`roomId` int not null,
	`winnerName` nvarchar(255) not null,
	`winnerFighter` varchar(255) not null,
	`winnerRate` int(4),
	`loserName` nvarchar(255) not null,
	`loserFighter` nvarchar(255) not null,
	`loserRate` int(4),
	primary key (`roomId`)
);

