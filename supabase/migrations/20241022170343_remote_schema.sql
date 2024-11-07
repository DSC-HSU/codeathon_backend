drop policy "Delete profile." on "public"."profiles";

drop policy "Users can insert their own profile." on "public"."profiles";

revoke delete on table "public"."challenges" from "anon";

revoke insert on table "public"."challenges" from "anon";

revoke references on table "public"."challenges" from "anon";

revoke select on table "public"."challenges" from "anon";

revoke trigger on table "public"."challenges" from "anon";

revoke truncate on table "public"."challenges" from "anon";

revoke update on table "public"."challenges" from "anon";

revoke delete on table "public"."challenges" from "authenticated";

revoke insert on table "public"."challenges" from "authenticated";

revoke references on table "public"."challenges" from "authenticated";

revoke select on table "public"."challenges" from "authenticated";

revoke trigger on table "public"."challenges" from "authenticated";

revoke truncate on table "public"."challenges" from "authenticated";

revoke update on table "public"."challenges" from "authenticated";

revoke delete on table "public"."challenges" from "service_role";

revoke insert on table "public"."challenges" from "service_role";

revoke references on table "public"."challenges" from "service_role";

revoke select on table "public"."challenges" from "service_role";

revoke trigger on table "public"."challenges" from "service_role";

revoke truncate on table "public"."challenges" from "service_role";

revoke update on table "public"."challenges" from "service_role";

alter table "public"."challenges" drop constraint "challenges_pkey";

drop index if exists "public"."challenges_pkey";

drop table "public"."challenges";

create table "public"."submissions" (
    "id" uuid not null,
    "created_at" timestamp with time zone not null default now(),
    "submitted_at" timestamp with time zone,
    "output_file_urls" text[],
    "challenge_id" text,
    "user_id" uuid,
    "score" double precision,
    "rank_score" double precision
);


alter table "public"."profiles" add column "updated_at" timestamp with time zone;

CREATE UNIQUE INDEX submissions_pkey ON public.submissions USING btree (id);

alter table "public"."submissions" add constraint "submissions_pkey" PRIMARY KEY using index "submissions_pkey";

set check_function_bodies = off;

CREATE OR REPLACE FUNCTION public.get_global_leaderboard(limit_value integer, offset_value integer)
 RETURNS TABLE(data jsonb, end_page integer)
 LANGUAGE plpgsql
AS $function$
DECLARE
    total_records INT;
    total_pages INT;
BEGIN
    -- Tính tổng số bản ghi (số lượng user_id duy nhất) trong bảng submissions
    SELECT COUNT(DISTINCT submissions.user_id) INTO total_records
    FROM submissions;

    -- Tính tổng số trang cuối cùng (end_page)
    total_pages := CEIL(total_records::FLOAT / limit_value);

    -- Trả về danh sách dữ liệu và trang cuối cùng
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
        SELECT submissions.user_id, SUM(submissions.rank_score) AS global_rank_score
        FROM submissions
        GROUP BY submissions.user_id
        ORDER BY global_rank_score DESC
        LIMIT limit_value OFFSET offset_value
    ) AS subquery;
END;
$function$
;

CREATE OR REPLACE FUNCTION public.get_leaderboard(cid text, limit_value integer, offset_value integer)
 RETURNS TABLE(data jsonb, end_page integer)
 LANGUAGE plpgsql
AS $function$
DECLARE
    total_records INT;
    total_pages INT;
BEGIN
    -- Tính tổng số bản ghi cho challenge cụ thể
    SELECT COUNT(*) INTO total_records
    FROM submissions
    WHERE challenge_id = cid;

    -- Tính tổng số trang cuối cùng (end_page)
    total_pages := CEIL(total_records::FLOAT / limit_value);

    -- Trả về danh sách dữ liệu và trang cuối cùng
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

CREATE OR REPLACE FUNCTION public.recalculate_rank_score(cid text)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
DECLARE
    max_score FLOAT8;
    min_score FLOAT8;
    score_range FLOAT8;
    submission RECORD;
BEGIN
    -- Lấy điểm cao nhất và thấp nhất trong challenge cụ thể
    SELECT MAX(score), MIN(score) INTO max_score, min_score
    FROM submissions
    WHERE challenge_id = cid;

    -- Nếu tất cả các điểm đều bằng nhau, gán rank_score là 1000 cho tất cả
    IF max_score = min_score THEN
        UPDATE submissions
        SET rank_score = 1000
        WHERE challenge_id = cid;
    ELSE
        -- Tính khoảng điểm
        score_range := max_score - min_score;

        -- Duyệt qua từng submission để cập nhật rank_score
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

CREATE OR REPLACE FUNCTION public.handle_new_user()
 RETURNS trigger
 LANGUAGE plpgsql
 SECURITY DEFINER
 SET search_path TO ''
AS $function$
begin
  insert into public.profiles (id, email)
  values (new.id, new.email);
  return new;
end;
$function$
;

grant delete on table "public"."submissions" to "anon";

grant insert on table "public"."submissions" to "anon";

grant references on table "public"."submissions" to "anon";

grant select on table "public"."submissions" to "anon";

grant trigger on table "public"."submissions" to "anon";

grant truncate on table "public"."submissions" to "anon";

grant update on table "public"."submissions" to "anon";

grant delete on table "public"."submissions" to "authenticated";

grant insert on table "public"."submissions" to "authenticated";

grant references on table "public"."submissions" to "authenticated";

grant select on table "public"."submissions" to "authenticated";

grant trigger on table "public"."submissions" to "authenticated";

grant truncate on table "public"."submissions" to "authenticated";

grant update on table "public"."submissions" to "authenticated";

grant delete on table "public"."submissions" to "service_role";

grant insert on table "public"."submissions" to "service_role";

grant references on table "public"."submissions" to "service_role";

grant select on table "public"."submissions" to "service_role";

grant trigger on table "public"."submissions" to "service_role";

grant truncate on table "public"."submissions" to "service_role";

grant update on table "public"."submissions" to "service_role";

create policy "Enable delete for users based on user_id"
on "public"."profiles"
as permissive
for delete
to public
using (true);


create policy "Users can insert their own profile."
on "public"."profiles"
as permissive
for insert
to public
with check (true);



