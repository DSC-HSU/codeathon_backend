create policy "Give anon users access to JPG images in folder 1enxk17_0"
on "storage"."objects"
as permissive
for select
to public
using ((bucket_id = 'output-files'::text));


create policy "Give anon users access to JPG images in folder 1enxk17_1"
on "storage"."objects"
as permissive
for insert
to public
with check ((bucket_id = 'output-files'::text));


create policy "Give anon users access to JPG images in folder 1enxk17_2"
on "storage"."objects"
as permissive
for update
to public
using ((bucket_id = 'output-files'::text));


create policy "Give anon users access to JPG images in folder 1enxk17_3"
on "storage"."objects"
as permissive
for delete
to public
using ((bucket_id = 'output-files'::text));


create policy "Give anon users access to JPG images in folder ykphhz_0"
on "storage"."objects"
as permissive
for insert
to public
with check ((bucket_id = 'eval-scripts'::text));


create policy "Give anon users access to JPG images in folder ykphhz_1"
on "storage"."objects"
as permissive
for select
to public
using ((bucket_id = 'eval-scripts'::text));


create policy "Give anon users access to JPG images in folder ykphhz_2"
on "storage"."objects"
as permissive
for update
to public
using ((bucket_id = 'eval-scripts'::text));


create policy "Give anon users access to JPG images in folder ykphhz_3"
on "storage"."objects"
as permissive
for delete
to public
using ((bucket_id = 'eval-scripts'::text));



