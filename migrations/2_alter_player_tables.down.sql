alter table player add column firstname varchar(255);
alter table player add column lastname varchar(255);

update player set firstname = name;

alter table player drop column name;