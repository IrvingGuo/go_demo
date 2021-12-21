import pandas as pd
import numpy as np

from db_ops import query, execute_sql


class DiffGroupUserTotal:

    def __init__(self) -> None:
        self.filename = "files/Master RD Resource Plan.xlsx"
        self.missed_user_filename = "files/diff_group_user_list.xlsx"
        pass

    def migrate(self):
        df = pd.read_excel(self.filename, sheet_name='Total', header=1)
        missed_user_df = pd.read_excel(self.missed_user_filename, header=0)['Name'].tolist()
        print(missed_user_df)
        df = df[df['Name'].isin(missed_user_df)]
        print(df)
        df.to_excel("files/diff_group_user_total.xlsx")


if __name__ == '__main__':
    # db must be empty before run this script
    DiffGroupUserTotal().migrate()
