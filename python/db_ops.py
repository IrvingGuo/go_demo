import mysql.connector
from mysql.connector import errorcode
from sqlalchemy import create_engine, NVARCHAR, inspect
import pandas as pd


config_without_db = {
    'host': '10.70.9.111',  # defined as container name in docker compose.yml
    'user': 'root',
    'password': 'root',
    'raise_on_warnings': True
}

database = 'jrs'

config_with_db = config_without_db.copy()
config_with_db['database'] = database


def conn(config=config_with_db):
    try:
        cnx = mysql.connector.connect(**config)
        cnx.autocommit = True
        return cnx
    except mysql.connector.Error as err:
        if err.errno == errorcode.ER_ACCESS_DENIED_ERROR:
            raise Exception("Wrong username or password", err)
        elif err.errno == errorcode.ER_BAD_DB_ERROR:
            print("Database does not exist, trying to create...")
            create_db()
            return conn()
        else:
            raise Exception(err)


# noinspection SqlNoDataSourceInspection
def create_db():
    cnx = conn(config_without_db)
    try:
        with cnx.cursor() as cursor:
            sql = "CREATE DATABASE IF NOT EXISTS {0};".format(database)
            cursor.execute(sql)
    except mysql.connector.Error as err:
        if err.errno != errorcode.ER_DB_CREATE_EXISTS:
            raise Exception('Fail to create database', err)
    finally:
        cnx.close()


def execute_sql(sql, params=()):
    cnx = conn()
    try:
        with cnx.cursor() as cursor:
            cursor.execute(sql, params)
    except Exception as e:
        print(e)
    finally:
        cnx.close()


def query(sql, params=()):
    cnx = conn()
    try:
        with cnx.cursor() as cursor:
            cursor.execute(sql, params)
            return cursor.fetchall()
    except Exception as e:
        print(e)
    finally:
        cnx.close()


def engine():
    return create_engine(
        'mysql+mysqlconnector://{0}:{1}@{2}/{3}'.format(
            config_without_db['user'],
            config_without_db['password'],
            config_without_db['host'],
            database
        )
    )


# db must be empty before run this script
def sql_within_df(data, table_name, primary_keys="", if_exists="append"):
    data.to_sql(
        name=table_name,
        con=engine(),
        if_exists=if_exists,
        index=False,
        dtype={col_name: NVARCHAR(length=255) for col_name in data.columns.tolist()}
    )


def read_within_df(sql, params=()):
    return pd.read_sql(sql=sql, con=engine(), params=params)
