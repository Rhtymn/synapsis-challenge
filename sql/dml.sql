\c synapsis_db

INSERT INTO shipment_methods(name)
    VALUES ('JNE'), -- 1
        ('JNT'), -- 2
        ('SICEPAT'); -- 3

INSERT INTO payment_methods(name)
    VALUES ('BNI'), -- 1
        ('BRI'), -- 2
        ('BCA'); -- 3

INSERT INTO accounts(email, email_verified, password, account_type, profile_set)
    VALUES ('admin@synapsis.id', true, '$2a$10$vvFy/rdK.GcwV5rwkRaa2e28ib9HGxyuEsigigUyuN9lME41Vse/m', 'admin', true), -- 1
        ('roihan@gmail.com', true, '$2a$10$vvFy/rdK.GcwV5rwkRaa2e28ib9HGxyuEsigigUyuN9lME41Vse/m', 'user', true), -- 2
        ('sellerone@gmail.com', true, '$2a$10$vvFy/rdK.GcwV5rwkRaa2e28ib9HGxyuEsigigUyuN9lME41Vse/m', 'seller', true), -- 3
        ('sellertwo@gmail.com', true, '$2a$10$vvFy/rdK.GcwV5rwkRaa2e28ib9HGxyuEsigigUyuN9lME41Vse/m', 'seller', true), -- 4
        ('sellerthree@gmail.com', true, '$2a$10$vvFy/rdK.GcwV5rwkRaa2e28ib9HGxyuEsigigUyuN9lME41Vse/m', 'seller', true), -- 5
        ('sellerfour@gmail.com', true, '$2a$10$vvFy/rdK.GcwV5rwkRaa2e28ib9HGxyuEsigigUyuN9lME41Vse/m', 'seller', true), -- 6
        ('sellerfive@gmail.com', true, '$2a$10$vvFy/rdK.GcwV5rwkRaa2e28ib9HGxyuEsigigUyuN9lME41Vse/m', 'seller', true); -- 7

INSERT INTO sellers(name, date_of_birth, gender, phone_number, id_account)
    VALUES ('andi herlambang', '1995-05-25', 'male', '085225121111', 3), -- 1
        ('bambang raharja', '1994-06-20', 'male', '085225122222', 4), -- 2
        ('willy', '1992-07-21', 'male', '085225123333', 5), -- 3
        ('angelina', '1993-01-18', 'female', '085225124444', 6), -- 4
        ('alina', '1994-09-02', 'female', '085225125555', 7); -- 5

INSERT INTO shops(shop_name, slug, phone_number, description, address, latitude, longitude, is_active, id_seller)
    VALUES ('andi shop', 'andi-shop', '082234567789', 
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            'Jl. Kenanga No. 15A, Jakarta 12345',
            -6.2088, 106.8456,
            true, 1),
            ('bambang shop', 'bambang-shop', '082234567779', 
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            'Perumahan Surya Indah Blok C2, Surabaya 67890',
            -7.2575, 112.7521,
            true, 2),
            ('willy shop', 'willy-shop', '082234567766', 
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            'Komplek Gading Permai 3, Medan 45678',
            3.5952, 98.6722,
            true, 3),
            ('angelina shop', 'angelina-shop', '082234567743', 
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            'Jl. Pahlawan Revolusi 28, Bandung 23456',
            -6.9175, 107.6191,
            true, 4),
            ('alina shop', 'alina-shop', '082234567711', 
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            'Graha Permata 2 Blok F5, Makassar 78901',
            -5.1477, 119.4327,
            true, 5);

INSERT INTO shop_shipment_methods(id_shop, id_shipment_method) 
    VALUES (1, 1), (1, 2), (1, 3),
        (2, 1), (2, 2), (2, 3),
        (3, 1), (3, 2), (3, 3),
        (4, 1), (4, 2), (4, 3),
        (5, 1), (5, 2), (5, 3);

INSERT INTO shop_payment_methods(id_shop, id_payment_method) 
    VALUES (1, 1), (1, 2), (1, 3),
        (2, 1), (2, 2), (2, 3),
        (3, 1), (3, 2), (3, 3),
        (4, 1), (4, 2), (4, 3),
        (5, 1), (5, 2), (5, 3);

INSERT INTO admins(name, id_account)
    VALUES('admin synapsis', 1);

INSERT INTO users(name, id_account, date_of_birth, gender, phone_number)
    VALUES('roihan', 2, '2002-04-29', 'male', '085225121403');

INSERT INTO user_addresses(name, phone_number, address, latitude, longitude, id_user)
    VALUES('rumah', '085225121403', 'kampung cemara indah, sukoharjo', 101.22, 22.39, 1),
        ('kantor', '085225121466', 'kampung melati, surabaya', 65.87, 12.49, 1);

UPDATE users SET main_address_id = 1 WHERE id = 1;

INSERT INTO categories(name, slug)
    VALUES ('buku', 'buku'), -- 1
        ('elektronik', 'elektronik'), -- 2
        ('dapur', 'dapur'), -- 3
        ('fashion', 'fashion'), -- 4
        ('handphone & tablet', 'handphone-dan-tablet'); -- 5

