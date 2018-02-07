# -*- coding:utf-8 -*-

__author__ = 'Demoliang'

import time,uuid
from orm import Model,StringField,BooleanField,FloatField,TextField

def next_id():
    return '%015d%s000' %(int(time.time()*1000),uuid.uuid4().hex)

class User(Model):
    '''
    create table users(
     id varchar(50),
     email varchar(50),
     passwd varchar(50),
     admin boolean,
     name varchar(50),
     image varchar(500),
     created_at date
     );
     insert into users(id,email,passwd,admin,name) values('1','abc@sina.cn','123456',0,'abc')
    '''
    __table__ = 'users'

    id = StringField(primary_key=True,default=next_id,ddl='varchar(50)')
    email = StringField(ddl='varchar(50)')
    passwd =StringField(ddl='varchar(50)')
    admin = BooleanField()
    name =StringField(ddl='varchar(50)')
    image = StringField(ddl='varchar(500)')
    created_at = FloatField(default=time.time)

class Blog(Model):
    __table__ = 'blogs'

    id = StringField(primary_key=True,default=next_id,ddl='varchar(50)')
    user_id = StringField(ddl='varchar(50)')
    user_name = StringField(ddl='varchar(50)')
    user_image = StringField(ddl='varchar(500)')
    name = StringField(ddl='varchar(50)')
    summary = StringField(ddl='varchar(200)')
    content = TextField()
    created_at = FloatField(default=time.time)

class Comment(Model):
    __table__ = 'comments'

    id = StringField(primary_key=True,default=next_id,ddl='varchar(50)')
    blog_id = StringField(ddl='varchar(50)')
    user_id = StringField(ddl='varchar(50)')
    user_name = StringField(ddl='varchar(50)')
    user_image = StringField(ddl='varchar(500)')
    content = TextField()
    created_at = FloatField(default=time.time)
