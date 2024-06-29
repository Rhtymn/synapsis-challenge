\c synapsis_db

INSERT INTO accounts(email, email_verified, password, account_type, profile_set)
    VALUES('admin@synapsis.id', true, '$2a$10$vvFy/rdK.GcwV5rwkRaa2e28ib9HGxyuEsigigUyuN9lME41Vse/m', 'admin', true);

INSERT INTO admins(name, id_account)
    VALUES('admin synapsis', 1)