create table "player" (
   "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   "firstname" varchar(255),
   "lastname" varchar(255)
);

create table "team" (
   "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   "player_1_id" UUID,
   "player_2_id" UUID
);

create table "match" (
   "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   "team_a_id" UUID,
   "team_a_score" smallint,
   "team_b_id" UUID,
   "team_b_score" smallint
);
