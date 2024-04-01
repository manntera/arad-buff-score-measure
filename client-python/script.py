import pyautogui
import time
from PIL import ImageGrab
import pygetwindow as gw
import sys
from pywinauto.application import Application
from scale_factor import ICON_POTISION_PARAMS
from settings import SETTINGS
import json
import argparse

#######################################
# 関数：指定した名前のwindowを取得する
#######################################
def activate_window(window_title):
    # ウィンドウのタイトルでウィンドウを検索する
    target_window = gw.getWindowsWithTitle(window_title)
    if target_window:
        window = target_window[0]
        # 最初のウィンドウをアクティブにする
        if window.isMinimized:
            window.restore()
        if not window.isActive:
            window.activate()
        time.sleep(1)
        return window
    else:
        print(f"ウィンドウ が見つかりませんでした。")
        sys.exit()

#######################################
# 関数：指定した名前のウィンドウ以外を最小化する
#######################################
def minimize_except_specified(title):
    all_windows = gw.getAllWindows()  # 開いているすべてのウィンドウを取得
    first_window = None
    for window in all_windows:
        if title.lower() not in window.title.lower():  # 指定したタイトルを含まないウィンドウを検出
            if first_window is None:
                first_window = window
            window.minimize()  # ウィンドウを最小化
    return first_window

#######################################
# 関数：ユーザー選択をプロンプト入力させる
#######################################
def get_user_choice(prompt, choices):
    """
    ユーザーに選択肢を提示し、有効な選択を行うまで繰り返します。
    """
    choice = input(prompt).upper()
    while choice not in choices:
        print(f"Invalid choice. Please choose from : [{', '.join(choices)}]")
        choice = input(prompt).upper()
    return choice

#######################################
# 関数：設定ファイルを更新
#######################################
# def update_settings(ui_size, row_count):
#     settings = {"UI_SIZE": ui_size, "ROW_COUNT": row_count}
#     with open('settings.py', 'w') as f:
#         f.write("SETTINGS = " + repr(settings))
#     print("設定ファイルを更新しました。")

#######################################
### 設定変更処理
#######################################
# def execSetting():
#     ui_size = get_user_choice("ゲーム内設定の<UIサイズ>を選択 [1=最高, 2=高, 3=中間, 4=低, 5=最低]: ", ["1","2","3","4","5"])
#     aspect_ratio = get_user_choice("ゲーム内設定の<画面比率>を選択 [1=4:3, 2=16:9, 3=16:10]: ", ["1","2","3"])
#     row_count = get_user_choice("キャプチャを撮るバフアイコンの行数を選択 [1, 2, 3]: ", ["1","2","3"])
#     update_settings(
#         list(UI_LEVEL_SCALE.keys())[int(ui_size)-1], 
#         # list(ASPECT_RATIO_PARAMS.keys())[int(aspect_ratio)-1],
#         row_count
#     )

#######################################
### メイン処理
#######################################
def main():
    print("キャプチャを開始します。マウスを触らないでお待ちください...")
    time.sleep(1)

    #アラド戦記ウィンドウを有効化し座標取得
    aradwindow = activate_window("アラド戦記")
    bbox = aradwindow.left, aradwindow.top, aradwindow.left+aradwindow.width, aradwindow.top+aradwindow.height
                
    # アラド戦記ウィンドウ以外をすべて最小化する
    first_window = minimize_except_specified("アラド戦記")

    # バフアイコンの位置算出
    x_start_position = aradwindow.left + round(aradwindow.width/2) + round(aradwindow.height * ICON_POTISION_PARAMS["X_ORIGIN_SCALE_FACTOR"])
    y_start_position = aradwindow.top + aradwindow.height + round(aradwindow.height * ICON_POTISION_PARAMS["Y_ORIGIN_SCALE_FACTOR"])
    span = aradwindow.height * ICON_POTISION_PARAMS["ICON_WIDTH"]

    # バフアイコンの縦方向のループ
    for j in range(3):
        # バフアイコンの横方向のループ
        for i in range(9):
            x_icon = x_start_position + round(span * i)
            y_icon = y_start_position - round(span * j)

            # マウスを指定した座標に移動します
            pyautogui.moveTo(x_icon, y_icon, duration=0)

            # アラド戦記ウィンドウのスクリーンショットを撮影
            screenshot = ImageGrab.grab(bbox)

            # スクリーンショットを保存
            screenshot.save(f"tmp/screenshot_{j+1}_{i+1}.png")

    # 完了メッセージを表示
    print("スクリーンショットを撮影しました。撮影が遅い場合、画面サイズを小さくしてお試しください。")

if __name__ == "__main__":
    # parser = argparse.ArgumentParser(description="Process some integers.")
    # parser.add_argument('--settings', action='store_true', help="Update UI size and capture count settings.")
    # args = parser.parse_args()

    # if args.settings:
    #     execSetting()
    # else:
    #    main()
    main()