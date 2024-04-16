from google.oauth2 import service_account
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request, AuthorizedSession
from requests_toolbelt.multipart.encoder import MultipartEncoder
from google.auth.exceptions import RefreshError
import requests
import sys
import json


# def callapi(files):
    
#     # APIエンドポイント
#     CLOUD_RUN_URL = 'https://arad-buff-score-measure.manntera.com/calculate-score'

#     # サービスアカウントの鍵ファイルのパス
#     SERVICE_ACCOUNT_FILE = 'calculate-score-api-credential.json'

#     # サービスアカウントキーを使用して認証情報を取得
#     credentials = service_account.Credentials.from_service_account_file(
#         SERVICE_ACCOUNT_FILE,
#         # Cloud Run APIを呼び出すために必要なスコープ
#         scopes=['https://www.googleapis.com/auth/cloud-platform'],
#     )

#     # 認証情報をリフレッシュしてアクセストークンを取得
#     request = Request()
#     credentials.refresh(request)
#     print("token: ", credentials.token)

#     # アクセストークンを返す
#     return credentials.token

#     # # 認証済みセッションを作成
#     # authed_session = requests.Session()
#     # auth_request = Request(session=authed_session)

#     # # 認証情報をセッションに適用
#     # credentials.refresh(auth_request)

#     # authed_session.headers.update({
#     #     'Authorization': f'Bearer {credentials.token}'
#     # })

#     # print("token: ", credentials.token)

#     # # Create an authorized session that automatically handles token refresh
#     # authed_session = AuthorizedSession(credentials)

#     # # 認証済みセッションを使用してPOSTリクエストを送信
#     # response = authed_session.post(CLOUD_RUN_URL, files=files)

#     response = None
#     return response

def callapi2(files):
    # サービスアカウントキーファイルへのパス
    KEY_PATH = "calculate-score-api-credential.json"

    # Cloud Run サービスの URL
    API_URL = "https://arad-buff-score-measure.manntera.com/calculate-score"

    # サービスアカウントの認証情報をロード
    credentials = service_account.IDTokenCredentials.from_service_account_file(
        KEY_PATH,
        target_audience='https://arad-buff-score-measure.manntera.com/calculate-score'
    )

    # 認証情報をリフレッシュしてIdentity Tokenを取得
    request = Request()
    credentials.refresh(request)

    # Identity Tokenの表示
    print("Identity Token:", credentials.token)

    # アクセストークンを使用して API リクエスト
    headers = {
        'Authorization': f'Bearer {credentials.token}',
        'Content-Type': 'application/json'
    }

    # POST リクエスト
    response = requests.post(API_URL, headers=headers, files=files)

    # レスポンスの表示
    print(response.status_code)
    print(response.text)
    
# def print_access_token():
#     # 認証情報ファイルのパスを指定
#     credentials_path = 'calculate-score-api-credential.json'

#     # 認証情報をロード
#     creds = Credentials.from_authorized_user_file(credentials_path)

#     # トークンが有効か、またはリフレッシュ可能か確認
#     if not creds.valid:
#         if creds.expired and creds.refresh_token:
#             try:
#                 creds.refresh(Request())
#             except RefreshError as e:
#                 print("Failed to refresh access token:", str(e))
#                 return
#         else:
#             print("No valid refresh token is available.")
#             return

#     # アクセストークンを出力
#     print("Access Token:", creds.token)

# 複数の画像ファイルをフォームデータとして準備
# files = [
#    ('images', open('test_image/test1.png', 'rb')),
#    ('images', open('test_image/test2.png', 'rb')),
#    ('images', open('test_image/test3.png', 'rb')),
#    ('images', open('test_image/test4.png', 'rb')),
#    ('images', open('test_image/test5.png', 'rb'))
#  ]

# file_paths = [
#     'test_image/test1.png', 
#     'test_image/test2.png', 
#     'test_image/test3.png'
# ]
# files = {'images': [(open(file_path, 'rb')) for file_path in file_paths]}

# 画像ファイルのリスト
# files = {
#     'images': [('test1.png', open('test_image/test1.png', 'rb')),
#                 ('test2.png', open('test_image/test2.png', 'rb')),
#                 ('test3.png', open('test_image/test3.png', 'rb')),
#                 ('test4.png', open('test_image/test4.png', 'rb')),
#                 ('test5.png', open('test_image/test5.png', 'rb'))]
# }

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

files = {
    'file': ('filename', open('test_image/test1.png', 'rb'))
}

try:
    # POSTリクエストを送信
    response = callapi2(files)

finally:
    # 使用後にファイルを閉じる
    # for file in files['images']:
    #     file[1].close()
    print("end")


# with open(file_paths[0], 'rb') as file1, open(file_paths[1], 'rb') as file2, open(file_paths[2], 'rb') as file3:
#     files = {
#         'image1': ('image1.png', file1),
#         'image2': ('image2.png', file2),
#         'image3': ('image3.png', file3)
#     }

#     # Send the POST request using the authenticated session
#     response = callapi(files)
#     print(response.status_code)
#     print(response.text)

