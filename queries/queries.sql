-- 1. create --
create table karyawan (
    id varchar(100),
    nama varchar(100),
    no_hp int,
    alamat text,
    departemen varchar(100)
);

alter table karyawan
add primary key (id);

alter type tingkat_kesulitan as enum ('easy', 'medium', 'hard');

create table departemen (
    id varchar(100),
    bidang varchar(100),
    category tingkat_kesulitan,
);

add primary key (id)

-- 2. insert --
insert into karyawan(id, nama, no_hp, alamat, departemen) values (
    '02',
    'alex',
    3123123,
    'sawojajar',
    'lead'
);

alter table karyawan
add column age int;

insert into departemen(id, bidang, category) values (
    '02',
    'designer',
    medium
);

alter table departemen
add column duration varchar(100);


-- 3. select --
select * from karyawan;

select * from karyawan where id='01';

select * from departemen;

select * from departemen where category=easy;

-- 4. update --
update karyawan set perusahan="telkom" where nama="agus"

update departemen set bidang="lead" where nama="jonah"

-- 5. delete --
delete from karyawan where nama="agus"

drop table karyawan;

delete from departemen where id="02"

drop table departemen;

-- 6. where --
select * from karyawan
where nama = 'budi';

select bidang from departemen
where id = '01';

-- 7. join --
select karyawan.nama, departemen.bidang
from karyawan
join departemen
on karyawan.id = departemen.id;