INSERT INTO categories(name, slug, parent_id)
    VALUES ('buku ekonomi & bisnis', 'buku-ekonomi-dan-bisnis', 1), -- 6
        ('buku hobi', 'buku-hobi', 1), -- 7
        ('buku hukum', 'buku-hukum', 1), -- 8
        ('alat pendingin ruangan', 'alat-pendingin-ruangan', 2), -- 9
        ('elektronik dapur', 'elektronik-dapur', 2), -- 10
        ('elektronik kantor', 'elektronik-kantor', 2), -- 11
        ('aksesoris dapur', 'aksesoris-dapur', 3), -- 12
        ('alat masak khusus', 'alat-masak-khusus', 3), -- 13
        ('peralatan baking', 'peralatan-baking', 3), -- 14
        ('aksesoris anak', 'aksesoris-anak', 4), -- 15
        ('fashion muslim', 'fashion-muslim', 4), -- 16
        ('sepatu', 'sepatu', 4), -- 17
        ('aksesoris handphone', 'aksesoris-handphone', 5), -- 18
        ('handphone', 'handphone', 5), -- 19
        ('tablet', 'tablet', 5); -- 20

INSERT INTO products(name, slug, price, description, stock, id_category, id_shop)
    VALUES ('Zero to One', 'zero-to-one', 57120,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            6, 1),
            ('The Psychology Money', 'the-psychology-money', 76500,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            6, 1),
            ('The Visual MBA', 'the-visual-mba', 107100,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            6, 1),
            ('Diary of CEO', 'diary-of-ceo', 394000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            6, 1),
            ('Good to Great', 'good-to-great', 120000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            6, 1), -- 1
            ('BUKU THE GEOGRAPHY OF BLISS', 'buku-the-geography-of-bliss', 57120,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            7, 1),
            ('BUKU MAN SEEKS GOD', 'buku-man-seeks-god', 76500,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            7, 1),
            ('Buku Teknik Dasar Fotografi Digital', 'buku-teknik-dasar-fotografi-digital', 107100,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            7, 1),
            ('Buku Belajar Bermain Keyboard Otodidak', 'buku-belajar-bermain-keyboard-otodidak', 394000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            7, 1),
            ('How to Win at Chess', 'how-to-win-at-chess', 120000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            7, 1), -- 2
            ('Buku Segi-Segi Hukum', 'buku-segi-segi-hukum', 57120,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            8, 1),
            ('Prinsip Prinsip Hukum Pidana', 'prinsip-prinsip-hukum-pidana', 76500,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            8, 1),
            ('KUHAP dan KUHP', 'kuhap-dan-kuhp', 107100,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            8, 1),
            ('Ilmu Hukum dan Filsafat Hukum', 'ilmu-hukum-dan-filsafat-hukum', 394000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            8, 1),
            ('Buku Ajaran Pemidanaan', 'buku-ajaran-pemidanaan', 120000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            8, 1); -- 3

INSERT INTO products(name, slug, price, description, stock, id_category, id_shop)
    VALUES ('UPHOME Kipas Angin Lipat Portable Mini', 'uphome-kipas-angin-lipat-portable-mini', 79000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            9, 2),
            ('Covenant Air Purifier AP-06 Pembersih', 'covenant-air-purifier-ap-06-pembersih', 687500,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            9, 2),
            ('MOVIO Reflektor AC - Talang AC', 'movio-reflektor-ac-talang-ac', 55000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            9, 2),
            ('Xiaomi Smart Air Purifier 4', 'xiaomi-smart-air-purifier-4', 1412000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            9, 2),
            ('Cover Pipa AC / Protective Pipe Rifeng', 'cover-pipa-ac-protective-pipe-rifeng', 98900,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            9, 2), -- 1
            ('TUVE TRC-001 Rice Cooker Low Carbo', 'tuve-trc-001-rice-cooker-low-carbo', 650000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            10, 2),
            ('NEOZEN BREAD MASTER AUTO', 'NEOZEN BREAD MASTER AUTO', 899000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            10, 2),
            ('Pensonic Air Fryer Low Watt PDFI-1305', 'pensonic-air-fryer-low-watt-pdfi-1305', 425000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            10, 2),
            ('Philips Blender', 'philips-blender', 394000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            10, 2),
            ('Oven Listrik', 'oven-listrik', 178000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            10,
            10, 2), -- 2
            ('Fingerprint Mesin Absensi Solution P208', 'fingerprint-mesin-absensi-solution-p208', 469000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            11,
            11, 2),
            ('Deli Mesin Absen Sidik Jari 1000', 'deli-mesin-absen-sidik-jari-1000', 472000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            11,
            11, 2),
            ('Deli Penghancur Kertas', 'deli-penghancur-kertas', 515000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            11,
            11, 2),
            ('SMARTCOM Cash Drawer', 'smartcom-cash-drawer', 272000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            11,
            11, 2),
            ('MOVIO Arm Rest Komputer', 'movio-arm-rest-komputer', 120000,
            'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ac risus a lacus feugiat tempor sit amet at justo. Fusce nec nibh sed nisi lacinia rutrum. Proin id augue sit amet dui fermentum blandit. Integer a felis nec augue ultrices dapibus eu id nulla.',
            11,
            11, 2); -- 3