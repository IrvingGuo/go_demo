import pandas as pd

from db_ops import query, execute_sql


class ResPlanMigrator:
    def __init__(self) -> None:
        self.filename = "files/Master RD Resource Plan.xlsx"
        pass

    def migrate(self):
        # raw res plan data
        df = pd.read_excel(self.filename, sheet_name='Total', header=1)
        df = df[['Name', 'Program', 'Year', 'Month', 'Allocation']]

        # db data
        user_result = query("select id, cn from users")
        user_df = pd.DataFrame(user_result, columns=['user_id', 'Name'])
        prog_result = query("select id, name from programs")
        prog_df = pd.DataFrame(prog_result, columns=['program_id', 'Program'])

        # merged data
        df_with_ids = pd.merge(df, user_df, how='left', on=['Name'])
        df_with_ids = pd.merge(df_with_ids, prog_df, how='left', on=['Program'])

        # merge year and month
        df_with_ids['Time'] = df_with_ids.apply(lambda x: '%d-%02d-01' % (x['Year'], x['Month']), axis=1)
        print(df_with_ids)

        # check Nan values
        missed_name_df = df_with_ids[df_with_ids['user_id'].isna()][['Name']].drop_duplicates().reset_index(drop=True)
        missed_name_df.to_excel("files/missing_user.xlsx")
        print(missed_name_df)
        print(df_with_ids[df_with_ids['program_id'].isna()][['Program']].drop_duplicates().reset_index())

        # add
        df_with_ids = df_with_ids.dropna().reset_index(drop=True)
        latest_assignment_id = query("select id from assignments order by id desc limit 1")[0][0]
        for i, row in df_with_ids.iterrows():
            sql = "insert into assignments(`id`, `user_id`, `program_id`, `allocation`, `allocation_time`, `status`) VALUES (%d, %d, %d, %f, '%s', %d)" % \
                (latest_assignment_id + i + 1, row['user_id'], row['program_id'], row['Allocation'], row['Time'], 2)
            print(sql)
            execute_sql(sql)


if __name__ == '__main__':
    # db must be empty before run this script
    ResPlanMigrator().migrate()
