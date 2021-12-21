import pandas as pd
import numpy as np

from db_ops import query, execute_sql


class JrsResAnonymousUserMigrater:

    def __init__(self) -> None:
        self.filename = "files/Master RD Resource Plan.xlsx"
        pass

    def migrate(self):
        df = pd.read_excel(self.filename, sheet_name='Total', header=1).drop_duplicates(['Name'])
        df = df[df['Name'].str.startswith('NN') |
                df['Name'].str.startswith('NBRD') |
                df['Name'].str.startswith('SHRD') |
                df['Name'].str.startswith('DLRD')
                ]
        print(df)
        df = df[['Division', 'Group', 'Name', 'Location', 'Resource Type']]

        dept_result = query("select id, name from departments where level like '%.%.%'")
        dept_df = pd.DataFrame(dept_result, columns=['dept_id', 'Group'])

        user_group_id_df = pd.merge(df, dept_df, how='left', on=['Group'])
        print(user_group_id_df[user_group_id_df['dept_id'].isna()][['Division', 'Group']].drop_duplicates().sort_values(['Division']))
        user_group_id_df = user_group_id_df.dropna()
        user_group_id_df['dept_id'] = user_group_id_df['dept_id'].astype(int)
        print(user_group_id_df)

        latest_user_id = query("select id from users order by id desc limit 1")[0][0]
        print(latest_user_id)
        for i, row in user_group_id_df.iterrows():
            sql = "insert into users(`id`, `cn`, `dept_id`, `location`, `resource_type`, `entry_date`, `resign_date`) VALUES (%d, '%s', %d, '%s', '%s', '%s', '%s')" % \
                (latest_user_id + i + 1, row['Name'], row['dept_id'], row['Location'], row['Resource Type'], "2016-09-07", "2099-12-31")
            print(sql)
            execute_sql(sql)


if __name__ == '__main__':
    # db must be empty before run this script
    JrsResAnonymousUserMigrater().migrate()
