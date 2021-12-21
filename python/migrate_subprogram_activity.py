import pandas as pd
import numpy as np

from db_ops import sql_within_df


class JrsSubprogramMigrater:

    def __init__(self) -> None:
        self.subprogram_filename = ".xlsx"
        self.activity_filename = "files/Activity.xlsx"
        pass

    # db must be empty before run this script
    def migrate_subprogram(self):
        df = pd.read_excel(self.subprogram_filename, sheet_name='Sheet', header=0)
        df = df.rename(columns={'activity_name': 'name'})
        df['status'] = df.apply(lambda row: 0 if row.status == 'Disabled' else 1, axis=1)
        df['id'] = np.arange(1, len(df) + 1)
        print(df)
        sql_within_df(df, 'programs', 'id')
    
    def migrate_activity(self):
        df = pd.read_excel(self.activity_filename, sheet_name='Sheet', header=0)
        df = df.rename(columns={'activity_name': 'name'})
        df['status'] = df.apply(lambda row: 0 if row.status == 'Disabled' else 1, axis=1)
        print(df)
        sql_within_df(df, 'activities', 'id')


if __name__ == '__main__':
    # db must be empty before run this script
    JrsSubprogramMigrater().migrate_activity()
