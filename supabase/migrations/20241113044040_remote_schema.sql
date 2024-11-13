create policy "Give anon users access folder 1brdfhw_0"
on "storage"."objects"
as permissive
for insert
to public
with check ((bucket_id = 'input-files'::text));


create policy "Give anon users access folder 1brdfhw_1"
on "storage"."objects"
as permissive
for select
to public
using ((bucket_id = 'input-files'::text));


create policy "Give anon users access folder 1brdfhw_2"
on "storage"."objects"
as permissive
for update
to public
using ((bucket_id = 'input-files'::text));


create policy "Give anon users access folder 1brdfhw_3"
on "storage"."objects"
as permissive
for delete
to public
using ((bucket_id = 'input-files'::text));


create policy "Give anon users access to JPG images in folder 1b6a26h_0"
on "storage"."objects"
as permissive
for insert
to public
with check ((bucket_id = 'source-code-files'::text));


create policy "Give anon users access to JPG images in folder 1b6a26h_1"
on "storage"."objects"
as permissive
for select
to public
using ((bucket_id = 'source-code-files'::text));


create policy "Give anon users access to JPG images in folder 1b6a26h_2"
on "storage"."objects"
as permissive
for update
to public
using ((bucket_id = 'source-code-files'::text));


create policy "Give anon users access to JPG images in folder 1b6a26h_3"
on "storage"."objects"
as permissive
for delete
to public
using ((bucket_id = 'source-code-files'::text));



