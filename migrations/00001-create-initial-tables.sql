CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- CREATE TYPE segment as enum (
--     'avito_voice_messages',
--     'avito_performance_vas',
--     'avito_discount_30',
--     'avito_discount_50');

CREATE TABLE IF NOT EXISTS "user"(
    id uuid DEFAULT uuid_generate_v4(),
    primary key (id)
);

CREATE TABLE IF NOT EXISTS segment(
    id int generated always as identity,
    segment_type varchar(100) unique not null,
    primary key (id)
);

CREATE TABLE IF NOT EXISTS user_segment(
    user_id uuid,
    segment_id int,
    primary key (user_id,segment_id),
    foreign key (user_id)
                    references "user" (id) on delete cascade,
    foreign key (segment_id)
                     references segment (id) on delete cascade

);

INSERT INTO "user" (id) VALUES (uuid_generate_v4());

INSERT INTO segment (segment_type) VALUES ('AVITO_VOICE_MESSAGES');
INSERT INTO segment (segment_type) VALUES ('AVITO_PERFORMANCE_VAS');
INSERT INTO segment (segment_type) VALUES ('AVITO_DISCOUNT_30');
INSERT INTO segment (segment_type) VALUES ('AVITO_DISCOUNT_50');


INSERT INTO user_segment (user_id,segment_id) VALUES ('573e2e34-9143-404a-84c1-d60ed8a06dc7',1);
INSERT INTO user_segment (user_id,segment_id) VALUES ('573e2e34-9143-404a-84c1-d60ed8a06dc7',2);
INSERT INTO user_segment (user_id,segment_id) VALUES ('573e2e34-9143-404a-84c1-d60ed8a06dc7',3);
INSERT INTO user_segment (user_id,segment_id) VALUES ('573e2e34-9143-404a-84c1-d60ed8a06dc7',4);


INSERT INTO user_segment (user_id,segment_id) VALUES ('e83942f6-fef2-43e3-93cd-f9ba652c8f9a',1);
INSERT INTO user_segment (user_id,segment_id) VALUES ('e83942f6-fef2-43e3-93cd-f9ba652c8f9a',2);
INSERT INTO user_segment (user_id,segment_id) VALUES ('e83942f6-fef2-43e3-93cd-f9ba652c8f9a',3);
INSERT INTO user_segment (user_id,segment_id) VALUES ('e83942f6-fef2-43e3-93cd-f9ba652c8f9a',4);


SELECT "user".id, ARRAY_AGG(segment.segment_type) AS segment_types
FROM "user"
         JOIN user_segment ON "user".id = user_segment.user_id
         JOIN segment ON segment.id = user_segment.segment_id
WHERE "user".id = '70c247da-377a-42ac-97f6-316abfc43722'
GROUP BY "user".id;
