create table project__terraform_inventory_alias(
  `alias` varchar(100) primary key,
  `project_id` int,
  `inventory_id` int,
  `auth_key_id` int,
  foreign key (`project_id`) references project(`id`) on delete cascade,
  foreign key (`inventory_id`) references project__inventory(`id`) on delete cascade,
  foreign key (`auth_key_id`) references access_key(`id`)
);

create table project__terraform_inventory_state(
  `id` integer primary key autoincrement,
  `project_id` int,
  `inventory_id` int,
  `state` text,
  `created` datetime NOT NULL,
  foreign key (`project_id`) references project(`id`) on delete cascade,
  foreign key (`inventory_id`) references project__inventory(`id`) on delete cascade
);