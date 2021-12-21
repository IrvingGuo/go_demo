import pandas as pd
import numpy as np

from db_ops import sql_within_df


class JrsProgramsMigrater:

    def __init__(self) -> None:
        self.filename = "Master RD Resource Plan.xlsx"
        pass

    # db must be empty before run this script
    def migrate(self):
        df = pd.read_excel(self.filename, sheet_name='Summary', header=None)
        df = df[[6, 7]].rename(columns={6: 'name', 7: 'type'})
        df['id'] = np.arange(1, len(df) + 1)
        print(df)
        sql_within_df(df, 'programs', 'id')


if __name__ == '__main__':
    # db must be empty before run this script
    JrsProgramsMigrater().migrate()
