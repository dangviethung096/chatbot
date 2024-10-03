export PGPASSWORD='chatbot'
psql --host localhost --username chatbot --dbname chatbot -f zalo.sql
psql --host localhost --username chatbot --dbname chatbot -f facebook_database.sql