from google.oauth2 import service_account
from google.auth.transport.requests import Request, AuthorizedSession
from requests_toolbelt.multipart.encoder import MultipartEncoder
import requests
import sys
import json


def getIDTokenCredentials():
    # サービスアカウントの鍵ファイルパス
    key_path = "calculate-score-api-credential.json"

    # サービスアカウント認証情報のロード
    credentials = service_account.IDTokenCredentials.from_service_account_file(
        key_path,
        target_audience='https://arad-buff-score-measure.manntera.com/calculate-score'
    )

    # 認証情報を使用してIDトークンを取得するためのリクエストを行う
    request = Request()
    credentials.refresh(request)

    return credentials

def callapi(files):
    
    # APIエンドポイント
    CLOUD_RUN_URL = 'https://arad-buff-score-measure.manntera.com/calculate-score'

    # サービスアカウントの鍵ファイルのパス
    SERVICE_ACCOUNT_FILE = 'calculate-score-api-credential.json'

    # サービスアカウントキーを使用して認証情報を取得
    credentials = service_account.Credentials.from_service_account_file(
        SERVICE_ACCOUNT_FILE,
        # Cloud Run APIを呼び出すために必要なスコープ
        scopes=['https://www.googleapis.com/auth/cloud-platform'],
    )

    # 認証済みセッションを作成
    authed_session = AuthorizedSession(credentials)

    # 認証済みセッションを使用してPOSTリクエストを送信
    response = authed_session.post(CLOUD_RUN_URL)

    return response

    
# 複数の画像ファイルをフォームデータとして準備
files = [
   ('images', open('test_image/test1.png', 'rb')),
   ('images', open('test_image/test2.png', 'rb')),
#    ('images', open('test_image/test3.png', 'rb')),
#    ('images', open('test_image/test4.png', 'rb')),
#    ('images', open('test_image/test5.png', 'rb')),
#    ('images', open('test_image/test6.png', 'rb'))
]

# マルチパートエンコーダーの作成
# files = MultipartEncoder(
#     fields={
#         # フォームデータとして送信するファイル
#         'images': ('test1.png', open('test_image/test1.png', 'rb'), 'image/png'),
#         'images': ('test2.png', open('test_image/test2.png', 'rb'), 'image/png'),
#         'images': ('test3.png', open('test_image/test3.png', 'rb'), 'image/png'),
#         'images': ('test4.png', open('test_image/test4.png', 'rb'), 'image/png'),
#         'images': ('test5.png', open('test_image/test5.png', 'rb'), 'image/png'),
#         'images': ('test6.png', open('test_image/test6.png', 'rb'), 'image/png')
#     }
# )

# POSTリクエストを送信
response = callapi(files)

# レスポンスを表示
print(response.text)

# 使用後にファイルを閉じる
for _, file in files:
    file.close()