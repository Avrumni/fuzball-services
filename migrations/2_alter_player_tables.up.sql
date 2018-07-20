alter table player add column name varchar(255);

update player set name = firstname || ' ' || lastname;

alter table player drop column firstname;
alter table player drop column lastname;