import ldap
import pandas as pd

from db_ops import sql_within_df

AD_URI = 'ldap://10.70.8.3:389/'
AD_USER_BASE_DN = 'OU=Internal,OU=User,OU=NB,OU=JNN,DC=Joynext,DC=com'
AD_BIND_USER = 'svc_jnn_gerrit@joynext.com'
AD_BIND_PWD = 'hgbQQedtfF6qvrTX'

AD_USER_FILTER = '(&(objectCategory=CN=Person,CN=Schema,CN=Configuration,DC=Joynext,DC=com))'


class JrsUsersMigrater:

    def __init__(self) -> None:
        self.filename = "Master RD Resource Plan.xlsx"
        pass

    # db must be empty before run this script
    def migrate(self):
        ad_df = self.get_ad_info()
        other_info_df = self.get_location_and_res_type()
        df = pd.merge(ad_df, other_info_df, how='left', on=['cn'])
        print(df)
        sql_within_df(df, 'users', 'id')

    def get_ad_info(self):
        """
        get ad info including cn, title, sam_account
        """
        ad_conn = JrsUsersMigrater.conn_ldap()
        results = ad_conn.search_s(AD_USER_BASE_DN, ldap.SCOPE_SUBTREE, AD_USER_FILTER)
        attr_keys = ['cn', 'title', 'sAMAccountName']
        df = pd.DataFrame(columns={"id", "cn", "title", "sam_account_name"})
        for i, (_, attr) in enumerate(results):
            missing_keys = self.check_attr_keys(attr, attr_keys)
            if len(missing_keys) != 0:
                print("missing key for", attr['cn'], missing_keys)
            cn = attr['cn'][0].decode("utf-8") if 'cn' in attr else ""
            title = attr['title'][0].decode("utf-8") if 'title' in attr else ""
            sam_accountName = attr['sAMAccountName'][0].decode("utf-8") if 'sAMAccountName' in attr else ""
            df = df.append({"id": i + 1, "cn": cn, "title": title, "sam_account_name": sam_accountName}, ignore_index=True)
        return df

    @staticmethod
    def conn_ldap():
        conn = ldap.initialize(AD_URI)
        conn.protocol_version = 3
        conn.set_option(ldap.OPT_REFERRALS, 0)
        conn.simple_bind_s(AD_BIND_USER, AD_BIND_PWD)
        return conn

    @staticmethod
    def check_attr_keys(attr, keys):
        missing_keys = []
        for key in keys:
            if key not in attr:
                missing_keys.append(key)
        return missing_keys

    def get_location_and_res_type(self):
        df = pd.read_excel(self.filename, sheet_name='Total', header=1)
        return df[['Name', 'Location', 'Resource Type']].\
            rename(columns={'Name': 'cn', 'Location': 'location', 'Resource Type': 'resource_type'}).\
            drop_duplicates(subset='cn', keep="first")


if __name__ == '__main__':
    # db must be empty before run this script
    JrsUsersMigrater().migrate()
