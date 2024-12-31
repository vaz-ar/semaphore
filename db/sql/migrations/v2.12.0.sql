create table user__totp(
  `id` integer primary key autoincrement,
  `user_id` int NOT NULL,
  `secret` varchar(200) NOT NULL,
  `created` datetime NOT NULL,
  unique (`user_id`),
  foreign key (`user_id`) references task(`id`) on delete cascade
);
