revoke delete on table "public"."quests" from "anon";

revoke insert on table "public"."quests" from "anon";

revoke references on table "public"."quests" from "anon";

revoke select on table "public"."quests" from "anon";

revoke trigger on table "public"."quests" from "anon";

revoke truncate on table "public"."quests" from "anon";

revoke update on table "public"."quests" from "anon";

revoke delete on table "public"."quests" from "authenticated";

revoke insert on table "public"."quests" from "authenticated";

revoke references on table "public"."quests" from "authenticated";

revoke select on table "public"."quests" from "authenticated";

revoke trigger on table "public"."quests" from "authenticated";

revoke truncate on table "public"."quests" from "authenticated";

revoke update on table "public"."quests" from "authenticated";

revoke delete on table "public"."quests" from "service_role";

revoke insert on table "public"."quests" from "service_role";

revoke references on table "public"."quests" from "service_role";

revoke select on table "public"."quests" from "service_role";

revoke trigger on table "public"."quests" from "service_role";

revoke truncate on table "public"."quests" from "service_role";

revoke update on table "public"."quests" from "service_role";

alter table "public"."quests" drop constraint "quests_pkey";

drop index if exists "public"."quests_pkey";

drop table "public"."quests";

set check_function_bodies = off;

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


