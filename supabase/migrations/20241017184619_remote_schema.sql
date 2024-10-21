drop policy "Users can update own profile." on "public"."profiles";

create table "public"."quests" (
    "id" uuid not null default gen_random_uuid(),
    "created_at" timestamp with time zone not null default now(),
    "title" text,
    "description" text,
    "photo_url" text,
    "points" bigint,
    "statement" text,
    "eval_script" text,
    "input_file_urls" text[]
);


alter table "public"."quests" enable row level security;

alter table "public"."profiles" drop column "updated_at";

CREATE UNIQUE INDEX quests_pkey ON public.quests USING btree (id);

alter table "public"."quests" add constraint "quests_pkey" PRIMARY KEY using index "quests_pkey";

grant delete on table "public"."quests" to "anon";

grant insert on table "public"."quests" to "anon";

grant references on table "public"."quests" to "anon";

grant select on table "public"."quests" to "anon";

grant trigger on table "public"."quests" to "anon";

grant truncate on table "public"."quests" to "anon";

grant update on table "public"."quests" to "anon";

grant delete on table "public"."quests" to "authenticated";

grant insert on table "public"."quests" to "authenticated";

grant references on table "public"."quests" to "authenticated";

grant select on table "public"."quests" to "authenticated";

grant trigger on table "public"."quests" to "authenticated";

grant truncate on table "public"."quests" to "authenticated";

grant update on table "public"."quests" to "authenticated";

grant delete on table "public"."quests" to "service_role";

grant insert on table "public"."quests" to "service_role";

grant references on table "public"."quests" to "service_role";

grant select on table "public"."quests" to "service_role";

grant trigger on table "public"."quests" to "service_role";

grant truncate on table "public"."quests" to "service_role";

grant update on table "public"."quests" to "service_role";

create policy "Delete profile."
on "public"."profiles"
as permissive
for delete
to public
using (true);


create policy "Users can update own profile."
on "public"."profiles"
as permissive
for update
to public
using (true);



