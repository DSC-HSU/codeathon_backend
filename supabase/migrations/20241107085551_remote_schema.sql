drop function if exists "public"."get_leaderboard"(cid text, limit_value integer, offset_value integer);

drop function if exists "public"."recalculate_rank_score"(cid text);

create table "public"."challenges" (
    "id" uuid not null,
    "created_at" timestamp with time zone not null default now(),
    "title" text,
    "description" text,
    "statement" text,
    "eval_script" text
);


alter table "public"."submissions" drop column "created_at";

alter table "public"."submissions" alter column "challenge_id" set data type uuid using "challenge_id"::uuid;

alter table "public"."submissions" alter column "id" set not null ;

alter table "public"."submissions" alter column "output_file_urls" set data type text using "output_file_urls"::text;

alter table "public"."submissions" alter column "submitted_at" set default now();

alter table "public"."submissions" alter column "submitted_at" set not null;

CREATE UNIQUE INDEX challenges_pkey ON public.challenges USING btree (id);

alter table "public"."challenges" add constraint "challenges_pkey" PRIMARY KEY using index "challenges_pkey";

alter table "public"."submissions" add constraint "submissions_challenge_id_fkey" FOREIGN KEY (challenge_id) REFERENCES challenges(id) ON UPDATE CASCADE ON DELETE CASCADE not valid;

alter table "public"."submissions" validate constraint "submissions_challenge_id_fkey";

set check_function_bodies = off;

CREATE OR REPLACE FUNCTION public.get_leaderboard(cid uuid, limit_value integer, offset_value integer)
 RETURNS TABLE(data jsonb, end_page integer)
 LANGUAGE plpgsql
AS $function$
DECLARE
    total_records INT;
    total_pages INT;
BEGIN
    -- Calculate the total number of records for the specified challenge
    SELECT COUNT(*) INTO total_records
    FROM submissions
    WHERE challenge_id = cid;

    -- Calculate the total number of pages based on limit_value
    total_pages := CEIL(total_records::FLOAT / limit_value);

    -- Return the aggregated data and the total number of pages
    RETURN QUERY
    SELECT 
        JSONB_AGG(
            JSONB_BUILD_OBJECT(
                'user_id', user_id,
                'score', score,
                'rank_score', rank_score
            )
        ) AS data,
        total_pages AS end_page
    FROM (
        SELECT user_id, score, rank_score
        FROM submissions
        WHERE challenge_id = cid
        ORDER BY rank_score DESC
        LIMIT limit_value OFFSET offset_value
    ) AS subquery;
END;
$function$
;

CREATE OR REPLACE FUNCTION public.recalculate_rank_score(cid uuid)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
DECLARE
    max_score FLOAT8;
    min_score FLOAT8;
    score_range FLOAT8;
    submission RECORD;
BEGIN
    -- Retrieve the highest and lowest scores for the specified challenge
    SELECT MAX(score), MIN(score) INTO max_score, min_score
    FROM submissions
    WHERE challenge_id = cid;

    -- If all scores are equal, set rank_score to 1000 for all submissions
    IF max_score = min_score THEN
        UPDATE submissions
        SET rank_score = 1000
        WHERE challenge_id = cid;
    ELSE
        -- Calculate the score range
        score_range := max_score - min_score;

        -- Loop through each submission to update rank_score
        FOR submission IN
            SELECT id, score FROM submissions
            WHERE challenge_id = cid
        LOOP
            UPDATE submissions
            SET rank_score = ((submission.score - min_score) / score_range) * 1000
            WHERE id = submission.id;
        END LOOP;
    END IF;
END;
$function$
;

grant delete on table "public"."challenges" to "anon";

grant insert on table "public"."challenges" to "anon";

grant references on table "public"."challenges" to "anon";

grant select on table "public"."challenges" to "anon";

grant trigger on table "public"."challenges" to "anon";

grant truncate on table "public"."challenges" to "anon";

grant update on table "public"."challenges" to "anon";

grant delete on table "public"."challenges" to "authenticated";

grant insert on table "public"."challenges" to "authenticated";

grant references on table "public"."challenges" to "authenticated";

grant select on table "public"."challenges" to "authenticated";

grant trigger on table "public"."challenges" to "authenticated";

grant truncate on table "public"."challenges" to "authenticated";

grant update on table "public"."challenges" to "authenticated";

grant delete on table "public"."challenges" to "service_role";

grant insert on table "public"."challenges" to "service_role";

grant references on table "public"."challenges" to "service_role";

grant select on table "public"."challenges" to "service_role";

grant trigger on table "public"."challenges" to "service_role";

grant truncate on table "public"."challenges" to "service_role";

grant update on table "public"."challenges" to "service_role";


