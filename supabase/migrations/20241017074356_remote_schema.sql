CREATE TRIGGER on_auth_user AFTER INSERT ON auth.users FOR EACH ROW EXECUTE FUNCTION handle_new_user();


