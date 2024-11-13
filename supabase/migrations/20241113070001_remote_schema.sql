alter table "public"."submissions" drop constraint "submissions_challenge_id_fkey";

alter table "public"."challenges" add column "input_file_urls" text[];

alter table "public"."challenges" add column "start_date_time" timestamp with time zone;

alter table "public"."profiles" add column "access_level" smallint not null default '0'::smallint;

alter table "public"."submissions" drop column "output_file_urls";

alter table "public"."submissions" drop column "submitted_at";

alter table "public"."submissions" add column "created_at" timestamp with time zone not null default now();

alter table "public"."submissions" add column "input_file_id" smallint;

alter table "public"."submissions" add column "output_file_url" text;

alter table "public"."submissions" add column "source_code_url" text;

alter table "public"."submissions" alter column "id" set default gen_random_uuid();

set check_function_bodies = off;

CREATE OR REPLACE FUNCTION public.get_global_leaderboard(limit_value integer, offset_value integer)
 RETURNS TABLE(data jsonb, end_page integer)
 LANGUAGE plpgsql
AS $function$DECLARE
    total_records INT;
    total_pages INT;
BEGIN
    -- Calculate the total number of unique users in the submissions table
    SELECT COUNT(DISTINCT user_id) INTO total_records
    FROM submissions;

    -- Calculate the total number of pages based on limit_value
    total_pages := CEIL(total_records::FLOAT / limit_value);

    -- Return the aggregated data and the total number of pages
    RETURN QUERY
    SELECT 
        JSONB_AGG(
            JSONB_BUILD_OBJECT(
                'user_id', user_id,
                'global_rank_score', global_rank_score
            )
        ) AS data,
        total_pages AS end_page
    FROM (
        SELECT 
            user_id,
            SUM(rank_score) AS global_rank_score
        FROM (
            SELECT user_id, rank_score
            FROM submissions s1
            WHERE created_at = (
                SELECT MAX(created_at)
                FROM submissions s2
                WHERE s2.user_id = s1.user_id
                  AND s2.input_file_id = s1.input_file_id
            )
        ) AS latest_submissions
        GROUP BY user_id
        ORDER BY global_rank_score DESC
        LIMIT limit_value OFFSET offset_value
    ) AS ranked_global_scores;
END;$function$
;

CREATE OR REPLACE FUNCTION public.get_leaderboard(cid uuid, limit_value integer, offset_value integer)
 RETURNS TABLE(data jsonb, end_page integer)
 LANGUAGE plpgsql
AS $function$DECLARE
    total_records INT;
    total_pages INT;
BEGIN
    -- Calculate the total number of unique users for the specified challenge
    SELECT COUNT(DISTINCT user_id) INTO total_records
    FROM (
        SELECT DISTINCT user_id
        FROM submissions
        WHERE challenge_id = cid
    ) AS user_submissions;

    -- Calculate the total number of pages based on limit_value
    total_pages := CEIL(total_records::FLOAT / limit_value);

    -- Return the aggregated data and the total number of pages
    RETURN QUERY
    SELECT 
        JSONB_AGG(
            JSONB_BUILD_OBJECT(
                'user_id', user_id,
                'score', total_score,
                'rank_score', rank_score
            )
        ) AS data,
        total_pages AS end_page
    FROM (
        SELECT 
            user_id,
            SUM(score) AS total_score,
            RANK() OVER (ORDER BY SUM(score) DESC) AS rank_score
        FROM (
            SELECT user_id, score
            FROM submissions s1
            WHERE challenge_id = cid
              AND created_at = (
                  SELECT MAX(created_at)
                  FROM submissions s2
                  WHERE s2.user_id = s1.user_id
                    AND s2.input_file_id = s1.input_file_id
                    AND s2.challenge_id = cid
              )
        ) AS latest_submissions
        GROUP BY user_id
        ORDER BY total_score DESC
        LIMIT limit_value OFFSET offset_value
    ) AS ranked_scores;
END;$function$
;


