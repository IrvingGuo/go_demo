import pandas as pd
import numpy as np

from db_ops import query, execute_sql


class MapUserDept:

    def __init__(self) -> None:
        self.filename = "files/R&D Staff List_20211207.xlsx"
        pass

    def migrate(self):
        # Division, Group, Name
        username_group_df = pd.read_excel(self.filename, sheet_name='Sheet1')
        print(username_group_df)

        # userId, Name
        user_result = query("select id, cn from users")
        user_df = pd.DataFrame(user_result, columns=['user_id', 'Name'])
        print(user_df)

        # deptId, Group
        dept_result = query("select id, name from departments where level like '%.%.%'")
        dept_df = pd.DataFrame(dept_result, columns=['dept_id', 'Group'])
        print(dept_df)

        # userId, Name, Division, Group
        user_group_df = pd.merge(user_df, username_group_df, how='left', on=['Name']).dropna()
        print(user_group_df)

        # userId, Name, Division, Group, deptId
        user_group_id_df = pd.merge(user_group_df, dept_df, how='left', on=['Group'])
        print(user_group_id_df[user_group_id_df['dept_id'].isna()][['Division', 'Group']].drop_duplicates().sort_values(['Division']))
        user_group_id_df = user_group_id_df.dropna()
        user_group_id_df['dept_id'] = user_group_id_df['dept_id'].astype(int)
        print(user_group_id_df)

        for _, row in user_group_id_df.iterrows():
            execute_sql("update users set dept_id = %d where id = %d" % (row['dept_id'], row['user_id']))


if __name__ == '__main__':
    # db must be empty before run this script
    MapUserDept().migrate()
